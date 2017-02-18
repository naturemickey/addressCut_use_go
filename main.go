// addressCut_use_go project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	tm1 := time.Now().UnixNano()
	tm2 := time.Now().UnixNano()
	fmt.Println("construct scanner use ", (tm2-tm1)/1000000, " ms.")
	address := scan("江西抚州市南昌大学抚州医学分院12级全科2班").info()
	tm3 := time.Now().UnixNano()
	fmt.Println("scan an address use ", (tm3-tm2)/1000000, " ms.")

	fmt.Println("province_address : ", address.provinceAddress)
	fmt.Println("city_address     : ", address.cityAddress)
	fmt.Println("area_address     : ", address.areaAddress)
	fmt.Println("town_address     : ", address.townAddress)
	fmt.Println("original_address : ", address.originalAddress)
	fmt.Println("detail_address   : ", address.detailAddress)

	tm4 := time.Now().UnixNano()

	filedata, _ := ioutil.ReadFile("测试地址.txt")
	lines := strings.Split(string(filedata), "\n")

	tm5 := time.Now().UnixNano()
	fmt.Println("读取所有测试地址消耗时长 ", (tm5-tm4)/1000000, " ms.")

	for _, line := range lines {
		scan(line).info()
	}
	tm6 := time.Now().UnixNano()
	fmt.Println("识别所有测试地址消耗时长 ", (tm6-tm5)/1000000, " ms.")

}
