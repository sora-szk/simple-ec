package productUsecases

import (
	domains "github.com/sora_szk/simple-ec/app/product/domains"
)

// productInteractor は商品に関連するユースケースの具体的な実装です。
type productInteractor struct {
	repo domains.ProductRepository
}

// NewProductInteractor は新しい商品インタラクタを初期化します。
func NewProductInteractor(repo domains.ProductRepository) ProductUsecase {
	return &productInteractor{repo: repo}
}

func (p *productInteractor) ListAllProducts() ([]*domains.Product, error) {
	return p.repo.FindAll()
}

func (p *productInteractor) ListActiveProducts() ([]*domains.Product, error) {
	return p.repo.FindAllActive()
}

func (p *productInteractor) GetProductByID(id domains.ProductID) (*domains.Product, error) {
	return p.repo.FindByID(id)
}

func (p *productInteractor) CreateProduct(input domains.NewProductInput) (domains.ProductID, error) {
	product := domains.NewProduct(input)
	err := p.repo.Create(product)
	if err != nil {
		return domains.IllegalProductID(), err
	}
	return product.ID, nil
}

func (p *productInteractor) UpdateProduct(id domains.ProductID, input domains.UpdateProductInput) error {
	return p.repo.Update(id, input)
}

func (p *productInteractor) DeleteProduct(id domains.ProductID) error {
	return p.repo.Delete(id)
}
