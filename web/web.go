package web

import (
	"io/ioutil"
	"net/http"
)

func ReadBody(resp *http.Request) []byte {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func BodytoString(resp *http.Request) string {
	return string(ReadBody(resp))
}
