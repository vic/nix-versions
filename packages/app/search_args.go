package app

import (
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/mattn/go-isatty"
)

type OutputType int

const (
	Text OutputType = iota
	Json
	Installable
	Flake
)

type SearchArgs struct {
	OnChannel     func(string) `long:"channel"`
	OnLazamar     func()       `long:"lazamar"`
	OnNixHub      func()       `long:"nixhub"`
	OnJson        func()       `long:"json"`
	OnText        func()       `long:"text"`
	OnInstallable func()       `long:"installable"`
	OnFlake       func()       `long:"flake"`
	OnWrite       func(string) `long:"out" short:"o"`
	Lazamar       bool
	Channel       string
	OutType       OutputType
	WriteTo       string
	Color         bool   `long:"color"`
	One           bool   `long:"assert-one"`
	Sort          bool   `long:"sort"`
	Reverse       bool   `long:"reverse"`
	Exact         bool   `long:"exact"`
	Limit         int    `long:"limit"`
	Constraint    string `long:"constraint"`
	Names         []string
}

func NewSearchArgs() *SearchArgs {
	var cliArgs = SearchArgs{
		Channel: "nixpkgs-unstable",
		Sort:    true,
		Color:   isatty.IsTerminal(os.Stdout.Fd()),
	}
	cliArgs.OnLazamar = func() {
		cliArgs.Lazamar = true
	}
	cliArgs.OnNixHub = func() {
		cliArgs.Lazamar = false
	}
	cliArgs.OnChannel = func(s string) {
		cliArgs.Channel = s
		cliArgs.Lazamar = true
	}
	cliArgs.OnJson = func() {
		cliArgs.OutType = Json
	}
	cliArgs.OnText = func() {
		cliArgs.OutType = Text
	}
	cliArgs.OnInstallable = func() {
		cliArgs.OutType = Installable
	}
	cliArgs.OnFlake = func() {
		cliArgs.OutType = Flake
	}
	cliArgs.OnWrite = func(s string) {
		cliArgs.One = true
		cliArgs.WriteTo = s
	}
	return &cliArgs
}

func (cliArgs *SearchArgs) Parse(args []string) error {
	parser := flags.NewParser(cliArgs, flags.AllowBoolValues|flags.IgnoreUnknown)
	extra, err := parser.ParseArgs(args)
	if err != nil {
		return err
	}
	cliArgs.Names = extra
	return nil
}

func (ctx *SearchArgs) ParseAndRun(args []string) error {
	err := ctx.Parse(args)
	if err != nil {
		return err
	}
	return ctx.SearchAction()
}
