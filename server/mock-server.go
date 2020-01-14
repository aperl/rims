package server

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

//New create a new mock server
func New() *MockServer {
	return &MockServer{
		Status: 200,
	}
}

// MockServer the server that spits back what it gets
type MockServer struct {
	Body        []byte
	Status      int
	ContentType string
}

//LoadMock load a response into the mock server
func (mock *MockServer) LoadMock(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	query := r.URL.Query()
	statusStr := query.Get("status")
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		status = 200
	}
	mock.Status = status

	t := r.Header.Get("Content-Type")
	if len(b) > 0 && len(t) == 0 {
		http.Error(w, "If a body is present then there must be a Content-Type", http.StatusBadRequest)
	}
	if len(t) > 0 {
		println("loading mock type: " + t)
	}
	mock.Body, mock.ContentType = b, t
	w.WriteHeader(http.StatusCreated)
	if len(mock.ContentType) > 0 {
		w.Header().Set("Content-Type", mock.ContentType)
		w.Write(mock.Body)
	}
}

//GetMock the response loaded into the mock server
func (mock *MockServer) GetMock(w http.ResponseWriter, r *http.Request) {
	suffix := ""
	if len(mock.ContentType) > 0 {
		suffix = " : " + mock.ContentType
	}
	println(r.Method + " " + r.URL.Path + suffix)
	w.WriteHeader(mock.Status)
	if len(mock.ContentType) > 0 {
		w.Header().Set("Content-Type", mock.ContentType)
		w.Write(mock.Body)
	}
}
