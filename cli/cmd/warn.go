/*
Copyright © 2022

*/
package cmd

import (
	"strings"
	"tideas/lggr"

	"github.com/muesli/coral"
)

// warnCmd represents the warn command
var warnCmd = &coral.Command{
	Use:   "warn",
	Short: "Prints in terminal formated Warning Message",
	Long: `Outputs in terminal formated message, with title "WARN" followed by warning message.
Example: stylog warn "Warning Message, You are warned!"`,
	Run: func(cmd *coral.Command, args []string) {
		m := strings.Join(args, " ")
    lggr.Warn(m)
	},
}

func init() {
	rootCmd.AddCommand(warnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// warnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// warnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
