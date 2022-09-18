package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var tclCmd = &cobra.Command{
	Use: "tcl",
	Short: "The Clean Ledger CLI",
	Long: "The Clean Ledger CLI Tool",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Hello CLI!")
	},
}

func main() {
	tclCmd.AddCommand(versionCmd)
	tclCmd.AddCommand(balancesCmd())
	tclCmd.AddCommand(txCmd())

	err := tclCmd.Execute()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}