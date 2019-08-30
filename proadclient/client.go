package proadclient

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

// apikey is needed for initializing the proad client
const apikey = "daa80e5383c27ce5c0c5434bbc699618960c67d0e26002e53a71d961c7091844d0ffbac4d643c4e2fa3eb98ee8c5bd2e4b817b4e4ac94224d4e590d687d0fd65"

// BaseURL defines the baseURL to the proad server
const BaseURL = "https://192.168.0.15/api/v5/"

// Client is a ProadClient singleton
var Client *ProadClient

// ProadClient is a http.Client without ssl verification
type ProadClient struct {
	httpClient *http.Client
	apiKey     string
}

func init() {
	Client = newProadClient(apikey)
}

// newProadClient returns a new ProadClient
func newProadClient(apiKey string) *ProadClient {
	return &ProadClient{
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		apiKey: apiKey,
	}
}

// modifyRequest modifies a request, so that Proad can handle it
func modifyRequest(req *http.Request) {
	req.Header.Add("apikey", Client.apiKey)
}

// Do functions the same as http.Client.Do but modifies the header of the request
func (pc *ProadClient) Do(request *http.Request) *http.Response {
	modifyRequest(request)
	resp, err := pc.httpClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return &http.Response{}
	}
	return resp
}

// CheckConnection checks wether the server is reachable and if the provided key is valid
func (pc *ProadClient) CheckConnection() {
	request, err := http.NewRequest("GET", "https://192.168.0.15/api/test/validate_key", http.NoBody)
	response := pc.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	respBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	fmt.Println(string(respBytes))
}
