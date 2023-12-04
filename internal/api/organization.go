package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/jsonapi"
)

var _ Organizations = (*DefaultOrganizations)(nil)

type Organizations interface {
	ReadEntitlements(ctx context.Context, organization string) (*OrganizationReadEntitlementsResponse, error)
}

type OrganizationReadEntitlementsResponse jsonapi.Document[
	*jsonapi.Resource[*OrganizationReadEntitlementsResponseResourceAttributes],
	OrganizationReadEntitlementsResponseResourceAttributes,
]

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

func (o *OrganizationReadEntitlementsResponse) MarshalJSON() ([]byte, error) {
	type Alias OrganizationReadEntitlementsResponse

	data, err := json.Marshal((*Alias)(o))
	if err != nil {
		return nil, err
	}

	return data, nil
}

type DefaultOrganizations struct{}

func (o *DefaultOrganizations) ReadEntitlements(ctx context.Context, organization string) (*OrganizationReadEntitlementsResponse, error) {
	return &OrganizationReadEntitlementsResponse{
		Data: &jsonapi.Resource[*OrganizationReadEntitlementsResponseResourceAttributes]{
			ID:   organization,
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
