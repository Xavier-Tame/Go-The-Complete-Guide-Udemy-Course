package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	Products "example.com/price-calculator/products"
)

func ListProducts(products []Products.Product) {
	for _, p := range products {
		fmt.Printf("Price: %.2f, Tax: %.2f\n", p.Price, p.Tax)
	}
}

func ShowProductsWithTax(products []Products.Product) {
	for _, p := range products {
		fmt.Printf("Price Before Tax: %.2f, Price After Tax: %.2f\n", p.Price, p.PriceWithTax())
	}
}

func NameOfFile() string {
	var fileName string
	for {
		fmt.Print("Please enter name of file (Without .json): ")

		reader := bufio.NewReader(os.Stdin)
		fileName, _ = reader.ReadString('\n')

		fileName = strings.TrimSpace(fileName)
		fileName = strings.ReplaceAll(fileName, " ", "_")

		if fileName != "" {
			return fileName + ".json"
		}

		fmt.Println("Filename cannot be empty.")
	}
}

func main() {
	products := []Products.Product{}

	prettyJSONStore := &Products.PrettyJSONStore{}
	normalJSONStore := &Products.NormalJSONStore{}

	prettyJSONStore.Products = &products
	normalJSONStore.Products = &products

	fmt.Printf(" Welcome to Products Tax Calculator\n Please Choose Option Using the Numbers Shown \n")
	var option int

	for option != 6 {
		fmt.Printf(" 1: Add Product\n 2: List Products Added\n 3: Calculate Tax Of All Products\n 4: Save Products To File \n 5: Load From File \n 6: Exit \n")
		fmt.Printf("Option: ")
		_, err := fmt.Scanln(&option)

		switch option {
		case 1:
			{

				fmt.Printf("Please provide price of product: ")
				var productPrice float64
				_, err := fmt.Scanln(&productPrice)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					break
				}

				fmt.Printf("Please provide tax of product (%%): ")
				var productTax float64
				_, err = fmt.Scanln(&productTax)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					break
				}

				products = append(products, Products.Product{Price: productPrice, Tax: productTax})
			}
		case 2:
			{
				for _, v := range products {
					fmt.Printf("Product Price: %.2f Product Tax: %.2f\n", v.Price, v.Tax)
				}
			}
		case 3:
			{
				for _, v := range products {
					fmt.Printf("Product Price Before Tax: %.2f, Product Tax: %.2f, Product Price After Tax: %.2f\n", v.Price, v.Tax, v.PriceWithTax())
				}
			}
		case 4:
			{
				var prettyOption int
				for prettyOption != 3 {
					fmt.Printf(" How Would You Like It Stored?\n 1: Pretty JSON \n 2: Compact JSON \n 3: Exit\n Option: ")
					_, err := fmt.Scanln(&prettyOption)
					if err != nil {
						fmt.Printf("Error: %v\n", err)
						break
					}

					switch prettyOption {
					case 1:
						fileName := NameOfFile()
						err := prettyJSONStore.SaveToFile(fileName)

						if err != nil {
							fmt.Printf("Error: %v", err)
							return
						} else {
							fmt.Printf("File Saved Successfully\n")
						}
					case 2:
						fileName := NameOfFile()
						err := normalJSONStore.SaveToFile(fileName)

						if err != nil {
							fmt.Printf("Error: %v", err)
							return
						} else {
							fmt.Printf("File Saved Successfully\n")
						}
					case 3:
						continue
					default:
						fmt.Printf("Please enter an appropiate option, option entered: %v", prettyOption)
					}
				}

			}
		case 5:
			{
				nameOfFile := NameOfFile()
				err := Products.LoadFromFile(nameOfFile, &products)
				// err := prettyJSONStore.LoadFromFile(nameOfFile)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					continue
				}
			}
		case 6:
			{
				return
			}
		default:
			{
				if err != nil {
					fmt.Println("Please Enter Valid Option, Error: ", err)
				} else {
					fmt.Println("Please Enter Number From Options, Number Entered: ", option)
				}
				continue
			}
		}
	}
}
