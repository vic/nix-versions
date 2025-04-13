# -*- mode: sh -*-
# shellcheck shell=bash

use_nix_installables() {
    direnv_load nix shell "${@}" -c $direnv dump
}

use_nix_tools() {
   declare -a args
   if test -z "${1:-}"; then
       watch_file "$PWD/.nix_tools"
       args+=("--read" "$PWD/.nix_tools")
   fi
   while test -n "$1"; do
       if test -f "$1"; then
           watch_file "$1"
           args+=("--read" $1)
       else
           args+=("$1")
       fi
       shift
   done
   use_nix_installables $(nix-versions --installable "${args[@]}")
}
