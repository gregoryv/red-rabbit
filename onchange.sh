#!/bin/sh
#
# Jarvis runs this script each time a file changes within the project
#
clear
path=$1
filename=$(basename "$path")
extension="${filename##*.}"

case $extension in
  go)
    go test github.com/gregoryv/red-rabbit/cursor
    if [ $? -eq 0 ]; then
      gofmt -w $path
    fi

    ;;
  *)
    ;;
esac

