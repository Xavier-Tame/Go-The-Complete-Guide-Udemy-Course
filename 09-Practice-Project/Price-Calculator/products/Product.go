package products

import (
	"encoding/json"
	"log"
	"os"
)

type Product struct {
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

func (p Product) PriceWithTax() (PriceWithTax float64) {
	return (p.Price * (p.Tax / 100)) + p.Price
}

type TaxCalculator interface {
	PriceWithTax() float64
}

type Saver interface {
	SaveToFile(filename string) error
}

type Loader interface {
	LoadFromFile(fileName string) error
}

type PrettyJSONStore struct {
	Products *[]Product `json:"products"`
}

func (products *PrettyJSONStore) SaveToFile(filename string) error {
	jsonData, err := json.MarshalIndent(products.Products, " ", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (products *PrettyJSONStore) LoadFromFile(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, products.Products)
	if err != nil {
		return err
	}

	return err
}

type NormalJSONStore struct {
	Products *[]Product `json:"products"`
}

func (products *NormalJSONStore) SaveToFile(filename string) error {
	jsonData, err := json.Marshal(products.Products)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (products *NormalJSONStore) LoadFromFile(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, products.Products)
	if err != nil {
		return err
	}

	return err
}

func LoadFromFile(fileName string, products *[]Product) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, products)
	if err != nil {
		return err
	}

	return err
}
