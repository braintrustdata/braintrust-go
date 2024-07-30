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

// Log feedback for a set of prompt events
func (r *PromptService) Feedback(ctx context.Context, promptID string, body PromptFeedbackParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if promptID == "" {
		err = errors.New("missing required prompt_id parameter")
		return
	}
	path := fmt.Sprintf("v1/prompt/%s/feedback", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
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
	// [PromptPromptDataOptionsParamsObjectResponseFormat].
	ResponseFormat interface{} `json:"response_format,required"`
	// This field can have the runtime type of
	// [PromptPromptDataOptionsParamsObjectToolChoiceUnion].
	ToolChoice interface{} `json:"tool_choice,required"`
	// This field can have the runtime type of
	// [PromptPromptDataOptionsParamsObjectFunctionCallUnion].
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
// Possible runtime types of the union are [PromptPromptDataOptionsParamsObject],
// [PromptPromptDataOptionsParamsObject], [PromptPromptDataOptionsParamsObject],
// [PromptPromptDataOptionsParamsObject], [PromptPromptDataOptionsParamsObject].
func (r PromptPromptDataOptionsParams) AsUnion() PromptPromptDataOptionsParamsUnion {
	return r.union
}

// Union satisfied by [PromptPromptDataOptionsParamsObject],
// [PromptPromptDataOptionsParamsObject], [PromptPromptDataOptionsParamsObject],
// [PromptPromptDataOptionsParamsObject] or [PromptPromptDataOptionsParamsObject].
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
	implementsPromptPromptDataOptionsParamsObjectFunctionCallUnion()
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

func (r PromptPromptDataOptionsParamsObjectFunctionCallString) implementsPromptPromptDataOptionsParamsObjectFunctionCallUnion() {
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

func (r PromptPromptDataOptionsParamsObjectFunctionCallObject) implementsPromptPromptDataOptionsParamsObjectFunctionCallUnion() {
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
	implementsPromptPromptDataOptionsParamsObjectToolChoiceUnion()
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

func (r PromptPromptDataOptionsParamsObjectToolChoiceString) implementsPromptPromptDataOptionsParamsObjectToolChoiceUnion() {
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

func (r PromptPromptDataOptionsParamsObjectToolChoiceObject) implementsPromptPromptDataOptionsParamsObjectToolChoiceUnion() {
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
// [PromptPromptDataPromptChat], [PromptPromptDataPromptNullVariant].
func (r PromptPromptDataPrompt) AsUnion() PromptPromptDataPromptUnion {
	return r.union
}

// Union satisfied by [PromptPromptDataPromptCompletion],
// [PromptPromptDataPromptChat] or [PromptPromptDataPromptNullVariant].
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
			Type:       reflect.TypeOf(PromptPromptDataPromptNullVariant{}),
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
	// [PromptPromptDataPromptChatMessagesObjectContentUnion].
	Content interface{}                            `json:"content,required"`
	Role    PromptPromptDataPromptChatMessagesRole `json:"role,required"`
	Name    string                                 `json:"name"`
	// This field can have the runtime type of
	// [PromptPromptDataPromptChatMessagesObjectFunctionCall].
	FunctionCall interface{} `json:"function_call,required"`
	// This field can have the runtime type of
	// [[]PromptPromptDataPromptChatMessagesObjectToolCall].
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
// [PromptPromptDataPromptChatMessagesObject],
// [PromptPromptDataPromptChatMessagesObject],
// [PromptPromptDataPromptChatMessagesObject],
// [PromptPromptDataPromptChatMessagesObject],
// [PromptPromptDataPromptChatMessagesObject],
// [PromptPromptDataPromptChatMessagesObject].
func (r PromptPromptDataPromptChatMessage) AsUnion() PromptPromptDataPromptChatMessagesUnion {
	return r.union
}

// Union satisfied by [PromptPromptDataPromptChatMessagesObject],
// [PromptPromptDataPromptChatMessagesObject],
// [PromptPromptDataPromptChatMessagesObject],
// [PromptPromptDataPromptChatMessagesObject],
// [PromptPromptDataPromptChatMessagesObject] or
// [PromptPromptDataPromptChatMessagesObject].
type PromptPromptDataPromptChatMessagesUnion interface {
	implementsPromptPromptDataPromptChatMessage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptPromptDataPromptChatMessagesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptPromptDataPromptChatMessagesObject{}),
		},
	)
}

type PromptPromptDataPromptChatMessagesObject struct {
	Role    PromptPromptDataPromptChatMessagesObjectRole `json:"role,required"`
	Content string                                       `json:"content"`
	Name    string                                       `json:"name"`
	JSON    promptPromptDataPromptChatMessagesObjectJSON `json:"-"`
}

// promptPromptDataPromptChatMessagesObjectJSON contains the JSON metadata for the
// struct [PromptPromptDataPromptChatMessagesObject]
type promptPromptDataPromptChatMessagesObjectJSON struct {
	Role        apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptChatMessagesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptChatMessagesObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptChatMessagesObject) implementsPromptPromptDataPromptChatMessage() {}

type PromptPromptDataPromptChatMessagesObjectRole string

const (
	PromptPromptDataPromptChatMessagesObjectRoleSystem PromptPromptDataPromptChatMessagesObjectRole = "system"
)

func (r PromptPromptDataPromptChatMessagesObjectRole) IsKnown() bool {
	switch r {
	case PromptPromptDataPromptChatMessagesObjectRoleSystem:
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

type PromptPromptDataPromptNullVariant struct {
	JSON promptPromptDataPromptNullVariantJSON `json:"-"`
}

// promptPromptDataPromptNullVariantJSON contains the JSON metadata for the struct
// [PromptPromptDataPromptNullVariant]
type promptPromptDataPromptNullVariantJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptPromptDataPromptNullVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptPromptDataPromptNullVariantJSON) RawJSON() string {
	return r.raw
}

func (r PromptPromptDataPromptNullVariant) implementsPromptPromptDataPrompt() {}

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

// Satisfied by [PromptNewParamsPromptDataOptionsParamsObjectFunctionCallString],
// [PromptNewParamsPromptDataOptionsParamsObjectFunctionCallString],
// [PromptNewParamsPromptDataOptionsParamsObjectFunctionCallObject].
type PromptNewParamsPromptDataOptionsParamsObjectFunctionCallUnion interface {
	implementsPromptNewParamsPromptDataOptionsParamsObjectFunctionCallUnion()
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

func (r PromptNewParamsPromptDataOptionsParamsObjectFunctionCallString) implementsPromptNewParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
}

type PromptNewParamsPromptDataOptionsParamsObjectFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsObjectFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsObjectFunctionCallObject) implementsPromptNewParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
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

// Satisfied by [PromptNewParamsPromptDataOptionsParamsObjectToolChoiceString],
// [PromptNewParamsPromptDataOptionsParamsObjectToolChoiceString],
// [PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObject].
type PromptNewParamsPromptDataOptionsParamsObjectToolChoiceUnion interface {
	implementsPromptNewParamsPromptDataOptionsParamsObjectToolChoiceUnion()
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

func (r PromptNewParamsPromptDataOptionsParamsObjectToolChoiceString) implementsPromptNewParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
}

type PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObject struct {
	Function param.Field[PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction] `json:"function,required"`
	Type     param.Field[PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObjectType]     `json:"type,required"`
}

func (r PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataOptionsParamsObjectToolChoiceObject) implementsPromptNewParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
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
// [PromptNewParamsPromptDataPromptNullVariant], [PromptNewParamsPromptDataPrompt].
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

// Satisfied by [PromptNewParamsPromptDataPromptChatMessagesObject],
// [PromptNewParamsPromptDataPromptChatMessagesObject],
// [PromptNewParamsPromptDataPromptChatMessagesObject],
// [PromptNewParamsPromptDataPromptChatMessagesObject],
// [PromptNewParamsPromptDataPromptChatMessagesObject],
// [PromptNewParamsPromptDataPromptChatMessagesObject],
// [PromptNewParamsPromptDataPromptChatMessage].
type PromptNewParamsPromptDataPromptChatMessageUnion interface {
	implementsPromptNewParamsPromptDataPromptChatMessageUnion()
}

type PromptNewParamsPromptDataPromptChatMessagesObject struct {
	Role    param.Field[PromptNewParamsPromptDataPromptChatMessagesObjectRole] `json:"role,required"`
	Content param.Field[string]                                                `json:"content"`
	Name    param.Field[string]                                                `json:"name"`
}

func (r PromptNewParamsPromptDataPromptChatMessagesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptChatMessagesObject) implementsPromptNewParamsPromptDataPromptChatMessageUnion() {
}

type PromptNewParamsPromptDataPromptChatMessagesObjectRole string

const (
	PromptNewParamsPromptDataPromptChatMessagesObjectRoleSystem PromptNewParamsPromptDataPromptChatMessagesObjectRole = "system"
)

func (r PromptNewParamsPromptDataPromptChatMessagesObjectRole) IsKnown() bool {
	switch r {
	case PromptNewParamsPromptDataPromptChatMessagesObjectRoleSystem:
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

type PromptNewParamsPromptDataPromptNullVariant struct {
}

func (r PromptNewParamsPromptDataPromptNullVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptNewParamsPromptDataPromptNullVariant) implementsPromptNewParamsPromptDataPromptUnion() {
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

// Satisfied by
// [PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallString],
// [PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallString],
// [PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallObject].
type PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallUnion interface {
	implementsPromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallUnion()
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

func (r PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallString) implementsPromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallObject) implementsPromptUpdateParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
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

// Satisfied by [PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceString],
// [PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceString],
// [PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObject].
type PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceUnion interface {
	implementsPromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceUnion()
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

func (r PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceString) implementsPromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
}

type PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObject struct {
	Function param.Field[PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction] `json:"function,required"`
	Type     param.Field[PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObjectType]     `json:"type,required"`
}

func (r PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceObject) implementsPromptUpdateParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
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
// [PromptUpdateParamsPromptDataPromptNullVariant],
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

// Satisfied by [PromptUpdateParamsPromptDataPromptChatMessagesObject],
// [PromptUpdateParamsPromptDataPromptChatMessagesObject],
// [PromptUpdateParamsPromptDataPromptChatMessagesObject],
// [PromptUpdateParamsPromptDataPromptChatMessagesObject],
// [PromptUpdateParamsPromptDataPromptChatMessagesObject],
// [PromptUpdateParamsPromptDataPromptChatMessagesObject],
// [PromptUpdateParamsPromptDataPromptChatMessage].
type PromptUpdateParamsPromptDataPromptChatMessageUnion interface {
	implementsPromptUpdateParamsPromptDataPromptChatMessageUnion()
}

type PromptUpdateParamsPromptDataPromptChatMessagesObject struct {
	Role    param.Field[PromptUpdateParamsPromptDataPromptChatMessagesObjectRole] `json:"role,required"`
	Content param.Field[string]                                                   `json:"content"`
	Name    param.Field[string]                                                   `json:"name"`
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptChatMessagesObject) implementsPromptUpdateParamsPromptDataPromptChatMessageUnion() {
}

type PromptUpdateParamsPromptDataPromptChatMessagesObjectRole string

const (
	PromptUpdateParamsPromptDataPromptChatMessagesObjectRoleSystem PromptUpdateParamsPromptDataPromptChatMessagesObjectRole = "system"
)

func (r PromptUpdateParamsPromptDataPromptChatMessagesObjectRole) IsKnown() bool {
	switch r {
	case PromptUpdateParamsPromptDataPromptChatMessagesObjectRoleSystem:
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

type PromptUpdateParamsPromptDataPromptNullVariant struct {
}

func (r PromptUpdateParamsPromptDataPromptNullVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptUpdateParamsPromptDataPromptNullVariant) implementsPromptUpdateParamsPromptDataPromptUnion() {
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

// Satisfied by
// [PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallString],
// [PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallString],
// [PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallObject].
type PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallUnion interface {
	implementsPromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallUnion()
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

func (r PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallString) implementsPromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallObject) implementsPromptReplaceParamsPromptDataOptionsParamsObjectFunctionCallUnion() {
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

// Satisfied by [PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceString],
// [PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceString],
// [PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObject].
type PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceUnion interface {
	implementsPromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceUnion()
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

func (r PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceString) implementsPromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
}

type PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObject struct {
	Function param.Field[PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectFunction] `json:"function,required"`
	Type     param.Field[PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObjectType]     `json:"type,required"`
}

func (r PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceObject) implementsPromptReplaceParamsPromptDataOptionsParamsObjectToolChoiceUnion() {
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
// [PromptReplaceParamsPromptDataPromptNullVariant],
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

// Satisfied by [PromptReplaceParamsPromptDataPromptChatMessagesObject],
// [PromptReplaceParamsPromptDataPromptChatMessagesObject],
// [PromptReplaceParamsPromptDataPromptChatMessagesObject],
// [PromptReplaceParamsPromptDataPromptChatMessagesObject],
// [PromptReplaceParamsPromptDataPromptChatMessagesObject],
// [PromptReplaceParamsPromptDataPromptChatMessagesObject],
// [PromptReplaceParamsPromptDataPromptChatMessage].
type PromptReplaceParamsPromptDataPromptChatMessageUnion interface {
	implementsPromptReplaceParamsPromptDataPromptChatMessageUnion()
}

type PromptReplaceParamsPromptDataPromptChatMessagesObject struct {
	Role    param.Field[PromptReplaceParamsPromptDataPromptChatMessagesObjectRole] `json:"role,required"`
	Content param.Field[string]                                                    `json:"content"`
	Name    param.Field[string]                                                    `json:"name"`
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptChatMessagesObject) implementsPromptReplaceParamsPromptDataPromptChatMessageUnion() {
}

type PromptReplaceParamsPromptDataPromptChatMessagesObjectRole string

const (
	PromptReplaceParamsPromptDataPromptChatMessagesObjectRoleSystem PromptReplaceParamsPromptDataPromptChatMessagesObjectRole = "system"
)

func (r PromptReplaceParamsPromptDataPromptChatMessagesObjectRole) IsKnown() bool {
	switch r {
	case PromptReplaceParamsPromptDataPromptChatMessagesObjectRoleSystem:
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

type PromptReplaceParamsPromptDataPromptNullVariant struct {
}

func (r PromptReplaceParamsPromptDataPromptNullVariant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptReplaceParamsPromptDataPromptNullVariant) implementsPromptReplaceParamsPromptDataPromptUnion() {
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
