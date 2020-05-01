package swarm

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func LoadSecret(variable string) {

	dat, err := ioutil.ReadFile("/run/secrets/" + variable)
	if err != nil {
		log.Fatal(err)
	} else {
		os.Setenv(variable, strings.Replace(string(dat), "\n", "", -1))

	}
}
