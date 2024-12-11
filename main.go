package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jeki00/UrlShortener_go_vue/controller"
	"github.com/rs/cors"
)

func main() {
	router := http.NewServeMux()

	// mux.Handle("/api/", apiHandler{})
	router.HandleFunc("GET /url", controller.GetAll)
	router.HandleFunc("GET /{url}", controller.GetUrl)
	router.HandleFunc("GET /detail/{shortUrl}", controller.DetailUrl)
	router.HandleFunc("POST /url", controller.PostUrl)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", handler))
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	w.Write([]byte("ini halaman get"))
}
