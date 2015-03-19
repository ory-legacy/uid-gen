package main

import (
    "github.com/stretchr/testify/assert"
    "github.com/gorilla/mux"
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
    "strconv"
    "log"
    "errors")

func TestCreate(t *testing.T) {
    mockUidCreate(t)
}

func BenchmarkTestCreate(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = NewUid()
    }
}

func TestCreateUnique(t *testing.T) {
    id1 := mockUidCreate(t)
    id2 := mockUidCreate(t)
    assert.NotEqual(t, id1.Data.Uid, id2.Data.Uid)
    log.Printf("Created Uid %d and %d", id1.Data.Uid, id2.Data.Uid)
}

func TestJsonError(t *testing.T) {
    var ec ErrorCarrier

    err := errors.New("foo")
    b, err := jsonError(err)
    assert.Nil(t, err)

    err = json.Unmarshal(b, &ec)
    assert.Nil(t, err)
    assert.Equal(t, "foo", ec.Error.Message)
    assert.Equal(t, 500, ec.Error.Code)

}

func mockUidCreate(t *testing.T) DataCarrier {
    m := mux.NewRouter()
    recorder := httptest.NewRecorder()
    req, err := http.NewRequest("POST", "http://example.com/uids", nil)
    assert.Nil(t, err)
    m.HandleFunc("/uids", createHandler).Methods("POST")
    m.ServeHTTP(recorder, req)

    assert.Nil(t, err)
    assert.Equal(t, 200, recorder.Code)
    assert.NotEmpty(t, recorder.Body)

    var e DataCarrier
    decoder := json.NewDecoder(recorder.Body)
    err = decoder.Decode(&e)

    assert.Nil(t, err)
    assert.NotEmpty(t, e.Id)
    assert.NotEmpty(t, e.Data.Uid)
    assert.NotEmpty(t, e.Data.UidStr)
    assert.Equal(t, strconv.FormatUint(e.Data.Uid, 10), e.Data.UidStr)

    return e
}
