{
  perSystem =
    { pkgs, ... }:
    let

      nix-versions = pkgs.buildGoModule {
        pname = "nix-versions";
        src = ./..;
        version = "1.0.0";
        vendorHash = builtins.readFile ./vendor-hash;
        meta = with pkgs.lib; {
          description = "Go CLI for searching nix packages versions using lazamar or nixhub";
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
