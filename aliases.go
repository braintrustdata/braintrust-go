// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"github.com/braintrustdata/braintrust-go/internal/apierror"
	"github.com/braintrustdata/braintrust-go/packages/param"
	"github.com/braintrustdata/braintrust-go/packages/resp"
	"github.com/braintrustdata/braintrust-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type AISecret = shared.AISecret

// An ACL grants a certain permission or role to a certain user or group on an
// object.
//
// ACLs are inherited across the object hierarchy. So for example, if a user has
// read permissions on a project, they will also have read permissions on any
// experiment, dataset, etc. created within that project.
//
// To restrict a grant to a particular sub-object, you may specify
// `restrict_object_type` in the ACL, as part of a direct permission grant or as
// part of a role.
//
// This is an alias to an internal type.
type ACL = shared.ACL

// This is an alias to an internal type.
type ACLBatchUpdateResponse = shared.ACLBatchUpdateResponse

// The object type that the ACL applies to
//
// This is an alias to an internal type.
type ACLObjectType = shared.ACLObjectType

// Equals "organization"
const ACLObjectTypeOrganization = shared.ACLObjectTypeOrganization

// Equals "project"
const ACLObjectTypeProject = shared.ACLObjectTypeProject

// Equals "experiment"
const ACLObjectTypeExperiment = shared.ACLObjectTypeExperiment

// Equals "dataset"
const ACLObjectTypeDataset = shared.ACLObjectTypeDataset

// Equals "prompt"
const ACLObjectTypePrompt = shared.ACLObjectTypePrompt

// Equals "prompt_session"
const ACLObjectTypePromptSession = shared.ACLObjectTypePromptSession

// Equals "group"
const ACLObjectTypeGroup = shared.ACLObjectTypeGroup

// Equals "role"
const ACLObjectTypeRole = shared.ACLObjectTypeRole

// Equals "org_member"
const ACLObjectTypeOrgMember = shared.ACLObjectTypeOrgMember

// Equals "project_log"
const ACLObjectTypeProjectLog = shared.ACLObjectTypeProjectLog

// Equals "org_project"
const ACLObjectTypeOrgProject = shared.ACLObjectTypeOrgProject

// This is an alias to an internal type.
type APIKey = shared.APIKey

// This is an alias to an internal type.
type ChatCompletionContentPartImage = shared.ChatCompletionContentPartImage

// This is an alias to an internal type.
type ChatCompletionContentPartImageImageURL = shared.ChatCompletionContentPartImageImageURL

// This is an alias to an internal type.
type ChatCompletionContentPartImageImageURLDetail = shared.ChatCompletionContentPartImageImageURLDetail

// Equals "auto"
const ChatCompletionContentPartImageImageURLDetailAuto = shared.ChatCompletionContentPartImageImageURLDetailAuto

// Equals "low"
const ChatCompletionContentPartImageImageURLDetailLow = shared.ChatCompletionContentPartImageImageURLDetailLow

// Equals "high"
const ChatCompletionContentPartImageImageURLDetailHigh = shared.ChatCompletionContentPartImageImageURLDetailHigh

// This is an alias to an internal type.
type ChatCompletionContentPartImageType = shared.ChatCompletionContentPartImageType

// Equals "image_url"
const ChatCompletionContentPartImageTypeImageURL = shared.ChatCompletionContentPartImageTypeImageURL

// This is an alias to an internal type.
type ChatCompletionContentPartImageParam = shared.ChatCompletionContentPartImageParam

// This is an alias to an internal type.
type ChatCompletionContentPartImageImageURLParam = shared.ChatCompletionContentPartImageImageURLParam

// This is an alias to an internal type.
type ChatCompletionContentPartText = shared.ChatCompletionContentPartText

// This is an alias to an internal type.
type ChatCompletionContentPartTextType = shared.ChatCompletionContentPartTextType

// Equals "text"
const ChatCompletionContentPartTextTypeText = shared.ChatCompletionContentPartTextTypeText

// This is an alias to an internal type.
type ChatCompletionContentPartTextParam = shared.ChatCompletionContentPartTextParam

// This is an alias to an internal type.
type ChatCompletionMessageToolCall = shared.ChatCompletionMessageToolCall

// This is an alias to an internal type.
type ChatCompletionMessageToolCallFunction = shared.ChatCompletionMessageToolCallFunction

// This is an alias to an internal type.
type ChatCompletionMessageToolCallType = shared.ChatCompletionMessageToolCallType

// Equals "function"
const ChatCompletionMessageToolCallTypeFunction = shared.ChatCompletionMessageToolCallTypeFunction

// This is an alias to an internal type.
type ChatCompletionMessageToolCallParam = shared.ChatCompletionMessageToolCallParam

// This is an alias to an internal type.
type ChatCompletionMessageToolCallFunctionParam = shared.ChatCompletionMessageToolCallFunctionParam

// This is an alias to an internal type.
type CodeBundle = shared.CodeBundle

// This is an alias to an internal type.
type CodeBundleLocationUnion = shared.CodeBundleLocationUnion

// This is an alias to an internal type.
type CodeBundleLocationExperiment = shared.CodeBundleLocationExperiment

// This is an alias to an internal type.
type CodeBundleLocationExperimentPositionUnion = shared.CodeBundleLocationExperimentPositionUnion

// This is an alias to an internal type.
type CodeBundleLocationExperimentPositionType = shared.CodeBundleLocationExperimentPositionType

// This is an alias to an internal type.
type CodeBundleLocationExperimentPositionScorer = shared.CodeBundleLocationExperimentPositionScorer

// This is an alias to an internal type.
type CodeBundleLocationFunction = shared.CodeBundleLocationFunction

// This is an alias to an internal type.
type CodeBundleRuntimeContext = shared.CodeBundleRuntimeContext

// This is an alias to an internal type.
type CodeBundleParam = shared.CodeBundleParam

// This is an alias to an internal type.
type CodeBundleLocationUnionParam = shared.CodeBundleLocationUnionParam

// This is an alias to an internal type.
type CodeBundleLocationExperimentParam = shared.CodeBundleLocationExperimentParam

// This is an alias to an internal type.
type CodeBundleLocationExperimentPositionUnionParam = shared.CodeBundleLocationExperimentPositionUnionParam

// This is an alias to an internal type.
type CodeBundleLocationExperimentPositionTypeParam = shared.CodeBundleLocationExperimentPositionTypeParam

// This is an alias to an internal type.
type CodeBundleLocationExperimentPositionScorerParam = shared.CodeBundleLocationExperimentPositionScorerParam

// This is an alias to an internal type.
type CodeBundleLocationFunctionParam = shared.CodeBundleLocationFunctionParam

// This is an alias to an internal type.
type CodeBundleRuntimeContextParam = shared.CodeBundleRuntimeContextParam

// This is an alias to an internal type.
type CreateAPIKeyOutput = shared.CreateAPIKeyOutput

// Summary of a dataset's data
//
// This is an alias to an internal type.
type DataSummary = shared.DataSummary

// This is an alias to an internal type.
type Dataset = shared.Dataset

// This is an alias to an internal type.
type DatasetEvent = shared.DatasetEvent

// A dictionary with additional data about the test example, model outputs, or just
// about anything else that's relevant, that you can use to help find and analyze
// examples later. For example, you could log the `prompt`, example's `id`, or
// anything else that would be useful to slice/dice later. The values in `metadata`
// can be any JSON-serializable type, but its keys must be strings
//
// This is an alias to an internal type.
type DatasetEventMetadata = shared.DatasetEventMetadata

// This is an alias to an internal type.
type EnvVar = shared.EnvVar

// The type of the object the environment variable is scoped for
//
// This is an alias to an internal type.
type EnvVarObjectType = shared.EnvVarObjectType

// Equals "organization"
const EnvVarObjectTypeOrganization = shared.EnvVarObjectTypeOrganization

// Equals "project"
const EnvVarObjectTypeProject = shared.EnvVarObjectTypeProject

// Equals "function"
const EnvVarObjectTypeFunction = shared.EnvVarObjectTypeFunction

// This is an alias to an internal type.
type Experiment = shared.Experiment

// This is an alias to an internal type.
type ExperimentEvent = shared.ExperimentEvent

// Context is additional information about the code that produced the experiment
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the
// experiment event
//
// This is an alias to an internal type.
type ExperimentEventContext = shared.ExperimentEventContext

// A dictionary with additional data about the test example, model outputs, or just
// about anything else that's relevant, that you can use to help find and analyze
// examples later. For example, you could log the `prompt`, example's `id`, or
// anything else that would be useful to slice/dice later. The values in `metadata`
// can be any JSON-serializable type, but its keys must be strings
//
// This is an alias to an internal type.
type ExperimentEventMetadata = shared.ExperimentEventMetadata

// Metrics are numerical measurements tracking the execution of the code that
// produced the experiment event. Use "start" and "end" to track the time span over
// which the experiment event was produced
//
// This is an alias to an internal type.
type ExperimentEventMetrics = shared.ExperimentEventMetrics

// This is an alias to an internal type.
type FeedbackDatasetItemParam = shared.FeedbackDatasetItemParam

// The source of the feedback. Must be one of "external" (default), "app", or "api"
//
// This is an alias to an internal type.
type FeedbackDatasetItemSource = shared.FeedbackDatasetItemSource

// Equals "app"
const FeedbackDatasetItemSourceApp = shared.FeedbackDatasetItemSourceApp

// Equals "api"
const FeedbackDatasetItemSourceAPI = shared.FeedbackDatasetItemSourceAPI

// Equals "external"
const FeedbackDatasetItemSourceExternal = shared.FeedbackDatasetItemSourceExternal

// This is an alias to an internal type.
type FeedbackExperimentItemParam = shared.FeedbackExperimentItemParam

// The source of the feedback. Must be one of "external" (default), "app", or "api"
//
// This is an alias to an internal type.
type FeedbackExperimentItemSource = shared.FeedbackExperimentItemSource

// Equals "app"
const FeedbackExperimentItemSourceApp = shared.FeedbackExperimentItemSourceApp

// Equals "api"
const FeedbackExperimentItemSourceAPI = shared.FeedbackExperimentItemSourceAPI

// Equals "external"
const FeedbackExperimentItemSourceExternal = shared.FeedbackExperimentItemSourceExternal

// This is an alias to an internal type.
type FeedbackProjectLogsItemParam = shared.FeedbackProjectLogsItemParam

// The source of the feedback. Must be one of "external" (default), "app", or "api"
//
// This is an alias to an internal type.
type FeedbackProjectLogsItemSource = shared.FeedbackProjectLogsItemSource

// Equals "app"
const FeedbackProjectLogsItemSourceApp = shared.FeedbackProjectLogsItemSourceApp

// Equals "api"
const FeedbackProjectLogsItemSourceAPI = shared.FeedbackProjectLogsItemSourceAPI

// Equals "external"
const FeedbackProjectLogsItemSourceExternal = shared.FeedbackProjectLogsItemSourceExternal

// This is an alias to an internal type.
type FeedbackResponseSchema = shared.FeedbackResponseSchema

// This is an alias to an internal type.
type FeedbackResponseSchemaStatus = shared.FeedbackResponseSchemaStatus

// Equals "success"
const FeedbackResponseSchemaStatusSuccess = shared.FeedbackResponseSchemaStatusSuccess

// This is an alias to an internal type.
type FetchDatasetEventsResponse = shared.FetchDatasetEventsResponse

// This is an alias to an internal type.
type FetchExperimentEventsResponse = shared.FetchExperimentEventsResponse

// This is an alias to an internal type.
type FetchProjectLogsEventsResponse = shared.FetchProjectLogsEventsResponse

// This is an alias to an internal type.
type Function = shared.Function

// This is an alias to an internal type.
type FunctionFunctionDataUnion = shared.FunctionFunctionDataUnion

// This is an alias to an internal type.
type FunctionFunctionDataPrompt = shared.FunctionFunctionDataPrompt

// This is an alias to an internal type.
type FunctionFunctionDataCode = shared.FunctionFunctionDataCode

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataUnion = shared.FunctionFunctionDataCodeDataUnion

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataBundle = shared.FunctionFunctionDataCodeDataBundle

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataInline = shared.FunctionFunctionDataCodeDataInline

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataInlineRuntimeContext = shared.FunctionFunctionDataCodeDataInlineRuntimeContext

// This is an alias to an internal type.
type FunctionFunctionDataGlobal = shared.FunctionFunctionDataGlobal

// A literal 'p' which identifies the object as a project prompt
//
// This is an alias to an internal type.
type FunctionLogID = shared.FunctionLogID

// Equals "p"
const FunctionLogIDP = shared.FunctionLogIDP

// JSON schema for the function's parameters and return type
//
// This is an alias to an internal type.
type FunctionFunctionSchema = shared.FunctionFunctionSchema

// This is an alias to an internal type.
type FunctionFunctionType = shared.FunctionFunctionType

// Equals "llm"
const FunctionFunctionTypeLlm = shared.FunctionFunctionTypeLlm

// Equals "scorer"
const FunctionFunctionTypeScorer = shared.FunctionFunctionTypeScorer

// Equals "task"
const FunctionFunctionTypeTask = shared.FunctionFunctionTypeTask

// Equals "tool"
const FunctionFunctionTypeTool = shared.FunctionFunctionTypeTool

// This is an alias to an internal type.
type FunctionOrigin = shared.FunctionOrigin

// A group is a collection of users which can be assigned an ACL
//
// Groups can consist of individual users, as well as a set of groups they inherit
// from
//
// This is an alias to an internal type.
type Group = shared.Group

// A dataset event
//
// This is an alias to an internal type.
type InsertDatasetEventParam = shared.InsertDatasetEventParam

// A dictionary with additional data about the test example, model outputs, or just
// about anything else that's relevant, that you can use to help find and analyze
// examples later. For example, you could log the `prompt`, example's `id`, or
// anything else that would be useful to slice/dice later. The values in `metadata`
// can be any JSON-serializable type, but its keys must be strings
//
// This is an alias to an internal type.
type InsertDatasetEventMetadataParam = shared.InsertDatasetEventMetadataParam

// This is an alias to an internal type.
type InsertEventsResponse = shared.InsertEventsResponse

// An experiment event
//
// This is an alias to an internal type.
type InsertExperimentEventParam = shared.InsertExperimentEventParam

// Context is additional information about the code that produced the experiment
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the
// experiment event
//
// This is an alias to an internal type.
type InsertExperimentEventContextParam = shared.InsertExperimentEventContextParam

// A dictionary with additional data about the test example, model outputs, or just
// about anything else that's relevant, that you can use to help find and analyze
// examples later. For example, you could log the `prompt`, example's `id`, or
// anything else that would be useful to slice/dice later. The values in `metadata`
// can be any JSON-serializable type, but its keys must be strings
//
// This is an alias to an internal type.
type InsertExperimentEventMetadataParam = shared.InsertExperimentEventMetadataParam

// Metrics are numerical measurements tracking the execution of the code that
// produced the experiment event. Use "start" and "end" to track the time span over
// which the experiment event was produced
//
// This is an alias to an internal type.
type InsertExperimentEventMetricsParam = shared.InsertExperimentEventMetricsParam

// A project logs event
//
// This is an alias to an internal type.
type InsertProjectLogsEventParam = shared.InsertProjectLogsEventParam

// Context is additional information about the code that produced the project logs
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the project
// logs event
//
// This is an alias to an internal type.
type InsertProjectLogsEventContextParam = shared.InsertProjectLogsEventContextParam

// A dictionary with additional data about the test example, model outputs, or just
// about anything else that's relevant, that you can use to help find and analyze
// examples later. For example, you could log the `prompt`, example's `id`, or
// anything else that would be useful to slice/dice later. The values in `metadata`
// can be any JSON-serializable type, but its keys must be strings
//
// This is an alias to an internal type.
type InsertProjectLogsEventMetadataParam = shared.InsertProjectLogsEventMetadataParam

// Metrics are numerical measurements tracking the execution of the code that
// produced the project logs event. Use "start" and "end" to track the time span
// over which the project logs event was produced
//
// This is an alias to an internal type.
type InsertProjectLogsEventMetricsParam = shared.InsertProjectLogsEventMetricsParam

// Summary of a metric's performance
//
// This is an alias to an internal type.
type MetricSummary = shared.MetricSummary

// Indicates the event was copied from another object.
//
// This is an alias to an internal type.
type ObjectReference = shared.ObjectReference

// Type of the object the event is originating from.
//
// This is an alias to an internal type.
type ObjectReferenceObjectType = shared.ObjectReferenceObjectType

// Equals "experiment"
const ObjectReferenceObjectTypeExperiment = shared.ObjectReferenceObjectTypeExperiment

// Equals "dataset"
const ObjectReferenceObjectTypeDataset = shared.ObjectReferenceObjectTypeDataset

// Equals "prompt"
const ObjectReferenceObjectTypePrompt = shared.ObjectReferenceObjectTypePrompt

// Equals "function"
const ObjectReferenceObjectTypeFunction = shared.ObjectReferenceObjectTypeFunction

// Equals "prompt_session"
const ObjectReferenceObjectTypePromptSession = shared.ObjectReferenceObjectTypePromptSession

// Equals "project_logs"
const ObjectReferenceObjectTypeProjectLogs = shared.ObjectReferenceObjectTypeProjectLogs

// Indicates the event was copied from another object.
//
// This is an alias to an internal type.
type ObjectReferenceParam = shared.ObjectReferenceParam

// This is an alias to an internal type.
type OnlineScoreConfig = shared.OnlineScoreConfig

// This is an alias to an internal type.
type OnlineScoreConfigScorerUnion = shared.OnlineScoreConfigScorerUnion

// This is an alias to an internal type.
type OnlineScoreConfigScorerFunction = shared.OnlineScoreConfigScorerFunction

// This is an alias to an internal type.
type OnlineScoreConfigScorerGlobal = shared.OnlineScoreConfigScorerGlobal

// This is an alias to an internal type.
type OnlineScoreConfigParam = shared.OnlineScoreConfigParam

// This is an alias to an internal type.
type OnlineScoreConfigScorerUnionParam = shared.OnlineScoreConfigScorerUnionParam

// This is an alias to an internal type.
type OnlineScoreConfigScorerFunctionParam = shared.OnlineScoreConfigScorerFunctionParam

// This is an alias to an internal type.
type OnlineScoreConfigScorerGlobalParam = shared.OnlineScoreConfigScorerGlobalParam

// This is an alias to an internal type.
type Organization = shared.Organization

// This is an alias to an internal type.
type PatchOrganizationMembersOutput = shared.PatchOrganizationMembersOutput

// This is an alias to an internal type.
type PatchOrganizationMembersOutputStatus = shared.PatchOrganizationMembersOutputStatus

// Equals "success"
const PatchOrganizationMembersOutputStatusSuccess = shared.PatchOrganizationMembersOutputStatusSuccess

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
//
// This is an alias to an internal type.
type Permission = shared.Permission

// Equals "create"
const PermissionCreate = shared.PermissionCreate

// Equals "read"
const PermissionRead = shared.PermissionRead

// Equals "update"
const PermissionUpdate = shared.PermissionUpdate

// Equals "delete"
const PermissionDelete = shared.PermissionDelete

// Equals "create_acls"
const PermissionCreateACLs = shared.PermissionCreateACLs

// Equals "read_acls"
const PermissionReadACLs = shared.PermissionReadACLs

// Equals "update_acls"
const PermissionUpdateACLs = shared.PermissionUpdateACLs

// Equals "delete_acls"
const PermissionDeleteACLs = shared.PermissionDeleteACLs

// This is an alias to an internal type.
type Project = shared.Project

// This is an alias to an internal type.
type ProjectLogsEvent = shared.ProjectLogsEvent

// A literal 'g' which identifies the log as a project log
//
// This is an alias to an internal type.
type ProjectLogsEventLogID = shared.ProjectLogsEventLogID

// Equals "g"
const ProjectLogsEventLogIDG = shared.ProjectLogsEventLogIDG

// Context is additional information about the code that produced the project logs
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the project
// logs event
//
// This is an alias to an internal type.
type ProjectLogsEventContext = shared.ProjectLogsEventContext

// A dictionary with additional data about the test example, model outputs, or just
// about anything else that's relevant, that you can use to help find and analyze
// examples later. For example, you could log the `prompt`, example's `id`, or
// anything else that would be useful to slice/dice later. The values in `metadata`
// can be any JSON-serializable type, but its keys must be strings
//
// This is an alias to an internal type.
type ProjectLogsEventMetadata = shared.ProjectLogsEventMetadata

// Metrics are numerical measurements tracking the execution of the code that
// produced the project logs event. Use "start" and "end" to track the time span
// over which the project logs event was produced
//
// This is an alias to an internal type.
type ProjectLogsEventMetrics = shared.ProjectLogsEventMetrics

// A project score is a user-configured score, which can be manually-labeled
// through the UI
//
// This is an alias to an internal type.
type ProjectScore = shared.ProjectScore

// For categorical-type project scores, the list of all categories
//
// This is an alias to an internal type.
type ProjectScoreCategoriesUnion = shared.ProjectScoreCategoriesUnion

// For categorical-type project scores, defines a single category
//
// This is an alias to an internal type.
type ProjectScoreCategory = shared.ProjectScoreCategory

// For categorical-type project scores, defines a single category
//
// This is an alias to an internal type.
type ProjectScoreCategoryParam = shared.ProjectScoreCategoryParam

// This is an alias to an internal type.
type ProjectScoreConfig = shared.ProjectScoreConfig

// This is an alias to an internal type.
type ProjectScoreConfigParam = shared.ProjectScoreConfigParam

// The type of the configured score
//
// This is an alias to an internal type.
type ProjectScoreType = shared.ProjectScoreType

// Equals "slider"
const ProjectScoreTypeSlider = shared.ProjectScoreTypeSlider

// Equals "categorical"
const ProjectScoreTypeCategorical = shared.ProjectScoreTypeCategorical

// Equals "weighted"
const ProjectScoreTypeWeighted = shared.ProjectScoreTypeWeighted

// Equals "minimum"
const ProjectScoreTypeMinimum = shared.ProjectScoreTypeMinimum

// Equals "maximum"
const ProjectScoreTypeMaximum = shared.ProjectScoreTypeMaximum

// Equals "online"
const ProjectScoreTypeOnline = shared.ProjectScoreTypeOnline

// Equals "free-form"
const ProjectScoreTypeFreeForm = shared.ProjectScoreTypeFreeForm

// This is an alias to an internal type.
type ProjectSettings = shared.ProjectSettings

// This is an alias to an internal type.
type ProjectSettingsSpanFieldOrder = shared.ProjectSettingsSpanFieldOrder

// This is an alias to an internal type.
type ProjectSettingsSpanFieldOrderLayout = shared.ProjectSettingsSpanFieldOrderLayout

// Equals "full"
const ProjectSettingsSpanFieldOrderLayoutFull = shared.ProjectSettingsSpanFieldOrderLayoutFull

// Equals "two_column"
const ProjectSettingsSpanFieldOrderLayoutTwoColumn = shared.ProjectSettingsSpanFieldOrderLayoutTwoColumn

// This is an alias to an internal type.
type ProjectSettingsParam = shared.ProjectSettingsParam

// This is an alias to an internal type.
type ProjectSettingsSpanFieldOrderParam = shared.ProjectSettingsSpanFieldOrderParam

// A project tag is a user-configured tag for tracking and filtering your
// experiments, logs, and other data
//
// This is an alias to an internal type.
type ProjectTag = shared.ProjectTag

// This is an alias to an internal type.
type Prompt = shared.Prompt

// A literal 'p' which identifies the object as a project prompt
//
// This is an alias to an internal type.
type PromptLogID = shared.PromptLogID

// Equals "p"
const PromptLogIDP = shared.PromptLogIDP

// This is an alias to an internal type.
type PromptFunctionType = shared.PromptFunctionType

// Equals "llm"
const PromptFunctionTypeLlm = shared.PromptFunctionTypeLlm

// Equals "scorer"
const PromptFunctionTypeScorer = shared.PromptFunctionTypeScorer

// Equals "task"
const PromptFunctionTypeTask = shared.PromptFunctionTypeTask

// Equals "tool"
const PromptFunctionTypeTool = shared.PromptFunctionTypeTool

// The prompt, model, and its parameters
//
// This is an alias to an internal type.
type PromptData = shared.PromptData

// This is an alias to an internal type.
type PromptDataOrigin = shared.PromptDataOrigin

// This is an alias to an internal type.
type PromptDataParser = shared.PromptDataParser

// This is an alias to an internal type.
type PromptDataPromptUnion = shared.PromptDataPromptUnion

// This is an alias to an internal type.
type PromptDataPromptCompletion = shared.PromptDataPromptCompletion

// This is an alias to an internal type.
type PromptDataPromptChat = shared.PromptDataPromptChat

// This is an alias to an internal type.
type PromptDataPromptChatMessageUnion = shared.PromptDataPromptChatMessageUnion

// This is an alias to an internal type.
type PromptDataPromptChatMessageSystem = shared.PromptDataPromptChatMessageSystem

// This is an alias to an internal type.
type PromptDataPromptChatMessageUser = shared.PromptDataPromptChatMessageUser

// This is an alias to an internal type.
type PromptDataPromptChatMessageUserContentUnion = shared.PromptDataPromptChatMessageUserContentUnion

// This is an alias to an internal type.
type PromptDataPromptChatMessageUserContentArrayItemUnion = shared.PromptDataPromptChatMessageUserContentArrayItemUnion

// This is an alias to an internal type.
type PromptDataPromptChatMessageAssistant = shared.PromptDataPromptChatMessageAssistant

// This is an alias to an internal type.
type PromptDataPromptChatMessageAssistantFunctionCall = shared.PromptDataPromptChatMessageAssistantFunctionCall

// This is an alias to an internal type.
type PromptDataPromptChatMessageTool = shared.PromptDataPromptChatMessageTool

// This is an alias to an internal type.
type PromptDataPromptChatMessageFunction = shared.PromptDataPromptChatMessageFunction

// This is an alias to an internal type.
type PromptDataPromptChatMessageFallback = shared.PromptDataPromptChatMessageFallback

// This is an alias to an internal type.
type PromptDataToolFunctionUnion = shared.PromptDataToolFunctionUnion

// This is an alias to an internal type.
type PromptDataToolFunctionFunction = shared.PromptDataToolFunctionFunction

// This is an alias to an internal type.
type PromptDataToolFunctionGlobal = shared.PromptDataToolFunctionGlobal

// The prompt, model, and its parameters
//
// This is an alias to an internal type.
type PromptDataParam = shared.PromptDataParam

// This is an alias to an internal type.
type PromptDataOriginParam = shared.PromptDataOriginParam

// This is an alias to an internal type.
type PromptDataParserParam = shared.PromptDataParserParam

// This is an alias to an internal type.
type PromptDataPromptUnionParam = shared.PromptDataPromptUnionParam

// This is an alias to an internal type.
type PromptDataPromptCompletionParam = shared.PromptDataPromptCompletionParam

// This is an alias to an internal type.
type PromptDataPromptChatParam = shared.PromptDataPromptChatParam

// This is an alias to an internal type.
type PromptDataPromptChatMessageUnionParam = shared.PromptDataPromptChatMessageUnionParam

// This is an alias to an internal type.
type PromptDataPromptChatMessageSystemParam = shared.PromptDataPromptChatMessageSystemParam

// This is an alias to an internal type.
type PromptDataPromptChatMessageUserParam = shared.PromptDataPromptChatMessageUserParam

// This is an alias to an internal type.
type PromptDataPromptChatMessageUserContentUnionParam = shared.PromptDataPromptChatMessageUserContentUnionParam

// This is an alias to an internal type.
type PromptDataPromptChatMessageUserContentArrayItemUnionParam = shared.PromptDataPromptChatMessageUserContentArrayItemUnionParam

// This is an alias to an internal type.
type PromptDataPromptChatMessageAssistantParam = shared.PromptDataPromptChatMessageAssistantParam

// This is an alias to an internal type.
type PromptDataPromptChatMessageAssistantFunctionCallParam = shared.PromptDataPromptChatMessageAssistantFunctionCallParam

// This is an alias to an internal type.
type PromptDataPromptChatMessageToolParam = shared.PromptDataPromptChatMessageToolParam

// This is an alias to an internal type.
type PromptDataPromptChatMessageFunctionParam = shared.PromptDataPromptChatMessageFunctionParam

// This is an alias to an internal type.
type PromptDataPromptChatMessageFallbackParam = shared.PromptDataPromptChatMessageFallbackParam

// This is an alias to an internal type.
type PromptDataToolFunctionUnionParam = shared.PromptDataToolFunctionUnionParam

// This is an alias to an internal type.
type PromptDataToolFunctionFunctionParam = shared.PromptDataToolFunctionFunctionParam

// This is an alias to an internal type.
type PromptDataToolFunctionGlobalParam = shared.PromptDataToolFunctionGlobalParam

// This is an alias to an internal type.
type PromptOptions = shared.PromptOptions

// This is an alias to an internal type.
type PromptOptionsParamsUnion = shared.PromptOptionsParamsUnion

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParams = shared.PromptOptionsParamsOpenAIModelParams

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsFunctionCallUnion = shared.PromptOptionsParamsOpenAIModelParamsFunctionCallUnion

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsFunctionCallString = shared.PromptOptionsParamsOpenAIModelParamsFunctionCallString

// Equals "auto"
const PromptOptionsParamsOpenAIModelParamsFunctionCallStringAuto = shared.PromptOptionsParamsOpenAIModelParamsFunctionCallStringAuto

// Equals "none"
const PromptOptionsParamsOpenAIModelParamsFunctionCallStringNone = shared.PromptOptionsParamsOpenAIModelParamsFunctionCallStringNone

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsFunctionCallFunction = shared.PromptOptionsParamsOpenAIModelParamsFunctionCallFunction

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatUnion = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatUnion

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObject = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObject

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchema = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchema

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchema = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchema

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatText = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatText

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsToolChoiceUnion = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceUnion

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsToolChoiceString = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceString

// Equals "auto"
const PromptOptionsParamsOpenAIModelParamsToolChoiceStringAuto = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceStringAuto

// Equals "none"
const PromptOptionsParamsOpenAIModelParamsToolChoiceStringNone = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceStringNone

// Equals "required"
const PromptOptionsParamsOpenAIModelParamsToolChoiceStringRequired = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceStringRequired

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsToolChoiceFunction = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceFunction

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction

// This is an alias to an internal type.
type PromptOptionsParamsAnthropicModelParams = shared.PromptOptionsParamsAnthropicModelParams

// This is an alias to an internal type.
type PromptOptionsParamsGoogleModelParams = shared.PromptOptionsParamsGoogleModelParams

// This is an alias to an internal type.
type PromptOptionsParamsWindowAIModelParams = shared.PromptOptionsParamsWindowAIModelParams

// This is an alias to an internal type.
type PromptOptionsParamsJsCompletionParams = shared.PromptOptionsParamsJsCompletionParams

// This is an alias to an internal type.
type PromptOptionsParam = shared.PromptOptionsParam

// This is an alias to an internal type.
type PromptOptionsParamsUnionParam = shared.PromptOptionsParamsUnionParam

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsParam = shared.PromptOptionsParamsOpenAIModelParamsParam

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam = shared.PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsFunctionCallFunctionParam = shared.PromptOptionsParamsOpenAIModelParamsFunctionCallFunctionParam

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectParam = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectParam

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaParam = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaParam

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchemaParam = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchemaParam

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatTextParam = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatTextParam

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionParam = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionParam

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionParam = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionParam

// This is an alias to an internal type.
type PromptOptionsParamsAnthropicModelParamsParam = shared.PromptOptionsParamsAnthropicModelParamsParam

// This is an alias to an internal type.
type PromptOptionsParamsGoogleModelParamsParam = shared.PromptOptionsParamsGoogleModelParamsParam

// This is an alias to an internal type.
type PromptOptionsParamsWindowAIModelParamsParam = shared.PromptOptionsParamsWindowAIModelParamsParam

// This is an alias to an internal type.
type PromptOptionsParamsJsCompletionParamsParam = shared.PromptOptionsParamsJsCompletionParamsParam

// Metadata about the state of the repo when the experiment was created
//
// This is an alias to an internal type.
type RepoInfo = shared.RepoInfo

// Metadata about the state of the repo when the experiment was created
//
// This is an alias to an internal type.
type RepoInfoParam = shared.RepoInfoParam

// A role is a collection of permissions which can be granted as part of an ACL
//
// Roles can consist of individual permissions, as well as a set of roles they
// inherit from
//
// This is an alias to an internal type.
type Role = shared.Role

// This is an alias to an internal type.
type RoleMemberPermission = shared.RoleMemberPermission

// Summary of a score's performance
//
// This is an alias to an internal type.
type ScoreSummary = shared.ScoreSummary

// Human-identifying attributes of the span, such as name, type, etc.
//
// This is an alias to an internal type.
type SpanAttributes = shared.SpanAttributes

// Human-identifying attributes of the span, such as name, type, etc.
//
// This is an alias to an internal type.
type SpanAttributesParam = shared.SpanAttributesParam

// This is an alias to an internal type.
type SpanIFrame = shared.SpanIFrame

// Type of the span, for display purposes only
//
// This is an alias to an internal type.
type SpanType = shared.SpanType

// Equals "llm"
const SpanTypeLlm = shared.SpanTypeLlm

// Equals "score"
const SpanTypeScore = shared.SpanTypeScore

// Equals "function"
const SpanTypeFunction = shared.SpanTypeFunction

// Equals "eval"
const SpanTypeEval = shared.SpanTypeEval

// Equals "task"
const SpanTypeTask = shared.SpanTypeTask

// Equals "tool"
const SpanTypeTool = shared.SpanTypeTool

// Summary of a dataset
//
// This is an alias to an internal type.
type SummarizeDatasetResponse = shared.SummarizeDatasetResponse

// Summary of an experiment
//
// This is an alias to an internal type.
type SummarizeExperimentResponse = shared.SummarizeExperimentResponse

// This is an alias to an internal type.
type User = shared.User

// This is an alias to an internal type.
type View = shared.View

// Type of table that the view corresponds to.
//
// This is an alias to an internal type.
type ViewViewType = shared.ViewViewType

// Equals "projects"
const ViewViewTypeProjects = shared.ViewViewTypeProjects

// Equals "experiments"
const ViewViewTypeExperiments = shared.ViewViewTypeExperiments

// Equals "experiment"
const ViewViewTypeExperiment = shared.ViewViewTypeExperiment

// Equals "playgrounds"
const ViewViewTypePlaygrounds = shared.ViewViewTypePlaygrounds

// Equals "playground"
const ViewViewTypePlayground = shared.ViewViewTypePlayground

// Equals "datasets"
const ViewViewTypeDatasets = shared.ViewViewTypeDatasets

// Equals "dataset"
const ViewViewTypeDataset = shared.ViewViewTypeDataset

// Equals "prompts"
const ViewViewTypePrompts = shared.ViewViewTypePrompts

// Equals "tools"
const ViewViewTypeTools = shared.ViewViewTypeTools

// Equals "scorers"
const ViewViewTypeScorers = shared.ViewViewTypeScorers

// Equals "logs"
const ViewViewTypeLogs = shared.ViewViewTypeLogs

// The view definition
//
// This is an alias to an internal type.
type ViewData = shared.ViewData

// The view definition
//
// This is an alias to an internal type.
type ViewDataParam = shared.ViewDataParam

// This is an alias to an internal type.
type ViewDataSearch = shared.ViewDataSearch

// This is an alias to an internal type.
type ViewDataSearchParam = shared.ViewDataSearchParam

// Options for the view in the app
//
// This is an alias to an internal type.
type ViewOptions = shared.ViewOptions

// Options for the view in the app
//
// This is an alias to an internal type.
type ViewOptionsParam = shared.ViewOptionsParam

// Type of table that the view corresponds to.
//
// This is an alias to an internal type.
type ViewType = shared.ViewType

// Equals "projects"
const ViewTypeProjects = shared.ViewTypeProjects

// Equals "experiments"
const ViewTypeExperiments = shared.ViewTypeExperiments

// Equals "experiment"
const ViewTypeExperiment = shared.ViewTypeExperiment

// Equals "playgrounds"
const ViewTypePlaygrounds = shared.ViewTypePlaygrounds

// Equals "playground"
const ViewTypePlayground = shared.ViewTypePlayground

// Equals "datasets"
const ViewTypeDatasets = shared.ViewTypeDatasets

// Equals "dataset"
const ViewTypeDataset = shared.ViewTypeDataset

// Equals "prompts"
const ViewTypePrompts = shared.ViewTypePrompts

// Equals "tools"
const ViewTypeTools = shared.ViewTypeTools

// Equals "scorers"
const ViewTypeScorers = shared.ViewTypeScorers

// Equals "logs"
const ViewTypeLogs = shared.ViewTypeLogs

func toParam[T comparable](value T, meta resp.Field) param.Opt[T] {
	if meta.IsPresent() {
		return param.NewOpt(value)
	}
	if meta.IsExplicitNull() {
		return param.NullOpt[T]()
	}
	return param.Opt[T]{}
}
