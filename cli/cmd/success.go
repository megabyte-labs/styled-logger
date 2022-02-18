/*
Copyright © 2022

*/
package cmd

import (
	"strings"
	"tideas/lggr"

	"github.com/muesli/coral"
)

// successCmd represents the success command
var successCmd = &coral.Command{
	Use:   "success",
	Short: "Prints in terminal formated success message",
	Long: `Outputs in terminal formated message, preceeded by green check sign symbol.
Example: stylog success "Success Message Congrats!"`,
	Run: func(cmd *coral.Command, args []string) {
		m := strings.Join(args, " ")
    lggr.Success(m)
	},
}

func init() {
	rootCmd.AddCommand(successCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// successCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// successCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
