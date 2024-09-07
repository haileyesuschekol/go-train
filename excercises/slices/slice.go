//create assambly line that contains type part
package main

import "fmt"

type part string

func shoeLine(line []part){
for i:=0; i<len(line);i++{
	part:= line[i]
	fmt.Println(part)
}
}

func main(){
assamblyLibe := make([]part,3)
assamblyLibe[0] = "pipe"
assamblyLibe[1] = "bolt"
assamblyLibe[2] = "sheet"

fmt.Println("3 parts:")
shoeLine(assamblyLibe)

assamblyLibe = append(assamblyLibe, "washer", "wheel")
fmt.Println("\nadded two parts")
shoeLine(assamblyLibe)

assamblyLibe = assamblyLibe[3:]
fmt.Println("\n sliced")
shoeLine(assamblyLibe)
}