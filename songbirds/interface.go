package songbirds

import (
	"regexp"
	"strings"
)

func EnumInterfaces(os string) ([]string, error) {

	switch os {
	case "linux":
		return enumEnableInterfacesLinux()
	}

	return []string{}, nil
}

func enumEnableInterfacesLinux() (results []string, err error) {
	stdout, _, err := RunCommand("ip", "a")

	if err != nil {
		return
	}

	re, _ := regexp.Compile(`^\d:\s([^:]+):.*`)
	for _, line := range strings.Split(stdout, "\n") {

		if strings.Contains(line, "BROADCAST") {
			if !re.MatchString(line) {
				continue
			}
			matchGroups := re.FindStringSubmatch(line)

			results = append(results, matchGroups[1])
		}
	}

	return
}
