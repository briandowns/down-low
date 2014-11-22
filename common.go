package main

const (
	PUB_KEY_TEXT = "ssh-rsa AAAAB3NzaC1"
)

type Msg interface {
	Send(*Configuration)
}

//privkey := usr.HomeDir + "/.ssh/id_rsa"

type Configuration struct {
	OS             string
	Username       string
	HomeDir        string
	ConfigFile     string
	GmailConfig    *GmailConf
	encryptionConf *EncryptionConf
	sshConf        *SSHConf
	aesConf        *AESConf
}

func NewSSH() *SSHConf {
	return &SSHConf{
		publicKeyPath:  "",
		privateKeyPath: "",
	}
}

type SSHConf struct {
	publicKeyPath  string
	privateKeyPath string
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
