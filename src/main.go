package main

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func main() {

	var cmdLog = &cobra.Command{
		Use:                "log [operation a,s] [amount] [reason]",
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

			addTransaction(tempTrans)
		},
	}

	var cmdInfo = &cobra.Command{
		Use:   "info",
		Short: "Show all information about the wallet.",
		Run: func(cmd *cobra.Command, args []string) {
			transactions := getAllTransactions()
			transTotal := len(transactions)

			var transDeposit int
			var transWithdrawl int
			var totalBalance int
			var transactionList string

			for index, element := range transactions {
				_ = index
				if element.Amount > 0 {
					transDeposit = transDeposit + 1
				}
				if element.Amount < 0 {
					transWithdrawl = transWithdrawl + 1
				}
				totalBalance = totalBalance + element.Amount

				transactionList += fmt.Sprintf("  {id: %s, amount: %d, reason: %s}\r\n", element.ID, element.Amount, element.Reason)

			}

			fmt.Printf(`##################
  COMMAND WALLET
##################


Stats
-------
  Total Transactions: %d
  Total Deposits: %d
  Total Withdrawls: %d
		  
  Total Balance: $%d
		  
Transactions
-------------
%s`, transTotal, transDeposit, transWithdrawl, totalBalance, transactionList)
		},
	}
	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "List transactions",
		Run: func(cmd *cobra.Command, args []string) {
			transactions := getAllTransactions()

			for index, element := range transactions {
				fmt.Printf("Transaction %d: {id: %s, amount: %d, reason: %s}\r\n", index, element.ID, element.Amount, element.Reason)
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "money"}
	rootCmd.AddCommand(cmdLog, cmdInfo, cmdList)
	InitDatastore()
	rootCmd.Execute()
}
