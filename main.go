package main

import (
	"auto-jj/api"
	"auto-jj/model" 
	"auto-jj/service"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {

	api := api.NetJVApi()
	js := service.NewJVPostSrv()

	client := &http.Client{}
	var data = strings.NewReader(`{"pageSize":30,"requestedPage":1,"sortBy":1,"searchId":null}`)
	req, err := http.NewRequest("POST", "https://candidateapi.jobvision.ir/api/v1/JobPost/List", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "candidateapi.jobvision.ir")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "en-US,en;q=0.9,fa-IR;q=0.8,fa;q=0.7")
	req.Header.Set("authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IkZEMjczNDJEQjU4RTFFREZEMzJBODQ4MUVGODU0REUwQzI5Q0I3MzdSUzI1NiIsIng1dCI6Il9TYzBMYldPSHRfVEtvU0I3NFZONE1LY3R6YyIsInR5cCI6ImF0K2p3dCJ9.eyJpc3MiOiJodHRwczovL2FjY291bnQuam9idmlzaW9uLmlyIiwibmJmIjoxNzA0NjYwMjE2LCJpYXQiOjE3MDQ2NjAyMTYsImV4cCI6MTcwNDc0NjYxNiwiYXVkIjpbIkpvYlZpc2lvbkFwaSIsIklkZW50aXR5U2VydmVyQXBpIl0sInNjb3BlIjpbIm9wZW5pZCIsInByb2ZpbGUiLCJKb2JWaXNpb25BcGkiLCJyb2xlcyIsIklkZW50aXR5U2VydmVyQXBpIiwib2ZmbGluZV9hY2Nlc3MiXSwiYW1yIjpbInB3ZCJdLCJjbGllbnRfaWQiOiJKb2JWaXNpb25Bbmd1bGFyQ2xpZW50Iiwic3ViIjoiMjQxNzUxNyIsImF1dGhfdGltZSI6MTY5NzA0OTY0NiwiaWRwIjoibG9jYWwiLCJlbWFpbCI6ImFiZnpsYWJmemxAZ21haWwuY29tIiwiaHR0cDovL3NjaGVtYXMubWljcm9zb2Z0LmNvbS93cy8yMDA4LzA2L2lkZW50aXR5L2NsYWltcy9yb2xlIjoiRW1wbG95ZWUiLCJSb2xlIjoiRW1wbG95ZWUiLCJGdWxsTmFtZSI6Itin2KjZiNin2YTZgdi22YQg2YLZh9ix2YXYp9mG24wiLCJMYXN0V29ya1RpdGxlIjoiYmFja2VuZCBkZXZlbG9wZXIiLCJTdWJtaXRUaW1lIjoiMjAyMy0wNy0wOFQwOTozMDoxNCIsInNpZCI6IjYxNzk0NzU1NEU0QjVCNTIwNEQ0Q0Q2NEQxQjYyMzdBIiwianRpIjoiQjAzMEU5RkNCRjhFQjk2OTNGOTQ5RkJGRkNBMjFEQTcifQ.siZ1Bk1Bzi2vrizZZbvsMOUX9kj6xQlzMyiiuFPR1x_2L70Rs6atxUdwsPFdtmCt6QrNB6pOQeswtR_cmkXy0Oa2SG9Qec-jL17B2BVEJvP4_e2ZougGgG_A7YGaKSyiCvyG_WyCYcEQHf8p7fjp5lePFWBYClb-5zwjz4DAgCtU2DygaLho1v6PfLGfLKS-pLo1w7fEYNMSsMvBDICDEoFA7aSmk7HaRIcg3cCO4Z4qn6eQjSulL80u4Y82NMCUfxyrhhaEQqu0xT77c49gkYqq38DIz2y-yVMeNVgym-ZAwFkC0iI0tzB-rN1GgFofyr1SVJlyAG35RJb7pn3I4g")
	req.Header.Set("clientid", "20308581")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("ngsw-bypass", "true")
	req.Header.Set("origin", "https://jobvision.ir")
	req.Header.Set("page-route", "https://jobvision.ir/jobs")
	req.Header.Set("referer", "https://jobvision.ir/")
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("web-app-version", "17.0.4")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

	var j model.JobsResp

	err2 := json.Unmarshal(bodyText, &j)

	if err2 != nil {
		fmt.Printf("j.Data.JobPosts: %v\n", j.Data.JobPosts)

	}
	for i, jp := range j.Data.JobPosts {

		
		
		
		
		fmt.Printf("i: %v\n", i)
		
		jpid := jp.ID
		
		jp2, err4 := js.GetByjobIf(jpid)

		if err4 != nil {
			fmt.Printf("err4: %v\n", err4)
		}
		fmt.Printf("jp2.ID: %v\n", jp2.ID)


		aj, err3 := api.Apply(jpid)

		if err3 != nil {
			fmt.Printf("err3: %v\n", err3)
		}

		fmt.Printf("aj.IsSuccess: %v\n", aj.IsSuccess)

	}
}
