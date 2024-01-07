package api

import (
	"auto-jj/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type JVApi struct {
	client http.Client
}

func NetJVApi() *JVApi {
	return &JVApi{client: http.Client{}}

}

func (receiver *JVApi) GetUserJobPostRecommendations() (*model.GetUserJobPostRecommendationResp, error) {

	client := receiver.client
 
	req, err := http.NewRequest("GET", "https://candidateapi.jobvision.ir/api/v1/JobPost/Recommendation/Page?sortType=1", nil)
	if err != nil {
		return nil,err
	}

	req.Header.Set("authority", "candidateapi.jobvision.ir")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "en-US,en;q=0.9,fa-IR;q=0.8,fa;q=0.7")
	req.Header.Set("authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IkZEMjczNDJEQjU4RTFFREZEMzJBODQ4MUVGODU0REUwQzI5Q0I3MzdSUzI1NiIsIng1dCI6Il9TYzBMYldPSHRfVEtvU0I3NFZONE1LY3R6YyIsInR5cCI6ImF0K2p3dCJ9.eyJpc3MiOiJodHRwczovL2FjY291bnQuam9idmlzaW9uLmlyIiwibmJmIjoxNzA0NjYwMjE2LCJpYXQiOjE3MDQ2NjAyMTYsImV4cCI6MTcwNDc0NjYxNiwiYXVkIjpbIkpvYlZpc2lvbkFwaSIsIklkZW50aXR5U2VydmVyQXBpIl0sInNjb3BlIjpbIm9wZW5pZCIsInByb2ZpbGUiLCJKb2JWaXNpb25BcGkiLCJyb2xlcyIsIklkZW50aXR5U2VydmVyQXBpIiwib2ZmbGluZV9hY2Nlc3MiXSwiYW1yIjpbInB3ZCJdLCJjbGllbnRfaWQiOiJKb2JWaXNpb25Bbmd1bGFyQ2xpZW50Iiwic3ViIjoiMjQxNzUxNyIsImF1dGhfdGltZSI6MTY5NzA0OTY0NiwiaWRwIjoibG9jYWwiLCJlbWFpbCI6ImFiZnpsYWJmemxAZ21haWwuY29tIiwiaHR0cDovL3NjaGVtYXMubWljcm9zb2Z0LmNvbS93cy8yMDA4LzA2L2lkZW50aXR5L2NsYWltcy9yb2xlIjoiRW1wbG95ZWUiLCJSb2xlIjoiRW1wbG95ZWUiLCJGdWxsTmFtZSI6Itin2KjZiNin2YTZgdi22YQg2YLZh9ix2YXYp9mG24wiLCJMYXN0V29ya1RpdGxlIjoiYmFja2VuZCBkZXZlbG9wZXIiLCJTdWJtaXRUaW1lIjoiMjAyMy0wNy0wOFQwOTozMDoxNCIsInNpZCI6IjYxNzk0NzU1NEU0QjVCNTIwNEQ0Q0Q2NEQxQjYyMzdBIiwianRpIjoiQjAzMEU5RkNCRjhFQjk2OTNGOTQ5RkJGRkNBMjFEQTcifQ.siZ1Bk1Bzi2vrizZZbvsMOUX9kj6xQlzMyiiuFPR1x_2L70Rs6atxUdwsPFdtmCt6QrNB6pOQeswtR_cmkXy0Oa2SG9Qec-jL17B2BVEJvP4_e2ZougGgG_A7YGaKSyiCvyG_WyCYcEQHf8p7fjp5lePFWBYClb-5zwjz4DAgCtU2DygaLho1v6PfLGfLKS-pLo1w7fEYNMSsMvBDICDEoFA7aSmk7HaRIcg3cCO4Z4qn6eQjSulL80u4Y82NMCUfxyrhhaEQqu0xT77c49gkYqq38DIz2y-yVMeNVgym-ZAwFkC0iI0tzB-rN1GgFofyr1SVJlyAG35RJb7pn3I4g")
	req.Header.Set("ngsw-bypass", "true")
	req.Header.Set("origin", "https://jobvision.ir")
	req.Header.Set("page-route", "https://jobvision.ir/recommended-jobs")
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
		return nil, err

	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	} 

	var jjj model.GetUserJobPostRecommendationResp

	err2 := json.Unmarshal(bodyText, &jjj)

	if err2 != nil {
		return nil, err2
	}


	return &jjj, err

}

func (receiver *JVApi) GetDeveloperJobs(currentPage int) (*model.JobsResp, error) {


	client := receiver.client
 
	pl:=fmt.Sprintf(`{
		"pageSize": 30,
		"requestedPage": %d,
		"sortBy": 1,
		"jobCategoryUrlTitle": "developer",
		"searchId": null
	  }`,currentPage)
	var data = strings.NewReader(pl)

	req, err := http.NewRequest("POST", "https://candidateapi.jobvision.ir/api/v1/JobPost/List", data)
	if err != nil {
		return nil,err

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
		return nil,err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,err

	} 

	var j model.JobsResp

	err2 := json.Unmarshal(bodyText, &j)

	if err2 != nil {
		fmt.Printf("j.Data.JobPosts: %v\n", j.Data.JobPosts)
		return nil,err2

	}

	return &j, err2

}


func (receiver *JVApi) Apply(jobid int) (*ApplyJVResp, error) {

	client := receiver.client

	reqstr := fmt.Sprintf(`{"jobPostId":%d,"referHireCode":"","userJobPostMatchScore":66}`, jobid)

	var data = strings.NewReader(reqstr)
	req, err := http.NewRequest("POST", "https://candidateapi.jobvision.ir/api/v1/Application/Apply", data)
	if err != nil {

		return nil, err

	}
	req.Header.Set("authority", "candidateapi.jobvision.ir")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "en-US,en;q=0.9,fa-IR;q=0.8,fa;q=0.7")
	req.Header.Set("authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IkZEMjczNDJEQjU4RTFFREZEMzJBODQ4MUVGODU0REUwQzI5Q0I3MzdSUzI1NiIsIng1dCI6Il9TYzBMYldPSHRfVEtvU0I3NFZONE1LY3R6YyIsInR5cCI6ImF0K2p3dCJ9.eyJpc3MiOiJodHRwczovL2FjY291bnQuam9idmlzaW9uLmlyIiwibmJmIjoxNzA0NjYwMjE2LCJpYXQiOjE3MDQ2NjAyMTYsImV4cCI6MTcwNDc0NjYxNiwiYXVkIjpbIkpvYlZpc2lvbkFwaSIsIklkZW50aXR5U2VydmVyQXBpIl0sInNjb3BlIjpbIm9wZW5pZCIsInByb2ZpbGUiLCJKb2JWaXNpb25BcGkiLCJyb2xlcyIsIklkZW50aXR5U2VydmVyQXBpIiwib2ZmbGluZV9hY2Nlc3MiXSwiYW1yIjpbInB3ZCJdLCJjbGllbnRfaWQiOiJKb2JWaXNpb25Bbmd1bGFyQ2xpZW50Iiwic3ViIjoiMjQxNzUxNyIsImF1dGhfdGltZSI6MTY5NzA0OTY0NiwiaWRwIjoibG9jYWwiLCJlbWFpbCI6ImFiZnpsYWJmemxAZ21haWwuY29tIiwiaHR0cDovL3NjaGVtYXMubWljcm9zb2Z0LmNvbS93cy8yMDA4LzA2L2lkZW50aXR5L2NsYWltcy9yb2xlIjoiRW1wbG95ZWUiLCJSb2xlIjoiRW1wbG95ZWUiLCJGdWxsTmFtZSI6Itin2KjZiNin2YTZgdi22YQg2YLZh9ix2YXYp9mG24wiLCJMYXN0V29ya1RpdGxlIjoiYmFja2VuZCBkZXZlbG9wZXIiLCJTdWJtaXRUaW1lIjoiMjAyMy0wNy0wOFQwOTozMDoxNCIsInNpZCI6IjYxNzk0NzU1NEU0QjVCNTIwNEQ0Q0Q2NEQxQjYyMzdBIiwianRpIjoiQjAzMEU5RkNCRjhFQjk2OTNGOTQ5RkJGRkNBMjFEQTcifQ.siZ1Bk1Bzi2vrizZZbvsMOUX9kj6xQlzMyiiuFPR1x_2L70Rs6atxUdwsPFdtmCt6QrNB6pOQeswtR_cmkXy0Oa2SG9Qec-jL17B2BVEJvP4_e2ZougGgG_A7YGaKSyiCvyG_WyCYcEQHf8p7fjp5lePFWBYClb-5zwjz4DAgCtU2DygaLho1v6PfLGfLKS-pLo1w7fEYNMSsMvBDICDEoFA7aSmk7HaRIcg3cCO4Z4qn6eQjSulL80u4Y82NMCUfxyrhhaEQqu0xT77c49gkYqq38DIz2y-yVMeNVgym-ZAwFkC0iI0tzB-rN1GgFofyr1SVJlyAG35RJb7pn3I4g")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("ngsw-bypass", "true")
	req.Header.Set("origin", "https://jobvision.ir")
	req.Header.Set("page-route", "https://jobvision.ir/jobs/626269/%D8%A7%D8%B3%D8%AA%D8%AE%D8%AF%D8%A7%D9%85-%DA%A9%D8%A7%D8%B1%D8%B4%D9%86%D8%A7%D8%B3-%D8%A7%D8%B1%D8%B4%D8%AF-%D9%86%D8%B1%D9%85-%D8%A7%D9%81%D8%B2%D8%A7%D8%B1---%D8%A2%D9%82%D8%A7")
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
		return nil, err

	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	} 

	var jjj ApplyJVResp

	err2 := json.Unmarshal(bodyText, &jjj)

	if err2 != nil {
		return nil, err2
	}


	return &jjj, err

}
