package main

import "fmt"

func average(a, b int) float32{
return float32(a+b) /3
}

func main(){

	a , b := 80,75

	if(a>b){
		fmt.Println("quize a")
	}else if(b>a){
		fmt.Println("quiz b")
	}else {
		fmt.Println("Equal")
	}

if average(a,b) > 50 {
	fmt.Println("pass")
}else{
	fmt.Println("fail")
}
}