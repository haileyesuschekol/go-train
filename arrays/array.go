package main

import "fmt"

type Room struct{
	name string
	cleaned bool
}

func checkClean(rooms [3]Room){
	for i:=0; i< len(rooms); i++{
		room := rooms[i]
		if room.cleaned{
			fmt.Println(room.name, " is clean")
		}else{
			fmt.Println(room.name, " is dirty")
		}
	}
}

func main(){
	rooms := [...]Room{
		{name: "office"},
		{name: "warehouse"},
		{name: "reception"},
	}

	checkClean(rooms)
	fmt.Println("performing cleaning...")

	rooms[0].cleaned = true
	rooms[2].cleaned =true

	checkClean(rooms)
	
}
