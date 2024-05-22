package main

func ShouldUpgradeWallet() bool {
	walletVersion := GetWalletVersion()
	if walletVersion == "" {
		// pre 0.0.1 wallet
		return true
	} else if walletVersion == "v0.0.1" {
		return false
	} else {
		// unknown
		return false
	}
}
