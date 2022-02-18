/*
Copyright © 2022

*/
package cmd

import (
	"strings"
	"tideas/lggr"

	"github.com/muesli/coral"
)

// errorCmd represents the error command
var errorCmd = &coral.Command{
	Use:   "error",
	Short: "Prints in terminal formated error message.",
	Long: `Outputs in terminal formated message, with title "ERROR" followed by error message.
and usage of using your command. For example:
Example: stylog error "Error Message, Fatal Error!"`,
	Run: func(cmd *coral.Command, args []string) {
    m := strings.Join(args, " ")
    lggr.Error(m)
	},
}

func init() {
	rootCmd.AddCommand(errorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// errorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// errorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
