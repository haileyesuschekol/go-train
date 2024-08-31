package main

import "fmt"

func main(){
	var myName = "jack"
	fmt.Println("my name is", myName, myName)

	var name string = "kathy"
	fmt.Println("name = ", name)

	userName :=  "admin"
	fmt.Println("admin Name = ", userName)

	var sum int 

	var (
		num1 = 2
		num2 = 3
	)

	sum = num1 + num2
	fmt.Println(num1 ,"+", num2, "=", sum)

	word1, word2, _ := "Hello","world" ,"!"
	fmt.Println(word1, word2)

}
