package gonlt

import (
	"encoding/json"
	"fmt"

	"github.com/douglaszuqueto/gonlt/nlttypes"
)

// handle errors
func handleError(body []byte, statusCode int) error {
	switch statusCode {
	case 401:
		return handleUnauthorizedError(body)
	case 403:
		return handleAuthError(body)
	case 404:
		return fmt.Errorf("not found")
	case 422:
		return handleDeviceCreateError(body)
	default:
		return fmt.Errorf("unknown error: %s", body)
	}
}

// handle unauthorized errors
func handleUnauthorizedError(body []byte) error {
	var errResp nlttypes.AuthError

	err := json.Unmarshal(body, &errResp)
	if err != nil {
		return err
	}

	return fmt.Errorf("invalid token: %s", errResp.Detail)
}

// handle auth errors
func handleAuthError(body []byte) error {
	var errResp nlttypes.AuthError

	err := json.Unmarshal(body, &errResp)
	if err != nil {
		return err
	}

	return fmt.Errorf("invalid credentials: %s", errResp.Detail)
}

// handle device errors
func handleDeviceCreateError(body []byte) error {
	var errResp nlttypes.DeviceCreateError

	err := json.Unmarshal(body, &errResp)
	if err != nil {
		return err
	}

	return fmt.Errorf("device: %s", errResp.Detail)
}
