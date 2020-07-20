package pipeline

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

<<<<<<< HEAD:tools/monorepo-ci/pipeline/jenkins.go
=======
type jenkins struct {
	httpClient httpCall
}

>>>>>>> 9a20343... [monorepo] Refactoring:tools/monorepo/pkg/pipeline/jenkins/jenkins.go
type build struct {
	lastRef string
}

type JobResponse struct {
	LastSuccessfulBuild *LastSuccessfulBuild `json:"lastSuccessfulBuild"`
}

type LastSuccessfulBuild struct {
	Number int    `json:"number"`
	Url    string `json:"url"`
}

type BuildResponse struct {
	Actions []*Action `json:"actions"`
}

type Action struct {
	LastBuiltRevision *LastBuiltRevision `json:"lastBuiltRevision"`
}

type LastBuiltRevision struct {
	Sha1   string    `json:"SHA1"`
	Branch []*Branch `json:"branch"`
}

type Branch struct {
	Name string `json:"name"`
}

<<<<<<< HEAD:tools/monorepo-ci/pipeline/jenkins.go
func getTargetHash(branch string) string {
=======
// GetTargetHash retrieve the git sha1 of the last successful build
func GetTargetHash(branch string) string {
>>>>>>> 9a20343... [monorepo] Refactoring:tools/monorepo/pkg/pipeline/jenkins/jenkins.go
	if branch != "" {
		build, err := getLastBuild(branch)
		if err == nil {
			return build.lastRef
		}
	}
	return ""
}

func getLastBuild(branch string) (*build, error) {
	lastBuild, err := getLastSuccessfulBuild()
	if err != nil {
		return nil, err
	}
	hash, err := getGitHash(branch, lastBuild)
	if err != nil {
		return nil, err
	}

	return &build{
		lastRef: *hash,
	}, nil
}

func getLastSuccessfulBuild() (*LastSuccessfulBuild, error) {
	response, err := http.Get("http://localhost:8080/view/K8S/job/CD/job/microservices-pipeline/api/json")
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var jobResponse JobResponse
	err = json.Unmarshal(responseData, &jobResponse)
	if err != nil {
		return nil, err
	}
	return jobResponse.LastSuccessfulBuild, nil
}

func getGitHash(branch string, build *LastSuccessfulBuild) (*string, error) {

	if build != nil {
		response, err := http.Get(fmt.Sprintf("http://localhost:8080/view/K8S/job/CD/job/microservices-pipeline/%d/api/json", build.Number))
		if err != nil {
			return nil, err
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		var buildResponse BuildResponse
		err = json.Unmarshal(responseData, &buildResponse)
		if err != nil {
			return nil, err
		}
		for _, action := range buildResponse.Actions {
			if action.LastBuiltRevision != nil && sameBranch(branch, action.LastBuiltRevision.Branch) {
				return &action.LastBuiltRevision.Sha1, nil
			}
		}
	}

	return nil, errors.New("revision not found")
}

func sameBranch(branchName string, branch []*Branch) bool {
	return branchName != "" && len(branch) == 1 && strings.Contains(branch[0].Name, branchName)
}
