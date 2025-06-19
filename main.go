package main

import (
	"fmt"
	ik "lawAgent/indianKanoon"
)

func main() {
	searchData := ik.IKSearchData{
		FormInput: "ishrat jahan encounter",
		DocTypes:  ik.DOCTYPE,
	}

	ikSearchClient := ik.GetIKApiClient()
	response := ikSearchClient.SearchQuery(searchData)
	fmt.Println(response)
}
