package main

type Address struct {
	originalAddress string
	value           *CityToken
	addrReal        string
	addrList        []string
	children        []*Address
	parent          *Address
}

func newAddress2(ct *CityToken, addrStr string, parent *Address) *Address {
	return &Address{value: ct, addrReal: addrStr, parent: parent, children: []*Address{}}
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
	for {
		if ct1.level < ct2.level && ct2.parent != nil {
			if ct2.parent.id == ct1.id {
				return 1
			}
			ct2 = ct2.parent
		}
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

	// TODO
	res := [][]*Address{}

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
					fl = append([]*Address{this}, this)
				}
				res = append(res, fl)
			}
		}
	}
	return res
}
