package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestJsonHandler3(t *testing.T) {
	assert := assert.New(t)

	mux := MakeWebHandler()
	var student Student

	// create new student
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students", strings.NewReader(`{"Id":0,"Name":"ccc","Age":15,"Score":78}`))
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	// check new student
	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/3", nil)
	mux.ServeHTTP(res, req)
	err := json.NewDecoder(res.Body).Decode(&student)

	assert.Equal(http.StatusOK, res.Code)
	assert.Nil(err)
	assert.Equal("ccc", student.Name)

	// without data
	res = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/students", nil)
	mux.ServeHTTP(res, req)
	err = json.NewDecoder(res.Body).Decode(&student)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.NotNil(err)
}

func TestJsonHandler4(t *testing.T) {
	assert := assert.New(t)

	mux := MakeWebHandler()
	var student Student
	var list []Student

	// delete /students/1
	res := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/students/1", nil)
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	// check deleted /students/1
	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/1", nil)
	mux.ServeHTTP(res, req)
	err := json.NewDecoder(res.Body).Decode(&student)

	assert.Equal(http.StatusNotFound, res.Code)
	assert.NotNil(err)

	// check student list
	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students", nil)
	mux.ServeHTTP(res, req)
	err = json.NewDecoder(res.Body).Decode(&list)

	assert.Nil(err)
	assert.Equal(1, len(list))
	assert.Equal("bbb", list[0].Name)

	// create new student for incremental number
	res = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/students", strings.NewReader(`{"Id":0,"Name":"ccc","Age":15,"Score":78}`))
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	// check incremental number
	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/3", nil)
	mux.ServeHTTP(res, req)
	err = json.NewDecoder(res.Body).Decode(&student)

	assert.Equal(http.StatusOK, res.Code)
	assert.Nil(err)
	assert.Equal(3, student.Id)
	assert.Equal("ccc", student.Name)
}
