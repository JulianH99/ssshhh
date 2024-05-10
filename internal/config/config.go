package config

import (
	"os"
	"path"
	"regexp"
	"strings"
)

type SshConfig struct {
	Host             string
	User             string
	Domain           string
	Key              string
	AdditionalConfig map[string]string
}

func parseLinesIntoSshConfig(lines []string) SshConfig {

	re := regexp.MustCompile("\\s|=")
	sshConfig := SshConfig{}
	extras := make(map[string]string)
	for _, line := range lines {
		if line == "" {
			continue
		}

		segments := re.Split(line, 2)

		key := strings.ToLower(segments[0])

		switch key {
		case "host":
			sshConfig.Host = segments[1]
		case "user":
			sshConfig.User = segments[1]
		case "hostname":
			sshConfig.Domain = segments[1]
		case "identityfile":
			sshConfig.Key = segments[1]
		default:
			extras[segments[0]] = segments[1]
		}
	}

	sshConfig.AdditionalConfig = extras
	return sshConfig
}

func getAvailableConfigs(body []byte) []SshConfig {
	sshConfigs := make([]SshConfig, 0)

	lines := strings.Split(string(body), "\n")
	secIndexes := make([]int, 0)
	for i, line := range lines {
		if strings.HasPrefix(line, "Host ") {
			secIndexes = append(secIndexes, i)
		}
	}

	for i := 0; i < len(secIndexes)-1; i++ {
		if i < len(secIndexes)-1 {
			sshConfigs = append(sshConfigs, parseLinesIntoSshConfig(lines[secIndexes[i]:secIndexes[i+1]-1]))
		}
	}
	return sshConfigs
}

func SshConfigs() ([]SshConfig, error) {
	// get .ssh/config file
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	sshConfigPath := path.Join(homedir, ".ssh", "config")

	body, err := os.ReadFile(sshConfigPath)

	if err != nil {
		return nil, err
	}

	sshConfigs := getAvailableConfigs(body)

	return sshConfigs, nil
}