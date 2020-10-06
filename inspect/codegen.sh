#!/bin/sh
set -e
cd $( dirname $0 )

go run markdown/main.go > markdown/README.md
