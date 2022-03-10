//This program aims to calculate interest rate, interest, annual balance and years it would take to earn certain amount of money in a bank account
//It also demonstrates how to use recursion on the yearsBeforeDesiredBalance() function.

package main

import "fmt"

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interest float32

	if balance < 0 {
		interest = 3.213
	} else if balance < 1000 {
		interest = 0.5
	} else if balance >= 1000 && balance < 5000 {
		interest = 1.621
	} else {
		interest = 2.475
	}
	return interest
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)/100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
// Notice the RECURSION in this function.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {

	if balance > targetBalance {
		return 0
	} else {
		return 1 + YearsBeforeDesiredBalance(AnnualBalanceUpdate(balance), targetBalance)
	}

}

func main() {
	yearsToDesiredBal := YearsBeforeDesiredBalance(2345.67, 12345.6789)
	fmt.Println("Years needed to reach balance goal is", yearsToDesiredBal)
}
