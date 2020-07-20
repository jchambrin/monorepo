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

func newPipelineDiffCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use: "diff",
		Short: "Scan modified code to discover pipelines to trigger",
		Long:  `Scan the code of the specified workspace. Return the pipeline that need to be ran based on the content of BUILD files`,
		Run: func(cmd *cobra.Command, args []string) {
			workspace, err := filepath.Abs(cmd.Flag("workspace").Value.String())
			utils.CheckIfError(err)
			diff := &pipeline.Diff{
				Workspace: workspace,
				Branch:    cmd.Flag("branch").Value.String(),
				Type:      cmd.Flag("type").Value.String(),
				Version:   cmd.Flag("from-version").Value.String(),
				Push:      utils.ParseBool(cmd.Flag("push-version").Value.String()),
				FullScan:  utils.ParseBool(cmd.Flag("full-scan").Value.String()),
			}
			fmt.Fprintln(out, strings.Join(diff.Diff(), ","))
		},
	}

	cmd.Flags().StringP("workspace", "w", "", "Workspace path")
	cmd.Flags().StringP("branch", "b", "", "Branch of current build")
	cmd.Flags().StringP("type", "t", "CD", "Pipeline type specified in BUILD.yaml files (under pipeline)")
	cmd.Flags().String("from-version", "", "From which version it will scan the changes")
	cmd.Flags().Bool("push-version", false, "Automatically create a commit for modified charts")
	cmd.Flags().Bool("full-scan", false, "Will return all the pipelines no matter if modified or not")

	err := cmd.MarkFlagRequired("branch")
	utils.CheckIfError(err)

	return cmd
}

