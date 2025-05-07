// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"net/http"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/braintrustdata/braintrust-go/packages/param"
	"github.com/braintrustdata/braintrust-go/shared"
)

// OrganizationMemberService contains methods and other services that help with
// interacting with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationMemberService] method instead.
type OrganizationMemberService struct {
	Options []option.RequestOption
}

// NewOrganizationMemberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationMemberService(opts ...option.RequestOption) (r OrganizationMemberService) {
	r = OrganizationMemberService{}
	r.Options = opts
	return
}

// Modify organization membership
func (r *OrganizationMemberService) Update(ctx context.Context, body OrganizationMemberUpdateParams, opts ...option.RequestOption) (res *shared.PatchOrganizationMembersOutput, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/organization/members"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type OrganizationMemberUpdateParams struct {
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, or in case you want to
	// explicitly assert the organization you are modifying, you may specify the id of
	// the organization.
	OrgID param.Opt[string] `json:"org_id,omitzero"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, or in case you want to
	// explicitly assert the organization you are modifying, you may specify the name
	// of the organization.
	OrgName param.Opt[string] `json:"org_name,omitzero"`
	// Users to invite to the organization
	InviteUsers OrganizationMemberUpdateParamsInviteUsers `json:"invite_users,omitzero"`
	// Users to remove from the organization
	RemoveUsers OrganizationMemberUpdateParamsRemoveUsers `json:"remove_users,omitzero"`
	paramObj
}

func (r OrganizationMemberUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationMemberUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationMemberUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Users to invite to the organization
type OrganizationMemberUpdateParamsInviteUsers struct {
	// Singular form of group_ids
	GroupID param.Opt[string] `json:"group_id,omitzero" format:"uuid"`
	// Singular form of group_names
	GroupName param.Opt[string] `json:"group_name,omitzero"`
	// If true, send invite emails to the users who wore actually added
	SendInviteEmails param.Opt[bool] `json:"send_invite_emails,omitzero"`
	// Emails of users to invite
	Emails []string `json:"emails,omitzero"`
	// Optional list of group ids to add newly-invited users to.
	GroupIDs []string `json:"group_ids,omitzero" format:"uuid"`
	// Optional list of group names to add newly-invited users to.
	GroupNames []string `json:"group_names,omitzero"`
	// Ids of existing users to invite
	IDs []string `json:"ids,omitzero" format:"uuid"`
	paramObj
}

func (r OrganizationMemberUpdateParamsInviteUsers) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationMemberUpdateParamsInviteUsers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationMemberUpdateParamsInviteUsers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Users to remove from the organization
type OrganizationMemberUpdateParamsRemoveUsers struct {
	// Emails of users to remove
	Emails []string `json:"emails,omitzero"`
	// Ids of users to remove
	IDs []string `json:"ids,omitzero" format:"uuid"`
	paramObj
}

func (r OrganizationMemberUpdateParamsRemoveUsers) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationMemberUpdateParamsRemoveUsers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationMemberUpdateParamsRemoveUsers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
