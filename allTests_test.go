package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"path"
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

func TestCli1(t *testing.T) {
	options = 2
	if options == 2 {
		name := path.Base(linkForTest)
		file, _ := os.Create(name)
		assert.Equal(t, name, file.Name())

	}
}
