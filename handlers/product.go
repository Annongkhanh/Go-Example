package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Annongkhanh/Go-Example/data"
)

type Product struct{
	l *log.Logger
}

func NewProduct(l*log.Logger) *Product{
	return &Product{l}
}

func(p*Product) ServeHTTP(rw http.ResponseWriter, r*http.Request){
	lp := data.GetProducts()
	
	d, err := 	json.Marshal(lp)
	if err != nil{
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	// fmt.Fprintf(rw, "Data %s", d)
	rw.Write(d)
}

