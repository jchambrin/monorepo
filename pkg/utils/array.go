package utils

import "path/filepath"

func AppendAll(results, arr []string) []string {
	for _, path := range arr {
		if !Contains(results, path) {
			results = append(results, path)
		}
	}
	return results
}

func AppendFile(files []string, file string) []string {
	name := filepath.Base(file)
	if name != "version" && name != "Chart.yaml" && !Contains(files, file) {
		files = append(files, file)
	}
	return files
}

func AppendAllFiles(files, filesToAppend []string) []string {
	for _, fileToAppend := range filesToAppend {
		files = AppendFile(files, fileToAppend)
	}
	return files
}

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

