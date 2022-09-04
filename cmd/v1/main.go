package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/gmarcha/notion-trigger/internal/v1/client"
	_ "github.com/gmarcha/notion-trigger/internal/v1/env"
	"github.com/gmarcha/notion-trigger/internal/v1/poll"
)

func main() {

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)

	notionClient := client.Notion()
	apiClient := client.Api()

	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for t := range ticker.C {
			poll.Campus(notionClient, apiClient)
			poll.Pool(notionClient, apiClient)
			poll.LastPollTime = t
		}
	}()
	<-sc
	ticker.Stop()
}
