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

	// searchDoc := ik.IKSearchDocument{
	// 	DocId: "653797",
	// }

	ikSearchClient := ik.GetIKApiClient()
	response := ikSearchClient.SearchQuery(searchData)
	//response := ikSearchClient.DocumentFetch(searchDoc)
	// response := ikSearchClient.DocumentFetchOriginal(searchDoc)

	fmt.Println(len(response.Docs))
}
