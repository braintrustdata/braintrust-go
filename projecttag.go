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
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/braintrustdata/braintrust-go/packages/pagination"
	"github.com/braintrustdata/braintrust-go/packages/param"
	"github.com/braintrustdata/braintrust-go/shared"
)

// ProjectTagService contains methods and other services that help with interacting
// with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectTagService] method instead.
type ProjectTagService struct {
	Options []option.RequestOption
}

// NewProjectTagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectTagService(opts ...option.RequestOption) (r ProjectTagService) {
	r = ProjectTagService{}
	r.Options = opts
	return
}

// Create a new project_tag. If there is an existing project_tag in the project
// with the same name as the one specified in the request, will return the existing
// project_tag unmodified
func (r *ProjectTagService) New(ctx context.Context, body ProjectTagNewParams, opts ...option.RequestOption) (res *shared.ProjectTag, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/project_tag"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a project_tag object by its id
func (r *ProjectTagService) Get(ctx context.Context, projectTagID string, opts ...option.RequestOption) (res *shared.ProjectTag, err error) {
	opts = append(r.Options[:], opts...)
	if projectTagID == "" {
		err = errors.New("missing required project_tag_id parameter")
		return
	}
	path := fmt.Sprintf("v1/project_tag/%s", projectTagID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update a project_tag object. Specify the fields to update in the
// payload. Any object-type fields will be deep-merged with existing content.
// Currently we do not support removing fields or setting them to null.
func (r *ProjectTagService) Update(ctx context.Context, projectTagID string, body ProjectTagUpdateParams, opts ...option.RequestOption) (res *shared.ProjectTag, err error) {
	opts = append(r.Options[:], opts...)
	if projectTagID == "" {
		err = errors.New("missing required project_tag_id parameter")
		return
	}
	path := fmt.Sprintf("v1/project_tag/%s", projectTagID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all project_tags. The project_tags are sorted by creation date, with
// the most recently-created project_tags coming first
func (r *ProjectTagService) List(ctx context.Context, query ProjectTagListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.ProjectTag], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/project_tag"
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

// List out all project_tags. The project_tags are sorted by creation date, with
// the most recently-created project_tags coming first
func (r *ProjectTagService) ListAutoPaging(ctx context.Context, query ProjectTagListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.ProjectTag] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a project_tag object by its id
func (r *ProjectTagService) Delete(ctx context.Context, projectTagID string, opts ...option.RequestOption) (res *shared.ProjectTag, err error) {
	opts = append(r.Options[:], opts...)
	if projectTagID == "" {
		err = errors.New("missing required project_tag_id parameter")
		return
	}
	path := fmt.Sprintf("v1/project_tag/%s", projectTagID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create or replace project_tag. If there is an existing project_tag in the
// project with the same name as the one specified in the request, will replace the
// existing project_tag with the provided fields
func (r *ProjectTagService) Replace(ctx context.Context, body ProjectTagReplaceParams, opts ...option.RequestOption) (res *shared.ProjectTag, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/project_tag"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ProjectTagNewParams struct {
	// Name of the project tag
	Name string `json:"name,required"`
	// Unique identifier for the project that the project tag belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// Color of the tag for the UI
	Color param.Opt[string] `json:"color,omitzero"`
	// Textual description of the project tag
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ProjectTagNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectTagNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectTagNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectTagUpdateParams struct {
	// Color of the tag for the UI
	Color param.Opt[string] `json:"color,omitzero"`
	// Textual description of the project tag
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the project tag
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ProjectTagUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectTagUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectTagUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectTagListParams struct {
	// Limit the number of objects to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Opt[string] `query:"ending_before,omitzero" format:"uuid" json:"-"`
	// Filter search results to within a particular organization
	OrgName param.Opt[string] `query:"org_name,omitzero" json:"-"`
	// Project id
	ProjectID param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Name of the project to search for
	ProjectName param.Opt[string] `query:"project_name,omitzero" json:"-"`
	// Name of the project_tag to search for
	ProjectTagName param.Opt[string] `query:"project_tag_name,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Opt[string] `query:"starting_after,omitzero" format:"uuid" json:"-"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs ProjectTagListParamsIDsUnion `query:"ids,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ProjectTagListParams]'s query parameters as `url.Values`.
func (r ProjectTagListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProjectTagListParamsIDsUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *ProjectTagListParamsIDsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type ProjectTagReplaceParams struct {
	// Name of the project tag
	Name string `json:"name,required"`
	// Unique identifier for the project that the project tag belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// Color of the tag for the UI
	Color param.Opt[string] `json:"color,omitzero"`
	// Textual description of the project tag
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ProjectTagReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectTagReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectTagReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
