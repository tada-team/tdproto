#!/bin/sh
set -e
cd $( dirname $0 )

dest=../../tdproto_dart/lib/src
go run dart/main.go $dest

cd $dest
dartfmt --line-length 120 -w .
flutter pub get
flutter packages pub run build_runner build --delete-conflicting-outputs

