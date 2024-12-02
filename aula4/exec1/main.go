package main

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func main() {
	produtos := make(Product, 0, 10)

}
