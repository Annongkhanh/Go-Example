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
	gh := handlers.NewGoodbye(l)

	// log.Println("Hello from", hh)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	http.ListenAndServe(":9090", sm)
}