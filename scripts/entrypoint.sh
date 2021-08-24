#!/bin/bash

set -e

# Setup requirements
make gen
sleep 1
/go/bin/CompileDaemon --build="go build -o bin/goer ./cmd/goer" --exclude-dir=".git" --exclude-dir="**/**/test" --exclude-dir="**/**/gomock*" --command=./bin/goer -verbose
