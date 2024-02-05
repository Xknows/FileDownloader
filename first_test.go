package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstPart(t *testing.T) {
	forTest := "https://golikhanoom.ir/wp-content/uploads/2023/11/aks-gol-v-manzare-ziba-baraye-porofail.jpg"
	err := download(forTest)
	assert.Equal(t, nil, err)
}
