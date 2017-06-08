## Encoding 

Go as of today (June 2017) does not yet have a yaml parser in its standard library. The most common library used for this purpose is: https://github.com/go-yaml/yaml 

### YAML vs JSON

YAML should be preferred whenever possible for the following reasons: 

1. YAML is a superset of JSON. Therefore, YAML parsers should be able to "understand" JSON format as well
2. YAML can use anchors 
3. YAML allows writing JSON in YAML files
4. Many other small [features](https://learnxinyminutes.com/docs/yaml/)

**Note**: YAML does not allow `tab` 

### Usage of yaml 

```go 
package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func main() {
	var m map[string]interface{}
	var err error
	var data string

	data = `name: yerken
details:
  school: NTU
  major: CS
age: 25`
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		panic(err)
	}
	fmt.Println(m) //map[name:yerken age:25 details:map[major:CS school:NTU]]
}
```

### Why not use only YAML format ? 
1. Many libraries define only JSON serialization tags for its internal structs
2. JSON is still a lot more popular and the only acceptable format for many APIs