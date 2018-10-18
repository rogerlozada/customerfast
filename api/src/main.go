package main

import (
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"customerfast/api/src/customerContext/application/customer"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Couldn't load .env file")
	}

	router := mux.NewRouter()
	router = customerapplication.GetCustomerRouter(router)

	configuracaoCors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	n := negroni.Classic()
	n.Use(configuracaoCors)
	n.UseHandler(router)
	n.Run(os.Getenv("APPLICATION_PORT"))
}
