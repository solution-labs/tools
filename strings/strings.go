package strings

import (
	"regexp"
	"strings"
)

func Empty(data string) bool {
	return len(data) > 0
}

func Trim(data string) string {
	return strings.TrimSpace(data)
}

func AlphaNumberOnlyWithSpace(data string) (string, error) {

	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")

	if err != nil {
		return "", err
	}

	return reg.ReplaceAllString(data, ""), err

}

func AlphaNumberOnlyWithSpaceArray(data []string) ([]string, error) {

	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")

	if err != nil {
		return data, err
	}

	for i, _ := range data {
		data[i] = reg.ReplaceAllString(data[i], "")
	}

	return data, err

}
