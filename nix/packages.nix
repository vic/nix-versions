{
  perSystem =
    { pkgs, ... }:
    let
      nix-versions = pkgs.buildGoModule {
        pname = "nix-versions";
        version = "1.0.0";
        src = ./..;
        vendorHash = "sha256-JDqKwcKyVKR/iMBEbKtPV7GL/xhw8h/plL+B0KTZwLY=";
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
