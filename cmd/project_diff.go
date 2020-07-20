package cmd

import (
	"com.reservit/devops/monorepo/pkg/project"
	"com.reservit/devops/monorepo/pkg/utils"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"path/filepath"
	"strings"
)

func newProjectDiffCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "diff",
		Short: "Search modified projects since last commit",
		Long:  `Scan code to retrieve the modified projects since the last commit`,
		Run: func(cmd *cobra.Command, args []string) {
			workspace, err := filepath.Abs(cmd.Flag("workspace").Value.String())
			utils.CheckIfError(err)
			diff := &project.Diff{Workspace: workspace}
			fmt.Fprintln(out, strings.Join(diff.Diff(), ","))
		},
	}

	cmd.Flags().StringP("workspace", "w", ".", "Workspace path")

	return cmd
}
