package indianKanoon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IKApiClient struct {
	Client *http.Client
}

func GetIKApiClient() *IKApiClient {
	IKclient := IKApiClient{}
	IKclient.Client = &http.Client{}
	return &IKclient
}

func (IKA IKApiClient) SearchQuery(IKSearchData IKSearchData) IKSearchResponse {
	req, err := http.NewRequest("POST", INDIAN_KANOON_BASE_URL+"search/", nil)
	searchDocumentResponse := IKSearchResponse{}
	if err != nil {
		fmt.Println("Error creating request for IK search query")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Token"+" "+AUTH_TOKEN)
	q := req.URL.Query()
	q.Add("formInput", IKSearchData.FormInput)
	q.Add("doctypes", IKSearchData.DocTypes)
	req.URL.RawQuery = q.Encode()

	res, err := IKA.Client.Do(req)
	if err != nil {
		fmt.Println("Error sending request")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read the response body")
	}

	err = json.Unmarshal(body, &searchDocumentResponse)
	if err != nil {
		fmt.Println("Error parsing the response")
	}

	return searchDocumentResponse

}

func (IKD IKApiClient) DocumentFetch(IKSearchdoc IKSearchDocument) IKFetchDocumentType {
	req, err := http.NewRequest("POST", INDIAN_KANOON_BASE_URL+"doc/"+IKSearchdoc.DocId+"/", nil)
	req.Header.Add("Authorization", "Token"+" "+AUTH_TOKEN)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error creating request for fetching document")
	}

	res, err := IKD.Client.Do(req)
	if err != nil {
		fmt.Println("Error while getting response for fetching document")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read the response body")
	}
	documentFetchResponse := IKFetchDocumentType{}
	err = json.Unmarshal(body, &documentFetchResponse)
	if err != nil {
		fmt.Println("Error parsing response")
	}
	return documentFetchResponse
}

func (IKOD IKApiClient) DocumentFetchOriginal(IKSearchdoc IKSearchDocument) string {
	req, err := http.NewRequest("POST", INDIAN_KANOON_BASE_URL+"origdoc/"+IKSearchdoc.DocId+"/", nil)
	req.Header.Add("Authorization", "Token"+" "+AUTH_TOKEN)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error creating request for fetching document")
	}

	res, err := IKOD.Client.Do(req)
	if err != nil {
		fmt.Println("Error while getting response for fetching original document")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read the response body")
	}
	return string(body)

}
