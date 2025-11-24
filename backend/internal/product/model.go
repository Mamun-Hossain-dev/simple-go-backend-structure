package product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

type Response struct {
	Message string    `json:"message"`
	Data    []Product `json:"data"`
}

type ProductResponse struct {
	Message string  `json:"message"`
	Data    Product `json:"data"`
}
