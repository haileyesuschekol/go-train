package main

import "fmt"

func printSlice(title string, slice []string){
	fmt.Println("---", title, "---")
	for i:=0; i<len(slice); i++{
		element:= slice[i]
		fmt.Println(element)
	}
}

func main(){
	 route := []string{"Grosery", "Department store", "salon" }
	printSlice("Route 1", route)

	route = append(route, "home" )
	printSlice("Route 2", route)
 
	route = route[2:]

	printSlice("Route 3",route)


	}