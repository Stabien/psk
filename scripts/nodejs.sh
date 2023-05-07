#!/bin/bash

JWT_KEY=$(openssl rand -base64 35)
folderName="$1"
language="$2"
path="git@github.com:Stabien/template_nodejs"

if [ $language = "JavaScript" ] 
then
  path="git@github.com:Stabien/template_nodejs_js"
fi

git clone "$path" "$folderName"

cd "$folderName"

rm -rf .git
git init

cat > .env << EOF
JWT_KEY=$JWT_KEY
EOF