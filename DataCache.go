package main

var nameMap = map[string][]*CityToken{}
var idMap = map[int64][]*CityToken{}
var names = []string{}

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
