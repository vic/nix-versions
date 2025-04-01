# `nix-versions` - All versions of nixpkgs at your fingerprints.

`nix-versions` is a CLI tool that can help you [find the nixpkgs revision for particular versions of a package](https://nix-versions.alwaysdata.net/getting-started/listing-versions.html).

It does so by integrating the following backend services available on the nix ecosystem:

- https://search.nixos.org (via the [nix-search-cli](https://github.com/peterldowns/nix-search-cli) elastic-search client). Used to search packages by program name or attribute name wildcards.

- https://nixhub.io - as default backend for package version resolution.

- https://lazamar.co.uk/nix-versions/ - search package versions by nixpkgs channel.

- https://history.nix-packages.con - another community provided versions index.

Any of these backends can be selected as default or combined on a per package basis.


It also features [version constraints](https://github.com/Masterminds/semver?tab=readme-ov-file#hyphen-range-comparisons) filters, letting you restrict packages to a know set of compatible/stable versions that work for you.

When used in conjuction with `nix shell` and `direnv`, `nix-versions` can also double as a plain-text [development shell and tools version manager with automatic environment loading](https://nix-versions.alwaysdata.net/tools-version-manager.html)

Our website provides a flakes generator endpoint that you can use as inputs on your own flakes, [providing constrained version updates](https://nix-versions.alwaysdata.net/flake-generator.html) that can guarantee you can get new updates as long as they match your specified constraints.

Read more at our [website](https://nix-versions.alwaysdata.net/)

### https://nix-versions.alwaysdata.net
