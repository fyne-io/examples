#!/bin/sh

file="bundled.go"
cmd=`go env GOPATH`/bin/bundler

rm $file
for suit in D H C S; do
	for card in A 2 3 4 5 6 7 8 9 10 J Q K; do
		append="--append"
		if [[ $suit = "D" ]] && [[ $card = "A" ]] ; then
			append=""
		fi

		$cmd $append --name=card$card$suit --package=faces $card$suit.svg >> $file
	done
done

$cmd --append --name=back back.svg >> $file
$cmd --append --name=space space.svg >> $file
