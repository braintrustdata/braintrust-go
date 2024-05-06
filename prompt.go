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
// with the braintrust API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPromptService] method instead.
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
func (r *PromptService) New(ctx context.Context, body PromptNewParams, opts ...option.RequestOption) (res *PromptNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/prompt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a prompt object by its id
func (r *PromptService) Get(ctx context.Context, promptID string, opts ...option.RequestOption) (res *PromptGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/prompt/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update a prompt object. Specify the fields to update in the payload.
// Any object-type fields will be deep-merged with existing content. Currently we
// do not support removing fields or setting them to null.
func (r *PromptService) Update(ctx context.Context, promptID string, body PromptUpdateParams, opts ...option.RequestOption) (res *PromptUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/prompt/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all prompts. The prompts are sorted by creation date, with the most
// recently-created prompts coming first
func (r *PromptService) List(ctx context.Context, query PromptListParams, opts ...option.RequestOption) (res *pagination.ListObjects[PromptListResponse], err error) {
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
func (r *PromptService) ListAutoPaging(ctx context.Context, query PromptListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[PromptListResponse] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a prompt object by its id
func (r *PromptService) Delete(ctx context.Context, promptID string, opts ...option.RequestOption) (res *PromptDeleteResponse, err error) {
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
func (r *PromptService) Replace(ctx context.Context, body PromptReplaceParams, opts ...option.RequestOption) (res *PromptReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/prompt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type PromptNewResponse struct {
	// Unique identifier for the prompt
	ID string `json:"id,required" format:"uuid"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the prompt (see the `version`
	// parameter)
	XactID string `json:"_xact_id,required"`
	// A literal 'p' which identifies the object as a project prompt
	LogID PromptNewResponseLogID `json:"log_id,required"`
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
	PromptData PromptNewResponsePromptData `json:"prompt_data,nullable"`
	// A list of tags for the prompt
	Tags []string              `json:"tags,nullable"`
	JSON promptNewResponseJSON `json:"-"`
}

// promptNewResponseJSON contains the JSON metadata for the struct
// [PromptNewResponse]
type promptNewResponseJSON struct {
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

func (r *PromptNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptNewResponseJSON) RawJSON() string {
	return r.raw
}

// A literal 'p' which identifies the object as a project prompt
type PromptNewResponseLogID string

const (
	PromptNewResponseLogIDP PromptNewResponseLogID = "p"
)

func (r PromptNewResponseLogID) IsKnown() bool {
	switch r {
	case PromptNewResponseLogIDP:
		return true
	}
	return false
}

// The prompt, model, and its parameters
type PromptNewResponsePromptData struct {
	Options PromptNewResponsePromptDataOptions `json:"options,nullable"`
	Origin  PromptNewResponsePromptDataOrigin  `json:"origin,nullable"`
	Prompt  PromptNewResponsePromptDataPrompt  `json:"prompt"`
	JSON    promptNewResponsePromptDataJSON    `json:"-"`
}

// promptNewResponsePromptDataJSON contains the JSON metadata for the struct
// [PromptNewResponsePromptData]
type promptNewResponsePromptDataJSON struct {
	Options     apijson.Field
	Origin      apijson.Field
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptNewResponsePromptData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptNewResponsePromptDataJSON) RawJSON() string {
	return r.raw
}

type PromptNewResponsePromptDataOptions struct {
	Model    string                                   `json:"model"`
	Params   PromptNewResponsePromptDataOptionsParams `json:"params"`
	Position string                                   `json:"position"`
	JSON     promptNewResponsePromptDataOptionsJSON   `json:"-"`
}

// promptNewResponsePromptDataOptionsJSON contains the JSON metadata for the struct
// [PromptNewResponsePromptDataOptions]
type promptNewResponsePromptDataOptionsJSON struct {
	Model       apijson.Field
	Params      apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptNewResponsePromptDataOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptNewResponsePromptDataOptionsJSON) RawJSON() string {
	return r.raw
}

type PromptNewResponsePromptDataOptionsParams struct {
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
	MaxTokensToSample float64                                      `json:"max_tokens_to_sample"`
	MaxOutputTokens   float64                                      `json:"maxOutputTokens"`
	TopP              float64                                      `json:"topP"`
	TopK              float64                                      `json:"topK"`
	JSON              promptNewResponsePromptDataOptionsParamsJSON `json:"-"`
	union             PromptNewResponsePromptDataOptionsParamsUnion
}

// promptNewResponsePromptDataOptionsParamsJSON contains the JSON metadata for the
// struct [PromptNewResponsePromptDataOptionsParams]
type promptNewResponsePromptDataOptionsParamsJSON struct {
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

func (r promptNewResponsePromptDataOptionsParamsJSON) RawJSON() string {
	return r.raw
}

func (r *PromptNewResponsePromptDataOptionsParams) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PromptNewResponsePromptDataOptionsParams) AsUnion() PromptNewResponsePromptDataOptionsParamsUnion {
	return r.union
}

// Union satisfied by [PromptNewResponsePromptDataOptionsParamsObject],
// [PromptNewResponsePromptDataOptionsParamsObject],
// [PromptNewResponsePromptDataOptionsParamsObject] or
// [PromptNewResponsePromptDataOptionsParamsObject].
type PromptNewResponsePromptDataOptionsParamsUnion interface {
	implementsPromptNewResponsePromptDataOptionsParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptNewResponsePromptDataOptionsParamsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptNewResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptNewResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptNewResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptNewResponsePromptDataOptionsParamsObject{}),
		},
	)
}

type PromptNewResponsePromptDataOptionsParamsObject struct {
	FrequencyPenalty float64                                                         `json:"frequency_penalty"`
	FunctionCall     PromptNewResponsePromptDataOptionsParamsObjectFunctionCallUnion `json:"function_call"`
	MaxTokens        float64                                                         `json:"max_tokens"`
	N                float64                                                         `json:"n"`
	PresencePenalty  float64                                                         `json:"presence_penalty"`
	ResponseFormat   PromptNewResponsePromptDataOptionsParamsObjectResponseFormat    `json:"response_format,nullable"`
	Stop             []string                                                        `json:"stop"`
	Temperature      float64                                                         `json:"temperature"`
	ToolChoice       PromptNewResponsePromptDataOptionsParamsObjectToolChoiceUnion   `json:"tool_choice"`
	TopP             float64                                                         `json:"top_p"`
	UseCache         bool                                                            `json:"use_cache"`
	JSON             promptNewResponsePromptDataOptionsParamsObjectJSON              `json:"-"`
}

// promptNewResponsePromptDataOptionsParamsObjectJSON contains the JSON metadata
// for the struct [PromptNewResponsePromptDataOptionsParamsObject]
type promptNewResponsePromptDataOptionsParamsObjectJSON struct {
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

func (r *PromptNewResponsePromptDataOptionsParamsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptNewResponsePromptDataOptionsParamsObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptNewResponsePromptDataOptionsParamsObject) implementsPromptNewResponsePromptDataOptionsParams() {
}

// Union satisfied by
// [PromptNewResponsePromptDataOptionsParamsObjectFunctionCallString],
// [PromptNewResponsePromptDataOptionsParamsObjectFunctionCallString] or
// [PromptNewResponsePromptDataOptionsParamsObjectFunctionCallObject].
type PromptNewResponsePromptDataOptionsParamsObjectFunctionCallUnion interface {
	ImplementsPromptNewResponsePromptDataOptionsParamsObjectFunctionCallUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptNewResponsePromptDataOptionsParamsObjectFunctionCallUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptNewResponsePromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptNewResponsePromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptNewResponsePromptDataOptionsParamsObjectFunctionCallObject{}),
		},
	)
}

type PromptNewResponsePromptDataOptionsParamsObjectFunctionCallString string

const (
	PromptNewResponsePromptDataOptionsParamsObjectFunctionCallStringAuto PromptNewResponsePromptDataOptionsParamsObjectFunctionCallString = "auto"
)

func (r PromptNewResponsePromptDataOptionsParamsObjectFunctionCallString) IsKnown() bool {
	switch r {
	case PromptNewResponsePromptDataOptionsParamsObjectFunctionCallStringAuto:
		return true
	}
	return false
}

type PromptNewResponsePromptDataOptionsParamsObjectFunctionCallObject struct {
	Name string                                                               `json:"name,required"`
	JSON promptNewResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON `json:"-"`
}

// promptNewResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON contains
// the JSON metadata for the struct
// [PromptNewResponsePromptDataOptionsParamsObjectFunctionCallObject]
type promptNewResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptNewResponsePromptDataOptionsParamsObjectFunctionCallObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptNewResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptNewResponsePromptDataOptionsParamsObjectFunctionCallObject) ImplementsPromptNewResponsePromptDataOptionsParamsObjectFunctionCallUnion() {
}

type PromptNewResponsePromptDataOptionsParamsObjectResponseFormat struct {
	Type PromptNewResponsePromptDataOptionsParamsObjectResponseFormatType `json:"type,required"`
	JSON promptNewResponsePromptDataOptionsParamsObjectResponseFormatJSON `json:"-"`
}

// promptNewResponsePromptDataOptionsParamsObjectResponseFormatJSON contains the
// JSON metadata for the struct
// [PromptNewResponsePromptDataOptionsParamsObjectResponseFormat]
type promptNewResponsePromptDataOptionsParamsObjectResponseFormatJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptNewResponsePromptDataOptionsParamsObjectResponseFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptNewResponsePromptDataOptionsParamsObjectResponseFormatJSON) RawJSON() string {
	return r.raw
}

type PromptNewResponsePromptDataOptionsParamsObjectResponseFormatType string

const (
	PromptNewResponsePromptDataOptionsParamsObjectResponseFormatTypeJsonObject PromptNewResponsePromptDataOptionsParamsObjectResponseFormatType = "json_object"
)

func (r PromptNewResponsePromptDataOptionsParamsObjectResponseFormatType) IsKnown() bool {
	switch r {
	case PromptNewResponsePromptDataOptionsParamsObjectResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Union satisfied by
// [PromptNewResponsePromptDataOptionsParamsObjectToolChoiceString],
// [PromptNewResponsePromptDataOptionsParamsObjectToolChoiceString] or
// [PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObject].
type PromptNewResponsePromptDataOptionsParamsObjectToolChoiceUnion interface {
	ImplementsPromptNewResponsePromptDataOptionsParamsObjectToolChoiceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptNewResponsePromptDataOptionsParamsObjectToolChoiceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptNewResponsePromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptNewResponsePromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObject{}),
		},
	)
}

type PromptNewResponsePromptDataOptionsParamsObjectToolChoiceString string

const (
	PromptNewResponsePromptDataOptionsParamsObjectToolChoiceStringAuto PromptNewResponsePromptDataOptionsParamsObjectToolChoiceString = "auto"
)

func (r PromptNewResponsePromptDataOptionsParamsObjectToolChoiceString) IsKnown() bool {
	switch r {
	case PromptNewResponsePromptDataOptionsParamsObjectToolChoiceStringAuto:
		return true
	}
	return false
}

type PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObject struct {
	Function PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction `json:"function,required"`
	Type     PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectType     `json:"type,required"`
	JSON     promptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON     `json:"-"`
}

// promptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON contains the
// JSON metadata for the struct
// [PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObject]
type promptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON struct {
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObject) ImplementsPromptNewResponsePromptDataOptionsParamsObjectToolChoiceUnion() {
}

type PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction struct {
	Name string                                                                     `json:"name,required"`
	JSON promptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON `json:"-"`
}

// promptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON
// contains the JSON metadata for the struct
// [PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction]
type promptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON) RawJSON() string {
	return r.raw
}

type PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectType string

const (
	PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectTypeFunction PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectType = "function"
)

func (r PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectType) IsKnown() bool {
	switch r {
	case PromptNewResponsePromptDataOptionsParamsObjectToolChoiceObjectTypeFunction:
		return true
	}
	return false
}

type PromptNewResponsePromptDataOrigin struct {
	ProjectID     string                                `json:"project_id"`
	PromptID      string                                `json:"prompt_id"`
	PromptVersion string                                `json:"prompt_version"`
	JSON          promptNewResponsePromptDataOriginJSON `json:"-"`
}

// promptNewResponsePromptDataOriginJSON contains the JSON metadata for the struct
// [PromptNewResponsePromptDataOrigin]
type promptNewResponsePromptDataOriginJSON struct {
	ProjectID     apijson.Field
	PromptID      apijson.Field
	PromptVersion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PromptNewResponsePromptDataOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptNewResponsePromptDataOriginJSON) RawJSON() string {
	return r.raw
}

type PromptNewResponsePromptDataPrompt struct {
	Type                  PromptNewResponsePromptDataPromptType `json:"type"`
	Content               string                                `json:"content"`
	Messages              interface{}                           `json:"messages,required"`
	Tools                 string                                `json:"tools"`
	ReservedOnlyAllowNull interface{}                           `json:"__reserved_only_allow_null,required"`
	JSON                  promptNewResponsePromptDataPromptJSON `json:"-"`
	union                 PromptNewResponsePromptDataPromptUnion
}

// promptNewResponsePromptDataPromptJSON contains the JSON metadata for the struct
// [PromptNewResponsePromptDataPrompt]
type promptNewResponsePromptDataPromptJSON struct {
	Type                  apijson.Field
	Content               apijson.Field
	Messages              apijson.Field
	Tools                 apijson.Field
	ReservedOnlyAllowNull apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r promptNewResponsePromptDataPromptJSON) RawJSON() string {
	return r.raw
}

func (r *PromptNewResponsePromptDataPrompt) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PromptNewResponsePromptDataPrompt) AsUnion() PromptNewResponsePromptDataPromptUnion {
	return r.union
}

// Union satisfied by [PromptNewResponsePromptDataPromptObject],
// [PromptNewResponsePromptDataPromptObject] or
// [PromptNewResponsePromptDataPromptObject].
type PromptNewResponsePromptDataPromptUnion interface {
	implementsPromptNewResponsePromptDataPrompt()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptNewResponsePromptDataPromptUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptNewResponsePromptDataPromptObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptNewResponsePromptDataPromptObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptNewResponsePromptDataPromptObject{}),
		},
	)
}

type PromptNewResponsePromptDataPromptObject struct {
	Content string                                      `json:"content,required"`
	Type    PromptNewResponsePromptDataPromptObjectType `json:"type,required"`
	JSON    promptNewResponsePromptDataPromptObjectJSON `json:"-"`
}

// promptNewResponsePromptDataPromptObjectJSON contains the JSON metadata for the
// struct [PromptNewResponsePromptDataPromptObject]
type promptNewResponsePromptDataPromptObjectJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptNewResponsePromptDataPromptObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptNewResponsePromptDataPromptObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptNewResponsePromptDataPromptObject) implementsPromptNewResponsePromptDataPrompt() {}

type PromptNewResponsePromptDataPromptObjectType string

const (
	PromptNewResponsePromptDataPromptObjectTypeCompletion PromptNewResponsePromptDataPromptObjectType = "completion"
)

func (r PromptNewResponsePromptDataPromptObjectType) IsKnown() bool {
	switch r {
	case PromptNewResponsePromptDataPromptObjectTypeCompletion:
		return true
	}
	return false
}

type PromptNewResponsePromptDataPromptType string

const (
	PromptNewResponsePromptDataPromptTypeCompletion PromptNewResponsePromptDataPromptType = "completion"
	PromptNewResponsePromptDataPromptTypeChat       PromptNewResponsePromptDataPromptType = "chat"
)

func (r PromptNewResponsePromptDataPromptType) IsKnown() bool {
	switch r {
	case PromptNewResponsePromptDataPromptTypeCompletion, PromptNewResponsePromptDataPromptTypeChat:
		return true
	}
	return false
}

type PromptGetResponse struct {
	// Unique identifier for the prompt
	ID string `json:"id,required" format:"uuid"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the prompt (see the `version`
	// parameter)
	XactID string `json:"_xact_id,required"`
	// A literal 'p' which identifies the object as a project prompt
	LogID PromptGetResponseLogID `json:"log_id,required"`
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
	PromptData PromptGetResponsePromptData `json:"prompt_data,nullable"`
	// A list of tags for the prompt
	Tags []string              `json:"tags,nullable"`
	JSON promptGetResponseJSON `json:"-"`
}

// promptGetResponseJSON contains the JSON metadata for the struct
// [PromptGetResponse]
type promptGetResponseJSON struct {
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

func (r *PromptGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptGetResponseJSON) RawJSON() string {
	return r.raw
}

// A literal 'p' which identifies the object as a project prompt
type PromptGetResponseLogID string

const (
	PromptGetResponseLogIDP PromptGetResponseLogID = "p"
)

func (r PromptGetResponseLogID) IsKnown() bool {
	switch r {
	case PromptGetResponseLogIDP:
		return true
	}
	return false
}

// The prompt, model, and its parameters
type PromptGetResponsePromptData struct {
	Options PromptGetResponsePromptDataOptions `json:"options,nullable"`
	Origin  PromptGetResponsePromptDataOrigin  `json:"origin,nullable"`
	Prompt  PromptGetResponsePromptDataPrompt  `json:"prompt"`
	JSON    promptGetResponsePromptDataJSON    `json:"-"`
}

// promptGetResponsePromptDataJSON contains the JSON metadata for the struct
// [PromptGetResponsePromptData]
type promptGetResponsePromptDataJSON struct {
	Options     apijson.Field
	Origin      apijson.Field
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptGetResponsePromptData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptGetResponsePromptDataJSON) RawJSON() string {
	return r.raw
}

type PromptGetResponsePromptDataOptions struct {
	Model    string                                   `json:"model"`
	Params   PromptGetResponsePromptDataOptionsParams `json:"params"`
	Position string                                   `json:"position"`
	JSON     promptGetResponsePromptDataOptionsJSON   `json:"-"`
}

// promptGetResponsePromptDataOptionsJSON contains the JSON metadata for the struct
// [PromptGetResponsePromptDataOptions]
type promptGetResponsePromptDataOptionsJSON struct {
	Model       apijson.Field
	Params      apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptGetResponsePromptDataOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptGetResponsePromptDataOptionsJSON) RawJSON() string {
	return r.raw
}

type PromptGetResponsePromptDataOptionsParams struct {
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
	MaxTokensToSample float64                                      `json:"max_tokens_to_sample"`
	MaxOutputTokens   float64                                      `json:"maxOutputTokens"`
	TopP              float64                                      `json:"topP"`
	TopK              float64                                      `json:"topK"`
	JSON              promptGetResponsePromptDataOptionsParamsJSON `json:"-"`
	union             PromptGetResponsePromptDataOptionsParamsUnion
}

// promptGetResponsePromptDataOptionsParamsJSON contains the JSON metadata for the
// struct [PromptGetResponsePromptDataOptionsParams]
type promptGetResponsePromptDataOptionsParamsJSON struct {
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

func (r promptGetResponsePromptDataOptionsParamsJSON) RawJSON() string {
	return r.raw
}

func (r *PromptGetResponsePromptDataOptionsParams) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PromptGetResponsePromptDataOptionsParams) AsUnion() PromptGetResponsePromptDataOptionsParamsUnion {
	return r.union
}

// Union satisfied by [PromptGetResponsePromptDataOptionsParamsObject],
// [PromptGetResponsePromptDataOptionsParamsObject],
// [PromptGetResponsePromptDataOptionsParamsObject] or
// [PromptGetResponsePromptDataOptionsParamsObject].
type PromptGetResponsePromptDataOptionsParamsUnion interface {
	implementsPromptGetResponsePromptDataOptionsParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptGetResponsePromptDataOptionsParamsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptGetResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptGetResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptGetResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptGetResponsePromptDataOptionsParamsObject{}),
		},
	)
}

type PromptGetResponsePromptDataOptionsParamsObject struct {
	FrequencyPenalty float64                                                         `json:"frequency_penalty"`
	FunctionCall     PromptGetResponsePromptDataOptionsParamsObjectFunctionCallUnion `json:"function_call"`
	MaxTokens        float64                                                         `json:"max_tokens"`
	N                float64                                                         `json:"n"`
	PresencePenalty  float64                                                         `json:"presence_penalty"`
	ResponseFormat   PromptGetResponsePromptDataOptionsParamsObjectResponseFormat    `json:"response_format,nullable"`
	Stop             []string                                                        `json:"stop"`
	Temperature      float64                                                         `json:"temperature"`
	ToolChoice       PromptGetResponsePromptDataOptionsParamsObjectToolChoiceUnion   `json:"tool_choice"`
	TopP             float64                                                         `json:"top_p"`
	UseCache         bool                                                            `json:"use_cache"`
	JSON             promptGetResponsePromptDataOptionsParamsObjectJSON              `json:"-"`
}

// promptGetResponsePromptDataOptionsParamsObjectJSON contains the JSON metadata
// for the struct [PromptGetResponsePromptDataOptionsParamsObject]
type promptGetResponsePromptDataOptionsParamsObjectJSON struct {
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

func (r *PromptGetResponsePromptDataOptionsParamsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptGetResponsePromptDataOptionsParamsObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptGetResponsePromptDataOptionsParamsObject) implementsPromptGetResponsePromptDataOptionsParams() {
}

// Union satisfied by
// [PromptGetResponsePromptDataOptionsParamsObjectFunctionCallString],
// [PromptGetResponsePromptDataOptionsParamsObjectFunctionCallString] or
// [PromptGetResponsePromptDataOptionsParamsObjectFunctionCallObject].
type PromptGetResponsePromptDataOptionsParamsObjectFunctionCallUnion interface {
	ImplementsPromptGetResponsePromptDataOptionsParamsObjectFunctionCallUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptGetResponsePromptDataOptionsParamsObjectFunctionCallUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptGetResponsePromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptGetResponsePromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptGetResponsePromptDataOptionsParamsObjectFunctionCallObject{}),
		},
	)
}

type PromptGetResponsePromptDataOptionsParamsObjectFunctionCallString string

const (
	PromptGetResponsePromptDataOptionsParamsObjectFunctionCallStringAuto PromptGetResponsePromptDataOptionsParamsObjectFunctionCallString = "auto"
)

func (r PromptGetResponsePromptDataOptionsParamsObjectFunctionCallString) IsKnown() bool {
	switch r {
	case PromptGetResponsePromptDataOptionsParamsObjectFunctionCallStringAuto:
		return true
	}
	return false
}

type PromptGetResponsePromptDataOptionsParamsObjectFunctionCallObject struct {
	Name string                                                               `json:"name,required"`
	JSON promptGetResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON `json:"-"`
}

// promptGetResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON contains
// the JSON metadata for the struct
// [PromptGetResponsePromptDataOptionsParamsObjectFunctionCallObject]
type promptGetResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptGetResponsePromptDataOptionsParamsObjectFunctionCallObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptGetResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptGetResponsePromptDataOptionsParamsObjectFunctionCallObject) ImplementsPromptGetResponsePromptDataOptionsParamsObjectFunctionCallUnion() {
}

type PromptGetResponsePromptDataOptionsParamsObjectResponseFormat struct {
	Type PromptGetResponsePromptDataOptionsParamsObjectResponseFormatType `json:"type,required"`
	JSON promptGetResponsePromptDataOptionsParamsObjectResponseFormatJSON `json:"-"`
}

// promptGetResponsePromptDataOptionsParamsObjectResponseFormatJSON contains the
// JSON metadata for the struct
// [PromptGetResponsePromptDataOptionsParamsObjectResponseFormat]
type promptGetResponsePromptDataOptionsParamsObjectResponseFormatJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptGetResponsePromptDataOptionsParamsObjectResponseFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptGetResponsePromptDataOptionsParamsObjectResponseFormatJSON) RawJSON() string {
	return r.raw
}

type PromptGetResponsePromptDataOptionsParamsObjectResponseFormatType string

const (
	PromptGetResponsePromptDataOptionsParamsObjectResponseFormatTypeJsonObject PromptGetResponsePromptDataOptionsParamsObjectResponseFormatType = "json_object"
)

func (r PromptGetResponsePromptDataOptionsParamsObjectResponseFormatType) IsKnown() bool {
	switch r {
	case PromptGetResponsePromptDataOptionsParamsObjectResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Union satisfied by
// [PromptGetResponsePromptDataOptionsParamsObjectToolChoiceString],
// [PromptGetResponsePromptDataOptionsParamsObjectToolChoiceString] or
// [PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObject].
type PromptGetResponsePromptDataOptionsParamsObjectToolChoiceUnion interface {
	ImplementsPromptGetResponsePromptDataOptionsParamsObjectToolChoiceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptGetResponsePromptDataOptionsParamsObjectToolChoiceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptGetResponsePromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptGetResponsePromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObject{}),
		},
	)
}

type PromptGetResponsePromptDataOptionsParamsObjectToolChoiceString string

const (
	PromptGetResponsePromptDataOptionsParamsObjectToolChoiceStringAuto PromptGetResponsePromptDataOptionsParamsObjectToolChoiceString = "auto"
)

func (r PromptGetResponsePromptDataOptionsParamsObjectToolChoiceString) IsKnown() bool {
	switch r {
	case PromptGetResponsePromptDataOptionsParamsObjectToolChoiceStringAuto:
		return true
	}
	return false
}

type PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObject struct {
	Function PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction `json:"function,required"`
	Type     PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectType     `json:"type,required"`
	JSON     promptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON     `json:"-"`
}

// promptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON contains the
// JSON metadata for the struct
// [PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObject]
type promptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON struct {
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObject) ImplementsPromptGetResponsePromptDataOptionsParamsObjectToolChoiceUnion() {
}

type PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction struct {
	Name string                                                                     `json:"name,required"`
	JSON promptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON `json:"-"`
}

// promptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON
// contains the JSON metadata for the struct
// [PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction]
type promptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON) RawJSON() string {
	return r.raw
}

type PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectType string

const (
	PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectTypeFunction PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectType = "function"
)

func (r PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectType) IsKnown() bool {
	switch r {
	case PromptGetResponsePromptDataOptionsParamsObjectToolChoiceObjectTypeFunction:
		return true
	}
	return false
}

type PromptGetResponsePromptDataOrigin struct {
	ProjectID     string                                `json:"project_id"`
	PromptID      string                                `json:"prompt_id"`
	PromptVersion string                                `json:"prompt_version"`
	JSON          promptGetResponsePromptDataOriginJSON `json:"-"`
}

// promptGetResponsePromptDataOriginJSON contains the JSON metadata for the struct
// [PromptGetResponsePromptDataOrigin]
type promptGetResponsePromptDataOriginJSON struct {
	ProjectID     apijson.Field
	PromptID      apijson.Field
	PromptVersion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PromptGetResponsePromptDataOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptGetResponsePromptDataOriginJSON) RawJSON() string {
	return r.raw
}

type PromptGetResponsePromptDataPrompt struct {
	Type                  PromptGetResponsePromptDataPromptType `json:"type"`
	Content               string                                `json:"content"`
	Messages              interface{}                           `json:"messages,required"`
	Tools                 string                                `json:"tools"`
	ReservedOnlyAllowNull interface{}                           `json:"__reserved_only_allow_null,required"`
	JSON                  promptGetResponsePromptDataPromptJSON `json:"-"`
	union                 PromptGetResponsePromptDataPromptUnion
}

// promptGetResponsePromptDataPromptJSON contains the JSON metadata for the struct
// [PromptGetResponsePromptDataPrompt]
type promptGetResponsePromptDataPromptJSON struct {
	Type                  apijson.Field
	Content               apijson.Field
	Messages              apijson.Field
	Tools                 apijson.Field
	ReservedOnlyAllowNull apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r promptGetResponsePromptDataPromptJSON) RawJSON() string {
	return r.raw
}

func (r *PromptGetResponsePromptDataPrompt) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PromptGetResponsePromptDataPrompt) AsUnion() PromptGetResponsePromptDataPromptUnion {
	return r.union
}

// Union satisfied by [PromptGetResponsePromptDataPromptObject],
// [PromptGetResponsePromptDataPromptObject] or
// [PromptGetResponsePromptDataPromptObject].
type PromptGetResponsePromptDataPromptUnion interface {
	implementsPromptGetResponsePromptDataPrompt()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptGetResponsePromptDataPromptUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptGetResponsePromptDataPromptObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptGetResponsePromptDataPromptObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptGetResponsePromptDataPromptObject{}),
		},
	)
}

type PromptGetResponsePromptDataPromptObject struct {
	Content string                                      `json:"content,required"`
	Type    PromptGetResponsePromptDataPromptObjectType `json:"type,required"`
	JSON    promptGetResponsePromptDataPromptObjectJSON `json:"-"`
}

// promptGetResponsePromptDataPromptObjectJSON contains the JSON metadata for the
// struct [PromptGetResponsePromptDataPromptObject]
type promptGetResponsePromptDataPromptObjectJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptGetResponsePromptDataPromptObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptGetResponsePromptDataPromptObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptGetResponsePromptDataPromptObject) implementsPromptGetResponsePromptDataPrompt() {}

type PromptGetResponsePromptDataPromptObjectType string

const (
	PromptGetResponsePromptDataPromptObjectTypeCompletion PromptGetResponsePromptDataPromptObjectType = "completion"
)

func (r PromptGetResponsePromptDataPromptObjectType) IsKnown() bool {
	switch r {
	case PromptGetResponsePromptDataPromptObjectTypeCompletion:
		return true
	}
	return false
}

type PromptGetResponsePromptDataPromptType string

const (
	PromptGetResponsePromptDataPromptTypeCompletion PromptGetResponsePromptDataPromptType = "completion"
	PromptGetResponsePromptDataPromptTypeChat       PromptGetResponsePromptDataPromptType = "chat"
)

func (r PromptGetResponsePromptDataPromptType) IsKnown() bool {
	switch r {
	case PromptGetResponsePromptDataPromptTypeCompletion, PromptGetResponsePromptDataPromptTypeChat:
		return true
	}
	return false
}

type PromptUpdateResponse struct {
	// Unique identifier for the prompt
	ID string `json:"id,required" format:"uuid"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the prompt (see the `version`
	// parameter)
	XactID string `json:"_xact_id,required"`
	// A literal 'p' which identifies the object as a project prompt
	LogID PromptUpdateResponseLogID `json:"log_id,required"`
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
	PromptData PromptUpdateResponsePromptData `json:"prompt_data,nullable"`
	// A list of tags for the prompt
	Tags []string                 `json:"tags,nullable"`
	JSON promptUpdateResponseJSON `json:"-"`
}

// promptUpdateResponseJSON contains the JSON metadata for the struct
// [PromptUpdateResponse]
type promptUpdateResponseJSON struct {
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

func (r *PromptUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A literal 'p' which identifies the object as a project prompt
type PromptUpdateResponseLogID string

const (
	PromptUpdateResponseLogIDP PromptUpdateResponseLogID = "p"
)

func (r PromptUpdateResponseLogID) IsKnown() bool {
	switch r {
	case PromptUpdateResponseLogIDP:
		return true
	}
	return false
}

// The prompt, model, and its parameters
type PromptUpdateResponsePromptData struct {
	Options PromptUpdateResponsePromptDataOptions `json:"options,nullable"`
	Origin  PromptUpdateResponsePromptDataOrigin  `json:"origin,nullable"`
	Prompt  PromptUpdateResponsePromptDataPrompt  `json:"prompt"`
	JSON    promptUpdateResponsePromptDataJSON    `json:"-"`
}

// promptUpdateResponsePromptDataJSON contains the JSON metadata for the struct
// [PromptUpdateResponsePromptData]
type promptUpdateResponsePromptDataJSON struct {
	Options     apijson.Field
	Origin      apijson.Field
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptUpdateResponsePromptData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptUpdateResponsePromptDataJSON) RawJSON() string {
	return r.raw
}

type PromptUpdateResponsePromptDataOptions struct {
	Model    string                                      `json:"model"`
	Params   PromptUpdateResponsePromptDataOptionsParams `json:"params"`
	Position string                                      `json:"position"`
	JSON     promptUpdateResponsePromptDataOptionsJSON   `json:"-"`
}

// promptUpdateResponsePromptDataOptionsJSON contains the JSON metadata for the
// struct [PromptUpdateResponsePromptDataOptions]
type promptUpdateResponsePromptDataOptionsJSON struct {
	Model       apijson.Field
	Params      apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptUpdateResponsePromptDataOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptUpdateResponsePromptDataOptionsJSON) RawJSON() string {
	return r.raw
}

type PromptUpdateResponsePromptDataOptionsParams struct {
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
	MaxTokensToSample float64                                         `json:"max_tokens_to_sample"`
	MaxOutputTokens   float64                                         `json:"maxOutputTokens"`
	TopP              float64                                         `json:"topP"`
	TopK              float64                                         `json:"topK"`
	JSON              promptUpdateResponsePromptDataOptionsParamsJSON `json:"-"`
	union             PromptUpdateResponsePromptDataOptionsParamsUnion
}

// promptUpdateResponsePromptDataOptionsParamsJSON contains the JSON metadata for
// the struct [PromptUpdateResponsePromptDataOptionsParams]
type promptUpdateResponsePromptDataOptionsParamsJSON struct {
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

func (r promptUpdateResponsePromptDataOptionsParamsJSON) RawJSON() string {
	return r.raw
}

func (r *PromptUpdateResponsePromptDataOptionsParams) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PromptUpdateResponsePromptDataOptionsParams) AsUnion() PromptUpdateResponsePromptDataOptionsParamsUnion {
	return r.union
}

// Union satisfied by [PromptUpdateResponsePromptDataOptionsParamsObject],
// [PromptUpdateResponsePromptDataOptionsParamsObject],
// [PromptUpdateResponsePromptDataOptionsParamsObject] or
// [PromptUpdateResponsePromptDataOptionsParamsObject].
type PromptUpdateResponsePromptDataOptionsParamsUnion interface {
	implementsPromptUpdateResponsePromptDataOptionsParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptUpdateResponsePromptDataOptionsParamsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptUpdateResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptUpdateResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptUpdateResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptUpdateResponsePromptDataOptionsParamsObject{}),
		},
	)
}

type PromptUpdateResponsePromptDataOptionsParamsObject struct {
	FrequencyPenalty float64                                                            `json:"frequency_penalty"`
	FunctionCall     PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallUnion `json:"function_call"`
	MaxTokens        float64                                                            `json:"max_tokens"`
	N                float64                                                            `json:"n"`
	PresencePenalty  float64                                                            `json:"presence_penalty"`
	ResponseFormat   PromptUpdateResponsePromptDataOptionsParamsObjectResponseFormat    `json:"response_format,nullable"`
	Stop             []string                                                           `json:"stop"`
	Temperature      float64                                                            `json:"temperature"`
	ToolChoice       PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceUnion   `json:"tool_choice"`
	TopP             float64                                                            `json:"top_p"`
	UseCache         bool                                                               `json:"use_cache"`
	JSON             promptUpdateResponsePromptDataOptionsParamsObjectJSON              `json:"-"`
}

// promptUpdateResponsePromptDataOptionsParamsObjectJSON contains the JSON metadata
// for the struct [PromptUpdateResponsePromptDataOptionsParamsObject]
type promptUpdateResponsePromptDataOptionsParamsObjectJSON struct {
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

func (r *PromptUpdateResponsePromptDataOptionsParamsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptUpdateResponsePromptDataOptionsParamsObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptUpdateResponsePromptDataOptionsParamsObject) implementsPromptUpdateResponsePromptDataOptionsParams() {
}

// Union satisfied by
// [PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallString],
// [PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallString] or
// [PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallObject].
type PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallUnion interface {
	ImplementsPromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallObject{}),
		},
	)
}

type PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallString string

const (
	PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallStringAuto PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallString = "auto"
)

func (r PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallString) IsKnown() bool {
	switch r {
	case PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallStringAuto:
		return true
	}
	return false
}

type PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallObject struct {
	Name string                                                                  `json:"name,required"`
	JSON promptUpdateResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON `json:"-"`
}

// promptUpdateResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON contains
// the JSON metadata for the struct
// [PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallObject]
type promptUpdateResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptUpdateResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallObject) ImplementsPromptUpdateResponsePromptDataOptionsParamsObjectFunctionCallUnion() {
}

type PromptUpdateResponsePromptDataOptionsParamsObjectResponseFormat struct {
	Type PromptUpdateResponsePromptDataOptionsParamsObjectResponseFormatType `json:"type,required"`
	JSON promptUpdateResponsePromptDataOptionsParamsObjectResponseFormatJSON `json:"-"`
}

// promptUpdateResponsePromptDataOptionsParamsObjectResponseFormatJSON contains the
// JSON metadata for the struct
// [PromptUpdateResponsePromptDataOptionsParamsObjectResponseFormat]
type promptUpdateResponsePromptDataOptionsParamsObjectResponseFormatJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptUpdateResponsePromptDataOptionsParamsObjectResponseFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptUpdateResponsePromptDataOptionsParamsObjectResponseFormatJSON) RawJSON() string {
	return r.raw
}

type PromptUpdateResponsePromptDataOptionsParamsObjectResponseFormatType string

const (
	PromptUpdateResponsePromptDataOptionsParamsObjectResponseFormatTypeJsonObject PromptUpdateResponsePromptDataOptionsParamsObjectResponseFormatType = "json_object"
)

func (r PromptUpdateResponsePromptDataOptionsParamsObjectResponseFormatType) IsKnown() bool {
	switch r {
	case PromptUpdateResponsePromptDataOptionsParamsObjectResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Union satisfied by
// [PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceString],
// [PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceString] or
// [PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObject].
type PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceUnion interface {
	ImplementsPromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObject{}),
		},
	)
}

type PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceString string

const (
	PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceStringAuto PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceString = "auto"
)

func (r PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceString) IsKnown() bool {
	switch r {
	case PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceStringAuto:
		return true
	}
	return false
}

type PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObject struct {
	Function PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction `json:"function,required"`
	Type     PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectType     `json:"type,required"`
	JSON     promptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON     `json:"-"`
}

// promptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON contains
// the JSON metadata for the struct
// [PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObject]
type promptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON struct {
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObject) ImplementsPromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceUnion() {
}

type PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction struct {
	Name string                                                                        `json:"name,required"`
	JSON promptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON `json:"-"`
}

// promptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON
// contains the JSON metadata for the struct
// [PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction]
type promptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON) RawJSON() string {
	return r.raw
}

type PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectType string

const (
	PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectTypeFunction PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectType = "function"
)

func (r PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectType) IsKnown() bool {
	switch r {
	case PromptUpdateResponsePromptDataOptionsParamsObjectToolChoiceObjectTypeFunction:
		return true
	}
	return false
}

type PromptUpdateResponsePromptDataOrigin struct {
	ProjectID     string                                   `json:"project_id"`
	PromptID      string                                   `json:"prompt_id"`
	PromptVersion string                                   `json:"prompt_version"`
	JSON          promptUpdateResponsePromptDataOriginJSON `json:"-"`
}

// promptUpdateResponsePromptDataOriginJSON contains the JSON metadata for the
// struct [PromptUpdateResponsePromptDataOrigin]
type promptUpdateResponsePromptDataOriginJSON struct {
	ProjectID     apijson.Field
	PromptID      apijson.Field
	PromptVersion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PromptUpdateResponsePromptDataOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptUpdateResponsePromptDataOriginJSON) RawJSON() string {
	return r.raw
}

type PromptUpdateResponsePromptDataPrompt struct {
	Type                  PromptUpdateResponsePromptDataPromptType `json:"type"`
	Content               string                                   `json:"content"`
	Messages              interface{}                              `json:"messages,required"`
	Tools                 string                                   `json:"tools"`
	ReservedOnlyAllowNull interface{}                              `json:"__reserved_only_allow_null,required"`
	JSON                  promptUpdateResponsePromptDataPromptJSON `json:"-"`
	union                 PromptUpdateResponsePromptDataPromptUnion
}

// promptUpdateResponsePromptDataPromptJSON contains the JSON metadata for the
// struct [PromptUpdateResponsePromptDataPrompt]
type promptUpdateResponsePromptDataPromptJSON struct {
	Type                  apijson.Field
	Content               apijson.Field
	Messages              apijson.Field
	Tools                 apijson.Field
	ReservedOnlyAllowNull apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r promptUpdateResponsePromptDataPromptJSON) RawJSON() string {
	return r.raw
}

func (r *PromptUpdateResponsePromptDataPrompt) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PromptUpdateResponsePromptDataPrompt) AsUnion() PromptUpdateResponsePromptDataPromptUnion {
	return r.union
}

// Union satisfied by [PromptUpdateResponsePromptDataPromptObject],
// [PromptUpdateResponsePromptDataPromptObject] or
// [PromptUpdateResponsePromptDataPromptObject].
type PromptUpdateResponsePromptDataPromptUnion interface {
	implementsPromptUpdateResponsePromptDataPrompt()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptUpdateResponsePromptDataPromptUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptUpdateResponsePromptDataPromptObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptUpdateResponsePromptDataPromptObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptUpdateResponsePromptDataPromptObject{}),
		},
	)
}

type PromptUpdateResponsePromptDataPromptObject struct {
	Content string                                         `json:"content,required"`
	Type    PromptUpdateResponsePromptDataPromptObjectType `json:"type,required"`
	JSON    promptUpdateResponsePromptDataPromptObjectJSON `json:"-"`
}

// promptUpdateResponsePromptDataPromptObjectJSON contains the JSON metadata for
// the struct [PromptUpdateResponsePromptDataPromptObject]
type promptUpdateResponsePromptDataPromptObjectJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptUpdateResponsePromptDataPromptObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptUpdateResponsePromptDataPromptObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptUpdateResponsePromptDataPromptObject) implementsPromptUpdateResponsePromptDataPrompt() {
}

type PromptUpdateResponsePromptDataPromptObjectType string

const (
	PromptUpdateResponsePromptDataPromptObjectTypeCompletion PromptUpdateResponsePromptDataPromptObjectType = "completion"
)

func (r PromptUpdateResponsePromptDataPromptObjectType) IsKnown() bool {
	switch r {
	case PromptUpdateResponsePromptDataPromptObjectTypeCompletion:
		return true
	}
	return false
}

type PromptUpdateResponsePromptDataPromptType string

const (
	PromptUpdateResponsePromptDataPromptTypeCompletion PromptUpdateResponsePromptDataPromptType = "completion"
	PromptUpdateResponsePromptDataPromptTypeChat       PromptUpdateResponsePromptDataPromptType = "chat"
)

func (r PromptUpdateResponsePromptDataPromptType) IsKnown() bool {
	switch r {
	case PromptUpdateResponsePromptDataPromptTypeCompletion, PromptUpdateResponsePromptDataPromptTypeChat:
		return true
	}
	return false
}

type PromptListResponse struct {
	// Unique identifier for the prompt
	ID string `json:"id,required" format:"uuid"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the prompt (see the `version`
	// parameter)
	XactID string `json:"_xact_id,required"`
	// A literal 'p' which identifies the object as a project prompt
	LogID PromptListResponseLogID `json:"log_id,required"`
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
	PromptData PromptListResponsePromptData `json:"prompt_data,nullable"`
	// A list of tags for the prompt
	Tags []string               `json:"tags,nullable"`
	JSON promptListResponseJSON `json:"-"`
}

// promptListResponseJSON contains the JSON metadata for the struct
// [PromptListResponse]
type promptListResponseJSON struct {
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

func (r *PromptListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptListResponseJSON) RawJSON() string {
	return r.raw
}

// A literal 'p' which identifies the object as a project prompt
type PromptListResponseLogID string

const (
	PromptListResponseLogIDP PromptListResponseLogID = "p"
)

func (r PromptListResponseLogID) IsKnown() bool {
	switch r {
	case PromptListResponseLogIDP:
		return true
	}
	return false
}

// The prompt, model, and its parameters
type PromptListResponsePromptData struct {
	Options PromptListResponsePromptDataOptions `json:"options,nullable"`
	Origin  PromptListResponsePromptDataOrigin  `json:"origin,nullable"`
	Prompt  PromptListResponsePromptDataPrompt  `json:"prompt"`
	JSON    promptListResponsePromptDataJSON    `json:"-"`
}

// promptListResponsePromptDataJSON contains the JSON metadata for the struct
// [PromptListResponsePromptData]
type promptListResponsePromptDataJSON struct {
	Options     apijson.Field
	Origin      apijson.Field
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptListResponsePromptData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptListResponsePromptDataJSON) RawJSON() string {
	return r.raw
}

type PromptListResponsePromptDataOptions struct {
	Model    string                                    `json:"model"`
	Params   PromptListResponsePromptDataOptionsParams `json:"params"`
	Position string                                    `json:"position"`
	JSON     promptListResponsePromptDataOptionsJSON   `json:"-"`
}

// promptListResponsePromptDataOptionsJSON contains the JSON metadata for the
// struct [PromptListResponsePromptDataOptions]
type promptListResponsePromptDataOptionsJSON struct {
	Model       apijson.Field
	Params      apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptListResponsePromptDataOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptListResponsePromptDataOptionsJSON) RawJSON() string {
	return r.raw
}

type PromptListResponsePromptDataOptionsParams struct {
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
	MaxTokensToSample float64                                       `json:"max_tokens_to_sample"`
	MaxOutputTokens   float64                                       `json:"maxOutputTokens"`
	TopP              float64                                       `json:"topP"`
	TopK              float64                                       `json:"topK"`
	JSON              promptListResponsePromptDataOptionsParamsJSON `json:"-"`
	union             PromptListResponsePromptDataOptionsParamsUnion
}

// promptListResponsePromptDataOptionsParamsJSON contains the JSON metadata for the
// struct [PromptListResponsePromptDataOptionsParams]
type promptListResponsePromptDataOptionsParamsJSON struct {
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

func (r promptListResponsePromptDataOptionsParamsJSON) RawJSON() string {
	return r.raw
}

func (r *PromptListResponsePromptDataOptionsParams) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PromptListResponsePromptDataOptionsParams) AsUnion() PromptListResponsePromptDataOptionsParamsUnion {
	return r.union
}

// Union satisfied by [PromptListResponsePromptDataOptionsParamsObject],
// [PromptListResponsePromptDataOptionsParamsObject],
// [PromptListResponsePromptDataOptionsParamsObject] or
// [PromptListResponsePromptDataOptionsParamsObject].
type PromptListResponsePromptDataOptionsParamsUnion interface {
	implementsPromptListResponsePromptDataOptionsParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptListResponsePromptDataOptionsParamsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptListResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptListResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptListResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptListResponsePromptDataOptionsParamsObject{}),
		},
	)
}

type PromptListResponsePromptDataOptionsParamsObject struct {
	FrequencyPenalty float64                                                          `json:"frequency_penalty"`
	FunctionCall     PromptListResponsePromptDataOptionsParamsObjectFunctionCallUnion `json:"function_call"`
	MaxTokens        float64                                                          `json:"max_tokens"`
	N                float64                                                          `json:"n"`
	PresencePenalty  float64                                                          `json:"presence_penalty"`
	ResponseFormat   PromptListResponsePromptDataOptionsParamsObjectResponseFormat    `json:"response_format,nullable"`
	Stop             []string                                                         `json:"stop"`
	Temperature      float64                                                          `json:"temperature"`
	ToolChoice       PromptListResponsePromptDataOptionsParamsObjectToolChoiceUnion   `json:"tool_choice"`
	TopP             float64                                                          `json:"top_p"`
	UseCache         bool                                                             `json:"use_cache"`
	JSON             promptListResponsePromptDataOptionsParamsObjectJSON              `json:"-"`
}

// promptListResponsePromptDataOptionsParamsObjectJSON contains the JSON metadata
// for the struct [PromptListResponsePromptDataOptionsParamsObject]
type promptListResponsePromptDataOptionsParamsObjectJSON struct {
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

func (r *PromptListResponsePromptDataOptionsParamsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptListResponsePromptDataOptionsParamsObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptListResponsePromptDataOptionsParamsObject) implementsPromptListResponsePromptDataOptionsParams() {
}

// Union satisfied by
// [PromptListResponsePromptDataOptionsParamsObjectFunctionCallString],
// [PromptListResponsePromptDataOptionsParamsObjectFunctionCallString] or
// [PromptListResponsePromptDataOptionsParamsObjectFunctionCallObject].
type PromptListResponsePromptDataOptionsParamsObjectFunctionCallUnion interface {
	ImplementsPromptListResponsePromptDataOptionsParamsObjectFunctionCallUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptListResponsePromptDataOptionsParamsObjectFunctionCallUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptListResponsePromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptListResponsePromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptListResponsePromptDataOptionsParamsObjectFunctionCallObject{}),
		},
	)
}

type PromptListResponsePromptDataOptionsParamsObjectFunctionCallString string

const (
	PromptListResponsePromptDataOptionsParamsObjectFunctionCallStringAuto PromptListResponsePromptDataOptionsParamsObjectFunctionCallString = "auto"
)

func (r PromptListResponsePromptDataOptionsParamsObjectFunctionCallString) IsKnown() bool {
	switch r {
	case PromptListResponsePromptDataOptionsParamsObjectFunctionCallStringAuto:
		return true
	}
	return false
}

type PromptListResponsePromptDataOptionsParamsObjectFunctionCallObject struct {
	Name string                                                                `json:"name,required"`
	JSON promptListResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON `json:"-"`
}

// promptListResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON contains
// the JSON metadata for the struct
// [PromptListResponsePromptDataOptionsParamsObjectFunctionCallObject]
type promptListResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptListResponsePromptDataOptionsParamsObjectFunctionCallObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptListResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptListResponsePromptDataOptionsParamsObjectFunctionCallObject) ImplementsPromptListResponsePromptDataOptionsParamsObjectFunctionCallUnion() {
}

type PromptListResponsePromptDataOptionsParamsObjectResponseFormat struct {
	Type PromptListResponsePromptDataOptionsParamsObjectResponseFormatType `json:"type,required"`
	JSON promptListResponsePromptDataOptionsParamsObjectResponseFormatJSON `json:"-"`
}

// promptListResponsePromptDataOptionsParamsObjectResponseFormatJSON contains the
// JSON metadata for the struct
// [PromptListResponsePromptDataOptionsParamsObjectResponseFormat]
type promptListResponsePromptDataOptionsParamsObjectResponseFormatJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptListResponsePromptDataOptionsParamsObjectResponseFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptListResponsePromptDataOptionsParamsObjectResponseFormatJSON) RawJSON() string {
	return r.raw
}

type PromptListResponsePromptDataOptionsParamsObjectResponseFormatType string

const (
	PromptListResponsePromptDataOptionsParamsObjectResponseFormatTypeJsonObject PromptListResponsePromptDataOptionsParamsObjectResponseFormatType = "json_object"
)

func (r PromptListResponsePromptDataOptionsParamsObjectResponseFormatType) IsKnown() bool {
	switch r {
	case PromptListResponsePromptDataOptionsParamsObjectResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Union satisfied by
// [PromptListResponsePromptDataOptionsParamsObjectToolChoiceString],
// [PromptListResponsePromptDataOptionsParamsObjectToolChoiceString] or
// [PromptListResponsePromptDataOptionsParamsObjectToolChoiceObject].
type PromptListResponsePromptDataOptionsParamsObjectToolChoiceUnion interface {
	ImplementsPromptListResponsePromptDataOptionsParamsObjectToolChoiceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptListResponsePromptDataOptionsParamsObjectToolChoiceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptListResponsePromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptListResponsePromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptListResponsePromptDataOptionsParamsObjectToolChoiceObject{}),
		},
	)
}

type PromptListResponsePromptDataOptionsParamsObjectToolChoiceString string

const (
	PromptListResponsePromptDataOptionsParamsObjectToolChoiceStringAuto PromptListResponsePromptDataOptionsParamsObjectToolChoiceString = "auto"
)

func (r PromptListResponsePromptDataOptionsParamsObjectToolChoiceString) IsKnown() bool {
	switch r {
	case PromptListResponsePromptDataOptionsParamsObjectToolChoiceStringAuto:
		return true
	}
	return false
}

type PromptListResponsePromptDataOptionsParamsObjectToolChoiceObject struct {
	Function PromptListResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction `json:"function,required"`
	Type     PromptListResponsePromptDataOptionsParamsObjectToolChoiceObjectType     `json:"type,required"`
	JSON     promptListResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON     `json:"-"`
}

// promptListResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON contains the
// JSON metadata for the struct
// [PromptListResponsePromptDataOptionsParamsObjectToolChoiceObject]
type promptListResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON struct {
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptListResponsePromptDataOptionsParamsObjectToolChoiceObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptListResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptListResponsePromptDataOptionsParamsObjectToolChoiceObject) ImplementsPromptListResponsePromptDataOptionsParamsObjectToolChoiceUnion() {
}

type PromptListResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction struct {
	Name string                                                                      `json:"name,required"`
	JSON promptListResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON `json:"-"`
}

// promptListResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON
// contains the JSON metadata for the struct
// [PromptListResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction]
type promptListResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptListResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptListResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON) RawJSON() string {
	return r.raw
}

type PromptListResponsePromptDataOptionsParamsObjectToolChoiceObjectType string

const (
	PromptListResponsePromptDataOptionsParamsObjectToolChoiceObjectTypeFunction PromptListResponsePromptDataOptionsParamsObjectToolChoiceObjectType = "function"
)

func (r PromptListResponsePromptDataOptionsParamsObjectToolChoiceObjectType) IsKnown() bool {
	switch r {
	case PromptListResponsePromptDataOptionsParamsObjectToolChoiceObjectTypeFunction:
		return true
	}
	return false
}

type PromptListResponsePromptDataOrigin struct {
	ProjectID     string                                 `json:"project_id"`
	PromptID      string                                 `json:"prompt_id"`
	PromptVersion string                                 `json:"prompt_version"`
	JSON          promptListResponsePromptDataOriginJSON `json:"-"`
}

// promptListResponsePromptDataOriginJSON contains the JSON metadata for the struct
// [PromptListResponsePromptDataOrigin]
type promptListResponsePromptDataOriginJSON struct {
	ProjectID     apijson.Field
	PromptID      apijson.Field
	PromptVersion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PromptListResponsePromptDataOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptListResponsePromptDataOriginJSON) RawJSON() string {
	return r.raw
}

type PromptListResponsePromptDataPrompt struct {
	Type                  PromptListResponsePromptDataPromptType `json:"type"`
	Content               string                                 `json:"content"`
	Messages              interface{}                            `json:"messages,required"`
	Tools                 string                                 `json:"tools"`
	ReservedOnlyAllowNull interface{}                            `json:"__reserved_only_allow_null,required"`
	JSON                  promptListResponsePromptDataPromptJSON `json:"-"`
	union                 PromptListResponsePromptDataPromptUnion
}

// promptListResponsePromptDataPromptJSON contains the JSON metadata for the struct
// [PromptListResponsePromptDataPrompt]
type promptListResponsePromptDataPromptJSON struct {
	Type                  apijson.Field
	Content               apijson.Field
	Messages              apijson.Field
	Tools                 apijson.Field
	ReservedOnlyAllowNull apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r promptListResponsePromptDataPromptJSON) RawJSON() string {
	return r.raw
}

func (r *PromptListResponsePromptDataPrompt) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PromptListResponsePromptDataPrompt) AsUnion() PromptListResponsePromptDataPromptUnion {
	return r.union
}

// Union satisfied by [PromptListResponsePromptDataPromptObject],
// [PromptListResponsePromptDataPromptObject] or
// [PromptListResponsePromptDataPromptObject].
type PromptListResponsePromptDataPromptUnion interface {
	implementsPromptListResponsePromptDataPrompt()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptListResponsePromptDataPromptUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptListResponsePromptDataPromptObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptListResponsePromptDataPromptObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptListResponsePromptDataPromptObject{}),
		},
	)
}

type PromptListResponsePromptDataPromptObject struct {
	Content string                                       `json:"content,required"`
	Type    PromptListResponsePromptDataPromptObjectType `json:"type,required"`
	JSON    promptListResponsePromptDataPromptObjectJSON `json:"-"`
}

// promptListResponsePromptDataPromptObjectJSON contains the JSON metadata for the
// struct [PromptListResponsePromptDataPromptObject]
type promptListResponsePromptDataPromptObjectJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptListResponsePromptDataPromptObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptListResponsePromptDataPromptObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptListResponsePromptDataPromptObject) implementsPromptListResponsePromptDataPrompt() {}

type PromptListResponsePromptDataPromptObjectType string

const (
	PromptListResponsePromptDataPromptObjectTypeCompletion PromptListResponsePromptDataPromptObjectType = "completion"
)

func (r PromptListResponsePromptDataPromptObjectType) IsKnown() bool {
	switch r {
	case PromptListResponsePromptDataPromptObjectTypeCompletion:
		return true
	}
	return false
}

type PromptListResponsePromptDataPromptType string

const (
	PromptListResponsePromptDataPromptTypeCompletion PromptListResponsePromptDataPromptType = "completion"
	PromptListResponsePromptDataPromptTypeChat       PromptListResponsePromptDataPromptType = "chat"
)

func (r PromptListResponsePromptDataPromptType) IsKnown() bool {
	switch r {
	case PromptListResponsePromptDataPromptTypeCompletion, PromptListResponsePromptDataPromptTypeChat:
		return true
	}
	return false
}

type PromptDeleteResponse struct {
	// Unique identifier for the prompt
	ID string `json:"id,required" format:"uuid"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the prompt (see the `version`
	// parameter)
	XactID string `json:"_xact_id,required"`
	// A literal 'p' which identifies the object as a project prompt
	LogID PromptDeleteResponseLogID `json:"log_id,required"`
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
	PromptData PromptDeleteResponsePromptData `json:"prompt_data,nullable"`
	// A list of tags for the prompt
	Tags []string                 `json:"tags,nullable"`
	JSON promptDeleteResponseJSON `json:"-"`
}

// promptDeleteResponseJSON contains the JSON metadata for the struct
// [PromptDeleteResponse]
type promptDeleteResponseJSON struct {
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

func (r *PromptDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A literal 'p' which identifies the object as a project prompt
type PromptDeleteResponseLogID string

const (
	PromptDeleteResponseLogIDP PromptDeleteResponseLogID = "p"
)

func (r PromptDeleteResponseLogID) IsKnown() bool {
	switch r {
	case PromptDeleteResponseLogIDP:
		return true
	}
	return false
}

// The prompt, model, and its parameters
type PromptDeleteResponsePromptData struct {
	Options PromptDeleteResponsePromptDataOptions `json:"options,nullable"`
	Origin  PromptDeleteResponsePromptDataOrigin  `json:"origin,nullable"`
	Prompt  PromptDeleteResponsePromptDataPrompt  `json:"prompt"`
	JSON    promptDeleteResponsePromptDataJSON    `json:"-"`
}

// promptDeleteResponsePromptDataJSON contains the JSON metadata for the struct
// [PromptDeleteResponsePromptData]
type promptDeleteResponsePromptDataJSON struct {
	Options     apijson.Field
	Origin      apijson.Field
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDeleteResponsePromptData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDeleteResponsePromptDataJSON) RawJSON() string {
	return r.raw
}

type PromptDeleteResponsePromptDataOptions struct {
	Model    string                                      `json:"model"`
	Params   PromptDeleteResponsePromptDataOptionsParams `json:"params"`
	Position string                                      `json:"position"`
	JSON     promptDeleteResponsePromptDataOptionsJSON   `json:"-"`
}

// promptDeleteResponsePromptDataOptionsJSON contains the JSON metadata for the
// struct [PromptDeleteResponsePromptDataOptions]
type promptDeleteResponsePromptDataOptionsJSON struct {
	Model       apijson.Field
	Params      apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDeleteResponsePromptDataOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDeleteResponsePromptDataOptionsJSON) RawJSON() string {
	return r.raw
}

type PromptDeleteResponsePromptDataOptionsParams struct {
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
	MaxTokensToSample float64                                         `json:"max_tokens_to_sample"`
	MaxOutputTokens   float64                                         `json:"maxOutputTokens"`
	TopP              float64                                         `json:"topP"`
	TopK              float64                                         `json:"topK"`
	JSON              promptDeleteResponsePromptDataOptionsParamsJSON `json:"-"`
	union             PromptDeleteResponsePromptDataOptionsParamsUnion
}

// promptDeleteResponsePromptDataOptionsParamsJSON contains the JSON metadata for
// the struct [PromptDeleteResponsePromptDataOptionsParams]
type promptDeleteResponsePromptDataOptionsParamsJSON struct {
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

func (r promptDeleteResponsePromptDataOptionsParamsJSON) RawJSON() string {
	return r.raw
}

func (r *PromptDeleteResponsePromptDataOptionsParams) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PromptDeleteResponsePromptDataOptionsParams) AsUnion() PromptDeleteResponsePromptDataOptionsParamsUnion {
	return r.union
}

// Union satisfied by [PromptDeleteResponsePromptDataOptionsParamsObject],
// [PromptDeleteResponsePromptDataOptionsParamsObject],
// [PromptDeleteResponsePromptDataOptionsParamsObject] or
// [PromptDeleteResponsePromptDataOptionsParamsObject].
type PromptDeleteResponsePromptDataOptionsParamsUnion interface {
	implementsPromptDeleteResponsePromptDataOptionsParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptDeleteResponsePromptDataOptionsParamsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDeleteResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDeleteResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDeleteResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDeleteResponsePromptDataOptionsParamsObject{}),
		},
	)
}

type PromptDeleteResponsePromptDataOptionsParamsObject struct {
	FrequencyPenalty float64                                                            `json:"frequency_penalty"`
	FunctionCall     PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallUnion `json:"function_call"`
	MaxTokens        float64                                                            `json:"max_tokens"`
	N                float64                                                            `json:"n"`
	PresencePenalty  float64                                                            `json:"presence_penalty"`
	ResponseFormat   PromptDeleteResponsePromptDataOptionsParamsObjectResponseFormat    `json:"response_format,nullable"`
	Stop             []string                                                           `json:"stop"`
	Temperature      float64                                                            `json:"temperature"`
	ToolChoice       PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceUnion   `json:"tool_choice"`
	TopP             float64                                                            `json:"top_p"`
	UseCache         bool                                                               `json:"use_cache"`
	JSON             promptDeleteResponsePromptDataOptionsParamsObjectJSON              `json:"-"`
}

// promptDeleteResponsePromptDataOptionsParamsObjectJSON contains the JSON metadata
// for the struct [PromptDeleteResponsePromptDataOptionsParamsObject]
type promptDeleteResponsePromptDataOptionsParamsObjectJSON struct {
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

func (r *PromptDeleteResponsePromptDataOptionsParamsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDeleteResponsePromptDataOptionsParamsObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptDeleteResponsePromptDataOptionsParamsObject) implementsPromptDeleteResponsePromptDataOptionsParams() {
}

// Union satisfied by
// [PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallString],
// [PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallString] or
// [PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallObject].
type PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallUnion interface {
	ImplementsPromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallObject{}),
		},
	)
}

type PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallString string

const (
	PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallStringAuto PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallString = "auto"
)

func (r PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallString) IsKnown() bool {
	switch r {
	case PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallStringAuto:
		return true
	}
	return false
}

type PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallObject struct {
	Name string                                                                  `json:"name,required"`
	JSON promptDeleteResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON `json:"-"`
}

// promptDeleteResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON contains
// the JSON metadata for the struct
// [PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallObject]
type promptDeleteResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDeleteResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallObject) ImplementsPromptDeleteResponsePromptDataOptionsParamsObjectFunctionCallUnion() {
}

type PromptDeleteResponsePromptDataOptionsParamsObjectResponseFormat struct {
	Type PromptDeleteResponsePromptDataOptionsParamsObjectResponseFormatType `json:"type,required"`
	JSON promptDeleteResponsePromptDataOptionsParamsObjectResponseFormatJSON `json:"-"`
}

// promptDeleteResponsePromptDataOptionsParamsObjectResponseFormatJSON contains the
// JSON metadata for the struct
// [PromptDeleteResponsePromptDataOptionsParamsObjectResponseFormat]
type promptDeleteResponsePromptDataOptionsParamsObjectResponseFormatJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDeleteResponsePromptDataOptionsParamsObjectResponseFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDeleteResponsePromptDataOptionsParamsObjectResponseFormatJSON) RawJSON() string {
	return r.raw
}

type PromptDeleteResponsePromptDataOptionsParamsObjectResponseFormatType string

const (
	PromptDeleteResponsePromptDataOptionsParamsObjectResponseFormatTypeJsonObject PromptDeleteResponsePromptDataOptionsParamsObjectResponseFormatType = "json_object"
)

func (r PromptDeleteResponsePromptDataOptionsParamsObjectResponseFormatType) IsKnown() bool {
	switch r {
	case PromptDeleteResponsePromptDataOptionsParamsObjectResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Union satisfied by
// [PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceString],
// [PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceString] or
// [PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObject].
type PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceUnion interface {
	ImplementsPromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObject{}),
		},
	)
}

type PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceString string

const (
	PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceStringAuto PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceString = "auto"
)

func (r PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceString) IsKnown() bool {
	switch r {
	case PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceStringAuto:
		return true
	}
	return false
}

type PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObject struct {
	Function PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction `json:"function,required"`
	Type     PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectType     `json:"type,required"`
	JSON     promptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON     `json:"-"`
}

// promptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON contains
// the JSON metadata for the struct
// [PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObject]
type promptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON struct {
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObject) ImplementsPromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceUnion() {
}

type PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction struct {
	Name string                                                                        `json:"name,required"`
	JSON promptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON `json:"-"`
}

// promptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON
// contains the JSON metadata for the struct
// [PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction]
type promptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON) RawJSON() string {
	return r.raw
}

type PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectType string

const (
	PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectTypeFunction PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectType = "function"
)

func (r PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectType) IsKnown() bool {
	switch r {
	case PromptDeleteResponsePromptDataOptionsParamsObjectToolChoiceObjectTypeFunction:
		return true
	}
	return false
}

type PromptDeleteResponsePromptDataOrigin struct {
	ProjectID     string                                   `json:"project_id"`
	PromptID      string                                   `json:"prompt_id"`
	PromptVersion string                                   `json:"prompt_version"`
	JSON          promptDeleteResponsePromptDataOriginJSON `json:"-"`
}

// promptDeleteResponsePromptDataOriginJSON contains the JSON metadata for the
// struct [PromptDeleteResponsePromptDataOrigin]
type promptDeleteResponsePromptDataOriginJSON struct {
	ProjectID     apijson.Field
	PromptID      apijson.Field
	PromptVersion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PromptDeleteResponsePromptDataOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDeleteResponsePromptDataOriginJSON) RawJSON() string {
	return r.raw
}

type PromptDeleteResponsePromptDataPrompt struct {
	Type                  PromptDeleteResponsePromptDataPromptType `json:"type"`
	Content               string                                   `json:"content"`
	Messages              interface{}                              `json:"messages,required"`
	Tools                 string                                   `json:"tools"`
	ReservedOnlyAllowNull interface{}                              `json:"__reserved_only_allow_null,required"`
	JSON                  promptDeleteResponsePromptDataPromptJSON `json:"-"`
	union                 PromptDeleteResponsePromptDataPromptUnion
}

// promptDeleteResponsePromptDataPromptJSON contains the JSON metadata for the
// struct [PromptDeleteResponsePromptDataPrompt]
type promptDeleteResponsePromptDataPromptJSON struct {
	Type                  apijson.Field
	Content               apijson.Field
	Messages              apijson.Field
	Tools                 apijson.Field
	ReservedOnlyAllowNull apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r promptDeleteResponsePromptDataPromptJSON) RawJSON() string {
	return r.raw
}

func (r *PromptDeleteResponsePromptDataPrompt) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PromptDeleteResponsePromptDataPrompt) AsUnion() PromptDeleteResponsePromptDataPromptUnion {
	return r.union
}

// Union satisfied by [PromptDeleteResponsePromptDataPromptObject],
// [PromptDeleteResponsePromptDataPromptObject] or
// [PromptDeleteResponsePromptDataPromptObject].
type PromptDeleteResponsePromptDataPromptUnion interface {
	implementsPromptDeleteResponsePromptDataPrompt()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptDeleteResponsePromptDataPromptUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDeleteResponsePromptDataPromptObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDeleteResponsePromptDataPromptObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDeleteResponsePromptDataPromptObject{}),
		},
	)
}

type PromptDeleteResponsePromptDataPromptObject struct {
	Content string                                         `json:"content,required"`
	Type    PromptDeleteResponsePromptDataPromptObjectType `json:"type,required"`
	JSON    promptDeleteResponsePromptDataPromptObjectJSON `json:"-"`
}

// promptDeleteResponsePromptDataPromptObjectJSON contains the JSON metadata for
// the struct [PromptDeleteResponsePromptDataPromptObject]
type promptDeleteResponsePromptDataPromptObjectJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDeleteResponsePromptDataPromptObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDeleteResponsePromptDataPromptObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptDeleteResponsePromptDataPromptObject) implementsPromptDeleteResponsePromptDataPrompt() {
}

type PromptDeleteResponsePromptDataPromptObjectType string

const (
	PromptDeleteResponsePromptDataPromptObjectTypeCompletion PromptDeleteResponsePromptDataPromptObjectType = "completion"
)

func (r PromptDeleteResponsePromptDataPromptObjectType) IsKnown() bool {
	switch r {
	case PromptDeleteResponsePromptDataPromptObjectTypeCompletion:
		return true
	}
	return false
}

type PromptDeleteResponsePromptDataPromptType string

const (
	PromptDeleteResponsePromptDataPromptTypeCompletion PromptDeleteResponsePromptDataPromptType = "completion"
	PromptDeleteResponsePromptDataPromptTypeChat       PromptDeleteResponsePromptDataPromptType = "chat"
)

func (r PromptDeleteResponsePromptDataPromptType) IsKnown() bool {
	switch r {
	case PromptDeleteResponsePromptDataPromptTypeCompletion, PromptDeleteResponsePromptDataPromptTypeChat:
		return true
	}
	return false
}

type PromptReplaceResponse struct {
	// Unique identifier for the prompt
	ID string `json:"id,required" format:"uuid"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the prompt (see the `version`
	// parameter)
	XactID string `json:"_xact_id,required"`
	// A literal 'p' which identifies the object as a project prompt
	LogID PromptReplaceResponseLogID `json:"log_id,required"`
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
	PromptData PromptReplaceResponsePromptData `json:"prompt_data,nullable"`
	// A list of tags for the prompt
	Tags []string                  `json:"tags,nullable"`
	JSON promptReplaceResponseJSON `json:"-"`
}

// promptReplaceResponseJSON contains the JSON metadata for the struct
// [PromptReplaceResponse]
type promptReplaceResponseJSON struct {
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

func (r *PromptReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptReplaceResponseJSON) RawJSON() string {
	return r.raw
}

// A literal 'p' which identifies the object as a project prompt
type PromptReplaceResponseLogID string

const (
	PromptReplaceResponseLogIDP PromptReplaceResponseLogID = "p"
)

func (r PromptReplaceResponseLogID) IsKnown() bool {
	switch r {
	case PromptReplaceResponseLogIDP:
		return true
	}
	return false
}

// The prompt, model, and its parameters
type PromptReplaceResponsePromptData struct {
	Options PromptReplaceResponsePromptDataOptions `json:"options,nullable"`
	Origin  PromptReplaceResponsePromptDataOrigin  `json:"origin,nullable"`
	Prompt  PromptReplaceResponsePromptDataPrompt  `json:"prompt"`
	JSON    promptReplaceResponsePromptDataJSON    `json:"-"`
}

// promptReplaceResponsePromptDataJSON contains the JSON metadata for the struct
// [PromptReplaceResponsePromptData]
type promptReplaceResponsePromptDataJSON struct {
	Options     apijson.Field
	Origin      apijson.Field
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptReplaceResponsePromptData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptReplaceResponsePromptDataJSON) RawJSON() string {
	return r.raw
}

type PromptReplaceResponsePromptDataOptions struct {
	Model    string                                       `json:"model"`
	Params   PromptReplaceResponsePromptDataOptionsParams `json:"params"`
	Position string                                       `json:"position"`
	JSON     promptReplaceResponsePromptDataOptionsJSON   `json:"-"`
}

// promptReplaceResponsePromptDataOptionsJSON contains the JSON metadata for the
// struct [PromptReplaceResponsePromptDataOptions]
type promptReplaceResponsePromptDataOptionsJSON struct {
	Model       apijson.Field
	Params      apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptReplaceResponsePromptDataOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptReplaceResponsePromptDataOptionsJSON) RawJSON() string {
	return r.raw
}

type PromptReplaceResponsePromptDataOptionsParams struct {
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
	MaxTokensToSample float64                                          `json:"max_tokens_to_sample"`
	MaxOutputTokens   float64                                          `json:"maxOutputTokens"`
	TopP              float64                                          `json:"topP"`
	TopK              float64                                          `json:"topK"`
	JSON              promptReplaceResponsePromptDataOptionsParamsJSON `json:"-"`
	union             PromptReplaceResponsePromptDataOptionsParamsUnion
}

// promptReplaceResponsePromptDataOptionsParamsJSON contains the JSON metadata for
// the struct [PromptReplaceResponsePromptDataOptionsParams]
type promptReplaceResponsePromptDataOptionsParamsJSON struct {
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

func (r promptReplaceResponsePromptDataOptionsParamsJSON) RawJSON() string {
	return r.raw
}

func (r *PromptReplaceResponsePromptDataOptionsParams) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PromptReplaceResponsePromptDataOptionsParams) AsUnion() PromptReplaceResponsePromptDataOptionsParamsUnion {
	return r.union
}

// Union satisfied by [PromptReplaceResponsePromptDataOptionsParamsObject],
// [PromptReplaceResponsePromptDataOptionsParamsObject],
// [PromptReplaceResponsePromptDataOptionsParamsObject] or
// [PromptReplaceResponsePromptDataOptionsParamsObject].
type PromptReplaceResponsePromptDataOptionsParamsUnion interface {
	implementsPromptReplaceResponsePromptDataOptionsParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptReplaceResponsePromptDataOptionsParamsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptReplaceResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptReplaceResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptReplaceResponsePromptDataOptionsParamsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptReplaceResponsePromptDataOptionsParamsObject{}),
		},
	)
}

type PromptReplaceResponsePromptDataOptionsParamsObject struct {
	FrequencyPenalty float64                                                             `json:"frequency_penalty"`
	FunctionCall     PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallUnion `json:"function_call"`
	MaxTokens        float64                                                             `json:"max_tokens"`
	N                float64                                                             `json:"n"`
	PresencePenalty  float64                                                             `json:"presence_penalty"`
	ResponseFormat   PromptReplaceResponsePromptDataOptionsParamsObjectResponseFormat    `json:"response_format,nullable"`
	Stop             []string                                                            `json:"stop"`
	Temperature      float64                                                             `json:"temperature"`
	ToolChoice       PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceUnion   `json:"tool_choice"`
	TopP             float64                                                             `json:"top_p"`
	UseCache         bool                                                                `json:"use_cache"`
	JSON             promptReplaceResponsePromptDataOptionsParamsObjectJSON              `json:"-"`
}

// promptReplaceResponsePromptDataOptionsParamsObjectJSON contains the JSON
// metadata for the struct [PromptReplaceResponsePromptDataOptionsParamsObject]
type promptReplaceResponsePromptDataOptionsParamsObjectJSON struct {
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

func (r *PromptReplaceResponsePromptDataOptionsParamsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptReplaceResponsePromptDataOptionsParamsObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptReplaceResponsePromptDataOptionsParamsObject) implementsPromptReplaceResponsePromptDataOptionsParams() {
}

// Union satisfied by
// [PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallString],
// [PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallString] or
// [PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallObject].
type PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallUnion interface {
	ImplementsPromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallObject{}),
		},
	)
}

type PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallString string

const (
	PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallStringAuto PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallString = "auto"
)

func (r PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallString) IsKnown() bool {
	switch r {
	case PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallStringAuto:
		return true
	}
	return false
}

type PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallObject struct {
	Name string                                                                   `json:"name,required"`
	JSON promptReplaceResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON `json:"-"`
}

// promptReplaceResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON
// contains the JSON metadata for the struct
// [PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallObject]
type promptReplaceResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptReplaceResponsePromptDataOptionsParamsObjectFunctionCallObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallObject) ImplementsPromptReplaceResponsePromptDataOptionsParamsObjectFunctionCallUnion() {
}

type PromptReplaceResponsePromptDataOptionsParamsObjectResponseFormat struct {
	Type PromptReplaceResponsePromptDataOptionsParamsObjectResponseFormatType `json:"type,required"`
	JSON promptReplaceResponsePromptDataOptionsParamsObjectResponseFormatJSON `json:"-"`
}

// promptReplaceResponsePromptDataOptionsParamsObjectResponseFormatJSON contains
// the JSON metadata for the struct
// [PromptReplaceResponsePromptDataOptionsParamsObjectResponseFormat]
type promptReplaceResponsePromptDataOptionsParamsObjectResponseFormatJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptReplaceResponsePromptDataOptionsParamsObjectResponseFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptReplaceResponsePromptDataOptionsParamsObjectResponseFormatJSON) RawJSON() string {
	return r.raw
}

type PromptReplaceResponsePromptDataOptionsParamsObjectResponseFormatType string

const (
	PromptReplaceResponsePromptDataOptionsParamsObjectResponseFormatTypeJsonObject PromptReplaceResponsePromptDataOptionsParamsObjectResponseFormatType = "json_object"
)

func (r PromptReplaceResponsePromptDataOptionsParamsObjectResponseFormatType) IsKnown() bool {
	switch r {
	case PromptReplaceResponsePromptDataOptionsParamsObjectResponseFormatTypeJsonObject:
		return true
	}
	return false
}

// Union satisfied by
// [PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceString],
// [PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceString] or
// [PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObject].
type PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceUnion interface {
	ImplementsPromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObject{}),
		},
	)
}

type PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceString string

const (
	PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceStringAuto PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceString = "auto"
)

func (r PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceString) IsKnown() bool {
	switch r {
	case PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceStringAuto:
		return true
	}
	return false
}

type PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObject struct {
	Function PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction `json:"function,required"`
	Type     PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectType     `json:"type,required"`
	JSON     promptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON     `json:"-"`
}

// promptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON contains
// the JSON metadata for the struct
// [PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObject]
type promptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON struct {
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObject) ImplementsPromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceUnion() {
}

type PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction struct {
	Name string                                                                         `json:"name,required"`
	JSON promptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON `json:"-"`
}

// promptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON
// contains the JSON metadata for the struct
// [PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction]
type promptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectFunctionJSON) RawJSON() string {
	return r.raw
}

type PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectType string

const (
	PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectTypeFunction PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectType = "function"
)

func (r PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectType) IsKnown() bool {
	switch r {
	case PromptReplaceResponsePromptDataOptionsParamsObjectToolChoiceObjectTypeFunction:
		return true
	}
	return false
}

type PromptReplaceResponsePromptDataOrigin struct {
	ProjectID     string                                    `json:"project_id"`
	PromptID      string                                    `json:"prompt_id"`
	PromptVersion string                                    `json:"prompt_version"`
	JSON          promptReplaceResponsePromptDataOriginJSON `json:"-"`
}

// promptReplaceResponsePromptDataOriginJSON contains the JSON metadata for the
// struct [PromptReplaceResponsePromptDataOrigin]
type promptReplaceResponsePromptDataOriginJSON struct {
	ProjectID     apijson.Field
	PromptID      apijson.Field
	PromptVersion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PromptReplaceResponsePromptDataOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptReplaceResponsePromptDataOriginJSON) RawJSON() string {
	return r.raw
}

type PromptReplaceResponsePromptDataPrompt struct {
	Type                  PromptReplaceResponsePromptDataPromptType `json:"type"`
	Content               string                                    `json:"content"`
	Messages              interface{}                               `json:"messages,required"`
	Tools                 string                                    `json:"tools"`
	ReservedOnlyAllowNull interface{}                               `json:"__reserved_only_allow_null,required"`
	JSON                  promptReplaceResponsePromptDataPromptJSON `json:"-"`
	union                 PromptReplaceResponsePromptDataPromptUnion
}

// promptReplaceResponsePromptDataPromptJSON contains the JSON metadata for the
// struct [PromptReplaceResponsePromptDataPrompt]
type promptReplaceResponsePromptDataPromptJSON struct {
	Type                  apijson.Field
	Content               apijson.Field
	Messages              apijson.Field
	Tools                 apijson.Field
	ReservedOnlyAllowNull apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r promptReplaceResponsePromptDataPromptJSON) RawJSON() string {
	return r.raw
}

func (r *PromptReplaceResponsePromptDataPrompt) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PromptReplaceResponsePromptDataPrompt) AsUnion() PromptReplaceResponsePromptDataPromptUnion {
	return r.union
}

// Union satisfied by [PromptReplaceResponsePromptDataPromptObject],
// [PromptReplaceResponsePromptDataPromptObject] or
// [PromptReplaceResponsePromptDataPromptObject].
type PromptReplaceResponsePromptDataPromptUnion interface {
	implementsPromptReplaceResponsePromptDataPrompt()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptReplaceResponsePromptDataPromptUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptReplaceResponsePromptDataPromptObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptReplaceResponsePromptDataPromptObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptReplaceResponsePromptDataPromptObject{}),
		},
	)
}

type PromptReplaceResponsePromptDataPromptObject struct {
	Content string                                          `json:"content,required"`
	Type    PromptReplaceResponsePromptDataPromptObjectType `json:"type,required"`
	JSON    promptReplaceResponsePromptDataPromptObjectJSON `json:"-"`
}

// promptReplaceResponsePromptDataPromptObjectJSON contains the JSON metadata for
// the struct [PromptReplaceResponsePromptDataPromptObject]
type promptReplaceResponsePromptDataPromptObjectJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptReplaceResponsePromptDataPromptObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptReplaceResponsePromptDataPromptObjectJSON) RawJSON() string {
	return r.raw
}

func (r PromptReplaceResponsePromptDataPromptObject) implementsPromptReplaceResponsePromptDataPrompt() {
}

type PromptReplaceResponsePromptDataPromptObjectType string

const (
	PromptReplaceResponsePromptDataPromptObjectTypeCompletion PromptReplaceResponsePromptDataPromptObjectType = "completion"
)

func (r PromptReplaceResponsePromptDataPromptObjectType) IsKnown() bool {
	switch r {
	case PromptReplaceResponsePromptDataPromptObjectTypeCompletion:
		return true
	}
	return false
}

type PromptReplaceResponsePromptDataPromptType string

const (
	PromptReplaceResponsePromptDataPromptTypeCompletion PromptReplaceResponsePromptDataPromptType = "completion"
	PromptReplaceResponsePromptDataPromptTypeChat       PromptReplaceResponsePromptDataPromptType = "chat"
)

func (r PromptReplaceResponsePromptDataPromptType) IsKnown() bool {
	switch r {
	case PromptReplaceResponsePromptDataPromptTypeCompletion, PromptReplaceResponsePromptDataPromptTypeChat:
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
