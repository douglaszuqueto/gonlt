package gonlt

import (
	"fmt"
	"time"

	"github.com/douglaszuqueto/gonlt/credentials"
	"github.com/vingarcia/krest"
)

const (
	baseURL = "https://lora.nlt-iot.com"
)

type Client struct {
	creds credentials.Credentials
	rest  krest.Client

	// Services
	Auth       AuthServiceOp
	Tag        TagsServiceOp
	Connection ConnectionServiceOp
	Device     DeviceServiceOp
	Message    MessageServiceOp
}

func NewClient(creds credentials.Credentials) (*Client, error) {
	if err := creds.Validate(); err != nil {
		return nil, err
	}

	rest := krest.New(10 * time.Second)

	return &Client{
		creds: creds,
		rest:  rest,

		Auth:       NewAuthService(rest, &creds),
		Tag:        NewTagsService(rest, &creds),
		Connection: NewConnectionService(rest, &creds),
		Device:     NewDeviceService(rest, &creds),
		Message:    NewMessageService(rest, &creds),
	}, nil
}

// build endpoint
func buildEndpoint(endpoint string) string {
	return fmt.Sprintf("%s/%s", baseURL, endpoint)
}
