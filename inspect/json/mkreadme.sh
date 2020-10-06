#!/bin/sh
set -e
cd $( dirname $0 )

dest=README.md
echo "## Protocol description in json:" > ${dest}
echo "\`\`\`json" >> ${dest}
go run . >> ${dest}
echo "\`\`\`" >> ${dest}

