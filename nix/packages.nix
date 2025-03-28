{
  perSystem =
    { pkgs, ... }:
    let

      nix-versions = pkgs.buildGoModule {
        pname = "nix-versions";
        src = ./../cli;
        version = "1.0.0";
        vendorHash = builtins.readFile ./../cli/vendor-hash;
        meta = with pkgs.lib; {
          description = "Go CLI for searching nix packages versions using lazamar or nixhub";
          homepage = "https://nix-versions.alwaysdata.net";
          mainProgram = "nix-versions";
        };
      };

      # Fix me, esbuild hangs, dunno why.
      docs = pkgs.buildNpmPackage {
        name = "nix-versions-site";
        src = ./../docs;
        npmDepsHash = builtins.readFile ./../docs/vendor-hash;
        buildPhase = ''
          export SASS_EMBEDDED_BIN_PATH="${pkgs.dart-sass}/bin/sass"
          mkdir -p $HOME/{temp,cache}
          npm run build -- --debug --clean-cache --clean-temp --temp $HOME/temp --cache $HOME/cache --dest $out
        '';
        dontInstall = true;
        meta = with pkgs.lib; {
          description = "Site for docs and flake generation services.";
          homepage = "https://nix-versions.alwaysdata.net";
        };
      };

      web = pkgs.buildGoModule {
        pname = "nix-versions-web";
        src = ./../web;
        version = "1.0.0";
        vendorHash = builtins.readFile ./../web/vendor-hash;
        env.CGO_ENABLED = 0; # static build
        meta = with pkgs.lib; {
          description = "Web UI for docs and flake generation services.";
          homepage = "https://nix-versions.alwaysdata.net";
          mainProgram = "web";
        };
      };

    in
    {

      packages = {
        default = nix-versions;
        inherit nix-versions;

        nix-versions-web = web;
        nix-versions-docs = docs;
      };

      checks = {
        inherit web nix-versions;
      };

    };
}
