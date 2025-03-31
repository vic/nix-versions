---
order: -1
---

# Compatibility with non-flakes

We also provide a couple of experimental endpoints, `default.{nix,zip}`.

They are intended for usage on non-flake environments.

```shell
curl https://nix-versions.alwaysdata.net/default.nix/ruby@latest -o default.nix

nix-shell https://nix-versions.alwaysdata.net/default.zip/ruby@latest -A devShell
```
The reason we mention them as *experimental* is because generating a `default.nix` file
requires our server to download and generate a sri-checksum for fetching each tools' 
nixpkgs with `fetchurl`. Doing so can consume resources (time&network) on our server.
And given that many people is moving to using flakes, we are marking non-flakes features
as _experimental_ :).


