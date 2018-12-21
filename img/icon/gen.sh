#!/bin/sh

DIR=`dirname "$0"`
FILE=bundled.go
BIN=`go env GOPATH`/bin

cd $DIR
rm $FILE

$BIN/bundler -package icon -name life life.svg > $FILE
$BIN/bundler -package icon -append -name lifeBitmap life.png >> $FILE

$BIN/bundler -package icon -append -name fractalBitmap fractal.png >> $FILE
$BIN/bundler -package icon -append -name bugBitmap bug.png >> $FILE
$BIN/bundler -package icon -append -name xkcdBitmap xkcd.png >> $FILE

