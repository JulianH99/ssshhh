package config_test

import (
	"fmt"
	"testing"

	"github.com/JulianH99/ssshhh/internal/config"
)

func TestSshConfigs(t *testing.T) {

	sshConfigs, err := config.SshConfigs()

	fmt.Println(sshConfigs, err)

}
