package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

var link string
var filePath string
var options int
var wg sync.WaitGroup

func main() {

	fmt.Printf("choose:\n 1: enter direct link\n 2: enter file.txt of links \nenter your choose: ...\n")
	_, err := fmt.Scan(&options)
	if err != nil {
		log.Fatal(err)
	}

	switch options {
	case 1:
		fmt.Print("enter your link: ....\n")
		_, err := fmt.Scan(&link)
		if err != nil {
			log.Fatal(err)
		}
		go download(link, &wg)
		wg.Add(1)
		wg.Wait()
	case 2:
		fmt.Printf("enter path: ....\n")
		_, err := fmt.Scan(&filePath)
		if err != nil {
			fmt.Println(err)
		}
		reader, err := os.Open(filePath)
		if err != nil {
			fmt.Println(err)
		}
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			go download(scanner.Text(), &wg)
			wg.Add(1)
		}
		wg.Wait()
		fmt.Println("***************\n" + "all done")

	default:
		fmt.Print("non valid options!")
	}

}

func download(link string, wg *sync.WaitGroup) {

	res, err := http.Get(link)
	start := time.Now()
	fmt.Println("****************")
	fmt.Printf("Starting \n")
	defer res.Body.Close()

	file, err := os.Create(path.Base(link))
	end := time.Now()

	size, err := io.Copy(file, res.Body)
	if err != nil {
		log.Fatal(err)
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
	totalTime := end.Sub(start)
	resLength := res.ContentLength
	_, err = speed(resLength, totalTime)
	if err != nil {
		log.Fatal(err)
	}
	defer wg.Done()
}
func speed(resLength int64, time time.Duration) (speed float64, err error) {
	calculate := float64(resLength) / time.Seconds() / 1024 / 1024
	if speed < 1 {
		log.Fatal("speed can not be zero. check connection\n")
	}
	fmt.Printf("Approximate download Speed: %.3f Mbps\n"+"total time: %.2f second \n", calculate, time.Seconds())
	return calculate, nil
}