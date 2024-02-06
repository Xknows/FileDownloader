package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
	"time"
)

const linkForTest = "https://golikhanoom.ir/wp-content/uploads/2023/11/aks-gol-v-manzare-ziba-baraye-porofail.jpg"
const fileFotTest = "test.txt"

func TestFirstPart(t *testing.T) {
	_, _, err := download(linkForTest)
	assert.Equal(t, nil, err)
}

func TestSecondPart(t *testing.T) {
	res, _ := http.Get(linkForTest)
	someNumber := time.Duration(124)

	err := speed(res.ContentLength, someNumber)
	assert.Equal(t, nil, err)
}

func TestCli1(t *testing.T) {
	options = 2
	if options == 1 {
		_, _, err := download(linkForTest)
		assert.Equal(t, nil, err)
	} else if options == 2 {
		reader, _ := os.Open(fileFotTest)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			_, _, err := download(scanner.Text())
			assert.Equal(t, nil, err)
		}
	}
}
