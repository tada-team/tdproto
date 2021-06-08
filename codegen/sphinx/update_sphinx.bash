#!/bin/bash
set -euo pipefail
shopt -s extdebug
IFS=$'\n\t'

TDPROTO_SPHINX_DIR="$1"

if [ ! -d "$TDPROTO_SPHINX_DIR" ]; then
    echo "Passed path is not a directory"
    exit 1
fi

DOCS_DIR=$(readlink -f "$TDPROTO_SPHINX_DIR")"/docs"

if [ ! -d "$DOCS_DIR" ]; then
    echo "docs direcroty missing"
    exit 1
fi

go run ./events > "${DOCS_DIR}/events.rst"
go run ./json_index > "${DOCS_DIR}/data_index.rst"
go run ./paths_doc team > "${DOCS_DIR}/team_paths.rst"
go run ./paths_doc group > "${DOCS_DIR}/group_paths.rst"
go run ./paths_doc chat > "${DOCS_DIR}/chat_paths.rst"
go run ./paths_doc misc > "${DOCS_DIR}/misc_paths.rst"
go run ./paths_doc task > "${DOCS_DIR}/task_paths.rst"
