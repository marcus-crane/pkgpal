package search

import httpclient "github.com/ddliu/go-httpclient"

func Query(url string) []byte {
	res, err := httpclient.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := res.ReadAll()
	if err != nil {
		panic(err)
	}
	return body
}
