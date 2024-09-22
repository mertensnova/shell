let
  pkgs = import <nixpkgs> {};
in
  pkgs.mkShell {
    packages = [
    pkgs.gdb
    pkgs.cmake
    pkgs.pkg-config
    pkgs.meson
    pkgs.gtest
    pkgs.ninja
];
  }
