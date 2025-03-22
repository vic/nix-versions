package app

import (
	"fmt"

	"github.com/vic/nix-versions/packages/flake"
	"github.com/vic/nix-versions/packages/search"
)

func (a *InitArgs) Run() error {
	f := flake.New()

	if a.NVFlake != "" {
		f.Flake.OverrideInput("nix-versions", a.NVFlake)
	}

	err := a.addPackages(f)
	if err != nil {
		return err
	}

	code, err := f.Render()
	if err != nil {
		return err
	}
	fmt.Println(code)

	return nil
}

func (a *InitArgs) addPackages(f *flake.Context) error {
	var specs search.PackageSearchSpecs
	for _, pkg := range a.rest {
		s, err := search.NewPackageSearchSpec(pkg)
		if err != nil {
			return err
		}
		specs = append(specs, s)
		s.LatestUnlessExplicitConstraint()
	}

	res, err := specs.Search()
	if err != nil {
		return err
	}
	if err = res.EnsureUniquePackageNames(); err != nil {
		return err
	}

	for _, r := range res {
		f.AddTool(r)
	}

	return nil
}
