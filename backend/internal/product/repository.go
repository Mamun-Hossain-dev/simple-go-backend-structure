package product

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(id int) (*Product, error)
	Store(p Product) (*Product, error)
	Update(id int, p Product) (*Product, error)
	Delete(id int) (*Product, error)
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) GetAll() ([]Product, error) {
	var products []Product
	err := r.db.Select(&products, "SELECT * FROM products;")
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepo) GetByID(id int) (*Product, error) {
	query := `SELECT * FROM products WHERE id=$1`

	var p Product
	err := r.db.Get(&p, query, id)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Store(p Product) (*Product, error) {
	query := `
	INSERT INTO products (title, description, price, img_url) 
	VALUES ($1, $2, $3, $4)
	RETURNING id, title, description, price, img_url
	`

	var createdProduct Product

	err := r.db.Get(&createdProduct, query, p.Title, p.Description, p.Price, p.ImgUrl)
	if err != nil {
		return nil, err
	}
	return &createdProduct, nil
}

func (r *productRepo) Update(id int, p Product) (*Product, error) {
	query := `
	UPDATE products 
	SET title=$1, description=$2, price=$3, img_url=$4
	WHERE id=$5
	RETURNING id, title, description, price, img_url 
	`

	var updated Product
	err := r.db.Get(&updated, query, p.Title, p.Description, p.Price, p.ImgUrl, id)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("product not found")
	}
	return &updated, nil
}

func (r *productRepo) Delete(id int) (*Product, error) {
	query := `
	DELETE FROM products
	WHERE id=$1 
	RETURNING id, title, description, price, img_url
	`

	var deleted Product
	err := r.db.Get(&deleted, query, id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return &deleted, nil
}
