package main

import (
	"auto-jj/api"
	"auto-jj/service"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {

	api := api.NetJVApi()
	js := service.NewJVPostSrv()

	for {

		for i := 0; i < 100; i++ {

			jr, err := api.GetDeveloperJobs(i)

			if err != nil {
				log.Fatal(err)

			}

			if !jr.IsSuccess {

			}

			for _, jp := range jr.Data.JobPosts {

				fmt.Printf("jp.Title: %v\n", jp.Title)
				if strings.Contains(strings.ToLower(jp.Title), "laravel") {
					continue

				}
				if strings.Contains(strings.ToLower(jp.Title), "php") {
					continue

				}

				jpid := jp.ID

				jp2, err4 := js.GetByjobIf(jpid)

				if err4 != nil {
					fmt.Printf("err4: %v\n", err4)
					if err4.Error() != "mongo: no documents in result" {
						continue

					}
				}
				fmt.Printf("jp2.ID: %v\n", jp2.ID)
				if jp2.ID == jpid {
					continue

				}

				aj, err3 := api.Apply(jpid)

				if err3 != nil {
					fmt.Printf("err3: %v\n", err3)
				}

				fmt.Printf("aj: %v\n", aj)

				if aj.IsSuccess {
					i2, err5 := js.CmnSrv.Create(&jp)

					if err5 != nil {

					}
					fmt.Printf("i2: %v\n", i2)

				}
				time.Sleep(5*time.Second)

			}

		}

		time.Sleep(24 * time.Hour)

	}

}
