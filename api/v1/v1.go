package v1

import (
	"github.com/merlindorin/go-shared/pkg/net/rest"
)

type V1 struct {
	Zone      *APIZone
	System    *APISystem
	Version   *APIVersion
	Webserver *APIWebserver
}

func NewV1(cl rest.Requester) *V1 {
	return &V1{
		Zone:      NewAPIZone(cl),
		System:    NewAPISystem(cl),
		Version:   NewAPIVersion(cl),
		Webserver: NewAPIWebserver(cl),
	}
}
