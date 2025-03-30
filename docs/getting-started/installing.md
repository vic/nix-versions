---
order: 1
---

# All you need is Nix

Since you will be installing packages from the [nixpkgs repository](https://github.com/nixos/nixpkgs),
the only requirement is to have the `nix` command on your system.

If you are running on a [NixOS System](https://nixos.org/download/) you are all set.

Otherwise you can install it on any Linux, MacOS, Windows-WSL2. We recommend either the [Lix](https://git.lix.systems/lix-project/lix-installer) or [Determinate](https://github.com/DeterminateSystems/nix-installer) nix installers.

## Running `nix-versions` directly.

Just run this command and skip all of this section.

```shell
nix run github:vic/nix-versions
```

::: note Enabling Nix Flakes and the Modern Nix command.
 If this is your first time using nix, it's possible that you might need to enable:
 `--extra-experimental-features 'flakes nix-command'`
 at the command line or [enable those features at your nix.conf file](https://www.tweag.io/blog/2020-05-25-flakes/#trying-out-flakes)
 :::

## Installing on your system.

If you want to avoid typing `nix run` everytime, you might consider
installing `nix-versions` on your profile.

```shell
nix profile install github:vic/nix-versions
nix-versions --help
```

