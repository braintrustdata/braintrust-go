// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"
	"time"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/internal/param"
	"github.com/tidwall/gjson"
)

type AISecret struct {
	// Unique identifier for the AI secret
	ID string `json:"id,required" format:"uuid"`
	// Name of the AI secret
	Name string `json:"name,required"`
	// Unique identifier for the organization
	OrgID string `json:"org_id,required" format:"uuid"`
	// Date of AI secret creation
	Created       time.Time              `json:"created,nullable" format:"date-time"`
	Metadata      map[string]interface{} `json:"metadata,nullable"`
	PreviewSecret string                 `json:"preview_secret,nullable"`
	Type          string                 `json:"type,nullable"`
	JSON          aiSecretJSON           `json:"-"`
}

// aiSecretJSON contains the JSON metadata for the struct [AISecret]
type aiSecretJSON struct {
	ID            apijson.Field
	Name          apijson.Field
	OrgID         apijson.Field
	Created       apijson.Field
	Metadata      apijson.Field
	PreviewSecret apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AISecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiSecretJSON) RawJSON() string {
	return r.raw
}

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
type ACL struct {
	// Unique identifier for the acl
	ID string `json:"id,required" format:"uuid"`
	// The organization the ACL's referred object belongs to
	ObjectOrgID string `json:"_object_org_id,required" format:"uuid"`
	// The id of the object the ACL applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType ACLObjectType `json:"object_type,required,nullable"`
	// Date of acl creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Id of the group the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	GroupID string `json:"group_id,nullable" format:"uuid"`
	// Permission the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	Permission ACLPermission `json:"permission,nullable"`
	// When setting a permission directly, optionally restricts the permission grant to
	// just the specified object type. Cannot be set alongside a `role_id`.
	RestrictObjectType ACLRestrictObjectType `json:"restrict_object_type,nullable"`
	// Id of the role the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	RoleID string `json:"role_id,nullable" format:"uuid"`
	// Id of the user the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	UserID string  `json:"user_id,nullable" format:"uuid"`
	JSON   aclJSON `json:"-"`
}

// aclJSON contains the JSON metadata for the struct [ACL]
type aclJSON struct {
	ID                 apijson.Field
	ObjectOrgID        apijson.Field
	ObjectID           apijson.Field
	ObjectType         apijson.Field
	Created            apijson.Field
	GroupID            apijson.Field
	Permission         apijson.Field
	RestrictObjectType apijson.Field
	RoleID             apijson.Field
	UserID             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclJSON) RawJSON() string {
	return r.raw
}

// The object type that the ACL applies to
type ACLObjectType string

const (
	ACLObjectTypeOrganization  ACLObjectType = "organization"
	ACLObjectTypeProject       ACLObjectType = "project"
	ACLObjectTypeExperiment    ACLObjectType = "experiment"
	ACLObjectTypeDataset       ACLObjectType = "dataset"
	ACLObjectTypePrompt        ACLObjectType = "prompt"
	ACLObjectTypePromptSession ACLObjectType = "prompt_session"
	ACLObjectTypeGroup         ACLObjectType = "group"
	ACLObjectTypeRole          ACLObjectType = "role"
	ACLObjectTypeOrgMember     ACLObjectType = "org_member"
	ACLObjectTypeProjectLog    ACLObjectType = "project_log"
	ACLObjectTypeOrgProject    ACLObjectType = "org_project"
)

func (r ACLObjectType) IsKnown() bool {
	switch r {
	case ACLObjectTypeOrganization, ACLObjectTypeProject, ACLObjectTypeExperiment, ACLObjectTypeDataset, ACLObjectTypePrompt, ACLObjectTypePromptSession, ACLObjectTypeGroup, ACLObjectTypeRole, ACLObjectTypeOrgMember, ACLObjectTypeProjectLog, ACLObjectTypeOrgProject:
		return true
	}
	return false
}

// Permission the ACL grants. Exactly one of `permission` and `role_id` will be
// provided
type ACLPermission string

const (
	ACLPermissionCreate     ACLPermission = "create"
	ACLPermissionRead       ACLPermission = "read"
	ACLPermissionUpdate     ACLPermission = "update"
	ACLPermissionDelete     ACLPermission = "delete"
	ACLPermissionCreateACLs ACLPermission = "create_acls"
	ACLPermissionReadACLs   ACLPermission = "read_acls"
	ACLPermissionUpdateACLs ACLPermission = "update_acls"
	ACLPermissionDeleteACLs ACLPermission = "delete_acls"
)

func (r ACLPermission) IsKnown() bool {
	switch r {
	case ACLPermissionCreate, ACLPermissionRead, ACLPermissionUpdate, ACLPermissionDelete, ACLPermissionCreateACLs, ACLPermissionReadACLs, ACLPermissionUpdateACLs, ACLPermissionDeleteACLs:
		return true
	}
	return false
}

// When setting a permission directly, optionally restricts the permission grant to
// just the specified object type. Cannot be set alongside a `role_id`.
type ACLRestrictObjectType string

const (
	ACLRestrictObjectTypeOrganization  ACLRestrictObjectType = "organization"
	ACLRestrictObjectTypeProject       ACLRestrictObjectType = "project"
	ACLRestrictObjectTypeExperiment    ACLRestrictObjectType = "experiment"
	ACLRestrictObjectTypeDataset       ACLRestrictObjectType = "dataset"
	ACLRestrictObjectTypePrompt        ACLRestrictObjectType = "prompt"
	ACLRestrictObjectTypePromptSession ACLRestrictObjectType = "prompt_session"
	ACLRestrictObjectTypeGroup         ACLRestrictObjectType = "group"
	ACLRestrictObjectTypeRole          ACLRestrictObjectType = "role"
	ACLRestrictObjectTypeOrgMember     ACLRestrictObjectType = "org_member"
	ACLRestrictObjectTypeProjectLog    ACLRestrictObjectType = "project_log"
	ACLRestrictObjectTypeOrgProject    ACLRestrictObjectType = "org_project"
)

func (r ACLRestrictObjectType) IsKnown() bool {
	switch r {
	case ACLRestrictObjectTypeOrganization, ACLRestrictObjectTypeProject, ACLRestrictObjectTypeExperiment, ACLRestrictObjectTypeDataset, ACLRestrictObjectTypePrompt, ACLRestrictObjectTypePromptSession, ACLRestrictObjectTypeGroup, ACLRestrictObjectTypeRole, ACLRestrictObjectTypeOrgMember, ACLRestrictObjectTypeProjectLog, ACLRestrictObjectTypeOrgProject:
		return true
	}
	return false
}

type APIKey struct {
	// Unique identifier for the api key
	ID string `json:"id,required" format:"uuid"`
	// Name of the api key
	Name        string `json:"name,required"`
	PreviewName string `json:"preview_name,required"`
	// Date of api key creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Unique identifier for the organization
	OrgID string `json:"org_id,nullable" format:"uuid"`
	// Unique identifier for the user
	UserID string     `json:"user_id,nullable" format:"uuid"`
	JSON   apiKeyJSON `json:"-"`
}

// apiKeyJSON contains the JSON metadata for the struct [APIKey]
type apiKeyJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	PreviewName apijson.Field
	Created     apijson.Field
	OrgID       apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyJSON) RawJSON() string {
	return r.raw
}

type CreateAPIKeyOutput struct {
	// Unique identifier for the api key
	ID string `json:"id,required" format:"uuid"`
	// The raw API key. It will only be exposed this one time
	Key string `json:"key,required"`
	// Name of the api key
	Name        string `json:"name,required"`
	PreviewName string `json:"preview_name,required"`
	// Date of api key creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Unique identifier for the organization
	OrgID string `json:"org_id,nullable" format:"uuid"`
	// Unique identifier for the user
	UserID string                 `json:"user_id,nullable" format:"uuid"`
	JSON   createAPIKeyOutputJSON `json:"-"`
}

// createAPIKeyOutputJSON contains the JSON metadata for the struct
// [CreateAPIKeyOutput]
type createAPIKeyOutputJSON struct {
	ID          apijson.Field
	Key         apijson.Field
	Name        apijson.Field
	PreviewName apijson.Field
	Created     apijson.Field
	OrgID       apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateAPIKeyOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createAPIKeyOutputJSON) RawJSON() string {
	return r.raw
}

// Summary of a dataset's data
type DataSummary struct {
	// Total number of records in the dataset
	TotalRecords int64           `json:"total_records,required"`
	JSON         dataSummaryJSON `json:"-"`
}

// dataSummaryJSON contains the JSON metadata for the struct [DataSummary]
type dataSummaryJSON struct {
	TotalRecords apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DataSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataSummaryJSON) RawJSON() string {
	return r.raw
}

type Dataset struct {
	// Unique identifier for the dataset
	ID string `json:"id,required" format:"uuid"`
	// Name of the dataset. Within a project, dataset names are unique
	Name string `json:"name,required"`
	// Date of dataset creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Date of dataset deletion, or null if the dataset is still active
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Textual description of the dataset
	Description string `json:"description,nullable"`
	// User-controlled metadata about the dataset
	Metadata map[string]interface{} `json:"metadata,nullable"`
	// Unique identifier for the project that the dataset belongs under
	ProjectID string `json:"project_id,nullable" format:"uuid"`
	// Identifies the user who created the dataset
	UserID string      `json:"user_id,nullable" format:"uuid"`
	JSON   datasetJSON `json:"-"`
}

// datasetJSON contains the JSON metadata for the struct [Dataset]
type datasetJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Created     apijson.Field
	DeletedAt   apijson.Field
	Description apijson.Field
	Metadata    apijson.Field
	ProjectID   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Dataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetJSON) RawJSON() string {
	return r.raw
}

type DatasetEvent struct {
	// A unique identifier for the dataset event. If you don't provide one, BrainTrust
	// will generate one for you
	ID string `json:"id,required"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the dataset (see the `version`
	// parameter)
	XactID string `json:"_xact_id,required"`
	// The timestamp the dataset event was created
	Created time.Time `json:"created,required" format:"date-time"`
	// Unique identifier for the dataset
	DatasetID string `json:"dataset_id,required" format:"uuid"`
	// The `span_id` of the root of the trace this dataset event belongs to
	RootSpanID string `json:"root_span_id,required"`
	// A unique identifier used to link different dataset events together as part of a
	// full trace. See the
	// [tracing guide](https://www.braintrust.dev/docs/guides/tracing) for full details
	// on tracing
	SpanID string `json:"span_id,required"`
	// The output of your application, including post-processing (an arbitrary, JSON
	// serializable object)
	Expected interface{} `json:"expected"`
	// The argument that uniquely define an input case (an arbitrary, JSON serializable
	// object)
	Input interface{} `json:"input"`
	// A dictionary with additional data about the test example, model outputs, or just
	// about anything else that's relevant, that you can use to help find and analyze
	// examples later. For example, you could log the `prompt`, example's `id`, or
	// anything else that would be useful to slice/dice later. The values in `metadata`
	// can be any JSON-serializable type, but its keys must be strings
	Metadata map[string]interface{} `json:"metadata,nullable"`
	// Unique identifier for the project that the dataset belongs under
	ProjectID string `json:"project_id,nullable" format:"uuid"`
	// A list of tags to log
	Tags []string         `json:"tags,nullable"`
	JSON datasetEventJSON `json:"-"`
}

// datasetEventJSON contains the JSON metadata for the struct [DatasetEvent]
type datasetEventJSON struct {
	ID          apijson.Field
	XactID      apijson.Field
	Created     apijson.Field
	DatasetID   apijson.Field
	RootSpanID  apijson.Field
	SpanID      apijson.Field
	Expected    apijson.Field
	Input       apijson.Field
	Metadata    apijson.Field
	ProjectID   apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetEventJSON) RawJSON() string {
	return r.raw
}

type Experiment struct {
	// Unique identifier for the experiment
	ID string `json:"id,required" format:"uuid"`
	// Name of the experiment. Within a project, experiment names are unique
	Name string `json:"name,required"`
	// Unique identifier for the project that the experiment belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// Whether or not the experiment is public. Public experiments can be viewed by
	// anybody inside or outside the organization
	Public bool `json:"public,required"`
	// Id of default base experiment to compare against when viewing this experiment
	BaseExpID string `json:"base_exp_id,nullable" format:"uuid"`
	// Commit, taken directly from `repo_info.commit`
	Commit string `json:"commit,nullable"`
	// Date of experiment creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Identifier of the linked dataset, or null if the experiment is not linked to a
	// dataset
	DatasetID string `json:"dataset_id,nullable" format:"uuid"`
	// Version number of the linked dataset the experiment was run against. This can be
	// used to reproduce the experiment after the dataset has been modified.
	DatasetVersion string `json:"dataset_version,nullable"`
	// Date of experiment deletion, or null if the experiment is still active
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Textual description of the experiment
	Description string `json:"description,nullable"`
	// User-controlled metadata about the experiment
	Metadata map[string]interface{} `json:"metadata,nullable"`
	// Metadata about the state of the repo when the experiment was created
	RepoInfo RepoInfo `json:"repo_info,nullable"`
	// Identifies the user who created the experiment
	UserID string         `json:"user_id,nullable" format:"uuid"`
	JSON   experimentJSON `json:"-"`
}

// experimentJSON contains the JSON metadata for the struct [Experiment]
type experimentJSON struct {
	ID             apijson.Field
	Name           apijson.Field
	ProjectID      apijson.Field
	Public         apijson.Field
	BaseExpID      apijson.Field
	Commit         apijson.Field
	Created        apijson.Field
	DatasetID      apijson.Field
	DatasetVersion apijson.Field
	DeletedAt      apijson.Field
	Description    apijson.Field
	Metadata       apijson.Field
	RepoInfo       apijson.Field
	UserID         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Experiment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r experimentJSON) RawJSON() string {
	return r.raw
}

type ExperimentEvent struct {
	// A unique identifier for the experiment event. If you don't provide one,
	// BrainTrust will generate one for you
	ID string `json:"id,required"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the experiment (see the
	// `version` parameter)
	XactID string `json:"_xact_id,required"`
	// The timestamp the experiment event was created
	Created time.Time `json:"created,required" format:"date-time"`
	// Unique identifier for the experiment
	ExperimentID string `json:"experiment_id,required" format:"uuid"`
	// Unique identifier for the project that the experiment belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// The `span_id` of the root of the trace this experiment event belongs to
	RootSpanID string `json:"root_span_id,required"`
	// A unique identifier used to link different experiment events together as part of
	// a full trace. See the
	// [tracing guide](https://www.braintrust.dev/docs/guides/tracing) for full details
	// on tracing
	SpanID string `json:"span_id,required"`
	// Context is additional information about the code that produced the experiment
	// event. It is essentially the textual counterpart to `metrics`. Use the
	// `caller_*` attributes to track the location in code which produced the
	// experiment event
	Context ExperimentEventContext `json:"context,nullable"`
	// If the experiment is associated to a dataset, this is the event-level dataset id
	// this experiment event is tied to
	DatasetRecordID string `json:"dataset_record_id,nullable"`
	// The error that occurred, if any.
	Error interface{} `json:"error"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not.
	// Braintrust currently does not compare `output` to `expected` for you, since
	// there are so many different ways to do that correctly. Instead, these values are
	// just used to help you navigate your experiments while digging into analyses.
	// However, we may later use these values to re-score outputs or fine-tune your
	// models
	Expected interface{} `json:"expected"`
	// The arguments that uniquely define a test case (an arbitrary, JSON serializable
	// object). Later on, Braintrust will use the `input` to know whether two test
	// cases are the same between experiments, so they should not contain
	// experiment-specific state. A simple rule of thumb is that if you run the same
	// experiment twice, the `input` should be identical
	Input interface{} `json:"input"`
	// A dictionary with additional data about the test example, model outputs, or just
	// about anything else that's relevant, that you can use to help find and analyze
	// examples later. For example, you could log the `prompt`, example's `id`, or
	// anything else that would be useful to slice/dice later. The values in `metadata`
	// can be any JSON-serializable type, but its keys must be strings
	Metadata map[string]interface{} `json:"metadata,nullable"`
	// Metrics are numerical measurements tracking the execution of the code that
	// produced the experiment event. Use "start" and "end" to track the time span over
	// which the experiment event was produced
	Metrics ExperimentEventMetrics `json:"metrics,nullable"`
	// The output of your application, including post-processing (an arbitrary, JSON
	// serializable object), that allows you to determine whether the result is correct
	// or not. For example, in an app that generates SQL queries, the `output` should
	// be the _result_ of the SQL query generated by the model, not the query itself,
	// because there may be multiple valid queries that answer a single question
	Output interface{} `json:"output"`
	// A dictionary of numeric values (between 0 and 1) to log. The scores should give
	// you a variety of signals that help you determine how accurate the outputs are
	// compared to what you expect and diagnose failures. For example, a summarization
	// app might have one score that tells you how accurate the summary is, and another
	// that measures the word similarity between the generated and grouth truth
	// summary. The word similarity score could help you determine whether the
	// summarization was covering similar concepts or not. You can use these scores to
	// help you sort, filter, and compare experiments
	Scores map[string]float64 `json:"scores,nullable"`
	// Human-identifying attributes of the span, such as name, type, etc.
	SpanAttributes ExperimentEventSpanAttributes `json:"span_attributes,nullable"`
	// An array of the parent `span_ids` of this experiment event. This should be empty
	// for the root span of a trace, and should most often contain just one parent
	// element for subspans
	SpanParents []string `json:"span_parents,nullable"`
	// A list of tags to log
	Tags []string            `json:"tags,nullable"`
	JSON experimentEventJSON `json:"-"`
}

// experimentEventJSON contains the JSON metadata for the struct [ExperimentEvent]
type experimentEventJSON struct {
	ID              apijson.Field
	XactID          apijson.Field
	Created         apijson.Field
	ExperimentID    apijson.Field
	ProjectID       apijson.Field
	RootSpanID      apijson.Field
	SpanID          apijson.Field
	Context         apijson.Field
	DatasetRecordID apijson.Field
	Error           apijson.Field
	Expected        apijson.Field
	Input           apijson.Field
	Metadata        apijson.Field
	Metrics         apijson.Field
	Output          apijson.Field
	Scores          apijson.Field
	SpanAttributes  apijson.Field
	SpanParents     apijson.Field
	Tags            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ExperimentEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r experimentEventJSON) RawJSON() string {
	return r.raw
}

// Context is additional information about the code that produced the experiment
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the
// experiment event
type ExperimentEventContext struct {
	// Name of the file in code where the experiment event was created
	CallerFilename string `json:"caller_filename,nullable"`
	// The function in code which created the experiment event
	CallerFunctionname string `json:"caller_functionname,nullable"`
	// Line of code where the experiment event was created
	CallerLineno int64                      `json:"caller_lineno,nullable"`
	ExtraFields  map[string]interface{}     `json:"-,extras"`
	JSON         experimentEventContextJSON `json:"-"`
}

// experimentEventContextJSON contains the JSON metadata for the struct
// [ExperimentEventContext]
type experimentEventContextJSON struct {
	CallerFilename     apijson.Field
	CallerFunctionname apijson.Field
	CallerLineno       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ExperimentEventContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r experimentEventContextJSON) RawJSON() string {
	return r.raw
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the experiment event. Use "start" and "end" to track the time span over
// which the experiment event was produced
type ExperimentEventMetrics struct {
	// The number of tokens in the completion generated by the model (only set if this
	// is an LLM span)
	CompletionTokens int64 `json:"completion_tokens,nullable"`
	// A unix timestamp recording when the section of code which produced the
	// experiment event finished
	End float64 `json:"end,nullable"`
	// The number of tokens in the prompt used to generate the experiment event (only
	// set if this is an LLM span)
	PromptTokens int64 `json:"prompt_tokens,nullable"`
	// A unix timestamp recording when the section of code which produced the
	// experiment event started
	Start float64 `json:"start,nullable"`
	// The total number of tokens in the input and output of the experiment event.
	Tokens      int64                      `json:"tokens,nullable"`
	ExtraFields map[string]interface{}     `json:"-,extras"`
	JSON        experimentEventMetricsJSON `json:"-"`
}

// experimentEventMetricsJSON contains the JSON metadata for the struct
// [ExperimentEventMetrics]
type experimentEventMetricsJSON struct {
	CompletionTokens apijson.Field
	End              apijson.Field
	PromptTokens     apijson.Field
	Start            apijson.Field
	Tokens           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ExperimentEventMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r experimentEventMetricsJSON) RawJSON() string {
	return r.raw
}

// Human-identifying attributes of the span, such as name, type, etc.
type ExperimentEventSpanAttributes struct {
	// Name of the span, for display purposes only
	Name string `json:"name,nullable"`
	// Type of the span, for display purposes only
	Type        ExperimentEventSpanAttributesType `json:"type,nullable"`
	ExtraFields map[string]interface{}            `json:"-,extras"`
	JSON        experimentEventSpanAttributesJSON `json:"-"`
}

// experimentEventSpanAttributesJSON contains the JSON metadata for the struct
// [ExperimentEventSpanAttributes]
type experimentEventSpanAttributesJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExperimentEventSpanAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r experimentEventSpanAttributesJSON) RawJSON() string {
	return r.raw
}

// Type of the span, for display purposes only
type ExperimentEventSpanAttributesType string

const (
	ExperimentEventSpanAttributesTypeLlm      ExperimentEventSpanAttributesType = "llm"
	ExperimentEventSpanAttributesTypeScore    ExperimentEventSpanAttributesType = "score"
	ExperimentEventSpanAttributesTypeFunction ExperimentEventSpanAttributesType = "function"
	ExperimentEventSpanAttributesTypeEval     ExperimentEventSpanAttributesType = "eval"
	ExperimentEventSpanAttributesTypeTask     ExperimentEventSpanAttributesType = "task"
	ExperimentEventSpanAttributesTypeTool     ExperimentEventSpanAttributesType = "tool"
)

func (r ExperimentEventSpanAttributesType) IsKnown() bool {
	switch r {
	case ExperimentEventSpanAttributesTypeLlm, ExperimentEventSpanAttributesTypeScore, ExperimentEventSpanAttributesTypeFunction, ExperimentEventSpanAttributesTypeEval, ExperimentEventSpanAttributesTypeTask, ExperimentEventSpanAttributesTypeTool:
		return true
	}
	return false
}

type FeedbackDatasetItemParam struct {
	// The id of the dataset event to log feedback for. This is the row `id` returned
	// by `POST /v1/dataset/{dataset_id}/insert`
	ID param.Field[string] `json:"id,required"`
	// An optional comment string to log about the dataset event
	Comment param.Field[string] `json:"comment"`
	// A dictionary with additional data about the feedback. If you have a `user_id`,
	// you can log it here and access it in the Braintrust UI.
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// The source of the feedback. Must be one of "external" (default), "app", or "api"
	Source param.Field[FeedbackDatasetItemSource] `json:"source"`
}

func (r FeedbackDatasetItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The source of the feedback. Must be one of "external" (default), "app", or "api"
type FeedbackDatasetItemSource string

const (
	FeedbackDatasetItemSourceApp      FeedbackDatasetItemSource = "app"
	FeedbackDatasetItemSourceAPI      FeedbackDatasetItemSource = "api"
	FeedbackDatasetItemSourceExternal FeedbackDatasetItemSource = "external"
)

func (r FeedbackDatasetItemSource) IsKnown() bool {
	switch r {
	case FeedbackDatasetItemSourceApp, FeedbackDatasetItemSourceAPI, FeedbackDatasetItemSourceExternal:
		return true
	}
	return false
}

type FeedbackExperimentItemParam struct {
	// The id of the experiment event to log feedback for. This is the row `id`
	// returned by `POST /v1/experiment/{experiment_id}/insert`
	ID param.Field[string] `json:"id,required"`
	// An optional comment string to log about the experiment event
	Comment param.Field[string] `json:"comment"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not
	Expected param.Field[interface{}] `json:"expected"`
	// A dictionary with additional data about the feedback. If you have a `user_id`,
	// you can log it here and access it in the Braintrust UI.
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// A dictionary of numeric values (between 0 and 1) to log. These scores will be
	// merged into the existing scores for the experiment event
	Scores param.Field[map[string]float64] `json:"scores"`
	// The source of the feedback. Must be one of "external" (default), "app", or "api"
	Source param.Field[FeedbackExperimentItemSource] `json:"source"`
}

func (r FeedbackExperimentItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The source of the feedback. Must be one of "external" (default), "app", or "api"
type FeedbackExperimentItemSource string

const (
	FeedbackExperimentItemSourceApp      FeedbackExperimentItemSource = "app"
	FeedbackExperimentItemSourceAPI      FeedbackExperimentItemSource = "api"
	FeedbackExperimentItemSourceExternal FeedbackExperimentItemSource = "external"
)

func (r FeedbackExperimentItemSource) IsKnown() bool {
	switch r {
	case FeedbackExperimentItemSourceApp, FeedbackExperimentItemSourceAPI, FeedbackExperimentItemSourceExternal:
		return true
	}
	return false
}

type FeedbackProjectLogsItemParam struct {
	// The id of the project logs event to log feedback for. This is the row `id`
	// returned by `POST /v1/project_logs/{project_id}/insert`
	ID param.Field[string] `json:"id,required"`
	// An optional comment string to log about the project logs event
	Comment param.Field[string] `json:"comment"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not
	Expected param.Field[interface{}] `json:"expected"`
	// A dictionary with additional data about the feedback. If you have a `user_id`,
	// you can log it here and access it in the Braintrust UI.
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// A dictionary of numeric values (between 0 and 1) to log. These scores will be
	// merged into the existing scores for the project logs event
	Scores param.Field[map[string]float64] `json:"scores"`
	// The source of the feedback. Must be one of "external" (default), "app", or "api"
	Source param.Field[FeedbackProjectLogsItemSource] `json:"source"`
}

func (r FeedbackProjectLogsItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The source of the feedback. Must be one of "external" (default), "app", or "api"
type FeedbackProjectLogsItemSource string

const (
	FeedbackProjectLogsItemSourceApp      FeedbackProjectLogsItemSource = "app"
	FeedbackProjectLogsItemSourceAPI      FeedbackProjectLogsItemSource = "api"
	FeedbackProjectLogsItemSourceExternal FeedbackProjectLogsItemSource = "external"
)

func (r FeedbackProjectLogsItemSource) IsKnown() bool {
	switch r {
	case FeedbackProjectLogsItemSourceApp, FeedbackProjectLogsItemSourceAPI, FeedbackProjectLogsItemSourceExternal:
		return true
	}
	return false
}

type FeedbackResponseSchema struct {
	Status FeedbackResponseSchemaStatus `json:"status,required"`
	JSON   feedbackResponseSchemaJSON   `json:"-"`
}

// feedbackResponseSchemaJSON contains the JSON metadata for the struct
// [FeedbackResponseSchema]
type feedbackResponseSchemaJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FeedbackResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r feedbackResponseSchemaJSON) RawJSON() string {
	return r.raw
}

type FeedbackResponseSchemaStatus string

const (
	FeedbackResponseSchemaStatusSuccess FeedbackResponseSchemaStatus = "success"
)

func (r FeedbackResponseSchemaStatus) IsKnown() bool {
	switch r {
	case FeedbackResponseSchemaStatusSuccess:
		return true
	}
	return false
}

type FetchDatasetEventsResponse struct {
	// A list of fetched events
	Events []DatasetEvent `json:"events,required"`
	// Pagination cursor
	//
	// Pass this string directly as the `cursor` param to your next fetch request to
	// get the next page of results. Not provided if the returned result set is empty.
	Cursor string                         `json:"cursor,nullable"`
	JSON   fetchDatasetEventsResponseJSON `json:"-"`
}

// fetchDatasetEventsResponseJSON contains the JSON metadata for the struct
// [FetchDatasetEventsResponse]
type fetchDatasetEventsResponseJSON struct {
	Events      apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FetchDatasetEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fetchDatasetEventsResponseJSON) RawJSON() string {
	return r.raw
}

type FetchExperimentEventsResponse struct {
	// A list of fetched events
	Events []ExperimentEvent `json:"events,required"`
	// Pagination cursor
	//
	// Pass this string directly as the `cursor` param to your next fetch request to
	// get the next page of results. Not provided if the returned result set is empty.
	Cursor string                            `json:"cursor,nullable"`
	JSON   fetchExperimentEventsResponseJSON `json:"-"`
}

// fetchExperimentEventsResponseJSON contains the JSON metadata for the struct
// [FetchExperimentEventsResponse]
type fetchExperimentEventsResponseJSON struct {
	Events      apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FetchExperimentEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fetchExperimentEventsResponseJSON) RawJSON() string {
	return r.raw
}

type FetchProjectLogsEventsResponse struct {
	// A list of fetched events
	Events []ProjectLogsEvent `json:"events,required"`
	// Pagination cursor
	//
	// Pass this string directly as the `cursor` param to your next fetch request to
	// get the next page of results. Not provided if the returned result set is empty.
	Cursor string                             `json:"cursor,nullable"`
	JSON   fetchProjectLogsEventsResponseJSON `json:"-"`
}

// fetchProjectLogsEventsResponseJSON contains the JSON metadata for the struct
// [FetchProjectLogsEventsResponse]
type fetchProjectLogsEventsResponseJSON struct {
	Events      apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FetchProjectLogsEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fetchProjectLogsEventsResponseJSON) RawJSON() string {
	return r.raw
}

type Function struct {
	// Unique identifier for the prompt
	ID string `json:"id,required" format:"uuid"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the prompt (see the `version`
	// parameter)
	XactID       string               `json:"_xact_id,required"`
	FunctionData FunctionFunctionData `json:"function_data,required"`
	// A literal 'p' which identifies the object as a project prompt
	LogID FunctionLogID `json:"log_id,required"`
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
	// JSON schema for the function's parameters and return type
	FunctionSchema FunctionFunctionSchema `json:"function_schema,nullable"`
	FunctionType   FunctionFunctionType   `json:"function_type,nullable"`
	// User-controlled metadata about the prompt
	Metadata map[string]interface{} `json:"metadata,nullable"`
	Origin   FunctionOrigin         `json:"origin,nullable"`
	// The prompt, model, and its parameters
	PromptData PromptData `json:"prompt_data,nullable"`
	// A list of tags for the prompt
	Tags []string     `json:"tags,nullable"`
	JSON functionJSON `json:"-"`
}

// functionJSON contains the JSON metadata for the struct [Function]
type functionJSON struct {
	ID             apijson.Field
	XactID         apijson.Field
	FunctionData   apijson.Field
	LogID          apijson.Field
	Name           apijson.Field
	OrgID          apijson.Field
	ProjectID      apijson.Field
	Slug           apijson.Field
	Created        apijson.Field
	Description    apijson.Field
	FunctionSchema apijson.Field
	FunctionType   apijson.Field
	Metadata       apijson.Field
	Origin         apijson.Field
	PromptData     apijson.Field
	Tags           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Function) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionJSON) RawJSON() string {
	return r.raw
}

type FunctionFunctionData struct {
	Type FunctionFunctionDataType `json:"type,required"`
	// This field can have the runtime type of [FunctionFunctionDataCodeData].
	Data  interface{}              `json:"data,required"`
	Name  string                   `json:"name"`
	JSON  functionFunctionDataJSON `json:"-"`
	union FunctionFunctionDataUnion
}

// functionFunctionDataJSON contains the JSON metadata for the struct
// [FunctionFunctionData]
type functionFunctionDataJSON struct {
	Type        apijson.Field
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r functionFunctionDataJSON) RawJSON() string {
	return r.raw
}

func (r *FunctionFunctionData) UnmarshalJSON(data []byte) (err error) {
	*r = FunctionFunctionData{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FunctionFunctionDataUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [shared.FunctionFunctionDataPrompt],
// [shared.FunctionFunctionDataCode], [shared.FunctionFunctionDataGlobal].
func (r FunctionFunctionData) AsUnion() FunctionFunctionDataUnion {
	return r.union
}

// Union satisfied by [shared.FunctionFunctionDataPrompt],
// [shared.FunctionFunctionDataCode] or [shared.FunctionFunctionDataGlobal].
type FunctionFunctionDataUnion interface {
	implementsSharedFunctionFunctionData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionFunctionDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataPrompt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataCode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataGlobal{}),
		},
	)
}

type FunctionFunctionDataPrompt struct {
	Type FunctionFunctionDataPromptType `json:"type,required"`
	JSON functionFunctionDataPromptJSON `json:"-"`
}

// functionFunctionDataPromptJSON contains the JSON metadata for the struct
// [FunctionFunctionDataPrompt]
type functionFunctionDataPromptJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataPromptJSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataPrompt) implementsSharedFunctionFunctionData() {}

type FunctionFunctionDataPromptType string

const (
	FunctionFunctionDataPromptTypePrompt FunctionFunctionDataPromptType = "prompt"
)

func (r FunctionFunctionDataPromptType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataPromptTypePrompt:
		return true
	}
	return false
}

type FunctionFunctionDataCode struct {
	Data FunctionFunctionDataCodeData `json:"data,required"`
	Type FunctionFunctionDataCodeType `json:"type,required"`
	JSON functionFunctionDataCodeJSON `json:"-"`
}

// functionFunctionDataCodeJSON contains the JSON metadata for the struct
// [FunctionFunctionDataCode]
type functionFunctionDataCodeJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataCodeJSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataCode) implementsSharedFunctionFunctionData() {}

type FunctionFunctionDataCodeData struct {
	Type FunctionFunctionDataCodeDataType `json:"type,required"`
	// This field can have the runtime type of
	// [FunctionFunctionDataCodeDataBundleRuntimeContext],
	// [FunctionFunctionDataCodeDataInlineRuntimeContext].
	RuntimeContext interface{} `json:"runtime_context"`
	// This field can have the runtime type of
	// [FunctionFunctionDataCodeDataBundleLocation].
	Location interface{} `json:"location,required"`
	BundleID string      `json:"bundle_id"`
	// A preview of the code
	Preview string                           `json:"preview,nullable"`
	Code    string                           `json:"code"`
	JSON    functionFunctionDataCodeDataJSON `json:"-"`
	union   FunctionFunctionDataCodeDataUnion
}

// functionFunctionDataCodeDataJSON contains the JSON metadata for the struct
// [FunctionFunctionDataCodeData]
type functionFunctionDataCodeDataJSON struct {
	Type           apijson.Field
	RuntimeContext apijson.Field
	Location       apijson.Field
	BundleID       apijson.Field
	Preview        apijson.Field
	Code           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r functionFunctionDataCodeDataJSON) RawJSON() string {
	return r.raw
}

func (r *FunctionFunctionDataCodeData) UnmarshalJSON(data []byte) (err error) {
	*r = FunctionFunctionDataCodeData{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FunctionFunctionDataCodeDataUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.FunctionFunctionDataCodeDataBundle],
// [shared.FunctionFunctionDataCodeDataInline].
func (r FunctionFunctionDataCodeData) AsUnion() FunctionFunctionDataCodeDataUnion {
	return r.union
}

// Union satisfied by [shared.FunctionFunctionDataCodeDataBundle] or
// [shared.FunctionFunctionDataCodeDataInline].
type FunctionFunctionDataCodeDataUnion interface {
	implementsSharedFunctionFunctionDataCodeData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionFunctionDataCodeDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataCodeDataBundle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataCodeDataInline{}),
		},
	)
}

type FunctionFunctionDataCodeDataBundle struct {
	BundleID       string                                           `json:"bundle_id,required"`
	Location       FunctionFunctionDataCodeDataBundleLocation       `json:"location,required"`
	RuntimeContext FunctionFunctionDataCodeDataBundleRuntimeContext `json:"runtime_context,required"`
	Type           FunctionFunctionDataCodeDataBundleType           `json:"type,required"`
	// A preview of the code
	Preview string                                 `json:"preview,nullable"`
	JSON    functionFunctionDataCodeDataBundleJSON `json:"-"`
}

// functionFunctionDataCodeDataBundleJSON contains the JSON metadata for the struct
// [FunctionFunctionDataCodeDataBundle]
type functionFunctionDataCodeDataBundleJSON struct {
	BundleID       apijson.Field
	Location       apijson.Field
	RuntimeContext apijson.Field
	Type           apijson.Field
	Preview        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FunctionFunctionDataCodeDataBundle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataCodeDataBundleJSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataCodeDataBundle) implementsSharedFunctionFunctionDataCodeData() {}

type FunctionFunctionDataCodeDataBundleLocation struct {
	Type     FunctionFunctionDataCodeDataBundleLocationType `json:"type,required"`
	EvalName string                                         `json:"eval_name"`
	// This field can have the runtime type of
	// [FunctionFunctionDataCodeDataBundleLocationExperimentPosition].
	Position interface{}                                    `json:"position,required"`
	Index    int64                                          `json:"index"`
	JSON     functionFunctionDataCodeDataBundleLocationJSON `json:"-"`
	union    FunctionFunctionDataCodeDataBundleLocationUnion
}

// functionFunctionDataCodeDataBundleLocationJSON contains the JSON metadata for
// the struct [FunctionFunctionDataCodeDataBundleLocation]
type functionFunctionDataCodeDataBundleLocationJSON struct {
	Type        apijson.Field
	EvalName    apijson.Field
	Position    apijson.Field
	Index       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r functionFunctionDataCodeDataBundleLocationJSON) RawJSON() string {
	return r.raw
}

func (r *FunctionFunctionDataCodeDataBundleLocation) UnmarshalJSON(data []byte) (err error) {
	*r = FunctionFunctionDataCodeDataBundleLocation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FunctionFunctionDataCodeDataBundleLocationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.FunctionFunctionDataCodeDataBundleLocationExperiment],
// [shared.FunctionFunctionDataCodeDataBundleLocationFunction].
func (r FunctionFunctionDataCodeDataBundleLocation) AsUnion() FunctionFunctionDataCodeDataBundleLocationUnion {
	return r.union
}

// Union satisfied by [shared.FunctionFunctionDataCodeDataBundleLocationExperiment]
// or [shared.FunctionFunctionDataCodeDataBundleLocationFunction].
type FunctionFunctionDataCodeDataBundleLocationUnion interface {
	implementsSharedFunctionFunctionDataCodeDataBundleLocation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionFunctionDataCodeDataBundleLocationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataCodeDataBundleLocationExperiment{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataCodeDataBundleLocationFunction{}),
		},
	)
}

type FunctionFunctionDataCodeDataBundleLocationExperiment struct {
	EvalName string                                                       `json:"eval_name,required"`
	Position FunctionFunctionDataCodeDataBundleLocationExperimentPosition `json:"position,required"`
	Type     FunctionFunctionDataCodeDataBundleLocationExperimentType     `json:"type,required"`
	JSON     functionFunctionDataCodeDataBundleLocationExperimentJSON     `json:"-"`
}

// functionFunctionDataCodeDataBundleLocationExperimentJSON contains the JSON
// metadata for the struct [FunctionFunctionDataCodeDataBundleLocationExperiment]
type functionFunctionDataCodeDataBundleLocationExperimentJSON struct {
	EvalName    apijson.Field
	Position    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataCodeDataBundleLocationExperiment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataCodeDataBundleLocationExperimentJSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataCodeDataBundleLocationExperiment) implementsSharedFunctionFunctionDataCodeDataBundleLocation() {
}

type FunctionFunctionDataCodeDataBundleLocationExperimentPosition struct {
	Type  FunctionFunctionDataCodeDataBundleLocationExperimentPositionType `json:"type,required"`
	Index int64                                                            `json:"index"`
	JSON  functionFunctionDataCodeDataBundleLocationExperimentPositionJSON `json:"-"`
	union FunctionFunctionDataCodeDataBundleLocationExperimentPositionUnion
}

// functionFunctionDataCodeDataBundleLocationExperimentPositionJSON contains the
// JSON metadata for the struct
// [FunctionFunctionDataCodeDataBundleLocationExperimentPosition]
type functionFunctionDataCodeDataBundleLocationExperimentPositionJSON struct {
	Type        apijson.Field
	Index       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r functionFunctionDataCodeDataBundleLocationExperimentPositionJSON) RawJSON() string {
	return r.raw
}

func (r *FunctionFunctionDataCodeDataBundleLocationExperimentPosition) UnmarshalJSON(data []byte) (err error) {
	*r = FunctionFunctionDataCodeDataBundleLocationExperimentPosition{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [FunctionFunctionDataCodeDataBundleLocationExperimentPositionUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.FunctionFunctionDataCodeDataBundleLocationExperimentPositionType],
// [shared.FunctionFunctionDataCodeDataBundleLocationExperimentPositionScorer].
func (r FunctionFunctionDataCodeDataBundleLocationExperimentPosition) AsUnion() FunctionFunctionDataCodeDataBundleLocationExperimentPositionUnion {
	return r.union
}

// Union satisfied by
// [shared.FunctionFunctionDataCodeDataBundleLocationExperimentPositionType] or
// [shared.FunctionFunctionDataCodeDataBundleLocationExperimentPositionScorer].
type FunctionFunctionDataCodeDataBundleLocationExperimentPositionUnion interface {
	implementsSharedFunctionFunctionDataCodeDataBundleLocationExperimentPosition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FunctionFunctionDataCodeDataBundleLocationExperimentPositionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataCodeDataBundleLocationExperimentPositionType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FunctionFunctionDataCodeDataBundleLocationExperimentPositionScorer{}),
		},
	)
}

type FunctionFunctionDataCodeDataBundleLocationExperimentPositionType struct {
	Type FunctionFunctionDataCodeDataBundleLocationExperimentPositionTypeType `json:"type,required"`
	JSON functionFunctionDataCodeDataBundleLocationExperimentPositionTypeJSON `json:"-"`
}

// functionFunctionDataCodeDataBundleLocationExperimentPositionTypeJSON contains
// the JSON metadata for the struct
// [FunctionFunctionDataCodeDataBundleLocationExperimentPositionType]
type functionFunctionDataCodeDataBundleLocationExperimentPositionTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataCodeDataBundleLocationExperimentPositionType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataCodeDataBundleLocationExperimentPositionTypeJSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataCodeDataBundleLocationExperimentPositionType) implementsSharedFunctionFunctionDataCodeDataBundleLocationExperimentPosition() {
}

type FunctionFunctionDataCodeDataBundleLocationExperimentPositionTypeType string

const (
	FunctionFunctionDataCodeDataBundleLocationExperimentPositionTypeTypeTask FunctionFunctionDataCodeDataBundleLocationExperimentPositionTypeType = "task"
)

func (r FunctionFunctionDataCodeDataBundleLocationExperimentPositionTypeType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataCodeDataBundleLocationExperimentPositionTypeTypeTask:
		return true
	}
	return false
}

type FunctionFunctionDataCodeDataBundleLocationExperimentPositionScorer struct {
	Index int64                                                                  `json:"index,required"`
	Type  FunctionFunctionDataCodeDataBundleLocationExperimentPositionScorerType `json:"type,required"`
	JSON  functionFunctionDataCodeDataBundleLocationExperimentPositionScorerJSON `json:"-"`
}

// functionFunctionDataCodeDataBundleLocationExperimentPositionScorerJSON contains
// the JSON metadata for the struct
// [FunctionFunctionDataCodeDataBundleLocationExperimentPositionScorer]
type functionFunctionDataCodeDataBundleLocationExperimentPositionScorerJSON struct {
	Index       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataCodeDataBundleLocationExperimentPositionScorer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataCodeDataBundleLocationExperimentPositionScorerJSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataCodeDataBundleLocationExperimentPositionScorer) implementsSharedFunctionFunctionDataCodeDataBundleLocationExperimentPosition() {
}

type FunctionFunctionDataCodeDataBundleLocationExperimentPositionScorerType string

const (
	FunctionFunctionDataCodeDataBundleLocationExperimentPositionScorerTypeScorer FunctionFunctionDataCodeDataBundleLocationExperimentPositionScorerType = "scorer"
)

func (r FunctionFunctionDataCodeDataBundleLocationExperimentPositionScorerType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataCodeDataBundleLocationExperimentPositionScorerTypeScorer:
		return true
	}
	return false
}

type FunctionFunctionDataCodeDataBundleLocationExperimentType string

const (
	FunctionFunctionDataCodeDataBundleLocationExperimentTypeExperiment FunctionFunctionDataCodeDataBundleLocationExperimentType = "experiment"
)

func (r FunctionFunctionDataCodeDataBundleLocationExperimentType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataCodeDataBundleLocationExperimentTypeExperiment:
		return true
	}
	return false
}

type FunctionFunctionDataCodeDataBundleLocationFunction struct {
	Index int64                                                  `json:"index,required"`
	Type  FunctionFunctionDataCodeDataBundleLocationFunctionType `json:"type,required"`
	JSON  functionFunctionDataCodeDataBundleLocationFunctionJSON `json:"-"`
}

// functionFunctionDataCodeDataBundleLocationFunctionJSON contains the JSON
// metadata for the struct [FunctionFunctionDataCodeDataBundleLocationFunction]
type functionFunctionDataCodeDataBundleLocationFunctionJSON struct {
	Index       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataCodeDataBundleLocationFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataCodeDataBundleLocationFunctionJSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataCodeDataBundleLocationFunction) implementsSharedFunctionFunctionDataCodeDataBundleLocation() {
}

type FunctionFunctionDataCodeDataBundleLocationFunctionType string

const (
	FunctionFunctionDataCodeDataBundleLocationFunctionTypeFunction FunctionFunctionDataCodeDataBundleLocationFunctionType = "function"
)

func (r FunctionFunctionDataCodeDataBundleLocationFunctionType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataCodeDataBundleLocationFunctionTypeFunction:
		return true
	}
	return false
}

type FunctionFunctionDataCodeDataBundleLocationType string

const (
	FunctionFunctionDataCodeDataBundleLocationTypeExperiment FunctionFunctionDataCodeDataBundleLocationType = "experiment"
	FunctionFunctionDataCodeDataBundleLocationTypeFunction   FunctionFunctionDataCodeDataBundleLocationType = "function"
)

func (r FunctionFunctionDataCodeDataBundleLocationType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataCodeDataBundleLocationTypeExperiment, FunctionFunctionDataCodeDataBundleLocationTypeFunction:
		return true
	}
	return false
}

type FunctionFunctionDataCodeDataBundleRuntimeContext struct {
	Runtime FunctionFunctionDataCodeDataBundleRuntimeContextRuntime `json:"runtime,required"`
	Version string                                                  `json:"version,required"`
	JSON    functionFunctionDataCodeDataBundleRuntimeContextJSON    `json:"-"`
}

// functionFunctionDataCodeDataBundleRuntimeContextJSON contains the JSON metadata
// for the struct [FunctionFunctionDataCodeDataBundleRuntimeContext]
type functionFunctionDataCodeDataBundleRuntimeContextJSON struct {
	Runtime     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataCodeDataBundleRuntimeContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataCodeDataBundleRuntimeContextJSON) RawJSON() string {
	return r.raw
}

type FunctionFunctionDataCodeDataBundleRuntimeContextRuntime string

const (
	FunctionFunctionDataCodeDataBundleRuntimeContextRuntimeNode   FunctionFunctionDataCodeDataBundleRuntimeContextRuntime = "node"
	FunctionFunctionDataCodeDataBundleRuntimeContextRuntimePython FunctionFunctionDataCodeDataBundleRuntimeContextRuntime = "python"
)

func (r FunctionFunctionDataCodeDataBundleRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionFunctionDataCodeDataBundleRuntimeContextRuntimeNode, FunctionFunctionDataCodeDataBundleRuntimeContextRuntimePython:
		return true
	}
	return false
}

type FunctionFunctionDataCodeDataBundleType string

const (
	FunctionFunctionDataCodeDataBundleTypeBundle FunctionFunctionDataCodeDataBundleType = "bundle"
)

func (r FunctionFunctionDataCodeDataBundleType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataCodeDataBundleTypeBundle:
		return true
	}
	return false
}

type FunctionFunctionDataCodeDataInline struct {
	Code           string                                           `json:"code,required"`
	RuntimeContext FunctionFunctionDataCodeDataInlineRuntimeContext `json:"runtime_context,required"`
	Type           FunctionFunctionDataCodeDataInlineType           `json:"type,required"`
	JSON           functionFunctionDataCodeDataInlineJSON           `json:"-"`
}

// functionFunctionDataCodeDataInlineJSON contains the JSON metadata for the struct
// [FunctionFunctionDataCodeDataInline]
type functionFunctionDataCodeDataInlineJSON struct {
	Code           apijson.Field
	RuntimeContext apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FunctionFunctionDataCodeDataInline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataCodeDataInlineJSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataCodeDataInline) implementsSharedFunctionFunctionDataCodeData() {}

type FunctionFunctionDataCodeDataInlineRuntimeContext struct {
	Runtime FunctionFunctionDataCodeDataInlineRuntimeContextRuntime `json:"runtime,required"`
	Version string                                                  `json:"version,required"`
	JSON    functionFunctionDataCodeDataInlineRuntimeContextJSON    `json:"-"`
}

// functionFunctionDataCodeDataInlineRuntimeContextJSON contains the JSON metadata
// for the struct [FunctionFunctionDataCodeDataInlineRuntimeContext]
type functionFunctionDataCodeDataInlineRuntimeContextJSON struct {
	Runtime     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataCodeDataInlineRuntimeContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataCodeDataInlineRuntimeContextJSON) RawJSON() string {
	return r.raw
}

type FunctionFunctionDataCodeDataInlineRuntimeContextRuntime string

const (
	FunctionFunctionDataCodeDataInlineRuntimeContextRuntimeNode   FunctionFunctionDataCodeDataInlineRuntimeContextRuntime = "node"
	FunctionFunctionDataCodeDataInlineRuntimeContextRuntimePython FunctionFunctionDataCodeDataInlineRuntimeContextRuntime = "python"
)

func (r FunctionFunctionDataCodeDataInlineRuntimeContextRuntime) IsKnown() bool {
	switch r {
	case FunctionFunctionDataCodeDataInlineRuntimeContextRuntimeNode, FunctionFunctionDataCodeDataInlineRuntimeContextRuntimePython:
		return true
	}
	return false
}

type FunctionFunctionDataCodeDataInlineType string

const (
	FunctionFunctionDataCodeDataInlineTypeInline FunctionFunctionDataCodeDataInlineType = "inline"
)

func (r FunctionFunctionDataCodeDataInlineType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataCodeDataInlineTypeInline:
		return true
	}
	return false
}

type FunctionFunctionDataCodeDataType string

const (
	FunctionFunctionDataCodeDataTypeBundle FunctionFunctionDataCodeDataType = "bundle"
	FunctionFunctionDataCodeDataTypeInline FunctionFunctionDataCodeDataType = "inline"
)

func (r FunctionFunctionDataCodeDataType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataCodeDataTypeBundle, FunctionFunctionDataCodeDataTypeInline:
		return true
	}
	return false
}

type FunctionFunctionDataCodeType string

const (
	FunctionFunctionDataCodeTypeCode FunctionFunctionDataCodeType = "code"
)

func (r FunctionFunctionDataCodeType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataCodeTypeCode:
		return true
	}
	return false
}

type FunctionFunctionDataGlobal struct {
	Name string                         `json:"name,required"`
	Type FunctionFunctionDataGlobalType `json:"type,required"`
	JSON functionFunctionDataGlobalJSON `json:"-"`
}

// functionFunctionDataGlobalJSON contains the JSON metadata for the struct
// [FunctionFunctionDataGlobal]
type functionFunctionDataGlobalJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionDataGlobal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionDataGlobalJSON) RawJSON() string {
	return r.raw
}

func (r FunctionFunctionDataGlobal) implementsSharedFunctionFunctionData() {}

type FunctionFunctionDataGlobalType string

const (
	FunctionFunctionDataGlobalTypeGlobal FunctionFunctionDataGlobalType = "global"
)

func (r FunctionFunctionDataGlobalType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataGlobalTypeGlobal:
		return true
	}
	return false
}

type FunctionFunctionDataType string

const (
	FunctionFunctionDataTypePrompt FunctionFunctionDataType = "prompt"
	FunctionFunctionDataTypeCode   FunctionFunctionDataType = "code"
	FunctionFunctionDataTypeGlobal FunctionFunctionDataType = "global"
)

func (r FunctionFunctionDataType) IsKnown() bool {
	switch r {
	case FunctionFunctionDataTypePrompt, FunctionFunctionDataTypeCode, FunctionFunctionDataTypeGlobal:
		return true
	}
	return false
}

// A literal 'p' which identifies the object as a project prompt
type FunctionLogID string

const (
	FunctionLogIDP FunctionLogID = "p"
)

func (r FunctionLogID) IsKnown() bool {
	switch r {
	case FunctionLogIDP:
		return true
	}
	return false
}

// JSON schema for the function's parameters and return type
type FunctionFunctionSchema struct {
	Parameters interface{}                `json:"parameters"`
	Returns    interface{}                `json:"returns"`
	JSON       functionFunctionSchemaJSON `json:"-"`
}

// functionFunctionSchemaJSON contains the JSON metadata for the struct
// [FunctionFunctionSchema]
type functionFunctionSchemaJSON struct {
	Parameters  apijson.Field
	Returns     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionFunctionSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionFunctionSchemaJSON) RawJSON() string {
	return r.raw
}

type FunctionFunctionType string

const (
	FunctionFunctionTypeLlm    FunctionFunctionType = "llm"
	FunctionFunctionTypeScorer FunctionFunctionType = "scorer"
	FunctionFunctionTypeTask   FunctionFunctionType = "task"
	FunctionFunctionTypeTool   FunctionFunctionType = "tool"
)

func (r FunctionFunctionType) IsKnown() bool {
	switch r {
	case FunctionFunctionTypeLlm, FunctionFunctionTypeScorer, FunctionFunctionTypeTask, FunctionFunctionTypeTool:
		return true
	}
	return false
}

type FunctionOrigin struct {
	// Id of the object the function is originating from
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType FunctionOriginObjectType `json:"object_type,required,nullable"`
	// The function exists for internal purposes and should not be displayed in the
	// list of functions.
	Internal bool               `json:"internal,nullable"`
	JSON     functionOriginJSON `json:"-"`
}

// functionOriginJSON contains the JSON metadata for the struct [FunctionOrigin]
type functionOriginJSON struct {
	ObjectID    apijson.Field
	ObjectType  apijson.Field
	Internal    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionOriginJSON) RawJSON() string {
	return r.raw
}

// The object type that the ACL applies to
type FunctionOriginObjectType string

const (
	FunctionOriginObjectTypeOrganization  FunctionOriginObjectType = "organization"
	FunctionOriginObjectTypeProject       FunctionOriginObjectType = "project"
	FunctionOriginObjectTypeExperiment    FunctionOriginObjectType = "experiment"
	FunctionOriginObjectTypeDataset       FunctionOriginObjectType = "dataset"
	FunctionOriginObjectTypePrompt        FunctionOriginObjectType = "prompt"
	FunctionOriginObjectTypePromptSession FunctionOriginObjectType = "prompt_session"
	FunctionOriginObjectTypeGroup         FunctionOriginObjectType = "group"
	FunctionOriginObjectTypeRole          FunctionOriginObjectType = "role"
	FunctionOriginObjectTypeOrgMember     FunctionOriginObjectType = "org_member"
	FunctionOriginObjectTypeProjectLog    FunctionOriginObjectType = "project_log"
	FunctionOriginObjectTypeOrgProject    FunctionOriginObjectType = "org_project"
)

func (r FunctionOriginObjectType) IsKnown() bool {
	switch r {
	case FunctionOriginObjectTypeOrganization, FunctionOriginObjectTypeProject, FunctionOriginObjectTypeExperiment, FunctionOriginObjectTypeDataset, FunctionOriginObjectTypePrompt, FunctionOriginObjectTypePromptSession, FunctionOriginObjectTypeGroup, FunctionOriginObjectTypeRole, FunctionOriginObjectTypeOrgMember, FunctionOriginObjectTypeProjectLog, FunctionOriginObjectTypeOrgProject:
		return true
	}
	return false
}

// A group is a collection of users which can be assigned an ACL
//
// Groups can consist of individual users, as well as a set of groups they inherit
// from
type Group struct {
	// Unique identifier for the group
	ID string `json:"id,required" format:"uuid"`
	// Name of the group
	Name string `json:"name,required"`
	// Unique id for the organization that the group belongs under
	//
	// It is forbidden to change the org after creating a group
	OrgID string `json:"org_id,required" format:"uuid"`
	// Date of group creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Date of group deletion, or null if the group is still active
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Textual description of the group
	Description string `json:"description,nullable"`
	// Ids of the groups this group inherits from
	//
	// An inheriting group has all the users contained in its member groups, as well as
	// all of their inherited users
	MemberGroups []string `json:"member_groups,nullable" format:"uuid"`
	// Ids of users which belong to this group
	MemberUsers []string `json:"member_users,nullable" format:"uuid"`
	// Identifies the user who created the group
	UserID string    `json:"user_id,nullable" format:"uuid"`
	JSON   groupJSON `json:"-"`
}

// groupJSON contains the JSON metadata for the struct [Group]
type groupJSON struct {
	ID           apijson.Field
	Name         apijson.Field
	OrgID        apijson.Field
	Created      apijson.Field
	DeletedAt    apijson.Field
	Description  apijson.Field
	MemberGroups apijson.Field
	MemberUsers  apijson.Field
	UserID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Group) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r groupJSON) RawJSON() string {
	return r.raw
}

type InsertDatasetEventMergeParam struct {
	// The `_is_merge` field controls how the row is merged with any existing row with
	// the same id in the DB. By default (or when set to `false`), the existing row is
	// completely replaced by the new row. When set to `true`, the new row is
	// deep-merged into the existing row
	//
	// For example, say there is an existing row in the DB
	// `{"id": "foo", "input": {"a": 5, "b": 10}}`. If we merge a new row as
	// `{"_is_merge": true, "id": "foo", "input": {"b": 11, "c": 20}}`, the new row
	// will be `{"id": "foo", "input": {"a": 5, "b": 11, "c": 20}}`. If we replace the
	// new row as `{"id": "foo", "input": {"b": 11, "c": 20}}`, the new row will be
	// `{"id": "foo", "input": {"b": 11, "c": 20}}`
	IsMerge param.Field[bool] `json:"_is_merge,required"`
	// A unique identifier for the dataset event. If you don't provide one, BrainTrust
	// will generate one for you
	ID param.Field[string] `json:"id"`
	// The `_merge_paths` field allows controlling the depth of the merge. It can only
	// be specified alongside `_is_merge=true`. `_merge_paths` is a list of paths,
	// where each path is a list of field names. The deep merge will not descend below
	// any of the specified merge paths.
	//
	// For example, say there is an existing row in the DB
	// `{"id": "foo", "input": {"a": {"b": 10}, "c": {"d": 20}}, "output": {"a": 20}}`.
	// If we merge a new row as
	// `{"_is_merge": true, "_merge_paths": [["input", "a"], ["output"]], "input": {"a": {"q": 30}, "c": {"e": 30}, "bar": "baz"}, "output": {"d": 40}}`,
	// the new row will be
	// `{"id": "foo": "input": {"a": {"q": 30}, "c": {"d": 20, "e": 30}, "bar": "baz"}, "output": {"d": 40}}`.
	// In this case, due to the merge paths, we have replaced `input.a` and `output`,
	// but have still deep-merged `input` and `input.c`.
	MergePaths param.Field[[][]string] `json:"_merge_paths"`
	// Pass `_object_delete=true` to mark the dataset event deleted. Deleted events
	// will not show up in subsequent fetches for this dataset
	ObjectDelete param.Field[bool] `json:"_object_delete"`
	// The timestamp the dataset event was created
	Created param.Field[time.Time] `json:"created" format:"date-time"`
	// The output of your application, including post-processing (an arbitrary, JSON
	// serializable object)
	Expected param.Field[interface{}] `json:"expected"`
	// The argument that uniquely define an input case (an arbitrary, JSON serializable
	// object)
	Input param.Field[interface{}] `json:"input"`
	// A dictionary with additional data about the test example, model outputs, or just
	// about anything else that's relevant, that you can use to help find and analyze
	// examples later. For example, you could log the `prompt`, example's `id`, or
	// anything else that would be useful to slice/dice later. The values in `metadata`
	// can be any JSON-serializable type, but its keys must be strings
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// A list of tags to log
	Tags param.Field[[]string] `json:"tags"`
}

func (r InsertDatasetEventMergeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InsertDatasetEventMergeParam) ImplementsDatasetInsertParamsEventUnion() {}

type InsertDatasetEventReplaceParam struct {
	// A unique identifier for the dataset event. If you don't provide one, BrainTrust
	// will generate one for you
	ID param.Field[string] `json:"id"`
	// The `_is_merge` field controls how the row is merged with any existing row with
	// the same id in the DB. By default (or when set to `false`), the existing row is
	// completely replaced by the new row. When set to `true`, the new row is
	// deep-merged into the existing row
	//
	// For example, say there is an existing row in the DB
	// `{"id": "foo", "input": {"a": 5, "b": 10}}`. If we merge a new row as
	// `{"_is_merge": true, "id": "foo", "input": {"b": 11, "c": 20}}`, the new row
	// will be `{"id": "foo", "input": {"a": 5, "b": 11, "c": 20}}`. If we replace the
	// new row as `{"id": "foo", "input": {"b": 11, "c": 20}}`, the new row will be
	// `{"id": "foo", "input": {"b": 11, "c": 20}}`
	IsMerge param.Field[bool] `json:"_is_merge"`
	// Pass `_object_delete=true` to mark the dataset event deleted. Deleted events
	// will not show up in subsequent fetches for this dataset
	ObjectDelete param.Field[bool] `json:"_object_delete"`
	// Use the `_parent_id` field to create this row as a subspan of an existing row.
	// It cannot be specified alongside `_is_merge=true`. Tracking hierarchical
	// relationships are important for tracing (see the
	// [guide](https://www.braintrust.dev/docs/guides/tracing) for full details).
	//
	// For example, say we have logged a row
	// `{"id": "abc", "input": "foo", "output": "bar", "expected": "boo", "scores": {"correctness": 0.33}}`.
	// We can create a sub-span of the parent row by logging
	// `{"_parent_id": "abc", "id": "llm_call", "input": {"prompt": "What comes after foo?"}, "output": "bar", "metrics": {"tokens": 1}}`.
	// In the webapp, only the root span row `"abc"` will show up in the summary view.
	// You can view the full trace hierarchy (in this case, the `"llm_call"` row) by
	// clicking on the "abc" row.
	ParentID param.Field[string] `json:"_parent_id"`
	// The timestamp the dataset event was created
	Created param.Field[time.Time] `json:"created" format:"date-time"`
	// The output of your application, including post-processing (an arbitrary, JSON
	// serializable object)
	Expected param.Field[interface{}] `json:"expected"`
	// The argument that uniquely define an input case (an arbitrary, JSON serializable
	// object)
	Input param.Field[interface{}] `json:"input"`
	// A dictionary with additional data about the test example, model outputs, or just
	// about anything else that's relevant, that you can use to help find and analyze
	// examples later. For example, you could log the `prompt`, example's `id`, or
	// anything else that would be useful to slice/dice later. The values in `metadata`
	// can be any JSON-serializable type, but its keys must be strings
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// A list of tags to log
	Tags param.Field[[]string] `json:"tags"`
}

func (r InsertDatasetEventReplaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InsertDatasetEventReplaceParam) ImplementsDatasetInsertParamsEventUnion() {}

type InsertEventsResponse struct {
	// The ids of all rows that were inserted, aligning one-to-one with the rows
	// provided as input
	RowIDs []string                 `json:"row_ids,required"`
	JSON   insertEventsResponseJSON `json:"-"`
}

// insertEventsResponseJSON contains the JSON metadata for the struct
// [InsertEventsResponse]
type insertEventsResponseJSON struct {
	RowIDs      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsertEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insertEventsResponseJSON) RawJSON() string {
	return r.raw
}

type InsertExperimentEventMergeParam struct {
	// The `_is_merge` field controls how the row is merged with any existing row with
	// the same id in the DB. By default (or when set to `false`), the existing row is
	// completely replaced by the new row. When set to `true`, the new row is
	// deep-merged into the existing row
	//
	// For example, say there is an existing row in the DB
	// `{"id": "foo", "input": {"a": 5, "b": 10}}`. If we merge a new row as
	// `{"_is_merge": true, "id": "foo", "input": {"b": 11, "c": 20}}`, the new row
	// will be `{"id": "foo", "input": {"a": 5, "b": 11, "c": 20}}`. If we replace the
	// new row as `{"id": "foo", "input": {"b": 11, "c": 20}}`, the new row will be
	// `{"id": "foo", "input": {"b": 11, "c": 20}}`
	IsMerge param.Field[bool] `json:"_is_merge,required"`
	// A unique identifier for the experiment event. If you don't provide one,
	// BrainTrust will generate one for you
	ID param.Field[string] `json:"id"`
	// The `_merge_paths` field allows controlling the depth of the merge. It can only
	// be specified alongside `_is_merge=true`. `_merge_paths` is a list of paths,
	// where each path is a list of field names. The deep merge will not descend below
	// any of the specified merge paths.
	//
	// For example, say there is an existing row in the DB
	// `{"id": "foo", "input": {"a": {"b": 10}, "c": {"d": 20}}, "output": {"a": 20}}`.
	// If we merge a new row as
	// `{"_is_merge": true, "_merge_paths": [["input", "a"], ["output"]], "input": {"a": {"q": 30}, "c": {"e": 30}, "bar": "baz"}, "output": {"d": 40}}`,
	// the new row will be
	// `{"id": "foo": "input": {"a": {"q": 30}, "c": {"d": 20, "e": 30}, "bar": "baz"}, "output": {"d": 40}}`.
	// In this case, due to the merge paths, we have replaced `input.a` and `output`,
	// but have still deep-merged `input` and `input.c`.
	MergePaths param.Field[[][]string] `json:"_merge_paths"`
	// Pass `_object_delete=true` to mark the experiment event deleted. Deleted events
	// will not show up in subsequent fetches for this experiment
	ObjectDelete param.Field[bool] `json:"_object_delete"`
	// Context is additional information about the code that produced the experiment
	// event. It is essentially the textual counterpart to `metrics`. Use the
	// `caller_*` attributes to track the location in code which produced the
	// experiment event
	Context param.Field[InsertExperimentEventMergeContextParam] `json:"context"`
	// The timestamp the experiment event was created
	Created param.Field[time.Time] `json:"created" format:"date-time"`
	// If the experiment is associated to a dataset, this is the event-level dataset id
	// this experiment event is tied to
	DatasetRecordID param.Field[string] `json:"dataset_record_id"`
	// The error that occurred, if any.
	Error param.Field[interface{}] `json:"error"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not.
	// Braintrust currently does not compare `output` to `expected` for you, since
	// there are so many different ways to do that correctly. Instead, these values are
	// just used to help you navigate your experiments while digging into analyses.
	// However, we may later use these values to re-score outputs or fine-tune your
	// models
	Expected param.Field[interface{}] `json:"expected"`
	// The arguments that uniquely define a test case (an arbitrary, JSON serializable
	// object). Later on, Braintrust will use the `input` to know whether two test
	// cases are the same between experiments, so they should not contain
	// experiment-specific state. A simple rule of thumb is that if you run the same
	// experiment twice, the `input` should be identical
	Input param.Field[interface{}] `json:"input"`
	// A dictionary with additional data about the test example, model outputs, or just
	// about anything else that's relevant, that you can use to help find and analyze
	// examples later. For example, you could log the `prompt`, example's `id`, or
	// anything else that would be useful to slice/dice later. The values in `metadata`
	// can be any JSON-serializable type, but its keys must be strings
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// Metrics are numerical measurements tracking the execution of the code that
	// produced the experiment event. Use "start" and "end" to track the time span over
	// which the experiment event was produced
	Metrics param.Field[InsertExperimentEventMergeMetricsParam] `json:"metrics"`
	// The output of your application, including post-processing (an arbitrary, JSON
	// serializable object), that allows you to determine whether the result is correct
	// or not. For example, in an app that generates SQL queries, the `output` should
	// be the _result_ of the SQL query generated by the model, not the query itself,
	// because there may be multiple valid queries that answer a single question
	Output param.Field[interface{}] `json:"output"`
	// A dictionary of numeric values (between 0 and 1) to log. The scores should give
	// you a variety of signals that help you determine how accurate the outputs are
	// compared to what you expect and diagnose failures. For example, a summarization
	// app might have one score that tells you how accurate the summary is, and another
	// that measures the word similarity between the generated and grouth truth
	// summary. The word similarity score could help you determine whether the
	// summarization was covering similar concepts or not. You can use these scores to
	// help you sort, filter, and compare experiments
	Scores param.Field[map[string]float64] `json:"scores"`
	// Human-identifying attributes of the span, such as name, type, etc.
	SpanAttributes param.Field[InsertExperimentEventMergeSpanAttributesParam] `json:"span_attributes"`
	// A list of tags to log
	Tags param.Field[[]string] `json:"tags"`
}

func (r InsertExperimentEventMergeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InsertExperimentEventMergeParam) ImplementsExperimentInsertParamsEventUnion() {}

// Context is additional information about the code that produced the experiment
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the
// experiment event
type InsertExperimentEventMergeContextParam struct {
	// Name of the file in code where the experiment event was created
	CallerFilename param.Field[string] `json:"caller_filename"`
	// The function in code which created the experiment event
	CallerFunctionname param.Field[string] `json:"caller_functionname"`
	// Line of code where the experiment event was created
	CallerLineno param.Field[int64]     `json:"caller_lineno"`
	ExtraFields  map[string]interface{} `json:"-,extras"`
}

func (r InsertExperimentEventMergeContextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the experiment event. Use "start" and "end" to track the time span over
// which the experiment event was produced
type InsertExperimentEventMergeMetricsParam struct {
	// The number of tokens in the completion generated by the model (only set if this
	// is an LLM span)
	CompletionTokens param.Field[int64] `json:"completion_tokens"`
	// A unix timestamp recording when the section of code which produced the
	// experiment event finished
	End param.Field[float64] `json:"end"`
	// The number of tokens in the prompt used to generate the experiment event (only
	// set if this is an LLM span)
	PromptTokens param.Field[int64] `json:"prompt_tokens"`
	// A unix timestamp recording when the section of code which produced the
	// experiment event started
	Start param.Field[float64] `json:"start"`
	// The total number of tokens in the input and output of the experiment event.
	Tokens      param.Field[int64]     `json:"tokens"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r InsertExperimentEventMergeMetricsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Human-identifying attributes of the span, such as name, type, etc.
type InsertExperimentEventMergeSpanAttributesParam struct {
	// Name of the span, for display purposes only
	Name param.Field[string] `json:"name"`
	// Type of the span, for display purposes only
	Type        param.Field[InsertExperimentEventMergeSpanAttributesType] `json:"type"`
	ExtraFields map[string]interface{}                                    `json:"-,extras"`
}

func (r InsertExperimentEventMergeSpanAttributesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the span, for display purposes only
type InsertExperimentEventMergeSpanAttributesType string

const (
	InsertExperimentEventMergeSpanAttributesTypeLlm      InsertExperimentEventMergeSpanAttributesType = "llm"
	InsertExperimentEventMergeSpanAttributesTypeScore    InsertExperimentEventMergeSpanAttributesType = "score"
	InsertExperimentEventMergeSpanAttributesTypeFunction InsertExperimentEventMergeSpanAttributesType = "function"
	InsertExperimentEventMergeSpanAttributesTypeEval     InsertExperimentEventMergeSpanAttributesType = "eval"
	InsertExperimentEventMergeSpanAttributesTypeTask     InsertExperimentEventMergeSpanAttributesType = "task"
	InsertExperimentEventMergeSpanAttributesTypeTool     InsertExperimentEventMergeSpanAttributesType = "tool"
)

func (r InsertExperimentEventMergeSpanAttributesType) IsKnown() bool {
	switch r {
	case InsertExperimentEventMergeSpanAttributesTypeLlm, InsertExperimentEventMergeSpanAttributesTypeScore, InsertExperimentEventMergeSpanAttributesTypeFunction, InsertExperimentEventMergeSpanAttributesTypeEval, InsertExperimentEventMergeSpanAttributesTypeTask, InsertExperimentEventMergeSpanAttributesTypeTool:
		return true
	}
	return false
}

type InsertExperimentEventReplaceParam struct {
	// A unique identifier for the experiment event. If you don't provide one,
	// BrainTrust will generate one for you
	ID param.Field[string] `json:"id"`
	// The `_is_merge` field controls how the row is merged with any existing row with
	// the same id in the DB. By default (or when set to `false`), the existing row is
	// completely replaced by the new row. When set to `true`, the new row is
	// deep-merged into the existing row
	//
	// For example, say there is an existing row in the DB
	// `{"id": "foo", "input": {"a": 5, "b": 10}}`. If we merge a new row as
	// `{"_is_merge": true, "id": "foo", "input": {"b": 11, "c": 20}}`, the new row
	// will be `{"id": "foo", "input": {"a": 5, "b": 11, "c": 20}}`. If we replace the
	// new row as `{"id": "foo", "input": {"b": 11, "c": 20}}`, the new row will be
	// `{"id": "foo", "input": {"b": 11, "c": 20}}`
	IsMerge param.Field[bool] `json:"_is_merge"`
	// Pass `_object_delete=true` to mark the experiment event deleted. Deleted events
	// will not show up in subsequent fetches for this experiment
	ObjectDelete param.Field[bool] `json:"_object_delete"`
	// Use the `_parent_id` field to create this row as a subspan of an existing row.
	// It cannot be specified alongside `_is_merge=true`. Tracking hierarchical
	// relationships are important for tracing (see the
	// [guide](https://www.braintrust.dev/docs/guides/tracing) for full details).
	//
	// For example, say we have logged a row
	// `{"id": "abc", "input": "foo", "output": "bar", "expected": "boo", "scores": {"correctness": 0.33}}`.
	// We can create a sub-span of the parent row by logging
	// `{"_parent_id": "abc", "id": "llm_call", "input": {"prompt": "What comes after foo?"}, "output": "bar", "metrics": {"tokens": 1}}`.
	// In the webapp, only the root span row `"abc"` will show up in the summary view.
	// You can view the full trace hierarchy (in this case, the `"llm_call"` row) by
	// clicking on the "abc" row.
	ParentID param.Field[string] `json:"_parent_id"`
	// Context is additional information about the code that produced the experiment
	// event. It is essentially the textual counterpart to `metrics`. Use the
	// `caller_*` attributes to track the location in code which produced the
	// experiment event
	Context param.Field[InsertExperimentEventReplaceContextParam] `json:"context"`
	// The timestamp the experiment event was created
	Created param.Field[time.Time] `json:"created" format:"date-time"`
	// If the experiment is associated to a dataset, this is the event-level dataset id
	// this experiment event is tied to
	DatasetRecordID param.Field[string] `json:"dataset_record_id"`
	// The error that occurred, if any.
	Error param.Field[interface{}] `json:"error"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not.
	// Braintrust currently does not compare `output` to `expected` for you, since
	// there are so many different ways to do that correctly. Instead, these values are
	// just used to help you navigate your experiments while digging into analyses.
	// However, we may later use these values to re-score outputs or fine-tune your
	// models
	Expected param.Field[interface{}] `json:"expected"`
	// The arguments that uniquely define a test case (an arbitrary, JSON serializable
	// object). Later on, Braintrust will use the `input` to know whether two test
	// cases are the same between experiments, so they should not contain
	// experiment-specific state. A simple rule of thumb is that if you run the same
	// experiment twice, the `input` should be identical
	Input param.Field[interface{}] `json:"input"`
	// A dictionary with additional data about the test example, model outputs, or just
	// about anything else that's relevant, that you can use to help find and analyze
	// examples later. For example, you could log the `prompt`, example's `id`, or
	// anything else that would be useful to slice/dice later. The values in `metadata`
	// can be any JSON-serializable type, but its keys must be strings
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// Metrics are numerical measurements tracking the execution of the code that
	// produced the experiment event. Use "start" and "end" to track the time span over
	// which the experiment event was produced
	Metrics param.Field[InsertExperimentEventReplaceMetricsParam] `json:"metrics"`
	// The output of your application, including post-processing (an arbitrary, JSON
	// serializable object), that allows you to determine whether the result is correct
	// or not. For example, in an app that generates SQL queries, the `output` should
	// be the _result_ of the SQL query generated by the model, not the query itself,
	// because there may be multiple valid queries that answer a single question
	Output param.Field[interface{}] `json:"output"`
	// A dictionary of numeric values (between 0 and 1) to log. The scores should give
	// you a variety of signals that help you determine how accurate the outputs are
	// compared to what you expect and diagnose failures. For example, a summarization
	// app might have one score that tells you how accurate the summary is, and another
	// that measures the word similarity between the generated and grouth truth
	// summary. The word similarity score could help you determine whether the
	// summarization was covering similar concepts or not. You can use these scores to
	// help you sort, filter, and compare experiments
	Scores param.Field[map[string]float64] `json:"scores"`
	// Human-identifying attributes of the span, such as name, type, etc.
	SpanAttributes param.Field[InsertExperimentEventReplaceSpanAttributesParam] `json:"span_attributes"`
	// A list of tags to log
	Tags param.Field[[]string] `json:"tags"`
}

func (r InsertExperimentEventReplaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InsertExperimentEventReplaceParam) ImplementsExperimentInsertParamsEventUnion() {}

// Context is additional information about the code that produced the experiment
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the
// experiment event
type InsertExperimentEventReplaceContextParam struct {
	// Name of the file in code where the experiment event was created
	CallerFilename param.Field[string] `json:"caller_filename"`
	// The function in code which created the experiment event
	CallerFunctionname param.Field[string] `json:"caller_functionname"`
	// Line of code where the experiment event was created
	CallerLineno param.Field[int64]     `json:"caller_lineno"`
	ExtraFields  map[string]interface{} `json:"-,extras"`
}

func (r InsertExperimentEventReplaceContextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the experiment event. Use "start" and "end" to track the time span over
// which the experiment event was produced
type InsertExperimentEventReplaceMetricsParam struct {
	// The number of tokens in the completion generated by the model (only set if this
	// is an LLM span)
	CompletionTokens param.Field[int64] `json:"completion_tokens"`
	// A unix timestamp recording when the section of code which produced the
	// experiment event finished
	End param.Field[float64] `json:"end"`
	// The number of tokens in the prompt used to generate the experiment event (only
	// set if this is an LLM span)
	PromptTokens param.Field[int64] `json:"prompt_tokens"`
	// A unix timestamp recording when the section of code which produced the
	// experiment event started
	Start param.Field[float64] `json:"start"`
	// The total number of tokens in the input and output of the experiment event.
	Tokens      param.Field[int64]     `json:"tokens"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r InsertExperimentEventReplaceMetricsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Human-identifying attributes of the span, such as name, type, etc.
type InsertExperimentEventReplaceSpanAttributesParam struct {
	// Name of the span, for display purposes only
	Name param.Field[string] `json:"name"`
	// Type of the span, for display purposes only
	Type        param.Field[InsertExperimentEventReplaceSpanAttributesType] `json:"type"`
	ExtraFields map[string]interface{}                                      `json:"-,extras"`
}

func (r InsertExperimentEventReplaceSpanAttributesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the span, for display purposes only
type InsertExperimentEventReplaceSpanAttributesType string

const (
	InsertExperimentEventReplaceSpanAttributesTypeLlm      InsertExperimentEventReplaceSpanAttributesType = "llm"
	InsertExperimentEventReplaceSpanAttributesTypeScore    InsertExperimentEventReplaceSpanAttributesType = "score"
	InsertExperimentEventReplaceSpanAttributesTypeFunction InsertExperimentEventReplaceSpanAttributesType = "function"
	InsertExperimentEventReplaceSpanAttributesTypeEval     InsertExperimentEventReplaceSpanAttributesType = "eval"
	InsertExperimentEventReplaceSpanAttributesTypeTask     InsertExperimentEventReplaceSpanAttributesType = "task"
	InsertExperimentEventReplaceSpanAttributesTypeTool     InsertExperimentEventReplaceSpanAttributesType = "tool"
)

func (r InsertExperimentEventReplaceSpanAttributesType) IsKnown() bool {
	switch r {
	case InsertExperimentEventReplaceSpanAttributesTypeLlm, InsertExperimentEventReplaceSpanAttributesTypeScore, InsertExperimentEventReplaceSpanAttributesTypeFunction, InsertExperimentEventReplaceSpanAttributesTypeEval, InsertExperimentEventReplaceSpanAttributesTypeTask, InsertExperimentEventReplaceSpanAttributesTypeTool:
		return true
	}
	return false
}

type InsertProjectLogsEventMergeParam struct {
	// The `_is_merge` field controls how the row is merged with any existing row with
	// the same id in the DB. By default (or when set to `false`), the existing row is
	// completely replaced by the new row. When set to `true`, the new row is
	// deep-merged into the existing row
	//
	// For example, say there is an existing row in the DB
	// `{"id": "foo", "input": {"a": 5, "b": 10}}`. If we merge a new row as
	// `{"_is_merge": true, "id": "foo", "input": {"b": 11, "c": 20}}`, the new row
	// will be `{"id": "foo", "input": {"a": 5, "b": 11, "c": 20}}`. If we replace the
	// new row as `{"id": "foo", "input": {"b": 11, "c": 20}}`, the new row will be
	// `{"id": "foo", "input": {"b": 11, "c": 20}}`
	IsMerge param.Field[bool] `json:"_is_merge,required"`
	// A unique identifier for the project logs event. If you don't provide one,
	// BrainTrust will generate one for you
	ID param.Field[string] `json:"id"`
	// The `_merge_paths` field allows controlling the depth of the merge. It can only
	// be specified alongside `_is_merge=true`. `_merge_paths` is a list of paths,
	// where each path is a list of field names. The deep merge will not descend below
	// any of the specified merge paths.
	//
	// For example, say there is an existing row in the DB
	// `{"id": "foo", "input": {"a": {"b": 10}, "c": {"d": 20}}, "output": {"a": 20}}`.
	// If we merge a new row as
	// `{"_is_merge": true, "_merge_paths": [["input", "a"], ["output"]], "input": {"a": {"q": 30}, "c": {"e": 30}, "bar": "baz"}, "output": {"d": 40}}`,
	// the new row will be
	// `{"id": "foo": "input": {"a": {"q": 30}, "c": {"d": 20, "e": 30}, "bar": "baz"}, "output": {"d": 40}}`.
	// In this case, due to the merge paths, we have replaced `input.a` and `output`,
	// but have still deep-merged `input` and `input.c`.
	MergePaths param.Field[[][]string] `json:"_merge_paths"`
	// Pass `_object_delete=true` to mark the project logs event deleted. Deleted
	// events will not show up in subsequent fetches for this project logs
	ObjectDelete param.Field[bool] `json:"_object_delete"`
	// Context is additional information about the code that produced the project logs
	// event. It is essentially the textual counterpart to `metrics`. Use the
	// `caller_*` attributes to track the location in code which produced the project
	// logs event
	Context param.Field[InsertProjectLogsEventMergeContextParam] `json:"context"`
	// The timestamp the project logs event was created
	Created param.Field[time.Time] `json:"created" format:"date-time"`
	// The error that occurred, if any.
	Error param.Field[interface{}] `json:"error"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not.
	// Braintrust currently does not compare `output` to `expected` for you, since
	// there are so many different ways to do that correctly. Instead, these values are
	// just used to help you navigate while digging into analyses. However, we may
	// later use these values to re-score outputs or fine-tune your models.
	Expected param.Field[interface{}] `json:"expected"`
	// The arguments that uniquely define a user input (an arbitrary, JSON serializable
	// object).
	Input param.Field[interface{}] `json:"input"`
	// A dictionary with additional data about the test example, model outputs, or just
	// about anything else that's relevant, that you can use to help find and analyze
	// examples later. For example, you could log the `prompt`, example's `id`, or
	// anything else that would be useful to slice/dice later. The values in `metadata`
	// can be any JSON-serializable type, but its keys must be strings
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// Metrics are numerical measurements tracking the execution of the code that
	// produced the project logs event. Use "start" and "end" to track the time span
	// over which the project logs event was produced
	Metrics param.Field[InsertProjectLogsEventMergeMetricsParam] `json:"metrics"`
	// The output of your application, including post-processing (an arbitrary, JSON
	// serializable object), that allows you to determine whether the result is correct
	// or not. For example, in an app that generates SQL queries, the `output` should
	// be the _result_ of the SQL query generated by the model, not the query itself,
	// because there may be multiple valid queries that answer a single question.
	Output param.Field[interface{}] `json:"output"`
	// A dictionary of numeric values (between 0 and 1) to log. The scores should give
	// you a variety of signals that help you determine how accurate the outputs are
	// compared to what you expect and diagnose failures. For example, a summarization
	// app might have one score that tells you how accurate the summary is, and another
	// that measures the word similarity between the generated and grouth truth
	// summary. The word similarity score could help you determine whether the
	// summarization was covering similar concepts or not. You can use these scores to
	// help you sort, filter, and compare logs.
	Scores param.Field[map[string]float64] `json:"scores"`
	// Human-identifying attributes of the span, such as name, type, etc.
	SpanAttributes param.Field[InsertProjectLogsEventMergeSpanAttributesParam] `json:"span_attributes"`
	// A list of tags to log
	Tags param.Field[[]string] `json:"tags"`
}

func (r InsertProjectLogsEventMergeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InsertProjectLogsEventMergeParam) ImplementsProjectLogInsertParamsEventUnion() {}

// Context is additional information about the code that produced the project logs
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the project
// logs event
type InsertProjectLogsEventMergeContextParam struct {
	// Name of the file in code where the project logs event was created
	CallerFilename param.Field[string] `json:"caller_filename"`
	// The function in code which created the project logs event
	CallerFunctionname param.Field[string] `json:"caller_functionname"`
	// Line of code where the project logs event was created
	CallerLineno param.Field[int64]     `json:"caller_lineno"`
	ExtraFields  map[string]interface{} `json:"-,extras"`
}

func (r InsertProjectLogsEventMergeContextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the project logs event. Use "start" and "end" to track the time span
// over which the project logs event was produced
type InsertProjectLogsEventMergeMetricsParam struct {
	// The number of tokens in the completion generated by the model (only set if this
	// is an LLM span)
	CompletionTokens param.Field[int64] `json:"completion_tokens"`
	// A unix timestamp recording when the section of code which produced the project
	// logs event finished
	End param.Field[float64] `json:"end"`
	// The number of tokens in the prompt used to generate the project logs event (only
	// set if this is an LLM span)
	PromptTokens param.Field[int64] `json:"prompt_tokens"`
	// A unix timestamp recording when the section of code which produced the project
	// logs event started
	Start param.Field[float64] `json:"start"`
	// The total number of tokens in the input and output of the project logs event.
	Tokens      param.Field[int64]     `json:"tokens"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r InsertProjectLogsEventMergeMetricsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Human-identifying attributes of the span, such as name, type, etc.
type InsertProjectLogsEventMergeSpanAttributesParam struct {
	// Name of the span, for display purposes only
	Name param.Field[string] `json:"name"`
	// Type of the span, for display purposes only
	Type        param.Field[InsertProjectLogsEventMergeSpanAttributesType] `json:"type"`
	ExtraFields map[string]interface{}                                     `json:"-,extras"`
}

func (r InsertProjectLogsEventMergeSpanAttributesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the span, for display purposes only
type InsertProjectLogsEventMergeSpanAttributesType string

const (
	InsertProjectLogsEventMergeSpanAttributesTypeLlm      InsertProjectLogsEventMergeSpanAttributesType = "llm"
	InsertProjectLogsEventMergeSpanAttributesTypeScore    InsertProjectLogsEventMergeSpanAttributesType = "score"
	InsertProjectLogsEventMergeSpanAttributesTypeFunction InsertProjectLogsEventMergeSpanAttributesType = "function"
	InsertProjectLogsEventMergeSpanAttributesTypeEval     InsertProjectLogsEventMergeSpanAttributesType = "eval"
	InsertProjectLogsEventMergeSpanAttributesTypeTask     InsertProjectLogsEventMergeSpanAttributesType = "task"
	InsertProjectLogsEventMergeSpanAttributesTypeTool     InsertProjectLogsEventMergeSpanAttributesType = "tool"
)

func (r InsertProjectLogsEventMergeSpanAttributesType) IsKnown() bool {
	switch r {
	case InsertProjectLogsEventMergeSpanAttributesTypeLlm, InsertProjectLogsEventMergeSpanAttributesTypeScore, InsertProjectLogsEventMergeSpanAttributesTypeFunction, InsertProjectLogsEventMergeSpanAttributesTypeEval, InsertProjectLogsEventMergeSpanAttributesTypeTask, InsertProjectLogsEventMergeSpanAttributesTypeTool:
		return true
	}
	return false
}

type InsertProjectLogsEventReplaceParam struct {
	// A unique identifier for the project logs event. If you don't provide one,
	// BrainTrust will generate one for you
	ID param.Field[string] `json:"id"`
	// The `_is_merge` field controls how the row is merged with any existing row with
	// the same id in the DB. By default (or when set to `false`), the existing row is
	// completely replaced by the new row. When set to `true`, the new row is
	// deep-merged into the existing row
	//
	// For example, say there is an existing row in the DB
	// `{"id": "foo", "input": {"a": 5, "b": 10}}`. If we merge a new row as
	// `{"_is_merge": true, "id": "foo", "input": {"b": 11, "c": 20}}`, the new row
	// will be `{"id": "foo", "input": {"a": 5, "b": 11, "c": 20}}`. If we replace the
	// new row as `{"id": "foo", "input": {"b": 11, "c": 20}}`, the new row will be
	// `{"id": "foo", "input": {"b": 11, "c": 20}}`
	IsMerge param.Field[bool] `json:"_is_merge"`
	// Pass `_object_delete=true` to mark the project logs event deleted. Deleted
	// events will not show up in subsequent fetches for this project logs
	ObjectDelete param.Field[bool] `json:"_object_delete"`
	// Use the `_parent_id` field to create this row as a subspan of an existing row.
	// It cannot be specified alongside `_is_merge=true`. Tracking hierarchical
	// relationships are important for tracing (see the
	// [guide](https://www.braintrust.dev/docs/guides/tracing) for full details).
	//
	// For example, say we have logged a row
	// `{"id": "abc", "input": "foo", "output": "bar", "expected": "boo", "scores": {"correctness": 0.33}}`.
	// We can create a sub-span of the parent row by logging
	// `{"_parent_id": "abc", "id": "llm_call", "input": {"prompt": "What comes after foo?"}, "output": "bar", "metrics": {"tokens": 1}}`.
	// In the webapp, only the root span row `"abc"` will show up in the summary view.
	// You can view the full trace hierarchy (in this case, the `"llm_call"` row) by
	// clicking on the "abc" row.
	ParentID param.Field[string] `json:"_parent_id"`
	// Context is additional information about the code that produced the project logs
	// event. It is essentially the textual counterpart to `metrics`. Use the
	// `caller_*` attributes to track the location in code which produced the project
	// logs event
	Context param.Field[InsertProjectLogsEventReplaceContextParam] `json:"context"`
	// The timestamp the project logs event was created
	Created param.Field[time.Time] `json:"created" format:"date-time"`
	// The error that occurred, if any.
	Error param.Field[interface{}] `json:"error"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not.
	// Braintrust currently does not compare `output` to `expected` for you, since
	// there are so many different ways to do that correctly. Instead, these values are
	// just used to help you navigate while digging into analyses. However, we may
	// later use these values to re-score outputs or fine-tune your models.
	Expected param.Field[interface{}] `json:"expected"`
	// The arguments that uniquely define a user input (an arbitrary, JSON serializable
	// object).
	Input param.Field[interface{}] `json:"input"`
	// A dictionary with additional data about the test example, model outputs, or just
	// about anything else that's relevant, that you can use to help find and analyze
	// examples later. For example, you could log the `prompt`, example's `id`, or
	// anything else that would be useful to slice/dice later. The values in `metadata`
	// can be any JSON-serializable type, but its keys must be strings
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// Metrics are numerical measurements tracking the execution of the code that
	// produced the project logs event. Use "start" and "end" to track the time span
	// over which the project logs event was produced
	Metrics param.Field[InsertProjectLogsEventReplaceMetricsParam] `json:"metrics"`
	// The output of your application, including post-processing (an arbitrary, JSON
	// serializable object), that allows you to determine whether the result is correct
	// or not. For example, in an app that generates SQL queries, the `output` should
	// be the _result_ of the SQL query generated by the model, not the query itself,
	// because there may be multiple valid queries that answer a single question.
	Output param.Field[interface{}] `json:"output"`
	// A dictionary of numeric values (between 0 and 1) to log. The scores should give
	// you a variety of signals that help you determine how accurate the outputs are
	// compared to what you expect and diagnose failures. For example, a summarization
	// app might have one score that tells you how accurate the summary is, and another
	// that measures the word similarity between the generated and grouth truth
	// summary. The word similarity score could help you determine whether the
	// summarization was covering similar concepts or not. You can use these scores to
	// help you sort, filter, and compare logs.
	Scores param.Field[map[string]float64] `json:"scores"`
	// Human-identifying attributes of the span, such as name, type, etc.
	SpanAttributes param.Field[InsertProjectLogsEventReplaceSpanAttributesParam] `json:"span_attributes"`
	// A list of tags to log
	Tags param.Field[[]string] `json:"tags"`
}

func (r InsertProjectLogsEventReplaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InsertProjectLogsEventReplaceParam) ImplementsProjectLogInsertParamsEventUnion() {}

// Context is additional information about the code that produced the project logs
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the project
// logs event
type InsertProjectLogsEventReplaceContextParam struct {
	// Name of the file in code where the project logs event was created
	CallerFilename param.Field[string] `json:"caller_filename"`
	// The function in code which created the project logs event
	CallerFunctionname param.Field[string] `json:"caller_functionname"`
	// Line of code where the project logs event was created
	CallerLineno param.Field[int64]     `json:"caller_lineno"`
	ExtraFields  map[string]interface{} `json:"-,extras"`
}

func (r InsertProjectLogsEventReplaceContextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the project logs event. Use "start" and "end" to track the time span
// over which the project logs event was produced
type InsertProjectLogsEventReplaceMetricsParam struct {
	// The number of tokens in the completion generated by the model (only set if this
	// is an LLM span)
	CompletionTokens param.Field[int64] `json:"completion_tokens"`
	// A unix timestamp recording when the section of code which produced the project
	// logs event finished
	End param.Field[float64] `json:"end"`
	// The number of tokens in the prompt used to generate the project logs event (only
	// set if this is an LLM span)
	PromptTokens param.Field[int64] `json:"prompt_tokens"`
	// A unix timestamp recording when the section of code which produced the project
	// logs event started
	Start param.Field[float64] `json:"start"`
	// The total number of tokens in the input and output of the project logs event.
	Tokens      param.Field[int64]     `json:"tokens"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r InsertProjectLogsEventReplaceMetricsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Human-identifying attributes of the span, such as name, type, etc.
type InsertProjectLogsEventReplaceSpanAttributesParam struct {
	// Name of the span, for display purposes only
	Name param.Field[string] `json:"name"`
	// Type of the span, for display purposes only
	Type        param.Field[InsertProjectLogsEventReplaceSpanAttributesType] `json:"type"`
	ExtraFields map[string]interface{}                                       `json:"-,extras"`
}

func (r InsertProjectLogsEventReplaceSpanAttributesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the span, for display purposes only
type InsertProjectLogsEventReplaceSpanAttributesType string

const (
	InsertProjectLogsEventReplaceSpanAttributesTypeLlm      InsertProjectLogsEventReplaceSpanAttributesType = "llm"
	InsertProjectLogsEventReplaceSpanAttributesTypeScore    InsertProjectLogsEventReplaceSpanAttributesType = "score"
	InsertProjectLogsEventReplaceSpanAttributesTypeFunction InsertProjectLogsEventReplaceSpanAttributesType = "function"
	InsertProjectLogsEventReplaceSpanAttributesTypeEval     InsertProjectLogsEventReplaceSpanAttributesType = "eval"
	InsertProjectLogsEventReplaceSpanAttributesTypeTask     InsertProjectLogsEventReplaceSpanAttributesType = "task"
	InsertProjectLogsEventReplaceSpanAttributesTypeTool     InsertProjectLogsEventReplaceSpanAttributesType = "tool"
)

func (r InsertProjectLogsEventReplaceSpanAttributesType) IsKnown() bool {
	switch r {
	case InsertProjectLogsEventReplaceSpanAttributesTypeLlm, InsertProjectLogsEventReplaceSpanAttributesTypeScore, InsertProjectLogsEventReplaceSpanAttributesTypeFunction, InsertProjectLogsEventReplaceSpanAttributesTypeEval, InsertProjectLogsEventReplaceSpanAttributesTypeTask, InsertProjectLogsEventReplaceSpanAttributesTypeTool:
		return true
	}
	return false
}

// Summary of a metric's performance
type MetricSummary struct {
	// Number of improvements in the metric
	Improvements int64 `json:"improvements,required"`
	// Average metric across all examples
	Metric float64 `json:"metric,required"`
	// Name of the metric
	Name string `json:"name,required"`
	// Number of regressions in the metric
	Regressions int64 `json:"regressions,required"`
	// Unit label for the metric
	Unit string `json:"unit,required"`
	// Difference in metric between the current and comparison experiment
	Diff float64           `json:"diff"`
	JSON metricSummaryJSON `json:"-"`
}

// metricSummaryJSON contains the JSON metadata for the struct [MetricSummary]
type metricSummaryJSON struct {
	Improvements apijson.Field
	Metric       apijson.Field
	Name         apijson.Field
	Regressions  apijson.Field
	Unit         apijson.Field
	Diff         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MetricSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricSummaryJSON) RawJSON() string {
	return r.raw
}

type Organization struct {
	// Unique identifier for the organization
	ID string `json:"id,required" format:"uuid"`
	// Name of the organization
	Name   string `json:"name,required"`
	APIURL string `json:"api_url,nullable"`
	// Date of organization creation
	Created        time.Time        `json:"created,nullable" format:"date-time"`
	IsUniversalAPI bool             `json:"is_universal_api,nullable"`
	ProxyURL       string           `json:"proxy_url,nullable"`
	RealtimeURL    string           `json:"realtime_url,nullable"`
	JSON           organizationJSON `json:"-"`
}

// organizationJSON contains the JSON metadata for the struct [Organization]
type organizationJSON struct {
	ID             apijson.Field
	Name           apijson.Field
	APIURL         apijson.Field
	Created        apijson.Field
	IsUniversalAPI apijson.Field
	ProxyURL       apijson.Field
	RealtimeURL    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Organization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationJSON) RawJSON() string {
	return r.raw
}

// A path-lookup filter describes an equality comparison against a specific
// sub-field in the event row. For instance, if you wish to filter on the value of
// `c` in `{"input": {"a": {"b": {"c": "hello"}}}}`, pass
// `path=["input", "a", "b", "c"]` and `value="hello"`
type PathLookupFilterParam struct {
	// List of fields describing the path to the value to be checked against. For
	// instance, if you wish to filter on the value of `c` in
	// `{"input": {"a": {"b": {"c": "hello"}}}}`, pass `path=["input", "a", "b", "c"]`
	Path param.Field[[]string] `json:"path,required"`
	// Denotes the type of filter as a path-lookup filter
	Type param.Field[PathLookupFilterType] `json:"type,required"`
	// The value to compare equality-wise against the event value at the specified
	// `path`. The value must be a "primitive", that is, any JSON-serializable object
	// except for objects and arrays. For instance, if you wish to filter on the value
	// of "input.a.b.c" in the object `{"input": {"a": {"b": {"c": "hello"}}}}`, pass
	// `value="hello"`
	Value param.Field[interface{}] `json:"value"`
}

func (r PathLookupFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Denotes the type of filter as a path-lookup filter
type PathLookupFilterType string

const (
	PathLookupFilterTypePathLookup PathLookupFilterType = "path_lookup"
)

func (r PathLookupFilterType) IsKnown() bool {
	switch r {
	case PathLookupFilterTypePathLookup:
		return true
	}
	return false
}

type Project struct {
	// Unique identifier for the project
	ID string `json:"id,required" format:"uuid"`
	// Name of the project
	Name string `json:"name,required"`
	// Unique id for the organization that the project belongs under
	OrgID string `json:"org_id,required" format:"uuid"`
	// Date of project creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Date of project deletion, or null if the project is still active
	DeletedAt time.Time       `json:"deleted_at,nullable" format:"date-time"`
	Settings  ProjectSettings `json:"settings,nullable"`
	// Identifies the user who created the project
	UserID string      `json:"user_id,nullable" format:"uuid"`
	JSON   projectJSON `json:"-"`
}

// projectJSON contains the JSON metadata for the struct [Project]
type projectJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	OrgID       apijson.Field
	Created     apijson.Field
	DeletedAt   apijson.Field
	Settings    apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Project) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectJSON) RawJSON() string {
	return r.raw
}

type ProjectSettings struct {
	// The key used to join two experiments (defaults to `input`).
	ComparisonKey string              `json:"comparison_key,nullable"`
	JSON          projectSettingsJSON `json:"-"`
}

// projectSettingsJSON contains the JSON metadata for the struct [ProjectSettings]
type projectSettingsJSON struct {
	ComparisonKey apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectSettingsJSON) RawJSON() string {
	return r.raw
}

type ProjectLogsEvent struct {
	// A unique identifier for the project logs event. If you don't provide one,
	// BrainTrust will generate one for you
	ID string `json:"id,required"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the project logs (see the
	// `version` parameter)
	XactID string `json:"_xact_id,required"`
	// The timestamp the project logs event was created
	Created time.Time `json:"created,required" format:"date-time"`
	// A literal 'g' which identifies the log as a project log
	LogID ProjectLogsEventLogID `json:"log_id,required"`
	// Unique id for the organization that the project belongs under
	OrgID string `json:"org_id,required" format:"uuid"`
	// Unique identifier for the project
	ProjectID string `json:"project_id,required" format:"uuid"`
	// The `span_id` of the root of the trace this project logs event belongs to
	RootSpanID string `json:"root_span_id,required"`
	// A unique identifier used to link different project logs events together as part
	// of a full trace. See the
	// [tracing guide](https://www.braintrust.dev/docs/guides/tracing) for full details
	// on tracing
	SpanID string `json:"span_id,required"`
	// Context is additional information about the code that produced the project logs
	// event. It is essentially the textual counterpart to `metrics`. Use the
	// `caller_*` attributes to track the location in code which produced the project
	// logs event
	Context ProjectLogsEventContext `json:"context,nullable"`
	// The error that occurred, if any.
	Error interface{} `json:"error"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not.
	// Braintrust currently does not compare `output` to `expected` for you, since
	// there are so many different ways to do that correctly. Instead, these values are
	// just used to help you navigate while digging into analyses. However, we may
	// later use these values to re-score outputs or fine-tune your models.
	Expected interface{} `json:"expected"`
	// The arguments that uniquely define a user input (an arbitrary, JSON serializable
	// object).
	Input interface{} `json:"input"`
	// A dictionary with additional data about the test example, model outputs, or just
	// about anything else that's relevant, that you can use to help find and analyze
	// examples later. For example, you could log the `prompt`, example's `id`, or
	// anything else that would be useful to slice/dice later. The values in `metadata`
	// can be any JSON-serializable type, but its keys must be strings
	Metadata map[string]interface{} `json:"metadata,nullable"`
	// Metrics are numerical measurements tracking the execution of the code that
	// produced the project logs event. Use "start" and "end" to track the time span
	// over which the project logs event was produced
	Metrics ProjectLogsEventMetrics `json:"metrics,nullable"`
	// The output of your application, including post-processing (an arbitrary, JSON
	// serializable object), that allows you to determine whether the result is correct
	// or not. For example, in an app that generates SQL queries, the `output` should
	// be the _result_ of the SQL query generated by the model, not the query itself,
	// because there may be multiple valid queries that answer a single question.
	Output interface{} `json:"output"`
	// A dictionary of numeric values (between 0 and 1) to log. The scores should give
	// you a variety of signals that help you determine how accurate the outputs are
	// compared to what you expect and diagnose failures. For example, a summarization
	// app might have one score that tells you how accurate the summary is, and another
	// that measures the word similarity between the generated and grouth truth
	// summary. The word similarity score could help you determine whether the
	// summarization was covering similar concepts or not. You can use these scores to
	// help you sort, filter, and compare logs.
	Scores map[string]float64 `json:"scores,nullable"`
	// Human-identifying attributes of the span, such as name, type, etc.
	SpanAttributes ProjectLogsEventSpanAttributes `json:"span_attributes,nullable"`
	// An array of the parent `span_ids` of this project logs event. This should be
	// empty for the root span of a trace, and should most often contain just one
	// parent element for subspans
	SpanParents []string `json:"span_parents,nullable"`
	// A list of tags to log
	Tags []string             `json:"tags,nullable"`
	JSON projectLogsEventJSON `json:"-"`
}

// projectLogsEventJSON contains the JSON metadata for the struct
// [ProjectLogsEvent]
type projectLogsEventJSON struct {
	ID             apijson.Field
	XactID         apijson.Field
	Created        apijson.Field
	LogID          apijson.Field
	OrgID          apijson.Field
	ProjectID      apijson.Field
	RootSpanID     apijson.Field
	SpanID         apijson.Field
	Context        apijson.Field
	Error          apijson.Field
	Expected       apijson.Field
	Input          apijson.Field
	Metadata       apijson.Field
	Metrics        apijson.Field
	Output         apijson.Field
	Scores         apijson.Field
	SpanAttributes apijson.Field
	SpanParents    apijson.Field
	Tags           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectLogsEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectLogsEventJSON) RawJSON() string {
	return r.raw
}

// A literal 'g' which identifies the log as a project log
type ProjectLogsEventLogID string

const (
	ProjectLogsEventLogIDG ProjectLogsEventLogID = "g"
)

func (r ProjectLogsEventLogID) IsKnown() bool {
	switch r {
	case ProjectLogsEventLogIDG:
		return true
	}
	return false
}

// Context is additional information about the code that produced the project logs
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the project
// logs event
type ProjectLogsEventContext struct {
	// Name of the file in code where the project logs event was created
	CallerFilename string `json:"caller_filename,nullable"`
	// The function in code which created the project logs event
	CallerFunctionname string `json:"caller_functionname,nullable"`
	// Line of code where the project logs event was created
	CallerLineno int64                       `json:"caller_lineno,nullable"`
	ExtraFields  map[string]interface{}      `json:"-,extras"`
	JSON         projectLogsEventContextJSON `json:"-"`
}

// projectLogsEventContextJSON contains the JSON metadata for the struct
// [ProjectLogsEventContext]
type projectLogsEventContextJSON struct {
	CallerFilename     apijson.Field
	CallerFunctionname apijson.Field
	CallerLineno       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProjectLogsEventContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectLogsEventContextJSON) RawJSON() string {
	return r.raw
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the project logs event. Use "start" and "end" to track the time span
// over which the project logs event was produced
type ProjectLogsEventMetrics struct {
	// The number of tokens in the completion generated by the model (only set if this
	// is an LLM span)
	CompletionTokens int64 `json:"completion_tokens,nullable"`
	// A unix timestamp recording when the section of code which produced the project
	// logs event finished
	End float64 `json:"end,nullable"`
	// The number of tokens in the prompt used to generate the project logs event (only
	// set if this is an LLM span)
	PromptTokens int64 `json:"prompt_tokens,nullable"`
	// A unix timestamp recording when the section of code which produced the project
	// logs event started
	Start float64 `json:"start,nullable"`
	// The total number of tokens in the input and output of the project logs event.
	Tokens      int64                       `json:"tokens,nullable"`
	ExtraFields map[string]interface{}      `json:"-,extras"`
	JSON        projectLogsEventMetricsJSON `json:"-"`
}

// projectLogsEventMetricsJSON contains the JSON metadata for the struct
// [ProjectLogsEventMetrics]
type projectLogsEventMetricsJSON struct {
	CompletionTokens apijson.Field
	End              apijson.Field
	PromptTokens     apijson.Field
	Start            apijson.Field
	Tokens           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectLogsEventMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectLogsEventMetricsJSON) RawJSON() string {
	return r.raw
}

// Human-identifying attributes of the span, such as name, type, etc.
type ProjectLogsEventSpanAttributes struct {
	// Name of the span, for display purposes only
	Name string `json:"name,nullable"`
	// Type of the span, for display purposes only
	Type        ProjectLogsEventSpanAttributesType `json:"type,nullable"`
	ExtraFields map[string]interface{}             `json:"-,extras"`
	JSON        projectLogsEventSpanAttributesJSON `json:"-"`
}

// projectLogsEventSpanAttributesJSON contains the JSON metadata for the struct
// [ProjectLogsEventSpanAttributes]
type projectLogsEventSpanAttributesJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectLogsEventSpanAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectLogsEventSpanAttributesJSON) RawJSON() string {
	return r.raw
}

// Type of the span, for display purposes only
type ProjectLogsEventSpanAttributesType string

const (
	ProjectLogsEventSpanAttributesTypeLlm      ProjectLogsEventSpanAttributesType = "llm"
	ProjectLogsEventSpanAttributesTypeScore    ProjectLogsEventSpanAttributesType = "score"
	ProjectLogsEventSpanAttributesTypeFunction ProjectLogsEventSpanAttributesType = "function"
	ProjectLogsEventSpanAttributesTypeEval     ProjectLogsEventSpanAttributesType = "eval"
	ProjectLogsEventSpanAttributesTypeTask     ProjectLogsEventSpanAttributesType = "task"
	ProjectLogsEventSpanAttributesTypeTool     ProjectLogsEventSpanAttributesType = "tool"
)

func (r ProjectLogsEventSpanAttributesType) IsKnown() bool {
	switch r {
	case ProjectLogsEventSpanAttributesTypeLlm, ProjectLogsEventSpanAttributesTypeScore, ProjectLogsEventSpanAttributesTypeFunction, ProjectLogsEventSpanAttributesTypeEval, ProjectLogsEventSpanAttributesTypeTask, ProjectLogsEventSpanAttributesTypeTool:
		return true
	}
	return false
}

// A project score is a user-configured score, which can be manually-labeled
// through the UI
type ProjectScore struct {
	// Unique identifier for the project score
	ID string `json:"id,required" format:"uuid"`
	// Name of the project score
	Name string `json:"name,required"`
	// Unique identifier for the project that the project score belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// The type of the configured score
	ScoreType ProjectScoreScoreType `json:"score_type,required,nullable"`
	UserID    string                `json:"user_id,required" format:"uuid"`
	// For categorical-type project scores, the list of all categories
	Categories ProjectScoreCategoriesUnion `json:"categories"`
	Config     ProjectScoreConfig          `json:"config,nullable"`
	// Date of project score creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Textual description of the project score
	Description string `json:"description,nullable"`
	// An optional LexoRank-based string that sets the sort position for the score in
	// the UI
	Position string           `json:"position,nullable"`
	JSON     projectScoreJSON `json:"-"`
}

// projectScoreJSON contains the JSON metadata for the struct [ProjectScore]
type projectScoreJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	ProjectID   apijson.Field
	ScoreType   apijson.Field
	UserID      apijson.Field
	Categories  apijson.Field
	Config      apijson.Field
	Created     apijson.Field
	Description apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectScore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectScoreJSON) RawJSON() string {
	return r.raw
}

// The type of the configured score
type ProjectScoreScoreType string

const (
	ProjectScoreScoreTypeSlider      ProjectScoreScoreType = "slider"
	ProjectScoreScoreTypeCategorical ProjectScoreScoreType = "categorical"
	ProjectScoreScoreTypeWeighted    ProjectScoreScoreType = "weighted"
	ProjectScoreScoreTypeMinimum     ProjectScoreScoreType = "minimum"
	ProjectScoreScoreTypeOnline      ProjectScoreScoreType = "online"
)

func (r ProjectScoreScoreType) IsKnown() bool {
	switch r {
	case ProjectScoreScoreTypeSlider, ProjectScoreScoreTypeCategorical, ProjectScoreScoreTypeWeighted, ProjectScoreScoreTypeMinimum, ProjectScoreScoreTypeOnline:
		return true
	}
	return false
}

// For categorical-type project scores, the list of all categories
//
// Union satisfied by [shared.ProjectScoreCategoriesCategorical],
// [shared.ProjectScoreCategoriesMinimum] or
// [shared.ProjectScoreCategoriesNullableVariant].
type ProjectScoreCategoriesUnion interface {
	implementsSharedProjectScoreCategoriesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectScoreCategoriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectScoreCategoriesCategorical{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectScoreCategoriesMinimum{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectScoreCategoriesNullableVariant{}),
		},
	)
}

type ProjectScoreCategoriesCategorical []ProjectScoreCategory

func (r ProjectScoreCategoriesCategorical) implementsSharedProjectScoreCategoriesUnion() {}

type ProjectScoreCategoriesMinimum []string

func (r ProjectScoreCategoriesMinimum) implementsSharedProjectScoreCategoriesUnion() {}

type ProjectScoreCategoriesNullableVariant struct {
	JSON projectScoreCategoriesNullableVariantJSON `json:"-"`
}

// projectScoreCategoriesNullableVariantJSON contains the JSON metadata for the
// struct [ProjectScoreCategoriesNullableVariant]
type projectScoreCategoriesNullableVariantJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectScoreCategoriesNullableVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectScoreCategoriesNullableVariantJSON) RawJSON() string {
	return r.raw
}

func (r ProjectScoreCategoriesNullableVariant) implementsSharedProjectScoreCategoriesUnion() {}

type ProjectScoreConfig struct {
	Destination ProjectScoreConfigDestination `json:"destination,nullable"`
	MultiSelect bool                          `json:"multi_select,nullable"`
	Online      ProjectScoreConfigOnline      `json:"online,nullable"`
	JSON        projectScoreConfigJSON        `json:"-"`
}

// projectScoreConfigJSON contains the JSON metadata for the struct
// [ProjectScoreConfig]
type projectScoreConfigJSON struct {
	Destination apijson.Field
	MultiSelect apijson.Field
	Online      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectScoreConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectScoreConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectScoreConfigDestination string

const (
	ProjectScoreConfigDestinationExpected ProjectScoreConfigDestination = "expected"
)

func (r ProjectScoreConfigDestination) IsKnown() bool {
	switch r {
	case ProjectScoreConfigDestinationExpected:
		return true
	}
	return false
}

type ProjectScoreConfigOnline struct {
	// The sampling rate for online scoring
	SamplingRate float64 `json:"sampling_rate,required"`
	// The list of scorers to use for online scoring
	Scorers []ProjectScoreConfigOnlineScorer `json:"scorers,required"`
	// Whether to trigger online scoring on the root span of each trace
	ApplyToRootSpan bool `json:"apply_to_root_span,nullable"`
	// Trigger online scoring on any spans with a name in this list
	ApplyToSpanNames []string                     `json:"apply_to_span_names,nullable"`
	JSON             projectScoreConfigOnlineJSON `json:"-"`
}

// projectScoreConfigOnlineJSON contains the JSON metadata for the struct
// [ProjectScoreConfigOnline]
type projectScoreConfigOnlineJSON struct {
	SamplingRate     apijson.Field
	Scorers          apijson.Field
	ApplyToRootSpan  apijson.Field
	ApplyToSpanNames apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectScoreConfigOnline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectScoreConfigOnlineJSON) RawJSON() string {
	return r.raw
}

type ProjectScoreConfigOnlineScorer struct {
	Type  ProjectScoreConfigOnlineScorersType `json:"type,required"`
	ID    string                              `json:"id"`
	Name  string                              `json:"name"`
	JSON  projectScoreConfigOnlineScorerJSON  `json:"-"`
	union ProjectScoreConfigOnlineScorersUnion
}

// projectScoreConfigOnlineScorerJSON contains the JSON metadata for the struct
// [ProjectScoreConfigOnlineScorer]
type projectScoreConfigOnlineScorerJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectScoreConfigOnlineScorerJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectScoreConfigOnlineScorer) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectScoreConfigOnlineScorer{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectScoreConfigOnlineScorersUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.ProjectScoreConfigOnlineScorersFunction],
// [shared.ProjectScoreConfigOnlineScorersGlobal].
func (r ProjectScoreConfigOnlineScorer) AsUnion() ProjectScoreConfigOnlineScorersUnion {
	return r.union
}

// Union satisfied by [shared.ProjectScoreConfigOnlineScorersFunction] or
// [shared.ProjectScoreConfigOnlineScorersGlobal].
type ProjectScoreConfigOnlineScorersUnion interface {
	implementsSharedProjectScoreConfigOnlineScorer()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectScoreConfigOnlineScorersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectScoreConfigOnlineScorersFunction{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectScoreConfigOnlineScorersGlobal{}),
		},
	)
}

type ProjectScoreConfigOnlineScorersFunction struct {
	ID   string                                      `json:"id,required"`
	Type ProjectScoreConfigOnlineScorersFunctionType `json:"type,required"`
	JSON projectScoreConfigOnlineScorersFunctionJSON `json:"-"`
}

// projectScoreConfigOnlineScorersFunctionJSON contains the JSON metadata for the
// struct [ProjectScoreConfigOnlineScorersFunction]
type projectScoreConfigOnlineScorersFunctionJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectScoreConfigOnlineScorersFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectScoreConfigOnlineScorersFunctionJSON) RawJSON() string {
	return r.raw
}

func (r ProjectScoreConfigOnlineScorersFunction) implementsSharedProjectScoreConfigOnlineScorer() {}

type ProjectScoreConfigOnlineScorersFunctionType string

const (
	ProjectScoreConfigOnlineScorersFunctionTypeFunction ProjectScoreConfigOnlineScorersFunctionType = "function"
)

func (r ProjectScoreConfigOnlineScorersFunctionType) IsKnown() bool {
	switch r {
	case ProjectScoreConfigOnlineScorersFunctionTypeFunction:
		return true
	}
	return false
}

type ProjectScoreConfigOnlineScorersGlobal struct {
	Name string                                    `json:"name,required"`
	Type ProjectScoreConfigOnlineScorersGlobalType `json:"type,required"`
	JSON projectScoreConfigOnlineScorersGlobalJSON `json:"-"`
}

// projectScoreConfigOnlineScorersGlobalJSON contains the JSON metadata for the
// struct [ProjectScoreConfigOnlineScorersGlobal]
type projectScoreConfigOnlineScorersGlobalJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectScoreConfigOnlineScorersGlobal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectScoreConfigOnlineScorersGlobalJSON) RawJSON() string {
	return r.raw
}

func (r ProjectScoreConfigOnlineScorersGlobal) implementsSharedProjectScoreConfigOnlineScorer() {}

type ProjectScoreConfigOnlineScorersGlobalType string

const (
	ProjectScoreConfigOnlineScorersGlobalTypeGlobal ProjectScoreConfigOnlineScorersGlobalType = "global"
)

func (r ProjectScoreConfigOnlineScorersGlobalType) IsKnown() bool {
	switch r {
	case ProjectScoreConfigOnlineScorersGlobalTypeGlobal:
		return true
	}
	return false
}

type ProjectScoreConfigOnlineScorersType string

const (
	ProjectScoreConfigOnlineScorersTypeFunction ProjectScoreConfigOnlineScorersType = "function"
	ProjectScoreConfigOnlineScorersTypeGlobal   ProjectScoreConfigOnlineScorersType = "global"
)

func (r ProjectScoreConfigOnlineScorersType) IsKnown() bool {
	switch r {
	case ProjectScoreConfigOnlineScorersTypeFunction, ProjectScoreConfigOnlineScorersTypeGlobal:
		return true
	}
	return false
}

// For categorical-type project scores, defines a single category
type ProjectScoreCategory struct {
	// Name of the category
	Name string `json:"name,required"`
	// Numerical value of the category. Must be between 0 and 1, inclusive
	Value float64                  `json:"value,required"`
	JSON  projectScoreCategoryJSON `json:"-"`
}

// projectScoreCategoryJSON contains the JSON metadata for the struct
// [ProjectScoreCategory]
type projectScoreCategoryJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectScoreCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectScoreCategoryJSON) RawJSON() string {
	return r.raw
}

// For categorical-type project scores, defines a single category
type ProjectScoreCategoryParam struct {
	// Name of the category
	Name param.Field[string] `json:"name,required"`
	// Numerical value of the category. Must be between 0 and 1, inclusive
	Value param.Field[float64] `json:"value,required"`
}

func (r ProjectScoreCategoryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A project tag is a user-configured tag for tracking and filtering your
// experiments, logs, and other data
type ProjectTag struct {
	// Unique identifier for the project tag
	ID string `json:"id,required" format:"uuid"`
	// Name of the project tag
	Name string `json:"name,required"`
	// Unique identifier for the project that the project tag belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	UserID    string `json:"user_id,required" format:"uuid"`
	// Color of the tag for the UI
	Color string `json:"color,nullable"`
	// Date of project tag creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Textual description of the project tag
	Description string         `json:"description,nullable"`
	JSON        projectTagJSON `json:"-"`
}

// projectTagJSON contains the JSON metadata for the struct [ProjectTag]
type projectTagJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	ProjectID   apijson.Field
	UserID      apijson.Field
	Color       apijson.Field
	Created     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectTagJSON) RawJSON() string {
	return r.raw
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
	Description  string             `json:"description,nullable"`
	FunctionType PromptFunctionType `json:"function_type,nullable"`
	// User-controlled metadata about the prompt
	Metadata map[string]interface{} `json:"metadata,nullable"`
	// The prompt, model, and its parameters
	PromptData PromptData `json:"prompt_data,nullable"`
	// A list of tags for the prompt
	Tags []string   `json:"tags,nullable"`
	JSON promptJSON `json:"-"`
}

// promptJSON contains the JSON metadata for the struct [Prompt]
type promptJSON struct {
	ID           apijson.Field
	XactID       apijson.Field
	LogID        apijson.Field
	Name         apijson.Field
	OrgID        apijson.Field
	ProjectID    apijson.Field
	Slug         apijson.Field
	Created      apijson.Field
	Description  apijson.Field
	FunctionType apijson.Field
	Metadata     apijson.Field
	PromptData   apijson.Field
	Tags         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
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

type PromptFunctionType string

const (
	PromptFunctionTypeLlm    PromptFunctionType = "llm"
	PromptFunctionTypeScorer PromptFunctionType = "scorer"
	PromptFunctionTypeTask   PromptFunctionType = "task"
	PromptFunctionTypeTool   PromptFunctionType = "tool"
)

func (r PromptFunctionType) IsKnown() bool {
	switch r {
	case PromptFunctionTypeLlm, PromptFunctionTypeScorer, PromptFunctionTypeTask, PromptFunctionTypeTool:
		return true
	}
	return false
}

// The prompt, model, and its parameters
type PromptData struct {
	Options       PromptDataOptions        `json:"options,nullable"`
	Origin        PromptDataOrigin         `json:"origin,nullable"`
	Parser        PromptDataParser         `json:"parser,nullable"`
	Prompt        PromptDataPrompt         `json:"prompt"`
	ToolFunctions []PromptDataToolFunction `json:"tool_functions,nullable"`
	JSON          promptDataJSON           `json:"-"`
}

// promptDataJSON contains the JSON metadata for the struct [PromptData]
type promptDataJSON struct {
	Options       apijson.Field
	Origin        apijson.Field
	Parser        apijson.Field
	Prompt        apijson.Field
	ToolFunctions apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

type PromptDataParser struct {
	ChoiceScores map[string]float64   `json:"choice_scores,required"`
	Type         PromptDataParserType `json:"type,required"`
	UseCot       bool                 `json:"use_cot,required"`
	JSON         promptDataParserJSON `json:"-"`
}

// promptDataParserJSON contains the JSON metadata for the struct
// [PromptDataParser]
type promptDataParserJSON struct {
	ChoiceScores apijson.Field
	Type         apijson.Field
	UseCot       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PromptDataParser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataParserJSON) RawJSON() string {
	return r.raw
}

type PromptDataParserType string

const (
	PromptDataParserTypeLlmClassifier PromptDataParserType = "llm_classifier"
)

func (r PromptDataParserType) IsKnown() bool {
	switch r {
	case PromptDataParserTypeLlmClassifier:
		return true
	}
	return false
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
	FunctionCall PromptDataPromptChatMessagesAssistantFunctionCall `json:"function_call,nullable"`
	Name         string                                            `json:"name,nullable"`
	ToolCalls    []PromptDataPromptChatMessagesAssistantToolCall   `json:"tool_calls,nullable"`
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

type PromptDataToolFunction struct {
	Type  PromptDataToolFunctionsType `json:"type,required"`
	ID    string                      `json:"id"`
	Name  string                      `json:"name"`
	JSON  promptDataToolFunctionJSON  `json:"-"`
	union PromptDataToolFunctionsUnion
}

// promptDataToolFunctionJSON contains the JSON metadata for the struct
// [PromptDataToolFunction]
type promptDataToolFunctionJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r promptDataToolFunctionJSON) RawJSON() string {
	return r.raw
}

func (r *PromptDataToolFunction) UnmarshalJSON(data []byte) (err error) {
	*r = PromptDataToolFunction{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PromptDataToolFunctionsUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.PromptDataToolFunctionsFunction],
// [shared.PromptDataToolFunctionsGlobal].
func (r PromptDataToolFunction) AsUnion() PromptDataToolFunctionsUnion {
	return r.union
}

// Union satisfied by [shared.PromptDataToolFunctionsFunction] or
// [shared.PromptDataToolFunctionsGlobal].
type PromptDataToolFunctionsUnion interface {
	implementsSharedPromptDataToolFunction()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PromptDataToolFunctionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataToolFunctionsFunction{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PromptDataToolFunctionsGlobal{}),
		},
	)
}

type PromptDataToolFunctionsFunction struct {
	ID   string                              `json:"id,required"`
	Type PromptDataToolFunctionsFunctionType `json:"type,required"`
	JSON promptDataToolFunctionsFunctionJSON `json:"-"`
}

// promptDataToolFunctionsFunctionJSON contains the JSON metadata for the struct
// [PromptDataToolFunctionsFunction]
type promptDataToolFunctionsFunctionJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataToolFunctionsFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataToolFunctionsFunctionJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataToolFunctionsFunction) implementsSharedPromptDataToolFunction() {}

type PromptDataToolFunctionsFunctionType string

const (
	PromptDataToolFunctionsFunctionTypeFunction PromptDataToolFunctionsFunctionType = "function"
)

func (r PromptDataToolFunctionsFunctionType) IsKnown() bool {
	switch r {
	case PromptDataToolFunctionsFunctionTypeFunction:
		return true
	}
	return false
}

type PromptDataToolFunctionsGlobal struct {
	Name string                            `json:"name,required"`
	Type PromptDataToolFunctionsGlobalType `json:"type,required"`
	JSON promptDataToolFunctionsGlobalJSON `json:"-"`
}

// promptDataToolFunctionsGlobalJSON contains the JSON metadata for the struct
// [PromptDataToolFunctionsGlobal]
type promptDataToolFunctionsGlobalJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PromptDataToolFunctionsGlobal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptDataToolFunctionsGlobalJSON) RawJSON() string {
	return r.raw
}

func (r PromptDataToolFunctionsGlobal) implementsSharedPromptDataToolFunction() {}

type PromptDataToolFunctionsGlobalType string

const (
	PromptDataToolFunctionsGlobalTypeGlobal PromptDataToolFunctionsGlobalType = "global"
)

func (r PromptDataToolFunctionsGlobalType) IsKnown() bool {
	switch r {
	case PromptDataToolFunctionsGlobalTypeGlobal:
		return true
	}
	return false
}

type PromptDataToolFunctionsType string

const (
	PromptDataToolFunctionsTypeFunction PromptDataToolFunctionsType = "function"
	PromptDataToolFunctionsTypeGlobal   PromptDataToolFunctionsType = "global"
)

func (r PromptDataToolFunctionsType) IsKnown() bool {
	switch r {
	case PromptDataToolFunctionsTypeFunction, PromptDataToolFunctionsTypeGlobal:
		return true
	}
	return false
}

// The prompt, model, and its parameters
type PromptDataParam struct {
	Options       param.Field[PromptDataOptionsParam]              `json:"options"`
	Origin        param.Field[PromptDataOriginParam]               `json:"origin"`
	Parser        param.Field[PromptDataParserParam]               `json:"parser"`
	Prompt        param.Field[PromptDataPromptUnionParam]          `json:"prompt"`
	ToolFunctions param.Field[[]PromptDataToolFunctionsUnionParam] `json:"tool_functions"`
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

type PromptDataParserParam struct {
	ChoiceScores param.Field[map[string]float64]   `json:"choice_scores,required"`
	Type         param.Field[PromptDataParserType] `json:"type,required"`
	UseCot       param.Field[bool]                 `json:"use_cot,required"`
}

func (r PromptDataParserParam) MarshalJSON() (data []byte, err error) {
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

type PromptDataToolFunctionParam struct {
	Type param.Field[PromptDataToolFunctionsType] `json:"type,required"`
	ID   param.Field[string]                      `json:"id"`
	Name param.Field[string]                      `json:"name"`
}

func (r PromptDataToolFunctionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataToolFunctionParam) implementsSharedPromptDataToolFunctionsUnionParam() {}

// Satisfied by [shared.PromptDataToolFunctionsFunctionParam],
// [shared.PromptDataToolFunctionsGlobalParam], [PromptDataToolFunctionParam].
type PromptDataToolFunctionsUnionParam interface {
	implementsSharedPromptDataToolFunctionsUnionParam()
}

type PromptDataToolFunctionsFunctionParam struct {
	ID   param.Field[string]                              `json:"id,required"`
	Type param.Field[PromptDataToolFunctionsFunctionType] `json:"type,required"`
}

func (r PromptDataToolFunctionsFunctionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataToolFunctionsFunctionParam) implementsSharedPromptDataToolFunctionsUnionParam() {}

type PromptDataToolFunctionsGlobalParam struct {
	Name param.Field[string]                            `json:"name,required"`
	Type param.Field[PromptDataToolFunctionsGlobalType] `json:"type,required"`
}

func (r PromptDataToolFunctionsGlobalParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PromptDataToolFunctionsGlobalParam) implementsSharedPromptDataToolFunctionsUnionParam() {}

// Metadata about the state of the repo when the experiment was created
type RepoInfo struct {
	// Email of the author of the most recent commit
	AuthorEmail string `json:"author_email,nullable"`
	// Name of the author of the most recent commit
	AuthorName string `json:"author_name,nullable"`
	// Name of the branch the most recent commit belongs to
	Branch string `json:"branch,nullable"`
	// SHA of most recent commit
	Commit string `json:"commit,nullable"`
	// Most recent commit message
	CommitMessage string `json:"commit_message,nullable"`
	// Time of the most recent commit
	CommitTime string `json:"commit_time,nullable"`
	// Whether or not the repo had uncommitted changes when snapshotted
	Dirty bool `json:"dirty,nullable"`
	// If the repo was dirty when run, this includes the diff between the current state
	// of the repo and the most recent commit.
	GitDiff string `json:"git_diff,nullable"`
	// Name of the tag on the most recent commit
	Tag  string       `json:"tag,nullable"`
	JSON repoInfoJSON `json:"-"`
}

// repoInfoJSON contains the JSON metadata for the struct [RepoInfo]
type repoInfoJSON struct {
	AuthorEmail   apijson.Field
	AuthorName    apijson.Field
	Branch        apijson.Field
	Commit        apijson.Field
	CommitMessage apijson.Field
	CommitTime    apijson.Field
	Dirty         apijson.Field
	GitDiff       apijson.Field
	Tag           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RepoInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r repoInfoJSON) RawJSON() string {
	return r.raw
}

// Metadata about the state of the repo when the experiment was created
type RepoInfoParam struct {
	// Email of the author of the most recent commit
	AuthorEmail param.Field[string] `json:"author_email"`
	// Name of the author of the most recent commit
	AuthorName param.Field[string] `json:"author_name"`
	// Name of the branch the most recent commit belongs to
	Branch param.Field[string] `json:"branch"`
	// SHA of most recent commit
	Commit param.Field[string] `json:"commit"`
	// Most recent commit message
	CommitMessage param.Field[string] `json:"commit_message"`
	// Time of the most recent commit
	CommitTime param.Field[string] `json:"commit_time"`
	// Whether or not the repo had uncommitted changes when snapshotted
	Dirty param.Field[bool] `json:"dirty"`
	// If the repo was dirty when run, this includes the diff between the current state
	// of the repo and the most recent commit.
	GitDiff param.Field[string] `json:"git_diff"`
	// Name of the tag on the most recent commit
	Tag param.Field[string] `json:"tag"`
}

func (r RepoInfoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A role is a collection of permissions which can be granted as part of an ACL
//
// Roles can consist of individual permissions, as well as a set of roles they
// inherit from
type Role struct {
	// Unique identifier for the role
	ID string `json:"id,required" format:"uuid"`
	// Name of the role
	Name string `json:"name,required"`
	// Date of role creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Date of role deletion, or null if the role is still active
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Textual description of the role
	Description string `json:"description,nullable"`
	// (permission, restrict_object_type) tuples which belong to this role
	MemberPermissions []RoleMemberPermission `json:"member_permissions,nullable"`
	// Ids of the roles this role inherits from
	//
	// An inheriting role has all the permissions contained in its member roles, as
	// well as all of their inherited permissions
	MemberRoles []string `json:"member_roles,nullable" format:"uuid"`
	// Unique id for the organization that the role belongs under
	//
	// A null org_id indicates a system role, which may be assigned to anybody and
	// inherited by any other role, but cannot be edited.
	//
	// It is forbidden to change the org after creating a role
	OrgID string `json:"org_id,nullable" format:"uuid"`
	// Identifies the user who created the role
	UserID string   `json:"user_id,nullable" format:"uuid"`
	JSON   roleJSON `json:"-"`
}

// roleJSON contains the JSON metadata for the struct [Role]
type roleJSON struct {
	ID                apijson.Field
	Name              apijson.Field
	Created           apijson.Field
	DeletedAt         apijson.Field
	Description       apijson.Field
	MemberPermissions apijson.Field
	MemberRoles       apijson.Field
	OrgID             apijson.Field
	UserID            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Role) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleJSON) RawJSON() string {
	return r.raw
}

type RoleMemberPermission struct {
	// Each permission permits a certain type of operation on an object in the system
	//
	// Permissions can be assigned to to objects on an individual basis, or grouped
	// into roles
	Permission RoleMemberPermissionsPermission `json:"permission,required,nullable"`
	// The object type that the ACL applies to
	RestrictObjectType RoleMemberPermissionsRestrictObjectType `json:"restrict_object_type,nullable"`
	JSON               roleMemberPermissionJSON                `json:"-"`
}

// roleMemberPermissionJSON contains the JSON metadata for the struct
// [RoleMemberPermission]
type roleMemberPermissionJSON struct {
	Permission         apijson.Field
	RestrictObjectType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RoleMemberPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleMemberPermissionJSON) RawJSON() string {
	return r.raw
}

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
type RoleMemberPermissionsPermission string

const (
	RoleMemberPermissionsPermissionCreate     RoleMemberPermissionsPermission = "create"
	RoleMemberPermissionsPermissionRead       RoleMemberPermissionsPermission = "read"
	RoleMemberPermissionsPermissionUpdate     RoleMemberPermissionsPermission = "update"
	RoleMemberPermissionsPermissionDelete     RoleMemberPermissionsPermission = "delete"
	RoleMemberPermissionsPermissionCreateACLs RoleMemberPermissionsPermission = "create_acls"
	RoleMemberPermissionsPermissionReadACLs   RoleMemberPermissionsPermission = "read_acls"
	RoleMemberPermissionsPermissionUpdateACLs RoleMemberPermissionsPermission = "update_acls"
	RoleMemberPermissionsPermissionDeleteACLs RoleMemberPermissionsPermission = "delete_acls"
)

func (r RoleMemberPermissionsPermission) IsKnown() bool {
	switch r {
	case RoleMemberPermissionsPermissionCreate, RoleMemberPermissionsPermissionRead, RoleMemberPermissionsPermissionUpdate, RoleMemberPermissionsPermissionDelete, RoleMemberPermissionsPermissionCreateACLs, RoleMemberPermissionsPermissionReadACLs, RoleMemberPermissionsPermissionUpdateACLs, RoleMemberPermissionsPermissionDeleteACLs:
		return true
	}
	return false
}

// The object type that the ACL applies to
type RoleMemberPermissionsRestrictObjectType string

const (
	RoleMemberPermissionsRestrictObjectTypeOrganization  RoleMemberPermissionsRestrictObjectType = "organization"
	RoleMemberPermissionsRestrictObjectTypeProject       RoleMemberPermissionsRestrictObjectType = "project"
	RoleMemberPermissionsRestrictObjectTypeExperiment    RoleMemberPermissionsRestrictObjectType = "experiment"
	RoleMemberPermissionsRestrictObjectTypeDataset       RoleMemberPermissionsRestrictObjectType = "dataset"
	RoleMemberPermissionsRestrictObjectTypePrompt        RoleMemberPermissionsRestrictObjectType = "prompt"
	RoleMemberPermissionsRestrictObjectTypePromptSession RoleMemberPermissionsRestrictObjectType = "prompt_session"
	RoleMemberPermissionsRestrictObjectTypeGroup         RoleMemberPermissionsRestrictObjectType = "group"
	RoleMemberPermissionsRestrictObjectTypeRole          RoleMemberPermissionsRestrictObjectType = "role"
	RoleMemberPermissionsRestrictObjectTypeOrgMember     RoleMemberPermissionsRestrictObjectType = "org_member"
	RoleMemberPermissionsRestrictObjectTypeProjectLog    RoleMemberPermissionsRestrictObjectType = "project_log"
	RoleMemberPermissionsRestrictObjectTypeOrgProject    RoleMemberPermissionsRestrictObjectType = "org_project"
)

func (r RoleMemberPermissionsRestrictObjectType) IsKnown() bool {
	switch r {
	case RoleMemberPermissionsRestrictObjectTypeOrganization, RoleMemberPermissionsRestrictObjectTypeProject, RoleMemberPermissionsRestrictObjectTypeExperiment, RoleMemberPermissionsRestrictObjectTypeDataset, RoleMemberPermissionsRestrictObjectTypePrompt, RoleMemberPermissionsRestrictObjectTypePromptSession, RoleMemberPermissionsRestrictObjectTypeGroup, RoleMemberPermissionsRestrictObjectTypeRole, RoleMemberPermissionsRestrictObjectTypeOrgMember, RoleMemberPermissionsRestrictObjectTypeProjectLog, RoleMemberPermissionsRestrictObjectTypeOrgProject:
		return true
	}
	return false
}

// Summary of a score's performance
type ScoreSummary struct {
	// Number of improvements in the score
	Improvements int64 `json:"improvements,required"`
	// Name of the score
	Name string `json:"name,required"`
	// Number of regressions in the score
	Regressions int64 `json:"regressions,required"`
	// Average score across all examples
	Score float64 `json:"score,required"`
	// Difference in score between the current and comparison experiment
	Diff float64          `json:"diff"`
	JSON scoreSummaryJSON `json:"-"`
}

// scoreSummaryJSON contains the JSON metadata for the struct [ScoreSummary]
type scoreSummaryJSON struct {
	Improvements apijson.Field
	Name         apijson.Field
	Regressions  apijson.Field
	Score        apijson.Field
	Diff         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScoreSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scoreSummaryJSON) RawJSON() string {
	return r.raw
}

// Summary of a dataset
type SummarizeDatasetResponse struct {
	// Name of the dataset
	DatasetName string `json:"dataset_name,required"`
	// URL to the dataset's page in the Braintrust app
	DatasetURL string `json:"dataset_url,required" format:"uri"`
	// Name of the project that the dataset belongs to
	ProjectName string `json:"project_name,required"`
	// URL to the project's page in the Braintrust app
	ProjectURL string `json:"project_url,required" format:"uri"`
	// Summary of a dataset's data
	DataSummary DataSummary                  `json:"data_summary,nullable"`
	JSON        summarizeDatasetResponseJSON `json:"-"`
}

// summarizeDatasetResponseJSON contains the JSON metadata for the struct
// [SummarizeDatasetResponse]
type summarizeDatasetResponseJSON struct {
	DatasetName apijson.Field
	DatasetURL  apijson.Field
	ProjectName apijson.Field
	ProjectURL  apijson.Field
	DataSummary apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SummarizeDatasetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r summarizeDatasetResponseJSON) RawJSON() string {
	return r.raw
}

// Summary of an experiment
type SummarizeExperimentResponse struct {
	// Name of the experiment
	ExperimentName string `json:"experiment_name,required"`
	// URL to the experiment's page in the Braintrust app
	ExperimentURL string `json:"experiment_url,required" format:"uri"`
	// Name of the project that the experiment belongs to
	ProjectName string `json:"project_name,required"`
	// URL to the project's page in the Braintrust app
	ProjectURL string `json:"project_url,required" format:"uri"`
	// The experiment which scores are baselined against
	ComparisonExperimentName string `json:"comparison_experiment_name,nullable"`
	// Summary of the experiment's metrics
	Metrics map[string]MetricSummary `json:"metrics,nullable"`
	// Summary of the experiment's scores
	Scores map[string]ScoreSummary         `json:"scores,nullable"`
	JSON   summarizeExperimentResponseJSON `json:"-"`
}

// summarizeExperimentResponseJSON contains the JSON metadata for the struct
// [SummarizeExperimentResponse]
type summarizeExperimentResponseJSON struct {
	ExperimentName           apijson.Field
	ExperimentURL            apijson.Field
	ProjectName              apijson.Field
	ProjectURL               apijson.Field
	ComparisonExperimentName apijson.Field
	Metrics                  apijson.Field
	Scores                   apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *SummarizeExperimentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r summarizeExperimentResponseJSON) RawJSON() string {
	return r.raw
}

type User struct {
	// Unique identifier for the user
	ID string `json:"id,required" format:"uuid"`
	// URL of the user's Avatar image
	AvatarURL string `json:"avatar_url,nullable"`
	// Date of user creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// The user's email
	Email string `json:"email,nullable"`
	// Family name of the user
	FamilyName string `json:"family_name,nullable"`
	// Given name of the user
	GivenName string   `json:"given_name,nullable"`
	JSON      userJSON `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	ID          apijson.Field
	AvatarURL   apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	FamilyName  apijson.Field
	GivenName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *User) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userJSON) RawJSON() string {
	return r.raw
}

type View struct {
	// Unique identifier for the view
	ID string `json:"id,required" format:"uuid"`
	// Name of the view
	Name string `json:"name,required"`
	// The id of the object the view applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType ViewObjectType `json:"object_type,required,nullable"`
	// Type of table that the view corresponds to.
	ViewType ViewViewType `json:"view_type,required,nullable"`
	// Date of view creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Date of role deletion, or null if the role is still active
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Options for the view in the app
	Options ViewOptions `json:"options,nullable"`
	// Identifies the user who created the view
	UserID string `json:"user_id,nullable" format:"uuid"`
	// The view definition
	ViewData ViewData `json:"view_data,nullable"`
	JSON     viewJSON `json:"-"`
}

// viewJSON contains the JSON metadata for the struct [View]
type viewJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	ObjectID    apijson.Field
	ObjectType  apijson.Field
	ViewType    apijson.Field
	Created     apijson.Field
	DeletedAt   apijson.Field
	Options     apijson.Field
	UserID      apijson.Field
	ViewData    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *View) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewJSON) RawJSON() string {
	return r.raw
}

// The object type that the ACL applies to
type ViewObjectType string

const (
	ViewObjectTypeOrganization  ViewObjectType = "organization"
	ViewObjectTypeProject       ViewObjectType = "project"
	ViewObjectTypeExperiment    ViewObjectType = "experiment"
	ViewObjectTypeDataset       ViewObjectType = "dataset"
	ViewObjectTypePrompt        ViewObjectType = "prompt"
	ViewObjectTypePromptSession ViewObjectType = "prompt_session"
	ViewObjectTypeGroup         ViewObjectType = "group"
	ViewObjectTypeRole          ViewObjectType = "role"
	ViewObjectTypeOrgMember     ViewObjectType = "org_member"
	ViewObjectTypeProjectLog    ViewObjectType = "project_log"
	ViewObjectTypeOrgProject    ViewObjectType = "org_project"
)

func (r ViewObjectType) IsKnown() bool {
	switch r {
	case ViewObjectTypeOrganization, ViewObjectTypeProject, ViewObjectTypeExperiment, ViewObjectTypeDataset, ViewObjectTypePrompt, ViewObjectTypePromptSession, ViewObjectTypeGroup, ViewObjectTypeRole, ViewObjectTypeOrgMember, ViewObjectTypeProjectLog, ViewObjectTypeOrgProject:
		return true
	}
	return false
}

// Type of table that the view corresponds to.
type ViewViewType string

const (
	ViewViewTypeProjects    ViewViewType = "projects"
	ViewViewTypeLogs        ViewViewType = "logs"
	ViewViewTypeExperiments ViewViewType = "experiments"
	ViewViewTypeDatasets    ViewViewType = "datasets"
	ViewViewTypePrompts     ViewViewType = "prompts"
	ViewViewTypePlaygrounds ViewViewType = "playgrounds"
	ViewViewTypeExperiment  ViewViewType = "experiment"
	ViewViewTypeDataset     ViewViewType = "dataset"
)

func (r ViewViewType) IsKnown() bool {
	switch r {
	case ViewViewTypeProjects, ViewViewTypeLogs, ViewViewTypeExperiments, ViewViewTypeDatasets, ViewViewTypePrompts, ViewViewTypePlaygrounds, ViewViewTypeExperiment, ViewViewTypeDataset:
		return true
	}
	return false
}

// The view definition
type ViewData struct {
	Search ViewDataSearch `json:"search,nullable"`
	JSON   viewDataJSON   `json:"-"`
}

// viewDataJSON contains the JSON metadata for the struct [ViewData]
type viewDataJSON struct {
	Search      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ViewData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewDataJSON) RawJSON() string {
	return r.raw
}

// The view definition
type ViewDataParam struct {
	Search param.Field[ViewDataSearchParam] `json:"search"`
}

func (r ViewDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ViewDataSearch struct {
	Filter []interface{}      `json:"filter,nullable"`
	Match  []interface{}      `json:"match,nullable"`
	Sort   []interface{}      `json:"sort,nullable"`
	Tag    []interface{}      `json:"tag,nullable"`
	JSON   viewDataSearchJSON `json:"-"`
}

// viewDataSearchJSON contains the JSON metadata for the struct [ViewDataSearch]
type viewDataSearchJSON struct {
	Filter      apijson.Field
	Match       apijson.Field
	Sort        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ViewDataSearch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewDataSearchJSON) RawJSON() string {
	return r.raw
}

type ViewDataSearchParam struct {
	Filter param.Field[[]interface{}] `json:"filter"`
	Match  param.Field[[]interface{}] `json:"match"`
	Sort   param.Field[[]interface{}] `json:"sort"`
	Tag    param.Field[[]interface{}] `json:"tag"`
}

func (r ViewDataSearchParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options for the view in the app
type ViewOptions struct {
	ColumnOrder      []string           `json:"columnOrder,nullable"`
	ColumnSizing     map[string]float64 `json:"columnSizing,nullable"`
	ColumnVisibility map[string]bool    `json:"columnVisibility,nullable"`
	JSON             viewOptionsJSON    `json:"-"`
}

// viewOptionsJSON contains the JSON metadata for the struct [ViewOptions]
type viewOptionsJSON struct {
	ColumnOrder      apijson.Field
	ColumnSizing     apijson.Field
	ColumnVisibility apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ViewOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewOptionsJSON) RawJSON() string {
	return r.raw
}

// Options for the view in the app
type ViewOptionsParam struct {
	ColumnOrder      param.Field[[]string]           `json:"columnOrder"`
	ColumnSizing     param.Field[map[string]float64] `json:"columnSizing"`
	ColumnVisibility param.Field[map[string]bool]    `json:"columnVisibility"`
}

func (r ViewOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
