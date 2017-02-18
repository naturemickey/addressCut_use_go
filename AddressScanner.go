package main

func scan(txt string) *Address {
	tmp := []rune{}
	for _, c := range txt {
		if c != ' ' && c != '\t' && c != '\n' && c != '\r' {
			tmp = append(tmp, c)
		}
	}
	txt = string(tmp)
	addrList := dfa.scan(txt)

	if len(addrList) > 0 {
		fc := []rune(addrList[0])
		if len(fc) == 3 {
			c := fc[0]
			if (c == '北' || c == '上' || c == '天' || c == '重') && fc[2] == '市' {
				c = fc[1]
				if c == '京' {
					addrList = append([]string{"北京"}, addrList...)
				} else if c == '海' {
					addrList = append([]string{"上海"}, addrList...)
				} else if c == '津' {
					addrList = append([]string{"天津"}, addrList...)
				} else if c == '庆' {
					addrList = append([]string{"重庆"}, addrList...)
				}
			}
		}
	}

	return newAddress(txt, addrList)
}
