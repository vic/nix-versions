#!/usr/bin/env bash
set -e -a -u -o pipefail

# Only run on CI
if [ -z "${CI:-}" ]; then
  echo "This script is meant to be run on CI only"
  exit 1
fi

cd docs
npm run build
cd -
ssh-add <(echo "$WEB_ADMIN_DEPLOY_KEY") 2>&1>/dev/null
rsync -avPz -e ssh --delete docs/.vitepress/dist/* "$WEB_ADMIN_HOST":"~/www"
scp "$WEB/bin/web" "$WEB_ADMIN_HOST":"~/new"
ssh "$WEB_ADMIN_HOST" -C "mv ~/new ~/web"
curl --basic --user "$WEB_ADMIN_API_KEY:" -X POST "$WEB_ADMIN_API_URL"
