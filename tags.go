package gonlt

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/douglaszuqueto/gonlt/credentials"
	"github.com/douglaszuqueto/gonlt/nlttypes"
	"github.com/vingarcia/krest"
)

type TagsService interface {
	// List all tags
	List(ctx context.Context) ([]nlttypes.Tag, error)
}

type TagsServiceOp struct {
	rest krest.Client

	// Credentials
	creds *credentials.Credentials
}

var _ TagsService = &TagsServiceOp{}

func NewTagsService(rest krest.Client, creds *credentials.Credentials) TagsServiceOp {
	return TagsServiceOp{
		rest:  rest,
		creds: creds,
	}
}

func (s TagsServiceOp) List(ctx context.Context) ([]nlttypes.Tag, error) {
	endpoint := buildEndpoint("tags")

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

	var tags []nlttypes.Tag

	err = json.Unmarshal(resp.Body, &tags)
	if err != nil {
		return nil, err
	}

	return tags, nil
}
