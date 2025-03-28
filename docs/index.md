---
# https://vitepress.dev/reference/default-theme-home-page
layout: home

hero:
  name: "<code>$ nix-versions</code>"
  text: '<small>Install pkgs@any version.</small>'
  tagline: |
    Any version of <code>&gt;200_000</code> nixpkgs at your disposal.
  actions:
    - theme: brand
      text: Get Started
      link: /guide/installing
    - theme: alt
      text: View on GitHub
      link: https://github.com/vic/nix-versions

features:
  - title: Works for everyone, at any level.
    details: |
      <br/>
      As a friendly CLI to explore available versions.<br/>
      As a zero nix-knowledge tools version manager.<br/>
      As a flake generator for pinned-version pkgs.
  - title: Combine bleeding-edge and dependable, stable tools.
    details: |
      Use compilers and tooling in known, stable releases that work(tm) with your code.<br/>
      And the most recent version for other nixpkgs.
  - title: Integrates with flakes and nix modules.
    details: |
      <br />
      Focuses on giving you access to packages versions and lets other
      tools shine on their own.<br/>
      Friendly with NixOS, devenv and devshell.
  - title: No-brainer development environment.
    details: |
      Combine with <code>direnv</code> and <code>nix shell</code> and get a devshell
      that can read your existing <code>.ruby-version</code>, <code>.node-version</code>, etc files.
  - title: Secure and reproducible environments.
    details: |
      Powered by Nix. All packages and their dependencies get hashed.
      Gone are the days of same versions but different bytes.
  - title: Part of a bigger self.
    details: |
      nix-versions is part of the Ntv suite.
      With a mission of providing tools that can attract more people to
      the amazing Nix ecosystem.
---
<br/>
<br/>

# Try it now.
<br/>

###### As a command line utility
```
nix run github:vic/nix-versions -- 'emacs@~27 || ~29' --all
```

###### As a flake generator webservice.

```
nix develop 'https://nix-versions.alwaysdata.net/flake.zip/cowsay@latest/go@1.24.x' --output-lock-file /dev/null
```

same as `nix-versions --flake cowsay@latest > flake.nix'`, but the webservice can be use as input on your own flakes!.
