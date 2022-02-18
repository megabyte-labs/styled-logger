/*
Copyright © 2022

*/
package cmd

import (
	"strings"
	"tideas/lggr"

	"github.com/muesli/coral"
)

// starCmd represents the star command
var starCmd = &coral.Command{
	Use:   "star",
	Short: "Prints in terminal Stared Message.",
	Long: `Outputs in terminal formated message, preceeded by star sign symbol
and usage of using your command. For example:
"Example: stylog star "Star Message, Hey!"`,
	Run: func(cmd *coral.Command, args []string) {
    m := strings.Join(args, " ")
    lggr.Star(m)
	},
}

func init() {
	rootCmd.AddCommand(starCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// starCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// starCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
