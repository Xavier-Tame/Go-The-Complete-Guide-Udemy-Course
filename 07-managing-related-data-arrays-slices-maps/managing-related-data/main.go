package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"
	userNames[1] = "Peter"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Charlie")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.6

	courseRatings.output()

	// fmt.Println(courseRatings)

	for i, val := range userNames {
		fmt.Println("Index: ", i)
		fmt.Println("Value: ", val)
	}

	for key, value := range courseRatings {
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)
	}
}
