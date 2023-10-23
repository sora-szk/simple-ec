package productDomains

// UpdateProductInput は商品の更新のための入力データを表します。
// 各フィールドがポインタ型になっているのは、データが提供されていない場合を示すためです。
type UpdateProductInput struct {
	Name        *string      `json:"name,omitempty"`
	Description *string      `json:"description,omitempty"`
	Price       *float64     `json:"price,omitempty"`
	Type        *ProductType `json:"type,omitempty"`
	Quantity    *int         `json:"quantity,omitempty"`
	IsActive    *bool        `json:"is_active,omitempty"`
}

// ProductRepository はデータストアでの商品へのアクセスや変更のためのインターフェースを提供します。
type ProductRepository interface {
	// FindByID は指定されたIDを持つ商品を探し出します。
	FindByID(id ProductID) (*Product, error)

	// FindAll はすべての商品を取得します。
	FindAll() ([]*Product, error)

	// FindAllActive はアクティブな商品のみを取得します。
	FindAllActive() ([]*Product, error)

	// Create は新しい商品をデータストアに追加します。
	Create(product *Product) error

	// Update は指定されたIDを持つ商品のデータを更新します。
	Update(id ProductID, input UpdateProductInput) error

	// Delete は指定されたIDを持つ商品を削除します。
	Delete(id ProductID) error
}
