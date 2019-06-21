#!/bin/sh

DIR=`dirname "$0"`
FILE=bundled.go
BIN=`go env GOPATH`/bin

cd $DIR
rm $FILE

$BIN/fyne bundle -package bugs -name codeIcon code.svg > $FILE
$BIN/fyne bundle -package bugs -name bugIcon -append bug.svg >> $FILE
$BIN/fyne bundle -package bugs -name flagIcon -append flag.svg >> $FILE

