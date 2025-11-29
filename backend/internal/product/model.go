package product

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"imageUrl" db:"img_url"`
}

type Response struct {
	Message string    `json:"message"`
	Data    []Product `json:"data"`
}

type ProductResponse struct {
	Message string  `json:"message"`
	Data    Product `json:"data"`
}
