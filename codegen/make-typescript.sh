#!/bin/sh
set -e
cd $( dirname $0 )

go run typescript/main.go > ../../tdproto-ts/src/index.ts
