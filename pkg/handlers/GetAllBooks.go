package handlers

import (
	"crud/pkg/mocks"
	"crud/pkg/models"
	"encoding/json"
	"net/http"
)

// swagger:response BookRes
type BookRes struct {
	// in: body
	Body []models.Book
}

// swagger:route GET /books books ListBook
// responses:
// 200:BookRes

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Books)
}
