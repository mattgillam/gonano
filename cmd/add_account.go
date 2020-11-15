package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addAccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Add an account within a wallet",
	Run: func(cmd *cobra.Command, args []string) {
		checkWalletIndex()
		wi := wallets[walletIndex]
		account, err := wi.w.NewAccount()
		fatalIf(err)
		wi.Accounts = append(wi.Accounts, account.Address())
		wi.save()
		fmt.Println("Added account", account.Address())
	},
}

func init() {
	addCmd.AddCommand(addAccountCmd)
}
