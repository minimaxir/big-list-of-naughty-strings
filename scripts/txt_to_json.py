###	Quick Python Script to convert the Big List of Naughty Strings into a JSON file
### 
###	By Max Woolf

import json

with open('../blns.txt', 'r') as f:

    # put all lines in the file into a Python list
    content = f.readlines()
    
    # above line leaves trailing newline characters; strip them out
    content = [x.strip('\n') for x in content]
    
    # remove empty-lines and comments
    content = [x for x in content if x and not x.startswith('#')]
    
    # insert empty string since all are being removed
    content.insert(0, "")
    
    # special case: convert "\" to "\\" for valid JSON
    #content = map(lambda x: x.replace('\','\\'), content)
    
with open('../blns.json', 'wb') as f:

	# write JSON to file; note the ensure_ascii parameter
	json.dump(content, f, indent=2, ensure_ascii=False)
    
