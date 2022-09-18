package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Major = "0"
const Minor = "1"
const Fix = "0"
const Verbal = "Tx Add && Balances List"

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Describes version.",
	Run: func(cmd *cobra.Command, args[] string) {
		fmt.Printf("Version: %s.%s.%s-beta %s", Major, Minor, Fix, Verbal)
	},
}