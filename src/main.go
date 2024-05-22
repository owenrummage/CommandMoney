package main

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func main() {

	var cmdLog = &cobra.Command{
		Use:                "log [amount] [reason]",
		Short:              "Log money coming in or out of the account",
		Long:               `Logs the money coming in or out of your account with a reason listed.`,
		Args:               cobra.RangeArgs(2, 2),
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			var tempTrans Transaction
			tempTrans.ID = uuid.NewString()

			amountInt, err := strconv.Atoi(args[0])
			if err != nil {
				// ... handle error
				panic(err)
			}
			tempTrans.Amount = amountInt

			tempTrans.Reason = args[1]

			AddTransaction(tempTrans)
		},
	}

	var cmdInfo = &cobra.Command{
		Use:   "info",
		Short: "Show all information about the wallet.",
		Run: func(cmd *cobra.Command, args []string) {
			transactions := GetAllTransactions()
			transTotal := len(transactions)
			upgrade := ShouldUpgradeWallet()
			walletVersion := GetWalletVersion()

			var transDeposit int
			var transWithdrawl int
			var totalBalance int

			for index, element := range transactions {
				_ = index
				if element.Amount > 0 {
					transDeposit = transDeposit + 1
				}
				if element.Amount < 0 {
					transWithdrawl = transWithdrawl + 1
				}
				totalBalance = totalBalance + element.Amount
			}

			fmt.Printf(
				`##################
  COMMAND WALLET
##################

Wallet Version: %s
CommandMoney Version: %s
Should Upgrade Wallet? %t

Stats
-------
  Total Transactions: %d
  Total Deposits: %d
  Total Withdrawls: %d
		  
  Total Balance: $%d
-------------
`, walletVersion, COMMANDMONEY_VER_STR, upgrade, transTotal, transDeposit, transWithdrawl, totalBalance)
		},
	}
	var cmdList = &cobra.Command{
		Use:                "list <number to list> <starting offset>",
		Short:              "List a number (default of 10, can be any positive number or the string \"all\") of transactions.",
		Args:               cobra.MaximumNArgs(2),
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			transactions := GetAllTransactions()
			var max int = 10
			var start int = 0
			var err error
			if len(args) > 0 {
				max, err = strconv.Atoi(args[0])
				if err != nil {
					// ... handle error
					panic(err)
				}
			}
			if len(args) > 1 {
				start, err = strconv.Atoi(args[1])
				if err != nil {
					// ... handle error
					panic(err)
				}
			}

			if max < 0 {
				max = len(transactions) - start // if negative, print until end of array
			}

			for i := 0; i <= max && i <= len(transactions)-start-1; i++ {
				element := transactions[i+start]
				fmt.Printf("Transaction %d: {id: \"%s\", amount: %d, reason: \"%s\"}\r\n", i+start, element.ID, element.Amount, element.Reason)
			}
		},
	}
	var cmdVer = &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("CommandMoney %s\r\n", COMMANDMONEY_VER_STR)
		},
	}

	var rootCmd = &cobra.Command{
		Use:     "money",
		Version: COMMANDMONEY_VER_STR,
		Short:   "CommandMoney is an efficient money management program written in Go.",
	}
	rootCmd.SetVersionTemplate("CommandMoney " + COMMANDMONEY_VER_STR + "\r\n")
	rootCmd.AddCommand(cmdLog, cmdInfo, cmdList, cmdVer)
	InitDatastore()
	rootCmd.Execute()
}
