# `nix-versions` - All versions of nixpkgs at your fingerprints.

`nix-versions` is a CLI tool that can help you [find the nixpkgs revision for historical versions of a package](https://nix-versions.oeiuwq.com/getting-started/listing-versions.html).

It does so by integrating the following services available to the nix ecosystem:

- https://search.nixos.org (via the [nix-search-cli](https://github.com/peterldowns/nix-search-cli) elastic-search client). Used to search packages by program name or attribute name wildcards.

- https://nixhub.io - as default backend for package version resolution.

- https://lazamar.co.uk/nix-versions/ - search package versions by nixpkgs channel.

- https://history.nix-packages.con - another community provided versions index.

Any of these backends can be selected as default or combined on a per-package basis.

It also features [version constraint](https://github.com/Masterminds/semver?tab=readme-ov-file#hyphen-range-comparisons) filters, allowing you to restrict packages to a known set of compatible/stable versions that work with you.

When used in conjuction with `nix shell` and `direnv`, `nix-versions`, it can also double as a plain-text [development shell and tools version manager with automatic environment loading](https://nix-versions.oeiuwq.com/tools-version-manager.html) targeting people new to nix, but familiar with tools like npm, rvm, asdf, mise.

Our website provides a flakes generator for version-pinned packages. You can use these generated flakes as inputs on your own flakes, [allowing version constrained package updates](https://nix-versions.oeiuwq.com/flake-generator.html).

Read more at our [website](https://nix-versions.oeiuwq.com/)

### https://nix-versions.oeiuwq.com
