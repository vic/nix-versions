{ inputs, ... }:

{
  imports = [ inputs.devshell.flakeModule ];

  perSystem =
    { pkgs, self', ... }:
    {
      devshells.default =
        { ... }:
        {
          packages = [
            pkgs.gopls
          ];
          packagesFrom = [
            self'.packages.default
          ];
        };
    };
}
