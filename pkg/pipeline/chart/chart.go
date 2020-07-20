package chart

import (
	"fmt"
	"io/ioutil"
	"strings"

	"com.reservit/devops/monorepo/pkg/pipeline/lookup"
	"com.reservit/devops/monorepo/pkg/repo"
	"com.reservit/devops/monorepo/pkg/utils"
)

// UpdateCharts update the version of charts
func UpdateCharts(workspace, branch, version string, lookupResults lookup.ProjectLookup, projects []string) {
	var modifiedCharts []string
	for _, p := range projects {
		charts := lookupResults.GetProjectCharts(p)
		for _, c := range charts {
			if applyVersion(version, c) {
				modifiedCharts = append(modifiedCharts, c)
			}
		}
	}

	if len(modifiedCharts) > 0 {
		repo.PushVersion(workspace, branch, version, modifiedCharts)
	}
}

// applyVersion apply the given version to the chart file
func applyVersion(version, chart string) bool {
	input, err := ioutil.ReadFile(chart)
	utils.CheckIfError(err)

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "version: ") {
			lines[i] = fmt.Sprintf("version: %s", version)
		} else if strings.Contains(line, "appVersion: ") {
			lines[i] = fmt.Sprintf("appVersion: %s", version)
		}
	}

	output := strings.Join(lines, "\n")
	if string(input) != output {
		err = ioutil.WriteFile(chart, []byte(output), 0644)
		utils.CheckIfError(err)
		return true
	}

	return false
}
