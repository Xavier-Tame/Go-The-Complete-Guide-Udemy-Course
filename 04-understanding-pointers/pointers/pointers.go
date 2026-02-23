package main

import "fmt"

func main() {
	age := 32
	var agePointer *int = &age

	// fmt.Println("Age Address: ", agePointer)
	fmt.Println("Age: ", *agePointer)

	// fmt.Println("Age: ", age)

	// adultYears := getAdultYears(agePointer)
	editAgeToAdultYears(agePointer)
	fmt.Println("Adult Years: ", age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18

	*age = *age - 18
}
