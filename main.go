package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const templatePath = "./templates"
const layoutPath = templatePath + "/layout.html"

var (
	indexTemplate = template.Must(template.ParseFiles(layoutPath, templatePath+"/index.html"))
)

func main() {
	http.HandleFunc("/", indexHandler)
	fmt.Println("Server is listening on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"PageTitle":  "Hello World",
		"Text":   "Hello World",
	})
}
	
