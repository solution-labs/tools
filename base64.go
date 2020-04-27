package solutionlabs_tools

import "encoding/base64"

func Base64ToString(str string) (string, error) {
	d, err := base64.StdEncoding.DecodeString(str)
	return string(d), err

}
