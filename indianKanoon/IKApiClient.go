package indianKanoon

import "net/http"

type IKApiClient struct {
	Client *http.Client
}

func GetIKApiClient() *IKApiClient {
	return &IKApiClient{
		Client: &http.Client{},
	}
}
