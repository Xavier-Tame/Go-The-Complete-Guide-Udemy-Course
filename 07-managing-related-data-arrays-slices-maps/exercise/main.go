package main

import (
	"fmt"

	"example.com/exercise/product"
)

func main() {
	hobbies := [3]string{"Robotics", "Cooking", "Gaming"}

	fmt.Printf("My hobbies are: %v\n", hobbies)
	fmt.Printf("First element: %s\n", hobbies[0])

	hobbiesEnd := hobbies[1:]
	fmt.Printf("Last two hobbies are: %v\n", hobbiesEnd)

	hobbiesSliceBeginning := hobbies[:2]
	fmt.Printf("First two hobbies are: %v\n", hobbiesSliceBeginning)
	hobbiesSliceBeginningTwo := hobbies[0:2]
	fmt.Printf("First two hobbies (method 2) are: %v\n", hobbiesSliceBeginningTwo)

	hobbiesSliceBeginning = hobbies[:]
	fmt.Printf("All hobbies are: %v\n", hobbiesSliceBeginning)

	courseGoalsDynamic := []string{"Build Projects", "Learn Golang"}
	fmt.Printf("Course goals: %v\n", courseGoalsDynamic)
	courseGoalsDynamic[1] = "Master Golang"
	fmt.Printf("Updated course goals: %v\n", courseGoalsDynamic)
	courseGoalsDynamic = append(courseGoalsDynamic, "Get a Job")
	fmt.Printf("Final course goals: %v\n", courseGoalsDynamic)

	_ = product.AddProduct("Laptop", 1, 999.99)
	_ = product.AddProduct("Smartphone", 2, 499.49)

	products := product.GetProducts()
	fmt.Printf("Products: %v\n", products)

	products = append(products, product.Product{
		Title: "Tablet",
		Id:    3,
		Price: 299.99,
	})
	fmt.Printf("Updated Products: %v\n", products)
}
