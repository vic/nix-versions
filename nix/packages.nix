{
  perSystem =
    { pkgs, ... }:
    let

      nix-versions = pkgs.buildGoModule {
        pname = "nix-versions";
        src = ./../cli;
        version = "1.0.1";
        vendorHash = builtins.readFile ./../cli/vendor-hash;
        meta = with pkgs.lib; {
          description = "Go CLI for searching nix packages versions using lazamar or nixhub";
          homepage = "https://nix-versions.oeiuwq.com";
          mainProgram = "nix-versions";
        };
      };

      web = pkgs.buildGoModule {
        pname = "nix-versions-web";
        src = ./../web;
        version = "1.0.1";
        vendorHash = builtins.readFile ./../web/vendor-hash;
        env.CGO_ENABLED = 0; # static build
        meta = with pkgs.lib; {
          description = "Web UI for docs and flake generation services.";
          homepage = "https://nix-versions.oeiuwq.com";
          mainProgram = "web";
        };
      };

      deploy-docs = pkgs.writeShellApplication {
        name = "deploy-docs";
        meta.description = "Deploy docs";
        runtimeInputs = with pkgs; [
          pnpm
          nodejs
          rsync
          openssh
        ];
        #runtimeEnv.DOCS = docs;
        text = ''
          ${pkgs.openssh}/bin/ssh-agent ${pkgs.bash}/bin/bash ${./docs.bash}
        '';
      };

      deploy-web = pkgs.writeShellApplication {
        name = "deploy-web";
        meta.description = "Deploy web";
        runtimeInputs = with pkgs; [
          curl
          openssh
        ];
        runtimeEnv.WEB = web;
        text = ''
          ${pkgs.openssh}/bin/ssh-agent ${pkgs.bash}/bin/bash ${./web.bash}
        '';
      };

    in
    {

      packages = {
        default = nix-versions;
        nix-versions-web = web;
        nix-versions-docs = docs;

        inherit
          nix-versions
          deploy-docs
          deploy-web
          ;
      };

      checks = {
        inherit web nix-versions;
      };

    };
}
