#!/bin/sh
set -e
cd $( dirname $0 )

cd openapi
go run ./main.go > v4.json

if [[ -z $( which redoc-cli ) ]]; then
  npm install -g redoc-cli
fi

redoc-cli bundle -o v4.html v4.json
