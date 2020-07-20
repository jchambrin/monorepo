package utils

func CheckIfError(err error) {
	if err == nil {
		return
	}
	panic(err)
}
