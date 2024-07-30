// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/internal/apiquery"
	"github.com/braintrustdata/braintrust-go/internal/pagination"
	"github.com/braintrustdata/braintrust-go/internal/param"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/tidwall/gjson"
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
func (r *ProjectScoreService) New(ctx context.Context, body ProjectScoreNewParams, opts ...option.RequestOption) (res *ProjectScore, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/project_score"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a project_score object by its id
func (r *ProjectScoreService) Get(ctx context.Context, projectScoreID string, opts ...option.RequestOption) (res *ProjectScore, err error) {
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
func (r *ProjectScoreService) Update(ctx context.Context, projectScoreID string, body ProjectScoreUpdateParams, opts ...option.RequestOption) (res *ProjectScore, err error) {
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
func (r *ProjectScoreService) List(ctx context.Context, query ProjectScoreListParams, opts ...option.RequestOption) (res *pagination.ListObjects[ProjectScore], err error) {
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
func (r *ProjectScoreService) ListAutoPaging(ctx context.Context, query ProjectScoreListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[ProjectScore] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a project_score object by its id
func (r *ProjectScoreService) Delete(ctx context.Context, projectScoreID string, opts ...option.RequestOption) (res *ProjectScore, err error) {
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
func (r *ProjectScoreService) Replace(ctx context.Context, body ProjectScoreReplaceParams, opts ...option.RequestOption) (res *ProjectScore, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/project_score"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// A project score is a user-configured score, which can be manually-labeled
// through the UI
type ProjectScore struct {
	// Unique identifier for the project score
	ID string `json:"id,required" format:"uuid"`
	// Name of the project score
	Name string `json:"name,required"`
	// Unique identifier for the project that the project score belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// The type of the configured score
	ScoreType ProjectScoreScoreType `json:"score_type,required,nullable"`
	UserID    string                `json:"user_id,required" format:"uuid"`
	// For categorical-type project scores, the list of all categories
	Categories ProjectScoreCategoriesUnion `json:"categories"`
	Config     ProjectScoreConfig          `json:"config,nullable"`
	// Date of project score creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Textual description of the project score
	Description string `json:"description,nullable"`
	// An optional LexoRank-based string that sets the sort position for the score in
	// the UI
	Position string           `json:"position,nullable"`
	JSON     projectScoreJSON `json:"-"`
}

// projectScoreJSON contains the JSON metadata for the struct [ProjectScore]
type projectScoreJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	ProjectID   apijson.Field
	ScoreType   apijson.Field
	UserID      apijson.Field
	Categories  apijson.Field
	Config      apijson.Field
	Created     apijson.Field
	Description apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectScore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectScoreJSON) RawJSON() string {
	return r.raw
}

// The type of the configured score
type ProjectScoreScoreType string

const (
	ProjectScoreScoreTypeSlider      ProjectScoreScoreType = "slider"
	ProjectScoreScoreTypeCategorical ProjectScoreScoreType = "categorical"
	ProjectScoreScoreTypeWeighted    ProjectScoreScoreType = "weighted"
	ProjectScoreScoreTypeMinimum     ProjectScoreScoreType = "minimum"
)

func (r ProjectScoreScoreType) IsKnown() bool {
	switch r {
	case ProjectScoreScoreTypeSlider, ProjectScoreScoreTypeCategorical, ProjectScoreScoreTypeWeighted, ProjectScoreScoreTypeMinimum:
		return true
	}
	return false
}

// For categorical-type project scores, the list of all categories
//
// Union satisfied by [ProjectScoreCategoriesArray], [ProjectScoreCategoriesArray]
// or [ProjectScoreCategoriesNullableVariant].
type ProjectScoreCategoriesUnion interface {
	implementsProjectScoreCategoriesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectScoreCategoriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectScoreCategoriesArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectScoreCategoriesArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectScoreCategoriesNullableVariant{}),
		},
	)
}

type ProjectScoreCategoriesArray []ProjectScoreCategoriesArrayItem

func (r ProjectScoreCategoriesArray) implementsProjectScoreCategoriesUnion() {}

// For categorical-type project scores, defines a single category
type ProjectScoreCategoriesArrayItem struct {
	// Name of the category
	Name string `json:"name,required"`
	// Numerical value of the category. Must be between 0 and 1, inclusive
	Value float64                             `json:"value,required"`
	JSON  projectScoreCategoriesArrayItemJSON `json:"-"`
}

// projectScoreCategoriesArrayItemJSON contains the JSON metadata for the struct
// [ProjectScoreCategoriesArrayItem]
type projectScoreCategoriesArrayItemJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectScoreCategoriesArrayItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectScoreCategoriesArrayItemJSON) RawJSON() string {
	return r.raw
}

type ProjectScoreCategoriesNullableVariant struct {
	JSON projectScoreCategoriesNullableVariantJSON `json:"-"`
}

// projectScoreCategoriesNullableVariantJSON contains the JSON metadata for the
// struct [ProjectScoreCategoriesNullableVariant]
type projectScoreCategoriesNullableVariantJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectScoreCategoriesNullableVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectScoreCategoriesNullableVariantJSON) RawJSON() string {
	return r.raw
}

func (r ProjectScoreCategoriesNullableVariant) implementsProjectScoreCategoriesUnion() {}

type ProjectScoreConfig struct {
	Destination ProjectScoreConfigDestination `json:"destination,nullable"`
	MultiSelect bool                          `json:"multi_select,nullable"`
	JSON        projectScoreConfigJSON        `json:"-"`
}

// projectScoreConfigJSON contains the JSON metadata for the struct
// [ProjectScoreConfig]
type projectScoreConfigJSON struct {
	Destination apijson.Field
	MultiSelect apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectScoreConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectScoreConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectScoreConfigDestination string

const (
	ProjectScoreConfigDestinationExpected ProjectScoreConfigDestination = "expected"
)

func (r ProjectScoreConfigDestination) IsKnown() bool {
	switch r {
	case ProjectScoreConfigDestinationExpected:
		return true
	}
	return false
}

type ProjectScoreNewParams struct {
	// Name of the project score
	Name param.Field[string] `json:"name,required"`
	// Unique identifier for the project that the project score belongs under
	ProjectID param.Field[string] `json:"project_id,required" format:"uuid"`
	// The type of the configured score
	ScoreType param.Field[ProjectScoreNewParamsScoreType] `json:"score_type,required"`
	// For categorical-type project scores, the list of all categories
	Categories param.Field[ProjectScoreNewParamsCategoriesUnion] `json:"categories"`
	// Textual description of the project score
	Description param.Field[string] `json:"description"`
}

func (r ProjectScoreNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the configured score
type ProjectScoreNewParamsScoreType string

const (
	ProjectScoreNewParamsScoreTypeSlider      ProjectScoreNewParamsScoreType = "slider"
	ProjectScoreNewParamsScoreTypeCategorical ProjectScoreNewParamsScoreType = "categorical"
	ProjectScoreNewParamsScoreTypeWeighted    ProjectScoreNewParamsScoreType = "weighted"
	ProjectScoreNewParamsScoreTypeMinimum     ProjectScoreNewParamsScoreType = "minimum"
)

func (r ProjectScoreNewParamsScoreType) IsKnown() bool {
	switch r {
	case ProjectScoreNewParamsScoreTypeSlider, ProjectScoreNewParamsScoreTypeCategorical, ProjectScoreNewParamsScoreTypeWeighted, ProjectScoreNewParamsScoreTypeMinimum:
		return true
	}
	return false
}

// For categorical-type project scores, the list of all categories
//
// Satisfied by [ProjectScoreNewParamsCategoriesArray],
// [ProjectScoreNewParamsCategoriesArray],
// [ProjectScoreNewParamsCategoriesNullableVariant].
type ProjectScoreNewParamsCategoriesUnion interface {
	implementsProjectScoreNewParamsCategoriesUnion()
}

type ProjectScoreNewParamsCategoriesArray []ProjectScoreNewParamsCategoriesArray

func (r ProjectScoreNewParamsCategoriesArray) implementsProjectScoreNewParamsCategoriesUnion() {}

type ProjectScoreNewParamsCategoriesNullableVariant struct {
}

func (r ProjectScoreNewParamsCategoriesNullableVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectScoreNewParamsCategoriesNullableVariant) implementsProjectScoreNewParamsCategoriesUnion() {
}

type ProjectScoreUpdateParams struct {
	// For categorical-type project scores, the list of all categories
	Categories param.Field[ProjectScoreUpdateParamsCategoriesUnion] `json:"categories"`
	// Textual description of the project score
	Description param.Field[string] `json:"description"`
	// Name of the project score
	Name param.Field[string] `json:"name"`
	// The type of the configured score
	ScoreType param.Field[ProjectScoreUpdateParamsScoreType] `json:"score_type"`
}

func (r ProjectScoreUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For categorical-type project scores, the list of all categories
//
// Satisfied by [ProjectScoreUpdateParamsCategoriesArray],
// [ProjectScoreUpdateParamsCategoriesArray],
// [ProjectScoreUpdateParamsCategoriesNullableVariant].
type ProjectScoreUpdateParamsCategoriesUnion interface {
	implementsProjectScoreUpdateParamsCategoriesUnion()
}

type ProjectScoreUpdateParamsCategoriesArray []ProjectScoreUpdateParamsCategoriesArray

func (r ProjectScoreUpdateParamsCategoriesArray) implementsProjectScoreUpdateParamsCategoriesUnion() {
}

type ProjectScoreUpdateParamsCategoriesNullableVariant struct {
}

func (r ProjectScoreUpdateParamsCategoriesNullableVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectScoreUpdateParamsCategoriesNullableVariant) implementsProjectScoreUpdateParamsCategoriesUnion() {
}

// The type of the configured score
type ProjectScoreUpdateParamsScoreType string

const (
	ProjectScoreUpdateParamsScoreTypeSlider      ProjectScoreUpdateParamsScoreType = "slider"
	ProjectScoreUpdateParamsScoreTypeCategorical ProjectScoreUpdateParamsScoreType = "categorical"
	ProjectScoreUpdateParamsScoreTypeWeighted    ProjectScoreUpdateParamsScoreType = "weighted"
	ProjectScoreUpdateParamsScoreTypeMinimum     ProjectScoreUpdateParamsScoreType = "minimum"
)

func (r ProjectScoreUpdateParamsScoreType) IsKnown() bool {
	switch r {
	case ProjectScoreUpdateParamsScoreTypeSlider, ProjectScoreUpdateParamsScoreTypeCategorical, ProjectScoreUpdateParamsScoreTypeWeighted, ProjectScoreUpdateParamsScoreTypeMinimum:
		return true
	}
	return false
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
	// Name of the project to search for
	ProjectName param.Field[string] `query:"project_name"`
	// Name of the project_score to search for
	ProjectScoreName param.Field[string] `query:"project_score_name"`
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

type ProjectScoreReplaceParams struct {
	// Name of the project score
	Name param.Field[string] `json:"name,required"`
	// Unique identifier for the project that the project score belongs under
	ProjectID param.Field[string] `json:"project_id,required" format:"uuid"`
	// The type of the configured score
	ScoreType param.Field[ProjectScoreReplaceParamsScoreType] `json:"score_type,required"`
	// For categorical-type project scores, the list of all categories
	Categories param.Field[ProjectScoreReplaceParamsCategoriesUnion] `json:"categories"`
	// Textual description of the project score
	Description param.Field[string] `json:"description"`
}

func (r ProjectScoreReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the configured score
type ProjectScoreReplaceParamsScoreType string

const (
	ProjectScoreReplaceParamsScoreTypeSlider      ProjectScoreReplaceParamsScoreType = "slider"
	ProjectScoreReplaceParamsScoreTypeCategorical ProjectScoreReplaceParamsScoreType = "categorical"
	ProjectScoreReplaceParamsScoreTypeWeighted    ProjectScoreReplaceParamsScoreType = "weighted"
	ProjectScoreReplaceParamsScoreTypeMinimum     ProjectScoreReplaceParamsScoreType = "minimum"
)

func (r ProjectScoreReplaceParamsScoreType) IsKnown() bool {
	switch r {
	case ProjectScoreReplaceParamsScoreTypeSlider, ProjectScoreReplaceParamsScoreTypeCategorical, ProjectScoreReplaceParamsScoreTypeWeighted, ProjectScoreReplaceParamsScoreTypeMinimum:
		return true
	}
	return false
}

// For categorical-type project scores, the list of all categories
//
// Satisfied by [ProjectScoreReplaceParamsCategoriesArray],
// [ProjectScoreReplaceParamsCategoriesArray],
// [ProjectScoreReplaceParamsCategoriesNullableVariant].
type ProjectScoreReplaceParamsCategoriesUnion interface {
	implementsProjectScoreReplaceParamsCategoriesUnion()
}

type ProjectScoreReplaceParamsCategoriesArray []ProjectScoreReplaceParamsCategoriesArray

func (r ProjectScoreReplaceParamsCategoriesArray) implementsProjectScoreReplaceParamsCategoriesUnion() {
}

type ProjectScoreReplaceParamsCategoriesNullableVariant struct {
}

func (r ProjectScoreReplaceParamsCategoriesNullableVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectScoreReplaceParamsCategoriesNullableVariant) implementsProjectScoreReplaceParamsCategoriesUnion() {
}
