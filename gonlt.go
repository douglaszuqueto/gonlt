package gonlt

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/douglaszuqueto/gonlt/credentials"
	"github.com/vingarcia/krest"
)

const (
	baseURL = "https://lora.nlt-iot.com"
)

type Client struct {
	ctx    context.Context
	cancel context.CancelFunc

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

	ctx, cancel := context.WithCancel(context.Background())

	client := &Client{
		ctx:    ctx,
		cancel: cancel,

		creds: creds,
		rest:  rest,

		Auth:       NewAuthService(rest, &creds),
		Tag:        NewTagsService(rest, &creds),
		Connection: NewConnectionService(rest, &creds),
		Device:     NewDeviceService(rest, &creds),
		Message:    NewMessageService(rest, &creds),
	}

	if err := client.autoLogin(); err != nil {
		return nil, err
	}

	return client, nil
}

// autoLogin will try to login if the credentials has the AutoLogin flag set to true
func (c *Client) autoLogin() error {
	if !c.creds.AutoLogin {
		return nil
	}

	_, err := c.Auth.Login(context.Background(), c.creds.Email, c.creds.Passwd)
	if err != nil {
		return err
	}

	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-c.ctx.Done():
				log.Println("gonlt: client stopping")
				return
			case <-ticker.C:
				log.Println("gonlt: trying to login")

				_, err := c.Auth.Login(context.Background(), c.creds.Email, c.creds.Passwd)
				if err != nil {
					log.Println(err)
				}

				log.Println("gonlt: token refreshed")
			}
		}
	}()

	return nil
}

// Stop will stop the client
func (c *Client) Stop() {
	c.cancel()

	log.Println("gonlt: client stopped")
}

// build endpoint
func buildEndpoint(endpoint string) string {
	return fmt.Sprintf("%s/%s", baseURL, endpoint)
}
