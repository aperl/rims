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
	println(`
Use: rims <PORT>

Setup:
	To setup data to be mocked post it as the body to http://host:PORT/mock
Requests:
	All requests sent to the mock server will respond with the values the values
	posted to /mock.


Response Body:
	Whatever body that is posted to /mock with be the body of all other responses.
Content-Type:
	Whatever content-type header is present on the post to /mock will be the
	content-type on all other responses
Status Code:
	The status code of a request by default will be 200. If another status code is
	desired then the status code can be set as a query param on the post to /mock.

	EX: /mock?status=404

`)
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
