package repo

import (
	"com.reservit/devops/monorepo/pkg/utils"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
	"path/filepath"
	"time"
)

func PushVersion(workspace, branch, version string, files []string) {

	r, err := git.PlainOpen(workspace)
	utils.CheckIfError(err)

	w, err := r.Worktree()
	utils.CheckIfError(err)

	for _, f := range files {
		f, err := filepath.Rel(workspace, f)
		_, err = w.Add(f)
		utils.CheckIfError(err)
	}

	msg := fmt.Sprintf("[CI] charts version %s", version)
	commit, err := w.Commit(msg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "integration",
			Email: "it.dev.integration@reservit.com",
			When:  time.Now(),
		},
	})
	utils.CheckIfError(err)

	_, err = r.CommitObject(commit)
	utils.CheckIfError(err)

	h, err := r.Head()
	utils.CheckIfError(err)
	refSpec := fmt.Sprintf("%s:refs/heads/%s", h.Name().String(), branch)

	err = r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &gerritSSH{},
		RefSpecs: []config.RefSpec{config.RefSpec(refSpec)},

	})
	utils.CheckIfError(err)
}