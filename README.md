# `nix-versions` - List Nix installable packages versions.

This tool can help you [find the nixpkgs revision](#examples) where a specific version of a package was available.

It can use https://search.nixos.org (via the [nix-search-cli](https://github.com/peterldowns/nix-search-cli) elastic-search client) to search packages by program name. And https://nixhub.io API or https://lazamar.co.uk/nix-versions/ as backend for finding available versions. It also features filtering by [version constraints](https://github.com/Masterminds/semver?tab=readme-ov-file#hyphen-range-comparisons) letting you restrict to an specific release series when needed.

When used in conjuction with `nix shell`, `nix-versions`
can also double as a plain-text [development shell](#creating-a-shell-where-latest-ruby-and-cargo-are-available)
and [tools versions manager](#reading-packages-and-version-constraints-plain-text-files)
with [direnv automatic environment](#using-direnv-to-automatically-load-an-environment).

Read [usage](#usage) for a description of `nix-versions` command line options.

`nix-versions` is part of the nascent [ntv](https://github.com/vic/ntv) suite, but can be used independently.

# Installation

The only runtime-requirement is to have [`nix`](https://nixos.org/download/#download-nix) with flakes feature installed on your operating system (any Linux, MacOS, Windows-WSL2). If you are new to Nix, we recommend using the [Determinate Nix Installer](https://determinate.systems/nix-installer/) since it enables the `nix-command flakes` features by default.

Once Nix is on your system, you can install `nix-versions` using:

```
nix profile install github:vic/nix-versions
```

Or using Go (if you have it available)

```
go install github.com/vic/nix-versions
```

# Examples

###### Show the latest known version of emacs in the local nixpkgs registry.

When a version constraint is not specified (no `@` symbol on package spec), you can query
what version of a package is current in your local nixpkgs tree.

![image](https://github.com/user-attachments/assets/0b479a15-7755-45b7-8c92-7ca827371126)

###### List available emacs versions that are part of the 29 release series.

Green is latest, Cyan are those versions also matching the constraint.
Use `--color=false` to turn off coloring visual aid.
![image](https://github.com/user-attachments/assets/72a3bea8-4e66-4407-b7e5-8e29d9d71ccd)

###### List all available emacs versions, even those not matching the version constraint.

![image](https://github.com/user-attachments/assets/5d1fb7b5-0af1-4c95-b058-f71d1470da41)

###### Show only one (the latest) version matching a constraint.

The `--one` option also has a short version `-1` (the number one)
![image](https://github.com/user-attachments/assets/52ab2515-6dba-404d-86f1-360796cc0e3d)

###### Use Lazamar index to show what the latest emacs was in `nixos-23.05` and `nixos-24.05` releases of nixpkgs.

By default package versions are searched using NixHub, but you can use Lazamar index and any
of the channels to find versions at previous nixpkgs releases. If no channel is specified, Lazamar will be
searched with the `nixpkgs-unstable` channel.

![image](https://github.com/user-attachments/assets/29db968c-2ccc-45b5-bcb5-97f0c0d7fbce)

###### Creating a shell where latest ruby and cargo are available.

Using the `--installable` (short `-i`) output format, you can create a `nix shell` from a list of installables.
![image](https://github.com/user-attachments/assets/f1fccd9f-18a8-4470-9fa5-4cd8de16758f)

###### Reading packages and version constraints plain text files.

Version constraints can be read from files. eg. `ruby@.ruby-versions` will read the `.ruby-versions` file
to actually obtain the constraint.

The `--read` (short `-r`) option can be used to read package specs from a file.
Here, `.nix-tools`, but name is not significant. Shell like comments are ignored from read files.

You can use `nix-versions` to create a `nix shell` containing those programs.

![image](https://github.com/user-attachments/assets/f0609287-ea24-4835-8391-2b685c655a64)

###### Install a set of tools into a particular directory (profile)

Using the same `.nix-tools` file from the last example, install those tools into the `my-tools` directory,
including all its support files.

![image](https://github.com/user-attachments/assets/ec40778a-bbed-485b-a16e-9ed9f0251032)

###### Using direnv to automatically load an environment

If you are already using [direnv](https://direnv.net/), just add the following to your `.envrc`:

```
direnv_load nix shell $(nix-versions -1ir .nix-tools) -c direnv dump
watch_file .nix-tools .node-version .ruby-version # or any files your versions depend upon.
```

![image](https://github.com/user-attachments/assets/c57e4d1b-f000-4d14-8937-7f40a4c32da7)

# Usage

```man
SYNOPSIS

    nix-versions [<options>] <package-spec>...
    
DESCRIPTION

    List available Nix package versions.


PACKAGE SPEC

   A package spec looks like: `<package-name>@<version-constraint>`.
   
   For example, having `emacs@~29`, `emacs` is the package-name, 
   and `~29` is a version constraint.  Using this spec, `nix-versions` will 
   list available versions of emacs that match the 29 release series.


   PACKAGE NAME

   A package name can be one of:

   * The attribute path of an installable in the nixpkgs tree.
     These are normally guessable, but some packages like pip are nested inside a package-set.

     `go`
     `emacs`
     `nodejs`
     `python312Packages.pip`
     `cargo`

     Use https://search.nixos.org or `nix run nixpkgs#nix-search-cil` to 
     find the package name for those not guessable at first try.

   * A flake installable.

     Any package provided by a flake. This will bypass version search but
     will still try to validate the installable against any specified version
     constraint.

     `nixpkgs#cargo`
     `nixpkgs/nixos-24.11#ruby`
     `github:vic/gleam-nix/main#gleam`

   * A program name.

     `bin/rustc`  - Packages providing the `rustc` program.
     `bin/*rust*` - Packages with any program containing `rust` on its name.

   VERSION CONSTRAINT

     See https://github.com/Masterminds/semver for more details on the
     syntax of constraints.

     If the value after `@` is the path to a file. That file will be read
     and its content will be used as version constraint.
     eg, for ruby you could use: `ruby@.ruby-version`

   SEARCH BACKEND

     When a spec has a version constraint (includes `@`) a backend will be
     used to search for available versions. If no version constraint is
     present, only the most-recent version known to the local nixpkgs
     instance will be shown.
     
     Default backend is https://nixhub.io. The default backend for specs
     that are not explicit about one can be changed using the 
     `--nixhub` or `--lazamar` options.

     An spec can specify a particular backend to use for it.
     `nixhub:go@latest` or `lazamar:emacs@latest`.


UNCOMPLICATED VERSION MANAGER AND DEVELOPMENT SHELL

You can store package specs in a plain-text file and
use the `--one --installable --read FILE` options to
load it into a `nix shell`.

Also, since version constraints can be read from plain
text files, you can keep using your `.java-version`/`.node-versions`/etc
files.

See the README for more examples.


OPTIONS

    --help  -h          Print this help and exit.

    --read  -r FILE     Package specs are read from FILE.

  SEARCH BACKEND

     --nixhub  -n       Set default to https://nixhub.io for version search.

     --lazamar -l       Set default to https://lazamar.co.uk/nix-versions/.

     --channel -c CHAN  Use CHAN as when searching with Lazamar.
                        Default is `nixpkgs-unstable`.

  OUTPUT FORMAT

    --json  -j          Output a JSON array of resolved packages.

    --text  -t          Output as a text table. [default]

    --installable -i    Print as a list of Nix installables.

    --flake  -f         Generate a flake. See also: `ntv init`

  TEXT OUTPUT OPTIONS

    --color -C   Use colors on text table to highlight selected versions.

    --all  -a    Show all versions even those not matching a constraint

    --one  -1    Show only the latest version matching a constraint.

```
