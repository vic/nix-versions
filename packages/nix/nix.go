package nix

import (
	"os"
	"os/exec"
	"slices"
)

var (
	flakes_enabled []string
)

func init() {
	flakes_enabled = []string{
		"--extra-experimental-features",
		"flakes nix-command",
	}
}

func Run(bin string, args ...string) (string, error) {
	cmd := exec.Command(bin, args...)
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func FlakeRun(args ...string) (string, error) {
	return Run("nix", slices.Concat(flakes_enabled, args)...)
}

func JsonToNix(json string) (string, error) {
	tmpFile, err := os.CreateTemp("", "json-to-nix-*.json")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(json); err != nil {
		return "", err
	}

	// flush
	if err := tmpFile.Close(); err != nil {
		return "", err
	}

	return Run(
		"nix-instantiate",
		"--eval",
		"--expr",
		"{f}: builtins.fromJSON (builtins.readFile f)",
		"--arg", "f", tmpFile.Name(),
	)
}

func Nixfmt(args ...string) error {
	_, err := FlakeRun(
		slices.Concat(
			[]string{"run", "nixpkgs#nixfmt-rfc-style", "--"},
			args,
		)...,
	)
	return err
}

func NixfmtCode(code string) (string, error) {
	tmpFile, err := os.CreateTemp("", "code-*.nix")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(code); err != nil {
		return "", err
	}

	// flush
	if err := tmpFile.Close(); err != nil {
		return "", err
	}

	Nixfmt(tmpFile.Name())

	res, err := os.ReadFile(tmpFile.Name())
	return string(res), err
}

func NvJSON(flakePath string) (string, error) {
	return FlakeRun("eval", "--json", (flakePath + "#lib.nix-versions"))
}
