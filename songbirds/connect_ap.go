package songbirds

import (
	"errors"
	"fmt"
	"os"
)

func ConnectSwitchAP(accessPoint AccessPoint, wpaPassword string, wlanInterface string, os string) (err error) {
	switch os {
	case "linux":
		err = connectSwitchAPLinux(accessPoint, wpaPassword, wlanInterface)
	case "windows":
		err = connectSwitchAPWindows(accessPoint, wpaPassword, wlanInterface)
	case "darwin":
		err = connectSwitchAPDarwin(accessPoint, wpaPassword, wlanInterface)
	}
	return
}

func connectSwitchAPLinux(accessPoint AccessPoint, wpaPassword string, wlanInterface string) (err error) {

	connectCmd := []string{"wpa_supplicant", "-B", "-i", wlanInterface}
	if accessPoint.Encrypted {
		passwdCmd := []string{"wpa_passphrase", accessPoint.ESSID, wpaPassword}
		conf, _, err := RunCommand(passwdCmd...)
		if err != nil {
			return errors.New(fmt.Sprintf("Failed to run wpa_passphrase: %v", err))
		}
		fp, err := os.OpenFile("/tmp/wpa_supplicant.conf", os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0600)
		if err != nil {
			return errors.New(fmt.Sprintf("Failed to open /tmp/wpa_supplicant.conf: %v", err))
		}

		fp.WriteString(conf)
		fp.Close()

		connectCmd = append(connectCmd, "-c", "/tmp/wpa_supplicant.conf")
	}

	fmt.Printf("%v\n", connectCmd)

	stdout, stderr, err := RunCommand(connectCmd...)

	if err != nil {
		fmt.Printf("e: %v\nstdout: %s\nstderr: %s\n", err, stdout, stderr)
		return
	}

	fmt.Printf("stdout: %s\n", stdout)

	return
}

func connectSwitchAPWindows(accessPoint AccessPoint, wpaPassword string, wlanInterface string) error {
	// err = createWpaSupplicantConf(accessPoint, wpaPassword)
	// if err != nil {
	// 	return
	// }

	// err = connectWpaSupplicant(wlanInterface)
	// if err != nil {
	// 	return
	// }

	// err = removeWpaSupplicantConf()
	// if err != nil {
	// 	return
	// }

	return errors.New("not implemented")
}

func connectSwitchAPDarwin(accessPoint AccessPoint, wpaPassword string, wlanInterface string) error {

	return errors.New("not implemented")
}
