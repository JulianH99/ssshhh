package config

import (
	"fmt"
	"os"
	"strings"
)

func sshConfigToBytes(sshConfig SshConfig) string {
	var strb strings.Builder

	// set the path to have ~ on the start, better for windows
	// configurations
	sshRelativeIndex := strings.Index(sshConfig.Key, ".ssh")
	sshTildePath := fmt.Sprintf("~/%s", sshConfig.Key[sshRelativeIndex:])

	strb.WriteString("\n")
	strb.WriteString(fmt.Sprintf("Host %s\n", sshConfig.Host))
	strb.WriteString(fmt.Sprintf("User %s\n", sshConfig.User))
	strb.WriteString(fmt.Sprintf("Hostname %s\n", sshConfig.Domain))
	strb.WriteString(fmt.Sprintf("IdentityFile %s\n", sshTildePath))
	for settingName, value := range sshConfig.AdditionalConfig {
		strb.WriteString(fmt.Sprintf("%s %s\n", settingName, value))
	}

	strb.WriteString("\n")

	return strb.String()
}

func AddEntryToSshConfig(sshConfig SshConfig) error {
	sshConfigPath, err := sshConfigFilePath()

	if err != nil {
		return err
	}

	content := sshConfigToBytes(sshConfig)

	file, err := os.OpenFile(sshConfigPath, os.O_APPEND, os.ModeAppend)

	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return err
	}

	return nil
}
