// print list items include name and price 
// last item , total cost, total item  
package main

import "fmt"

type List struct {
	name  string
	price int
}

func printInfo(info [3]List) {
	cost := 0
	numItem := 0
	for i := 0; i < len(info); i++ {
		item := info[i]
		numItem++
		cost += item.price
		fmt.Println(item.name, " ",item.price)
	}
	
	fmt.Println("Totol Item ", numItem)
	fmt.Println("Total cost ", cost)
	fmt.Println("Last item ", info[numItem - 1] )
}

func main() {

	product := [...]List{
		{name: "car", price: 10000},
		{name: "electronics", price: 1000},
		{name: "textiles", price: 500},
	}

	printInfo(product)
}