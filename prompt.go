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
	PromptData PromptData `json:"prompt_data,nullable"`
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
type PromptData struct {
	Options PromptDataOptions `json:"options,nullable"`
	Origin  PromptDataOrigin  `json:"origin,nullable"`
	Prompt  PromptDataPrompt  `json:"prompt"`
	JSON    promptDataJSON    `json:"-"`
}

// promptDataJSON contains the JSON metadata for the struct [PromptData]
type promptDataJSON struct {
	Options     apijson.Field
	Origin      apijson.Field
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataJSON) RawJSON() string {
	return r.raw
}

type PromptDataOptions struct {
	Model    string                       `json:"model"`
	Params   PromptDataOptionsParamsUnion `json:"params"`
	Position string                       `json:"position"`
	JSON     promptDataOptionsJSON        `json:"-"`
}

// promptDataOptionsJSON contains the JSON metadata for the struct
// [PromptDataOptions]
type promptDataOptionsJSON struct {
	Model       apijson.Field
	Params      apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataOptionsJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [PromptDataOptionsParamsOpenAIModelParams],
// [PromptDataOptionsParamsAnthropicModelParams],
// [PromptDataOptionsParamsGoogleModelParams],
// [PromptDataOptionsParamsWindowAIModelParams] or
// [PromptDataOptionsParamsJsCompletionParams].
type PromptDataOptionsParamsUnion interface {
	implementsPromptDataOptionsParamsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptDataOptionsParamsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataOptionsParamsOpenAIModelParams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataOptionsParamsAnthropicModelParams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataOptionsParamsGoogleModelParams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataOptionsParamsWindowAIModelParams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataOptionsParamsJsCompletionParams{}),
		},
	)
}

type PromptDataOptionsParamsOpenAIModelParams struct {
	FrequencyPenalty float64                                                   `json:"frequency_penalty"`
	FunctionCall     PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion `json:"function_call"`
	MaxTokens        float64                                                   `json:"max_tokens"`
	N                float64                                                   `json:"n"`
	PresencePenalty  float64                                                   `json:"presence_penalty"`
	ResponseFormat   PromptDataOptionsParamsOpenAIModelParamsResponseFormat    `json:"response_format,nullable"`
	Stop             []string                                                  `json:"stop"`
	Temperature      float64                                                   `json:"temperature"`
	ToolChoice       PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion   `json:"tool_choice"`
	TopP             float64                                                   `json:"top_p"`
	UseCache         bool                                                      `json:"use_cache"`
	JSON             promptDataOptionsParamsOpenAIModelParamsJSON              `json:"-"`
}

// promptDataOptionsParamsOpenAIModelParamsJSON contains the JSON metadata for the
// struct [PromptDataOptionsParamsOpenAIModelParams]
type promptDataOptionsParamsOpenAIModelParamsJSON struct {
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

func (r *PromptDataOptionsParamsOpenAIModelParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataOptionsParamsOpenAIModelParamsJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataOptionsParamsOpenAIModelParams) implementsPromptDataOptionsParamsUnion() {}

// Union satisfied by [PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto],
// [PromptDataOptionsParamsOpenAIModelParamsFunctionCallNone] or
// [PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction].
type PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptDataOptionsParamsOpenAIModelParamsFunctionCallNone("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction{}),
		},
	)
}

type PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto string

const (
	PromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto = "auto"
)

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) IsKnown() bool {
	switch r {
	case PromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto:
		return true
	}
	return false
}

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) implementsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) implementsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam() {
}

type PromptDataOptionsParamsOpenAIModelParamsFunctionCallNone string

const (
	PromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone PromptDataOptionsParamsOpenAIModelParamsFunctionCallNone = "none"
)

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) IsKnown() bool {
	switch r {
	case PromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone:
		return true
	}
	return false
}

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) implementsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) implementsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam() {
}

type PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction struct {
	Name string                                                           `json:"name,required"`
	JSON promptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionJSON `json:"-"`
}

// promptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionJSON contains the
// JSON metadata for the struct
// [PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction]
type promptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) implementsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

type PromptDataOptionsParamsOpenAIModelParamsResponseFormat struct {
	Type PromptDataOptionsParamsOpenAIModelParamsResponseFormatType `json:"type,required"`
	JSON promptDataOptionsParamsOpenAIModelParamsResponseFormatJSON `json:"-"`
}

// promptDataOptionsParamsOpenAIModelParamsResponseFormatJSON contains the JSON
// metadata for the struct [PromptDataOptionsParamsOpenAIModelParamsResponseFormat]
type promptDataOptionsParamsOpenAIModelParamsResponseFormatJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataOptionsParamsOpenAIModelParamsResponseFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataOptionsParamsOpenAIModelParamsResponseFormatJSON) RawJSON() string {
	return r.raw
}

type PromptDataOptionsParamsOpenAIModelParamsResponseFormatType string

const (
	PromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject PromptDataOptionsParamsOpenAIModelParamsResponseFormatType = "json_object"
)

func (r PromptDataOptionsParamsOpenAIModelParamsResponseFormatType) IsKnown() bool {
	switch r {
	case PromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Union satisfied by [PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto],
// [PromptDataOptionsParamsOpenAIModelParamsToolChoiceNone] or
// [PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction].
type PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptDataOptionsParamsOpenAIModelParamsToolChoiceNone("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction{}),
		},
	)
}

type PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto string

const (
	PromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto = "auto"
)

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) IsKnown() bool {
	switch r {
	case PromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto:
		return true
	}
	return false
}

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) implementsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) implementsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam() {
}

type PromptDataOptionsParamsOpenAIModelParamsToolChoiceNone string

const (
	PromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone PromptDataOptionsParamsOpenAIModelParamsToolChoiceNone = "none"
)

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) IsKnown() bool {
	switch r {
	case PromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone:
		return true
	}
	return false
}

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) implementsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) implementsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam() {
}

type PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction struct {
	Function PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction `json:"function,required"`
	Type     PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType     `json:"type,required"`
	JSON     promptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionJSON     `json:"-"`
}

// promptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionJSON contains the JSON
// metadata for the struct
// [PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction]
type promptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionJSON struct {
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) implementsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

type PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction struct {
	Name string                                                                 `json:"name,required"`
	JSON promptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionJSON `json:"-"`
}

// promptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionJSON contains
// the JSON metadata for the struct
// [PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction]
type promptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionJSON) RawJSON() string {
	return r.raw
}

type PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType string

const (
	PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType = "function"
)

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType) IsKnown() bool {
	switch r {
	case PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction:
		return true
	}
	return false
}

type PromptDataOptionsParamsAnthropicModelParams struct {
	MaxTokens   float64 `json:"max_tokens,required"`
	Temperature float64 `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample float64                                         `json:"max_tokens_to_sample"`
	StopSequences     []string                                        `json:"stop_sequences"`
	TopK              float64                                         `json:"top_k"`
	TopP              float64                                         `json:"top_p"`
	UseCache          bool                                            `json:"use_cache"`
	JSON              promptDataOptionsParamsAnthropicModelParamsJSON `json:"-"`
}

// promptDataOptionsParamsAnthropicModelParamsJSON contains the JSON metadata for
// the struct [PromptDataOptionsParamsAnthropicModelParams]
type promptDataOptionsParamsAnthropicModelParamsJSON struct {
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

func (r *PromptDataOptionsParamsAnthropicModelParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataOptionsParamsAnthropicModelParamsJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataOptionsParamsAnthropicModelParams) implementsPromptDataOptionsParamsUnion() {}

type PromptDataOptionsParamsGoogleModelParams struct {
	MaxOutputTokens float64                                      `json:"maxOutputTokens"`
	Temperature     float64                                      `json:"temperature"`
	TopK            float64                                      `json:"topK"`
	TopP            float64                                      `json:"topP"`
	UseCache        bool                                         `json:"use_cache"`
	JSON            promptDataOptionsParamsGoogleModelParamsJSON `json:"-"`
}

// promptDataOptionsParamsGoogleModelParamsJSON contains the JSON metadata for the
// struct [PromptDataOptionsParamsGoogleModelParams]
type promptDataOptionsParamsGoogleModelParamsJSON struct {
	MaxOutputTokens apijson.Field
	Temperature     apijson.Field
	TopK            apijson.Field
	TopP            apijson.Field
	UseCache        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PromptDataOptionsParamsGoogleModelParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataOptionsParamsGoogleModelParamsJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataOptionsParamsGoogleModelParams) implementsPromptDataOptionsParamsUnion() {}

type PromptDataOptionsParamsWindowAIModelParams struct {
	Temperature float64                                        `json:"temperature"`
	TopK        float64                                        `json:"topK"`
	UseCache    bool                                           `json:"use_cache"`
	JSON        promptDataOptionsParamsWindowAIModelParamsJSON `json:"-"`
}

// promptDataOptionsParamsWindowAIModelParamsJSON contains the JSON metadata for
// the struct [PromptDataOptionsParamsWindowAIModelParams]
type promptDataOptionsParamsWindowAIModelParamsJSON struct {
	Temperature apijson.Field
	TopK        apijson.Field
	UseCache    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataOptionsParamsWindowAIModelParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataOptionsParamsWindowAIModelParamsJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataOptionsParamsWindowAIModelParams) implementsPromptDataOptionsParamsUnion() {}

type PromptDataOptionsParamsJsCompletionParams struct {
	UseCache bool                                          `json:"use_cache"`
	JSON     promptDataOptionsParamsJsCompletionParamsJSON `json:"-"`
}

// promptDataOptionsParamsJsCompletionParamsJSON contains the JSON metadata for the
// struct [PromptDataOptionsParamsJsCompletionParams]
type promptDataOptionsParamsJsCompletionParamsJSON struct {
	UseCache    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataOptionsParamsJsCompletionParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataOptionsParamsJsCompletionParamsJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataOptionsParamsJsCompletionParams) implementsPromptDataOptionsParamsUnion() {}

type PromptDataOrigin struct {
	ProjectID     string               `json:"project_id"`
	PromptID      string               `json:"prompt_id"`
	PromptVersion string               `json:"prompt_version"`
	JSON          promptDataOriginJSON `json:"-"`
}

// promptDataOriginJSON contains the JSON metadata for the struct
// [PromptDataOrigin]
type promptDataOriginJSON struct {
	ProjectID     apijson.Field
	PromptID      apijson.Field
	PromptVersion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PromptDataOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataOriginJSON) RawJSON() string {
	return r.raw
}

type PromptDataPrompt struct {
	Type    PromptDataPromptType `json:"type"`
	Content string               `json:"content"`
	// This field can have the runtime type of [[]PromptDataPromptChatMessage].
	Messages interface{}          `json:"messages,required"`
	Tools    string               `json:"tools"`
	JSON     promptDataPromptJSON `json:"-"`
	union    PromptDataPromptUnion
}

// promptDataPromptJSON contains the JSON metadata for the struct
// [PromptDataPrompt]
type promptDataPromptJSON struct {
	Type        apijson.Field
	Content     apijson.Field
	Messages    apijson.Field
	Tools       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r promptDataPromptJSON) RawJSON() string {
	return r.raw
}

func (r *PromptDataPrompt) UnmarshalJSON(data []byte) (err error) {
	*r = PromptDataPrompt{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PromptDataPromptUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [PromptDataPromptCompletion],
// [PromptDataPromptChat], [PromptDataPromptNullableVariant].
func (r PromptDataPrompt) AsUnion() PromptDataPromptUnion {
	return r.union
}

// Union satisfied by [PromptDataPromptCompletion], [PromptDataPromptChat] or
// [PromptDataPromptNullableVariant].
type PromptDataPromptUnion interface {
	implementsPromptDataPrompt()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptDataPromptUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataPromptCompletion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataPromptChat{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataPromptNullableVariant{}),
		},
	)
}

type PromptDataPromptCompletion struct {
	Content string                         `json:"content,required"`
	Type    PromptDataPromptCompletionType `json:"type,required"`
	JSON    promptDataPromptCompletionJSON `json:"-"`
}

// promptDataPromptCompletionJSON contains the JSON metadata for the struct
// [PromptDataPromptCompletion]
type promptDataPromptCompletionJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataPromptCompletion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataPromptCompletionJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataPromptCompletion) implementsPromptDataPrompt() {}

type PromptDataPromptCompletionType string

const (
	PromptDataPromptCompletionTypeCompletion PromptDataPromptCompletionType = "completion"
)

func (r PromptDataPromptCompletionType) IsKnown() bool {
	switch r {
	case PromptDataPromptCompletionTypeCompletion:
		return true
	}
	return false
}

type PromptDataPromptChat struct {
	Messages []PromptDataPromptChatMessage `json:"messages,required"`
	Type     PromptDataPromptChatType      `json:"type,required"`
	Tools    string                        `json:"tools"`
	JSON     promptDataPromptChatJSON      `json:"-"`
}

// promptDataPromptChatJSON contains the JSON metadata for the struct
// [PromptDataPromptChat]
type promptDataPromptChatJSON struct {
	Messages    apijson.Field
	Type        apijson.Field
	Tools       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataPromptChat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataPromptChatJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataPromptChat) implementsPromptDataPrompt() {}

type PromptDataPromptChatMessage struct {
	// This field can have the runtime type of [string],
	// [PromptDataPromptChatMessagesUserContentUnion].
	Content interface{}                      `json:"content,required"`
	Role    PromptDataPromptChatMessagesRole `json:"role,required"`
	Name    string                           `json:"name"`
	// This field can have the runtime type of
	// [PromptDataPromptChatMessagesAssistantFunctionCall].
	FunctionCall interface{} `json:"function_call,required"`
	// This field can have the runtime type of
	// [[]PromptDataPromptChatMessagesAssistantToolCall].
	ToolCalls  interface{}                     `json:"tool_calls,required"`
	ToolCallID string                          `json:"tool_call_id"`
	JSON       promptDataPromptChatMessageJSON `json:"-"`
	union      PromptDataPromptChatMessagesUnion
}

// promptDataPromptChatMessageJSON contains the JSON metadata for the struct
// [PromptDataPromptChatMessage]
type promptDataPromptChatMessageJSON struct {
	Content      apijson.Field
	Role         apijson.Field
	Name         apijson.Field
	FunctionCall apijson.Field
	ToolCalls    apijson.Field
	ToolCallID   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r promptDataPromptChatMessageJSON) RawJSON() string {
	return r.raw
}

func (r *PromptDataPromptChatMessage) UnmarshalJSON(data []byte) (err error) {
	*r = PromptDataPromptChatMessage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PromptDataPromptChatMessagesUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [PromptDataPromptChatMessagesSystem],
// [PromptDataPromptChatMessagesUser], [PromptDataPromptChatMessagesAssistant],
// [PromptDataPromptChatMessagesTool], [PromptDataPromptChatMessagesFunction],
// [PromptDataPromptChatMessagesFallback].
func (r PromptDataPromptChatMessage) AsUnion() PromptDataPromptChatMessagesUnion {
	return r.union
}

// Union satisfied by [PromptDataPromptChatMessagesSystem],
// [PromptDataPromptChatMessagesUser], [PromptDataPromptChatMessagesAssistant],
// [PromptDataPromptChatMessagesTool], [PromptDataPromptChatMessagesFunction] or
// [PromptDataPromptChatMessagesFallback].
type PromptDataPromptChatMessagesUnion interface {
	implementsPromptDataPromptChatMessage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptDataPromptChatMessagesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataPromptChatMessagesSystem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataPromptChatMessagesUser{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataPromptChatMessagesAssistant{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataPromptChatMessagesTool{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataPromptChatMessagesFunction{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataPromptChatMessagesFallback{}),
		},
	)
}

type PromptDataPromptChatMessagesSystem struct {
	Role    PromptDataPromptChatMessagesSystemRole `json:"role,required"`
	Content string                                 `json:"content"`
	Name    string                                 `json:"name"`
	JSON    promptDataPromptChatMessagesSystemJSON `json:"-"`
}

// promptDataPromptChatMessagesSystemJSON contains the JSON metadata for the struct
// [PromptDataPromptChatMessagesSystem]
type promptDataPromptChatMessagesSystemJSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataPromptChatMessagesSystem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataPromptChatMessagesSystemJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataPromptChatMessagesSystem) implementsPromptDataPromptChatMessage() {}

type PromptDataPromptChatMessagesSystemRole string

const (
	PromptDataPromptChatMessagesSystemRoleSystem PromptDataPromptChatMessagesSystemRole = "system"
)

func (r PromptDataPromptChatMessagesSystemRole) IsKnown() bool {
	switch r {
	case PromptDataPromptChatMessagesSystemRoleSystem:
		return true
	}
	return false
}

type PromptDataPromptChatMessagesUser struct {
	Role    PromptDataPromptChatMessagesUserRole         `json:"role,required"`
	Content PromptDataPromptChatMessagesUserContentUnion `json:"content"`
	Name    string                                       `json:"name"`
	JSON    promptDataPromptChatMessagesUserJSON         `json:"-"`
}

// promptDataPromptChatMessagesUserJSON contains the JSON metadata for the struct
// [PromptDataPromptChatMessagesUser]
type promptDataPromptChatMessagesUserJSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataPromptChatMessagesUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataPromptChatMessagesUserJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataPromptChatMessagesUser) implementsPromptDataPromptChatMessage() {}

type PromptDataPromptChatMessagesUserRole string

const (
	PromptDataPromptChatMessagesUserRoleUser PromptDataPromptChatMessagesUserRole = "user"
)

func (r PromptDataPromptChatMessagesUserRole) IsKnown() bool {
	switch r {
	case PromptDataPromptChatMessagesUserRoleUser:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or
// [PromptDataPromptChatMessagesUserContentArray].
type PromptDataPromptChatMessagesUserContentUnion interface {
	ImplementsPromptDataPromptChatMessagesUserContentUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptDataPromptChatMessagesUserContentUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataPromptChatMessagesUserContentArray{}),
		},
	)
}

type PromptDataPromptChatMessagesUserContentArray []PromptDataPromptChatMessagesUserContentArrayItem

func (r PromptDataPromptChatMessagesUserContentArray) ImplementsPromptDataPromptChatMessagesUserContentUnion() {
}

type PromptDataPromptChatMessagesUserContentArrayItem struct {
	Text string                                           `json:"text"`
	Type PromptDataPromptChatMessagesUserContentArrayType `json:"type,required"`
	// This field can have the runtime type of
	// [PromptDataPromptChatMessagesUserContentArrayImageURLImageURL].
	ImageURL interface{}                                          `json:"image_url,required"`
	JSON     promptDataPromptChatMessagesUserContentArrayItemJSON `json:"-"`
	union    PromptDataPromptChatMessagesUserContentArrayUnionItem
}

// promptDataPromptChatMessagesUserContentArrayItemJSON contains the JSON metadata
// for the struct [PromptDataPromptChatMessagesUserContentArrayItem]
type promptDataPromptChatMessagesUserContentArrayItemJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	ImageURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r promptDataPromptChatMessagesUserContentArrayItemJSON) RawJSON() string {
	return r.raw
}

func (r *PromptDataPromptChatMessagesUserContentArrayItem) UnmarshalJSON(data []byte) (err error) {
	*r = PromptDataPromptChatMessagesUserContentArrayItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PromptDataPromptChatMessagesUserContentArrayUnionItem]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PromptDataPromptChatMessagesUserContentArrayText],
// [PromptDataPromptChatMessagesUserContentArrayImageURL].
func (r PromptDataPromptChatMessagesUserContentArrayItem) AsUnion() PromptDataPromptChatMessagesUserContentArrayUnionItem {
	return r.union
}

// Union satisfied by [PromptDataPromptChatMessagesUserContentArrayText] or
// [PromptDataPromptChatMessagesUserContentArrayImageURL].
type PromptDataPromptChatMessagesUserContentArrayUnionItem interface {
	implementsPromptDataPromptChatMessagesUserContentArrayItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptDataPromptChatMessagesUserContentArrayUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataPromptChatMessagesUserContentArrayText{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataPromptChatMessagesUserContentArrayImageURL{}),
		},
	)
}

type PromptDataPromptChatMessagesUserContentArrayText struct {
	Type PromptDataPromptChatMessagesUserContentArrayTextType `json:"type,required"`
	Text string                                               `json:"text"`
	JSON promptDataPromptChatMessagesUserContentArrayTextJSON `json:"-"`
}

// promptDataPromptChatMessagesUserContentArrayTextJSON contains the JSON metadata
// for the struct [PromptDataPromptChatMessagesUserContentArrayText]
type promptDataPromptChatMessagesUserContentArrayTextJSON struct {
	Type        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataPromptChatMessagesUserContentArrayText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataPromptChatMessagesUserContentArrayTextJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataPromptChatMessagesUserContentArrayText) implementsPromptDataPromptChatMessagesUserContentArrayItem() {
}

type PromptDataPromptChatMessagesUserContentArrayTextType string

const (
	PromptDataPromptChatMessagesUserContentArrayTextTypeText PromptDataPromptChatMessagesUserContentArrayTextType = "text"
)

func (r PromptDataPromptChatMessagesUserContentArrayTextType) IsKnown() bool {
	switch r {
	case PromptDataPromptChatMessagesUserContentArrayTextTypeText:
		return true
	}
	return false
}

type PromptDataPromptChatMessagesUserContentArrayImageURL struct {
	ImageURL PromptDataPromptChatMessagesUserContentArrayImageURLImageURL `json:"image_url,required"`
	Type     PromptDataPromptChatMessagesUserContentArrayImageURLType     `json:"type,required"`
	JSON     promptDataPromptChatMessagesUserContentArrayImageURLJSON     `json:"-"`
}

// promptDataPromptChatMessagesUserContentArrayImageURLJSON contains the JSON
// metadata for the struct [PromptDataPromptChatMessagesUserContentArrayImageURL]
type promptDataPromptChatMessagesUserContentArrayImageURLJSON struct {
	ImageURL    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataPromptChatMessagesUserContentArrayImageURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataPromptChatMessagesUserContentArrayImageURLJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataPromptChatMessagesUserContentArrayImageURL) implementsPromptDataPromptChatMessagesUserContentArrayItem() {
}

type PromptDataPromptChatMessagesUserContentArrayImageURLImageURL struct {
	URL    string                                                             `json:"url,required"`
	Detail PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail `json:"detail"`
	JSON   promptDataPromptChatMessagesUserContentArrayImageURLImageURLJSON   `json:"-"`
}

// promptDataPromptChatMessagesUserContentArrayImageURLImageURLJSON contains the
// JSON metadata for the struct
// [PromptDataPromptChatMessagesUserContentArrayImageURLImageURL]
type promptDataPromptChatMessagesUserContentArrayImageURLImageURLJSON struct {
	URL         apijson.Field
	Detail      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataPromptChatMessagesUserContentArrayImageURLImageURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataPromptChatMessagesUserContentArrayImageURLImageURLJSON) RawJSON() string {
	return r.raw
}

type PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail string

const (
	PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "auto"
	PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow  PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "low"
	PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = "high"
)

func (r PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail) IsKnown() bool {
	switch r {
	case PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto, PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow, PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh:
		return true
	}
	return false
}

type PromptDataPromptChatMessagesUserContentArrayImageURLType string

const (
	PromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL PromptDataPromptChatMessagesUserContentArrayImageURLType = "image_url"
)

func (r PromptDataPromptChatMessagesUserContentArrayImageURLType) IsKnown() bool {
	switch r {
	case PromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL:
		return true
	}
	return false
}

type PromptDataPromptChatMessagesUserContentArrayType string

const (
	PromptDataPromptChatMessagesUserContentArrayTypeText     PromptDataPromptChatMessagesUserContentArrayType = "text"
	PromptDataPromptChatMessagesUserContentArrayTypeImageURL PromptDataPromptChatMessagesUserContentArrayType = "image_url"
)

func (r PromptDataPromptChatMessagesUserContentArrayType) IsKnown() bool {
	switch r {
	case PromptDataPromptChatMessagesUserContentArrayTypeText, PromptDataPromptChatMessagesUserContentArrayTypeImageURL:
		return true
	}
	return false
}

type PromptDataPromptChatMessagesAssistant struct {
	Role         PromptDataPromptChatMessagesAssistantRole         `json:"role,required"`
	Content      string                                            `json:"content,nullable"`
	FunctionCall PromptDataPromptChatMessagesAssistantFunctionCall `json:"function_call"`
	Name         string                                            `json:"name"`
	ToolCalls    []PromptDataPromptChatMessagesAssistantToolCall   `json:"tool_calls"`
	JSON         promptDataPromptChatMessagesAssistantJSON         `json:"-"`
}

// promptDataPromptChatMessagesAssistantJSON contains the JSON metadata for the
// struct [PromptDataPromptChatMessagesAssistant]
type promptDataPromptChatMessagesAssistantJSON struct {
	Role         apijson.Field
	Content      apijson.Field
	FunctionCall apijson.Field
	Name         apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PromptDataPromptChatMessagesAssistant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataPromptChatMessagesAssistantJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataPromptChatMessagesAssistant) implementsPromptDataPromptChatMessage() {}

type PromptDataPromptChatMessagesAssistantRole string

const (
	PromptDataPromptChatMessagesAssistantRoleAssistant PromptDataPromptChatMessagesAssistantRole = "assistant"
)

func (r PromptDataPromptChatMessagesAssistantRole) IsKnown() bool {
	switch r {
	case PromptDataPromptChatMessagesAssistantRoleAssistant:
		return true
	}
	return false
}

type PromptDataPromptChatMessagesAssistantFunctionCall struct {
	Arguments string                                                `json:"arguments,required"`
	Name      string                                                `json:"name,required"`
	JSON      promptDataPromptChatMessagesAssistantFunctionCallJSON `json:"-"`
}

// promptDataPromptChatMessagesAssistantFunctionCallJSON contains the JSON metadata
// for the struct [PromptDataPromptChatMessagesAssistantFunctionCall]
type promptDataPromptChatMessagesAssistantFunctionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataPromptChatMessagesAssistantFunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataPromptChatMessagesAssistantFunctionCallJSON) RawJSON() string {
	return r.raw
}

type PromptDataPromptChatMessagesAssistantToolCall struct {
	ID       string                                                 `json:"id,required"`
	Function PromptDataPromptChatMessagesAssistantToolCallsFunction `json:"function,required"`
	Type     PromptDataPromptChatMessagesAssistantToolCallsType     `json:"type,required"`
	JSON     promptDataPromptChatMessagesAssistantToolCallJSON      `json:"-"`
}

// promptDataPromptChatMessagesAssistantToolCallJSON contains the JSON metadata for
// the struct [PromptDataPromptChatMessagesAssistantToolCall]
type promptDataPromptChatMessagesAssistantToolCallJSON struct {
	ID          apijson.Field
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataPromptChatMessagesAssistantToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataPromptChatMessagesAssistantToolCallJSON) RawJSON() string {
	return r.raw
}

type PromptDataPromptChatMessagesAssistantToolCallsFunction struct {
	Arguments string                                                     `json:"arguments,required"`
	Name      string                                                     `json:"name,required"`
	JSON      promptDataPromptChatMessagesAssistantToolCallsFunctionJSON `json:"-"`
}

// promptDataPromptChatMessagesAssistantToolCallsFunctionJSON contains the JSON
// metadata for the struct [PromptDataPromptChatMessagesAssistantToolCallsFunction]
type promptDataPromptChatMessagesAssistantToolCallsFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataPromptChatMessagesAssistantToolCallsFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataPromptChatMessagesAssistantToolCallsFunctionJSON) RawJSON() string {
	return r.raw
}

type PromptDataPromptChatMessagesAssistantToolCallsType string

const (
	PromptDataPromptChatMessagesAssistantToolCallsTypeFunction PromptDataPromptChatMessagesAssistantToolCallsType = "function"
)

func (r PromptDataPromptChatMessagesAssistantToolCallsType) IsKnown() bool {
	switch r {
	case PromptDataPromptChatMessagesAssistantToolCallsTypeFunction:
		return true
	}
	return false
}

type PromptDataPromptChatMessagesTool struct {
	Role       PromptDataPromptChatMessagesToolRole `json:"role,required"`
	Content    string                               `json:"content"`
	ToolCallID string                               `json:"tool_call_id"`
	JSON       promptDataPromptChatMessagesToolJSON `json:"-"`
}

// promptDataPromptChatMessagesToolJSON contains the JSON metadata for the struct
// [PromptDataPromptChatMessagesTool]
type promptDataPromptChatMessagesToolJSON struct {
	Role        apijson.Field
	Content     apijson.Field
	ToolCallID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataPromptChatMessagesTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataPromptChatMessagesToolJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataPromptChatMessagesTool) implementsPromptDataPromptChatMessage() {}

type PromptDataPromptChatMessagesToolRole string

const (
	PromptDataPromptChatMessagesToolRoleTool PromptDataPromptChatMessagesToolRole = "tool"
)

func (r PromptDataPromptChatMessagesToolRole) IsKnown() bool {
	switch r {
	case PromptDataPromptChatMessagesToolRoleTool:
		return true
	}
	return false
}

type PromptDataPromptChatMessagesFunction struct {
	Name    string                                   `json:"name,required"`
	Role    PromptDataPromptChatMessagesFunctionRole `json:"role,required"`
	Content string                                   `json:"content"`
	JSON    promptDataPromptChatMessagesFunctionJSON `json:"-"`
}

// promptDataPromptChatMessagesFunctionJSON contains the JSON metadata for the
// struct [PromptDataPromptChatMessagesFunction]
type promptDataPromptChatMessagesFunctionJSON struct {
	Name        apijson.Field
	Role        apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataPromptChatMessagesFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataPromptChatMessagesFunctionJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataPromptChatMessagesFunction) implementsPromptDataPromptChatMessage() {}

type PromptDataPromptChatMessagesFunctionRole string

const (
	PromptDataPromptChatMessagesFunctionRoleFunction PromptDataPromptChatMessagesFunctionRole = "function"
)

func (r PromptDataPromptChatMessagesFunctionRole) IsKnown() bool {
	switch r {
	case PromptDataPromptChatMessagesFunctionRoleFunction:
		return true
	}
	return false
}

type PromptDataPromptChatMessagesFallback struct {
	Role    PromptDataPromptChatMessagesFallbackRole `json:"role,required"`
	Content string                                   `json:"content,nullable"`
	JSON    promptDataPromptChatMessagesFallbackJSON `json:"-"`
}

// promptDataPromptChatMessagesFallbackJSON contains the JSON metadata for the
// struct [PromptDataPromptChatMessagesFallback]
type promptDataPromptChatMessagesFallbackJSON struct {
	Role        apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataPromptChatMessagesFallback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataPromptChatMessagesFallbackJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataPromptChatMessagesFallback) implementsPromptDataPromptChatMessage() {}

type PromptDataPromptChatMessagesFallbackRole string

const (
	PromptDataPromptChatMessagesFallbackRoleModel PromptDataPromptChatMessagesFallbackRole = "model"
)

func (r PromptDataPromptChatMessagesFallbackRole) IsKnown() bool {
	switch r {
	case PromptDataPromptChatMessagesFallbackRoleModel:
		return true
	}
	return false
}

type PromptDataPromptChatMessagesRole string

const (
	PromptDataPromptChatMessagesRoleSystem    PromptDataPromptChatMessagesRole = "system"
	PromptDataPromptChatMessagesRoleUser      PromptDataPromptChatMessagesRole = "user"
	PromptDataPromptChatMessagesRoleAssistant PromptDataPromptChatMessagesRole = "assistant"
	PromptDataPromptChatMessagesRoleTool      PromptDataPromptChatMessagesRole = "tool"
	PromptDataPromptChatMessagesRoleFunction  PromptDataPromptChatMessagesRole = "function"
	PromptDataPromptChatMessagesRoleModel     PromptDataPromptChatMessagesRole = "model"
)

func (r PromptDataPromptChatMessagesRole) IsKnown() bool {
	switch r {
	case PromptDataPromptChatMessagesRoleSystem, PromptDataPromptChatMessagesRoleUser, PromptDataPromptChatMessagesRoleAssistant, PromptDataPromptChatMessagesRoleTool, PromptDataPromptChatMessagesRoleFunction, PromptDataPromptChatMessagesRoleModel:
		return true
	}
	return false
}

type PromptDataPromptChatType string

const (
	PromptDataPromptChatTypeChat PromptDataPromptChatType = "chat"
)

func (r PromptDataPromptChatType) IsKnown() bool {
	switch r {
	case PromptDataPromptChatTypeChat:
		return true
	}
	return false
}

type PromptDataPromptNullableVariant struct {
	JSON promptDataPromptNullableVariantJSON `json:"-"`
}

// promptDataPromptNullableVariantJSON contains the JSON metadata for the struct
// [PromptDataPromptNullableVariant]
type promptDataPromptNullableVariantJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataPromptNullableVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataPromptNullableVariantJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataPromptNullableVariant) implementsPromptDataPrompt() {}

type PromptDataPromptType string

const (
	PromptDataPromptTypeCompletion PromptDataPromptType = "completion"
	PromptDataPromptTypeChat       PromptDataPromptType = "chat"
)

func (r PromptDataPromptType) IsKnown() bool {
	switch r {
	case PromptDataPromptTypeCompletion, PromptDataPromptTypeChat:
		return true
	}
	return false
}

// The prompt, model, and its parameters
type PromptDataParam struct {
	Options param.Field[PromptDataOptionsParam]     `json:"options"`
	Origin  param.Field[PromptDataOriginParam]      `json:"origin"`
	Prompt  param.Field[PromptDataPromptUnionParam] `json:"prompt"`
}

func (r PromptDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptDataOptionsParam struct {
	Model    param.Field[string]                            `json:"model"`
	Params   param.Field[PromptDataOptionsParamsUnionParam] `json:"params"`
	Position param.Field[string]                            `json:"position"`
}

func (r PromptDataOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [PromptDataOptionsParamsOpenAIModelParamsParam],
// [PromptDataOptionsParamsAnthropicModelParamsParam],
// [PromptDataOptionsParamsGoogleModelParamsParam],
// [PromptDataOptionsParamsWindowAIModelParamsParam],
// [PromptDataOptionsParamsJsCompletionParamsParam].
type PromptDataOptionsParamsUnionParam interface {
	implementsPromptDataOptionsParamsUnionParam()
}

type PromptDataOptionsParamsOpenAIModelParamsParam struct {
	FrequencyPenalty param.Field[float64]                                                        `json:"frequency_penalty"`
	FunctionCall     param.Field[PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                        `json:"max_tokens"`
	N                param.Field[float64]                                                        `json:"n"`
	PresencePenalty  param.Field[float64]                                                        `json:"presence_penalty"`
	ResponseFormat   param.Field[PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam]    `json:"response_format"`
	Stop             param.Field[[]string]                                                       `json:"stop"`
	Temperature      param.Field[float64]                                                        `json:"temperature"`
	ToolChoice       param.Field[PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                        `json:"top_p"`
	UseCache         param.Field[bool]                                                           `json:"use_cache"`
}

func (r PromptDataOptionsParamsOpenAIModelParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataOptionsParamsOpenAIModelParamsParam) implementsPromptDataOptionsParamsUnionParam() {
}

// Satisfied by [PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto],
// [PromptDataOptionsParamsOpenAIModelParamsFunctionCallNone],
// [PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionParam].
type PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam interface {
	implementsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam()
}

type PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionParam struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionParam) implementsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam() {
}

type PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam struct {
	Type param.Field[PromptDataOptionsParamsOpenAIModelParamsResponseFormatType] `json:"type,required"`
}

func (r PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto],
// [PromptDataOptionsParamsOpenAIModelParamsToolChoiceNone],
// [PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionParam].
type PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam interface {
	implementsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam()
}

type PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionParam struct {
	Function param.Field[PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionParam] `json:"function,required"`
	Type     param.Field[PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType]          `json:"type,required"`
}

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionParam) implementsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam() {
}

type PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionParam struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptDataOptionsParamsAnthropicModelParamsParam struct {
	MaxTokens   param.Field[float64] `json:"max_tokens,required"`
	Temperature param.Field[float64] `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64]  `json:"max_tokens_to_sample"`
	StopSequences     param.Field[[]string] `json:"stop_sequences"`
	TopK              param.Field[float64]  `json:"top_k"`
	TopP              param.Field[float64]  `json:"top_p"`
	UseCache          param.Field[bool]     `json:"use_cache"`
}

func (r PromptDataOptionsParamsAnthropicModelParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataOptionsParamsAnthropicModelParamsParam) implementsPromptDataOptionsParamsUnionParam() {
}

type PromptDataOptionsParamsGoogleModelParamsParam struct {
	MaxOutputTokens param.Field[float64] `json:"maxOutputTokens"`
	Temperature     param.Field[float64] `json:"temperature"`
	TopK            param.Field[float64] `json:"topK"`
	TopP            param.Field[float64] `json:"topP"`
	UseCache        param.Field[bool]    `json:"use_cache"`
}

func (r PromptDataOptionsParamsGoogleModelParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataOptionsParamsGoogleModelParamsParam) implementsPromptDataOptionsParamsUnionParam() {
}

type PromptDataOptionsParamsWindowAIModelParamsParam struct {
	Temperature param.Field[float64] `json:"temperature"`
	TopK        param.Field[float64] `json:"topK"`
	UseCache    param.Field[bool]    `json:"use_cache"`
}

func (r PromptDataOptionsParamsWindowAIModelParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataOptionsParamsWindowAIModelParamsParam) implementsPromptDataOptionsParamsUnionParam() {
}

type PromptDataOptionsParamsJsCompletionParamsParam struct {
	UseCache param.Field[bool] `json:"use_cache"`
}

func (r PromptDataOptionsParamsJsCompletionParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataOptionsParamsJsCompletionParamsParam) implementsPromptDataOptionsParamsUnionParam() {
}

type PromptDataOriginParam struct {
	ProjectID     param.Field[string] `json:"project_id"`
	PromptID      param.Field[string] `json:"prompt_id"`
	PromptVersion param.Field[string] `json:"prompt_version"`
}

func (r PromptDataOriginParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptDataPromptParam struct {
	Type     param.Field[PromptDataPromptType] `json:"type"`
	Content  param.Field[string]               `json:"content"`
	Messages param.Field[interface{}]          `json:"messages,required"`
	Tools    param.Field[string]               `json:"tools"`
}

func (r PromptDataPromptParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptParam) implementsPromptDataPromptUnionParam() {}

// Satisfied by [PromptDataPromptCompletionParam], [PromptDataPromptChatParam],
// [PromptDataPromptNullableVariantParam], [PromptDataPromptParam].
type PromptDataPromptUnionParam interface {
	implementsPromptDataPromptUnionParam()
}

type PromptDataPromptCompletionParam struct {
	Content param.Field[string]                         `json:"content,required"`
	Type    param.Field[PromptDataPromptCompletionType] `json:"type,required"`
}

func (r PromptDataPromptCompletionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptCompletionParam) implementsPromptDataPromptUnionParam() {}

type PromptDataPromptChatParam struct {
	Messages param.Field[[]PromptDataPromptChatMessagesUnionParam] `json:"messages,required"`
	Type     param.Field[PromptDataPromptChatType]                 `json:"type,required"`
	Tools    param.Field[string]                                   `json:"tools"`
}

func (r PromptDataPromptChatParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatParam) implementsPromptDataPromptUnionParam() {}

type PromptDataPromptChatMessageParam struct {
	Content      param.Field[interface{}]                      `json:"content,required"`
	Role         param.Field[PromptDataPromptChatMessagesRole] `json:"role,required"`
	Name         param.Field[string]                           `json:"name"`
	FunctionCall param.Field[interface{}]                      `json:"function_call,required"`
	ToolCalls    param.Field[interface{}]                      `json:"tool_calls,required"`
	ToolCallID   param.Field[string]                           `json:"tool_call_id"`
}

func (r PromptDataPromptChatMessageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessageParam) implementsPromptDataPromptChatMessagesUnionParam() {}

// Satisfied by [PromptDataPromptChatMessagesSystemParam],
// [PromptDataPromptChatMessagesUserParam],
// [PromptDataPromptChatMessagesAssistantParam],
// [PromptDataPromptChatMessagesToolParam],
// [PromptDataPromptChatMessagesFunctionParam],
// [PromptDataPromptChatMessagesFallbackParam], [PromptDataPromptChatMessageParam].
type PromptDataPromptChatMessagesUnionParam interface {
	implementsPromptDataPromptChatMessagesUnionParam()
}

type PromptDataPromptChatMessagesSystemParam struct {
	Role    param.Field[PromptDataPromptChatMessagesSystemRole] `json:"role,required"`
	Content param.Field[string]                                 `json:"content"`
	Name    param.Field[string]                                 `json:"name"`
}

func (r PromptDataPromptChatMessagesSystemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesSystemParam) implementsPromptDataPromptChatMessagesUnionParam() {}

type PromptDataPromptChatMessagesUserParam struct {
	Role    param.Field[PromptDataPromptChatMessagesUserRole]              `json:"role,required"`
	Content param.Field[PromptDataPromptChatMessagesUserContentUnionParam] `json:"content"`
	Name    param.Field[string]                                            `json:"name"`
}

func (r PromptDataPromptChatMessagesUserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesUserParam) implementsPromptDataPromptChatMessagesUnionParam() {}

// Satisfied by [shared.UnionString],
// [PromptDataPromptChatMessagesUserContentArrayParam].
type PromptDataPromptChatMessagesUserContentUnionParam interface {
	ImplementsPromptDataPromptChatMessagesUserContentUnionParam()
}

type PromptDataPromptChatMessagesUserContentArrayParam []PromptDataPromptChatMessagesUserContentArrayUnionItemParam

func (r PromptDataPromptChatMessagesUserContentArrayParam) ImplementsPromptDataPromptChatMessagesUserContentUnionParam() {
}

type PromptDataPromptChatMessagesUserContentArrayItemParam struct {
	Text     param.Field[string]                                           `json:"text"`
	Type     param.Field[PromptDataPromptChatMessagesUserContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                      `json:"image_url,required"`
}

func (r PromptDataPromptChatMessagesUserContentArrayItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesUserContentArrayItemParam) implementsPromptDataPromptChatMessagesUserContentArrayUnionItemParam() {
}

// Satisfied by [PromptDataPromptChatMessagesUserContentArrayTextParam],
// [PromptDataPromptChatMessagesUserContentArrayImageURLParam],
// [PromptDataPromptChatMessagesUserContentArrayItemParam].
type PromptDataPromptChatMessagesUserContentArrayUnionItemParam interface {
	implementsPromptDataPromptChatMessagesUserContentArrayUnionItemParam()
}

type PromptDataPromptChatMessagesUserContentArrayTextParam struct {
	Type param.Field[PromptDataPromptChatMessagesUserContentArrayTextType] `json:"type,required"`
	Text param.Field[string]                                               `json:"text"`
}

func (r PromptDataPromptChatMessagesUserContentArrayTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesUserContentArrayTextParam) implementsPromptDataPromptChatMessagesUserContentArrayUnionItemParam() {
}

type PromptDataPromptChatMessagesUserContentArrayImageURLParam struct {
	ImageURL param.Field[PromptDataPromptChatMessagesUserContentArrayImageURLImageURLParam] `json:"image_url,required"`
	Type     param.Field[PromptDataPromptChatMessagesUserContentArrayImageURLType]          `json:"type,required"`
}

func (r PromptDataPromptChatMessagesUserContentArrayImageURLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesUserContentArrayImageURLParam) implementsPromptDataPromptChatMessagesUserContentArrayUnionItemParam() {
}

type PromptDataPromptChatMessagesUserContentArrayImageURLImageURLParam struct {
	URL    param.Field[string]                                                             `json:"url,required"`
	Detail param.Field[PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail] `json:"detail"`
}

func (r PromptDataPromptChatMessagesUserContentArrayImageURLImageURLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptDataPromptChatMessagesAssistantParam struct {
	Role         param.Field[PromptDataPromptChatMessagesAssistantRole]              `json:"role,required"`
	Content      param.Field[string]                                                 `json:"content"`
	FunctionCall param.Field[PromptDataPromptChatMessagesAssistantFunctionCallParam] `json:"function_call"`
	Name         param.Field[string]                                                 `json:"name"`
	ToolCalls    param.Field[[]PromptDataPromptChatMessagesAssistantToolCallParam]   `json:"tool_calls"`
}

func (r PromptDataPromptChatMessagesAssistantParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesAssistantParam) implementsPromptDataPromptChatMessagesUnionParam() {
}

type PromptDataPromptChatMessagesAssistantFunctionCallParam struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptDataPromptChatMessagesAssistantFunctionCallParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptDataPromptChatMessagesAssistantToolCallParam struct {
	ID       param.Field[string]                                                      `json:"id,required"`
	Function param.Field[PromptDataPromptChatMessagesAssistantToolCallsFunctionParam] `json:"function,required"`
	Type     param.Field[PromptDataPromptChatMessagesAssistantToolCallsType]          `json:"type,required"`
}

func (r PromptDataPromptChatMessagesAssistantToolCallParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptDataPromptChatMessagesAssistantToolCallsFunctionParam struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptDataPromptChatMessagesAssistantToolCallsFunctionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptDataPromptChatMessagesToolParam struct {
	Role       param.Field[PromptDataPromptChatMessagesToolRole] `json:"role,required"`
	Content    param.Field[string]                               `json:"content"`
	ToolCallID param.Field[string]                               `json:"tool_call_id"`
}

func (r PromptDataPromptChatMessagesToolParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesToolParam) implementsPromptDataPromptChatMessagesUnionParam() {}

type PromptDataPromptChatMessagesFunctionParam struct {
	Name    param.Field[string]                                   `json:"name,required"`
	Role    param.Field[PromptDataPromptChatMessagesFunctionRole] `json:"role,required"`
	Content param.Field[string]                                   `json:"content"`
}

func (r PromptDataPromptChatMessagesFunctionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesFunctionParam) implementsPromptDataPromptChatMessagesUnionParam() {
}

type PromptDataPromptChatMessagesFallbackParam struct {
	Role    param.Field[PromptDataPromptChatMessagesFallbackRole] `json:"role,required"`
	Content param.Field[string]                                   `json:"content"`
}

func (r PromptDataPromptChatMessagesFallbackParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesFallbackParam) implementsPromptDataPromptChatMessagesUnionParam() {
}

type PromptDataPromptNullableVariantParam struct {
}

func (r PromptDataPromptNullableVariantParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptNullableVariantParam) implementsPromptDataPromptUnionParam() {}

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
	PromptData param.Field[PromptDataParam] `json:"prompt_data"`
	// A list of tags for the prompt
	Tags param.Field[[]string] `json:"tags"`
}

func (r PromptNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParams struct {
	// Textual description of the prompt
	Description param.Field[string] `json:"description"`
	// Name of the prompt
	Name param.Field[string] `json:"name"`
	// The prompt, model, and its parameters
	PromptData param.Field[PromptDataParam] `json:"prompt_data"`
	// A list of tags for the prompt
	Tags param.Field[[]string] `json:"tags"`
}

func (r PromptUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	PromptData param.Field[PromptDataParam] `json:"prompt_data"`
	// A list of tags for the prompt
	Tags param.Field[[]string] `json:"tags"`
}

func (r PromptReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
