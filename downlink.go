package gonlt

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/douglaszuqueto/gonlt/credentials"
	"github.com/douglaszuqueto/gonlt/nlttypes"
	"github.com/vingarcia/krest"
)

type DownlinkService interface {
	// Send
	Send(ctx context.Context, deviceEui string, params nlttypes.DownlinkRequest) (*nlttypes.DownlinkResponse, error)
}

type DownlinkServiceOp struct {
	rest  krest.Client
	creds *credentials.Credentials
}

var _ DownlinkService = &DownlinkServiceOp{}

func NewDownlinkService(rest krest.Client, creds *credentials.Credentials) *DownlinkServiceOp {
	return &DownlinkServiceOp{
		rest:  rest,
		creds: creds,
	}
}

func (s *DownlinkServiceOp) Send(ctx context.Context, deviceEui string, params nlttypes.DownlinkRequest) (*nlttypes.DownlinkResponse, error) {
	endpoint := buildEndpoint("messages/" + deviceEui + "/send-downlink-claim")

	resp, err := s.rest.Get(ctx, endpoint, krest.RequestData{
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", s.creds.Token),
		},
		MaxRetries: 3,
		Body: map[string]interface{}{
			"payload":   params.Payload,
			"port":      params.Port,
			"confirmed": params.Confirmed,
		},
	})
	if err != nil {
		if resp.StatusCode == 401 {
			return nil, handleUnauthorizedError(resp.Body)
		}

		return nil, err
	}

	var body nlttypes.DownlinkResponse

	err = json.Unmarshal(resp.Body, &body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}
