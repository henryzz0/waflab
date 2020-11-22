package docker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	scheduleInterval = time.Millisecond * 50
)

type Worker struct {
	mux sync.Mutex
	master *Master
	dockerClient *client.Client
	httpClient *http.Client
	context context.Context
	port string
	numJob *safeCounter // number of ongoing jobs for this worker
	jobCapacity int // max number of jobs
}


func (w *Worker) Run() {
	for {
		if w.numJob.Value() < w.jobCapacity {
			fmt.Printf("%s start find job with %v ongoing jobs\n", w.port, w.numJob.Value())
			task := w.master.getTask()
			fmt.Printf("%s get job %s\n", w.port, task.ID)
			go w.doTask(task)
		}
		time.Sleep(scheduleInterval)
	}
}

func (w *Worker) doTask(task *Task) {
	fmt.Printf("Start task %s on %s\n", task.ID, w.port)
	w.numJob.Increment()
	defer w.numJob.Decrement()

	resp, err := sendRequest(w.httpClient,
		fmt.Sprintf("http://127.0.0.1:%s", w.port),
		task.Hostname, task.YAMLPath)
	if err != nil {
		w.master.reportTask(task, false, nil, err)
		return
	}
	defer resp.Body.Close()

	var result []Response
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &result)
	fmt.Println(result)
	if err != nil {
		w.master.reportTask(task, false, nil, err)
	}

	w.master.reportTask(task, true, result, nil)
}


// https://stackoverflow.com/questions/20205796/post-data-using-the-content-type-multipart-form-data
func sendRequest(client *http.Client, url string, target string, filename string) (res *http.Response, err error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	values := map[string]io.Reader{
		"file":  file,
		"hostname": strings.NewReader(target),
	}

	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return nil, err
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return nil, err
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return nil, err
		}
	}

	w.Close()

	req, err := http.NewRequest("GET", url, &b)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	return client.Do(req)
}

func MakeWorker(master *Master, cli *client.Client, ctx context.Context, port string) (*Worker, error) {
	w := Worker{}
	w.mux.Lock()
	defer w.mux.Unlock()

	fmt.Println(port)

	containerName := fmt.Sprintf("wafbench-%s", port)
	hasContainer, err := checkContainerWithNameExist(cli, ctx, containerName)
	if err != nil {
		return nil, err
	}
	if !hasContainer {
		// create and start container
		resp, err := cli.ContainerCreate(ctx, &container.Config{
			Image: "olament/wafbench",
			WorkingDir: "/WAFBench/ftw_compatible_tool",
			ExposedPorts: nat.PortSet{
				"5000": struct{}{},
			},
			Cmd: []string{
				"gunicorn",
				"--bind", "0.0.0.0:5000",
				"web_interface:app",
			},
			Tty: true,
		},
			&container.HostConfig{
				PortBindings: nat.PortMap{
					"5000": []nat.PortBinding{
						{
							HostIP: "0.0.0.0",
							HostPort: port,
						},
					},
				},
			}, nil, containerName)
		if err != nil {
			return nil, err
		}

		err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
		if err != nil {
			return nil, err
		}

		time.Sleep(time.Second * 2) // TODO: wait gunicorn start
	}

	// http client
	httpClient := http.Client{}

	// counter
	numJob := safeCounter{
		value: 0,
	}

	w.master = master
	w.dockerClient = cli
	w.httpClient = &httpClient
	w.context = ctx
	w.port = port
	w.numJob = &numJob
	w.jobCapacity = MaxJobPerWorker

	return &w, nil
}

func checkContainerWithNameExist(client *client.Client, ctx context.Context, name string) (bool, error) {
	containerList, err := client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return false, err
	}
	for _, instance := range containerList {
		for _, partialName := range instance.Names {
			if strings.Contains(partialName, name) {
				return true, nil
			}
		}
	}
	return false, nil
}

