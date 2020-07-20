package cmd

import (
	"github.com/spf13/cobra"
	"io"
)

func newPipelineCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use: "pipeline list|diff",
		Short: "Manage pipelines configured in source code",
		Long: `Manage pipelines configured in source code`,
	}

	cmd.AddCommand(newPipelineDiffCmd(out))
	cmd.AddCommand(newPipelineListCmd(out))

	return cmd
}
