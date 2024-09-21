// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"net/http"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/internal/param"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
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
func NewOrganizationMemberService(opts ...option.RequestOption) (r *OrganizationMemberService) {
	r = &OrganizationMemberService{}
	r.Options = opts
	return
}

// Modify organization membership
func (r *OrganizationMemberService) Update(ctx context.Context, body OrganizationMemberUpdateParams, opts ...option.RequestOption) (res *OrganizationMemberUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/organization/members"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type OrganizationMemberUpdateResponse struct {
	Status OrganizationMemberUpdateResponseStatus `json:"status,required"`
	// If invite emails failed to send for some reason, the patch operation will still
	// complete, but we will return an error message here
	SendEmailError string                               `json:"send_email_error,nullable"`
	JSON           organizationMemberUpdateResponseJSON `json:"-"`
}

// organizationMemberUpdateResponseJSON contains the JSON metadata for the struct
// [OrganizationMemberUpdateResponse]
type organizationMemberUpdateResponseJSON struct {
	Status         apijson.Field
	SendEmailError apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMemberUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationMemberUpdateResponseStatus string

const (
	OrganizationMemberUpdateResponseStatusSuccess OrganizationMemberUpdateResponseStatus = "success"
)

func (r OrganizationMemberUpdateResponseStatus) IsKnown() bool {
	switch r {
	case OrganizationMemberUpdateResponseStatusSuccess:
		return true
	}
	return false
}

type OrganizationMemberUpdateParams struct {
	// Users to invite to the organization
	InviteUsers param.Field[OrganizationMemberUpdateParamsInviteUsers] `json:"invite_users"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, or in case you want to
	// explicitly assert the organization you are modifying, you may specify the id of
	// the organization.
	OrgID param.Field[string] `json:"org_id"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, or in case you want to
	// explicitly assert the organization you are modifying, you may specify the name
	// of the organization.
	OrgName param.Field[string] `json:"org_name"`
	// Users to remove from the organization
	RemoveUsers param.Field[OrganizationMemberUpdateParamsRemoveUsers] `json:"remove_users"`
}

func (r OrganizationMemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Users to invite to the organization
type OrganizationMemberUpdateParamsInviteUsers struct {
	// Emails of users to invite
	Emails param.Field[[]string] `json:"emails"`
	// Singular form of group_ids
	GroupID param.Field[string] `json:"group_id" format:"uuid"`
	// Optional list of group ids to add newly-invited users to.
	GroupIDs param.Field[[]string] `json:"group_ids" format:"uuid"`
	// Singular form of group_names
	GroupName param.Field[string] `json:"group_name"`
	// Optional list of group names to add newly-invited users to.
	GroupNames param.Field[[]string] `json:"group_names"`
	// Ids of existing users to invite
	IDs param.Field[[]string] `json:"ids" format:"uuid"`
	// If true, send invite emails to the users who wore actually added
	SendInviteEmails param.Field[bool] `json:"send_invite_emails"`
}

func (r OrganizationMemberUpdateParamsInviteUsers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Users to remove from the organization
type OrganizationMemberUpdateParamsRemoveUsers struct {
	// Emails of users to remove
	Emails param.Field[[]string] `json:"emails"`
	// Ids of users to remove
	IDs param.Field[[]string] `json:"ids" format:"uuid"`
}

func (r OrganizationMemberUpdateParamsRemoveUsers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
