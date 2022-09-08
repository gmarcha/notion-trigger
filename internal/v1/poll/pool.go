package poll

import (
	"context"

	"github.com/gmarcha/notion-trigger/internal/v1/goswagger/client"
	"github.com/gmarcha/notion-trigger/internal/v1/goswagger/client/task"
	"github.com/gmarcha/notion-trigger/internal/v1/log"
	"github.com/jomei/notionapi"
	"golang.org/x/sync/errgroup"
)

func Pool(notion *notionapi.Client, api *client.Notion) error {

	ctx := context.Background()
	body := &notionapi.DatabaseQueryRequest{
		Filter: &notionapi.PropertyFilter{
			Property: "Category", Select: &notionapi.SelectFilterCondition{Equals: "Pool"},
		},
	}
	resp, err := notion.Database.Query(ctx, eventDbId, body)
	if err != nil {
		log.Logger.Error("Notion API event database query")
		return err
	}
	log.Logger.Debug("Notion API event database query")

	var eg errgroup.Group
	for _, page := range resp.Results {
		currentPage := page
		eg.Go(func() error {
			if currentPage.CreatedTime.After(LastPollTime) {
				_, err := api.Task.PoolCreate(&task.PoolCreateParams{ID: currentPage.ID.String(), Context: ctx})
				if err != nil {
					log.Logger.Error("Internal API create pool tasks")
					return err
				}
				log.Logger.Debug("Internal API create pool tasks")
			}
			return nil
		})
	}
	return eg.Wait()
}
