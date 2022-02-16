package main

import (
	"fmt"
	"os"
	"tideas/lggr"

	flag "github.com/spf13/pflag"
)

func exit(){
  flag.Usage()
  os.Exit(2)
}

func main() {

  i18init()

  flag.Usage = func() {
    fmt.Fprintf(os.Stderr, lggr.HelpTitle("USAGE") + "\n  %s <command> [flags] <argument>\n\n", os.Args[0])

    fmt.Fprint(os.Stderr, lggr.HelpTitle("COMMANDS") + "\n")

    // for _, cmd := range cr.cmds {
    //   cmd.Flags().Usage()
    //   fmt.Fprint(os.Stderr, "\n    flags:\n")
    //   fmt.Fprintf(os.Stderr, cmd.Flags().FlagUsages() + "\n")
    // }

    fmt.Fprint(os.Stderr, lggr.HelpTitle("LEARN MORE") + "\n")
    fmt.Fprint(os.Stderr, "  " + getLocalized("learn_more_text") + "\n\n")
  }

  args := os.Args[1:]
  if len(args) < 1 {
    exit()
  }

  cm := NewLogginCommand(logCommandName(args[0]))
  cm.Init(args[1:])
  cm.Run()
}
