commentChar="#"
while read p; do 
        firstChar=${p:0:1}
        if [[ "$firstChar" != "$commentChar" && "$firstChar" != "" ]] ; then
		echo $p | base64; 
	else
		echo $p;
	fi
done <blns.txt
