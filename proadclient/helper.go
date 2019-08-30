package proadclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func makeGETRequest(relURL string, query map[string]string) *http.Request {
	req, err := http.NewRequest("GET", BaseURL+relURL, http.NoBody)
	if err != nil {
		fmt.Printf("[proadclient/emplyees.go] error: %s\n", err.Error())
		return &http.Request{}
	}
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	return req
}

// PrettyPrintResponse prints the json body of a response as human readable data
func PrettyPrintResponse(resp *http.Response) {
	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, bb, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("responsebody from requesturl %s\n", resp.Request.URL)
	fmt.Println(string(prettyJSON.Bytes()))
}
