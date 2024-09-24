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

// Invoke a function.
func (r *FunctionService) Invoke(ctx context.Context, functionID string, body FunctionInvokeParams, opts ...option.RequestOption) (res *FunctionInvokeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required function_id parameter")
		return
	}
	path := fmt.Sprintf("v1/function/%s/invoke", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

type FunctionInvokeResponse = interface{}

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
	// JSON schema for the function's parameters and return type
	FunctionSchema param.Field[FunctionNewParamsFunctionSchema] `json:"function_schema"`
	FunctionType   param.Field[FunctionNewParamsFunctionType]   `json:"function_type"`
	Origin         param.Field[FunctionNewParamsOrigin]         `json:"origin"`
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
	Location       param.Field[FunctionNewParamsFunctionDataCodeDataBundleLocationUnion]  `json:"location,required"`
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
	Type     param.Field[FunctionNewParamsFunctionDataCodeDataBundleLocationType] `json:"type,required"`
	EvalName param.Field[string]                                                  `json:"eval_name"`
	Position param.Field[interface{}]                                             `json:"position,required"`
	Index    param.Field[int64]                                                   `json:"index"`
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocation) implementsFunctionNewParamsFunctionDataCodeDataBundleLocationUnion() {
}

// Satisfied by [FunctionNewParamsFunctionDataCodeDataBundleLocationExperiment],
// [FunctionNewParamsFunctionDataCodeDataBundleLocationFunction],
// [FunctionNewParamsFunctionDataCodeDataBundleLocation].
type FunctionNewParamsFunctionDataCodeDataBundleLocationUnion interface {
	implementsFunctionNewParamsFunctionDataCodeDataBundleLocationUnion()
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationExperiment struct {
	EvalName param.Field[string]                                                                     `json:"eval_name,required"`
	Position param.Field[FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionUnion] `json:"position,required"`
	Type     param.Field[FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentType]          `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationExperiment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationExperiment) implementsFunctionNewParamsFunctionDataCodeDataBundleLocationUnion() {
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPosition struct {
	Type  param.Field[FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionType] `json:"type,required"`
	Index param.Field[int64]                                                                     `json:"index"`
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPosition) implementsFunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionUnion() {
}

// Satisfied by
// [FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionType],
// [FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionScorer],
// [FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPosition].
type FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionUnion interface {
	implementsFunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionUnion()
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionType struct {
	Type param.Field[FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionTypeType] `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionType) implementsFunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionUnion() {
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionTypeType string

const (
	FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionTypeTypeTask FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionTypeType = "task"
)

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionTypeType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionTypeTypeTask:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionScorer struct {
	Index param.Field[int64]                                                                           `json:"index,required"`
	Type  param.Field[FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionScorerType] `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionScorer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionScorer) implementsFunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionUnion() {
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionScorerType string

const (
	FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionScorerTypeScorer FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionScorerType = "scorer"
)

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionScorerType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentPositionScorerTypeScorer:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentType string

const (
	FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentTypeExperiment FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentType = "experiment"
)

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataBundleLocationExperimentTypeExperiment:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationFunction struct {
	Index param.Field[int64]                                                           `json:"index,required"`
	Type  param.Field[FunctionNewParamsFunctionDataCodeDataBundleLocationFunctionType] `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationFunction) implementsFunctionNewParamsFunctionDataCodeDataBundleLocationUnion() {
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationFunctionType string

const (
	FunctionNewParamsFunctionDataCodeDataBundleLocationFunctionTypeFunction FunctionNewParamsFunctionDataCodeDataBundleLocationFunctionType = "function"
)

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationFunctionType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataBundleLocationFunctionTypeFunction:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataCodeDataBundleLocationType string

const (
	FunctionNewParamsFunctionDataCodeDataBundleLocationTypeExperiment FunctionNewParamsFunctionDataCodeDataBundleLocationType = "experiment"
	FunctionNewParamsFunctionDataCodeDataBundleLocationTypeFunction   FunctionNewParamsFunctionDataCodeDataBundleLocationType = "function"
)

func (r FunctionNewParamsFunctionDataCodeDataBundleLocationType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataCodeDataBundleLocationTypeExperiment, FunctionNewParamsFunctionDataCodeDataBundleLocationTypeFunction:
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

// JSON schema for the function's parameters and return type
type FunctionNewParamsFunctionSchema struct {
	Parameters param.Field[interface{}] `json:"parameters"`
	Returns    param.Field[interface{}] `json:"returns"`
}

func (r FunctionNewParamsFunctionSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsFunctionType string

const (
	FunctionNewParamsFunctionTypeLlm    FunctionNewParamsFunctionType = "llm"
	FunctionNewParamsFunctionTypeScorer FunctionNewParamsFunctionType = "scorer"
	FunctionNewParamsFunctionTypeTask   FunctionNewParamsFunctionType = "task"
	FunctionNewParamsFunctionTypeTool   FunctionNewParamsFunctionType = "tool"
)

func (r FunctionNewParamsFunctionType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionTypeLlm, FunctionNewParamsFunctionTypeScorer, FunctionNewParamsFunctionTypeTask, FunctionNewParamsFunctionTypeTool:
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
	Description  param.Field[string]                           `json:"description"`
	FunctionData param.Field[FunctionUpdateParamsFunctionData] `json:"function_data"`
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
}

func (r FunctionUpdateParamsFunctionData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type FunctionInvokeParams struct {
	// Argument to the function, which can be any JSON serializable value
	Input param.Field[interface{}] `json:"input"`
	// If the function is an LLM, additional messages to pass along to it
	Messages param.Field[[]FunctionInvokeParamsMessageUnion] `json:"messages"`
	// The mode format of the returned value (defaults to 'auto')
	Mode param.Field[FunctionInvokeParamsMode] `json:"mode"`
	// Options for tracing the function call
	Parent param.Field[FunctionInvokeParamsParentUnion] `json:"parent"`
	// Whether to stream the response. If true, results will be returned in the
	// Braintrust SSE format.
	Stream param.Field[bool] `json:"stream"`
	// The version of the function
	Version param.Field[string] `json:"version"`
}

func (r FunctionInvokeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionInvokeParamsMessage struct {
	Content      param.Field[interface{}]                      `json:"content,required"`
	Role         param.Field[FunctionInvokeParamsMessagesRole] `json:"role,required"`
	Name         param.Field[string]                           `json:"name"`
	FunctionCall param.Field[interface{}]                      `json:"function_call,required"`
	ToolCalls    param.Field[interface{}]                      `json:"tool_calls,required"`
	ToolCallID   param.Field[string]                           `json:"tool_call_id"`
}

func (r FunctionInvokeParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionInvokeParamsMessage) implementsFunctionInvokeParamsMessageUnion() {}

// Satisfied by [FunctionInvokeParamsMessagesSystem],
// [FunctionInvokeParamsMessagesUser], [FunctionInvokeParamsMessagesAssistant],
// [FunctionInvokeParamsMessagesTool], [FunctionInvokeParamsMessagesFunction],
// [FunctionInvokeParamsMessagesFallback], [FunctionInvokeParamsMessage].
type FunctionInvokeParamsMessageUnion interface {
	implementsFunctionInvokeParamsMessageUnion()
}

type FunctionInvokeParamsMessagesSystem struct {
	Role    param.Field[FunctionInvokeParamsMessagesSystemRole] `json:"role,required"`
	Content param.Field[string]                                 `json:"content"`
	Name    param.Field[string]                                 `json:"name"`
}

func (r FunctionInvokeParamsMessagesSystem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionInvokeParamsMessagesSystem) implementsFunctionInvokeParamsMessageUnion() {}

type FunctionInvokeParamsMessagesSystemRole string

const (
	FunctionInvokeParamsMessagesSystemRoleSystem FunctionInvokeParamsMessagesSystemRole = "system"
)

func (r FunctionInvokeParamsMessagesSystemRole) IsKnown() bool {
	switch r {
	case FunctionInvokeParamsMessagesSystemRoleSystem:
		return true
	}
	return false
}

type FunctionInvokeParamsMessagesUser struct {
	Role    param.Field[FunctionInvokeParamsMessagesUserRole]         `json:"role,required"`
	Content param.Field[FunctionInvokeParamsMessagesUserContentUnion] `json:"content"`
	Name    param.Field[string]                                       `json:"name"`
}

func (r FunctionInvokeParamsMessagesUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionInvokeParamsMessagesUser) implementsFunctionInvokeParamsMessageUnion() {}

type FunctionInvokeParamsMessagesUserRole string

const (
	FunctionInvokeParamsMessagesUserRoleUser FunctionInvokeParamsMessagesUserRole = "user"
)

func (r FunctionInvokeParamsMessagesUserRole) IsKnown() bool {
	switch r {
	case FunctionInvokeParamsMessagesUserRoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [FunctionInvokeParamsMessagesUserContentArray].
type FunctionInvokeParamsMessagesUserContentUnion interface {
	ImplementsFunctionInvokeParamsMessagesUserContentUnion()
}

type FunctionInvokeParamsMessagesUserContentArray []FunctionInvokeParamsMessagesUserContentArrayUnion

func (r FunctionInvokeParamsMessagesUserContentArray) ImplementsFunctionInvokeParamsMessagesUserContentUnion() {
}

type FunctionInvokeParamsMessagesUserContentArray struct {
	Text     param.Field[string]                                           `json:"text"`
	Type     param.Field[FunctionInvokeParamsMessagesUserContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                      `json:"image_url,required"`
}

func (r FunctionInvokeParamsMessagesUserContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionInvokeParamsMessagesUserContentArray) implementsFunctionInvokeParamsMessagesUserContentArrayUnion() {
}

// Satisfied by [FunctionInvokeParamsMessagesUserContentArrayText],
// [FunctionInvokeParamsMessagesUserContentArrayImageURL],
// [FunctionInvokeParamsMessagesUserContentArray].
type FunctionInvokeParamsMessagesUserContentArrayUnion interface {
	implementsFunctionInvokeParamsMessagesUserContentArrayUnion()
}

type FunctionInvokeParamsMessagesUserContentArrayText struct {
	Type param.Field[FunctionInvokeParamsMessagesUserContentArrayTextType] `json:"type,required"`
	Text param.Field[string]                                               `json:"text"`
}

func (r FunctionInvokeParamsMessagesUserContentArrayText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionInvokeParamsMessagesUserContentArrayText) implementsFunctionInvokeParamsMessagesUserContentArrayUnion() {
}

type FunctionInvokeParamsMessagesUserContentArrayTextType string

const (
	FunctionInvokeParamsMessagesUserContentArrayTextTypeText FunctionInvokeParamsMessagesUserContentArrayTextType = "text"
)

func (r FunctionInvokeParamsMessagesUserContentArrayTextType) IsKnown() bool {
	switch r {
	case FunctionInvokeParamsMessagesUserContentArrayTextTypeText:
		return true
	}
	return false
}

type FunctionInvokeParamsMessagesUserContentArrayImageURL struct {
	ImageURL param.Field[FunctionInvokeParamsMessagesUserContentArrayImageURLImageURL] `json:"image_url,required"`
	Type     param.Field[FunctionInvokeParamsMessagesUserContentArrayImageURLType]     `json:"type,required"`
}

func (r FunctionInvokeParamsMessagesUserContentArrayImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionInvokeParamsMessagesUserContentArrayImageURL) implementsFunctionInvokeParamsMessagesUserContentArrayUnion() {
}

type FunctionInvokeParamsMessagesUserContentArrayImageURLImageURL struct {
	URL    param.Field[string]                                                             `json:"url,required"`
	Detail param.Field[FunctionInvokeParamsMessagesUserContentArrayImageURLImageURLDetail] `json:"detail"`
}

func (r FunctionInvokeParamsMessagesUserContentArrayImageURLImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionInvokeParamsMessagesUserContentArrayImageURLImageURLDetail string

const (
	FunctionInvokeParamsMessagesUserContentArrayImageURLImageURLDetailAuto FunctionInvokeParamsMessagesUserContentArrayImageURLImageURLDetail = "auto"
	FunctionInvokeParamsMessagesUserContentArrayImageURLImageURLDetailLow  FunctionInvokeParamsMessagesUserContentArrayImageURLImageURLDetail = "low"
	FunctionInvokeParamsMessagesUserContentArrayImageURLImageURLDetailHigh FunctionInvokeParamsMessagesUserContentArrayImageURLImageURLDetail = "high"
)

func (r FunctionInvokeParamsMessagesUserContentArrayImageURLImageURLDetail) IsKnown() bool {
	switch r {
	case FunctionInvokeParamsMessagesUserContentArrayImageURLImageURLDetailAuto, FunctionInvokeParamsMessagesUserContentArrayImageURLImageURLDetailLow, FunctionInvokeParamsMessagesUserContentArrayImageURLImageURLDetailHigh:
		return true
	}
	return false
}

type FunctionInvokeParamsMessagesUserContentArrayImageURLType string

const (
	FunctionInvokeParamsMessagesUserContentArrayImageURLTypeImageURL FunctionInvokeParamsMessagesUserContentArrayImageURLType = "image_url"
)

func (r FunctionInvokeParamsMessagesUserContentArrayImageURLType) IsKnown() bool {
	switch r {
	case FunctionInvokeParamsMessagesUserContentArrayImageURLTypeImageURL:
		return true
	}
	return false
}

type FunctionInvokeParamsMessagesUserContentArrayType string

const (
	FunctionInvokeParamsMessagesUserContentArrayTypeText     FunctionInvokeParamsMessagesUserContentArrayType = "text"
	FunctionInvokeParamsMessagesUserContentArrayTypeImageURL FunctionInvokeParamsMessagesUserContentArrayType = "image_url"
)

func (r FunctionInvokeParamsMessagesUserContentArrayType) IsKnown() bool {
	switch r {
	case FunctionInvokeParamsMessagesUserContentArrayTypeText, FunctionInvokeParamsMessagesUserContentArrayTypeImageURL:
		return true
	}
	return false
}

type FunctionInvokeParamsMessagesAssistant struct {
	Role         param.Field[FunctionInvokeParamsMessagesAssistantRole]         `json:"role,required"`
	Content      param.Field[string]                                            `json:"content"`
	FunctionCall param.Field[FunctionInvokeParamsMessagesAssistantFunctionCall] `json:"function_call"`
	Name         param.Field[string]                                            `json:"name"`
	ToolCalls    param.Field[[]FunctionInvokeParamsMessagesAssistantToolCall]   `json:"tool_calls"`
}

func (r FunctionInvokeParamsMessagesAssistant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionInvokeParamsMessagesAssistant) implementsFunctionInvokeParamsMessageUnion() {}

type FunctionInvokeParamsMessagesAssistantRole string

const (
	FunctionInvokeParamsMessagesAssistantRoleAssistant FunctionInvokeParamsMessagesAssistantRole = "assistant"
)

func (r FunctionInvokeParamsMessagesAssistantRole) IsKnown() bool {
	switch r {
	case FunctionInvokeParamsMessagesAssistantRoleAssistant:
		return true
	}
	return false
}

type FunctionInvokeParamsMessagesAssistantFunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionInvokeParamsMessagesAssistantFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionInvokeParamsMessagesAssistantToolCall struct {
	ID       param.Field[string]                                                 `json:"id,required"`
	Function param.Field[FunctionInvokeParamsMessagesAssistantToolCallsFunction] `json:"function,required"`
	Type     param.Field[FunctionInvokeParamsMessagesAssistantToolCallsType]     `json:"type,required"`
}

func (r FunctionInvokeParamsMessagesAssistantToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionInvokeParamsMessagesAssistantToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionInvokeParamsMessagesAssistantToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionInvokeParamsMessagesAssistantToolCallsType string

const (
	FunctionInvokeParamsMessagesAssistantToolCallsTypeFunction FunctionInvokeParamsMessagesAssistantToolCallsType = "function"
)

func (r FunctionInvokeParamsMessagesAssistantToolCallsType) IsKnown() bool {
	switch r {
	case FunctionInvokeParamsMessagesAssistantToolCallsTypeFunction:
		return true
	}
	return false
}

type FunctionInvokeParamsMessagesTool struct {
	Role       param.Field[FunctionInvokeParamsMessagesToolRole] `json:"role,required"`
	Content    param.Field[string]                               `json:"content"`
	ToolCallID param.Field[string]                               `json:"tool_call_id"`
}

func (r FunctionInvokeParamsMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionInvokeParamsMessagesTool) implementsFunctionInvokeParamsMessageUnion() {}

type FunctionInvokeParamsMessagesToolRole string

const (
	FunctionInvokeParamsMessagesToolRoleTool FunctionInvokeParamsMessagesToolRole = "tool"
)

func (r FunctionInvokeParamsMessagesToolRole) IsKnown() bool {
	switch r {
	case FunctionInvokeParamsMessagesToolRoleTool:
		return true
	}
	return false
}

type FunctionInvokeParamsMessagesFunction struct {
	Name    param.Field[string]                                   `json:"name,required"`
	Role    param.Field[FunctionInvokeParamsMessagesFunctionRole] `json:"role,required"`
	Content param.Field[string]                                   `json:"content"`
}

func (r FunctionInvokeParamsMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionInvokeParamsMessagesFunction) implementsFunctionInvokeParamsMessageUnion() {}

type FunctionInvokeParamsMessagesFunctionRole string

const (
	FunctionInvokeParamsMessagesFunctionRoleFunction FunctionInvokeParamsMessagesFunctionRole = "function"
)

func (r FunctionInvokeParamsMessagesFunctionRole) IsKnown() bool {
	switch r {
	case FunctionInvokeParamsMessagesFunctionRoleFunction:
		return true
	}
	return false
}

type FunctionInvokeParamsMessagesFallback struct {
	Role    param.Field[FunctionInvokeParamsMessagesFallbackRole] `json:"role,required"`
	Content param.Field[string]                                   `json:"content"`
}

func (r FunctionInvokeParamsMessagesFallback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionInvokeParamsMessagesFallback) implementsFunctionInvokeParamsMessageUnion() {}

type FunctionInvokeParamsMessagesFallbackRole string

const (
	FunctionInvokeParamsMessagesFallbackRoleModel FunctionInvokeParamsMessagesFallbackRole = "model"
)

func (r FunctionInvokeParamsMessagesFallbackRole) IsKnown() bool {
	switch r {
	case FunctionInvokeParamsMessagesFallbackRoleModel:
		return true
	}
	return false
}

type FunctionInvokeParamsMessagesRole string

const (
	FunctionInvokeParamsMessagesRoleSystem    FunctionInvokeParamsMessagesRole = "system"
	FunctionInvokeParamsMessagesRoleUser      FunctionInvokeParamsMessagesRole = "user"
	FunctionInvokeParamsMessagesRoleAssistant FunctionInvokeParamsMessagesRole = "assistant"
	FunctionInvokeParamsMessagesRoleTool      FunctionInvokeParamsMessagesRole = "tool"
	FunctionInvokeParamsMessagesRoleFunction  FunctionInvokeParamsMessagesRole = "function"
	FunctionInvokeParamsMessagesRoleModel     FunctionInvokeParamsMessagesRole = "model"
)

func (r FunctionInvokeParamsMessagesRole) IsKnown() bool {
	switch r {
	case FunctionInvokeParamsMessagesRoleSystem, FunctionInvokeParamsMessagesRoleUser, FunctionInvokeParamsMessagesRoleAssistant, FunctionInvokeParamsMessagesRoleTool, FunctionInvokeParamsMessagesRoleFunction, FunctionInvokeParamsMessagesRoleModel:
		return true
	}
	return false
}

// The mode format of the returned value (defaults to 'auto')
type FunctionInvokeParamsMode string

const (
	FunctionInvokeParamsModeAuto     FunctionInvokeParamsMode = "auto"
	FunctionInvokeParamsModeParallel FunctionInvokeParamsMode = "parallel"
)

func (r FunctionInvokeParamsMode) IsKnown() bool {
	switch r {
	case FunctionInvokeParamsModeAuto, FunctionInvokeParamsModeParallel:
		return true
	}
	return false
}

// Options for tracing the function call
//
// Satisfied by [FunctionInvokeParamsParentSpanParentStruct], [shared.UnionString].
type FunctionInvokeParamsParentUnion interface {
	ImplementsFunctionInvokeParamsParentUnion()
}

// Span parent properties
type FunctionInvokeParamsParentSpanParentStruct struct {
	// The id of the container object you are logging to
	ObjectID   param.Field[string]                                               `json:"object_id,required"`
	ObjectType param.Field[FunctionInvokeParamsParentSpanParentStructObjectType] `json:"object_type,required"`
	// Include these properties in every span created under this parent
	PropagatedEvent param.Field[map[string]interface{}] `json:"propagated_event"`
	// Identifiers for the row to to log a subspan under
	RowIDs param.Field[FunctionInvokeParamsParentSpanParentStructRowIDs] `json:"row_ids"`
}

func (r FunctionInvokeParamsParentSpanParentStruct) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionInvokeParamsParentSpanParentStruct) ImplementsFunctionInvokeParamsParentUnion() {}

type FunctionInvokeParamsParentSpanParentStructObjectType string

const (
	FunctionInvokeParamsParentSpanParentStructObjectTypeProjectLogs FunctionInvokeParamsParentSpanParentStructObjectType = "project_logs"
	FunctionInvokeParamsParentSpanParentStructObjectTypeExperiment  FunctionInvokeParamsParentSpanParentStructObjectType = "experiment"
)

func (r FunctionInvokeParamsParentSpanParentStructObjectType) IsKnown() bool {
	switch r {
	case FunctionInvokeParamsParentSpanParentStructObjectTypeProjectLogs, FunctionInvokeParamsParentSpanParentStructObjectTypeExperiment:
		return true
	}
	return false
}

// Identifiers for the row to to log a subspan under
type FunctionInvokeParamsParentSpanParentStructRowIDs struct {
	// The id of the row
	ID param.Field[string] `json:"id,required"`
	// The root_span_id of the row
	RootSpanID param.Field[string] `json:"root_span_id,required"`
	// The span_id of the row
	SpanID param.Field[string] `json:"span_id,required"`
}

func (r FunctionInvokeParamsParentSpanParentStructRowIDs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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
	// JSON schema for the function's parameters and return type
	FunctionSchema param.Field[FunctionReplaceParamsFunctionSchema] `json:"function_schema"`
	FunctionType   param.Field[FunctionReplaceParamsFunctionType]   `json:"function_type"`
	Origin         param.Field[FunctionReplaceParamsOrigin]         `json:"origin"`
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
	Location       param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleLocationUnion]  `json:"location,required"`
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
	Type     param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleLocationType] `json:"type,required"`
	EvalName param.Field[string]                                                      `json:"eval_name"`
	Position param.Field[interface{}]                                                 `json:"position,required"`
	Index    param.Field[int64]                                                       `json:"index"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocation) implementsFunctionReplaceParamsFunctionDataCodeDataBundleLocationUnion() {
}

// Satisfied by
// [FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperiment],
// [FunctionReplaceParamsFunctionDataCodeDataBundleLocationFunction],
// [FunctionReplaceParamsFunctionDataCodeDataBundleLocation].
type FunctionReplaceParamsFunctionDataCodeDataBundleLocationUnion interface {
	implementsFunctionReplaceParamsFunctionDataCodeDataBundleLocationUnion()
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperiment struct {
	EvalName param.Field[string]                                                                         `json:"eval_name,required"`
	Position param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionUnion] `json:"position,required"`
	Type     param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentType]          `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperiment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperiment) implementsFunctionReplaceParamsFunctionDataCodeDataBundleLocationUnion() {
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPosition struct {
	Type  param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionType] `json:"type,required"`
	Index param.Field[int64]                                                                         `json:"index"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPosition) implementsFunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionUnion() {
}

// Satisfied by
// [FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionType],
// [FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionScorer],
// [FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPosition].
type FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionUnion interface {
	implementsFunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionUnion()
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionType struct {
	Type param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionTypeType] `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionType) implementsFunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionUnion() {
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionTypeType string

const (
	FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionTypeTypeTask FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionTypeType = "task"
)

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionTypeType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionTypeTypeTask:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionScorer struct {
	Index param.Field[int64]                                                                               `json:"index,required"`
	Type  param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionScorerType] `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionScorer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionScorer) implementsFunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionUnion() {
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionScorerType string

const (
	FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionScorerTypeScorer FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionScorerType = "scorer"
)

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionScorerType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentPositionScorerTypeScorer:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentType string

const (
	FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentTypeExperiment FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentType = "experiment"
)

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataBundleLocationExperimentTypeExperiment:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationFunction struct {
	Index param.Field[int64]                                                               `json:"index,required"`
	Type  param.Field[FunctionReplaceParamsFunctionDataCodeDataBundleLocationFunctionType] `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationFunction) implementsFunctionReplaceParamsFunctionDataCodeDataBundleLocationUnion() {
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationFunctionType string

const (
	FunctionReplaceParamsFunctionDataCodeDataBundleLocationFunctionTypeFunction FunctionReplaceParamsFunctionDataCodeDataBundleLocationFunctionType = "function"
)

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationFunctionType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataBundleLocationFunctionTypeFunction:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataCodeDataBundleLocationType string

const (
	FunctionReplaceParamsFunctionDataCodeDataBundleLocationTypeExperiment FunctionReplaceParamsFunctionDataCodeDataBundleLocationType = "experiment"
	FunctionReplaceParamsFunctionDataCodeDataBundleLocationTypeFunction   FunctionReplaceParamsFunctionDataCodeDataBundleLocationType = "function"
)

func (r FunctionReplaceParamsFunctionDataCodeDataBundleLocationType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataCodeDataBundleLocationTypeExperiment, FunctionReplaceParamsFunctionDataCodeDataBundleLocationTypeFunction:
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

// JSON schema for the function's parameters and return type
type FunctionReplaceParamsFunctionSchema struct {
	Parameters param.Field[interface{}] `json:"parameters"`
	Returns    param.Field[interface{}] `json:"returns"`
}

func (r FunctionReplaceParamsFunctionSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsFunctionType string

const (
	FunctionReplaceParamsFunctionTypeLlm    FunctionReplaceParamsFunctionType = "llm"
	FunctionReplaceParamsFunctionTypeScorer FunctionReplaceParamsFunctionType = "scorer"
	FunctionReplaceParamsFunctionTypeTask   FunctionReplaceParamsFunctionType = "task"
	FunctionReplaceParamsFunctionTypeTool   FunctionReplaceParamsFunctionType = "tool"
)

func (r FunctionReplaceParamsFunctionType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionTypeLlm, FunctionReplaceParamsFunctionTypeScorer, FunctionReplaceParamsFunctionTypeTask, FunctionReplaceParamsFunctionTypeTool:
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
