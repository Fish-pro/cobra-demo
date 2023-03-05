package testctl

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewTestCtl(use string) *cobra.Command {
	cmd := cobra.Command{
		Use:   use,
		Short: "tool test",
		Long:  "tool test long",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("tool test")
		},
	}
	return &cmd
}
