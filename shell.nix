{pkgs ? import <nixpkgs> {}, ...}:
pkgs.mkShell {
  packages = with pkgs; [
    go
    gopls
    pre-commit
  ];
}
