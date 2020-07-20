package project

import (
	"os"
	"path/filepath"
	"strings"

	"com.reservit/devops/monorepo/pkg/utils"
)

type Diff struct {
	Workspace string
}

// Diff projects since last commit
func (d *Diff) Diff() []string {
	files := d.getFiles()
	return d.closestProjects(files)
}

// closestProjects find all projects which belong the files paths given in parameter
func (d *Diff) closestProjects(files []string) []string {
	allProjects := d.allProjects()
	var projects []string
	if len(files) > 0 {
		for _, file := range files {
			project := findProject(file, allProjects)
			if project != "" && !utils.Contains(projects, project) {
				projects = append(projects, project)
			}
		}
	}
	return projects
}

// allProjects look for BUILD file to identify all projects
func (d *Diff) allProjects() []string {
	var projects []string
	err := filepath.Walk(d.Workspace,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == utils.BUILD {
				dir, err := filepath.Rel(d.Workspace, filepath.Dir(path))
				utils.CheckIfError(err)
				projects = append(projects, dir)
			}
			return nil
		})
	utils.CheckIfError(err)
	return projects
}

// findProject find which project belongs to the file given in parameter
func findProject(file string, projects []string) string {
	for _, project := range projects {
		if strings.HasPrefix(file, project+"/") {
			return project
		}
	}
	return ""
}
