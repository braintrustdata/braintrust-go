// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/internal/apiquery"
	"github.com/braintrustdata/braintrust-go/internal/pagination"
	"github.com/braintrustdata/braintrust-go/internal/param"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
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
func NewGroupService(opts ...option.RequestOption) (r *GroupService) {
	r = &GroupService{}
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
func (r *GroupService) Get(ctx context.Context, groupID shared.GroupIDParam, opts ...option.RequestOption) (res *shared.Group, err error) {
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
func (r *GroupService) Update(ctx context.Context, groupID shared.GroupIDParam, body GroupUpdateParams, opts ...option.RequestOption) (res *shared.Group, err error) {
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
func (r *GroupService) Delete(ctx context.Context, groupID shared.GroupIDParam, opts ...option.RequestOption) (res *shared.Group, err error) {
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
	CreateGroup shared.CreateGroupParam `json:"create_group,required"`
}

func (r GroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateGroup)
}

type GroupUpdateParams struct {
	PatchGroup shared.PatchGroupParam `json:"patch_group,required"`
}

func (r GroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PatchGroup)
}

type GroupListParams struct {
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[shared.EndingBeforeParam] `query:"ending_before" format:"uuid"`
	// Name of the group to search for
	GroupName param.Field[shared.GroupNameParam] `query:"group_name"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[shared.IDsUnionParam] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[shared.AppLimitParam] `query:"limit"`
	// Filter search results to within a particular organization
	OrgName param.Field[shared.OrgNameParam] `query:"org_name"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Field[shared.StartingAfterParam] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [GroupListParams]'s query parameters as `url.Values`.
func (r GroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GroupReplaceParams struct {
	CreateGroup shared.CreateGroupParam `json:"create_group,required"`
}

func (r GroupReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateGroup)
}
