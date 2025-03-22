{ lib, config, ... }:
{
  options.nix-versions.flake =
    with lib;
    with types;
    mkOption {
      type = unspecified;
    };

  config.flake.lib.nix-versions.flake = config.nix-versions.flake;
}
