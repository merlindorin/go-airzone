package pkg

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	v1 "github.com/merlindorin/go-airzone/api/v1"
	"github.com/merlindorin/go-shared/pkg/net/do"
	"github.com/merlindorin/go-shared/pkg/net/rest"
	"go.uber.org/zap"
)

type API interface {
	V1() *v1.V1
}
type APIClient struct {
	rest.Requester

	httpClient *http.Client

	v1 *v1.V1
}

func (c *APIClient) V1() *v1.V1 {
	return c.v1
}

func NewClient(baseURL *url.URL, logger *zap.Logger, options ...Option) API {
	apiClient := &APIClient{}

	for _, option := range append([]Option{WithDefaultHTTPClient()}, options...) {
		option.apply(apiClient)
	}

	apiClient.Requester = rest.NewRest(
		baseURL,
		do.WithJSONRequest(),
		do.WithLogger(logger),
		do.WithClient(apiClient.httpClient),
		WithDefaultHTTPErrorCodeHandler(),
	)

	apiClient.v1 = v1.NewV1(apiClient)

	return apiClient
}

type ErrorMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Level   string `json:"level"`
}

func WithDefaultHTTPErrorCodeHandler() do.Option {
	return do.WithPostRequestHandler("http_response_errorCode_handler", defaultRequestHandler)
}

func defaultRequestHandler(_ context.Context, _ *http.Request, response *http.Response) error {
	if response.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf(
			"unexcpected response status code %d: %s",
			response.StatusCode,
			http.StatusText(response.StatusCode),
		)
	}
	return nil
}
