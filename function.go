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
	Description  param.Field[string]                        `json:"description"`
	FunctionType param.Field[FunctionNewParamsFunctionType] `json:"function_type"`
	Origin       param.Field[FunctionNewParamsOrigin]       `json:"origin"`
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
	Data param.Field[FunctionNewParamsFunctionDataCodeDataUnion] `json:"data,required"`
	Type param.Field[FunctionNewParamsFunctionDataCodeType]      `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataCode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataCode) implementsFunctionNewParamsFunctionDataUnion() {}

type FunctionNewParamsFunctionDataCodeData struct {
	Type           param.Field[FunctionNewParamsFunctionDataCodeDataType] `json:"type,required"`
	RuntimeContext param.Field[interface{}]                               `json:"runtime_context"`
	Location       param.Field[interface{}]                               `json:"location,required"`
	BundleID       param.Field[string]                                    `json:"bundle_id"`
	// A preview of the code
	Preview param.Field[string] `json:"preview"`
	Code    param.Field[string] `json:"code"`
}

func (r FunctionNewParamsFunctionDataCodeData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataCodeData) implementsFunctionNewParamsFunctionDataCodeDataUnion() {
}

// Satisfied by [FunctionNewParamsFunctionDataCodeDataBundle],
// [FunctionNewParamsFunctionDataCodeDataInline],
// [FunctionNewParamsFunctionDataCodeData].
type FunctionNewParamsFunctionDataCodeDataUnion interface {
	implementsFunctionNewParamsFunctionDataCodeDataUnion()
}

type FunctionNewParamsFunctionDataCodeDataBundle struct {
	BundleID       param.Field[string]                                                    `json:"bundle_id,required"`
	Location       param.Field[FunctionNewParamsFunctionDataCodeDataBundleLocation]       `json:"location,required"`
	RuntimeContext param.Field[FunctionNewParamsFunctionDataCodeDataBundleRuntimeContext] `json:"runtime_context,required"`
	Type           param.Field[FunctionNewParamsFunctionDataCodeDataBundleType]           `json:"type,required"`
	// A preview of the code
	Preview param.Field[string] `json:"preview"`
}

func (r FunctionNewParamsFunctionDataCodeDataBundle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataCodeDataBundle) implementsFunctionNewParamsFunctionDataCodeDataUnion() {
}

type FunctionNewParamsFunctionDataCodeDataBundleLocation struct {
	EvalName param.Field[string]                                                           `json:"eval_name,required"`
	Position param.Field[FunctionNewParamsFunctionDataCodeDataBundleLocationPositionUnion] `json:"position,required"`
	Type     param.Field[FunctionNewParamsFunctionDataCodeDataBundleLocationType]          `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationPosition struct {
	Type  param.Field[FunctionNewParamsFunctionDataCodeDataBundleLocationPositionType] `json:"type,required"`
	Index param.Field[float64]                                                         `json:"index"`
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationPosition) implementsFunctionNewParamsFunctionDataCodeDataBundleLocationPositionUnion() {
}

// Satisfied by [FunctionNewParamsFunctionDataCodeDataBundleLocationPositionType],
// [FunctionNewParamsFunctionDataCodeDataBundleLocationPositionScorer],
// [FunctionNewParamsFunctionDataCodeDataBundleLocationPosition].
type FunctionNewParamsFunctionDataCodeDataBundleLocationPositionUnion interface {
	implementsFunctionNewParamsFunctionDataCodeDataBundleLocationPositionUnion()
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationPositionType struct {
	Type param.Field[FunctionNewParamsFunctionDataCodeDataBundleLocationPositionTypeType] `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationPositionType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationPositionType) implementsFunctionNewParamsFunctionDataCodeDataBundleLocationPositionUnion() {
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationPositionTypeType string

const (
	FunctionNewParamsFunctionDataCodeDataBundleLocationPositionTypeTypeTask FunctionNewParamsFunctionDataCodeDataBundleLocationPositionTypeType = "task"
)

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationPositionTypeType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataBundleLocationPositionTypeTypeTask:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationPositionScorer struct {
	Index param.Field[float64]                                                               `json:"index,required"`
	Type  param.Field[FunctionNewParamsFunctionDataCodeDataBundleLocationPositionScorerType] `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationPositionScorer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationPositionScorer) implementsFunctionNewParamsFunctionDataCodeDataBundleLocationPositionUnion() {
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationPositionScorerType string

const (
	FunctionNewParamsFunctionDataCodeDataBundleLocationPositionScorerTypeScorer FunctionNewParamsFunctionDataCodeDataBundleLocationPositionScorerType = "scorer"
)

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationPositionScorerType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataBundleLocationPositionScorerTypeScorer:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationType string

const (
	FunctionNewParamsFunctionDataCodeDataBundleLocationTypeExperiment FunctionNewParamsFunctionDataCodeDataBundleLocationType = "experiment"
)

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataBundleLocationTypeExperiment:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataCodeDataBundleRuntimeContext struct {
	Runtime param.Field[FunctionNewParamsFunctionDataCodeDataBundleRuntimeContextRuntime] `json:"runtime,required"`
	Version param.Field[string]                                                           `json:"version,required"`
}

func (r FunctionNewParamsFunctionDataCodeDataBundleRuntimeContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsFunctionDataCodeDataBundleRuntimeContextRuntime string

const (
	FunctionNewParamsFunctionDataCodeDataBundleRuntimeContextRuntimeNode   FunctionNewParamsFunctionDataCodeDataBundleRuntimeContextRuntime = "node"
	FunctionNewParamsFunctionDataCodeDataBundleRuntimeContextRuntimePython FunctionNewParamsFunctionDataCodeDataBundleRuntimeContextRuntime = "python"
)

func (r FunctionNewParamsFunctionDataCodeDataBundleRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataBundleRuntimeContextRuntimeNode, FunctionNewParamsFunctionDataCodeDataBundleRuntimeContextRuntimePython:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataCodeDataBundleType string

const (
	FunctionNewParamsFunctionDataCodeDataBundleTypeBundle FunctionNewParamsFunctionDataCodeDataBundleType = "bundle"
)

func (r FunctionNewParamsFunctionDataCodeDataBundleType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataBundleTypeBundle:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataCodeDataInline struct {
	Code           param.Field[string]                                                    `json:"code,required"`
	RuntimeContext param.Field[FunctionNewParamsFunctionDataCodeDataInlineRuntimeContext] `json:"runtime_context,required"`
	Type           param.Field[FunctionNewParamsFunctionDataCodeDataInlineType]           `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataCodeDataInline) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataCodeDataInline) implementsFunctionNewParamsFunctionDataCodeDataUnion() {
}

type FunctionNewParamsFunctionDataCodeDataInlineRuntimeContext struct {
	Runtime param.Field[FunctionNewParamsFunctionDataCodeDataInlineRuntimeContextRuntime] `json:"runtime,required"`
	Version param.Field[string]                                                           `json:"version,required"`
}

func (r FunctionNewParamsFunctionDataCodeDataInlineRuntimeContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsFunctionDataCodeDataInlineRuntimeContextRuntime string

const (
	FunctionNewParamsFunctionDataCodeDataInlineRuntimeContextRuntimeNode   FunctionNewParamsFunctionDataCodeDataInlineRuntimeContextRuntime = "node"
	FunctionNewParamsFunctionDataCodeDataInlineRuntimeContextRuntimePython FunctionNewParamsFunctionDataCodeDataInlineRuntimeContextRuntime = "python"
)

func (r FunctionNewParamsFunctionDataCodeDataInlineRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataInlineRuntimeContextRuntimeNode, FunctionNewParamsFunctionDataCodeDataInlineRuntimeContextRuntimePython:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataCodeDataInlineType string

const (
	FunctionNewParamsFunctionDataCodeDataInlineTypeInline FunctionNewParamsFunctionDataCodeDataInlineType = "inline"
)

func (r FunctionNewParamsFunctionDataCodeDataInlineType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataInlineTypeInline:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataCodeDataType string

const (
	FunctionNewParamsFunctionDataCodeDataTypeBundle FunctionNewParamsFunctionDataCodeDataType = "bundle"
	FunctionNewParamsFunctionDataCodeDataTypeInline FunctionNewParamsFunctionDataCodeDataType = "inline"
)

func (r FunctionNewParamsFunctionDataCodeDataType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataTypeBundle, FunctionNewParamsFunctionDataCodeDataTypeInline:
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

type FunctionNewParamsFunctionType string

const (
	FunctionNewParamsFunctionTypeTask   FunctionNewParamsFunctionType = "task"
	FunctionNewParamsFunctionTypeLlm    FunctionNewParamsFunctionType = "llm"
	FunctionNewParamsFunctionTypeScorer FunctionNewParamsFunctionType = "scorer"
)

func (r FunctionNewParamsFunctionType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionTypeTask, FunctionNewParamsFunctionTypeLlm, FunctionNewParamsFunctionTypeScorer:
		return true
	}
	return false
}

type FunctionNewParamsOrigin struct {
	// Id of the object the function is originating from
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[FunctionNewParamsOriginObjectType] `json:"object_type,required"`
	// The function exists for internal purposes and should not be displayed in the
	// list of functions.
	Internal param.Field[bool] `json:"internal"`
}

func (r FunctionNewParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The object type that the ACL applies to
type FunctionNewParamsOriginObjectType string

const (
	FunctionNewParamsOriginObjectTypeOrganization  FunctionNewParamsOriginObjectType = "organization"
	FunctionNewParamsOriginObjectTypeProject       FunctionNewParamsOriginObjectType = "project"
	FunctionNewParamsOriginObjectTypeExperiment    FunctionNewParamsOriginObjectType = "experiment"
	FunctionNewParamsOriginObjectTypeDataset       FunctionNewParamsOriginObjectType = "dataset"
	FunctionNewParamsOriginObjectTypePrompt        FunctionNewParamsOriginObjectType = "prompt"
	FunctionNewParamsOriginObjectTypePromptSession FunctionNewParamsOriginObjectType = "prompt_session"
	FunctionNewParamsOriginObjectTypeGroup         FunctionNewParamsOriginObjectType = "group"
	FunctionNewParamsOriginObjectTypeRole          FunctionNewParamsOriginObjectType = "role"
	FunctionNewParamsOriginObjectTypeOrgMember     FunctionNewParamsOriginObjectType = "org_member"
	FunctionNewParamsOriginObjectTypeProjectLog    FunctionNewParamsOriginObjectType = "project_log"
	FunctionNewParamsOriginObjectTypeOrgProject    FunctionNewParamsOriginObjectType = "org_project"
)

func (r FunctionNewParamsOriginObjectType) IsKnown() bool {
	switch r {
	case FunctionNewParamsOriginObjectTypeOrganization, FunctionNewParamsOriginObjectTypeProject, FunctionNewParamsOriginObjectTypeExperiment, FunctionNewParamsOriginObjectTypeDataset, FunctionNewParamsOriginObjectTypePrompt, FunctionNewParamsOriginObjectTypePromptSession, FunctionNewParamsOriginObjectTypeGroup, FunctionNewParamsOriginObjectTypeRole, FunctionNewParamsOriginObjectTypeOrgMember, FunctionNewParamsOriginObjectTypeProjectLog, FunctionNewParamsOriginObjectTypeOrgProject:
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
	Data param.Field[FunctionUpdateParamsFunctionDataCodeDataUnion] `json:"data,required"`
	Type param.Field[FunctionUpdateParamsFunctionDataCodeType]      `json:"type,required"`
}

func (r FunctionUpdateParamsFunctionDataCode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataCode) implementsFunctionUpdateParamsFunctionDataUnion() {}

type FunctionUpdateParamsFunctionDataCodeData struct {
	Type           param.Field[FunctionUpdateParamsFunctionDataCodeDataType] `json:"type,required"`
	RuntimeContext param.Field[interface{}]                                  `json:"runtime_context"`
	Location       param.Field[interface{}]                                  `json:"location,required"`
	BundleID       param.Field[string]                                       `json:"bundle_id"`
	// A preview of the code
	Preview param.Field[string] `json:"preview"`
	Code    param.Field[string] `json:"code"`
}

func (r FunctionUpdateParamsFunctionDataCodeData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataCodeData) implementsFunctionUpdateParamsFunctionDataCodeDataUnion() {
}

// Satisfied by [FunctionUpdateParamsFunctionDataCodeDataBundle],
// [FunctionUpdateParamsFunctionDataCodeDataInline],
// [FunctionUpdateParamsFunctionDataCodeData].
type FunctionUpdateParamsFunctionDataCodeDataUnion interface {
	implementsFunctionUpdateParamsFunctionDataCodeDataUnion()
}

type FunctionUpdateParamsFunctionDataCodeDataBundle struct {
	BundleID       param.Field[string]                                                       `json:"bundle_id,required"`
	Location       param.Field[FunctionUpdateParamsFunctionDataCodeDataBundleLocation]       `json:"location,required"`
	RuntimeContext param.Field[FunctionUpdateParamsFunctionDataCodeDataBundleRuntimeContext] `json:"runtime_context,required"`
	Type           param.Field[FunctionUpdateParamsFunctionDataCodeDataBundleType]           `json:"type,required"`
	// A preview of the code
	Preview param.Field[string] `json:"preview"`
}

func (r FunctionUpdateParamsFunctionDataCodeDataBundle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataCodeDataBundle) implementsFunctionUpdateParamsFunctionDataCodeDataUnion() {
}

type FunctionUpdateParamsFunctionDataCodeDataBundleLocation struct {
	EvalName param.Field[string]                                                              `json:"eval_name,required"`
	Position param.Field[FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionUnion] `json:"position,required"`
	Type     param.Field[FunctionUpdateParamsFunctionDataCodeDataBundleLocationType]          `json:"type,required"`
}

func (r FunctionUpdateParamsFunctionDataCodeDataBundleLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsFunctionDataCodeDataBundleLocationPosition struct {
	Type  param.Field[FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionType] `json:"type,required"`
	Index param.Field[float64]                                                            `json:"index"`
}

func (r FunctionUpdateParamsFunctionDataCodeDataBundleLocationPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataCodeDataBundleLocationPosition) implementsFunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionUnion() {
}

// Satisfied by
// [FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionType],
// [FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionScorer],
// [FunctionUpdateParamsFunctionDataCodeDataBundleLocationPosition].
type FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionUnion interface {
	implementsFunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionUnion()
}

type FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionType struct {
	Type param.Field[FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionTypeType] `json:"type,required"`
}

func (r FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionType) implementsFunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionUnion() {
}

type FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionTypeType string

const (
	FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionTypeTypeTask FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionTypeType = "task"
)

func (r FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionTypeType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionTypeTypeTask:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionScorer struct {
	Index param.Field[float64]                                                                  `json:"index,required"`
	Type  param.Field[FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionScorerType] `json:"type,required"`
}

func (r FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionScorer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionScorer) implementsFunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionUnion() {
}

type FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionScorerType string

const (
	FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionScorerTypeScorer FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionScorerType = "scorer"
)

func (r FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionScorerType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataCodeDataBundleLocationPositionScorerTypeScorer:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataCodeDataBundleLocationType string

const (
	FunctionUpdateParamsFunctionDataCodeDataBundleLocationTypeExperiment FunctionUpdateParamsFunctionDataCodeDataBundleLocationType = "experiment"
)

func (r FunctionUpdateParamsFunctionDataCodeDataBundleLocationType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataCodeDataBundleLocationTypeExperiment:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataCodeDataBundleRuntimeContext struct {
	Runtime param.Field[FunctionUpdateParamsFunctionDataCodeDataBundleRuntimeContextRuntime] `json:"runtime,required"`
	Version param.Field[string]                                                              `json:"version,required"`
}

func (r FunctionUpdateParamsFunctionDataCodeDataBundleRuntimeContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsFunctionDataCodeDataBundleRuntimeContextRuntime string

const (
	FunctionUpdateParamsFunctionDataCodeDataBundleRuntimeContextRuntimeNode   FunctionUpdateParamsFunctionDataCodeDataBundleRuntimeContextRuntime = "node"
	FunctionUpdateParamsFunctionDataCodeDataBundleRuntimeContextRuntimePython FunctionUpdateParamsFunctionDataCodeDataBundleRuntimeContextRuntime = "python"
)

func (r FunctionUpdateParamsFunctionDataCodeDataBundleRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataCodeDataBundleRuntimeContextRuntimeNode, FunctionUpdateParamsFunctionDataCodeDataBundleRuntimeContextRuntimePython:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataCodeDataBundleType string

const (
	FunctionUpdateParamsFunctionDataCodeDataBundleTypeBundle FunctionUpdateParamsFunctionDataCodeDataBundleType = "bundle"
)

func (r FunctionUpdateParamsFunctionDataCodeDataBundleType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataCodeDataBundleTypeBundle:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataCodeDataInline struct {
	Code           param.Field[string]                                                       `json:"code,required"`
	RuntimeContext param.Field[FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContext] `json:"runtime_context,required"`
	Type           param.Field[FunctionUpdateParamsFunctionDataCodeDataInlineType]           `json:"type,required"`
}

func (r FunctionUpdateParamsFunctionDataCodeDataInline) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataCodeDataInline) implementsFunctionUpdateParamsFunctionDataCodeDataUnion() {
}

type FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContext struct {
	Runtime param.Field[FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContextRuntime] `json:"runtime,required"`
	Version param.Field[string]                                                              `json:"version,required"`
}

func (r FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContextRuntime string

const (
	FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContextRuntimeNode   FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContextRuntime = "node"
	FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContextRuntimePython FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContextRuntime = "python"
)

func (r FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContextRuntimeNode, FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContextRuntimePython:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataCodeDataInlineType string

const (
	FunctionUpdateParamsFunctionDataCodeDataInlineTypeInline FunctionUpdateParamsFunctionDataCodeDataInlineType = "inline"
)

func (r FunctionUpdateParamsFunctionDataCodeDataInlineType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataCodeDataInlineTypeInline:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataCodeDataType string

const (
	FunctionUpdateParamsFunctionDataCodeDataTypeBundle FunctionUpdateParamsFunctionDataCodeDataType = "bundle"
	FunctionUpdateParamsFunctionDataCodeDataTypeInline FunctionUpdateParamsFunctionDataCodeDataType = "inline"
)

func (r FunctionUpdateParamsFunctionDataCodeDataType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataCodeDataTypeBundle, FunctionUpdateParamsFunctionDataCodeDataTypeInline:
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
	Description  param.Field[string]                            `json:"description"`
	FunctionType param.Field[FunctionReplaceParamsFunctionType] `json:"function_type"`
	Origin       param.Field[FunctionReplaceParamsOrigin]       `json:"origin"`
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
	Data param.Field[FunctionReplaceParamsFunctionDataCodeDataUnion] `json:"data,required"`
	Type param.Field[FunctionReplaceParamsFunctionDataCodeType]      `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataCode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataCode) implementsFunctionReplaceParamsFunctionDataUnion() {}

type FunctionReplaceParamsFunctionDataCodeData struct {
	Type           param.Field[FunctionReplaceParamsFunctionDataCodeDataType] `json:"type,required"`
	RuntimeContext param.Field[interface{}]                                   `json:"runtime_context"`
	Location       param.Field[interface{}]                                   `json:"location,required"`
	BundleID       param.Field[string]                                        `json:"bundle_id"`
	// A preview of the code
	Preview param.Field[string] `json:"preview"`
	Code    param.Field[string] `json:"code"`
}

func (r FunctionReplaceParamsFunctionDataCodeData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataCodeData) implementsFunctionReplaceParamsFunctionDataCodeDataUnion() {
}

// Satisfied by [FunctionReplaceParamsFunctionDataCodeDataBundle],
// [FunctionReplaceParamsFunctionDataCodeDataInline],
// [FunctionReplaceParamsFunctionDataCodeData].
type FunctionReplaceParamsFunctionDataCodeDataUnion interface {
	implementsFunctionReplaceParamsFunctionDataCodeDataUnion()
}

type FunctionReplaceParamsFunctionDataCodeDataBundle struct {
	BundleID       param.Field[string]                                                        `json:"bundle_id,required"`
	Location       param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleLocation]       `json:"location,required"`
	RuntimeContext param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleRuntimeContext] `json:"runtime_context,required"`
	Type           param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleType]           `json:"type,required"`
	// A preview of the code
	Preview param.Field[string] `json:"preview"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundle) implementsFunctionReplaceParamsFunctionDataCodeDataUnion() {
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocation struct {
	EvalName param.Field[string]                                                               `json:"eval_name,required"`
	Position param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionUnion] `json:"position,required"`
	Type     param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleLocationType]          `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationPosition struct {
	Type  param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionType] `json:"type,required"`
	Index param.Field[float64]                                                             `json:"index"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationPosition) implementsFunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionUnion() {
}

// Satisfied by
// [FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionType],
// [FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionScorer],
// [FunctionReplaceParamsFunctionDataCodeDataBundleLocationPosition].
type FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionUnion interface {
	implementsFunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionUnion()
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionType struct {
	Type param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionTypeType] `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionType) implementsFunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionUnion() {
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionTypeType string

const (
	FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionTypeTypeTask FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionTypeType = "task"
)

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionTypeType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionTypeTypeTask:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionScorer struct {
	Index param.Field[float64]                                                                   `json:"index,required"`
	Type  param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionScorerType] `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionScorer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionScorer) implementsFunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionUnion() {
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionScorerType string

const (
	FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionScorerTypeScorer FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionScorerType = "scorer"
)

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionScorerType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataBundleLocationPositionScorerTypeScorer:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationType string

const (
	FunctionReplaceParamsFunctionDataCodeDataBundleLocationTypeExperiment FunctionReplaceParamsFunctionDataCodeDataBundleLocationType = "experiment"
)

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataBundleLocationTypeExperiment:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataCodeDataBundleRuntimeContext struct {
	Runtime param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleRuntimeContextRuntime] `json:"runtime,required"`
	Version param.Field[string]                                                               `json:"version,required"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleRuntimeContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsFunctionDataCodeDataBundleRuntimeContextRuntime string

const (
	FunctionReplaceParamsFunctionDataCodeDataBundleRuntimeContextRuntimeNode   FunctionReplaceParamsFunctionDataCodeDataBundleRuntimeContextRuntime = "node"
	FunctionReplaceParamsFunctionDataCodeDataBundleRuntimeContextRuntimePython FunctionReplaceParamsFunctionDataCodeDataBundleRuntimeContextRuntime = "python"
)

func (r FunctionReplaceParamsFunctionDataCodeDataBundleRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataBundleRuntimeContextRuntimeNode, FunctionReplaceParamsFunctionDataCodeDataBundleRuntimeContextRuntimePython:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataCodeDataBundleType string

const (
	FunctionReplaceParamsFunctionDataCodeDataBundleTypeBundle FunctionReplaceParamsFunctionDataCodeDataBundleType = "bundle"
)

func (r FunctionReplaceParamsFunctionDataCodeDataBundleType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataBundleTypeBundle:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataCodeDataInline struct {
	Code           param.Field[string]                                                        `json:"code,required"`
	RuntimeContext param.Field[FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContext] `json:"runtime_context,required"`
	Type           param.Field[FunctionReplaceParamsFunctionDataCodeDataInlineType]           `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataInline) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataCodeDataInline) implementsFunctionReplaceParamsFunctionDataCodeDataUnion() {
}

type FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContext struct {
	Runtime param.Field[FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContextRuntime] `json:"runtime,required"`
	Version param.Field[string]                                                               `json:"version,required"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContextRuntime string

const (
	FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContextRuntimeNode   FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContextRuntime = "node"
	FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContextRuntimePython FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContextRuntime = "python"
)

func (r FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContextRuntimeNode, FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContextRuntimePython:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataCodeDataInlineType string

const (
	FunctionReplaceParamsFunctionDataCodeDataInlineTypeInline FunctionReplaceParamsFunctionDataCodeDataInlineType = "inline"
)

func (r FunctionReplaceParamsFunctionDataCodeDataInlineType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataInlineTypeInline:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataCodeDataType string

const (
	FunctionReplaceParamsFunctionDataCodeDataTypeBundle FunctionReplaceParamsFunctionDataCodeDataType = "bundle"
	FunctionReplaceParamsFunctionDataCodeDataTypeInline FunctionReplaceParamsFunctionDataCodeDataType = "inline"
)

func (r FunctionReplaceParamsFunctionDataCodeDataType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataTypeBundle, FunctionReplaceParamsFunctionDataCodeDataTypeInline:
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

type FunctionReplaceParamsFunctionType string

const (
	FunctionReplaceParamsFunctionTypeTask   FunctionReplaceParamsFunctionType = "task"
	FunctionReplaceParamsFunctionTypeLlm    FunctionReplaceParamsFunctionType = "llm"
	FunctionReplaceParamsFunctionTypeScorer FunctionReplaceParamsFunctionType = "scorer"
)

func (r FunctionReplaceParamsFunctionType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionTypeTask, FunctionReplaceParamsFunctionTypeLlm, FunctionReplaceParamsFunctionTypeScorer:
		return true
	}
	return false
}

type FunctionReplaceParamsOrigin struct {
	// Id of the object the function is originating from
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[FunctionReplaceParamsOriginObjectType] `json:"object_type,required"`
	// The function exists for internal purposes and should not be displayed in the
	// list of functions.
	Internal param.Field[bool] `json:"internal"`
}

func (r FunctionReplaceParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The object type that the ACL applies to
type FunctionReplaceParamsOriginObjectType string

const (
	FunctionReplaceParamsOriginObjectTypeOrganization  FunctionReplaceParamsOriginObjectType = "organization"
	FunctionReplaceParamsOriginObjectTypeProject       FunctionReplaceParamsOriginObjectType = "project"
	FunctionReplaceParamsOriginObjectTypeExperiment    FunctionReplaceParamsOriginObjectType = "experiment"
	FunctionReplaceParamsOriginObjectTypeDataset       FunctionReplaceParamsOriginObjectType = "dataset"
	FunctionReplaceParamsOriginObjectTypePrompt        FunctionReplaceParamsOriginObjectType = "prompt"
	FunctionReplaceParamsOriginObjectTypePromptSession FunctionReplaceParamsOriginObjectType = "prompt_session"
	FunctionReplaceParamsOriginObjectTypeGroup         FunctionReplaceParamsOriginObjectType = "group"
	FunctionReplaceParamsOriginObjectTypeRole          FunctionReplaceParamsOriginObjectType = "role"
	FunctionReplaceParamsOriginObjectTypeOrgMember     FunctionReplaceParamsOriginObjectType = "org_member"
	FunctionReplaceParamsOriginObjectTypeProjectLog    FunctionReplaceParamsOriginObjectType = "project_log"
	FunctionReplaceParamsOriginObjectTypeOrgProject    FunctionReplaceParamsOriginObjectType = "org_project"
)

func (r FunctionReplaceParamsOriginObjectType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsOriginObjectTypeOrganization, FunctionReplaceParamsOriginObjectTypeProject, FunctionReplaceParamsOriginObjectTypeExperiment, FunctionReplaceParamsOriginObjectTypeDataset, FunctionReplaceParamsOriginObjectTypePrompt, FunctionReplaceParamsOriginObjectTypePromptSession, FunctionReplaceParamsOriginObjectTypeGroup, FunctionReplaceParamsOriginObjectTypeRole, FunctionReplaceParamsOriginObjectTypeOrgMember, FunctionReplaceParamsOriginObjectTypeProjectLog, FunctionReplaceParamsOriginObjectTypeOrgProject:
		return true
	}
	return false
}
