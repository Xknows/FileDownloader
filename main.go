package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

var link string

func main() {

	fmt.Println("input link :")
	fmt.Scanln(&link)
	download(link)
}
func download(link string) {
	res, _ := http.Get(link)
	file, _ := os.Create(path.Base(link))
	size, err := io.Copy(file, res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%dkb \n", size/1000)
	fmt.Printf("finished")
	defer res.Body.Close()
}
