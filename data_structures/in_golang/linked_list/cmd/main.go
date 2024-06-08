package main

import (
	mypackage "app/my-package"
	"fmt"
)

func main() {

	mypackage.Foo()
	mypackage.Call()

	r := mypackage.Add(23, 34)
	fmt.Println(r)

}
