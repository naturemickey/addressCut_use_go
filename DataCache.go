package main

import "io/ioutil"
import "strings"
import "fmt"
import "strconv"

var nameMap = make(map[string][]*CityToken)
var idMap = make(map[int64][]*CityToken)

func init() {
	filedata, _ := ioutil.ReadFile("citybasedata_v3.config")
	lines := strings.Split(string(filedata), "\n")
	for _, line := range lines {
		ss := strings.Split(line, ",")
		if len(ss) <= 2 {
			continue
		}
		//fmt.Println(ss)
		id, _ := strconv.ParseInt(ss[0], 10, 64)
		parentId, _ := strconv.ParseInt(ss[1], 10, 64)
		level, _ := strconv.Atoi(ss[2])

		for i := 3; i < len(ss); i++ {
			name := ss[i]

			if len(name) == 0 || name == "null" {
				continue
			}

			ct := &CityToken{id: id, parentId: parentId, name: name, level: level, parent: nil}

			fmt.Println(ct)
		}
	}
}
