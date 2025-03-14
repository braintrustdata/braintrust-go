// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"github.com/braintrustdata/braintrust-go/internal/apierror"
	"github.com/braintrustdata/braintrust-go/shared"
)

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

// The object type that the ACL applies to
//
// This is an alias to an internal type.
type ACLObjectType = shared.ACLObjectType

// This is an alias to an internal value.
const ACLObjectTypeOrganization = shared.ACLObjectTypeOrganization

// This is an alias to an internal value.
const ACLObjectTypeProject = shared.ACLObjectTypeProject

// This is an alias to an internal value.
const ACLObjectTypeExperiment = shared.ACLObjectTypeExperiment

// This is an alias to an internal value.
const ACLObjectTypeDataset = shared.ACLObjectTypeDataset

// This is an alias to an internal value.
const ACLObjectTypePrompt = shared.ACLObjectTypePrompt

// This is an alias to an internal value.
const ACLObjectTypePromptSession = shared.ACLObjectTypePromptSession

// This is an alias to an internal value.
const ACLObjectTypeGroup = shared.ACLObjectTypeGroup

// This is an alias to an internal value.
const ACLObjectTypeRole = shared.ACLObjectTypeRole

// This is an alias to an internal value.
const ACLObjectTypeOrgMember = shared.ACLObjectTypeOrgMember

// This is an alias to an internal value.
const ACLObjectTypeProjectLog = shared.ACLObjectTypeProjectLog

// This is an alias to an internal value.
const ACLObjectTypeOrgProject = shared.ACLObjectTypeOrgProject

// Permission the ACL grants. Exactly one of `permission` and `role_id` will be
// provided
//
// This is an alias to an internal type.
type ACLPermission = shared.ACLPermission

// This is an alias to an internal value.
const ACLPermissionCreate = shared.ACLPermissionCreate

// This is an alias to an internal value.
const ACLPermissionRead = shared.ACLPermissionRead

// This is an alias to an internal value.
const ACLPermissionUpdate = shared.ACLPermissionUpdate

// This is an alias to an internal value.
const ACLPermissionDelete = shared.ACLPermissionDelete

// This is an alias to an internal value.
const ACLPermissionCreateACLs = shared.ACLPermissionCreateACLs

// This is an alias to an internal value.
const ACLPermissionReadACLs = shared.ACLPermissionReadACLs

// This is an alias to an internal value.
const ACLPermissionUpdateACLs = shared.ACLPermissionUpdateACLs

// This is an alias to an internal value.
const ACLPermissionDeleteACLs = shared.ACLPermissionDeleteACLs

// When setting a permission directly, optionally restricts the permission grant to
// just the specified object type. Cannot be set alongside a `role_id`.
//
// This is an alias to an internal type.
type ACLRestrictObjectType = shared.ACLRestrictObjectType

// This is an alias to an internal value.
const ACLRestrictObjectTypeOrganization = shared.ACLRestrictObjectTypeOrganization

// This is an alias to an internal value.
const ACLRestrictObjectTypeProject = shared.ACLRestrictObjectTypeProject

// This is an alias to an internal value.
const ACLRestrictObjectTypeExperiment = shared.ACLRestrictObjectTypeExperiment

// This is an alias to an internal value.
const ACLRestrictObjectTypeDataset = shared.ACLRestrictObjectTypeDataset

// This is an alias to an internal value.
const ACLRestrictObjectTypePrompt = shared.ACLRestrictObjectTypePrompt

// This is an alias to an internal value.
const ACLRestrictObjectTypePromptSession = shared.ACLRestrictObjectTypePromptSession

// This is an alias to an internal value.
const ACLRestrictObjectTypeGroup = shared.ACLRestrictObjectTypeGroup

// This is an alias to an internal value.
const ACLRestrictObjectTypeRole = shared.ACLRestrictObjectTypeRole

// This is an alias to an internal value.
const ACLRestrictObjectTypeOrgMember = shared.ACLRestrictObjectTypeOrgMember

// This is an alias to an internal value.
const ACLRestrictObjectTypeProjectLog = shared.ACLRestrictObjectTypeProjectLog

// This is an alias to an internal value.
const ACLRestrictObjectTypeOrgProject = shared.ACLRestrictObjectTypeOrgProject

// This is an alias to an internal type.
type APIKey = shared.APIKey

// This is an alias to an internal type.
type ChatCompletionContentPartImage = shared.ChatCompletionContentPartImage

// This is an alias to an internal type.
type ChatCompletionContentPartImageImageURL = shared.ChatCompletionContentPartImageImageURL

// This is an alias to an internal type.
type ChatCompletionContentPartImageImageURLDetail = shared.ChatCompletionContentPartImageImageURLDetail

// This is an alias to an internal value.
const ChatCompletionContentPartImageImageURLDetailAuto = shared.ChatCompletionContentPartImageImageURLDetailAuto

// This is an alias to an internal value.
const ChatCompletionContentPartImageImageURLDetailLow = shared.ChatCompletionContentPartImageImageURLDetailLow

// This is an alias to an internal value.
const ChatCompletionContentPartImageImageURLDetailHigh = shared.ChatCompletionContentPartImageImageURLDetailHigh

// This is an alias to an internal type.
type ChatCompletionContentPartImageType = shared.ChatCompletionContentPartImageType

// This is an alias to an internal value.
const ChatCompletionContentPartImageTypeImageURL = shared.ChatCompletionContentPartImageTypeImageURL

// This is an alias to an internal type.
type ChatCompletionContentPartImageParam = shared.ChatCompletionContentPartImageParam

// This is an alias to an internal type.
type ChatCompletionContentPartImageImageURLParam = shared.ChatCompletionContentPartImageImageURLParam

// This is an alias to an internal type.
type ChatCompletionContentPartText = shared.ChatCompletionContentPartText

// This is an alias to an internal type.
type ChatCompletionContentPartTextType = shared.ChatCompletionContentPartTextType

// This is an alias to an internal value.
const ChatCompletionContentPartTextTypeText = shared.ChatCompletionContentPartTextTypeText

// This is an alias to an internal type.
type ChatCompletionContentPartTextParam = shared.ChatCompletionContentPartTextParam

// This is an alias to an internal type.
type ChatCompletionMessageToolCall = shared.ChatCompletionMessageToolCall

// This is an alias to an internal type.
type ChatCompletionMessageToolCallFunction = shared.ChatCompletionMessageToolCallFunction

// This is an alias to an internal type.
type ChatCompletionMessageToolCallType = shared.ChatCompletionMessageToolCallType

// This is an alias to an internal value.
const ChatCompletionMessageToolCallTypeFunction = shared.ChatCompletionMessageToolCallTypeFunction

// This is an alias to an internal type.
type ChatCompletionMessageToolCallParam = shared.ChatCompletionMessageToolCallParam

// This is an alias to an internal type.
type ChatCompletionMessageToolCallFunctionParam = shared.ChatCompletionMessageToolCallFunctionParam

// This is an alias to an internal type.
type CodeBundle = shared.CodeBundle

// This is an alias to an internal type.
type CodeBundleLocation = shared.CodeBundleLocation

// This is an alias to an internal type.
type CodeBundleLocationExperiment = shared.CodeBundleLocationExperiment

// This is an alias to an internal type.
type CodeBundleLocationExperimentPosition = shared.CodeBundleLocationExperimentPosition

// This is an alias to an internal type.
type CodeBundleLocationExperimentPositionType = shared.CodeBundleLocationExperimentPositionType

// This is an alias to an internal type.
type CodeBundleLocationExperimentPositionTypeType = shared.CodeBundleLocationExperimentPositionTypeType

// This is an alias to an internal value.
const CodeBundleLocationExperimentPositionTypeTypeTask = shared.CodeBundleLocationExperimentPositionTypeTypeTask

// This is an alias to an internal type.
type CodeBundleLocationExperimentPositionScorer = shared.CodeBundleLocationExperimentPositionScorer

// This is an alias to an internal type.
type CodeBundleLocationExperimentPositionScorerType = shared.CodeBundleLocationExperimentPositionScorerType

// This is an alias to an internal value.
const CodeBundleLocationExperimentPositionScorerTypeScorer = shared.CodeBundleLocationExperimentPositionScorerTypeScorer

// This is an alias to an internal type.
type CodeBundleLocationExperimentType = shared.CodeBundleLocationExperimentType

// This is an alias to an internal value.
const CodeBundleLocationExperimentTypeExperiment = shared.CodeBundleLocationExperimentTypeExperiment

// This is an alias to an internal type.
type CodeBundleLocationFunction = shared.CodeBundleLocationFunction

// This is an alias to an internal type.
type CodeBundleLocationFunctionType = shared.CodeBundleLocationFunctionType

// This is an alias to an internal value.
const CodeBundleLocationFunctionTypeFunction = shared.CodeBundleLocationFunctionTypeFunction

// This is an alias to an internal type.
type CodeBundleLocationType = shared.CodeBundleLocationType

// This is an alias to an internal value.
const CodeBundleLocationTypeExperiment = shared.CodeBundleLocationTypeExperiment

// This is an alias to an internal value.
const CodeBundleLocationTypeFunction = shared.CodeBundleLocationTypeFunction

// This is an alias to an internal type.
type CodeBundleRuntimeContext = shared.CodeBundleRuntimeContext

// This is an alias to an internal type.
type CodeBundleRuntimeContextRuntime = shared.CodeBundleRuntimeContextRuntime

// This is an alias to an internal value.
const CodeBundleRuntimeContextRuntimeNode = shared.CodeBundleRuntimeContextRuntimeNode

// This is an alias to an internal value.
const CodeBundleRuntimeContextRuntimePython = shared.CodeBundleRuntimeContextRuntimePython

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

// This is an alias to an internal value.
const EnvVarObjectTypeOrganization = shared.EnvVarObjectTypeOrganization

// This is an alias to an internal value.
const EnvVarObjectTypeProject = shared.EnvVarObjectTypeProject

// This is an alias to an internal value.
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

// This is an alias to an internal value.
const FeedbackDatasetItemSourceApp = shared.FeedbackDatasetItemSourceApp

// This is an alias to an internal value.
const FeedbackDatasetItemSourceAPI = shared.FeedbackDatasetItemSourceAPI

// This is an alias to an internal value.
const FeedbackDatasetItemSourceExternal = shared.FeedbackDatasetItemSourceExternal

// This is an alias to an internal type.
type FeedbackExperimentItemParam = shared.FeedbackExperimentItemParam

// The source of the feedback. Must be one of "external" (default), "app", or "api"
//
// This is an alias to an internal type.
type FeedbackExperimentItemSource = shared.FeedbackExperimentItemSource

// This is an alias to an internal value.
const FeedbackExperimentItemSourceApp = shared.FeedbackExperimentItemSourceApp

// This is an alias to an internal value.
const FeedbackExperimentItemSourceAPI = shared.FeedbackExperimentItemSourceAPI

// This is an alias to an internal value.
const FeedbackExperimentItemSourceExternal = shared.FeedbackExperimentItemSourceExternal

// This is an alias to an internal type.
type FeedbackProjectLogsItemParam = shared.FeedbackProjectLogsItemParam

// The source of the feedback. Must be one of "external" (default), "app", or "api"
//
// This is an alias to an internal type.
type FeedbackProjectLogsItemSource = shared.FeedbackProjectLogsItemSource

// This is an alias to an internal value.
const FeedbackProjectLogsItemSourceApp = shared.FeedbackProjectLogsItemSourceApp

// This is an alias to an internal value.
const FeedbackProjectLogsItemSourceAPI = shared.FeedbackProjectLogsItemSourceAPI

// This is an alias to an internal value.
const FeedbackProjectLogsItemSourceExternal = shared.FeedbackProjectLogsItemSourceExternal

// This is an alias to an internal type.
type FeedbackResponseSchema = shared.FeedbackResponseSchema

// This is an alias to an internal type.
type FeedbackResponseSchemaStatus = shared.FeedbackResponseSchemaStatus

// This is an alias to an internal value.
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
type FunctionFunctionData = shared.FunctionFunctionData

// This is an alias to an internal type.
type FunctionFunctionDataPrompt = shared.FunctionFunctionDataPrompt

// This is an alias to an internal type.
type FunctionFunctionDataPromptType = shared.FunctionFunctionDataPromptType

// This is an alias to an internal value.
const FunctionFunctionDataPromptTypePrompt = shared.FunctionFunctionDataPromptTypePrompt

// This is an alias to an internal type.
type FunctionFunctionDataCode = shared.FunctionFunctionDataCode

// This is an alias to an internal type.
type FunctionFunctionDataCodeData = shared.FunctionFunctionDataCodeData

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataBundle = shared.FunctionFunctionDataCodeDataBundle

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataBundleType = shared.FunctionFunctionDataCodeDataBundleType

// This is an alias to an internal value.
const FunctionFunctionDataCodeDataBundleTypeBundle = shared.FunctionFunctionDataCodeDataBundleTypeBundle

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataInline = shared.FunctionFunctionDataCodeDataInline

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataInlineRuntimeContext = shared.FunctionFunctionDataCodeDataInlineRuntimeContext

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataInlineRuntimeContextRuntime = shared.FunctionFunctionDataCodeDataInlineRuntimeContextRuntime

// This is an alias to an internal value.
const FunctionFunctionDataCodeDataInlineRuntimeContextRuntimeNode = shared.FunctionFunctionDataCodeDataInlineRuntimeContextRuntimeNode

// This is an alias to an internal value.
const FunctionFunctionDataCodeDataInlineRuntimeContextRuntimePython = shared.FunctionFunctionDataCodeDataInlineRuntimeContextRuntimePython

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataInlineType = shared.FunctionFunctionDataCodeDataInlineType

// This is an alias to an internal value.
const FunctionFunctionDataCodeDataInlineTypeInline = shared.FunctionFunctionDataCodeDataInlineTypeInline

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataType = shared.FunctionFunctionDataCodeDataType

// This is an alias to an internal value.
const FunctionFunctionDataCodeDataTypeBundle = shared.FunctionFunctionDataCodeDataTypeBundle

// This is an alias to an internal value.
const FunctionFunctionDataCodeDataTypeInline = shared.FunctionFunctionDataCodeDataTypeInline

// This is an alias to an internal type.
type FunctionFunctionDataCodeType = shared.FunctionFunctionDataCodeType

// This is an alias to an internal value.
const FunctionFunctionDataCodeTypeCode = shared.FunctionFunctionDataCodeTypeCode

// This is an alias to an internal type.
type FunctionFunctionDataGlobal = shared.FunctionFunctionDataGlobal

// This is an alias to an internal type.
type FunctionFunctionDataGlobalType = shared.FunctionFunctionDataGlobalType

// This is an alias to an internal value.
const FunctionFunctionDataGlobalTypeGlobal = shared.FunctionFunctionDataGlobalTypeGlobal

// This is an alias to an internal type.
type FunctionFunctionDataType = shared.FunctionFunctionDataType

// This is an alias to an internal value.
const FunctionFunctionDataTypePrompt = shared.FunctionFunctionDataTypePrompt

// This is an alias to an internal value.
const FunctionFunctionDataTypeCode = shared.FunctionFunctionDataTypeCode

// This is an alias to an internal value.
const FunctionFunctionDataTypeGlobal = shared.FunctionFunctionDataTypeGlobal

// A literal 'p' which identifies the object as a project prompt
//
// This is an alias to an internal type.
type FunctionLogID = shared.FunctionLogID

// This is an alias to an internal value.
const FunctionLogIDP = shared.FunctionLogIDP

// JSON schema for the function's parameters and return type
//
// This is an alias to an internal type.
type FunctionFunctionSchema = shared.FunctionFunctionSchema

// This is an alias to an internal type.
type FunctionFunctionType = shared.FunctionFunctionType

// This is an alias to an internal value.
const FunctionFunctionTypeLlm = shared.FunctionFunctionTypeLlm

// This is an alias to an internal value.
const FunctionFunctionTypeScorer = shared.FunctionFunctionTypeScorer

// This is an alias to an internal value.
const FunctionFunctionTypeTask = shared.FunctionFunctionTypeTask

// This is an alias to an internal value.
const FunctionFunctionTypeTool = shared.FunctionFunctionTypeTool

// This is an alias to an internal type.
type FunctionOrigin = shared.FunctionOrigin

// The object type that the ACL applies to
//
// This is an alias to an internal type.
type FunctionOriginObjectType = shared.FunctionOriginObjectType

// This is an alias to an internal value.
const FunctionOriginObjectTypeOrganization = shared.FunctionOriginObjectTypeOrganization

// This is an alias to an internal value.
const FunctionOriginObjectTypeProject = shared.FunctionOriginObjectTypeProject

// This is an alias to an internal value.
const FunctionOriginObjectTypeExperiment = shared.FunctionOriginObjectTypeExperiment

// This is an alias to an internal value.
const FunctionOriginObjectTypeDataset = shared.FunctionOriginObjectTypeDataset

// This is an alias to an internal value.
const FunctionOriginObjectTypePrompt = shared.FunctionOriginObjectTypePrompt

// This is an alias to an internal value.
const FunctionOriginObjectTypePromptSession = shared.FunctionOriginObjectTypePromptSession

// This is an alias to an internal value.
const FunctionOriginObjectTypeGroup = shared.FunctionOriginObjectTypeGroup

// This is an alias to an internal value.
const FunctionOriginObjectTypeRole = shared.FunctionOriginObjectTypeRole

// This is an alias to an internal value.
const FunctionOriginObjectTypeOrgMember = shared.FunctionOriginObjectTypeOrgMember

// This is an alias to an internal value.
const FunctionOriginObjectTypeProjectLog = shared.FunctionOriginObjectTypeProjectLog

// This is an alias to an internal value.
const FunctionOriginObjectTypeOrgProject = shared.FunctionOriginObjectTypeOrgProject

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

// This is an alias to an internal value.
const ObjectReferenceObjectTypeExperiment = shared.ObjectReferenceObjectTypeExperiment

// This is an alias to an internal value.
const ObjectReferenceObjectTypeDataset = shared.ObjectReferenceObjectTypeDataset

// This is an alias to an internal value.
const ObjectReferenceObjectTypePrompt = shared.ObjectReferenceObjectTypePrompt

// This is an alias to an internal value.
const ObjectReferenceObjectTypeFunction = shared.ObjectReferenceObjectTypeFunction

// This is an alias to an internal value.
const ObjectReferenceObjectTypePromptSession = shared.ObjectReferenceObjectTypePromptSession

// This is an alias to an internal value.
const ObjectReferenceObjectTypeProjectLogs = shared.ObjectReferenceObjectTypeProjectLogs

// Indicates the event was copied from another object.
//
// This is an alias to an internal type.
type ObjectReferenceParam = shared.ObjectReferenceParam

// This is an alias to an internal type.
type OnlineScoreConfig = shared.OnlineScoreConfig

// This is an alias to an internal type.
type OnlineScoreConfigScorer = shared.OnlineScoreConfigScorer

// This is an alias to an internal type.
type OnlineScoreConfigScorersFunction = shared.OnlineScoreConfigScorersFunction

// This is an alias to an internal type.
type OnlineScoreConfigScorersFunctionType = shared.OnlineScoreConfigScorersFunctionType

// This is an alias to an internal value.
const OnlineScoreConfigScorersFunctionTypeFunction = shared.OnlineScoreConfigScorersFunctionTypeFunction

// This is an alias to an internal type.
type OnlineScoreConfigScorersGlobal = shared.OnlineScoreConfigScorersGlobal

// This is an alias to an internal type.
type OnlineScoreConfigScorersGlobalType = shared.OnlineScoreConfigScorersGlobalType

// This is an alias to an internal value.
const OnlineScoreConfigScorersGlobalTypeGlobal = shared.OnlineScoreConfigScorersGlobalTypeGlobal

// This is an alias to an internal type.
type OnlineScoreConfigScorersType = shared.OnlineScoreConfigScorersType

// This is an alias to an internal value.
const OnlineScoreConfigScorersTypeFunction = shared.OnlineScoreConfigScorersTypeFunction

// This is an alias to an internal value.
const OnlineScoreConfigScorersTypeGlobal = shared.OnlineScoreConfigScorersTypeGlobal

// This is an alias to an internal type.
type OnlineScoreConfigParam = shared.OnlineScoreConfigParam

// This is an alias to an internal type.
type OnlineScoreConfigScorersUnionParam = shared.OnlineScoreConfigScorersUnionParam

// This is an alias to an internal type.
type OnlineScoreConfigScorersFunctionParam = shared.OnlineScoreConfigScorersFunctionParam

// This is an alias to an internal type.
type OnlineScoreConfigScorersGlobalParam = shared.OnlineScoreConfigScorersGlobalParam

// This is an alias to an internal type.
type Organization = shared.Organization

// This is an alias to an internal type.
type PatchOrganizationMembersOutput = shared.PatchOrganizationMembersOutput

// This is an alias to an internal type.
type PatchOrganizationMembersOutputStatus = shared.PatchOrganizationMembersOutputStatus

// This is an alias to an internal value.
const PatchOrganizationMembersOutputStatusSuccess = shared.PatchOrganizationMembersOutputStatusSuccess

// This is an alias to an internal type.
type Project = shared.Project

// This is an alias to an internal type.
type ProjectLogsEvent = shared.ProjectLogsEvent

// A literal 'g' which identifies the log as a project log
//
// This is an alias to an internal type.
type ProjectLogsEventLogID = shared.ProjectLogsEventLogID

// This is an alias to an internal value.
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

// The type of the configured score
//
// This is an alias to an internal type.
type ProjectScoreScoreType = shared.ProjectScoreScoreType

// This is an alias to an internal value.
const ProjectScoreScoreTypeSlider = shared.ProjectScoreScoreTypeSlider

// This is an alias to an internal value.
const ProjectScoreScoreTypeCategorical = shared.ProjectScoreScoreTypeCategorical

// This is an alias to an internal value.
const ProjectScoreScoreTypeWeighted = shared.ProjectScoreScoreTypeWeighted

// This is an alias to an internal value.
const ProjectScoreScoreTypeMinimum = shared.ProjectScoreScoreTypeMinimum

// This is an alias to an internal value.
const ProjectScoreScoreTypeMaximum = shared.ProjectScoreScoreTypeMaximum

// This is an alias to an internal value.
const ProjectScoreScoreTypeOnline = shared.ProjectScoreScoreTypeOnline

// This is an alias to an internal value.
const ProjectScoreScoreTypeFreeForm = shared.ProjectScoreScoreTypeFreeForm

// For categorical-type project scores, the list of all categories
//
// This is an alias to an internal type.
type ProjectScoreCategoriesUnion = shared.ProjectScoreCategoriesUnion

// For categorical-type project scores, the list of all categories
//
// This is an alias to an internal type.
type ProjectScoreCategoriesCategorical = shared.ProjectScoreCategoriesCategorical

// For minimum-type project scores, the list of included scores
//
// This is an alias to an internal type.
type ProjectScoreCategoriesMinimum = shared.ProjectScoreCategoriesMinimum

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

// This is an alias to an internal type.
type ProjectSettings = shared.ProjectSettings

// This is an alias to an internal type.
type ProjectSettingsSpanFieldOrder = shared.ProjectSettingsSpanFieldOrder

// This is an alias to an internal type.
type ProjectSettingsSpanFieldOrderLayout = shared.ProjectSettingsSpanFieldOrderLayout

// This is an alias to an internal value.
const ProjectSettingsSpanFieldOrderLayoutFull = shared.ProjectSettingsSpanFieldOrderLayoutFull

// This is an alias to an internal value.
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

// This is an alias to an internal value.
const PromptLogIDP = shared.PromptLogIDP

// This is an alias to an internal type.
type PromptFunctionType = shared.PromptFunctionType

// This is an alias to an internal value.
const PromptFunctionTypeLlm = shared.PromptFunctionTypeLlm

// This is an alias to an internal value.
const PromptFunctionTypeScorer = shared.PromptFunctionTypeScorer

// This is an alias to an internal value.
const PromptFunctionTypeTask = shared.PromptFunctionTypeTask

// This is an alias to an internal value.
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
type PromptDataParserType = shared.PromptDataParserType

// This is an alias to an internal value.
const PromptDataParserTypeLlmClassifier = shared.PromptDataParserTypeLlmClassifier

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
type PromptDataPromptChatMessagesUserContentArrayUnionItem = shared.PromptDataPromptChatMessagesUserContentArrayUnionItem

// This is an alias to an internal type.
type PromptDataPromptChatMessagesAssistant = shared.PromptDataPromptChatMessagesAssistant

// This is an alias to an internal type.
type PromptDataPromptChatMessagesAssistantRole = shared.PromptDataPromptChatMessagesAssistantRole

// This is an alias to an internal value.
const PromptDataPromptChatMessagesAssistantRoleAssistant = shared.PromptDataPromptChatMessagesAssistantRoleAssistant

// This is an alias to an internal type.
type PromptDataPromptChatMessagesAssistantFunctionCall = shared.PromptDataPromptChatMessagesAssistantFunctionCall

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
type PromptDataPromptType = shared.PromptDataPromptType

// This is an alias to an internal value.
const PromptDataPromptTypeCompletion = shared.PromptDataPromptTypeCompletion

// This is an alias to an internal value.
const PromptDataPromptTypeChat = shared.PromptDataPromptTypeChat

// This is an alias to an internal type.
type PromptDataToolFunction = shared.PromptDataToolFunction

// This is an alias to an internal type.
type PromptDataToolFunctionsFunction = shared.PromptDataToolFunctionsFunction

// This is an alias to an internal type.
type PromptDataToolFunctionsFunctionType = shared.PromptDataToolFunctionsFunctionType

// This is an alias to an internal value.
const PromptDataToolFunctionsFunctionTypeFunction = shared.PromptDataToolFunctionsFunctionTypeFunction

// This is an alias to an internal type.
type PromptDataToolFunctionsGlobal = shared.PromptDataToolFunctionsGlobal

// This is an alias to an internal type.
type PromptDataToolFunctionsGlobalType = shared.PromptDataToolFunctionsGlobalType

// This is an alias to an internal value.
const PromptDataToolFunctionsGlobalTypeGlobal = shared.PromptDataToolFunctionsGlobalTypeGlobal

// This is an alias to an internal type.
type PromptDataToolFunctionsType = shared.PromptDataToolFunctionsType

// This is an alias to an internal value.
const PromptDataToolFunctionsTypeFunction = shared.PromptDataToolFunctionsTypeFunction

// This is an alias to an internal value.
const PromptDataToolFunctionsTypeGlobal = shared.PromptDataToolFunctionsTypeGlobal

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
type PromptDataPromptChatMessagesAssistantParam = shared.PromptDataPromptChatMessagesAssistantParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesAssistantFunctionCallParam = shared.PromptDataPromptChatMessagesAssistantFunctionCallParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesToolParam = shared.PromptDataPromptChatMessagesToolParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesFunctionParam = shared.PromptDataPromptChatMessagesFunctionParam

// This is an alias to an internal type.
type PromptDataPromptChatMessagesFallbackParam = shared.PromptDataPromptChatMessagesFallbackParam

// This is an alias to an internal type.
type PromptDataToolFunctionsUnionParam = shared.PromptDataToolFunctionsUnionParam

// This is an alias to an internal type.
type PromptDataToolFunctionsFunctionParam = shared.PromptDataToolFunctionsFunctionParam

// This is an alias to an internal type.
type PromptDataToolFunctionsGlobalParam = shared.PromptDataToolFunctionsGlobalParam

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

// This is an alias to an internal value.
const PromptOptionsParamsOpenAIModelParamsFunctionCallStringAuto = shared.PromptOptionsParamsOpenAIModelParamsFunctionCallStringAuto

// This is an alias to an internal value.
const PromptOptionsParamsOpenAIModelParamsFunctionCallStringNone = shared.PromptOptionsParamsOpenAIModelParamsFunctionCallStringNone

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsFunctionCallFunction = shared.PromptOptionsParamsOpenAIModelParamsFunctionCallFunction

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsReasoningEffort = shared.PromptOptionsParamsOpenAIModelParamsReasoningEffort

// This is an alias to an internal value.
const PromptOptionsParamsOpenAIModelParamsReasoningEffortLow = shared.PromptOptionsParamsOpenAIModelParamsReasoningEffortLow

// This is an alias to an internal value.
const PromptOptionsParamsOpenAIModelParamsReasoningEffortMedium = shared.PromptOptionsParamsOpenAIModelParamsReasoningEffortMedium

// This is an alias to an internal value.
const PromptOptionsParamsOpenAIModelParamsReasoningEffortHigh = shared.PromptOptionsParamsOpenAIModelParamsReasoningEffortHigh

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormat = shared.PromptOptionsParamsOpenAIModelParamsResponseFormat

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObject = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObject

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectType = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectType

// This is an alias to an internal value.
const PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectTypeJsonObject = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectTypeJsonObject

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchema = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchema

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchema = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchema

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaType = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaType

// This is an alias to an internal value.
const PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaTypeJsonSchema = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaTypeJsonSchema

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatText = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatText

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatTextType = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatTextType

// This is an alias to an internal value.
const PromptOptionsParamsOpenAIModelParamsResponseFormatTextTypeText = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatTextTypeText

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsResponseFormatType = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatType

// This is an alias to an internal value.
const PromptOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject

// This is an alias to an internal value.
const PromptOptionsParamsOpenAIModelParamsResponseFormatTypeJsonSchema = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatTypeJsonSchema

// This is an alias to an internal value.
const PromptOptionsParamsOpenAIModelParamsResponseFormatTypeText = shared.PromptOptionsParamsOpenAIModelParamsResponseFormatTypeText

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsToolChoiceUnion = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceUnion

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsToolChoiceString = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceString

// This is an alias to an internal value.
const PromptOptionsParamsOpenAIModelParamsToolChoiceStringAuto = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceStringAuto

// This is an alias to an internal value.
const PromptOptionsParamsOpenAIModelParamsToolChoiceStringNone = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceStringNone

// This is an alias to an internal value.
const PromptOptionsParamsOpenAIModelParamsToolChoiceStringRequired = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceStringRequired

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsToolChoiceFunction = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceFunction

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction

// This is an alias to an internal type.
type PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionType = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionType

// This is an alias to an internal value.
const PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction = shared.PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionTypeFunction

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

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
//
// This is an alias to an internal type.
type RoleMemberPermissionsPermission = shared.RoleMemberPermissionsPermission

// This is an alias to an internal value.
const RoleMemberPermissionsPermissionCreate = shared.RoleMemberPermissionsPermissionCreate

// This is an alias to an internal value.
const RoleMemberPermissionsPermissionRead = shared.RoleMemberPermissionsPermissionRead

// This is an alias to an internal value.
const RoleMemberPermissionsPermissionUpdate = shared.RoleMemberPermissionsPermissionUpdate

// This is an alias to an internal value.
const RoleMemberPermissionsPermissionDelete = shared.RoleMemberPermissionsPermissionDelete

// This is an alias to an internal value.
const RoleMemberPermissionsPermissionCreateACLs = shared.RoleMemberPermissionsPermissionCreateACLs

// This is an alias to an internal value.
const RoleMemberPermissionsPermissionReadACLs = shared.RoleMemberPermissionsPermissionReadACLs

// This is an alias to an internal value.
const RoleMemberPermissionsPermissionUpdateACLs = shared.RoleMemberPermissionsPermissionUpdateACLs

// This is an alias to an internal value.
const RoleMemberPermissionsPermissionDeleteACLs = shared.RoleMemberPermissionsPermissionDeleteACLs

// The object type that the ACL applies to
//
// This is an alias to an internal type.
type RoleMemberPermissionsRestrictObjectType = shared.RoleMemberPermissionsRestrictObjectType

// This is an alias to an internal value.
const RoleMemberPermissionsRestrictObjectTypeOrganization = shared.RoleMemberPermissionsRestrictObjectTypeOrganization

// This is an alias to an internal value.
const RoleMemberPermissionsRestrictObjectTypeProject = shared.RoleMemberPermissionsRestrictObjectTypeProject

// This is an alias to an internal value.
const RoleMemberPermissionsRestrictObjectTypeExperiment = shared.RoleMemberPermissionsRestrictObjectTypeExperiment

// This is an alias to an internal value.
const RoleMemberPermissionsRestrictObjectTypeDataset = shared.RoleMemberPermissionsRestrictObjectTypeDataset

// This is an alias to an internal value.
const RoleMemberPermissionsRestrictObjectTypePrompt = shared.RoleMemberPermissionsRestrictObjectTypePrompt

// This is an alias to an internal value.
const RoleMemberPermissionsRestrictObjectTypePromptSession = shared.RoleMemberPermissionsRestrictObjectTypePromptSession

// This is an alias to an internal value.
const RoleMemberPermissionsRestrictObjectTypeGroup = shared.RoleMemberPermissionsRestrictObjectTypeGroup

// This is an alias to an internal value.
const RoleMemberPermissionsRestrictObjectTypeRole = shared.RoleMemberPermissionsRestrictObjectTypeRole

// This is an alias to an internal value.
const RoleMemberPermissionsRestrictObjectTypeOrgMember = shared.RoleMemberPermissionsRestrictObjectTypeOrgMember

// This is an alias to an internal value.
const RoleMemberPermissionsRestrictObjectTypeProjectLog = shared.RoleMemberPermissionsRestrictObjectTypeProjectLog

// This is an alias to an internal value.
const RoleMemberPermissionsRestrictObjectTypeOrgProject = shared.RoleMemberPermissionsRestrictObjectTypeOrgProject

// Summary of a score's performance
//
// This is an alias to an internal type.
type ScoreSummary = shared.ScoreSummary

// Human-identifying attributes of the span, such as name, type, etc.
//
// This is an alias to an internal type.
type SpanAttributes = shared.SpanAttributes

// Type of the span, for display purposes only
//
// This is an alias to an internal type.
type SpanAttributesType = shared.SpanAttributesType

// This is an alias to an internal value.
const SpanAttributesTypeLlm = shared.SpanAttributesTypeLlm

// This is an alias to an internal value.
const SpanAttributesTypeScore = shared.SpanAttributesTypeScore

// This is an alias to an internal value.
const SpanAttributesTypeFunction = shared.SpanAttributesTypeFunction

// This is an alias to an internal value.
const SpanAttributesTypeEval = shared.SpanAttributesTypeEval

// This is an alias to an internal value.
const SpanAttributesTypeTask = shared.SpanAttributesTypeTask

// This is an alias to an internal value.
const SpanAttributesTypeTool = shared.SpanAttributesTypeTool

// Human-identifying attributes of the span, such as name, type, etc.
//
// This is an alias to an internal type.
type SpanAttributesParam = shared.SpanAttributesParam

// This is an alias to an internal type.
type SpanIFrame = shared.SpanIFrame

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

// The object type that the ACL applies to
//
// This is an alias to an internal type.
type ViewObjectType = shared.ViewObjectType

// This is an alias to an internal value.
const ViewObjectTypeOrganization = shared.ViewObjectTypeOrganization

// This is an alias to an internal value.
const ViewObjectTypeProject = shared.ViewObjectTypeProject

// This is an alias to an internal value.
const ViewObjectTypeExperiment = shared.ViewObjectTypeExperiment

// This is an alias to an internal value.
const ViewObjectTypeDataset = shared.ViewObjectTypeDataset

// This is an alias to an internal value.
const ViewObjectTypePrompt = shared.ViewObjectTypePrompt

// This is an alias to an internal value.
const ViewObjectTypePromptSession = shared.ViewObjectTypePromptSession

// This is an alias to an internal value.
const ViewObjectTypeGroup = shared.ViewObjectTypeGroup

// This is an alias to an internal value.
const ViewObjectTypeRole = shared.ViewObjectTypeRole

// This is an alias to an internal value.
const ViewObjectTypeOrgMember = shared.ViewObjectTypeOrgMember

// This is an alias to an internal value.
const ViewObjectTypeProjectLog = shared.ViewObjectTypeProjectLog

// This is an alias to an internal value.
const ViewObjectTypeOrgProject = shared.ViewObjectTypeOrgProject

// Type of table that the view corresponds to.
//
// This is an alias to an internal type.
type ViewViewType = shared.ViewViewType

// This is an alias to an internal value.
const ViewViewTypeProjects = shared.ViewViewTypeProjects

// This is an alias to an internal value.
const ViewViewTypeExperiments = shared.ViewViewTypeExperiments

// This is an alias to an internal value.
const ViewViewTypeExperiment = shared.ViewViewTypeExperiment

// This is an alias to an internal value.
const ViewViewTypePlaygrounds = shared.ViewViewTypePlaygrounds

// This is an alias to an internal value.
const ViewViewTypePlayground = shared.ViewViewTypePlayground

// This is an alias to an internal value.
const ViewViewTypeDatasets = shared.ViewViewTypeDatasets

// This is an alias to an internal value.
const ViewViewTypeDataset = shared.ViewViewTypeDataset

// This is an alias to an internal value.
const ViewViewTypePrompts = shared.ViewViewTypePrompts

// This is an alias to an internal value.
const ViewViewTypeTools = shared.ViewViewTypeTools

// This is an alias to an internal value.
const ViewViewTypeScorers = shared.ViewViewTypeScorers

// This is an alias to an internal value.
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
