package database

// Product type
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

// product slice
var productList []Product

type ProductsResponse struct {
	Message string    `json:"message"`
	Data    []Product `json:"data"`
}

type ProductResponse struct {
	Message string  `json:"message"`
	Data    Product `json:"data"`
}

func GetProductList() []Product {
	return productList
}

func StoreProduct(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func GetProductById(pId int) *Product {
	for i := range productList {
		if productList[i].ID == pId {
			return &productList[i]
		}
	}
	return nil
}

func UpdateProduct(pId int, updatedProduct Product) *Product {
	for i, p := range productList {
		if p.ID == pId {
			updatedProduct.ID = p.ID
			productList[i] = updatedProduct
			return &productList[i]
		}
	}
	return nil
}

func DeleteProduct(pId int) *Product {
	for i, p := range productList {
		if p.ID == pId {
			deletedProduct := p

			// remove the product from slice
			productList = append(productList[:i], productList[i+1:]...)

			return &deletedProduct
		}
	}
	return nil
}

func LoadProducts() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is one of my fav. fruit. this is an Orange",
		Price:       120.2423,
		ImageUrl:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR9yz1d_OrZBKB6TIWRyUtCIPBIjgyDOpybxw&s",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Banana",
		Description: "Banana is one of my fav. fruit. this is full of K",
		Price:       20.2423,
		ImageUrl:    "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAVm9biNM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/ar-new-banana-adobe-ar-2x1-917fdde58d194b529b41042ebff1c031.jpg",
	}

	prd3 := Product{
		ID:          3,
		Title:       "PineApple",
		Description: "PineApple is one of my fav. fruit. this is an Orange",
		Price:       70.2423,
		ImageUrl:    "https://img.freepik.com/free-photo/pineapple-fruit_1203-7746.jpg?semt=ais_hybrid&w=740&q=80",
	}

	// append the above products to slice productList
	productList = append(productList, prd1, prd2, prd3)
}
