package lists

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

/*func main() {
	productNames := [4]string{"Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println("Prices:", prices)

	productNames[2] = "A carpet"

	fmt.Println("Product Names:", productNames)

	fmt.Println(prices[0])

	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99
	highlightedPrices := featuredPrices[:1]
	fmt.Println("Featured Prices:", featuredPrices)
	fmt.Println("Highlighted Prices:", highlightedPrices)
	fmt.Println(prices)
	fmt.Println("Length of prices array:", len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println("Length of prices array:", len(highlightedPrices), cap(highlightedPrices))
}*/
