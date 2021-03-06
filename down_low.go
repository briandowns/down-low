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

const usage = ``

// New builds a new Message object.
func New(from, to, subject string) *Message {
	return &Message{From: from, To: to, Subject: subject}
}

// processArgs processes the CLI arguments.
func processArgs() {
	// -key="/path/to/key" -service="service" -to="user@service" -m
	var keyPath = flag.String("k", "", "Path to key file.")
	var service = flag.String("s", "", "Service to send message through.")
	var to = flag.String("t", "", "User to send message to.")
	var message = flag.Bool("m", false, "Message to send.")
}

// Determine the type of key given by the user.
func detectKeyType(config *CLIParameters) {
	//
}

// Setup the application with the needed state from the environment and from
// the user defined state file.
func buildState(keyPath string) (*State, error) {
	var state State
	var gmconf GmailConf

	userData, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	state = &State{
		ConfigFile: fmt.Sprintf("%s/%s", userData.HomeDir, ".down-low.json"),
		OS:         runtime.GOOS,
		Username:   userData.Username,
		HomeDir:    userData.HomeDir,
		KeyFile:    []byte(fmt.Sprintf("%s%s%s", userData.HomeDir, dirSeperator, "")),
	}

	file, err := os.Open(state.ConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	gmconf = GmailConf{}

	err = decoder.Decode(&gmconf)
	if err != nil {
		log.Fatalln(err)
	}

	state.GmailConfig = &gmconf
	return state, nil
}

func main() {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	processArgs()
	conf, err := buildState("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf)

	os.Exit(0)
}
