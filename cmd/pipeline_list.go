package cmd

import (
	"com.reservit/devops/monorepo/pkg/pipeline"
	"com.reservit/devops/monorepo/pkg/utils"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"path/filepath"
	"strings"
)

func newPipelineListCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		Short: "Scan modified code to discover pipelines to trigger",
		Long:  `Scan the code of the specified workspace. Return the pipeline that need to be ran based on the content of BUILD files`,
		Run: func(cmd *cobra.Command, args []string) {
			workspace, err := filepath.Abs(cmd.Flag("workspace").Value.String())
			utils.CheckIfError(err)
			list := &pipeline.List{
				Workspace: workspace,
				Version: cmd.Flag("version").Value.String(),
			}
			fmt.Fprintln(out, strings.Join(list.List(), ","))
		},
	}

	cmd.Flags().StringP("version", "v", "", "Version")

	return cmd
}