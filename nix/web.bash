#!/usr/bin/env bash
set -e -a -u -o pipefail

# Only run on CI
if [ -z "${CI:-}" ]; then
  echo "This script is meant to be run on CI only"
  exit 1
fi

mkdir -p $HOME/.ssh
echo "StrictHostKeyChecking no" >>$HOME/.ssh/config
ssh-add <(echo "$WEB_ADMIN_DEPLOY_KEY") 2>&1 >/dev/null
scp "$WEB/bin/web" "$WEB_ADMIN_HOST":"~/new"
ssh "$WEB_ADMIN_HOST" -C "mv ~/new ~/web"
sleep 3
curl --basic --user "$WEB_ADMIN_API_KEY account=nix-versions:" -X POST "$WEB_ADMIN_API_URL"
