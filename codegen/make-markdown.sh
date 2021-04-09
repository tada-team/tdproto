#!/bin/sh
set -e
cd $( dirname $0 )

go run markdown/ > markdown/README.md
