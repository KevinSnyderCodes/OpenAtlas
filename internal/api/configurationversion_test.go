package api

import (
	"os"
	"testing"

	xjson "github.com/KevinSnyderCodes/OpenAtlas/internal/x/json"
	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/jsonapi"
	"github.com/stretchr/testify/require"
)

var testConfigurationVersionReadResponseMarshalJSONStruct = ConfigurationVersionReadResponse{
	Data: &jsonapi.Resource[*ConfigurationVersionReadResponseResourceAttributes]{
		ID:   "cv-ntv3HbhJqvFzamy7",
		Type: "configuration-versions",
		Attributes: &ConfigurationVersionReadResponseResourceAttributes{
			Error:            nil,
			ErrorMessage:     nil,
			Source:           "gitlab",
			Speculative:      false,
			Status:           "uploaded",
			StatusTimestamps: &ConfigurationVersionReadResponseResourceAttributesStatusTimestamps{},
			Provisional:      false,
		},
		Relationships: jsonapi.Relationships{
			"ingress-attributes": &jsonapi.Relationship{
				Data: map[string]any{
					"id":   "ia-i4MrTxmQXYxH2nYD",
					"type": "ingress-attributes",
				},
				Links: jsonapi.Links{
					"related": "/api/v2/configuration-versions/cv-ntv3HbhJqvFzamy7/ingress-attributes",
				},
			},
		},
		Links: jsonapi.Links{
			"self":     "/api/v2/configuration-versions/cv-ntv3HbhJqvFzamy7",
			"download": "/api/v2/configuration-versions/cv-ntv3HbhJqvFzamy7/download",
		},
	},
}

var testConfigurationVersionCreateRequestUnmarshalJSONStruct = ConfigurationVersionCreateRequest{
	Data: &jsonapi.Resource[*ConfigurationVersionCreateRequestResourceAttributes]{
		Type: "configuration-versions",
		Attributes: &ConfigurationVersionCreateRequestResourceAttributes{
			AutoQueueRuns: true,
		},
	},
}

var testConfigurationVersionCreateResponseMarshalJSONStruct = ConfigurationVersionCreateResponse{
	Data: &jsonapi.Resource[*ConfigurationVersionCreateResponseResourceAttributes]{
		ID:   "cv-UYwHEakurukz85nW",
		Type: "configuration-versions",
		Attributes: &ConfigurationVersionCreateResponseResourceAttributes{
			AutoQueueRuns:    true,
			Error:            nil,
			ErrorMessage:     nil,
			Source:           "tfe-api",
			Speculative:      false,
			Status:           "pending",
			StatusTimestamps: &ConfigurationVersionCreateResponseResourceAttributesStatusTimestamps{},
			UploadURL:        "https://archivist.terraform.io/v1/object/9224c6b3-2e14-4cd7-adff-ed484d7294c2",
			Provisional:      false,
		},
		Relationships: jsonapi.Relationships{
			"ingress-attributes": &jsonapi.Relationship{
				Data: xjson.Null,
				Links: jsonapi.Links{
					"related": "/api/v2/configuration-versions/cv-UYwHEakurukz85nW/ingress-attributes",
				},
			},
		},
		Links: jsonapi.Links{
			"self": "/api/v2/configuration-versions/cv-UYwHEakurukz85nW",
		},
	},
}

func TestConfigurationVersionReadResponse_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		o        *ConfigurationVersionReadResponse
		wantFile string
		wantErr  bool
	}{
		{
			name:     "success",
			o:        &testConfigurationVersionReadResponseMarshalJSONStruct,
			wantFile: "./testdata/configurationversion_read_response.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := os.ReadFile(tt.wantFile)
			if err != nil {
				t.Errorf("error reading file: %v", err)
				return
			}

			got, err := tt.o.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigurationVersionReadResponse.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.JSONEq(t, string(want), string(got))
		})
	}
}

func TestConfigurationVersionCreateRequest_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		dataFile string
		want     *ConfigurationVersionCreateRequest
		wantErr  bool
	}{
		{
			name:     "success",
			dataFile: "./testdata/configurationversion_create_request.json",
			want:     &testConfigurationVersionCreateRequestUnmarshalJSONStruct,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(tt.dataFile)
			if err != nil {
				t.Errorf("error reading file: %v", err)
				return
			}

			o := ConfigurationVersionCreateRequest{}
			if err := o.UnmarshalJSON(data); (err != nil) != tt.wantErr {
				t.Errorf("ConfigurationVersionCreateRequest.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			require.Equal(t, tt.want, &o)
		})
	}
}

func TestConfigurationVersionCreateResponse_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		o        *ConfigurationVersionCreateResponse
		wantFile string
		wantErr  bool
	}{
		{
			name:     "success",
			o:        &testConfigurationVersionCreateResponseMarshalJSONStruct,
			wantFile: "./testdata/configurationversion_create_response.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := os.ReadFile(tt.wantFile)
			if err != nil {
				t.Errorf("error reading file: %v", err)
				return
			}

			got, err := tt.o.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganizationReadResponse.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.JSONEq(t, string(want), string(got))
		})
	}
}
