package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	lgr "tideas/lggr"

	"github.com/charmbracelet/glamour"
	flag "github.com/spf13/pflag"
)

type CommandRunner interface {
  Init(args []string)
  Run()
  Name() string
  Help() string
  Flags() *flag.FlagSet
}

type Command struct{
  fs *flag.FlagSet
  name string
}

func (c *Command) Flags() *flag.FlagSet {
  return c.fs
}

func (c *Command) Init(args []string) {
  c.fs.Parse(args)
}

func (c *Command) Name() string {
  return c.name
}

func (c *Command) Help() string {
  return c.fs.FlagUsages()
}

type MDCommand struct {
  Command
  style string
  fname string
}

func (mdc *MDCommand) Run() {
  fn := mdc.fname
  if _, err := filepath.Match("*.md", filepath.Base(fn)); err != nil {
		log.Fatal("Wrong file format, only *.md files are supported")
	}

	md, err := os.ReadFile(fn)
	if err != nil {
		log.Fatalf( "Unable to read from file: %v\n", err)
	}

	gout, err := glamour.RenderBytes(md, mdc.style)
	fmt.Printf("%s", gout)
}

func NewMDCommand() *MDCommand {
  mdc := &MDCommand{
    Command{
      fs: flag.NewFlagSet("md", flag.ExitOnError),
      name: "md",
  }, "dark", ""}

  mdc.fs.Usage = func() {
    fmt.Fprint(os.Stderr, "Outputs in terminal formated content of an *.md file\nExample: log md -f README.md -s dark")
  }

  mdc.fs.StringVar(&mdc.fname, "f", "", "An *.md file to output.")
  mdc.fs.StringVar(&mdc.style, "s", "dark", "Style to use for ouputing fomated MD file\nDefaults to \"dark\"\nAvailable options are: ascii|dark|light|notty|dracula")

  return mdc
}

type InfoCommand struct {
	Command
  message string
}

func (i *InfoCommand) Run() {
  lgr.Info(i.message)
}

func NewInfoCommand() *InfoCommand {
  inc := &InfoCommand{
    Command{
      fs: flag.NewFlagSet("info", flag.ExitOnError),
      name: "info",
  }, ""}

  inc.fs.Usage = func() {
    fmt.Fprint(os.Stderr, "Outputs in terminal formated message, preceeded by blue dot symbol\nExample: log info -m \"Informational Message FYI\"\n")
    fmt.Fprintln(os.Stderr, "Available flags:")
    inc.fs.VisitAll(func(f *flag.Flag) {
      fmt.Fprintf(os.Stderr, "\t--%v %v\n", f.Name, f.Usage)
    })
  }

  inc.fs.StringVar(&inc.message, "m", "", "Informational Message to be printed")

  return inc
}

type SuccessCommand struct {
	Command
  message string
}

func (s *SuccessCommand) Run() {
  lgr.Success(s.message)
}

func NewSuccessCommand() *SuccessCommand {
  scc := &SuccessCommand{Command{
    fs: flag.NewFlagSet("success", flag.ExitOnError),
    name: "success",
  }, ""}

  scc.fs.Usage = func() {
    fmt.Fprint(os.Stderr, "Outputs in terminal formated message, preceeded by green check sign symbol\nExample: log -success \"Success Message Congrats!\"")
  }

  scc.fs.StringVar(&scc.message, "m", "", "Success Message to be printed")

  return scc
}

type ErrorCommand struct {
	Command
  message string
}

func (e *ErrorCommand) Run() {
  lgr.Error(e.message)
}

func NewErrorCommand() *ErrorCommand {
  erc := &ErrorCommand{Command{
    fs: flag.NewFlagSet("error", flag.ExitOnError),
    name: "error",
  }, ""}

  erc.fs.Usage = func() {
    fmt.Fprint(os.Stderr, "Outputs in terminal formated message, with title \"ERROR\" followed by error message\nExample: log -error \"Error Message, Fatal Error!\"")
  }

  erc.fs.StringVar(&erc.message, "m", "", "Error Message to be printed")

  return erc
}

type StarCommand struct {
	Command
  message string
}

func (s *StarCommand) Run() {
  lgr.Star(s.message)
}

func NewStarCommand() *StarCommand {
  stc := &StarCommand{Command{
    fs: flag.NewFlagSet("star", flag.ExitOnError),
    name: "star",
  }, ""}

  stc.fs.Usage = func() {
    fmt.Fprint(os.Stderr, "Outputs in terminal formated message, preceeded by star sign symbol\nExample: log -star \"Star Message, Hey!\"")
  }

  stc.fs.StringVar(&stc.message, "m", "", "Stared Message to be printed")

  return stc
}

type WarnCommand struct {
	Command
  message string
}

func (s *WarnCommand) Run() {
  lgr.Star(s.message)
}

func NewWarnCommand() *WarnCommand {
  wrc := &WarnCommand{Command{
    fs: flag.NewFlagSet("warn", flag.ExitOnError),
    name: "warn",
  }, ""}

  wrc.fs.Usage = func() {
    fmt.Fprint(os.Stderr, "Outputs in terminal formated message, with title \"WARN\" followed by warning message\nExample: log -warn \"Warning Message, Hey!\"")
  }

  wrc.fs.StringVar(&wrc.message, "m", "", "Warning Message to be printed")

  return wrc
}

type CommandRepository struct {
  cmds map[string]CommandRunner
}

func (cr *CommandRepository) Register(id string, r CommandRunner) {
	cr.cmds[id] = r
}

func (cr *CommandRepository) Get(id string) (CommandRunner, bool) {
	c, ok := cr.cmds[id]
	return c, ok
}

func (cr *CommandRepository) Init() {
  flag.Usage = func() {
    fmt.Fprintf(os.Stderr, "Custom help %s:\n", os.Args[0])
    for _, cmd := range cr.cmds {
      cmd.Flags().Usage()
      cmd.Flags().VisitAll(func(f *flag.Flag) {
        fmt.Fprintf(os.Stderr, "    %v\n", f.Usage) // f.Name, f.Value
      })
    }

  }
}

func NewCommandRepository() *CommandRepository {
  cr := &CommandRepository{cmds: make(map[string]CommandRunner)} // rememeber
  cr.Register("md", NewMDCommand())
  cr.Register("info", NewInfoCommand())
  cr.Register("success", NewSuccessCommand())
  cr.Register("error", NewErrorCommand())
  cr.Register("star", NewStarCommand())
  cr.Register("warn", NewWarnCommand())
  cr.Init()
  return cr
}


