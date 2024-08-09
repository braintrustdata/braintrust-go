// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"github.com/braintrustdata/braintrust-go/internal/apierror"
	"github.com/braintrustdata/braintrust-go/shared"
)

type Error = apierror.Error

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
type CreateACLParam = shared.CreateACLParam

// The object type that the ACL applies to
//
// This is an alias to an internal type.
type CreateACLObjectType = shared.CreateACLObjectType

// This is an alias to an internal value.
const CreateACLObjectTypeOrganization = shared.CreateACLObjectTypeOrganization

// This is an alias to an internal value.
const CreateACLObjectTypeProject = shared.CreateACLObjectTypeProject

// This is an alias to an internal value.
const CreateACLObjectTypeExperiment = shared.CreateACLObjectTypeExperiment

// This is an alias to an internal value.
const CreateACLObjectTypeDataset = shared.CreateACLObjectTypeDataset

// This is an alias to an internal value.
const CreateACLObjectTypePrompt = shared.CreateACLObjectTypePrompt

// This is an alias to an internal value.
const CreateACLObjectTypePromptSession = shared.CreateACLObjectTypePromptSession

// This is an alias to an internal value.
const CreateACLObjectTypeGroup = shared.CreateACLObjectTypeGroup

// This is an alias to an internal value.
const CreateACLObjectTypeRole = shared.CreateACLObjectTypeRole

// This is an alias to an internal value.
const CreateACLObjectTypeOrgMember = shared.CreateACLObjectTypeOrgMember

// This is an alias to an internal value.
const CreateACLObjectTypeProjectLog = shared.CreateACLObjectTypeProjectLog

// This is an alias to an internal value.
const CreateACLObjectTypeOrgProject = shared.CreateACLObjectTypeOrgProject

// Permission the ACL grants. Exactly one of `permission` and `role_id` will be
// provided
//
// This is an alias to an internal type.
type CreateACLPermission = shared.CreateACLPermission

// This is an alias to an internal value.
const CreateACLPermissionCreate = shared.CreateACLPermissionCreate

// This is an alias to an internal value.
const CreateACLPermissionRead = shared.CreateACLPermissionRead

// This is an alias to an internal value.
const CreateACLPermissionUpdate = shared.CreateACLPermissionUpdate

// This is an alias to an internal value.
const CreateACLPermissionDelete = shared.CreateACLPermissionDelete

// This is an alias to an internal value.
const CreateACLPermissionCreateACLs = shared.CreateACLPermissionCreateACLs

// This is an alias to an internal value.
const CreateACLPermissionReadACLs = shared.CreateACLPermissionReadACLs

// This is an alias to an internal value.
const CreateACLPermissionUpdateACLs = shared.CreateACLPermissionUpdateACLs

// This is an alias to an internal value.
const CreateACLPermissionDeleteACLs = shared.CreateACLPermissionDeleteACLs

// When setting a permission directly, optionally restricts the permission grant to
// just the specified object type. Cannot be set alongside a `role_id`.
//
// This is an alias to an internal type.
type CreateACLRestrictObjectType = shared.CreateACLRestrictObjectType

// This is an alias to an internal value.
const CreateACLRestrictObjectTypeOrganization = shared.CreateACLRestrictObjectTypeOrganization

// This is an alias to an internal value.
const CreateACLRestrictObjectTypeProject = shared.CreateACLRestrictObjectTypeProject

// This is an alias to an internal value.
const CreateACLRestrictObjectTypeExperiment = shared.CreateACLRestrictObjectTypeExperiment

// This is an alias to an internal value.
const CreateACLRestrictObjectTypeDataset = shared.CreateACLRestrictObjectTypeDataset

// This is an alias to an internal value.
const CreateACLRestrictObjectTypePrompt = shared.CreateACLRestrictObjectTypePrompt

// This is an alias to an internal value.
const CreateACLRestrictObjectTypePromptSession = shared.CreateACLRestrictObjectTypePromptSession

// This is an alias to an internal value.
const CreateACLRestrictObjectTypeGroup = shared.CreateACLRestrictObjectTypeGroup

// This is an alias to an internal value.
const CreateACLRestrictObjectTypeRole = shared.CreateACLRestrictObjectTypeRole

// This is an alias to an internal value.
const CreateACLRestrictObjectTypeOrgMember = shared.CreateACLRestrictObjectTypeOrgMember

// This is an alias to an internal value.
const CreateACLRestrictObjectTypeProjectLog = shared.CreateACLRestrictObjectTypeProjectLog

// This is an alias to an internal value.
const CreateACLRestrictObjectTypeOrgProject = shared.CreateACLRestrictObjectTypeOrgProject

// This is an alias to an internal type.
type CreateAPIKeyOutput = shared.CreateAPIKeyOutput

// This is an alias to an internal type.
type CreateDatasetParam = shared.CreateDatasetParam

// This is an alias to an internal type.
type CreateExperimentParam = shared.CreateExperimentParam

// This is an alias to an internal type.
type CreateFunctionParam = shared.CreateFunctionParam

// This is an alias to an internal type.
type CreateFunctionFunctionDataUnionParam = shared.CreateFunctionFunctionDataUnionParam

// This is an alias to an internal type.
type CreateFunctionFunctionDataPromptParam = shared.CreateFunctionFunctionDataPromptParam

// This is an alias to an internal type.
type CreateFunctionFunctionDataPromptType = shared.CreateFunctionFunctionDataPromptType

// This is an alias to an internal value.
const CreateFunctionFunctionDataPromptTypePrompt = shared.CreateFunctionFunctionDataPromptTypePrompt

// This is an alias to an internal type.
type CreateFunctionFunctionDataCodeParam = shared.CreateFunctionFunctionDataCodeParam

// This is an alias to an internal type.
type CreateFunctionFunctionDataCodeDataParam = shared.CreateFunctionFunctionDataCodeDataParam

// This is an alias to an internal type.
type CreateFunctionFunctionDataCodeDataLocationParam = shared.CreateFunctionFunctionDataCodeDataLocationParam

// This is an alias to an internal type.
type CreateFunctionFunctionDataCodeDataLocationPositionUnionParam = shared.CreateFunctionFunctionDataCodeDataLocationPositionUnionParam

// This is an alias to an internal type.
type CreateFunctionFunctionDataCodeDataLocationPositionTask = shared.CreateFunctionFunctionDataCodeDataLocationPositionTask

// This is an alias to an internal value.
const CreateFunctionFunctionDataCodeDataLocationPositionTaskTask = shared.CreateFunctionFunctionDataCodeDataLocationPositionTaskTask

// This is an alias to an internal type.
type CreateFunctionFunctionDataCodeDataLocationPositionScoreParam = shared.CreateFunctionFunctionDataCodeDataLocationPositionScoreParam

// This is an alias to an internal type.
type CreateFunctionFunctionDataCodeDataLocationType = shared.CreateFunctionFunctionDataCodeDataLocationType

// This is an alias to an internal value.
const CreateFunctionFunctionDataCodeDataLocationTypeExperiment = shared.CreateFunctionFunctionDataCodeDataLocationTypeExperiment

// This is an alias to an internal type.
type CreateFunctionFunctionDataCodeDataRuntimeContextParam = shared.CreateFunctionFunctionDataCodeDataRuntimeContextParam

// This is an alias to an internal type.
type CreateFunctionFunctionDataCodeDataRuntimeContextRuntime = shared.CreateFunctionFunctionDataCodeDataRuntimeContextRuntime

// This is an alias to an internal value.
const CreateFunctionFunctionDataCodeDataRuntimeContextRuntimeNode = shared.CreateFunctionFunctionDataCodeDataRuntimeContextRuntimeNode

// This is an alias to an internal type.
type CreateFunctionFunctionDataCodeType = shared.CreateFunctionFunctionDataCodeType

// This is an alias to an internal value.
const CreateFunctionFunctionDataCodeTypeCode = shared.CreateFunctionFunctionDataCodeTypeCode

// This is an alias to an internal type.
type CreateFunctionFunctionDataGlobalParam = shared.CreateFunctionFunctionDataGlobalParam

// This is an alias to an internal type.
type CreateFunctionFunctionDataGlobalType = shared.CreateFunctionFunctionDataGlobalType

// This is an alias to an internal value.
const CreateFunctionFunctionDataGlobalTypeGlobal = shared.CreateFunctionFunctionDataGlobalTypeGlobal

// This is an alias to an internal type.
type CreateFunctionFunctionDataType = shared.CreateFunctionFunctionDataType

// This is an alias to an internal value.
const CreateFunctionFunctionDataTypePrompt = shared.CreateFunctionFunctionDataTypePrompt

// This is an alias to an internal value.
const CreateFunctionFunctionDataTypeCode = shared.CreateFunctionFunctionDataTypeCode

// This is an alias to an internal value.
const CreateFunctionFunctionDataTypeGlobal = shared.CreateFunctionFunctionDataTypeGlobal

// This is an alias to an internal type.
type CreateGroupParam = shared.CreateGroupParam

// This is an alias to an internal type.
type CreateProjectParam = shared.CreateProjectParam

// This is an alias to an internal type.
type CreateProjectScoreParam = shared.CreateProjectScoreParam

// The type of the configured score
//
// This is an alias to an internal type.
type CreateProjectScoreScoreType = shared.CreateProjectScoreScoreType

// This is an alias to an internal value.
const CreateProjectScoreScoreTypeSlider = shared.CreateProjectScoreScoreTypeSlider

// This is an alias to an internal value.
const CreateProjectScoreScoreTypeCategorical = shared.CreateProjectScoreScoreTypeCategorical

// This is an alias to an internal value.
const CreateProjectScoreScoreTypeWeighted = shared.CreateProjectScoreScoreTypeWeighted

// This is an alias to an internal value.
const CreateProjectScoreScoreTypeMinimum = shared.CreateProjectScoreScoreTypeMinimum

// For categorical-type project scores, the list of all categories
//
// This is an alias to an internal type.
type CreateProjectScoreCategoriesUnionParam = shared.CreateProjectScoreCategoriesUnionParam

// For categorical-type project scores, the list of all categories
//
// This is an alias to an internal type.
type CreateProjectScoreCategoriesCategoricalParam = shared.CreateProjectScoreCategoriesCategoricalParam

// For minimum-type project scores, the list of included scores
//
// This is an alias to an internal type.
type CreateProjectScoreCategoriesMinimumParam = shared.CreateProjectScoreCategoriesMinimumParam

// This is an alias to an internal type.
type CreateProjectScoreCategoriesNullableVariantParam = shared.CreateProjectScoreCategoriesNullableVariantParam

// This is an alias to an internal type.
type CreateProjectTagParam = shared.CreateProjectTagParam

// This is an alias to an internal type.
type CreatePromptParam = shared.CreatePromptParam

// This is an alias to an internal type.
type CreateRoleParam = shared.CreateRoleParam

// This is an alias to an internal type.
type CreateRoleMemberPermissionParam = shared.CreateRoleMemberPermissionParam

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
//
// This is an alias to an internal type.
type CreateRoleMemberPermissionsPermission = shared.CreateRoleMemberPermissionsPermission

// This is an alias to an internal value.
const CreateRoleMemberPermissionsPermissionCreate = shared.CreateRoleMemberPermissionsPermissionCreate

// This is an alias to an internal value.
const CreateRoleMemberPermissionsPermissionRead = shared.CreateRoleMemberPermissionsPermissionRead

// This is an alias to an internal value.
const CreateRoleMemberPermissionsPermissionUpdate = shared.CreateRoleMemberPermissionsPermissionUpdate

// This is an alias to an internal value.
const CreateRoleMemberPermissionsPermissionDelete = shared.CreateRoleMemberPermissionsPermissionDelete

// This is an alias to an internal value.
const CreateRoleMemberPermissionsPermissionCreateACLs = shared.CreateRoleMemberPermissionsPermissionCreateACLs

// This is an alias to an internal value.
const CreateRoleMemberPermissionsPermissionReadACLs = shared.CreateRoleMemberPermissionsPermissionReadACLs

// This is an alias to an internal value.
const CreateRoleMemberPermissionsPermissionUpdateACLs = shared.CreateRoleMemberPermissionsPermissionUpdateACLs

// This is an alias to an internal value.
const CreateRoleMemberPermissionsPermissionDeleteACLs = shared.CreateRoleMemberPermissionsPermissionDeleteACLs

// The object type that the ACL applies to
//
// This is an alias to an internal type.
type CreateRoleMemberPermissionsRestrictObjectType = shared.CreateRoleMemberPermissionsRestrictObjectType

// This is an alias to an internal value.
const CreateRoleMemberPermissionsRestrictObjectTypeOrganization = shared.CreateRoleMemberPermissionsRestrictObjectTypeOrganization

// This is an alias to an internal value.
const CreateRoleMemberPermissionsRestrictObjectTypeProject = shared.CreateRoleMemberPermissionsRestrictObjectTypeProject

// This is an alias to an internal value.
const CreateRoleMemberPermissionsRestrictObjectTypeExperiment = shared.CreateRoleMemberPermissionsRestrictObjectTypeExperiment

// This is an alias to an internal value.
const CreateRoleMemberPermissionsRestrictObjectTypeDataset = shared.CreateRoleMemberPermissionsRestrictObjectTypeDataset

// This is an alias to an internal value.
const CreateRoleMemberPermissionsRestrictObjectTypePrompt = shared.CreateRoleMemberPermissionsRestrictObjectTypePrompt

// This is an alias to an internal value.
const CreateRoleMemberPermissionsRestrictObjectTypePromptSession = shared.CreateRoleMemberPermissionsRestrictObjectTypePromptSession

// This is an alias to an internal value.
const CreateRoleMemberPermissionsRestrictObjectTypeGroup = shared.CreateRoleMemberPermissionsRestrictObjectTypeGroup

// This is an alias to an internal value.
const CreateRoleMemberPermissionsRestrictObjectTypeRole = shared.CreateRoleMemberPermissionsRestrictObjectTypeRole

// This is an alias to an internal value.
const CreateRoleMemberPermissionsRestrictObjectTypeOrgMember = shared.CreateRoleMemberPermissionsRestrictObjectTypeOrgMember

// This is an alias to an internal value.
const CreateRoleMemberPermissionsRestrictObjectTypeProjectLog = shared.CreateRoleMemberPermissionsRestrictObjectTypeProjectLog

// This is an alias to an internal value.
const CreateRoleMemberPermissionsRestrictObjectTypeOrgProject = shared.CreateRoleMemberPermissionsRestrictObjectTypeOrgProject

// This is an alias to an internal type.
type CreateViewParam = shared.CreateViewParam

// The object type that the ACL applies to
//
// This is an alias to an internal type.
type CreateViewObjectType = shared.CreateViewObjectType

// This is an alias to an internal value.
const CreateViewObjectTypeOrganization = shared.CreateViewObjectTypeOrganization

// This is an alias to an internal value.
const CreateViewObjectTypeProject = shared.CreateViewObjectTypeProject

// This is an alias to an internal value.
const CreateViewObjectTypeExperiment = shared.CreateViewObjectTypeExperiment

// This is an alias to an internal value.
const CreateViewObjectTypeDataset = shared.CreateViewObjectTypeDataset

// This is an alias to an internal value.
const CreateViewObjectTypePrompt = shared.CreateViewObjectTypePrompt

// This is an alias to an internal value.
const CreateViewObjectTypePromptSession = shared.CreateViewObjectTypePromptSession

// This is an alias to an internal value.
const CreateViewObjectTypeGroup = shared.CreateViewObjectTypeGroup

// This is an alias to an internal value.
const CreateViewObjectTypeRole = shared.CreateViewObjectTypeRole

// This is an alias to an internal value.
const CreateViewObjectTypeOrgMember = shared.CreateViewObjectTypeOrgMember

// This is an alias to an internal value.
const CreateViewObjectTypeProjectLog = shared.CreateViewObjectTypeProjectLog

// This is an alias to an internal value.
const CreateViewObjectTypeOrgProject = shared.CreateViewObjectTypeOrgProject

// Type of table that the view corresponds to.
//
// This is an alias to an internal type.
type CreateViewViewType = shared.CreateViewViewType

// This is an alias to an internal value.
const CreateViewViewTypeProjects = shared.CreateViewViewTypeProjects

// This is an alias to an internal value.
const CreateViewViewTypeLogs = shared.CreateViewViewTypeLogs

// This is an alias to an internal value.
const CreateViewViewTypeExperiments = shared.CreateViewViewTypeExperiments

// This is an alias to an internal value.
const CreateViewViewTypeDatasets = shared.CreateViewViewTypeDatasets

// This is an alias to an internal value.
const CreateViewViewTypePrompts = shared.CreateViewViewTypePrompts

// This is an alias to an internal value.
const CreateViewViewTypePlaygrounds = shared.CreateViewViewTypePlaygrounds

// This is an alias to an internal value.
const CreateViewViewTypeExperiment = shared.CreateViewViewTypeExperiment

// This is an alias to an internal value.
const CreateViewViewTypeDataset = shared.CreateViewViewTypeDataset

// Summary of a dataset's data
//
// This is an alias to an internal type.
type DataSummary = shared.DataSummary

// This is an alias to an internal type.
type Dataset = shared.Dataset

// This is an alias to an internal type.
type DatasetEvent = shared.DatasetEvent

// This is an alias to an internal type.
type DeleteViewParam = shared.DeleteViewParam

// The object type that the ACL applies to
//
// This is an alias to an internal type.
type DeleteViewObjectType = shared.DeleteViewObjectType

// This is an alias to an internal value.
const DeleteViewObjectTypeOrganization = shared.DeleteViewObjectTypeOrganization

// This is an alias to an internal value.
const DeleteViewObjectTypeProject = shared.DeleteViewObjectTypeProject

// This is an alias to an internal value.
const DeleteViewObjectTypeExperiment = shared.DeleteViewObjectTypeExperiment

// This is an alias to an internal value.
const DeleteViewObjectTypeDataset = shared.DeleteViewObjectTypeDataset

// This is an alias to an internal value.
const DeleteViewObjectTypePrompt = shared.DeleteViewObjectTypePrompt

// This is an alias to an internal value.
const DeleteViewObjectTypePromptSession = shared.DeleteViewObjectTypePromptSession

// This is an alias to an internal value.
const DeleteViewObjectTypeGroup = shared.DeleteViewObjectTypeGroup

// This is an alias to an internal value.
const DeleteViewObjectTypeRole = shared.DeleteViewObjectTypeRole

// This is an alias to an internal value.
const DeleteViewObjectTypeOrgMember = shared.DeleteViewObjectTypeOrgMember

// This is an alias to an internal value.
const DeleteViewObjectTypeProjectLog = shared.DeleteViewObjectTypeProjectLog

// This is an alias to an internal value.
const DeleteViewObjectTypeOrgProject = shared.DeleteViewObjectTypeOrgProject

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
type FeedbackDatasetEventRequestParam = shared.FeedbackDatasetEventRequestParam

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
type FeedbackExperimentEventRequestParam = shared.FeedbackExperimentEventRequestParam

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
type FeedbackProjectLogsEventRequestParam = shared.FeedbackProjectLogsEventRequestParam

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
type FetchDatasetEventsResponse = shared.FetchDatasetEventsResponse

// This is an alias to an internal type.
type FetchEventsRequestParam = shared.FetchEventsRequestParam

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
type FunctionFunctionDataCodeDataLocation = shared.FunctionFunctionDataCodeDataLocation

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataLocationPositionUnion = shared.FunctionFunctionDataCodeDataLocationPositionUnion

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataLocationPositionTask = shared.FunctionFunctionDataCodeDataLocationPositionTask

// This is an alias to an internal value.
const FunctionFunctionDataCodeDataLocationPositionTaskTask = shared.FunctionFunctionDataCodeDataLocationPositionTaskTask

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataLocationPositionScore = shared.FunctionFunctionDataCodeDataLocationPositionScore

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataLocationType = shared.FunctionFunctionDataCodeDataLocationType

// This is an alias to an internal value.
const FunctionFunctionDataCodeDataLocationTypeExperiment = shared.FunctionFunctionDataCodeDataLocationTypeExperiment

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataRuntimeContext = shared.FunctionFunctionDataCodeDataRuntimeContext

// This is an alias to an internal type.
type FunctionFunctionDataCodeDataRuntimeContextRuntime = shared.FunctionFunctionDataCodeDataRuntimeContextRuntime

// This is an alias to an internal value.
const FunctionFunctionDataCodeDataRuntimeContextRuntimeNode = shared.FunctionFunctionDataCodeDataRuntimeContextRuntimeNode

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
type InsertDatasetEventRequestParam = shared.InsertDatasetEventRequestParam

// A dataset event
//
// This is an alias to an internal type.
type InsertDatasetEventRequestEventsUnionParam = shared.InsertDatasetEventRequestEventsUnionParam

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
type InsertExperimentEventRequestParam = shared.InsertExperimentEventRequestParam

// An experiment event
//
// This is an alias to an internal type.
type InsertExperimentEventRequestEventsUnionParam = shared.InsertExperimentEventRequestEventsUnionParam

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
type InsertProjectLogsEventRequestParam = shared.InsertProjectLogsEventRequestParam

// A project logs event
//
// This is an alias to an internal type.
type InsertProjectLogsEventRequestEventsUnionParam = shared.InsertProjectLogsEventRequestEventsUnionParam

// Summary of a metric's performance
//
// This is an alias to an internal type.
type MetricSummary = shared.MetricSummary

// This is an alias to an internal type.
type Organization = shared.Organization

// This is an alias to an internal type.
type PatchDatasetParam = shared.PatchDatasetParam

// This is an alias to an internal type.
type PatchExperimentParam = shared.PatchExperimentParam

// This is an alias to an internal type.
type PatchFunctionParam = shared.PatchFunctionParam

// This is an alias to an internal type.
type PatchFunctionFunctionDataUnionParam = shared.PatchFunctionFunctionDataUnionParam

// This is an alias to an internal type.
type PatchFunctionFunctionDataPromptParam = shared.PatchFunctionFunctionDataPromptParam

// This is an alias to an internal type.
type PatchFunctionFunctionDataPromptType = shared.PatchFunctionFunctionDataPromptType

// This is an alias to an internal value.
const PatchFunctionFunctionDataPromptTypePrompt = shared.PatchFunctionFunctionDataPromptTypePrompt

// This is an alias to an internal type.
type PatchFunctionFunctionDataCodeParam = shared.PatchFunctionFunctionDataCodeParam

// This is an alias to an internal type.
type PatchFunctionFunctionDataCodeDataParam = shared.PatchFunctionFunctionDataCodeDataParam

// This is an alias to an internal type.
type PatchFunctionFunctionDataCodeDataLocationParam = shared.PatchFunctionFunctionDataCodeDataLocationParam

// This is an alias to an internal type.
type PatchFunctionFunctionDataCodeDataLocationPositionUnionParam = shared.PatchFunctionFunctionDataCodeDataLocationPositionUnionParam

// This is an alias to an internal type.
type PatchFunctionFunctionDataCodeDataLocationPositionTask = shared.PatchFunctionFunctionDataCodeDataLocationPositionTask

// This is an alias to an internal value.
const PatchFunctionFunctionDataCodeDataLocationPositionTaskTask = shared.PatchFunctionFunctionDataCodeDataLocationPositionTaskTask

// This is an alias to an internal type.
type PatchFunctionFunctionDataCodeDataLocationPositionScoreParam = shared.PatchFunctionFunctionDataCodeDataLocationPositionScoreParam

// This is an alias to an internal type.
type PatchFunctionFunctionDataCodeDataLocationType = shared.PatchFunctionFunctionDataCodeDataLocationType

// This is an alias to an internal value.
const PatchFunctionFunctionDataCodeDataLocationTypeExperiment = shared.PatchFunctionFunctionDataCodeDataLocationTypeExperiment

// This is an alias to an internal type.
type PatchFunctionFunctionDataCodeDataRuntimeContextParam = shared.PatchFunctionFunctionDataCodeDataRuntimeContextParam

// This is an alias to an internal type.
type PatchFunctionFunctionDataCodeDataRuntimeContextRuntime = shared.PatchFunctionFunctionDataCodeDataRuntimeContextRuntime

// This is an alias to an internal value.
const PatchFunctionFunctionDataCodeDataRuntimeContextRuntimeNode = shared.PatchFunctionFunctionDataCodeDataRuntimeContextRuntimeNode

// This is an alias to an internal type.
type PatchFunctionFunctionDataCodeType = shared.PatchFunctionFunctionDataCodeType

// This is an alias to an internal value.
const PatchFunctionFunctionDataCodeTypeCode = shared.PatchFunctionFunctionDataCodeTypeCode

// This is an alias to an internal type.
type PatchFunctionFunctionDataGlobalParam = shared.PatchFunctionFunctionDataGlobalParam

// This is an alias to an internal type.
type PatchFunctionFunctionDataGlobalType = shared.PatchFunctionFunctionDataGlobalType

// This is an alias to an internal value.
const PatchFunctionFunctionDataGlobalTypeGlobal = shared.PatchFunctionFunctionDataGlobalTypeGlobal

// This is an alias to an internal type.
type PatchFunctionFunctionDataNullableVariantParam = shared.PatchFunctionFunctionDataNullableVariantParam

// This is an alias to an internal type.
type PatchFunctionFunctionDataType = shared.PatchFunctionFunctionDataType

// This is an alias to an internal value.
const PatchFunctionFunctionDataTypePrompt = shared.PatchFunctionFunctionDataTypePrompt

// This is an alias to an internal value.
const PatchFunctionFunctionDataTypeCode = shared.PatchFunctionFunctionDataTypeCode

// This is an alias to an internal value.
const PatchFunctionFunctionDataTypeGlobal = shared.PatchFunctionFunctionDataTypeGlobal

// This is an alias to an internal type.
type PatchGroupParam = shared.PatchGroupParam

// This is an alias to an internal type.
type PatchOrganizationParam = shared.PatchOrganizationParam

// This is an alias to an internal type.
type PatchOrganizationMembersParam = shared.PatchOrganizationMembersParam

// Users to invite to the organization
//
// This is an alias to an internal type.
type PatchOrganizationMembersInviteUsersParam = shared.PatchOrganizationMembersInviteUsersParam

// Users to remove from the organization
//
// This is an alias to an internal type.
type PatchOrganizationMembersRemoveUsersParam = shared.PatchOrganizationMembersRemoveUsersParam

// This is an alias to an internal type.
type PatchProjectParam = shared.PatchProjectParam

// Project settings. Patch operations replace all settings, so make sure you
// include all settings you want to keep.
//
// This is an alias to an internal type.
type PatchProjectSettingsParam = shared.PatchProjectSettingsParam

// This is an alias to an internal type.
type PatchProjectScoreParam = shared.PatchProjectScoreParam

// For categorical-type project scores, the list of all categories
//
// This is an alias to an internal type.
type PatchProjectScoreCategoriesUnionParam = shared.PatchProjectScoreCategoriesUnionParam

// For categorical-type project scores, the list of all categories
//
// This is an alias to an internal type.
type PatchProjectScoreCategoriesCategoricalParam = shared.PatchProjectScoreCategoriesCategoricalParam

// For minimum-type project scores, the list of included scores
//
// This is an alias to an internal type.
type PatchProjectScoreCategoriesMinimumParam = shared.PatchProjectScoreCategoriesMinimumParam

// This is an alias to an internal type.
type PatchProjectScoreCategoriesNullableVariantParam = shared.PatchProjectScoreCategoriesNullableVariantParam

// The type of the configured score
//
// This is an alias to an internal type.
type PatchProjectScoreScoreType = shared.PatchProjectScoreScoreType

// This is an alias to an internal value.
const PatchProjectScoreScoreTypeSlider = shared.PatchProjectScoreScoreTypeSlider

// This is an alias to an internal value.
const PatchProjectScoreScoreTypeCategorical = shared.PatchProjectScoreScoreTypeCategorical

// This is an alias to an internal value.
const PatchProjectScoreScoreTypeWeighted = shared.PatchProjectScoreScoreTypeWeighted

// This is an alias to an internal value.
const PatchProjectScoreScoreTypeMinimum = shared.PatchProjectScoreScoreTypeMinimum

// This is an alias to an internal type.
type PatchProjectTagParam = shared.PatchProjectTagParam

// This is an alias to an internal type.
type PatchPromptParam = shared.PatchPromptParam

// This is an alias to an internal type.
type PatchRoleParam = shared.PatchRoleParam

// This is an alias to an internal type.
type PatchRoleAddMemberPermissionParam = shared.PatchRoleAddMemberPermissionParam

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
//
// This is an alias to an internal type.
type PatchRoleAddMemberPermissionsPermission = shared.PatchRoleAddMemberPermissionsPermission

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsPermissionCreate = shared.PatchRoleAddMemberPermissionsPermissionCreate

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsPermissionRead = shared.PatchRoleAddMemberPermissionsPermissionRead

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsPermissionUpdate = shared.PatchRoleAddMemberPermissionsPermissionUpdate

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsPermissionDelete = shared.PatchRoleAddMemberPermissionsPermissionDelete

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsPermissionCreateACLs = shared.PatchRoleAddMemberPermissionsPermissionCreateACLs

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsPermissionReadACLs = shared.PatchRoleAddMemberPermissionsPermissionReadACLs

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsPermissionUpdateACLs = shared.PatchRoleAddMemberPermissionsPermissionUpdateACLs

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsPermissionDeleteACLs = shared.PatchRoleAddMemberPermissionsPermissionDeleteACLs

// The object type that the ACL applies to
//
// This is an alias to an internal type.
type PatchRoleAddMemberPermissionsRestrictObjectType = shared.PatchRoleAddMemberPermissionsRestrictObjectType

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsRestrictObjectTypeOrganization = shared.PatchRoleAddMemberPermissionsRestrictObjectTypeOrganization

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsRestrictObjectTypeProject = shared.PatchRoleAddMemberPermissionsRestrictObjectTypeProject

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsRestrictObjectTypeExperiment = shared.PatchRoleAddMemberPermissionsRestrictObjectTypeExperiment

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsRestrictObjectTypeDataset = shared.PatchRoleAddMemberPermissionsRestrictObjectTypeDataset

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsRestrictObjectTypePrompt = shared.PatchRoleAddMemberPermissionsRestrictObjectTypePrompt

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsRestrictObjectTypePromptSession = shared.PatchRoleAddMemberPermissionsRestrictObjectTypePromptSession

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsRestrictObjectTypeGroup = shared.PatchRoleAddMemberPermissionsRestrictObjectTypeGroup

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsRestrictObjectTypeRole = shared.PatchRoleAddMemberPermissionsRestrictObjectTypeRole

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsRestrictObjectTypeOrgMember = shared.PatchRoleAddMemberPermissionsRestrictObjectTypeOrgMember

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsRestrictObjectTypeProjectLog = shared.PatchRoleAddMemberPermissionsRestrictObjectTypeProjectLog

// This is an alias to an internal value.
const PatchRoleAddMemberPermissionsRestrictObjectTypeOrgProject = shared.PatchRoleAddMemberPermissionsRestrictObjectTypeOrgProject

// This is an alias to an internal type.
type PatchRoleRemoveMemberPermissionParam = shared.PatchRoleRemoveMemberPermissionParam

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
//
// This is an alias to an internal type.
type PatchRoleRemoveMemberPermissionsPermission = shared.PatchRoleRemoveMemberPermissionsPermission

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsPermissionCreate = shared.PatchRoleRemoveMemberPermissionsPermissionCreate

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsPermissionRead = shared.PatchRoleRemoveMemberPermissionsPermissionRead

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsPermissionUpdate = shared.PatchRoleRemoveMemberPermissionsPermissionUpdate

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsPermissionDelete = shared.PatchRoleRemoveMemberPermissionsPermissionDelete

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsPermissionCreateACLs = shared.PatchRoleRemoveMemberPermissionsPermissionCreateACLs

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsPermissionReadACLs = shared.PatchRoleRemoveMemberPermissionsPermissionReadACLs

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsPermissionUpdateACLs = shared.PatchRoleRemoveMemberPermissionsPermissionUpdateACLs

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsPermissionDeleteACLs = shared.PatchRoleRemoveMemberPermissionsPermissionDeleteACLs

// The object type that the ACL applies to
//
// This is an alias to an internal type.
type PatchRoleRemoveMemberPermissionsRestrictObjectType = shared.PatchRoleRemoveMemberPermissionsRestrictObjectType

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsRestrictObjectTypeOrganization = shared.PatchRoleRemoveMemberPermissionsRestrictObjectTypeOrganization

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsRestrictObjectTypeProject = shared.PatchRoleRemoveMemberPermissionsRestrictObjectTypeProject

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsRestrictObjectTypeExperiment = shared.PatchRoleRemoveMemberPermissionsRestrictObjectTypeExperiment

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsRestrictObjectTypeDataset = shared.PatchRoleRemoveMemberPermissionsRestrictObjectTypeDataset

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsRestrictObjectTypePrompt = shared.PatchRoleRemoveMemberPermissionsRestrictObjectTypePrompt

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsRestrictObjectTypePromptSession = shared.PatchRoleRemoveMemberPermissionsRestrictObjectTypePromptSession

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsRestrictObjectTypeGroup = shared.PatchRoleRemoveMemberPermissionsRestrictObjectTypeGroup

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsRestrictObjectTypeRole = shared.PatchRoleRemoveMemberPermissionsRestrictObjectTypeRole

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsRestrictObjectTypeOrgMember = shared.PatchRoleRemoveMemberPermissionsRestrictObjectTypeOrgMember

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsRestrictObjectTypeProjectLog = shared.PatchRoleRemoveMemberPermissionsRestrictObjectTypeProjectLog

// This is an alias to an internal value.
const PatchRoleRemoveMemberPermissionsRestrictObjectTypeOrgProject = shared.PatchRoleRemoveMemberPermissionsRestrictObjectTypeOrgProject

// This is an alias to an internal type.
type PatchViewParam = shared.PatchViewParam

// The object type that the ACL applies to
//
// This is an alias to an internal type.
type PatchViewObjectType = shared.PatchViewObjectType

// This is an alias to an internal value.
const PatchViewObjectTypeOrganization = shared.PatchViewObjectTypeOrganization

// This is an alias to an internal value.
const PatchViewObjectTypeProject = shared.PatchViewObjectTypeProject

// This is an alias to an internal value.
const PatchViewObjectTypeExperiment = shared.PatchViewObjectTypeExperiment

// This is an alias to an internal value.
const PatchViewObjectTypeDataset = shared.PatchViewObjectTypeDataset

// This is an alias to an internal value.
const PatchViewObjectTypePrompt = shared.PatchViewObjectTypePrompt

// This is an alias to an internal value.
const PatchViewObjectTypePromptSession = shared.PatchViewObjectTypePromptSession

// This is an alias to an internal value.
const PatchViewObjectTypeGroup = shared.PatchViewObjectTypeGroup

// This is an alias to an internal value.
const PatchViewObjectTypeRole = shared.PatchViewObjectTypeRole

// This is an alias to an internal value.
const PatchViewObjectTypeOrgMember = shared.PatchViewObjectTypeOrgMember

// This is an alias to an internal value.
const PatchViewObjectTypeProjectLog = shared.PatchViewObjectTypeProjectLog

// This is an alias to an internal value.
const PatchViewObjectTypeOrgProject = shared.PatchViewObjectTypeOrgProject

// Type of table that the view corresponds to.
//
// This is an alias to an internal type.
type PatchViewViewType = shared.PatchViewViewType

// This is an alias to an internal value.
const PatchViewViewTypeProjects = shared.PatchViewViewTypeProjects

// This is an alias to an internal value.
const PatchViewViewTypeLogs = shared.PatchViewViewTypeLogs

// This is an alias to an internal value.
const PatchViewViewTypeExperiments = shared.PatchViewViewTypeExperiments

// This is an alias to an internal value.
const PatchViewViewTypeDatasets = shared.PatchViewViewTypeDatasets

// This is an alias to an internal value.
const PatchViewViewTypePrompts = shared.PatchViewViewTypePrompts

// This is an alias to an internal value.
const PatchViewViewTypePlaygrounds = shared.PatchViewViewTypePlaygrounds

// This is an alias to an internal value.
const PatchViewViewTypeExperiment = shared.PatchViewViewTypeExperiment

// This is an alias to an internal value.
const PatchViewViewTypeDataset = shared.PatchViewViewTypeDataset

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
