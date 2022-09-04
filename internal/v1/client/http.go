package client

import (
	"net/http"
)

type NotionTransport struct{}

func (t *NotionTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+notionToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	return http.DefaultTransport.RoundTrip(req)
}

func Http() *http.Client {
	return &http.Client{Transport: &NotionTransport{}}
}
