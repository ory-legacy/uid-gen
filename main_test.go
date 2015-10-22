package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

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
	assert.NotEqual(t, id1.Uid, id2.Uid)
	log.Printf("Created Uid %d and %d", id1.Uid, id2.Uid)
}

func TestJsonError(t *testing.T) {
	var ec errorResponse

	err := errors.New("foo")
	b, err := jsonError(err)
	assert.Nil(t, err)

	err = json.Unmarshal(b, &ec)
	assert.Nil(t, err)
	assert.Equal(t, "foo", ec.Message)
	assert.Equal(t, 500, ec.Code)

}

func mockUidCreate(t *testing.T) successResponse {
	m := mux.NewRouter()
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "http://example.com/uids", nil)
	assert.Nil(t, err)
	m.HandleFunc("/uids", createHandler).Methods("POST")
	m.ServeHTTP(recorder, req)

	assert.Nil(t, err)
	assert.Equal(t, 200, recorder.Code)
	assert.NotEmpty(t, recorder.Body)

	var e successResponse
	decoder := json.NewDecoder(recorder.Body)
	err = decoder.Decode(&e)

	assert.Nil(t, err)
	assert.NotEmpty(t, e.Uid)
	assert.NotEmpty(t, e.UidStr)
	assert.Equal(t, strconv.FormatUint(e.Uid, 10), e.UidStr)

	return e
}
