// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/internal/param"
	"github.com/tidwall/gjson"
)

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

// Union satisfied by [shared.PromptDataOptionsParamsOpenAIModelParams],
// [shared.PromptDataOptionsParamsAnthropicModelParams],
// [shared.PromptDataOptionsParamsGoogleModelParams],
// [shared.PromptDataOptionsParamsWindowAIModelParams] or
// [shared.PromptDataOptionsParamsJsCompletionParams].
type PromptDataOptionsParamsUnion interface {
	implementsSharedPromptDataOptionsParamsUnion()
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

func (r PromptDataOptionsParamsOpenAIModelParams) implementsSharedPromptDataOptionsParamsUnion() {}

// Union satisfied by
// [shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto],
// [shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallNone] or
// [shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction].
type PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion interface {
	implementsSharedPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion()
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

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) implementsSharedPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto) implementsSharedPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam() {
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

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) implementsSharedPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
}

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallNone) implementsSharedPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam() {
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

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction) implementsSharedPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion() {
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

// Union satisfied by
// [shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto],
// [shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceNone] or
// [shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction].
type PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion interface {
	implementsSharedPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion()
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

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) implementsSharedPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto) implementsSharedPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam() {
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

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) implementsSharedPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
}

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceNone) implementsSharedPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam() {
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

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction) implementsSharedPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion() {
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

func (r PromptDataOptionsParamsAnthropicModelParams) implementsSharedPromptDataOptionsParamsUnion() {}

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

func (r PromptDataOptionsParamsGoogleModelParams) implementsSharedPromptDataOptionsParamsUnion() {}

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

func (r PromptDataOptionsParamsWindowAIModelParams) implementsSharedPromptDataOptionsParamsUnion() {}

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

func (r PromptDataOptionsParamsJsCompletionParams) implementsSharedPromptDataOptionsParamsUnion() {}

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
// Possible runtime types of the union are [shared.PromptDataPromptCompletion],
// [shared.PromptDataPromptChat], [shared.PromptDataPromptNullableVariant].
func (r PromptDataPrompt) AsUnion() PromptDataPromptUnion {
	return r.union
}

// Union satisfied by [shared.PromptDataPromptCompletion],
// [shared.PromptDataPromptChat] or [shared.PromptDataPromptNullableVariant].
type PromptDataPromptUnion interface {
	implementsSharedPromptDataPrompt()
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

func (r PromptDataPromptCompletion) implementsSharedPromptDataPrompt() {}

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

func (r PromptDataPromptChat) implementsSharedPromptDataPrompt() {}

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
// Possible runtime types of the union are
// [shared.PromptDataPromptChatMessagesSystem],
// [shared.PromptDataPromptChatMessagesUser],
// [shared.PromptDataPromptChatMessagesAssistant],
// [shared.PromptDataPromptChatMessagesTool],
// [shared.PromptDataPromptChatMessagesFunction],
// [shared.PromptDataPromptChatMessagesFallback].
func (r PromptDataPromptChatMessage) AsUnion() PromptDataPromptChatMessagesUnion {
	return r.union
}

// Union satisfied by [shared.PromptDataPromptChatMessagesSystem],
// [shared.PromptDataPromptChatMessagesUser],
// [shared.PromptDataPromptChatMessagesAssistant],
// [shared.PromptDataPromptChatMessagesTool],
// [shared.PromptDataPromptChatMessagesFunction] or
// [shared.PromptDataPromptChatMessagesFallback].
type PromptDataPromptChatMessagesUnion interface {
	implementsSharedPromptDataPromptChatMessage()
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

func (r PromptDataPromptChatMessagesSystem) implementsSharedPromptDataPromptChatMessage() {}

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

func (r PromptDataPromptChatMessagesUser) implementsSharedPromptDataPromptChatMessage() {}

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
// [shared.PromptDataPromptChatMessagesUserContentArray].
type PromptDataPromptChatMessagesUserContentUnion interface {
	ImplementsSharedPromptDataPromptChatMessagesUserContentUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptDataPromptChatMessagesUserContentUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataPromptChatMessagesUserContentArray{}),
		},
	)
}

type PromptDataPromptChatMessagesUserContentArray []PromptDataPromptChatMessagesUserContentArrayItem

func (r PromptDataPromptChatMessagesUserContentArray) ImplementsSharedPromptDataPromptChatMessagesUserContentUnion() {
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
// [shared.PromptDataPromptChatMessagesUserContentArrayText],
// [shared.PromptDataPromptChatMessagesUserContentArrayImageURL].
func (r PromptDataPromptChatMessagesUserContentArrayItem) AsUnion() PromptDataPromptChatMessagesUserContentArrayUnionItem {
	return r.union
}

// Union satisfied by [shared.PromptDataPromptChatMessagesUserContentArrayText] or
// [shared.PromptDataPromptChatMessagesUserContentArrayImageURL].
type PromptDataPromptChatMessagesUserContentArrayUnionItem interface {
	implementsSharedPromptDataPromptChatMessagesUserContentArrayItem()
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

func (r PromptDataPromptChatMessagesUserContentArrayText) implementsSharedPromptDataPromptChatMessagesUserContentArrayItem() {
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

func (r PromptDataPromptChatMessagesUserContentArrayImageURL) implementsSharedPromptDataPromptChatMessagesUserContentArrayItem() {
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

func (r PromptDataPromptChatMessagesAssistant) implementsSharedPromptDataPromptChatMessage() {}

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

func (r PromptDataPromptChatMessagesTool) implementsSharedPromptDataPromptChatMessage() {}

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

func (r PromptDataPromptChatMessagesFunction) implementsSharedPromptDataPromptChatMessage() {}

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

func (r PromptDataPromptChatMessagesFallback) implementsSharedPromptDataPromptChatMessage() {}

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

func (r PromptDataPromptNullableVariant) implementsSharedPromptDataPrompt() {}

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

// Satisfied by [shared.PromptDataOptionsParamsOpenAIModelParamsParam],
// [shared.PromptDataOptionsParamsAnthropicModelParamsParam],
// [shared.PromptDataOptionsParamsGoogleModelParamsParam],
// [shared.PromptDataOptionsParamsWindowAIModelParamsParam],
// [shared.PromptDataOptionsParamsJsCompletionParamsParam].
type PromptDataOptionsParamsUnionParam interface {
	implementsSharedPromptDataOptionsParamsUnionParam()
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

func (r PromptDataOptionsParamsOpenAIModelParamsParam) implementsSharedPromptDataOptionsParamsUnionParam() {
}

// Satisfied by [shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto],
// [shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallNone],
// [shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionParam].
type PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam interface {
	implementsSharedPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam()
}

type PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionParam struct {
	Name param.Field[string] `json:"name,required"`
}

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionParam) implementsSharedPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam() {
}

type PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam struct {
	Type param.Field[PromptDataOptionsParamsOpenAIModelParamsResponseFormatType] `json:"type,required"`
}

func (r PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto],
// [shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceNone],
// [shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionParam].
type PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam interface {
	implementsSharedPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam()
}

type PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionParam struct {
	Function param.Field[PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionParam] `json:"function,required"`
	Type     param.Field[PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType]          `json:"type,required"`
}

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionParam) implementsSharedPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam() {
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

func (r PromptDataOptionsParamsAnthropicModelParamsParam) implementsSharedPromptDataOptionsParamsUnionParam() {
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

func (r PromptDataOptionsParamsGoogleModelParamsParam) implementsSharedPromptDataOptionsParamsUnionParam() {
}

type PromptDataOptionsParamsWindowAIModelParamsParam struct {
	Temperature param.Field[float64] `json:"temperature"`
	TopK        param.Field[float64] `json:"topK"`
	UseCache    param.Field[bool]    `json:"use_cache"`
}

func (r PromptDataOptionsParamsWindowAIModelParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataOptionsParamsWindowAIModelParamsParam) implementsSharedPromptDataOptionsParamsUnionParam() {
}

type PromptDataOptionsParamsJsCompletionParamsParam struct {
	UseCache param.Field[bool] `json:"use_cache"`
}

func (r PromptDataOptionsParamsJsCompletionParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataOptionsParamsJsCompletionParamsParam) implementsSharedPromptDataOptionsParamsUnionParam() {
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

func (r PromptDataPromptParam) implementsSharedPromptDataPromptUnionParam() {}

// Satisfied by [shared.PromptDataPromptCompletionParam],
// [shared.PromptDataPromptChatParam],
// [shared.PromptDataPromptNullableVariantParam], [PromptDataPromptParam].
type PromptDataPromptUnionParam interface {
	implementsSharedPromptDataPromptUnionParam()
}

type PromptDataPromptCompletionParam struct {
	Content param.Field[string]                         `json:"content,required"`
	Type    param.Field[PromptDataPromptCompletionType] `json:"type,required"`
}

func (r PromptDataPromptCompletionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptCompletionParam) implementsSharedPromptDataPromptUnionParam() {}

type PromptDataPromptChatParam struct {
	Messages param.Field[[]PromptDataPromptChatMessagesUnionParam] `json:"messages,required"`
	Type     param.Field[PromptDataPromptChatType]                 `json:"type,required"`
	Tools    param.Field[string]                                   `json:"tools"`
}

func (r PromptDataPromptChatParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatParam) implementsSharedPromptDataPromptUnionParam() {}

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

func (r PromptDataPromptChatMessageParam) implementsSharedPromptDataPromptChatMessagesUnionParam() {}

// Satisfied by [shared.PromptDataPromptChatMessagesSystemParam],
// [shared.PromptDataPromptChatMessagesUserParam],
// [shared.PromptDataPromptChatMessagesAssistantParam],
// [shared.PromptDataPromptChatMessagesToolParam],
// [shared.PromptDataPromptChatMessagesFunctionParam],
// [shared.PromptDataPromptChatMessagesFallbackParam],
// [PromptDataPromptChatMessageParam].
type PromptDataPromptChatMessagesUnionParam interface {
	implementsSharedPromptDataPromptChatMessagesUnionParam()
}

type PromptDataPromptChatMessagesSystemParam struct {
	Role    param.Field[PromptDataPromptChatMessagesSystemRole] `json:"role,required"`
	Content param.Field[string]                                 `json:"content"`
	Name    param.Field[string]                                 `json:"name"`
}

func (r PromptDataPromptChatMessagesSystemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesSystemParam) implementsSharedPromptDataPromptChatMessagesUnionParam() {
}

type PromptDataPromptChatMessagesUserParam struct {
	Role    param.Field[PromptDataPromptChatMessagesUserRole]              `json:"role,required"`
	Content param.Field[PromptDataPromptChatMessagesUserContentUnionParam] `json:"content"`
	Name    param.Field[string]                                            `json:"name"`
}

func (r PromptDataPromptChatMessagesUserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesUserParam) implementsSharedPromptDataPromptChatMessagesUnionParam() {
}

// Satisfied by [shared.UnionString],
// [shared.PromptDataPromptChatMessagesUserContentArrayParam].
type PromptDataPromptChatMessagesUserContentUnionParam interface {
	ImplementsSharedPromptDataPromptChatMessagesUserContentUnionParam()
}

type PromptDataPromptChatMessagesUserContentArrayParam []PromptDataPromptChatMessagesUserContentArrayUnionItemParam

func (r PromptDataPromptChatMessagesUserContentArrayParam) ImplementsSharedPromptDataPromptChatMessagesUserContentUnionParam() {
}

type PromptDataPromptChatMessagesUserContentArrayItemParam struct {
	Text     param.Field[string]                                           `json:"text"`
	Type     param.Field[PromptDataPromptChatMessagesUserContentArrayType] `json:"type,required"`
	ImageURL param.Field[interface{}]                                      `json:"image_url,required"`
}

func (r PromptDataPromptChatMessagesUserContentArrayItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesUserContentArrayItemParam) implementsSharedPromptDataPromptChatMessagesUserContentArrayUnionItemParam() {
}

// Satisfied by [shared.PromptDataPromptChatMessagesUserContentArrayTextParam],
// [shared.PromptDataPromptChatMessagesUserContentArrayImageURLParam],
// [PromptDataPromptChatMessagesUserContentArrayItemParam].
type PromptDataPromptChatMessagesUserContentArrayUnionItemParam interface {
	implementsSharedPromptDataPromptChatMessagesUserContentArrayUnionItemParam()
}

type PromptDataPromptChatMessagesUserContentArrayTextParam struct {
	Type param.Field[PromptDataPromptChatMessagesUserContentArrayTextType] `json:"type,required"`
	Text param.Field[string]                                               `json:"text"`
}

func (r PromptDataPromptChatMessagesUserContentArrayTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesUserContentArrayTextParam) implementsSharedPromptDataPromptChatMessagesUserContentArrayUnionItemParam() {
}

type PromptDataPromptChatMessagesUserContentArrayImageURLParam struct {
	ImageURL param.Field[PromptDataPromptChatMessagesUserContentArrayImageURLImageURLParam] `json:"image_url,required"`
	Type     param.Field[PromptDataPromptChatMessagesUserContentArrayImageURLType]          `json:"type,required"`
}

func (r PromptDataPromptChatMessagesUserContentArrayImageURLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesUserContentArrayImageURLParam) implementsSharedPromptDataPromptChatMessagesUserContentArrayUnionItemParam() {
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

func (r PromptDataPromptChatMessagesAssistantParam) implementsSharedPromptDataPromptChatMessagesUnionParam() {
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

func (r PromptDataPromptChatMessagesToolParam) implementsSharedPromptDataPromptChatMessagesUnionParam() {
}

type PromptDataPromptChatMessagesFunctionParam struct {
	Name    param.Field[string]                                   `json:"name,required"`
	Role    param.Field[PromptDataPromptChatMessagesFunctionRole] `json:"role,required"`
	Content param.Field[string]                                   `json:"content"`
}

func (r PromptDataPromptChatMessagesFunctionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesFunctionParam) implementsSharedPromptDataPromptChatMessagesUnionParam() {
}

type PromptDataPromptChatMessagesFallbackParam struct {
	Role    param.Field[PromptDataPromptChatMessagesFallbackRole] `json:"role,required"`
	Content param.Field[string]                                   `json:"content"`
}

func (r PromptDataPromptChatMessagesFallbackParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptChatMessagesFallbackParam) implementsSharedPromptDataPromptChatMessagesUnionParam() {
}

type PromptDataPromptNullableVariantParam struct {
}

func (r PromptDataPromptNullableVariantParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataPromptNullableVariantParam) implementsSharedPromptDataPromptUnionParam() {}
