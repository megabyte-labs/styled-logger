package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"tideas/lggr"
	lgr "tideas/lggr"

	"github.com/charmbracelet/glamour"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	flag "github.com/spf13/pflag"
	"golang.org/x/text/language"
)

var (
  localizer *i18n.Localizer
  bundle *i18n.Bundle
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

func NewMDCommand() *MDCommand {
  mdc := &MDCommand{
    Command{
      fs: flag.NewFlagSet("md", flag.ExitOnError),
      name: "md",
  }, "dark", ""}

  mdc.fs.Usage = func() {
    fmt.Fprintf(os.Stderr, "  %-8s" + getLocalized("md_usage_out") + "\n%10s"+ getLocalized("md_usage_ex") + "\n", mdc.name, "")
  }

  mdc.fs.StringVar(&mdc.fname, "f", "", getLocalized("md_flag_f"))
  mdc.fs.StringVar(&mdc.style, "s", "dark", getLocalized("md_flag_s"))

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
    fmt.Fprintf(os.Stderr, "  %-8s" + getLocalized("info_usage_out") + "\n%10s" + getLocalized("info_usage_ex") + "\n", inc.name, "")
  }

  inc.fs.StringVar(&inc.message, "m", "", getLocalized("info_flag_m"))

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
    fmt.Fprintf(os.Stderr, "  %-8s" + getLocalized("success_usage_out") + "\n%10s" + getLocalized("success_usage_ex") + "\n", scc.name, "")
  }

  scc.fs.StringVar(&scc.message, "m", "", getLocalized("success_flag_m"))

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
    fmt.Fprintf(os.Stderr, "  %-8s" + getLocalized("error_usage_out") + "\n%10s" + getLocalized("error_usage_ex") + "\n", erc.name, "")
  }

  erc.fs.StringVar(&erc.message, "m", "", getLocalized("error_flag_m"))

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
    fmt.Fprintf(os.Stderr, "  %-8s" + getLocalized("star_usage_out") + "\n%10s" + getLocalized("star_usage_ex") + "\n", stc.name, "")
  }

  stc.fs.StringVar(&stc.message, "m", "", getLocalized("star_flag_m"))

  return stc
}

type WarnCommand struct {
	Command
  message string
}

func (w *WarnCommand) Run() {
  lgr.Warn(w.message)
}

func NewWarnCommand() *WarnCommand {
  wrc := &WarnCommand{Command{
    fs: flag.NewFlagSet("warn", flag.ExitOnError),
    name: "warn",
  }, ""}

  wrc.fs.Usage = func() {
    fmt.Fprintf(os.Stderr, "  %-8s" + getLocalized("warn_usage_out") + "\n%10s" + getLocalized("warn_usage_ex") + "\n", wrc.name, "")
  }

  wrc.fs.StringVar(&wrc.message, "m", "", getLocalized("warn_flag_m"))

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
  _, err := bundle.LoadMessageFile("cmd/resources/en.json")
  if err != nil {
    log.Fatalf("Unable to read language file: %v", err)
  }
  localizer = i18n.NewLocalizer(bundle)
}

func (cr *CommandRepository) Init() {
  flag.Usage = func() {
    fmt.Fprintf(os.Stderr, lggr.HelpTitle("USAGE") + "\n  %s <command> [flags] <argument>\n\n", os.Args[0])

    fmt.Fprint(os.Stderr, lggr.HelpTitle("COMMANDS") + "\n")

    for _, cmd := range cr.cmds {
      cmd.Flags().Usage()
      fmt.Fprint(os.Stderr, "\n    flags:\n")
      fmt.Fprintf(os.Stderr, cmd.Flags().FlagUsages() + "\n")
    }

    fmt.Fprint(os.Stderr, lggr.HelpTitle("LEARN MORE") + "\n")
    fmt.Fprint(os.Stderr, "  " + getLocalized("learn_more_text") + "\n\n")
  }
}

func NewCommandRepository() *CommandRepository {
  i18init()
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


