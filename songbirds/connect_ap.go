package songbirds

import (
	"errors"
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
		passwdCmd := []string{"-c", "wpa_passphrase", accessPoint.ESSID, wpaPassword}
		connectCmd = append(connectCmd, passwdCmd...)
	}

	stdout, _, err := RunCommand(connectCmd...)

	if err != nil {
		return
	}

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
