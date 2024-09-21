package commands

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Quiz CLI",
}

func init() {
	RootCmd.AddCommand(GetQuestionsCmd, SubmitAnswersCmd)
}
