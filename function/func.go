package main

import "fmt"

func double(nums int) int{
return nums * 2
}

func greet(){
	fmt.Println("Hellooo")
}

func multiReturn() (int, string, int){
return 1, "value", 2
}

  func main(){
	doubleNum := double(3)
	fmt.Println(doubleNum)

	greet()

	a, b, _ := multiReturn()
	fmt.Println(a, b)
  }