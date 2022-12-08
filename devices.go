package gonlt

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/douglaszuqueto/gonlt/credentials"
	"github.com/douglaszuqueto/gonlt/nlttypes"
	"github.com/vingarcia/krest"
)

type DeviceService interface {
	List(ctx context.Context) (*nlttypes.DeviceListResponse, error)
	Find(ctx context.Context, deviceID string) (*nlttypes.Device, error)
	Create(ctx context.Context, device nlttypes.DeviceCreateRequest) (*nlttypes.Device, error)
	Activate(ctx context.Context, deviceID string) error
	Deactivate(ctx context.Context, deviceID string) error
	Delete(ctx context.Context, deviceID string) error
}

type DeviceServiceOp struct {
	rest  krest.Client
	creds *credentials.Credentials
}

var _ DeviceService = &DeviceServiceOp{}

func NewDeviceService(rest krest.Client, creds *credentials.Credentials) DeviceServiceOp {
	return DeviceServiceOp{
		rest:  rest,
		creds: creds,
	}
}

func (s DeviceServiceOp) List(ctx context.Context) (*nlttypes.DeviceListResponse, error) {
	endpoint := buildEndpoint("devices?offset=0&limit=100")

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

	var devices nlttypes.DeviceListResponse

	err = json.Unmarshal(resp.Body, &devices)
	if err != nil {
		return nil, err
	}

	return &devices, nil
}

func (s DeviceServiceOp) Find(ctx context.Context, deviceID string) (*nlttypes.Device, error) {
	endpoint := buildEndpoint(fmt.Sprintf("devices/%s", deviceID))

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

		if resp.StatusCode == 404 {
			return nil, errors.New("device not found")
		}

		return nil, err
	}

	var device nlttypes.Device

	err = json.Unmarshal(resp.Body, &device)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (s DeviceServiceOp) Create(ctx context.Context, device nlttypes.DeviceCreateRequest) (*nlttypes.Device, error) {
	endpoint := buildEndpoint("devices/create-device")

	resp, err := s.rest.Post(ctx, endpoint, krest.RequestData{
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", s.creds.Token),
		},
		MaxRetries: 3,
		Body:       device,
	})
	if err != nil {
		if resp.StatusCode == 401 {
			return nil, handleUnauthorizedError(resp.Body)
		}

		if resp.StatusCode == 400 {
			return nil, handleDeviceCreateError(resp.Body)
		}

		return nil, err
	}

	var createdDevice nlttypes.Device

	err = json.Unmarshal(resp.Body, &createdDevice)
	if err != nil {
		return nil, err
	}

	return &createdDevice, nil
}

func (s DeviceServiceOp) Update(ctx context.Context, device nlttypes.Device) (*nlttypes.Device, error) {
	endpoint := buildEndpoint("devices/" + device.DevEui)

	resp, err := s.rest.Patch(ctx, endpoint, krest.RequestData{
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", s.creds.Token),
		},
		MaxRetries: 3,
		Body:       device,
	})
	if err != nil {
		if resp.StatusCode == 401 {
			return nil, handleUnauthorizedError(resp.Body)
		}

		if resp.StatusCode == 400 {
			return nil, handleDeviceCreateError(resp.Body)
		}

		return nil, err
	}

	var result nlttypes.Device

	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s DeviceServiceOp) Activate(ctx context.Context, deviceID string) error {
	endpoint := buildEndpoint(fmt.Sprintf("devices/%s/activation", deviceID))

	resp, err := s.rest.Post(ctx, endpoint, krest.RequestData{
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", s.creds.Token),
		},
		MaxRetries: 3,
		Body: map[string]interface{}{
			"is_active": true,
		},
	})
	if err != nil {
		if resp.StatusCode == 401 {
			return handleUnauthorizedError(resp.Body)
		}

		return err
	}

	var device nlttypes.Device

	err = json.Unmarshal(resp.Body, &device)
	if err != nil {
		return err
	}

	if device.Detail != "" {
		return fmt.Errorf("error activating device: %s", device.Detail)
	}

	return nil
}

func (s DeviceServiceOp) Deactivate(ctx context.Context, deviceID string) error {
	endpoint := buildEndpoint(fmt.Sprintf("devices/%s/activation", deviceID))

	resp, err := s.rest.Post(ctx, endpoint, krest.RequestData{
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", s.creds.Token),
		},
		MaxRetries: 3,
		Body: map[string]interface{}{
			"is_active": false,
		},
	})
	if err != nil {
		if resp.StatusCode == 401 {
			return handleUnauthorizedError(resp.Body)
		}

		return err
	}

	var device nlttypes.Device

	err = json.Unmarshal(resp.Body, &device)
	if err != nil {
		return err
	}

	if device.Detail != "" {
		return fmt.Errorf("error deactivating device: %s", device.Detail)
	}

	return nil
}

func (s DeviceServiceOp) Delete(ctx context.Context, deviceID string) error {
	endpoint := buildEndpoint(fmt.Sprintf("devices/%s", deviceID))

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

		if resp.StatusCode == 404 {
			return errors.New("device not found")
		}

		return err
	}

	var device nlttypes.Device

	err = json.Unmarshal(resp.Body, &device)
	if err != nil {
		return err
	}

	if device.Message != "The device was deleted" {
		return fmt.Errorf("error deleting device: %s", device.Message)
	}

	return nil
}
