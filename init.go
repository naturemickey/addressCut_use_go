package main

import "io/ioutil"
import "strings"
import "fmt"
import "strconv"
import "sort"
import "time"

func init() {
	start := time.Now().UnixNano()

	filedata, _ := ioutil.ReadFile("citybasedata_v3.config")
	lines := strings.Split(string(filedata), "\n")

	for _, line := range lines {
		ss := strings.Split(line, ",")
		if len(ss) <= 2 {
			continue
		}

		id, _ := strconv.ParseInt(ss[0], 10, 64)
		parentId, _ := strconv.ParseInt(ss[1], 10, 64)
		level, _ := strconv.Atoi(ss[2])

		for i := 3; i < len(ss); i++ {
			name := ss[i]

			if len(name) == 0 || name == "null" {
				continue
			}

			ct := &CityToken{id: id, parentId: parentId, name: name, level: level, parent: nil}

			setNameMap(name, ct)
			setIdMap(id, ct)
		}
	}

	for _, value := range idMap {
		sort.Sort(pCityTokeSliceForName(value))
	}

	for key, value := range nameMap {
		sort.Sort(pCityTokeSliceForLevel(value))

		for _, v := range value {
			if cityTokenSlice, ok := idMap[v.parentId]; ok {
				if len(cityTokenSlice) > 0 {
					v.parent = cityTokenSlice[0]
				}
			}
		}

		names = append(names, key)
	}
	fmt.Println("DataCache init cost:", (time.Now().UnixNano()-start)/1000000, "ms")
}

func init() {
	start := time.Now().UnixNano()
	dfa = newDFA(names)
	fmt.Println("DFA init cost:", (time.Now().UnixNano()-start)/1000000, "ms")
}
