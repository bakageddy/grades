#!/usr/bin/env sh

set -xe
ls *.go */*.go | entr -rcs "echo 'Started Server' && go run main.go"
