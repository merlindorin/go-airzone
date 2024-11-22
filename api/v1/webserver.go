package v1

import (
	"context"

	"github.com/merlindorin/go-shared/pkg/net/do"
	"github.com/merlindorin/go-shared/pkg/net/rest"
)

type APIWebserver struct {
	cl rest.Requester
}

func NewAPIWebserver(cl rest.Requester) *APIWebserver {
	return &APIWebserver{cl: cl}
}

func (receiver *APIWebserver) Get(ctx context.Context, webserver *Webserver) error {
	return receiver.cl.POST(ctx, do.WithPath("/api/v1/webserver"), do.WithUnmarshalBody(&webserver))
}
