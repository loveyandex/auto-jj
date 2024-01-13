package main

import "fmt"

func maimn() {
	intChan := make(chan int)

	for {

		for i := 0; i < 100000000; i++ {
			go func(i int) {
				Send(i)
				intChan <- i
			}(i)

		}
		receivedValue := <-intChan

		if receivedValue == 100 {

			fmt.Printf("receivedValue: %v\n", receivedValue)

		}

	}
}

func Send(i int) {

}
