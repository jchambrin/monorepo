package project

import (
	"com.reservit/devops/monorepo/pkg/repo"
	"com.reservit/devops/monorepo/pkg/utils"
	"github.com/go-git/go-git/v5"
)

// getFiles get diff files between HEAD and HEAD^
func (d *Diff) getFiles() []string {
	r, err := git.PlainOpen(d.Workspace)
	utils.CheckIfError(err)

	ref, err := r.Head()
	utils.CheckIfError(err)

	fromCommit, err := r.CommitObject(ref.Hash())
	utils.CheckIfError(err)

	toCommit, err := fromCommit.Parent(0)
	utils.CheckIfError(err)

	return repo.GetDiffBetween(fromCommit, toCommit)
}
