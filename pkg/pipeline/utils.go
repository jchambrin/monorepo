package pipeline

import (
	"com.reservit/devops/monorepo/pkg/build"
	"path/filepath"

	"com.reservit/devops/monorepo/pkg/utils"
)

func findPipelines(pipelineType string, projects []string) []string {
	var results []string
	for _, p := range projects {
		results = utils.AppendAll(results, getPipelineNames(pipelineType, filepath.Join(p, utils.BUILD)))
	}
	return results
}

func getPipelineNames(pipelineType, path string) []string {
	switch pipelineType {
	case "CD":
		return build.PipelinesCD(path)
	case "deploy":
		return build.PipelineDeploy(path)
	default:
		return nil
	}
}
