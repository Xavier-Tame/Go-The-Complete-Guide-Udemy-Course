package product

type Product struct {
	Title string
	Id    int
	Price float64
}

var products []Product

func AddProduct(title string, id int, price float64) Product {
	product := Product{
		Title: title,
		Id:    id,
		Price: price,
	}
	products = append(products, product)
	return product
}

func GetProducts() []Product {
	return products
}
