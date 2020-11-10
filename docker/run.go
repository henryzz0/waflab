package object

import (
	"bytes"
	"context"
	"io"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

func runContainer(folder string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	imageName := "hsluoyz/wafbench"

	//out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//io.Copy(os.Stdout, out)

	err = cli.ContainerRemove(ctx, "wafbench", types.ContainerRemoveOptions{})
	if err != nil {
		if !strings.Contains(err.Error(), "No such container") {
			panic(err)
		}
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd: []string{
			"ftw_compatible_tool",
			//"-d", "/data/regression.db",
			"-x", "load /testcase | gen | start http://test.waflab.org:7080| report | exit"},
		Tty:          true,
		//AttachStderr: true,
		//AttachStdout: true,
	},
		//&container.HostConfig{
		//	Mounts: []mount.Mount{
		//		{
		//			Type:   mount.TypeBind,
		//			Source: "C:/wafbench",
		//			Target: "/data",
		//		},
		//	},
		//},
		&container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: folder,
					Target: "/testcase",
				},
			},
		},
		nil, "wafbench")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	time.Sleep(5 * time.Second)

	containerId := resp.ID
	reader, err := cli.ContainerLogs(ctx, containerId, types.ContainerLogsOptions{
		ShowStdout: true,
		//ShowStderr: true,
	})
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	stdoutput := new(bytes.Buffer)
	io.Copy(stdoutput, reader)
	s := stdoutput.String()
	println(s)

	//fmt.Println(resp.ID)
}
