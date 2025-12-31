package main

import "fmt"

func main() {
	var name string
	var num int
	fmt.Println("Enter your name")
	fmt.Scanln(&name)
	fmt.Println("Enter your Number")
	fmt.Scanln(&num)

	if num <= 100 && num >= 80 {
		println("You got A+", name)

	} else if num <= 79 && num >= 70 {
		println("You got A", name)

	} else if num <= 69 && num >= 60 {
		println("you got A-", name)

	} else if num <= 59 && num >= 50 {
		println("You got B",name)

	} else if num <= 49 && num >= 40 {
		println("You got C", name)

	} else if num <= 39 && num >= 33 {
		println("You got D")

	} else if num <= 32 && num >= 00 {
		println("You fail exam",name)
	} else {
		println("Enter your valid number",name)
	}

}
