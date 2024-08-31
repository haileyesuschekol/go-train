package main

import "fmt"

func double(nums int) int{
return nums * 2
}

func greet(){
	fmt.Println("Hellooo")
}

  func main(){
	doubleNum := double(3)
	fmt.Println(doubleNum)

	greet()
  }