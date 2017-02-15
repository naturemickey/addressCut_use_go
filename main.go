// addressCut_use_go project main.go
package main

import (
	"fmt"
)

var o1 = createDFAState()

func main() {
	var o2 = o1
	o3 := o1
	fmt.Printf("%p\n", o1)
	fmt.Printf("%p\n", o2)
	fmt.Printf("%p\n", o3)
}
