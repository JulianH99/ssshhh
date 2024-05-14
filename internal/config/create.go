package config

import (
	"os"
	"os/exec"
	"path"
	"strings"
)

type sshKeyCreateConfig struct {
	path    string
	comment string
	keyType string
}

func (s sshKeyCreateConfig) Path() string {
	return s.path
}

func CreateSshKeyConfig(p, keyType, comment string) sshKeyCreateConfig {
	home, err := os.UserHomeDir()

	if err != nil {
		home = os.Expand("$HOME", os.Getenv)
	}

	sshKeyPath := path.Join(home, ".ssh", strings.TrimSpace(p))
	comment = strings.TrimSpace(comment)
	return sshKeyCreateConfig{sshKeyPath, comment, keyType}
}

func CreateSshKey(config sshKeyCreateConfig) *exec.Cmd {
	cmd := exec.Command("ssh-keygen", "-t", config.keyType, "-f", config.path, "-C", config.comment)
	return cmd
}
