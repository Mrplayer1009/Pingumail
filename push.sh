#!/bin/bash

git add .
git commit --amend --no-edit
git fetch
git pull
git push origin main