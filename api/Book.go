package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
	"httprouter/models"
	"httprouter/lib"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, "hello httprouter!")
}

func BookCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	book := &models.Book{}
	if err := lib.PopulateModelFromHandler(w, r, book); err != nil {
		models.SendJsonErrResponse(w, http.StatusInternalServerError, "Unprocessible Entity!")
		return
	}

	//存储book数据

	bookstore := models.GetBookstore()
	models.SendJsonResponse(w, bookstore[book.ISDN])
}

func BookIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bookstore := models.GetBookstore()
	books := [] *models.Book{}
	for _, book := range bookstore{
		books = append(books, book)
	}
	//models.SendJsonResponse(w, bookstore)
	models.SendJsonResponse(w, books)
}

func BookShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	isdn := params.ByName("isdn")
	bookstore := models.GetBookstore()
	book, ok := bookstore[isdn]
	if !ok {
		models.SendJsonErrResponse(w, http.StatusInternalServerError, "Record Not Found")
		return
	}

	models.SendJsonResponse(w, book)
}