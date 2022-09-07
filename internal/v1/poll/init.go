package poll

import (
	"os"
	"time"

	"github.com/gmarcha/notion-trigger/internal/v1/log"
	"github.com/jomei/notionapi"
)

var LastPollTime time.Time

func init() {
	LastPollTime = time.Now().Add(-time.Minute)
}

var (
	campusDbId notionapi.DatabaseID
	eventDbId  notionapi.DatabaseID
)

func init() {
	campusDbId = notionapi.DatabaseID(os.Getenv("NOTION_CAMPUS_DB_ID"))
	if campusDbId == "" {
		log.Logger.Warn("Environment variable NOTION_CAMPUS_DB_ID is empty")
	}
	eventDbId = notionapi.DatabaseID(os.Getenv("NOTION_EVENT_DB_ID"))
	if eventDbId == "" {
		log.Logger.Warn("Environment variable NOTION_EVENT_DB_ID is empty")
	}
}
