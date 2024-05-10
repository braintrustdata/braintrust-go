// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
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
	path := fmt.Sprintf("v1/prompt/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update a prompt object. Specify the fields to update in the payload.
// Any object-type fields will be deep-merged with existing content. Currently we
// do not support removing fields or setting them to null.
func (r *PromptService) Update(ctx context.Context, promptID string, body PromptUpdateParams, opts ...option.RequestOption) (res *Prompt, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/prompt/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all prompts. The prompts are sorted by creation date, with the most
// recently-created prompts coming first
func (r *PromptService) List(ctx context.Context, query PromptListParams, opts ...option.RequestOption) (res *pagination.ListObjects[Prompt], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
	path := fmt.Sprintf("v1/prompt/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Log feedback for a set of prompt events
func (r *PromptService) Feedback(ctx context.Context, promptID string, body PromptFeedbackParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("v1/prompt/%s/feedback", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// NOTE: This operation is deprecated and will be removed in a future revision of
// the API. Create or replace a new prompt. If there is an existing prompt in the
// project with the same slug as the one specified in the request, will return the
// existing prompt unmodified, will replace the existing prompt with the provided
// fields
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
	UseCache         bool        `json:"use_cache"`
	Temperature      float64     `json:"temperature"`
	TopP             float64     `json:"top_p"`
	MaxTokens        float64     `json:"max_tokens"`
	FrequencyPenalty float64     `json:"frequency_penalty"`
	PresencePenalty  float64     `json:"presence_penalty"`
	ResponseFormat   interface{} `json:"response_format,required"`
	ToolChoice       interface{} `json:"tool_choice,required"`
	FunctionCall     interface{} `json:"function_call,required"`
	N                float64     `json:"n"`
	Stop             interface{} `json:"stop,required"`
	TopK             float64     `json:"top_k"`
	StopSequences    interface{} `json:"stop_sequences,required"`
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
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PromptPromptDataOptionsParams) AsUnion() PromptPromptDataOptionsParamsUnion {
	return r.union
}

// Union satisfied by [PromptPromptDataOptionsParamsObject],
// [PromptPromptDataOptionsParamsObject], [PromptPromptDataOptionsParamsObject] or
// [PromptPromptDataOptionsParamsObject].
type PromptPromptDataOptionsParamsUnion interface {
	implementsPromptPromptDataOptionsParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataOptionsParamsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsObject{}),
		},
	)
}

type PromptPromptDataOptionsParamsObject struct {
	FrequencyPenalty float64                                              `json:"frequency_penalty"`
	FunctionCall     PromptPromptDataOptionsParamsObjectFunctionCallUnion `json:"function_call"`
	MaxTokens        float64                                              `json:"max_tokens"`
	N                float64                                              `json:"n"`
	PresencePenalty  float64                                              `json:"presence_penalty"`
	ResponseFormat   PromptPromptDataOptionsParamsObjectResponseFormat    `json:"response_format,nullable"`
	Stop             []string                                             `json:"stop"`
	Temperature      float64                                              `json:"temperature"`
	ToolChoice       PromptPromptDataOptionsParamsObjectToolChoiceUnion   `json:"tool_choice"`
	TopP             float64                                              `json:"top_p"`
	UseCache         bool                                                 `json:"use_cache"`
	JSON             promptPromptDataOptionsParamsObjectJSON              `json:"-"`
}

// promptPromptDataOptionsParamsObjectJSON contains the JSON metadata for the
// struct [PromptPromptDataOptionsParamsObject]
type promptPromptDataOptionsParamsObjectJSON struct {
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

func (r *PromptPromptDataOptionsParamsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsObject) implementsPromptPromptDataOptionsParams() {}

// Union satisfied by [PromptPromptDataOptionsParamsObjectFunctionCallString],
// [PromptPromptDataOptionsParamsObjectFunctionCallString] or
// [PromptPromptDataOptionsParamsObjectFunctionCallObject].
type PromptPromptDataOptionsParamsObjectFunctionCallUnion interface {
	ImplementsPromptPromptDataOptionsParamsObjectFunctionCallUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataOptionsParamsObjectFunctionCallUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsObjectFunctionCallObject{}),
		},
	)
}

type PromptPromptDataOptionsParamsObjectFunctionCallString string

const (
	PromptPromptDataOptionsParamsObjectFunctionCallStringAuto PromptPromptDataOptionsParamsObjectFunctionCallString = "auto"
)

func (r PromptPromptDataOptionsParamsObjectFunctionCallString) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsObjectFunctionCallStringAuto:
		return true
	}
	return false
}

type PromptPromptDataOptionsParamsObjectFunctionCallObject struct {
	Name string                                                    `json:"name,required"`
	JSON promptPromptDataOptionsParamsObjectFunctionCallObjectJSON `json:"-"`
}

// promptPromptDataOptionsParamsObjectFunctionCallObjectJSON contains the JSON
// metadata for the struct [PromptPromptDataOptionsParamsObjectFunctionCallObject]
type promptPromptDataOptionsParamsObjectFunctionCallObjectJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsObjectFunctionCallObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsObjectFunctionCallObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsObjectFunctionCallObject) ImplementsPromptPromptDataOptionsParamsObjectFunctionCallUnion() {
}

type PromptPromptDataOptionsParamsObjectResponseFormat struct {
	Type PromptPromptDataOptionsParamsObjectResponseFormatType `json:"type,required"`
	JSON promptPromptDataOptionsParamsObjectResponseFormatJSON `json:"-"`
}

// promptPromptDataOptionsParamsObjectResponseFormatJSON contains the JSON metadata
// for the struct [PromptPromptDataOptionsParamsObjectResponseFormat]
type promptPromptDataOptionsParamsObjectResponseFormatJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsObjectResponseFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsObjectResponseFormatJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataOptionsParamsObjectResponseFormatType string

const (
	PromptPromptDataOptionsParamsObjectResponseFormatTypeJsonObject PromptPromptDataOptionsParamsObjectResponseFormatType = "json_object"
)

func (r PromptPromptDataOptionsParamsObjectResponseFormatType) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsObjectResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Union satisfied by [PromptPromptDataOptionsParamsObjectToolChoiceString],
// [PromptPromptDataOptionsParamsObjectToolChoiceString] or
// [PromptPromptDataOptionsParamsObjectToolChoiceObject].
type PromptPromptDataOptionsParamsObjectToolChoiceUnion interface {
	ImplementsPromptPromptDataOptionsParamsObjectToolChoiceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataOptionsParamsObjectToolChoiceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataOptionsParamsObjectToolChoiceObject{}),
		},
	)
}

type PromptPromptDataOptionsParamsObjectToolChoiceString string

const (
	PromptPromptDataOptionsParamsObjectToolChoiceStringAuto PromptPromptDataOptionsParamsObjectToolChoiceString = "auto"
)

func (r PromptPromptDataOptionsParamsObjectToolChoiceString) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsObjectToolChoiceStringAuto:
		return true
	}
	return false
}

type PromptPromptDataOptionsParamsObjectToolChoiceObject struct {
	Function PromptPromptDataOptionsParamsObjectToolChoiceObjectFunction `json:"function,required"`
	Type     PromptPromptDataOptionsParamsObjectToolChoiceObjectType     `json:"type,required"`
	JSON     promptPromptDataOptionsParamsObjectToolChoiceObjectJSON     `json:"-"`
}

// promptPromptDataOptionsParamsObjectToolChoiceObjectJSON contains the JSON
// metadata for the struct [PromptPromptDataOptionsParamsObjectToolChoiceObject]
type promptPromptDataOptionsParamsObjectToolChoiceObjectJSON struct {
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsObjectToolChoiceObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsObjectToolChoiceObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataOptionsParamsObjectToolChoiceObject) ImplementsPromptPromptDataOptionsParamsObjectToolChoiceUnion() {
}

type PromptPromptDataOptionsParamsObjectToolChoiceObjectFunction struct {
	Name string                                                          `json:"name,required"`
	JSON promptPromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON `json:"-"`
}

// promptPromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON contains the
// JSON metadata for the struct
// [PromptPromptDataOptionsParamsObjectToolChoiceObjectFunction]
type promptPromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataOptionsParamsObjectToolChoiceObjectFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON) RawJSON() string {
	return r.raw
}

type PromptPromptDataOptionsParamsObjectToolChoiceObjectType string

const (
	PromptPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction PromptPromptDataOptionsParamsObjectToolChoiceObjectType = "function"
)

func (r PromptPromptDataOptionsParamsObjectToolChoiceObjectType) IsKnown() bool {
	switch r {
	case PromptPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction:
		return true
	}
	return false
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
	Type                  PromptPromptDataPromptType `json:"type"`
	Content               string                     `json:"content"`
	Messages              interface{}                `json:"messages,required"`
	Tools                 string                     `json:"tools"`
	ReservedOnlyAllowNull interface{}                `json:"__reserved_only_allow_null,required"`
	JSON                  promptPromptDataPromptJSON `json:"-"`
	union                 PromptPromptDataPromptUnion
}

// promptPromptDataPromptJSON contains the JSON metadata for the struct
// [PromptPromptDataPrompt]
type promptPromptDataPromptJSON struct {
	Type                  apijson.Field
	Content               apijson.Field
	Messages              apijson.Field
	Tools                 apijson.Field
	ReservedOnlyAllowNull apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r promptPromptDataPromptJSON) RawJSON() string {
	return r.raw
}

func (r *PromptPromptDataPrompt) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PromptPromptDataPrompt) AsUnion() PromptPromptDataPromptUnion {
	return r.union
}

// Union satisfied by [PromptPromptDataPromptObject],
// [PromptPromptDataPromptObject] or [PromptPromptDataPromptObject].
type PromptPromptDataPromptUnion interface {
	implementsPromptPromptDataPrompt()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataPromptUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptObject{}),
		},
	)
}

type PromptPromptDataPromptObject struct {
	Content string                           `json:"content,required"`
	Type    PromptPromptDataPromptObjectType `json:"type,required"`
	JSON    promptPromptDataPromptObjectJSON `json:"-"`
}

// promptPromptDataPromptObjectJSON contains the JSON metadata for the struct
// [PromptPromptDataPromptObject]
type promptPromptDataPromptObjectJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptObject) implementsPromptPromptDataPrompt() {}

type PromptPromptDataPromptObjectType string

const (
	PromptPromptDataPromptObjectTypeCompletion PromptPromptDataPromptObjectType = "completion"
)

func (r PromptPromptDataPromptObjectType) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptObjectTypeCompletion:
		return true
	}
	return false
}

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

// Satisfied by [PromptNewParamsPromptDataOptionsParamsObject],
// [PromptNewParamsPromptDataOptionsParamsObject],
// [PromptNewParamsPromptDataOptionsParamsObject],
// [PromptNewParamsPromptDataOptionsParamsObject],
// [PromptNewParamsPromptDataOptionsParams].
type PromptNewParamsPromptDataOptionsParamsUnion interface {
	implementsPromptNewParamsPromptDataOptionsParamsUnion()
}

type PromptNewParamsPromptDataOptionsParamsObject struct {
	FrequencyPenalty param.Field[float64]                                                       `json:"frequency_penalty"`
	FunctionCall     param.Field[PromptNewParamsPromptDataOptionsParamsObjectFunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                       `json:"max_tokens"`
	N                param.Field[float64]                                                       `json:"n"`
	PresencePenalty  param.Field[float64]                                                       `json:"presence_penalty"`
	ResponseFormat   param.Field[PromptNewParamsPromptDataOptionsParamsObjectResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                      `json:"stop"`
	Temperature      param.Field[float64]                                                       `json:"temperature"`
	ToolChoice       param.Field[PromptNewParamsPromptDataOptionsParamsObjectToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                       `json:"top_p"`
	UseCache         param.Field[bool]                                                          `json:"use_cache"`
}

func (r PromptNewParamsPromptDataOptionsParamsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsObject) implementsPromptNewParamsPromptDataOptionsParamsUnion() {
}

type PromptNewParamsPromptDataOptionsParamsObjectFunctionCall struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsObjectFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsObjectFunctionCall) ImplementsPromptNewParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
}

// Satisfied by [PromptNewParamsPromptDataOptionsParamsObjectFunctionCallString],
// [PromptNewParamsPromptDataOptionsParamsObjectFunctionCallString],
// [PromptNewParamsPromptDataOptionsParamsObjectFunctionCallObject],
// [PromptNewParamsPromptDataOptionsParamsObjectFunctionCall].
type PromptNewParamsPromptDataOptionsParamsObjectFunctionCallUnion interface {
	ImplementsPromptNewParamsPromptDataOptionsParamsObjectFunctionCallUnion()
}

type PromptNewParamsPromptDataOptionsParamsObjectFunctionCallString string

const (
	PromptNewParamsPromptDataOptionsParamsObjectFunctionCallStringAuto PromptNewParamsPromptDataOptionsParamsObjectFunctionCallString = "auto"
)

func (r PromptNewParamsPromptDataOptionsParamsObjectFunctionCallString) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsObjectFunctionCallStringAuto:
		return true
	}
	return false
}

type PromptNewParamsPromptDataOptionsParamsObjectFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsObjectFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsObjectFunctionCallObject) ImplementsPromptNewParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
}

type PromptNewParamsPromptDataOptionsParamsObjectResponseFormat struct {
	Type param.Field[PromptNewParamsPromptDataOptionsParamsObjectResponseFormatType] `json:"type,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsObjectResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataOptionsParamsObjectResponseFormatType string

const (
	PromptNewParamsPromptDataOptionsParamsObjectResponseFormatTypeJsonObject PromptNewParamsPromptDataOptionsParamsObjectResponseFormatType = "json_object"
)

func (r PromptNewParamsPromptDataOptionsParamsObjectResponseFormatType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsObjectResponseFormatTypeJsonObject:
		return true
	}
	return false
}

type PromptNewParamsPromptDataOptionsParamsObjectToolChoice struct {
	Type     param.Field[PromptNewParamsPromptDataOptionsParamsObjectToolChoiceType] `json:"type,required"`
	Function param.Field[interface{}]                                                `json:"function"`
}

func (r PromptNewParamsPromptDataOptionsParamsObjectToolChoice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsObjectToolChoice) ImplementsPromptNewParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
}

// Satisfied by [PromptNewParamsPromptDataOptionsParamsObjectToolChoiceString],
// [PromptNewParamsPromptDataOptionsParamsObjectToolChoiceString],
// [PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObject],
// [PromptNewParamsPromptDataOptionsParamsObjectToolChoice].
type PromptNewParamsPromptDataOptionsParamsObjectToolChoiceUnion interface {
	ImplementsPromptNewParamsPromptDataOptionsParamsObjectToolChoiceUnion()
}

type PromptNewParamsPromptDataOptionsParamsObjectToolChoiceString string

const (
	PromptNewParamsPromptDataOptionsParamsObjectToolChoiceStringAuto PromptNewParamsPromptDataOptionsParamsObjectToolChoiceString = "auto"
)

func (r PromptNewParamsPromptDataOptionsParamsObjectToolChoiceString) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsObjectToolChoiceStringAuto:
		return true
	}
	return false
}

type PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObject struct {
	Function param.Field[PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction] `json:"function,required"`
	Type     param.Field[PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObjectType]     `json:"type,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObject) ImplementsPromptNewParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
}

type PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObjectType string

const (
	PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObjectType = "function"
)

func (r PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObjectType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction:
		return true
	}
	return false
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
	Type                  param.Field[PromptNewParamsPromptDataPromptType] `json:"type"`
	Content               param.Field[string]                              `json:"content"`
	Messages              param.Field[interface{}]                         `json:"messages,required"`
	Tools                 param.Field[string]                              `json:"tools"`
	ReservedOnlyAllowNull param.Field[interface{}]                         `json:"__reserved_only_allow_null,required"`
}

func (r PromptNewParamsPromptDataPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPrompt) implementsPromptNewParamsPromptDataPromptUnion() {}

// Satisfied by [PromptNewParamsPromptDataPromptObject],
// [PromptNewParamsPromptDataPromptObject],
// [PromptNewParamsPromptDataPromptObject], [PromptNewParamsPromptDataPrompt].
type PromptNewParamsPromptDataPromptUnion interface {
	implementsPromptNewParamsPromptDataPromptUnion()
}

type PromptNewParamsPromptDataPromptObject struct {
	Content param.Field[string]                                    `json:"content,required"`
	Type    param.Field[PromptNewParamsPromptDataPromptObjectType] `json:"type,required"`
}

func (r PromptNewParamsPromptDataPromptObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptObject) implementsPromptNewParamsPromptDataPromptUnion() {}

type PromptNewParamsPromptDataPromptObjectType string

const (
	PromptNewParamsPromptDataPromptObjectTypeCompletion PromptNewParamsPromptDataPromptObjectType = "completion"
)

func (r PromptNewParamsPromptDataPromptObjectType) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptObjectTypeCompletion:
		return true
	}
	return false
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

// Satisfied by [PromptUpdateParamsPromptDataOptionsParamsObject],
// [PromptUpdateParamsPromptDataOptionsParamsObject],
// [PromptUpdateParamsPromptDataOptionsParamsObject],
// [PromptUpdateParamsPromptDataOptionsParamsObject],
// [PromptUpdateParamsPromptDataOptionsParams].
type PromptUpdateParamsPromptDataOptionsParamsUnion interface {
	implementsPromptUpdateParamsPromptDataOptionsParamsUnion()
}

type PromptUpdateParamsPromptDataOptionsParamsObject struct {
	FrequencyPenalty param.Field[float64]                                                          `json:"frequency_penalty"`
	FunctionCall     param.Field[PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                          `json:"max_tokens"`
	N                param.Field[float64]                                                          `json:"n"`
	PresencePenalty  param.Field[float64]                                                          `json:"presence_penalty"`
	ResponseFormat   param.Field[PromptUpdateParamsPromptDataOptionsParamsObjectResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                         `json:"stop"`
	Temperature      param.Field[float64]                                                          `json:"temperature"`
	ToolChoice       param.Field[PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                          `json:"top_p"`
	UseCache         param.Field[bool]                                                             `json:"use_cache"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsObject) implementsPromptUpdateParamsPromptDataOptionsParamsUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCall struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCall) ImplementsPromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
}

// Satisfied by
// [PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallString],
// [PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallString],
// [PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallObject],
// [PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCall].
type PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallUnion interface {
	ImplementsPromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallUnion()
}

type PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallString string

const (
	PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallStringAuto PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallString = "auto"
)

func (r PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallString) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallStringAuto:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallObject) ImplementsPromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsObjectResponseFormat struct {
	Type param.Field[PromptUpdateParamsPromptDataOptionsParamsObjectResponseFormatType] `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsObjectResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataOptionsParamsObjectResponseFormatType string

const (
	PromptUpdateParamsPromptDataOptionsParamsObjectResponseFormatTypeJsonObject PromptUpdateParamsPromptDataOptionsParamsObjectResponseFormatType = "json_object"
)

func (r PromptUpdateParamsPromptDataOptionsParamsObjectResponseFormatType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsObjectResponseFormatTypeJsonObject:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataOptionsParamsObjectToolChoice struct {
	Type     param.Field[PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceType] `json:"type,required"`
	Function param.Field[interface{}]                                                   `json:"function"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsObjectToolChoice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsObjectToolChoice) ImplementsPromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
}

// Satisfied by [PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceString],
// [PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceString],
// [PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObject],
// [PromptUpdateParamsPromptDataOptionsParamsObjectToolChoice].
type PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceUnion interface {
	ImplementsPromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceUnion()
}

type PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceString string

const (
	PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceStringAuto PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceString = "auto"
)

func (r PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceString) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceStringAuto:
		return true
	}
	return false
}

type PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObject struct {
	Function param.Field[PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction] `json:"function,required"`
	Type     param.Field[PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectType]     `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObject) ImplementsPromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectType string

const (
	PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectType = "function"
)

func (r PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction:
		return true
	}
	return false
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
	Type                  param.Field[PromptUpdateParamsPromptDataPromptType] `json:"type"`
	Content               param.Field[string]                                 `json:"content"`
	Messages              param.Field[interface{}]                            `json:"messages,required"`
	Tools                 param.Field[string]                                 `json:"tools"`
	ReservedOnlyAllowNull param.Field[interface{}]                            `json:"__reserved_only_allow_null,required"`
}

func (r PromptUpdateParamsPromptDataPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPrompt) implementsPromptUpdateParamsPromptDataPromptUnion() {}

// Satisfied by [PromptUpdateParamsPromptDataPromptObject],
// [PromptUpdateParamsPromptDataPromptObject],
// [PromptUpdateParamsPromptDataPromptObject],
// [PromptUpdateParamsPromptDataPrompt].
type PromptUpdateParamsPromptDataPromptUnion interface {
	implementsPromptUpdateParamsPromptDataPromptUnion()
}

type PromptUpdateParamsPromptDataPromptObject struct {
	Content param.Field[string]                                       `json:"content,required"`
	Type    param.Field[PromptUpdateParamsPromptDataPromptObjectType] `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataPromptObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptObject) implementsPromptUpdateParamsPromptDataPromptUnion() {
}

type PromptUpdateParamsPromptDataPromptObjectType string

const (
	PromptUpdateParamsPromptDataPromptObjectTypeCompletion PromptUpdateParamsPromptDataPromptObjectType = "completion"
)

func (r PromptUpdateParamsPromptDataPromptObjectType) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptObjectTypeCompletion:
		return true
	}
	return false
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
	// Retrieve a snapshot of events from a past time
	//
	// The version id is essentially a filter on the latest event transaction id. You
	// can use the `max_xact_id` returned by a past fetch as the version to reproduce
	// that exact fetch.
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

type PromptFeedbackParams struct {
	// A list of prompt feedback items
	Feedback param.Field[[]PromptFeedbackParamsFeedback] `json:"feedback,required"`
}

func (r PromptFeedbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptFeedbackParamsFeedback struct {
	// The id of the prompt event to log feedback for. This is the row `id` returned by
	// `POST /v1/prompt/{prompt_id}/insert`
	ID param.Field[string] `json:"id,required"`
	// An optional comment string to log about the prompt event
	Comment param.Field[string] `json:"comment"`
	// A dictionary with additional data about the feedback. If you have a `user_id`,
	// you can log it here and access it in the Braintrust UI.
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// The source of the feedback. Must be one of "external" (default), "app", or "api"
	Source param.Field[PromptFeedbackParamsFeedbackSource] `json:"source"`
}

func (r PromptFeedbackParamsFeedback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The source of the feedback. Must be one of "external" (default), "app", or "api"
type PromptFeedbackParamsFeedbackSource string

const (
	PromptFeedbackParamsFeedbackSourceApp      PromptFeedbackParamsFeedbackSource = "app"
	PromptFeedbackParamsFeedbackSourceAPI      PromptFeedbackParamsFeedbackSource = "api"
	PromptFeedbackParamsFeedbackSourceExternal PromptFeedbackParamsFeedbackSource = "external"
)

func (r PromptFeedbackParamsFeedbackSource) IsKnown() bool {
	switch r {
	case PromptFeedbackParamsFeedbackSourceApp, PromptFeedbackParamsFeedbackSourceAPI, PromptFeedbackParamsFeedbackSourceExternal:
		return true
	}
	return false
}

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

// Satisfied by [PromptReplaceParamsPromptDataOptionsParamsObject],
// [PromptReplaceParamsPromptDataOptionsParamsObject],
// [PromptReplaceParamsPromptDataOptionsParamsObject],
// [PromptReplaceParamsPromptDataOptionsParamsObject],
// [PromptReplaceParamsPromptDataOptionsParams].
type PromptReplaceParamsPromptDataOptionsParamsUnion interface {
	implementsPromptReplaceParamsPromptDataOptionsParamsUnion()
}

type PromptReplaceParamsPromptDataOptionsParamsObject struct {
	FrequencyPenalty param.Field[float64]                                                           `json:"frequency_penalty"`
	FunctionCall     param.Field[PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallUnion] `json:"function_call"`
	MaxTokens        param.Field[float64]                                                           `json:"max_tokens"`
	N                param.Field[float64]                                                           `json:"n"`
	PresencePenalty  param.Field[float64]                                                           `json:"presence_penalty"`
	ResponseFormat   param.Field[PromptReplaceParamsPromptDataOptionsParamsObjectResponseFormat]    `json:"response_format"`
	Stop             param.Field[[]string]                                                          `json:"stop"`
	Temperature      param.Field[float64]                                                           `json:"temperature"`
	ToolChoice       param.Field[PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceUnion]   `json:"tool_choice"`
	TopP             param.Field[float64]                                                           `json:"top_p"`
	UseCache         param.Field[bool]                                                              `json:"use_cache"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsObject) implementsPromptReplaceParamsPromptDataOptionsParamsUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCall struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCall) ImplementsPromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
}

// Satisfied by
// [PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallString],
// [PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallString],
// [PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallObject],
// [PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCall].
type PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallUnion interface {
	ImplementsPromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallUnion()
}

type PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallString string

const (
	PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallStringAuto PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallString = "auto"
)

func (r PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallString) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallStringAuto:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallObject) ImplementsPromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsObjectResponseFormat struct {
	Type param.Field[PromptReplaceParamsPromptDataOptionsParamsObjectResponseFormatType] `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsObjectResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataOptionsParamsObjectResponseFormatType string

const (
	PromptReplaceParamsPromptDataOptionsParamsObjectResponseFormatTypeJsonObject PromptReplaceParamsPromptDataOptionsParamsObjectResponseFormatType = "json_object"
)

func (r PromptReplaceParamsPromptDataOptionsParamsObjectResponseFormatType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsObjectResponseFormatTypeJsonObject:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataOptionsParamsObjectToolChoice struct {
	Type     param.Field[PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceType] `json:"type,required"`
	Function param.Field[interface{}]                                                    `json:"function"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsObjectToolChoice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsObjectToolChoice) ImplementsPromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
}

// Satisfied by [PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceString],
// [PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceString],
// [PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObject],
// [PromptReplaceParamsPromptDataOptionsParamsObjectToolChoice].
type PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceUnion interface {
	ImplementsPromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceUnion()
}

type PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceString string

const (
	PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceStringAuto PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceString = "auto"
)

func (r PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceString) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceStringAuto:
		return true
	}
	return false
}

type PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObject struct {
	Function param.Field[PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction] `json:"function,required"`
	Type     param.Field[PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectType]     `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObject) ImplementsPromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectType string

const (
	PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectType = "function"
)

func (r PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectTypeFunction:
		return true
	}
	return false
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
	Type                  param.Field[PromptReplaceParamsPromptDataPromptType] `json:"type"`
	Content               param.Field[string]                                  `json:"content"`
	Messages              param.Field[interface{}]                             `json:"messages,required"`
	Tools                 param.Field[string]                                  `json:"tools"`
	ReservedOnlyAllowNull param.Field[interface{}]                             `json:"__reserved_only_allow_null,required"`
}

func (r PromptReplaceParamsPromptDataPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPrompt) implementsPromptReplaceParamsPromptDataPromptUnion() {}

// Satisfied by [PromptReplaceParamsPromptDataPromptObject],
// [PromptReplaceParamsPromptDataPromptObject],
// [PromptReplaceParamsPromptDataPromptObject],
// [PromptReplaceParamsPromptDataPrompt].
type PromptReplaceParamsPromptDataPromptUnion interface {
	implementsPromptReplaceParamsPromptDataPromptUnion()
}

type PromptReplaceParamsPromptDataPromptObject struct {
	Content param.Field[string]                                        `json:"content,required"`
	Type    param.Field[PromptReplaceParamsPromptDataPromptObjectType] `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataPromptObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptObject) implementsPromptReplaceParamsPromptDataPromptUnion() {
}

type PromptReplaceParamsPromptDataPromptObjectType string

const (
	PromptReplaceParamsPromptDataPromptObjectTypeCompletion PromptReplaceParamsPromptDataPromptObjectType = "completion"
)

func (r PromptReplaceParamsPromptDataPromptObjectType) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptObjectTypeCompletion:
		return true
	}
	return false
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
