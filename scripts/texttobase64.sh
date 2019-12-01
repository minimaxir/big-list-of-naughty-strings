#!/usr/bin/env bash

set -o errexit -o nounset -o pipefail

commentChar="#"
while read -r line
do
    firstChar="${line:0:1}"
    if [[ "$firstChar" != "$commentChar" ]] && [[ -n "$firstChar" ]]
    then
		echo -n "$line" | base64
	else
		echo "$line"
	fi
done <blns.txt
