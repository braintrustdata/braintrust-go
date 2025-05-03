// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/braintrustdata/braintrust-go/internal/apiquery"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/braintrustdata/braintrust-go/packages/pagination"
	"github.com/braintrustdata/braintrust-go/packages/param"
	"github.com/braintrustdata/braintrust-go/shared"
)

// GroupService contains methods and other services that help with interacting with
// the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGroupService] method instead.
type GroupService struct {
	Options []option.RequestOption
}

// NewGroupService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGroupService(opts ...option.RequestOption) (r GroupService) {
	r = GroupService{}
	r.Options = opts
	return
}

// Create a new group. If there is an existing group with the same name as the one
// specified in the request, will return the existing group unmodified
func (r *GroupService) New(ctx context.Context, body GroupNewParams, opts ...option.RequestOption) (res *shared.Group, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a group object by its id
func (r *GroupService) Get(ctx context.Context, groupID string, opts ...option.RequestOption) (res *shared.Group, err error) {
	opts = append(r.Options[:], opts...)
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("v1/group/%s", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update a group object. Specify the fields to update in the payload.
// Any object-type fields will be deep-merged with existing content. Currently we
// do not support removing fields or setting them to null.
func (r *GroupService) Update(ctx context.Context, groupID string, body GroupUpdateParams, opts ...option.RequestOption) (res *shared.Group, err error) {
	opts = append(r.Options[:], opts...)
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("v1/group/%s", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all groups. The groups are sorted by creation date, with the most
// recently-created groups coming first
func (r *GroupService) List(ctx context.Context, query GroupListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.Group], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/group"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List out all groups. The groups are sorted by creation date, with the most
// recently-created groups coming first
func (r *GroupService) ListAutoPaging(ctx context.Context, query GroupListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.Group] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a group object by its id
func (r *GroupService) Delete(ctx context.Context, groupID string, opts ...option.RequestOption) (res *shared.Group, err error) {
	opts = append(r.Options[:], opts...)
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("v1/group/%s", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create or replace group. If there is an existing group with the same name as the
// one specified in the request, will replace the existing group with the provided
// fields
func (r *GroupService) Replace(ctx context.Context, body GroupReplaceParams, opts ...option.RequestOption) (res *shared.Group, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type GroupNewParams struct {
	// Name of the group
	Name string `json:"name,required"`
	// Textual description of the group
	Description param.Opt[string] `json:"description,omitzero"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the group belongs in.
	OrgName param.Opt[string] `json:"org_name,omitzero"`
	// Ids of the groups this group inherits from
	//
	// An inheriting group has all the users contained in its member groups, as well as
	// all of their inherited users
	MemberGroups []string `json:"member_groups,omitzero" format:"uuid"`
	// Ids of users which belong to this group
	MemberUsers []string `json:"member_users,omitzero" format:"uuid"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GroupNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r GroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type GroupUpdateParams struct {
	// Textual description of the group
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the group
	Name param.Opt[string] `json:"name,omitzero"`
	// A list of group IDs to add to the group's inheriting-from set
	AddMemberGroups []string `json:"add_member_groups,omitzero" format:"uuid"`
	// A list of user IDs to add to the group
	AddMemberUsers []string `json:"add_member_users,omitzero" format:"uuid"`
	// A list of group IDs to remove from the group's inheriting-from set
	RemoveMemberGroups []string `json:"remove_member_groups,omitzero" format:"uuid"`
	// A list of user IDs to remove from the group
	RemoveMemberUsers []string `json:"remove_member_users,omitzero" format:"uuid"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GroupUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r GroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow GroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type GroupListParams struct {
	// Limit the number of objects to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Opt[string] `query:"ending_before,omitzero" format:"uuid" json:"-"`
	// Name of the group to search for
	GroupName param.Opt[string] `query:"group_name,omitzero" json:"-"`
	// Filter search results to within a particular organization
	OrgName param.Opt[string] `query:"org_name,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Opt[string] `query:"starting_after,omitzero" format:"uuid" json:"-"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs GroupListParamsIDsUnion `query:"ids,omitzero" format:"uuid" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GroupListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [GroupListParams]'s query parameters as `url.Values`.
func (r GroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GroupListParamsIDsUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u GroupListParamsIDsUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }

func (u *GroupListParamsIDsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type GroupReplaceParams struct {
	// Name of the group
	Name string `json:"name,required"`
	// Textual description of the group
	Description param.Opt[string] `json:"description,omitzero"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the group belongs in.
	OrgName param.Opt[string] `json:"org_name,omitzero"`
	// Ids of the groups this group inherits from
	//
	// An inheriting group has all the users contained in its member groups, as well as
	// all of their inherited users
	MemberGroups []string `json:"member_groups,omitzero" format:"uuid"`
	// Ids of users which belong to this group
	MemberUsers []string `json:"member_users,omitzero" format:"uuid"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GroupReplaceParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r GroupReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow GroupReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
