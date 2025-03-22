package app

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
)

//go:embed APP_HELP
var AppHelp string

//go:embed VERSION
var AppVersion string

//go:embed REVISION
var AppRevision string

type AppArgs struct {
	OnHelp    func() `long:"help" short:"h"`
	OnVersion func() `long:"version"`
}

func NewAppArgs() *AppArgs {
	var cliArgs AppArgs
	cliArgs.OnHelp = func() {
		fmt.Println(AppHelp)
		os.Exit(0)
	}
	cliArgs.OnVersion = func() {
		fmt.Print(strings.TrimSpace(AppVersion))
		revision := strings.TrimSpace(AppRevision)
		if revision != "" {
			fmt.Printf(" (%s)", revision)
		}
		fmt.Println()
		os.Exit(0)
	}
	return &cliArgs
}

func (cliArgs *AppArgs) ParseAndRun(args []string) error {
	parser := flags.NewParser(cliArgs, flags.IgnoreUnknown)
	extra, err := parser.ParseArgs(args)
	if err != nil {
		return err
	}

	var cmd string
	if len(extra) > 0 {
		cmd = extra[0]
	}

	if cmd == "help" {
		cliArgs.OnHelp()
		return nil
	}

	if cmd == "new" || cmd == "init" {
		return NewInitArgs().ParseAndRun(extra[1:])
	}

	if cmd == "list" || cmd == "search" {
		return NewSearchArgs().ParseAndRun(extra[1:])
	}

	// Default action is search.
	return NewSearchArgs().ParseAndRun(extra)
}
