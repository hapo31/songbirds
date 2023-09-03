package songbirds

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type AccessPoint struct {
	SSID      string `json:"ssid"`
	RSSI      int    `json:"rssi"`
	ESSID     string `json:"essid"`
	Encrypted bool   `json:"encrypted"`
	Channel   int    `json:"channel"`
}

func ScanAccessPoint(wlanInterface string, os string) (accessPoints []AccessPoint, err error) {

	switch os {
	case "linux":
		return scanLinux(wlanInterface)

	}

	return []AccessPoint{}, nil
}

func scanLinux(wlanInterface string) (accessPoints []AccessPoint, err error) {
	stdout, _, err := RunCommand("iwlist", wlanInterface, "scan")

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(strings.NewReader(stdout))

	var buf bytes.Buffer

	startEntryRegExp, err := regexp.Compile(`Cell\s\d+\s-\s(Address:.*)`)
	if err != nil {
		return nil, err
	}

	for scanner.Scan() {
		line := scanner.Text()
		if startEntryRegExp.MatchString(line) {
			entryOutput := buf.String()
			if len(entryOutput) > 0 {
				accessPoint, err := parseIwlistEntryLinux(entryOutput)
				if err != nil {
					return nil, err
				}
				accessPoints = append(accessPoints, accessPoint)
				buf.Reset()
			}
			matches := startEntryRegExp.FindStringSubmatch(line)
			buf.WriteString(matches[1] + "\n")
		} else {
			buf.WriteString(strings.Trim(line, " ") + "\n")
		}
	}

	return
}

func parseIwlistEntryLinux(output string) (result AccessPoint, err error) {

	fmt.Println(output)

	if err != nil {
		return
	}

	for _, line := range strings.Split(output, "\n") {
		entries := strings.Split(line, ":")

		if len(entries) <= 0 {
			continue
		}

		entry := entries[0]
		value := strings.Join(entries[1:], "")

		switch entry {
		case "Address":
			result.SSID = value
		case "Channel":
			parsed, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				continue
			}
			result.Channel = int(parsed)

		case "Encryption key":
			if value == "on" {
				result.Encrypted = true
			} else {
				result.Encrypted = false
			}

		case "ESSID":
			result.ESSID = strings.Trim(value, "\"")
		}
	}

	return
}
