#!/bin/bash

JWT_KEY=$(openssl rand -base64 35)
folderName="$1"

git clone "git@github.com:Stabien/template_nodejs" "$folderName"

cd "$folderName"

rm -rf .git
git init

cat > .env << EOF
JWT_KEY=$JWT_KEY
EOF