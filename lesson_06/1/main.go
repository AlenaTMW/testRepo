package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response map[string]interface{}

func (r Response) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		s := ""
		return s
	}
	s := string(b)
	return s
}

func handler(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, Response{"Host": r.Host, "UserAgent": r.UserAgent(), "RequestURI": r.RequestURI, "Header": r.Header})
	default:
		fmt.Fprintf(w, "Sorry, only GET method is supported.")
	}
}

func main() {
	// server port number
	const port = 8080

	fmt.Printf("Launching server on port: %d \n\n", port)

	// set handler for route '/'
	http.HandleFunc("/", handler)
	// start server without ending
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
