package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

const forTest = "https://golikhanoom.ir/wp-content/uploads/2023/11/aks-gol-v-manzare-ziba-baraye-porofail.jpg"

func TestFirstPart(t *testing.T) {
	_, _, err := download(forTest)
	assert.Equal(t, nil, err)
}

func TestSecondPart(t *testing.T) {
	res, _ := http.Get(forTest)
	someNumber := time.Duration(124)

	err := speed(res.ContentLength, someNumber)
	assert.Equal(t, nil, err)
}
