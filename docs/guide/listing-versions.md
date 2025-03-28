______________________________________________________________________

## layout: doc outline: 1 title: Listing package versions prev: text: Installing link: /installing next: text: Finding packages link: /finding-package-names

# Listing available package versions.

The following command will list known versions of the `emacs` package.

```shell
nix-versions emacs@
```

::: details see command output
:::

In the previous example, `emacs@` is a package spec, and is equivalent to `emacs@*`, that is, emacs at any version (no-constraint).

If you dont include an `@` symbol, `nix-versions` will show what version is available on your system's
nixpkgs tree.

```shell
nix-versions emacs
```

::: details see command output
:::

# Version Constraints.

A version constraint lets you filter by only those versions that match a particular release set.

This is a very useful feature when you need to keep your tools in a particular range of stable versions.
For example, you might need that your compiler/interpreter is in a known version-range compatible with your current code, even if the nixpkgs tree is bleeding edge and contains latest versions that you might be not be ready to use.

For more information on the supported constraint syntax, read the documentation of the library we use:  [semver constraints](https://github.com/Masterminds/semver?tab=readme-ov-file#checking-version-constraints).

Using the previous emacs example, lets filter by just a pair of release series.

```shell
# show only emacs 27 and 29 release series.
nix-versions 'emacs@~27 || ~29'
```

::: details see command output
:::

::: tip Tip: use `--all` (short: `-a`) to visualize the matching versions compared to all others.
:::

```shell
nix-versions 'emacs@~27 || ~29' --all
```

::: details see command output
:::

As you can see, coloring can help visualising the selected versions matching an specified constraint and also
the latest version in all the set. You can turn off colors using the `--color=false` flag.

::: tip Tip: use `--one` (short: `-1`) to show only the latest version that matches an specified constraint.
:::

```shell
nix-versions 'emacs@~27 || ~29' --one
```

::: details see command output
:::
