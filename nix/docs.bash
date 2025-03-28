#!/usr/bin/env bash
set -e -a -u -o pipefail

# Only run on CI
if [ -z "${CI:-}" ]; then
  echo "This script is meant to be run on CI only"
  exit 1
fi

mkdir -p "$HOME"/.ssh
echo "StrictHostKeyChecking no" >> "$HOME"/.ssh/config
ssh-add <(echo "$WEB_ADMIN_DEPLOY_KEY") 2>&1>/dev/null
rsync -avPz -e ssh --delete "$DOCS"/* "$WEB_ADMIN_HOST":"~/www"
