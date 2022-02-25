package handlers

import (
	"crud/pkg/mocks"
	"crud/pkg/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

// swagger:route POST /books books Book
// responses:
// 200:bookResponse

//swagger:parameters Book
type BookReq struct {
	// in: body
	//required: true
	Name models.Book
}

// swagger:response bookResponse
type bookResponse struct {
	// in: body
	Body []models.Book
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book

	json.Unmarshal(body, &book)

	// Append to the Book mocks
	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	// Send 201 created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")

}
