package songbirds

type AccessPoint struct {
	MacAddr   string `json:"macaddr"`
	RSSI      int    `json:"rssi"`
	ESSID     string `json:"essid"`
	Encrypted bool   `json:"encrypted"`
	Channel   int    `json:"channel"`
}

func lookupWlanInterface() {

}

func makeCommand() {

}
