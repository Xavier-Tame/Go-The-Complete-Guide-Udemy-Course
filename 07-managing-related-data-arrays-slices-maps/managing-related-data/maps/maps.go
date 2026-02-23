package maps

import "fmt"

type Product struct {
	id    string
	title float32
	price float64
}

func main() {
	websites := map[string]string{
		"Google":              "https://www.google.com",
		"GitHub":              "https://www.github.com",
		"StackOverflow":       "https://stackoverflow.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])
	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
