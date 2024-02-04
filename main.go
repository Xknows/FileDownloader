package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

var link string

func main() {

	fmt.Println("input link :")
	_, err := fmt.Scanln(&link)
	if err != nil {
		log.Fatal(err)
	}
	download(link)
}
func download(link string) {
	res, err := http.Get(link)
	fmt.Printf("Starting \n")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	
	file, err := os.Create(path.Base(link))
	if err != nil {
		log.Fatal(err)
	}
	
	size, err := io.Copy(file, res.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	sizeKb := size / 1000
	sizeMb := sizeKb / 1000

	if size < 1024 {
		fmt.Printf("Finished. size: %db", size)
	} else if sizeKb < 1024 {
		fmt.Printf("Finished. size: %dKb", sizeKb)
	} else if sizeKb >= 1024 {
		fmt.Printf("Finished. size: %dMb", sizeMb)
	}

}
