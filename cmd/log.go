package main

import (
	"os"

	flag "github.com/spf13/pflag"
)

func exit(){
  flag.Usage()
  os.Exit(2)
}

func main() {
  cr := NewCommandRepository()
  args := os.Args[1:]
  if len(args) < 1 {
    exit()
  }
  cm, ok := cr.Get(args[0])
  if !ok {
    exit()
  }

  cm.Init(args[1:])
  cm.Run()

}
