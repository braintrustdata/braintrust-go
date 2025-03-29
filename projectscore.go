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

// ProjectScoreService contains methods and other services that help with
// interacting with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectScoreService] method instead.
type ProjectScoreService struct {
	Options []option.RequestOption
}

// NewProjectScoreService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectScoreService(opts ...option.RequestOption) (r ProjectScoreService) {
	r = ProjectScoreService{}
	r.Options = opts
	return
}

// Create a new project_score. If there is an existing project_score in the project
// with the same name as the one specified in the request, will return the existing
// project_score unmodified
func (r *ProjectScoreService) New(ctx context.Context, body ProjectScoreNewParams, opts ...option.RequestOption) (res *shared.ProjectScore, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/project_score"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a project_score object by its id
func (r *ProjectScoreService) Get(ctx context.Context, projectScoreID string, opts ...option.RequestOption) (res *shared.ProjectScore, err error) {
	opts = append(r.Options[:], opts...)
	if projectScoreID == "" {
		err = errors.New("missing required project_score_id parameter")
		return
	}
	path := fmt.Sprintf("v1/project_score/%s", projectScoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update a project_score object. Specify the fields to update in the
// payload. Any object-type fields will be deep-merged with existing content.
// Currently we do not support removing fields or setting them to null.
func (r *ProjectScoreService) Update(ctx context.Context, projectScoreID string, body ProjectScoreUpdateParams, opts ...option.RequestOption) (res *shared.ProjectScore, err error) {
	opts = append(r.Options[:], opts...)
	if projectScoreID == "" {
		err = errors.New("missing required project_score_id parameter")
		return
	}
	path := fmt.Sprintf("v1/project_score/%s", projectScoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all project_scores. The project_scores are sorted by creation date,
// with the most recently-created project_scores coming first
func (r *ProjectScoreService) List(ctx context.Context, query ProjectScoreListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.ProjectScore], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/project_score"
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

// List out all project_scores. The project_scores are sorted by creation date,
// with the most recently-created project_scores coming first
func (r *ProjectScoreService) ListAutoPaging(ctx context.Context, query ProjectScoreListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.ProjectScore] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a project_score object by its id
func (r *ProjectScoreService) Delete(ctx context.Context, projectScoreID string, opts ...option.RequestOption) (res *shared.ProjectScore, err error) {
	opts = append(r.Options[:], opts...)
	if projectScoreID == "" {
		err = errors.New("missing required project_score_id parameter")
		return
	}
	path := fmt.Sprintf("v1/project_score/%s", projectScoreID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create or replace project_score. If there is an existing project_score in the
// project with the same name as the one specified in the request, will replace the
// existing project_score with the provided fields
func (r *ProjectScoreService) Replace(ctx context.Context, body ProjectScoreReplaceParams, opts ...option.RequestOption) (res *shared.ProjectScore, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/project_score"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ProjectScoreNewParams struct {
	// Name of the project score
	Name string `json:"name,required"`
	// Unique identifier for the project that the project score belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// The type of the configured score
	//
	// Any of "slider", "categorical", "weighted", "minimum", "maximum", "online",
	// "free-form".
	ScoreType shared.ProjectScoreType `json:"score_type,omitzero,required"`
	// Textual description of the project score
	Description param.Opt[string] `json:"description,omitzero"`
	// For categorical-type project scores, the list of all categories
	Categories ProjectScoreNewParamsCategoriesUnion `json:"categories,omitzero"`
	Config     shared.ProjectScoreConfigParam       `json:"config,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectScoreNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ProjectScoreNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectScoreNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProjectScoreNewParamsCategoriesUnion struct {
	OfCategorical []shared.ProjectScoreCategoryParam `json:",omitzero,inline"`
	OfMinimum     []string                           `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u ProjectScoreNewParamsCategoriesUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u ProjectScoreNewParamsCategoriesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[ProjectScoreNewParamsCategoriesUnion](u.OfCategorical, u.OfMinimum)
}

func (u *ProjectScoreNewParamsCategoriesUnion) asAny() any {
	if !param.IsOmitted(u.OfCategorical) {
		return &u.OfCategorical
	} else if !param.IsOmitted(u.OfMinimum) {
		return &u.OfMinimum
	}
	return nil
}

type ProjectScoreUpdateParams struct {
	// Textual description of the project score
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the project score
	Name param.Opt[string] `json:"name,omitzero"`
	// For categorical-type project scores, the list of all categories
	Categories ProjectScoreUpdateParamsCategoriesUnion `json:"categories,omitzero"`
	Config     shared.ProjectScoreConfigParam          `json:"config,omitzero"`
	// The type of the configured score
	//
	// Any of "slider", "categorical", "weighted", "minimum", "maximum", "online",
	// "free-form".
	ScoreType shared.ProjectScoreType `json:"score_type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectScoreUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ProjectScoreUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectScoreUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProjectScoreUpdateParamsCategoriesUnion struct {
	OfCategorical []shared.ProjectScoreCategoryParam `json:",omitzero,inline"`
	OfMinimum     []string                           `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u ProjectScoreUpdateParamsCategoriesUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u ProjectScoreUpdateParamsCategoriesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[ProjectScoreUpdateParamsCategoriesUnion](u.OfCategorical, u.OfMinimum)
}

func (u *ProjectScoreUpdateParamsCategoriesUnion) asAny() any {
	if !param.IsOmitted(u.OfCategorical) {
		return &u.OfCategorical
	} else if !param.IsOmitted(u.OfMinimum) {
		return &u.OfMinimum
	}
	return nil
}

type ProjectScoreListParams struct {
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
	// Name of the project_score to search for
	ProjectScoreName param.Opt[string] `query:"project_score_name,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Opt[string] `query:"starting_after,omitzero" format:"uuid" json:"-"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs ProjectScoreListParamsIDsUnion `query:"ids,omitzero" format:"uuid" json:"-"`
	// The type of the configured score
	ScoreType ProjectScoreListParamsScoreTypeUnion `query:"score_type,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectScoreListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ProjectScoreListParams]'s query parameters as `url.Values`.
func (r ProjectScoreListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProjectScoreListParamsIDsUnion struct {
	OfString                    param.Opt[string] `json:",omitzero,inline"`
	OfProjectScoreListsIDsArray []string          `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u ProjectScoreListParamsIDsUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u ProjectScoreListParamsIDsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[ProjectScoreListParamsIDsUnion](u.OfString, u.OfProjectScoreListsIDsArray)
}

func (u *ProjectScoreListParamsIDsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfProjectScoreListsIDsArray) {
		return &u.OfProjectScoreListsIDsArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProjectScoreListParamsScoreTypeUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfProjectScoreTypeSingle)
	OfProjectScoreTypeSingle          param.Opt[shared.ProjectScoreType] `json:",omitzero,inline"`
	OfProjectScoreListsScoreTypeArray []string                           `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u ProjectScoreListParamsScoreTypeUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u ProjectScoreListParamsScoreTypeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[ProjectScoreListParamsScoreTypeUnion](u.OfProjectScoreTypeSingle, u.OfProjectScoreListsScoreTypeArray)
}

func (u *ProjectScoreListParamsScoreTypeUnion) asAny() any {
	if !param.IsOmitted(u.OfProjectScoreTypeSingle) {
		return &u.OfProjectScoreTypeSingle
	} else if !param.IsOmitted(u.OfProjectScoreListsScoreTypeArray) {
		return &u.OfProjectScoreListsScoreTypeArray
	}
	return nil
}

type ProjectScoreReplaceParams struct {
	// Name of the project score
	Name string `json:"name,required"`
	// Unique identifier for the project that the project score belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// The type of the configured score
	//
	// Any of "slider", "categorical", "weighted", "minimum", "maximum", "online",
	// "free-form".
	ScoreType shared.ProjectScoreType `json:"score_type,omitzero,required"`
	// Textual description of the project score
	Description param.Opt[string] `json:"description,omitzero"`
	// For categorical-type project scores, the list of all categories
	Categories ProjectScoreReplaceParamsCategoriesUnion `json:"categories,omitzero"`
	Config     shared.ProjectScoreConfigParam           `json:"config,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectScoreReplaceParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ProjectScoreReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectScoreReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProjectScoreReplaceParamsCategoriesUnion struct {
	OfCategorical []shared.ProjectScoreCategoryParam `json:",omitzero,inline"`
	OfMinimum     []string                           `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u ProjectScoreReplaceParamsCategoriesUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u ProjectScoreReplaceParamsCategoriesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[ProjectScoreReplaceParamsCategoriesUnion](u.OfCategorical, u.OfMinimum)
}

func (u *ProjectScoreReplaceParamsCategoriesUnion) asAny() any {
	if !param.IsOmitted(u.OfCategorical) {
		return &u.OfCategorical
	} else if !param.IsOmitted(u.OfMinimum) {
		return &u.OfMinimum
	}
	return nil
}
