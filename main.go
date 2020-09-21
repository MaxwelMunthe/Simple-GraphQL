package main

import (
	"github.com/go-chi/chi"
	"graphql/book"
	"graphql/infrastruktur"
	"log"
	"net/http"
	"net/url"
)


func main() {
	routes := chi.NewRouter()
	r := book.RegisterRoutes(routes)
	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func init() {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	env := infrastruktur.Environment{}
	env.SetEnvironment()
	env.LoadConfig()
	env.InitMongoDB()
}
