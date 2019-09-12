package basecampclient

import (
	"fmt"
	"io"
	"net/http"
	urlutil "net/url"
	"path"
)

// Do makes a request to basecamp with the specified url
func (c *Basecampclient) Do(method, url string, body io.Reader) (*http.Response, error) {
	fmt.Printf("[basecampclient/get.go/Do] method: %s, url: %s\n", method, url)
	requestURL, err := urlutil.Parse(url)
	if err != nil {
		fmt.Println("could not parse url")
		return nil, err
	}
	if !requestURL.IsAbs() {
		fmt.Printf("%s is not absolute\n", url)
		requestURL = c.baseURL()
		requestURL.Path = path.Join(requestURL.Path, url)
	}
	fmt.Println(requestURL.String())
	req, err := http.NewRequest(method, requestURL.String(), body)
	if err != nil {
		fmt.Println("error making request " + err.Error())
		return nil, err
	}
	c.addHeader(req)
	resp, err := c.httpclient.Do(req)
	if err != nil {
		fmt.Println("error making request " + err.Error())
		return nil, err
	}
	return resp, nil
}
