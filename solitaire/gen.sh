#!/bin/sh

file="faces/bundled.go"
cmd=`go env GOPATH`/bin/fyne

$cmd bundle --package=faces --prefix=card faces > $file

