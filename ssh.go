package main

import (
	"io/ioutil"
	"log"
	"os/user"
	"strings"
)

// findSSHKeys will look for RSA public and private
// keys in the users home directory.
func findSSHKeys() []string {

}

func isPublicSSHKey(key string) bool {
	var result bool
	ud, uErr := user.Current()
	if uErr != nil {
		log.Fatalln(uErr)
	}

	keyData, err := ioutil.ReadFile(ud.HomeDir + "/.ssh/" + key)
	if err != nil {
		log.Println(err)
	}

	for _, line := range strings.Split(string(keyData), "\n") {
		if strings.Contains(line, PUB_KEY_TEXT) {
			result = true
			break
		}
	}
	return result
}
