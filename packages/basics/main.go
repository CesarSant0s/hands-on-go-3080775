// packages/basics/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world!")

	fmt.Printf("Today is %s", time.Now().Weekday())
}
