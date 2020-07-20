package version

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"com.reservit/devops/monorepo/pkg/utils"
)

const versionFile = "version"

func Get(workspacePath string) string {
	content, err := ioutil.ReadFile(filepath.Join(workspacePath, versionFile))
	utils.CheckIfError(err)

	return Parse(string(content))
}

func Parse(content string) string {
	return strings.TrimSpace(content)
}
