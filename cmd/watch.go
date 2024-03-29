package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/mostdobby/watch"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a directory argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		directory := args[0]
		err := watch.RunTest(directory)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
