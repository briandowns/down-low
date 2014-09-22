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

//
// Provides the ability to send messages via the user's Gmail account.
//

package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strconv"
	"text/template"
)

const (
	emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}
`
)

type GmailSender struct {
	Username    string
	Password    string
	EmailServer string
	Port        int
}

type GmailMessageData struct {
	From    string
	To      string
	Subject string
	Body    []byte
}

func New(from string, to string, subject string) *GmailMessageData {
	return &GmailMessageData{From: from, To: to, Subject: subject}
}

func (g *GmailMessageData) Send(c *Configuration) {
	emailUser := &GmailSender{c.GmailUser, c.GmailPassword, c.GmailServer, c.GmailPort}

	auth := smtp.PlainAuth("", emailUser.Username, emailUser.Password, emailUser.EmailServer)
	var err error
	var doc bytes.Buffer

	err = smtp.SendMail(fmt.Sprintf("%s:%s", emailUser.EmailServer, strconv.Itoa(emailUser.Port)),
		auth,
		emailUser.Username,
		[]string{"brian.downs@gmail.com"},
		doc.Bytes())
	if err != nil {
		fmt.Println(err)
	}

	t := template.New("emailTemplate")

	t, err = t.Parse(emailTemplate)
	if err != nil {
		fmt.Println("error trying to parse mail template")
	}

	err = t.Execute(&doc, g)
	if err != nil {
		fmt.Println(err)
	}
}
