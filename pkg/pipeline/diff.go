package pipeline

import (
	"com.reservit/devops/monorepo/pkg/pipeline/chart"
	"com.reservit/devops/monorepo/pkg/pipeline/jenkins"
	"com.reservit/devops/monorepo/pkg/pipeline/lookup"
	"com.reservit/devops/monorepo/pkg/repo"
	"com.reservit/devops/monorepo/pkg/utils"
	"com.reservit/devops/monorepo/pkg/version"
	"strings"
)

type Diff struct {
	Workspace string
	Branch string
	Type string
	Version string
	Push bool
	FullScan bool
}

func (d *Diff) Diff() []string {
	hash := jenkins.GetTargetHash(d.Branch)

	gitCrawler := &repo.GitCrawler{
		Path: d.Workspace,
		Version: d.Version,
		Hash: hash,
		FullScan: d.FullScan,
	}
	files := gitCrawler.Diff()

	v := version.Get(d.Workspace)
	lookup := lookup.FilesLookup(d.Workspace)
	items := dependenciesTree(d.Type, lookup)
	var projects []string
	for _, item := range items {
		_, impactedProjects := item.impactedProjects(d.Workspace, files)
		projects = utils.AppendAll(projects, impactedProjects)
	}

	pipelines := findPipelines(d.Type, projects)
	var results []string
	for _, p := range pipelines {
		name := strings.TrimSpace(p + ":" + v)
		if !utils.Contains(pipelines, name) {
			results = append(results, name)
		}
	}

	if d.Push {
		chart.UpdateCharts(d.Workspace, d.Branch, v, lookup, projects)
	}

	return results
}
