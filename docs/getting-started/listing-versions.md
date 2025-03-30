---
title: Listing package versions 
order: 2
---

# Listing package versions.

In the following example, `emacs@` is a package spec, and is equivalent to `emacs@*`, that is, emacs at any version (no-constraint).

```shell
nix-versions emacs@
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./emacs-all.ansi.html -->
</pre>
</details>


If you dont include an `@` symbol, `nix-versions` will show what version is available on your system's
nixpkgs tree.

```shell
nix-versions emacs
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./emacs-local.ansi.html -->
</pre>
</details>

# Version Constraints.

A version constraint lets you filter by only those versions that match a particular release set.

This is a very useful feature when you need to keep your tools in a range of known, stable versions.
For example, you might need that your compiler/interpreter is always compatible with your current code, even if the nixpkgs tree is bleeding edge and contains latest versions that you might be not be ready to use.

::: info Constraint syntax
For more information on the supported constraint syntax, read the documentation of the library we use:  [semver constraints](https://github.com/Masterminds/semver?tab=readme-ov-file#checking-version-constraints).
:::

Using the previous emacs example, lets filter by just a pair of release series.

```shell
# show only emacs 27 and 29 release series.
nix-versions 'emacs@~27 || ~29'
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./emacs-27-29.ansi.html -->
</pre>
</details>

::: tip Tip: use `--all` (short: `-a`) to visualize the matching versions compared to all others.
:::

```shell
nix-versions 'emacs@~27 || ~29' --all
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./emacs-27-29-all.ansi.html -->
</pre>
</details>


As you can see, coloring can help visualising the selected versions matching an specified constraint and also
the latest version in all the set. You can turn off colors using the `--color=false` [option](cli-help.html).

::: tip Tip: use `--one` (short: `-1`) to show only the latest version that matches an specified constraint.
:::

```shell
nix-versions 'emacs@~27 || ~29' --one
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./emacs-27-29-one.ansi.html -->
</pre>
</details>

