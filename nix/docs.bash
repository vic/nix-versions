#!/usr/bin/env bash
set -e -a -u -o pipefail

# Only run on CI
if [ -z "${CI:-}" ]; then
  echo "This script is meant to be run on CI only"
  exit 1
fi

mkdir -p $HOME/.ssh
echo "StrictHostKeyChecking no" >> $HOME/.ssh/config
ssh-add <(echo "$WEB_ADMIN_DEPLOY_KEY") 2>&1>/dev/null
cd docs
npm ci
mkdir -p node_modules/vitepress/lib/app/temp
npm run build -- --outDir $HOME/docs-built
rsync -avPz -e ssh --delete $HOME/docs-built/* "$WEB_ADMIN_HOST":"~/www"
