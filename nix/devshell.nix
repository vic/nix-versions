{ inputs, ... }:

{
  imports = [ inputs.devshell.flakeModule ];

  perSystem =
    { pkgs, self', ... }:
    let
      go-nix-versions = pkgs.writeShellApplication {
        name = "nix-versions";
        runtimeInputs = [
          pkgs.go
        ];
        text = ''
          (
            cd "''${PROJECT_ROOT:-"$(git rev-parse --show-toplevel)"}"/cli
            go build
            env PATH="$PWD:$PATH" nix-versions --color=true "$@"
          )
        '';
        meta.description = "nix-versions (development version)";
      };

      go-web = pkgs.writeShellApplication {
        name = "web";
        runtimeInputs = [
          pkgs.go
        ];
        text = ''
          (
            cd "''${PROJECT_ROOT:-"$(git rev-parse --show-toplevel)"}"/web
            go run main.go "$@"
          )
        '';
        meta.description = "Go web backend (development version)";
      };

      restart-web = pkgs.writeShellApplication {
        name = "restart-web";
        text = ''
          curl --basic --user "$WEB_ADMIN_API_KEY account=nix-versions:" -X POST "$WEB_ADMIN_API_URL" "$@"
        '';
        meta.description = "Restart production web server";
      };
    in
    {
      devshells.default =
        { ... }:
        {
          imports = [ "${inputs.devshell}/extra/git/hooks.nix" ];

          git.hooks.enable = true;
          git.hooks.pre-push.text = "nix flake check";

          commands = [
            { package = go-nix-versions; }
            { package = go-web; }
            { package = restart-web; }
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
            pkgs.pnpm
          ];

          packagesFrom = [
            self'.packages.default
          ];
        };
    };
}
