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
	"os/user"
	"runtime"
	"log"
	"fmt"
)

type Message interface {
	Send()
}

const dirSeperator string = "/"

type Config struct {
	OS string
	Username string
	HomeDir string
	KeyFile []byte

}

func buildConfig() *Config {
	userData, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		OS: runtime.GOOS,
		Username: userData.Name,
		HomeDir: userData.HomeDir,
		KeyFile: []byte(fmt.Sprintf("%s%s%s", userData.HomeDir, dirSeperator)),
	}
}

func main() {
	fmt.Println("Sending Message")
	gm := New("Brian Downs", "brian.downs@gmail.com", "Down-Low Message")
	gm.Body = []byte("This is the message")
	//gm.SendMessage()
	m := Message(gm)
	m.Send()
}
