// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/braintrustdata/braintrust-go/internal/apiquery"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/braintrustdata/braintrust-go/packages/pagination"
	"github.com/braintrustdata/braintrust-go/packages/param"
	"github.com/braintrustdata/braintrust-go/shared"
)

// ViewService contains methods and other services that help with interacting with
// the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewViewService] method instead.
type ViewService struct {
	Options []option.RequestOption
}

// NewViewService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewViewService(opts ...option.RequestOption) (r ViewService) {
	r = ViewService{}
	r.Options = opts
	return
}

// Create a new view. If there is an existing view with the same name as the one
// specified in the request, will return the existing view unmodified
func (r *ViewService) New(ctx context.Context, body ViewNewParams, opts ...option.RequestOption) (res *shared.View, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/view"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a view object by its id
func (r *ViewService) Get(ctx context.Context, viewID string, query ViewGetParams, opts ...option.RequestOption) (res *shared.View, err error) {
	opts = append(r.Options[:], opts...)
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("v1/view/%s", viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Partially update a view object. Specify the fields to update in the payload. Any
// object-type fields will be deep-merged with existing content. Currently we do
// not support removing fields or setting them to null.
func (r *ViewService) Update(ctx context.Context, viewID string, body ViewUpdateParams, opts ...option.RequestOption) (res *shared.View, err error) {
	opts = append(r.Options[:], opts...)
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("v1/view/%s", viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all views. The views are sorted by creation date, with the most
// recently-created views coming first
func (r *ViewService) List(ctx context.Context, query ViewListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.View], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/view"
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

// List out all views. The views are sorted by creation date, with the most
// recently-created views coming first
func (r *ViewService) ListAutoPaging(ctx context.Context, query ViewListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.View] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a view object by its id
func (r *ViewService) Delete(ctx context.Context, viewID string, body ViewDeleteParams, opts ...option.RequestOption) (res *shared.View, err error) {
	opts = append(r.Options[:], opts...)
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("v1/view/%s", viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Create or replace view. If there is an existing view with the same name as the
// one specified in the request, will replace the existing view with the provided
// fields
func (r *ViewService) Replace(ctx context.Context, body ViewReplaceParams, opts ...option.RequestOption) (res *shared.View, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/view"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ViewNewParams struct {
	// Type of table that the view corresponds to.
	//
	// Any of "projects", "experiments", "experiment", "playgrounds", "playground",
	// "datasets", "dataset", "prompts", "tools", "scorers", "logs".
	ViewType ViewNewParamsViewType `json:"view_type,omitzero,required"`
	// Name of the view
	Name string `json:"name,required"`
	// The id of the object the view applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType shared.ACLObjectType `json:"object_type,omitzero,required"`
	// Date of role deletion, or null if the role is still active
	DeletedAt param.Opt[time.Time] `json:"deleted_at,omitzero" format:"date-time"`
	// Identifies the user who created the view
	UserID param.Opt[string] `json:"user_id,omitzero" format:"uuid"`
	// Options for the view in the app
	Options shared.ViewOptionsParam `json:"options,omitzero"`
	// The view definition
	ViewData shared.ViewDataParam `json:"view_data,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ViewNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ViewNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ViewNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Type of table that the view corresponds to.
type ViewNewParamsViewType string

const (
	ViewNewParamsViewTypeProjects    ViewNewParamsViewType = "projects"
	ViewNewParamsViewTypeExperiments ViewNewParamsViewType = "experiments"
	ViewNewParamsViewTypeExperiment  ViewNewParamsViewType = "experiment"
	ViewNewParamsViewTypePlaygrounds ViewNewParamsViewType = "playgrounds"
	ViewNewParamsViewTypePlayground  ViewNewParamsViewType = "playground"
	ViewNewParamsViewTypeDatasets    ViewNewParamsViewType = "datasets"
	ViewNewParamsViewTypeDataset     ViewNewParamsViewType = "dataset"
	ViewNewParamsViewTypePrompts     ViewNewParamsViewType = "prompts"
	ViewNewParamsViewTypeTools       ViewNewParamsViewType = "tools"
	ViewNewParamsViewTypeScorers     ViewNewParamsViewType = "scorers"
	ViewNewParamsViewTypeLogs        ViewNewParamsViewType = "logs"
)

type ViewGetParams struct {
	// The id of the object the ACL applies to
	ObjectID string `query:"object_id,required" format:"uuid" json:"-"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType shared.ACLObjectType `query:"object_type,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ViewGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ViewGetParams]'s query parameters as `url.Values`.
func (r ViewGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ViewUpdateParams struct {
	// The id of the object the view applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType shared.ACLObjectType `json:"object_type,omitzero,required"`
	// Name of the view
	Name param.Opt[string] `json:"name,omitzero"`
	// Identifies the user who created the view
	UserID param.Opt[string] `json:"user_id,omitzero" format:"uuid"`
	// Options for the view in the app
	Options shared.ViewOptionsParam `json:"options,omitzero"`
	// The view definition
	ViewData shared.ViewDataParam `json:"view_data,omitzero"`
	// Type of table that the view corresponds to.
	//
	// Any of "projects", "experiments", "experiment", "playgrounds", "playground",
	// "datasets", "dataset", "prompts", "tools", "scorers", "logs".
	ViewType ViewUpdateParamsViewType `json:"view_type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ViewUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ViewUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ViewUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Type of table that the view corresponds to.
type ViewUpdateParamsViewType string

const (
	ViewUpdateParamsViewTypeProjects    ViewUpdateParamsViewType = "projects"
	ViewUpdateParamsViewTypeExperiments ViewUpdateParamsViewType = "experiments"
	ViewUpdateParamsViewTypeExperiment  ViewUpdateParamsViewType = "experiment"
	ViewUpdateParamsViewTypePlaygrounds ViewUpdateParamsViewType = "playgrounds"
	ViewUpdateParamsViewTypePlayground  ViewUpdateParamsViewType = "playground"
	ViewUpdateParamsViewTypeDatasets    ViewUpdateParamsViewType = "datasets"
	ViewUpdateParamsViewTypeDataset     ViewUpdateParamsViewType = "dataset"
	ViewUpdateParamsViewTypePrompts     ViewUpdateParamsViewType = "prompts"
	ViewUpdateParamsViewTypeTools       ViewUpdateParamsViewType = "tools"
	ViewUpdateParamsViewTypeScorers     ViewUpdateParamsViewType = "scorers"
	ViewUpdateParamsViewTypeLogs        ViewUpdateParamsViewType = "logs"
)

type ViewListParams struct {
	// The id of the object the ACL applies to
	ObjectID string `query:"object_id,required" format:"uuid" json:"-"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType shared.ACLObjectType `query:"object_type,omitzero,required" json:"-"`
	// Limit the number of objects to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Opt[string] `query:"ending_before,omitzero" format:"uuid" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Opt[string] `query:"starting_after,omitzero" format:"uuid" json:"-"`
	// Name of the view to search for
	ViewName param.Opt[string] `query:"view_name,omitzero" json:"-"`
	// Type of table that the view corresponds to.
	//
	// Any of "projects", "experiments", "experiment", "playgrounds", "playground",
	// "datasets", "dataset", "prompts", "tools", "scorers", "logs".
	ViewType shared.ViewType `query:"view_type,omitzero" json:"-"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs ViewListParamsIDsUnion `query:"ids,omitzero" format:"uuid" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ViewListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ViewListParams]'s query parameters as `url.Values`.
func (r ViewListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ViewListParamsIDsUnion struct {
	OfString            param.Opt[string] `json:",omitzero,inline"`
	OfViewListsIDsArray []string          `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u ViewListParamsIDsUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u ViewListParamsIDsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[ViewListParamsIDsUnion](u.OfString, u.OfViewListsIDsArray)
}

func (u *ViewListParamsIDsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfViewListsIDsArray) {
		return &u.OfViewListsIDsArray
	}
	return nil
}

type ViewDeleteParams struct {
	// The id of the object the view applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType shared.ACLObjectType `json:"object_type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ViewDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ViewDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow ViewDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type ViewReplaceParams struct {
	// Type of table that the view corresponds to.
	//
	// Any of "projects", "experiments", "experiment", "playgrounds", "playground",
	// "datasets", "dataset", "prompts", "tools", "scorers", "logs".
	ViewType ViewReplaceParamsViewType `json:"view_type,omitzero,required"`
	// Name of the view
	Name string `json:"name,required"`
	// The id of the object the view applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType shared.ACLObjectType `json:"object_type,omitzero,required"`
	// Date of role deletion, or null if the role is still active
	DeletedAt param.Opt[time.Time] `json:"deleted_at,omitzero" format:"date-time"`
	// Identifies the user who created the view
	UserID param.Opt[string] `json:"user_id,omitzero" format:"uuid"`
	// Options for the view in the app
	Options shared.ViewOptionsParam `json:"options,omitzero"`
	// The view definition
	ViewData shared.ViewDataParam `json:"view_data,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ViewReplaceParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ViewReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow ViewReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Type of table that the view corresponds to.
type ViewReplaceParamsViewType string

const (
	ViewReplaceParamsViewTypeProjects    ViewReplaceParamsViewType = "projects"
	ViewReplaceParamsViewTypeExperiments ViewReplaceParamsViewType = "experiments"
	ViewReplaceParamsViewTypeExperiment  ViewReplaceParamsViewType = "experiment"
	ViewReplaceParamsViewTypePlaygrounds ViewReplaceParamsViewType = "playgrounds"
	ViewReplaceParamsViewTypePlayground  ViewReplaceParamsViewType = "playground"
	ViewReplaceParamsViewTypeDatasets    ViewReplaceParamsViewType = "datasets"
	ViewReplaceParamsViewTypeDataset     ViewReplaceParamsViewType = "dataset"
	ViewReplaceParamsViewTypePrompts     ViewReplaceParamsViewType = "prompts"
	ViewReplaceParamsViewTypeTools       ViewReplaceParamsViewType = "tools"
	ViewReplaceParamsViewTypeScorers     ViewReplaceParamsViewType = "scorers"
	ViewReplaceParamsViewTypeLogs        ViewReplaceParamsViewType = "logs"
)
