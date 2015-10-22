package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ory-am/common/env"
	"github.com/ory-am/common/rand/numeric"
	"log"
	"net/http"
	"strconv"
)

type successResponse struct {
	Uid    uint64 `json:"uid"`
	UidStr string `json:"uidStr"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	host := env.Getenv("HOST", "")
	port := env.Getenv("PORT", "80")
	listen := fmt.Sprintf("%s:%s", host, port)
	r := mux.NewRouter()
	r.HandleFunc("/uids", createHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(listen, r))
}

// createHandler is a HTTP handler for returning random uint64s
func createHandler(w http.ResponseWriter, r *http.Request) {
	i := NewUid()
	e := successResponse{i, strconv.FormatUint(i, 10)}
	j, err := json.Marshal(e)
	if err != nil {
		je, fatal := jsonError(err)
		if fatal != nil {
			log.Fatal(fatal)
		}
		w.Write(je)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// jsonError is creates an ErrorCarrier and returns a byte slice
func jsonError(err error) ([]byte, error) {
	return json.Marshal(errorResponse{500, err.Error()})
}

// NewUid returns a cryptographically strong pseudo-random uint64
func NewUid() uint64 {
	return numeric.UInt64()
}
