package main

import (
	"os"

	"github.com/yedamao/mcqbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
