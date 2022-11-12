#!/bin/bash

JWT_KEY=$(openssl rand -base64 35)
folderName="template_nodejs"

if [ $# -eq 0 ]
  then
    git clone "git@github.com:Stabien/template_nodejs"
else
    folderName="$1"
    git clone "git@github.com:Stabien/template_nodejs" "$folderName"
fi

cd "$folderName"

npm update
npm install

rm -rf .git
git init

cat > .env << EOF
JWT_KEY=$JWT_KEY
EOF