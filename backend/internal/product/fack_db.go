package product

var ProductList []Product

func LoadFakeProducts() {
	prod1 := Product{ID: 1, Title: "Banana", Description: "High potassium", Price: 125.22}
	prod2 := Product{ID: 2, Title: "Orange", Description: "Vitamin C", Price: 90.00}
	prod3 := Product{ID: 3, Title: "Apple", Description: "Healthy", Price: 150.99}

	ProductList = append(ProductList, prod1, prod2, prod3)
}
