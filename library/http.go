package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// resp, _ := http.Get("http://example.com")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	base, err := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)
	fmt.Println(base, err)

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match", `W/"wyzzz"`)
	q := req.URL.Query()
	q.Add("c", "3")
	req.URL.RawQuery = q.Encode()
	fmt.Println(q)

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
