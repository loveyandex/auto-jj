package api

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {

	api:=NetJVApi()


	gujprr, err := api.GetUserJobPostRecommendations()


	fmt.Printf("err: %v\n", err)


	for _, jp := range gujprr.Data {

		fmt.Printf("jp.Title: %v\n", jp.Title)
		
	}

}