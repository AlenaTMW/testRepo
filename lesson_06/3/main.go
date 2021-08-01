package main

import (
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
		// save cookie
		name = r.PostFormValue("name")
		address = r.PostFormValue("address")
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

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
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
