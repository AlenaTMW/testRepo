package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	name    = ""
	address = ""
)

func handler(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "<!DOCTYPE html>"+
			"<html>"+
			"<head>"+
			"<meta charset=\"UTF-8\" />"+
			"</head>"+
			"<body>"+
			"<div>Known token: <span id=\"known-token\"></span></div>"+
			"<div>"+
			"<form action=\"http://localhost:8080\" method=\"POST\" action=\"/\">"+
			"<label>Name</label><input name=\"name\" type=\"text\" value=\"\" />"+
			"<label>Address</label><input name=\"address\" type=\"text\" value=\"\" />"+
			"<input type=\"submit\" value=\"submit\" />"+
			"</form>"+
			"</div>"+
			"</body>"+
			"<script type=\"text/javascript\">"+
			"const elem = document.querySelector(\"#known-token\");"+
			"let token = document.cookie.split(';').filter((item) => item.trim().startsWith('token='))[0];"+
			"elem.innerHTML = token.replace('token=', '');"+
			"</script>"+
			"</html>")
	case "POST":

		name = r.PostFormValue("name")
		address = r.PostFormValue("address")

		//http.Post() to second app
		message := map[string]interface{}{
			"token": name + ":" + address,
		}

		bytesRepresentation, err := json.Marshal(message)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := http.Post("http://app2:8081", "application/json", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			log.Fatalln(err)
		}

		// save cookie
		cookie := &http.Cookie{
			Name:       "token",
			Value:      name + ":" + address,
			Path:       "",
			Domain:     "",
			Expires:    time.Time{},
			RawExpires: "",
			MaxAge:     0,
			Secure:     false,
			HttpOnly:   false,
			SameSite:   0,
			Raw:        "",
			Unparsed:   []string{},
		}
		http.SetCookie(w, cookie)
		fmt.Fprint(w, "<!DOCTYPE html>"+
			"<html>"+
			"<head>"+
			"<meta charset=\"UTF-8\" />"+
			"</head>"+
			"<body>"+
			"<div>Known token: <span id=\"known-token\"></span></div>"+
			"<div>"+
			"<form action=\"http://localhost:8080\" method=\"POST\" action=\"/\">"+
			"<label>Name</label><input name=\"name\" type=\"text\" value=\"\" />"+
			"<label>Address</label><input name=\"address\" type=\"text\" value=\"\" />"+
			"<input type=\"submit\" value=\"submit\" />"+
			"</form>"+
			"</div>"+
			"</body>"+
			"<script type=\"text/javascript\">"+
			"const elem = document.querySelector(\"#known-token\");"+
			"let token = document.cookie.split(';').filter((item) => item.trim().startsWith('token='))[0];"+
			"elem.innerHTML = token.replace('token=', '');"+
			"</script>"+
			"</html>")
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)

		log.Println(result)
		//fmt.Println(resp.Body)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	// server port number
	const port = 8080

	log.Printf("Launching server on port: %d \n\n", port)

	// set handler for route '/'
	http.HandleFunc("/", handler)
	// start server without ending
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
