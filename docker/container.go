// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package docker

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"

	"github.com/waflab/waflab/util"
)

func GetStatusFromContainer(folder string, url string) []int {
	runContainer(folder, url)
	statuses := readDbResult(folder + "/db/result.db")
	util.RemovePath(folder)
	return statuses
}

func runContainer(folder string, url string) {
	dbFolder := folder + "/db"
	util.RemovePath(dbFolder)
	util.EnsureFileFolderExists(dbFolder)

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

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd: []string{
			"ftw_compatible_tool",
			"-d", "/data/result.db",
			"-x", fmt.Sprintf("load /testcase | gen | start %s| report | exit", url)},
		Tty: true,
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
			//AutoRemove: true,
			NetworkMode: "host",
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: folder,
					Target: "/testcase",
				},
				{
					Type:   mount.TypeBind,
					Source: dbFolder,
					Target: "/data",
				},
			},
		},
		nil, "")
	if err != nil {
		panic(err)
	}

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	err = cli.ContainerResize(ctx, resp.ID, types.ResizeOptions{
		Width:  800,
		Height: 600,
	})
	if err != nil {
		panic(err)
	}

	// https://stackoverflow.com/questions/58732588/accept-user-input-os-stdin-to-container-using-golang-docker-sdk-interactive-co
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
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

		err = cli.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{
			//RemoveVolumes: true,
			//RemoveLinks:   true,
			//Force:         true,
		})
		if err != nil {
			panic(err)
		}
	}
}
