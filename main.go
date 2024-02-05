package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

var link string

func main() {

	fmt.Println("input link :")
	_, err := fmt.Scanln(&link)
	if err != nil {
		log.Fatal(err)
	}
	lenght, tTime, error := download(link)
	if error != nil {
		fmt.Println(error)
	}
	error2 := speed(lenght, tTime)
	if error2 != nil {
		fmt.Println(error2)
	}

}
func download(link string) (lenght int64, totalTime time.Duration, error error) {
	res, err := http.Get(link)
	start := time.Now()
	fmt.Printf("Starting \n")
	defer res.Body.Close()

	file, err := os.Create(path.Base(link))
	if err != nil {
		//log.Fatal(err)
		return 0, 0, err
	}
	end := time.Now()

	size, err := io.Copy(file, res.Body)
	if err != nil {
		//log.Fatal(err)
		return 0, 0, err
	}

	sizeKb := size / 1000
	sizeMb := sizeKb / 1000

	if size < 1024 {
		fmt.Printf("Finished. file size: %db \n", size)
	} else if sizeKb < 1024 {
		fmt.Printf("Finished. file size: %dKb \n", sizeKb)
	} else if sizeKb >= 1024 {
		fmt.Printf("Finished. file size: %dMb \n", sizeMb)
	}
	totalTime = end.Sub(start)
	return res.ContentLength, totalTime, nil
}

func speed(resLength int64, time time.Duration) error {
	speed := float64(resLength) / time.Seconds() / 1024 / 1024
	if speed <= 0 {
		log.Fatal("speed can not be zero. check connection")
	}
	fmt.Printf("Approximate download Speed: %.3f Mbps\n"+"total time: %.2f second \n", speed, time.Seconds())
	return nil
}
