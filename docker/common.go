// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package docker

const MaxJobPerWorker = 4
const MaxRetryTime = 0

type Task struct {
	ID       string
	Hostname string
	YAMLFile string
	Res      chan []Response
	Err      chan error
}

type Response struct {
	Title   string `json:"title"`
	Status  string `json:"result"`
	HitRule string `json:"hit_rule"`
}
