// types/structs/methods/begin/main.go
package main

import "fmt"

type author struct {
	first string
	last  string
}

// fullName returns the full name of the author
func (a author) fullName() string {
	return a.first + " " + a.last
}

//

func main() {
	// initialize author
	a := author{
		first: "César",
		last:  "Santos",
	}

	// print the author's full name
	fmt.Println(a.fullName())
}
