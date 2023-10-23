package productDomains

import (
	"time"
)

// ProductType は商品の種類を表します。
type ProductType int

const (
	// PhysicalProduct は物理的な商品を表します。
	PhysicalProduct ProductType = iota
	// DigitalProduct はデジタルな商品を表します。
	DigitalProduct
)

// Product は商品ドメインの中心エンティティを表します。
type Product struct {
	ID          ProductID   `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Type        ProductType `json:"type"`
	Quantity    int         `json:"quantity"`
	IsActive    bool        `json:"is_active"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// NewProductInput は新しい商品を作成するための入力データを表します。
type NewProductInput struct {
	Name        string
	Description string
	Price       float64
	ProductType ProductType
	Quantity    int
	IsActive    bool
}

// NewProduct は新しい商品エンティティを初期化します。
// デジタル商品の場合、数量は0に設定されます。
func NewProduct(input NewProductInput) *Product {
	if input.ProductType == DigitalProduct {
		input.Quantity = 0
	}

	return &Product{
		ID:          NewProductID(),
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Type:        input.ProductType,
		Quantity:    input.Quantity,
		IsActive:    input.IsActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
