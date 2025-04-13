---
home: true
title: Home

actions:
  - text: Get Started
    link: /getting-started/installing.html
    type: primary

  - text: Visit at GitHub
    link: https://github.com/vic/nix-versions
    type: secondary

features:
  - title: Works for everyone, at any level.
    details: |
      <br/>
      As a friendly CLI to explore available versions.<br/><br/>
      As a tools version manager and development shell.<br/><br/>
      As a flake generator for pinned-version packages.
  - title: Combine latest and dependable, stable versions.
    details: |
      <br/>
      Keep your compilers and tooling at <em>known, stable releases</em> that work with your current code.<br/><br/>
      And use the most recent version for other nixpkgs.
  - title: Plays well with existing tools.
    details: |
      <br/>
      Friendly with <code>nix shell</code>, <code>direnv</code> and advanced nix environments like <code>devenv</code> and <code>devshell</code>.<br/><br/>
      Usable as an <code>input</code> for any Flake and with <code>fetchurl</code> for non-flakes.
  - title: Powered by nix but requires <small>no-wizardy&trade;</small>
    details: |
      <br/>
      All the advantages of Nix without much effort.
      Reproducibility, security checksums, sandboxed builds, caching,
      remote building.
      <br/>
      <br/>
      Re-use your existing <code>.ruby-version</code>, <code>.node-version</code> project files.

footer: Kindly hosted by <a href="https://alwaysdata.com">AlwaysData</a> | Made with <3 by <a href="https://x.com/oeiuwq">@oeiuwq</a> and <a href="https://github.com/vic/nix-versions/graphs/contributors">contributors</a>.
---

<br/>
<br/>

# Try it now.
<br/>

### As a command line utility
```shell
nix run github:vic/nix-versions -- 'emacs@~27 || ~29' --all
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./getting-started/emacs-27-29-all.ansi.html -->
</pre>
</details>


### As a flake.nix generator
> You can use our `flake.zip` endpoint as an input on your own `flake.nix` or `devenv.yaml`.<br/>
> There's also `flake.nix` endpoint that outputs a text file.
```shell
nix develop 'https://nix-versions.alwaysdata.net/flake.zip/cowsay@latest/go@1.24.x' --output-lock-file /dev/null
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./flake-zip-cowsay-develop.ansi.html -->
</pre>
</details>

### As a direnv shell generator
> Our `use_nix_tools.sh` endpoint can get you a `direnv` shell in no time!<br/>
> You don't even need `nix-versions` installed, just `nix` and `direnv`.
```shell
direnv fetchurl "https://nix-versions.alwaysdata.net/use_nix_tools.sh/ruby/cowsay"
```
