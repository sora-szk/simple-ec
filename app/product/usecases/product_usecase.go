package productUsecases

import (
	domains "github.com/sora_szk/simple-ec/app/product/domains"
)

// ProductUsecase は商品に関連するユースケースを表します。
type ProductUsecase interface {
	// ListAllProducts は全ての商品をリストとして取得します。
	ListAllProducts() ([]*domains.Product, error)

	// ListActiveProducts はアクティブな商品のみをリストとして取得します。
	ListActiveProducts() ([]*domains.Product, error)

	// GetProductByID は指定されたIDに該当する商品を取得します。
	GetProductByID(id domains.ProductID) (*domains.Product, error)

	// CreateProduct は新しい商品を作成します。
	// 成功時には新しく作成された商品のIDを返します。
	CreateProduct(input domains.NewProductInput) (domains.ProductID, error)

	// UpdateProduct は指定されたIDを持つ商品の情報を更新します。
	UpdateProduct(id domains.ProductID, input domains.UpdateProductInput) error

	// DeleteProduct は指定されたIDを持つ商品を削除します。
	DeleteProduct(id domains.ProductID) error
}
