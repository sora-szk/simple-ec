package productPresenters

import (
	domains "github.com/sora_szk/simple-ec/app/product/domains"
)

// HttpProductPresenter は、製品関連の出力を提示するためのインターフェースを提供します。
type HttpProductPresenter interface {
	// PresentProducts は複数の製品情報を提示します。
	PresentProducts(products []*domains.Product) ([]byte, error)

	// PresentProduct は1つの製品情報を提示します。
	PresentProduct(product *domains.Product) ([]byte, error)

	// PresentProductCreation は新しく作成された製品のIDを提示します。
	PresentProductCreation(productID domains.ProductID) ([]byte, error)

	// PresentError はエラー情報を提示します。
	PresentError(err error) ([]byte, error)
}
