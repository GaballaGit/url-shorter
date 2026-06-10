{
  description = "Basic flake for url shorter awip;hgwaiopgh";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = nixpkgs.legacyPackages.${system};

	go-migrate-postgres = pkgs.go-migrate.overrideAttrs (old: {
					tags = [ "postgres" ];
				});
      in {
        devShell = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            go-migrate-postgres
          ];
        };
      }
    );
}
