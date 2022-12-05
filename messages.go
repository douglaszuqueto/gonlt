package gonlt

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/douglaszuqueto/gonlt/credentials"
	"github.com/douglaszuqueto/gonlt/nlttypes"
	"github.com/vingarcia/krest"
)

type MessageService interface {
	List(ctx context.Context, deviceEui string, filter MessageFilter) (*nlttypes.Messages, error)
}

type MessageServiceOp struct {
	rest  krest.Client
	creds *credentials.Credentials
}

type MessageFilter struct {
	Type      string
	StartDate time.Time
	EndDate   time.Time
}

// Build a filter for messages
func (s MessageFilter) Build() (string, error) {
	filter := ""

	filter += fmt.Sprintf("message_type=%s", s.Type)
	filter += fmt.Sprintf("&initial_date=%s", s.StartDate.Format("2006-01-02 15:04"))
	filter += fmt.Sprintf("&final_date=%s", s.EndDate.Format("2006-01-02 15:04"))

	return filter, nil
}

var _ MessageService = &MessageServiceOp{}

func NewMessageService(rest krest.Client, creds *credentials.Credentials) MessageServiceOp {
	return MessageServiceOp{
		rest:  rest,
		creds: creds,
	}
}

func (s MessageServiceOp) List(ctx context.Context, deviceEui string, filter MessageFilter) (*nlttypes.Messages, error) {
	filterStr, err := filter.Build()
	if err != nil {
		return nil, err
	}

	urlQuery, err := url.ParseQuery(filterStr)
	if err != nil {
		return nil, err
	}

	endpoint := buildEndpoint("messages/" + deviceEui + "?" + urlQuery.Encode())

	resp, err := s.rest.Get(ctx, endpoint, krest.RequestData{
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", s.creds.Token),
		},
		MaxRetries: 3,
	})
	if err != nil {
		return nil, err
	}

	var messages nlttypes.Messages

	err = json.Unmarshal(resp.Body, &messages)
	if err != nil {
		return nil, err
	}

	return &messages, nil
}
