// create a program to manage lending of library books
package main

import (
	"fmt"
	"time"
)



type Title string
type Name string

type LendAudit struct{
	checkOut time.Time
	checkIn time.Time
}

type Member struct{
	name Name
	book map[Title]LendAudit
}

type BookEntry struct{
	total int 
	lended int
}

type Library struct{
	member map[Name]Member
	books map[Title]BookEntry
}

func printMemberAudit(member *Member){
for title, audit :=range member.book{
	var returnTime string
	if audit.checkIn.IsZero(){
		returnTime = "[not return yrt]"
	}else{
		returnTime = audit.checkIn.String()
	}
	fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
}
}

func printMemberAudits(library *Library){
	for _, member := range library.member{
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library){
	fmt.Println()
	for title, book := range library.books{
		fmt.Println(title, "/ total", book.total, "/ lended", book.lended)
	}
}

func checkoutBook(library *Library, title Title, member *Member) bool{
	book, found :=library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	if book.lended == book.total{
		fmt.Println(("No more books available to lend"))
		return false
	}
book.lended +=1
library.books[title] =book
member.book[title] = LendAudit{checkOut: time.Now()}
return true
}

func returnBook(library *Library, title Title, member *Member) bool{
book, found := library.books[title]
if !found{
	fmt.Println("Book not part of library")
}

audit, found := member.book[title]
if !found{
	fmt.Println("Member did not check out this book")
	return false
}

book.lended -= 1
library.books[title] =book
audit.checkIn = time.Now()
member.book[title] =audit
return true
}


func main(){
	library := Library{
		books: make(map[Title]BookEntry),
		member: make(map[Name]Member),
	}

	library.books["Webapps i Go"] = BookEntry{
		total: 4,
		lended: 0,
	}
	library.books["The Little GO Book"] = BookEntry{
		total: 3,
		lended: 0,
	}
	library.books["Let's learn Go"] = BookEntry{
		total: 2,
		lended: 0,
	}
	library.books["Go Bootcamp"] = BookEntry{
		total: 1,
		lended: 0,
	}

	library.member["Jayson"] = Member{"Jayson", make(map[Title]LendAudit)}
	library.member["Billy"] = Member{"Billy", make(map[Title]LendAudit)}
	library.member["Sussan"] = Member{"Sussan", make(map[Title]LendAudit)}

	fmt.Println("\nInitial:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	member := library.member["Jayson"]
	checkedOut := checkoutBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck out a book:")
	if checkedOut{
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

	returned := returnBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck out a book:")
	if returned{
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

}