package songbirds

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func LookUpWlanInterface(os string) (string, error) {
	switch os {
	case "linux":
		return lookupWlanInterfaceLinux()
	}

	return "", errors.New(fmt.Sprintf("Not implements LookUpWlanInterface() for your Operation system: '%s'", os))
}

func EnumInterfaces(os string) ([]string, error) {

	switch os {
	case "linux":
		return enumEnableInterfacesLinux()
	}

	return []string{}, nil
}

func lookupWlanInterfaceLinux() (string, error) {

	if wlanInterface, err := runIwConfig(); err == nil {
		return wlanInterface, nil
	}

	interfaces, err := enumEnableInterfacesLinux()

	if err != nil {
		return "", err
	}

	for _, name := range interfaces {
		if strings.Contains(name, "wl") {
			return name, nil
		}
	}

	return "", errors.New("Failed lookup wlan like interface, should be specify yourself")
}

func runIwConfig() (string, error) {
	stdout, _, err := RunCommand("iwconfig")

	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(stdout, "\n") {
		if !strings.Contains(line, "no wireless extensions.") {
			strs := strings.Split(line, " ")
			if len(strs) > 0 && len(strs[0]) > 0 {
				return strs[0], nil
			}
		}
	}

	return "", errors.New("Cannot found enable wlan interface from iwconfig")
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
