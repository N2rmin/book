package handlers

import (
	"crud/pkg/mocks"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all mock Books

	for index, book := range mocks.Books {

		if book.Id == id {
			// Delete book and send response if the book Id matches dynamic Id
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break

		}
	}

}
