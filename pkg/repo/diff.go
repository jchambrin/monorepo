package repo

import (
	"com.reservit/devops/monorepo/pkg/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"path/filepath"
)

type GitCrawler struct {
	Path string
	Version string
	Hash string
	FullScan bool
}

func (c *GitCrawler) Diff() []string {
	r, err := git.PlainOpen(c.Path)
	utils.CheckIfError(err)

	ref, err := r.Head()
	utils.CheckIfError(err)

	fromCommit, err := r.CommitObject(ref.Hash())
	utils.CheckIfError(err)

	var toCommit *object.Commit
	if c.Hash != "" && c.Version == "" {
		toCommit = getHashCommit(r, c.Hash)
	}

	if toCommit == nil {
		toCommit = getVersionCommit(c.Version, r, fromCommit)
	}

	if toCommit == nil || c.FullScan {
		return allFiles(c.Path)
	}

	return GetDiffBetween(fromCommit, toCommit)
}

func GetDiffBetween(from *object.Commit, to *object.Commit) []string {
	patch, err := from.Patch(to)
	utils.CheckIfError(err)
	var files []string
	for _, filePatch := range patch.FilePatches() {
		from, to := filePatch.Files()
		if from != nil {
			files = utils.AppendFile(files, from.Path())
		}
		if to != nil {
			files = utils.AppendFile(files, to.Path())
		}
	}
	return files
}

func getHashCommit(r *git.Repository, hash string) *object.Commit {
	to, err := r.CommitObject(plumbing.NewHash(hash))
	if err != nil {
		return nil
	}
	return to
}

func getVersionCommit(targetVersion string, r *git.Repository, c *object.Commit) *object.Commit {
	versionFile := "version"
	commitIter, err := r.Log(&git.LogOptions{From: c.Hash, FileName: &versionFile})
	utils.CheckIfError(err)

	stop := false
	for !stop {
		commit, err := commitIter.Next()
		if err != nil {
			stop = true
			break
		}
		if containsVersionFile(targetVersion, commit) {
			return commit
		}
	}

	return nil
}

func allFiles(path string) []string {
	var files []string
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	return files
}
