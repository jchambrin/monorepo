package utils

import "strconv"

func ParseBool(str string) bool {
	bool, err := strconv.ParseBool(str)
	CheckIfError(err)
	return bool
}
