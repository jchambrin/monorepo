package repo

import (
	"com.reservit/devops/monorepo/pkg/version"
	"path/filepath"

	"com.reservit/devops/monorepo/pkg/utils"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func containsVersionFile(targetVersion string, c *object.Commit) bool {
	result := false
	filesIter, err := c.Files()
	utils.CheckIfError(err)

	filesIter.ForEach(func(file *object.File) error {
		if filepath.Base(file.Name) == "version" {
			content, err := file.Contents()
			utils.CheckIfError(err)
			if targetVersion == "" || version.Parse(content) == targetVersion {
				result = true
			}
		}
		return nil
	})

	return result
}
