package main 

import (
	"fmt"
)

func main() {
	var name string = "Cephus"
	fmt.Printf("Hello World I'm, %s", name)
	
	energy := []int{45,42,41,40}
	for _, v := range energy{
		fmt.Println("checking Energy Levels -- @ %v", v)
	}

}
var energyLevel int = 5
var energyPointer *int = &energyLevel

var quantumCodes [5]int = [5]int{11,33,44,5,66}






