package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	
  //Question 1
	fmt.Println("\nQUESTION 1")
	var name = "Hendra Huang"
	country := "United States"
	fmt.Printf("Your name: %s\nCountry: %s\n", name, country)

	fmt.Println("He says: \"Hello\"")
	fmt.Println(`C:\Users\a.txt`)
	fmt.Println()

	//Question 2
	fmt.Println("Question 2")
	r := []rune("ă")
	fmt.Printf("%T\n", r)
	strOne, strTwo := "ma", "m"
	word := strOne + strTwo + string(r)

	fmt.Println("The combined word for exercise 2 is:", word)

	//Question 3
	fmt.Println("Question 3")
	s1 := "țară means country in Romanian"

	fmt.Println("Byte by byte:")
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v ", s1[i])
	}

	fmt.Println()

	for i, r := range s1 {
		fmt.Printf("%v: rune [%c]\n", i, r)
	}

	//Exercise 4
	fmt.Println("Exercise 4")
	fact := "Go is cool!"
	x := fact[0]
	fmt.Println(x)

	fact = strings.Replace(fact, "G", "X", -1)
	fmt.Println(fact)

	// printing the number of runes in the string
	fmt.Println(len(fact))
	fmt.Println(utf8.RuneCountInString(fact))

	//Question 5
	s := "你好 Go!"
	chinese := []byte(s)

	fmt.Printf("%#v\n", chinese)

	for i, val := range chinese {
		fmt.Printf(" %d => byte value [%v]\n", i, val)
	}

	//Question 6
	sOne := "你好 Go!"
	eachChar := []rune(sOne)

	fmt.Printf("%#v ", eachChar)

	fmt.Println()
	for i, r := range eachChar {
		fmt.Printf("%d: %c\n", i, r)
  }
}
