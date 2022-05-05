package main

import (
	"fmt"
	"time"
)

func main() {

	language := "Rust"

	switch language {
	case "Java":
		fmt.Println("You are learning Java! You have to do semi colon for every line.")

	case "Python":
		fmt.Println("Wow, they don't have semi colons at all!")

	case "golang":
		fmt.Println("Good, Go will get you a job at YouTube!")

	default:
		fmt.Println("What kind of languages are you using?")
	}

	//do a switch statement where switch statement is for bool (true)
	//cases would be number modulos (%) to even or odd number, and then

	value := 57

	switch {
	case value%2 == 0:
		fmt.Println("It's even number ")
	case value%2 != 0:
		fmt.Println("It's not an even number")
	}
	//create a variable with name hour and do a switch statement and print out a string to say
	//good morning if it's not yet noon and good afternoon if before 6pm and good evening if it's night time
	hour := time.Now().Hour()

	switch {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening") //prints good evening
	}
}
