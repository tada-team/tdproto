#!/bin/bash
set -euo pipefail
shopt -s extdebug
IFS=$'\n\t'

TEMP_DIR="$(mktemp --directory --suffix=tdproto_openapi)"
TEMP_FILE_PATH="${TEMP_DIR}/tdproto.json"

go run ./ > "$TEMP_FILE_PATH"

openapi-spec-validator "$TEMP_FILE_PATH"

rm -r "$TEMP_DIR"
