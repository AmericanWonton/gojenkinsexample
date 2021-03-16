package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func init() {
	template1 = template.Must(template.ParseGlob("./static/templates/*"))
}

/* TEMPLATE DEFINITION BEGINNING */
var template1 *template.Template

func index(w http.ResponseWriter, r *http.Request) {
	/* Execute template, handle error */
	err1 := template1.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err1)
}

// Handle Errors passing templates
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	fmt.Println("DEBUG: We are handling requests on 8080")

	http.Handle("/favicon.ico", http.NotFoundHandler()) //For missing FavIcon
	myRouter.HandleFunc("/", index)
	//Serve our static files
	myRouter.Handle("/", http.FileServer(http.Dir("./static")))
	myRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Here we are in main")
	//Handle Requests
	handleRequests()
}
