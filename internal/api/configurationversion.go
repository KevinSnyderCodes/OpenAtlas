package api

import (
	"context"
	"encoding/json"
	"io"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/jsonapi"
)

var _ ConfigurationVersions = (*DefaultConfigurationVersions)(nil)

type ConfigurationVersions interface {
	Read(ctx context.Context, req *ConfigurationVersionReadRequest) (*ConfigurationVersionCreateResponse, error)
	Create(ctx context.Context, req *ConfigurationVersionCreateRequest) (*ConfigurationVersionCreateResponse, error)
	Upload(ctx context.Context, r io.Reader) error
}

type ConfigurationVersionReadRequest struct {
	// TODO: Populate
}

type ConfigurationVersionReadResponse jsonapi.Document[
	*jsonapi.Resource[*ConfigurationVersionReadResponseResourceAttributes],
	ConfigurationVersionReadResponseResourceAttributes,
]

type ConfigurationVersionReadResponseResourceAttributes struct {
	Error            any                                                                 `json:"error"`
	ErrorMessage     any                                                                 `json:"error-message"`
	Source           string                                                              `json:"source"`
	Speculative      bool                                                                `json:"speculative"`
	Status           string                                                              `json:"status"`
	StatusTimestamps *ConfigurationVersionReadResponseResourceAttributesStatusTimestamps `json:"status-timestamps"`
	Provisional      bool                                                                `json:"provisional"`
}

type ConfigurationVersionReadResponseResourceAttributesStatusTimestamps struct{}

func (o *ConfigurationVersionReadResponse) MarshalJSON() ([]byte, error) {
	type Alias ConfigurationVersionReadResponse

	data, err := json.Marshal((*Alias)(o))
	if err != nil {
		return nil, err
	}

	return data, nil
}

type ConfigurationVersionCreateRequest jsonapi.Document[
	*jsonapi.Resource[*ConfigurationVersionCreateRequestResourceAttributes],
	ConfigurationVersionCreateRequestResourceAttributes,
]

type ConfigurationVersionCreateRequestResourceAttributes struct {
	AutoQueueRuns bool `json:"auto-queue-runs"`
	Speculative   bool `json:"speculative"`
	Provisional   bool `json:"provisional"`
}

func (o *ConfigurationVersionCreateRequest) UnmarshalJSON(data []byte) error {
	type Alias ConfigurationVersionCreateRequest

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	*o = ConfigurationVersionCreateRequest(alias)

	return nil
}

type ConfigurationVersionCreateResponse jsonapi.Document[
	*jsonapi.Resource[*ConfigurationVersionCreateResponseResourceAttributes],
	ConfigurationVersionCreateResponseResourceAttributes,
]

type ConfigurationVersionCreateResponseResourceAttributes struct {
	AutoQueueRuns    bool                                                                  `json:"auto-queue-runs"`
	Error            any                                                                   `json:"error"`
	ErrorMessage     any                                                                   `json:"error-message"`
	Source           string                                                                `json:"source"`
	Speculative      bool                                                                  `json:"speculative"`
	Status           string                                                                `json:"status"`
	StatusTimestamps *ConfigurationVersionCreateResponseResourceAttributesStatusTimestamps `json:"status-timestamps"`
	UploadURL        string                                                                `json:"upload-url"`
	Provisional      bool                                                                  `json:"provisional"`
}

type ConfigurationVersionCreateResponseResourceAttributesStatusTimestamps struct{}

func (o *ConfigurationVersionCreateResponse) MarshalJSON() ([]byte, error) {
	type Alias ConfigurationVersionCreateResponse

	data, err := json.Marshal((*Alias)(o))
	if err != nil {
		return nil, err
	}

	return data, nil
}

type DefaultConfigurationVersions struct{}

func (o *DefaultConfigurationVersions) Read(ctx context.Context, req *ConfigurationVersionReadRequest) (*ConfigurationVersionCreateResponse, error) {
	return nil, ErrNotImplemented
}

func (o *DefaultConfigurationVersions) Create(ctx context.Context, req *ConfigurationVersionCreateRequest) (*ConfigurationVersionCreateResponse, error) {
	return nil, ErrNotImplemented
}

func (o *DefaultConfigurationVersions) Upload(ctx context.Context, r io.Reader) error {
	return ErrNotImplemented
}
