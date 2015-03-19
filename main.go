package main

import (
    "code.google.com/p/go-uuid/uuid"
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "github.com/ory-libs/env"
    "github.com/ory-libs/rand/numeric"
    "strconv"
)

type Data struct {
    Uid    uint64 `json:"uid"`
    UidStr string `json:"uidStr"`
}

type Error struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

type DataCarrier struct {
    ApiVersion string    `json:"apiVersion"`
    Id         uuid.UUID `json:"id"`
    Data       Data      `json:"data"`
}

type ErrorCarrier struct {
    ApiVersion string    `json:"apiVersion"`
    Id         uuid.UUID `json:"id"`
    Error      Error     `json:"error"`
}

const (
    ApiVersion = "1.0"
)

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
    e := DataCarrier{
        ApiVersion: ApiVersion,
        Id: uuid.NewRandom(),
        Data: Data{
            Uid: i,
            UidStr: strconv.FormatUint(i, 10),
        },
    }

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
    return json.Marshal(ErrorCarrier{
        ApiVersion: ApiVersion,
        Id: uuid.NewRandom(),
        Error: Error{
            Code: 500,
            Message: err.Error(),
        },
    })
}

// NewUid returns a cryptographically strong pseudo-random uint64
func NewUid() uint64 {
    return numeric.UInt64()
}
