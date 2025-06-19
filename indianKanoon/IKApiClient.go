package indianKanoon

import (
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

func (IKA IKApiClient) SearchQuery(IKSearchData IKSearchData) string {
	req, err := http.NewRequest("POST", INDIAN_KANOON_BASE_URL+"search/", nil)
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

	return string(body)

}

func (IKA IKApiClient) DocumentFetch(IKSearchdoc IKSearchDocument) string {
	req, error := http.NewRequest("POST", INDIAN_KANOON_BASE_URL+"doc/", nil)
	req.Header.Add("Authorization", AUTH_TOKEN)
	req.Header.Add("Content-Type", "application/json")
	if error != nil {
		fmt.Println("Error creating request for fetching document")
	}
	u := req.URL.Query()
	u.Add("docid", IKSearchdoc.DocId)
	req.URL.RawQuery = u.Encode()
	res, error := IKA.Client.Do(req)
	if error != nil {
		fmt.Println("Error while getting response for fetching document")
	}
	body, error := io.ReadAll(res.Body)
	if error != nil {
		fmt.Println("Failed to read the response body")
	}
	return string(body)

}
