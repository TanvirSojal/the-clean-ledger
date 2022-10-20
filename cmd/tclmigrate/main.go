package main

import (
	"fmt"
	"go-chain/database"
	"os"
	"time"
)

func main() {
	fmt.Println("Migrating transactions to block db...")

	state, err := database.NewStateFromDisk()

	if err != nil { 
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer state.Close()

	block0 := database.NewBlock(
		database.Hash{},
		uint64(time.Now().Unix()),
		[]database.Tx {
			database.NewTx("sojal", "sojal", 2, ""),
			database.NewTx("sojal", "sojal", 50, ""),
		},
	)

	state.AddBlock(block0)

	block0Hash, _ := state.Persist()

	block1 := database.NewBlock(
		block0Hash,
		uint64(time.Now().Unix()),
		[]database.Tx {
			database.NewTx("sojal", "sojal", 100, ""),
			database.NewTx("sojal", "sojal", 150, ""),
			database.NewTx("sojal", "sojal", 100, "reward"),
			database.NewTx("sojal", "sojal", 100, "reward"),
			database.NewTx("sojal", "sojal", 100, "reward"),
			database.NewTx("sojal", "sojal", 100, "reward"),
			database.NewTx("sojal", "sojal", 100, "reward"),
			database.NewTx("sojal", "mark", 200, ""),
		},
	)

	state.AddBlock(block1)

	state.Persist()
}