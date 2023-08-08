package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successfull\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request is successfull\n")
	fn := r.FormValue("First name")
	ln := r.FormValue("Last Name")
	sx := r.FormValue("sex")
	fmt.Fprintf(w, "Name = %s\n", fn)
	fmt.Fprintf(w, "Last Name = %s\n", ln)
	switch sx {
	case "1":
		fmt.Fprintf(w, "Sex = F\n")
	case "2":
		fmt.Fprintf(w, "Sex = M\n")
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusFound)
		return
	}
	fmt.Fprint(w, "hello")
}
func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/data", dataHandler)
	//http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server is listening")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
