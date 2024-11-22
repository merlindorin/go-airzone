package v1

import (
	"context"
	"strconv"

	"github.com/merlindorin/go-shared/pkg/net/do"
	"github.com/merlindorin/go-shared/pkg/net/rest"
)

type APIZone struct {
	cl rest.Requester
}

func NewAPIZone(cl rest.Requester) *APIZone {
	return &APIZone{cl: cl}
}

func (receiver *APIZone) Get(ctx context.Context, systemID string, zoneID string, zone *Zone) error {
	res := struct {
		Data []Zone `json:"data"`
	}{
		Data: []Zone{},
	}

	err := receiver.cl.GET(
		ctx,
		do.WithPath("/api/v1/hvac"),
		do.WithQuery("systemid", systemID),
		do.WithQuery("zoneid", zoneID),
		do.WithUnmarshalBody(&res),
	)
	if err != nil {
		return err
	}

	if len(res.Data) > 0 {
		*zone = res.Data[0]
	}

	return nil
}

func (receiver *APIZone) List(ctx context.Context, system int, zones *[]Zone) error {
	res := struct {
		Data *[]Zone `json:"data"`
	}{
		zones,
	}

	return receiver.cl.GET(
		ctx,
		do.WithPath("/api/v1/hvac"),
		do.WithQuery("systemid", strconv.Itoa(system)),
		do.WithQuery("zoneid", "0"),
		do.WithUnmarshalBody(&res),
	)
}

type ZoneSetRequest struct {
	Name     *string  `json:"name,omitempty"`
	Power    *int     `json:"power,omitempty"`
	SetPoint *float64 `json:"setpoint,omitempty"`
	Sleep    *int     `json:"sleep,omitempty"`
	Mode     *int     `json:"speed,omitempty"`
}

func (receiver *APIZone) Set(ctx context.Context, system int, zone int, opts ...Option) error {
	r := ZoneSetRequest{}

	for _, opt := range opts {
		opt.Apply(&r)
	}

	req := struct {
		SystemID int `json:"systemID"`
		ZoneID   int `json:"zoneID"`
		ZoneSetRequest
	}{
		SystemID:       system,
		ZoneID:         zone,
		ZoneSetRequest: r,
	}

	return receiver.cl.PUT(ctx, do.WithPath("/api/v1/hvac"), do.WithMarshalBody(req))
}
