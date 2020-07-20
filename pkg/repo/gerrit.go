package repo

import (
	"errors"
	"github.com/kevinburke/ssh_config"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"com.reservit/devops/monorepo/pkg/utils"
	transportssh "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"golang.org/x/crypto/ssh"
)

var preferredKexAlgos = []string{
	"diffie-hellman-group1-sha1",
}

var preferredHostKeyAlgorithms = []string{
	"ssh-rsa",
}

type gerritSSH struct {
	transportssh.AuthMethod
}

func (a *gerritSSH) ClientConfig() (*ssh.ClientConfig, error) {
	cb, err := transportssh.NewKnownHostsCallback()
	utils.CheckIfError(err)
	username, err := username()
	signer := signer()
	utils.CheckIfError(err)

	return &ssh.ClientConfig{
		User:              username,
		Auth:              []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyAlgorithms: preferredHostKeyAlgorithms,
		HostKeyCallback:   cb,
		Config:            ssh.Config{KeyExchanges: preferredKexAlgos},
	}, nil
}

func (a *gerritSSH) Name() string {
	return "gerrit-public-key"
}

func (a *gerritSSH) String() string {
	return a.Name()
}

func username() (string, error) {
	username := ssh_config.Get(gerritHost, "User")
	if username == "" {
		if u, err := user.Current(); err == nil {
			username = u.Username
		} else {
			username = os.Getenv("USER")
		}
	}

	if username == "" {
		return "", errors.New("failed to get username")
	}

	return username, nil
}

func signer() ssh.Signer {
	bytes, err := ioutil.ReadFile(filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa"))
	utils.CheckIfError(err)

	signer, err := ssh.ParsePrivateKey(bytes)
	utils.CheckIfError(err)
	return signer
}
