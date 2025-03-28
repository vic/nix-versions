{ inputs, ... }:

{
  imports = [ inputs.devshell.flakeModule ];

  perSystem =
    { pkgs, self', ... }:
    let
      deploy = pkgs.writeShellApplication {
        name = "deploy";
        meta.description = "Deploy web";
        runtimeInputs = with pkgs; [
          coreutils
          nodejs
          curl
          rsync
          openssh
        ];
        runtimeEnv.WEB = self'.packages.nix-versions-web;
        text = ''
          ${pkgs.openssh}/bin/ssh-agent ${pkgs.bash}/bin/bash ${./deploy.bash}
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
            { package = deploy; }
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
