package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/db"
	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/id"
	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/jsonapi"
)

var _ ConfigurationVersions = (*DefaultConfigurationVersions)(nil)

type ConfigurationVersions interface {
	Read(ctx context.Context, configurationVersionID string) (*ConfigurationVersionDocument, error)
	Create(ctx context.Context, workspaceID string, req *ConfigurationVersionCreateRequest) (*ConfigurationVersionDocument, error)
	Upload(ctx context.Context, configurationVersionID string, r io.Reader) error
}

type ConfigurationVersionDocument jsonapi.Document[
	*jsonapi.Resource[*ConfigurationVersionResourceAttributes],
	ConfigurationVersionResourceAttributes,
]

func (o *ConfigurationVersionDocument) MarshalJSON() ([]byte, error) {
	type Alias ConfigurationVersionDocument

	data, err := json.Marshal((*Alias)(o))
	if err != nil {
		return nil, err
	}

	return data, nil
}

type ConfigurationVersionResourceAttributes struct {
	AutoQueueRuns    bool                                                    `json:"auto-queue-runs"`
	Error            any                                                     `json:"error"`
	ErrorMessage     any                                                     `json:"error-message"`
	Source           string                                                  `json:"source"`
	Speculative      bool                                                    `json:"speculative"`
	Status           string                                                  `json:"status"`
	StatusTimestamps *ConfigurationVersionResourceAttributesStatusTimestamps `json:"status-timestamps"`
	UploadURL        string                                                  `json:"upload-url"`
	Provisional      bool                                                    `json:"provisional"`
}

type ConfigurationVersionResourceAttributesStatusTimestamps struct{}

type ConfigurationVersionReadRequest struct {
	// TODO: Populate
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

type TFEConfigurationVersionDB interface {
	CreateTFEConfigurationVersion(ctx context.Context, arg db.CreateTFEConfigurationVersionParams) (db.TFEConfigurationVersion, error)
	GetTFEConfigurationVersion(ctx context.Context, id string) (db.TFEConfigurationVersion, error)
	UploadTFEConfigurationVersion(ctx context.Context, arg db.UploadTFEConfigurationVersionParams) (db.TFEConfigurationVersion, error)
}

type TFEConfigurationVersion db.TFEConfigurationVersion

func (o TFEConfigurationVersion) ConfigurationVersionID() (id.ConfigurationVersionID, error) {
	v, err := id.NewConfigurationVersionIDFromInternalID(o.ID)
	if err != nil {
		return v, fmt.Errorf("error creating configuration version id: %w", err)
	}

	return v, nil
}

func (o TFEConfigurationVersion) ConfigurationVersionDocument() (*ConfigurationVersionDocument, error) {
	configurationVersionID, err := o.ConfigurationVersionID()
	if err != nil {
		return nil, fmt.Errorf("error getting configuration version id: %w", err)
	}

	return &ConfigurationVersionDocument{
		Data: &jsonapi.Resource[*ConfigurationVersionResourceAttributes]{
			ID:   configurationVersionID.ExternalID(),
			Type: "configuration-versions",
			Attributes: &ConfigurationVersionResourceAttributes{
				AutoQueueRuns: o.AutoQueueRuns,
				Speculative:   o.Speculative,
				Provisional:   o.Provisional,
				Status:        string(o.Status),
				UploadURL:     fmt.Sprintf("https://localhost:8443/api/upload/%s", configurationVersionID.ExternalID()),
			},
		},
	}, nil
}

type DefaultConfigurationVersions struct {
	db TFEConfigurationVersionDB
}

func NewDefaultConfigurationVersions(db TFEConfigurationVersionDB) *DefaultConfigurationVersions {
	return &DefaultConfigurationVersions{
		db: db,
	}
}

func (o *DefaultConfigurationVersions) Read(ctx context.Context, configurationVersionExternalID string) (*ConfigurationVersionDocument, error) {
	configurationVersionID, err := id.NewConfigurationVersionIDFromExternalID(configurationVersionExternalID)
	if err != nil {
		return nil, fmt.Errorf("error creating configuration version id: %w", err)
	}

	row, err := o.db.GetTFEConfigurationVersion(ctx, configurationVersionID.InternalID())
	if err != nil {
		return nil, fmt.Errorf("error getting configuration version by external id: %w", err)
	}

	resp, err := (TFEConfigurationVersion)(row).ConfigurationVersionDocument()
	if err != nil {
		return nil, fmt.Errorf("error creating configuration version document: %w", err)
	}

	return resp, nil
}

func (o *DefaultConfigurationVersions) Create(ctx context.Context, workspaceExternalID string, req *ConfigurationVersionCreateRequest) (*ConfigurationVersionDocument, error) {
	workspaceID, err := id.NewWorkspaceIDFromExternalID(workspaceExternalID)
	if err != nil {
		return nil, fmt.Errorf("error creating workspace id: %w", err)
	}

	if workspaceID.ExternalID() != defaultWorkspaceID {
		return nil, fmt.Errorf("workspace not found: %s", workspaceExternalID)
	}

	arg := db.CreateTFEConfigurationVersionParams{
		ID:            GenerateID(),
		AutoQueueRuns: false,
		Speculative:   false,
		Provisional:   false,
		Status:        db.TFEConfigurationVersionStatusPending,
	}
	row, err := o.db.CreateTFEConfigurationVersion(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("error creating configuration version: %w", err)
	}

	resp, err := (TFEConfigurationVersion)(row).ConfigurationVersionDocument()
	if err != nil {
		return nil, fmt.Errorf("error creating configuration version document: %w", err)
	}

	return resp, nil
}

func (o *DefaultConfigurationVersions) Upload(ctx context.Context, configurationVersionExternalID string, r io.Reader) error {
	configurationVersionID, err := id.NewConfigurationVersionIDFromExternalID(configurationVersionExternalID)
	if err != nil {
		return fmt.Errorf("error creating configuration version id: %w", err)
	}

	row, err := o.db.GetTFEConfigurationVersion(ctx, configurationVersionID.InternalID())
	if err != nil {
		return fmt.Errorf("error getting configuration version by external id: %w", err)
	}
	if row.Status != db.TFEConfigurationVersionStatusPending {
		return fmt.Errorf("invalid configuration version status: %s", row.Status)
	}

	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("error reading upload data: %w", err)
	}

	arg := db.UploadTFEConfigurationVersionParams{
		ID:         row.ID,
		UploadData: data,
	}
	if _, err := o.db.UploadTFEConfigurationVersion(ctx, arg); err != nil {
		return fmt.Errorf("error uploading configuration version: %w", err)
	}

	return nil
}
