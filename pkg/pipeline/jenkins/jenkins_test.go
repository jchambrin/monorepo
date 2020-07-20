package jenkins

import (
	"io/ioutil"
	"testing"
)

type clientTest struct {}

func (c *clientTest) httpGET(URL string) ([]byte, error) {
	return ioutil.ReadFile("data_test/microservices-pipeline.json")
}

type clientTest203 struct {}

func (c *clientTest203) httpGET(URL string) ([]byte, error) {
	return ioutil.ReadFile("data_test/microservices-pipeline-203.json")
}

func TestLastSuccessfulBuild(T *testing.T) {
	jenkins := &jenkins{httpClient: &clientTest{}}
	results, err := jenkins.lastSuccessfulBuild()
	if err != nil || results.Number != 203 {
		T.Error(err)
	}
}

func TestGitHash(T *testing.T) {
	build := &LastSuccessfulBuild{
		Number: 203,
		Url:    "http://localhost:8080",
	}
	jenkins := &jenkins{httpClient: &clientTest203{}}
	result, err := jenkins.gitHash("hotfix", build)
	if err != nil || *result != "e5bfec821ff2b9cfe444d3849dc40095afa14b4b" {
		T.Error(err)
	}
}
