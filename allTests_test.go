package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

const linkForTest = "https://golikhanoom.ir/wp-content/uploads/2023/11/aks-gol-v-manzare-ziba-baraye-porofail.jpg"
const fileFotTest = "test.txt"

func TestMerged(t *testing.T) {
	res, _ := http.Get(linkForTest)
	someNumber := time.Duration(124)

	_, error := speed(res.ContentLength, someNumber)
	assert.Equal(t, nil, error)
}
