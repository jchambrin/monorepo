package pipeline

import (
	"com.reservit/devops/monorepo/pkg/build"
	"path/filepath"
	"strings"

	"com.reservit/devops/monorepo/pkg/pipeline/lookup"
	"com.reservit/devops/monorepo/pkg/utils"
)

type item struct {
	path      string
	pipelines []string
	parents   []*item
	child     *item
}

func (i *item) allPipelines(basePath string, files []string) (bool, []string) {
	found := false
	var pipelines []string
	if i.parents != nil {
		for _, parent := range i.parents {
			foundParent, pipelinesParent := parent.allPipelines(basePath, files)
			pipelines = utils.AppendAll(pipelines, pipelinesParent)
			if foundParent {
				found = true
			}
		}
	}
	if i.pipelines != nil && (found || i.fileInPath(basePath, files)) {
		pipelines = utils.AppendAll(pipelines, i.pipelines)
		found = true
	}
	return found, pipelines
}

func (i *item) impactedProjects(basePath string, files []string) (bool, []string) {
	found := false
	var results []string
	if i.parents != nil {
		for _, parent := range i.parents {
			foundParent, resultsParent := parent.impactedProjects(basePath, files)
			results = utils.AppendAll(results, resultsParent)
			if foundParent {
				found = true
			}
		}
	}
	if found || i.fileInPath(basePath, files) {
		if !utils.Contains(results, i.path) {
			results = append(results, i.path)
		}
		found = true
	}
	return found, results
}

func (i *item) fileInPath(basePath string, files []string) bool {
	path, err := filepath.Rel(basePath, i.path)
	utils.CheckIfError(err)
	for _, file := range files {
		if strings.HasPrefix(file, path+"/") {
			return true
		}
	}
	return false
}

func (i *item) addParent(parent *item) {
	parent.child = i
	i.parents = append(i.parents, parent)
}

func dependenciesTree(pipelineType string, lookupResults *lookup.Results) []*item {
	var items []*item
	for _, path := range lookupResults.Builds {
		dir := filepath.Dir(path)
		items = append(items, analyzeSettings(pipelineType, dir))
	}
	return items
}

func analyzeSettings(pipelineType, path string) *item {
	buildFile := filepath.Join(path, utils.BUILD)
	item := &item{
		path:      path,
		pipelines: getPipelineNames(pipelineType, buildFile),
	}
	parents := build.Dependencies(buildFile)
	for _, parent := range parents {
		parentItem := analyzeSettings(pipelineType, filepath.Join(path, parent))
		item.addParent(parentItem)
	}
	return item
}

