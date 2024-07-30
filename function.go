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

// Log feedback for a set of function events
func (r *FunctionService) Feedback(ctx context.Context, functionID string, body FunctionFeedbackParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if functionID == "" {
		err = errors.New("missing required function_id parameter")
		return
	}
	path := fmt.Sprintf("v1/function/%s/feedback", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
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
	// This field can have the runtime type of [FunctionFunctionDataObjectData].
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
// Possible runtime types of the union are [FunctionFunctionDataObject],
// [FunctionFunctionDataObject], [FunctionFunctionDataObject].
func (r FunctionFunctionData) AsUnion() FunctionFunctionDataUnion {
	return r.union
}

// Union satisfied by [FunctionFunctionDataObject], [FunctionFunctionDataObject] or
// [FunctionFunctionDataObject].
type FunctionFunctionDataUnion interface {
	implementsFunctionFunctionData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionFunctionDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataObject{}),
		},
	)
}

type FunctionFunctionDataObject struct {
	Type FunctionFunctionDataObjectType `json:"type,required"`
	JSON functionFunctionDataObjectJSON `json:"-"`
}

// functionFunctionDataObjectJSON contains the JSON metadata for the struct
// [FunctionFunctionDataObject]
type functionFunctionDataObjectJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataObjectJSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataObject) implementsFunctionFunctionData() {}

type FunctionFunctionDataObjectType string

const (
	FunctionFunctionDataObjectTypePrompt FunctionFunctionDataObjectType = "prompt"
)

func (r FunctionFunctionDataObjectType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataObjectTypePrompt:
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
	// [FunctionPromptDataOptionsParamsObjectResponseFormat].
	ResponseFormat interface{} `json:"response_format,required"`
	// This field can have the runtime type of
	// [FunctionPromptDataOptionsParamsObjectToolChoiceUnion].
	ToolChoice interface{} `json:"tool_choice,required"`
	// This field can have the runtime type of
	// [FunctionPromptDataOptionsParamsObjectFunctionCallUnion].
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
// Possible runtime types of the union are [FunctionPromptDataOptionsParamsObject],
// [FunctionPromptDataOptionsParamsObject],
// [FunctionPromptDataOptionsParamsObject],
// [FunctionPromptDataOptionsParamsObject],
// [FunctionPromptDataOptionsParamsObject].
func (r FunctionPromptDataOptionsParams) AsUnion() FunctionPromptDataOptionsParamsUnion {
	return r.union
}

// Union satisfied by [FunctionPromptDataOptionsParamsObject],
// [FunctionPromptDataOptionsParamsObject],
// [FunctionPromptDataOptionsParamsObject], [FunctionPromptDataOptionsParamsObject]
// or [FunctionPromptDataOptionsParamsObject].
type FunctionPromptDataOptionsParamsUnion interface {
	implementsFunctionPromptDataOptionsParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataOptionsParamsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsObject{}),
		},
	)
}

type FunctionPromptDataOptionsParamsObject struct {
	FrequencyPenalty float64                                                `json:"frequency_penalty"`
	FunctionCall     FunctionPromptDataOptionsParamsObjectFunctionCallUnion `json:"function_call"`
	MaxTokens        float64                                                `json:"max_tokens"`
	N                float64                                                `json:"n"`
	PresencePenalty  float64                                                `json:"presence_penalty"`
	ResponseFormat   FunctionPromptDataOptionsParamsObjectResponseFormat    `json:"response_format,nullable"`
	Stop             []string                                               `json:"stop"`
	Temperature      float64                                                `json:"temperature"`
	ToolChoice       FunctionPromptDataOptionsParamsObjectToolChoiceUnion   `json:"tool_choice"`
	TopP             float64                                                `json:"top_p"`
	UseCache         bool                                                   `json:"use_cache"`
	JSON             functionPromptDataOptionsParamsObjectJSON              `json:"-"`
}

// functionPromptDataOptionsParamsObjectJSON contains the JSON metadata for the
// struct [FunctionPromptDataOptionsParamsObject]
type functionPromptDataOptionsParamsObjectJSON struct {
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

func (r *FunctionPromptDataOptionsParamsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsObjectJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsObject) implementsFunctionPromptDataOptionsParams() {}

// Union satisfied by [FunctionPromptDataOptionsParamsObjectFunctionCallString],
// [FunctionPromptDataOptionsParamsObjectFunctionCallString] or
// [FunctionPromptDataOptionsParamsObjectFunctionCallObject].
type FunctionPromptDataOptionsParamsObjectFunctionCallUnion interface {
	implementsFunctionPromptDataOptionsParamsObjectFunctionCallUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataOptionsParamsObjectFunctionCallUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsObjectFunctionCallObject{}),
		},
	)
}

type FunctionPromptDataOptionsParamsObjectFunctionCallString string

const (
	FunctionPromptDataOptionsParamsObjectFunctionCallStringAuto FunctionPromptDataOptionsParamsObjectFunctionCallString = "auto"
)

func (r FunctionPromptDataOptionsParamsObjectFunctionCallString) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsObjectFunctionCallStringAuto:
		return true
	}
	return false
}

func (r FunctionPromptDataOptionsParamsObjectFunctionCallString) implementsFunctionPromptDataOptionsParamsObjectFunctionCallUnion() {
}

type FunctionPromptDataOptionsParamsObjectFunctionCallObject struct {
	Name string                                                      `json:"name,required"`
	JSON functionPromptDataOptionsParamsObjectFunctionCallObjectJSON `json:"-"`
}

// functionPromptDataOptionsParamsObjectFunctionCallObjectJSON contains the JSON
// metadata for the struct
// [FunctionPromptDataOptionsParamsObjectFunctionCallObject]
type functionPromptDataOptionsParamsObjectFunctionCallObjectJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsObjectFunctionCallObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsObjectFunctionCallObjectJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsObjectFunctionCallObject) implementsFunctionPromptDataOptionsParamsObjectFunctionCallUnion() {
}

type FunctionPromptDataOptionsParamsObjectResponseFormat struct {
	Type FunctionPromptDataOptionsParamsObjectResponseFormatType `json:"type,required"`
	JSON functionPromptDataOptionsParamsObjectResponseFormatJSON `json:"-"`
}

// functionPromptDataOptionsParamsObjectResponseFormatJSON contains the JSON
// metadata for the struct [FunctionPromptDataOptionsParamsObjectResponseFormat]
type functionPromptDataOptionsParamsObjectResponseFormatJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsObjectResponseFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsObjectResponseFormatJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataOptionsParamsObjectResponseFormatType string

const (
	FunctionPromptDataOptionsParamsObjectResponseFormatTypeJsonObject FunctionPromptDataOptionsParamsObjectResponseFormatType = "json_object"
)

func (r FunctionPromptDataOptionsParamsObjectResponseFormatType) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsObjectResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Union satisfied by [FunctionPromptDataOptionsParamsObjectToolChoiceString],
// [FunctionPromptDataOptionsParamsObjectToolChoiceString] or
// [FunctionPromptDataOptionsParamsObjectToolChoiceObject].
type FunctionPromptDataOptionsParamsObjectToolChoiceUnion interface {
	implementsFunctionPromptDataOptionsParamsObjectToolChoiceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataOptionsParamsObjectToolChoiceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataOptionsParamsObjectToolChoiceObject{}),
		},
	)
}

type FunctionPromptDataOptionsParamsObjectToolChoiceString string

const (
	FunctionPromptDataOptionsParamsObjectToolChoiceStringAuto FunctionPromptDataOptionsParamsObjectToolChoiceString = "auto"
)

func (r FunctionPromptDataOptionsParamsObjectToolChoiceString) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsObjectToolChoiceStringAuto:
		return true
	}
	return false
}

func (r FunctionPromptDataOptionsParamsObjectToolChoiceString) implementsFunctionPromptDataOptionsParamsObjectToolChoiceUnion() {
}

type FunctionPromptDataOptionsParamsObjectToolChoiceObject struct {
	Function FunctionPromptDataOptionsParamsObjectToolChoiceObjectFunction `json:"function,required"`
	Type     FunctionPromptDataOptionsParamsObjectToolChoiceObjectType     `json:"type,required"`
	JSON     functionPromptDataOptionsParamsObjectToolChoiceObjectJSON     `json:"-"`
}

// functionPromptDataOptionsParamsObjectToolChoiceObjectJSON contains the JSON
// metadata for the struct [FunctionPromptDataOptionsParamsObjectToolChoiceObject]
type functionPromptDataOptionsParamsObjectToolChoiceObjectJSON struct {
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsObjectToolChoiceObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsObjectToolChoiceObjectJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataOptionsParamsObjectToolChoiceObject) implementsFunctionPromptDataOptionsParamsObjectToolChoiceUnion() {
}

type FunctionPromptDataOptionsParamsObjectToolChoiceObjectFunction struct {
	Name string                                                            `json:"name,required"`
	JSON functionPromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON `json:"-"`
}

// functionPromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON contains the
// JSON metadata for the struct
// [FunctionPromptDataOptionsParamsObjectToolChoiceObjectFunction]
type functionPromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataOptionsParamsObjectToolChoiceObjectFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON) RawJSON() string {
	return r.raw
}

type FunctionPromptDataOptionsParamsObjectToolChoiceObjectType string

const (
	FunctionPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction FunctionPromptDataOptionsParamsObjectToolChoiceObjectType = "function"
)

func (r FunctionPromptDataOptionsParamsObjectToolChoiceObjectType) IsKnown() bool {
	switch r {
	case FunctionPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction:
		return true
	}
	return false
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
// [FunctionPromptDataPromptChat], [FunctionPromptDataPromptNullVariant].
func (r FunctionPromptDataPrompt) AsUnion() FunctionPromptDataPromptUnion {
	return r.union
}

// Union satisfied by [FunctionPromptDataPromptCompletion],
// [FunctionPromptDataPromptChat] or [FunctionPromptDataPromptNullVariant].
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
			Type:       reflect.TypeOf(FunctionPromptDataPromptNullVariant{}),
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
	// [FunctionPromptDataPromptChatMessagesObjectContentUnion].
	Content interface{}                              `json:"content,required"`
	Role    FunctionPromptDataPromptChatMessagesRole `json:"role,required"`
	Name    string                                   `json:"name"`
	// This field can have the runtime type of
	// [FunctionPromptDataPromptChatMessagesObjectFunctionCall].
	FunctionCall interface{} `json:"function_call,required"`
	// This field can have the runtime type of
	// [[]FunctionPromptDataPromptChatMessagesObjectToolCall].
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
// [FunctionPromptDataPromptChatMessagesObject],
// [FunctionPromptDataPromptChatMessagesObject],
// [FunctionPromptDataPromptChatMessagesObject],
// [FunctionPromptDataPromptChatMessagesObject],
// [FunctionPromptDataPromptChatMessagesObject],
// [FunctionPromptDataPromptChatMessagesObject].
func (r FunctionPromptDataPromptChatMessage) AsUnion() FunctionPromptDataPromptChatMessagesUnion {
	return r.union
}

// Union satisfied by [FunctionPromptDataPromptChatMessagesObject],
// [FunctionPromptDataPromptChatMessagesObject],
// [FunctionPromptDataPromptChatMessagesObject],
// [FunctionPromptDataPromptChatMessagesObject],
// [FunctionPromptDataPromptChatMessagesObject] or
// [FunctionPromptDataPromptChatMessagesObject].
type FunctionPromptDataPromptChatMessagesUnion interface {
	implementsFunctionPromptDataPromptChatMessage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionPromptDataPromptChatMessagesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionPromptDataPromptChatMessagesObject{}),
		},
	)
}

type FunctionPromptDataPromptChatMessagesObject struct {
	Role    FunctionPromptDataPromptChatMessagesObjectRole `json:"role,required"`
	Content string                                         `json:"content"`
	Name    string                                         `json:"name"`
	JSON    functionPromptDataPromptChatMessagesObjectJSON `json:"-"`
}

// functionPromptDataPromptChatMessagesObjectJSON contains the JSON metadata for
// the struct [FunctionPromptDataPromptChatMessagesObject]
type functionPromptDataPromptChatMessagesObjectJSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptChatMessagesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptChatMessagesObjectJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptChatMessagesObject) implementsFunctionPromptDataPromptChatMessage() {}

type FunctionPromptDataPromptChatMessagesObjectRole string

const (
	FunctionPromptDataPromptChatMessagesObjectRoleSystem FunctionPromptDataPromptChatMessagesObjectRole = "system"
)

func (r FunctionPromptDataPromptChatMessagesObjectRole) IsKnown() bool {
	switch r {
	case FunctionPromptDataPromptChatMessagesObjectRoleSystem:
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

type FunctionPromptDataPromptNullVariant struct {
	JSON functionPromptDataPromptNullVariantJSON `json:"-"`
}

// functionPromptDataPromptNullVariantJSON contains the JSON metadata for the
// struct [FunctionPromptDataPromptNullVariant]
type functionPromptDataPromptNullVariantJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionPromptDataPromptNullVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionPromptDataPromptNullVariantJSON) RawJSON() string {
	return r.raw
}

func (r FunctionPromptDataPromptNullVariant) implementsFunctionPromptDataPrompt() {}

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

// Satisfied by [FunctionNewParamsFunctionDataObject],
// [FunctionNewParamsFunctionDataObject], [FunctionNewParamsFunctionDataObject],
// [FunctionNewParamsFunctionData].
type FunctionNewParamsFunctionDataUnion interface {
	implementsFunctionNewParamsFunctionDataUnion()
}

type FunctionNewParamsFunctionDataObject struct {
	Type param.Field[FunctionNewParamsFunctionDataObjectType] `json:"type,required"`
}

func (r FunctionNewParamsFunctionDataObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsFunctionDataObject) implementsFunctionNewParamsFunctionDataUnion() {}

type FunctionNewParamsFunctionDataObjectType string

const (
	FunctionNewParamsFunctionDataObjectTypePrompt FunctionNewParamsFunctionDataObjectType = "prompt"
)

func (r FunctionNewParamsFunctionDataObjectType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionDataObjectTypePrompt:
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

// Satisfied by [FunctionNewParamsPromptDataOptionsParamsObject],
// [FunctionNewParamsPromptDataOptionsParamsObject],
// [FunctionNewParamsPromptDataOptionsParamsObject],
// [FunctionNewParamsPromptDataOptionsParamsObject],
// [FunctionNewParamsPromptDataOptionsParamsObject],
// [FunctionNewParamsPromptDataOptionsParams].
type FunctionNewParamsPromptDataOptionsParamsUnion interface {
	implementsFunctionNewParamsPromptDataOptionsParamsUnion()
}

type FunctionNewParamsPromptDataOptionsParamsObject struct {
	FrequencyPenalty param.Field[float64]                                                         `json:"frequency_penalty"`
	FunctionCall     param.Field[FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                         `json:"max_tokens"`
	N                param.Field[float64]                                                         `json:"n"`
	PresencePenalty  param.Field[float64]                                                         `json:"presence_penalty"`
	ResponseFormat   param.Field[FunctionNewParamsPromptDataOptionsParamsObjectResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                        `json:"stop"`
	Temperature      param.Field[float64]                                                         `json:"temperature"`
	ToolChoice       param.Field[FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                         `json:"top_p"`
	UseCache         param.Field[bool]                                                            `json:"use_cache"`
}

func (r FunctionNewParamsPromptDataOptionsParamsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsObject) implementsFunctionNewParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by [FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallString],
// [FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallString],
// [FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallObject].
type FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallUnion interface {
	implementsFunctionNewParamsPromptDataOptionsParamsObjectFunctionCallUnion()
}

type FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallString string

const (
	FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallStringAuto FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallString = "auto"
)

func (r FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallString) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallStringAuto:
		return true
	}
	return false
}

func (r FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallString) implementsFunctionNewParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallObject) implementsFunctionNewParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsObjectResponseFormat struct {
	Type param.Field[FunctionNewParamsPromptDataOptionsParamsObjectResponseFormatType] `json:"type,required"`
}

func (r FunctionNewParamsPromptDataOptionsParamsObjectResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataOptionsParamsObjectResponseFormatType string

const (
	FunctionNewParamsPromptDataOptionsParamsObjectResponseFormatTypeJsonObject FunctionNewParamsPromptDataOptionsParamsObjectResponseFormatType = "json_object"
)

func (r FunctionNewParamsPromptDataOptionsParamsObjectResponseFormatType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsObjectResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Satisfied by [FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceString],
// [FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceString],
// [FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceObject].
type FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceUnion interface {
	implementsFunctionNewParamsPromptDataOptionsParamsObjectToolChoiceUnion()
}

type FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceString string

const (
	FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceStringAuto FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceString = "auto"
)

func (r FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceString) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceStringAuto:
		return true
	}
	return false
}

func (r FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceString) implementsFunctionNewParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceObject struct {
	Function param.Field[FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction] `json:"function,required"`
	Type     param.Field[FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceObjectType]     `json:"type,required"`
}

func (r FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceObject) implementsFunctionNewParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
}

type FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceObjectType string

const (
	FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceObjectType = "function"
)

func (r FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceObjectType) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction:
		return true
	}
	return false
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
// [FunctionNewParamsPromptDataPromptNullVariant],
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

// Satisfied by [FunctionNewParamsPromptDataPromptChatMessagesObject],
// [FunctionNewParamsPromptDataPromptChatMessagesObject],
// [FunctionNewParamsPromptDataPromptChatMessagesObject],
// [FunctionNewParamsPromptDataPromptChatMessagesObject],
// [FunctionNewParamsPromptDataPromptChatMessagesObject],
// [FunctionNewParamsPromptDataPromptChatMessagesObject],
// [FunctionNewParamsPromptDataPromptChatMessage].
type FunctionNewParamsPromptDataPromptChatMessageUnion interface {
	implementsFunctionNewParamsPromptDataPromptChatMessageUnion()
}

type FunctionNewParamsPromptDataPromptChatMessagesObject struct {
	Role    param.Field[FunctionNewParamsPromptDataPromptChatMessagesObjectRole] `json:"role,required"`
	Content param.Field[string]                                                  `json:"content"`
	Name    param.Field[string]                                                  `json:"name"`
}

func (r FunctionNewParamsPromptDataPromptChatMessagesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptChatMessagesObject) implementsFunctionNewParamsPromptDataPromptChatMessageUnion() {
}

type FunctionNewParamsPromptDataPromptChatMessagesObjectRole string

const (
	FunctionNewParamsPromptDataPromptChatMessagesObjectRoleSystem FunctionNewParamsPromptDataPromptChatMessagesObjectRole = "system"
)

func (r FunctionNewParamsPromptDataPromptChatMessagesObjectRole) IsKnown() bool {
	switch r {
	case FunctionNewParamsPromptDataPromptChatMessagesObjectRoleSystem:
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

type FunctionNewParamsPromptDataPromptNullVariant struct {
}

func (r FunctionNewParamsPromptDataPromptNullVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionNewParamsPromptDataPromptNullVariant) implementsFunctionNewParamsPromptDataPromptUnion() {
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

// Satisfied by [FunctionUpdateParamsFunctionDataObject],
// [FunctionUpdateParamsFunctionDataObject],
// [FunctionUpdateParamsFunctionDataObject],
// [FunctionUpdateParamsFunctionDataObject], [FunctionUpdateParamsFunctionData].
type FunctionUpdateParamsFunctionDataUnion interface {
	implementsFunctionUpdateParamsFunctionDataUnion()
}

type FunctionUpdateParamsFunctionDataObject struct {
	Type param.Field[FunctionUpdateParamsFunctionDataObjectType] `json:"type,required"`
}

func (r FunctionUpdateParamsFunctionDataObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsFunctionDataObject) implementsFunctionUpdateParamsFunctionDataUnion() {}

type FunctionUpdateParamsFunctionDataObjectType string

const (
	FunctionUpdateParamsFunctionDataObjectTypePrompt FunctionUpdateParamsFunctionDataObjectType = "prompt"
)

func (r FunctionUpdateParamsFunctionDataObjectType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsFunctionDataObjectTypePrompt:
		return true
	}
	return false
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

// Satisfied by [FunctionUpdateParamsPromptDataOptionsParamsObject],
// [FunctionUpdateParamsPromptDataOptionsParamsObject],
// [FunctionUpdateParamsPromptDataOptionsParamsObject],
// [FunctionUpdateParamsPromptDataOptionsParamsObject],
// [FunctionUpdateParamsPromptDataOptionsParamsObject],
// [FunctionUpdateParamsPromptDataOptionsParams].
type FunctionUpdateParamsPromptDataOptionsParamsUnion interface {
	implementsFunctionUpdateParamsPromptDataOptionsParamsUnion()
}

type FunctionUpdateParamsPromptDataOptionsParamsObject struct {
	FrequencyPenalty param.Field[float64]                                                            `json:"frequency_penalty"`
	FunctionCall     param.Field[FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                            `json:"max_tokens"`
	N                param.Field[float64]                                                            `json:"n"`
	PresencePenalty  param.Field[float64]                                                            `json:"presence_penalty"`
	ResponseFormat   param.Field[FunctionUpdateParamsPromptDataOptionsParamsObjectResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                           `json:"stop"`
	Temperature      param.Field[float64]                                                            `json:"temperature"`
	ToolChoice       param.Field[FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                            `json:"top_p"`
	UseCache         param.Field[bool]                                                               `json:"use_cache"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsObject) implementsFunctionUpdateParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by
// [FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallString],
// [FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallString],
// [FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallObject].
type FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallUnion interface {
	implementsFunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallUnion()
}

type FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallString string

const (
	FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallStringAuto FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallString = "auto"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallString) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallStringAuto:
		return true
	}
	return false
}

func (r FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallString) implementsFunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallObject) implementsFunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsObjectResponseFormat struct {
	Type param.Field[FunctionUpdateParamsPromptDataOptionsParamsObjectResponseFormatType] `json:"type,required"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsObjectResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataOptionsParamsObjectResponseFormatType string

const (
	FunctionUpdateParamsPromptDataOptionsParamsObjectResponseFormatTypeJsonObject FunctionUpdateParamsPromptDataOptionsParamsObjectResponseFormatType = "json_object"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsObjectResponseFormatType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsObjectResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Satisfied by
// [FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceString],
// [FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceString],
// [FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceObject].
type FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceUnion interface {
	implementsFunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceUnion()
}

type FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceString string

const (
	FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceStringAuto FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceString = "auto"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceString) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceStringAuto:
		return true
	}
	return false
}

func (r FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceString) implementsFunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceObject struct {
	Function param.Field[FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction] `json:"function,required"`
	Type     param.Field[FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectType]     `json:"type,required"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceObject) implementsFunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
}

type FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectType string

const (
	FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectType = "function"
)

func (r FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectType) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction:
		return true
	}
	return false
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
// [FunctionUpdateParamsPromptDataPromptNullVariant],
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

// Satisfied by [FunctionUpdateParamsPromptDataPromptChatMessagesObject],
// [FunctionUpdateParamsPromptDataPromptChatMessagesObject],
// [FunctionUpdateParamsPromptDataPromptChatMessagesObject],
// [FunctionUpdateParamsPromptDataPromptChatMessagesObject],
// [FunctionUpdateParamsPromptDataPromptChatMessagesObject],
// [FunctionUpdateParamsPromptDataPromptChatMessagesObject],
// [FunctionUpdateParamsPromptDataPromptChatMessage].
type FunctionUpdateParamsPromptDataPromptChatMessageUnion interface {
	implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion()
}

type FunctionUpdateParamsPromptDataPromptChatMessagesObject struct {
	Role    param.Field[FunctionUpdateParamsPromptDataPromptChatMessagesObjectRole] `json:"role,required"`
	Content param.Field[string]                                                     `json:"content"`
	Name    param.Field[string]                                                     `json:"name"`
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptChatMessagesObject) implementsFunctionUpdateParamsPromptDataPromptChatMessageUnion() {
}

type FunctionUpdateParamsPromptDataPromptChatMessagesObjectRole string

const (
	FunctionUpdateParamsPromptDataPromptChatMessagesObjectRoleSystem FunctionUpdateParamsPromptDataPromptChatMessagesObjectRole = "system"
)

func (r FunctionUpdateParamsPromptDataPromptChatMessagesObjectRole) IsKnown() bool {
	switch r {
	case FunctionUpdateParamsPromptDataPromptChatMessagesObjectRoleSystem:
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

type FunctionUpdateParamsPromptDataPromptNullVariant struct {
}

func (r FunctionUpdateParamsPromptDataPromptNullVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionUpdateParamsPromptDataPromptNullVariant) implementsFunctionUpdateParamsPromptDataPromptUnion() {
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

type FunctionFeedbackParams struct {
	// A list of function feedback items
	Feedback param.Field[[]FunctionFeedbackParamsFeedback] `json:"feedback,required"`
}

func (r FunctionFeedbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionFeedbackParamsFeedback struct {
	// The id of the function event to log feedback for. This is the row `id` returned
	// by `POST /v1/function/{function_id}/insert`
	ID param.Field[string] `json:"id,required"`
	// An optional comment string to log about the function event
	Comment param.Field[string] `json:"comment"`
	// A dictionary with additional data about the feedback. If you have a `user_id`,
	// you can log it here and access it in the Braintrust UI.
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// The source of the feedback. Must be one of "external" (default), "app", or "api"
	Source param.Field[FunctionFeedbackParamsFeedbackSource] `json:"source"`
}

func (r FunctionFeedbackParamsFeedback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The source of the feedback. Must be one of "external" (default), "app", or "api"
type FunctionFeedbackParamsFeedbackSource string

const (
	FunctionFeedbackParamsFeedbackSourceApp      FunctionFeedbackParamsFeedbackSource = "app"
	FunctionFeedbackParamsFeedbackSourceAPI      FunctionFeedbackParamsFeedbackSource = "api"
	FunctionFeedbackParamsFeedbackSourceExternal FunctionFeedbackParamsFeedbackSource = "external"
)

func (r FunctionFeedbackParamsFeedbackSource) IsKnown() bool {
	switch r {
	case FunctionFeedbackParamsFeedbackSourceApp, FunctionFeedbackParamsFeedbackSourceAPI, FunctionFeedbackParamsFeedbackSourceExternal:
		return true
	}
	return false
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

// Satisfied by [FunctionReplaceParamsFunctionDataObject],
// [FunctionReplaceParamsFunctionDataObject],
// [FunctionReplaceParamsFunctionDataObject], [FunctionReplaceParamsFunctionData].
type FunctionReplaceParamsFunctionDataUnion interface {
	implementsFunctionReplaceParamsFunctionDataUnion()
}

type FunctionReplaceParamsFunctionDataObject struct {
	Type param.Field[FunctionReplaceParamsFunctionDataObjectType] `json:"type,required"`
}

func (r FunctionReplaceParamsFunctionDataObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsFunctionDataObject) implementsFunctionReplaceParamsFunctionDataUnion() {}

type FunctionReplaceParamsFunctionDataObjectType string

const (
	FunctionReplaceParamsFunctionDataObjectTypePrompt FunctionReplaceParamsFunctionDataObjectType = "prompt"
)

func (r FunctionReplaceParamsFunctionDataObjectType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsFunctionDataObjectTypePrompt:
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

// Satisfied by [FunctionReplaceParamsPromptDataOptionsParamsObject],
// [FunctionReplaceParamsPromptDataOptionsParamsObject],
// [FunctionReplaceParamsPromptDataOptionsParamsObject],
// [FunctionReplaceParamsPromptDataOptionsParamsObject],
// [FunctionReplaceParamsPromptDataOptionsParamsObject],
// [FunctionReplaceParamsPromptDataOptionsParams].
type FunctionReplaceParamsPromptDataOptionsParamsUnion interface {
	implementsFunctionReplaceParamsPromptDataOptionsParamsUnion()
}

type FunctionReplaceParamsPromptDataOptionsParamsObject struct {
	FrequencyPenalty param.Field[float64]                                                             `json:"frequency_penalty"`
	FunctionCall     param.Field[FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                             `json:"max_tokens"`
	N                param.Field[float64]                                                             `json:"n"`
	PresencePenalty  param.Field[float64]                                                             `json:"presence_penalty"`
	ResponseFormat   param.Field[FunctionReplaceParamsPromptDataOptionsParamsObjectResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                            `json:"stop"`
	Temperature      param.Field[float64]                                                             `json:"temperature"`
	ToolChoice       param.Field[FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                             `json:"top_p"`
	UseCache         param.Field[bool]                                                                `json:"use_cache"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsObject) implementsFunctionReplaceParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by
// [FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallString],
// [FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallString],
// [FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallObject].
type FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallUnion interface {
	implementsFunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallUnion()
}

type FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallString string

const (
	FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallStringAuto FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallString = "auto"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallString) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallStringAuto:
		return true
	}
	return false
}

func (r FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallString) implementsFunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallObject) implementsFunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsObjectResponseFormat struct {
	Type param.Field[FunctionReplaceParamsPromptDataOptionsParamsObjectResponseFormatType] `json:"type,required"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsObjectResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataOptionsParamsObjectResponseFormatType string

const (
	FunctionReplaceParamsPromptDataOptionsParamsObjectResponseFormatTypeJsonObject FunctionReplaceParamsPromptDataOptionsParamsObjectResponseFormatType = "json_object"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsObjectResponseFormatType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsObjectResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Satisfied by
// [FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceString],
// [FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceString],
// [FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceObject].
type FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceUnion interface {
	implementsFunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceUnion()
}

type FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceString string

const (
	FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceStringAuto FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceString = "auto"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceString) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceStringAuto:
		return true
	}
	return false
}

func (r FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceString) implementsFunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceObject struct {
	Function param.Field[FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction] `json:"function,required"`
	Type     param.Field[FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectType]     `json:"type,required"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceObject) implementsFunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
}

type FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectType string

const (
	FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectType = "function"
)

func (r FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectType) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction:
		return true
	}
	return false
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
// [FunctionReplaceParamsPromptDataPromptNullVariant],
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

// Satisfied by [FunctionReplaceParamsPromptDataPromptChatMessagesObject],
// [FunctionReplaceParamsPromptDataPromptChatMessagesObject],
// [FunctionReplaceParamsPromptDataPromptChatMessagesObject],
// [FunctionReplaceParamsPromptDataPromptChatMessagesObject],
// [FunctionReplaceParamsPromptDataPromptChatMessagesObject],
// [FunctionReplaceParamsPromptDataPromptChatMessagesObject],
// [FunctionReplaceParamsPromptDataPromptChatMessage].
type FunctionReplaceParamsPromptDataPromptChatMessageUnion interface {
	implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion()
}

type FunctionReplaceParamsPromptDataPromptChatMessagesObject struct {
	Role    param.Field[FunctionReplaceParamsPromptDataPromptChatMessagesObjectRole] `json:"role,required"`
	Content param.Field[string]                                                      `json:"content"`
	Name    param.Field[string]                                                      `json:"name"`
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptChatMessagesObject) implementsFunctionReplaceParamsPromptDataPromptChatMessageUnion() {
}

type FunctionReplaceParamsPromptDataPromptChatMessagesObjectRole string

const (
	FunctionReplaceParamsPromptDataPromptChatMessagesObjectRoleSystem FunctionReplaceParamsPromptDataPromptChatMessagesObjectRole = "system"
)

func (r FunctionReplaceParamsPromptDataPromptChatMessagesObjectRole) IsKnown() bool {
	switch r {
	case FunctionReplaceParamsPromptDataPromptChatMessagesObjectRoleSystem:
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

type FunctionReplaceParamsPromptDataPromptNullVariant struct {
}

func (r FunctionReplaceParamsPromptDataPromptNullVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FunctionReplaceParamsPromptDataPromptNullVariant) implementsFunctionReplaceParamsPromptDataPromptUnion() {
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
