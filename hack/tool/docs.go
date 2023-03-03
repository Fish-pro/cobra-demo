package main

import (
	"log"
	"os"

	"github.com/spf13/cobra/doc"

	"github.com/Fish-pro/cobra-demo/pkg/testctl"
)

func main() {
	// set HOME env var so that default values involve user's home directory do not depend on the running user.
	os.Setenv("HOME", "/home/user")
	os.Setenv("XDG_CONFIG_HOME", "/home/user/.config")

	err := doc.GenMarkdownTree(testctl.NewTestCtl("testctl"), "./docs/")
	if err != nil {
		log.Fatal(err)
	}
}
