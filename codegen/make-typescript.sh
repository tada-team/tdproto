#!/bin/sh
set -e
cd $( dirname $0 )

go run typescript/ > ../../tdproto-ts/src/index.ts
