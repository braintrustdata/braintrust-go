// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"net/http"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
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
	PatchOrganizationMembers shared.PatchOrganizationMembersParam `json:"patch_organization_members,required"`
}

func (r OrganizationMemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PatchOrganizationMembers)
}
