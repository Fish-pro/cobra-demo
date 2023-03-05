package main

import (
	"os"

	"github.com/Fish-pro/cobra-demo/pkg/testctl"
)

func main() {
	cmd := testctl.NewTestCtl("testctl")
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
