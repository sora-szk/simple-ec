package productPresenters

import (
	"encoding/json"

	"github.com/google/uuid"
	domains "github.com/sora_szk/simple-ec/app/product/domains"
)

type HttpProductPresenterImpl struct{}

// NewHTTPProductPresenter は新しいHttpProductPresenterImplを初期化して返します。
func NewHTTPProductPresenter() HttpProductPresenter {
	return &HttpProductPresenterImpl{}
}

func (h *HttpProductPresenterImpl) PresentProducts(products []*domains.Product) ([]byte, error) {
	return json.Marshal(products)
}

func (h *HttpProductPresenterImpl) PresentProduct(product *domains.Product) ([]byte, error) {
	return json.Marshal(product)
}

func (h *HttpProductPresenterImpl) PresentProductCreation(productID domains.ProductID) ([]byte, error) {
	response := map[string]string{
		"product_id": uuid.UUID(productID).String(),
	}
	return json.Marshal(response)
}

func (h *HttpProductPresenterImpl) PresentError(err error) ([]byte, error) {
	response := map[string]string{
		"error": err.Error(),
	}
	return json.Marshal(response)
}
