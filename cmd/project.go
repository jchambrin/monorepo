package cmd

import (
	"github.com/spf13/cobra"
	"io"
)

func newProjectCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project list|diff",
		Short: "Manage projects",
		Long:  `Manage projects`,
	}

	cmd.AddCommand(newProjectDiffCmd(out))

	return cmd
}
