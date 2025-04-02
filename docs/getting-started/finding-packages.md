---
title: Finding packages 
order: 3
---

# Finding packages

In order to install a package, you need to know its attribute-path. 

`nix-versions` is only possible thanks to the awesome people who have built public web services that allow us to search 
not only historic versions, but also packages by matching on their attribute-path, and provided programs.

::: note On nixpkgs attribute-paths
You can think of the nixpkgs collection as a tree of json-like objects but implemented in the nix-language. Installable packages (`derivations`) are leafs of that tree. An `attribute-path` is an string of dot-separated keys leading to an installable package.

Most of the installable packages have pretty much simple guessable attribute-paths, available from the root of the tree, like: 

> `emacs`, `ruby`, `go`, `cargo`, `zig`, `nodejs`

However other packages like `pip` are accessible under a nested package-set, and their attribute-path looks like `python312Packages.pip`. This is because sometimes a package is bound to a particular version of their runtime or simply organized as part of a particular package set, eg. there's also `rubyPackages_3_4.*`, `kdePackages.*`, etc. It pretty much depends on how the package maintainers decide to organize that tree of 200_000+ packages.
:::

## Search Backends

Currently, `nix-versions` can interface to the following backends:

- [search.nixos.org](https://search.nixos.org) - This is the official NixOS website used to search for packages. Most people use it via ther web interface, but thanks to [nix-search-cli](https://github.com/peterldowns/nix-search-cli) we can query their ElasticSearch database and find packages by name or provided programs (more on this bellow).

- [Lazamar index](https://lazamar.co.uk/nix-versions/) - Thanks to Lazamar, people can search historical versions of packages by channel. (You can think of a nixpkgs channel as a particular branch, eg `nixpkgs-unstable` or a particular release `nixos-24.05`, see their website for available channels.)

- [NixHub API](https://www.nixhub.io/) - The nice guys at [Jetify](https://www.jetify.com/) built this versions index for their wesome [devbox](https://www.jetify.com/devbox) product. And kindly have provided a public [API](https://www.jetify.com/docs/nixhub/).

- [history.nix-packages.com API](https://history.nix-packages.com) - Another community provided API providing historic versions for nixpkgs.


When searching for packages by wildcards on their attribute-name or provided program-names, we use `search.nixos.org`.

For searching versions, we can use any of the other backends. 
NixHub is pre-selected as default versions backend, since it seems to be updated more frequently.

## Backend Prefixes

For each package spec it is possible to specify the backend that will be used to search for versions of it 
by using a prefix like `nixhub:`, `lazamar:<channel>`, `history:`, `system:`.

Of these `system:` is the only one that is not a remote service. It will query your local `nixpkgs` tree and return the current version
that is available for some package.

You can mix them on a single command line to compare version availability on the different indexes:

```shell
<!-- @include: ./emacs-all-backends.ansi.bash -->
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./emacs-all-backends.ansi.html -->
</pre>
</details>


You can also change the default backend for those that do not use a particular backend-prefix by using the `--nixhub`, `--lazamar` and `--history` [options](cli-help.html).

#### Interlude (`nix-search-cli`)

We recommend the amazing [nix-search-cli](https://github.com/peterldowns/nix-search-cli) tool 
which we use internally to interact with [search.nixos.org](https://search.nixos.org), but that
you can use from your terminal to search for packages by name, description, programs, license, homepage, and more.

```shell
# Search programs whose name matches pip
nix run nixpkgs#nix-search-cli -- --name pip --max-results 3
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./nix-search-cli-pip.ansi.html -->
</pre>
</details>

Be sure to checkout their help documentation `nix-search-cli --help` for more advanced examples.

## Wildcards in package names

Thanks to our integration with `nix-search-cli`,
It is possible to include the `*` wildcard in a package name.

```shell
<!-- @include: ./cargo-star.ansi.bash -->
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./cargo-star.ansi.html -->
</pre>
</details>



## Packages by program name

It is possible to list packages that provide a given program by
prefixing with `bin/`.

```shell
<!-- @include: ./emacsclient-one.ansi.bash -->
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./emacsclient-one.ansi.html -->
</pre>
</details>


Listing different ruby implementations.
```shell
<!-- @include: ./find-ruby-program.ansi.bash -->
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./find-ruby-program.ansi.html -->
</pre>
</details>


## Packages by program wildcard

The program can also include `*` wildcards.

```shell
<!-- @include: ./bin-cargo-star.ansi.bash -->
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./bin-cargo-star.ansi.html -->
</pre>
</details>

## Finding versions on a particular nixpkgs channel.

Using the Lazamar index, you can search for versions of programs at a particular NixOS release or nixpkgs branch. See [their webpage](https://lazamar.co.uk/nix-versions/) for existing channels you can use.

The following example lists the latest versions of emacs at `nixos-21.05`, `nixos-23.05` and `nixpkgs-unstable`.

```shell
<!-- @include: ./lazamar-channels.ansi.bash -->
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./lazamar-channels.ansi.html -->
</pre>
</details>


## Setting default versions backend

If no prefix has been specified, the default versions backend is used (NixHub). But you can also
override the default backend via the `--nixhub`, `--history`, `--channel <lazamar-channel>` [options](cli-help.html).

```shell
<!-- @include: ./versions-backends-lazamar.ansi.bash -->
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./versions-backends-lazamar.ansi.html -->
</pre>
</details>
