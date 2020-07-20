package pipeline

import "com.reservit/devops/monorepo/pkg/version"

type List struct {
	Workspace string
	Version string
}

// List the pipelines with the version given in parameter (or the last version if none provided)
func (l *List) List() []string {
	l.fillVersion()

	// TODO

	return make([]string, 1)
}

func (l *List) fillVersion() {
	if l.Version == "" {
		l.Version = version.Get(l.Workspace)
	}
}
