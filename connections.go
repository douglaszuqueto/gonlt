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

	// Create a new connection
	Create(ctx context.Context, req nlttypes.CreateConnectionRequest) (*nlttypes.CreateConnectionResponse, error)

	// Update a connection
	Update(ctx context.Context, req nlttypes.UpdateConnectionRequest) (*nlttypes.UpdateConnectionResponse, error)

	// Delete a connection
	Delete(ctx context.Context, id int) error
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

// Create a new connection
func (s ConnectionServiceOp) Create(ctx context.Context, req nlttypes.CreateConnectionRequest) (*nlttypes.CreateConnectionResponse, error) {
	endpoint := buildEndpoint("connections")

	resp, err := s.rest.Post(ctx, endpoint, krest.RequestData{
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", s.creds.Token),
		},
		MaxRetries: 3,
		Body:       req,
	})
	if err != nil {
		if resp.StatusCode == 401 {
			return nil, handleUnauthorizedError(resp.Body)
		}

		return nil, err
	}

	var connection nlttypes.CreateConnectionResponse

	err = json.Unmarshal(resp.Body, &connection)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

// Update a connection
func (s ConnectionServiceOp) Update(ctx context.Context, req nlttypes.UpdateConnectionRequest) (*nlttypes.UpdateConnectionResponse, error) {
	endpoint := buildEndpoint(fmt.Sprintf("connections/%d", req.Connectionmodel.ID))

	resp, err := s.rest.Patch(ctx, endpoint, krest.RequestData{
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", s.creds.Token),
		},
		MaxRetries: 3,
		Body:       req,
	})
	if err != nil {
		if resp.StatusCode == 401 {
			return nil, handleUnauthorizedError(resp.Body)
		}

		return nil, err
	}

	var connection nlttypes.UpdateConnectionResponse

	err = json.Unmarshal(resp.Body, &connection)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

// Delete a connection
func (s ConnectionServiceOp) Delete(ctx context.Context, id int) error {
	endpoint := buildEndpoint(fmt.Sprintf("connections/%d", id))

	resp, err := s.rest.Delete(ctx, endpoint, krest.RequestData{
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", s.creds.Token),
		},
		MaxRetries: 3,
	})
	if err != nil {
		if resp.StatusCode == 401 {
			return handleUnauthorizedError(resp.Body)
		}

		return err
	}

	var response nlttypes.DeleteConnectionResponse

	err = json.Unmarshal(resp.Body, &response)
	if err != nil {
		return err
	}

	if response.Message != "Connection deleted." {
		return fmt.Errorf("unexpected response: %s", response.Message)
	}

	return nil
}
