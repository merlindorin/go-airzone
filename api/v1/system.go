package v1

import (
	"context"

	"github.com/merlindorin/go-shared/pkg/net/do"
	"github.com/merlindorin/go-shared/pkg/net/rest"
)

type APISystem struct {
	cl rest.Requester
}

func NewAPISystem(cl rest.Requester) *APISystem {
	return &APISystem{cl: cl}
}

func (receiver *APISystem) Get(ctx context.Context, systemID string, system *System) error {
	res := struct {
		Data *System `json:"data"`
	}{
		Data: system,
	}

	return receiver.cl.GET(
		ctx,
		do.WithPath("/api/v1/hvac"),
		do.WithQuery("systemid", systemID),
		do.WithUnmarshalBody(&res),
	)
}

func (receiver *APISystem) List(ctx context.Context, systems *[]System) error {
	res := struct {
		Systems *[]System `json:"systems"`
	}{
		systems,
	}

	return receiver.cl.GET(
		ctx,
		do.WithPath("/api/v1/hvac"),
		do.WithQuery("systemid", "127"),
		do.WithUnmarshalBody(&res),
	)
}
