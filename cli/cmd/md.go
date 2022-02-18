/*
Copyright © 2022

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/glamour"
	"github.com/muesli/coral"
)

// mdCmd represents the md command
var mdCmd = &coral.Command{
	Use:   "md",
	Short: "Prints in terminal formated md file.",
	Long: `Outputs in terminal formated content of a Markdown document (*.md) file with Glamour.
Example: stylog md README.md`,
	Run: func(cmd *coral.Command, args []string) {
		fn := args[0]
  if m, err := filepath.Match("*.md", filepath.Base(fn)); !m || err != nil {
		log.Fatal("Wrong file format, only *.md files are supported")
    return
	}

	md, err := os.ReadFile(fn)
	if err != nil {
		log.Fatalf( "Unable to read from file: %v\n", err)
	}

	gout, err := glamour.RenderBytes(md, "dark")
	fmt.Printf("%s", gout)
	},
}

func init() {
	rootCmd.AddCommand(mdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
