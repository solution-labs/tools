package solutionlabs_tools

import "encoding/base64"

// Base64ToString converts a base64 encoded string to a standard string
func Base64ToString(str string) (string, error) {
	d, err := base64.StdEncoding.DecodeString(str)
	return string(d), err

}
