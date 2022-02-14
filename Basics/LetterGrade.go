//This program calculates the final letter grade and prints onto the terminal after taking in user's input of all their scores
//It demonstrates how to use if..else if statements of Golang

/* REQUIREMENTS FOR THIS PROJECT
 so there are 3 exams, add them up and add it to a variable called "totalScore".
 create conditions using if and else if based on the totalScore earned. The score would be using letter grade
 print out the result
*/

package main

import "fmt"

func main() {

	var examOne, examTwo, examThree int
	var finalResult int
	fmt.Println("Please enter the first exam score you obtained: ")
	fmt.Scan(&examOne)
	fmt.Println("Now, second exam?")
	fmt.Scan(&examTwo)
	fmt.Println("Ok, enter your final exam grade: ")
	fmt.Scan(&examThree)

	if finalResult = (examOne + examTwo + examThree) / 3; finalResult >= 90 {
		fmt.Println("Your final letter grade is A, well done star student!")
	} else if finalResult >= 80 && finalResult <= 89 {
		fmt.Println("Your final letter grade is B, good job.")
	} else if finalResult >= 70 && finalResult <= 79 {
		fmt.Println("Your final letter grade is C, that was close!")
	} else {
		fmt.Println("Your final letter grade is D or F, let's get you the assistance you need!")
	}
}
