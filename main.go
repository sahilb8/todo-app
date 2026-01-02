package main

import (
	"fmt"
	"strconv"
)

type hello struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello, World!")

	var name string = "Sahil"
	var age = 19
	isGuitarist := true

	fmt.Println(name, age, isGuitarist)
	fmt.Println(sayHi(name))
	fmt.Println(isPersonGuitarist(isGuitarist))

	if practiceMonths, guitarChord := 10, "fe3f"; isGuitarist {
		fmt.Println("Play the barrrr " + guitarChord + " chordsssss!!!!")
		fmt.Println("Practiced for: " + strconv.Itoa(practiceMonths))
	} else {
		fmt.Println("go learnnnnnnn musiccccc!!!!")
	}

	switch age {
	case 19:
		fmt.Println("You are 19 years old")
		fallthrough
	case 20:
		fmt.Println("You are 20 years old")
	default:
		fmt.Println("You are not 19 or 20 years old")
	}

	items := []string{"pen", "pencil", "notebook"}
	for index, value := range items {
		fmt.Println(index, value)
	}

	mySlice := make([]bool, 4)
	fmt.Printf("Len: %d, Cap: %d, Pointer: %p\n", len(mySlice), cap(mySlice), mySlice)
	mySlice = append(mySlice, true)
	fmt.Printf("Len: %d, Cap: %d, Pointer: %p\n", len(mySlice), cap(mySlice), mySlice)

	myMap := map[string]int{
		"pen":      1,
		"pencil":   0,
		"notebook": 1,
	}
	value, ok := myMap["fwef"]
	fmt.Println(value, ok)

	ubuu := hello{
		name: "Sahil",
		age:  19,
	}
	fmt.Println(ubuu.age)
	fmt.Println(ubuu.name)

	ubbuAdress := &ubuu
	explorePointer(*ubbuAdress)
	fmt.Println(*ubbuAdress)

}

func (data hello) greetPeople() string {
	return "Hello " + data.name
}

func explorePointer(ptr hello) {
	fmt.Println("insideeeeee")
	fmt.Println(ptr.name)
	fmt.Println(ptr.age)
	fmt.Println(ptr.greetPeople())
	fmt.Printf("%p\n", ptr)
}

func sayHi(name string) string {
	return "Hi " + name
}

func isPersonGuitarist(isGuitarist bool) (string, error) {
	if isGuitarist {
		return "Guitarist", nil
	}
	return "Not a Guitarist", nil
}
