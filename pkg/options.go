package pkg

import (
	"crypto/tls"
	"net/http"
)

type Option func(cl *APIClient)

func (w Option) apply(cl *APIClient) {
	w(cl)
}

func WithHTTPClient(cl *http.Client) Option {
	return func(a *APIClient) {
		a.httpClient = cl
	}
}

func WithDefaultHTTPClient() Option {
	cl := http.DefaultClient
	cl.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // not supported yet
	}
	return WithHTTPClient(cl)
}
