package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
)

var dataPath = ""
var decoded Wallet
var walletName = ".config/wallet.json"
var walletMode = 0755
var initalWallet = `{
	"id": "0",
	"version": "` + COMMANDMONEY_VER_STR + `",
	"name": "Default Wallet",
	"currentBalance": 0,
	"transactions": []
}`

type Transaction struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
	Reason string `json:"reason"`
}

type Wallet struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	CurrentBalance int           `json:"currentBalance"`
	Transactions   []Transaction `json:"transactions"`
}

func InitDatastore() {
	var dir, err = os.UserHomeDir()

	if err != nil {
		panic("Error getting user home directory")
	}

	dataPath = path.Join(dir, walletName)

	data, err := os.ReadFile(dataPath)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Data does not exist. Initializing with default configuration.")
		createWallet()

		data = []byte(initalWallet)
	} else if err != nil {
		panic("Unable to read data file: " + err.Error())
	}
	decodeData(data)
}

func decodeData(data []byte) {
	var err = json.Unmarshal([]byte(data), &decoded)
	if err != nil {
		panic("JSON Parse Error: " + err.Error())
	}
}
func encodeData() []byte {
	var enc, err = json.Marshal(decoded)
	if err != nil {
		panic("JSON Encoding Error: " + err.Error())
	}
	return enc

}
func createWallet() {
	var err = os.WriteFile(dataPath, []byte(initalWallet), fs.FileMode(walletMode))
	if err != nil {
		panic("Unable to write to data file: " + err.Error())
	}
}

func writeData(data []byte) {
	var dir, err = os.UserHomeDir()

	if err != nil {
		panic("Error getting user home directory")
	}

	dataPath = path.Join(dir, walletName)

	err = os.WriteFile(dataPath, data, fs.FileMode(walletMode))
	if err != nil {
		panic("Unable to write data file: " + err.Error())
	}
	decodeData(data)
}

func addTransaction(transaction Transaction) {
	decoded.Transactions = append(decoded.Transactions, transaction)
	var encoded = encodeData()
	writeData(encoded)
}

func getAllTransactions() []Transaction {
	return decoded.Transactions
}
