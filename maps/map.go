package main

import "fmt"

func main(){
	 items := make(map[string] int)	 
	 items["milk"] = 1
	 items["bread"] = 1
	 items["eggs"] = 12

	 fmt.Println(items)
	
	delete(items, "milk")
	fmt.Println("milk deleted, new list",items)

	cereal, found := items["cereal"]
	fmt.Println("need cereal?")
	if !found{
		fmt.Println("nope")
	}else {
		fmt.Println(cereal,"found")
	}

	totalItem := 0

	for item, val:= range items{
		fmt.Println(val, item )
		totalItem +=val
	}

	fmt.Println("total", totalItem)

}