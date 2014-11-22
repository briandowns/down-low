package main

import (
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
	"strings"
)

// findSSHKeys will look for RSA public and private
// keys in the users home directory.
func findSSHKeys() ([]string, error) {
	keys, err := filepath.Glob(fmt.Sprintf("%s/.ssh/*.pub", Configuration.HomeDir))
	if err != nil {
		return "", log.Fatalln(err)
	}
	return keys, ""
}

func isPublicRSAKey(key string) bool {
	ud, err := user.Current()
	if err != nil {
		log.Fatalln(uErr)
	}

	keyData, err := ioutil.ReadFile(ud.HomeDir + "/.ssh/" + key)
	if err != nil {
		log.Println(err)
	}

	for _, line := range strings.Split(string(keyData), "\n") {
		if strings.Contains(line, PUB_KEY_TEXT) {
			return result
		}
	}
}
