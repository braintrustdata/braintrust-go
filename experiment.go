// File generated from our OpenAPI spec by Stainless.

package braintrust

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/internal/apiquery"
	"github.com/braintrustdata/braintrust-go/internal/param"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/internal/shared"
	"github.com/braintrustdata/braintrust-go/option"
)

// ExperimentService contains methods and other services that help with interacting
// with the braintrust API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewExperimentService] method instead.
type ExperimentService struct {
	Options []option.RequestOption
}

// NewExperimentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExperimentService(opts ...option.RequestOption) (r *ExperimentService) {
	r = &ExperimentService{}
	r.Options = opts
	return
}

// Create a new experiment. If there is an existing experiment in the project with
// the same name as the one specified in the request, will create a new experiment
// from `name`, suffixed with a unique identifier
func (r *ExperimentService) New(ctx context.Context, body ExperimentNewParams, opts ...option.RequestOption) (res *Experiment, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/experiment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an experiment object by its id
func (r *ExperimentService) Get(ctx context.Context, experimentID string, opts ...option.RequestOption) (res *Experiment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/experiment/%s", experimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update an experiment object. Specify the fields to update in the
// payload. Any object-type fields will be deep-merged with existing content.
// Currently we do not support removing fields or setting them to null. As a
// workaround, you may retrieve the full object with `GET /experiment/{id}` and
// then replace it with `PUT /experiment`.
func (r *ExperimentService) Update(ctx context.Context, experimentID string, body ExperimentUpdateParams, opts ...option.RequestOption) (res *Experiment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/experiment/%s", experimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all experiments. The experiments are sorted by creation date, with the
// most recently-created experiments coming first
func (r *ExperimentService) List(ctx context.Context, query ExperimentListParams, opts ...option.RequestOption) (res *shared.ListObjects[Experiment], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/experiment"
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

// List out all experiments. The experiments are sorted by creation date, with the
// most recently-created experiments coming first
func (r *ExperimentService) ListAutoPaging(ctx context.Context, query ExperimentListParams, opts ...option.RequestOption) *shared.ListObjectsAutoPager[Experiment] {
	return shared.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete an experiment object by its id
func (r *ExperimentService) Delete(ctx context.Context, experimentID string, opts ...option.RequestOption) (res *Experiment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/experiment/%s", experimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Log feedback for a set of experiment events
func (r *ExperimentService) Feedback(ctx context.Context, experimentID string, body ExperimentFeedbackParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("v1/experiment/%s/feedback", experimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Fetch the events in an experiment. Equivalent to the POST form of the same path,
// but with the parameters in the URL query rather than in the request body
func (r *ExperimentService) Fetch(ctx context.Context, experimentID string, query ExperimentFetchParams, opts ...option.RequestOption) (res *ExperimentFetchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/experiment/%s/fetch", experimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch the events in an experiment. Equivalent to the GET form of the same path,
// but with the parameters in the request body rather than in the URL query
func (r *ExperimentService) FetchPost(ctx context.Context, experimentID string, body ExperimentFetchPostParams, opts ...option.RequestOption) (res *ExperimentFetchPostResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/experiment/%s/fetch", experimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Insert a set of events into the experiment
func (r *ExperimentService) Insert(ctx context.Context, experimentID string, body ExperimentInsertParams, opts ...option.RequestOption) (res *ExperimentInsertResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/experiment/%s/insert", experimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create or replace a new experiment. If there is an existing experiment in the
// project with the same name as the one specified in the request, will replace the
// existing experiment with the provided fields
func (r *ExperimentService) Replace(ctx context.Context, body ExperimentReplaceParams, opts ...option.RequestOption) (res *Experiment, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/experiment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
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
	RepoInfo ExperimentRepoInfo `json:"repo_info,nullable"`
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

// Metadata about the state of the repo when the experiment was created
type ExperimentRepoInfo struct {
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
	Tag  string                 `json:"tag,nullable"`
	JSON experimentRepoInfoJSON `json:"-"`
}

// experimentRepoInfoJSON contains the JSON metadata for the struct
// [ExperimentRepoInfo]
type experimentRepoInfoJSON struct {
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

func (r *ExperimentRepoInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentFetchResponse struct {
	// A list of fetched events
	Events []ExperimentFetchResponseEvent `json:"events,required"`
	JSON   experimentFetchResponseJSON    `json:"-"`
}

// experimentFetchResponseJSON contains the JSON metadata for the struct
// [ExperimentFetchResponse]
type experimentFetchResponseJSON struct {
	Events      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExperimentFetchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentFetchResponseEvent struct {
	// A unique identifier for the experiment event. If you don't provide one,
	// BrainTrust will generate one for you
	ID string `json:"id,required"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the experiment (see the
	// `version` parameter)
	XactID int64 `json:"_xact_id,required"`
	// Unique identifier for the experiment
	ExperimentID string `json:"experiment_id,required" format:"uuid"`
	// Unique identifier for the project that the experiment belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// The `span_id` of the root of the trace this experiment event belongs to
	RootSpanID string `json:"root_span_id,required"`
	// A unique identifier used to link different experiment events together as part of
	// a full trace. See the
	// [tracing guide](https://www.braintrustdata.com/docs/guides/tracing) for full
	// details on tracing
	SpanID string `json:"span_id,required"`
	// Context is additional information about the code that produced the experiment
	// event. It is essentially the textual counterpart to `metrics`. Use the
	// `caller_*` attributes to track the location in code which produced the
	// experiment event
	Context ExperimentFetchResponseEventsContext `json:"context,nullable"`
	// The timestamp the experiment event was created
	Created time.Time `json:"created,nullable" format:"date-time"`
	// If the experiment is associated to a dataset, this is the event-level dataset id
	// this experiment event is tied to
	DatasetRecordID string `json:"dataset_record_id,nullable"`
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
	Metrics ExperimentFetchResponseEventsMetrics `json:"metrics,nullable"`
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
	SpanAttributes ExperimentFetchResponseEventsSpanAttributes `json:"span_attributes,nullable"`
	// An array of the parent `span_ids` of this experiment event. This should be empty
	// for the root span of a trace, and should most often contain just one parent
	// element for subspans
	SpanParents []string                         `json:"span_parents,nullable"`
	JSON        experimentFetchResponseEventJSON `json:"-"`
}

// experimentFetchResponseEventJSON contains the JSON metadata for the struct
// [ExperimentFetchResponseEvent]
type experimentFetchResponseEventJSON struct {
	ID              apijson.Field
	XactID          apijson.Field
	ExperimentID    apijson.Field
	ProjectID       apijson.Field
	RootSpanID      apijson.Field
	SpanID          apijson.Field
	Context         apijson.Field
	Created         apijson.Field
	DatasetRecordID apijson.Field
	Expected        apijson.Field
	Input           apijson.Field
	Metadata        apijson.Field
	Metrics         apijson.Field
	Output          apijson.Field
	Scores          apijson.Field
	SpanAttributes  apijson.Field
	SpanParents     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ExperimentFetchResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Context is additional information about the code that produced the experiment
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the
// experiment event
type ExperimentFetchResponseEventsContext struct {
	// Name of the file in code where the experiment event was created
	CallerFilename string `json:"caller_filename,nullable"`
	// The function in code which created the experiment event
	CallerFunctionname string `json:"caller_functionname,nullable"`
	// Line of code where the experiment event was created
	CallerLineno int64                                    `json:"caller_lineno,nullable"`
	ExtraFields  map[string]interface{}                   `json:"-,extras"`
	JSON         experimentFetchResponseEventsContextJSON `json:"-"`
}

// experimentFetchResponseEventsContextJSON contains the JSON metadata for the
// struct [ExperimentFetchResponseEventsContext]
type experimentFetchResponseEventsContextJSON struct {
	CallerFilename     apijson.Field
	CallerFunctionname apijson.Field
	CallerLineno       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ExperimentFetchResponseEventsContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the experiment event. Use "start" and "end" to track the time span over
// which the experiment event was produced
type ExperimentFetchResponseEventsMetrics struct {
	// A unix timestamp recording when the section of code which produced the
	// experiment event finished
	End float64 `json:"end,nullable"`
	// A unix timestamp recording when the section of code which produced the
	// experiment event started
	Start       float64                                  `json:"start,nullable"`
	ExtraFields map[string]interface{}                   `json:"-,extras"`
	JSON        experimentFetchResponseEventsMetricsJSON `json:"-"`
}

// experimentFetchResponseEventsMetricsJSON contains the JSON metadata for the
// struct [ExperimentFetchResponseEventsMetrics]
type experimentFetchResponseEventsMetricsJSON struct {
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExperimentFetchResponseEventsMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Human-identifying attributes of the span, such as name, type, etc.
type ExperimentFetchResponseEventsSpanAttributes struct {
	// Name of the span, for display purposes only
	Name string `json:"name,nullable"`
	// Type of the span, for display purposes only
	Type        ExperimentFetchResponseEventsSpanAttributesType `json:"type,nullable"`
	ExtraFields map[string]interface{}                          `json:"-,extras"`
	JSON        experimentFetchResponseEventsSpanAttributesJSON `json:"-"`
}

// experimentFetchResponseEventsSpanAttributesJSON contains the JSON metadata for
// the struct [ExperimentFetchResponseEventsSpanAttributes]
type experimentFetchResponseEventsSpanAttributesJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExperimentFetchResponseEventsSpanAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the span, for display purposes only
type ExperimentFetchResponseEventsSpanAttributesType string

const (
	ExperimentFetchResponseEventsSpanAttributesTypeLlm      ExperimentFetchResponseEventsSpanAttributesType = "llm"
	ExperimentFetchResponseEventsSpanAttributesTypeScore    ExperimentFetchResponseEventsSpanAttributesType = "score"
	ExperimentFetchResponseEventsSpanAttributesTypeFunction ExperimentFetchResponseEventsSpanAttributesType = "function"
	ExperimentFetchResponseEventsSpanAttributesTypeEval     ExperimentFetchResponseEventsSpanAttributesType = "eval"
	ExperimentFetchResponseEventsSpanAttributesTypeTask     ExperimentFetchResponseEventsSpanAttributesType = "task"
	ExperimentFetchResponseEventsSpanAttributesTypeTool     ExperimentFetchResponseEventsSpanAttributesType = "tool"
)

type ExperimentFetchPostResponse struct {
	// A list of fetched events
	Events []ExperimentFetchPostResponseEvent `json:"events,required"`
	JSON   experimentFetchPostResponseJSON    `json:"-"`
}

// experimentFetchPostResponseJSON contains the JSON metadata for the struct
// [ExperimentFetchPostResponse]
type experimentFetchPostResponseJSON struct {
	Events      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExperimentFetchPostResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentFetchPostResponseEvent struct {
	// A unique identifier for the experiment event. If you don't provide one,
	// BrainTrust will generate one for you
	ID string `json:"id,required"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the experiment (see the
	// `version` parameter)
	XactID int64 `json:"_xact_id,required"`
	// Unique identifier for the experiment
	ExperimentID string `json:"experiment_id,required" format:"uuid"`
	// Unique identifier for the project that the experiment belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// The `span_id` of the root of the trace this experiment event belongs to
	RootSpanID string `json:"root_span_id,required"`
	// A unique identifier used to link different experiment events together as part of
	// a full trace. See the
	// [tracing guide](https://www.braintrustdata.com/docs/guides/tracing) for full
	// details on tracing
	SpanID string `json:"span_id,required"`
	// Context is additional information about the code that produced the experiment
	// event. It is essentially the textual counterpart to `metrics`. Use the
	// `caller_*` attributes to track the location in code which produced the
	// experiment event
	Context ExperimentFetchPostResponseEventsContext `json:"context,nullable"`
	// The timestamp the experiment event was created
	Created time.Time `json:"created,nullable" format:"date-time"`
	// If the experiment is associated to a dataset, this is the event-level dataset id
	// this experiment event is tied to
	DatasetRecordID string `json:"dataset_record_id,nullable"`
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
	Metrics ExperimentFetchPostResponseEventsMetrics `json:"metrics,nullable"`
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
	SpanAttributes ExperimentFetchPostResponseEventsSpanAttributes `json:"span_attributes,nullable"`
	// An array of the parent `span_ids` of this experiment event. This should be empty
	// for the root span of a trace, and should most often contain just one parent
	// element for subspans
	SpanParents []string                             `json:"span_parents,nullable"`
	JSON        experimentFetchPostResponseEventJSON `json:"-"`
}

// experimentFetchPostResponseEventJSON contains the JSON metadata for the struct
// [ExperimentFetchPostResponseEvent]
type experimentFetchPostResponseEventJSON struct {
	ID              apijson.Field
	XactID          apijson.Field
	ExperimentID    apijson.Field
	ProjectID       apijson.Field
	RootSpanID      apijson.Field
	SpanID          apijson.Field
	Context         apijson.Field
	Created         apijson.Field
	DatasetRecordID apijson.Field
	Expected        apijson.Field
	Input           apijson.Field
	Metadata        apijson.Field
	Metrics         apijson.Field
	Output          apijson.Field
	Scores          apijson.Field
	SpanAttributes  apijson.Field
	SpanParents     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ExperimentFetchPostResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Context is additional information about the code that produced the experiment
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the
// experiment event
type ExperimentFetchPostResponseEventsContext struct {
	// Name of the file in code where the experiment event was created
	CallerFilename string `json:"caller_filename,nullable"`
	// The function in code which created the experiment event
	CallerFunctionname string `json:"caller_functionname,nullable"`
	// Line of code where the experiment event was created
	CallerLineno int64                                        `json:"caller_lineno,nullable"`
	ExtraFields  map[string]interface{}                       `json:"-,extras"`
	JSON         experimentFetchPostResponseEventsContextJSON `json:"-"`
}

// experimentFetchPostResponseEventsContextJSON contains the JSON metadata for the
// struct [ExperimentFetchPostResponseEventsContext]
type experimentFetchPostResponseEventsContextJSON struct {
	CallerFilename     apijson.Field
	CallerFunctionname apijson.Field
	CallerLineno       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ExperimentFetchPostResponseEventsContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the experiment event. Use "start" and "end" to track the time span over
// which the experiment event was produced
type ExperimentFetchPostResponseEventsMetrics struct {
	// A unix timestamp recording when the section of code which produced the
	// experiment event finished
	End float64 `json:"end,nullable"`
	// A unix timestamp recording when the section of code which produced the
	// experiment event started
	Start       float64                                      `json:"start,nullable"`
	ExtraFields map[string]interface{}                       `json:"-,extras"`
	JSON        experimentFetchPostResponseEventsMetricsJSON `json:"-"`
}

// experimentFetchPostResponseEventsMetricsJSON contains the JSON metadata for the
// struct [ExperimentFetchPostResponseEventsMetrics]
type experimentFetchPostResponseEventsMetricsJSON struct {
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExperimentFetchPostResponseEventsMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Human-identifying attributes of the span, such as name, type, etc.
type ExperimentFetchPostResponseEventsSpanAttributes struct {
	// Name of the span, for display purposes only
	Name string `json:"name,nullable"`
	// Type of the span, for display purposes only
	Type        ExperimentFetchPostResponseEventsSpanAttributesType `json:"type,nullable"`
	ExtraFields map[string]interface{}                              `json:"-,extras"`
	JSON        experimentFetchPostResponseEventsSpanAttributesJSON `json:"-"`
}

// experimentFetchPostResponseEventsSpanAttributesJSON contains the JSON metadata
// for the struct [ExperimentFetchPostResponseEventsSpanAttributes]
type experimentFetchPostResponseEventsSpanAttributesJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExperimentFetchPostResponseEventsSpanAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the span, for display purposes only
type ExperimentFetchPostResponseEventsSpanAttributesType string

const (
	ExperimentFetchPostResponseEventsSpanAttributesTypeLlm      ExperimentFetchPostResponseEventsSpanAttributesType = "llm"
	ExperimentFetchPostResponseEventsSpanAttributesTypeScore    ExperimentFetchPostResponseEventsSpanAttributesType = "score"
	ExperimentFetchPostResponseEventsSpanAttributesTypeFunction ExperimentFetchPostResponseEventsSpanAttributesType = "function"
	ExperimentFetchPostResponseEventsSpanAttributesTypeEval     ExperimentFetchPostResponseEventsSpanAttributesType = "eval"
	ExperimentFetchPostResponseEventsSpanAttributesTypeTask     ExperimentFetchPostResponseEventsSpanAttributesType = "task"
	ExperimentFetchPostResponseEventsSpanAttributesTypeTool     ExperimentFetchPostResponseEventsSpanAttributesType = "tool"
)

type ExperimentInsertResponse struct {
	// The ids of all rows that were inserted, aligning one-to-one with the rows
	// provided as input
	RowIDs []string                     `json:"row_ids,required"`
	JSON   experimentInsertResponseJSON `json:"-"`
}

// experimentInsertResponseJSON contains the JSON metadata for the struct
// [ExperimentInsertResponse]
type experimentInsertResponseJSON struct {
	RowIDs      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExperimentInsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentNewParams struct {
	// Unique identifier for the project that the experiment belongs under
	ProjectID param.Field[string] `json:"project_id,required" format:"uuid"`
	// Id of default base experiment to compare against when viewing this experiment
	BaseExpID param.Field[string] `json:"base_exp_id" format:"uuid"`
	// Identifier of the linked dataset, or null if the experiment is not linked to a
	// dataset
	DatasetID param.Field[string] `json:"dataset_id" format:"uuid"`
	// Version number of the linked dataset the experiment was run against. This can be
	// used to reproduce the experiment after the dataset has been modified.
	DatasetVersion param.Field[string] `json:"dataset_version"`
	// Textual description of the experiment
	Description param.Field[string] `json:"description"`
	// User-controlled metadata about the experiment
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// Name of the experiment. Within a project, experiment names are unique
	Name param.Field[string] `json:"name"`
	// Whether or not the experiment is public. Public experiments can be viewed by
	// anybody inside or outside the organization
	Public param.Field[bool] `json:"public"`
	// Metadata about the state of the repo when the experiment was created
	RepoInfo param.Field[ExperimentNewParamsRepoInfo] `json:"repo_info"`
}

func (r ExperimentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metadata about the state of the repo when the experiment was created
type ExperimentNewParamsRepoInfo struct {
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

func (r ExperimentNewParamsRepoInfo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExperimentUpdateParams struct {
	// Id of default base experiment to compare against when viewing this experiment
	BaseExpID param.Field[string] `json:"base_exp_id" format:"uuid"`
	// Identifier of the linked dataset, or null if the experiment is not linked to a
	// dataset
	DatasetID param.Field[string] `json:"dataset_id" format:"uuid"`
	// Version number of the linked dataset the experiment was run against. This can be
	// used to reproduce the experiment after the dataset has been modified.
	DatasetVersion param.Field[string] `json:"dataset_version"`
	// Textual description of the experiment
	Description param.Field[string] `json:"description"`
	// User-controlled metadata about the experiment
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// Name of the experiment. Within a project, experiment names are unique
	Name param.Field[string] `json:"name"`
	// Whether or not the experiment is public. Public experiments can be viewed by
	// anybody inside or outside the organization
	Public param.Field[bool] `json:"public"`
	// Metadata about the state of the repo when the experiment was created
	RepoInfo param.Field[ExperimentUpdateParamsRepoInfo] `json:"repo_info"`
}

func (r ExperimentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metadata about the state of the repo when the experiment was created
type ExperimentUpdateParamsRepoInfo struct {
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

func (r ExperimentUpdateParamsRepoInfo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExperimentListParams struct {
	// A cursor for pagination. For example, if the initial item in the last page you
	// fetched had an id of `foo`, pass `ending_before=foo` to fetch the previous page.
	// Note: you may only pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Name of the experiment to search for
	ExperimentName param.Field[string] `query:"experiment_name"`
	// Limit the number of objects to return
	Limit param.Field[int64] `query:"limit"`
	// Filter search results to within a particular organization
	OrgName param.Field[string] `query:"org_name"`
	// Name of the project to search for
	ProjectName param.Field[string] `query:"project_name"`
	// A cursor for pagination. For example, if the final item in the last page you
	// fetched had an id of `foo`, pass `starting_after=foo` to fetch the next page.
	// Note: you may only pass one of `starting_after` and `ending_before`
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [ExperimentListParams]'s query parameters as `url.Values`.
func (r ExperimentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExperimentFeedbackParams struct {
	// A list of experiment feedback items
	Feedback param.Field[[]ExperimentFeedbackParamsFeedback] `json:"feedback,required"`
}

func (r ExperimentFeedbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExperimentFeedbackParamsFeedback struct {
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
	Source param.Field[ExperimentFeedbackParamsFeedbackSource] `json:"source"`
}

func (r ExperimentFeedbackParamsFeedback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The source of the feedback. Must be one of "external" (default), "app", or "api"
type ExperimentFeedbackParamsFeedbackSource string

const (
	ExperimentFeedbackParamsFeedbackSourceApp      ExperimentFeedbackParamsFeedbackSource = "app"
	ExperimentFeedbackParamsFeedbackSourceAPI      ExperimentFeedbackParamsFeedbackSource = "api"
	ExperimentFeedbackParamsFeedbackSourceExternal ExperimentFeedbackParamsFeedbackSource = "external"
)

type ExperimentFetchParams struct {
	// Fetch queries may be paginated if the total result size is expected to be large
	// (e.g. project_logs which accumulate over a long time). Note that fetch queries
	// only support pagination in descending time order (from latest to earliest
	// `_xact_id`. Furthermore, later pages may return rows which showed up in earlier
	// pages, except with an earlier `_xact_id`. This happens because pagination occurs
	// over the whole version history of the event log. You will most likely want to
	// exclude any such duplicate, outdated rows (by `id`) from your combined result
	// set.
	//
	// The `limit` parameter controls the number of full traces to return. So you may
	// end up with more individual rows than the specified limit if you are fetching
	// events containing traces.
	Limit param.Field[int64] `query:"limit"`
	// Together, `max_xact_id` and `max_root_span_id` form a cursor for paginating
	// event fetches. Given a previous fetch with a list of rows, you can determine
	// `max_root_span_id` as the maximum of the `root_span_id` field over all rows. See
	// the documentation for `limit` for an overview of paginating fetch queries.
	MaxRootSpanID param.Field[string] `query:"max_root_span_id"`
	// Together, `max_xact_id` and `max_root_span_id` form a cursor for paginating
	// event fetches. Given a previous fetch with a list of rows, you can determine
	// `max_xact_id` as the maximum of the `_xact_id` field over all rows. See the
	// documentation for `limit` for an overview of paginating fetch queries.
	MaxXactID param.Field[int64] `query:"max_xact_id"`
	// You may specify a version id to retrieve a snapshot of the events from a past
	// time. The version id is essentially a filter on the latest event transaction id.
	// You can use the `max_xact_id` returned by a past fetch as the version to
	// reproduce that exact fetch.
	Version param.Field[int64] `query:"version"`
}

// URLQuery serializes [ExperimentFetchParams]'s query parameters as `url.Values`.
func (r ExperimentFetchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExperimentFetchPostParams struct {
	// A list of filters on the events to fetch. Currently, only path-lookup type
	// filters are supported, but we may add more in the future
	Filters param.Field[[]ExperimentFetchPostParamsFilter] `json:"filters"`
	// Fetch queries may be paginated if the total result size is expected to be large
	// (e.g. project_logs which accumulate over a long time). Note that fetch queries
	// only support pagination in descending time order (from latest to earliest
	// `_xact_id`. Furthermore, later pages may return rows which showed up in earlier
	// pages, except with an earlier `_xact_id`. This happens because pagination occurs
	// over the whole version history of the event log. You will most likely want to
	// exclude any such duplicate, outdated rows (by `id`) from your combined result
	// set.
	//
	// The `limit` parameter controls the number of full traces to return. So you may
	// end up with more individual rows than the specified limit if you are fetching
	// events containing traces.
	Limit param.Field[int64] `json:"limit"`
	// Together, `max_xact_id` and `max_root_span_id` form a cursor for paginating
	// event fetches. Given a previous fetch with a list of rows, you can determine
	// `max_root_span_id` as the maximum of the `root_span_id` field over all rows. See
	// the documentation for `limit` for an overview of paginating fetch queries.
	MaxRootSpanID param.Field[string] `json:"max_root_span_id"`
	// Together, `max_xact_id` and `max_root_span_id` form a cursor for paginating
	// event fetches. Given a previous fetch with a list of rows, you can determine
	// `max_xact_id` as the maximum of the `_xact_id` field over all rows. See the
	// documentation for `limit` for an overview of paginating fetch queries.
	MaxXactID param.Field[int64] `json:"max_xact_id"`
	// You may specify a version id to retrieve a snapshot of the events from a past
	// time. The version id is essentially a filter on the latest event transaction id.
	// You can use the `max_xact_id` returned by a past fetch as the version to
	// reproduce that exact fetch.
	Version param.Field[int64] `json:"version"`
}

func (r ExperimentFetchPostParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A path-lookup filter describes an equality comparison against a specific
// sub-field in the event row. For instance, if you wish to filter on the value of
// `c` in `{"input": {"a": {"b": {"c": "hello"}}}}`, pass
// `path=["input", "a", "b", "c"]` and `value="hello"`
type ExperimentFetchPostParamsFilter struct {
	// List of fields describing the path to the value to be checked against. For
	// instance, if you wish to filter on the value of `c` in
	// `{"input": {"a": {"b": {"c": "hello"}}}}`, pass `path=["input", "a", "b", "c"]`
	Path param.Field[[]string] `json:"path,required"`
	// Denotes the type of filter as a path-lookup filter
	Type param.Field[ExperimentFetchPostParamsFiltersType] `json:"type,required"`
	// The value to compare equality-wise against the event value at the specified
	// `path`. The value must be a "primitive", that is, any JSON-serializable object
	// except for objects and arrays. For instance, if you wish to filter on the value
	// of "input.a.b.c" in the object `{"input": {"a": {"b": {"c": "hello"}}}}`, pass
	// `value="hello"`
	Value param.Field[interface{}] `json:"value"`
}

func (r ExperimentFetchPostParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Denotes the type of filter as a path-lookup filter
type ExperimentFetchPostParamsFiltersType string

const (
	ExperimentFetchPostParamsFiltersTypePathLookup ExperimentFetchPostParamsFiltersType = "path_lookup"
)

type ExperimentInsertParams struct {
	// A list of experiment events to insert
	Events param.Field[[]ExperimentInsertParamsEvent] `json:"events,required"`
}

func (r ExperimentInsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [ExperimentInsertParamsEventsInsertExperimentEventReplace],
// [ExperimentInsertParamsEventsInsertExperimentEventMerge].
type ExperimentInsertParamsEvent interface {
	implementsExperimentInsertParamsEvent()
}

type ExperimentInsertParamsEventsInsertExperimentEventReplace struct {
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
	// [guide](https://www.braintrustdata.com/docs/guides/tracing) for full details).
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
	Context param.Field[ExperimentInsertParamsEventsInsertExperimentEventReplaceContext] `json:"context"`
	// If the experiment is associated to a dataset, this is the event-level dataset id
	// this experiment event is tied to
	DatasetRecordID param.Field[string] `json:"dataset_record_id"`
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
	Metrics param.Field[ExperimentInsertParamsEventsInsertExperimentEventReplaceMetrics] `json:"metrics"`
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
	SpanAttributes param.Field[ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributes] `json:"span_attributes"`
}

func (r ExperimentInsertParamsEventsInsertExperimentEventReplace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExperimentInsertParamsEventsInsertExperimentEventReplace) implementsExperimentInsertParamsEvent() {
}

// Context is additional information about the code that produced the experiment
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the
// experiment event
type ExperimentInsertParamsEventsInsertExperimentEventReplaceContext struct {
	// Name of the file in code where the experiment event was created
	CallerFilename param.Field[string] `json:"caller_filename"`
	// The function in code which created the experiment event
	CallerFunctionname param.Field[string] `json:"caller_functionname"`
	// Line of code where the experiment event was created
	CallerLineno param.Field[int64]     `json:"caller_lineno"`
	ExtraFields  map[string]interface{} `json:"-,extras"`
}

func (r ExperimentInsertParamsEventsInsertExperimentEventReplaceContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the experiment event. Use "start" and "end" to track the time span over
// which the experiment event was produced
type ExperimentInsertParamsEventsInsertExperimentEventReplaceMetrics struct {
	// A unix timestamp recording when the section of code which produced the
	// experiment event finished
	End param.Field[float64] `json:"end"`
	// A unix timestamp recording when the section of code which produced the
	// experiment event started
	Start       param.Field[float64]   `json:"start"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ExperimentInsertParamsEventsInsertExperimentEventReplaceMetrics) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Human-identifying attributes of the span, such as name, type, etc.
type ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributes struct {
	// Name of the span, for display purposes only
	Name param.Field[string] `json:"name"`
	// Type of the span, for display purposes only
	Type        param.Field[ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesType] `json:"type"`
	ExtraFields map[string]interface{}                                                                  `json:"-,extras"`
}

func (r ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the span, for display purposes only
type ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesType string

const (
	ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesTypeLlm      ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesType = "llm"
	ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesTypeScore    ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesType = "score"
	ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesTypeFunction ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesType = "function"
	ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesTypeEval     ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesType = "eval"
	ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesTypeTask     ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesType = "task"
	ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesTypeTool     ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesType = "tool"
)

type ExperimentInsertParamsEventsInsertExperimentEventMerge struct {
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
	Context param.Field[ExperimentInsertParamsEventsInsertExperimentEventMergeContext] `json:"context"`
	// If the experiment is associated to a dataset, this is the event-level dataset id
	// this experiment event is tied to
	DatasetRecordID param.Field[string] `json:"dataset_record_id"`
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
	Metrics param.Field[ExperimentInsertParamsEventsInsertExperimentEventMergeMetrics] `json:"metrics"`
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
	SpanAttributes param.Field[ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributes] `json:"span_attributes"`
}

func (r ExperimentInsertParamsEventsInsertExperimentEventMerge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExperimentInsertParamsEventsInsertExperimentEventMerge) implementsExperimentInsertParamsEvent() {
}

// Context is additional information about the code that produced the experiment
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the
// experiment event
type ExperimentInsertParamsEventsInsertExperimentEventMergeContext struct {
	// Name of the file in code where the experiment event was created
	CallerFilename param.Field[string] `json:"caller_filename"`
	// The function in code which created the experiment event
	CallerFunctionname param.Field[string] `json:"caller_functionname"`
	// Line of code where the experiment event was created
	CallerLineno param.Field[int64]     `json:"caller_lineno"`
	ExtraFields  map[string]interface{} `json:"-,extras"`
}

func (r ExperimentInsertParamsEventsInsertExperimentEventMergeContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the experiment event. Use "start" and "end" to track the time span over
// which the experiment event was produced
type ExperimentInsertParamsEventsInsertExperimentEventMergeMetrics struct {
	// A unix timestamp recording when the section of code which produced the
	// experiment event finished
	End param.Field[float64] `json:"end"`
	// A unix timestamp recording when the section of code which produced the
	// experiment event started
	Start       param.Field[float64]   `json:"start"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ExperimentInsertParamsEventsInsertExperimentEventMergeMetrics) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Human-identifying attributes of the span, such as name, type, etc.
type ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributes struct {
	// Name of the span, for display purposes only
	Name param.Field[string] `json:"name"`
	// Type of the span, for display purposes only
	Type        param.Field[ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributesType] `json:"type"`
	ExtraFields map[string]interface{}                                                                `json:"-,extras"`
}

func (r ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the span, for display purposes only
type ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributesType string

const (
	ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributesTypeLlm      ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributesType = "llm"
	ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributesTypeScore    ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributesType = "score"
	ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributesTypeFunction ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributesType = "function"
	ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributesTypeEval     ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributesType = "eval"
	ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributesTypeTask     ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributesType = "task"
	ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributesTypeTool     ExperimentInsertParamsEventsInsertExperimentEventMergeSpanAttributesType = "tool"
)

type ExperimentReplaceParams struct {
	// Unique identifier for the project that the experiment belongs under
	ProjectID param.Field[string] `json:"project_id,required" format:"uuid"`
	// Id of default base experiment to compare against when viewing this experiment
	BaseExpID param.Field[string] `json:"base_exp_id" format:"uuid"`
	// Identifier of the linked dataset, or null if the experiment is not linked to a
	// dataset
	DatasetID param.Field[string] `json:"dataset_id" format:"uuid"`
	// Version number of the linked dataset the experiment was run against. This can be
	// used to reproduce the experiment after the dataset has been modified.
	DatasetVersion param.Field[string] `json:"dataset_version"`
	// Textual description of the experiment
	Description param.Field[string] `json:"description"`
	// User-controlled metadata about the experiment
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// Name of the experiment. Within a project, experiment names are unique
	Name param.Field[string] `json:"name"`
	// Whether or not the experiment is public. Public experiments can be viewed by
	// anybody inside or outside the organization
	Public param.Field[bool] `json:"public"`
	// Metadata about the state of the repo when the experiment was created
	RepoInfo param.Field[ExperimentReplaceParamsRepoInfo] `json:"repo_info"`
}

func (r ExperimentReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metadata about the state of the repo when the experiment was created
type ExperimentReplaceParamsRepoInfo struct {
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

func (r ExperimentReplaceParamsRepoInfo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
