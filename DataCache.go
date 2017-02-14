package main

import "io/ioutil"
import "strings"
import "fmt"
import "strconv"
import "sort"
import "time"

var nameMap = make(map[string][]*CityToken)
var idMap = make(map[int64][]*CityToken)

func init() {
	start := time.Now().Unix()

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

	for _, value := range nameMap {
		sort.Sort(pCityTokeSliceForLevel(value))

		for _, v := range value {
			if cityTokenSlice, ok := idMap[v.parentId]; ok {
				if len(cityTokenSlice) > 0 {
					v.parent = cityTokenSlice[0]
				}
			}
		}
	}

	fmt.Println(time.Now().Unix() - start)
}

// CityToken按name排序
type pCityTokeSliceForName []*CityToken

func (a pCityTokeSliceForName) Len() int {
	return len(a)
}
func (a pCityTokeSliceForName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a pCityTokeSliceForName) Less(i, j int) bool {
	return len(a[j].name) < len(a[i].name)
}

// CityToken按level排序
type pCityTokeSliceForLevel []*CityToken

func (a pCityTokeSliceForLevel) Len() int {
	return len(a)
}
func (a pCityTokeSliceForLevel) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a pCityTokeSliceForLevel) Less(i, j int) bool {
	return a[j].level < a[i].level
}

func setIdMap(key int64, ct *CityToken) {
	if v, ok := idMap[key]; ok {
		idMap[key] = append(v, ct)
	} else {
		idMap[key] = []*CityToken{ct}
	}
}
func setNameMap(key string, ct *CityToken) {
	if v, ok := nameMap[key]; ok {
		nameMap[key] = append(v, ct)
	} else {
		nameMap[key] = []*CityToken{ct}
	}
}
