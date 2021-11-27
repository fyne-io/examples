#!/bin/sh

DIR=`dirname "$0"`
FILE=bundled.go
BIN=`go env GOPATH`/bin

cd $DIR
rm $FILE

$BIN/fyne bundle -package icon -name bugBitmap bug.png > $FILE
$BIN/fyne bundle -package icon -append -name clockBitmap clock.png >> $FILE
$BIN/fyne bundle -package icon -append -name fractalBitmap fractal.png >> $FILE
$BIN/fyne bundle -package icon -append -name textEditorBitmap texteditor.png >> $FILE
$BIN/fyne bundle -package icon -append -name xkcdBitmap xkcd.png >> $FILE
$BIN/fyne bundle -package icon -append -name widgetBitmap widgetsIcon.png >> $FILE

