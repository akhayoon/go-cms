package main

import (
	"net/http"

	"goprojects/cms"
)

func main() {
	http.HandleFunc("/", cms.ServeIndex)
	http.HandleFunc("/page/", cms.ServePage)
	http.HandleFunc("/post/", cms.ServePost)
	http.HandleFunc("/new", cms.HandleNew)
	http.ListenAndServe(":3000", nil)
}
