package poll

import (
	"context"

	"github.com/gmarcha/notion-trigger/internal/v1/goswagger/client"
	"github.com/gmarcha/notion-trigger/internal/v1/goswagger/client/task"
	"github.com/gmarcha/notion-trigger/internal/v1/log"
	"github.com/jomei/notionapi"
	"golang.org/x/sync/errgroup"
)

func Campus(notion *notionapi.Client, api *client.Notion) error {

	ctx := context.Background()
	resp, err := notion.Database.Query(ctx, campusDbId, nil)
	if err != nil {
		log.Logger.Error("Notion API campus database query")
		return err
	}
	log.Logger.Debug("Notion API campus database query")

	var eg errgroup.Group
	for _, page := range resp.Results {
		currentPage := page
		eg.Go(func() error {
			if currentPage.CreatedTime.After(LastPollTime) {
				_, err := api.Task.CampusCreate(&task.CampusCreateParams{ID: currentPage.ID.String(), Context: ctx})
				if err != nil {
					log.Logger.Error("Internal API create campus tasks")
					return err
				}
				log.Logger.Debug("Internal API create campus tasks")
			}
			return nil
		})
	}
	return eg.Wait()
}
