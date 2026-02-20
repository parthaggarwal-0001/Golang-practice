package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello from server")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}

	fmt.Fprintf(w, "Post request successful\n")

	name := r.FormValue("name")
	password := r.FormValue("Password")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Password = %s\n", password)
}

func main() {
	fileServer := http.FileServer(http.Dir("./"))

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	http.Handle("/", fileServer)

	fmt.Println("Your server is running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
