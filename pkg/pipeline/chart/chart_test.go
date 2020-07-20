package chart

import (
	"testing"
)

var mockProjects func(project string) []string

type lookupResultsMock struct{}

func (l *lookupResultsMock) GetProjectCharts(project string) []string {
	return mockProjects(project)
}

func TestUpdateCharts(T *testing.T) {
	lookupResults := &lookupResultsMock{}
	mockProjects = func(project string) []string {
		return []string{"test"}
	}
	UpdateCharts("/tmp", "hotfix", "1.0.0", lookupResults, []string{"project1", "project2"})
}
