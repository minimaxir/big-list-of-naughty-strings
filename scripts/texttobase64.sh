commentChar="#"
while read line
do
    firstChar="${line:0:1}"
    if [[ "$firstChar" != "$commentChar" ]] && [[ -n "$firstChar" ]]
    then
		echo -n $line | base64
	else
		echo $line
	fi
done <blns.txt
