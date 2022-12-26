package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chris124567/requiz/api"
	"github.com/chris124567/requiz/web"
)

func main() {
	fmt.Println("Starting...")
	server := &http.Server{
		Addr:         "0.0.0.0:8000",
		Handler:      web.NewHandler("./data/static", "./data/template", api.NewDefaultClient()),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	panic(server.ListenAndServe())
}
