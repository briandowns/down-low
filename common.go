package main

const (
	PubKeyText = "ssh-rsa AAAAB3NzaC1"
)

type Msg interface {
	Send(*State)
}

//privkey := usr.HomeDir + "/.ssh/id_rsa"

type State struct {
	OS             string
	Username       string
	HomeDir        string
	ConfigFile     string
	GmailConfig    *GmailConf
	encryptionConf *EncryptionConf
	sshConf        *SSHConf
	aesConf        *AESConf
}

type SSHConf struct {
	publicKeyPath  string
	privateKeyPath string
}

func NewSSH() *SSHConf {
	return &SSHConf{publicKeyPath: "", privateKeyPath: ""}
}

type AESConf struct {
	aesKeyPath string
	ivPath     string
}

// Primary message data structure.
type Message struct {
	From    string
	To      string
	Secret  []byte
	Subject string
}
