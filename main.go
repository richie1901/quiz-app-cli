package main

import (
	"richard_adekponya_fasttrack_cli_quizapp.com/app/commands"
	"os"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
