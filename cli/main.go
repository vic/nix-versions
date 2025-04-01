package main

import (
	"log"
	"os"
	"os/exec"

	list "github.com/vic/ntv/packages/app/list"
)

func main() {
	args, err := list.Help.ParseAndRun(os.Args)
	if err == nil {
		err = list.NewListArgs().ParseAndRun(args)
	}
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			log.Fatal(string(ee.Stderr))
			os.Exit(ee.ExitCode())
		}
		log.Fatal(err)
		os.Exit(2)
	}
}
