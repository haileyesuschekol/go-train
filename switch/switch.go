package main

import (
	"fmt"
)

func main(){

	const (
		economy = 0
		buisness = 1
		firstClass = 2
	)

	ticket := buisness
	switch ticket{
	case economy:
		fmt.Println("Economy seat")
	case buisness:
		fmt.Println("Buisness seat")	
	case firstClass: 
		fmt.Println("First class sear")
	default:
		fmt.Println("other")	
	}
	
}