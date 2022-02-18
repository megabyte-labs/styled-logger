/*
Copyright © 2022

*/
package cmd

import (
	"strings"
	"tideas/lggr"

	"github.com/muesli/coral"
)

// infoCmd represents the info command
var infoCmd = &coral.Command{
	Use:   "info",
	Short: "Outputs in terminal formated Info Message",
	Long: `Outputs in terminal formated message, preceeded by blue dot symbol
Example: stylog info "Informational Message FYI"`,

  Args: coral.MinimumNArgs(1),

	Run: func(cmd *coral.Command, args []string) {
    m := strings.Join(args, " ")
    lggr.Info(m)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
