package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// fmt.Println("The secret number is ", secretNumber)
	value := 0
	fmt.Println("Please input your guess")
	for {
		_, err := fmt.Scanf("%d", &value)
		if err != nil {
			log.Println("invaild input , try again")
		}
		fmt.Println("You guess is", value)
		if value > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if value < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
