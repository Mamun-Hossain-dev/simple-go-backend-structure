package product

type ProductRepository interface {
	GetAll() []Product
	GetByID(id int) *Product
	Store(p Product) Product
	Update(id int, p Product) *Product
	Delete(id int) *Product
}

type productRepo struct{}

func NewProductRepository() ProductRepository {
	return &productRepo{}
}

func (r *productRepo) GetAll() []Product {
	return ProductList
}

func (r *productRepo) GetByID(id int) *Product {
	for _, p := range ProductList {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

func (r *productRepo) Store(p Product) Product {
	p.ID = len(ProductList) + 1
	ProductList = append(ProductList, p)
	return p
}

func (r *productRepo) Update(id int, p Product) *Product {
	for i, prod := range ProductList {
		if prod.ID == id {
			ProductList[i] = p
			return &ProductList[i]
		}
	}
	return nil
}

func (r *productRepo) Delete(id int) *Product {
	for i, p := range ProductList {
		if p.ID == id {
			deletedProduct := p
			ProductList = append(ProductList[:i], ProductList[i+1:]...)
			return &deletedProduct
		}
	}
	return nil
}
