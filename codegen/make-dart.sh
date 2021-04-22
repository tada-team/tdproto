#!/bin/bash
set -euo pipefail
shopt -s extdebug
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

rm -r "${DART_LIB_FOLDER}/lib/src/models"
rm -r "${DART_LIB_FOLDER}/lib/src/enums"

go run "./dart" "$DART_LIB_FOLDER"

cd "$DART_LIB_FOLDER"

if [ "$RUN_BUILD" = true ]; then
    flutter pub get
    flutter packages pub run build_runner build --delete-conflicting-outputs
fi
if [ "$RUN_FORMAT" = true ]; then
    dartfmt -l 120 -w .
fi
