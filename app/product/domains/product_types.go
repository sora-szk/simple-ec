package productDomains

import "github.com/google/uuid"

// ProductID は商品の一意の識別子を表します。
type ProductID uuid.UUID

// NewProductID は新しい一意の商品IDを生成します。
func NewProductID() ProductID {
	return ProductID(uuid.New())
}

// IllegalProductID は無効な商品IDを表します。
// このIDは商品が存在しない場合や不正なIDが与えられた場合などに使用されます。
func IllegalProductID() ProductID {
	return ProductID(uuid.Nil)
}
