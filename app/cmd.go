package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	flog "tideas/lggr/logger"

	"github.com/charmbracelet/glamour"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	flag "github.com/spf13/pflag"
	"golang.org/x/text/language"
)

var (
  localizer *i18n.Localizer
  bundle *i18n.Bundle
)

type logCommandName string

const (
  infoCommandName logCommandName = "info"
  warnCommandName logCommandName = "warn"
  starCommandName logCommandName = "star"
  successCommandName logCommandName = "success"
  errorCommandName logCommandName = "error"
  mdCommandName logCommandName = "md"
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
  message string
  runner func()
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

func (c *Command) Run() {
  c.runner()
}

type MDCommand struct {
  Command
  style string
  fname string
}

func (mdc *MDCommand) Run() {
  fn := mdc.fname
  if m, err := filepath.Match("*.md", filepath.Base(fn)); !m || err != nil {
		log.Fatal(getLocalized("bad_file"))
    return
	}

	md, err := os.ReadFile(fn)
	if err != nil {
		log.Fatalf( getLocalized("cant_read"), err)
	}

	gout, err := glamour.RenderBytes(md, mdc.style)
	fmt.Printf("%s", gout)
}



func NewLogginCommand(n logCommandName) CommandRunner {

  var inc CommandRunner

  lf := flog.NewLogFactory()

  command := Command{
    fs: flag.NewFlagSet(string(n), flag.ExitOnError),
      name: string(n),
  }

  // command.fs.Usage = func() {
  //   fmt.Fprintf(os.Stderr, "  %-8s" + getLocalized(string(n) + "_usage_out") + "\n%10s" + getLocalized(string(n) + "_usage_ex") + "\n", command.name, "")
  // }

  switch n {

    case infoCommandName, warnCommandName, starCommandName, successCommandName, errorCommandName:
      runner, ok := lf.Get(string(n))
      if !ok {
        panic("Not implemented")
      }

      command.runner = func() {
        runner.Log(command.message)
      }

      command.fs.StringVar(&command.message, "m", "", getLocalized(string(n) + "_flag_m"))
      inc = &command

    case mdCommandName:
      mdc := &MDCommand{
        Command{
          fs: flag.NewFlagSet("md", flag.ExitOnError),
          name: "md",
      }, "dark", ""}

      mdc.fs.StringVar(&mdc.fname, "f", "", getLocalized("md_flag_f"))
      mdc.fs.StringVar(&mdc.style, "s", "dark", getLocalized("md_flag_s"))

      inc = mdc

  }

  return inc
}

func getLocalized(id string) string {
  str, err := localizer.LocalizeMessage(&i18n.Message{
    ID: id,
  })
  if err != nil {
    log.Fatalf("Cant localize")
  }

  return str
}

func i18init() {
  bundle = i18n.NewBundle(language.English)
  bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
  _, err := bundle.LoadMessageFile("app/resources/en.json")
  if err != nil {
    log.Fatalf("Unable to read language file: %v", err)
  }
  localizer = i18n.NewLocalizer(bundle)
}
