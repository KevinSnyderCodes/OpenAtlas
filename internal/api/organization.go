package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/db"
	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/jsonapi"
)

const defaultOrganizationName = "default"
const defaultOrganizationID = "org-0000000000000000"

var _ Organizations = (*DefaultOrganizations)(nil)

type Organizations interface {
	ReadEntitlements(ctx context.Context, organization string) (*OrganizationReadEntitlementsResponse, error)
	ReadRunQueue(ctx context.Context, organization string) (*OrganizationReadRunQueueResponse, error)
}

type OrganizationReadEntitlementsResponse jsonapi.Document[
	*jsonapi.Resource[*OrganizationReadEntitlementsResponseResourceAttributes],
	OrganizationReadEntitlementsResponseResourceAttributes,
]

func (o *OrganizationReadEntitlementsResponse) MarshalJSON() ([]byte, error) {
	type Alias OrganizationReadEntitlementsResponse

	data, err := json.Marshal((*Alias)(o))
	if err != nil {
		return nil, err
	}

	return data, nil
}

type OrganizationReadEntitlementsResponseResourceAttributes struct {
	CostEstimation                   bool `json:"cost-estimation"`
	ConfigurationDesigner            bool `json:"configuration-designer"`
	ModuleTestsGeneration            bool `json:"module-tests-generation"`
	Operations                       bool `json:"operations"`
	PrivateModuleRegistry            bool `json:"private-module-registry"`
	PolicyEnforcement                bool `json:"policy-enforcement"`
	Sentinel                         bool `json:"sentinel"`
	RunTasks                         bool `json:"run-tasks"`
	StateStorage                     bool `json:"state-storage"`
	Teams                            bool `json:"teams"`
	VCSIntegrations                  bool `json:"vcs-integrations"`
	UsageReporting                   bool `json:"usage-reporting"`
	UserLimit                        int  `json:"user-limit"`
	SelfServeBilling                 bool `json:"self-serve-billing"`
	AuditLogging                     bool `json:"audit-logging"`
	Agents                           bool `json:"agents"`
	SSO                              bool `json:"sso"`
	RunTaskLimit                     int  `json:"run-task-limit"`
	RunTaskWorkspaceLimit            int  `json:"run-task-workspace-limit"`
	RunTaskMandatoryEnforcementLimit int  `json:"run-task-mandatory-enforcement-limit"`
	PolicySetLimit                   int  `json:"policy-set-limit"`
	PolicyLimit                      int  `json:"policy-limit"`
	PolicyMandatoryEnforcementLimit  any  `json:"policy-mandatory-enforcement-limit"`
	VersionedPolicySetLimit          any  `json:"versioned-policy-set-limit"`
}

type OrganizationReadRunQueueResponse RunListResponse

func (o *OrganizationReadRunQueueResponse) MarshalJSON() ([]byte, error) {
	type Alias OrganizationReadRunQueueResponse

	data, err := json.Marshal((*Alias)(o))
	if err != nil {
		return nil, err
	}

	return data, nil
}

type OrganizationsDB interface {
	ListTFERuns(ctx context.Context) ([]db.TFERun, error)
}

type DefaultOrganizations struct {
	db OrganizationsDB
}

func NewDefaultOrganization(db OrganizationsDB) *DefaultOrganizations {
	return &DefaultOrganizations{
		db: db,
	}
}

func (o *DefaultOrganizations) ReadEntitlements(ctx context.Context, organization string) (*OrganizationReadEntitlementsResponse, error) {
	if organization != defaultOrganizationName {
		return nil, fmt.Errorf("organization not found: %s", organization)
	}

	return &OrganizationReadEntitlementsResponse{
		Data: &jsonapi.Resource[*OrganizationReadEntitlementsResponseResourceAttributes]{
			ID:   defaultOrganizationID,
			Type: "entitlement-sets",
			Attributes: &OrganizationReadEntitlementsResponseResourceAttributes{
				Operations: true,
			},
			Links: jsonapi.Links{
				"self": fmt.Sprintf("/api/v2/entitlement-sets/%s", organization),
			},
		},
	}, nil
}

func (o *DefaultOrganizations) ReadRunQueue(ctx context.Context, organization string) (*OrganizationReadRunQueueResponse, error) {
	if organization != defaultOrganizationName {
		return nil, fmt.Errorf("organization not found: %s", organization)
	}

	rows, err := o.db.ListTFERuns(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing runs: %w", err)
	}

	resp, err := (TFERuns)(rows).GetRunListResponse()
	if err != nil {
		return nil, fmt.Errorf("error getting run list response: %w", err)
	}

	return (*OrganizationReadRunQueueResponse)(resp), nil
}
