package client

import (
	"github.com/jomei/notionapi"
)

func Notion() *notionapi.Client {
	return notionapi.NewClient(notionapi.Token(notionToken))
}
