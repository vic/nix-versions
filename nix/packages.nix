{
  perSystem =
    { pkgs, ... }:
    let

      nix-versions = pkgs.buildGoModule {
        pname = "nix-versions";
        src = ./..;
        version = "1.0.0";
        vendorHash = "sha256-yrQlhu3eDLgSJ37WyhAZSzmKpiZQYcsr4Pvbk3T4usM=";
        meta = with pkgs.lib; {
          description = "CLI for searching nix packages versions using lazamar or nixhub, written in Go";
          homepage = "https://github.com/vic/nix-versions";
          mainProgram = "nix-versions";
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
