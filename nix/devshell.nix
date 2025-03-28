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
          nodejs
          rsync
          openssh
        ];
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

    in
    {
      devshells.default =
        { ... }:
        {
          imports = [ "${inputs.devshell}/extra/git/hooks.nix" ];

          git.hooks.pre-commit.text = "nix flake check";

          commands = [
            { package = deploy-docs; }
            { package = deploy-web; }
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
