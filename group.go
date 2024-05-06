// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/internal/apiquery"
	"github.com/braintrustdata/braintrust-go/internal/pagination"
	"github.com/braintrustdata/braintrust-go/internal/param"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
)

// GroupService contains methods and other services that help with interacting with
// the braintrust API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewGroupService] method instead.
type GroupService struct {
	Options []option.RequestOption
}

// NewGroupService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGroupService(opts ...option.RequestOption) (r *GroupService) {
	r = &GroupService{}
	r.Options = opts
	return
}

// Create a new group. If there is an existing group with the same name as the one
// specified in the request, will return the existing group unmodified
func (r *GroupService) New(ctx context.Context, body GroupNewParams, opts ...option.RequestOption) (res *Group, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a group object by its id
func (r *GroupService) Get(ctx context.Context, groupID string, opts ...option.RequestOption) (res *Group, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/group/%s", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update a group object. Specify the fields to update in the payload.
// Any object-type fields will be deep-merged with existing content. Currently we
// do not support removing fields or setting them to null.
func (r *GroupService) Update(ctx context.Context, groupID string, body GroupUpdateParams, opts ...option.RequestOption) (res *Group, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/group/%s", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all groups. The groups are sorted by creation date, with the most
// recently-created groups coming first
func (r *GroupService) List(ctx context.Context, query GroupListParams, opts ...option.RequestOption) (res *pagination.ListObjects[Group], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *GroupService) ListAutoPaging(ctx context.Context, query GroupListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[Group] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a group object by its id
func (r *GroupService) Delete(ctx context.Context, groupID string, opts ...option.RequestOption) (res *Group, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/group/%s", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// NOTE: This operation is deprecated and will be removed in a future revision of
// the API. Create or replace a new group. If there is an existing group with the
// same name as the one specified in the request, will return the existing group
// unmodified, will replace the existing group with the provided fields
func (r *GroupService) Replace(ctx context.Context, body GroupReplaceParams, opts ...option.RequestOption) (res *Group, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/group"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// A group is a collection of users which can be assigned an ACL
//
// Groups can consist of individual users, as well as a set of groups they inherit
// from
type Group struct {
	// Unique identifier for the group
	ID string `json:"id,required" format:"uuid"`
	// Name of the group
	Name string `json:"name,required"`
	// Unique id for the organization that the group belongs under
	//
	// It is forbidden to change the org after creating a group
	OrgID string `json:"org_id,required" format:"uuid"`
	// Date of group creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Date of group deletion, or null if the group is still active
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Textual description of the group
	Description string `json:"description,nullable"`
	// Ids of the groups this group inherits from
	//
	// An inheriting group has all the users contained in its member groups, as well as
	// all of their inherited users
	MemberGroups []string `json:"member_groups,nullable" format:"uuid"`
	// Ids of users which belong to this group
	MemberUsers []string `json:"member_users,nullable" format:"uuid"`
	// Identifies the user who created the group
	UserID string    `json:"user_id,nullable" format:"uuid"`
	JSON   groupJSON `json:"-"`
}

// groupJSON contains the JSON metadata for the struct [Group]
type groupJSON struct {
	ID           apijson.Field
	Name         apijson.Field
	OrgID        apijson.Field
	Created      apijson.Field
	DeletedAt    apijson.Field
	Description  apijson.Field
	MemberGroups apijson.Field
	MemberUsers  apijson.Field
	UserID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Group) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r groupJSON) RawJSON() string {
	return r.raw
}

type GroupNewParams struct {
	// Name of the group
	Name param.Field[string] `json:"name,required"`
	// Textual description of the group
	Description param.Field[string] `json:"description"`
	// Ids of the groups this group inherits from
	//
	// An inheriting group has all the users contained in its member groups, as well as
	// all of their inherited users
	MemberGroups param.Field[[]string] `json:"member_groups" format:"uuid"`
	// Ids of users which belong to this group
	MemberUsers param.Field[[]string] `json:"member_users" format:"uuid"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the group belongs in.
	OrgName param.Field[string] `json:"org_name"`
}

func (r GroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GroupUpdateParams struct {
	// Textual description of the group
	Description param.Field[string] `json:"description"`
	// Ids of the groups this group inherits from
	//
	// An inheriting group has all the users contained in its member groups, as well as
	// all of their inherited users
	MemberGroups param.Field[[]string] `json:"member_groups" format:"uuid"`
	// Ids of users which belong to this group
	MemberUsers param.Field[[]string] `json:"member_users" format:"uuid"`
	// Name of the group
	Name param.Field[string] `json:"name"`
}

func (r GroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GroupListParams struct {
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Name of the group to search for
	GroupName param.Field[string] `query:"group_name"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[GroupListParamsIDsUnion] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[int64] `query:"limit"`
	// Filter search results to within a particular organization
	OrgName param.Field[string] `query:"org_name"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [GroupListParams]'s query parameters as `url.Values`.
func (r GroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [GroupListParamsIDsArray].
type GroupListParamsIDsUnion interface {
	ImplementsGroupListParamsIDsUnion()
}

type GroupListParamsIDsArray []string

func (r GroupListParamsIDsArray) ImplementsGroupListParamsIDsUnion() {}

type GroupReplaceParams struct {
	// Name of the group
	Name param.Field[string] `json:"name,required"`
	// Textual description of the group
	Description param.Field[string] `json:"description"`
	// Ids of the groups this group inherits from
	//
	// An inheriting group has all the users contained in its member groups, as well as
	// all of their inherited users
	MemberGroups param.Field[[]string] `json:"member_groups" format:"uuid"`
	// Ids of users which belong to this group
	MemberUsers param.Field[[]string] `json:"member_users" format:"uuid"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the group belongs in.
	OrgName param.Field[string] `json:"org_name"`
}

func (r GroupReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
