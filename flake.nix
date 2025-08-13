{
  description = "A Nix flake for hydectl";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs =
    { self, nixpkgs }:
    {
      packages.x86_64-linux.default = nixpkgs.legacyPackages.x86_64-linux.buildGoModule {
        pname = "hydectl";
        version = "unstable"; # You might want to get this from git or a file
        src = ./.;
        vendorHash = "sha256-43umwv0IITBzNdiObxDROPoklZkrEd8Qb2F+739B6nM="; # Replace with actual hash
        proxyVendor = true;
      };

      devShells.x86_64-linux.default = nixpkgs.legacyPackages.x86_64-linux.mkShell {
        packages = with nixpkgs.legacyPackages.x86_64-linux; [
          go
          direnv
        ];
      };
    };
}
