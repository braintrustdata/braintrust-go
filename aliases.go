// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"github.com/braintrustdata/braintrust-go/internal/apierror"
	"github.com/braintrustdata/braintrust-go/shared"
)

type Error = apierror.Error

// The prompt, model, and its parameters
//
// This is an alias to an internal type.
type PromptData = shared.PromptData

// This is an alias to an internal type.
type PromptDataOptions = shared.PromptDataOptions

// This is an alias to an internal type.
type PromptDataOptionsParamsUnion = shared.PromptDataOptionsParamsUnion

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParams = shared.PromptDataOptionsParamsOpenAIModelParams

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion = shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto = shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto

// This is an alias to an internal value.
const PromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto = shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsFunctionCallNone = shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallNone

// This is an alias to an internal value.
const PromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone = shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallNoneNone

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction = shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunction

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsResponseFormat = shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormat

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsResponseFormatType = shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormatType

// This is an alias to an internal value.
const PromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject = shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion = shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto = shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto

// This is an alias to an internal value.
const PromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto = shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsToolChoiceNone = shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceNone

// This is an alias to an internal value.
const PromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone = shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceNoneNone

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction = shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunction

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction = shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType = shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionType

// This is an alias to an internal value.
const PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction = shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction

// This is an alias to an internal type.
type PromptDataOptionsParamsAnthropicModelParams = shared.PromptDataOptionsParamsAnthropicModelParams

// This is an alias to an internal type.
type PromptDataOptionsParamsGoogleModelParams = shared.PromptDataOptionsParamsGoogleModelParams

// This is an alias to an internal type.
type PromptDataOptionsParamsWindowAIModelParams = shared.PromptDataOptionsParamsWindowAIModelParams

// This is an alias to an internal type.
type PromptDataOptionsParamsJsCompletionParams = shared.PromptDataOptionsParamsJsCompletionParams

// This is an alias to an internal type.
type PromptDataOrigin = shared.PromptDataOrigin

// This is an alias to an internal type.
type PromptDataPrompt = shared.PromptDataPrompt

// This is an alias to an internal type.
type PromptDataPromptCompletion = shared.PromptDataPromptCompletion

// This is an alias to an internal type.
type PromptDataPromptCompletionType = shared.PromptDataPromptCompletionType

// This is an alias to an internal value.
const PromptDataPromptCompletionTypeCompletion = shared.PromptDataPromptCompletionTypeCompletion

// This is an alias to an internal type.
type PromptDataPromptChat = shared.PromptDataPromptChat

// This is an alias to an internal type.
type PromptDataPromptChatMessage = shared.PromptDataPromptChatMessage

// This is an alias to an internal type.
type PromptDataPromptChatMessagesSystem = shared.PromptDataPromptChatMessagesSystem

// This is an alias to an internal type.
type PromptDataPromptChatMessagesSystemRole = shared.PromptDataPromptChatMessagesSystemRole

// This is an alias to an internal value.
const PromptDataPromptChatMessagesSystemRoleSystem = shared.PromptDataPromptChatMessagesSystemRoleSystem

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUser = shared.PromptDataPromptChatMessagesUser

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserRole = shared.PromptDataPromptChatMessagesUserRole

// This is an alias to an internal value.
const PromptDataPromptChatMessagesUserRoleUser = shared.PromptDataPromptChatMessagesUserRoleUser

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentUnion = shared.PromptDataPromptChatMessagesUserContentUnion

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentArray = shared.PromptDataPromptChatMessagesUserContentArray

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentArrayItem = shared.PromptDataPromptChatMessagesUserContentArrayItem

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentArrayText = shared.PromptDataPromptChatMessagesUserContentArrayText

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentArrayTextType = shared.PromptDataPromptChatMessagesUserContentArrayTextType

// This is an alias to an internal value.
const PromptDataPromptChatMessagesUserContentArrayTextTypeText = shared.PromptDataPromptChatMessagesUserContentArrayTextTypeText

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentArrayImageURL = shared.PromptDataPromptChatMessagesUserContentArrayImageURL

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentArrayImageURLImageURL = shared.PromptDataPromptChatMessagesUserContentArrayImageURLImageURL

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail = shared.PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetail

// This is an alias to an internal value.
const PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto = shared.PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailAuto

// This is an alias to an internal value.
const PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow = shared.PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailLow

// This is an alias to an internal value.
const PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh = shared.PromptDataPromptChatMessagesUserContentArrayImageURLImageURLDetailHigh

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentArrayImageURLType = shared.PromptDataPromptChatMessagesUserContentArrayImageURLType

// This is an alias to an internal value.
const PromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL = shared.PromptDataPromptChatMessagesUserContentArrayImageURLTypeImageURL

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentArrayType = shared.PromptDataPromptChatMessagesUserContentArrayType

// This is an alias to an internal value.
const PromptDataPromptChatMessagesUserContentArrayTypeText = shared.PromptDataPromptChatMessagesUserContentArrayTypeText

// This is an alias to an internal value.
const PromptDataPromptChatMessagesUserContentArrayTypeImageURL = shared.PromptDataPromptChatMessagesUserContentArrayTypeImageURL

// This is an alias to an internal type.
type PromptDataPromptChatMessagesAssistant = shared.PromptDataPromptChatMessagesAssistant

// This is an alias to an internal type.
type PromptDataPromptChatMessagesAssistantRole = shared.PromptDataPromptChatMessagesAssistantRole

// This is an alias to an internal value.
const PromptDataPromptChatMessagesAssistantRoleAssistant = shared.PromptDataPromptChatMessagesAssistantRoleAssistant

// This is an alias to an internal type.
type PromptDataPromptChatMessagesAssistantFunctionCall = shared.PromptDataPromptChatMessagesAssistantFunctionCall

// This is an alias to an internal type.
type PromptDataPromptChatMessagesAssistantToolCall = shared.PromptDataPromptChatMessagesAssistantToolCall

// This is an alias to an internal type.
type PromptDataPromptChatMessagesAssistantToolCallsFunction = shared.PromptDataPromptChatMessagesAssistantToolCallsFunction

// This is an alias to an internal type.
type PromptDataPromptChatMessagesAssistantToolCallsType = shared.PromptDataPromptChatMessagesAssistantToolCallsType

// This is an alias to an internal value.
const PromptDataPromptChatMessagesAssistantToolCallsTypeFunction = shared.PromptDataPromptChatMessagesAssistantToolCallsTypeFunction

// This is an alias to an internal type.
type PromptDataPromptChatMessagesTool = shared.PromptDataPromptChatMessagesTool

// This is an alias to an internal type.
type PromptDataPromptChatMessagesToolRole = shared.PromptDataPromptChatMessagesToolRole

// This is an alias to an internal value.
const PromptDataPromptChatMessagesToolRoleTool = shared.PromptDataPromptChatMessagesToolRoleTool

// This is an alias to an internal type.
type PromptDataPromptChatMessagesFunction = shared.PromptDataPromptChatMessagesFunction

// This is an alias to an internal type.
type PromptDataPromptChatMessagesFunctionRole = shared.PromptDataPromptChatMessagesFunctionRole

// This is an alias to an internal value.
const PromptDataPromptChatMessagesFunctionRoleFunction = shared.PromptDataPromptChatMessagesFunctionRoleFunction

// This is an alias to an internal type.
type PromptDataPromptChatMessagesFallback = shared.PromptDataPromptChatMessagesFallback

// This is an alias to an internal type.
type PromptDataPromptChatMessagesFallbackRole = shared.PromptDataPromptChatMessagesFallbackRole

// This is an alias to an internal value.
const PromptDataPromptChatMessagesFallbackRoleModel = shared.PromptDataPromptChatMessagesFallbackRoleModel

// This is an alias to an internal type.
type PromptDataPromptChatMessagesRole = shared.PromptDataPromptChatMessagesRole

// This is an alias to an internal value.
const PromptDataPromptChatMessagesRoleSystem = shared.PromptDataPromptChatMessagesRoleSystem

// This is an alias to an internal value.
const PromptDataPromptChatMessagesRoleUser = shared.PromptDataPromptChatMessagesRoleUser

// This is an alias to an internal value.
const PromptDataPromptChatMessagesRoleAssistant = shared.PromptDataPromptChatMessagesRoleAssistant

// This is an alias to an internal value.
const PromptDataPromptChatMessagesRoleTool = shared.PromptDataPromptChatMessagesRoleTool

// This is an alias to an internal value.
const PromptDataPromptChatMessagesRoleFunction = shared.PromptDataPromptChatMessagesRoleFunction

// This is an alias to an internal value.
const PromptDataPromptChatMessagesRoleModel = shared.PromptDataPromptChatMessagesRoleModel

// This is an alias to an internal type.
type PromptDataPromptChatType = shared.PromptDataPromptChatType

// This is an alias to an internal value.
const PromptDataPromptChatTypeChat = shared.PromptDataPromptChatTypeChat

// This is an alias to an internal type.
type PromptDataPromptNullableVariant = shared.PromptDataPromptNullableVariant

// This is an alias to an internal type.
type PromptDataPromptType = shared.PromptDataPromptType

// This is an alias to an internal value.
const PromptDataPromptTypeCompletion = shared.PromptDataPromptTypeCompletion

// This is an alias to an internal value.
const PromptDataPromptTypeChat = shared.PromptDataPromptTypeChat

// The prompt, model, and its parameters
//
// This is an alias to an internal type.
type PromptDataParam = shared.PromptDataParam

// This is an alias to an internal type.
type PromptDataOptionsParam = shared.PromptDataOptionsParam

// This is an alias to an internal type.
type PromptDataOptionsParamsUnionParam = shared.PromptDataOptionsParamsUnionParam

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsParam = shared.PromptDataOptionsParamsOpenAIModelParamsParam

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam = shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionParam = shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallFunctionParam

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam = shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam = shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionParam = shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionParam

// This is an alias to an internal type.
type PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionParam = shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionParam

// This is an alias to an internal type.
type PromptDataOptionsParamsAnthropicModelParamsParam = shared.PromptDataOptionsParamsAnthropicModelParamsParam

// This is an alias to an internal type.
type PromptDataOptionsParamsGoogleModelParamsParam = shared.PromptDataOptionsParamsGoogleModelParamsParam

// This is an alias to an internal type.
type PromptDataOptionsParamsWindowAIModelParamsParam = shared.PromptDataOptionsParamsWindowAIModelParamsParam

// This is an alias to an internal type.
type PromptDataOptionsParamsJsCompletionParamsParam = shared.PromptDataOptionsParamsJsCompletionParamsParam

// This is an alias to an internal type.
type PromptDataOriginParam = shared.PromptDataOriginParam

// This is an alias to an internal type.
type PromptDataPromptUnionParam = shared.PromptDataPromptUnionParam

// This is an alias to an internal type.
type PromptDataPromptCompletionParam = shared.PromptDataPromptCompletionParam

// This is an alias to an internal type.
type PromptDataPromptChatParam = shared.PromptDataPromptChatParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUnionParam = shared.PromptDataPromptChatMessagesUnionParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesSystemParam = shared.PromptDataPromptChatMessagesSystemParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserParam = shared.PromptDataPromptChatMessagesUserParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentUnionParam = shared.PromptDataPromptChatMessagesUserContentUnionParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentArrayParam = shared.PromptDataPromptChatMessagesUserContentArrayParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentArrayUnionItemParam = shared.PromptDataPromptChatMessagesUserContentArrayUnionItemParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentArrayTextParam = shared.PromptDataPromptChatMessagesUserContentArrayTextParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentArrayImageURLParam = shared.PromptDataPromptChatMessagesUserContentArrayImageURLParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesUserContentArrayImageURLImageURLParam = shared.PromptDataPromptChatMessagesUserContentArrayImageURLImageURLParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesAssistantParam = shared.PromptDataPromptChatMessagesAssistantParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesAssistantFunctionCallParam = shared.PromptDataPromptChatMessagesAssistantFunctionCallParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesAssistantToolCallParam = shared.PromptDataPromptChatMessagesAssistantToolCallParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesAssistantToolCallsFunctionParam = shared.PromptDataPromptChatMessagesAssistantToolCallsFunctionParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesToolParam = shared.PromptDataPromptChatMessagesToolParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesFunctionParam = shared.PromptDataPromptChatMessagesFunctionParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesFallbackParam = shared.PromptDataPromptChatMessagesFallbackParam

// This is an alias to an internal type.
type PromptDataPromptNullableVariantParam = shared.PromptDataPromptNullableVariantParam
