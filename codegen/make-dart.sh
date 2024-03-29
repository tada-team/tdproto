#!/bin/bash
set -euo pipefail
shopt -s extdebug dotglob
IFS=$'\n\t'

# Setup command line arguments
TEMP=$(getopt --options 'bf' --long 'no-build,no-format' --name "make-dart.bash" -- "$@")

eval set -- "$TEMP"
unset TEMP

RUN_BUILD=true
RUN_FORMAT=true

while true; do
    case "$1" in
        '-b'|'--no-build')
            RUN_BUILD=false
            shift
            continue
        ;;
        '-f'|'--no-format')
            RUN_FORMAT=false
            shift
            continue
        ;;
        '--')
            shift
            break
        ;;
        *)
            echo 'Internal error!' >&2
            exit 1
        ;;
    esac
done

if [ 1 -ne "${#}" ]; then
    echo "Wrong number of arguments. Should pass only 1 which is the dart library folder." >&2
    exit 1
fi

DART_LIB_FOLDER="$1"

if [ ! -d $DART_LIB_FOLDER ]; then
    echo "$DART_LIB_FOLDER is not a folder" >&2
    exit 1
fi

DART_LIB_FOLDER="$(readlink --canonicalize $DART_LIB_FOLDER)"

LIB_ENUMS_PATH="${DART_LIB_FOLDER}/lib/src/enums"
LIB_MODELS_PATH="${DART_LIB_FOLDER}/lib/src/models"
LIB_FILE_PATH="${DART_LIB_FOLDER}/lib/tdproto_dart.dart"

LIB_TEMPLATE_DIR="./dart/lib_template"

cp -r "$LIB_TEMPLATE_DIR"/* "$DART_LIB_FOLDER"

if [ -d "$LIB_ENUMS_PATH" ]; then rm -r "$LIB_ENUMS_PATH"; fi
if [ -d "$LIB_MODELS_PATH" ]; then rm -r "$LIB_MODELS_PATH"; fi
if [ -f "$LIB_FILE_PATH" ]; then rm "$LIB_FILE_PATH"; fi

if [ -d "${DART_LIB_FOLDER}/.git" ]; then git -C "${DART_LIB_FOLDER}" rm --cached -r . ; fi

go run "./dart" "$DART_LIB_FOLDER"

cd "$DART_LIB_FOLDER"

if [ "$RUN_BUILD" = true ]; then
    flutter pub get
    flutter packages pub run build_runner build --delete-conflicting-outputs
fi
#if [ "$RUN_FORMAT" = true ]; then
#    dartfmt -l 120 -w .
#fi
