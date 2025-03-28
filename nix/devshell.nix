{ inputs, ... }:

{
  imports = [ inputs.devshell.flakeModule ];

  perSystem =
    { pkgs, self', ... }:
    let

      deploy-docs = pkgs.writeShellApplication {
        name = "deploy-docs";
        meta.description = "Deploy docs";
        runtimeInputs = with pkgs; [
          rsync
          openssh
        ];
        runtimeEnv.DOCS = self'.packages.nix-versions-docs;
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
        runtimeEnv.WEB = self'.packages.nix-versions-web;
        text = ''
          ${pkgs.openssh}/bin/ssh-agent ${pkgs.bash}/bin/bash ${./web.bash}
        '';
      };

      develop-docs = pkgs.writeShellApplication {
        name = "develop-docs";
        meta.description = "Develop docs";
        runtimeInputs = with pkgs; [
          nodejs
        ];
        text = ''
          (cd docs && npm run dev)
        '';
      };

    in
    {
      devshells.default =
        { ... }:
        {
          imports = [ "${inputs.devshell}/extra/git/hooks.nix" ];

          git.hooks.pre-commit.text = "nix flake check";

          commands = [
            { package = develop-docs; }
          ];

          packages = [
            deploy-docs
            deploy-web
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
