package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
)

var dataPath = ""
var decoded Wallet
var initalWallet = `{
	"id": 0,
	"name": "Default Wallet",
	"currentBalance": 0,
	"transactions": []
}`

type Transaction struct {
	ID     int    `json:"id"`
	Amount int    `json:"amount"`
	Reason string `json:"reason"`
}

type Wallet struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	CurrentBalance int           `json:"currentBalance"`
	Transactions   []Transaction `json:"transactions"`
}

func InitDatastore() {
	var dir, err = os.UserHomeDir()

	if err != nil {
		panic("Error getting user home directory")
	}

	dataPath = path.Join(dir, ".config/wallet.json")

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
func createWallet() {
	var err = os.WriteFile(dataPath, []byte(initalWallet), 0755)
	if err != nil {
		panic("Unable to write to data file: " + err.Error())
	}
}

func getAllTransactions() {
	for index, element := range decoded.Transactions {
		fmt.Printf("Transaction %d: {id: %d, amount: %d, reason: %s}\r\n", index, element.ID, element.Amount, element.Reason)
	}
}
