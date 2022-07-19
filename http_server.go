package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Message struct {
	Name      string
	Last_name string
	Age       int
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
    <head>
        <title>Hello World</title>
    </head>
    <body>
        Hello World!
    `,
	)
	io.WriteString(
		res,
		`<a href='http://127.0.0.1:9000/assets/'>PON</a>`,
	)
	io.WriteString(res, `</body>
	</html>`)
}

func json_sender(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"application/json",
	)
	m := Message{"Valera", "rabchenkov", 16}
	b, err := json.Marshal(m)
	fmt.Println(err)
	io.WriteString(res, string(b))
}

func json_receiver(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	//fmt.Println(html.EscapeString(req.GetBody()))
	var arr_body []byte
	var body string
	arr_body, _ = ioutil.ReadAll(req.Body)
	body = string(arr_body)
	fmt.Println(body)
	var m Message
	encode_err := json.Unmarshal(arr_body, &m)
	fmt.Println(m, encode_err)
	//fmt.Println(req.Header["Content-Type"])
	output, decode_err := json.Marshal(m)
	fmt.Println(decode_err)
	io.WriteString(res, string(output))
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/json/get", json_sender)
	http.HandleFunc("/json/post", json_receiver)
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)
	http.ListenAndServe(":9000", nil)
}
