package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonHandler1(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)

	assert.Equal(http.StatusOK, res.Code)
	assert.Nil(err)
	assert.Equal(2, len(list))
	assert.Equal("aaa", list[0].Name)
	assert.Equal("bbb", list[1].Name)
}

func TestJsonHandler2(t *testing.T) {
	assert := assert.New(t)

	mux := MakeWebHandler()
	var student Student

	// /students/1
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil)
	mux.ServeHTTP(res, req)
	err := json.NewDecoder(res.Body).Decode(&student)

	assert.Equal(http.StatusOK, res.Code)
	assert.Nil(err)
	assert.Equal("aaa", student.Name)

	// /students/2
	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/2", nil)
	mux.ServeHTTP(res, req)
	err = json.NewDecoder(res.Body).Decode(&student)

	assert.Equal(http.StatusOK, res.Code)
	assert.Nil(err)
	assert.Equal("bbb", student.Name)

	// /students/0
	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/0", nil)
	mux.ServeHTTP(res, req)
	err = json.NewDecoder(res.Body).Decode(&student)

	assert.Equal(http.StatusNotFound, res.Code)
	assert.NotNil(err)
}
