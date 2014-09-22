// Copyright 2014 Brian J. Downs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"runtime"
)

const dirSeperator string = "/"

type Message interface {
	Send(*Configuration)
}

type Configuration struct {
	OS            string
	Username      string
	HomeDir       string
	ConfigFile    string
	GmailAddress  string `json:"gmail_address"`
	GmailUser     string `json:"gmail_user"`
	GmailPassword string `json:"gmail_password"`
	GmailServer   string `json:"gmail_server_address"`
	GmailPort     int    `json:"gmail_server_port"`
	KeyFile       []byte
}

func parseArgs() {
	//
}

// Setup the applicaiton with the needed configuration from
// the environment and from the user defined confuration
// file.
func buildConfig() (*Configuration, error) {
	var configuration Configuration

	confFile := ".down-low.json"

	userData, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	results, err := ioutil.ReadDir(userData.HomeDir)
	if err != nil {
		log.Fatal(err)
	}

	// Look in the user's home dir to find a down-low config file.
	for _, i := range results {
		if i.Name() == confFile {
			file, err := os.Open(confFile)
			if err != nil {
				log.Fatal(err)
			}

			decoder := json.NewDecoder(file)
			configuration = Configuration{}

			err = decoder.Decode(&configuration)
			if err != nil {
				log.Fatalln(err)
			}

			configuration.ConfigFile = confFile
			configuration.OS = runtime.GOOS
			configuration.Username = userData.Username
			configuration.HomeDir = userData.HomeDir
			configuration.KeyFile = []byte(fmt.Sprintf("%s%s%s", userData.HomeDir, dirSeperator))

			break
		} else {
			return nil, errors.New("Config file not found!")
		}
	}

	return &configuration, nil
}

func main() {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	parseArgs()
	conf, err := buildConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf)
}
