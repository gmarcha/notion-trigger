package poll

import (
	"context"
	"log"
	"os"

	"github.com/gmarcha/notion-trigger/internal/v1/goswagger/client"
	"github.com/gmarcha/notion-trigger/internal/v1/goswagger/client/task"
	"github.com/jomei/notionapi"
)

var (
	campusDbId notionapi.DatabaseID
)

func init() {
	campusDbId = notionapi.DatabaseID(os.Getenv("NOTION_CAMPUS_DB_ID"))
}

func Campus(notion *notionapi.Client, api *client.Notion) error {

	ctx := context.Background()
	resp, err := notion.Database.Query(ctx, campusDbId, nil)
	if err != nil {
		return err
	}
	for _, page := range resp.Results {
		if page.CreatedTime.After(LastPollTime) {
			_, err := api.Task.CampusCreate(&task.CampusCreateParams{ID: page.ID.String(), Context: ctx})
			if err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}
