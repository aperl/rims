package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/aperl/rims/server"
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

	server := server.New()

	r := mux.NewRouter()
	r.HandleFunc("/mock", server.LoadMock).Methods(http.MethodPost)
	r.PathPrefix("/").HandlerFunc(server.GetMock)
	println("Listening on port ", args[0])

	if err := http.ListenAndServe(":"+args[0], r); err != nil {
		println("Error: " + err.Error())
	}
	println("Closeing")
}
