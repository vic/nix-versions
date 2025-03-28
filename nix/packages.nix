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

      # hm this is not building `vitepress build` hangs
      docs = pkgs.buildNpmPackage {
        name = "nix-versions-site";
        src = ./../docs;
        npmDepsHash = "sha256-DiFgdB7XMCWpemqcifoVpEsMskZTFltB6fSh7frwjq0=";
        buildPhase = ''
          ls -la
          node_modules/.bin/vitepress build | tee build.log
          mv .vitepress/dist $out
        '';
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
