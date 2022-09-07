package client

import (
	"os"

	"github.com/gmarcha/notion-trigger/internal/v1/log"
)

var notionToken string

func init() {
	notionToken = os.Getenv("NOTION_TOKEN")
	if notionToken == "" {
		log.Logger.Warn("Environment variable NOTION_TOKEN is empty")
	}
}
