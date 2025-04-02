---
order: -3
---

# Installing packages

As shown in the previous section, you can produce a list of Nix installables by using the `--installable` [option](getting-started/cli-help.html).

Not only you can feed that list into `nix shell` but also to any other nix-cli like `nix profile install`.
Allowing you to install a set of tools into a directory of your choose along with all its dependencies in just one command.


The following example will install `cargo` and `glibc`. Notice that for packages providing many outputs (like glibc `static`, and `dev` headers)
you can specify [output selectors](https://nix.dev/manual/nix/2.18/command-ref/new-cli/nix.html?highlight=installable#derivation-output-selection).


```shell
> mkdir tmp
> nix profile install --profile $PWD/tmp/my-tools $(nix-versions -i cargo@latest glibc^static,dev@latest)
> export PATH=$PWD/tmp/my-tools/bin:$PATH
> which cargo
/home/foo/tmp/my-tools/bin/cargo
```
