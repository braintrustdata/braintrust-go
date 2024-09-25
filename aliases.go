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
type Code = shared.Code

// This is an alias to an internal type.
type CodeData = shared.CodeData

// This is an alias to an internal type.
type CodeDataBundle = shared.CodeDataBundle

// This is an alias to an internal type.
type CodeDataBundleLocation = shared.CodeDataBundleLocation

// This is an alias to an internal type.
type CodeDataBundleLocationExperiment = shared.CodeDataBundleLocationExperiment

// This is an alias to an internal type.
type CodeDataBundleLocationExperimentPosition = shared.CodeDataBundleLocationExperimentPosition

// This is an alias to an internal type.
type CodeDataBundleLocationExperimentPositionType = shared.CodeDataBundleLocationExperimentPositionType

// This is an alias to an internal value.
const CodeDataBundleLocationExperimentPositionTypeTask = shared.CodeDataBundleLocationExperimentPositionTypeTask

// This is an alias to an internal value.
const CodeDataBundleLocationExperimentPositionTypeScorer = shared.CodeDataBundleLocationExperimentPositionTypeScorer

// This is an alias to an internal type.
type CodeDataBundleLocationExperimentType = shared.CodeDataBundleLocationExperimentType

// This is an alias to an internal value.
const CodeDataBundleLocationExperimentTypeExperiment = shared.CodeDataBundleLocationExperimentTypeExperiment

// This is an alias to an internal type.
type CodeDataBundleLocationFunction = shared.CodeDataBundleLocationFunction

// This is an alias to an internal type.
type CodeDataBundleLocationFunctionType = shared.CodeDataBundleLocationFunctionType

// This is an alias to an internal value.
const CodeDataBundleLocationFunctionTypeFunction = shared.CodeDataBundleLocationFunctionTypeFunction

// This is an alias to an internal type.
type CodeDataBundleLocationType = shared.CodeDataBundleLocationType

// This is an alias to an internal value.
const CodeDataBundleLocationTypeExperiment = shared.CodeDataBundleLocationTypeExperiment

// This is an alias to an internal value.
const CodeDataBundleLocationTypeFunction = shared.CodeDataBundleLocationTypeFunction

// This is an alias to an internal type.
type CodeDataBundleRuntimeContext = shared.CodeDataBundleRuntimeContext

// This is an alias to an internal type.
type CodeDataBundleRuntimeContextRuntime = shared.CodeDataBundleRuntimeContextRuntime

// This is an alias to an internal value.
const CodeDataBundleRuntimeContextRuntimeNode = shared.CodeDataBundleRuntimeContextRuntimeNode

// This is an alias to an internal value.
const CodeDataBundleRuntimeContextRuntimePython = shared.CodeDataBundleRuntimeContextRuntimePython

// This is an alias to an internal type.
type CodeDataBundleType = shared.CodeDataBundleType

// This is an alias to an internal value.
const CodeDataBundleTypeBundle = shared.CodeDataBundleTypeBundle

// This is an alias to an internal type.
type CodeDataInline = shared.CodeDataInline

// This is an alias to an internal type.
type CodeDataInlineRuntimeContext = shared.CodeDataInlineRuntimeContext

// This is an alias to an internal type.
type CodeDataInlineRuntimeContextRuntime = shared.CodeDataInlineRuntimeContextRuntime

// This is an alias to an internal value.
const CodeDataInlineRuntimeContextRuntimeNode = shared.CodeDataInlineRuntimeContextRuntimeNode

// This is an alias to an internal value.
const CodeDataInlineRuntimeContextRuntimePython = shared.CodeDataInlineRuntimeContextRuntimePython

// This is an alias to an internal type.
type CodeDataInlineType = shared.CodeDataInlineType

// This is an alias to an internal value.
const CodeDataInlineTypeInline = shared.CodeDataInlineTypeInline

// This is an alias to an internal type.
type CodeDataType = shared.CodeDataType

// This is an alias to an internal value.
const CodeDataTypeBundle = shared.CodeDataTypeBundle

// This is an alias to an internal value.
const CodeDataTypeInline = shared.CodeDataTypeInline

// This is an alias to an internal type.
type CodeType = shared.CodeType

// This is an alias to an internal value.
const CodeTypeCode = shared.CodeTypeCode

// This is an alias to an internal type.
type CodeParam = shared.CodeParam

// This is an alias to an internal type.
type CodeDataUnionParam = shared.CodeDataUnionParam

// This is an alias to an internal type.
type CodeDataBundleParam = shared.CodeDataBundleParam

// This is an alias to an internal type.
type CodeDataBundleLocationUnionParam = shared.CodeDataBundleLocationUnionParam

// This is an alias to an internal type.
type CodeDataBundleLocationExperimentParam = shared.CodeDataBundleLocationExperimentParam

// This is an alias to an internal type.
type CodeDataBundleLocationExperimentPositionUnionParam = shared.CodeDataBundleLocationExperimentPositionUnionParam

// This is an alias to an internal type.
type CodeDataBundleLocationFunctionParam = shared.CodeDataBundleLocationFunctionParam

// This is an alias to an internal type.
type CodeDataBundleRuntimeContextParam = shared.CodeDataBundleRuntimeContextParam

// This is an alias to an internal type.
type CodeDataInlineParam = shared.CodeDataInlineParam

// This is an alias to an internal type.
type CodeDataInlineRuntimeContextParam = shared.CodeDataInlineRuntimeContextParam

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

// Metrics are numerical measurements tracking the execution of the code that
// produced the experiment event. Use "start" and "end" to track the time span over
// which the experiment event was produced
//
// This is an alias to an internal type.
type ExperimentEventMetrics = shared.ExperimentEventMetrics

// Human-identifying attributes of the span, such as name, type, etc.
//
// This is an alias to an internal type.
type ExperimentEventSpanAttributes = shared.ExperimentEventSpanAttributes

// Type of the span, for display purposes only
//
// This is an alias to an internal type.
type ExperimentEventSpanAttributesType = shared.ExperimentEventSpanAttributesType

// This is an alias to an internal value.
const ExperimentEventSpanAttributesTypeLlm = shared.ExperimentEventSpanAttributesTypeLlm

// This is an alias to an internal value.
const ExperimentEventSpanAttributesTypeScore = shared.ExperimentEventSpanAttributesTypeScore

// This is an alias to an internal value.
const ExperimentEventSpanAttributesTypeFunction = shared.ExperimentEventSpanAttributesTypeFunction

// This is an alias to an internal value.
const ExperimentEventSpanAttributesTypeEval = shared.ExperimentEventSpanAttributesTypeEval

// This is an alias to an internal value.
const ExperimentEventSpanAttributesTypeTask = shared.ExperimentEventSpanAttributesTypeTask

// This is an alias to an internal value.
const ExperimentEventSpanAttributesTypeTool = shared.ExperimentEventSpanAttributesTypeTool

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

// This is an alias to an internal type.
type InsertDatasetEventMergeParam = shared.InsertDatasetEventMergeParam

// This is an alias to an internal type.
type InsertDatasetEventReplaceParam = shared.InsertDatasetEventReplaceParam

// This is an alias to an internal type.
type InsertEventsResponse = shared.InsertEventsResponse

// This is an alias to an internal type.
type InsertExperimentEventMergeParam = shared.InsertExperimentEventMergeParam

// Context is additional information about the code that produced the experiment
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the
// experiment event
//
// This is an alias to an internal type.
type InsertExperimentEventMergeContextParam = shared.InsertExperimentEventMergeContextParam

// Metrics are numerical measurements tracking the execution of the code that
// produced the experiment event. Use "start" and "end" to track the time span over
// which the experiment event was produced
//
// This is an alias to an internal type.
type InsertExperimentEventMergeMetricsParam = shared.InsertExperimentEventMergeMetricsParam

// Human-identifying attributes of the span, such as name, type, etc.
//
// This is an alias to an internal type.
type InsertExperimentEventMergeSpanAttributesParam = shared.InsertExperimentEventMergeSpanAttributesParam

// Type of the span, for display purposes only
//
// This is an alias to an internal type.
type InsertExperimentEventMergeSpanAttributesType = shared.InsertExperimentEventMergeSpanAttributesType

// This is an alias to an internal value.
const InsertExperimentEventMergeSpanAttributesTypeLlm = shared.InsertExperimentEventMergeSpanAttributesTypeLlm

// This is an alias to an internal value.
const InsertExperimentEventMergeSpanAttributesTypeScore = shared.InsertExperimentEventMergeSpanAttributesTypeScore

// This is an alias to an internal value.
const InsertExperimentEventMergeSpanAttributesTypeFunction = shared.InsertExperimentEventMergeSpanAttributesTypeFunction

// This is an alias to an internal value.
const InsertExperimentEventMergeSpanAttributesTypeEval = shared.InsertExperimentEventMergeSpanAttributesTypeEval

// This is an alias to an internal value.
const InsertExperimentEventMergeSpanAttributesTypeTask = shared.InsertExperimentEventMergeSpanAttributesTypeTask

// This is an alias to an internal value.
const InsertExperimentEventMergeSpanAttributesTypeTool = shared.InsertExperimentEventMergeSpanAttributesTypeTool

// This is an alias to an internal type.
type InsertExperimentEventReplaceParam = shared.InsertExperimentEventReplaceParam

// Context is additional information about the code that produced the experiment
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the
// experiment event
//
// This is an alias to an internal type.
type InsertExperimentEventReplaceContextParam = shared.InsertExperimentEventReplaceContextParam

// Metrics are numerical measurements tracking the execution of the code that
// produced the experiment event. Use "start" and "end" to track the time span over
// which the experiment event was produced
//
// This is an alias to an internal type.
type InsertExperimentEventReplaceMetricsParam = shared.InsertExperimentEventReplaceMetricsParam

// Human-identifying attributes of the span, such as name, type, etc.
//
// This is an alias to an internal type.
type InsertExperimentEventReplaceSpanAttributesParam = shared.InsertExperimentEventReplaceSpanAttributesParam

// Type of the span, for display purposes only
//
// This is an alias to an internal type.
type InsertExperimentEventReplaceSpanAttributesType = shared.InsertExperimentEventReplaceSpanAttributesType

// This is an alias to an internal value.
const InsertExperimentEventReplaceSpanAttributesTypeLlm = shared.InsertExperimentEventReplaceSpanAttributesTypeLlm

// This is an alias to an internal value.
const InsertExperimentEventReplaceSpanAttributesTypeScore = shared.InsertExperimentEventReplaceSpanAttributesTypeScore

// This is an alias to an internal value.
const InsertExperimentEventReplaceSpanAttributesTypeFunction = shared.InsertExperimentEventReplaceSpanAttributesTypeFunction

// This is an alias to an internal value.
const InsertExperimentEventReplaceSpanAttributesTypeEval = shared.InsertExperimentEventReplaceSpanAttributesTypeEval

// This is an alias to an internal value.
const InsertExperimentEventReplaceSpanAttributesTypeTask = shared.InsertExperimentEventReplaceSpanAttributesTypeTask

// This is an alias to an internal value.
const InsertExperimentEventReplaceSpanAttributesTypeTool = shared.InsertExperimentEventReplaceSpanAttributesTypeTool

// This is an alias to an internal type.
type InsertProjectLogsEventMergeParam = shared.InsertProjectLogsEventMergeParam

// Context is additional information about the code that produced the project logs
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the project
// logs event
//
// This is an alias to an internal type.
type InsertProjectLogsEventMergeContextParam = shared.InsertProjectLogsEventMergeContextParam

// Metrics are numerical measurements tracking the execution of the code that
// produced the project logs event. Use "start" and "end" to track the time span
// over which the project logs event was produced
//
// This is an alias to an internal type.
type InsertProjectLogsEventMergeMetricsParam = shared.InsertProjectLogsEventMergeMetricsParam

// Human-identifying attributes of the span, such as name, type, etc.
//
// This is an alias to an internal type.
type InsertProjectLogsEventMergeSpanAttributesParam = shared.InsertProjectLogsEventMergeSpanAttributesParam

// Type of the span, for display purposes only
//
// This is an alias to an internal type.
type InsertProjectLogsEventMergeSpanAttributesType = shared.InsertProjectLogsEventMergeSpanAttributesType

// This is an alias to an internal value.
const InsertProjectLogsEventMergeSpanAttributesTypeLlm = shared.InsertProjectLogsEventMergeSpanAttributesTypeLlm

// This is an alias to an internal value.
const InsertProjectLogsEventMergeSpanAttributesTypeScore = shared.InsertProjectLogsEventMergeSpanAttributesTypeScore

// This is an alias to an internal value.
const InsertProjectLogsEventMergeSpanAttributesTypeFunction = shared.InsertProjectLogsEventMergeSpanAttributesTypeFunction

// This is an alias to an internal value.
const InsertProjectLogsEventMergeSpanAttributesTypeEval = shared.InsertProjectLogsEventMergeSpanAttributesTypeEval

// This is an alias to an internal value.
const InsertProjectLogsEventMergeSpanAttributesTypeTask = shared.InsertProjectLogsEventMergeSpanAttributesTypeTask

// This is an alias to an internal value.
const InsertProjectLogsEventMergeSpanAttributesTypeTool = shared.InsertProjectLogsEventMergeSpanAttributesTypeTool

// This is an alias to an internal type.
type InsertProjectLogsEventReplaceParam = shared.InsertProjectLogsEventReplaceParam

// Context is additional information about the code that produced the project logs
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the project
// logs event
//
// This is an alias to an internal type.
type InsertProjectLogsEventReplaceContextParam = shared.InsertProjectLogsEventReplaceContextParam

// Metrics are numerical measurements tracking the execution of the code that
// produced the project logs event. Use "start" and "end" to track the time span
// over which the project logs event was produced
//
// This is an alias to an internal type.
type InsertProjectLogsEventReplaceMetricsParam = shared.InsertProjectLogsEventReplaceMetricsParam

// Human-identifying attributes of the span, such as name, type, etc.
//
// This is an alias to an internal type.
type InsertProjectLogsEventReplaceSpanAttributesParam = shared.InsertProjectLogsEventReplaceSpanAttributesParam

// Type of the span, for display purposes only
//
// This is an alias to an internal type.
type InsertProjectLogsEventReplaceSpanAttributesType = shared.InsertProjectLogsEventReplaceSpanAttributesType

// This is an alias to an internal value.
const InsertProjectLogsEventReplaceSpanAttributesTypeLlm = shared.InsertProjectLogsEventReplaceSpanAttributesTypeLlm

// This is an alias to an internal value.
const InsertProjectLogsEventReplaceSpanAttributesTypeScore = shared.InsertProjectLogsEventReplaceSpanAttributesTypeScore

// This is an alias to an internal value.
const InsertProjectLogsEventReplaceSpanAttributesTypeFunction = shared.InsertProjectLogsEventReplaceSpanAttributesTypeFunction

// This is an alias to an internal value.
const InsertProjectLogsEventReplaceSpanAttributesTypeEval = shared.InsertProjectLogsEventReplaceSpanAttributesTypeEval

// This is an alias to an internal value.
const InsertProjectLogsEventReplaceSpanAttributesTypeTask = shared.InsertProjectLogsEventReplaceSpanAttributesTypeTask

// This is an alias to an internal value.
const InsertProjectLogsEventReplaceSpanAttributesTypeTool = shared.InsertProjectLogsEventReplaceSpanAttributesTypeTool

// This is an alias to an internal type.
type Messages = shared.Messages

// This is an alias to an internal type.
type MessagesSystem = shared.MessagesSystem

// This is an alias to an internal type.
type MessagesSystemRole = shared.MessagesSystemRole

// This is an alias to an internal value.
const MessagesSystemRoleSystem = shared.MessagesSystemRoleSystem

// This is an alias to an internal type.
type MessagesUser = shared.MessagesUser

// This is an alias to an internal type.
type MessagesUserRole = shared.MessagesUserRole

// This is an alias to an internal value.
const MessagesUserRoleUser = shared.MessagesUserRoleUser

// This is an alias to an internal type.
type MessagesAssistant = shared.MessagesAssistant

// This is an alias to an internal type.
type MessagesAssistantRole = shared.MessagesAssistantRole

// This is an alias to an internal value.
const MessagesAssistantRoleAssistant = shared.MessagesAssistantRoleAssistant

// This is an alias to an internal type.
type MessagesAssistantFunctionCall = shared.MessagesAssistantFunctionCall

// This is an alias to an internal type.
type MessagesAssistantToolCall = shared.MessagesAssistantToolCall

// This is an alias to an internal type.
type MessagesAssistantToolCallsFunction = shared.MessagesAssistantToolCallsFunction

// This is an alias to an internal type.
type MessagesAssistantToolCallsType = shared.MessagesAssistantToolCallsType

// This is an alias to an internal value.
const MessagesAssistantToolCallsTypeFunction = shared.MessagesAssistantToolCallsTypeFunction

// This is an alias to an internal type.
type MessagesTool = shared.MessagesTool

// This is an alias to an internal type.
type MessagesToolRole = shared.MessagesToolRole

// This is an alias to an internal value.
const MessagesToolRoleTool = shared.MessagesToolRoleTool

// This is an alias to an internal type.
type MessagesFunction = shared.MessagesFunction

// This is an alias to an internal type.
type MessagesFunctionRole = shared.MessagesFunctionRole

// This is an alias to an internal value.
const MessagesFunctionRoleFunction = shared.MessagesFunctionRoleFunction

// This is an alias to an internal type.
type MessagesFallback = shared.MessagesFallback

// This is an alias to an internal type.
type MessagesFallbackRole = shared.MessagesFallbackRole

// This is an alias to an internal value.
const MessagesFallbackRoleModel = shared.MessagesFallbackRoleModel

// This is an alias to an internal type.
type MessagesRole = shared.MessagesRole

// This is an alias to an internal value.
const MessagesRoleSystem = shared.MessagesRoleSystem

// This is an alias to an internal value.
const MessagesRoleUser = shared.MessagesRoleUser

// This is an alias to an internal value.
const MessagesRoleAssistant = shared.MessagesRoleAssistant

// This is an alias to an internal value.
const MessagesRoleTool = shared.MessagesRoleTool

// This is an alias to an internal value.
const MessagesRoleFunction = shared.MessagesRoleFunction

// This is an alias to an internal value.
const MessagesRoleModel = shared.MessagesRoleModel

// This is an alias to an internal type.
type MessagesUnionParam = shared.MessagesUnionParam

// This is an alias to an internal type.
type MessagesSystemParam = shared.MessagesSystemParam

// This is an alias to an internal type.
type MessagesUserParam = shared.MessagesUserParam

// This is an alias to an internal type.
type MessagesAssistantParam = shared.MessagesAssistantParam

// This is an alias to an internal type.
type MessagesAssistantFunctionCallParam = shared.MessagesAssistantFunctionCallParam

// This is an alias to an internal type.
type MessagesAssistantToolCallParam = shared.MessagesAssistantToolCallParam

// This is an alias to an internal type.
type MessagesAssistantToolCallsFunctionParam = shared.MessagesAssistantToolCallsFunctionParam

// This is an alias to an internal type.
type MessagesToolParam = shared.MessagesToolParam

// This is an alias to an internal type.
type MessagesFunctionParam = shared.MessagesFunctionParam

// This is an alias to an internal type.
type MessagesFallbackParam = shared.MessagesFallbackParam

// Summary of a metric's performance
//
// This is an alias to an internal type.
type MetricSummary = shared.MetricSummary

// This is an alias to an internal type.
type Organization = shared.Organization

// A path-lookup filter describes an equality comparison against a specific
// sub-field in the event row. For instance, if you wish to filter on the value of
// `c` in `{"input": {"a": {"b": {"c": "hello"}}}}`, pass
// `path=["input", "a", "b", "c"]` and `value="hello"`
//
// This is an alias to an internal type.
type PathLookupFilterParam = shared.PathLookupFilterParam

// Denotes the type of filter as a path-lookup filter
//
// This is an alias to an internal type.
type PathLookupFilterType = shared.PathLookupFilterType

// This is an alias to an internal value.
const PathLookupFilterTypePathLookup = shared.PathLookupFilterTypePathLookup

// This is an alias to an internal type.
type Project = shared.Project

// This is an alias to an internal type.
type ProjectSettings = shared.ProjectSettings

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

// Metrics are numerical measurements tracking the execution of the code that
// produced the project logs event. Use "start" and "end" to track the time span
// over which the project logs event was produced
//
// This is an alias to an internal type.
type ProjectLogsEventMetrics = shared.ProjectLogsEventMetrics

// Human-identifying attributes of the span, such as name, type, etc.
//
// This is an alias to an internal type.
type ProjectLogsEventSpanAttributes = shared.ProjectLogsEventSpanAttributes

// Type of the span, for display purposes only
//
// This is an alias to an internal type.
type ProjectLogsEventSpanAttributesType = shared.ProjectLogsEventSpanAttributesType

// This is an alias to an internal value.
const ProjectLogsEventSpanAttributesTypeLlm = shared.ProjectLogsEventSpanAttributesTypeLlm

// This is an alias to an internal value.
const ProjectLogsEventSpanAttributesTypeScore = shared.ProjectLogsEventSpanAttributesTypeScore

// This is an alias to an internal value.
const ProjectLogsEventSpanAttributesTypeFunction = shared.ProjectLogsEventSpanAttributesTypeFunction

// This is an alias to an internal value.
const ProjectLogsEventSpanAttributesTypeEval = shared.ProjectLogsEventSpanAttributesTypeEval

// This is an alias to an internal value.
const ProjectLogsEventSpanAttributesTypeTask = shared.ProjectLogsEventSpanAttributesTypeTask

// This is an alias to an internal value.
const ProjectLogsEventSpanAttributesTypeTool = shared.ProjectLogsEventSpanAttributesTypeTool

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
const ProjectScoreScoreTypeOnline = shared.ProjectScoreScoreTypeOnline

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

// This is an alias to an internal type.
type ProjectScoreCategoriesNullableVariant = shared.ProjectScoreCategoriesNullableVariant

// This is an alias to an internal type.
type ProjectScoreConfig = shared.ProjectScoreConfig

// This is an alias to an internal type.
type ProjectScoreConfigDestination = shared.ProjectScoreConfigDestination

// This is an alias to an internal value.
const ProjectScoreConfigDestinationExpected = shared.ProjectScoreConfigDestinationExpected

// This is an alias to an internal type.
type ProjectScoreConfigOnline = shared.ProjectScoreConfigOnline

// This is an alias to an internal type.
type ProjectScoreConfigOnlineScorer = shared.ProjectScoreConfigOnlineScorer

// This is an alias to an internal type.
type ProjectScoreConfigOnlineScorersFunction = shared.ProjectScoreConfigOnlineScorersFunction

// This is an alias to an internal type.
type ProjectScoreConfigOnlineScorersFunctionType = shared.ProjectScoreConfigOnlineScorersFunctionType

// This is an alias to an internal value.
const ProjectScoreConfigOnlineScorersFunctionTypeFunction = shared.ProjectScoreConfigOnlineScorersFunctionTypeFunction

// This is an alias to an internal type.
type ProjectScoreConfigOnlineScorersGlobal = shared.ProjectScoreConfigOnlineScorersGlobal

// This is an alias to an internal type.
type ProjectScoreConfigOnlineScorersGlobalType = shared.ProjectScoreConfigOnlineScorersGlobalType

// This is an alias to an internal value.
const ProjectScoreConfigOnlineScorersGlobalTypeGlobal = shared.ProjectScoreConfigOnlineScorersGlobalTypeGlobal

// This is an alias to an internal type.
type ProjectScoreConfigOnlineScorersType = shared.ProjectScoreConfigOnlineScorersType

// This is an alias to an internal value.
const ProjectScoreConfigOnlineScorersTypeFunction = shared.ProjectScoreConfigOnlineScorersTypeFunction

// This is an alias to an internal value.
const ProjectScoreConfigOnlineScorersTypeGlobal = shared.ProjectScoreConfigOnlineScorersTypeGlobal

// For categorical-type project scores, defines a single category
//
// This is an alias to an internal type.
type ProjectScoreCategory = shared.ProjectScoreCategory

// For categorical-type project scores, defines a single category
//
// This is an alias to an internal type.
type ProjectScoreCategoryParam = shared.ProjectScoreCategoryParam

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
type PromptDataParserParam = shared.PromptDataParserParam

// This is an alias to an internal type.
type PromptDataPromptUnionParam = shared.PromptDataPromptUnionParam

// This is an alias to an internal type.
type PromptDataPromptCompletionParam = shared.PromptDataPromptCompletionParam

// This is an alias to an internal type.
type PromptDataPromptChatParam = shared.PromptDataPromptChatParam

// This is an alias to an internal type.
type PromptDataPromptNullableVariantParam = shared.PromptDataPromptNullableVariantParam

// This is an alias to an internal type.
type PromptDataToolFunctionsUnionParam = shared.PromptDataToolFunctionsUnionParam

// This is an alias to an internal type.
type PromptDataToolFunctionsFunctionParam = shared.PromptDataToolFunctionsFunctionParam

// This is an alias to an internal type.
type PromptDataToolFunctionsGlobalParam = shared.PromptDataToolFunctionsGlobalParam

// This is an alias to an internal type.
type PromptImageURL = shared.PromptImageURL

// This is an alias to an internal type.
type PromptImageURLDetail = shared.PromptImageURLDetail

// This is an alias to an internal value.
const PromptImageURLDetailAuto = shared.PromptImageURLDetailAuto

// This is an alias to an internal value.
const PromptImageURLDetailLow = shared.PromptImageURLDetailLow

// This is an alias to an internal value.
const PromptImageURLDetailHigh = shared.PromptImageURLDetailHigh

// This is an alias to an internal type.
type PromptImageURLParam = shared.PromptImageURLParam

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

// This is an alias to an internal type.
type Scorer = shared.Scorer

// This is an alias to an internal type.
type ScorerType = shared.ScorerType

// This is an alias to an internal value.
const ScorerTypeScorer = shared.ScorerTypeScorer

// Summary of a dataset
//
// This is an alias to an internal type.
type SummarizeDatasetResponse = shared.SummarizeDatasetResponse

// Summary of an experiment
//
// This is an alias to an internal type.
type SummarizeExperimentResponse = shared.SummarizeExperimentResponse

// This is an alias to an internal type.
type Task = shared.Task

// This is an alias to an internal type.
type TaskType = shared.TaskType

// This is an alias to an internal value.
const TaskTypeTask = shared.TaskTypeTask

// This is an alias to an internal type.
type ToolChoiceFunction = shared.ToolChoiceFunction

// This is an alias to an internal type.
type ToolChoiceFunctionParam = shared.ToolChoiceFunctionParam

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
const ViewViewTypeLogs = shared.ViewViewTypeLogs

// This is an alias to an internal value.
const ViewViewTypeExperiments = shared.ViewViewTypeExperiments

// This is an alias to an internal value.
const ViewViewTypeDatasets = shared.ViewViewTypeDatasets

// This is an alias to an internal value.
const ViewViewTypePrompts = shared.ViewViewTypePrompts

// This is an alias to an internal value.
const ViewViewTypePlaygrounds = shared.ViewViewTypePlaygrounds

// This is an alias to an internal value.
const ViewViewTypeExperiment = shared.ViewViewTypeExperiment

// This is an alias to an internal value.
const ViewViewTypeDataset = shared.ViewViewTypeDataset

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
