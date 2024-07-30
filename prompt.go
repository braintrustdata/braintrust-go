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
	// [PromptPromptDataOptionsParamsPromptDataOptions0ResponseFormat].
	ResponseFormat interface{} `json:"response_format,required"`
	// This field can have the runtime type of
	// [PromptPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion].
	ToolChoice interface{} `json:"tool_choice,required"`
	// This field can have the runtime type of
	// [PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion].
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
// [PromptPromptDataOptionsParamsPromptDataOptions0],
// [PromptPromptDataOptionsParamsPromptDataOptions1],
// [PromptPromptDataOptionsParamsPromptDataOptions2],
// [PromptPromptDataOptionsParamsPromptDataOptions3],
// [PromptPromptDataOptionsParamsJsCompletionParams].
func (r PromptPromptDataOptionsParams) AsUnion() PromptPromptDataOptionsParamsUnion {
	return r.union
}

// Union satisfied by [PromptPromptDataOptionsParamsPromptDataOptions0],
// [PromptPromptDataOptionsParamsPromptDataOptions1],
// [PromptPromptDataOptionsParamsPromptDataOptions2],
// [PromptPromptDataOptionsParamsPromptDataOptions3] or
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
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsPromptDataOptions0{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsPromptDataOptions1{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsPromptDataOptions2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsPromptDataOptions3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsJsCompletionParams{}),
		},
	)
}

type PromptPromptDataOptionsParamsPromptDataOptions0 struct {
	FrequencyPenalty float64                                                          `json:"frequency_penalty"`
	FunctionCall     PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion `json:"function_call"`
	MaxTokens        float64                                                          `json:"max_tokens"`
	N                float64                                                          `json:"n"`
	PresencePenalty  float64                                                          `json:"presence_penalty"`
	ResponseFormat   PromptPromptDataOptionsParamsPromptDataOptions0ResponseFormat    `json:"response_format,nullable"`
	Stop             []string                                                         `json:"stop"`
	Temperature      float64                                                          `json:"temperature"`
	ToolChoice       PromptPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion   `json:"tool_choice"`
	TopP             float64                                                          `json:"top_p"`
	UseCache         bool                                                             `json:"use_cache"`
	JSON             promptPromptDataOptionsParamsPromptDataOptions0JSON              `json:"-"`
}

// promptPromptDataOptionsParamsPromptDataOptions0JSON contains the JSON metadata
// for the struct [PromptPromptDataOptionsParamsPromptDataOptions0]
type promptPromptDataOptionsParamsPromptDataOptions0JSON struct {
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

func (r *PromptPromptDataOptionsParamsPromptDataOptions0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsPromptDataOptions0JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsPromptDataOptions0) implementsPromptPromptDataOptionsParams() {}

// Union satisfied by
// [PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0],
// [PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1]
// or [PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallObject].
type PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion interface {
	implementsPromptPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallObject{}),
		},
	)
}

type PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 string

const (
	PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 = "auto"
)

func (r PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto:
		return true
	}
	return false
}

func (r PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) implementsPromptPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 string

const (
	PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 = "none"
)

func (r PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None:
		return true
	}
	return false
}

func (r PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) implementsPromptPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallObject struct {
	Name string                                                                `json:"name,required"`
	JSON promptPromptDataOptionsParamsPromptDataOptions0FunctionCallObjectJSON `json:"-"`
}

// promptPromptDataOptionsParamsPromptDataOptions0FunctionCallObjectJSON contains
// the JSON metadata for the struct
// [PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallObject]
type promptPromptDataOptionsParamsPromptDataOptions0FunctionCallObjectJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsPromptDataOptions0FunctionCallObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) implementsPromptPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type PromptPromptDataOptionsParamsPromptDataOptions0ResponseFormat struct {
	Type PromptPromptDataOptionsParamsPromptDataOptions0ResponseFormatType `json:"type,required"`
	JSON promptPromptDataOptionsParamsPromptDataOptions0ResponseFormatJSON `json:"-"`
}

// promptPromptDataOptionsParamsPromptDataOptions0ResponseFormatJSON contains the
// JSON metadata for the struct
// [PromptPromptDataOptionsParamsPromptDataOptions0ResponseFormat]
type promptPromptDataOptionsParamsPromptDataOptions0ResponseFormatJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsPromptDataOptions0ResponseFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsPromptDataOptions0ResponseFormatJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataOptionsParamsPromptDataOptions0ResponseFormatType string

const (
	PromptPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject PromptPromptDataOptionsParamsPromptDataOptions0ResponseFormatType = "json_object"
)

func (r PromptPromptDataOptionsParamsPromptDataOptions0ResponseFormatType) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Union satisfied by
// [PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0],
// [PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1]
// or
// [PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2].
type PromptPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion interface {
	implementsPromptPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2{}),
		},
	)
}

type PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 string

const (
	PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 = "auto"
)

func (r PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto:
		return true
	}
	return false
}

func (r PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) implementsPromptPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 string

const (
	PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 = "none"
)

func (r PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None:
		return true
	}
	return false
}

func (r PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) implementsPromptPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2 struct {
	Function PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function `json:"function,required"`
	Type     PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type     `json:"type,required"`
	JSON     promptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2JSON     `json:"-"`
}

// promptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2JSON
// contains the JSON metadata for the struct
// [PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2]
type promptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2JSON struct {
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) implementsPromptPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function struct {
	Name string                                                                                     `json:"name,required"`
	JSON promptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2FunctionJSON `json:"-"`
}

// promptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2FunctionJSON
// contains the JSON metadata for the struct
// [PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function]
type promptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2FunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2FunctionJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type string

const (
	PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type = "function"
)

func (r PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction:
		return true
	}
	return false
}

type PromptPromptDataOptionsParamsPromptDataOptions1 struct {
	MaxTokens   float64 `json:"max_tokens,required"`
	Temperature float64 `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample float64                                             `json:"max_tokens_to_sample"`
	StopSequences     []string                                            `json:"stop_sequences"`
	TopK              float64                                             `json:"top_k"`
	TopP              float64                                             `json:"top_p"`
	UseCache          bool                                                `json:"use_cache"`
	JSON              promptPromptDataOptionsParamsPromptDataOptions1JSON `json:"-"`
}

// promptPromptDataOptionsParamsPromptDataOptions1JSON contains the JSON metadata
// for the struct [PromptPromptDataOptionsParamsPromptDataOptions1]
type promptPromptDataOptionsParamsPromptDataOptions1JSON struct {
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

func (r *PromptPromptDataOptionsParamsPromptDataOptions1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsPromptDataOptions1JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsPromptDataOptions1) implementsPromptPromptDataOptionsParams() {}

type PromptPromptDataOptionsParamsPromptDataOptions2 struct {
	MaxOutputTokens float64                                             `json:"maxOutputTokens"`
	Temperature     float64                                             `json:"temperature"`
	TopK            float64                                             `json:"topK"`
	TopP            float64                                             `json:"topP"`
	UseCache        bool                                                `json:"use_cache"`
	JSON            promptPromptDataOptionsParamsPromptDataOptions2JSON `json:"-"`
}

// promptPromptDataOptionsParamsPromptDataOptions2JSON contains the JSON metadata
// for the struct [PromptPromptDataOptionsParamsPromptDataOptions2]
type promptPromptDataOptionsParamsPromptDataOptions2JSON struct {
	MaxOutputTokens apijson.Field
	Temperature     apijson.Field
	TopK            apijson.Field
	TopP            apijson.Field
	UseCache        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsPromptDataOptions2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsPromptDataOptions2JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsPromptDataOptions2) implementsPromptPromptDataOptionsParams() {}

type PromptPromptDataOptionsParamsPromptDataOptions3 struct {
	Temperature float64                                             `json:"temperature"`
	TopK        float64                                             `json:"topK"`
	UseCache    bool                                                `json:"use_cache"`
	JSON        promptPromptDataOptionsParamsPromptDataOptions3JSON `json:"-"`
}

// promptPromptDataOptionsParamsPromptDataOptions3JSON contains the JSON metadata
// for the struct [PromptPromptDataOptionsParamsPromptDataOptions3]
type promptPromptDataOptionsParamsPromptDataOptions3JSON struct {
	Temperature apijson.Field
	TopK        apijson.Field
	UseCache    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsPromptDataOptions3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsPromptDataOptions3JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsPromptDataOptions3) implementsPromptPromptDataOptionsParams() {}

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
	// This field can have the runtime type of
	// [[]PromptPromptDataPromptPromptDataPrompt1Message].
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
// Possible runtime types of the union are
// [PromptPromptDataPromptPromptDataPrompt0],
// [PromptPromptDataPromptPromptDataPrompt1],
// [PromptPromptDataPromptPromptDataPrompt2].
func (r PromptPromptDataPrompt) AsUnion() PromptPromptDataPromptUnion {
	return r.union
}

// Union satisfied by [PromptPromptDataPromptPromptDataPrompt0],
// [PromptPromptDataPromptPromptDataPrompt1] or
// [PromptPromptDataPromptPromptDataPrompt2].
type PromptPromptDataPromptUnion interface {
	implementsPromptPromptDataPrompt()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataPromptUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptPromptDataPrompt0{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptPromptDataPrompt1{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptPromptDataPrompt2{}),
		},
	)
}

type PromptPromptDataPromptPromptDataPrompt0 struct {
	Content string                                      `json:"content,required"`
	Type    PromptPromptDataPromptPromptDataPrompt0Type `json:"type,required"`
	JSON    promptPromptDataPromptPromptDataPrompt0JSON `json:"-"`
}

// promptPromptDataPromptPromptDataPrompt0JSON contains the JSON metadata for the
// struct [PromptPromptDataPromptPromptDataPrompt0]
type promptPromptDataPromptPromptDataPrompt0JSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptPromptDataPrompt0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptPromptDataPrompt0JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptPromptDataPrompt0) implementsPromptPromptDataPrompt() {}

type PromptPromptDataPromptPromptDataPrompt0Type string

const (
	PromptPromptDataPromptPromptDataPrompt0TypeCompletion PromptPromptDataPromptPromptDataPrompt0Type = "completion"
)

func (r PromptPromptDataPromptPromptDataPrompt0Type) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptPromptDataPrompt0TypeCompletion:
		return true
	}
	return false
}

type PromptPromptDataPromptPromptDataPrompt1 struct {
	Messages []PromptPromptDataPromptPromptDataPrompt1Message `json:"messages,required"`
	Type     PromptPromptDataPromptPromptDataPrompt1Type      `json:"type,required"`
	Tools    string                                           `json:"tools"`
	JSON     promptPromptDataPromptPromptDataPrompt1JSON      `json:"-"`
}

// promptPromptDataPromptPromptDataPrompt1JSON contains the JSON metadata for the
// struct [PromptPromptDataPromptPromptDataPrompt1]
type promptPromptDataPromptPromptDataPrompt1JSON struct {
	Messages    apijson.Field
	Type        apijson.Field
	Tools       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptPromptDataPrompt1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptPromptDataPrompt1JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptPromptDataPrompt1) implementsPromptPromptDataPrompt() {}

type PromptPromptDataPromptPromptDataPrompt1Message struct {
	// This field can have the runtime type of [string],
	// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion].
	Content interface{}                                         `json:"content,required"`
	Role    PromptPromptDataPromptPromptDataPrompt1MessagesRole `json:"role,required"`
	Name    string                                              `json:"name"`
	// This field can have the runtime type of
	// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall].
	FunctionCall interface{} `json:"function_call,required"`
	// This field can have the runtime type of
	// [[]PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall].
	ToolCalls  interface{}                                        `json:"tool_calls,required"`
	ToolCallID string                                             `json:"tool_call_id"`
	JSON       promptPromptDataPromptPromptDataPrompt1MessageJSON `json:"-"`
	union      PromptPromptDataPromptPromptDataPrompt1MessagesUnion
}

// promptPromptDataPromptPromptDataPrompt1MessageJSON contains the JSON metadata
// for the struct [PromptPromptDataPromptPromptDataPrompt1Message]
type promptPromptDataPromptPromptDataPrompt1MessageJSON struct {
	Content      apijson.Field
	Role         apijson.Field
	Name         apijson.Field
	FunctionCall apijson.Field
	ToolCalls    apijson.Field
	ToolCallID   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r promptPromptDataPromptPromptDataPrompt1MessageJSON) RawJSON() string {
	return r.raw
}

func (r *PromptPromptDataPromptPromptDataPrompt1Message) UnmarshalJSON(data []byte) (err error) {
	*r = PromptPromptDataPromptPromptDataPrompt1Message{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PromptPromptDataPromptPromptDataPrompt1MessagesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0],
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1],
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2],
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3],
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4],
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5].
func (r PromptPromptDataPromptPromptDataPrompt1Message) AsUnion() PromptPromptDataPromptPromptDataPrompt1MessagesUnion {
	return r.union
}

// Union satisfied by
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0],
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1],
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2],
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3],
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4] or
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5].
type PromptPromptDataPromptPromptDataPrompt1MessagesUnion interface {
	implementsPromptPromptDataPromptPromptDataPrompt1Message()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataPromptPromptDataPrompt1MessagesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5{}),
		},
	)
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0 struct {
	Role    PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role `json:"role,required"`
	Content string                                                                      `json:"content"`
	Name    string                                                                      `json:"name"`
	JSON    promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0JSON `json:"-"`
}

// promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0JSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0]
type promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) implementsPromptPromptDataPromptPromptDataPrompt1Message() {
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role string

const (
	PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role = "system"
)

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1 struct {
	Role    PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role         `json:"role,required"`
	Content PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion `json:"content"`
	Name    string                                                                              `json:"name"`
	JSON    promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1JSON         `json:"-"`
}

// promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1JSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1]
type promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) implementsPromptPromptDataPromptPromptDataPrompt1Message() {
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role string

const (
	PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role = "user"
)

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray].
type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsPromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray{}),
		},
	)
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray []PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) ImplementsPromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion() {
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem struct {
	Text string                                                                                  `json:"text"`
	Type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType `json:"type,required"`
	// This field can have the runtime type of
	// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL].
	ImageURL interface{}                                                                                 `json:"image_url,required"`
	JSON     promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItemJSON `json:"-"`
	union    PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnionItem
}

// promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItemJSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem]
type promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItemJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	ImageURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItemJSON) RawJSON() string {
	return r.raw
}

func (r *PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem) UnmarshalJSON(data []byte) (err error) {
	*r = PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnionItem]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList].
func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem) AsUnion() PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnionItem {
	return r.union
}

// Union satisfied by
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent]
// or
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList].
type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnionItem interface {
	implementsPromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList{}),
		},
	)
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType `json:"type,required"`
	Text string                                                                                                                `json:"text"`
	JSON promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON `json:"-"`
}

// promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent]
type promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON struct {
	Type        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsPromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem() {
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL `json:"image_url,required"`
	Type     PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType     `json:"type,required"`
	JSON     promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON     `json:"-"`
}

// promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList]
type promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON struct {
	ImageURL    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsPromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayItem() {
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    string                                                                                                                       `json:"url,required"`
	Detail PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail `json:"detail"`
	JSON   promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON   `json:"-"`
}

// promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL]
type promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON struct {
	URL         apijson.Field
	Detail      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType string

const (
	PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText     PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "text"
	PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText, PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2 struct {
	Role         PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role         `json:"role,required"`
	Content      string                                                                              `json:"content,nullable"`
	FunctionCall PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall `json:"function_call"`
	Name         string                                                                              `json:"name"`
	ToolCalls    []PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall   `json:"tool_calls"`
	JSON         promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2JSON         `json:"-"`
}

// promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2JSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2]
type promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2JSON struct {
	Role         apijson.Field
	Content      apijson.Field
	FunctionCall apijson.Field
	Name         apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) implementsPromptPromptDataPromptPromptDataPrompt1Message() {
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role string

const (
	PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role = "assistant"
)

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments string                                                                                  `json:"arguments,required"`
	Name      string                                                                                  `json:"name,required"`
	JSON      promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCallJSON `json:"-"`
}

// promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCallJSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall]
type promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCallJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall struct {
	ID       string                                                                                   `json:"id,required"`
	Function PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction `json:"function,required"`
	Type     PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType     `json:"type,required"`
	JSON     promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallJSON      `json:"-"`
}

// promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallJSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall]
type promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallJSON struct {
	ID          apijson.Field
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments string                                                                                       `json:"arguments,required"`
	Name      string                                                                                       `json:"name,required"`
	JSON      promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunctionJSON `json:"-"`
}

// promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunctionJSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction]
type promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunctionJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType string

const (
	PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3 struct {
	Role       PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role `json:"role,required"`
	Content    string                                                                      `json:"content"`
	ToolCallID string                                                                      `json:"tool_call_id"`
	JSON       promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3JSON `json:"-"`
}

// promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3JSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3]
type promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	ToolCallID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) implementsPromptPromptDataPromptPromptDataPrompt1Message() {
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role string

const (
	PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role = "tool"
)

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4 struct {
	Name    string                                                                      `json:"name,required"`
	Role    PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role `json:"role,required"`
	Content string                                                                      `json:"content"`
	JSON    promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4JSON `json:"-"`
}

// promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4JSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4]
type promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4JSON struct {
	Name        apijson.Field
	Role        apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) implementsPromptPromptDataPromptPromptDataPrompt1Message() {
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role string

const (
	PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role = "function"
)

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5 struct {
	Role    PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role `json:"role,required"`
	Content string                                                                      `json:"content,nullable"`
	JSON    promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5JSON `json:"-"`
}

// promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5JSON
// contains the JSON metadata for the struct
// [PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5]
type promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5JSON struct {
	Role        apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) implementsPromptPromptDataPromptPromptDataPrompt1Message() {
}

type PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role string

const (
	PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role = "model"
)

func (r PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type PromptPromptDataPromptPromptDataPrompt1MessagesRole string

const (
	PromptPromptDataPromptPromptDataPrompt1MessagesRoleSystem    PromptPromptDataPromptPromptDataPrompt1MessagesRole = "system"
	PromptPromptDataPromptPromptDataPrompt1MessagesRoleUser      PromptPromptDataPromptPromptDataPrompt1MessagesRole = "user"
	PromptPromptDataPromptPromptDataPrompt1MessagesRoleAssistant PromptPromptDataPromptPromptDataPrompt1MessagesRole = "assistant"
	PromptPromptDataPromptPromptDataPrompt1MessagesRoleTool      PromptPromptDataPromptPromptDataPrompt1MessagesRole = "tool"
	PromptPromptDataPromptPromptDataPrompt1MessagesRoleFunction  PromptPromptDataPromptPromptDataPrompt1MessagesRole = "function"
	PromptPromptDataPromptPromptDataPrompt1MessagesRoleModel     PromptPromptDataPromptPromptDataPrompt1MessagesRole = "model"
)

func (r PromptPromptDataPromptPromptDataPrompt1MessagesRole) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptPromptDataPrompt1MessagesRoleSystem, PromptPromptDataPromptPromptDataPrompt1MessagesRoleUser, PromptPromptDataPromptPromptDataPrompt1MessagesRoleAssistant, PromptPromptDataPromptPromptDataPrompt1MessagesRoleTool, PromptPromptDataPromptPromptDataPrompt1MessagesRoleFunction, PromptPromptDataPromptPromptDataPrompt1MessagesRoleModel:
		return true
	}
	return false
}

type PromptPromptDataPromptPromptDataPrompt1Type string

const (
	PromptPromptDataPromptPromptDataPrompt1TypeChat PromptPromptDataPromptPromptDataPrompt1Type = "chat"
)

func (r PromptPromptDataPromptPromptDataPrompt1Type) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptPromptDataPrompt1TypeChat:
		return true
	}
	return false
}

type PromptPromptDataPromptPromptDataPrompt2 struct {
	JSON promptPromptDataPromptPromptDataPrompt2JSON `json:"-"`
}

// promptPromptDataPromptPromptDataPrompt2JSON contains the JSON metadata for the
// struct [PromptPromptDataPromptPromptDataPrompt2]
type promptPromptDataPromptPromptDataPrompt2JSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptPromptDataPrompt2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptPromptDataPrompt2JSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptPromptDataPrompt2) implementsPromptPromptDataPrompt() {}

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

// Satisfied by [PromptNewParamsPromptDataOptionsParamsPromptDataOptions0],
// [PromptNewParamsPromptDataOptionsParamsPromptDataOptions1],
// [PromptNewParamsPromptDataOptionsParamsPromptDataOptions2],
// [PromptNewParamsPromptDataOptionsParamsPromptDataOptions3],
// [PromptNewParamsPromptDataOptionsParamsJsCompletionParams],
// [PromptNewParamsPromptDataOptionsParams].
type PromptNewParamsPromptDataOptionsParamsUnion interface {
	implementsPromptNewParamsPromptDataOptionsParamsUnion()
}

type PromptNewParamsPromptDataOptionsParamsPromptDataOptions0 struct {
	FrequencyPenalty param.Field[float64]                                                                   `json:"frequency_penalty"`
	FunctionCall     param.Field[PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                                   `json:"max_tokens"`
	N                param.Field[float64]                                                                   `json:"n"`
	PresencePenalty  param.Field[float64]                                                                   `json:"presence_penalty"`
	ResponseFormat   param.Field[PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                                  `json:"stop"`
	Temperature      param.Field[float64]                                                                   `json:"temperature"`
	ToolChoice       param.Field[PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                                   `json:"top_p"`
	UseCache         param.Field[bool]                                                                      `json:"use_cache"`
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0) implementsPromptNewParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by
// [PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0],
// [PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1],
// [PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject].
type PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion interface {
	implementsPromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion()
}

type PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 string

const (
	PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 = "auto"
)

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto:
		return true
	}
	return false
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) implementsPromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 string

const (
	PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 = "none"
)

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None:
		return true
	}
	return false
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) implementsPromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) implementsPromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat struct {
	Type param.Field[PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType] `json:"type,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType string

const (
	PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType = "json_object"
)

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Satisfied by
// [PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0],
// [PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1],
// [PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2].
type PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion interface {
	implementsPromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion()
}

type PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 string

const (
	PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 = "auto"
)

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto:
		return true
	}
	return false
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) implementsPromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 string

const (
	PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 = "none"
)

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None:
		return true
	}
	return false
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) implementsPromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2 struct {
	Function param.Field[PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function] `json:"function,required"`
	Type     param.Field[PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type]     `json:"type,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) implementsPromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type string

const (
	PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type = "function"
)

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction:
		return true
	}
	return false
}

type PromptNewParamsPromptDataOptionsParamsPromptDataOptions1 struct {
	MaxTokens   param.Field[float64] `json:"max_tokens,required"`
	Temperature param.Field[float64] `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64]  `json:"max_tokens_to_sample"`
	StopSequences     param.Field[[]string] `json:"stop_sequences"`
	TopK              param.Field[float64]  `json:"top_k"`
	TopP              param.Field[float64]  `json:"top_p"`
	UseCache          param.Field[bool]     `json:"use_cache"`
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions1) implementsPromptNewParamsPromptDataOptionsParamsUnion() {
}

type PromptNewParamsPromptDataOptionsParamsPromptDataOptions2 struct {
	MaxOutputTokens param.Field[float64] `json:"maxOutputTokens"`
	Temperature     param.Field[float64] `json:"temperature"`
	TopK            param.Field[float64] `json:"topK"`
	TopP            param.Field[float64] `json:"topP"`
	UseCache        param.Field[bool]    `json:"use_cache"`
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions2) implementsPromptNewParamsPromptDataOptionsParamsUnion() {
}

type PromptNewParamsPromptDataOptionsParamsPromptDataOptions3 struct {
	Temperature param.Field[float64] `json:"temperature"`
	TopK        param.Field[float64] `json:"topK"`
	UseCache    param.Field[bool]    `json:"use_cache"`
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsPromptDataOptions3) implementsPromptNewParamsPromptDataOptionsParamsUnion() {
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

// Satisfied by [PromptNewParamsPromptDataPromptPromptDataPrompt0],
// [PromptNewParamsPromptDataPromptPromptDataPrompt1],
// [PromptNewParamsPromptDataPromptPromptDataPrompt2],
// [PromptNewParamsPromptDataPrompt].
type PromptNewParamsPromptDataPromptUnion interface {
	implementsPromptNewParamsPromptDataPromptUnion()
}

type PromptNewParamsPromptDataPromptPromptDataPrompt0 struct {
	Content param.Field[string]                                               `json:"content,required"`
	Type    param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt0Type] `json:"type,required"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt0) implementsPromptNewParamsPromptDataPromptUnion() {
}

type PromptNewParamsPromptDataPromptPromptDataPrompt0Type string

const (
	PromptNewParamsPromptDataPromptPromptDataPrompt0TypeCompletion PromptNewParamsPromptDataPromptPromptDataPrompt0Type = "completion"
)

func (r PromptNewParamsPromptDataPromptPromptDataPrompt0Type) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptPromptDataPrompt0TypeCompletion:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1 struct {
	Messages param.Field[[]PromptNewParamsPromptDataPromptPromptDataPrompt1MessageUnion] `json:"messages,required"`
	Type     param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1Type]           `json:"type,required"`
	Tools    param.Field[string]                                                         `json:"tools"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1) implementsPromptNewParamsPromptDataPromptUnion() {
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1Message struct {
	Content      param.Field[interface{}]                                                  `json:"content,required"`
	Role         param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRole] `json:"role,required"`
	Name         param.Field[string]                                                       `json:"name"`
	FunctionCall param.Field[interface{}]                                                  `json:"function_call,required"`
	ToolCalls    param.Field[interface{}]                                                  `json:"tool_calls,required"`
	ToolCallID   param.Field[string]                                                       `json:"tool_call_id"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1Message) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1Message) implementsPromptNewParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

// Satisfied by
// [PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0],
// [PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1],
// [PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2],
// [PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3],
// [PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4],
// [PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5],
// [PromptNewParamsPromptDataPromptPromptDataPrompt1Message].
type PromptNewParamsPromptDataPromptPromptDataPrompt1MessageUnion interface {
	implementsPromptNewParamsPromptDataPromptPromptDataPrompt1MessageUnion()
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0 struct {
	Role    param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role] `json:"role,required"`
	Content param.Field[string]                                                                               `json:"content"`
	Name    param.Field[string]                                                                               `json:"name"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) implementsPromptNewParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role string

const (
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role = "system"
)

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1 struct {
	Role    param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role]         `json:"role,required"`
	Content param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion] `json:"content"`
	Name    param.Field[string]                                                                                       `json:"name"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) implementsPromptNewParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role string

const (
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role = "user"
)

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray].
type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsPromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion()
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray []PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) ImplementsPromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion() {
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray struct {
	Text     param.Field[string]                                                                                           `json:"text"`
	Type     param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                                                      `json:"image_url,required"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) implementsPromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

// Satisfied by
// [PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList],
// [PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray].
type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion interface {
	implementsPromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion()
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType] `json:"type,required"`
	Text param.Field[string]                                                                                                                         `json:"text"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsPromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL] `json:"image_url,required"`
	Type     param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType]     `json:"type,required"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsPromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    param.Field[string]                                                                                                                                `json:"url,required"`
	Detail param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail] `json:"detail"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType string

const (
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText     PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "text"
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText, PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2 struct {
	Role         param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role]         `json:"role,required"`
	Content      param.Field[string]                                                                                       `json:"content"`
	FunctionCall param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                                                       `json:"name"`
	ToolCalls    param.Field[[]PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall]   `json:"tool_calls"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) implementsPromptNewParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role string

const (
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role = "assistant"
)

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall struct {
	ID       param.Field[string]                                                                                            `json:"id,required"`
	Function param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction] `json:"function,required"`
	Type     param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType]     `json:"type,required"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType string

const (
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3 struct {
	Role       param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role] `json:"role,required"`
	Content    param.Field[string]                                                                               `json:"content"`
	ToolCallID param.Field[string]                                                                               `json:"tool_call_id"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) implementsPromptNewParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role string

const (
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role = "tool"
)

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4 struct {
	Name    param.Field[string]                                                                               `json:"name,required"`
	Role    param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role] `json:"role,required"`
	Content param.Field[string]                                                                               `json:"content"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) implementsPromptNewParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role string

const (
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role = "function"
)

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5 struct {
	Role    param.Field[PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role] `json:"role,required"`
	Content param.Field[string]                                                                               `json:"content"`
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) implementsPromptNewParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role string

const (
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role = "model"
)

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRole string

const (
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleSystem    PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRole = "system"
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleUser      PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRole = "user"
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleAssistant PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRole = "assistant"
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleTool      PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRole = "tool"
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleFunction  PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRole = "function"
	PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleModel     PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRole = "model"
)

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRole) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleSystem, PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleUser, PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleAssistant, PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleTool, PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleFunction, PromptNewParamsPromptDataPromptPromptDataPrompt1MessagesRoleModel:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptPromptDataPrompt1Type string

const (
	PromptNewParamsPromptDataPromptPromptDataPrompt1TypeChat PromptNewParamsPromptDataPromptPromptDataPrompt1Type = "chat"
)

func (r PromptNewParamsPromptDataPromptPromptDataPrompt1Type) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptPromptDataPrompt1TypeChat:
		return true
	}
	return false
}

type PromptNewParamsPromptDataPromptPromptDataPrompt2 struct {
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptPromptDataPrompt2) implementsPromptNewParamsPromptDataPromptUnion() {
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

// Satisfied by [PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0],
// [PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions1],
// [PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions2],
// [PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions3],
// [PromptUpdateParamsPromptDataOptionsParamsJsCompletionParams],
// [PromptUpdateParamsPromptDataOptionsParams].
type PromptUpdateParamsPromptDataOptionsParamsUnion interface {
	implementsPromptUpdateParamsPromptDataOptionsParamsUnion()
}

type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0 struct {
	FrequencyPenalty param.Field[float64]                                                                      `json:"frequency_penalty"`
	FunctionCall     param.Field[PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                                      `json:"max_tokens"`
	N                param.Field[float64]                                                                      `json:"n"`
	PresencePenalty  param.Field[float64]                                                                      `json:"presence_penalty"`
	ResponseFormat   param.Field[PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                                     `json:"stop"`
	Temperature      param.Field[float64]                                                                      `json:"temperature"`
	ToolChoice       param.Field[PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                                      `json:"top_p"`
	UseCache         param.Field[bool]                                                                         `json:"use_cache"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0) implementsPromptUpdateParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by
// [PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0],
// [PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1],
// [PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject].
type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion interface {
	implementsPromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion()
}

type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 string

const (
	PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 = "auto"
)

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto:
		return true
	}
	return false
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) implementsPromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 string

const (
	PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 = "none"
)

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None:
		return true
	}
	return false
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) implementsPromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) implementsPromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat struct {
	Type param.Field[PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType] `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType string

const (
	PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType = "json_object"
)

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Satisfied by
// [PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0],
// [PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1],
// [PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2].
type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion interface {
	implementsPromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion()
}

type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 string

const (
	PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 = "auto"
)

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto:
		return true
	}
	return false
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) implementsPromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 string

const (
	PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 = "none"
)

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None:
		return true
	}
	return false
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) implementsPromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2 struct {
	Function param.Field[PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function] `json:"function,required"`
	Type     param.Field[PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type]     `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) implementsPromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type string

const (
	PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type = "function"
)

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions1 struct {
	MaxTokens   param.Field[float64] `json:"max_tokens,required"`
	Temperature param.Field[float64] `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64]  `json:"max_tokens_to_sample"`
	StopSequences     param.Field[[]string] `json:"stop_sequences"`
	TopK              param.Field[float64]  `json:"top_k"`
	TopP              param.Field[float64]  `json:"top_p"`
	UseCache          param.Field[bool]     `json:"use_cache"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions1) implementsPromptUpdateParamsPromptDataOptionsParamsUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions2 struct {
	MaxOutputTokens param.Field[float64] `json:"maxOutputTokens"`
	Temperature     param.Field[float64] `json:"temperature"`
	TopK            param.Field[float64] `json:"topK"`
	TopP            param.Field[float64] `json:"topP"`
	UseCache        param.Field[bool]    `json:"use_cache"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions2) implementsPromptUpdateParamsPromptDataOptionsParamsUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions3 struct {
	Temperature param.Field[float64] `json:"temperature"`
	TopK        param.Field[float64] `json:"topK"`
	UseCache    param.Field[bool]    `json:"use_cache"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions3) implementsPromptUpdateParamsPromptDataOptionsParamsUnion() {
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

// Satisfied by [PromptUpdateParamsPromptDataPromptPromptDataPrompt0],
// [PromptUpdateParamsPromptDataPromptPromptDataPrompt1],
// [PromptUpdateParamsPromptDataPromptPromptDataPrompt2],
// [PromptUpdateParamsPromptDataPrompt].
type PromptUpdateParamsPromptDataPromptUnion interface {
	implementsPromptUpdateParamsPromptDataPromptUnion()
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt0 struct {
	Content param.Field[string]                                                  `json:"content,required"`
	Type    param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt0Type] `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt0) implementsPromptUpdateParamsPromptDataPromptUnion() {
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt0Type string

const (
	PromptUpdateParamsPromptDataPromptPromptDataPrompt0TypeCompletion PromptUpdateParamsPromptDataPromptPromptDataPrompt0Type = "completion"
)

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt0Type) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptPromptDataPrompt0TypeCompletion:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1 struct {
	Messages param.Field[[]PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion] `json:"messages,required"`
	Type     param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1Type]           `json:"type,required"`
	Tools    param.Field[string]                                                            `json:"tools"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1) implementsPromptUpdateParamsPromptDataPromptUnion() {
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1Message struct {
	Content      param.Field[interface{}]                                                     `json:"content,required"`
	Role         param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole] `json:"role,required"`
	Name         param.Field[string]                                                          `json:"name"`
	FunctionCall param.Field[interface{}]                                                     `json:"function_call,required"`
	ToolCalls    param.Field[interface{}]                                                     `json:"tool_calls,required"`
	ToolCallID   param.Field[string]                                                          `json:"tool_call_id"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1Message) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1Message) implementsPromptUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

// Satisfied by
// [PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0],
// [PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1],
// [PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2],
// [PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3],
// [PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4],
// [PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5],
// [PromptUpdateParamsPromptDataPromptPromptDataPrompt1Message].
type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion interface {
	implementsPromptUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion()
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0 struct {
	Role    param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role] `json:"role,required"`
	Content param.Field[string]                                                                                  `json:"content"`
	Name    param.Field[string]                                                                                  `json:"name"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) implementsPromptUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role string

const (
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role = "system"
)

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1 struct {
	Role    param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role]         `json:"role,required"`
	Content param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion] `json:"content"`
	Name    param.Field[string]                                                                                          `json:"name"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) implementsPromptUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role string

const (
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role = "user"
)

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray].
type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsPromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion()
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray []PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) ImplementsPromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion() {
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray struct {
	Text     param.Field[string]                                                                                              `json:"text"`
	Type     param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                                                         `json:"image_url,required"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) implementsPromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

// Satisfied by
// [PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList],
// [PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray].
type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion interface {
	implementsPromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion()
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType] `json:"type,required"`
	Text param.Field[string]                                                                                                                            `json:"text"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsPromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL] `json:"image_url,required"`
	Type     param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType]     `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsPromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    param.Field[string]                                                                                                                                   `json:"url,required"`
	Detail param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail] `json:"detail"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType string

const (
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText     PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "text"
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText, PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2 struct {
	Role         param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role]         `json:"role,required"`
	Content      param.Field[string]                                                                                          `json:"content"`
	FunctionCall param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                                                          `json:"name"`
	ToolCalls    param.Field[[]PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall]   `json:"tool_calls"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) implementsPromptUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role string

const (
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role = "assistant"
)

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall struct {
	ID       param.Field[string]                                                                                               `json:"id,required"`
	Function param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction] `json:"function,required"`
	Type     param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType]     `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType string

const (
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3 struct {
	Role       param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role] `json:"role,required"`
	Content    param.Field[string]                                                                                  `json:"content"`
	ToolCallID param.Field[string]                                                                                  `json:"tool_call_id"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) implementsPromptUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role string

const (
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role = "tool"
)

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4 struct {
	Name    param.Field[string]                                                                                  `json:"name,required"`
	Role    param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role] `json:"role,required"`
	Content param.Field[string]                                                                                  `json:"content"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) implementsPromptUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role string

const (
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role = "function"
)

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5 struct {
	Role    param.Field[PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role] `json:"role,required"`
	Content param.Field[string]                                                                                  `json:"content"`
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) implementsPromptUpdateParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role string

const (
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role = "model"
)

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole string

const (
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleSystem    PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole = "system"
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleUser      PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole = "user"
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleAssistant PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole = "assistant"
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleTool      PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole = "tool"
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleFunction  PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole = "function"
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleModel     PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole = "model"
)

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRole) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleSystem, PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleUser, PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleAssistant, PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleTool, PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleFunction, PromptUpdateParamsPromptDataPromptPromptDataPrompt1MessagesRoleModel:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt1Type string

const (
	PromptUpdateParamsPromptDataPromptPromptDataPrompt1TypeChat PromptUpdateParamsPromptDataPromptPromptDataPrompt1Type = "chat"
)

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt1Type) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptPromptDataPrompt1TypeChat:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataPromptPromptDataPrompt2 struct {
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptPromptDataPrompt2) implementsPromptUpdateParamsPromptDataPromptUnion() {
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

// Satisfied by [PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0],
// [PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions1],
// [PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions2],
// [PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions3],
// [PromptReplaceParamsPromptDataOptionsParamsJsCompletionParams],
// [PromptReplaceParamsPromptDataOptionsParams].
type PromptReplaceParamsPromptDataOptionsParamsUnion interface {
	implementsPromptReplaceParamsPromptDataOptionsParamsUnion()
}

type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0 struct {
	FrequencyPenalty param.Field[float64]                                                                       `json:"frequency_penalty"`
	FunctionCall     param.Field[PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                                       `json:"max_tokens"`
	N                param.Field[float64]                                                                       `json:"n"`
	PresencePenalty  param.Field[float64]                                                                       `json:"presence_penalty"`
	ResponseFormat   param.Field[PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                                      `json:"stop"`
	Temperature      param.Field[float64]                                                                       `json:"temperature"`
	ToolChoice       param.Field[PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                                       `json:"top_p"`
	UseCache         param.Field[bool]                                                                          `json:"use_cache"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0) implementsPromptReplaceParamsPromptDataOptionsParamsUnion() {
}

// Satisfied by
// [PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0],
// [PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1],
// [PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject].
type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion interface {
	implementsPromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion()
}

type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 string

const (
	PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0 = "auto"
)

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto:
		return true
	}
	return false
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0) implementsPromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 string

const (
	PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1 = "none"
)

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1None:
		return true
	}
	return false
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall1) implementsPromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallObject) implementsPromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat struct {
	Type param.Field[PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType] `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType string

const (
	PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType = "json_object"
)

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Satisfied by
// [PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0],
// [PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1],
// [PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2].
type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion interface {
	implementsPromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion()
}

type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 string

const (
	PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0 = "auto"
)

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto:
		return true
	}
	return false
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0) implementsPromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 string

const (
	PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1 = "none"
)

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1None:
		return true
	}
	return false
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice1) implementsPromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2 struct {
	Function param.Field[PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function] `json:"function,required"`
	Type     param.Field[PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type]     `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2) implementsPromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Function) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type string

const (
	PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type = "function"
)

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2Type) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice2TypeFunction:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions1 struct {
	MaxTokens   param.Field[float64] `json:"max_tokens,required"`
	Temperature param.Field[float64] `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Field[float64]  `json:"max_tokens_to_sample"`
	StopSequences     param.Field[[]string] `json:"stop_sequences"`
	TopK              param.Field[float64]  `json:"top_k"`
	TopP              param.Field[float64]  `json:"top_p"`
	UseCache          param.Field[bool]     `json:"use_cache"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions1) implementsPromptReplaceParamsPromptDataOptionsParamsUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions2 struct {
	MaxOutputTokens param.Field[float64] `json:"maxOutputTokens"`
	Temperature     param.Field[float64] `json:"temperature"`
	TopK            param.Field[float64] `json:"topK"`
	TopP            param.Field[float64] `json:"topP"`
	UseCache        param.Field[bool]    `json:"use_cache"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions2) implementsPromptReplaceParamsPromptDataOptionsParamsUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions3 struct {
	Temperature param.Field[float64] `json:"temperature"`
	TopK        param.Field[float64] `json:"topK"`
	UseCache    param.Field[bool]    `json:"use_cache"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions3) implementsPromptReplaceParamsPromptDataOptionsParamsUnion() {
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

// Satisfied by [PromptReplaceParamsPromptDataPromptPromptDataPrompt0],
// [PromptReplaceParamsPromptDataPromptPromptDataPrompt1],
// [PromptReplaceParamsPromptDataPromptPromptDataPrompt2],
// [PromptReplaceParamsPromptDataPrompt].
type PromptReplaceParamsPromptDataPromptUnion interface {
	implementsPromptReplaceParamsPromptDataPromptUnion()
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt0 struct {
	Content param.Field[string]                                                   `json:"content,required"`
	Type    param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt0Type] `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt0) implementsPromptReplaceParamsPromptDataPromptUnion() {
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt0Type string

const (
	PromptReplaceParamsPromptDataPromptPromptDataPrompt0TypeCompletion PromptReplaceParamsPromptDataPromptPromptDataPrompt0Type = "completion"
)

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt0Type) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptPromptDataPrompt0TypeCompletion:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1 struct {
	Messages param.Field[[]PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion] `json:"messages,required"`
	Type     param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1Type]           `json:"type,required"`
	Tools    param.Field[string]                                                             `json:"tools"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1) implementsPromptReplaceParamsPromptDataPromptUnion() {
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1Message struct {
	Content      param.Field[interface{}]                                                      `json:"content,required"`
	Role         param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole] `json:"role,required"`
	Name         param.Field[string]                                                           `json:"name"`
	FunctionCall param.Field[interface{}]                                                      `json:"function_call,required"`
	ToolCalls    param.Field[interface{}]                                                      `json:"tool_calls,required"`
	ToolCallID   param.Field[string]                                                           `json:"tool_call_id"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1Message) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1Message) implementsPromptReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

// Satisfied by
// [PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0],
// [PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1],
// [PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2],
// [PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3],
// [PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4],
// [PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5],
// [PromptReplaceParamsPromptDataPromptPromptDataPrompt1Message].
type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion interface {
	implementsPromptReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion()
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0 struct {
	Role    param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role] `json:"role,required"`
	Content param.Field[string]                                                                                   `json:"content"`
	Name    param.Field[string]                                                                                   `json:"name"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0) implementsPromptReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role string

const (
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role = "system"
)

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0Role) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage0RoleSystem:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1 struct {
	Role    param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role]         `json:"role,required"`
	Content param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion] `json:"content"`
	Name    param.Field[string]                                                                                           `json:"name"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1) implementsPromptReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role string

const (
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role = "user"
)

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1Role) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1RoleUser:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray].
type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion interface {
	ImplementsPromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion()
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray []PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) ImplementsPromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentUnion() {
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray struct {
	Text     param.Field[string]                                                                                               `json:"text"`
	Type     param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                                                                          `json:"image_url,required"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray) implementsPromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

// Satisfied by
// [PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent],
// [PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList],
// [PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArray].
type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion interface {
	implementsPromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion()
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent struct {
	Type param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType] `json:"type,required"`
	Text param.Field[string]                                                                                                                             `json:"text"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContent) implementsPromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType string

const (
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType = "text"
)

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageContentTypeText:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList struct {
	ImageURL param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL] `json:"image_url,required"`
	Type     param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType]     `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageList) implementsPromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayUnion() {
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL struct {
	URL    param.Field[string]                                                                                                                                    `json:"url,required"`
	Detail param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail] `json:"detail"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail string

const (
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "auto"
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow  PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "low"
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail = "high"
)

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetail) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailAuto, PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailLow, PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListImageURLDetailHigh:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType string

const (
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType = "image_url"
)

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayPromptDataPromptMessageListTypeImageURL:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType string

const (
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText     PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "text"
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType = "image_url"
)

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeText, PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage1ContentArrayTypeImageURL:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2 struct {
	Role         param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role]         `json:"role,required"`
	Content      param.Field[string]                                                                                           `json:"content"`
	FunctionCall param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                                                           `json:"name"`
	ToolCalls    param.Field[[]PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall]   `json:"tool_calls"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2) implementsPromptReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role string

const (
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role = "assistant"
)

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2Role) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2RoleAssistant:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2FunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall struct {
	ID       param.Field[string]                                                                                                `json:"id,required"`
	Function param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction] `json:"function,required"`
	Type     param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType]     `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType string

const (
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType = "function"
)

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage2ToolCallsTypeFunction:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3 struct {
	Role       param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role] `json:"role,required"`
	Content    param.Field[string]                                                                                   `json:"content"`
	ToolCallID param.Field[string]                                                                                   `json:"tool_call_id"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3) implementsPromptReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role string

const (
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role = "tool"
)

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3Role) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage3RoleTool:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4 struct {
	Name    param.Field[string]                                                                                   `json:"name,required"`
	Role    param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role] `json:"role,required"`
	Content param.Field[string]                                                                                   `json:"content"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4) implementsPromptReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role string

const (
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role = "function"
)

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4Role) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage4RoleFunction:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5 struct {
	Role    param.Field[PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role] `json:"role,required"`
	Content param.Field[string]                                                                                   `json:"content"`
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5) implementsPromptReplaceParamsPromptDataPromptPromptDataPrompt1MessageUnion() {
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role string

const (
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role = "model"
)

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5Role) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesPromptDataPromptMessage5RoleModel:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole string

const (
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleSystem    PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole = "system"
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleUser      PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole = "user"
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleAssistant PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole = "assistant"
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleTool      PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole = "tool"
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleFunction  PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole = "function"
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleModel     PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole = "model"
)

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRole) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleSystem, PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleUser, PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleAssistant, PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleTool, PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleFunction, PromptReplaceParamsPromptDataPromptPromptDataPrompt1MessagesRoleModel:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt1Type string

const (
	PromptReplaceParamsPromptDataPromptPromptDataPrompt1TypeChat PromptReplaceParamsPromptDataPromptPromptDataPrompt1Type = "chat"
)

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt1Type) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptPromptDataPrompt1TypeChat:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataPromptPromptDataPrompt2 struct {
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptPromptDataPrompt2) implementsPromptReplaceParamsPromptDataPromptUnion() {
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
