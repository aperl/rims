package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

var emptyBody = make([]byte, 0)
var emptyString = ""

func printHelp() {
	println("Use: rims <PORT>")
	println("To set the mock response POST to host:PORT/mock")
	println("The Content-Type in the post will be returned in the body of the next request")
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		printHelp()
		return
	}

	port, err := strconv.Atoi(args[0])
	if err != nil {
		printHelp()
		return
	}

	if port < 0 || port > 499151 {
		printHelp()
		return
	}

	r := mux.NewRouter()
	body := &emptyBody
	contentType := &emptyString
	r.HandleFunc("/mock", loadMock(body, contentType)).Methods(http.MethodPost)
	r.PathPrefix("/").Handler(getMock(body, contentType))
	println("Listening on port ", args[0])

	if err := http.ListenAndServe(":"+args[0], r); err != nil {
		println("Error: " + err.Error())
	}
	println("Closeing")
}

func loadMock(body *[]byte, contentType *string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t := r.Header.Get("Content-Type")
		if len(b) > 0 && len(t) == 0 {
			http.Error(w, "If a body is present then there must be a Content-Type", http.StatusBadRequest)
		}
		if len(t) > 0 {
			println("loading mock type: " + t)
		}
		*body, *contentType = b, t
		w.WriteHeader(http.StatusCreated)
		if len(*contentType) > 0 {
			w.Header().Set("Content-Type", *contentType)
			w.Write(*body)
		}
	}
}

func getMock(body *[]byte, contentType *string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		suffix := ""
		if len(*contentType) > 0 {
			suffix = " : " + *contentType
		}
		println(r.Method + " " + r.URL.Path + suffix)
		w.WriteHeader(200)
		if len(*contentType) > 0 {
			w.Header().Set("Content-Type", *contentType)
			w.Write(*body)
		}
	}
}
