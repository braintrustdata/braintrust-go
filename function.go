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
	Model    string                               `json:"model"`
	Params   FunctionPromptDataOptionsParamsUnion `json:"params"`
	Position string                               `json:"position"`
	JSON     functionPromptDataOptionsJSON        `json:"-"`
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

// Union satisfied by [FunctionPromptDataOptionsParamsOpenAIModelParams],
// [FunctionPromptDataOptionsParamsAnthropicModelParams],
// [FunctionPromptDataOptionsParamsGoogleModelParams],
// [FunctionPromptDataOptionsParamsWindowAIModelParams] or
// [FunctionPromptDataOptionsParamsJsCompletionParams].
type FunctionPromptDataOptionsParamsUnion interface {
	implementsFunctionPromptDataOptionsParamsUnion()
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

func (r FunctionPromptDataOptionsParamsOpenAIModelParams) implementsFunctionPromptDataOptionsParamsUnion() {
}

// Union satisfied by
// [FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto],
// [FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone] or
// [FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction].
type FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsFunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction{}),
		},
	)
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto string

const (
	FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto = "auto"
)

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto:
		return true
	}
	return false
}

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) implementsFunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone string

const (
	FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone = "none"
)

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone:
		return true
	}
	return false
}

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) implementsFunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction struct {
	Name string                                                                   `json:"name,required"`
	JSON functionPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionJSON `json:"-"`
}

// functionPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction]
type functionPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) implementsFunctionPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
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
// [FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto],
// [FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone] or
// [FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction].
type FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsFunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction{}),
		},
	)
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto string

const (
	FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto = "auto"
)

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto:
		return true
	}
	return false
}

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) implementsFunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone string

const (
	FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone = "none"
)

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone:
		return true
	}
	return false
}

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) implementsFunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction struct {
	Function FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction `json:"function,required"`
	Type     FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType     `json:"type,required"`
	JSON     functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionJSON     `json:"-"`
}

// functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionJSON contains
// the JSON metadata for the struct
// [FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction]
type functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionJSON struct {
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) implementsFunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction struct {
	Name string                                                                         `json:"name,required"`
	JSON functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionJSON `json:"-"`
}

// functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction]
type functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType string

const (
	FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType = "function"
)

func (r FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction:
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

func (r FunctionPromptDataOptionsParamsAnthropicModelParams) implementsFunctionPromptDataOptionsParamsUnion() {
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

func (r FunctionPromptDataOptionsParamsGoogleModelParams) implementsFunctionPromptDataOptionsParamsUnion() {
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

func (r FunctionPromptDataOptionsParamsWindowAIModelParams) implementsFunctionPromptDataOptionsParamsUnion() {
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

func (r FunctionPromptDataOptionsParamsJsCompletionParams) implementsFunctionPromptDataOptionsParamsUnion() {
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
	// [FunctionPromptDataPromptChatMessagesUserContentUnion].
	Content interface{}                              `json:"content,required"`
	Role    FunctionPromptDataPromptChatMessagesRole `json:"role,required"`
	Name    string                                   `json:"name"`
	// This field can have the runtime type of
	// [FunctionPromptDataPromptChatMessagesAssistantFunctionCall].
	FunctionCall interface{} `json:"function_call,required"`
	// This field can have the runtime type of
	// [[]FunctionPromptDataPromptChatMessagesAssistantToolCall].
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
// [FunctionPromptDataPromptChatMessagesSystem],
// [FunctionPromptDataPromptChatMessagesUser],
// [FunctionPromptDataPromptChatMessagesAssistant],
// [FunctionPromptDataPromptChatMessagesTool],
// [FunctionPromptDataPromptChatMessagesFunction],
// [FunctionPromptDataPromptChatMessagesFallback].
func (r FunctionPromptDataPromptChatMessage) AsUnion() FunctionPromptDataPromptChatMessagesUnion {
	return r.union
}

// Union satisfied by [FunctionPromptDataPromptChatMessagesSystem],
// [FunctionPromptDataPromptChatMessagesUser],
// [FunctionPromptDataPromptChatMessagesAssistant],
// [FunctionPromptDataPromptChatMessagesTool],
// [FunctionPromptDataPromptChatMessagesFunction] or
// [FunctionPromptDataPromptChatMessagesFallback].
type FunctionPromptDataPromptChatMessagesUnion interface {
	implementsFunctionPromptDataPromptChatMessage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataPromptChatMessagesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesSystem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesUser{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesAssistant{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesTool{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesFunction{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesFallback{}),
		},
	)
}

type FunctionPromptDataPromptChatMessagesSystem struct {
	Role    FunctionPromptDataPromptChatMessagesSystemRole `json:"role,required"`
	Content string                                         `json:"content"`
	Name    string                                         `json:"name"`
	JSON    functionPromptDataPromptChatMessagesSystemJSON `json:"-"`
}

// functionPromptDataPromptChatMessagesSystemJSON contains the JSON metadata for
// the struct [FunctionPromptDataPromptChatMessagesSystem]
type functionPromptDataPromptChatMessagesSystemJSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesSystem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesSystemJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesSystem) implementsFunctionPromptDataPromptChatMessage() {}

type FunctionPromptDataPromptChatMessagesSystemRole string

const (
	FunctionPromptDataPromptChatMessagesSystemRoleSystem FunctionPromptDataPromptChatMessagesSystemRole = "system"
)

func (r FunctionPromptDataPromptChatMessagesSystemRole) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesSystemRoleSystem:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesUser struct {
	Role    FunctionPromptDataPromptChatMessagesUserRole         `json:"role,required"`
	Content FunctionPromptDataPromptChatMessagesUserContentUnion `json:"content"`
	Name    string                                               `json:"name"`
	JSON    functionPromptDataPromptChatMessagesUserJSON         `json:"-"`
}

// functionPromptDataPromptChatMessagesUserJSON contains the JSON metadata for the
// struct [FunctionPromptDataPromptChatMessagesUser]
type functionPromptDataPromptChatMessagesUserJSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesUserJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesUser) implementsFunctionPromptDataPromptChatMessage() {}

type FunctionPromptDataPromptChatMessagesUserRole string

const (
	FunctionPromptDataPromptChatMessagesUserRoleUser FunctionPromptDataPromptChatMessagesUserRole = "user"
)

func (r FunctionPromptDataPromptChatMessagesUserRole) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesUserRoleUser:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or
// [FunctionPromptDataPromptChatMessagesUserContentArray].
type FunctionPromptDataPromptChatMessagesUserContentUnion interface {
	ImplementsFunctionPromptDataPromptChatMessagesUserContentUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataPromptChatMessagesUserContentUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesUserContentArray{}),
		},
	)
}

type FunctionPromptDataPromptChatMessagesUserContentArray []FunctionPromptDataPromptChatMessagesUserContentArrayItem

func (r FunctionPromptDataPromptChatMessagesUserContentArray) ImplementsFunctionPromptDataPromptChatMessagesUserContentUnion() {
}

type FunctionPromptDataPromptChatMessagesUserContentArrayItem struct {
	Text string                                                   `json:"text"`
	Type FunctionPromptDataPromptChatMessagesUserContentArrayType `json:"type,required"`
	// This field can have the runtime type of
	// [FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURL].
	ImageURL interface{}                                                  `json:"image_url,required"`
	JSON     functionPromptDataPromptChatMessagesUserContentArrayItemJSON `json:"-"`
	union    FunctionPromptDataPromptChatMessagesUserContentArrayUnionItem
}

// functionPromptDataPromptChatMessagesUserContentArrayItemJSON contains the JSON
// metadata for the struct
// [FunctionPromptDataPromptChatMessagesUserContentArrayItem]
type functionPromptDataPromptChatMessagesUserContentArrayItemJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	ImageURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r functionPromptDataPromptChatMessagesUserContentArrayItemJSON) RawJSON() string {
	return r.raw
}

func (r *FunctionPromptDataPromptChatMessagesUserContentArrayItem) UnmarshalJSON(data []byte) (err error) {
	*r = FunctionPromptDataPromptChatMessagesUserContentArrayItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [FunctionPromptDataPromptChatMessagesUserContentArrayUnionItem] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [FunctionPromptDataPromptChatMessagesUserContentArrayText],
// [FunctionPromptDataPromptChatMessagesUserContentArrayImageURL].
func (r FunctionPromptDataPromptChatMessagesUserContentArrayItem) AsUnion() FunctionPromptDataPromptChatMessagesUserContentArrayUnionItem {
	return r.union
}

// Union satisfied by [FunctionPromptDataPromptChatMessagesUserContentArrayText] or
// [FunctionPromptDataPromptChatMessagesUserContentArrayImageURL].
type FunctionPromptDataPromptChatMessagesUserContentArrayUnionItem interface {
	implementsFunctionPromptDataPromptChatMessagesUserContentArrayItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataPromptChatMessagesUserContentArrayUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesUserContentArrayText{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesUserContentArrayImageURL{}),
		},
	)
}

type FunctionPromptDataPromptChatMessagesUserContentArrayText struct {
	Type FunctionPromptDataPromptChatMessagesUserContentArrayTextType `json:"type,required"`
	Text string                                                       `json:"text"`
	JSON functionPromptDataPromptChatMessagesUserContentArrayTextJSON `json:"-"`
}

// functionPromptDataPromptChatMessagesUserContentArrayTextJSON contains the JSON
// metadata for the struct
// [FunctionPromptDataPromptChatMessagesUserContentArrayText]
type functionPromptDataPromptChatMessagesUserContentArrayTextJSON struct {
	Type        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesUserContentArrayText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesUserContentArrayTextJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesUserContentArrayText) implementsFunctionPromptDataPromptChatMessagesUserContentArrayItem() {
}

type FunctionPromptDataPromptChatMessagesUserContentArrayTextType string

const (
	FunctionPromptDataPromptChatMessagesUserContentArrayTextTypeText FunctionPromptDataPromptChatMessagesUserContentArrayTextType = "text"
)

func (r FunctionPromptDataPromptChatMessagesUserContentArrayTextType) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesUserContentArrayTextTypeText:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesUserContentArrayImageURL struct {
	ImageURL FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURL `json:"image_url,required"`
	Type     FunctionPromptDataPromptChatMessagesUserContentArrayImageURLType     `json:"type,required"`
	JSON     functionPromptDataPromptChatMessagesUserContentArrayImageURLJSON     `json:"-"`
}

// functionPromptDataPromptChatMessagesUserContentArrayImageURLJSON contains the
// JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesUserContentArrayImageURL]
type functionPromptDataPromptChatMessagesUserContentArrayImageURLJSON struct {
	ImageURL    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesUserContentArrayImageURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesUserContentArrayImageURLJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesUserContentArrayImageURL) implementsFunctionPromptDataPromptChatMessagesUserContentArrayItem() {
}

type FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURL struct {
	URL    string                                                                     `json:"url,required"`
	Detail FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail `json:"detail"`
	JSON   functionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLJSON   `json:"-"`
}

// functionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURL]
type functionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLJSON struct {
	URL         apijson.Field
	Detail      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail string

const (
	FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "auto"
	FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow  FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "low"
	FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "high"
)

func (r FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto, FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow, FunctionPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesUserContentArrayImageURLType string

const (
	FunctionPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL FunctionPromptDataPromptChatMessagesUserContentArrayImageURLType = "image_url"
)

func (r FunctionPromptDataPromptChatMessagesUserContentArrayImageURLType) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesUserContentArrayType string

const (
	FunctionPromptDataPromptChatMessagesUserContentArrayTypeText     FunctionPromptDataPromptChatMessagesUserContentArrayType = "text"
	FunctionPromptDataPromptChatMessagesUserContentArrayTypeImageURL FunctionPromptDataPromptChatMessagesUserContentArrayType = "image_url"
)

func (r FunctionPromptDataPromptChatMessagesUserContentArrayType) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesUserContentArrayTypeText, FunctionPromptDataPromptChatMessagesUserContentArrayTypeImageURL:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesAssistant struct {
	Role         FunctionPromptDataPromptChatMessagesAssistantRole         `json:"role,required"`
	Content      string                                                    `json:"content,nullable"`
	FunctionCall FunctionPromptDataPromptChatMessagesAssistantFunctionCall `json:"function_call"`
	Name         string                                                    `json:"name"`
	ToolCalls    []FunctionPromptDataPromptChatMessagesAssistantToolCall   `json:"tool_calls"`
	JSON         functionPromptDataPromptChatMessagesAssistantJSON         `json:"-"`
}

// functionPromptDataPromptChatMessagesAssistantJSON contains the JSON metadata for
// the struct [FunctionPromptDataPromptChatMessagesAssistant]
type functionPromptDataPromptChatMessagesAssistantJSON struct {
	Role         apijson.Field
	Content      apijson.Field
	FunctionCall apijson.Field
	Name         apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesAssistant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesAssistantJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesAssistant) implementsFunctionPromptDataPromptChatMessage() {
}

type FunctionPromptDataPromptChatMessagesAssistantRole string

const (
	FunctionPromptDataPromptChatMessagesAssistantRoleAssistant FunctionPromptDataPromptChatMessagesAssistantRole = "assistant"
)

func (r FunctionPromptDataPromptChatMessagesAssistantRole) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesAssistantRoleAssistant:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesAssistantFunctionCall struct {
	Arguments string                                                        `json:"arguments,required"`
	Name      string                                                        `json:"name,required"`
	JSON      functionPromptDataPromptChatMessagesAssistantFunctionCallJSON `json:"-"`
}

// functionPromptDataPromptChatMessagesAssistantFunctionCallJSON contains the JSON
// metadata for the struct
// [FunctionPromptDataPromptChatMessagesAssistantFunctionCall]
type functionPromptDataPromptChatMessagesAssistantFunctionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesAssistantFunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesAssistantFunctionCallJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataPromptChatMessagesAssistantToolCall struct {
	ID       string                                                         `json:"id,required"`
	Function FunctionPromptDataPromptChatMessagesAssistantToolCallsFunction `json:"function,required"`
	Type     FunctionPromptDataPromptChatMessagesAssistantToolCallsType     `json:"type,required"`
	JSON     functionPromptDataPromptChatMessagesAssistantToolCallJSON      `json:"-"`
}

// functionPromptDataPromptChatMessagesAssistantToolCallJSON contains the JSON
// metadata for the struct [FunctionPromptDataPromptChatMessagesAssistantToolCall]
type functionPromptDataPromptChatMessagesAssistantToolCallJSON struct {
	ID          apijson.Field
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesAssistantToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesAssistantToolCallJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataPromptChatMessagesAssistantToolCallsFunction struct {
	Arguments string                                                             `json:"arguments,required"`
	Name      string                                                             `json:"name,required"`
	JSON      functionPromptDataPromptChatMessagesAssistantToolCallsFunctionJSON `json:"-"`
}

// functionPromptDataPromptChatMessagesAssistantToolCallsFunctionJSON contains the
// JSON metadata for the struct
// [FunctionPromptDataPromptChatMessagesAssistantToolCallsFunction]
type functionPromptDataPromptChatMessagesAssistantToolCallsFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesAssistantToolCallsFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesAssistantToolCallsFunctionJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataPromptChatMessagesAssistantToolCallsType string

const (
	FunctionPromptDataPromptChatMessagesAssistantToolCallsTypeFunction FunctionPromptDataPromptChatMessagesAssistantToolCallsType = "function"
)

func (r FunctionPromptDataPromptChatMessagesAssistantToolCallsType) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesAssistantToolCallsTypeFunction:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesTool struct {
	Role       FunctionPromptDataPromptChatMessagesToolRole `json:"role,required"`
	Content    string                                       `json:"content"`
	ToolCallID string                                       `json:"tool_call_id"`
	JSON       functionPromptDataPromptChatMessagesToolJSON `json:"-"`
}

// functionPromptDataPromptChatMessagesToolJSON contains the JSON metadata for the
// struct [FunctionPromptDataPromptChatMessagesTool]
type functionPromptDataPromptChatMessagesToolJSON struct {
	Role        apijson.Field
	Content     apijson.Field
	ToolCallID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesToolJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesTool) implementsFunctionPromptDataPromptChatMessage() {}

type FunctionPromptDataPromptChatMessagesToolRole string

const (
	FunctionPromptDataPromptChatMessagesToolRoleTool FunctionPromptDataPromptChatMessagesToolRole = "tool"
)

func (r FunctionPromptDataPromptChatMessagesToolRole) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesToolRoleTool:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesFunction struct {
	Name    string                                           `json:"name,required"`
	Role    FunctionPromptDataPromptChatMessagesFunctionRole `json:"role,required"`
	Content string                                           `json:"content"`
	JSON    functionPromptDataPromptChatMessagesFunctionJSON `json:"-"`
}

// functionPromptDataPromptChatMessagesFunctionJSON contains the JSON metadata for
// the struct [FunctionPromptDataPromptChatMessagesFunction]
type functionPromptDataPromptChatMessagesFunctionJSON struct {
	Name        apijson.Field
	Role        apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesFunctionJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesFunction) implementsFunctionPromptDataPromptChatMessage() {
}

type FunctionPromptDataPromptChatMessagesFunctionRole string

const (
	FunctionPromptDataPromptChatMessagesFunctionRoleFunction FunctionPromptDataPromptChatMessagesFunctionRole = "function"
)

func (r FunctionPromptDataPromptChatMessagesFunctionRole) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesFunctionRoleFunction:
		return true
	}
	return false
}

type FunctionPromptDataPromptChatMessagesFallback struct {
	Role    FunctionPromptDataPromptChatMessagesFallbackRole `json:"role,required"`
	Content string                                           `json:"content,nullable"`
	JSON    functionPromptDataPromptChatMessagesFallbackJSON `json:"-"`
}

// functionPromptDataPromptChatMessagesFallbackJSON contains the JSON metadata for
// the struct [FunctionPromptDataPromptChatMessagesFallback]
type functionPromptDataPromptChatMessagesFallbackJSON struct {
	Role        apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesFallback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesFallbackJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesFallback) implementsFunctionPromptDataPromptChatMessage() {
}

type FunctionPromptDataPromptChatMessagesFallbackRole string

const (
	FunctionPromptDataPromptChatMessagesFallbackRoleModel FunctionPromptDataPromptChatMessagesFallbackRole = "model"
)

func (r FunctionPromptDataPromptChatMessagesFallbackRole) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesFallbackRoleModel:
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

// Satisfied by [FunctionNewParamsPromptDataOptionsParamsOpenAIModelParams],
// [FunctionNewParamsPromptDataOptionsParamsAnthropicModelParams],
// [FunctionNewParamsPromptDataOptionsParamsGoogleModelParams],
// [FunctionNewParamsPromptDataOptionsParamsWindowAIModelParams],
// [FunctionNewParamsPromptDataOptionsParamsJsCompletionParams].
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
// [FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto],
// [FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone],
// [FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction].
type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsFunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto string

const (
	FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto = "auto"
)

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto:
		return true
	}
	return false
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) implementsFunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone string

const (
	FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone = "none"
)

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone:
		return true
	}
	return false
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) implementsFunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) implementsFunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
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
// [FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto],
// [FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone],
// [FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction].
type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsFunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto string

const (
	FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto = "auto"
)

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto:
		return true
	}
	return false
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) implementsFunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone string

const (
	FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone = "none"
)

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone:
		return true
	}
	return false
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) implementsFunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction struct {
	Function param.Field[FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction] `json:"function,required"`
	Type     param.Field[FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType]     `json:"type,required"`
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) implementsFunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType string

const (
	FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType = "function"
)

func (r FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction:
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

// Satisfied by [FunctionNewParamsPromptDataPromptChatMessagesSystem],
// [FunctionNewParamsPromptDataPromptChatMessagesUser],
// [FunctionNewParamsPromptDataPromptChatMessagesAssistant],
// [FunctionNewParamsPromptDataPromptChatMessagesTool],
// [FunctionNewParamsPromptDataPromptChatMessagesFunction],
// [FunctionNewParamsPromptDataPromptChatMessagesFallback],
// [FunctionNewParamsPromptDataPromptChatMessage].
type FunctionNewParamsPromptDataPromptChatMessageUnion interface {
	implementsFunctionNewParamsPromptDataPromptChatMessageUnion()
}

type FunctionNewParamsPromptDataPromptChatMessagesSystem struct {
	Role    param.Field[FunctionNewParamsPromptDataPromptChatMessagesSystemRole] `json:"role,required"`
	Content param.Field[string]                                                  `json:"content"`
	Name    param.Field[string]                                                  `json:"name"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesSystem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesSystem) implementsFunctionNewParamsPromptDataPromptChatMessageUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesSystemRole string

const (
	FunctionNewParamsPromptDataPromptChatMessagesSystemRoleSystem FunctionNewParamsPromptDataPromptChatMessagesSystemRole = "system"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesSystemRole) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesSystemRoleSystem:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesUser struct {
	Role    param.Field[FunctionNewParamsPromptDataPromptChatMessagesUserRole]         `json:"role,required"`
	Content param.Field[FunctionNewParamsPromptDataPromptChatMessagesUserContentUnion] `json:"content"`
	Name    param.Field[string]                                                        `json:"name"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesUser) implementsFunctionNewParamsPromptDataPromptChatMessageUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesUserRole string

const (
	FunctionNewParamsPromptDataPromptChatMessagesUserRoleUser FunctionNewParamsPromptDataPromptChatMessagesUserRole = "user"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesUserRole) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesUserRoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [FunctionNewParamsPromptDataPromptChatMessagesUserContentArray].
type FunctionNewParamsPromptDataPromptChatMessagesUserContentUnion interface {
	ImplementsFunctionNewParamsPromptDataPromptChatMessagesUserContentUnion()
}

type FunctionNewParamsPromptDataPromptChatMessagesUserContentArray []FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayUnion

func (r FunctionNewParamsPromptDataPromptChatMessagesUserContentArray) ImplementsFunctionNewParamsPromptDataPromptChatMessagesUserContentUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesUserContentArray struct {
	Text     param.Field[string]                                                            `json:"text"`
	Type     param.Field[FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                       `json:"image_url,required"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesUserContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesUserContentArray) implementsFunctionNewParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

// Satisfied by
// [FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayText],
// [FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURL],
// [FunctionNewParamsPromptDataPromptChatMessagesUserContentArray].
type FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayUnion interface {
	implementsFunctionNewParamsPromptDataPromptChatMessagesUserContentArrayUnion()
}

type FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayText struct {
	Type param.Field[FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayTextType] `json:"type,required"`
	Text param.Field[string]                                                                `json:"text"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayText) implementsFunctionNewParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayTextType string

const (
	FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayTextTypeText FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayTextType = "text"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayTextType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayTextTypeText:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURL struct {
	ImageURL param.Field[FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL] `json:"image_url,required"`
	Type     param.Field[FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLType]     `json:"type,required"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURL) implementsFunctionNewParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL struct {
	URL    param.Field[string]                                                                              `json:"url,required"`
	Detail param.Field[FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail] `json:"detail"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail string

const (
	FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "auto"
	FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow  FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "low"
	FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "high"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto, FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow, FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLType string

const (
	FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLType = "image_url"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayType string

const (
	FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayTypeText     FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayType = "text"
	FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayTypeImageURL FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayType = "image_url"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayTypeText, FunctionNewParamsPromptDataPromptChatMessagesUserContentArrayTypeImageURL:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesAssistant struct {
	Role         param.Field[FunctionNewParamsPromptDataPromptChatMessagesAssistantRole]         `json:"role,required"`
	Content      param.Field[string]                                                             `json:"content"`
	FunctionCall param.Field[FunctionNewParamsPromptDataPromptChatMessagesAssistantFunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                             `json:"name"`
	ToolCalls    param.Field[[]FunctionNewParamsPromptDataPromptChatMessagesAssistantToolCall]   `json:"tool_calls"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesAssistant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesAssistant) implementsFunctionNewParamsPromptDataPromptChatMessageUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesAssistantRole string

const (
	FunctionNewParamsPromptDataPromptChatMessagesAssistantRoleAssistant FunctionNewParamsPromptDataPromptChatMessagesAssistantRole = "assistant"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesAssistantRole) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesAssistantRoleAssistant:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesAssistantFunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesAssistantFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataPromptChatMessagesAssistantToolCall struct {
	ID       param.Field[string]                                                                  `json:"id,required"`
	Function param.Field[FunctionNewParamsPromptDataPromptChatMessagesAssistantToolCallsFunction] `json:"function,required"`
	Type     param.Field[FunctionNewParamsPromptDataPromptChatMessagesAssistantToolCallsType]     `json:"type,required"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesAssistantToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataPromptChatMessagesAssistantToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesAssistantToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataPromptChatMessagesAssistantToolCallsType string

const (
	FunctionNewParamsPromptDataPromptChatMessagesAssistantToolCallsTypeFunction FunctionNewParamsPromptDataPromptChatMessagesAssistantToolCallsType = "function"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesAssistantToolCallsType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesAssistantToolCallsTypeFunction:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesTool struct {
	Role       param.Field[FunctionNewParamsPromptDataPromptChatMessagesToolRole] `json:"role,required"`
	Content    param.Field[string]                                                `json:"content"`
	ToolCallID param.Field[string]                                                `json:"tool_call_id"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesTool) implementsFunctionNewParamsPromptDataPromptChatMessageUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesToolRole string

const (
	FunctionNewParamsPromptDataPromptChatMessagesToolRoleTool FunctionNewParamsPromptDataPromptChatMessagesToolRole = "tool"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesToolRole) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesToolRoleTool:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesFunction struct {
	Name    param.Field[string]                                                    `json:"name,required"`
	Role    param.Field[FunctionNewParamsPromptDataPromptChatMessagesFunctionRole] `json:"role,required"`
	Content param.Field[string]                                                    `json:"content"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesFunction) implementsFunctionNewParamsPromptDataPromptChatMessageUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesFunctionRole string

const (
	FunctionNewParamsPromptDataPromptChatMessagesFunctionRoleFunction FunctionNewParamsPromptDataPromptChatMessagesFunctionRole = "function"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesFunctionRole) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesFunctionRoleFunction:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptChatMessagesFallback struct {
	Role    param.Field[FunctionNewParamsPromptDataPromptChatMessagesFallbackRole] `json:"role,required"`
	Content param.Field[string]                                                    `json:"content"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesFallback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesFallback) implementsFunctionNewParamsPromptDataPromptChatMessageUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesFallbackRole string

const (
	FunctionNewParamsPromptDataPromptChatMessagesFallbackRoleModel FunctionNewParamsPromptDataPromptChatMessagesFallbackRole = "model"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesFallbackRole) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesFallbackRoleModel:
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

// Satisfied by [FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParams],
// [FunctionUpdateParamsPromptDataOptionsParamsAnthropicModelParams],
// [FunctionUpdateParamsPromptDataOptionsParamsGoogleModelParams],
// [FunctionUpdateParamsPromptDataOptionsParamsWindowAIModelParams],
// [FunctionUpdateParamsPromptDataOptionsParamsJsCompletionParams].
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
// [FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto],
// [FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone],
// [FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction].
type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsFunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto string

const (
	FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto = "auto"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto:
		return true
	}
	return false
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) implementsFunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone string

const (
	FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone = "none"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone:
		return true
	}
	return false
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) implementsFunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) implementsFunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
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
// [FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto],
// [FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone],
// [FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction].
type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsFunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto string

const (
	FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto = "auto"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto:
		return true
	}
	return false
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) implementsFunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone string

const (
	FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone = "none"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone:
		return true
	}
	return false
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) implementsFunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction struct {
	Function param.Field[FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction] `json:"function,required"`
	Type     param.Field[FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType]     `json:"type,required"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) implementsFunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType string

const (
	FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType = "function"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction:
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

// Satisfied by [FunctionUpdateParamsPromptDataPromptChatMessagesSystem],
// [FunctionUpdateParamsPromptDataPromptChatMessagesUser],
// [FunctionUpdateParamsPromptDataPromptChatMessagesAssistant],
// [FunctionUpdateParamsPromptDataPromptChatMessagesTool],
// [FunctionUpdateParamsPromptDataPromptChatMessagesFunction],
// [FunctionUpdateParamsPromptDataPromptChatMessagesFallback],
// [FunctionUpdateParamsPromptDataPromptChatMessage].
type FunctionUpdateParamsPromptDataPromptChatMessageUnion interface {
	implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion()
}

type FunctionUpdateParamsPromptDataPromptChatMessagesSystem struct {
	Role    param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesSystemRole] `json:"role,required"`
	Content param.Field[string]                                                     `json:"content"`
	Name    param.Field[string]                                                     `json:"name"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesSystem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesSystem) implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesSystemRole string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesSystemRoleSystem FunctionUpdateParamsPromptDataPromptChatMessagesSystemRole = "system"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesSystemRole) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesSystemRoleSystem:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesUser struct {
	Role    param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesUserRole]         `json:"role,required"`
	Content param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesUserContentUnion] `json:"content"`
	Name    param.Field[string]                                                           `json:"name"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesUser) implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesUserRole string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesUserRoleUser FunctionUpdateParamsPromptDataPromptChatMessagesUserRole = "user"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesUserRole) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesUserRoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArray].
type FunctionUpdateParamsPromptDataPromptChatMessagesUserContentUnion interface {
	ImplementsFunctionUpdateParamsPromptDataPromptChatMessagesUserContentUnion()
}

type FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArray []FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayUnion

func (r FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArray) ImplementsFunctionUpdateParamsPromptDataPromptChatMessagesUserContentUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArray struct {
	Text     param.Field[string]                                                               `json:"text"`
	Type     param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                          `json:"image_url,required"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArray) implementsFunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

// Satisfied by
// [FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayText],
// [FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURL],
// [FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArray].
type FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayUnion interface {
	implementsFunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayUnion()
}

type FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayText struct {
	Type param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayTextType] `json:"type,required"`
	Text param.Field[string]                                                                   `json:"text"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayText) implementsFunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayTextType string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayTextTypeText FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayTextType = "text"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayTextType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayTextTypeText:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURL struct {
	ImageURL param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL] `json:"image_url,required"`
	Type     param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLType]     `json:"type,required"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURL) implementsFunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL struct {
	URL    param.Field[string]                                                                                 `json:"url,required"`
	Detail param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail] `json:"detail"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "auto"
	FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow  FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "low"
	FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "high"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto, FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow, FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLType string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLType = "image_url"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayType string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayTypeText     FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayType = "text"
	FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayTypeImageURL FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayType = "image_url"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayTypeText, FunctionUpdateParamsPromptDataPromptChatMessagesUserContentArrayTypeImageURL:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesAssistant struct {
	Role         param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesAssistantRole]         `json:"role,required"`
	Content      param.Field[string]                                                                `json:"content"`
	FunctionCall param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesAssistantFunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                                `json:"name"`
	ToolCalls    param.Field[[]FunctionUpdateParamsPromptDataPromptChatMessagesAssistantToolCall]   `json:"tool_calls"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesAssistant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesAssistant) implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesAssistantRole string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesAssistantRoleAssistant FunctionUpdateParamsPromptDataPromptChatMessagesAssistantRole = "assistant"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesAssistantRole) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesAssistantRoleAssistant:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesAssistantFunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesAssistantFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataPromptChatMessagesAssistantToolCall struct {
	ID       param.Field[string]                                                                     `json:"id,required"`
	Function param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsFunction] `json:"function,required"`
	Type     param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsType]     `json:"type,required"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesAssistantToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsType string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsTypeFunction FunctionUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsType = "function"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsTypeFunction:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesTool struct {
	Role       param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesToolRole] `json:"role,required"`
	Content    param.Field[string]                                                   `json:"content"`
	ToolCallID param.Field[string]                                                   `json:"tool_call_id"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesTool) implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesToolRole string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesToolRoleTool FunctionUpdateParamsPromptDataPromptChatMessagesToolRole = "tool"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesToolRole) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesToolRoleTool:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesFunction struct {
	Name    param.Field[string]                                                       `json:"name,required"`
	Role    param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesFunctionRole] `json:"role,required"`
	Content param.Field[string]                                                       `json:"content"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesFunction) implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesFunctionRole string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesFunctionRoleFunction FunctionUpdateParamsPromptDataPromptChatMessagesFunctionRole = "function"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesFunctionRole) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesFunctionRoleFunction:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptChatMessagesFallback struct {
	Role    param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesFallbackRole] `json:"role,required"`
	Content param.Field[string]                                                       `json:"content"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesFallback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesFallback) implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesFallbackRole string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesFallbackRoleModel FunctionUpdateParamsPromptDataPromptChatMessagesFallbackRole = "model"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesFallbackRole) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesFallbackRoleModel:
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

// Satisfied by [FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParams],
// [FunctionReplaceParamsPromptDataOptionsParamsAnthropicModelParams],
// [FunctionReplaceParamsPromptDataOptionsParamsGoogleModelParams],
// [FunctionReplaceParamsPromptDataOptionsParamsWindowAIModelParams],
// [FunctionReplaceParamsPromptDataOptionsParamsJsCompletionParams].
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
// [FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto],
// [FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone],
// [FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction].
type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsFunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto string

const (
	FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto = "auto"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto:
		return true
	}
	return false
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) implementsFunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone string

const (
	FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone = "none"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone:
		return true
	}
	return false
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) implementsFunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) implementsFunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
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
// [FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto],
// [FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone],
// [FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction].
type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsFunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto string

const (
	FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto = "auto"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto:
		return true
	}
	return false
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) implementsFunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone string

const (
	FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone = "none"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone:
		return true
	}
	return false
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) implementsFunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction struct {
	Function param.Field[FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction] `json:"function,required"`
	Type     param.Field[FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType]     `json:"type,required"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) implementsFunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType string

const (
	FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType = "function"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction:
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

// Satisfied by [FunctionReplaceParamsPromptDataPromptChatMessagesSystem],
// [FunctionReplaceParamsPromptDataPromptChatMessagesUser],
// [FunctionReplaceParamsPromptDataPromptChatMessagesAssistant],
// [FunctionReplaceParamsPromptDataPromptChatMessagesTool],
// [FunctionReplaceParamsPromptDataPromptChatMessagesFunction],
// [FunctionReplaceParamsPromptDataPromptChatMessagesFallback],
// [FunctionReplaceParamsPromptDataPromptChatMessage].
type FunctionReplaceParamsPromptDataPromptChatMessageUnion interface {
	implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion()
}

type FunctionReplaceParamsPromptDataPromptChatMessagesSystem struct {
	Role    param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesSystemRole] `json:"role,required"`
	Content param.Field[string]                                                      `json:"content"`
	Name    param.Field[string]                                                      `json:"name"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesSystem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesSystem) implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesSystemRole string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesSystemRoleSystem FunctionReplaceParamsPromptDataPromptChatMessagesSystemRole = "system"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesSystemRole) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesSystemRoleSystem:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesUser struct {
	Role    param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesUserRole]         `json:"role,required"`
	Content param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesUserContentUnion] `json:"content"`
	Name    param.Field[string]                                                            `json:"name"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesUser) implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesUserRole string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesUserRoleUser FunctionReplaceParamsPromptDataPromptChatMessagesUserRole = "user"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesUserRole) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesUserRoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArray].
type FunctionReplaceParamsPromptDataPromptChatMessagesUserContentUnion interface {
	ImplementsFunctionReplaceParamsPromptDataPromptChatMessagesUserContentUnion()
}

type FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArray []FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayUnion

func (r FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArray) ImplementsFunctionReplaceParamsPromptDataPromptChatMessagesUserContentUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArray struct {
	Text     param.Field[string]                                                                `json:"text"`
	Type     param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                           `json:"image_url,required"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArray) implementsFunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

// Satisfied by
// [FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayText],
// [FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURL],
// [FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArray].
type FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayUnion interface {
	implementsFunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayUnion()
}

type FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayText struct {
	Type param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayTextType] `json:"type,required"`
	Text param.Field[string]                                                                    `json:"text"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayText) implementsFunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayTextType string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayTextTypeText FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayTextType = "text"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayTextType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayTextTypeText:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURL struct {
	ImageURL param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL] `json:"image_url,required"`
	Type     param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLType]     `json:"type,required"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURL) implementsFunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL struct {
	URL    param.Field[string]                                                                                  `json:"url,required"`
	Detail param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail] `json:"detail"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "auto"
	FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow  FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "low"
	FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "high"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto, FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow, FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLType string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLType = "image_url"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayType string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayTypeText     FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayType = "text"
	FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayTypeImageURL FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayType = "image_url"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayTypeText, FunctionReplaceParamsPromptDataPromptChatMessagesUserContentArrayTypeImageURL:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesAssistant struct {
	Role         param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesAssistantRole]         `json:"role,required"`
	Content      param.Field[string]                                                                 `json:"content"`
	FunctionCall param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesAssistantFunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                                 `json:"name"`
	ToolCalls    param.Field[[]FunctionReplaceParamsPromptDataPromptChatMessagesAssistantToolCall]   `json:"tool_calls"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesAssistant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesAssistant) implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesAssistantRole string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesAssistantRoleAssistant FunctionReplaceParamsPromptDataPromptChatMessagesAssistantRole = "assistant"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesAssistantRole) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesAssistantRoleAssistant:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesAssistantFunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesAssistantFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataPromptChatMessagesAssistantToolCall struct {
	ID       param.Field[string]                                                                      `json:"id,required"`
	Function param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsFunction] `json:"function,required"`
	Type     param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsType]     `json:"type,required"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesAssistantToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsType string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsTypeFunction FunctionReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsType = "function"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsTypeFunction:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesTool struct {
	Role       param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesToolRole] `json:"role,required"`
	Content    param.Field[string]                                                    `json:"content"`
	ToolCallID param.Field[string]                                                    `json:"tool_call_id"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesTool) implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesToolRole string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesToolRoleTool FunctionReplaceParamsPromptDataPromptChatMessagesToolRole = "tool"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesToolRole) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesToolRoleTool:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesFunction struct {
	Name    param.Field[string]                                                        `json:"name,required"`
	Role    param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesFunctionRole] `json:"role,required"`
	Content param.Field[string]                                                        `json:"content"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesFunction) implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesFunctionRole string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesFunctionRoleFunction FunctionReplaceParamsPromptDataPromptChatMessagesFunctionRole = "function"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesFunctionRole) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesFunctionRoleFunction:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptChatMessagesFallback struct {
	Role    param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesFallbackRole] `json:"role,required"`
	Content param.Field[string]                                                        `json:"content"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesFallback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesFallback) implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesFallbackRole string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesFallbackRoleModel FunctionReplaceParamsPromptDataPromptChatMessagesFallbackRole = "model"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesFallbackRole) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesFallbackRoleModel:
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
