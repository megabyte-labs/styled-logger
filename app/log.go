package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	lggr "tideas/lggr/logger"

	"github.com/charmbracelet/glamour"
)

const (
	helpUsage = `NAME
	log - Format log message by type
EXAMPLE:
	log info "Hello World"
	availale types info, success, error, star, warn
	OR
	log --md README.md
	this will output formated README.md file contents
	type log --help for more info
	`

	wrongFileFormat = `Wrong file format, only *.md files are supported`
)

func renderMD(fn string) {

	if _, err := filepath.Match("*.md", filepath.Base(fn)); err != nil {
		log.Fatalf(wrongFileFormat)
	}

	md, err := os.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	gout, err := glamour.RenderBytes(md, "dark")
	fmt.Printf("%s", gout)
}

func renderLog(lt string, msg string) {
	lf := lggr.NewLogFactory()
	lg, ok := lf.Get(lt)
	if !ok {
		fmt.Println(helpUsage)
		return
	}
	lg.Log(msg)
}

type Runner interface {
  Init([]string) error
  Run() error
  Name() string
  Help() string
}

func main() {

	md := flag.Bool("md", false, "Outputs in terminal formated content of an *.md file\nExample: log -md README.md\n")
	_ = flag. Bool("info", false, "Outputs in terminal formated message, preceeded by blue dot symbol\nExample: log -info \"Informational Message FYI\"")
	_ = flag.Bool("success", false, "Outputs in terminal formated message, preceeded by green check sign symbol\nExample: log -success \"Success Message Congrats!\"")
  	_ = flag.Bool("error", false, "Outputs in terminal formated message, with title \"ERROR\" followed by error message\nExample: log -error \"Error Message, Fatal Error!\"")
  	_ = flag.Bool("star", false, "Outputs in terminal formated message, preceeded by star sign symbol\nExample: log -star \"Star Message, Hey!\"")
  	_ = flag.Bool("warn", false, "Outputs in terminal formated message, with title \"WARN\" followed by warning message\nExample: log -warn \"Warning Message, Hey!\"")

	  flag.Parse()

	args := flag.Args()

	if *md {
		renderMD(args[0])
		return
	}

	if len(args) < 2 {
		fmt.Println(helpUsage)
		return
	}

	renderLog(args[0], args[1])

}
