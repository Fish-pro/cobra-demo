package testctl

import "github.com/spf13/cobra"

func NewTestCtl(use string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: "tool test",
		Long:  "tool test long",
	}
}
