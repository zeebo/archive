#!/bin/bash

found=$(find ~/tmp -type d -depth 2 | grep 'tmp/\.' | grep "$1")

for dir in $found; do
	if [ ! -z "$1" ]; then
		echo resurrecting "$dir"
		touch "$dir"
		mv "$dir" ~/tmp
	else
		echo "$dir"
	fi
done