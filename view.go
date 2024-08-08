// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"errors"
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
func NewViewService(opts ...option.RequestOption) (r *ViewService) {
	r = &ViewService{}
	r.Options = opts
	return
}

// Create a new view. If there is an existing view in the project with the same
// name as the one specified in the request, will return the existing view
// unmodified
func (r *ViewService) New(ctx context.Context, body ViewNewParams, opts ...option.RequestOption) (res *View, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/view"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a view object by its id
func (r *ViewService) Get(ctx context.Context, viewID string, query ViewGetParams, opts ...option.RequestOption) (res *View, err error) {
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
func (r *ViewService) Update(ctx context.Context, viewID string, body ViewUpdateParams, opts ...option.RequestOption) (res *View, err error) {
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
func (r *ViewService) List(ctx context.Context, query ViewListParams, opts ...option.RequestOption) (res *pagination.ListObjects[View], err error) {
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
func (r *ViewService) ListAutoPaging(ctx context.Context, query ViewListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[View] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a view object by its id
func (r *ViewService) Delete(ctx context.Context, viewID string, body ViewDeleteParams, opts ...option.RequestOption) (res *View, err error) {
	opts = append(r.Options[:], opts...)
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("v1/view/%s", viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Create or replace view. If there is an existing view in the project with the
// same name as the one specified in the request, will replace the existing view
// with the provided fields
func (r *ViewService) Replace(ctx context.Context, body ViewReplaceParams, opts ...option.RequestOption) (res *View, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/view"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type View struct {
	// Unique identifier for the view
	ID string `json:"id,required" format:"uuid"`
	// Name of the view
	Name string `json:"name,required"`
	// The id of the object the view applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType ViewObjectType `json:"object_type,required,nullable"`
	// Type of table that the view corresponds to.
	ViewType ViewViewType `json:"view_type,required,nullable"`
	// Date of view creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Date of role deletion, or null if the role is still active
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Options for the view in the app
	Options ViewOptions `json:"options,nullable"`
	// Identifies the user who created the view
	UserID string `json:"user_id,nullable" format:"uuid"`
	// The view definition
	ViewData ViewData `json:"view_data,nullable"`
	JSON     viewJSON `json:"-"`
}

// viewJSON contains the JSON metadata for the struct [View]
type viewJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	ObjectID    apijson.Field
	ObjectType  apijson.Field
	ViewType    apijson.Field
	Created     apijson.Field
	DeletedAt   apijson.Field
	Options     apijson.Field
	UserID      apijson.Field
	ViewData    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *View) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewJSON) RawJSON() string {
	return r.raw
}

// The object type that the ACL applies to
type ViewObjectType string

const (
	ViewObjectTypeOrganization  ViewObjectType = "organization"
	ViewObjectTypeProject       ViewObjectType = "project"
	ViewObjectTypeExperiment    ViewObjectType = "experiment"
	ViewObjectTypeDataset       ViewObjectType = "dataset"
	ViewObjectTypePrompt        ViewObjectType = "prompt"
	ViewObjectTypePromptSession ViewObjectType = "prompt_session"
	ViewObjectTypeGroup         ViewObjectType = "group"
	ViewObjectTypeRole          ViewObjectType = "role"
	ViewObjectTypeOrgMember     ViewObjectType = "org_member"
	ViewObjectTypeProjectLog    ViewObjectType = "project_log"
	ViewObjectTypeOrgProject    ViewObjectType = "org_project"
)

func (r ViewObjectType) IsKnown() bool {
	switch r {
	case ViewObjectTypeOrganization, ViewObjectTypeProject, ViewObjectTypeExperiment, ViewObjectTypeDataset, ViewObjectTypePrompt, ViewObjectTypePromptSession, ViewObjectTypeGroup, ViewObjectTypeRole, ViewObjectTypeOrgMember, ViewObjectTypeProjectLog, ViewObjectTypeOrgProject:
		return true
	}
	return false
}

// Type of table that the view corresponds to.
type ViewViewType string

const (
	ViewViewTypeProjects    ViewViewType = "projects"
	ViewViewTypeLogs        ViewViewType = "logs"
	ViewViewTypeExperiments ViewViewType = "experiments"
	ViewViewTypeDatasets    ViewViewType = "datasets"
	ViewViewTypePrompts     ViewViewType = "prompts"
	ViewViewTypePlaygrounds ViewViewType = "playgrounds"
	ViewViewTypeExperiment  ViewViewType = "experiment"
	ViewViewTypeDataset     ViewViewType = "dataset"
)

func (r ViewViewType) IsKnown() bool {
	switch r {
	case ViewViewTypeProjects, ViewViewTypeLogs, ViewViewTypeExperiments, ViewViewTypeDatasets, ViewViewTypePrompts, ViewViewTypePlaygrounds, ViewViewTypeExperiment, ViewViewTypeDataset:
		return true
	}
	return false
}

// The view definition
type ViewData struct {
	Search ViewDataSearch `json:"search,nullable"`
	JSON   viewDataJSON   `json:"-"`
}

// viewDataJSON contains the JSON metadata for the struct [ViewData]
type viewDataJSON struct {
	Search      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ViewData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewDataJSON) RawJSON() string {
	return r.raw
}

// The view definition
type ViewDataParam struct {
	Search param.Field[ViewDataSearchParam] `json:"search"`
}

func (r ViewDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ViewDataSearch struct {
	Filter []interface{}      `json:"filter,nullable"`
	Match  []interface{}      `json:"match,nullable"`
	Sort   []interface{}      `json:"sort,nullable"`
	Tag    []interface{}      `json:"tag,nullable"`
	JSON   viewDataSearchJSON `json:"-"`
}

// viewDataSearchJSON contains the JSON metadata for the struct [ViewDataSearch]
type viewDataSearchJSON struct {
	Filter      apijson.Field
	Match       apijson.Field
	Sort        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ViewDataSearch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewDataSearchJSON) RawJSON() string {
	return r.raw
}

type ViewDataSearchParam struct {
	Filter param.Field[[]interface{}] `json:"filter"`
	Match  param.Field[[]interface{}] `json:"match"`
	Sort   param.Field[[]interface{}] `json:"sort"`
	Tag    param.Field[[]interface{}] `json:"tag"`
}

func (r ViewDataSearchParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options for the view in the app
type ViewOptions struct {
	ColumnOrder      []string           `json:"columnOrder,nullable"`
	ColumnSizing     map[string]float64 `json:"columnSizing,nullable"`
	ColumnVisibility map[string]bool    `json:"columnVisibility,nullable"`
	JSON             viewOptionsJSON    `json:"-"`
}

// viewOptionsJSON contains the JSON metadata for the struct [ViewOptions]
type viewOptionsJSON struct {
	ColumnOrder      apijson.Field
	ColumnSizing     apijson.Field
	ColumnVisibility apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ViewOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewOptionsJSON) RawJSON() string {
	return r.raw
}

// Options for the view in the app
type ViewOptionsParam struct {
	ColumnOrder      param.Field[[]string]           `json:"columnOrder"`
	ColumnSizing     param.Field[map[string]float64] `json:"columnSizing"`
	ColumnVisibility param.Field[map[string]bool]    `json:"columnVisibility"`
}

func (r ViewOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ViewNewParams struct {
	// Name of the view
	Name param.Field[string] `json:"name,required"`
	// The id of the object the view applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ViewNewParamsObjectType] `json:"object_type,required"`
	// Type of table that the view corresponds to.
	ViewType param.Field[ViewNewParamsViewType] `json:"view_type,required"`
	// Date of role deletion, or null if the role is still active
	DeletedAt param.Field[time.Time] `json:"deleted_at" format:"date-time"`
	// Options for the view in the app
	Options param.Field[ViewOptionsParam] `json:"options"`
	// Identifies the user who created the view
	UserID param.Field[string] `json:"user_id" format:"uuid"`
	// The view definition
	ViewData param.Field[ViewDataParam] `json:"view_data"`
}

func (r ViewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The object type that the ACL applies to
type ViewNewParamsObjectType string

const (
	ViewNewParamsObjectTypeOrganization  ViewNewParamsObjectType = "organization"
	ViewNewParamsObjectTypeProject       ViewNewParamsObjectType = "project"
	ViewNewParamsObjectTypeExperiment    ViewNewParamsObjectType = "experiment"
	ViewNewParamsObjectTypeDataset       ViewNewParamsObjectType = "dataset"
	ViewNewParamsObjectTypePrompt        ViewNewParamsObjectType = "prompt"
	ViewNewParamsObjectTypePromptSession ViewNewParamsObjectType = "prompt_session"
	ViewNewParamsObjectTypeGroup         ViewNewParamsObjectType = "group"
	ViewNewParamsObjectTypeRole          ViewNewParamsObjectType = "role"
	ViewNewParamsObjectTypeOrgMember     ViewNewParamsObjectType = "org_member"
	ViewNewParamsObjectTypeProjectLog    ViewNewParamsObjectType = "project_log"
	ViewNewParamsObjectTypeOrgProject    ViewNewParamsObjectType = "org_project"
)

func (r ViewNewParamsObjectType) IsKnown() bool {
	switch r {
	case ViewNewParamsObjectTypeOrganization, ViewNewParamsObjectTypeProject, ViewNewParamsObjectTypeExperiment, ViewNewParamsObjectTypeDataset, ViewNewParamsObjectTypePrompt, ViewNewParamsObjectTypePromptSession, ViewNewParamsObjectTypeGroup, ViewNewParamsObjectTypeRole, ViewNewParamsObjectTypeOrgMember, ViewNewParamsObjectTypeProjectLog, ViewNewParamsObjectTypeOrgProject:
		return true
	}
	return false
}

// Type of table that the view corresponds to.
type ViewNewParamsViewType string

const (
	ViewNewParamsViewTypeProjects    ViewNewParamsViewType = "projects"
	ViewNewParamsViewTypeLogs        ViewNewParamsViewType = "logs"
	ViewNewParamsViewTypeExperiments ViewNewParamsViewType = "experiments"
	ViewNewParamsViewTypeDatasets    ViewNewParamsViewType = "datasets"
	ViewNewParamsViewTypePrompts     ViewNewParamsViewType = "prompts"
	ViewNewParamsViewTypePlaygrounds ViewNewParamsViewType = "playgrounds"
	ViewNewParamsViewTypeExperiment  ViewNewParamsViewType = "experiment"
	ViewNewParamsViewTypeDataset     ViewNewParamsViewType = "dataset"
)

func (r ViewNewParamsViewType) IsKnown() bool {
	switch r {
	case ViewNewParamsViewTypeProjects, ViewNewParamsViewTypeLogs, ViewNewParamsViewTypeExperiments, ViewNewParamsViewTypeDatasets, ViewNewParamsViewTypePrompts, ViewNewParamsViewTypePlaygrounds, ViewNewParamsViewTypeExperiment, ViewNewParamsViewTypeDataset:
		return true
	}
	return false
}

type ViewGetParams struct {
	// The id of the object the ACL applies to
	ObjectID param.Field[string] `query:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ViewGetParamsObjectType] `query:"object_type,required"`
}

// URLQuery serializes [ViewGetParams]'s query parameters as `url.Values`.
func (r ViewGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The object type that the ACL applies to
type ViewGetParamsObjectType string

const (
	ViewGetParamsObjectTypeOrganization  ViewGetParamsObjectType = "organization"
	ViewGetParamsObjectTypeProject       ViewGetParamsObjectType = "project"
	ViewGetParamsObjectTypeExperiment    ViewGetParamsObjectType = "experiment"
	ViewGetParamsObjectTypeDataset       ViewGetParamsObjectType = "dataset"
	ViewGetParamsObjectTypePrompt        ViewGetParamsObjectType = "prompt"
	ViewGetParamsObjectTypePromptSession ViewGetParamsObjectType = "prompt_session"
	ViewGetParamsObjectTypeGroup         ViewGetParamsObjectType = "group"
	ViewGetParamsObjectTypeRole          ViewGetParamsObjectType = "role"
	ViewGetParamsObjectTypeOrgMember     ViewGetParamsObjectType = "org_member"
	ViewGetParamsObjectTypeProjectLog    ViewGetParamsObjectType = "project_log"
	ViewGetParamsObjectTypeOrgProject    ViewGetParamsObjectType = "org_project"
)

func (r ViewGetParamsObjectType) IsKnown() bool {
	switch r {
	case ViewGetParamsObjectTypeOrganization, ViewGetParamsObjectTypeProject, ViewGetParamsObjectTypeExperiment, ViewGetParamsObjectTypeDataset, ViewGetParamsObjectTypePrompt, ViewGetParamsObjectTypePromptSession, ViewGetParamsObjectTypeGroup, ViewGetParamsObjectTypeRole, ViewGetParamsObjectTypeOrgMember, ViewGetParamsObjectTypeProjectLog, ViewGetParamsObjectTypeOrgProject:
		return true
	}
	return false
}

type ViewUpdateParams struct {
	// The id of the object the view applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ViewUpdateParamsObjectType] `json:"object_type,required"`
	// Name of the view
	Name param.Field[string] `json:"name"`
	// Options for the view in the app
	Options param.Field[ViewOptionsParam] `json:"options"`
	// Identifies the user who created the view
	UserID param.Field[string] `json:"user_id" format:"uuid"`
	// The view definition
	ViewData param.Field[ViewDataParam] `json:"view_data"`
	// Type of table that the view corresponds to.
	ViewType param.Field[ViewUpdateParamsViewType] `json:"view_type"`
}

func (r ViewUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The object type that the ACL applies to
type ViewUpdateParamsObjectType string

const (
	ViewUpdateParamsObjectTypeOrganization  ViewUpdateParamsObjectType = "organization"
	ViewUpdateParamsObjectTypeProject       ViewUpdateParamsObjectType = "project"
	ViewUpdateParamsObjectTypeExperiment    ViewUpdateParamsObjectType = "experiment"
	ViewUpdateParamsObjectTypeDataset       ViewUpdateParamsObjectType = "dataset"
	ViewUpdateParamsObjectTypePrompt        ViewUpdateParamsObjectType = "prompt"
	ViewUpdateParamsObjectTypePromptSession ViewUpdateParamsObjectType = "prompt_session"
	ViewUpdateParamsObjectTypeGroup         ViewUpdateParamsObjectType = "group"
	ViewUpdateParamsObjectTypeRole          ViewUpdateParamsObjectType = "role"
	ViewUpdateParamsObjectTypeOrgMember     ViewUpdateParamsObjectType = "org_member"
	ViewUpdateParamsObjectTypeProjectLog    ViewUpdateParamsObjectType = "project_log"
	ViewUpdateParamsObjectTypeOrgProject    ViewUpdateParamsObjectType = "org_project"
)

func (r ViewUpdateParamsObjectType) IsKnown() bool {
	switch r {
	case ViewUpdateParamsObjectTypeOrganization, ViewUpdateParamsObjectTypeProject, ViewUpdateParamsObjectTypeExperiment, ViewUpdateParamsObjectTypeDataset, ViewUpdateParamsObjectTypePrompt, ViewUpdateParamsObjectTypePromptSession, ViewUpdateParamsObjectTypeGroup, ViewUpdateParamsObjectTypeRole, ViewUpdateParamsObjectTypeOrgMember, ViewUpdateParamsObjectTypeProjectLog, ViewUpdateParamsObjectTypeOrgProject:
		return true
	}
	return false
}

// Type of table that the view corresponds to.
type ViewUpdateParamsViewType string

const (
	ViewUpdateParamsViewTypeProjects    ViewUpdateParamsViewType = "projects"
	ViewUpdateParamsViewTypeLogs        ViewUpdateParamsViewType = "logs"
	ViewUpdateParamsViewTypeExperiments ViewUpdateParamsViewType = "experiments"
	ViewUpdateParamsViewTypeDatasets    ViewUpdateParamsViewType = "datasets"
	ViewUpdateParamsViewTypePrompts     ViewUpdateParamsViewType = "prompts"
	ViewUpdateParamsViewTypePlaygrounds ViewUpdateParamsViewType = "playgrounds"
	ViewUpdateParamsViewTypeExperiment  ViewUpdateParamsViewType = "experiment"
	ViewUpdateParamsViewTypeDataset     ViewUpdateParamsViewType = "dataset"
)

func (r ViewUpdateParamsViewType) IsKnown() bool {
	switch r {
	case ViewUpdateParamsViewTypeProjects, ViewUpdateParamsViewTypeLogs, ViewUpdateParamsViewTypeExperiments, ViewUpdateParamsViewTypeDatasets, ViewUpdateParamsViewTypePrompts, ViewUpdateParamsViewTypePlaygrounds, ViewUpdateParamsViewTypeExperiment, ViewUpdateParamsViewTypeDataset:
		return true
	}
	return false
}

type ViewListParams struct {
	// The id of the object the ACL applies to
	ObjectID param.Field[string] `query:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ViewListParamsObjectType] `query:"object_type,required"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[ViewListParamsIDsUnion] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[int64] `query:"limit"`
	// Name of the project to search for
	ProjectName param.Field[string] `query:"project_name"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
	// Name of the view to search for
	ViewName param.Field[string] `query:"view_name"`
	// Type of table that the view corresponds to.
	ViewType param.Field[ViewListParamsViewType] `query:"view_type"`
}

// URLQuery serializes [ViewListParams]'s query parameters as `url.Values`.
func (r ViewListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The object type that the ACL applies to
type ViewListParamsObjectType string

const (
	ViewListParamsObjectTypeOrganization  ViewListParamsObjectType = "organization"
	ViewListParamsObjectTypeProject       ViewListParamsObjectType = "project"
	ViewListParamsObjectTypeExperiment    ViewListParamsObjectType = "experiment"
	ViewListParamsObjectTypeDataset       ViewListParamsObjectType = "dataset"
	ViewListParamsObjectTypePrompt        ViewListParamsObjectType = "prompt"
	ViewListParamsObjectTypePromptSession ViewListParamsObjectType = "prompt_session"
	ViewListParamsObjectTypeGroup         ViewListParamsObjectType = "group"
	ViewListParamsObjectTypeRole          ViewListParamsObjectType = "role"
	ViewListParamsObjectTypeOrgMember     ViewListParamsObjectType = "org_member"
	ViewListParamsObjectTypeProjectLog    ViewListParamsObjectType = "project_log"
	ViewListParamsObjectTypeOrgProject    ViewListParamsObjectType = "org_project"
)

func (r ViewListParamsObjectType) IsKnown() bool {
	switch r {
	case ViewListParamsObjectTypeOrganization, ViewListParamsObjectTypeProject, ViewListParamsObjectTypeExperiment, ViewListParamsObjectTypeDataset, ViewListParamsObjectTypePrompt, ViewListParamsObjectTypePromptSession, ViewListParamsObjectTypeGroup, ViewListParamsObjectTypeRole, ViewListParamsObjectTypeOrgMember, ViewListParamsObjectTypeProjectLog, ViewListParamsObjectTypeOrgProject:
		return true
	}
	return false
}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [ViewListParamsIDsArray].
type ViewListParamsIDsUnion interface {
	ImplementsViewListParamsIDsUnion()
}

type ViewListParamsIDsArray []string

func (r ViewListParamsIDsArray) ImplementsViewListParamsIDsUnion() {}

// Type of table that the view corresponds to.
type ViewListParamsViewType string

const (
	ViewListParamsViewTypeProjects    ViewListParamsViewType = "projects"
	ViewListParamsViewTypeLogs        ViewListParamsViewType = "logs"
	ViewListParamsViewTypeExperiments ViewListParamsViewType = "experiments"
	ViewListParamsViewTypeDatasets    ViewListParamsViewType = "datasets"
	ViewListParamsViewTypePrompts     ViewListParamsViewType = "prompts"
	ViewListParamsViewTypePlaygrounds ViewListParamsViewType = "playgrounds"
	ViewListParamsViewTypeExperiment  ViewListParamsViewType = "experiment"
	ViewListParamsViewTypeDataset     ViewListParamsViewType = "dataset"
)

func (r ViewListParamsViewType) IsKnown() bool {
	switch r {
	case ViewListParamsViewTypeProjects, ViewListParamsViewTypeLogs, ViewListParamsViewTypeExperiments, ViewListParamsViewTypeDatasets, ViewListParamsViewTypePrompts, ViewListParamsViewTypePlaygrounds, ViewListParamsViewTypeExperiment, ViewListParamsViewTypeDataset:
		return true
	}
	return false
}

type ViewDeleteParams struct {
	// The id of the object the view applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ViewDeleteParamsObjectType] `json:"object_type,required"`
}

func (r ViewDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The object type that the ACL applies to
type ViewDeleteParamsObjectType string

const (
	ViewDeleteParamsObjectTypeOrganization  ViewDeleteParamsObjectType = "organization"
	ViewDeleteParamsObjectTypeProject       ViewDeleteParamsObjectType = "project"
	ViewDeleteParamsObjectTypeExperiment    ViewDeleteParamsObjectType = "experiment"
	ViewDeleteParamsObjectTypeDataset       ViewDeleteParamsObjectType = "dataset"
	ViewDeleteParamsObjectTypePrompt        ViewDeleteParamsObjectType = "prompt"
	ViewDeleteParamsObjectTypePromptSession ViewDeleteParamsObjectType = "prompt_session"
	ViewDeleteParamsObjectTypeGroup         ViewDeleteParamsObjectType = "group"
	ViewDeleteParamsObjectTypeRole          ViewDeleteParamsObjectType = "role"
	ViewDeleteParamsObjectTypeOrgMember     ViewDeleteParamsObjectType = "org_member"
	ViewDeleteParamsObjectTypeProjectLog    ViewDeleteParamsObjectType = "project_log"
	ViewDeleteParamsObjectTypeOrgProject    ViewDeleteParamsObjectType = "org_project"
)

func (r ViewDeleteParamsObjectType) IsKnown() bool {
	switch r {
	case ViewDeleteParamsObjectTypeOrganization, ViewDeleteParamsObjectTypeProject, ViewDeleteParamsObjectTypeExperiment, ViewDeleteParamsObjectTypeDataset, ViewDeleteParamsObjectTypePrompt, ViewDeleteParamsObjectTypePromptSession, ViewDeleteParamsObjectTypeGroup, ViewDeleteParamsObjectTypeRole, ViewDeleteParamsObjectTypeOrgMember, ViewDeleteParamsObjectTypeProjectLog, ViewDeleteParamsObjectTypeOrgProject:
		return true
	}
	return false
}

type ViewReplaceParams struct {
	// Name of the view
	Name param.Field[string] `json:"name,required"`
	// The id of the object the view applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ViewReplaceParamsObjectType] `json:"object_type,required"`
	// Type of table that the view corresponds to.
	ViewType param.Field[ViewReplaceParamsViewType] `json:"view_type,required"`
	// Date of role deletion, or null if the role is still active
	DeletedAt param.Field[time.Time] `json:"deleted_at" format:"date-time"`
	// Options for the view in the app
	Options param.Field[ViewOptionsParam] `json:"options"`
	// Identifies the user who created the view
	UserID param.Field[string] `json:"user_id" format:"uuid"`
	// The view definition
	ViewData param.Field[ViewDataParam] `json:"view_data"`
}

func (r ViewReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The object type that the ACL applies to
type ViewReplaceParamsObjectType string

const (
	ViewReplaceParamsObjectTypeOrganization  ViewReplaceParamsObjectType = "organization"
	ViewReplaceParamsObjectTypeProject       ViewReplaceParamsObjectType = "project"
	ViewReplaceParamsObjectTypeExperiment    ViewReplaceParamsObjectType = "experiment"
	ViewReplaceParamsObjectTypeDataset       ViewReplaceParamsObjectType = "dataset"
	ViewReplaceParamsObjectTypePrompt        ViewReplaceParamsObjectType = "prompt"
	ViewReplaceParamsObjectTypePromptSession ViewReplaceParamsObjectType = "prompt_session"
	ViewReplaceParamsObjectTypeGroup         ViewReplaceParamsObjectType = "group"
	ViewReplaceParamsObjectTypeRole          ViewReplaceParamsObjectType = "role"
	ViewReplaceParamsObjectTypeOrgMember     ViewReplaceParamsObjectType = "org_member"
	ViewReplaceParamsObjectTypeProjectLog    ViewReplaceParamsObjectType = "project_log"
	ViewReplaceParamsObjectTypeOrgProject    ViewReplaceParamsObjectType = "org_project"
)

func (r ViewReplaceParamsObjectType) IsKnown() bool {
	switch r {
	case ViewReplaceParamsObjectTypeOrganization, ViewReplaceParamsObjectTypeProject, ViewReplaceParamsObjectTypeExperiment, ViewReplaceParamsObjectTypeDataset, ViewReplaceParamsObjectTypePrompt, ViewReplaceParamsObjectTypePromptSession, ViewReplaceParamsObjectTypeGroup, ViewReplaceParamsObjectTypeRole, ViewReplaceParamsObjectTypeOrgMember, ViewReplaceParamsObjectTypeProjectLog, ViewReplaceParamsObjectTypeOrgProject:
		return true
	}
	return false
}

// Type of table that the view corresponds to.
type ViewReplaceParamsViewType string

const (
	ViewReplaceParamsViewTypeProjects    ViewReplaceParamsViewType = "projects"
	ViewReplaceParamsViewTypeLogs        ViewReplaceParamsViewType = "logs"
	ViewReplaceParamsViewTypeExperiments ViewReplaceParamsViewType = "experiments"
	ViewReplaceParamsViewTypeDatasets    ViewReplaceParamsViewType = "datasets"
	ViewReplaceParamsViewTypePrompts     ViewReplaceParamsViewType = "prompts"
	ViewReplaceParamsViewTypePlaygrounds ViewReplaceParamsViewType = "playgrounds"
	ViewReplaceParamsViewTypeExperiment  ViewReplaceParamsViewType = "experiment"
	ViewReplaceParamsViewTypeDataset     ViewReplaceParamsViewType = "dataset"
)

func (r ViewReplaceParamsViewType) IsKnown() bool {
	switch r {
	case ViewReplaceParamsViewTypeProjects, ViewReplaceParamsViewTypeLogs, ViewReplaceParamsViewTypeExperiments, ViewReplaceParamsViewTypeDatasets, ViewReplaceParamsViewTypePrompts, ViewReplaceParamsViewTypePlaygrounds, ViewReplaceParamsViewTypeExperiment, ViewReplaceParamsViewTypeDataset:
		return true
	}
	return false
}
