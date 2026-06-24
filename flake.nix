{
  description = "slime-chunks — a Minecraft slime chunk finder";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs =
    { self, nixpkgs }:
    let
      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];
      forAllSystems = f: nixpkgs.lib.genAttrs systems (system: f nixpkgs.legacyPackages.${system});
    in
    {
      packages = forAllSystems (pkgs: rec {
        slime-chunks = pkgs.buildGoModule {
          pname = "slime-chunks";
          version = "0.1.0";

          src = self;

          vendorHash = "sha256-UvuKONqCvOodWKBsPasP3XfLSHI+Sc9Dh40iL9HcO2Q=";

          meta = {
            description = "A Minecraft slime chunk finder";
            homepage = "https://github.com/abiriadev/slime-chunks";
            license = pkgs.lib.licenses.mit;
            mainProgram = "slime-chunks";
          };
        };

        default = slime-chunks;
      });

      devShells = forAllSystems (pkgs: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gopls
            gotools
            go-tools
          ];
        };
      });

      formatter = forAllSystems (pkgs: pkgs.nixfmt);
    };
}
