package main

import (
	"strings"
)

var dCity = NewStringSet()
var hmtCity = NewStringSet()

type Address struct {
	originalAddress string
	value           *CityToken
	addrReal        string
	addrList        []string
	children        []*Address
	parent          *Address
}

func newAddress(originalAddress string, addrList []string) *Address {
	if addrList == nil {
		addrList = []string{}
	}
	address := &Address{originalAddress: originalAddress, addrList: addrList}
	for _, addr := range addrList {
		address.add1(addr)
	}
	return address
}

func (this *Address) add1(addrStr string) {
	for _, ct := range nameMap[addrStr] {
		ct = idMap[ct.id][0]
		this.add2(ct, addrStr)
	}
}

func (this *Address) add2(ct *CityToken, addrStr string) {
	hasRelationship := false
	for i, addr := range this.children {
		if ct.id == addr.value.id {
			hasRelationship = true
			if addrStr == ct.name {
				addr.value = ct
				addr.addrReal = addrStr
			}
			break
		} else {
			relationship := getRelationship(ct, addr.value)
			if relationship != 0 {
				hasRelationship = true
				if relationship > 0 {
					addrNew := &Address{value: ct, addrReal: addrStr, parent: this, children: []*Address{}}
					this.children[i] = addrNew
					addrNew.children = append(addrNew.children, addr)
				} else {
					addr.add2(ct, addrStr)
				}
				break
			}
		}
	}
	if !hasRelationship {
		this.children = append(this.children, &Address{value: ct, addrReal: addrStr, parent: this, children: []*Address{}})
	}
}

func getRelationship(ct1 *CityToken, ct2 *CityToken) int {
	if ct1.level == ct2.level || ct1.name == ct2.name {
		// 1.同一级别不可能是上下级关系。
		// 2.名字相同的以级别高的为准。
		return 0
	}
	if ct1.level > ct2.level {
		return -1 * getRelationship(ct2, ct1)
	}
	for ct1.level < ct2.level && ct2.parent != nil {
		if ct2.parent.id == ct1.id {
			return 1
		}
		ct2 = ct2.parent
	}
	return 0
}

func (this *Address) breakTree() [][]*Address {
	childrenTmp := []*Address{}
	for _, a := range this.children {
		if a.value != nil && a.value.level < 4 {
			childrenTmp = append(childrenTmp, a)
		}
	}
	this.children = childrenTmp

	res := this.breakTreeRecu()
	for idx, resIt := range res {
		tmp := []*Address{}
		ns := NewStringSet()

		for _, a := range resIt {
			ar := a.addrReal
			if (a.parent != nil && a.parent.value != nil && a.value.level >= 4 && a.parent.value.level == 1 && !dCity.contains(a.parent.value.name)) || !ns.add(ar) {

			} else {
				tmp = append(tmp, a)
			}
		}
		res[idx] = tmp
	}

	return res
}

func (this *Address) breakTreeRecu() [][]*Address {
	res := [][]*Address{}
	if len(this.children) == 0 {
		l := []*Address{}
		if this.value != nil {
			l = append(l, this)
		}
		res = append(res, l)
	} else {
		for _, c := range this.children {
			for _, fl := range c.breakTreeRecu() {
				if this.value != nil {
					fl = append([]*Address{this}, fl...)
				}
				res = append(res, fl)
			}
		}
	}
	return res
}

func (this *Address) choose(ll [][]*Address, idx int) []*Address {
	if len(ll) == 0 {
		return []*Address{}
	}
	if len(ll) == 1 {
		return ll[0]
	}
	res1 := [][]*Address{}
	level := 999999
	for _, l := range ll {
		if len(l) > idx {
			a := l[idx]
			if a.value != nil {
				if level >= 3 && a.value.level < level {
					res1 = [][]*Address{}
					level = a.value.level
					res1 = append(res1, l)
				} else if a.value.level == level || (level <= 2 && a.value.level <= 2) {
					res1 = append(res1, l)
				}
			}
		}
	}

	if len(res1) == 1 {
		return res1[0]
	}

	isStd := false
	toRecu := false
	res2 := [][]*Address{}
	for _, l := range res1 {
		a := l[idx]
		if a.value.level < 2 {
			res2 = res1
			for _, l2 := range res1 {
				toRecu = toRecu || len(l2) > idx+1
			}
			break
		}
		if isStd {
			if len(a.value.name) == len(a.addrReal) {
				res2 = append(res2, l)
				toRecu = toRecu || len(l) > idx+1
			}
		} else {
			if len(a.value.name) == len(a.addrReal) {
				res2 = [][]*Address{}
				toRecu = false
				isStd = true
			}
			res2 = append(res2, l)
			toRecu = toRecu || len(l) > idx+1
		}
	}
	if toRecu {
		return this.choose(res2, idx+1)
	}
	if len(res2) == 0 {
		return []*Address{}
	}
	if len(res2) == 1 {
		return res2[0]
	}
	var res []*Address = nil
	for _, l := range res2 {
		if res == nil {
			res = l
		} else {
			a1 := res[idx]
			a2 := l[idx]
			if indexOf(this.addrList, a2.addrReal) < indexOf(this.addrList, a1.addrReal) {
				res = l
			}
		}
	}
	return res
}

func indexOf(l []string, s string) int {
	for i, x := range l {
		if x == s {
			return i
		}
	}
	return -1
}

func (this *Address) getCutRes() ([]*CityToken, string) {
	return this.fixToToken(this.choose(this.breakTree(), 0))
}

func (this *Address) fixToToken(addrList []*Address) ([]*CityToken, string) {
	detailAddress := this.originalAddress
	if len(addrList) == 0 {
		return []*CityToken{}, detailAddress
	}
	addr := addrList[len(addrList)-1]
	ct := idMap[addr.value.id][0]
	lastRealAddr := addr.addrReal
	lastStandardAddr := ct.name
	ctList := []*CityToken{}
	for ct != nil {
		ctList = append([]*CityToken{ct}, ctList...)
		ct = ct.parent
	}
	if lastRealAddr != "" || lastStandardAddr != "" {
		detailAddress = this.subOrigAddr(lastRealAddr, lastStandardAddr)
	}
	return ctList, detailAddress
}
func (this *Address) subOrigAddr(addr string, stdAddr string) string {
	idx := -1
	if stdAddr != "" {
		idx = strings.LastIndex(this.originalAddress, stdAddr)
	}
	if idx < 0 && addr != "" {
		idx = strings.LastIndex(this.originalAddress, addr)
	}
	if idx > 0 {
		buf := []byte(this.originalAddress)
		length := len(buf)
		buf2 := []byte{}
		for i := idx; i < length; i++ {
			buf2 = append(buf2, buf[i])
		}
		return string(buf2)
	}
	return this.originalAddress
}

func (this *Address) info() *Info {
	res := &Info{}
	ctList, detailAddress := this.getCutRes()
	res.detailAddress = detailAddress
	if ctList != nil {
		for _, ct := range ctList {
			switch ct.level {
			case 1:
				res.provinceAddress = ct.name
			case 2:
				res.cityAddress = ct.name
			case 3:
				res.areaAddress = ct.name
			case 4:
				res.townAddress = ct.name
			}
		}
		if len(ctList) == 1 {
			if dCity.contains(res.provinceAddress) {
				res.cityAddress = res.provinceAddress + "市"
			} else if hmtCity.contains(res.provinceAddress) {
				res.cityAddress = res.provinceAddress
			}
		}
	}
	res.originalAddress = this.originalAddress
	return res
}
