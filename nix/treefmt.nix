{ inputs, ... }:
{
  imports = [ inputs.treefmt-nix.flakeModule ];
  perSystem.treefmt = {
    projectRootFile = "flake.nix";
    programs.nixfmt.enable = true;
    programs.nixfmt.excludes = [ ".direnv" ];
    programs.deadnix.enable = true;
    programs.mdformat.enable = true;
    programs.mdformat.excludes = [ "docs/index.md" ];
    programs.yamlfmt.enable = true;
    programs.gofmt.enable = true;
  };
}
