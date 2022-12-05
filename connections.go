package gonlt

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/douglaszuqueto/gonlt/credentials"
	"github.com/douglaszuqueto/gonlt/nlttypes"
	"github.com/vingarcia/krest"
)

type ConnectionService interface {
	// List all connections
	List(ctx context.Context) (*nlttypes.ConnectionResponse, error)
}

type ConnectionServiceOp struct {
	rest  krest.Client
	creds *credentials.Credentials
}

var _ ConnectionService = &ConnectionServiceOp{}

func NewConnectionService(rest krest.Client, creds *credentials.Credentials) ConnectionServiceOp {
	return ConnectionServiceOp{
		rest:  rest,
		creds: creds,
	}
}

func (s ConnectionServiceOp) List(ctx context.Context) (*nlttypes.ConnectionResponse, error) {
	endpoint := buildEndpoint("connections")

	resp, err := s.rest.Get(ctx, endpoint, krest.RequestData{
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", s.creds.Token),
		},
		MaxRetries: 3,
	})
	if err != nil {
		if resp.StatusCode == 401 {
			return nil, handleUnauthorizedError(resp.Body)
		}

		return nil, err
	}

	var connections nlttypes.ConnectionResponse

	err = json.Unmarshal(resp.Body, &connections)
	if err != nil {
		return nil, err
	}

	return &connections, nil
}
