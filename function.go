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

// FunctionService contains methods and other services that help with interacting
// with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionService] method instead.
type FunctionService struct {
	Options []option.RequestOption
}

// NewFunctionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFunctionService(opts ...option.RequestOption) (r *FunctionService) {
	r = &FunctionService{}
	r.Options = opts
	return
}

// Create a new function. If there is an existing function in the project with the
// same slug as the one specified in the request, will return the existing function
// unmodified
func (r *FunctionService) New(ctx context.Context, body FunctionNewParams, opts ...option.RequestOption) (res *shared.Function, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/function"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a function object by its id
func (r *FunctionService) Get(ctx context.Context, functionID string, opts ...option.RequestOption) (res *shared.Function, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required function_id parameter")
		return
	}
	path := fmt.Sprintf("v1/function/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update a function object. Specify the fields to update in the payload.
// Any object-type fields will be deep-merged with existing content. Currently we
// do not support removing fields or setting them to null.
func (r *FunctionService) Update(ctx context.Context, functionID string, body FunctionUpdateParams, opts ...option.RequestOption) (res *shared.Function, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required function_id parameter")
		return
	}
	path := fmt.Sprintf("v1/function/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all functions. The functions are sorted by creation date, with the most
// recently-created functions coming first
func (r *FunctionService) List(ctx context.Context, query FunctionListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.Function], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/function"
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

// List out all functions. The functions are sorted by creation date, with the most
// recently-created functions coming first
func (r *FunctionService) ListAutoPaging(ctx context.Context, query FunctionListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.Function] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a function object by its id
func (r *FunctionService) Delete(ctx context.Context, functionID string, opts ...option.RequestOption) (res *shared.Function, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required function_id parameter")
		return
	}
	path := fmt.Sprintf("v1/function/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create or replace function. If there is an existing function in the project with
// the same slug as the one specified in the request, will replace the existing
// function with the provided fields
func (r *FunctionService) Replace(ctx context.Context, body FunctionReplaceParams, opts ...option.RequestOption) (res *shared.Function, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/function"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type FunctionNewParams struct {
	FunctionData param.Field[FunctionNewParamsFunctionDataUnion] `json:"function_data,required"`
	// Name of the prompt
	Name param.Field[string] `json:"name,required"`
	// Unique identifier for the project that the prompt belongs under
	ProjectID param.Field[string] `json:"project_id,required" format:"uuid"`
	// Unique identifier for the prompt
	Slug param.Field[string] `json:"slug,required"`
	// Textual description of the prompt
	Description param.Field[string] `json:"description"`
	// The prompt, model, and its parameters
	PromptData param.Field[shared.PromptDataParam] `json:"prompt_data"`
	// A list of tags for the prompt
	Tags param.Field[[]string] `json:"tags"`
}

func (r FunctionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsFunctionData struct {
	Type param.Field[FunctionNewParamsFunctionDataType] `json:"type,required"`
	Data param.Field[interface{}]                       `json:"data,required"`
	Name param.Field[string]                            `json:"name"`
}

func (r FunctionNewParamsFunctionData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionData) implementsFunctionNewParamsFunctionDataUnion() {}

// Satisfied by [FunctionNewParamsFunctionDataPrompt],
// [FunctionNewParamsFunctionDataCode], [FunctionNewParamsFunctionDataGlobal],
// [FunctionNewParamsFunctionData].
type FunctionNewParamsFunctionDataUnion interface {
	implementsFunctionNewParamsFunctionDataUnion()
}

type FunctionNewParamsFunctionDataPrompt struct {
	Type param.Field[FunctionNewParamsFunctionDataPromptType] `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataPrompt) implementsFunctionNewParamsFunctionDataUnion() {}

type FunctionNewParamsFunctionDataPromptType string

const (
	FunctionNewParamsFunctionDataPromptTypePrompt FunctionNewParamsFunctionDataPromptType = "prompt"
)

func (r FunctionNewParamsFunctionDataPromptType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataPromptTypePrompt:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataCode struct {
	Data param.Field[FunctionNewParamsFunctionDataCodeData] `json:"data,required"`
	Type param.Field[FunctionNewParamsFunctionDataCodeType] `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataCode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataCode) implementsFunctionNewParamsFunctionDataUnion() {}

type FunctionNewParamsFunctionDataCodeData struct {
	BundleID       param.Field[string]                                              `json:"bundle_id,required"`
	Location       param.Field[FunctionNewParamsFunctionDataCodeDataLocation]       `json:"location,required"`
	RuntimeContext param.Field[FunctionNewParamsFunctionDataCodeDataRuntimeContext] `json:"runtime_context,required"`
}

func (r FunctionNewParamsFunctionDataCodeData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsFunctionDataCodeDataLocation struct {
	EvalName param.Field[string]                                                     `json:"eval_name,required"`
	Position param.Field[FunctionNewParamsFunctionDataCodeDataLocationPositionUnion] `json:"position,required"`
	Type     param.Field[FunctionNewParamsFunctionDataCodeDataLocationType]          `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataCodeDataLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [FunctionNewParamsFunctionDataCodeDataLocationPositionTask],
// [FunctionNewParamsFunctionDataCodeDataLocationPositionScore].
type FunctionNewParamsFunctionDataCodeDataLocationPositionUnion interface {
	implementsFunctionNewParamsFunctionDataCodeDataLocationPositionUnion()
}

type FunctionNewParamsFunctionDataCodeDataLocationPositionTask string

const (
	FunctionNewParamsFunctionDataCodeDataLocationPositionTaskTask FunctionNewParamsFunctionDataCodeDataLocationPositionTask = "task"
)

func (r FunctionNewParamsFunctionDataCodeDataLocationPositionTask) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataLocationPositionTaskTask:
		return true
	}
	return false
}

func (r FunctionNewParamsFunctionDataCodeDataLocationPositionTask) implementsFunctionNewParamsFunctionDataCodeDataLocationPositionUnion() {
}

type FunctionNewParamsFunctionDataCodeDataLocationPositionScore struct {
	Score param.Field[float64] `json:"score,required"`
}

func (r FunctionNewParamsFunctionDataCodeDataLocationPositionScore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataCodeDataLocationPositionScore) implementsFunctionNewParamsFunctionDataCodeDataLocationPositionUnion() {
}

type FunctionNewParamsFunctionDataCodeDataLocationType string

const (
	FunctionNewParamsFunctionDataCodeDataLocationTypeExperiment FunctionNewParamsFunctionDataCodeDataLocationType = "experiment"
)

func (r FunctionNewParamsFunctionDataCodeDataLocationType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataLocationTypeExperiment:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataCodeDataRuntimeContext struct {
	Runtime param.Field[FunctionNewParamsFunctionDataCodeDataRuntimeContextRuntime] `json:"runtime,required"`
	Version param.Field[string]                                                     `json:"version,required"`
}

func (r FunctionNewParamsFunctionDataCodeDataRuntimeContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsFunctionDataCodeDataRuntimeContextRuntime string

const (
	FunctionNewParamsFunctionDataCodeDataRuntimeContextRuntimeNode FunctionNewParamsFunctionDataCodeDataRuntimeContextRuntime = "node"
)

func (r FunctionNewParamsFunctionDataCodeDataRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataRuntimeContextRuntimeNode:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataCodeType string

const (
	FunctionNewParamsFunctionDataCodeTypeCode FunctionNewParamsFunctionDataCodeType = "code"
)

func (r FunctionNewParamsFunctionDataCodeType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeTypeCode:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataGlobal struct {
	Name param.Field[string]                                  `json:"name,required"`
	Type param.Field[FunctionNewParamsFunctionDataGlobalType] `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataGlobal) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataGlobal) implementsFunctionNewParamsFunctionDataUnion() {}

type FunctionNewParamsFunctionDataGlobalType string

const (
	FunctionNewParamsFunctionDataGlobalTypeGlobal FunctionNewParamsFunctionDataGlobalType = "global"
)

func (r FunctionNewParamsFunctionDataGlobalType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataGlobalTypeGlobal:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataType string

const (
	FunctionNewParamsFunctionDataTypePrompt FunctionNewParamsFunctionDataType = "prompt"
	FunctionNewParamsFunctionDataTypeCode   FunctionNewParamsFunctionDataType = "code"
	FunctionNewParamsFunctionDataTypeGlobal FunctionNewParamsFunctionDataType = "global"
)

func (r FunctionNewParamsFunctionDataType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataTypePrompt, FunctionNewParamsFunctionDataTypeCode, FunctionNewParamsFunctionDataTypeGlobal:
		return true
	}
	return false
}

type FunctionUpdateParams struct {
	// Textual description of the prompt
	Description  param.Field[string]                                `json:"description"`
	FunctionData param.Field[FunctionUpdateParamsFunctionDataUnion] `json:"function_data"`
	// Name of the prompt
	Name param.Field[string] `json:"name"`
	// The prompt, model, and its parameters
	PromptData param.Field[shared.PromptDataParam] `json:"prompt_data"`
	// A list of tags for the prompt
	Tags param.Field[[]string] `json:"tags"`
}

func (r FunctionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsFunctionData struct {
	Type param.Field[FunctionUpdateParamsFunctionDataType] `json:"type"`
	Data param.Field[interface{}]                          `json:"data,required"`
	Name param.Field[string]                               `json:"name"`
}

func (r FunctionUpdateParamsFunctionData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionData) implementsFunctionUpdateParamsFunctionDataUnion() {}

// Satisfied by [FunctionUpdateParamsFunctionDataPrompt],
// [FunctionUpdateParamsFunctionDataCode],
// [FunctionUpdateParamsFunctionDataGlobal],
// [FunctionUpdateParamsFunctionDataNullableVariant],
// [FunctionUpdateParamsFunctionData].
type FunctionUpdateParamsFunctionDataUnion interface {
	implementsFunctionUpdateParamsFunctionDataUnion()
}

type FunctionUpdateParamsFunctionDataPrompt struct {
	Type param.Field[FunctionUpdateParamsFunctionDataPromptType] `json:"type,required"`
}

func (r FunctionUpdateParamsFunctionDataPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataPrompt) implementsFunctionUpdateParamsFunctionDataUnion() {}

type FunctionUpdateParamsFunctionDataPromptType string

const (
	FunctionUpdateParamsFunctionDataPromptTypePrompt FunctionUpdateParamsFunctionDataPromptType = "prompt"
)

func (r FunctionUpdateParamsFunctionDataPromptType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataPromptTypePrompt:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataCode struct {
	Data param.Field[FunctionUpdateParamsFunctionDataCodeData] `json:"data,required"`
	Type param.Field[FunctionUpdateParamsFunctionDataCodeType] `json:"type,required"`
}

func (r FunctionUpdateParamsFunctionDataCode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataCode) implementsFunctionUpdateParamsFunctionDataUnion() {}

type FunctionUpdateParamsFunctionDataCodeData struct {
	BundleID       param.Field[string]                                                 `json:"bundle_id,required"`
	Location       param.Field[FunctionUpdateParamsFunctionDataCodeDataLocation]       `json:"location,required"`
	RuntimeContext param.Field[FunctionUpdateParamsFunctionDataCodeDataRuntimeContext] `json:"runtime_context,required"`
}

func (r FunctionUpdateParamsFunctionDataCodeData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsFunctionDataCodeDataLocation struct {
	EvalName param.Field[string]                                                        `json:"eval_name,required"`
	Position param.Field[FunctionUpdateParamsFunctionDataCodeDataLocationPositionUnion] `json:"position,required"`
	Type     param.Field[FunctionUpdateParamsFunctionDataCodeDataLocationType]          `json:"type,required"`
}

func (r FunctionUpdateParamsFunctionDataCodeDataLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [FunctionUpdateParamsFunctionDataCodeDataLocationPositionTask],
// [FunctionUpdateParamsFunctionDataCodeDataLocationPositionScore].
type FunctionUpdateParamsFunctionDataCodeDataLocationPositionUnion interface {
	implementsFunctionUpdateParamsFunctionDataCodeDataLocationPositionUnion()
}

type FunctionUpdateParamsFunctionDataCodeDataLocationPositionTask string

const (
	FunctionUpdateParamsFunctionDataCodeDataLocationPositionTaskTask FunctionUpdateParamsFunctionDataCodeDataLocationPositionTask = "task"
)

func (r FunctionUpdateParamsFunctionDataCodeDataLocationPositionTask) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataCodeDataLocationPositionTaskTask:
		return true
	}
	return false
}

func (r FunctionUpdateParamsFunctionDataCodeDataLocationPositionTask) implementsFunctionUpdateParamsFunctionDataCodeDataLocationPositionUnion() {
}

type FunctionUpdateParamsFunctionDataCodeDataLocationPositionScore struct {
	Score param.Field[float64] `json:"score,required"`
}

func (r FunctionUpdateParamsFunctionDataCodeDataLocationPositionScore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataCodeDataLocationPositionScore) implementsFunctionUpdateParamsFunctionDataCodeDataLocationPositionUnion() {
}

type FunctionUpdateParamsFunctionDataCodeDataLocationType string

const (
	FunctionUpdateParamsFunctionDataCodeDataLocationTypeExperiment FunctionUpdateParamsFunctionDataCodeDataLocationType = "experiment"
)

func (r FunctionUpdateParamsFunctionDataCodeDataLocationType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataCodeDataLocationTypeExperiment:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataCodeDataRuntimeContext struct {
	Runtime param.Field[FunctionUpdateParamsFunctionDataCodeDataRuntimeContextRuntime] `json:"runtime,required"`
	Version param.Field[string]                                                        `json:"version,required"`
}

func (r FunctionUpdateParamsFunctionDataCodeDataRuntimeContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsFunctionDataCodeDataRuntimeContextRuntime string

const (
	FunctionUpdateParamsFunctionDataCodeDataRuntimeContextRuntimeNode FunctionUpdateParamsFunctionDataCodeDataRuntimeContextRuntime = "node"
)

func (r FunctionUpdateParamsFunctionDataCodeDataRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataCodeDataRuntimeContextRuntimeNode:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataCodeType string

const (
	FunctionUpdateParamsFunctionDataCodeTypeCode FunctionUpdateParamsFunctionDataCodeType = "code"
)

func (r FunctionUpdateParamsFunctionDataCodeType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataCodeTypeCode:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataGlobal struct {
	Name param.Field[string]                                     `json:"name,required"`
	Type param.Field[FunctionUpdateParamsFunctionDataGlobalType] `json:"type,required"`
}

func (r FunctionUpdateParamsFunctionDataGlobal) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataGlobal) implementsFunctionUpdateParamsFunctionDataUnion() {}

type FunctionUpdateParamsFunctionDataGlobalType string

const (
	FunctionUpdateParamsFunctionDataGlobalTypeGlobal FunctionUpdateParamsFunctionDataGlobalType = "global"
)

func (r FunctionUpdateParamsFunctionDataGlobalType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataGlobalTypeGlobal:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataNullableVariant struct {
}

func (r FunctionUpdateParamsFunctionDataNullableVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataNullableVariant) implementsFunctionUpdateParamsFunctionDataUnion() {
}

type FunctionUpdateParamsFunctionDataType string

const (
	FunctionUpdateParamsFunctionDataTypePrompt FunctionUpdateParamsFunctionDataType = "prompt"
	FunctionUpdateParamsFunctionDataTypeCode   FunctionUpdateParamsFunctionDataType = "code"
	FunctionUpdateParamsFunctionDataTypeGlobal FunctionUpdateParamsFunctionDataType = "global"
)

func (r FunctionUpdateParamsFunctionDataType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataTypePrompt, FunctionUpdateParamsFunctionDataTypeCode, FunctionUpdateParamsFunctionDataTypeGlobal:
		return true
	}
	return false
}

type FunctionListParams struct {
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Name of the function to search for
	FunctionName param.Field[string] `query:"function_name"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[FunctionListParamsIDsUnion] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[int64] `query:"limit"`
	// Filter search results to within a particular organization
	OrgName param.Field[string] `query:"org_name"`
	// Project id
	ProjectID param.Field[string] `query:"project_id" format:"uuid"`
	// Name of the project to search for
	ProjectName param.Field[string] `query:"project_name"`
	// Retrieve prompt with a specific slug
	Slug param.Field[string] `query:"slug"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
	// Retrieve prompt at a specific version.
	//
	// The version id can either be a transaction id (e.g. '1000192656880881099') or a
	// version identifier (e.g. '81cd05ee665fdfb3').
	Version param.Field[string] `query:"version"`
}

// URLQuery serializes [FunctionListParams]'s query parameters as `url.Values`.
func (r FunctionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [FunctionListParamsIDsArray].
type FunctionListParamsIDsUnion interface {
	ImplementsFunctionListParamsIDsUnion()
}

type FunctionListParamsIDsArray []string

func (r FunctionListParamsIDsArray) ImplementsFunctionListParamsIDsUnion() {}

type FunctionReplaceParams struct {
	FunctionData param.Field[FunctionReplaceParamsFunctionDataUnion] `json:"function_data,required"`
	// Name of the prompt
	Name param.Field[string] `json:"name,required"`
	// Unique identifier for the project that the prompt belongs under
	ProjectID param.Field[string] `json:"project_id,required" format:"uuid"`
	// Unique identifier for the prompt
	Slug param.Field[string] `json:"slug,required"`
	// Textual description of the prompt
	Description param.Field[string] `json:"description"`
	// The prompt, model, and its parameters
	PromptData param.Field[shared.PromptDataParam] `json:"prompt_data"`
	// A list of tags for the prompt
	Tags param.Field[[]string] `json:"tags"`
}

func (r FunctionReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsFunctionData struct {
	Type param.Field[FunctionReplaceParamsFunctionDataType] `json:"type,required"`
	Data param.Field[interface{}]                           `json:"data,required"`
	Name param.Field[string]                                `json:"name"`
}

func (r FunctionReplaceParamsFunctionData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionData) implementsFunctionReplaceParamsFunctionDataUnion() {}

// Satisfied by [FunctionReplaceParamsFunctionDataPrompt],
// [FunctionReplaceParamsFunctionDataCode],
// [FunctionReplaceParamsFunctionDataGlobal], [FunctionReplaceParamsFunctionData].
type FunctionReplaceParamsFunctionDataUnion interface {
	implementsFunctionReplaceParamsFunctionDataUnion()
}

type FunctionReplaceParamsFunctionDataPrompt struct {
	Type param.Field[FunctionReplaceParamsFunctionDataPromptType] `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataPrompt) implementsFunctionReplaceParamsFunctionDataUnion() {}

type FunctionReplaceParamsFunctionDataPromptType string

const (
	FunctionReplaceParamsFunctionDataPromptTypePrompt FunctionReplaceParamsFunctionDataPromptType = "prompt"
)

func (r FunctionReplaceParamsFunctionDataPromptType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataPromptTypePrompt:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataCode struct {
	Data param.Field[FunctionReplaceParamsFunctionDataCodeData] `json:"data,required"`
	Type param.Field[FunctionReplaceParamsFunctionDataCodeType] `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataCode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataCode) implementsFunctionReplaceParamsFunctionDataUnion() {}

type FunctionReplaceParamsFunctionDataCodeData struct {
	BundleID       param.Field[string]                                                  `json:"bundle_id,required"`
	Location       param.Field[FunctionReplaceParamsFunctionDataCodeDataLocation]       `json:"location,required"`
	RuntimeContext param.Field[FunctionReplaceParamsFunctionDataCodeDataRuntimeContext] `json:"runtime_context,required"`
}

func (r FunctionReplaceParamsFunctionDataCodeData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsFunctionDataCodeDataLocation struct {
	EvalName param.Field[string]                                                         `json:"eval_name,required"`
	Position param.Field[FunctionReplaceParamsFunctionDataCodeDataLocationPositionUnion] `json:"position,required"`
	Type     param.Field[FunctionReplaceParamsFunctionDataCodeDataLocationType]          `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [FunctionReplaceParamsFunctionDataCodeDataLocationPositionTask],
// [FunctionReplaceParamsFunctionDataCodeDataLocationPositionScore].
type FunctionReplaceParamsFunctionDataCodeDataLocationPositionUnion interface {
	implementsFunctionReplaceParamsFunctionDataCodeDataLocationPositionUnion()
}

type FunctionReplaceParamsFunctionDataCodeDataLocationPositionTask string

const (
	FunctionReplaceParamsFunctionDataCodeDataLocationPositionTaskTask FunctionReplaceParamsFunctionDataCodeDataLocationPositionTask = "task"
)

func (r FunctionReplaceParamsFunctionDataCodeDataLocationPositionTask) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataLocationPositionTaskTask:
		return true
	}
	return false
}

func (r FunctionReplaceParamsFunctionDataCodeDataLocationPositionTask) implementsFunctionReplaceParamsFunctionDataCodeDataLocationPositionUnion() {
}

type FunctionReplaceParamsFunctionDataCodeDataLocationPositionScore struct {
	Score param.Field[float64] `json:"score,required"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataLocationPositionScore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataCodeDataLocationPositionScore) implementsFunctionReplaceParamsFunctionDataCodeDataLocationPositionUnion() {
}

type FunctionReplaceParamsFunctionDataCodeDataLocationType string

const (
	FunctionReplaceParamsFunctionDataCodeDataLocationTypeExperiment FunctionReplaceParamsFunctionDataCodeDataLocationType = "experiment"
)

func (r FunctionReplaceParamsFunctionDataCodeDataLocationType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataLocationTypeExperiment:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataCodeDataRuntimeContext struct {
	Runtime param.Field[FunctionReplaceParamsFunctionDataCodeDataRuntimeContextRuntime] `json:"runtime,required"`
	Version param.Field[string]                                                         `json:"version,required"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataRuntimeContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsFunctionDataCodeDataRuntimeContextRuntime string

const (
	FunctionReplaceParamsFunctionDataCodeDataRuntimeContextRuntimeNode FunctionReplaceParamsFunctionDataCodeDataRuntimeContextRuntime = "node"
)

func (r FunctionReplaceParamsFunctionDataCodeDataRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataRuntimeContextRuntimeNode:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataCodeType string

const (
	FunctionReplaceParamsFunctionDataCodeTypeCode FunctionReplaceParamsFunctionDataCodeType = "code"
)

func (r FunctionReplaceParamsFunctionDataCodeType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeTypeCode:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataGlobal struct {
	Name param.Field[string]                                      `json:"name,required"`
	Type param.Field[FunctionReplaceParamsFunctionDataGlobalType] `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataGlobal) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataGlobal) implementsFunctionReplaceParamsFunctionDataUnion() {}

type FunctionReplaceParamsFunctionDataGlobalType string

const (
	FunctionReplaceParamsFunctionDataGlobalTypeGlobal FunctionReplaceParamsFunctionDataGlobalType = "global"
)

func (r FunctionReplaceParamsFunctionDataGlobalType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataGlobalTypeGlobal:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataType string

const (
	FunctionReplaceParamsFunctionDataTypePrompt FunctionReplaceParamsFunctionDataType = "prompt"
	FunctionReplaceParamsFunctionDataTypeCode   FunctionReplaceParamsFunctionDataType = "code"
	FunctionReplaceParamsFunctionDataTypeGlobal FunctionReplaceParamsFunctionDataType = "global"
)

func (r FunctionReplaceParamsFunctionDataType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataTypePrompt, FunctionReplaceParamsFunctionDataTypeCode, FunctionReplaceParamsFunctionDataTypeGlobal:
		return true
	}
	return false
}
