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
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"
)

type Msg interface {
	Send(*Configuration)
}

type Configuration struct {
	OS          string
	Username    string
	HomeDir     string
	ConfigFile  string
	GmailConfig *GmailConf
	KeyFile     []byte
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

// Process the CLI arguments.
func processArgs() {
	// -key="/path/to/key" -service="service" -to="user@service" -m
	var keyPath *string = flag.String("key", "", "Path to key file.")
	var service *string = flag.String("service", "", "Service to send message through.")
	var to *string = flag.String("to", "", "User to send message to.")
	var message *bool = flag.Bool("m", false, "Message to send.")
}

// Determine the type of key given by the user.
func detectKeyType() *CLIParameters {
	//
}

// Setup the application with the needed configuration from
// the environment and from the user defined configuration
// file.
func buildConfig(keyPath string) (*Configuration, error) {
	var configuration Configuration
	var gmconf GmailConf

	userData, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Redo this so it doesn't look so terrible.
	configuration = Configuration{}
	configuration.ConfigFile = fmt.Sprintf("%s/%s", userData.HomeDir, ".down-low.json")
	configuration.OS = runtime.GOOS
	configuration.Username = userData.Username
	configuration.HomeDir = userData.HomeDir
	configuration.KeyFile = []byte(fmt.Sprintf("%s%s%s", userData.HomeDir, dirSeperator))

	file, ofErr := os.Open(configuration.ConfigFile)
	if ofErr != nil {
		log.Fatal(ofErr)
	}

	decoder := json.NewDecoder(file)
	gmconf = GmailConf{}

	err = decoder.Decode(&gmconf)
	if err != nil {
		log.Fatalln(err)
	}

	configuration.GmailConfig = &gmconf
	return &configuration, nil
}

func main() {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	processArgs()
	conf, err := buildConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf)
}
