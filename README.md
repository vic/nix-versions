# `nix-versions` - All versions of nixpkgs at your fingerprints.

`nix-versions` is a CLI tool that can help you [find the nixpkgs revision for particular versions of a package](https://nix-versions.alwaysdata.net/getting-started/listing-versions.html).

It does so by integrating excelent services available on the nix ecosystem:

We use https://search.nixos.org (via the [nix-search-cli](https://github.com/peterldowns/nix-search-cli) elastic-search client) to search packages by program name.
And https://nixhub.io API or https://lazamar.co.uk/nix-versions/ as backend for finding available versions.
It also features filtering by [version constraints](https://github.com/Masterminds/semver?tab=readme-ov-file#hyphen-range-comparisons) letting you restrict to an specific release series when needed.

When used in conjuction with `nix shell` and `direnv`, `nix-versions` can also double as a plain-text [development shell and tools version manager with automatic environment loading](https://nix-versions.alwaysdata.net/tools-version-manager.html)

Our website provides a flakes for you to use as inputs on your flakes, [providing constrained version updates](https://nix-versions.alwaysdata.net/flake-generator.html) that can guarantee your environment stability.

Read more about these features on our [documentation website](https://nix-versions.alwaysdata.net/)

### https://nix-versions.alwaysdata.net
