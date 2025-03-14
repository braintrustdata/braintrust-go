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
	"github.com/braintrustdata/braintrust-go/internal/param"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/braintrustdata/braintrust-go/packages/pagination"
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
func NewProjectScoreService(opts ...option.RequestOption) (r *ProjectScoreService) {
	r = &ProjectScoreService{}
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
	Name param.Field[string] `json:"name,required"`
	// Unique identifier for the project that the project score belongs under
	ProjectID param.Field[string] `json:"project_id,required" format:"uuid"`
	// The type of the configured score
	ScoreType param.Field[shared.ProjectScoreType] `json:"score_type,required"`
	// For categorical-type project scores, the list of all categories
	Categories param.Field[ProjectScoreNewParamsCategoriesUnion] `json:"categories"`
	Config     param.Field[shared.ProjectScoreConfigParam]       `json:"config"`
	// Textual description of the project score
	Description param.Field[string] `json:"description"`
}

func (r ProjectScoreNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For categorical-type project scores, the list of all categories
//
// Satisfied by [ProjectScoreNewParamsCategoriesCategorical],
// [ProjectScoreNewParamsCategoriesMinimum].
type ProjectScoreNewParamsCategoriesUnion interface {
	implementsProjectScoreNewParamsCategoriesUnion()
}

type ProjectScoreNewParamsCategoriesCategorical []shared.ProjectScoreCategoryParam

func (r ProjectScoreNewParamsCategoriesCategorical) implementsProjectScoreNewParamsCategoriesUnion() {
}

type ProjectScoreNewParamsCategoriesMinimum []string

func (r ProjectScoreNewParamsCategoriesMinimum) implementsProjectScoreNewParamsCategoriesUnion() {}

type ProjectScoreUpdateParams struct {
	// For categorical-type project scores, the list of all categories
	Categories param.Field[ProjectScoreUpdateParamsCategoriesUnion] `json:"categories"`
	Config     param.Field[shared.ProjectScoreConfigParam]          `json:"config"`
	// Textual description of the project score
	Description param.Field[string] `json:"description"`
	// Name of the project score
	Name param.Field[string] `json:"name"`
	// The type of the configured score
	ScoreType param.Field[shared.ProjectScoreType] `json:"score_type"`
}

func (r ProjectScoreUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For categorical-type project scores, the list of all categories
//
// Satisfied by [ProjectScoreUpdateParamsCategoriesCategorical],
// [ProjectScoreUpdateParamsCategoriesMinimum].
type ProjectScoreUpdateParamsCategoriesUnion interface {
	implementsProjectScoreUpdateParamsCategoriesUnion()
}

type ProjectScoreUpdateParamsCategoriesCategorical []shared.ProjectScoreCategoryParam

func (r ProjectScoreUpdateParamsCategoriesCategorical) implementsProjectScoreUpdateParamsCategoriesUnion() {
}

type ProjectScoreUpdateParamsCategoriesMinimum []string

func (r ProjectScoreUpdateParamsCategoriesMinimum) implementsProjectScoreUpdateParamsCategoriesUnion() {
}

type ProjectScoreListParams struct {
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[ProjectScoreListParamsIDsUnion] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[int64] `query:"limit"`
	// Filter search results to within a particular organization
	OrgName param.Field[string] `query:"org_name"`
	// Project id
	ProjectID param.Field[string] `query:"project_id" format:"uuid"`
	// Name of the project to search for
	ProjectName param.Field[string] `query:"project_name"`
	// Name of the project_score to search for
	ProjectScoreName param.Field[string] `query:"project_score_name"`
	// The type of the configured score
	ScoreType param.Field[ProjectScoreListParamsScoreTypeUnion] `query:"score_type"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [ProjectScoreListParams]'s query parameters as `url.Values`.
func (r ProjectScoreListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [ProjectScoreListParamsIDsArray].
type ProjectScoreListParamsIDsUnion interface {
	ImplementsProjectScoreListParamsIDsUnion()
}

type ProjectScoreListParamsIDsArray []string

func (r ProjectScoreListParamsIDsArray) ImplementsProjectScoreListParamsIDsUnion() {}

// The type of the configured score
//
// Satisfied by [shared.ProjectScoreType], [ProjectScoreListParamsScoreTypeArray].
type ProjectScoreListParamsScoreTypeUnion interface {
	ImplementsProjectScoreListParamsScoreTypeUnion()
}

type ProjectScoreListParamsScoreTypeArray []shared.ProjectScoreType

func (r ProjectScoreListParamsScoreTypeArray) ImplementsProjectScoreListParamsScoreTypeUnion() {}

type ProjectScoreReplaceParams struct {
	// Name of the project score
	Name param.Field[string] `json:"name,required"`
	// Unique identifier for the project that the project score belongs under
	ProjectID param.Field[string] `json:"project_id,required" format:"uuid"`
	// The type of the configured score
	ScoreType param.Field[shared.ProjectScoreType] `json:"score_type,required"`
	// For categorical-type project scores, the list of all categories
	Categories param.Field[ProjectScoreReplaceParamsCategoriesUnion] `json:"categories"`
	Config     param.Field[shared.ProjectScoreConfigParam]           `json:"config"`
	// Textual description of the project score
	Description param.Field[string] `json:"description"`
}

func (r ProjectScoreReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For categorical-type project scores, the list of all categories
//
// Satisfied by [ProjectScoreReplaceParamsCategoriesCategorical],
// [ProjectScoreReplaceParamsCategoriesMinimum].
type ProjectScoreReplaceParamsCategoriesUnion interface {
	implementsProjectScoreReplaceParamsCategoriesUnion()
}

type ProjectScoreReplaceParamsCategoriesCategorical []shared.ProjectScoreCategoryParam

func (r ProjectScoreReplaceParamsCategoriesCategorical) implementsProjectScoreReplaceParamsCategoriesUnion() {
}

type ProjectScoreReplaceParamsCategoriesMinimum []string

func (r ProjectScoreReplaceParamsCategoriesMinimum) implementsProjectScoreReplaceParamsCategoriesUnion() {
}
