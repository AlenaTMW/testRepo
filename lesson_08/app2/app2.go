package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
//name    = ""
//address = ""
)

func handler(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "POST":
		var result map[string]interface{}
		json.NewDecoder(r.Body).Decode(&result)

		log.Println(result["token"])

		currentDate := time.Now()

		message := map[string]interface{}{
			"token":     result["token"],
			"createdAt": currentDate,
			"expiredAt": currentDate.Add(time.Hour * 24 * 10),
		}

		bytesRepresentation, err := json.Marshal(message)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(message)

		w.Write(bytesRepresentation)
		if err != nil {
			log.Fatalln(err)
		}

	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}

func main() {
	// server port number
	const port = 8081

	log.Printf("Launching server on port: %d \n\n", port)

	// set handler for route '/'
	http.HandleFunc("/", handler)
	// start server without ending
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
