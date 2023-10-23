package productControllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	domains "github.com/sora_szk/simple-ec/app/product/domains"
	presenters "github.com/sora_szk/simple-ec/app/product/presenters"
	usecases "github.com/sora_szk/simple-ec/app/product/usecases"
)

type HTTPProductHandlerImpl struct {
	usecase   usecases.ProductUsecase
	presenter presenters.HttpProductPresenter
}

func NewHTTPProductHandler(u usecases.ProductUsecase, p presenters.HttpProductPresenter) HTTPProductHandler {
	return &HTTPProductHandlerImpl{
		usecase:   u,
		presenter: p,
	}
}

func (h *HTTPProductHandlerImpl) jsonResponse(w http.ResponseWriter, statusCode int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func (h *HTTPProductHandlerImpl) errorResponse(w http.ResponseWriter, statusCode int, err error) {
	data, _ := h.presenter.PresentError(err)
	h.jsonResponse(w, statusCode, data)
}

func (h *HTTPProductHandlerImpl) HandleListAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.usecase.ListAllProducts()
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, _ := h.presenter.PresentProducts(products)
	h.jsonResponse(w, http.StatusOK, data)
}

func (h *HTTPProductHandlerImpl) HandleListActiveProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.usecase.ListActiveProducts()
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, _ := h.presenter.PresentProducts(products)
	h.jsonResponse(w, http.StatusOK, data)
}

func (h *HTTPProductHandlerImpl) HandleGetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, err)
		return
	}

	product, err := h.usecase.GetProductByID(domains.ProductID(id))
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, _ := h.presenter.PresentProduct(product)
	h.jsonResponse(w, http.StatusOK, data)
}

func (h *HTTPProductHandlerImpl) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var input domains.NewProductInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, err)
		return
	}

	productID, err := h.usecase.CreateProduct(input)
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, _ := h.presenter.PresentProductCreation(productID)
	h.jsonResponse(w, http.StatusCreated, data)
}

func (h *HTTPProductHandlerImpl) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		h.errorResponse(w, http.StatusBadRequest, errors.New("ID path parameter is required"))
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, err)
		return
	}

	var input domains.UpdateProductInput
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&input)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.usecase.UpdateProduct(domains.ProductID(id), input)
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPProductHandlerImpl) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		h.errorResponse(w, http.StatusBadRequest, errors.New("ID path parameter is required"))
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.usecase.DeleteProduct(domains.ProductID(id))
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
