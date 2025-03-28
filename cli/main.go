package main

import (
	"log"
	"os"
	"os/exec"

	list "github.com/vic/ntv/packages/app/list"
)

func main() {
	if len(os.Args) < 2 {
		list.HelpAndExit("nix-versions", 1)
	}
	err := list.NewListArgs().ParseAndRun(os.Args[1:])
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			log.Fatal(string(ee.Stderr))
			os.Exit(ee.ExitCode())
		}
		log.Fatal(err)
		os.Exit(2)
	}
}
