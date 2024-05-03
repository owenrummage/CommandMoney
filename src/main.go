package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func main() {

	var cmdLog = &cobra.Command{
		Use:   "log [amount] [reason]",
		Short: "Log money coming in or out of the account",
		Long:  `Logs the money coming in or out of your account with a reason listed.`,
		Args:  cobra.RangeArgs(2, 2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Logged Transaction: '" + args[1] + "' for amount " + args[0] + " at time " + time.Now().Format("2006-01-02 15:04:05"))
			// todo save transactions
		},
	}

	var cmdInfo = &cobra.Command{
		Use:   "info",
		Short: "Show all information about the wallet.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Wallet Information")
			// todo print information
		},
	}

	var rootCmd = &cobra.Command{Use: "money"}
	rootCmd.AddCommand(cmdLog, cmdInfo)
	rootCmd.Execute()
}
