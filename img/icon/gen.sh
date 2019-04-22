#!/bin/sh

DIR=`dirname "$0"`
FILE=bundled.go
BIN=`go env GOPATH`/bin

cd $DIR
rm $FILE

$BIN/fyne bundle -package icon -name life life.svg > $FILE
$BIN/fyne bundle -package icon -append -name lifeBitmap life.png >> $FILE

$BIN/fyne bundle -package icon -append -name bugBitmap bug.png >> $FILE
$BIN/fyne bundle -package icon -append -name calculatorBitmap calculator.png >> $FILE
$BIN/fyne bundle -package icon -append -name clockBitmap clock.png >> $FILE
$BIN/fyne bundle -package icon -append -name fractalBitmap fractal.png >> $FILE
$BIN/fyne bundle -package icon -append -name solitaireBitmap solitaire.png >> $FILE
$BIN/fyne bundle -package icon -append -name textEditorBitmap texteditor.png >> $FILE
$BIN/fyne bundle -package icon -append -name xkcdBitmap xkcd.png >> $FILE

