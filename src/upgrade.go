package main

import "fmt"

type UpgradeFunc func()
type VerInfo struct {
	version       string
	shouldUpgrade bool
	upgradeFunc   UpgradeFunc
}

var verinfo = []VerInfo{
	{"", true, v000_to_v001},
	{"v0.0.1", false, nil},
}

func ShouldUpgradeWallet() bool {
	walletVersion := decoded.Version
	for _, info := range verinfo {
		if walletVersion == info.version {
			return info.shouldUpgrade
		}
	}
	return false // Default case if version is not found
}

func doUpgrade(maxVer string) {
	for _, info := range verinfo {
		walletVersion := decoded.Version
		if walletVersion != "" && walletVersion == maxVer {
			break
		}
		if walletVersion == info.version && info.upgradeFunc != nil {
			info.upgradeFunc()
			break
		}
	}
}

func v000_to_v001() {
	fmt.Println("Upgrading v0.0.0 wallet to v0.0.1 wallet...")

	// no additional changes were made other than adding the version tag
	decoded.Version = "v0.0.1"
	writeData(encodeData())
}
