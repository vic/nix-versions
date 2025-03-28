#!/usr/bin/env bash
set -e -a -u -o pipefail

# Only run on CI
if [ -z "${CI:-}" ]; then
  echo "This script is meant to be run on CI only"
  exit 1
fi

ssh-add <(echo "$WEB_ADMIN_DEPLOY_KEY") 2>&1>/dev/null
cd docs
npm ci
npm run build
cd -
rsync -avPz -e ssh --delete docs/.vitepress/dist/* "$WEB_ADMIN_HOST":"~/www"
