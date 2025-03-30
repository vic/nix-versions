---
title: Finding packages 
order: 3
---

# Finding packages

In order to install a package, you need to know its attribute-path (read the note bellow). `nix-versions` integrates with some web services (also listed bellow) that allow you to find packages by their exposed executable programs and find their available versions. 

`nix-versions` is only possible thanks to the awesome people who have built such services.

::: note On nixpkgs attribute-paths
You can think of the nixpkgs collection as a tree of json-like objects but implemented in the nix-language. Installable packages (`derivations`) are leafs of that tree. An `attribute-path` is an string of dot-separated keys leading to an installable package.

Most of the installable packages have pretty much simple guessable attribute-paths, available from the root of the tree, like: 

> `emacs`, `ruby`, `go`, `cargo`, `zig`, `nodejs`

However other packages like `pip` are accessible under a nested package-set, and their attribute-path looks like `python312Packages.pip`. This is because sometimes a package is bound to a particular version of their runtime or simply organized as part of a particular package set, eg. there's also `rubyPackages_3_4.*`, `kdePackages.*`, etc. It pretty much depends on how the package maintainers decide to organize that tree of 200_000+ packages.
:::

## Search Backends

Currently, `nix-versions` connects to the following backends:

- [search.nixos.org](https://search.nixos.org) - This is the Official NixOS website to search for packages. Most people use it via ther web interface, but thanks to [nix-search-cli](https://github.com/peterldowns/nix-search-cli) we can query their ElasticSearch indexes for finding packages by name or provided programs (more on this bellow).

- [Lazamar index](https://lazamar.co.uk/nix-versions/) - Thanks to Lazamar, people can search historical versions of packages by channel. (You can think of a nixpkgs channel as a particular branch, eg `nixpkgs-unstable` or a particular release `nixos-24.05`, see their website for available channels.)

- [NixHub API](https://www.nixhub.io/) - The nice guys at [Jetify](https://www.jetify.com/) built this versions index for their wesome [devbox](https://www.jetify.com/devbox) product. And kindly have provided a public [API](https://www.jetify.com/docs/nixhub/).

Both Lazamar and NixHub can be used by `nix-versions` when searching for historical versions. By default, we use NixHub since it seems to be updated more frequently.

However each package spec can specify the backend used for it, and you can also change the default for those specs not specifying a particular backend with the `--nixhub` and `--lazamar` options. (read bellow)

## Finding by name / description / meta-data.

As stated above, the amazing [nix-search-cli](https://github.com/peterldowns/nix-search-cli) tool can be used from the terminal to query on [search.nixos.org](https://search.nixos.org).

```shell
# Search programs whose name matches pip
nix run nixpkgs#nix-search-cli -- --name pip --max-results 3
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./nix-search-cli-pip.ansi.html -->
</pre>
</details>

See `nix-search-cli --help` for more advanced examples.

## Finding by provided program.

This is a feature we have borrowed from `nix-search-cli`, since it might be useful to search directly by the name of a provided program.

```shell
# Packages that provide a program named exactly `gleam`
nix-versions bin/gleam@'>1.6'
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./find-gleam-program.ansi.html -->
</pre>
</details>

When multiple packages provide different implementations of a command (eg. different ruby interpreters)
```shell
nix-versions bin/ruby
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./find-ruby-program.ansi.html -->
</pre>
</details>

We recommend the [`nix-search-cli`](https://github.com/peterldowns/nix-search-cli) tool for finding the correct attribute-path of some package you need. It even lets you filter by license and other metadata, see the examples on their readme.

## Finding versions on a particular nixpkgs channel.

Using the Lazamar index, you can search for versions of programs at a particular NixOS release or nixpkgs branch. See [their webpage](https://lazamar.co.uk/nix-versions/) for existing channels you can use.

The following example lists the latest versions of emacs at `nixos-21.05`, `nixos-23.05` and `nixpkgs-unstable`.

```shell
nix-versions lazamar:nixos-21.05:emacs@ lazamar:nixos-23.05:emacs@ lazamar:nixpkgs-unstable:emacs@ --one
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./lazamar-channels.ansi.html -->
</pre>
</details>


## Setting default versions backend

A package spec can specify the backend used to search for versions with either `nixhub:` or `lazamar:<channel>:` prefixes. 

If no prefix has been specified, the default versions backend is used (NixHub). But you can also
override the default backend via the `--nixhub`, `--lazamar`, `--lazamar-channel <channel>` [options](cli-help.html).

```shell
nix-versions --lazamar emacs@ --one
nix-versions --nixhub  emacs@ --one
nix-versions lazamar:emacs@ nixhub:emacs@ --one
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./versions-backends.ansi.html -->
</pre>
</details>
