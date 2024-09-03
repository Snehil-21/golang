package main

import "fmt"

func main() {
	// declare and init a variable

	// var character_name string = "Pikachu"
	var character_name = "Pikachu"

	// decalre variable but assign value later
	var league_name string

	fmt.Println(character_name, league_name)

	league_name = "Diamond"

	fmt.Println(character_name, league_name)

	// short hand variable initialisation

	myName := "Snehil"

	fmt.Println(myName)

	// variable types include : int, unit, float, string, int16, int32 likewise , uint16, uint32 likewise , float32, float64 likewise

	// fmt.Println("My name is", myName, ". I have", character_name, "in the", league_name, "league")

	// %_ - format specifier
	fmt.Printf("My name is %v. I have %v in the %v league.\n", myName, character_name, league_name)

	// Sprintf - saves string as a variable
	myStr := fmt.Sprintf("My name is %v. I have %v in the %v league.", myName, character_name, league_name)
	fmt.Println(myStr)
}
