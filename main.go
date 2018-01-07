package main

import (
	"github.com/nmaupu/kube-ingwatcher/cli"
)

const (
	AppName = "kube-ingwatcher"
	AppDesc = "Acts as ingress sender or receiver and makes actions"
)

var (
	AppVersion string
)

func main() {
	if AppVersion == "" {
		AppVersion = "master"
	}

	cli.Process(AppName, AppDesc, AppVersion)
}
