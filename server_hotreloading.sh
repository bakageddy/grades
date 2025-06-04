#!/usr/bin/env sh

set -xe
ls *.go */*.go | entr -rcs "go build && echo 'Finished Rebuilding' && ./grades"
