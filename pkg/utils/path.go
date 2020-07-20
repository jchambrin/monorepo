package utils

import "os"

func AppendSlash(path string) string {
	if path[len(path)-1:] != string(os.PathSeparator) {
		return path + string(os.PathSeparator)
	}
	return path
}
