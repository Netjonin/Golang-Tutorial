package main

import "fmt"

//declared separately
var firstName string
var middleName, lastName string
var age int

//declared as a block of code
var (
	mySon string
	hisAge int
)

//when initialize a variable you don't need to specify the type
//it would be implicitly infered

var (
	myLove = "Ajike"
	myLang = "Go"
)

//intializing them on the same line
var (
	myName, myAge, myHeight = "Taofeek", 37, 178
)

//Declaring constant in go
//Done in mixed case or upper case
const (
    StatusOK              = 0
    StatusConnectionReset = 1
    StatusOtherError      = 2
)

func main() {
	//short declaration and initialization
	myFavLang := "Golang"
	fmt.Println(myFavLang)
}