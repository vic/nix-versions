package app

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/vic/nix-versions/packages/marshalling"
	"github.com/vic/nix-versions/packages/search"
	lib "github.com/vic/nix-versions/packages/versions"
)

func (ctx *SearchArgs) SearchAction() error {
	if len(ctx.Names) < 1 {
		fmt.Println(AppHelp)
		os.Exit(1)
		return nil
	}
	var (
		versions []lib.Version
		err      error
		str      string
	)

	opts := search.Opts{
		One:        ctx.One,
		Exact:      ctx.Exact,
		Constraint: ctx.Constraint,
		Limit:      ctx.Limit,
		Sort:       ctx.Sort,
		Reverse:    ctx.Reverse,
		Lazamar:    ctx.Lazamar,
		Channel:    ctx.Channel,
	}

	versions, err = search.FindReadingVersionsAll(opts, ctx.Names)
	if err != nil {
		return err
	}

	if ctx.OutType == Json {
		str, err = marshalling.VersionsJson(versions)
		if err != nil {
			return err
		}
	} else if ctx.OutType == Installable {
		str = marshalling.Installables(versions)
	} else if ctx.OutType == Flake {
		str, err = marshalling.Flake(versions)
		if err != nil {
			return err
		}
	} else {
		str = marshalling.VersionsTable(versions, ctx.Color)
	}

	if ctx.One {
		var seen = make(map[string]int)
		var anyFailed = false
		for _, v := range versions {
			seen[v.Name]++
			if seen[v.Name] > 1 {
				anyFailed = true
			}
		}
		if anyFailed {
			str = marshalling.VersionsTable(versions, ctx.Color)
			fmt.Fprint(os.Stderr, str, "\n")
			return fmt.Errorf("assertion failure: expected at most one version per package, but got %v - try using @latest or a more specific version constraint", seen)
		}
	}

	if ctx.WriteTo != "" && ctx.WriteTo != "-" {
		err := os.WriteFile(ctx.WriteTo, []byte(str), 0644)
		if err != nil {
			return err
		}
		return nil
	}

	fmt.Println(str)
	return nil
}
