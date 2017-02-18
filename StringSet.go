package main

var o = 0

type StringSet struct {
	innerMap map[string]int
}

func NewStringSet() *StringSet {
	return &StringSet{innerMap: map[string]int{}}
}

func (this *StringSet) add(s string) bool {
	if _, ok := this.innerMap[s]; ok {
		return false
	}
	this.innerMap[s] = o
	return true
}

func (this *StringSet) contains(s string) bool {
	_, ok := this.innerMap[s]
	return ok
}
