---
title: Listing package versions 
order: 2
---

# Listing package versions.

In the following example, `emacs@*` is a package spec. Requesting emacs at any version.
And is equivalent to just `emacs` or `emacs@` with an empty constraint.

```shell
<!-- @include: ./emacs-all.ansi.bash -->
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./emacs-all.ansi.html -->
</pre>
</details>

If you want to know the emacs version available on your system's nixpkgs tree,
use the `system:` backend prefix. More on backend prefixes [here](./finding-packages.html#backend-prefixes).

```shell
<!-- @include: ./emacs-local.ansi.bash -->
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./emacs-local.ansi.html -->
</pre>
</details>


## Semantic Version Constraints.

A version constraint lets you filter by only those versions that match a particular release set.

This is a very useful feature when you need to keep your tools in a range of known, stable versions.
For example, you might need that your compiler/interpreter is always compatible with your current code, even if the nixpkgs tree is bleeding edge and contains latest versions that you might be not be ready to use.

### SemVer Constraint Syntax

::: tip [SemVer Documentation](https://github.com/Masterminds/semver?tab=readme-ov-file#basic-comparisons).

For more information on the supported constraint syntax, read the documentation of the library we use:  [semver constraints](https://github.com/Masterminds/semver?tab=readme-ov-file#checking-version-constraints).

:::

Using the previous emacs example, lets filter by just a pair of release series.

```shell
<!-- @include: ./emacs-27-29.ansi.bash -->
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./emacs-27-29.ansi.html -->
</pre>
</details>

### `--all` (short `-a`)
Use `--all` to visualize the matching versions compared to all others.

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


### `--one` (short `-1`)

Use `--one` to show only the latest version that matches an specified constraint.

```shell
<!-- @include: ./emacs-27-29-one.ansi.bash -->
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./emacs-27-29-one.ansi.html -->
</pre>
</details>


## Regexp Version Constraints

Sometimes packages do not follow SemVer conventions. In those cases you can specify
a regular expression to match on versions.

::: info Enable Regexps by ending your constraint with the `$` symbol.

Since `$` is an invalid character on SemVer constraint syntax, we use it to identify
when a constraint should be matched as a regexp. So, always try to start your regexp
expression with `^` and end it with `$`.
:::


```shell
<!-- @include: ./leanify-regexp-all.ansi.bash -->
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
<!-- @include: ./leanify-regexp-all.ansi.html -->
</pre>
</details>

