package main

import (
	rice "github.com/GeertJohan/go.rice"
	"net/http"
)

func main() {
	box := rice.MustFindBox("../files")
	httpbox := box.HTTPBox()
	http.ListenAndServe(":8888", http.FileServer(httpbox))

	//box := rice.MustFindBox("templates")
	//templateStr, err := box.String("example.html")
	//if err != nil {
	//	log.Fatal(err)
	//}
}
