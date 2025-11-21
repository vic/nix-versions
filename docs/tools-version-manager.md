---
order: -4
---

# Tools Version Manager

The nix-versions tool-manager is so simple that it does not exist. It is more of a pattern of usage when
combined with [`nix shell`](https://nix.dev/manual/nix/2.26/command-ref/nix-shell.html) and [`direnv`](https://direnv.net/).

The idea is simple:
<em>No nix-wizardy required</em>, just use plain-text files to specify your tool requirements and let `nix-versions` produce output that `nix shell` and `direnv` can use to give you an stable development environment.


::: info [Flake Generator](flake-generator.html) For Advanced Nix users
If you already know Nix, and want to use pinned-version packages as inputs for your own Nix Flake
or integrate with state-of-the-art Nix environments like
[devenv](https://devenv.sh/) or [devshell](https://github.com/numtide/devshell), `NixOS/nix-darwin/home-manager` or any other nix module class.
<b>See our [flake generator](flake-generator.html) service.</b>
:::

#### Target Audience

As a Tools Version Manager, the pattern presented on this page can replace 90% of what tools like [asdf-vm](https://asdf-vm.com/) do,
but with all the benefits you can get from Nix:
All installable tools from Nixpkgs at your fingerprints, Reproducibility, Security Checksums, Sandboxed Builds, Remote Builders, Caching, etc.
And of course, pinned version packages by `nix-versions`.

If you are new to Nix but have used other version managers like `nvm`, `rvm`, `asdf`, `mise` we want to provide you with an integrated toolset that lets you take advantage of Nix
without mandating you to learn the nix-language. By editing plain-text files and reusing your existing `.ruby-version`, `.node-version`, etc files, you can cover most of your needs.

::: info ⚡ Fast Track ⚡ - 🏃 The quick-n-dirty `use_nix_tools.sh` endpoint.

If you already have Nix and direnv installed, you can quickly get an environment ready in no time.
Note that you don't even need `nix-versions` installed locally for this to work.
Because the endpoint already resolves the nix-installales for you.
(read bellow if you want to know more on how it works).

We recommend you to look at the downloaded script beforehand.
You will notice a `use_nix_installables` function, that you can use independently of `nix-versions`.

```bash
# Place this on your .envrc
source_url "https://nix-versions.oeiuwq.com/use_nix_tools.sh/go/ruby" HASH
```

Where `HASH` can be obtained with:

```bash
direnv fetchurl "https://nix-versions.oeiuwq.com/use_nix_tools.sh/go/ruby"
```

You can obtian package updates by doing `direnv reload`.
:::

## How it works

By playing well with others. Following the UNIX philosophy of doing just ONE thing (listing package versions) and produce plain-text output that can be used by other programs to become part of something bigger.

#### The `Nix Installables` output format

By using the `--installable` (short: `-i`) [option](../getting-started/cli-help.html), `nix-versions` can produce a list of [Nix Installables <small>(flake-output-attribute)</small>](https://nix.dev/manual/nix/2.24/command-ref/new-cli/nix#flake-output-attribute) that can be read by `nix shell`.

```shell
nix-versions --installable go@1.24.x ruby@latest
```
<details><summary>see command output</summary>
<pre class="ansi-to-html">
nixpkgs/de0fe301211c267807afd11b12613f5511ff7433#go_1_24
nixpkgs/0d534853a55b5d02a4ababa1d71921ce8f0aee4c#ruby_3_4
</pre>
</details>

#### Reading package specs from a plain-text file

Instead of giving package specs as command line arguments you can use the `--read` (short `-r`) [option](../getting-started/cli-help.html) for reading them from a file.

Name of the file is not special to `nix-versions`, but we use the convention of having a `.nix_tools` file.

```shell
nix-versions --read .nix_tools --installable
```

The `.nix_tools` file can look like this:
```text
# shell-like comments are ignored.
# You can use `pkg@version` as in nix-versions command line
ruby@latest              # same as `ruby@` or `ruby@*`, ie. no version restriction.

# You can use spaces/tabs after the package name, just like .tool-versions files from asdf.
# And it also looks much cleaner.

go       >= 1.24  <1.26   # white space/tabs are not significant.
nodejs   .node-version    # read version constraint from an existing file.
```

As you can see from the previous example, you can re-use your existing `.ruby-version`, `.node-version`, `.tool-versions`, etc files that might be already in use in your project.


## Entering a `nix shell` environment

Now that you have some files like `.nix_tools`, `.node-version` from the previous examples, you are ready to enter a nix development shell containing those tools, pinned to their constrained versions.

```shell
nix shell $(nix-versions -ir .nix_tools)
```

## Automatic environment loading with `direnv`

Now you might want to load the environment into your existing shell automatically, every time you enter your project directory.
The right tool for doing this is [direnv](https://direnv.net).

::: tip Quick direnv setup
If you are new to direnv. Read their [Getting-Started](https://direnv.net/#getting-started) documentation.
Or you can try using the following instructions to set it up:

```shell
nix profile install nixpkgs#direnv # install direnv on your local profile.
echo 'eval "$(direnv hook bash)"' >> $HOME/.bashrc # hook on your shell
```
:::


All you need now is to create the following file `$HOME/.config/direnv/lib/use_nix_tools.sh`. This file
will install a function that all your projects can use to load their respective environment.

```shell
mkdir -p ~/.config/direnv/lib
# You can always inspect the downloaded function before installing it
curl "https://nix-versions.oeiuwq.com/use_nix_tools.sh" -o ~/.config/direnv/lib/use_nix_tools.sh
```

Then, on your project directory, besides your `.nix_tools` file, create an `.envrc` file that will be
detected by `direnv`.

```bash
# .envrc
use nix_tools
```

::: tip Arguments to the `use_nix_tools` function.
If given no arguments, a `$PWD/.nix_tools` file will be read. But you can provide any
other file with the same format as described above or any package-spec as expected by the `nix-versions` cli.
:::

And you are set!, just `direnv allow` and enjoy using your tools.


## More advanced environments.

The Nix ecosystem has much more advanced development environments that those produced by `nix shell`.
A couple of them are [devenv](https://devenv.sh/) and [devshell](https://github.com/numtide/devshell),
that provide more advanced features than simply loading environment variables.

They have different features depending on your needs, but they can do process management, services, deployment of containers, git workflow hooks, and much more. Be sure to read their webpages for more info.

If you happen to need any of this features, `nix-versions` can provide pinned-packages inputs for them. Be sure to read our [Flake Generator](flake-generator.html) documentation.
