package proadclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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

func makeFilteredGetRequest(path string, code StatusCode, startDate, endDate time.Time) *http.Request {
	key := ""
	if strings.Contains(path, "project") {
		key = "order_date"
	} else {
		key = "from_date"
	}
	q := queryPair{
		key:   key,
		value: startDate.Format(proadTimeFormat) + "--" + endDate.Format(proadTimeFormat),
	}
	req := makeGETRequest(path, queryMap(code, q))
	return req
}

func queryMap(code StatusCode, pairs ...queryPair) map[string]string {
	sMap := map[string]string{}
	if code != StatusNone {
		sMap["status"] = code.String()
	}
	for _, qP := range pairs {
		sMap[qP.key] = qP.value
	}
	return sMap
}

func unmarshalResponse(request *http.Request, v interface{}) {
	resp := Client.Do(request)
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[proadclient/emplyees.go] error: %s\n", err.Error())
		return
	}
	defer resp.Body.Close()

	err = json.Unmarshal(respBodyBytes, &v)
	if err != nil {
		fmt.Printf("[proadclient/emplyees.go] error: %s\n", err.Error())
		return
	}
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
