package main

import (
	pc "curso_golang_platzi/src/mypackage"
	"fmt"
)

func main() {

	var myPC pc.MyPCPublic
	myPC.Brand = "aaa"
	myPC.Ram = 16
	myPC.Disk = 512
	fmt.Println(myPC)
	myPC.SeePC()
}
