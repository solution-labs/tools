package web

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func ReadBody(resp http.Request) []byte {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func BodytoString(resp http.Request) string {
	return string(ReadBody(resp))
}

func BodytoStruct(resp http.Request, obj *interface{}) error {

	body := ReadBody(resp)

	if len(body) > 0 {
		return json.Unmarshal(body, &obj)
	} else {
		return errors.New("empty body")
	}

}
