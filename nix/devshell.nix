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

      develop-docs = pkgs.writeShellApplication {
        name = "docs";
        runtimeInputs = [
          pkgs.go
        ];
        text = ''
          (
            cd "''${PROJECT_ROOT:-"$(git rev-parse --show-toplevel)"}"/docs
            npm run dev
          )
        '';
        meta.description = "Docs devserver";
      };

      gen-ansi-html = pkgs.writeShellApplication {
        name = "gen-ansi-html";
        runtimeInputs = [
          go-nix-versions
          pkgs.nodejs
          pkgs.findutils
          pkgs.ansi2html
        ];
        text = ''
          # shellcheck disable=SC2016
          find . -name '*.ansi.bash' -print0 | xargs -0 -I FILE bash -v -c 'bash FILE 2>&1 | ansi2html -i -a -p > $(echo FILE | sed -e s/.bash/.html/)'
        '';
        meta.description = "Generate HTML from ANSI (*.ansi.bash -> *.ansi.html)";
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
            { package = develop-docs; }
            { package = gen-ansi-html; }
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
          ];

          packagesFrom = [
            self'.packages.default
          ];
        };
    };
}
