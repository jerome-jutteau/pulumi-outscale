{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
    nativeBuildInputs = with pkgs.buildPackages; [ 
      gnused
      gnumake
      pulumictl
      pulumi
      pulumiPackages.pulumi-language-go
      pulumiPackages.pulumi-language-python
      pulumiPackages.pulumi-language-nodejs
      go
      nodejs_20
      yarn
      nodePackages.typescript
      python311Packages.setuptools
      python311
      dotnet-sdk_7];
}