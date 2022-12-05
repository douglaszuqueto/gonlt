package gonlt

import (
	"context"
	"encoding/json"

	"github.com/douglaszuqueto/gonlt/credentials"
	"github.com/douglaszuqueto/gonlt/nlttypes"
	"github.com/vingarcia/krest"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (*nlttypes.AuthResponse, error)
}

type AuthServiceOp struct {
	rest  krest.Client
	creds *credentials.Credentials
}

var _ AuthService = &AuthServiceOp{}

func NewAuthService(rest krest.Client, creds *credentials.Credentials) AuthServiceOp {
	return AuthServiceOp{
		rest:  rest,
		creds: creds,
	}
}

func (s AuthServiceOp) Login(ctx context.Context, email, password string) (*nlttypes.AuthResponse, error) {
	endpoint := buildEndpoint("token")

	resp, err := s.rest.Post(ctx, endpoint, krest.RequestData{
		Headers:    map[string]string{},
		MaxRetries: 3,
		Body: map[string]interface{}{
			"email":    email,
			"password": password,
		},
	})
	if err != nil {
		if resp.StatusCode == 400 {
			return nil, handleAuthError(resp.Body)
		}

		return nil, err
	}

	var account nlttypes.AuthResponse

	err = json.Unmarshal(resp.Body, &account)
	if err != nil {
		return nil, err
	}

	s.creds.Token = account.AccessToken

	return &account, nil
}
