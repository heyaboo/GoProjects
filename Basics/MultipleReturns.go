//This go file demonstrates how to do multiple return statement and also how to make and call a function to main function

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calcSquareArea(squareLen int) (int, int) { //when returning multiple variables, identify the return data type separated by a (,)
	rand.Seed(time.Now().Unix())
	luckyNum := rand.Intn(20)
	if squareLen <= 0 {
		fmt.Println("\nLength of square can't be 0.")
		return 0, luckyNum
	} else {
		return (squareLen * squareLen), luckyNum
	}
}

func main() {

	var squareLen, areaOfSquare, luckyNum int
	fmt.Print("What is the length of this square?")
	fmt.Scan(&squareLen)

	areaOfSquare, luckyNum = calcSquareArea(squareLen) //Since we are returning 2 variables, then we would have to also provide 2 variables for assignment.

	if areaOfSquare > 0 {
		fmt.Println("Area:", areaOfSquare, "square meter, and lucky number is", luckyNum)
	} else {
		fmt.Println("Please run and enter the length again.")
	}

}
