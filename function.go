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
	"github.com/braintrustdata/braintrust-go/shared"
	"github.com/tidwall/gjson"
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
func (r *FunctionService) New(ctx context.Context, body FunctionNewParams, opts ...option.RequestOption) (res *Function, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/function"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a function object by its id
func (r *FunctionService) Get(ctx context.Context, functionID string, opts ...option.RequestOption) (res *Function, err error) {
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
func (r *FunctionService) Update(ctx context.Context, functionID string, body FunctionUpdateParams, opts ...option.RequestOption) (res *Function, err error) {
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
func (r *FunctionService) List(ctx context.Context, query FunctionListParams, opts ...option.RequestOption) (res *pagination.ListObjects[Function], err error) {
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
func (r *FunctionService) ListAutoPaging(ctx context.Context, query FunctionListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[Function] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a function object by its id
func (r *FunctionService) Delete(ctx context.Context, functionID string, opts ...option.RequestOption) (res *Function, err error) {
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
func (r *FunctionService) Replace(ctx context.Context, body FunctionReplaceParams, opts ...option.RequestOption) (res *Function, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/function"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type Function struct {
	// Unique identifier for the prompt
	ID string `json:"id,required" format:"uuid"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the prompt (see the `version`
	// parameter)
	XactID       string               `json:"_xact_id,required"`
	FunctionData FunctionFunctionData `json:"function_data,required"`
	// A literal 'p' which identifies the object as a project prompt
	LogID FunctionLogID `json:"log_id,required"`
	// Name of the prompt
	Name string `json:"name,required"`
	// Unique identifier for the organization
	OrgID string `json:"org_id,required" format:"uuid"`
	// Unique identifier for the project that the prompt belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// Unique identifier for the prompt
	Slug string `json:"slug,required"`
	// Date of prompt creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Textual description of the prompt
	Description string `json:"description,nullable"`
	// User-controlled metadata about the prompt
	Metadata map[string]interface{} `json:"metadata,nullable"`
	// The prompt, model, and its parameters
	PromptData FunctionPromptData `json:"prompt_data,nullable"`
	// A list of tags for the prompt
	Tags []string     `json:"tags,nullable"`
	JSON functionJSON `json:"-"`
}

// functionJSON contains the JSON metadata for the struct [Function]
type functionJSON struct {
	ID           apijson.Field
	XactID       apijson.Field
	FunctionData apijson.Field
	LogID        apijson.Field
	Name         apijson.Field
	OrgID        apijson.Field
	ProjectID    apijson.Field
	Slug         apijson.Field
	Created      apijson.Field
	Description  apijson.Field
	Metadata     apijson.Field
	PromptData   apijson.Field
	Tags         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Function) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionJSON) RawJSON() string {
	return r.raw
}

type FunctionFunctionData struct {
	Type FunctionFunctionDataType `json:"type,required"`
	// This field can have the runtime type of [FunctionFunctionDataCodeData].
	Data  interface{}              `json:"data,required"`
	Name  string                   `json:"name"`
	JSON  functionFunctionDataJSON `json:"-"`
	union FunctionFunctionDataUnion
}

// functionFunctionDataJSON contains the JSON metadata for the struct
// [FunctionFunctionData]
type functionFunctionDataJSON struct {
	Type        apijson.Field
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r functionFunctionDataJSON) RawJSON() string {
	return r.raw
}

func (r *FunctionFunctionData) UnmarshalJSON(data []byte) (err error) {
	*r = FunctionFunctionData{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FunctionFunctionDataUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [FunctionFunctionDataPrompt],
// [FunctionFunctionDataCode], [FunctionFunctionDataGlobal].
func (r FunctionFunctionData) AsUnion() FunctionFunctionDataUnion {
	return r.union
}

// Union satisfied by [FunctionFunctionDataPrompt], [FunctionFunctionDataCode] or
// [FunctionFunctionDataGlobal].
type FunctionFunctionDataUnion interface {
	implementsFunctionFunctionData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionFunctionDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataPrompt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataCode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataGlobal{}),
		},
	)
}

type FunctionFunctionDataPrompt struct {
	Type FunctionFunctionDataPromptType `json:"type,required"`
	JSON functionFunctionDataPromptJSON `json:"-"`
}

// functionFunctionDataPromptJSON contains the JSON metadata for the struct
// [FunctionFunctionDataPrompt]
type functionFunctionDataPromptJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataPromptJSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataPrompt) implementsFunctionFunctionData() {}

type FunctionFunctionDataPromptType string

const (
	FunctionFunctionDataPromptTypePrompt FunctionFunctionDataPromptType = "prompt"
)

func (r FunctionFunctionDataPromptType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataPromptTypePrompt:
		return true
	}
	return false
}

type FunctionFunctionDataCode struct {
	Data FunctionFunctionDataCodeData `json:"data,required"`
	Type FunctionFunctionDataCodeType `json:"type,required"`
	JSON functionFunctionDataCodeJSON `json:"-"`
}

// functionFunctionDataCodeJSON contains the JSON metadata for the struct
// [FunctionFunctionDataCode]
type functionFunctionDataCodeJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataCodeJSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataCode) implementsFunctionFunctionData() {}

type FunctionFunctionDataCodeData struct {
	BundleID       string                                     `json:"bundle_id,required"`
	Location       FunctionFunctionDataCodeDataLocation       `json:"location,required"`
	RuntimeContext FunctionFunctionDataCodeDataRuntimeContext `json:"runtime_context,required"`
	JSON           functionFunctionDataCodeDataJSON           `json:"-"`
}

// functionFunctionDataCodeDataJSON contains the JSON metadata for the struct
// [FunctionFunctionDataCodeData]
type functionFunctionDataCodeDataJSON struct {
	BundleID       apijson.Field
	Location       apijson.Field
	RuntimeContext apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FunctionFunctionDataCodeData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataCodeDataJSON) RawJSON() string {
	return r.raw
}

type FunctionFunctionDataCodeDataLocation struct {
	EvalName string                                            `json:"eval_name,required"`
	Position FunctionFunctionDataCodeDataLocationPositionUnion `json:"position,required"`
	Type     FunctionFunctionDataCodeDataLocationType          `json:"type,required"`
	JSON     functionFunctionDataCodeDataLocationJSON          `json:"-"`
}

// functionFunctionDataCodeDataLocationJSON contains the JSON metadata for the
// struct [FunctionFunctionDataCodeDataLocation]
type functionFunctionDataCodeDataLocationJSON struct {
	EvalName    apijson.Field
	Position    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataCodeDataLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataCodeDataLocationJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [FunctionFunctionDataCodeDataLocationPositionTask] or
// [FunctionFunctionDataCodeDataLocationPositionScore].
type FunctionFunctionDataCodeDataLocationPositionUnion interface {
	implementsFunctionFunctionDataCodeDataLocationPositionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionFunctionDataCodeDataLocationPositionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionFunctionDataCodeDataLocationPositionTask("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataCodeDataLocationPositionScore{}),
		},
	)
}

type FunctionFunctionDataCodeDataLocationPositionTask string

const (
	FunctionFunctionDataCodeDataLocationPositionTaskTask FunctionFunctionDataCodeDataLocationPositionTask = "task"
)

func (r FunctionFunctionDataCodeDataLocationPositionTask) IsKnown() bool {
	switch r {
	case FunctionFunctionDataCodeDataLocationPositionTaskTask:
		return true
	}
	return false
}

func (r FunctionFunctionDataCodeDataLocationPositionTask) implementsFunctionFunctionDataCodeDataLocationPositionUnion() {
}

type FunctionFunctionDataCodeDataLocationPositionScore struct {
	Score float64                                               `json:"score,required"`
	JSON  functionFunctionDataCodeDataLocationPositionScoreJSON `json:"-"`
}

// functionFunctionDataCodeDataLocationPositionScoreJSON contains the JSON metadata
// for the struct [FunctionFunctionDataCodeDataLocationPositionScore]
type functionFunctionDataCodeDataLocationPositionScoreJSON struct {
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataCodeDataLocationPositionScore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataCodeDataLocationPositionScoreJSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataCodeDataLocationPositionScore) implementsFunctionFunctionDataCodeDataLocationPositionUnion() {
}

type FunctionFunctionDataCodeDataLocationType string

const (
	FunctionFunctionDataCodeDataLocationTypeExperiment FunctionFunctionDataCodeDataLocationType = "experiment"
)

func (r FunctionFunctionDataCodeDataLocationType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataCodeDataLocationTypeExperiment:
		return true
	}
	return false
}

type FunctionFunctionDataCodeDataRuntimeContext struct {
	Runtime FunctionFunctionDataCodeDataRuntimeContextRuntime `json:"runtime,required"`
	Version string                                            `json:"version,required"`
	JSON    functionFunctionDataCodeDataRuntimeContextJSON    `json:"-"`
}

// functionFunctionDataCodeDataRuntimeContextJSON contains the JSON metadata for
// the struct [FunctionFunctionDataCodeDataRuntimeContext]
type functionFunctionDataCodeDataRuntimeContextJSON struct {
	Runtime     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataCodeDataRuntimeContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataCodeDataRuntimeContextJSON) RawJSON() string {
	return r.raw
}

type FunctionFunctionDataCodeDataRuntimeContextRuntime string

const (
	FunctionFunctionDataCodeDataRuntimeContextRuntimeNode FunctionFunctionDataCodeDataRuntimeContextRuntime = "node"
)

func (r FunctionFunctionDataCodeDataRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionFunctionDataCodeDataRuntimeContextRuntimeNode:
		return true
	}
	return false
}

type FunctionFunctionDataCodeType string

const (
	FunctionFunctionDataCodeTypeCode FunctionFunctionDataCodeType = "code"
)

func (r FunctionFunctionDataCodeType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataCodeTypeCode:
		return true
	}
	return false
}

type FunctionFunctionDataGlobal struct {
	Name string                         `json:"name,required"`
	Type FunctionFunctionDataGlobalType `json:"type,required"`
	JSON functionFunctionDataGlobalJSON `json:"-"`
}

// functionFunctionDataGlobalJSON contains the JSON metadata for the struct
// [FunctionFunctionDataGlobal]
type functionFunctionDataGlobalJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataGlobal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataGlobalJSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataGlobal) implementsFunctionFunctionData() {}

type FunctionFunctionDataGlobalType string

const (
	FunctionFunctionDataGlobalTypeGlobal FunctionFunctionDataGlobalType = "global"
)

func (r FunctionFunctionDataGlobalType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataGlobalTypeGlobal:
		return true
	}
	return false
}

type FunctionFunctionDataType string

const (
	FunctionFunctionDataTypePrompt FunctionFunctionDataType = "prompt"
	FunctionFunctionDataTypeCode   FunctionFunctionDataType = "code"
	FunctionFunctionDataTypeGlobal FunctionFunctionDataType = "global"
)

func (r FunctionFunctionDataType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataTypePrompt, FunctionFunctionDataTypeCode, FunctionFunctionDataTypeGlobal:
		return true
	}
	return false
}

// A literal 'p' which identifies the object as a project prompt
type FunctionLogID string

const (
	FunctionLogIDP FunctionLogID = "p"
)

func (r FunctionLogID) IsKnown() bool {
	switch r {
	case FunctionLogIDP:
		return true
	}
	return false
}

// The prompt, model, and its parameters
type FunctionPromptData struct {
	Options FunctionPromptDataOptions `json:"options,nullable"`
	Origin  FunctionPromptDataOrigin  `json:"origin,nullable"`
	Prompt  FunctionPromptDataPrompt  `json:"prompt"`
	JSON    functionPromptDataJSON    `json:"-"`
}

// functionPromptDataJSON contains the JSON metadata for the struct
// [FunctionPromptData]
type functionPromptDataJSON struct {
	Options     apijson.Field
	Origin      apijson.Field
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataOptions struct {
	Model    string                          `json:"model"`
	Params   FunctionPromptDataOptionsParams `json:"params"`
	Position string                          `json:"position"`
	JSON     functionPromptDataOptionsJSON   `json:"-"`
}

// functionPromptDataOptionsJSON contains the JSON metadata for the struct
// [FunctionPromptDataOptions]
type functionPromptDataOptionsJSON struct {
	Model       apijson.Field
	Params      apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataOptionsParams struct {
	UseCache         bool    `json:"use_cache"`
	Temperature      float64 `json:"temperature"`
	TopP             float64 `json:"top_p"`
	MaxTokens        float64 `json:"max_tokens"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
	// This field can have the runtime type of
	// [FunctionPromptDataOptionsParamsOpenAIModelParamsResponseFormat].
	ResponseFormat interface{} `json:"response_format,required"`
	// This field can have the runtime type of
	// [FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion].
	ToolChoice interface{} `json:"tool_choice,required"`
	// This field can have the runtime type of
	// [FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion].
	FunctionCall interface{} `json:"function_call,required"`
	N            float64     `json:"n"`
	// This field can have the runtime type of [[]string].
	Stop interface{} `json:"stop,required"`
	TopK float64     `json:"top_k"`
	// This field can have the runtime type of [[]string].
	StopSequences interface{} `json:"stop_sequences,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample float64                             `json:"max_tokens_to_sample"`
	MaxOutputTokens   float64                             `json:"maxOutputTokens"`
	TopP              float64                             `json:"topP"`
	TopK              float64                             `json:"topK"`
	JSON              functionPromptDataOptionsParamsJSON `json:"-"`
	union             FunctionPromptDataOptionsParamsUnion
}

// functionPromptDataOptionsParamsJSON contains the JSON metadata for the struct
// [FunctionPromptDataOptionsParams]
type functionPromptDataOptionsParamsJSON struct {
	UseCache          apijson.Field
	Temperature       apijson.Field
	TopP              apijson.Field
	MaxTokens         apijson.Field
	FrequencyPenalty  apijson.Field
	PresencePenalty   apijson.Field
	ResponseFormat    apijson.Field
	ToolChoice        apijson.Field
	FunctionCall      apijson.Field
	N                 apijson.Field
	Stop              apijson.Field
	TopK              apijson.Field
	StopSequences     apijson.Field
	MaxTokensToSample apijson.Field
	MaxOutputTokens   apijson.Field
	TopP              apijson.Field
	TopK              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r functionPromptDataOptionsParamsJSON) RawJSON() string {
	return r.raw
}

func (r *FunctionPromptDataOptionsParams) UnmarshalJSON(data []byte) (err error) {
	*r = FunctionPromptDataOptionsParams{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FunctionPromptDataOptionsParamsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [FunctionPromptDataOptionsParamsOpenAIModelParams],
// [FunctionPromptDataOptionsParamsAnthropicModelParams],
// [FunctionPromptDataOptionsParamsGoogleModelParams],
// [FunctionPromptDataOptionsParamsWindowAIModelParams],
// [FunctionPromptDataOptionsParamsJsCompletionParams].
func (r FunctionPromptDataOptionsParams) AsUnion() FunctionPromptDataOptionsParamsUnion {
	return r.union
}

// Union satisfied by [FunctionPromptDataOptionsParamsOpenAIModelParams],
// [FunctionPromptDataOptionsParamsAnthropicModelParams],
// [FunctionPromptDataOptionsParamsGoogleModelParams],
// [FunctionPromptDataOptionsParamsWindowAIModelParams] or
// [FunctionPromptDataOptionsParamsJsCompletionParams].
type FunctionPromptDataOptionsParamsUnion interface {
	implementsFunctionPromptDataOptionsParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataOptionsParamsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsOpenAIModelParams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsAnthropicModelParams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsGoogleModelParams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsWindowAIModelParams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsJsCompletionParams{}),
		},
	)
}

type FunctionPromptDataOptionsParamsOpenAIModelParams struct {
	FrequencyPenalty float64                                                           `json:"frequency_penalty"`
	FunctionCall     FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion `json:"function_call"`
	MaxTokens        float64                                                           `json:"max_tokens"`
	N                float64                                                           `json:"n"`
	PresencePenalty  float64                                                           `json:"presence_penalty"`
	ResponseFormat   FunctionPromptDataOptionsParamsOpenAIModelParamsResponseFormat    `json:"response_format,nullable"`
	Stop             []string                                                          `json:"stop"`
	Temperature      float64                                                           `json:"temperature"`
	ToolChoice       FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion   `json:"tool_choice"`
	TopP             float64                                                           `json:"top_p"`
	UseCache         bool                                                              `json:"use_cache"`
	JSON             functionPromptDataOptionsParamsOpenAIModelParamsJSON              `json:"-"`
}

// functionPromptDataOptionsParamsOpenAIModelParamsJSON contains the JSON metadata
// for the struct [FunctionPromptDataOptionsParamsOpenAIModelParams]
type functionPromptDataOptionsParamsOpenAIModelParamsJSON struct {
	FrequencyPenalty apijson.Field
	FunctionCall     apijson.Field
	MaxTokens        apijson.Field
	N                apijson.Field
	PresencePenalty  apijson.Field
	ResponseFormat   apijson.Field
	Stop             apijson.Field
	Temperature      apijson.Field
	ToolChoice       apijson.Field
	TopP             apijson.Field
	UseCache         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsOpenAIModelParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsOpenAIModelParamsJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsOpenAIModelParams) implementsFunctionPromptDataOptionsParams() {
}

// Union satisfied by
// [FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallString],
// [FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallString] or
// [FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject].
type FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsFunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject{}),
		},
	)
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallString string

const (
	FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallString = "auto"
)

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto:
		return true
	}
	return false
}

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) implementsFunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject struct {
	Name string                                                                 `json:"name,required"`
	JSON functionPromptDataOptionsParamsOpenAIModelParamsFunctionCallObjectJSON `json:"-"`
}

// functionPromptDataOptionsParamsOpenAIModelParamsFunctionCallObjectJSON contains
// the JSON metadata for the struct
// [FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject]
type functionPromptDataOptionsParamsOpenAIModelParamsFunctionCallObjectJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsOpenAIModelParamsFunctionCallObjectJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) implementsFunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsResponseFormat struct {
	Type FunctionPromptDataOptionsParamsOpenAIModelParamsResponseFormatType `json:"type,required"`
	JSON functionPromptDataOptionsParamsOpenAIModelParamsResponseFormatJSON `json:"-"`
}

// functionPromptDataOptionsParamsOpenAIModelParamsResponseFormatJSON contains the
// JSON metadata for the struct
// [FunctionPromptDataOptionsParamsOpenAIModelParamsResponseFormat]
type functionPromptDataOptionsParamsOpenAIModelParamsResponseFormatJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsOpenAIModelParamsResponseFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsOpenAIModelParamsResponseFormatJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsResponseFormatType string

const (
	FunctionPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject FunctionPromptDataOptionsParamsOpenAIModelParamsResponseFormatType = "json_object"
)

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsResponseFormatType) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Union satisfied by
// [FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceString],
// [FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceString] or
// [FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject].
type FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsFunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject{}),
		},
	)
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceString string

const (
	FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceString = "auto"
)

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto:
		return true
	}
	return false
}

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) implementsFunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject struct {
	Function FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction `json:"function,required"`
	Type     FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType     `json:"type,required"`
	JSON     functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectJSON     `json:"-"`
}

// functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectJSON contains
// the JSON metadata for the struct
// [FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject]
type functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectJSON struct {
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) implementsFunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction struct {
	Name string                                                                       `json:"name,required"`
	JSON functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunctionJSON `json:"-"`
}

// functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunctionJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction]
type functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunctionJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType string

const (
	FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType = "function"
)

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction:
		return true
	}
	return false
}

type FunctionPromptDataOptionsParamsAnthropicModelParams struct {
	MaxTokens   float64 `json:"max_tokens,required"`
	Temperature float64 `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample float64                                                 `json:"max_tokens_to_sample"`
	StopSequences     []string                                                `json:"stop_sequences"`
	TopK              float64                                                 `json:"top_k"`
	TopP              float64                                                 `json:"top_p"`
	UseCache          bool                                                    `json:"use_cache"`
	JSON              functionPromptDataOptionsParamsAnthropicModelParamsJSON `json:"-"`
}

// functionPromptDataOptionsParamsAnthropicModelParamsJSON contains the JSON
// metadata for the struct [FunctionPromptDataOptionsParamsAnthropicModelParams]
type functionPromptDataOptionsParamsAnthropicModelParamsJSON struct {
	MaxTokens         apijson.Field
	Temperature       apijson.Field
	MaxTokensToSample apijson.Field
	StopSequences     apijson.Field
	TopK              apijson.Field
	TopP              apijson.Field
	UseCache          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsAnthropicModelParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsAnthropicModelParamsJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsAnthropicModelParams) implementsFunctionPromptDataOptionsParams() {
}

type FunctionPromptDataOptionsParamsGoogleModelParams struct {
	MaxOutputTokens float64                                              `json:"maxOutputTokens"`
	Temperature     float64                                              `json:"temperature"`
	TopK            float64                                              `json:"topK"`
	TopP            float64                                              `json:"topP"`
	UseCache        bool                                                 `json:"use_cache"`
	JSON            functionPromptDataOptionsParamsGoogleModelParamsJSON `json:"-"`
}

// functionPromptDataOptionsParamsGoogleModelParamsJSON contains the JSON metadata
// for the struct [FunctionPromptDataOptionsParamsGoogleModelParams]
type functionPromptDataOptionsParamsGoogleModelParamsJSON struct {
	MaxOutputTokens apijson.Field
	Temperature     apijson.Field
	TopK            apijson.Field
	TopP            apijson.Field
	UseCache        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsGoogleModelParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsGoogleModelParamsJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsGoogleModelParams) implementsFunctionPromptDataOptionsParams() {
}

type FunctionPromptDataOptionsParamsWindowAIModelParams struct {
	Temperature float64                                                `json:"temperature"`
	TopK        float64                                                `json:"topK"`
	UseCache    bool                                                   `json:"use_cache"`
	JSON        functionPromptDataOptionsParamsWindowAIModelParamsJSON `json:"-"`
}

// functionPromptDataOptionsParamsWindowAIModelParamsJSON contains the JSON
// metadata for the struct [FunctionPromptDataOptionsParamsWindowAIModelParams]
type functionPromptDataOptionsParamsWindowAIModelParamsJSON struct {
	Temperature apijson.Field
	TopK        apijson.Field
	UseCache    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsWindowAIModelParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsWindowAIModelParamsJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsWindowAIModelParams) implementsFunctionPromptDataOptionsParams() {
}

type FunctionPromptDataOptionsParamsJsCompletionParams struct {
	UseCache bool                                                  `json:"use_cache"`
	JSON     functionPromptDataOptionsParamsJsCompletionParamsJSON `json:"-"`
}

// functionPromptDataOptionsParamsJsCompletionParamsJSON contains the JSON metadata
// for the struct [FunctionPromptDataOptionsParamsJsCompletionParams]
type functionPromptDataOptionsParamsJsCompletionParamsJSON struct {
	UseCache    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsJsCompletionParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsJsCompletionParamsJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsJsCompletionParams) implementsFunctionPromptDataOptionsParams() {
}

type FunctionPromptDataOrigin struct {
	ProjectID     string                       `json:"project_id"`
	PromptID      string                       `json:"prompt_id"`
	PromptVersion string                       `json:"prompt_version"`
	JSON          functionPromptDataOriginJSON `json:"-"`
}

// functionPromptDataOriginJSON contains the JSON metadata for the struct
// [FunctionPromptDataOrigin]
type functionPromptDataOriginJSON struct {
	ProjectID     apijson.Field
	PromptID      apijson.Field
	PromptVersion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FunctionPromptDataOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOriginJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataPrompt struct {
	Type    FunctionPromptDataPromptType `json:"type"`
	Content string                       `json:"content"`
	// This field can have the runtime type of [[]FunctionPromptDataPromptChatMessage].
	Messages interface{}                  `json:"messages,required"`
	Tools    string                       `json:"tools"`
	JSON     functionPromptDataPromptJSON `json:"-"`
	union    FunctionPromptDataPromptUnion
}

// functionPromptDataPromptJSON contains the JSON metadata for the struct
// [FunctionPromptDataPrompt]
type functionPromptDataPromptJSON struct {
	Type        apijson.Field
	Content     apijson.Field
	Messages    apijson.Field
	Tools       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r functionPromptDataPromptJSON) RawJSON() string {
	return r.raw
}

func (r *FunctionPromptDataPrompt) UnmarshalJSON(data []byte) (err error) {
	*r = FunctionPromptDataPrompt{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FunctionPromptDataPromptUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [FunctionPromptDataPromptCompletion],
// [FunctionPromptDataPromptChat], [FunctionPromptDataPromptNullableVariant].
func (r FunctionPromptDataPrompt) AsUnion() FunctionPromptDataPromptUnion {
	return r.union
}

// Union satisfied by [FunctionPromptDataPromptCompletion],
// [FunctionPromptDataPromptChat] or [FunctionPromptDataPromptNullableVariant].
type FunctionPromptDataPromptUnion interface {
	implementsFunctionPromptDataPrompt()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataPromptUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptCompletion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChat{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptNullableVariant{}),
		},
	)
}

type FunctionPromptDataPromptCompletion struct {
	Content string                                 `json:"content,required"`
	Type    FunctionPromptDataPromptCompletionType `json:"type,required"`
	JSON    functionPromptDataPromptCompletionJSON `json:"-"`
}

// functionPromptDataPromptCompletionJSON contains the JSON metadata for the struct
// [FunctionPromptDataPromptCompletion]
type functionPromptDataPromptCompletionJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptCompletion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptCompletionJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptCompletion) implementsFunctionPromptDataPrompt() {}

type FunctionPromptDataPromptCompletionType string

const (
	FunctionPromptDataPromptCompletionTypeCompletion FunctionPromptDataPromptCompletionType = "completion"
)

func (r FunctionPromptDataPromptCompletionType) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptCompletionTypeCompletion:
		return true
	}
	return false
}

type FunctionPromptDataPromptChat struct {
	Messages []FunctionPromptDataPromptChatMessage `json:"messages,required"`
	Type     FunctionPromptDataPromptChatType      `json:"type,required"`
	Tools    string                                `json:"tools"`
	JSON     functionPromptDataPromptChatJSON      `json:"-"`
}

// functionPromptDataPromptChatJSON contains the JSON metadata for the struct
// [FunctionPromptDataPromptChat]
type functionPromptDataPromptChatJSON struct {
	Messages    apijson.Field
	Type        apijson.Field
	Tools       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChat) implementsFunctionPromptDataPrompt() {}

type FunctionPromptDataPromptChatMessage struct {
	// This field can have the runtime type of [string],
	// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion].
	Content interface{}                              `json:"content,required"`
	Role    FunctionPromptDataPromptChatMessagesRole `json:"role,required"`
	Name    string                                   `json:"name"`
	// This field can have the runtime type of
	// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall].
	FunctionCall interface{} `json:"function_call,required"`
	// This field can have the runtime type of
	// [[]FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall].
	ToolCalls  interface{}                             `json:"tool_calls,required"`
	ToolCallID string                                  `json:"tool_call_id"`
	JSON       functionPromptDataPromptChatMessageJSON `json:"-"`
	union      FunctionPromptDataPromptChatMessagesUnion
}

// functionPromptDataPromptChatMessageJSON contains the JSON metadata for the
// struct [FunctionPromptDataPromptChatMessage]
type functionPromptDataPromptChatMessageJSON struct {
	Content      apijson.Field
	Role         apijson.Field
	Name         apijson.Field
	FunctionCall apijson.Field
	ToolCalls    apijson.Field
	ToolCallID   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r functionPromptDataPromptChatMessageJSON) RawJSON() string {
	return r.raw
}

func (r *FunctionPromptDataPromptChatMessage) UnmarshalJSON(data []byte) (err error) {
	*r = FunctionPromptDataPromptChatMessage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FunctionPromptDataPromptChatMessagesUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage0],
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1],
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2],
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage3],
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage4],
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage5].
func (r FunctionPromptDataPromptChatMessage) AsUnion() FunctionPromptDataPromptChatMessagesUnion {
	return r.union
}

// Union satisfied by
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage0],
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1],
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2],
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage3],
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage4] or
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage5].
type FunctionPromptDataPromptChatMessagesUnion interface {
	implementsFunctionPromptDataPromptChatMessage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataPromptChatMessagesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesPromptDataPromptMessage0{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesPromptDataPromptMessage3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesPromptDataPromptMessage4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesPromptDataPromptMessage5{}),
		},
	)
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage0 struct {
	Role    FunctionPromptDataPromptChatMessagesPromptDataPromptMessage0Role `json:"role,required"`
	Content string                                                           `json:"content"`
	Name    string                                                           `json:"name"`
	JSON    functionPromptDataPromptChatMessagesPromptDataPromptMessage0JSON `json:"-"`
}

// functionPromptDataPromptChatMessagesPromptDataPromptMessage0JSON contains the
// JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage0]
type functionPromptDataPromptChatMessagesPromptDataPromptMessage0JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesPromptDataPromptMessage0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesPromptDataPromptMessage0JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage0) implementsFunctionPromptDataPromptChatMessage() {
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage0Role string

const (
	FunctionPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem FunctionPromptDataPromptChatMessagesPromptDataPromptMessage0Role = "system"
)

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1 struct {
	Role    FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1Role         `json:"role,required"`
	Content FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion `json:"content"`
	Name    string                                                                   `json:"name"`
	JSON    functionPromptDataPromptChatMessagesPromptDataPromptMessage1JSON         `json:"-"`
}

// functionPromptDataPromptChatMessagesPromptDataPromptMessage1JSON contains the
// JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1]
type functionPromptDataPromptChatMessagesPromptDataPromptMessage1JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesPromptDataPromptMessage1JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1) implementsFunctionPromptDataPromptChatMessage() {
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1Role string

const (
	FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1Role = "user"
)

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray].
type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsFunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray{}),
		},
	)
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray []FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) ImplementsFunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion() {
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem struct {
	Text string                                                                       `json:"text"`
	Type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType `json:"type,required"`
	// This field can have the runtime type of
	// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL].
	ImageURL interface{}                                                                      `json:"image_url,required"`
	JSON     functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItemJSON `json:"-"`
	union    FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnionItem
}

// functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItemJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem]
type functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItemJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	ImageURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItemJSON) RawJSON() string {
	return r.raw
}

func (r *FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem) UnmarshalJSON(data []byte) (err error) {
	*r = FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnionItem]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList].
func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem) AsUnion() FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnionItem {
	return r.union
}

// Union satisfied by
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent]
// or
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList].
type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnionItem interface {
	implementsFunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList{}),
		},
	)
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType `json:"type,required"`
	Text string                                                                                                     `json:"text"`
	JSON functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON `json:"-"`
}

// functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent]
type functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON struct {
	Type        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsFunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem() {
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL `json:"image_url,required"`
	Type     FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType     `json:"type,required"`
	JSON     functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON     `json:"-"`
}

// functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList]
type functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON struct {
	ImageURL    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsFunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem() {
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    string                                                                                                            `json:"url,required"`
	Detail FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail `json:"detail"`
	JSON   functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON   `json:"-"`
}

// functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL]
type functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON struct {
	URL         apijson.Field
	Detail      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType string

const (
	FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText     FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "text"
	FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText, FunctionPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2 struct {
	Role         FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2Role         `json:"role,required"`
	Content      string                                                                   `json:"content,nullable"`
	FunctionCall FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall `json:"function_call"`
	Name         string                                                                   `json:"name"`
	ToolCalls    []FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall   `json:"tool_calls"`
	JSON         functionPromptDataPromptChatMessagesPromptDataPromptMessage2JSON         `json:"-"`
}

// functionPromptDataPromptChatMessagesPromptDataPromptMessage2JSON contains the
// JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2]
type functionPromptDataPromptChatMessagesPromptDataPromptMessage2JSON struct {
	Role         apijson.Field
	Content      apijson.Field
	FunctionCall apijson.Field
	Name         apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesPromptDataPromptMessage2JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2) implementsFunctionPromptDataPromptChatMessage() {
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2Role string

const (
	FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2Role = "assistant"
)

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments string                                                                       `json:"arguments,required"`
	Name      string                                                                       `json:"name,required"`
	JSON      functionPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCallJSON `json:"-"`
}

// functionPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCallJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall]
type functionPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCallJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall struct {
	ID       string                                                                        `json:"id,required"`
	Function FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction `json:"function,required"`
	Type     FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType     `json:"type,required"`
	JSON     functionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallJSON      `json:"-"`
}

// functionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall]
type functionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallJSON struct {
	ID          apijson.Field
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments string                                                                            `json:"arguments,required"`
	Name      string                                                                            `json:"name,required"`
	JSON      functionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunctionJSON `json:"-"`
}

// functionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunctionJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction]
type functionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunctionJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType string

const (
	FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage3 struct {
	Role       FunctionPromptDataPromptChatMessagesPromptDataPromptMessage3Role `json:"role,required"`
	Content    string                                                           `json:"content"`
	ToolCallID string                                                           `json:"tool_call_id"`
	JSON       functionPromptDataPromptChatMessagesPromptDataPromptMessage3JSON `json:"-"`
}

// functionPromptDataPromptChatMessagesPromptDataPromptMessage3JSON contains the
// JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage3]
type functionPromptDataPromptChatMessagesPromptDataPromptMessage3JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	ToolCallID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesPromptDataPromptMessage3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesPromptDataPromptMessage3JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage3) implementsFunctionPromptDataPromptChatMessage() {
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage3Role string

const (
	FunctionPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool FunctionPromptDataPromptChatMessagesPromptDataPromptMessage3Role = "tool"
)

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage4 struct {
	Name    string                                                           `json:"name,required"`
	Role    FunctionPromptDataPromptChatMessagesPromptDataPromptMessage4Role `json:"role,required"`
	Content string                                                           `json:"content"`
	JSON    functionPromptDataPromptChatMessagesPromptDataPromptMessage4JSON `json:"-"`
}

// functionPromptDataPromptChatMessagesPromptDataPromptMessage4JSON contains the
// JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage4]
type functionPromptDataPromptChatMessagesPromptDataPromptMessage4JSON struct {
	Name        apijson.Field
	Role        apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesPromptDataPromptMessage4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesPromptDataPromptMessage4JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage4) implementsFunctionPromptDataPromptChatMessage() {
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage4Role string

const (
	FunctionPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction FunctionPromptDataPromptChatMessagesPromptDataPromptMessage4Role = "function"
)

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage5 struct {
	Role    FunctionPromptDataPromptChatMessagesPromptDataPromptMessage5Role `json:"role,required"`
	Content string                                                           `json:"content,nullable"`
	JSON    functionPromptDataPromptChatMessagesPromptDataPromptMessage5JSON `json:"-"`
}

// functionPromptDataPromptChatMessagesPromptDataPromptMessage5JSON contains the
// JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesPromptDataPromptMessage5]
type functionPromptDataPromptChatMessagesPromptDataPromptMessage5JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesPromptDataPromptMessage5) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesPromptDataPromptMessage5JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage5) implementsFunctionPromptDataPromptChatMessage() {
}

type FunctionPromptDataPromptChatMessagesPromptDataPromptMessage5Role string

const (
	FunctionPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel FunctionPromptDataPromptChatMessagesPromptDataPromptMessage5Role = "model"
)

func (r FunctionPromptDataPromptChatMessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesRole string

const (
	FunctionPromptDataPromptChatMessagesRoleSystem    FunctionPromptDataPromptChatMessagesRole = "system"
	FunctionPromptDataPromptChatMessagesRoleUser      FunctionPromptDataPromptChatMessagesRole = "user"
	FunctionPromptDataPromptChatMessagesRoleAssistant FunctionPromptDataPromptChatMessagesRole = "assistant"
	FunctionPromptDataPromptChatMessagesRoleTool      FunctionPromptDataPromptChatMessagesRole = "tool"
	FunctionPromptDataPromptChatMessagesRoleFunction  FunctionPromptDataPromptChatMessagesRole = "function"
	FunctionPromptDataPromptChatMessagesRoleModel     FunctionPromptDataPromptChatMessagesRole = "model"
)

func (r FunctionPromptDataPromptChatMessagesRole) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesRoleSystem, FunctionPromptDataPromptChatMessagesRoleUser, FunctionPromptDataPromptChatMessagesRoleAssistant, FunctionPromptDataPromptChatMessagesRoleTool, FunctionPromptDataPromptChatMessagesRoleFunction, FunctionPromptDataPromptChatMessagesRoleModel:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatType string

const (
	FunctionPromptDataPromptChatTypeChat FunctionPromptDataPromptChatType = "chat"
)

func (r FunctionPromptDataPromptChatType) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatTypeChat:
		return true
	}
	return false
}

type FunctionPromptDataPromptNullableVariant struct {
	JSON functionPromptDataPromptNullableVariantJSON `json:"-"`
}

// functionPromptDataPromptNullableVariantJSON contains the JSON metadata for the
// struct [FunctionPromptDataPromptNullableVariant]
type functionPromptDataPromptNullableVariantJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptNullableVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptNullableVariantJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptNullableVariant) implementsFunctionPromptDataPrompt() {}

type FunctionPromptDataPromptType string

const (
	FunctionPromptDataPromptTypeCompletion FunctionPromptDataPromptType = "completion"
	FunctionPromptDataPromptTypeChat       FunctionPromptDataPromptType = "chat"
)

func (r FunctionPromptDataPromptType) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptTypeCompletion, FunctionPromptDataPromptTypeChat:
		return true
	}
	return false
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
	PromptData param.Field[FunctionNewParamsPromptData] `json:"prompt_data"`
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

// The prompt, model, and its parameters
type FunctionNewParamsPromptData struct {
	Options param.Field[FunctionNewParamsPromptDataOptions]     `json:"options"`
	Origin  param.Field[FunctionNewParamsPromptDataOrigin]      `json:"origin"`
	Prompt  param.Field[FunctionNewParamsPromptDataPromptUnion] `json:"prompt"`
}

func (r FunctionNewParamsPromptData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataOptions struct {
	Model    param.Field[string]                                        `json:"model"`
	Params   param.Field[FunctionNewParamsPromptDataOptionsParamsUnion] `json:"params"`
	Position param.Field[string]                                        `json:"position"`
}

func (r FunctionNewParamsPromptDataOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataOptionsParams struct {
	UseCache         param.Field[bool]        `json:"use_cache"`
	Temperature      param.Field[float64]     `json:"temperature"`
	TopP             param.Field[float64]     `json:"top_p"`
	MaxTokens        param.Field[float64]     `json:"max_tokens"`
	FrequencyPenalty param.Field[float64]     `json:"frequency_penalty"`
	PresencePenalty  param.Field[float64]     `json:"presence_penalty"`
	ResponseFormat   param.Field[interface{}] `json:"response_format,required"`
	ToolChoice       param.Field[interface{}] `json:"tool_choice,required"`
	FunctionCall     param.Field[interface{}] `json:"function_call,required"`
	N                param.Field[float64]     `json:"n"`
	Stop             param.Field[interface{}] `json:"stop,required"`
	TopK             param.Field[float64]     `json:"top_k"`
	StopSequences    param.Field[interface{}] `json:"stop_sequences,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64] `json:"max_tokens_to_sample"`
	MaxOutputTokens   param.Field[float64] `json:"maxOutputTokens"`
	TopP              param.Field[float64] `json:"topP"`
	TopK              param.Field[float64] `json:"topK"`
}

func (r FunctionNewParamsPromptDataOptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParams) implementsFunctionNewParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by [FunctionNewParamsPromptDataOptionsParamsOpenAIModelParams],
// [FunctionNewParamsPromptDataOptionsParamsAnthropicModelParams],
// [FunctionNewParamsPromptDataOptionsParamsGoogleModelParams],
// [FunctionNewParamsPromptDataOptionsParamsWindowAIModelParams],
// [FunctionNewParamsPromptDataOptionsParamsJsCompletionParams],
// [FunctionNewParamsPromptDataOptionsParams].
type FunctionNewParamsPromptDataOptionsParamsUnion interface {
	implementsFunctionNewParamsPromptDataOptionsParamsUnion()
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParams struct {
	FrequencyPenalty param.Field[float64]                                                                    `json:"frequency_penalty"`
	FunctionCall     param.Field[FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                                    `json:"max_tokens"`
	N                param.Field[float64]                                                                    `json:"n"`
	PresencePenalty  param.Field[float64]                                                                    `json:"presence_penalty"`
	ResponseFormat   param.Field[FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                                   `json:"stop"`
	Temperature      param.Field[float64]                                                                    `json:"temperature"`
	ToolChoice       param.Field[FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                                    `json:"top_p"`
	UseCache         param.Field[bool]                                                                       `json:"use_cache"`
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParams) implementsFunctionNewParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by
// [FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString],
// [FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString],
// [FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject].
type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsFunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString string

const (
	FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString = "auto"
)

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto:
		return true
	}
	return false
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) implementsFunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) implementsFunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat struct {
	Type param.Field[FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType] `json:"type,required"`
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType string

const (
	FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType = "json_object"
)

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Satisfied by
// [FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString],
// [FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString],
// [FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject].
type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsFunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString string

const (
	FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString = "auto"
)

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto:
		return true
	}
	return false
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) implementsFunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject struct {
	Function param.Field[FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction] `json:"function,required"`
	Type     param.Field[FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType]     `json:"type,required"`
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) implementsFunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType string

const (
	FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType = "function"
)

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataOptionsParamsAnthropicModelParams struct {
	MaxTokens   param.Field[float64] `json:"max_tokens,required"`
	Temperature param.Field[float64] `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64]  `json:"max_tokens_to_sample"`
	StopSequences     param.Field[[]string] `json:"stop_sequences"`
	TopK              param.Field[float64]  `json:"top_k"`
	TopP              param.Field[float64]  `json:"top_p"`
	UseCache          param.Field[bool]     `json:"use_cache"`
}

func (r FunctionNewParamsPromptDataOptionsParamsAnthropicModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsAnthropicModelParams) implementsFunctionNewParamsPromptDataOptionsParamsUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsGoogleModelParams struct {
	MaxOutputTokens param.Field[float64] `json:"maxOutputTokens"`
	Temperature     param.Field[float64] `json:"temperature"`
	TopK            param.Field[float64] `json:"topK"`
	TopP            param.Field[float64] `json:"topP"`
	UseCache        param.Field[bool]    `json:"use_cache"`
}

func (r FunctionNewParamsPromptDataOptionsParamsGoogleModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsGoogleModelParams) implementsFunctionNewParamsPromptDataOptionsParamsUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsWindowAIModelParams struct {
	Temperature param.Field[float64] `json:"temperature"`
	TopK        param.Field[float64] `json:"topK"`
	UseCache    param.Field[bool]    `json:"use_cache"`
}

func (r FunctionNewParamsPromptDataOptionsParamsWindowAIModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsWindowAIModelParams) implementsFunctionNewParamsPromptDataOptionsParamsUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsJsCompletionParams struct {
	UseCache param.Field[bool] `json:"use_cache"`
}

func (r FunctionNewParamsPromptDataOptionsParamsJsCompletionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsJsCompletionParams) implementsFunctionNewParamsPromptDataOptionsParamsUnion() {
}

type FunctionNewParamsPromptDataOrigin struct {
	ProjectID     param.Field[string] `json:"project_id"`
	PromptID      param.Field[string] `json:"prompt_id"`
	PromptVersion param.Field[string] `json:"prompt_version"`
}

func (r FunctionNewParamsPromptDataOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataPrompt struct {
	Type     param.Field[FunctionNewParamsPromptDataPromptType] `json:"type"`
	Content  param.Field[string]                                `json:"content"`
	Messages param.Field[interface{}]                           `json:"messages,required"`
	Tools    param.Field[string]                                `json:"tools"`
}

func (r FunctionNewParamsPromptDataPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPrompt) implementsFunctionNewParamsPromptDataPromptUnion() {}

// Satisfied by [FunctionNewParamsPromptDataPromptCompletion],
// [FunctionNewParamsPromptDataPromptChat],
// [FunctionNewParamsPromptDataPromptNullableVariant],
// [FunctionNewParamsPromptDataPrompt].
type FunctionNewParamsPromptDataPromptUnion interface {
	implementsFunctionNewParamsPromptDataPromptUnion()
}

type FunctionNewParamsPromptDataPromptCompletion struct {
	Content param.Field[string]                                          `json:"content,required"`
	Type    param.Field[FunctionNewParamsPromptDataPromptCompletionType] `json:"type,required"`
}

func (r FunctionNewParamsPromptDataPromptCompletion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptCompletion) implementsFunctionNewParamsPromptDataPromptUnion() {
}

type FunctionNewParamsPromptDataPromptCompletionType string

const (
	FunctionNewParamsPromptDataPromptCompletionTypeCompletion FunctionNewParamsPromptDataPromptCompletionType = "completion"
)

func (r FunctionNewParamsPromptDataPromptCompletionType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptCompletionTypeCompletion:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChat struct {
	Messages param.Field[[]FunctionNewParamsPromptDataPromptChatMessageUnion] `json:"messages,required"`
	Type     param.Field[FunctionNewParamsPromptDataPromptChatType]           `json:"type,required"`
	Tools    param.Field[string]                                              `json:"tools"`
}

func (r FunctionNewParamsPromptDataPromptChat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChat) implementsFunctionNewParamsPromptDataPromptUnion() {}

type FunctionNewParamsPromptDataPromptChatMessage struct {
	Content      param.Field[interface{}]                                       `json:"content,required"`
	Role         param.Field[FunctionNewParamsPromptDataPromptChatMessagesRole] `json:"role,required"`
	Name         param.Field[string]                                            `json:"name"`
	FunctionCall param.Field[interface{}]                                       `json:"function_call,required"`
	ToolCalls    param.Field[interface{}]                                       `json:"tool_calls,required"`
	ToolCallID   param.Field[string]                                            `json:"tool_call_id"`
}

func (r FunctionNewParamsPromptDataPromptChatMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessage) implementsFunctionNewParamsPromptDataPromptChatMessageUnion() {
}

// Satisfied by
// [FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0],
// [FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1],
// [FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2],
// [FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3],
// [FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4],
// [FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5],
// [FunctionNewParamsPromptDataPromptChatMessage].
type FunctionNewParamsPromptDataPromptChatMessageUnion interface {
	implementsFunctionNewParamsPromptDataPromptChatMessageUnion()
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0 struct {
	Role    param.Field[FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role] `json:"role,required"`
	Content param.Field[string]                                                                    `json:"content"`
	Name    param.Field[string]                                                                    `json:"name"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0) implementsFunctionNewParamsPromptDataPromptChatMessageUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role string

const (
	FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role = "system"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1 struct {
	Role    param.Field[FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role]         `json:"role,required"`
	Content param.Field[FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion] `json:"content"`
	Name    param.Field[string]                                                                            `json:"name"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1) implementsFunctionNewParamsPromptDataPromptChatMessageUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role string

const (
	FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role = "user"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray].
type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsFunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion()
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray []FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) ImplementsFunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray struct {
	Text     param.Field[string]                                                                                `json:"text"`
	Type     param.Field[FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                                           `json:"image_url,required"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) implementsFunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

// Satisfied by
// [FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList],
// [FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray].
type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion interface {
	implementsFunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion()
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type param.Field[FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType] `json:"type,required"`
	Text param.Field[string]                                                                                                              `json:"text"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsFunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL param.Field[FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL] `json:"image_url,required"`
	Type     param.Field[FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType]     `json:"type,required"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsFunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    param.Field[string]                                                                                                                     `json:"url,required"`
	Detail param.Field[FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail] `json:"detail"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType string

const (
	FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText     FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "text"
	FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText, FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2 struct {
	Role         param.Field[FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role]         `json:"role,required"`
	Content      param.Field[string]                                                                            `json:"content"`
	FunctionCall param.Field[FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                                            `json:"name"`
	ToolCalls    param.Field[[]FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall]   `json:"tool_calls"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2) implementsFunctionNewParamsPromptDataPromptChatMessageUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role string

const (
	FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role = "assistant"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall struct {
	ID       param.Field[string]                                                                                 `json:"id,required"`
	Function param.Field[FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction] `json:"function,required"`
	Type     param.Field[FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType]     `json:"type,required"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType string

const (
	FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3 struct {
	Role       param.Field[FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role] `json:"role,required"`
	Content    param.Field[string]                                                                    `json:"content"`
	ToolCallID param.Field[string]                                                                    `json:"tool_call_id"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3) implementsFunctionNewParamsPromptDataPromptChatMessageUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role string

const (
	FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role = "tool"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4 struct {
	Name    param.Field[string]                                                                    `json:"name,required"`
	Role    param.Field[FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role] `json:"role,required"`
	Content param.Field[string]                                                                    `json:"content"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4) implementsFunctionNewParamsPromptDataPromptChatMessageUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role string

const (
	FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role = "function"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5 struct {
	Role    param.Field[FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role] `json:"role,required"`
	Content param.Field[string]                                                                    `json:"content"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5) implementsFunctionNewParamsPromptDataPromptChatMessageUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role string

const (
	FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role = "model"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesRole string

const (
	FunctionNewParamsPromptDataPromptChatMessagesRoleSystem    FunctionNewParamsPromptDataPromptChatMessagesRole = "system"
	FunctionNewParamsPromptDataPromptChatMessagesRoleUser      FunctionNewParamsPromptDataPromptChatMessagesRole = "user"
	FunctionNewParamsPromptDataPromptChatMessagesRoleAssistant FunctionNewParamsPromptDataPromptChatMessagesRole = "assistant"
	FunctionNewParamsPromptDataPromptChatMessagesRoleTool      FunctionNewParamsPromptDataPromptChatMessagesRole = "tool"
	FunctionNewParamsPromptDataPromptChatMessagesRoleFunction  FunctionNewParamsPromptDataPromptChatMessagesRole = "function"
	FunctionNewParamsPromptDataPromptChatMessagesRoleModel     FunctionNewParamsPromptDataPromptChatMessagesRole = "model"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesRole) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesRoleSystem, FunctionNewParamsPromptDataPromptChatMessagesRoleUser, FunctionNewParamsPromptDataPromptChatMessagesRoleAssistant, FunctionNewParamsPromptDataPromptChatMessagesRoleTool, FunctionNewParamsPromptDataPromptChatMessagesRoleFunction, FunctionNewParamsPromptDataPromptChatMessagesRoleModel:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatType string

const (
	FunctionNewParamsPromptDataPromptChatTypeChat FunctionNewParamsPromptDataPromptChatType = "chat"
)

func (r FunctionNewParamsPromptDataPromptChatType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatTypeChat:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptNullableVariant struct {
}

func (r FunctionNewParamsPromptDataPromptNullableVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptNullableVariant) implementsFunctionNewParamsPromptDataPromptUnion() {
}

type FunctionNewParamsPromptDataPromptType string

const (
	FunctionNewParamsPromptDataPromptTypeCompletion FunctionNewParamsPromptDataPromptType = "completion"
	FunctionNewParamsPromptDataPromptTypeChat       FunctionNewParamsPromptDataPromptType = "chat"
)

func (r FunctionNewParamsPromptDataPromptType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptTypeCompletion, FunctionNewParamsPromptDataPromptTypeChat:
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
	PromptData param.Field[FunctionUpdateParamsPromptData] `json:"prompt_data"`
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

// The prompt, model, and its parameters
type FunctionUpdateParamsPromptData struct {
	Options param.Field[FunctionUpdateParamsPromptDataOptions]     `json:"options"`
	Origin  param.Field[FunctionUpdateParamsPromptDataOrigin]      `json:"origin"`
	Prompt  param.Field[FunctionUpdateParamsPromptDataPromptUnion] `json:"prompt"`
}

func (r FunctionUpdateParamsPromptData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataOptions struct {
	Model    param.Field[string]                                           `json:"model"`
	Params   param.Field[FunctionUpdateParamsPromptDataOptionsParamsUnion] `json:"params"`
	Position param.Field[string]                                           `json:"position"`
}

func (r FunctionUpdateParamsPromptDataOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataOptionsParams struct {
	UseCache         param.Field[bool]        `json:"use_cache"`
	Temperature      param.Field[float64]     `json:"temperature"`
	TopP             param.Field[float64]     `json:"top_p"`
	MaxTokens        param.Field[float64]     `json:"max_tokens"`
	FrequencyPenalty param.Field[float64]     `json:"frequency_penalty"`
	PresencePenalty  param.Field[float64]     `json:"presence_penalty"`
	ResponseFormat   param.Field[interface{}] `json:"response_format,required"`
	ToolChoice       param.Field[interface{}] `json:"tool_choice,required"`
	FunctionCall     param.Field[interface{}] `json:"function_call,required"`
	N                param.Field[float64]     `json:"n"`
	Stop             param.Field[interface{}] `json:"stop,required"`
	TopK             param.Field[float64]     `json:"top_k"`
	StopSequences    param.Field[interface{}] `json:"stop_sequences,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64] `json:"max_tokens_to_sample"`
	MaxOutputTokens   param.Field[float64] `json:"maxOutputTokens"`
	TopP              param.Field[float64] `json:"topP"`
	TopK              param.Field[float64] `json:"topK"`
}

func (r FunctionUpdateParamsPromptDataOptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParams) implementsFunctionUpdateParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by [FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParams],
// [FunctionUpdateParamsPromptDataOptionsParamsAnthropicModelParams],
// [FunctionUpdateParamsPromptDataOptionsParamsGoogleModelParams],
// [FunctionUpdateParamsPromptDataOptionsParamsWindowAIModelParams],
// [FunctionUpdateParamsPromptDataOptionsParamsJsCompletionParams],
// [FunctionUpdateParamsPromptDataOptionsParams].
type FunctionUpdateParamsPromptDataOptionsParamsUnion interface {
	implementsFunctionUpdateParamsPromptDataOptionsParamsUnion()
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParams struct {
	FrequencyPenalty param.Field[float64]                                                                       `json:"frequency_penalty"`
	FunctionCall     param.Field[FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                                       `json:"max_tokens"`
	N                param.Field[float64]                                                                       `json:"n"`
	PresencePenalty  param.Field[float64]                                                                       `json:"presence_penalty"`
	ResponseFormat   param.Field[FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                                      `json:"stop"`
	Temperature      param.Field[float64]                                                                       `json:"temperature"`
	ToolChoice       param.Field[FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                                       `json:"top_p"`
	UseCache         param.Field[bool]                                                                          `json:"use_cache"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParams) implementsFunctionUpdateParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by
// [FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString],
// [FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString],
// [FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject].
type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsFunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString string

const (
	FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString = "auto"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto:
		return true
	}
	return false
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) implementsFunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) implementsFunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat struct {
	Type param.Field[FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType] `json:"type,required"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType string

const (
	FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType = "json_object"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Satisfied by
// [FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString],
// [FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString],
// [FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject].
type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsFunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString string

const (
	FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString = "auto"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto:
		return true
	}
	return false
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) implementsFunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject struct {
	Function param.Field[FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction] `json:"function,required"`
	Type     param.Field[FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType]     `json:"type,required"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) implementsFunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType string

const (
	FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType = "function"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataOptionsParamsAnthropicModelParams struct {
	MaxTokens   param.Field[float64] `json:"max_tokens,required"`
	Temperature param.Field[float64] `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64]  `json:"max_tokens_to_sample"`
	StopSequences     param.Field[[]string] `json:"stop_sequences"`
	TopK              param.Field[float64]  `json:"top_k"`
	TopP              param.Field[float64]  `json:"top_p"`
	UseCache          param.Field[bool]     `json:"use_cache"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsAnthropicModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsAnthropicModelParams) implementsFunctionUpdateParamsPromptDataOptionsParamsUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsGoogleModelParams struct {
	MaxOutputTokens param.Field[float64] `json:"maxOutputTokens"`
	Temperature     param.Field[float64] `json:"temperature"`
	TopK            param.Field[float64] `json:"topK"`
	TopP            param.Field[float64] `json:"topP"`
	UseCache        param.Field[bool]    `json:"use_cache"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsGoogleModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsGoogleModelParams) implementsFunctionUpdateParamsPromptDataOptionsParamsUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsWindowAIModelParams struct {
	Temperature param.Field[float64] `json:"temperature"`
	TopK        param.Field[float64] `json:"topK"`
	UseCache    param.Field[bool]    `json:"use_cache"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsWindowAIModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsWindowAIModelParams) implementsFunctionUpdateParamsPromptDataOptionsParamsUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsJsCompletionParams struct {
	UseCache param.Field[bool] `json:"use_cache"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsJsCompletionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsJsCompletionParams) implementsFunctionUpdateParamsPromptDataOptionsParamsUnion() {
}

type FunctionUpdateParamsPromptDataOrigin struct {
	ProjectID     param.Field[string] `json:"project_id"`
	PromptID      param.Field[string] `json:"prompt_id"`
	PromptVersion param.Field[string] `json:"prompt_version"`
}

func (r FunctionUpdateParamsPromptDataOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataPrompt struct {
	Type     param.Field[FunctionUpdateParamsPromptDataPromptType] `json:"type"`
	Content  param.Field[string]                                   `json:"content"`
	Messages param.Field[interface{}]                              `json:"messages,required"`
	Tools    param.Field[string]                                   `json:"tools"`
}

func (r FunctionUpdateParamsPromptDataPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPrompt) implementsFunctionUpdateParamsPromptDataPromptUnion() {}

// Satisfied by [FunctionUpdateParamsPromptDataPromptCompletion],
// [FunctionUpdateParamsPromptDataPromptChat],
// [FunctionUpdateParamsPromptDataPromptNullableVariant],
// [FunctionUpdateParamsPromptDataPrompt].
type FunctionUpdateParamsPromptDataPromptUnion interface {
	implementsFunctionUpdateParamsPromptDataPromptUnion()
}

type FunctionUpdateParamsPromptDataPromptCompletion struct {
	Content param.Field[string]                                             `json:"content,required"`
	Type    param.Field[FunctionUpdateParamsPromptDataPromptCompletionType] `json:"type,required"`
}

func (r FunctionUpdateParamsPromptDataPromptCompletion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptCompletion) implementsFunctionUpdateParamsPromptDataPromptUnion() {
}

type FunctionUpdateParamsPromptDataPromptCompletionType string

const (
	FunctionUpdateParamsPromptDataPromptCompletionTypeCompletion FunctionUpdateParamsPromptDataPromptCompletionType = "completion"
)

func (r FunctionUpdateParamsPromptDataPromptCompletionType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptCompletionTypeCompletion:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChat struct {
	Messages param.Field[[]FunctionUpdateParamsPromptDataPromptChatMessageUnion] `json:"messages,required"`
	Type     param.Field[FunctionUpdateParamsPromptDataPromptChatType]           `json:"type,required"`
	Tools    param.Field[string]                                                 `json:"tools"`
}

func (r FunctionUpdateParamsPromptDataPromptChat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChat) implementsFunctionUpdateParamsPromptDataPromptUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessage struct {
	Content      param.Field[interface{}]                                          `json:"content,required"`
	Role         param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesRole] `json:"role,required"`
	Name         param.Field[string]                                               `json:"name"`
	FunctionCall param.Field[interface{}]                                          `json:"function_call,required"`
	ToolCalls    param.Field[interface{}]                                          `json:"tool_calls,required"`
	ToolCallID   param.Field[string]                                               `json:"tool_call_id"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessage) implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion() {
}

// Satisfied by
// [FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0],
// [FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1],
// [FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2],
// [FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3],
// [FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4],
// [FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5],
// [FunctionUpdateParamsPromptDataPromptChatMessage].
type FunctionUpdateParamsPromptDataPromptChatMessageUnion interface {
	implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion()
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0 struct {
	Role    param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role] `json:"role,required"`
	Content param.Field[string]                                                                       `json:"content"`
	Name    param.Field[string]                                                                       `json:"name"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0) implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role = "system"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1 struct {
	Role    param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role]         `json:"role,required"`
	Content param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion] `json:"content"`
	Name    param.Field[string]                                                                               `json:"name"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1) implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role = "user"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray].
type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsFunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion()
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray []FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) ImplementsFunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray struct {
	Text     param.Field[string]                                                                                   `json:"text"`
	Type     param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                                              `json:"image_url,required"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) implementsFunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

// Satisfied by
// [FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList],
// [FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray].
type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion interface {
	implementsFunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion()
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType] `json:"type,required"`
	Text param.Field[string]                                                                                                                 `json:"text"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsFunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL] `json:"image_url,required"`
	Type     param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType]     `json:"type,required"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsFunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    param.Field[string]                                                                                                                        `json:"url,required"`
	Detail param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail] `json:"detail"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText     FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "text"
	FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText, FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2 struct {
	Role         param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role]         `json:"role,required"`
	Content      param.Field[string]                                                                               `json:"content"`
	FunctionCall param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                                               `json:"name"`
	ToolCalls    param.Field[[]FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall]   `json:"tool_calls"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2) implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role = "assistant"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall struct {
	ID       param.Field[string]                                                                                    `json:"id,required"`
	Function param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction] `json:"function,required"`
	Type     param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType]     `json:"type,required"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3 struct {
	Role       param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role] `json:"role,required"`
	Content    param.Field[string]                                                                       `json:"content"`
	ToolCallID param.Field[string]                                                                       `json:"tool_call_id"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3) implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role = "tool"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4 struct {
	Name    param.Field[string]                                                                       `json:"name,required"`
	Role    param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role] `json:"role,required"`
	Content param.Field[string]                                                                       `json:"content"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4) implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role = "function"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5 struct {
	Role    param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role] `json:"role,required"`
	Content param.Field[string]                                                                       `json:"content"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5) implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role = "model"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesRole string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesRoleSystem    FunctionUpdateParamsPromptDataPromptChatMessagesRole = "system"
	FunctionUpdateParamsPromptDataPromptChatMessagesRoleUser      FunctionUpdateParamsPromptDataPromptChatMessagesRole = "user"
	FunctionUpdateParamsPromptDataPromptChatMessagesRoleAssistant FunctionUpdateParamsPromptDataPromptChatMessagesRole = "assistant"
	FunctionUpdateParamsPromptDataPromptChatMessagesRoleTool      FunctionUpdateParamsPromptDataPromptChatMessagesRole = "tool"
	FunctionUpdateParamsPromptDataPromptChatMessagesRoleFunction  FunctionUpdateParamsPromptDataPromptChatMessagesRole = "function"
	FunctionUpdateParamsPromptDataPromptChatMessagesRoleModel     FunctionUpdateParamsPromptDataPromptChatMessagesRole = "model"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesRole) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesRoleSystem, FunctionUpdateParamsPromptDataPromptChatMessagesRoleUser, FunctionUpdateParamsPromptDataPromptChatMessagesRoleAssistant, FunctionUpdateParamsPromptDataPromptChatMessagesRoleTool, FunctionUpdateParamsPromptDataPromptChatMessagesRoleFunction, FunctionUpdateParamsPromptDataPromptChatMessagesRoleModel:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatType string

const (
	FunctionUpdateParamsPromptDataPromptChatTypeChat FunctionUpdateParamsPromptDataPromptChatType = "chat"
)

func (r FunctionUpdateParamsPromptDataPromptChatType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatTypeChat:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptNullableVariant struct {
}

func (r FunctionUpdateParamsPromptDataPromptNullableVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptNullableVariant) implementsFunctionUpdateParamsPromptDataPromptUnion() {
}

type FunctionUpdateParamsPromptDataPromptType string

const (
	FunctionUpdateParamsPromptDataPromptTypeCompletion FunctionUpdateParamsPromptDataPromptType = "completion"
	FunctionUpdateParamsPromptDataPromptTypeChat       FunctionUpdateParamsPromptDataPromptType = "chat"
)

func (r FunctionUpdateParamsPromptDataPromptType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptTypeCompletion, FunctionUpdateParamsPromptDataPromptTypeChat:
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
	PromptData param.Field[FunctionReplaceParamsPromptData] `json:"prompt_data"`
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

// The prompt, model, and its parameters
type FunctionReplaceParamsPromptData struct {
	Options param.Field[FunctionReplaceParamsPromptDataOptions]     `json:"options"`
	Origin  param.Field[FunctionReplaceParamsPromptDataOrigin]      `json:"origin"`
	Prompt  param.Field[FunctionReplaceParamsPromptDataPromptUnion] `json:"prompt"`
}

func (r FunctionReplaceParamsPromptData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataOptions struct {
	Model    param.Field[string]                                            `json:"model"`
	Params   param.Field[FunctionReplaceParamsPromptDataOptionsParamsUnion] `json:"params"`
	Position param.Field[string]                                            `json:"position"`
}

func (r FunctionReplaceParamsPromptDataOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataOptionsParams struct {
	UseCache         param.Field[bool]        `json:"use_cache"`
	Temperature      param.Field[float64]     `json:"temperature"`
	TopP             param.Field[float64]     `json:"top_p"`
	MaxTokens        param.Field[float64]     `json:"max_tokens"`
	FrequencyPenalty param.Field[float64]     `json:"frequency_penalty"`
	PresencePenalty  param.Field[float64]     `json:"presence_penalty"`
	ResponseFormat   param.Field[interface{}] `json:"response_format,required"`
	ToolChoice       param.Field[interface{}] `json:"tool_choice,required"`
	FunctionCall     param.Field[interface{}] `json:"function_call,required"`
	N                param.Field[float64]     `json:"n"`
	Stop             param.Field[interface{}] `json:"stop,required"`
	TopK             param.Field[float64]     `json:"top_k"`
	StopSequences    param.Field[interface{}] `json:"stop_sequences,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64] `json:"max_tokens_to_sample"`
	MaxOutputTokens   param.Field[float64] `json:"maxOutputTokens"`
	TopP              param.Field[float64] `json:"topP"`
	TopK              param.Field[float64] `json:"topK"`
}

func (r FunctionReplaceParamsPromptDataOptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParams) implementsFunctionReplaceParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by [FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParams],
// [FunctionReplaceParamsPromptDataOptionsParamsAnthropicModelParams],
// [FunctionReplaceParamsPromptDataOptionsParamsGoogleModelParams],
// [FunctionReplaceParamsPromptDataOptionsParamsWindowAIModelParams],
// [FunctionReplaceParamsPromptDataOptionsParamsJsCompletionParams],
// [FunctionReplaceParamsPromptDataOptionsParams].
type FunctionReplaceParamsPromptDataOptionsParamsUnion interface {
	implementsFunctionReplaceParamsPromptDataOptionsParamsUnion()
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParams struct {
	FrequencyPenalty param.Field[float64]                                                                        `json:"frequency_penalty"`
	FunctionCall     param.Field[FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                                        `json:"max_tokens"`
	N                param.Field[float64]                                                                        `json:"n"`
	PresencePenalty  param.Field[float64]                                                                        `json:"presence_penalty"`
	ResponseFormat   param.Field[FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                                       `json:"stop"`
	Temperature      param.Field[float64]                                                                        `json:"temperature"`
	ToolChoice       param.Field[FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                                        `json:"top_p"`
	UseCache         param.Field[bool]                                                                           `json:"use_cache"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParams) implementsFunctionReplaceParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by
// [FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString],
// [FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString],
// [FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject].
type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsFunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString string

const (
	FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString = "auto"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto:
		return true
	}
	return false
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) implementsFunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) implementsFunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat struct {
	Type param.Field[FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType] `json:"type,required"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType string

const (
	FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType = "json_object"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Satisfied by
// [FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString],
// [FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString],
// [FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject].
type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsFunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString string

const (
	FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString = "auto"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto:
		return true
	}
	return false
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) implementsFunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject struct {
	Function param.Field[FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction] `json:"function,required"`
	Type     param.Field[FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType]     `json:"type,required"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) implementsFunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType string

const (
	FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType = "function"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataOptionsParamsAnthropicModelParams struct {
	MaxTokens   param.Field[float64] `json:"max_tokens,required"`
	Temperature param.Field[float64] `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64]  `json:"max_tokens_to_sample"`
	StopSequences     param.Field[[]string] `json:"stop_sequences"`
	TopK              param.Field[float64]  `json:"top_k"`
	TopP              param.Field[float64]  `json:"top_p"`
	UseCache          param.Field[bool]     `json:"use_cache"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsAnthropicModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsAnthropicModelParams) implementsFunctionReplaceParamsPromptDataOptionsParamsUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsGoogleModelParams struct {
	MaxOutputTokens param.Field[float64] `json:"maxOutputTokens"`
	Temperature     param.Field[float64] `json:"temperature"`
	TopK            param.Field[float64] `json:"topK"`
	TopP            param.Field[float64] `json:"topP"`
	UseCache        param.Field[bool]    `json:"use_cache"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsGoogleModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsGoogleModelParams) implementsFunctionReplaceParamsPromptDataOptionsParamsUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsWindowAIModelParams struct {
	Temperature param.Field[float64] `json:"temperature"`
	TopK        param.Field[float64] `json:"topK"`
	UseCache    param.Field[bool]    `json:"use_cache"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsWindowAIModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsWindowAIModelParams) implementsFunctionReplaceParamsPromptDataOptionsParamsUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsJsCompletionParams struct {
	UseCache param.Field[bool] `json:"use_cache"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsJsCompletionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsJsCompletionParams) implementsFunctionReplaceParamsPromptDataOptionsParamsUnion() {
}

type FunctionReplaceParamsPromptDataOrigin struct {
	ProjectID     param.Field[string] `json:"project_id"`
	PromptID      param.Field[string] `json:"prompt_id"`
	PromptVersion param.Field[string] `json:"prompt_version"`
}

func (r FunctionReplaceParamsPromptDataOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataPrompt struct {
	Type     param.Field[FunctionReplaceParamsPromptDataPromptType] `json:"type"`
	Content  param.Field[string]                                    `json:"content"`
	Messages param.Field[interface{}]                               `json:"messages,required"`
	Tools    param.Field[string]                                    `json:"tools"`
}

func (r FunctionReplaceParamsPromptDataPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPrompt) implementsFunctionReplaceParamsPromptDataPromptUnion() {
}

// Satisfied by [FunctionReplaceParamsPromptDataPromptCompletion],
// [FunctionReplaceParamsPromptDataPromptChat],
// [FunctionReplaceParamsPromptDataPromptNullableVariant],
// [FunctionReplaceParamsPromptDataPrompt].
type FunctionReplaceParamsPromptDataPromptUnion interface {
	implementsFunctionReplaceParamsPromptDataPromptUnion()
}

type FunctionReplaceParamsPromptDataPromptCompletion struct {
	Content param.Field[string]                                              `json:"content,required"`
	Type    param.Field[FunctionReplaceParamsPromptDataPromptCompletionType] `json:"type,required"`
}

func (r FunctionReplaceParamsPromptDataPromptCompletion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptCompletion) implementsFunctionReplaceParamsPromptDataPromptUnion() {
}

type FunctionReplaceParamsPromptDataPromptCompletionType string

const (
	FunctionReplaceParamsPromptDataPromptCompletionTypeCompletion FunctionReplaceParamsPromptDataPromptCompletionType = "completion"
)

func (r FunctionReplaceParamsPromptDataPromptCompletionType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptCompletionTypeCompletion:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChat struct {
	Messages param.Field[[]FunctionReplaceParamsPromptDataPromptChatMessageUnion] `json:"messages,required"`
	Type     param.Field[FunctionReplaceParamsPromptDataPromptChatType]           `json:"type,required"`
	Tools    param.Field[string]                                                  `json:"tools"`
}

func (r FunctionReplaceParamsPromptDataPromptChat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChat) implementsFunctionReplaceParamsPromptDataPromptUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessage struct {
	Content      param.Field[interface{}]                                           `json:"content,required"`
	Role         param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesRole] `json:"role,required"`
	Name         param.Field[string]                                                `json:"name"`
	FunctionCall param.Field[interface{}]                                           `json:"function_call,required"`
	ToolCalls    param.Field[interface{}]                                           `json:"tool_calls,required"`
	ToolCallID   param.Field[string]                                                `json:"tool_call_id"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessage) implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion() {
}

// Satisfied by
// [FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0],
// [FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1],
// [FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2],
// [FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3],
// [FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4],
// [FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5],
// [FunctionReplaceParamsPromptDataPromptChatMessage].
type FunctionReplaceParamsPromptDataPromptChatMessageUnion interface {
	implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion()
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0 struct {
	Role    param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role] `json:"role,required"`
	Content param.Field[string]                                                                        `json:"content"`
	Name    param.Field[string]                                                                        `json:"name"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0) implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role = "system"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1 struct {
	Role    param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role]         `json:"role,required"`
	Content param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion] `json:"content"`
	Name    param.Field[string]                                                                                `json:"name"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1) implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role = "user"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray].
type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsFunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion()
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray []FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) ImplementsFunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray struct {
	Text     param.Field[string]                                                                                    `json:"text"`
	Type     param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                                               `json:"image_url,required"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) implementsFunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

// Satisfied by
// [FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList],
// [FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray].
type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion interface {
	implementsFunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion()
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType] `json:"type,required"`
	Text param.Field[string]                                                                                                                  `json:"text"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsFunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL] `json:"image_url,required"`
	Type     param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType]     `json:"type,required"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsFunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    param.Field[string]                                                                                                                         `json:"url,required"`
	Detail param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail] `json:"detail"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText     FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "text"
	FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText, FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2 struct {
	Role         param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role]         `json:"role,required"`
	Content      param.Field[string]                                                                                `json:"content"`
	FunctionCall param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                                                `json:"name"`
	ToolCalls    param.Field[[]FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall]   `json:"tool_calls"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2) implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role = "assistant"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall struct {
	ID       param.Field[string]                                                                                     `json:"id,required"`
	Function param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction] `json:"function,required"`
	Type     param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType]     `json:"type,required"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3 struct {
	Role       param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role] `json:"role,required"`
	Content    param.Field[string]                                                                        `json:"content"`
	ToolCallID param.Field[string]                                                                        `json:"tool_call_id"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3) implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role = "tool"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4 struct {
	Name    param.Field[string]                                                                        `json:"name,required"`
	Role    param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role] `json:"role,required"`
	Content param.Field[string]                                                                        `json:"content"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4) implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role = "function"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5 struct {
	Role    param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role] `json:"role,required"`
	Content param.Field[string]                                                                        `json:"content"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5) implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role = "model"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesRole string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesRoleSystem    FunctionReplaceParamsPromptDataPromptChatMessagesRole = "system"
	FunctionReplaceParamsPromptDataPromptChatMessagesRoleUser      FunctionReplaceParamsPromptDataPromptChatMessagesRole = "user"
	FunctionReplaceParamsPromptDataPromptChatMessagesRoleAssistant FunctionReplaceParamsPromptDataPromptChatMessagesRole = "assistant"
	FunctionReplaceParamsPromptDataPromptChatMessagesRoleTool      FunctionReplaceParamsPromptDataPromptChatMessagesRole = "tool"
	FunctionReplaceParamsPromptDataPromptChatMessagesRoleFunction  FunctionReplaceParamsPromptDataPromptChatMessagesRole = "function"
	FunctionReplaceParamsPromptDataPromptChatMessagesRoleModel     FunctionReplaceParamsPromptDataPromptChatMessagesRole = "model"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesRole) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesRoleSystem, FunctionReplaceParamsPromptDataPromptChatMessagesRoleUser, FunctionReplaceParamsPromptDataPromptChatMessagesRoleAssistant, FunctionReplaceParamsPromptDataPromptChatMessagesRoleTool, FunctionReplaceParamsPromptDataPromptChatMessagesRoleFunction, FunctionReplaceParamsPromptDataPromptChatMessagesRoleModel:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatType string

const (
	FunctionReplaceParamsPromptDataPromptChatTypeChat FunctionReplaceParamsPromptDataPromptChatType = "chat"
)

func (r FunctionReplaceParamsPromptDataPromptChatType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatTypeChat:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptNullableVariant struct {
}

func (r FunctionReplaceParamsPromptDataPromptNullableVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptNullableVariant) implementsFunctionReplaceParamsPromptDataPromptUnion() {
}

type FunctionReplaceParamsPromptDataPromptType string

const (
	FunctionReplaceParamsPromptDataPromptTypeCompletion FunctionReplaceParamsPromptDataPromptType = "completion"
	FunctionReplaceParamsPromptDataPromptTypeChat       FunctionReplaceParamsPromptDataPromptType = "chat"
)

func (r FunctionReplaceParamsPromptDataPromptType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptTypeCompletion, FunctionReplaceParamsPromptDataPromptTypeChat:
		return true
	}
	return false
}
