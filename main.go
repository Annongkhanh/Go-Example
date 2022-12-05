package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Annongkhanh/Go-Example/handlers"
)

func main()  {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	log.Println("Hello from", hh)
	http.ListenAndServe(":9090", nil)
}