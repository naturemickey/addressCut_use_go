package main

type CityToken struct {
	id       int64
	name     string
	parentId int64
	level    int
	parent   *CityToken
}
