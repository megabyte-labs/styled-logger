/*
Copyright © 2022

*/
package cmd

import (
	"os"

	"github.com/muesli/coral"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &coral.Command{
	Use:   "stylog",
	Short: "Styled Terminal Logger",
	Long: `This project was built to provide styled terminal log messages
  from within shell scripts (mainly housed in a Task project). It supports
  basic log message types that are prepended with an emoji and also provides
  the capability of logging markdown files styled by Glamour.`,
}

var (
  Localizer *i18n.Localizer
  Bundle *i18n.Bundle
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
