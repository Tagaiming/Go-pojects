package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("*** gussing game****")
	fmt.Println("You got only 10 attemp")
	fmt.Println("Enter your number 1-10")
	var User int

	robot := rand.IntN(10)+1
	for i := 0; i < 10; i++ {

		fmt.Scanln(&User)

		if User == robot {
			fmt.Println("User won the game")
			break

		} else if User > robot {
			fmt.Println("You guess is bigger than number")

		} else if User < robot {
			fmt.Println("Your guess is lower than number")

		} else {

			fmt.Println("You lose the game ",robot)
		}

	}

}
