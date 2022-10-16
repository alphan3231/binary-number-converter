package main

import (
	"fmt"
	"os"
)

func main() {
	var answer int
	fmt.Println("Do you want to learn binary equivalent of some integer ( 1 for yes, 0 for no).")
	fmt.Scan(&answer)
	if answer == 1 {
		var num int
		fmt.Println("type your number please")
		fmt.Scan(&num)
		fmt.Printf(" binary equivalent is %b", num)
	} else {
		os.Exit(0)
	}

}
