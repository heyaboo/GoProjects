package main

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}
// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}
// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    
	count,ok := units[unit]

	if !ok {
		return false
	}
	bill[item] += count
	return true
}
// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    
	unitExists, itemExists := units[unit], bill[item]
	qty := bill[item] //this will get the quantity of searched item
	remaining := qty - unitExists
	isRemoved := false
	if itemExists < 1 || unitExists < 1 || remaining < 0 {
	} else if remaining == 0 {
		delete(bill, item)
		isRemoved = true
	} else {
		bill[item] = remaining
		isRemoved = true
	}
	return isRemoved
}
// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	
    i, ok := bill[item] //if item isn't on the bill
	return i, ok
}

func main() {

	units := Units()
	bill := NewBill()
	ok := AddItem(bill, units, "carrot", "hendra")
	fmt.Println(ok)

	basket := map[string]int{
		"apple":    5,
		"orange":   6,
		"cucumber": 3,
	}

	RemoveItem(basket, Units(), "apple", "half_of_a_dozen")

	fmt.Println(basket)

}
