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
	Model    string                             `json:"model"`
	Params   PromptPromptDataOptionsParamsUnion `json:"params"`
	Position string                             `json:"position"`
	JSON     promptPromptDataOptionsJSON        `json:"-"`
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

// Union satisfied by [PromptPromptDataOptionsParamsOpenAIModelParams],
// [PromptPromptDataOptionsParamsAnthropicModelParams],
// [PromptPromptDataOptionsParamsGoogleModelParams],
// [PromptPromptDataOptionsParamsWindowAIModelParams] or
// [PromptPromptDataOptionsParamsJsCompletionParams].
type PromptPromptDataOptionsParamsUnion interface {
	implementsPromptPromptDataOptionsParamsUnion()
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

func (r PromptPromptDataOptionsParamsOpenAIModelParams) implementsPromptPromptDataOptionsParamsUnion() {
}

// Union satisfied by
// [PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto],
// [PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone] or
// [PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction].
type PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsPromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction{}),
		},
	)
}

type PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto string

const (
	PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto = "auto"
)

func (r PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto:
		return true
	}
	return false
}

func (r PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) implementsPromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone string

const (
	PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone = "none"
)

func (r PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone:
		return true
	}
	return false
}

func (r PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) implementsPromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction struct {
	Name string                                                                 `json:"name,required"`
	JSON promptPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionJSON `json:"-"`
}

// promptPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionJSON contains
// the JSON metadata for the struct
// [PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction]
type promptPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) implementsPromptPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
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
// [PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto],
// [PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone] or
// [PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction].
type PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsPromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction{}),
		},
	)
}

type PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto string

const (
	PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto = "auto"
)

func (r PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto:
		return true
	}
	return false
}

func (r PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) implementsPromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone string

const (
	PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone = "none"
)

func (r PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone:
		return true
	}
	return false
}

func (r PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) implementsPromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction struct {
	Function PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction `json:"function,required"`
	Type     PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType     `json:"type,required"`
	JSON     promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionJSON     `json:"-"`
}

// promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionJSON contains
// the JSON metadata for the struct
// [PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction]
type promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionJSON struct {
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) implementsPromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction struct {
	Name string                                                                       `json:"name,required"`
	JSON promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionJSON `json:"-"`
}

// promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionJSON
// contains the JSON metadata for the struct
// [PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction]
type promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType string

const (
	PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType = "function"
)

func (r PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction:
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

func (r PromptPromptDataOptionsParamsAnthropicModelParams) implementsPromptPromptDataOptionsParamsUnion() {
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

func (r PromptPromptDataOptionsParamsGoogleModelParams) implementsPromptPromptDataOptionsParamsUnion() {
}

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

func (r PromptPromptDataOptionsParamsWindowAIModelParams) implementsPromptPromptDataOptionsParamsUnion() {
}

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

func (r PromptPromptDataOptionsParamsJsCompletionParams) implementsPromptPromptDataOptionsParamsUnion() {
}

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
	// [PromptPromptDataPromptChatMessagesUserContentUnion].
	Content interface{}                            `json:"content,required"`
	Role    PromptPromptDataPromptChatMessagesRole `json:"role,required"`
	Name    string                                 `json:"name"`
	// This field can have the runtime type of
	// [PromptPromptDataPromptChatMessagesAssistantFunctionCall].
	FunctionCall interface{} `json:"function_call,required"`
	// This field can have the runtime type of
	// [[]PromptPromptDataPromptChatMessagesAssistantToolCall].
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
// [PromptPromptDataPromptChatMessagesSystem],
// [PromptPromptDataPromptChatMessagesUser],
// [PromptPromptDataPromptChatMessagesAssistant],
// [PromptPromptDataPromptChatMessagesTool],
// [PromptPromptDataPromptChatMessagesFunction],
// [PromptPromptDataPromptChatMessagesFallback].
func (r PromptPromptDataPromptChatMessage) AsUnion() PromptPromptDataPromptChatMessagesUnion {
	return r.union
}

// Union satisfied by [PromptPromptDataPromptChatMessagesSystem],
// [PromptPromptDataPromptChatMessagesUser],
// [PromptPromptDataPromptChatMessagesAssistant],
// [PromptPromptDataPromptChatMessagesTool],
// [PromptPromptDataPromptChatMessagesFunction] or
// [PromptPromptDataPromptChatMessagesFallback].
type PromptPromptDataPromptChatMessagesUnion interface {
	implementsPromptPromptDataPromptChatMessage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataPromptChatMessagesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesSystem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesUser{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesAssistant{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesTool{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesFunction{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesFallback{}),
		},
	)
}

type PromptPromptDataPromptChatMessagesSystem struct {
	Role    PromptPromptDataPromptChatMessagesSystemRole `json:"role,required"`
	Content string                                       `json:"content"`
	Name    string                                       `json:"name"`
	JSON    promptPromptDataPromptChatMessagesSystemJSON `json:"-"`
}

// promptPromptDataPromptChatMessagesSystemJSON contains the JSON metadata for the
// struct [PromptPromptDataPromptChatMessagesSystem]
type promptPromptDataPromptChatMessagesSystemJSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesSystem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesSystemJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesSystem) implementsPromptPromptDataPromptChatMessage() {}

type PromptPromptDataPromptChatMessagesSystemRole string

const (
	PromptPromptDataPromptChatMessagesSystemRoleSystem PromptPromptDataPromptChatMessagesSystemRole = "system"
)

func (r PromptPromptDataPromptChatMessagesSystemRole) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesSystemRoleSystem:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesUser struct {
	Role    PromptPromptDataPromptChatMessagesUserRole         `json:"role,required"`
	Content PromptPromptDataPromptChatMessagesUserContentUnion `json:"content"`
	Name    string                                             `json:"name"`
	JSON    promptPromptDataPromptChatMessagesUserJSON         `json:"-"`
}

// promptPromptDataPromptChatMessagesUserJSON contains the JSON metadata for the
// struct [PromptPromptDataPromptChatMessagesUser]
type promptPromptDataPromptChatMessagesUserJSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesUserJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesUser) implementsPromptPromptDataPromptChatMessage() {}

type PromptPromptDataPromptChatMessagesUserRole string

const (
	PromptPromptDataPromptChatMessagesUserRoleUser PromptPromptDataPromptChatMessagesUserRole = "user"
)

func (r PromptPromptDataPromptChatMessagesUserRole) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesUserRoleUser:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or
// [PromptPromptDataPromptChatMessagesUserContentArray].
type PromptPromptDataPromptChatMessagesUserContentUnion interface {
	ImplementsPromptPromptDataPromptChatMessagesUserContentUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataPromptChatMessagesUserContentUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesUserContentArray{}),
		},
	)
}

type PromptPromptDataPromptChatMessagesUserContentArray []PromptPromptDataPromptChatMessagesUserContentArrayItem

func (r PromptPromptDataPromptChatMessagesUserContentArray) ImplementsPromptPromptDataPromptChatMessagesUserContentUnion() {
}

type PromptPromptDataPromptChatMessagesUserContentArrayItem struct {
	Text string                                                 `json:"text"`
	Type PromptPromptDataPromptChatMessagesUserContentArrayType `json:"type,required"`
	// This field can have the runtime type of
	// [PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURL].
	ImageURL interface{}                                                `json:"image_url,required"`
	JSON     promptPromptDataPromptChatMessagesUserContentArrayItemJSON `json:"-"`
	union    PromptPromptDataPromptChatMessagesUserContentArrayUnionItem
}

// promptPromptDataPromptChatMessagesUserContentArrayItemJSON contains the JSON
// metadata for the struct [PromptPromptDataPromptChatMessagesUserContentArrayItem]
type promptPromptDataPromptChatMessagesUserContentArrayItemJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	ImageURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r promptPromptDataPromptChatMessagesUserContentArrayItemJSON) RawJSON() string {
	return r.raw
}

func (r *PromptPromptDataPromptChatMessagesUserContentArrayItem) UnmarshalJSON(data []byte) (err error) {
	*r = PromptPromptDataPromptChatMessagesUserContentArrayItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PromptPromptDataPromptChatMessagesUserContentArrayUnionItem]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PromptPromptDataPromptChatMessagesUserContentArrayText],
// [PromptPromptDataPromptChatMessagesUserContentArrayImageURL].
func (r PromptPromptDataPromptChatMessagesUserContentArrayItem) AsUnion() PromptPromptDataPromptChatMessagesUserContentArrayUnionItem {
	return r.union
}

// Union satisfied by [PromptPromptDataPromptChatMessagesUserContentArrayText] or
// [PromptPromptDataPromptChatMessagesUserContentArrayImageURL].
type PromptPromptDataPromptChatMessagesUserContentArrayUnionItem interface {
	implementsPromptPromptDataPromptChatMessagesUserContentArrayItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataPromptChatMessagesUserContentArrayUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesUserContentArrayText{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesUserContentArrayImageURL{}),
		},
	)
}

type PromptPromptDataPromptChatMessagesUserContentArrayText struct {
	Type PromptPromptDataPromptChatMessagesUserContentArrayTextType `json:"type,required"`
	Text string                                                     `json:"text"`
	JSON promptPromptDataPromptChatMessagesUserContentArrayTextJSON `json:"-"`
}

// promptPromptDataPromptChatMessagesUserContentArrayTextJSON contains the JSON
// metadata for the struct [PromptPromptDataPromptChatMessagesUserContentArrayText]
type promptPromptDataPromptChatMessagesUserContentArrayTextJSON struct {
	Type        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesUserContentArrayText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesUserContentArrayTextJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesUserContentArrayText) implementsPromptPromptDataPromptChatMessagesUserContentArrayItem() {
}

type PromptPromptDataPromptChatMessagesUserContentArrayTextType string

const (
	PromptPromptDataPromptChatMessagesUserContentArrayTextTypeText PromptPromptDataPromptChatMessagesUserContentArrayTextType = "text"
)

func (r PromptPromptDataPromptChatMessagesUserContentArrayTextType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesUserContentArrayTextTypeText:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesUserContentArrayImageURL struct {
	ImageURL PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURL `json:"image_url,required"`
	Type     PromptPromptDataPromptChatMessagesUserContentArrayImageURLType     `json:"type,required"`
	JSON     promptPromptDataPromptChatMessagesUserContentArrayImageURLJSON     `json:"-"`
}

// promptPromptDataPromptChatMessagesUserContentArrayImageURLJSON contains the JSON
// metadata for the struct
// [PromptPromptDataPromptChatMessagesUserContentArrayImageURL]
type promptPromptDataPromptChatMessagesUserContentArrayImageURLJSON struct {
	ImageURL    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesUserContentArrayImageURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesUserContentArrayImageURLJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesUserContentArrayImageURL) implementsPromptPromptDataPromptChatMessagesUserContentArrayItem() {
}

type PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURL struct {
	URL    string                                                                   `json:"url,required"`
	Detail PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail `json:"detail"`
	JSON   promptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLJSON   `json:"-"`
}

// promptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLJSON contains
// the JSON metadata for the struct
// [PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURL]
type promptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLJSON struct {
	URL         apijson.Field
	Detail      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail string

const (
	PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "auto"
	PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow  PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "low"
	PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "high"
)

func (r PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto, PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow, PromptPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesUserContentArrayImageURLType string

const (
	PromptPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL PromptPromptDataPromptChatMessagesUserContentArrayImageURLType = "image_url"
)

func (r PromptPromptDataPromptChatMessagesUserContentArrayImageURLType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesUserContentArrayType string

const (
	PromptPromptDataPromptChatMessagesUserContentArrayTypeText     PromptPromptDataPromptChatMessagesUserContentArrayType = "text"
	PromptPromptDataPromptChatMessagesUserContentArrayTypeImageURL PromptPromptDataPromptChatMessagesUserContentArrayType = "image_url"
)

func (r PromptPromptDataPromptChatMessagesUserContentArrayType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesUserContentArrayTypeText, PromptPromptDataPromptChatMessagesUserContentArrayTypeImageURL:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesAssistant struct {
	Role         PromptPromptDataPromptChatMessagesAssistantRole         `json:"role,required"`
	Content      string                                                  `json:"content,nullable"`
	FunctionCall PromptPromptDataPromptChatMessagesAssistantFunctionCall `json:"function_call"`
	Name         string                                                  `json:"name"`
	ToolCalls    []PromptPromptDataPromptChatMessagesAssistantToolCall   `json:"tool_calls"`
	JSON         promptPromptDataPromptChatMessagesAssistantJSON         `json:"-"`
}

// promptPromptDataPromptChatMessagesAssistantJSON contains the JSON metadata for
// the struct [PromptPromptDataPromptChatMessagesAssistant]
type promptPromptDataPromptChatMessagesAssistantJSON struct {
	Role         apijson.Field
	Content      apijson.Field
	FunctionCall apijson.Field
	Name         apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesAssistant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesAssistantJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesAssistant) implementsPromptPromptDataPromptChatMessage() {}

type PromptPromptDataPromptChatMessagesAssistantRole string

const (
	PromptPromptDataPromptChatMessagesAssistantRoleAssistant PromptPromptDataPromptChatMessagesAssistantRole = "assistant"
)

func (r PromptPromptDataPromptChatMessagesAssistantRole) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesAssistantRoleAssistant:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesAssistantFunctionCall struct {
	Arguments string                                                      `json:"arguments,required"`
	Name      string                                                      `json:"name,required"`
	JSON      promptPromptDataPromptChatMessagesAssistantFunctionCallJSON `json:"-"`
}

// promptPromptDataPromptChatMessagesAssistantFunctionCallJSON contains the JSON
// metadata for the struct
// [PromptPromptDataPromptChatMessagesAssistantFunctionCall]
type promptPromptDataPromptChatMessagesAssistantFunctionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesAssistantFunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesAssistantFunctionCallJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataPromptChatMessagesAssistantToolCall struct {
	ID       string                                                       `json:"id,required"`
	Function PromptPromptDataPromptChatMessagesAssistantToolCallsFunction `json:"function,required"`
	Type     PromptPromptDataPromptChatMessagesAssistantToolCallsType     `json:"type,required"`
	JSON     promptPromptDataPromptChatMessagesAssistantToolCallJSON      `json:"-"`
}

// promptPromptDataPromptChatMessagesAssistantToolCallJSON contains the JSON
// metadata for the struct [PromptPromptDataPromptChatMessagesAssistantToolCall]
type promptPromptDataPromptChatMessagesAssistantToolCallJSON struct {
	ID          apijson.Field
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesAssistantToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesAssistantToolCallJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataPromptChatMessagesAssistantToolCallsFunction struct {
	Arguments string                                                           `json:"arguments,required"`
	Name      string                                                           `json:"name,required"`
	JSON      promptPromptDataPromptChatMessagesAssistantToolCallsFunctionJSON `json:"-"`
}

// promptPromptDataPromptChatMessagesAssistantToolCallsFunctionJSON contains the
// JSON metadata for the struct
// [PromptPromptDataPromptChatMessagesAssistantToolCallsFunction]
type promptPromptDataPromptChatMessagesAssistantToolCallsFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesAssistantToolCallsFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesAssistantToolCallsFunctionJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataPromptChatMessagesAssistantToolCallsType string

const (
	PromptPromptDataPromptChatMessagesAssistantToolCallsTypeFunction PromptPromptDataPromptChatMessagesAssistantToolCallsType = "function"
)

func (r PromptPromptDataPromptChatMessagesAssistantToolCallsType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesAssistantToolCallsTypeFunction:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesTool struct {
	Role       PromptPromptDataPromptChatMessagesToolRole `json:"role,required"`
	Content    string                                     `json:"content"`
	ToolCallID string                                     `json:"tool_call_id"`
	JSON       promptPromptDataPromptChatMessagesToolJSON `json:"-"`
}

// promptPromptDataPromptChatMessagesToolJSON contains the JSON metadata for the
// struct [PromptPromptDataPromptChatMessagesTool]
type promptPromptDataPromptChatMessagesToolJSON struct {
	Role        apijson.Field
	Content     apijson.Field
	ToolCallID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesToolJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesTool) implementsPromptPromptDataPromptChatMessage() {}

type PromptPromptDataPromptChatMessagesToolRole string

const (
	PromptPromptDataPromptChatMessagesToolRoleTool PromptPromptDataPromptChatMessagesToolRole = "tool"
)

func (r PromptPromptDataPromptChatMessagesToolRole) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesToolRoleTool:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesFunction struct {
	Name    string                                         `json:"name,required"`
	Role    PromptPromptDataPromptChatMessagesFunctionRole `json:"role,required"`
	Content string                                         `json:"content"`
	JSON    promptPromptDataPromptChatMessagesFunctionJSON `json:"-"`
}

// promptPromptDataPromptChatMessagesFunctionJSON contains the JSON metadata for
// the struct [PromptPromptDataPromptChatMessagesFunction]
type promptPromptDataPromptChatMessagesFunctionJSON struct {
	Name        apijson.Field
	Role        apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesFunctionJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesFunction) implementsPromptPromptDataPromptChatMessage() {}

type PromptPromptDataPromptChatMessagesFunctionRole string

const (
	PromptPromptDataPromptChatMessagesFunctionRoleFunction PromptPromptDataPromptChatMessagesFunctionRole = "function"
)

func (r PromptPromptDataPromptChatMessagesFunctionRole) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesFunctionRoleFunction:
		return true
	}
	return false
}

type PromptPromptDataPromptChatMessagesFallback struct {
	Role    PromptPromptDataPromptChatMessagesFallbackRole `json:"role,required"`
	Content string                                         `json:"content,nullable"`
	JSON    promptPromptDataPromptChatMessagesFallbackJSON `json:"-"`
}

// promptPromptDataPromptChatMessagesFallbackJSON contains the JSON metadata for
// the struct [PromptPromptDataPromptChatMessagesFallback]
type promptPromptDataPromptChatMessagesFallbackJSON struct {
	Role        apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesFallback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesFallbackJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesFallback) implementsPromptPromptDataPromptChatMessage() {}

type PromptPromptDataPromptChatMessagesFallbackRole string

const (
	PromptPromptDataPromptChatMessagesFallbackRoleModel PromptPromptDataPromptChatMessagesFallbackRole = "model"
)

func (r PromptPromptDataPromptChatMessagesFallbackRole) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesFallbackRoleModel:
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

// Satisfied by [PromptNewParamsPromptDataOptionsParamsOpenAIModelParams],
// [PromptNewParamsPromptDataOptionsParamsAnthropicModelParams],
// [PromptNewParamsPromptDataOptionsParamsGoogleModelParams],
// [PromptNewParamsPromptDataOptionsParamsWindowAIModelParams],
// [PromptNewParamsPromptDataOptionsParamsJsCompletionParams].
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
// [PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto],
// [PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone],
// [PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction].
type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsPromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto string

const (
	PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto = "auto"
)

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto:
		return true
	}
	return false
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) implementsPromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone string

const (
	PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone = "none"
)

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone:
		return true
	}
	return false
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) implementsPromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) implementsPromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
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
// [PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto],
// [PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone],
// [PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction].
type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsPromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto string

const (
	PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto = "auto"
)

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto:
		return true
	}
	return false
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) implementsPromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone string

const (
	PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone = "none"
)

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone:
		return true
	}
	return false
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) implementsPromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction struct {
	Function param.Field[PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction] `json:"function,required"`
	Type     param.Field[PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType]     `json:"type,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) implementsPromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType string

const (
	PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType = "function"
)

func (r PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction:
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

// Satisfied by [PromptNewParamsPromptDataPromptChatMessagesSystem],
// [PromptNewParamsPromptDataPromptChatMessagesUser],
// [PromptNewParamsPromptDataPromptChatMessagesAssistant],
// [PromptNewParamsPromptDataPromptChatMessagesTool],
// [PromptNewParamsPromptDataPromptChatMessagesFunction],
// [PromptNewParamsPromptDataPromptChatMessagesFallback],
// [PromptNewParamsPromptDataPromptChatMessage].
type PromptNewParamsPromptDataPromptChatMessageUnion interface {
	implementsPromptNewParamsPromptDataPromptChatMessageUnion()
}

type PromptNewParamsPromptDataPromptChatMessagesSystem struct {
	Role    param.Field[PromptNewParamsPromptDataPromptChatMessagesSystemRole] `json:"role,required"`
	Content param.Field[string]                                                `json:"content"`
	Name    param.Field[string]                                                `json:"name"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesSystem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesSystem) implementsPromptNewParamsPromptDataPromptChatMessageUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesSystemRole string

const (
	PromptNewParamsPromptDataPromptChatMessagesSystemRoleSystem PromptNewParamsPromptDataPromptChatMessagesSystemRole = "system"
)

func (r PromptNewParamsPromptDataPromptChatMessagesSystemRole) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesSystemRoleSystem:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesUser struct {
	Role    param.Field[PromptNewParamsPromptDataPromptChatMessagesUserRole]         `json:"role,required"`
	Content param.Field[PromptNewParamsPromptDataPromptChatMessagesUserContentUnion] `json:"content"`
	Name    param.Field[string]                                                      `json:"name"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesUser) implementsPromptNewParamsPromptDataPromptChatMessageUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesUserRole string

const (
	PromptNewParamsPromptDataPromptChatMessagesUserRoleUser PromptNewParamsPromptDataPromptChatMessagesUserRole = "user"
)

func (r PromptNewParamsPromptDataPromptChatMessagesUserRole) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesUserRoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [PromptNewParamsPromptDataPromptChatMessagesUserContentArray].
type PromptNewParamsPromptDataPromptChatMessagesUserContentUnion interface {
	ImplementsPromptNewParamsPromptDataPromptChatMessagesUserContentUnion()
}

type PromptNewParamsPromptDataPromptChatMessagesUserContentArray []PromptNewParamsPromptDataPromptChatMessagesUserContentArrayUnion

func (r PromptNewParamsPromptDataPromptChatMessagesUserContentArray) ImplementsPromptNewParamsPromptDataPromptChatMessagesUserContentUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesUserContentArray struct {
	Text     param.Field[string]                                                          `json:"text"`
	Type     param.Field[PromptNewParamsPromptDataPromptChatMessagesUserContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                     `json:"image_url,required"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesUserContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesUserContentArray) implementsPromptNewParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

// Satisfied by [PromptNewParamsPromptDataPromptChatMessagesUserContentArrayText],
// [PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURL],
// [PromptNewParamsPromptDataPromptChatMessagesUserContentArray].
type PromptNewParamsPromptDataPromptChatMessagesUserContentArrayUnion interface {
	implementsPromptNewParamsPromptDataPromptChatMessagesUserContentArrayUnion()
}

type PromptNewParamsPromptDataPromptChatMessagesUserContentArrayText struct {
	Type param.Field[PromptNewParamsPromptDataPromptChatMessagesUserContentArrayTextType] `json:"type,required"`
	Text param.Field[string]                                                              `json:"text"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesUserContentArrayText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesUserContentArrayText) implementsPromptNewParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesUserContentArrayTextType string

const (
	PromptNewParamsPromptDataPromptChatMessagesUserContentArrayTextTypeText PromptNewParamsPromptDataPromptChatMessagesUserContentArrayTextType = "text"
)

func (r PromptNewParamsPromptDataPromptChatMessagesUserContentArrayTextType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesUserContentArrayTextTypeText:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURL struct {
	ImageURL param.Field[PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL] `json:"image_url,required"`
	Type     param.Field[PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLType]     `json:"type,required"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURL) implementsPromptNewParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL struct {
	URL    param.Field[string]                                                                            `json:"url,required"`
	Detail param.Field[PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail] `json:"detail"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail string

const (
	PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "auto"
	PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow  PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "low"
	PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "high"
)

func (r PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto, PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow, PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLType string

const (
	PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLType = "image_url"
)

func (r PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesUserContentArrayType string

const (
	PromptNewParamsPromptDataPromptChatMessagesUserContentArrayTypeText     PromptNewParamsPromptDataPromptChatMessagesUserContentArrayType = "text"
	PromptNewParamsPromptDataPromptChatMessagesUserContentArrayTypeImageURL PromptNewParamsPromptDataPromptChatMessagesUserContentArrayType = "image_url"
)

func (r PromptNewParamsPromptDataPromptChatMessagesUserContentArrayType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesUserContentArrayTypeText, PromptNewParamsPromptDataPromptChatMessagesUserContentArrayTypeImageURL:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesAssistant struct {
	Role         param.Field[PromptNewParamsPromptDataPromptChatMessagesAssistantRole]         `json:"role,required"`
	Content      param.Field[string]                                                           `json:"content"`
	FunctionCall param.Field[PromptNewParamsPromptDataPromptChatMessagesAssistantFunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                           `json:"name"`
	ToolCalls    param.Field[[]PromptNewParamsPromptDataPromptChatMessagesAssistantToolCall]   `json:"tool_calls"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesAssistant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesAssistant) implementsPromptNewParamsPromptDataPromptChatMessageUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesAssistantRole string

const (
	PromptNewParamsPromptDataPromptChatMessagesAssistantRoleAssistant PromptNewParamsPromptDataPromptChatMessagesAssistantRole = "assistant"
)

func (r PromptNewParamsPromptDataPromptChatMessagesAssistantRole) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesAssistantRoleAssistant:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesAssistantFunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesAssistantFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataPromptChatMessagesAssistantToolCall struct {
	ID       param.Field[string]                                                                `json:"id,required"`
	Function param.Field[PromptNewParamsPromptDataPromptChatMessagesAssistantToolCallsFunction] `json:"function,required"`
	Type     param.Field[PromptNewParamsPromptDataPromptChatMessagesAssistantToolCallsType]     `json:"type,required"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesAssistantToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataPromptChatMessagesAssistantToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesAssistantToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataPromptChatMessagesAssistantToolCallsType string

const (
	PromptNewParamsPromptDataPromptChatMessagesAssistantToolCallsTypeFunction PromptNewParamsPromptDataPromptChatMessagesAssistantToolCallsType = "function"
)

func (r PromptNewParamsPromptDataPromptChatMessagesAssistantToolCallsType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesAssistantToolCallsTypeFunction:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesTool struct {
	Role       param.Field[PromptNewParamsPromptDataPromptChatMessagesToolRole] `json:"role,required"`
	Content    param.Field[string]                                              `json:"content"`
	ToolCallID param.Field[string]                                              `json:"tool_call_id"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesTool) implementsPromptNewParamsPromptDataPromptChatMessageUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesToolRole string

const (
	PromptNewParamsPromptDataPromptChatMessagesToolRoleTool PromptNewParamsPromptDataPromptChatMessagesToolRole = "tool"
)

func (r PromptNewParamsPromptDataPromptChatMessagesToolRole) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesToolRoleTool:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesFunction struct {
	Name    param.Field[string]                                                  `json:"name,required"`
	Role    param.Field[PromptNewParamsPromptDataPromptChatMessagesFunctionRole] `json:"role,required"`
	Content param.Field[string]                                                  `json:"content"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesFunction) implementsPromptNewParamsPromptDataPromptChatMessageUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesFunctionRole string

const (
	PromptNewParamsPromptDataPromptChatMessagesFunctionRoleFunction PromptNewParamsPromptDataPromptChatMessagesFunctionRole = "function"
)

func (r PromptNewParamsPromptDataPromptChatMessagesFunctionRole) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesFunctionRoleFunction:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptChatMessagesFallback struct {
	Role    param.Field[PromptNewParamsPromptDataPromptChatMessagesFallbackRole] `json:"role,required"`
	Content param.Field[string]                                                  `json:"content"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesFallback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesFallback) implementsPromptNewParamsPromptDataPromptChatMessageUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesFallbackRole string

const (
	PromptNewParamsPromptDataPromptChatMessagesFallbackRoleModel PromptNewParamsPromptDataPromptChatMessagesFallbackRole = "model"
)

func (r PromptNewParamsPromptDataPromptChatMessagesFallbackRole) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesFallbackRoleModel:
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

// Satisfied by [PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParams],
// [PromptUpdateParamsPromptDataOptionsParamsAnthropicModelParams],
// [PromptUpdateParamsPromptDataOptionsParamsGoogleModelParams],
// [PromptUpdateParamsPromptDataOptionsParamsWindowAIModelParams],
// [PromptUpdateParamsPromptDataOptionsParamsJsCompletionParams].
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
// [PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto],
// [PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone],
// [PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction].
type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsPromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto string

const (
	PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto = "auto"
)

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto:
		return true
	}
	return false
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) implementsPromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone string

const (
	PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone = "none"
)

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone:
		return true
	}
	return false
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) implementsPromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) implementsPromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
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
// [PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto],
// [PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone],
// [PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction].
type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsPromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto string

const (
	PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto = "auto"
)

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto:
		return true
	}
	return false
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) implementsPromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone string

const (
	PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone = "none"
)

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone:
		return true
	}
	return false
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) implementsPromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction struct {
	Function param.Field[PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction] `json:"function,required"`
	Type     param.Field[PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType]     `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) implementsPromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType string

const (
	PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType = "function"
)

func (r PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction:
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

// Satisfied by [PromptUpdateParamsPromptDataPromptChatMessagesSystem],
// [PromptUpdateParamsPromptDataPromptChatMessagesUser],
// [PromptUpdateParamsPromptDataPromptChatMessagesAssistant],
// [PromptUpdateParamsPromptDataPromptChatMessagesTool],
// [PromptUpdateParamsPromptDataPromptChatMessagesFunction],
// [PromptUpdateParamsPromptDataPromptChatMessagesFallback],
// [PromptUpdateParamsPromptDataPromptChatMessage].
type PromptUpdateParamsPromptDataPromptChatMessageUnion interface {
	implementsPromptUpdateParamsPromptDataPromptChatMessageUnion()
}

type PromptUpdateParamsPromptDataPromptChatMessagesSystem struct {
	Role    param.Field[PromptUpdateParamsPromptDataPromptChatMessagesSystemRole] `json:"role,required"`
	Content param.Field[string]                                                   `json:"content"`
	Name    param.Field[string]                                                   `json:"name"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesSystem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesSystem) implementsPromptUpdateParamsPromptDataPromptChatMessageUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesSystemRole string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesSystemRoleSystem PromptUpdateParamsPromptDataPromptChatMessagesSystemRole = "system"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesSystemRole) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesSystemRoleSystem:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesUser struct {
	Role    param.Field[PromptUpdateParamsPromptDataPromptChatMessagesUserRole]         `json:"role,required"`
	Content param.Field[PromptUpdateParamsPromptDataPromptChatMessagesUserContentUnion] `json:"content"`
	Name    param.Field[string]                                                         `json:"name"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesUser) implementsPromptUpdateParamsPromptDataPromptChatMessageUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesUserRole string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesUserRoleUser PromptUpdateParamsPromptDataPromptChatMessagesUserRole = "user"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesUserRole) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesUserRoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [PromptUpdateParamsPromptDataPromptChatMessagesUserContentArray].
type PromptUpdateParamsPromptDataPromptChatMessagesUserContentUnion interface {
	ImplementsPromptUpdateParamsPromptDataPromptChatMessagesUserContentUnion()
}

type PromptUpdateParamsPromptDataPromptChatMessagesUserContentArray []PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayUnion

func (r PromptUpdateParamsPromptDataPromptChatMessagesUserContentArray) ImplementsPromptUpdateParamsPromptDataPromptChatMessagesUserContentUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesUserContentArray struct {
	Text     param.Field[string]                                                             `json:"text"`
	Type     param.Field[PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                        `json:"image_url,required"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesUserContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesUserContentArray) implementsPromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

// Satisfied by
// [PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayText],
// [PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURL],
// [PromptUpdateParamsPromptDataPromptChatMessagesUserContentArray].
type PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayUnion interface {
	implementsPromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayUnion()
}

type PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayText struct {
	Type param.Field[PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayTextType] `json:"type,required"`
	Text param.Field[string]                                                                 `json:"text"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayText) implementsPromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayTextType string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayTextTypeText PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayTextType = "text"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayTextType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayTextTypeText:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURL struct {
	ImageURL param.Field[PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL] `json:"image_url,required"`
	Type     param.Field[PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLType]     `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURL) implementsPromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL struct {
	URL    param.Field[string]                                                                               `json:"url,required"`
	Detail param.Field[PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail] `json:"detail"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "auto"
	PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow  PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "low"
	PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "high"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto, PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow, PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLType string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLType = "image_url"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayType string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayTypeText     PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayType = "text"
	PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayTypeImageURL PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayType = "image_url"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayTypeText, PromptUpdateParamsPromptDataPromptChatMessagesUserContentArrayTypeImageURL:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesAssistant struct {
	Role         param.Field[PromptUpdateParamsPromptDataPromptChatMessagesAssistantRole]         `json:"role,required"`
	Content      param.Field[string]                                                              `json:"content"`
	FunctionCall param.Field[PromptUpdateParamsPromptDataPromptChatMessagesAssistantFunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                              `json:"name"`
	ToolCalls    param.Field[[]PromptUpdateParamsPromptDataPromptChatMessagesAssistantToolCall]   `json:"tool_calls"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesAssistant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesAssistant) implementsPromptUpdateParamsPromptDataPromptChatMessageUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesAssistantRole string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesAssistantRoleAssistant PromptUpdateParamsPromptDataPromptChatMessagesAssistantRole = "assistant"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesAssistantRole) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesAssistantRoleAssistant:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesAssistantFunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesAssistantFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataPromptChatMessagesAssistantToolCall struct {
	ID       param.Field[string]                                                                   `json:"id,required"`
	Function param.Field[PromptUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsFunction] `json:"function,required"`
	Type     param.Field[PromptUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsType]     `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesAssistantToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsType string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsTypeFunction PromptUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsType = "function"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesAssistantToolCallsTypeFunction:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesTool struct {
	Role       param.Field[PromptUpdateParamsPromptDataPromptChatMessagesToolRole] `json:"role,required"`
	Content    param.Field[string]                                                 `json:"content"`
	ToolCallID param.Field[string]                                                 `json:"tool_call_id"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesTool) implementsPromptUpdateParamsPromptDataPromptChatMessageUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesToolRole string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesToolRoleTool PromptUpdateParamsPromptDataPromptChatMessagesToolRole = "tool"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesToolRole) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesToolRoleTool:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesFunction struct {
	Name    param.Field[string]                                                     `json:"name,required"`
	Role    param.Field[PromptUpdateParamsPromptDataPromptChatMessagesFunctionRole] `json:"role,required"`
	Content param.Field[string]                                                     `json:"content"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesFunction) implementsPromptUpdateParamsPromptDataPromptChatMessageUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesFunctionRole string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesFunctionRoleFunction PromptUpdateParamsPromptDataPromptChatMessagesFunctionRole = "function"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesFunctionRole) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesFunctionRoleFunction:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptChatMessagesFallback struct {
	Role    param.Field[PromptUpdateParamsPromptDataPromptChatMessagesFallbackRole] `json:"role,required"`
	Content param.Field[string]                                                     `json:"content"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesFallback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesFallback) implementsPromptUpdateParamsPromptDataPromptChatMessageUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesFallbackRole string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesFallbackRoleModel PromptUpdateParamsPromptDataPromptChatMessagesFallbackRole = "model"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesFallbackRole) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesFallbackRoleModel:
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

// Satisfied by [PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParams],
// [PromptReplaceParamsPromptDataOptionsParamsAnthropicModelParams],
// [PromptReplaceParamsPromptDataOptionsParamsGoogleModelParams],
// [PromptReplaceParamsPromptDataOptionsParamsWindowAIModelParams],
// [PromptReplaceParamsPromptDataOptionsParamsJsCompletionParams].
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
// [PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto],
// [PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone],
// [PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction].
type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsPromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto string

const (
	PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto = "auto"
)

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto:
		return true
	}
	return false
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) implementsPromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone string

const (
	PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone = "none"
)

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone:
		return true
	}
	return false
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) implementsPromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) implementsPromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
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
// [PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto],
// [PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone],
// [PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction].
type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsPromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto string

const (
	PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto = "auto"
)

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto:
		return true
	}
	return false
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) implementsPromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone string

const (
	PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone = "none"
)

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone:
		return true
	}
	return false
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) implementsPromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction struct {
	Function param.Field[PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction] `json:"function,required"`
	Type     param.Field[PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType]     `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) implementsPromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType string

const (
	PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType = "function"
)

func (r PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction:
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

// Satisfied by [PromptReplaceParamsPromptDataPromptChatMessagesSystem],
// [PromptReplaceParamsPromptDataPromptChatMessagesUser],
// [PromptReplaceParamsPromptDataPromptChatMessagesAssistant],
// [PromptReplaceParamsPromptDataPromptChatMessagesTool],
// [PromptReplaceParamsPromptDataPromptChatMessagesFunction],
// [PromptReplaceParamsPromptDataPromptChatMessagesFallback],
// [PromptReplaceParamsPromptDataPromptChatMessage].
type PromptReplaceParamsPromptDataPromptChatMessageUnion interface {
	implementsPromptReplaceParamsPromptDataPromptChatMessageUnion()
}

type PromptReplaceParamsPromptDataPromptChatMessagesSystem struct {
	Role    param.Field[PromptReplaceParamsPromptDataPromptChatMessagesSystemRole] `json:"role,required"`
	Content param.Field[string]                                                    `json:"content"`
	Name    param.Field[string]                                                    `json:"name"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesSystem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesSystem) implementsPromptReplaceParamsPromptDataPromptChatMessageUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesSystemRole string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesSystemRoleSystem PromptReplaceParamsPromptDataPromptChatMessagesSystemRole = "system"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesSystemRole) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesSystemRoleSystem:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesUser struct {
	Role    param.Field[PromptReplaceParamsPromptDataPromptChatMessagesUserRole]         `json:"role,required"`
	Content param.Field[PromptReplaceParamsPromptDataPromptChatMessagesUserContentUnion] `json:"content"`
	Name    param.Field[string]                                                          `json:"name"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesUser) implementsPromptReplaceParamsPromptDataPromptChatMessageUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesUserRole string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesUserRoleUser PromptReplaceParamsPromptDataPromptChatMessagesUserRole = "user"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesUserRole) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesUserRoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [PromptReplaceParamsPromptDataPromptChatMessagesUserContentArray].
type PromptReplaceParamsPromptDataPromptChatMessagesUserContentUnion interface {
	ImplementsPromptReplaceParamsPromptDataPromptChatMessagesUserContentUnion()
}

type PromptReplaceParamsPromptDataPromptChatMessagesUserContentArray []PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayUnion

func (r PromptReplaceParamsPromptDataPromptChatMessagesUserContentArray) ImplementsPromptReplaceParamsPromptDataPromptChatMessagesUserContentUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesUserContentArray struct {
	Text     param.Field[string]                                                              `json:"text"`
	Type     param.Field[PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                         `json:"image_url,required"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesUserContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesUserContentArray) implementsPromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

// Satisfied by
// [PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayText],
// [PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURL],
// [PromptReplaceParamsPromptDataPromptChatMessagesUserContentArray].
type PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayUnion interface {
	implementsPromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayUnion()
}

type PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayText struct {
	Type param.Field[PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayTextType] `json:"type,required"`
	Text param.Field[string]                                                                  `json:"text"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayText) implementsPromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayTextType string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayTextTypeText PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayTextType = "text"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayTextType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayTextTypeText:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURL struct {
	ImageURL param.Field[PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL] `json:"image_url,required"`
	Type     param.Field[PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLType]     `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURL) implementsPromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL struct {
	URL    param.Field[string]                                                                                `json:"url,required"`
	Detail param.Field[PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail] `json:"detail"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "auto"
	PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow  PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "low"
	PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "high"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto, PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow, PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLType string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLType = "image_url"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayType string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayTypeText     PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayType = "text"
	PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayTypeImageURL PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayType = "image_url"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayTypeText, PromptReplaceParamsPromptDataPromptChatMessagesUserContentArrayTypeImageURL:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesAssistant struct {
	Role         param.Field[PromptReplaceParamsPromptDataPromptChatMessagesAssistantRole]         `json:"role,required"`
	Content      param.Field[string]                                                               `json:"content"`
	FunctionCall param.Field[PromptReplaceParamsPromptDataPromptChatMessagesAssistantFunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                               `json:"name"`
	ToolCalls    param.Field[[]PromptReplaceParamsPromptDataPromptChatMessagesAssistantToolCall]   `json:"tool_calls"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesAssistant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesAssistant) implementsPromptReplaceParamsPromptDataPromptChatMessageUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesAssistantRole string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesAssistantRoleAssistant PromptReplaceParamsPromptDataPromptChatMessagesAssistantRole = "assistant"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesAssistantRole) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesAssistantRoleAssistant:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesAssistantFunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesAssistantFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataPromptChatMessagesAssistantToolCall struct {
	ID       param.Field[string]                                                                    `json:"id,required"`
	Function param.Field[PromptReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsFunction] `json:"function,required"`
	Type     param.Field[PromptReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsType]     `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesAssistantToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsType string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsTypeFunction PromptReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsType = "function"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesAssistantToolCallsTypeFunction:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesTool struct {
	Role       param.Field[PromptReplaceParamsPromptDataPromptChatMessagesToolRole] `json:"role,required"`
	Content    param.Field[string]                                                  `json:"content"`
	ToolCallID param.Field[string]                                                  `json:"tool_call_id"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesTool) implementsPromptReplaceParamsPromptDataPromptChatMessageUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesToolRole string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesToolRoleTool PromptReplaceParamsPromptDataPromptChatMessagesToolRole = "tool"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesToolRole) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesToolRoleTool:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesFunction struct {
	Name    param.Field[string]                                                      `json:"name,required"`
	Role    param.Field[PromptReplaceParamsPromptDataPromptChatMessagesFunctionRole] `json:"role,required"`
	Content param.Field[string]                                                      `json:"content"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesFunction) implementsPromptReplaceParamsPromptDataPromptChatMessageUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesFunctionRole string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesFunctionRoleFunction PromptReplaceParamsPromptDataPromptChatMessagesFunctionRole = "function"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesFunctionRole) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesFunctionRoleFunction:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptChatMessagesFallback struct {
	Role    param.Field[PromptReplaceParamsPromptDataPromptChatMessagesFallbackRole] `json:"role,required"`
	Content param.Field[string]                                                      `json:"content"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesFallback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesFallback) implementsPromptReplaceParamsPromptDataPromptChatMessageUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesFallbackRole string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesFallbackRoleModel PromptReplaceParamsPromptDataPromptChatMessagesFallbackRole = "model"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesFallbackRole) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesFallbackRoleModel:
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
