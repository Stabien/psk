#!/bin/bash

folderName="template_reactjs"

if [ $# -eq 0 ]
  then
    git clone "git@github.com:Stabien/template_reactjs"
else
    folderName="$1"
    git clone "git@github.com:Stabien/template_reactjs" "$folderName"
fi

cd "$folderName"

npm update
npm install

rm -rf .git
git init
