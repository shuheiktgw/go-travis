#!/bin/bash -eu

echo current version: $(gobump show -r)
read -p "input next version: " next_version

gobump set $next_version -w
git pull origin master --tag
git tag v$next_version
git-chglog -o CHANGELOG.md

git add version.go CHANGELOG.md
git commit -m 'Bumps up to 'v$next_version