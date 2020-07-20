package version

import "testing"

func TestParse(T *testing.T) {
	if Parse(" 3.1.10  ") != "3.1.10" {
		T.Error()
	}
}
