#!/bin/bash -eu

echo current version: $(gobump show -r)
read -p "input next version: " next_version

gobump set $next_version -w
git-chglog --next-tag v$next_version -o CHANGELOG.md

git add version.go CHANGELOG.md
git commit -m 'Bumps up to 'v$next_version