{ inputs, ... }:

{
  imports = [ inputs.devshell.flakeModule ];

  perSystem =
    { pkgs, self', ... }:
    {
      devshells.default =
        { ... }:
        {
          imports = [ "${inputs.devshell}/extra/git/hooks.nix" ];

          git.hooks.pre-commit.text = "nix flake check";

          packages = [
            pkgs.gopls
            pkgs.go
            pkgs.nodejs
          ];
          packagesFrom = [
            self'.packages.default
          ];
        };
    };
}
