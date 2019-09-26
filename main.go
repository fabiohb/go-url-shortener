package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.POST("/", short)
	router.GET("/assets/*filepath", assets)
	router.GET("/u/:urlRedirID", urlRedir)

	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("ERROS: ", err)
	}
}

func short(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("ERROS: ", err)
	}
}

func assets(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	filepath := ps.ByName("filepath")
	filepath = "./assets" + filepath
	http.ServeFile(w, req, filepath)
}

func urlRedir(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	urlRedirID := ps.ByName("urlRedirID")

	fmt.Fprintf(w, "USER, %s!\n", urlRedirID)
}
