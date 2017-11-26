package main

import (
	"net/http"

	"github.com/sasasaiki/web-templete-go-ts-pug-sass/src/go"
)

func main() {
	r := imageFileServer.CreateRoute(imageFileServer.NewProdHandler())
	//cssやjsを読み込めるようにするHandler
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	http.ListenAndServe(":8080", r)
}
