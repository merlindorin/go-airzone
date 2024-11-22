package v1

import (
	"context"

	"github.com/merlindorin/go-shared/pkg/net/do"
	"github.com/merlindorin/go-shared/pkg/net/rest"
)

type APIVersion struct {
	cl rest.Requester
}

func NewAPIVersion(cl rest.Requester) *APIVersion {
	return &APIVersion{cl: cl}
}

func (receiver *APIVersion) Get(ctx context.Context, version *Version) error {
	return receiver.cl.POST(ctx, do.WithPath("/api/v1/version"), do.WithUnmarshalBody(&version))
}
