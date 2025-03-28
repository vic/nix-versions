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

          git.hooks.enable = true;
          git.hooks.pre-push.text = "nix flake check";

          commands = [
            {
              name = "develop-docs";
              help = "Local docs devserver";
              command = "cd docs; npm run dev";
            }
          ];

          env = [
            {
              # Otherwise the sass compiler fails on nixos.
              # See patches/sass-embedded-1.62.0.patch
              name = "SASS_EMBEDDED_BIN_PATH";
              value = "${pkgs.dart-sass}/bin/sass";
            }
          ];

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
