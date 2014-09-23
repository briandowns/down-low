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
	"path/filepath"
	"runtime"
)

type Msg interface {
	Send(*Configuration)
}

// Prepping for a later version of the configuration

type Configuration struct {
	OS         string
	Username   string
	HomeDir    string
	ConfigFile string
	GmailConf  *GmailConf
	KeyFile    []byte
}


type Message struct {
	From    string
	To      string
	Body    []byte
	Subject string
}

func New(from string, to string, subject string) *Message {
	return &Message{From: from, To: to, Subject: subject}
}

type Configuration struct {
	OS         string
	Username   string
	HomeDir    string
	ConfigFile string
	//	KeyFile       []byte
}

func parseArgs() {
	//
}

// Setup the application with the needed configuration from
// the environment and from the user defined configuration
// file.
func buildConfig() (*Configuration, error) {
	var configuration Configuration

	userData, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	confFile := fmt.Sprintf("%s/%s", userData.HomeDir, ".down-low.json")
	results, err := ioutil.ReadDir(userData.HomeDir)
	if err != nil {
		log.Fatal(err)
	}

	// Look in the user's home dir to find a down-low config file.
	for _, i := range results {
		log.Println(i.Name())
		if i.Name() == confFile {
			absPath, pathErr := filepath.Abs(i.Name())
			log.Println(absPath)
			if pathErr != nil {
				log.Fatal(pathErr)
			}
			file, err := os.Open(absPath)
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
			//configuration.KeyFile = []byte(fmt.Sprintf("%s%s%s", userData.HomeDir, dirSeperator))

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
