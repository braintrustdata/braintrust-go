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
	// This field can have the runtime type of [FunctionFunctionDataFunctionData1Data].
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
// [FunctionFunctionDataFunctionData1], [FunctionFunctionDataFunctionData2].
func (r FunctionFunctionData) AsUnion() FunctionFunctionDataUnion {
	return r.union
}

// Union satisfied by [FunctionFunctionDataPrompt],
// [FunctionFunctionDataFunctionData1] or [FunctionFunctionDataFunctionData2].
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
			Type:       reflect.TypeOf(FunctionFunctionDataFunctionData1{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataFunctionData2{}),
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

type FunctionFunctionDataFunctionData1 struct {
	Data FunctionFunctionDataFunctionData1Data `json:"data,required"`
	Type FunctionFunctionDataFunctionData1Type `json:"type,required"`
	JSON functionFunctionDataFunctionData1JSON `json:"-"`
}

// functionFunctionDataFunctionData1JSON contains the JSON metadata for the struct
// [FunctionFunctionDataFunctionData1]
type functionFunctionDataFunctionData1JSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataFunctionData1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataFunctionData1JSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataFunctionData1) implementsFunctionFunctionData() {}

type FunctionFunctionDataFunctionData1Data struct {
	BundleID       string                                              `json:"bundle_id,required"`
	Location       FunctionFunctionDataFunctionData1DataLocation       `json:"location,required"`
	RuntimeContext FunctionFunctionDataFunctionData1DataRuntimeContext `json:"runtime_context,required"`
	JSON           functionFunctionDataFunctionData1DataJSON           `json:"-"`
}

// functionFunctionDataFunctionData1DataJSON contains the JSON metadata for the
// struct [FunctionFunctionDataFunctionData1Data]
type functionFunctionDataFunctionData1DataJSON struct {
	BundleID       apijson.Field
	Location       apijson.Field
	RuntimeContext apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FunctionFunctionDataFunctionData1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataFunctionData1DataJSON) RawJSON() string {
	return r.raw
}

type FunctionFunctionDataFunctionData1DataLocation struct {
	EvalName string                                                     `json:"eval_name,required"`
	Position FunctionFunctionDataFunctionData1DataLocationPositionUnion `json:"position,required"`
	Type     FunctionFunctionDataFunctionData1DataLocationType          `json:"type,required"`
	JSON     functionFunctionDataFunctionData1DataLocationJSON          `json:"-"`
}

// functionFunctionDataFunctionData1DataLocationJSON contains the JSON metadata for
// the struct [FunctionFunctionDataFunctionData1DataLocation]
type functionFunctionDataFunctionData1DataLocationJSON struct {
	EvalName    apijson.Field
	Position    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataFunctionData1DataLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataFunctionData1DataLocationJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [FunctionFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1]
// or [FunctionFunctionDataFunctionData1DataLocationPositionScore].
type FunctionFunctionDataFunctionData1DataLocationPositionUnion interface {
	implementsFunctionFunctionDataFunctionData1DataLocationPositionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionFunctionDataFunctionData1DataLocationPositionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataFunctionData1DataLocationPositionScore{}),
		},
	)
}

type FunctionFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1 string

const (
	FunctionFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1Task FunctionFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1 = "task"
)

func (r FunctionFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1) IsKnown() bool {
	switch r {
	case FunctionFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1Task:
		return true
	}
	return false
}

func (r FunctionFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1) implementsFunctionFunctionDataFunctionData1DataLocationPositionUnion() {
}

type FunctionFunctionDataFunctionData1DataLocationPositionScore struct {
	Score float64                                                        `json:"score,required"`
	JSON  functionFunctionDataFunctionData1DataLocationPositionScoreJSON `json:"-"`
}

// functionFunctionDataFunctionData1DataLocationPositionScoreJSON contains the JSON
// metadata for the struct
// [FunctionFunctionDataFunctionData1DataLocationPositionScore]
type functionFunctionDataFunctionData1DataLocationPositionScoreJSON struct {
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataFunctionData1DataLocationPositionScore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataFunctionData1DataLocationPositionScoreJSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataFunctionData1DataLocationPositionScore) implementsFunctionFunctionDataFunctionData1DataLocationPositionUnion() {
}

type FunctionFunctionDataFunctionData1DataLocationType string

const (
	FunctionFunctionDataFunctionData1DataLocationTypeExperiment FunctionFunctionDataFunctionData1DataLocationType = "experiment"
)

func (r FunctionFunctionDataFunctionData1DataLocationType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataFunctionData1DataLocationTypeExperiment:
		return true
	}
	return false
}

type FunctionFunctionDataFunctionData1DataRuntimeContext struct {
	Runtime FunctionFunctionDataFunctionData1DataRuntimeContextRuntime `json:"runtime,required"`
	Version string                                                     `json:"version,required"`
	JSON    functionFunctionDataFunctionData1DataRuntimeContextJSON    `json:"-"`
}

// functionFunctionDataFunctionData1DataRuntimeContextJSON contains the JSON
// metadata for the struct [FunctionFunctionDataFunctionData1DataRuntimeContext]
type functionFunctionDataFunctionData1DataRuntimeContextJSON struct {
	Runtime     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataFunctionData1DataRuntimeContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataFunctionData1DataRuntimeContextJSON) RawJSON() string {
	return r.raw
}

type FunctionFunctionDataFunctionData1DataRuntimeContextRuntime string

const (
	FunctionFunctionDataFunctionData1DataRuntimeContextRuntimeNode FunctionFunctionDataFunctionData1DataRuntimeContextRuntime = "node"
)

func (r FunctionFunctionDataFunctionData1DataRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionFunctionDataFunctionData1DataRuntimeContextRuntimeNode:
		return true
	}
	return false
}

type FunctionFunctionDataFunctionData1Type string

const (
	FunctionFunctionDataFunctionData1TypeCode FunctionFunctionDataFunctionData1Type = "code"
)

func (r FunctionFunctionDataFunctionData1Type) IsKnown() bool {
	switch r {
	case FunctionFunctionDataFunctionData1TypeCode:
		return true
	}
	return false
}

type FunctionFunctionDataFunctionData2 struct {
	Name string                                `json:"name,required"`
	Type FunctionFunctionDataFunctionData2Type `json:"type,required"`
	JSON functionFunctionDataFunctionData2JSON `json:"-"`
}

// functionFunctionDataFunctionData2JSON contains the JSON metadata for the struct
// [FunctionFunctionDataFunctionData2]
type functionFunctionDataFunctionData2JSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataFunctionData2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataFunctionData2JSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataFunctionData2) implementsFunctionFunctionData() {}

type FunctionFunctionDataFunctionData2Type string

const (
	FunctionFunctionDataFunctionData2TypeGlobal FunctionFunctionDataFunctionData2Type = "global"
)

func (r FunctionFunctionDataFunctionData2Type) IsKnown() bool {
	switch r {
	case FunctionFunctionDataFunctionData2TypeGlobal:
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
	// [FunctionPromptDataOptionsParamsPromptDataOptions0ResponseFormat].
	ResponseFormat interface{} `json:"response_format,required"`
	// This field can have the runtime type of
	// [FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion].
	ToolChoice interface{} `json:"tool_choice,required"`
	// This field can have the runtime type of
	// [FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion].
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
// [FunctionPromptDataOptionsParamsPromptDataOptions0],
// [FunctionPromptDataOptionsParamsPromptDataOptions1],
// [FunctionPromptDataOptionsParamsPromptDataOptions2],
// [FunctionPromptDataOptionsParamsPromptDataOptions3],
// [FunctionPromptDataOptionsParamsJsCompletionParams].
func (r FunctionPromptDataOptionsParams) AsUnion() FunctionPromptDataOptionsParamsUnion {
	return r.union
}

// Union satisfied by [FunctionPromptDataOptionsParamsPromptDataOptions0],
// [FunctionPromptDataOptionsParamsPromptDataOptions1],
// [FunctionPromptDataOptionsParamsPromptDataOptions2],
// [FunctionPromptDataOptionsParamsPromptDataOptions3] or
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
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsPromptDataOptions0{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsPromptDataOptions1{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsPromptDataOptions2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsPromptDataOptions3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsJsCompletionParams{}),
		},
	)
}

type FunctionPromptDataOptionsParamsPromptDataOptions0 struct {
	FrequencyPenalty float64                                                            `json:"frequency_penalty"`
	FunctionCall     FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion `json:"function_call"`
	MaxTokens        float64                                                            `json:"max_tokens"`
	N                float64                                                            `json:"n"`
	PresencePenalty  float64                                                            `json:"presence_penalty"`
	ResponseFormat   FunctionPromptDataOptionsParamsPromptDataOptions0ResponseFormat    `json:"response_format,nullable"`
	Stop             []string                                                           `json:"stop"`
	Temperature      float64                                                            `json:"temperature"`
	ToolChoice       FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion   `json:"tool_choice"`
	TopP             float64                                                            `json:"top_p"`
	UseCache         bool                                                               `json:"use_cache"`
	JSON             functionPromptDataOptionsParamsPromptDataOptions0JSON              `json:"-"`
}

// functionPromptDataOptionsParamsPromptDataOptions0JSON contains the JSON metadata
// for the struct [FunctionPromptDataOptionsParamsPromptDataOptions0]
type functionPromptDataOptionsParamsPromptDataOptions0JSON struct {
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

func (r *FunctionPromptDataOptionsParamsPromptDataOptions0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsPromptDataOptions0JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsPromptDataOptions0) implementsFunctionPromptDataOptionsParams() {
}

// Union satisfied by
// [FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0],
// [FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1]
// or [FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallObject].
type FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion interface {
	implementsFunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallObject{}),
		},
	)
}

type FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 string

const (
	FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 = "auto"
)

func (r FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto:
		return true
	}
	return false
}

func (r FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) implementsFunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 string

const (
	FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 = "none"
)

func (r FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None:
		return true
	}
	return false
}

func (r FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) implementsFunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallObject struct {
	Name string                                                                  `json:"name,required"`
	JSON functionPromptDataOptionsParamsPromptDataOptions0FunctionCallObjectJSON `json:"-"`
}

// functionPromptDataOptionsParamsPromptDataOptions0FunctionCallObjectJSON contains
// the JSON metadata for the struct
// [FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallObject]
type functionPromptDataOptionsParamsPromptDataOptions0FunctionCallObjectJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsPromptDataOptions0FunctionCallObjectJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) implementsFunctionPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type FunctionPromptDataOptionsParamsPromptDataOptions0ResponseFormat struct {
	Type FunctionPromptDataOptionsParamsPromptDataOptions0ResponseFormatType `json:"type,required"`
	JSON functionPromptDataOptionsParamsPromptDataOptions0ResponseFormatJSON `json:"-"`
}

// functionPromptDataOptionsParamsPromptDataOptions0ResponseFormatJSON contains the
// JSON metadata for the struct
// [FunctionPromptDataOptionsParamsPromptDataOptions0ResponseFormat]
type functionPromptDataOptionsParamsPromptDataOptions0ResponseFormatJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsPromptDataOptions0ResponseFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsPromptDataOptions0ResponseFormatJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataOptionsParamsPromptDataOptions0ResponseFormatType string

const (
	FunctionPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject FunctionPromptDataOptionsParamsPromptDataOptions0ResponseFormatType = "json_object"
)

func (r FunctionPromptDataOptionsParamsPromptDataOptions0ResponseFormatType) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Union satisfied by
// [FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0],
// [FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1]
// or
// [FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2].
type FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion interface {
	implementsFunctionPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2{}),
		},
	)
}

type FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 string

const (
	FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 = "auto"
)

func (r FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto:
		return true
	}
	return false
}

func (r FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) implementsFunctionPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 string

const (
	FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 = "none"
)

func (r FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None:
		return true
	}
	return false
}

func (r FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) implementsFunctionPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2 struct {
	Function FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function `json:"function,required"`
	Type     FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type     `json:"type,required"`
	JSON     functionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2JSON     `json:"-"`
}

// functionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2JSON
// contains the JSON metadata for the struct
// [FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2]
type functionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2JSON struct {
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) implementsFunctionPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function struct {
	Name string                                                                                       `json:"name,required"`
	JSON functionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2FunctionJSON `json:"-"`
}

// functionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2FunctionJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function]
type functionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2FunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2FunctionJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type string

const (
	FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type = "function"
)

func (r FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction:
		return true
	}
	return false
}

type FunctionPromptDataOptionsParamsPromptDataOptions1 struct {
	MaxTokens   float64 `json:"max_tokens,required"`
	Temperature float64 `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample float64                                               `json:"max_tokens_to_sample"`
	StopSequences     []string                                              `json:"stop_sequences"`
	TopK              float64                                               `json:"top_k"`
	TopP              float64                                               `json:"top_p"`
	UseCache          bool                                                  `json:"use_cache"`
	JSON              functionPromptDataOptionsParamsPromptDataOptions1JSON `json:"-"`
}

// functionPromptDataOptionsParamsPromptDataOptions1JSON contains the JSON metadata
// for the struct [FunctionPromptDataOptionsParamsPromptDataOptions1]
type functionPromptDataOptionsParamsPromptDataOptions1JSON struct {
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

func (r *FunctionPromptDataOptionsParamsPromptDataOptions1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsPromptDataOptions1JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsPromptDataOptions1) implementsFunctionPromptDataOptionsParams() {
}

type FunctionPromptDataOptionsParamsPromptDataOptions2 struct {
	MaxOutputTokens float64                                               `json:"maxOutputTokens"`
	Temperature     float64                                               `json:"temperature"`
	TopK            float64                                               `json:"topK"`
	TopP            float64                                               `json:"topP"`
	UseCache        bool                                                  `json:"use_cache"`
	JSON            functionPromptDataOptionsParamsPromptDataOptions2JSON `json:"-"`
}

// functionPromptDataOptionsParamsPromptDataOptions2JSON contains the JSON metadata
// for the struct [FunctionPromptDataOptionsParamsPromptDataOptions2]
type functionPromptDataOptionsParamsPromptDataOptions2JSON struct {
	MaxOutputTokens apijson.Field
	Temperature     apijson.Field
	TopK            apijson.Field
	TopP            apijson.Field
	UseCache        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsPromptDataOptions2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsPromptDataOptions2JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsPromptDataOptions2) implementsFunctionPromptDataOptionsParams() {
}

type FunctionPromptDataOptionsParamsPromptDataOptions3 struct {
	Temperature float64                                               `json:"temperature"`
	TopK        float64                                               `json:"topK"`
	UseCache    bool                                                  `json:"use_cache"`
	JSON        functionPromptDataOptionsParamsPromptDataOptions3JSON `json:"-"`
}

// functionPromptDataOptionsParamsPromptDataOptions3JSON contains the JSON metadata
// for the struct [FunctionPromptDataOptionsParamsPromptDataOptions3]
type functionPromptDataOptionsParamsPromptDataOptions3JSON struct {
	Temperature apijson.Field
	TopK        apijson.Field
	UseCache    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsPromptDataOptions3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsPromptDataOptions3JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsPromptDataOptions3) implementsFunctionPromptDataOptionsParams() {
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
	// This field can have the runtime type of
	// [[]FunctionPromptDataPromptPromptDataPrompt1Message].
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
// Possible runtime types of the union are
// [FunctionPromptDataPromptPromptDataPrompt0],
// [FunctionPromptDataPromptPromptDataPrompt1],
// [FunctionPromptDataPromptPromptDataPrompt2].
func (r FunctionPromptDataPrompt) AsUnion() FunctionPromptDataPromptUnion {
	return r.union
}

// Union satisfied by [FunctionPromptDataPromptPromptDataPrompt0],
// [FunctionPromptDataPromptPromptDataPrompt1] or
// [FunctionPromptDataPromptPromptDataPrompt2].
type FunctionPromptDataPromptUnion interface {
	implementsFunctionPromptDataPrompt()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataPromptUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptPromptDataPrompt0{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptPromptDataPrompt1{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptPromptDataPrompt2{}),
		},
	)
}

type FunctionPromptDataPromptPromptDataPrompt0 struct {
	Content string                                        `json:"content,required"`
	Type    FunctionPromptDataPromptPromptDataPrompt0Type `json:"type,required"`
	JSON    functionPromptDataPromptPromptDataPrompt0JSON `json:"-"`
}

// functionPromptDataPromptPromptDataPrompt0JSON contains the JSON metadata for the
// struct [FunctionPromptDataPromptPromptDataPrompt0]
type functionPromptDataPromptPromptDataPrompt0JSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptPromptDataPrompt0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptPromptDataPrompt0JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptPromptDataPrompt0) implementsFunctionPromptDataPrompt() {}

type FunctionPromptDataPromptPromptDataPrompt0Type string

const (
	FunctionPromptDataPromptPromptDataPrompt0TypeCompletion FunctionPromptDataPromptPromptDataPrompt0Type = "completion"
)

func (r FunctionPromptDataPromptPromptDataPrompt0Type) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptPromptDataPrompt0TypeCompletion:
		return true
	}
	return false
}

type FunctionPromptDataPromptPromptDataPrompt1 struct {
	Messages []FunctionPromptDataPromptPromptDataPrompt1Message `json:"messages,required"`
	Type     FunctionPromptDataPromptPromptDataPrompt1Type      `json:"type,required"`
	Tools    string                                             `json:"tools"`
	JSON     functionPromptDataPromptPromptDataPrompt1JSON      `json:"-"`
}

// functionPromptDataPromptPromptDataPrompt1JSON contains the JSON metadata for the
// struct [FunctionPromptDataPromptPromptDataPrompt1]
type functionPromptDataPromptPromptDataPrompt1JSON struct {
	Messages    apijson.Field
	Type        apijson.Field
	Tools       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptPromptDataPrompt1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptPromptDataPrompt1JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptPromptDataPrompt1) implementsFunctionPromptDataPrompt() {}

type FunctionPromptDataPromptPromptDataPrompt1Message struct {
	// This field can have the runtime type of [string],
	// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion].
	Content interface{}                                           `json:"content,required"`
	Role    FunctionPromptDataPromptPromptDataPrompt1MessagesRole `json:"role,required"`
	Name    string                                                `json:"name"`
	// This field can have the runtime type of
	// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall].
	FunctionCall interface{} `json:"function_call,required"`
	// This field can have the runtime type of
	// [[]FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall].
	ToolCalls  interface{}                                          `json:"tool_calls,required"`
	ToolCallID string                                               `json:"tool_call_id"`
	JSON       functionPromptDataPromptPromptDataPrompt1MessageJSON `json:"-"`
	union      FunctionPromptDataPromptPromptDataPrompt1MessagesUnion
}

// functionPromptDataPromptPromptDataPrompt1MessageJSON contains the JSON metadata
// for the struct [FunctionPromptDataPromptPromptDataPrompt1Message]
type functionPromptDataPromptPromptDataPrompt1MessageJSON struct {
	Content      apijson.Field
	Role         apijson.Field
	Name         apijson.Field
	FunctionCall apijson.Field
	ToolCalls    apijson.Field
	ToolCallID   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r functionPromptDataPromptPromptDataPrompt1MessageJSON) RawJSON() string {
	return r.raw
}

func (r *FunctionPromptDataPromptPromptDataPrompt1Message) UnmarshalJSON(data []byte) (err error) {
	*r = FunctionPromptDataPromptPromptDataPrompt1Message{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FunctionPromptDataPromptPromptDataPrompt1MessagesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0],
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1],
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2],
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3],
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4],
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5].
func (r FunctionPromptDataPromptPromptDataPrompt1Message) AsUnion() FunctionPromptDataPromptPromptDataPrompt1MessagesUnion {
	return r.union
}

// Union satisfied by
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0],
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1],
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2],
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3],
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4] or
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5].
type FunctionPromptDataPromptPromptDataPrompt1MessagesUnion interface {
	implementsFunctionPromptDataPromptPromptDataPrompt1Message()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataPromptPromptDataPrompt1MessagesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5{}),
		},
	)
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0 struct {
	Role    FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role `json:"role,required"`
	Content string                                                                        `json:"content"`
	Name    string                                                                        `json:"name"`
	JSON    functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0JSON `json:"-"`
}

// functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0JSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0]
type functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) implementsFunctionPromptDataPromptPromptDataPrompt1Message() {
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role string

const (
	FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role = "system"
)

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1 struct {
	Role    FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role         `json:"role,required"`
	Content FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion `json:"content"`
	Name    string                                                                                `json:"name"`
	JSON    functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1JSON         `json:"-"`
}

// functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1JSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1]
type functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) implementsFunctionPromptDataPromptPromptDataPrompt1Message() {
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role string

const (
	FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role = "user"
)

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray].
type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsFunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray{}),
		},
	)
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray []FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) ImplementsFunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion() {
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem struct {
	Text string                                                                                    `json:"text"`
	Type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType `json:"type,required"`
	// This field can have the runtime type of
	// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL].
	ImageURL interface{}                                                                                   `json:"image_url,required"`
	JSON     functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItemJSON `json:"-"`
	union    FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnionItem
}

// functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItemJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem]
type functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItemJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	ImageURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItemJSON) RawJSON() string {
	return r.raw
}

func (r *FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem) UnmarshalJSON(data []byte) (err error) {
	*r = FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnionItem]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList].
func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem) AsUnion() FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnionItem {
	return r.union
}

// Union satisfied by
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent]
// or
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList].
type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnionItem interface {
	implementsFunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList{}),
		},
	)
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType `json:"type,required"`
	Text string                                                                                                                  `json:"text"`
	JSON functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON `json:"-"`
}

// functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent]
type functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON struct {
	Type        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsFunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem() {
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL `json:"image_url,required"`
	Type     FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType     `json:"type,required"`
	JSON     functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON     `json:"-"`
}

// functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList]
type functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON struct {
	ImageURL    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsFunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem() {
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    string                                                                                                                         `json:"url,required"`
	Detail FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail `json:"detail"`
	JSON   functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON   `json:"-"`
}

// functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL]
type functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON struct {
	URL         apijson.Field
	Detail      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType string

const (
	FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText     FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "text"
	FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText, FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2 struct {
	Role         FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role         `json:"role,required"`
	Content      string                                                                                `json:"content,nullable"`
	FunctionCall FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall `json:"function_call"`
	Name         string                                                                                `json:"name"`
	ToolCalls    []FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall   `json:"tool_calls"`
	JSON         functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2JSON         `json:"-"`
}

// functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2JSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2]
type functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2JSON struct {
	Role         apijson.Field
	Content      apijson.Field
	FunctionCall apijson.Field
	Name         apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) implementsFunctionPromptDataPromptPromptDataPrompt1Message() {
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role string

const (
	FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role = "assistant"
)

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments string                                                                                    `json:"arguments,required"`
	Name      string                                                                                    `json:"name,required"`
	JSON      functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCallJSON `json:"-"`
}

// functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCallJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall]
type functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCallJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall struct {
	ID       string                                                                                     `json:"id,required"`
	Function FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction `json:"function,required"`
	Type     FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType     `json:"type,required"`
	JSON     functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallJSON      `json:"-"`
}

// functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall]
type functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallJSON struct {
	ID          apijson.Field
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments string                                                                                         `json:"arguments,required"`
	Name      string                                                                                         `json:"name,required"`
	JSON      functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunctionJSON `json:"-"`
}

// functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunctionJSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction]
type functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunctionJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType string

const (
	FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3 struct {
	Role       FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role `json:"role,required"`
	Content    string                                                                        `json:"content"`
	ToolCallID string                                                                        `json:"tool_call_id"`
	JSON       functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3JSON `json:"-"`
}

// functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3JSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3]
type functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	ToolCallID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) implementsFunctionPromptDataPromptPromptDataPrompt1Message() {
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role string

const (
	FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role = "tool"
)

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4 struct {
	Name    string                                                                        `json:"name,required"`
	Role    FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role `json:"role,required"`
	Content string                                                                        `json:"content"`
	JSON    functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4JSON `json:"-"`
}

// functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4JSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4]
type functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4JSON struct {
	Name        apijson.Field
	Role        apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) implementsFunctionPromptDataPromptPromptDataPrompt1Message() {
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role string

const (
	FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role = "function"
)

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5 struct {
	Role    FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role `json:"role,required"`
	Content string                                                                        `json:"content,nullable"`
	JSON    functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5JSON `json:"-"`
}

// functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5JSON
// contains the JSON metadata for the struct
// [FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5]
type functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) implementsFunctionPromptDataPromptPromptDataPrompt1Message() {
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role string

const (
	FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role = "model"
)

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type FunctionPromptDataPromptPromptDataPrompt1MessagesRole string

const (
	FunctionPromptDataPromptPromptDataPrompt1MessagesRoleSystem    FunctionPromptDataPromptPromptDataPrompt1MessagesRole = "system"
	FunctionPromptDataPromptPromptDataPrompt1MessagesRoleUser      FunctionPromptDataPromptPromptDataPrompt1MessagesRole = "user"
	FunctionPromptDataPromptPromptDataPrompt1MessagesRoleAssistant FunctionPromptDataPromptPromptDataPrompt1MessagesRole = "assistant"
	FunctionPromptDataPromptPromptDataPrompt1MessagesRoleTool      FunctionPromptDataPromptPromptDataPrompt1MessagesRole = "tool"
	FunctionPromptDataPromptPromptDataPrompt1MessagesRoleFunction  FunctionPromptDataPromptPromptDataPrompt1MessagesRole = "function"
	FunctionPromptDataPromptPromptDataPrompt1MessagesRoleModel     FunctionPromptDataPromptPromptDataPrompt1MessagesRole = "model"
)

func (r FunctionPromptDataPromptPromptDataPrompt1MessagesRole) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptPromptDataPrompt1MessagesRoleSystem, FunctionPromptDataPromptPromptDataPrompt1MessagesRoleUser, FunctionPromptDataPromptPromptDataPrompt1MessagesRoleAssistant, FunctionPromptDataPromptPromptDataPrompt1MessagesRoleTool, FunctionPromptDataPromptPromptDataPrompt1MessagesRoleFunction, FunctionPromptDataPromptPromptDataPrompt1MessagesRoleModel:
		return true
	}
	return false
}

type FunctionPromptDataPromptPromptDataPrompt1Type string

const (
	FunctionPromptDataPromptPromptDataPrompt1TypeChat FunctionPromptDataPromptPromptDataPrompt1Type = "chat"
)

func (r FunctionPromptDataPromptPromptDataPrompt1Type) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptPromptDataPrompt1TypeChat:
		return true
	}
	return false
}

type FunctionPromptDataPromptPromptDataPrompt2 struct {
	JSON functionPromptDataPromptPromptDataPrompt2JSON `json:"-"`
}

// functionPromptDataPromptPromptDataPrompt2JSON contains the JSON metadata for the
// struct [FunctionPromptDataPromptPromptDataPrompt2]
type functionPromptDataPromptPromptDataPrompt2JSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptPromptDataPrompt2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptPromptDataPrompt2JSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptPromptDataPrompt2) implementsFunctionPromptDataPrompt() {}

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
// [FunctionNewParamsFunctionDataFunctionData1],
// [FunctionNewParamsFunctionDataFunctionData2], [FunctionNewParamsFunctionData].
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

type FunctionNewParamsFunctionDataFunctionData1 struct {
	Data param.Field[FunctionNewParamsFunctionDataFunctionData1Data] `json:"data,required"`
	Type param.Field[FunctionNewParamsFunctionDataFunctionData1Type] `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataFunctionData1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataFunctionData1) implementsFunctionNewParamsFunctionDataUnion() {}

type FunctionNewParamsFunctionDataFunctionData1Data struct {
	BundleID       param.Field[string]                                                       `json:"bundle_id,required"`
	Location       param.Field[FunctionNewParamsFunctionDataFunctionData1DataLocation]       `json:"location,required"`
	RuntimeContext param.Field[FunctionNewParamsFunctionDataFunctionData1DataRuntimeContext] `json:"runtime_context,required"`
}

func (r FunctionNewParamsFunctionDataFunctionData1Data) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsFunctionDataFunctionData1DataLocation struct {
	EvalName param.Field[string]                                                              `json:"eval_name,required"`
	Position param.Field[FunctionNewParamsFunctionDataFunctionData1DataLocationPositionUnion] `json:"position,required"`
	Type     param.Field[FunctionNewParamsFunctionDataFunctionData1DataLocationType]          `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataFunctionData1DataLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [FunctionNewParamsFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1],
// [FunctionNewParamsFunctionDataFunctionData1DataLocationPositionScore].
type FunctionNewParamsFunctionDataFunctionData1DataLocationPositionUnion interface {
	implementsFunctionNewParamsFunctionDataFunctionData1DataLocationPositionUnion()
}

type FunctionNewParamsFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1 string

const (
	FunctionNewParamsFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1Task FunctionNewParamsFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1 = "task"
)

func (r FunctionNewParamsFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1Task:
		return true
	}
	return false
}

func (r FunctionNewParamsFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1) implementsFunctionNewParamsFunctionDataFunctionData1DataLocationPositionUnion() {
}

type FunctionNewParamsFunctionDataFunctionData1DataLocationPositionScore struct {
	Score param.Field[float64] `json:"score,required"`
}

func (r FunctionNewParamsFunctionDataFunctionData1DataLocationPositionScore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataFunctionData1DataLocationPositionScore) implementsFunctionNewParamsFunctionDataFunctionData1DataLocationPositionUnion() {
}

type FunctionNewParamsFunctionDataFunctionData1DataLocationType string

const (
	FunctionNewParamsFunctionDataFunctionData1DataLocationTypeExperiment FunctionNewParamsFunctionDataFunctionData1DataLocationType = "experiment"
)

func (r FunctionNewParamsFunctionDataFunctionData1DataLocationType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataFunctionData1DataLocationTypeExperiment:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataFunctionData1DataRuntimeContext struct {
	Runtime param.Field[FunctionNewParamsFunctionDataFunctionData1DataRuntimeContextRuntime] `json:"runtime,required"`
	Version param.Field[string]                                                              `json:"version,required"`
}

func (r FunctionNewParamsFunctionDataFunctionData1DataRuntimeContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsFunctionDataFunctionData1DataRuntimeContextRuntime string

const (
	FunctionNewParamsFunctionDataFunctionData1DataRuntimeContextRuntimeNode FunctionNewParamsFunctionDataFunctionData1DataRuntimeContextRuntime = "node"
)

func (r FunctionNewParamsFunctionDataFunctionData1DataRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataFunctionData1DataRuntimeContextRuntimeNode:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataFunctionData1Type string

const (
	FunctionNewParamsFunctionDataFunctionData1TypeCode FunctionNewParamsFunctionDataFunctionData1Type = "code"
)

func (r FunctionNewParamsFunctionDataFunctionData1Type) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataFunctionData1TypeCode:
		return true
	}
	return false
}

type FunctionNewParamsFunctionDataFunctionData2 struct {
	Name param.Field[string]                                         `json:"name,required"`
	Type param.Field[FunctionNewParamsFunctionDataFunctionData2Type] `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataFunctionData2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataFunctionData2) implementsFunctionNewParamsFunctionDataUnion() {}

type FunctionNewParamsFunctionDataFunctionData2Type string

const (
	FunctionNewParamsFunctionDataFunctionData2TypeGlobal FunctionNewParamsFunctionDataFunctionData2Type = "global"
)

func (r FunctionNewParamsFunctionDataFunctionData2Type) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataFunctionData2TypeGlobal:
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

// Satisfied by [FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0],
// [FunctionNewParamsPromptDataOptionsParamsPromptDataOptions1],
// [FunctionNewParamsPromptDataOptionsParamsPromptDataOptions2],
// [FunctionNewParamsPromptDataOptionsParamsPromptDataOptions3],
// [FunctionNewParamsPromptDataOptionsParamsJsCompletionParams],
// [FunctionNewParamsPromptDataOptionsParams].
type FunctionNewParamsPromptDataOptionsParamsUnion interface {
	implementsFunctionNewParamsPromptDataOptionsParamsUnion()
}

type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0 struct {
	FrequencyPenalty param.Field[float64]                                                                     `json:"frequency_penalty"`
	FunctionCall     param.Field[FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                                     `json:"max_tokens"`
	N                param.Field[float64]                                                                     `json:"n"`
	PresencePenalty  param.Field[float64]                                                                     `json:"presence_penalty"`
	ResponseFormat   param.Field[FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                                    `json:"stop"`
	Temperature      param.Field[float64]                                                                     `json:"temperature"`
	ToolChoice       param.Field[FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                                     `json:"top_p"`
	UseCache         param.Field[bool]                                                                        `json:"use_cache"`
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0) implementsFunctionNewParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by
// [FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0],
// [FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1],
// [FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject].
type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion interface {
	implementsFunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion()
}

type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 string

const (
	FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 = "auto"
)

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto:
		return true
	}
	return false
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) implementsFunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 string

const (
	FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 = "none"
)

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None:
		return true
	}
	return false
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) implementsFunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) implementsFunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat struct {
	Type param.Field[FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType] `json:"type,required"`
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType string

const (
	FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType = "json_object"
)

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Satisfied by
// [FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0],
// [FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1],
// [FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2].
type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion interface {
	implementsFunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion()
}

type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 string

const (
	FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 = "auto"
)

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto:
		return true
	}
	return false
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) implementsFunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 string

const (
	FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 = "none"
)

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None:
		return true
	}
	return false
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) implementsFunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2 struct {
	Function param.Field[FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function] `json:"function,required"`
	Type     param.Field[FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type]     `json:"type,required"`
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) implementsFunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type string

const (
	FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type = "function"
)

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions1 struct {
	MaxTokens   param.Field[float64] `json:"max_tokens,required"`
	Temperature param.Field[float64] `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64]  `json:"max_tokens_to_sample"`
	StopSequences     param.Field[[]string] `json:"stop_sequences"`
	TopK              param.Field[float64]  `json:"top_k"`
	TopP              param.Field[float64]  `json:"top_p"`
	UseCache          param.Field[bool]     `json:"use_cache"`
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions1) implementsFunctionNewParamsPromptDataOptionsParamsUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions2 struct {
	MaxOutputTokens param.Field[float64] `json:"maxOutputTokens"`
	Temperature     param.Field[float64] `json:"temperature"`
	TopK            param.Field[float64] `json:"topK"`
	TopP            param.Field[float64] `json:"topP"`
	UseCache        param.Field[bool]    `json:"use_cache"`
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions2) implementsFunctionNewParamsPromptDataOptionsParamsUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsPromptDataOptions3 struct {
	Temperature param.Field[float64] `json:"temperature"`
	TopK        param.Field[float64] `json:"topK"`
	UseCache    param.Field[bool]    `json:"use_cache"`
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsPromptDataOptions3) implementsFunctionNewParamsPromptDataOptionsParamsUnion() {
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

// Satisfied by [FunctionNewParamsPromptDataPromptPromptDataPrompt0],
// [FunctionNewParamsPromptDataPromptPromptDataPrompt1],
// [FunctionNewParamsPromptDataPromptPromptDataPrompt2],
// [FunctionNewParamsPromptDataPrompt].
type FunctionNewParamsPromptDataPromptUnion interface {
	implementsFunctionNewParamsPromptDataPromptUnion()
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt0 struct {
	Content param.Field[string]                                                 `json:"content,required"`
	Type    param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt0Type] `json:"type,required"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt0) implementsFunctionNewParamsPromptDataPromptUnion() {
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt0Type string

const (
	FunctionNewParamsPromptDataPromptPromptDataPrompt0TypeCompletion FunctionNewParamsPromptDataPromptPromptDataPrompt0Type = "completion"
)

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt0Type) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptPromptDataPrompt0TypeCompletion:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1 struct {
	Messages param.Field[[]FunctionNewParamsPromptDataPromptPromptDataPrompt1MessageUnion] `json:"messages,required"`
	Type     param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1Type]           `json:"type,required"`
	Tools    param.Field[string]                                                           `json:"tools"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1) implementsFunctionNewParamsPromptDataPromptUnion() {
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1Message struct {
	Content      param.Field[interface{}]                                                    `json:"content,required"`
	Role         param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRole] `json:"role,required"`
	Name         param.Field[string]                                                         `json:"name"`
	FunctionCall param.Field[interface{}]                                                    `json:"function_call,required"`
	ToolCalls    param.Field[interface{}]                                                    `json:"tool_calls,required"`
	ToolCallID   param.Field[string]                                                         `json:"tool_call_id"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1Message) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1Message) implementsFunctionNewParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

// Satisfied by
// [FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0],
// [FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1],
// [FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2],
// [FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3],
// [FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4],
// [FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5],
// [FunctionNewParamsPromptDataPromptPromptDataPrompt1Message].
type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessageUnion interface {
	implementsFunctionNewParamsPromptDataPromptPromptDataPrompt1MessageUnion()
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0 struct {
	Role    param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role] `json:"role,required"`
	Content param.Field[string]                                                                                 `json:"content"`
	Name    param.Field[string]                                                                                 `json:"name"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) implementsFunctionNewParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role string

const (
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role = "system"
)

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1 struct {
	Role    param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role]         `json:"role,required"`
	Content param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion] `json:"content"`
	Name    param.Field[string]                                                                                         `json:"name"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) implementsFunctionNewParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role string

const (
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role = "user"
)

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray].
type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsFunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion()
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray []FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) ImplementsFunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion() {
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray struct {
	Text     param.Field[string]                                                                                             `json:"text"`
	Type     param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                                                        `json:"image_url,required"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) implementsFunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

// Satisfied by
// [FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList],
// [FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray].
type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion interface {
	implementsFunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion()
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType] `json:"type,required"`
	Text param.Field[string]                                                                                                                           `json:"text"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsFunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL] `json:"image_url,required"`
	Type     param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType]     `json:"type,required"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsFunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    param.Field[string]                                                                                                                                  `json:"url,required"`
	Detail param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail] `json:"detail"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType string

const (
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText     FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "text"
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText, FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2 struct {
	Role         param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role]         `json:"role,required"`
	Content      param.Field[string]                                                                                         `json:"content"`
	FunctionCall param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                                                         `json:"name"`
	ToolCalls    param.Field[[]FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall]   `json:"tool_calls"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) implementsFunctionNewParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role string

const (
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role = "assistant"
)

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall struct {
	ID       param.Field[string]                                                                                              `json:"id,required"`
	Function param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction] `json:"function,required"`
	Type     param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType]     `json:"type,required"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType string

const (
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3 struct {
	Role       param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role] `json:"role,required"`
	Content    param.Field[string]                                                                                 `json:"content"`
	ToolCallID param.Field[string]                                                                                 `json:"tool_call_id"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) implementsFunctionNewParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role string

const (
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role = "tool"
)

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4 struct {
	Name    param.Field[string]                                                                                 `json:"name,required"`
	Role    param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role] `json:"role,required"`
	Content param.Field[string]                                                                                 `json:"content"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) implementsFunctionNewParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role string

const (
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role = "function"
)

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5 struct {
	Role    param.Field[FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role] `json:"role,required"`
	Content param.Field[string]                                                                                 `json:"content"`
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) implementsFunctionNewParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role string

const (
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role = "model"
)

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRole string

const (
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleSystem    FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRole = "system"
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleUser      FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRole = "user"
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleAssistant FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRole = "assistant"
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleTool      FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRole = "tool"
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleFunction  FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRole = "function"
	FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleModel     FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRole = "model"
)

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRole) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleSystem, FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleUser, FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleAssistant, FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleTool, FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleFunction, FunctionNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleModel:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt1Type string

const (
	FunctionNewParamsPromptDataPromptPromptDataPrompt1TypeChat FunctionNewParamsPromptDataPromptPromptDataPrompt1Type = "chat"
)

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt1Type) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptPromptDataPrompt1TypeChat:
		return true
	}
	return false
}

type FunctionNewParamsPromptDataPromptPromptDataPrompt2 struct {
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptPromptDataPrompt2) implementsFunctionNewParamsPromptDataPromptUnion() {
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
// [FunctionUpdateParamsFunctionDataPatchFunctionData1],
// [FunctionUpdateParamsFunctionDataPatchFunctionData2],
// [FunctionUpdateParamsFunctionDataPatchFunctionData3],
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

type FunctionUpdateParamsFunctionDataPatchFunctionData1 struct {
	Data param.Field[FunctionUpdateParamsFunctionDataPatchFunctionData1Data] `json:"data,required"`
	Type param.Field[FunctionUpdateParamsFunctionDataPatchFunctionData1Type] `json:"type,required"`
}

func (r FunctionUpdateParamsFunctionDataPatchFunctionData1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataPatchFunctionData1) implementsFunctionUpdateParamsFunctionDataUnion() {
}

type FunctionUpdateParamsFunctionDataPatchFunctionData1Data struct {
	BundleID       param.Field[string]                                                               `json:"bundle_id,required"`
	Location       param.Field[FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocation]       `json:"location,required"`
	RuntimeContext param.Field[FunctionUpdateParamsFunctionDataPatchFunctionData1DataRuntimeContext] `json:"runtime_context,required"`
}

func (r FunctionUpdateParamsFunctionDataPatchFunctionData1Data) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocation struct {
	EvalName param.Field[string]                                                                      `json:"eval_name,required"`
	Position param.Field[FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionUnion] `json:"position,required"`
	Type     param.Field[FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationType]          `json:"type,required"`
}

func (r FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionPatchFunctionDataProperties],
// [FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionScore].
type FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionUnion interface {
	implementsFunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionUnion()
}

type FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionPatchFunctionDataProperties string

const (
	FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionPatchFunctionDataPropertiesTask FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionPatchFunctionDataProperties = "task"
)

func (r FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionPatchFunctionDataProperties) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionPatchFunctionDataPropertiesTask:
		return true
	}
	return false
}

func (r FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionPatchFunctionDataProperties) implementsFunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionUnion() {
}

type FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionScore struct {
	Score param.Field[float64] `json:"score,required"`
}

func (r FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionScore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionScore) implementsFunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationPositionUnion() {
}

type FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationType string

const (
	FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationTypeExperiment FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationType = "experiment"
)

func (r FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataPatchFunctionData1DataLocationTypeExperiment:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataPatchFunctionData1DataRuntimeContext struct {
	Runtime param.Field[FunctionUpdateParamsFunctionDataPatchFunctionData1DataRuntimeContextRuntime] `json:"runtime,required"`
	Version param.Field[string]                                                                      `json:"version,required"`
}

func (r FunctionUpdateParamsFunctionDataPatchFunctionData1DataRuntimeContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsFunctionDataPatchFunctionData1DataRuntimeContextRuntime string

const (
	FunctionUpdateParamsFunctionDataPatchFunctionData1DataRuntimeContextRuntimeNode FunctionUpdateParamsFunctionDataPatchFunctionData1DataRuntimeContextRuntime = "node"
)

func (r FunctionUpdateParamsFunctionDataPatchFunctionData1DataRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataPatchFunctionData1DataRuntimeContextRuntimeNode:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataPatchFunctionData1Type string

const (
	FunctionUpdateParamsFunctionDataPatchFunctionData1TypeCode FunctionUpdateParamsFunctionDataPatchFunctionData1Type = "code"
)

func (r FunctionUpdateParamsFunctionDataPatchFunctionData1Type) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataPatchFunctionData1TypeCode:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataPatchFunctionData2 struct {
	Name param.Field[string]                                                 `json:"name,required"`
	Type param.Field[FunctionUpdateParamsFunctionDataPatchFunctionData2Type] `json:"type,required"`
}

func (r FunctionUpdateParamsFunctionDataPatchFunctionData2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataPatchFunctionData2) implementsFunctionUpdateParamsFunctionDataUnion() {
}

type FunctionUpdateParamsFunctionDataPatchFunctionData2Type string

const (
	FunctionUpdateParamsFunctionDataPatchFunctionData2TypeGlobal FunctionUpdateParamsFunctionDataPatchFunctionData2Type = "global"
)

func (r FunctionUpdateParamsFunctionDataPatchFunctionData2Type) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataPatchFunctionData2TypeGlobal:
		return true
	}
	return false
}

type FunctionUpdateParamsFunctionDataPatchFunctionData3 struct {
}

func (r FunctionUpdateParamsFunctionDataPatchFunctionData3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataPatchFunctionData3) implementsFunctionUpdateParamsFunctionDataUnion() {
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

// Satisfied by [FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0],
// [FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions1],
// [FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions2],
// [FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions3],
// [FunctionUpdateParamsPromptDataOptionsParamsJsCompletionParams],
// [FunctionUpdateParamsPromptDataOptionsParams].
type FunctionUpdateParamsPromptDataOptionsParamsUnion interface {
	implementsFunctionUpdateParamsPromptDataOptionsParamsUnion()
}

type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0 struct {
	FrequencyPenalty param.Field[float64]                                                                        `json:"frequency_penalty"`
	FunctionCall     param.Field[FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                                        `json:"max_tokens"`
	N                param.Field[float64]                                                                        `json:"n"`
	PresencePenalty  param.Field[float64]                                                                        `json:"presence_penalty"`
	ResponseFormat   param.Field[FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                                       `json:"stop"`
	Temperature      param.Field[float64]                                                                        `json:"temperature"`
	ToolChoice       param.Field[FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                                        `json:"top_p"`
	UseCache         param.Field[bool]                                                                           `json:"use_cache"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0) implementsFunctionUpdateParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by
// [FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0],
// [FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1],
// [FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject].
type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion interface {
	implementsFunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion()
}

type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 string

const (
	FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 = "auto"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto:
		return true
	}
	return false
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) implementsFunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 string

const (
	FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 = "none"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None:
		return true
	}
	return false
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) implementsFunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) implementsFunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat struct {
	Type param.Field[FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType] `json:"type,required"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType string

const (
	FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType = "json_object"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Satisfied by
// [FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0],
// [FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1],
// [FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2].
type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion interface {
	implementsFunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion()
}

type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 string

const (
	FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 = "auto"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto:
		return true
	}
	return false
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) implementsFunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 string

const (
	FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 = "none"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None:
		return true
	}
	return false
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) implementsFunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2 struct {
	Function param.Field[FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function] `json:"function,required"`
	Type     param.Field[FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type]     `json:"type,required"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) implementsFunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type string

const (
	FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type = "function"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions1 struct {
	MaxTokens   param.Field[float64] `json:"max_tokens,required"`
	Temperature param.Field[float64] `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64]  `json:"max_tokens_to_sample"`
	StopSequences     param.Field[[]string] `json:"stop_sequences"`
	TopK              param.Field[float64]  `json:"top_k"`
	TopP              param.Field[float64]  `json:"top_p"`
	UseCache          param.Field[bool]     `json:"use_cache"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions1) implementsFunctionUpdateParamsPromptDataOptionsParamsUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions2 struct {
	MaxOutputTokens param.Field[float64] `json:"maxOutputTokens"`
	Temperature     param.Field[float64] `json:"temperature"`
	TopK            param.Field[float64] `json:"topK"`
	TopP            param.Field[float64] `json:"topP"`
	UseCache        param.Field[bool]    `json:"use_cache"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions2) implementsFunctionUpdateParamsPromptDataOptionsParamsUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions3 struct {
	Temperature param.Field[float64] `json:"temperature"`
	TopK        param.Field[float64] `json:"topK"`
	UseCache    param.Field[bool]    `json:"use_cache"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions3) implementsFunctionUpdateParamsPromptDataOptionsParamsUnion() {
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

// Satisfied by [FunctionUpdateParamsPromptDataPromptPromptDataPrompt0],
// [FunctionUpdateParamsPromptDataPromptPromptDataPrompt1],
// [FunctionUpdateParamsPromptDataPromptPromptDataPrompt2],
// [FunctionUpdateParamsPromptDataPrompt].
type FunctionUpdateParamsPromptDataPromptUnion interface {
	implementsFunctionUpdateParamsPromptDataPromptUnion()
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt0 struct {
	Content param.Field[string]                                                    `json:"content,required"`
	Type    param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt0Type] `json:"type,required"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt0) implementsFunctionUpdateParamsPromptDataPromptUnion() {
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt0Type string

const (
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt0TypeCompletion FunctionUpdateParamsPromptDataPromptPromptDataPrompt0Type = "completion"
)

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt0Type) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptPromptDataPrompt0TypeCompletion:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1 struct {
	Messages param.Field[[]FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion] `json:"messages,required"`
	Type     param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1Type]           `json:"type,required"`
	Tools    param.Field[string]                                                              `json:"tools"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1) implementsFunctionUpdateParamsPromptDataPromptUnion() {
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1Message struct {
	Content      param.Field[interface{}]                                                       `json:"content,required"`
	Role         param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole] `json:"role,required"`
	Name         param.Field[string]                                                            `json:"name"`
	FunctionCall param.Field[interface{}]                                                       `json:"function_call,required"`
	ToolCalls    param.Field[interface{}]                                                       `json:"tool_calls,required"`
	ToolCallID   param.Field[string]                                                            `json:"tool_call_id"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1Message) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1Message) implementsFunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

// Satisfied by
// [FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0],
// [FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1],
// [FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2],
// [FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3],
// [FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4],
// [FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5],
// [FunctionUpdateParamsPromptDataPromptPromptDataPrompt1Message].
type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion interface {
	implementsFunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion()
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0 struct {
	Role    param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role] `json:"role,required"`
	Content param.Field[string]                                                                                    `json:"content"`
	Name    param.Field[string]                                                                                    `json:"name"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) implementsFunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role string

const (
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role = "system"
)

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1 struct {
	Role    param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role]         `json:"role,required"`
	Content param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion] `json:"content"`
	Name    param.Field[string]                                                                                            `json:"name"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) implementsFunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role string

const (
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role = "user"
)

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray].
type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsFunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion()
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray []FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) ImplementsFunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion() {
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray struct {
	Text     param.Field[string]                                                                                                `json:"text"`
	Type     param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                                                           `json:"image_url,required"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) implementsFunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

// Satisfied by
// [FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList],
// [FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray].
type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion interface {
	implementsFunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion()
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType] `json:"type,required"`
	Text param.Field[string]                                                                                                                              `json:"text"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsFunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL] `json:"image_url,required"`
	Type     param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType]     `json:"type,required"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsFunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    param.Field[string]                                                                                                                                     `json:"url,required"`
	Detail param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail] `json:"detail"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType string

const (
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText     FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "text"
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText, FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2 struct {
	Role         param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role]         `json:"role,required"`
	Content      param.Field[string]                                                                                            `json:"content"`
	FunctionCall param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                                                            `json:"name"`
	ToolCalls    param.Field[[]FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall]   `json:"tool_calls"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) implementsFunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role string

const (
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role = "assistant"
)

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall struct {
	ID       param.Field[string]                                                                                                 `json:"id,required"`
	Function param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction] `json:"function,required"`
	Type     param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType]     `json:"type,required"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType string

const (
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3 struct {
	Role       param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role] `json:"role,required"`
	Content    param.Field[string]                                                                                    `json:"content"`
	ToolCallID param.Field[string]                                                                                    `json:"tool_call_id"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) implementsFunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role string

const (
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role = "tool"
)

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4 struct {
	Name    param.Field[string]                                                                                    `json:"name,required"`
	Role    param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role] `json:"role,required"`
	Content param.Field[string]                                                                                    `json:"content"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) implementsFunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role string

const (
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role = "function"
)

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5 struct {
	Role    param.Field[FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role] `json:"role,required"`
	Content param.Field[string]                                                                                    `json:"content"`
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) implementsFunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role string

const (
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role = "model"
)

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole string

const (
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleSystem    FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole = "system"
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleUser      FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole = "user"
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleAssistant FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole = "assistant"
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleTool      FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole = "tool"
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleFunction  FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole = "function"
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleModel     FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole = "model"
)

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleSystem, FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleUser, FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleAssistant, FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleTool, FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleFunction, FunctionUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleModel:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt1Type string

const (
	FunctionUpdateParamsPromptDataPromptPromptDataPrompt1TypeChat FunctionUpdateParamsPromptDataPromptPromptDataPrompt1Type = "chat"
)

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt1Type) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptPromptDataPrompt1TypeChat:
		return true
	}
	return false
}

type FunctionUpdateParamsPromptDataPromptPromptDataPrompt2 struct {
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptPromptDataPrompt2) implementsFunctionUpdateParamsPromptDataPromptUnion() {
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
// [FunctionReplaceParamsFunctionDataFunctionData1],
// [FunctionReplaceParamsFunctionDataFunctionData2],
// [FunctionReplaceParamsFunctionData].
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

type FunctionReplaceParamsFunctionDataFunctionData1 struct {
	Data param.Field[FunctionReplaceParamsFunctionDataFunctionData1Data] `json:"data,required"`
	Type param.Field[FunctionReplaceParamsFunctionDataFunctionData1Type] `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataFunctionData1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataFunctionData1) implementsFunctionReplaceParamsFunctionDataUnion() {
}

type FunctionReplaceParamsFunctionDataFunctionData1Data struct {
	BundleID       param.Field[string]                                                           `json:"bundle_id,required"`
	Location       param.Field[FunctionReplaceParamsFunctionDataFunctionData1DataLocation]       `json:"location,required"`
	RuntimeContext param.Field[FunctionReplaceParamsFunctionDataFunctionData1DataRuntimeContext] `json:"runtime_context,required"`
}

func (r FunctionReplaceParamsFunctionDataFunctionData1Data) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsFunctionDataFunctionData1DataLocation struct {
	EvalName param.Field[string]                                                                  `json:"eval_name,required"`
	Position param.Field[FunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionUnion] `json:"position,required"`
	Type     param.Field[FunctionReplaceParamsFunctionDataFunctionData1DataLocationType]          `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataFunctionData1DataLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [FunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1],
// [FunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionScore].
type FunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionUnion interface {
	implementsFunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionUnion()
}

type FunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1 string

const (
	FunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1Task FunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1 = "task"
)

func (r FunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1Task:
		return true
	}
	return false
}

func (r FunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionFunctionDataProperties1) implementsFunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionUnion() {
}

type FunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionScore struct {
	Score param.Field[float64] `json:"score,required"`
}

func (r FunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionScore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionScore) implementsFunctionReplaceParamsFunctionDataFunctionData1DataLocationPositionUnion() {
}

type FunctionReplaceParamsFunctionDataFunctionData1DataLocationType string

const (
	FunctionReplaceParamsFunctionDataFunctionData1DataLocationTypeExperiment FunctionReplaceParamsFunctionDataFunctionData1DataLocationType = "experiment"
)

func (r FunctionReplaceParamsFunctionDataFunctionData1DataLocationType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataFunctionData1DataLocationTypeExperiment:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataFunctionData1DataRuntimeContext struct {
	Runtime param.Field[FunctionReplaceParamsFunctionDataFunctionData1DataRuntimeContextRuntime] `json:"runtime,required"`
	Version param.Field[string]                                                                  `json:"version,required"`
}

func (r FunctionReplaceParamsFunctionDataFunctionData1DataRuntimeContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsFunctionDataFunctionData1DataRuntimeContextRuntime string

const (
	FunctionReplaceParamsFunctionDataFunctionData1DataRuntimeContextRuntimeNode FunctionReplaceParamsFunctionDataFunctionData1DataRuntimeContextRuntime = "node"
)

func (r FunctionReplaceParamsFunctionDataFunctionData1DataRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataFunctionData1DataRuntimeContextRuntimeNode:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataFunctionData1Type string

const (
	FunctionReplaceParamsFunctionDataFunctionData1TypeCode FunctionReplaceParamsFunctionDataFunctionData1Type = "code"
)

func (r FunctionReplaceParamsFunctionDataFunctionData1Type) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataFunctionData1TypeCode:
		return true
	}
	return false
}

type FunctionReplaceParamsFunctionDataFunctionData2 struct {
	Name param.Field[string]                                             `json:"name,required"`
	Type param.Field[FunctionReplaceParamsFunctionDataFunctionData2Type] `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataFunctionData2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataFunctionData2) implementsFunctionReplaceParamsFunctionDataUnion() {
}

type FunctionReplaceParamsFunctionDataFunctionData2Type string

const (
	FunctionReplaceParamsFunctionDataFunctionData2TypeGlobal FunctionReplaceParamsFunctionDataFunctionData2Type = "global"
)

func (r FunctionReplaceParamsFunctionDataFunctionData2Type) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataFunctionData2TypeGlobal:
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

// Satisfied by [FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0],
// [FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions1],
// [FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions2],
// [FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions3],
// [FunctionReplaceParamsPromptDataOptionsParamsJsCompletionParams],
// [FunctionReplaceParamsPromptDataOptionsParams].
type FunctionReplaceParamsPromptDataOptionsParamsUnion interface {
	implementsFunctionReplaceParamsPromptDataOptionsParamsUnion()
}

type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0 struct {
	FrequencyPenalty param.Field[float64]                                                                         `json:"frequency_penalty"`
	FunctionCall     param.Field[FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                                         `json:"max_tokens"`
	N                param.Field[float64]                                                                         `json:"n"`
	PresencePenalty  param.Field[float64]                                                                         `json:"presence_penalty"`
	ResponseFormat   param.Field[FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                                        `json:"stop"`
	Temperature      param.Field[float64]                                                                         `json:"temperature"`
	ToolChoice       param.Field[FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                                         `json:"top_p"`
	UseCache         param.Field[bool]                                                                            `json:"use_cache"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0) implementsFunctionReplaceParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by
// [FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0],
// [FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1],
// [FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject].
type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion interface {
	implementsFunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion()
}

type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 string

const (
	FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 = "auto"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto:
		return true
	}
	return false
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) implementsFunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 string

const (
	FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 = "none"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None:
		return true
	}
	return false
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) implementsFunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) implementsFunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat struct {
	Type param.Field[FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType] `json:"type,required"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType string

const (
	FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType = "json_object"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Satisfied by
// [FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0],
// [FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1],
// [FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2].
type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion interface {
	implementsFunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion()
}

type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 string

const (
	FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 = "auto"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto:
		return true
	}
	return false
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) implementsFunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 string

const (
	FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 = "none"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None:
		return true
	}
	return false
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) implementsFunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2 struct {
	Function param.Field[FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function] `json:"function,required"`
	Type     param.Field[FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type]     `json:"type,required"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) implementsFunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type string

const (
	FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type = "function"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions1 struct {
	MaxTokens   param.Field[float64] `json:"max_tokens,required"`
	Temperature param.Field[float64] `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64]  `json:"max_tokens_to_sample"`
	StopSequences     param.Field[[]string] `json:"stop_sequences"`
	TopK              param.Field[float64]  `json:"top_k"`
	TopP              param.Field[float64]  `json:"top_p"`
	UseCache          param.Field[bool]     `json:"use_cache"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions1) implementsFunctionReplaceParamsPromptDataOptionsParamsUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions2 struct {
	MaxOutputTokens param.Field[float64] `json:"maxOutputTokens"`
	Temperature     param.Field[float64] `json:"temperature"`
	TopK            param.Field[float64] `json:"topK"`
	TopP            param.Field[float64] `json:"topP"`
	UseCache        param.Field[bool]    `json:"use_cache"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions2) implementsFunctionReplaceParamsPromptDataOptionsParamsUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions3 struct {
	Temperature param.Field[float64] `json:"temperature"`
	TopK        param.Field[float64] `json:"topK"`
	UseCache    param.Field[bool]    `json:"use_cache"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions3) implementsFunctionReplaceParamsPromptDataOptionsParamsUnion() {
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

// Satisfied by [FunctionReplaceParamsPromptDataPromptPromptDataPrompt0],
// [FunctionReplaceParamsPromptDataPromptPromptDataPrompt1],
// [FunctionReplaceParamsPromptDataPromptPromptDataPrompt2],
// [FunctionReplaceParamsPromptDataPrompt].
type FunctionReplaceParamsPromptDataPromptUnion interface {
	implementsFunctionReplaceParamsPromptDataPromptUnion()
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt0 struct {
	Content param.Field[string]                                                     `json:"content,required"`
	Type    param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt0Type] `json:"type,required"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt0) implementsFunctionReplaceParamsPromptDataPromptUnion() {
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt0Type string

const (
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt0TypeCompletion FunctionReplaceParamsPromptDataPromptPromptDataPrompt0Type = "completion"
)

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt0Type) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptPromptDataPrompt0TypeCompletion:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1 struct {
	Messages param.Field[[]FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion] `json:"messages,required"`
	Type     param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1Type]           `json:"type,required"`
	Tools    param.Field[string]                                                               `json:"tools"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1) implementsFunctionReplaceParamsPromptDataPromptUnion() {
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1Message struct {
	Content      param.Field[interface{}]                                                        `json:"content,required"`
	Role         param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole] `json:"role,required"`
	Name         param.Field[string]                                                             `json:"name"`
	FunctionCall param.Field[interface{}]                                                        `json:"function_call,required"`
	ToolCalls    param.Field[interface{}]                                                        `json:"tool_calls,required"`
	ToolCallID   param.Field[string]                                                             `json:"tool_call_id"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1Message) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1Message) implementsFunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

// Satisfied by
// [FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0],
// [FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1],
// [FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2],
// [FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3],
// [FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4],
// [FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5],
// [FunctionReplaceParamsPromptDataPromptPromptDataPrompt1Message].
type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion interface {
	implementsFunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion()
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0 struct {
	Role    param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role] `json:"role,required"`
	Content param.Field[string]                                                                                     `json:"content"`
	Name    param.Field[string]                                                                                     `json:"name"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) implementsFunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role string

const (
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role = "system"
)

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1 struct {
	Role    param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role]         `json:"role,required"`
	Content param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion] `json:"content"`
	Name    param.Field[string]                                                                                             `json:"name"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) implementsFunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role string

const (
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role = "user"
)

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray].
type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsFunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion()
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray []FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) ImplementsFunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion() {
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray struct {
	Text     param.Field[string]                                                                                                 `json:"text"`
	Type     param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                                                            `json:"image_url,required"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) implementsFunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

// Satisfied by
// [FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList],
// [FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray].
type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion interface {
	implementsFunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion()
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType] `json:"type,required"`
	Text param.Field[string]                                                                                                                               `json:"text"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsFunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL] `json:"image_url,required"`
	Type     param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType]     `json:"type,required"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsFunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    param.Field[string]                                                                                                                                      `json:"url,required"`
	Detail param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail] `json:"detail"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType string

const (
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText     FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "text"
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText, FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2 struct {
	Role         param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role]         `json:"role,required"`
	Content      param.Field[string]                                                                                             `json:"content"`
	FunctionCall param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                                                             `json:"name"`
	ToolCalls    param.Field[[]FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall]   `json:"tool_calls"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) implementsFunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role string

const (
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role = "assistant"
)

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall struct {
	ID       param.Field[string]                                                                                                  `json:"id,required"`
	Function param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction] `json:"function,required"`
	Type     param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType]     `json:"type,required"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType string

const (
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3 struct {
	Role       param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role] `json:"role,required"`
	Content    param.Field[string]                                                                                     `json:"content"`
	ToolCallID param.Field[string]                                                                                     `json:"tool_call_id"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) implementsFunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role string

const (
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role = "tool"
)

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4 struct {
	Name    param.Field[string]                                                                                     `json:"name,required"`
	Role    param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role] `json:"role,required"`
	Content param.Field[string]                                                                                     `json:"content"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) implementsFunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role string

const (
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role = "function"
)

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5 struct {
	Role    param.Field[FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role] `json:"role,required"`
	Content param.Field[string]                                                                                     `json:"content"`
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) implementsFunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role string

const (
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role = "model"
)

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole string

const (
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleSystem    FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole = "system"
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleUser      FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole = "user"
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleAssistant FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole = "assistant"
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleTool      FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole = "tool"
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleFunction  FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole = "function"
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleModel     FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole = "model"
)

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleSystem, FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleUser, FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleAssistant, FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleTool, FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleFunction, FunctionReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleModel:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt1Type string

const (
	FunctionReplaceParamsPromptDataPromptPromptDataPrompt1TypeChat FunctionReplaceParamsPromptDataPromptPromptDataPrompt1Type = "chat"
)

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt1Type) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptPromptDataPrompt1TypeChat:
		return true
	}
	return false
}

type FunctionReplaceParamsPromptDataPromptPromptDataPrompt2 struct {
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptPromptDataPrompt2) implementsFunctionReplaceParamsPromptDataPromptUnion() {
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
