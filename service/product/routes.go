package product

import (
	"fmt"
	"net/http"

	"github.com/donnykd/go-ecom/types"
	"github.com/donnykd/go-ecom/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	storage types.ProductStorage
}

func NewHandler(storage types.ProductStorage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) RegisterRoutes(subrouter *mux.Router) {
	subrouter.HandleFunc("/products", h.handleGetProduct).Methods("GET")
	subrouter.HandleFunc("/products", h.handleCreateProduct).Methods("POST")
}

func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	ps, err := h.storage.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateProduct
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.storage.CreateProduct(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, payload)
}
