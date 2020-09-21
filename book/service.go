package book

import (
	"context"
	"log"
	"net/http"

	"graphql/infrastruktur"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
)

/* Rest API */
func RestApiGetBook(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	books := &Book{}
	bookName := chi.URLParam(r, "bookname")

	cur, err := infrastruktur.Mongodb.Collection("booklist").Find(ctx, bson.M{"name": bookName})
	defer cur.Close(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	for cur.Next(ctx) {
		cur.Decode(&books)

	}

	if s := books; s != nil {
		HttpResponseSuccess(w, r, books)
		return
	}
	return
}
