{
  perSystem =
    { pkgs, ... }:
    let
      nix-versions = pkgs.buildGoModule {
        pname = "nix-versions";
        src = ./..;
        version = pkgs.lib.trim (builtins.readFile ./../packages/app/VERSION);
        vendorHash = "sha256-KZSWUaiG0hhhL13GxIue4CzWACmzFK96fAZqejAihqU=";
        meta = with pkgs.lib; {
          description = "CLI for searching nix packages versions using lazamar or nixhub, written in Go";
          homepage = "https://github.com/vic/nix-versions";
        };
      };

    in
    {

      packages = {
        default = nix-versions;
        inherit nix-versions;
      };

    };
}
