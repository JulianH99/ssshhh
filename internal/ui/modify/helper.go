package modify

import "strings"

func stringToMap(s string) map[string]string {
	configs := make(map[string]string)

	for _, config := range strings.Split(s, ",") {
		config = strings.TrimSpace(config)

		if config == "" {
			continue
		}
		keyValue := strings.Split(config, " ")

		configs[keyValue[0]] = keyValue[1]
	}

	return configs
}
