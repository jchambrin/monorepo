package lookup

import (
	"os"
	"path/filepath"
	"strings"

	"com.reservit/devops/monorepo/pkg/utils"
)

type ProjectLookup interface {
	GetProjectCharts(project string) []string
}

type Results struct {
	Builds []string
	Charts []string
}

func FilesLookup(rootPath string) *Results {
	results := &Results{}
	err := filepath.Walk(rootPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == utils.BUILD {
				results.Builds = append(results.Builds, path)
			} else if info.Name() == utils.CHART {
				results.Charts = append(results.Charts, path)
			}
			return nil
		})
	utils.CheckIfError(err)
	return results
}

func (r *Results) GetProjectCharts(project string) []string {
	var charts []string
	subProjects := r.subProjects(project)
	for _, c := range r.subCharts(project) {
		if !inSubProject(c, subProjects) {
			charts = append(charts, c)
		}
	}
	return charts
}

func (r *Results) subCharts(project string) []string {
	var charts []string
	project = utils.AppendSlash(project)
	for _, c := range r.Charts {
		if strings.HasPrefix(c, project) {
			charts = append(charts, c)
		}
	}
	return charts
}

func (r *Results) subProjects(project string) []string {
	var projects []string
	project = utils.AppendSlash(project)
	for _, b := range r.Builds {
		p := utils.AppendSlash(filepath.Dir(b))
		if p != project && strings.HasPrefix(p, project) {
			projects = append(projects, filepath.Dir(p))
		}
	}
	return projects
}

func inSubProject(chart string, subProjects []string) bool {
	for _, s := range subProjects {
		if strings.HasPrefix(chart, s) {
			return true
		}
	}
	return false
}
