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

// PromptService contains methods and other services that help with interacting
// with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPromptService] method instead.
type PromptService struct {
	Options []option.RequestOption
}

// NewPromptService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPromptService(opts ...option.RequestOption) (r *PromptService) {
	r = &PromptService{}
	r.Options = opts
	return
}

// Create a new prompt. If there is an existing prompt in the project with the same
// slug as the one specified in the request, will return the existing prompt
// unmodified
func (r *PromptService) New(ctx context.Context, body PromptNewParams, opts ...option.RequestOption) (res *Prompt, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/prompt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a prompt object by its id
func (r *PromptService) Get(ctx context.Context, promptID string, opts ...option.RequestOption) (res *Prompt, err error) {
	opts = append(r.Options[:], opts...)
	if promptID == "" {
		err = errors.New("missing required prompt_id parameter")
		return
	}
	path := fmt.Sprintf("v1/prompt/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update a prompt object. Specify the fields to update in the payload.
// Any object-type fields will be deep-merged with existing content. Currently we
// do not support removing fields or setting them to null.
func (r *PromptService) Update(ctx context.Context, promptID string, body PromptUpdateParams, opts ...option.RequestOption) (res *Prompt, err error) {
	opts = append(r.Options[:], opts...)
	if promptID == "" {
		err = errors.New("missing required prompt_id parameter")
		return
	}
	path := fmt.Sprintf("v1/prompt/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all prompts. The prompts are sorted by creation date, with the most
// recently-created prompts coming first
func (r *PromptService) List(ctx context.Context, query PromptListParams, opts ...option.RequestOption) (res *pagination.ListObjects[Prompt], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/prompt"
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

// List out all prompts. The prompts are sorted by creation date, with the most
// recently-created prompts coming first
func (r *PromptService) ListAutoPaging(ctx context.Context, query PromptListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[Prompt] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a prompt object by its id
func (r *PromptService) Delete(ctx context.Context, promptID string, opts ...option.RequestOption) (res *Prompt, err error) {
	opts = append(r.Options[:], opts...)
	if promptID == "" {
		err = errors.New("missing required prompt_id parameter")
		return
	}
	path := fmt.Sprintf("v1/prompt/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create or replace prompt. If there is an existing prompt in the project with the
// same slug as the one specified in the request, will replace the existing prompt
// with the provided fields
func (r *PromptService) Replace(ctx context.Context, body PromptReplaceParams, opts ...option.RequestOption) (res *Prompt, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/prompt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type Prompt struct {
	// Unique identifier for the prompt
	ID string `json:"id,required" format:"uuid"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the prompt (see the `version`
	// parameter)
	XactID string `json:"_xact_id,required"`
	// A literal 'p' which identifies the object as a project prompt
	LogID PromptLogID `json:"log_id,required"`
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
	PromptData PromptPromptData `json:"prompt_data,nullable"`
	// A list of tags for the prompt
	Tags []string   `json:"tags,nullable"`
	JSON promptJSON `json:"-"`
}

// promptJSON contains the JSON metadata for the struct [Prompt]
type promptJSON struct {
	ID          apijson.Field
	XactID      apijson.Field
	LogID       apijson.Field
	Name        apijson.Field
	OrgID       apijson.Field
	ProjectID   apijson.Field
	Slug        apijson.Field
	Created     apijson.Field
	Description apijson.Field
	Metadata    apijson.Field
	PromptData  apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Prompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptJSON) RawJSON() string {
	return r.raw
}

// A literal 'p' which identifies the object as a project prompt
type PromptLogID string

const (
	PromptLogIDP PromptLogID = "p"
)

func (r PromptLogID) IsKnown() bool {
	switch r {
	case PromptLogIDP:
		return true
	}
	return false
}

// The prompt, model, and its parameters
type PromptPromptData struct {
	Options PromptPromptDataOptions `json:"options,nullable"`
	Origin  PromptPromptDataOrigin  `json:"origin,nullable"`
	Prompt  PromptPromptDataPrompt  `json:"prompt"`
	JSON    promptPromptDataJSON    `json:"-"`
}

// promptPromptDataJSON contains the JSON metadata for the struct
// [PromptPromptData]
type promptPromptDataJSON struct {
	Options     apijson.Field
	Origin      apijson.Field
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataOptions struct {
	Model    string                        `json:"model"`
	Params   PromptPromptDataOptionsParams `json:"params"`
	Position string                        `json:"position"`
	JSON     promptPromptDataOptionsJSON   `json:"-"`
}

// promptPromptDataOptionsJSON contains the JSON metadata for the struct
// [PromptPromptDataOptions]
type promptPromptDataOptionsJSON struct {
	Model       apijson.Field
	Params      apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataOptionsParams struct {
	UseCache         bool    `json:"use_cache"`
	Temperature      float64 `json:"temperature"`
	TopP             float64 `json:"top_p"`
	MaxTokens        float64 `json:"max_tokens"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
	// This field can have the runtime type of
	// [PromptPromptDataOptionsParamsOpenAIModelParamsResponseFormat].
	ResponseFormat interface{} `json:"response_format,required"`
	// This field can have the runtime type of
	// [PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion].
	ToolChoice interface{} `json:"tool_choice,required"`
	// This field can have the runtime type of
	// [PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion].
	FunctionCall interface{} `json:"function_call,required"`
	N            float64     `json:"n"`
	// This field can have the runtime type of [[]string].
	Stop interface{} `json:"stop,required"`
	TopK float64     `json:"top_k"`
	// This field can have the runtime type of [[]string].
	StopSequences interface{} `json:"stop_sequences,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample float64                           `json:"max_tokens_to_sample"`
	MaxOutputTokens   float64                           `json:"maxOutputTokens"`
	TopP              float64                           `json:"topP"`
	TopK              float64                           `json:"topK"`
	JSON              promptPromptDataOptionsParamsJSON `json:"-"`
	union             PromptPromptDataOptionsParamsUnion
}

// promptPromptDataOptionsParamsJSON contains the JSON metadata for the struct
// [PromptPromptDataOptionsParams]
type promptPromptDataOptionsParamsJSON struct {
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

func (r promptPromptDataOptionsParamsJSON) RawJSON() string {
	return r.raw
}

func (r *PromptPromptDataOptionsParams) UnmarshalJSON(data []byte) (err error) {
	*r = PromptPromptDataOptionsParams{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PromptPromptDataOptionsParamsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PromptPromptDataOptionsParamsOpenAIModelParams],
// [PromptPromptDataOptionsParamsAnthropicModelParams],
// [PromptPromptDataOptionsParamsGoogleModelParams],
// [PromptPromptDataOptionsParamsWindowAIModelParams],
// [PromptPromptDataOptionsParamsJsCompletionParams].
func (r PromptPromptDataOptionsParams) AsUnion() PromptPromptDataOptionsParamsUnion {
	return r.union
}

// Union satisfied by [PromptPromptDataOptionsParamsOpenAIModelParams],
// [PromptPromptDataOptionsParamsAnthropicModelParams],
// [PromptPromptDataOptionsParamsGoogleModelParams],
// [PromptPromptDataOptionsParamsWindowAIModelParams] or
// [PromptPromptDataOptionsParamsJsCompletionParams].
type PromptPromptDataOptionsParamsUnion interface {
	implementsPromptPromptDataOptionsParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataOptionsParamsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsOpenAIModelParams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsAnthropicModelParams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsGoogleModelParams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsWindowAIModelParams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsJsCompletionParams{}),
		},
	)
}

type PromptPromptDataOptionsParamsOpenAIModelParams struct {
	FrequencyPenalty float64                                                         `json:"frequency_penalty"`
	FunctionCall     PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion `json:"function_call"`
	MaxTokens        float64                                                         `json:"max_tokens"`
	N                float64                                                         `json:"n"`
	PresencePenalty  float64                                                         `json:"presence_penalty"`
	ResponseFormat   PromptPromptDataOptionsParamsOpenAIModelParamsResponseFormat    `json:"response_format,nullable"`
	Stop             []string                                                        `json:"stop"`
	Temperature      float64                                                         `json:"temperature"`
	ToolChoice       PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion   `json:"tool_choice"`
	TopP             float64                                                         `json:"top_p"`
	UseCache         bool                                                            `json:"use_cache"`
	JSON             promptPromptDataOptionsParamsOpenAIModelParamsJSON              `json:"-"`
}

// promptPromptDataOptionsParamsOpenAIModelParamsJSON contains the JSON metadata
// for the struct [PromptPromptDataOptionsParamsOpenAIModelParams]
type promptPromptDataOptionsParamsOpenAIModelParamsJSON struct {
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

func (r *PromptPromptDataOptionsParamsOpenAIModelParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsOpenAIModelParamsJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsOpenAIModelParams) implementsPromptPromptDataOptionsParams() {}

// Union satisfied by
// [PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallString],
// [PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallString] or
// [PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject].
type PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsPromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject{}),
		},
	)
}

type PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallString string

const (
	PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallString = "auto"
)

func (r PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto:
		return true
	}
	return false
}

func (r PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) implementsPromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject struct {
	Name string                                                               `json:"name,required"`
	JSON promptPromptDataOptionsParamsOpenAIModelParamsFunctionCallObjectJSON `json:"-"`
}

// promptPromptDataOptionsParamsOpenAIModelParamsFunctionCallObjectJSON contains
// the JSON metadata for the struct
// [PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject]
type promptPromptDataOptionsParamsOpenAIModelParamsFunctionCallObjectJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsOpenAIModelParamsFunctionCallObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) implementsPromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptPromptDataOptionsParamsOpenAIModelParamsResponseFormat struct {
	Type PromptPromptDataOptionsParamsOpenAIModelParamsResponseFormatType `json:"type,required"`
	JSON promptPromptDataOptionsParamsOpenAIModelParamsResponseFormatJSON `json:"-"`
}

// promptPromptDataOptionsParamsOpenAIModelParamsResponseFormatJSON contains the
// JSON metadata for the struct
// [PromptPromptDataOptionsParamsOpenAIModelParamsResponseFormat]
type promptPromptDataOptionsParamsOpenAIModelParamsResponseFormatJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsOpenAIModelParamsResponseFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsOpenAIModelParamsResponseFormatJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataOptionsParamsOpenAIModelParamsResponseFormatType string

const (
	PromptPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject PromptPromptDataOptionsParamsOpenAIModelParamsResponseFormatType = "json_object"
)

func (r PromptPromptDataOptionsParamsOpenAIModelParamsResponseFormatType) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Union satisfied by
// [PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceString],
// [PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceString] or
// [PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject].
type PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsPromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject{}),
		},
	)
}

type PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceString string

const (
	PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceString = "auto"
)

func (r PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto:
		return true
	}
	return false
}

func (r PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) implementsPromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject struct {
	Function PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction `json:"function,required"`
	Type     PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType     `json:"type,required"`
	JSON     promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectJSON     `json:"-"`
}

// promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectJSON contains the
// JSON metadata for the struct
// [PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject]
type promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectJSON struct {
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) implementsPromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction struct {
	Name string                                                                     `json:"name,required"`
	JSON promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunctionJSON `json:"-"`
}

// promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunctionJSON
// contains the JSON metadata for the struct
// [PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction]
type promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunctionJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType string

const (
	PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType = "function"
)

func (r PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction:
		return true
	}
	return false
}

type PromptPromptDataOptionsParamsAnthropicModelParams struct {
	MaxTokens   float64 `json:"max_tokens,required"`
	Temperature float64 `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample float64                                               `json:"max_tokens_to_sample"`
	StopSequences     []string                                              `json:"stop_sequences"`
	TopK              float64                                               `json:"top_k"`
	TopP              float64                                               `json:"top_p"`
	UseCache          bool                                                  `json:"use_cache"`
	JSON              promptPromptDataOptionsParamsAnthropicModelParamsJSON `json:"-"`
}

// promptPromptDataOptionsParamsAnthropicModelParamsJSON contains the JSON metadata
// for the struct [PromptPromptDataOptionsParamsAnthropicModelParams]
type promptPromptDataOptionsParamsAnthropicModelParamsJSON struct {
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

func (r *PromptPromptDataOptionsParamsAnthropicModelParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsAnthropicModelParamsJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsAnthropicModelParams) implementsPromptPromptDataOptionsParams() {
}

type PromptPromptDataOptionsParamsGoogleModelParams struct {
	MaxOutputTokens float64                                            `json:"maxOutputTokens"`
	Temperature     float64                                            `json:"temperature"`
	TopK            float64                                            `json:"topK"`
	TopP            float64                                            `json:"topP"`
	UseCache        bool                                               `json:"use_cache"`
	JSON            promptPromptDataOptionsParamsGoogleModelParamsJSON `json:"-"`
}

// promptPromptDataOptionsParamsGoogleModelParamsJSON contains the JSON metadata
// for the struct [PromptPromptDataOptionsParamsGoogleModelParams]
type promptPromptDataOptionsParamsGoogleModelParamsJSON struct {
	MaxOutputTokens apijson.Field
	Temperature     apijson.Field
	TopK            apijson.Field
	TopP            apijson.Field
	UseCache        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsGoogleModelParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsGoogleModelParamsJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsGoogleModelParams) implementsPromptPromptDataOptionsParams() {}

type PromptPromptDataOptionsParamsWindowAIModelParams struct {
	Temperature float64                                              `json:"temperature"`
	TopK        float64                                              `json:"topK"`
	UseCache    bool                                                 `json:"use_cache"`
	JSON        promptPromptDataOptionsParamsWindowAIModelParamsJSON `json:"-"`
}

// promptPromptDataOptionsParamsWindowAIModelParamsJSON contains the JSON metadata
// for the struct [PromptPromptDataOptionsParamsWindowAIModelParams]
type promptPromptDataOptionsParamsWindowAIModelParamsJSON struct {
	Temperature apijson.Field
	TopK        apijson.Field
	UseCache    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsWindowAIModelParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsWindowAIModelParamsJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsWindowAIModelParams) implementsPromptPromptDataOptionsParams() {}

type PromptPromptDataOptionsParamsJsCompletionParams struct {
	UseCache bool                                                `json:"use_cache"`
	JSON     promptPromptDataOptionsParamsJsCompletionParamsJSON `json:"-"`
}

// promptPromptDataOptionsParamsJsCompletionParamsJSON contains the JSON metadata
// for the struct [PromptPromptDataOptionsParamsJsCompletionParams]
type promptPromptDataOptionsParamsJsCompletionParamsJSON struct {
	UseCache    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsJsCompletionParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsJsCompletionParamsJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsJsCompletionParams) implementsPromptPromptDataOptionsParams() {}

type PromptPromptDataOrigin struct {
	ProjectID     string                     `json:"project_id"`
	PromptID      string                     `json:"prompt_id"`
	PromptVersion string                     `json:"prompt_version"`
	JSON          promptPromptDataOriginJSON `json:"-"`
}

// promptPromptDataOriginJSON contains the JSON metadata for the struct
// [PromptPromptDataOrigin]
type promptPromptDataOriginJSON struct {
	ProjectID     apijson.Field
	PromptID      apijson.Field
	PromptVersion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PromptPromptDataOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOriginJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataPrompt struct {
	Type    PromptPromptDataPromptType `json:"type"`
	Content string                     `json:"content"`
	// This field can have the runtime type of [[]PromptPromptDataPromptChatMessage].
	Messages interface{}                `json:"messages,required"`
	Tools    string                     `json:"tools"`
	JSON     promptPromptDataPromptJSON `json:"-"`
	union    PromptPromptDataPromptUnion
}

// promptPromptDataPromptJSON contains the JSON metadata for the struct
// [PromptPromptDataPrompt]
type promptPromptDataPromptJSON struct {
	Type        apijson.Field
	Content     apijson.Field
	Messages    apijson.Field
	Tools       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r promptPromptDataPromptJSON) RawJSON() string {
	return r.raw
}

func (r *PromptPromptDataPrompt) UnmarshalJSON(data []byte) (err error) {
	*r = PromptPromptDataPrompt{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PromptPromptDataPromptUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [PromptPromptDataPromptCompletion],
// [PromptPromptDataPromptChat], [PromptPromptDataPromptNullableVariant].
func (r PromptPromptDataPrompt) AsUnion() PromptPromptDataPromptUnion {
	return r.union
}

// Union satisfied by [PromptPromptDataPromptCompletion],
// [PromptPromptDataPromptChat] or [PromptPromptDataPromptNullableVariant].
type PromptPromptDataPromptUnion interface {
	implementsPromptPromptDataPrompt()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataPromptUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptCompletion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChat{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptNullableVariant{}),
		},
	)
}

type PromptPromptDataPromptCompletion struct {
	Content string                               `json:"content,required"`
	Type    PromptPromptDataPromptCompletionType `json:"type,required"`
	JSON    promptPromptDataPromptCompletionJSON `json:"-"`
}

// promptPromptDataPromptCompletionJSON contains the JSON metadata for the struct
// [PromptPromptDataPromptCompletion]
type promptPromptDataPromptCompletionJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptCompletion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptCompletionJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptCompletion) implementsPromptPromptDataPrompt() {}

type PromptPromptDataPromptCompletionType string

const (
	PromptPromptDataPromptCompletionTypeCompletion PromptPromptDataPromptCompletionType = "completion"
)

func (r PromptPromptDataPromptCompletionType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptCompletionTypeCompletion:
		return true
	}
	return false
}

type PromptPromptDataPromptChat struct {
	Messages []PromptPromptDataPromptChatMessage `json:"messages,required"`
	Type     PromptPromptDataPromptChatType      `json:"type,required"`
	Tools    string                              `json:"tools"`
	JSON     promptPromptDataPromptChatJSON      `json:"-"`
}

// promptPromptDataPromptChatJSON contains the JSON metadata for the struct
// [PromptPromptDataPromptChat]
type promptPromptDataPromptChatJSON struct {
	Messages    apijson.Field
	Type        apijson.Field
	Tools       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChat) implementsPromptPromptDataPrompt() {}

type PromptPromptDataPromptChatMessage struct {
	// This field can have the runtime type of [string],
	// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion].
	Content interface{}                            `json:"content,required"`
	Role    PromptPromptDataPromptChatMessagesRole `json:"role,required"`
	Name    string                                 `json:"name"`
	// This field can have the runtime type of
	// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall].
	FunctionCall interface{} `json:"function_call,required"`
	// This field can have the runtime type of
	// [[]PromptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall].
	ToolCalls  interface{}                           `json:"tool_calls,required"`
	ToolCallID string                                `json:"tool_call_id"`
	JSON       promptPromptDataPromptChatMessageJSON `json:"-"`
	union      PromptPromptDataPromptChatMessagesUnion
}

// promptPromptDataPromptChatMessageJSON contains the JSON metadata for the struct
// [PromptPromptDataPromptChatMessage]
type promptPromptDataPromptChatMessageJSON struct {
	Content      apijson.Field
	Role         apijson.Field
	Name         apijson.Field
	FunctionCall apijson.Field
	ToolCalls    apijson.Field
	ToolCallID   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r promptPromptDataPromptChatMessageJSON) RawJSON() string {
	return r.raw
}

func (r *PromptPromptDataPromptChatMessage) UnmarshalJSON(data []byte) (err error) {
	*r = PromptPromptDataPromptChatMessage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PromptPromptDataPromptChatMessagesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage0],
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage1],
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage2],
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage3],
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage4],
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage5].
func (r PromptPromptDataPromptChatMessage) AsUnion() PromptPromptDataPromptChatMessagesUnion {
	return r.union
}

// Union satisfied by [PromptPromptDataPromptChatMessagesPromptDataPromptMessage0],
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage1],
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage2],
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage3],
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage4] or
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage5].
type PromptPromptDataPromptChatMessagesUnion interface {
	implementsPromptPromptDataPromptChatMessage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataPromptChatMessagesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesPromptDataPromptMessage0{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesPromptDataPromptMessage1{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesPromptDataPromptMessage2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesPromptDataPromptMessage3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesPromptDataPromptMessage4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesPromptDataPromptMessage5{}),
		},
	)
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage0 struct {
	Role    PromptPromptDataPromptChatMessagesPromptDataPromptMessage0Role `json:"role,required"`
	Content string                                                         `json:"content"`
	Name    string                                                         `json:"name"`
	JSON    promptPromptDataPromptChatMessagesPromptDataPromptMessage0JSON `json:"-"`
}

// promptPromptDataPromptChatMessagesPromptDataPromptMessage0JSON contains the JSON
// metadata for the struct
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage0]
type promptPromptDataPromptChatMessagesPromptDataPromptMessage0JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesPromptDataPromptMessage0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesPromptDataPromptMessage0JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage0) implementsPromptPromptDataPromptChatMessage() {
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage0Role string

const (
	PromptPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem PromptPromptDataPromptChatMessagesPromptDataPromptMessage0Role = "system"
)

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage1 struct {
	Role    PromptPromptDataPromptChatMessagesPromptDataPromptMessage1Role         `json:"role,required"`
	Content PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion `json:"content"`
	Name    string                                                                 `json:"name"`
	JSON    promptPromptDataPromptChatMessagesPromptDataPromptMessage1JSON         `json:"-"`
}

// promptPromptDataPromptChatMessagesPromptDataPromptMessage1JSON contains the JSON
// metadata for the struct
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage1]
type promptPromptDataPromptChatMessagesPromptDataPromptMessage1JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesPromptDataPromptMessage1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesPromptDataPromptMessage1JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage1) implementsPromptPromptDataPromptChatMessage() {
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage1Role string

const (
	PromptPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser PromptPromptDataPromptChatMessagesPromptDataPromptMessage1Role = "user"
)

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray].
type PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsPromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray{}),
		},
	)
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray []PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) ImplementsPromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion() {
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem struct {
	Text string                                                                     `json:"text"`
	Type PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType `json:"type,required"`
	// This field can have the runtime type of
	// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL].
	ImageURL interface{}                                                                    `json:"image_url,required"`
	JSON     promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItemJSON `json:"-"`
	union    PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnionItem
}

// promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItemJSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem]
type promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItemJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	ImageURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItemJSON) RawJSON() string {
	return r.raw
}

func (r *PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem) UnmarshalJSON(data []byte) (err error) {
	*r = PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnionItem]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList].
func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem) AsUnion() PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnionItem {
	return r.union
}

// Union satisfied by
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent]
// or
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList].
type PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnionItem interface {
	implementsPromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList{}),
		},
	)
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType `json:"type,required"`
	Text string                                                                                                   `json:"text"`
	JSON promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON `json:"-"`
}

// promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent]
type promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON struct {
	Type        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsPromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem() {
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL `json:"image_url,required"`
	Type     PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType     `json:"type,required"`
	JSON     promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON     `json:"-"`
}

// promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList]
type promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON struct {
	ImageURL    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsPromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayItem() {
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    string                                                                                                          `json:"url,required"`
	Detail PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail `json:"detail"`
	JSON   promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON   `json:"-"`
}

// promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL]
type promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON struct {
	URL         apijson.Field
	Detail      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType string

const (
	PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText     PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "text"
	PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText, PromptPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage2 struct {
	Role         PromptPromptDataPromptChatMessagesPromptDataPromptMessage2Role         `json:"role,required"`
	Content      string                                                                 `json:"content,nullable"`
	FunctionCall PromptPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall `json:"function_call"`
	Name         string                                                                 `json:"name"`
	ToolCalls    []PromptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall   `json:"tool_calls"`
	JSON         promptPromptDataPromptChatMessagesPromptDataPromptMessage2JSON         `json:"-"`
}

// promptPromptDataPromptChatMessagesPromptDataPromptMessage2JSON contains the JSON
// metadata for the struct
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage2]
type promptPromptDataPromptChatMessagesPromptDataPromptMessage2JSON struct {
	Role         apijson.Field
	Content      apijson.Field
	FunctionCall apijson.Field
	Name         apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesPromptDataPromptMessage2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesPromptDataPromptMessage2JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage2) implementsPromptPromptDataPromptChatMessage() {
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage2Role string

const (
	PromptPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant PromptPromptDataPromptChatMessagesPromptDataPromptMessage2Role = "assistant"
)

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments string                                                                     `json:"arguments,required"`
	Name      string                                                                     `json:"name,required"`
	JSON      promptPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCallJSON `json:"-"`
}

// promptPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCallJSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall]
type promptPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCallJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall struct {
	ID       string                                                                      `json:"id,required"`
	Function PromptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction `json:"function,required"`
	Type     PromptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType     `json:"type,required"`
	JSON     promptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallJSON      `json:"-"`
}

// promptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallJSON contains
// the JSON metadata for the struct
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall]
type promptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallJSON struct {
	ID          apijson.Field
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments string                                                                          `json:"arguments,required"`
	Name      string                                                                          `json:"name,required"`
	JSON      promptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunctionJSON `json:"-"`
}

// promptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunctionJSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction]
type promptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunctionJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType string

const (
	PromptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction PromptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage3 struct {
	Role       PromptPromptDataPromptChatMessagesPromptDataPromptMessage3Role `json:"role,required"`
	Content    string                                                         `json:"content"`
	ToolCallID string                                                         `json:"tool_call_id"`
	JSON       promptPromptDataPromptChatMessagesPromptDataPromptMessage3JSON `json:"-"`
}

// promptPromptDataPromptChatMessagesPromptDataPromptMessage3JSON contains the JSON
// metadata for the struct
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage3]
type promptPromptDataPromptChatMessagesPromptDataPromptMessage3JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	ToolCallID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesPromptDataPromptMessage3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesPromptDataPromptMessage3JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage3) implementsPromptPromptDataPromptChatMessage() {
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage3Role string

const (
	PromptPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool PromptPromptDataPromptChatMessagesPromptDataPromptMessage3Role = "tool"
)

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage4 struct {
	Name    string                                                         `json:"name,required"`
	Role    PromptPromptDataPromptChatMessagesPromptDataPromptMessage4Role `json:"role,required"`
	Content string                                                         `json:"content"`
	JSON    promptPromptDataPromptChatMessagesPromptDataPromptMessage4JSON `json:"-"`
}

// promptPromptDataPromptChatMessagesPromptDataPromptMessage4JSON contains the JSON
// metadata for the struct
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage4]
type promptPromptDataPromptChatMessagesPromptDataPromptMessage4JSON struct {
	Name        apijson.Field
	Role        apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesPromptDataPromptMessage4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesPromptDataPromptMessage4JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage4) implementsPromptPromptDataPromptChatMessage() {
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage4Role string

const (
	PromptPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction PromptPromptDataPromptChatMessagesPromptDataPromptMessage4Role = "function"
)

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage5 struct {
	Role    PromptPromptDataPromptChatMessagesPromptDataPromptMessage5Role `json:"role,required"`
	Content string                                                         `json:"content,nullable"`
	JSON    promptPromptDataPromptChatMessagesPromptDataPromptMessage5JSON `json:"-"`
}

// promptPromptDataPromptChatMessagesPromptDataPromptMessage5JSON contains the JSON
// metadata for the struct
// [PromptPromptDataPromptChatMessagesPromptDataPromptMessage5]
type promptPromptDataPromptChatMessagesPromptDataPromptMessage5JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesPromptDataPromptMessage5) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesPromptDataPromptMessage5JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage5) implementsPromptPromptDataPromptChatMessage() {
}

type PromptPromptDataPromptChatMessagesPromptDataPromptMessage5Role string

const (
	PromptPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel PromptPromptDataPromptChatMessagesPromptDataPromptMessage5Role = "model"
)

func (r PromptPromptDataPromptChatMessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesRole string

const (
	PromptPromptDataPromptChatMessagesRoleSystem    PromptPromptDataPromptChatMessagesRole = "system"
	PromptPromptDataPromptChatMessagesRoleUser      PromptPromptDataPromptChatMessagesRole = "user"
	PromptPromptDataPromptChatMessagesRoleAssistant PromptPromptDataPromptChatMessagesRole = "assistant"
	PromptPromptDataPromptChatMessagesRoleTool      PromptPromptDataPromptChatMessagesRole = "tool"
	PromptPromptDataPromptChatMessagesRoleFunction  PromptPromptDataPromptChatMessagesRole = "function"
	PromptPromptDataPromptChatMessagesRoleModel     PromptPromptDataPromptChatMessagesRole = "model"
)

func (r PromptPromptDataPromptChatMessagesRole) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesRoleSystem, PromptPromptDataPromptChatMessagesRoleUser, PromptPromptDataPromptChatMessagesRoleAssistant, PromptPromptDataPromptChatMessagesRoleTool, PromptPromptDataPromptChatMessagesRoleFunction, PromptPromptDataPromptChatMessagesRoleModel:
		return true
	}
	return false
}

type PromptPromptDataPromptChatType string

const (
	PromptPromptDataPromptChatTypeChat PromptPromptDataPromptChatType = "chat"
)

func (r PromptPromptDataPromptChatType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatTypeChat:
		return true
	}
	return false
}

type PromptPromptDataPromptNullableVariant struct {
	JSON promptPromptDataPromptNullableVariantJSON `json:"-"`
}

// promptPromptDataPromptNullableVariantJSON contains the JSON metadata for the
// struct [PromptPromptDataPromptNullableVariant]
type promptPromptDataPromptNullableVariantJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptNullableVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptNullableVariantJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptNullableVariant) implementsPromptPromptDataPrompt() {}

type PromptPromptDataPromptType string

const (
	PromptPromptDataPromptTypeCompletion PromptPromptDataPromptType = "completion"
	PromptPromptDataPromptTypeChat       PromptPromptDataPromptType = "chat"
)

func (r PromptPromptDataPromptType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptTypeCompletion, PromptPromptDataPromptTypeChat:
		return true
	}
	return false
}

type PromptNewParams struct {
	// Name of the prompt
	Name param.Field[string] `json:"name,required"`
	// Unique identifier for the project that the prompt belongs under
	ProjectID param.Field[string] `json:"project_id,required" format:"uuid"`
	// Unique identifier for the prompt
	Slug param.Field[string] `json:"slug,required"`
	// Textual description of the prompt
	Description param.Field[string] `json:"description"`
	// The prompt, model, and its parameters
	PromptData param.Field[PromptNewParamsPromptData] `json:"prompt_data"`
	// A list of tags for the prompt
	Tags param.Field[[]string] `json:"tags"`
}

func (r PromptNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The prompt, model, and its parameters
type PromptNewParamsPromptData struct {
	Options param.Field[PromptNewParamsPromptDataOptions]     `json:"options"`
	Origin  param.Field[PromptNewParamsPromptDataOrigin]      `json:"origin"`
	Prompt  param.Field[PromptNewParamsPromptDataPromptUnion] `json:"prompt"`
}

func (r PromptNewParamsPromptData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataOptions struct {
	Model    param.Field[string]                                      `json:"model"`
	Params   param.Field[PromptNewParamsPromptDataOptionsParamsUnion] `json:"params"`
	Position param.Field[string]                                      `json:"position"`
}

func (r PromptNewParamsPromptDataOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataOptionsParams struct {
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

func (r PromptNewParamsPromptDataOptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParams) implementsPromptNewParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by [PromptNewParamsPromptDataOptionsParamsOpenAIModelParams],
// [PromptNewParamsPromptDataOptionsParamsAnthropicModelParams],
// [PromptNewParamsPromptDataOptionsParamsGoogleModelParams],
// [PromptNewParamsPromptDataOptionsParamsWindowAIModelParams],
// [PromptNewParamsPromptDataOptionsParamsJsCompletionParams],
// [PromptNewParamsPromptDataOptionsParams].
type PromptNewParamsPromptDataOptionsParamsUnion interface {
	implementsPromptNewParamsPromptDataOptionsParamsUnion()
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParams struct {
	FrequencyPenalty param.Field[float64]                                                                  `json:"frequency_penalty"`
	FunctionCall     param.Field[PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                                  `json:"max_tokens"`
	N                param.Field[float64]                                                                  `json:"n"`
	PresencePenalty  param.Field[float64]                                                                  `json:"presence_penalty"`
	ResponseFormat   param.Field[PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                                 `json:"stop"`
	Temperature      param.Field[float64]                                                                  `json:"temperature"`
	ToolChoice       param.Field[PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                                  `json:"top_p"`
	UseCache         param.Field[bool]                                                                     `json:"use_cache"`
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParams) implementsPromptNewParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by
// [PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString],
// [PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString],
// [PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject].
type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsPromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString string

const (
	PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString = "auto"
)

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto:
		return true
	}
	return false
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) implementsPromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) implementsPromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat struct {
	Type param.Field[PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType] `json:"type,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType string

const (
	PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType = "json_object"
)

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Satisfied by
// [PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString],
// [PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString],
// [PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject].
type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsPromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString string

const (
	PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString = "auto"
)

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto:
		return true
	}
	return false
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) implementsPromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject struct {
	Function param.Field[PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction] `json:"function,required"`
	Type     param.Field[PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType]     `json:"type,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) implementsPromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType string

const (
	PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType = "function"
)

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction:
		return true
	}
	return false
}

type PromptNewParamsPromptDataOptionsParamsAnthropicModelParams struct {
	MaxTokens   param.Field[float64] `json:"max_tokens,required"`
	Temperature param.Field[float64] `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64]  `json:"max_tokens_to_sample"`
	StopSequences     param.Field[[]string] `json:"stop_sequences"`
	TopK              param.Field[float64]  `json:"top_k"`
	TopP              param.Field[float64]  `json:"top_p"`
	UseCache          param.Field[bool]     `json:"use_cache"`
}

func (r PromptNewParamsPromptDataOptionsParamsAnthropicModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsAnthropicModelParams) implementsPromptNewParamsPromptDataOptionsParamsUnion() {
}

type PromptNewParamsPromptDataOptionsParamsGoogleModelParams struct {
	MaxOutputTokens param.Field[float64] `json:"maxOutputTokens"`
	Temperature     param.Field[float64] `json:"temperature"`
	TopK            param.Field[float64] `json:"topK"`
	TopP            param.Field[float64] `json:"topP"`
	UseCache        param.Field[bool]    `json:"use_cache"`
}

func (r PromptNewParamsPromptDataOptionsParamsGoogleModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsGoogleModelParams) implementsPromptNewParamsPromptDataOptionsParamsUnion() {
}

type PromptNewParamsPromptDataOptionsParamsWindowAIModelParams struct {
	Temperature param.Field[float64] `json:"temperature"`
	TopK        param.Field[float64] `json:"topK"`
	UseCache    param.Field[bool]    `json:"use_cache"`
}

func (r PromptNewParamsPromptDataOptionsParamsWindowAIModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsWindowAIModelParams) implementsPromptNewParamsPromptDataOptionsParamsUnion() {
}

type PromptNewParamsPromptDataOptionsParamsJsCompletionParams struct {
	UseCache param.Field[bool] `json:"use_cache"`
}

func (r PromptNewParamsPromptDataOptionsParamsJsCompletionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsJsCompletionParams) implementsPromptNewParamsPromptDataOptionsParamsUnion() {
}

type PromptNewParamsPromptDataOrigin struct {
	ProjectID     param.Field[string] `json:"project_id"`
	PromptID      param.Field[string] `json:"prompt_id"`
	PromptVersion param.Field[string] `json:"prompt_version"`
}

func (r PromptNewParamsPromptDataOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataPrompt struct {
	Type     param.Field[PromptNewParamsPromptDataPromptType] `json:"type"`
	Content  param.Field[string]                              `json:"content"`
	Messages param.Field[interface{}]                         `json:"messages,required"`
	Tools    param.Field[string]                              `json:"tools"`
}

func (r PromptNewParamsPromptDataPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPrompt) implementsPromptNewParamsPromptDataPromptUnion() {}

// Satisfied by [PromptNewParamsPromptDataPromptCompletion],
// [PromptNewParamsPromptDataPromptChat],
// [PromptNewParamsPromptDataPromptNullableVariant],
// [PromptNewParamsPromptDataPrompt].
type PromptNewParamsPromptDataPromptUnion interface {
	implementsPromptNewParamsPromptDataPromptUnion()
}

type PromptNewParamsPromptDataPromptCompletion struct {
	Content param.Field[string]                                        `json:"content,required"`
	Type    param.Field[PromptNewParamsPromptDataPromptCompletionType] `json:"type,required"`
}

func (r PromptNewParamsPromptDataPromptCompletion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptCompletion) implementsPromptNewParamsPromptDataPromptUnion() {}

type PromptNewParamsPromptDataPromptCompletionType string

const (
	PromptNewParamsPromptDataPromptCompletionTypeCompletion PromptNewParamsPromptDataPromptCompletionType = "completion"
)

func (r PromptNewParamsPromptDataPromptCompletionType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptCompletionTypeCompletion:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChat struct {
	Messages param.Field[[]PromptNewParamsPromptDataPromptChatMessageUnion] `json:"messages,required"`
	Type     param.Field[PromptNewParamsPromptDataPromptChatType]           `json:"type,required"`
	Tools    param.Field[string]                                            `json:"tools"`
}

func (r PromptNewParamsPromptDataPromptChat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChat) implementsPromptNewParamsPromptDataPromptUnion() {}

type PromptNewParamsPromptDataPromptChatMessage struct {
	Content      param.Field[interface{}]                                     `json:"content,required"`
	Role         param.Field[PromptNewParamsPromptDataPromptChatMessagesRole] `json:"role,required"`
	Name         param.Field[string]                                          `json:"name"`
	FunctionCall param.Field[interface{}]                                     `json:"function_call,required"`
	ToolCalls    param.Field[interface{}]                                     `json:"tool_calls,required"`
	ToolCallID   param.Field[string]                                          `json:"tool_call_id"`
}

func (r PromptNewParamsPromptDataPromptChatMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessage) implementsPromptNewParamsPromptDataPromptChatMessageUnion() {
}

// Satisfied by
// [PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0],
// [PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1],
// [PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2],
// [PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3],
// [PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4],
// [PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5],
// [PromptNewParamsPromptDataPromptChatMessage].
type PromptNewParamsPromptDataPromptChatMessageUnion interface {
	implementsPromptNewParamsPromptDataPromptChatMessageUnion()
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0 struct {
	Role    param.Field[PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role] `json:"role,required"`
	Content param.Field[string]                                                                  `json:"content"`
	Name    param.Field[string]                                                                  `json:"name"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0) implementsPromptNewParamsPromptDataPromptChatMessageUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role string

const (
	PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role = "system"
)

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1 struct {
	Role    param.Field[PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role]         `json:"role,required"`
	Content param.Field[PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion] `json:"content"`
	Name    param.Field[string]                                                                          `json:"name"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1) implementsPromptNewParamsPromptDataPromptChatMessageUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role string

const (
	PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role = "user"
)

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray].
type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsPromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion()
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray []PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) ImplementsPromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray struct {
	Text     param.Field[string]                                                                              `json:"text"`
	Type     param.Field[PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                                         `json:"image_url,required"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) implementsPromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

// Satisfied by
// [PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList],
// [PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray].
type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion interface {
	implementsPromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion()
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type param.Field[PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType] `json:"type,required"`
	Text param.Field[string]                                                                                                            `json:"text"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsPromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL param.Field[PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL] `json:"image_url,required"`
	Type     param.Field[PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType]     `json:"type,required"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsPromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    param.Field[string]                                                                                                                   `json:"url,required"`
	Detail param.Field[PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail] `json:"detail"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType string

const (
	PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText     PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "text"
	PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText, PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2 struct {
	Role         param.Field[PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role]         `json:"role,required"`
	Content      param.Field[string]                                                                          `json:"content"`
	FunctionCall param.Field[PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                                          `json:"name"`
	ToolCalls    param.Field[[]PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall]   `json:"tool_calls"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2) implementsPromptNewParamsPromptDataPromptChatMessageUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role string

const (
	PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role = "assistant"
)

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall struct {
	ID       param.Field[string]                                                                               `json:"id,required"`
	Function param.Field[PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction] `json:"function,required"`
	Type     param.Field[PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType]     `json:"type,required"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType string

const (
	PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3 struct {
	Role       param.Field[PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role] `json:"role,required"`
	Content    param.Field[string]                                                                  `json:"content"`
	ToolCallID param.Field[string]                                                                  `json:"tool_call_id"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3) implementsPromptNewParamsPromptDataPromptChatMessageUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role string

const (
	PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role = "tool"
)

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4 struct {
	Name    param.Field[string]                                                                  `json:"name,required"`
	Role    param.Field[PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role] `json:"role,required"`
	Content param.Field[string]                                                                  `json:"content"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4) implementsPromptNewParamsPromptDataPromptChatMessageUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role string

const (
	PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role = "function"
)

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5 struct {
	Role    param.Field[PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role] `json:"role,required"`
	Content param.Field[string]                                                                  `json:"content"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5) implementsPromptNewParamsPromptDataPromptChatMessageUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role string

const (
	PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role = "model"
)

func (r PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesRole string

const (
	PromptNewParamsPromptDataPromptChatMessagesRoleSystem    PromptNewParamsPromptDataPromptChatMessagesRole = "system"
	PromptNewParamsPromptDataPromptChatMessagesRoleUser      PromptNewParamsPromptDataPromptChatMessagesRole = "user"
	PromptNewParamsPromptDataPromptChatMessagesRoleAssistant PromptNewParamsPromptDataPromptChatMessagesRole = "assistant"
	PromptNewParamsPromptDataPromptChatMessagesRoleTool      PromptNewParamsPromptDataPromptChatMessagesRole = "tool"
	PromptNewParamsPromptDataPromptChatMessagesRoleFunction  PromptNewParamsPromptDataPromptChatMessagesRole = "function"
	PromptNewParamsPromptDataPromptChatMessagesRoleModel     PromptNewParamsPromptDataPromptChatMessagesRole = "model"
)

func (r PromptNewParamsPromptDataPromptChatMessagesRole) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesRoleSystem, PromptNewParamsPromptDataPromptChatMessagesRoleUser, PromptNewParamsPromptDataPromptChatMessagesRoleAssistant, PromptNewParamsPromptDataPromptChatMessagesRoleTool, PromptNewParamsPromptDataPromptChatMessagesRoleFunction, PromptNewParamsPromptDataPromptChatMessagesRoleModel:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatType string

const (
	PromptNewParamsPromptDataPromptChatTypeChat PromptNewParamsPromptDataPromptChatType = "chat"
)

func (r PromptNewParamsPromptDataPromptChatType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatTypeChat:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptNullableVariant struct {
}

func (r PromptNewParamsPromptDataPromptNullableVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptNullableVariant) implementsPromptNewParamsPromptDataPromptUnion() {
}

type PromptNewParamsPromptDataPromptType string

const (
	PromptNewParamsPromptDataPromptTypeCompletion PromptNewParamsPromptDataPromptType = "completion"
	PromptNewParamsPromptDataPromptTypeChat       PromptNewParamsPromptDataPromptType = "chat"
)

func (r PromptNewParamsPromptDataPromptType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptTypeCompletion, PromptNewParamsPromptDataPromptTypeChat:
		return true
	}
	return false
}

type PromptUpdateParams struct {
	// Textual description of the prompt
	Description param.Field[string] `json:"description"`
	// Name of the prompt
	Name param.Field[string] `json:"name"`
	// The prompt, model, and its parameters
	PromptData param.Field[PromptUpdateParamsPromptData] `json:"prompt_data"`
	// A list of tags for the prompt
	Tags param.Field[[]string] `json:"tags"`
}

func (r PromptUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The prompt, model, and its parameters
type PromptUpdateParamsPromptData struct {
	Options param.Field[PromptUpdateParamsPromptDataOptions]     `json:"options"`
	Origin  param.Field[PromptUpdateParamsPromptDataOrigin]      `json:"origin"`
	Prompt  param.Field[PromptUpdateParamsPromptDataPromptUnion] `json:"prompt"`
}

func (r PromptUpdateParamsPromptData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataOptions struct {
	Model    param.Field[string]                                         `json:"model"`
	Params   param.Field[PromptUpdateParamsPromptDataOptionsParamsUnion] `json:"params"`
	Position param.Field[string]                                         `json:"position"`
}

func (r PromptUpdateParamsPromptDataOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataOptionsParams struct {
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

func (r PromptUpdateParamsPromptDataOptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParams) implementsPromptUpdateParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by [PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParams],
// [PromptUpdateParamsPromptDataOptionsParamsAnthropicModelParams],
// [PromptUpdateParamsPromptDataOptionsParamsGoogleModelParams],
// [PromptUpdateParamsPromptDataOptionsParamsWindowAIModelParams],
// [PromptUpdateParamsPromptDataOptionsParamsJsCompletionParams],
// [PromptUpdateParamsPromptDataOptionsParams].
type PromptUpdateParamsPromptDataOptionsParamsUnion interface {
	implementsPromptUpdateParamsPromptDataOptionsParamsUnion()
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParams struct {
	FrequencyPenalty param.Field[float64]                                                                     `json:"frequency_penalty"`
	FunctionCall     param.Field[PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                                     `json:"max_tokens"`
	N                param.Field[float64]                                                                     `json:"n"`
	PresencePenalty  param.Field[float64]                                                                     `json:"presence_penalty"`
	ResponseFormat   param.Field[PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                                    `json:"stop"`
	Temperature      param.Field[float64]                                                                     `json:"temperature"`
	ToolChoice       param.Field[PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                                     `json:"top_p"`
	UseCache         param.Field[bool]                                                                        `json:"use_cache"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParams) implementsPromptUpdateParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by
// [PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString],
// [PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString],
// [PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject].
type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsPromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString string

const (
	PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString = "auto"
)

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto:
		return true
	}
	return false
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) implementsPromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) implementsPromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat struct {
	Type param.Field[PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType] `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType string

const (
	PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType = "json_object"
)

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Satisfied by
// [PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString],
// [PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString],
// [PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject].
type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsPromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString string

const (
	PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString = "auto"
)

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto:
		return true
	}
	return false
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) implementsPromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject struct {
	Function param.Field[PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction] `json:"function,required"`
	Type     param.Field[PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType]     `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) implementsPromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType string

const (
	PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType = "function"
)

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataOptionsParamsAnthropicModelParams struct {
	MaxTokens   param.Field[float64] `json:"max_tokens,required"`
	Temperature param.Field[float64] `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64]  `json:"max_tokens_to_sample"`
	StopSequences     param.Field[[]string] `json:"stop_sequences"`
	TopK              param.Field[float64]  `json:"top_k"`
	TopP              param.Field[float64]  `json:"top_p"`
	UseCache          param.Field[bool]     `json:"use_cache"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsAnthropicModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsAnthropicModelParams) implementsPromptUpdateParamsPromptDataOptionsParamsUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsGoogleModelParams struct {
	MaxOutputTokens param.Field[float64] `json:"maxOutputTokens"`
	Temperature     param.Field[float64] `json:"temperature"`
	TopK            param.Field[float64] `json:"topK"`
	TopP            param.Field[float64] `json:"topP"`
	UseCache        param.Field[bool]    `json:"use_cache"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsGoogleModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsGoogleModelParams) implementsPromptUpdateParamsPromptDataOptionsParamsUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsWindowAIModelParams struct {
	Temperature param.Field[float64] `json:"temperature"`
	TopK        param.Field[float64] `json:"topK"`
	UseCache    param.Field[bool]    `json:"use_cache"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsWindowAIModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsWindowAIModelParams) implementsPromptUpdateParamsPromptDataOptionsParamsUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsJsCompletionParams struct {
	UseCache param.Field[bool] `json:"use_cache"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsJsCompletionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsJsCompletionParams) implementsPromptUpdateParamsPromptDataOptionsParamsUnion() {
}

type PromptUpdateParamsPromptDataOrigin struct {
	ProjectID     param.Field[string] `json:"project_id"`
	PromptID      param.Field[string] `json:"prompt_id"`
	PromptVersion param.Field[string] `json:"prompt_version"`
}

func (r PromptUpdateParamsPromptDataOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataPrompt struct {
	Type     param.Field[PromptUpdateParamsPromptDataPromptType] `json:"type"`
	Content  param.Field[string]                                 `json:"content"`
	Messages param.Field[interface{}]                            `json:"messages,required"`
	Tools    param.Field[string]                                 `json:"tools"`
}

func (r PromptUpdateParamsPromptDataPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPrompt) implementsPromptUpdateParamsPromptDataPromptUnion() {}

// Satisfied by [PromptUpdateParamsPromptDataPromptCompletion],
// [PromptUpdateParamsPromptDataPromptChat],
// [PromptUpdateParamsPromptDataPromptNullableVariant],
// [PromptUpdateParamsPromptDataPrompt].
type PromptUpdateParamsPromptDataPromptUnion interface {
	implementsPromptUpdateParamsPromptDataPromptUnion()
}

type PromptUpdateParamsPromptDataPromptCompletion struct {
	Content param.Field[string]                                           `json:"content,required"`
	Type    param.Field[PromptUpdateParamsPromptDataPromptCompletionType] `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataPromptCompletion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptCompletion) implementsPromptUpdateParamsPromptDataPromptUnion() {
}

type PromptUpdateParamsPromptDataPromptCompletionType string

const (
	PromptUpdateParamsPromptDataPromptCompletionTypeCompletion PromptUpdateParamsPromptDataPromptCompletionType = "completion"
)

func (r PromptUpdateParamsPromptDataPromptCompletionType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptCompletionTypeCompletion:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChat struct {
	Messages param.Field[[]PromptUpdateParamsPromptDataPromptChatMessageUnion] `json:"messages,required"`
	Type     param.Field[PromptUpdateParamsPromptDataPromptChatType]           `json:"type,required"`
	Tools    param.Field[string]                                               `json:"tools"`
}

func (r PromptUpdateParamsPromptDataPromptChat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChat) implementsPromptUpdateParamsPromptDataPromptUnion() {}

type PromptUpdateParamsPromptDataPromptChatMessage struct {
	Content      param.Field[interface{}]                                        `json:"content,required"`
	Role         param.Field[PromptUpdateParamsPromptDataPromptChatMessagesRole] `json:"role,required"`
	Name         param.Field[string]                                             `json:"name"`
	FunctionCall param.Field[interface{}]                                        `json:"function_call,required"`
	ToolCalls    param.Field[interface{}]                                        `json:"tool_calls,required"`
	ToolCallID   param.Field[string]                                             `json:"tool_call_id"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessage) implementsPromptUpdateParamsPromptDataPromptChatMessageUnion() {
}

// Satisfied by
// [PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0],
// [PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1],
// [PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2],
// [PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3],
// [PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4],
// [PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5],
// [PromptUpdateParamsPromptDataPromptChatMessage].
type PromptUpdateParamsPromptDataPromptChatMessageUnion interface {
	implementsPromptUpdateParamsPromptDataPromptChatMessageUnion()
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0 struct {
	Role    param.Field[PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role] `json:"role,required"`
	Content param.Field[string]                                                                     `json:"content"`
	Name    param.Field[string]                                                                     `json:"name"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0) implementsPromptUpdateParamsPromptDataPromptChatMessageUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role = "system"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1 struct {
	Role    param.Field[PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role]         `json:"role,required"`
	Content param.Field[PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion] `json:"content"`
	Name    param.Field[string]                                                                             `json:"name"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1) implementsPromptUpdateParamsPromptDataPromptChatMessageUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role = "user"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray].
type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsPromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion()
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray []PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) ImplementsPromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray struct {
	Text     param.Field[string]                                                                                 `json:"text"`
	Type     param.Field[PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                                            `json:"image_url,required"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) implementsPromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

// Satisfied by
// [PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList],
// [PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray].
type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion interface {
	implementsPromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion()
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type param.Field[PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType] `json:"type,required"`
	Text param.Field[string]                                                                                                               `json:"text"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsPromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL param.Field[PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL] `json:"image_url,required"`
	Type     param.Field[PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType]     `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsPromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    param.Field[string]                                                                                                                      `json:"url,required"`
	Detail param.Field[PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail] `json:"detail"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText     PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "text"
	PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText, PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2 struct {
	Role         param.Field[PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role]         `json:"role,required"`
	Content      param.Field[string]                                                                             `json:"content"`
	FunctionCall param.Field[PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                                             `json:"name"`
	ToolCalls    param.Field[[]PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall]   `json:"tool_calls"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2) implementsPromptUpdateParamsPromptDataPromptChatMessageUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role = "assistant"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall struct {
	ID       param.Field[string]                                                                                  `json:"id,required"`
	Function param.Field[PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction] `json:"function,required"`
	Type     param.Field[PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType]     `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3 struct {
	Role       param.Field[PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role] `json:"role,required"`
	Content    param.Field[string]                                                                     `json:"content"`
	ToolCallID param.Field[string]                                                                     `json:"tool_call_id"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3) implementsPromptUpdateParamsPromptDataPromptChatMessageUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role = "tool"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4 struct {
	Name    param.Field[string]                                                                     `json:"name,required"`
	Role    param.Field[PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role] `json:"role,required"`
	Content param.Field[string]                                                                     `json:"content"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4) implementsPromptUpdateParamsPromptDataPromptChatMessageUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role = "function"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5 struct {
	Role    param.Field[PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role] `json:"role,required"`
	Content param.Field[string]                                                                     `json:"content"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5) implementsPromptUpdateParamsPromptDataPromptChatMessageUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role = "model"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesRole string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesRoleSystem    PromptUpdateParamsPromptDataPromptChatMessagesRole = "system"
	PromptUpdateParamsPromptDataPromptChatMessagesRoleUser      PromptUpdateParamsPromptDataPromptChatMessagesRole = "user"
	PromptUpdateParamsPromptDataPromptChatMessagesRoleAssistant PromptUpdateParamsPromptDataPromptChatMessagesRole = "assistant"
	PromptUpdateParamsPromptDataPromptChatMessagesRoleTool      PromptUpdateParamsPromptDataPromptChatMessagesRole = "tool"
	PromptUpdateParamsPromptDataPromptChatMessagesRoleFunction  PromptUpdateParamsPromptDataPromptChatMessagesRole = "function"
	PromptUpdateParamsPromptDataPromptChatMessagesRoleModel     PromptUpdateParamsPromptDataPromptChatMessagesRole = "model"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesRole) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesRoleSystem, PromptUpdateParamsPromptDataPromptChatMessagesRoleUser, PromptUpdateParamsPromptDataPromptChatMessagesRoleAssistant, PromptUpdateParamsPromptDataPromptChatMessagesRoleTool, PromptUpdateParamsPromptDataPromptChatMessagesRoleFunction, PromptUpdateParamsPromptDataPromptChatMessagesRoleModel:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatType string

const (
	PromptUpdateParamsPromptDataPromptChatTypeChat PromptUpdateParamsPromptDataPromptChatType = "chat"
)

func (r PromptUpdateParamsPromptDataPromptChatType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatTypeChat:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptNullableVariant struct {
}

func (r PromptUpdateParamsPromptDataPromptNullableVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptNullableVariant) implementsPromptUpdateParamsPromptDataPromptUnion() {
}

type PromptUpdateParamsPromptDataPromptType string

const (
	PromptUpdateParamsPromptDataPromptTypeCompletion PromptUpdateParamsPromptDataPromptType = "completion"
	PromptUpdateParamsPromptDataPromptTypeChat       PromptUpdateParamsPromptDataPromptType = "chat"
)

func (r PromptUpdateParamsPromptDataPromptType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptTypeCompletion, PromptUpdateParamsPromptDataPromptTypeChat:
		return true
	}
	return false
}

type PromptListParams struct {
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[PromptListParamsIDsUnion] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[int64] `query:"limit"`
	// Filter search results to within a particular organization
	OrgName param.Field[string] `query:"org_name"`
	// Name of the project to search for
	ProjectName param.Field[string] `query:"project_name"`
	// Name of the prompt to search for
	PromptName param.Field[string] `query:"prompt_name"`
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

// URLQuery serializes [PromptListParams]'s query parameters as `url.Values`.
func (r PromptListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [PromptListParamsIDsArray].
type PromptListParamsIDsUnion interface {
	ImplementsPromptListParamsIDsUnion()
}

type PromptListParamsIDsArray []string

func (r PromptListParamsIDsArray) ImplementsPromptListParamsIDsUnion() {}

type PromptReplaceParams struct {
	// Name of the prompt
	Name param.Field[string] `json:"name,required"`
	// Unique identifier for the project that the prompt belongs under
	ProjectID param.Field[string] `json:"project_id,required" format:"uuid"`
	// Unique identifier for the prompt
	Slug param.Field[string] `json:"slug,required"`
	// Textual description of the prompt
	Description param.Field[string] `json:"description"`
	// The prompt, model, and its parameters
	PromptData param.Field[PromptReplaceParamsPromptData] `json:"prompt_data"`
	// A list of tags for the prompt
	Tags param.Field[[]string] `json:"tags"`
}

func (r PromptReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The prompt, model, and its parameters
type PromptReplaceParamsPromptData struct {
	Options param.Field[PromptReplaceParamsPromptDataOptions]     `json:"options"`
	Origin  param.Field[PromptReplaceParamsPromptDataOrigin]      `json:"origin"`
	Prompt  param.Field[PromptReplaceParamsPromptDataPromptUnion] `json:"prompt"`
}

func (r PromptReplaceParamsPromptData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataOptions struct {
	Model    param.Field[string]                                          `json:"model"`
	Params   param.Field[PromptReplaceParamsPromptDataOptionsParamsUnion] `json:"params"`
	Position param.Field[string]                                          `json:"position"`
}

func (r PromptReplaceParamsPromptDataOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataOptionsParams struct {
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

func (r PromptReplaceParamsPromptDataOptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParams) implementsPromptReplaceParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by [PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParams],
// [PromptReplaceParamsPromptDataOptionsParamsAnthropicModelParams],
// [PromptReplaceParamsPromptDataOptionsParamsGoogleModelParams],
// [PromptReplaceParamsPromptDataOptionsParamsWindowAIModelParams],
// [PromptReplaceParamsPromptDataOptionsParamsJsCompletionParams],
// [PromptReplaceParamsPromptDataOptionsParams].
type PromptReplaceParamsPromptDataOptionsParamsUnion interface {
	implementsPromptReplaceParamsPromptDataOptionsParamsUnion()
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParams struct {
	FrequencyPenalty param.Field[float64]                                                                      `json:"frequency_penalty"`
	FunctionCall     param.Field[PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                                      `json:"max_tokens"`
	N                param.Field[float64]                                                                      `json:"n"`
	PresencePenalty  param.Field[float64]                                                                      `json:"presence_penalty"`
	ResponseFormat   param.Field[PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                                     `json:"stop"`
	Temperature      param.Field[float64]                                                                      `json:"temperature"`
	ToolChoice       param.Field[PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                                      `json:"top_p"`
	UseCache         param.Field[bool]                                                                         `json:"use_cache"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParams) implementsPromptReplaceParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by
// [PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString],
// [PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString],
// [PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject].
type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsPromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString string

const (
	PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString = "auto"
)

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallStringAuto:
		return true
	}
	return false
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallString) implementsPromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallObject) implementsPromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat struct {
	Type param.Field[PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType] `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType string

const (
	PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType = "json_object"
)

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Satisfied by
// [PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString],
// [PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString],
// [PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject].
type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsPromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString string

const (
	PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString = "auto"
)

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceStringAuto:
		return true
	}
	return false
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceString) implementsPromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject struct {
	Function param.Field[PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction] `json:"function,required"`
	Type     param.Field[PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType]     `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObject) implementsPromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType string

const (
	PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType = "function"
)

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceObjectTypeFunction:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataOptionsParamsAnthropicModelParams struct {
	MaxTokens   param.Field[float64] `json:"max_tokens,required"`
	Temperature param.Field[float64] `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64]  `json:"max_tokens_to_sample"`
	StopSequences     param.Field[[]string] `json:"stop_sequences"`
	TopK              param.Field[float64]  `json:"top_k"`
	TopP              param.Field[float64]  `json:"top_p"`
	UseCache          param.Field[bool]     `json:"use_cache"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsAnthropicModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsAnthropicModelParams) implementsPromptReplaceParamsPromptDataOptionsParamsUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsGoogleModelParams struct {
	MaxOutputTokens param.Field[float64] `json:"maxOutputTokens"`
	Temperature     param.Field[float64] `json:"temperature"`
	TopK            param.Field[float64] `json:"topK"`
	TopP            param.Field[float64] `json:"topP"`
	UseCache        param.Field[bool]    `json:"use_cache"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsGoogleModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsGoogleModelParams) implementsPromptReplaceParamsPromptDataOptionsParamsUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsWindowAIModelParams struct {
	Temperature param.Field[float64] `json:"temperature"`
	TopK        param.Field[float64] `json:"topK"`
	UseCache    param.Field[bool]    `json:"use_cache"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsWindowAIModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsWindowAIModelParams) implementsPromptReplaceParamsPromptDataOptionsParamsUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsJsCompletionParams struct {
	UseCache param.Field[bool] `json:"use_cache"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsJsCompletionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsJsCompletionParams) implementsPromptReplaceParamsPromptDataOptionsParamsUnion() {
}

type PromptReplaceParamsPromptDataOrigin struct {
	ProjectID     param.Field[string] `json:"project_id"`
	PromptID      param.Field[string] `json:"prompt_id"`
	PromptVersion param.Field[string] `json:"prompt_version"`
}

func (r PromptReplaceParamsPromptDataOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataPrompt struct {
	Type     param.Field[PromptReplaceParamsPromptDataPromptType] `json:"type"`
	Content  param.Field[string]                                  `json:"content"`
	Messages param.Field[interface{}]                             `json:"messages,required"`
	Tools    param.Field[string]                                  `json:"tools"`
}

func (r PromptReplaceParamsPromptDataPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPrompt) implementsPromptReplaceParamsPromptDataPromptUnion() {}

// Satisfied by [PromptReplaceParamsPromptDataPromptCompletion],
// [PromptReplaceParamsPromptDataPromptChat],
// [PromptReplaceParamsPromptDataPromptNullableVariant],
// [PromptReplaceParamsPromptDataPrompt].
type PromptReplaceParamsPromptDataPromptUnion interface {
	implementsPromptReplaceParamsPromptDataPromptUnion()
}

type PromptReplaceParamsPromptDataPromptCompletion struct {
	Content param.Field[string]                                            `json:"content,required"`
	Type    param.Field[PromptReplaceParamsPromptDataPromptCompletionType] `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataPromptCompletion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptCompletion) implementsPromptReplaceParamsPromptDataPromptUnion() {
}

type PromptReplaceParamsPromptDataPromptCompletionType string

const (
	PromptReplaceParamsPromptDataPromptCompletionTypeCompletion PromptReplaceParamsPromptDataPromptCompletionType = "completion"
)

func (r PromptReplaceParamsPromptDataPromptCompletionType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptCompletionTypeCompletion:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChat struct {
	Messages param.Field[[]PromptReplaceParamsPromptDataPromptChatMessageUnion] `json:"messages,required"`
	Type     param.Field[PromptReplaceParamsPromptDataPromptChatType]           `json:"type,required"`
	Tools    param.Field[string]                                                `json:"tools"`
}

func (r PromptReplaceParamsPromptDataPromptChat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChat) implementsPromptReplaceParamsPromptDataPromptUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessage struct {
	Content      param.Field[interface{}]                                         `json:"content,required"`
	Role         param.Field[PromptReplaceParamsPromptDataPromptChatMessagesRole] `json:"role,required"`
	Name         param.Field[string]                                              `json:"name"`
	FunctionCall param.Field[interface{}]                                         `json:"function_call,required"`
	ToolCalls    param.Field[interface{}]                                         `json:"tool_calls,required"`
	ToolCallID   param.Field[string]                                              `json:"tool_call_id"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessage) implementsPromptReplaceParamsPromptDataPromptChatMessageUnion() {
}

// Satisfied by
// [PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0],
// [PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1],
// [PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2],
// [PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3],
// [PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4],
// [PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5],
// [PromptReplaceParamsPromptDataPromptChatMessage].
type PromptReplaceParamsPromptDataPromptChatMessageUnion interface {
	implementsPromptReplaceParamsPromptDataPromptChatMessageUnion()
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0 struct {
	Role    param.Field[PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role] `json:"role,required"`
	Content param.Field[string]                                                                      `json:"content"`
	Name    param.Field[string]                                                                      `json:"name"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0) implementsPromptReplaceParamsPromptDataPromptChatMessageUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role = "system"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1 struct {
	Role    param.Field[PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role]         `json:"role,required"`
	Content param.Field[PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion] `json:"content"`
	Name    param.Field[string]                                                                              `json:"name"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1) implementsPromptReplaceParamsPromptDataPromptChatMessageUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role = "user"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray].
type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsPromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion()
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray []PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) ImplementsPromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray struct {
	Text     param.Field[string]                                                                                  `json:"text"`
	Type     param.Field[PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                                             `json:"image_url,required"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray) implementsPromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

// Satisfied by
// [PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList],
// [PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArray].
type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion interface {
	implementsPromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion()
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type param.Field[PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType] `json:"type,required"`
	Text param.Field[string]                                                                                                                `json:"text"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsPromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL param.Field[PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL] `json:"image_url,required"`
	Type     param.Field[PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType]     `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsPromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    param.Field[string]                                                                                                                       `json:"url,required"`
	Detail param.Field[PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail] `json:"detail"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText     PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "text"
	PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeText, PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2 struct {
	Role         param.Field[PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role]         `json:"role,required"`
	Content      param.Field[string]                                                                              `json:"content"`
	FunctionCall param.Field[PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                                              `json:"name"`
	ToolCalls    param.Field[[]PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall]   `json:"tool_calls"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2) implementsPromptReplaceParamsPromptDataPromptChatMessageUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role = "assistant"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2FunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall struct {
	ID       param.Field[string]                                                                                   `json:"id,required"`
	Function param.Field[PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction] `json:"function,required"`
	Type     param.Field[PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType]     `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3 struct {
	Role       param.Field[PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role] `json:"role,required"`
	Content    param.Field[string]                                                                      `json:"content"`
	ToolCallID param.Field[string]                                                                      `json:"tool_call_id"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3) implementsPromptReplaceParamsPromptDataPromptChatMessageUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role = "tool"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4 struct {
	Name    param.Field[string]                                                                      `json:"name,required"`
	Role    param.Field[PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role] `json:"role,required"`
	Content param.Field[string]                                                                      `json:"content"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4) implementsPromptReplaceParamsPromptDataPromptChatMessageUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role = "function"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5 struct {
	Role    param.Field[PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role] `json:"role,required"`
	Content param.Field[string]                                                                      `json:"content"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5) implementsPromptReplaceParamsPromptDataPromptChatMessageUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role = "model"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesRole string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesRoleSystem    PromptReplaceParamsPromptDataPromptChatMessagesRole = "system"
	PromptReplaceParamsPromptDataPromptChatMessagesRoleUser      PromptReplaceParamsPromptDataPromptChatMessagesRole = "user"
	PromptReplaceParamsPromptDataPromptChatMessagesRoleAssistant PromptReplaceParamsPromptDataPromptChatMessagesRole = "assistant"
	PromptReplaceParamsPromptDataPromptChatMessagesRoleTool      PromptReplaceParamsPromptDataPromptChatMessagesRole = "tool"
	PromptReplaceParamsPromptDataPromptChatMessagesRoleFunction  PromptReplaceParamsPromptDataPromptChatMessagesRole = "function"
	PromptReplaceParamsPromptDataPromptChatMessagesRoleModel     PromptReplaceParamsPromptDataPromptChatMessagesRole = "model"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesRole) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesRoleSystem, PromptReplaceParamsPromptDataPromptChatMessagesRoleUser, PromptReplaceParamsPromptDataPromptChatMessagesRoleAssistant, PromptReplaceParamsPromptDataPromptChatMessagesRoleTool, PromptReplaceParamsPromptDataPromptChatMessagesRoleFunction, PromptReplaceParamsPromptDataPromptChatMessagesRoleModel:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatType string

const (
	PromptReplaceParamsPromptDataPromptChatTypeChat PromptReplaceParamsPromptDataPromptChatType = "chat"
)

func (r PromptReplaceParamsPromptDataPromptChatType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatTypeChat:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptNullableVariant struct {
}

func (r PromptReplaceParamsPromptDataPromptNullableVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptNullableVariant) implementsPromptReplaceParamsPromptDataPromptUnion() {
}

type PromptReplaceParamsPromptDataPromptType string

const (
	PromptReplaceParamsPromptDataPromptTypeCompletion PromptReplaceParamsPromptDataPromptType = "completion"
	PromptReplaceParamsPromptDataPromptTypeChat       PromptReplaceParamsPromptDataPromptType = "chat"
)

func (r PromptReplaceParamsPromptDataPromptType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptTypeCompletion, PromptReplaceParamsPromptDataPromptTypeChat:
		return true
	}
	return false
}
