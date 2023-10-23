package productControllers

import (
	"net/http"
)

// HTTPProductHandler は、製品関連のHTTPリクエストを処理するためのインターフェースを提供します。
// このインターフェースは、各HTTPエンドポイントに対応するハンドル関数を定義しています。
type HTTPProductHandler interface {
	// HandleListAllProducts は、すべての製品を一覧表示するリクエストを処理します。
	HandleListAllProducts(w http.ResponseWriter, r *http.Request)

	// HandleListActiveProducts は、アクティブな製品のみを一覧表示するリクエストを処理します。
	HandleListActiveProducts(w http.ResponseWriter, r *http.Request)

	// HandleGetProductByID は、指定されたIDの製品情報を取得するリクエストを処理します。
	HandleGetProductByID(w http.ResponseWriter, r *http.Request)

	// HandleCreateProduct は、新しい製品を作成するリクエストを処理します。
	HandleCreateProduct(w http.ResponseWriter, r *http.Request)

	// HandleUpdateProduct は、指定されたIDの製品情報を更新するリクエストを処理します。
	HandleUpdateProduct(w http.ResponseWriter, r *http.Request)

	// HandleDeleteProduct は、指定されたIDの製品を削除するリクエストを処理します。
	HandleDeleteProduct(w http.ResponseWriter, r *http.Request)
}
