package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func handleSearch(w http.ResponseWriter, r *http.Request) {
	getSearchTerms := r.URL.Query().Get("title")
	response := searchOpenLibrary(strings.Fields(getSearchTerms), 1, 20000)
	_, err := w.Write([]byte(response))
	if err != nil {
		fmt.Println("error handleMain", err)
	}
}

func applicationJSON() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	})
}

func main() {

	n := negroni.Classic()
	n.Use(applicationJSON())

	r := mux.NewRouter()
	n.UseHandler(r)

	r.HandleFunc("/search-by-title", handleSearch).Methods("GET")

	err := http.ListenAndServe(":3000", n)
	if err != nil {
		fmt.Println(err)
	}

}

func searchOpenLibrary(searchTerms []string, page int, perPage int) []byte {

	endpoint := "https://openlibrary.org/search.json?"+
		"title=" + strings.Join(searchTerms, "+") + 
		"&page=" + strconv.Itoa(page) + 
		"&limit=" + strconv.Itoa(perPage) + 
		"&fields=title" 

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}
	return []byte(body)
}
