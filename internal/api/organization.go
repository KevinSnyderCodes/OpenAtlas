package api

import (
	"context"
	"encoding/json"
	"time"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/jsonapi"
)

var _ Organizations = (*DefaultOrganizations)(nil)

type Organizations interface {
	List(ctx context.Context, options *OrganizationListRequest) (*OrganizationListResponse, error)
	// TODO: Add methods
}

type OrganizationListRequest struct {
	// TODO: Populate
}

type OrganizationListResponse jsonapi.Document[[]*jsonapi.Resource[*OrganizationListResponseResourceAttributes], OrganizationListResponseResourceAttributes]

func (o *OrganizationListResponse) MarshalJSON() ([]byte, error) {
	type Alias OrganizationListResponse

	data, err := json.Marshal((*Alias)(o))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (o *OrganizationListResponse) Validate() error {
	v := (*jsonapi.Document[[]*jsonapi.Resource[*OrganizationListResponseResourceAttributes], OrganizationListResponseResourceAttributes])(o)

	if err := v.Validate(); err != nil {
		return err
	}

	return nil
}

// https://developer.hashicorp.com/terraform/cloud-docs/api-docs/organizations#sample-response
type OrganizationListResponseResourceAttributes struct {
	ExternalID                                        string                                                 `json:"external-id"`
	CreatedAt                                         time.Time                                              `json:"created-at"`
	Email                                             string                                                 `json:"email"`
	SessionTimeout                                    any                                                    `json:"session-timeout"`
	SessionRemember                                   any                                                    `json:"session-remember"`
	CollaboratorAuthPolicy                            string                                                 `json:"collaborator-auth-policy"`
	PlanExpired                                       bool                                                   `json:"plan-expired"`
	PlanExpiresAt                                     any                                                    `json:"plan-expires-at"`
	PlanIsTrial                                       bool                                                   `json:"plan-is-trial"`
	PlanIsEnterprise                                  bool                                                   `json:"plan-is-enterprise"`
	PlanIdentifier                                    string                                                 `json:"plan-identifier"`
	CostEstimationEnabled                             bool                                                   `json:"cost-estimation-enabled"`
	SendPassingStatusesForUntriggeredSpeculativePlans bool                                                   `json:"send-passing-statuses-for-untriggered-speculative-plans"`
	AllowForceDeleteWorkspaces                        bool                                                   `json:"allow-force-delete-workspaces"`
	Name                                              string                                                 `json:"name"`
	Permissions                                       *OrganizationListResponseResourceAttributesPermissions `json:"permissions"`
	FairRunQueuingEnabled                             bool                                                   `json:"fair-run-queuing-enabled"`
	SAMLEnabled                                       bool                                                   `json:"saml-enabled"`
	OwnersTeamSAMLRoleID                              any                                                    `json:"owners-team-saml-role-id"`
	TwoFactorConformant                               bool                                                   `json:"two-factor-conformant"`
	AssessmentsEnforced                               bool                                                   `json:"assessments-enforced"`
	DefaultExecutionMode                              string                                                 `json:"default-execution-mode"`
}

// https://developer.hashicorp.com/terraform/cloud-docs/api-docs/organizations#sample-response
type OrganizationListResponseResourceAttributesPermissions struct {
	CanUpdate                bool `json:"can-update"`
	CanDestroy               bool `json:"can-destroy"`
	CanAccessViaTeams        bool `json:"can-access-via-teams"`
	CanCreateModule          bool `json:"can-create-module"`
	CanCreateTeam            bool `json:"can-create-team"`
	CanCreateWorkspace       bool `json:"can-create-workspace"`
	CanManageUsers           bool `json:"can-manage-users"`
	CanManageSubscription    bool `json:"can-manage-subscription"`
	CanManageSSO             bool `json:"can-manage-sso"`
	CanUpdateOAuth           bool `json:"can-update-oauth"`
	CanUpdateSentinel        bool `json:"can-update-sentinel"`
	CanUpdateSSHKeys         bool `json:"can-update-ssh-keys"`
	CanUpdateAPIToken        bool `json:"can-update-api-token"`
	CanTraverse              bool `json:"can-traverse"`
	CanStartTrial            bool `json:"can-start-trial"`
	CanUpdateAgentPools      bool `json:"can-update-agent-pools"`
	CanManageTags            bool `json:"can-manage-tags"`
	CanManageVarsets         bool `json:"can-manage-varsets"`
	CanReadVarsets           bool `json:"can-read-varsets"`
	CanManagePublicProviders bool `json:"can-manage-public-providers"`
	CanCreateProvider        bool `json:"can-create-provider"`
	CanManagePublicModules   bool `json:"can-manage-public-modules"`
	CanManageCustomProviders bool `json:"can-manage-custom-providers"`
	CanManageRunTasks        bool `json:"can-manage-run-tasks"`
	CanReadRunTasks          bool `json:"can-read-run-tasks"`
	CanCreateProject         bool `json:"can-create-project"`
}

type DefaultOrganizations struct{}

func (s *DefaultOrganizations) List(ctx context.Context, req *OrganizationListRequest) (*OrganizationListResponse, error) {
	return nil, ErrNotImplemented
}
