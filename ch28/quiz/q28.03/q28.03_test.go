package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi1(t *testing.T) {
	assert := assert.New(t)

	n, err := Atoi("0")
	assert.Equal(0, n, "Atoi(\"0\") should be 0")
	assert.Equal(err, nil, "cannont conver to int")

	n, err = Atoi("1")
	assert.Equal(1, n, "Atoi(\"1\") should be 1")
	assert.Equal(err, nil, "cannont conver to int")

	n, err = Atoi("100")
	assert.Equal(100, n, "Atoi(\"100\") should be 100")
	assert.Equal(err, nil, "cannont conver to int")

	n, err = Atoi("2022")
	assert.Equal(2022, n, "Atoi(\"2022\") should be 2022")
	assert.Equal(err, nil, "cannont conver to int")

	n, err = Atoi("12345")
	assert.Equal(12345, n, "Atoi(\"12345\") should be 12345")
	assert.Equal(err, nil, "cannont conver to int")
}

func TestAtoi2(t *testing.T) {
	assert := assert.New(t)

	n, err := Atoi("  1234")
	assert.Equal(1234, n, "Atoi(\"1234\") should be 1234")
	assert.Equal(err, nil, "cannont conver to int")

	n, err = Atoi("5678  ")
	assert.Equal(5678, n, "Atoi(\"1234\") should be 5678")
	assert.Equal(err, nil, "cannont conver to int")

	n, err = Atoi("-10")
	assert.Equal(-10, n, "Atoi(\"-10\") should be -10")
	assert.Equal(err, nil, "cannont conver to int")

	_, err = Atoi("asdf1234")
	assert.NotEqual(err, nil)
}
