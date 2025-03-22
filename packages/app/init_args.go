package app

import (
	"github.com/jessevdk/go-flags"
)

type InitArgs struct {
	NVFlake string `long:"override-nv"`
	rest    []string
}

func NewInitArgs() *InitArgs {
	args := InitArgs{}
	return &args
}

func (a *InitArgs) Parse(args []string) error {
	parser := flags.NewParser(a, flags.AllowBoolValues|flags.IgnoreUnknown)
	rest, err := parser.ParseArgs(args)
	if err != nil {
		return err
	}
	a.rest = rest
	return nil
}

func (a *InitArgs) ParseAndRun(args []string) error {
	err := a.Parse(args)
	if err != nil {
		return err
	}
	return a.Run()
}
