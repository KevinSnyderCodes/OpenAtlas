package api

import (
	"bytes"
	"context"
	"time"

	"github.com/google/jsonapi"
)

var _ Organizations = (*DefaultOrganizations)(nil)

type Organizations interface {
	List(ctx context.Context, options *OrganizationListRequest) (*OrganizationListResponse, error)
	// TODO: Add methods
}

type OrganizationListRequest struct {
	// TODO: Populate
}

// https://developer.hashicorp.com/terraform/cloud-docs/api-docs/organizations#sample-response
type OrganizationListResponse struct {
	Items []*OrganizationListResponseItem
	Links *jsonapi.Links
	Meta  *jsonapi.Meta
}

func (s *OrganizationListResponse) JSONAPIItems() interface{} {
	return s.Items
}

func (s *OrganizationListResponse) JSONAPILinks() *jsonapi.Links {
	return s.Links
}

func (s *OrganizationListResponse) JSONAPIMeta() *jsonapi.Meta {
	return s.Meta
}

// https://developer.hashicorp.com/terraform/cloud-docs/api-docs/organizations#sample-response
type OrganizationListResponseItem struct {
	ID string `jsonapi:"primary,organizations"`

	ExternalID                                        string                                   `jsonapi:"attr,external-id"`
	CreatedAt                                         time.Time                                `jsonapi:"attr,created-at,iso8601milli"`
	Email                                             string                                   `jsonapi:"attr,email"`
	SessionTimeout                                    any                                      `jsonapi:"attr,session-timeout"`
	SessionRemember                                   any                                      `jsonapi:"attr,session-remember"`
	CollaboratorAuthPolicy                            string                                   `jsonapi:"attr,collaborator-auth-policy"`
	PlanExpired                                       bool                                     `jsonapi:"attr,plan-expired"`
	PlanExpiresAt                                     any                                      `jsonapi:"attr,plan-expires-at"`
	PlanIsTrial                                       bool                                     `jsonapi:"attr,plan-is-trial"`
	PlanIsEnterprise                                  bool                                     `jsonapi:"attr,plan-is-enterprise"`
	PlanIdentifier                                    string                                   `jsonapi:"attr,plan-identifier"`
	CostEstimationEnabled                             bool                                     `jsonapi:"attr,cost-estimation-enabled"`
	SendPassingStatusesForUntriggeredSpeculativePlans bool                                     `jsonapi:"attr,send-passing-statuses-for-untriggered-speculative-plans"`
	AllowForceDeleteWorkspaces                        bool                                     `jsonapi:"attr,allow-force-delete-workspaces"`
	Name                                              string                                   `jsonapi:"attr,name"`
	Permissions                                       *OrganizationListResponseItemPermissions `jsonapi:"attr,permissions"`
	FairRunQueuingEnabled                             bool                                     `jsonapi:"attr,fair-run-queuing-enabled"`
	SAMLEnabled                                       bool                                     `jsonapi:"attr,saml-enabled"`
	OwnersTeamSAMLRoleID                              any                                      `jsonapi:"attr,owners-team-saml-role-id"`
	TwoFactorConformant                               bool                                     `jsonapi:"attr,two-factor-conformant"`
	AssessmentsEnforced                               bool                                     `jsonapi:"attr,assessments-enforced"`
	DefaultExecutionMode                              string                                   `jsonapi:"attr,default-execution-mode"`

	RelationshipDefaultAgentPool    any                                                     `jsonapi:"relation,default-agent-pool"`
	RelationshipOAuthTokens         any                                                     `jsonapi:"relation,oauth-tokens,omitempty"`
	RelationshipAuthenticationToken any                                                     `jsonapi:"relation,authentication-token,omitempty"`
	RelationshipEntitlementSet      *OrganizationListResponseItemRelationshipEntitlementSet `jsonapi:"relation,entitlement-set"`
	RelationshipSubscription        any                                                     `jsonapi:"relation,subscription,omitempty"`

	RelationshipLinksOAuthTokens         *jsonapi.Links
	RelationshipLinksAuthenticationToken *jsonapi.Links
	RelationshipLinksEntitlementSet      *jsonapi.Links
	RelationshipLinksSubscription        *jsonapi.Links

	Links *jsonapi.Links
}

func (s *OrganizationListResponseItem) JSONAPIRelationshipLinks(relation string) *jsonapi.Links {
	switch relation {
	case "oauth-tokens":
		return s.RelationshipLinksOAuthTokens
	case "authentication-token":
		return s.RelationshipLinksAuthenticationToken
	case "entitlement-set":
		return s.RelationshipLinksEntitlementSet
	case "subscription":
		return s.RelationshipLinksSubscription
	default:
		return nil
	}
}

func (s *OrganizationListResponseItem) JSONAPILinks() *jsonapi.Links {
	return s.Links
}

// https://developer.hashicorp.com/terraform/cloud-docs/api-docs/organizations#sample-response
type OrganizationListResponseItemPermissions struct {
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

// https://developer.hashicorp.com/terraform/cloud-docs/api-docs/organizations#sample-response
type OrganizationListResponseItemRelationshipEntitlementSet struct {
	ID string `jsonapi:"primary,entitlement-sets"`
}

func (s *OrganizationListResponse) Marshal() ([]byte, error) {
	var buf bytes.Buffer

	if err := jsonapi.MarshalPayloadWithoutIncluded(&buf, s); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type DefaultOrganizations struct{}

func (s *DefaultOrganizations) List(ctx context.Context, req *OrganizationListRequest) (*OrganizationListResponse, error) {
	return nil, ErrNotImplemented
}
