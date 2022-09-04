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
	eventDbId notionapi.DatabaseID
)

func init() {
	eventDbId = notionapi.DatabaseID(os.Getenv("NOTION_EVENT_DB_ID"))
}

func Pool(notion *notionapi.Client, api *client.Notion) error {

	ctx := context.Background()
	body := &notionapi.DatabaseQueryRequest{
		Filter: &notionapi.PropertyFilter{
			Property: "Category", Select: &notionapi.SelectFilterCondition{Equals: "Pool"},
		},
	}
	resp, err := notion.Database.Query(ctx, eventDbId, body)
	if err != nil {
		return err
	}
	for _, page := range resp.Results {
		if page.CreatedTime.After(LastPollTime) {
			_, err := api.Task.PoolCreate(&task.PoolCreateParams{ID: page.ID.String(), Context: ctx})
			if err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}
