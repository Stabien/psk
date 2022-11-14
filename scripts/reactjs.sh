#!/bin/bash

folderName="$1"

git clone "git@github.com:Stabien/template_reactjs" "$folderName"

cd "$folderName"

rm -rf .git
git init
