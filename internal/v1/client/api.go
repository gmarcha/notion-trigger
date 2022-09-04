package client

import (
	"github.com/gmarcha/notion-trigger/internal/v1/goswagger/client"
	swaggerClient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func Api() *client.Notion {

	httpTransport := swaggerClient.New(client.DefaultHost, client.DefaultBasePath, client.DefaultSchemes)
	return client.New(httpTransport, strfmt.Default)
}
