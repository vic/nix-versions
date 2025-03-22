{inputs, ...}: {

  imports = [
    ./nix-versions.nix
  ];

  perSystem = ({inputs', ...}: {
    packages.nvs = inputs'.nix-versions.packages.default;
  });

}