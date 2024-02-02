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
	"github.com/braintrustdata/braintrust-go/option"
)

// ProjectLogService contains methods and other services that help with interacting
// with the braintrust API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewProjectLogService] method instead.
type ProjectLogService struct {
	Options []option.RequestOption
}

// NewProjectLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectLogService(opts ...option.RequestOption) (r *ProjectLogService) {
	r = &ProjectLogService{}
	r.Options = opts
	return
}

// Log feedback for a set of project logs events
func (r *ProjectLogService) Feedback(ctx context.Context, projectID string, body ProjectLogFeedbackParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("v1/project_logs/%s/feedback", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Fetch the events in a project logs. Equivalent to the POST form of the same
// path, but with the parameters in the URL query rather than in the request body
func (r *ProjectLogService) Fetch(ctx context.Context, projectID string, query ProjectLogFetchParams, opts ...option.RequestOption) (res *ProjectLogFetchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/project_logs/%s/fetch", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch the events in a project logs. Equivalent to the GET form of the same path,
// but with the parameters in the request body rather than in the URL query
func (r *ProjectLogService) FetchPost(ctx context.Context, projectID string, body ProjectLogFetchPostParams, opts ...option.RequestOption) (res *ProjectLogFetchPostResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/project_logs/%s/fetch", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Insert a set of events into the project logs
func (r *ProjectLogService) Insert(ctx context.Context, projectID string, body ProjectLogInsertParams, opts ...option.RequestOption) (res *ProjectLogInsertResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/project_logs/%s/insert", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ProjectLogFetchResponse struct {
	// A list of fetched events
	Events []ProjectLogFetchResponseEvent `json:"events,required"`
	JSON   projectLogFetchResponseJSON    `json:"-"`
}

// projectLogFetchResponseJSON contains the JSON metadata for the struct
// [ProjectLogFetchResponse]
type projectLogFetchResponseJSON struct {
	Events      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectLogFetchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectLogFetchResponseEvent struct {
	// A unique identifier for the project logs event. If you don't provide one,
	// BrainTrust will generate one for you
	ID string `json:"id,required"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the project logs (see the
	// `version` parameter)
	XactID int64 `json:"_xact_id,required"`
	// A literal 'g' which identifies the log as a project log
	LogID ProjectLogFetchResponseEventsLogID `json:"log_id,required"`
	// Unique id for the organization that the project belongs under
	OrgID string `json:"org_id,required" format:"uuid"`
	// Unique identifier for the project
	ProjectID string `json:"project_id,required" format:"uuid"`
	// The `span_id` of the root of the trace this project logs event belongs to
	RootSpanID string `json:"root_span_id,required"`
	// A unique identifier used to link different project logs events together as part
	// of a full trace. See the
	// [tracing guide](https://www.braintrustdata.com/docs/guides/tracing) for full
	// details on tracing
	SpanID string `json:"span_id,required"`
	// Context is additional information about the code that produced the project logs
	// event. It is essentially the textual counterpart to `metrics`. Use the
	// `caller_*` attributes to track the location in code which produced the project
	// logs event
	Context ProjectLogFetchResponseEventsContext `json:"context,nullable"`
	// The timestamp the project logs event was created
	Created time.Time `json:"created,nullable" format:"date-time"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not.
	// Braintrust currently does not compare `output` to `expected` for you, since
	// there are so many different ways to do that correctly. Instead, these values are
	// just used to help you navigate while digging into analyses. However, we may
	// later use these values to re-score outputs or fine-tune your models.
	Expected interface{} `json:"expected"`
	// The arguments that uniquely define a user input(an arbitrary, JSON serializable
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
	Metrics ProjectLogFetchResponseEventsMetrics `json:"metrics,nullable"`
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
	SpanAttributes ProjectLogFetchResponseEventsSpanAttributes `json:"span_attributes,nullable"`
	// An array of the parent `span_ids` of this project logs event. This should be
	// empty for the root span of a trace, and should most often contain just one
	// parent element for subspans
	SpanParents []string                         `json:"span_parents,nullable"`
	JSON        projectLogFetchResponseEventJSON `json:"-"`
}

// projectLogFetchResponseEventJSON contains the JSON metadata for the struct
// [ProjectLogFetchResponseEvent]
type projectLogFetchResponseEventJSON struct {
	ID             apijson.Field
	XactID         apijson.Field
	LogID          apijson.Field
	OrgID          apijson.Field
	ProjectID      apijson.Field
	RootSpanID     apijson.Field
	SpanID         apijson.Field
	Context        apijson.Field
	Created        apijson.Field
	Expected       apijson.Field
	Input          apijson.Field
	Metadata       apijson.Field
	Metrics        apijson.Field
	Output         apijson.Field
	Scores         apijson.Field
	SpanAttributes apijson.Field
	SpanParents    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectLogFetchResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A literal 'g' which identifies the log as a project log
type ProjectLogFetchResponseEventsLogID string

const (
	ProjectLogFetchResponseEventsLogIDG ProjectLogFetchResponseEventsLogID = "g"
)

// Context is additional information about the code that produced the project logs
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the project
// logs event
type ProjectLogFetchResponseEventsContext struct {
	// Name of the file in code where the project logs event was created
	CallerFilename string `json:"caller_filename,nullable"`
	// The function in code which created the project logs event
	CallerFunctionname string `json:"caller_functionname,nullable"`
	// Line of code where the project logs event was created
	CallerLineno int64                                    `json:"caller_lineno,nullable"`
	ExtraFields  map[string]interface{}                   `json:"-,extras"`
	JSON         projectLogFetchResponseEventsContextJSON `json:"-"`
}

// projectLogFetchResponseEventsContextJSON contains the JSON metadata for the
// struct [ProjectLogFetchResponseEventsContext]
type projectLogFetchResponseEventsContextJSON struct {
	CallerFilename     apijson.Field
	CallerFunctionname apijson.Field
	CallerLineno       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProjectLogFetchResponseEventsContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the project logs event. Use "start" and "end" to track the time span
// over which the project logs event was produced
type ProjectLogFetchResponseEventsMetrics struct {
	// A unix timestamp recording when the section of code which produced the project
	// logs event finished
	End float64 `json:"end,nullable"`
	// A unix timestamp recording when the section of code which produced the project
	// logs event started
	Start       float64                                  `json:"start,nullable"`
	ExtraFields map[string]interface{}                   `json:"-,extras"`
	JSON        projectLogFetchResponseEventsMetricsJSON `json:"-"`
}

// projectLogFetchResponseEventsMetricsJSON contains the JSON metadata for the
// struct [ProjectLogFetchResponseEventsMetrics]
type projectLogFetchResponseEventsMetricsJSON struct {
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectLogFetchResponseEventsMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Human-identifying attributes of the span, such as name, type, etc.
type ProjectLogFetchResponseEventsSpanAttributes struct {
	// Name of the span, for display purposes only
	Name string `json:"name,nullable"`
	// Type of the span, for display purposes only
	Type        ProjectLogFetchResponseEventsSpanAttributesType `json:"type,nullable"`
	ExtraFields map[string]interface{}                          `json:"-,extras"`
	JSON        projectLogFetchResponseEventsSpanAttributesJSON `json:"-"`
}

// projectLogFetchResponseEventsSpanAttributesJSON contains the JSON metadata for
// the struct [ProjectLogFetchResponseEventsSpanAttributes]
type projectLogFetchResponseEventsSpanAttributesJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectLogFetchResponseEventsSpanAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the span, for display purposes only
type ProjectLogFetchResponseEventsSpanAttributesType string

const (
	ProjectLogFetchResponseEventsSpanAttributesTypeLlm      ProjectLogFetchResponseEventsSpanAttributesType = "llm"
	ProjectLogFetchResponseEventsSpanAttributesTypeScore    ProjectLogFetchResponseEventsSpanAttributesType = "score"
	ProjectLogFetchResponseEventsSpanAttributesTypeFunction ProjectLogFetchResponseEventsSpanAttributesType = "function"
	ProjectLogFetchResponseEventsSpanAttributesTypeEval     ProjectLogFetchResponseEventsSpanAttributesType = "eval"
	ProjectLogFetchResponseEventsSpanAttributesTypeTask     ProjectLogFetchResponseEventsSpanAttributesType = "task"
	ProjectLogFetchResponseEventsSpanAttributesTypeTool     ProjectLogFetchResponseEventsSpanAttributesType = "tool"
)

type ProjectLogFetchPostResponse struct {
	// A list of fetched events
	Events []ProjectLogFetchPostResponseEvent `json:"events,required"`
	JSON   projectLogFetchPostResponseJSON    `json:"-"`
}

// projectLogFetchPostResponseJSON contains the JSON metadata for the struct
// [ProjectLogFetchPostResponse]
type projectLogFetchPostResponseJSON struct {
	Events      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectLogFetchPostResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectLogFetchPostResponseEvent struct {
	// A unique identifier for the project logs event. If you don't provide one,
	// BrainTrust will generate one for you
	ID string `json:"id,required"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the project logs (see the
	// `version` parameter)
	XactID int64 `json:"_xact_id,required"`
	// A literal 'g' which identifies the log as a project log
	LogID ProjectLogFetchPostResponseEventsLogID `json:"log_id,required"`
	// Unique id for the organization that the project belongs under
	OrgID string `json:"org_id,required" format:"uuid"`
	// Unique identifier for the project
	ProjectID string `json:"project_id,required" format:"uuid"`
	// The `span_id` of the root of the trace this project logs event belongs to
	RootSpanID string `json:"root_span_id,required"`
	// A unique identifier used to link different project logs events together as part
	// of a full trace. See the
	// [tracing guide](https://www.braintrustdata.com/docs/guides/tracing) for full
	// details on tracing
	SpanID string `json:"span_id,required"`
	// Context is additional information about the code that produced the project logs
	// event. It is essentially the textual counterpart to `metrics`. Use the
	// `caller_*` attributes to track the location in code which produced the project
	// logs event
	Context ProjectLogFetchPostResponseEventsContext `json:"context,nullable"`
	// The timestamp the project logs event was created
	Created time.Time `json:"created,nullable" format:"date-time"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not.
	// Braintrust currently does not compare `output` to `expected` for you, since
	// there are so many different ways to do that correctly. Instead, these values are
	// just used to help you navigate while digging into analyses. However, we may
	// later use these values to re-score outputs or fine-tune your models.
	Expected interface{} `json:"expected"`
	// The arguments that uniquely define a user input(an arbitrary, JSON serializable
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
	Metrics ProjectLogFetchPostResponseEventsMetrics `json:"metrics,nullable"`
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
	SpanAttributes ProjectLogFetchPostResponseEventsSpanAttributes `json:"span_attributes,nullable"`
	// An array of the parent `span_ids` of this project logs event. This should be
	// empty for the root span of a trace, and should most often contain just one
	// parent element for subspans
	SpanParents []string                             `json:"span_parents,nullable"`
	JSON        projectLogFetchPostResponseEventJSON `json:"-"`
}

// projectLogFetchPostResponseEventJSON contains the JSON metadata for the struct
// [ProjectLogFetchPostResponseEvent]
type projectLogFetchPostResponseEventJSON struct {
	ID             apijson.Field
	XactID         apijson.Field
	LogID          apijson.Field
	OrgID          apijson.Field
	ProjectID      apijson.Field
	RootSpanID     apijson.Field
	SpanID         apijson.Field
	Context        apijson.Field
	Created        apijson.Field
	Expected       apijson.Field
	Input          apijson.Field
	Metadata       apijson.Field
	Metrics        apijson.Field
	Output         apijson.Field
	Scores         apijson.Field
	SpanAttributes apijson.Field
	SpanParents    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProjectLogFetchPostResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A literal 'g' which identifies the log as a project log
type ProjectLogFetchPostResponseEventsLogID string

const (
	ProjectLogFetchPostResponseEventsLogIDG ProjectLogFetchPostResponseEventsLogID = "g"
)

// Context is additional information about the code that produced the project logs
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the project
// logs event
type ProjectLogFetchPostResponseEventsContext struct {
	// Name of the file in code where the project logs event was created
	CallerFilename string `json:"caller_filename,nullable"`
	// The function in code which created the project logs event
	CallerFunctionname string `json:"caller_functionname,nullable"`
	// Line of code where the project logs event was created
	CallerLineno int64                                        `json:"caller_lineno,nullable"`
	ExtraFields  map[string]interface{}                       `json:"-,extras"`
	JSON         projectLogFetchPostResponseEventsContextJSON `json:"-"`
}

// projectLogFetchPostResponseEventsContextJSON contains the JSON metadata for the
// struct [ProjectLogFetchPostResponseEventsContext]
type projectLogFetchPostResponseEventsContextJSON struct {
	CallerFilename     apijson.Field
	CallerFunctionname apijson.Field
	CallerLineno       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProjectLogFetchPostResponseEventsContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the project logs event. Use "start" and "end" to track the time span
// over which the project logs event was produced
type ProjectLogFetchPostResponseEventsMetrics struct {
	// A unix timestamp recording when the section of code which produced the project
	// logs event finished
	End float64 `json:"end,nullable"`
	// A unix timestamp recording when the section of code which produced the project
	// logs event started
	Start       float64                                      `json:"start,nullable"`
	ExtraFields map[string]interface{}                       `json:"-,extras"`
	JSON        projectLogFetchPostResponseEventsMetricsJSON `json:"-"`
}

// projectLogFetchPostResponseEventsMetricsJSON contains the JSON metadata for the
// struct [ProjectLogFetchPostResponseEventsMetrics]
type projectLogFetchPostResponseEventsMetricsJSON struct {
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectLogFetchPostResponseEventsMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Human-identifying attributes of the span, such as name, type, etc.
type ProjectLogFetchPostResponseEventsSpanAttributes struct {
	// Name of the span, for display purposes only
	Name string `json:"name,nullable"`
	// Type of the span, for display purposes only
	Type        ProjectLogFetchPostResponseEventsSpanAttributesType `json:"type,nullable"`
	ExtraFields map[string]interface{}                              `json:"-,extras"`
	JSON        projectLogFetchPostResponseEventsSpanAttributesJSON `json:"-"`
}

// projectLogFetchPostResponseEventsSpanAttributesJSON contains the JSON metadata
// for the struct [ProjectLogFetchPostResponseEventsSpanAttributes]
type projectLogFetchPostResponseEventsSpanAttributesJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectLogFetchPostResponseEventsSpanAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the span, for display purposes only
type ProjectLogFetchPostResponseEventsSpanAttributesType string

const (
	ProjectLogFetchPostResponseEventsSpanAttributesTypeLlm      ProjectLogFetchPostResponseEventsSpanAttributesType = "llm"
	ProjectLogFetchPostResponseEventsSpanAttributesTypeScore    ProjectLogFetchPostResponseEventsSpanAttributesType = "score"
	ProjectLogFetchPostResponseEventsSpanAttributesTypeFunction ProjectLogFetchPostResponseEventsSpanAttributesType = "function"
	ProjectLogFetchPostResponseEventsSpanAttributesTypeEval     ProjectLogFetchPostResponseEventsSpanAttributesType = "eval"
	ProjectLogFetchPostResponseEventsSpanAttributesTypeTask     ProjectLogFetchPostResponseEventsSpanAttributesType = "task"
	ProjectLogFetchPostResponseEventsSpanAttributesTypeTool     ProjectLogFetchPostResponseEventsSpanAttributesType = "tool"
)

type ProjectLogInsertResponse struct {
	// The ids of all rows that were inserted, aligning one-to-one with the rows
	// provided as input
	RowIDs []string                     `json:"row_ids,required"`
	JSON   projectLogInsertResponseJSON `json:"-"`
}

// projectLogInsertResponseJSON contains the JSON metadata for the struct
// [ProjectLogInsertResponse]
type projectLogInsertResponseJSON struct {
	RowIDs      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectLogInsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectLogFeedbackParams struct {
	// A list of project logs feedback items
	Feedback param.Field[[]ProjectLogFeedbackParamsFeedback] `json:"feedback,required"`
}

func (r ProjectLogFeedbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectLogFeedbackParamsFeedback struct {
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
	Source param.Field[ProjectLogFeedbackParamsFeedbackSource] `json:"source"`
}

func (r ProjectLogFeedbackParamsFeedback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The source of the feedback. Must be one of "external" (default), "app", or "api"
type ProjectLogFeedbackParamsFeedbackSource string

const (
	ProjectLogFeedbackParamsFeedbackSourceApp      ProjectLogFeedbackParamsFeedbackSource = "app"
	ProjectLogFeedbackParamsFeedbackSourceAPI      ProjectLogFeedbackParamsFeedbackSource = "api"
	ProjectLogFeedbackParamsFeedbackSourceExternal ProjectLogFeedbackParamsFeedbackSource = "external"
)

type ProjectLogFetchParams struct {
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

// URLQuery serializes [ProjectLogFetchParams]'s query parameters as `url.Values`.
func (r ProjectLogFetchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectLogFetchPostParams struct {
	// A list of filters on the events to fetch. Currently, only path-lookup type
	// filters are supported, but we may add more in the future
	Filters param.Field[[]ProjectLogFetchPostParamsFilter] `json:"filters"`
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

func (r ProjectLogFetchPostParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A path-lookup filter describes an equality comparison against a specific
// sub-field in the event row. For instance, if you wish to filter on the value of
// `c` in `{"input": {"a": {"b": {"c": "hello"}}}}`, pass
// `path=["input", "a", "b", "c"]` and `value="hello"`
type ProjectLogFetchPostParamsFilter struct {
	// List of fields describing the path to the value to be checked against. For
	// instance, if you wish to filter on the value of `c` in
	// `{"input": {"a": {"b": {"c": "hello"}}}}`, pass `path=["input", "a", "b", "c"]`
	Path param.Field[[]string] `json:"path,required"`
	// Denotes the type of filter as a path-lookup filter
	Type param.Field[ProjectLogFetchPostParamsFiltersType] `json:"type,required"`
	// The value to compare equality-wise against the event value at the specified
	// `path`. The value must be a "primitive", that is, any JSON-serializable object
	// except for objects and arrays. For instance, if you wish to filter on the value
	// of "input.a.b.c" in the object `{"input": {"a": {"b": {"c": "hello"}}}}`, pass
	// `value="hello"`
	Value param.Field[interface{}] `json:"value"`
}

func (r ProjectLogFetchPostParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Denotes the type of filter as a path-lookup filter
type ProjectLogFetchPostParamsFiltersType string

const (
	ProjectLogFetchPostParamsFiltersTypePathLookup ProjectLogFetchPostParamsFiltersType = "path_lookup"
)

type ProjectLogInsertParams struct {
	// A list of project logs events to insert
	Events param.Field[[]ProjectLogInsertParamsEvent] `json:"events,required"`
}

func (r ProjectLogInsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [ProjectLogInsertParamsEventsInsertProjectLogsEventReplace],
// [ProjectLogInsertParamsEventsInsertProjectLogsEventMerge].
type ProjectLogInsertParamsEvent interface {
	implementsProjectLogInsertParamsEvent()
}

type ProjectLogInsertParamsEventsInsertProjectLogsEventReplace struct {
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
	IsMerge param.Field[ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceIsMerge] `json:"_is_merge"`
	// Pass `_object_delete=true` to mark the project logs event deleted. Deleted
	// events will not show up in subsequent fetches for this project logs
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
	// Context is additional information about the code that produced the project logs
	// event. It is essentially the textual counterpart to `metrics`. Use the
	// `caller_*` attributes to track the location in code which produced the project
	// logs event
	Context param.Field[ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceContext] `json:"context"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not.
	// Braintrust currently does not compare `output` to `expected` for you, since
	// there are so many different ways to do that correctly. Instead, these values are
	// just used to help you navigate while digging into analyses. However, we may
	// later use these values to re-score outputs or fine-tune your models.
	Expected param.Field[interface{}] `json:"expected"`
	// The arguments that uniquely define a user input(an arbitrary, JSON serializable
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
	Metrics param.Field[ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceMetrics] `json:"metrics"`
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
	SpanAttributes param.Field[ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributes] `json:"span_attributes"`
}

func (r ProjectLogInsertParamsEventsInsertProjectLogsEventReplace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectLogInsertParamsEventsInsertProjectLogsEventReplace) implementsProjectLogInsertParamsEvent() {
}

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
//
// Satisfied by
// [ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceIsMergeBoolean],
// [ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceIsMergeUnknown].
type ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceIsMerge interface {
	ImplementsProjectLogInsertParamsEventsInsertProjectLogsEventReplaceIsMerge()
}

type ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceIsMergeBoolean bool

const (
	ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceIsMergeBooleanFalse ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceIsMergeBoolean = false
)

// Context is additional information about the code that produced the project logs
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the project
// logs event
type ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceContext struct {
	// Name of the file in code where the project logs event was created
	CallerFilename param.Field[string] `json:"caller_filename"`
	// The function in code which created the project logs event
	CallerFunctionname param.Field[string] `json:"caller_functionname"`
	// Line of code where the project logs event was created
	CallerLineno param.Field[int64]     `json:"caller_lineno"`
	ExtraFields  map[string]interface{} `json:"-,extras"`
}

func (r ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the project logs event. Use "start" and "end" to track the time span
// over which the project logs event was produced
type ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceMetrics struct {
	// A unix timestamp recording when the section of code which produced the project
	// logs event finished
	End param.Field[float64] `json:"end"`
	// A unix timestamp recording when the section of code which produced the project
	// logs event started
	Start       param.Field[float64]   `json:"start"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceMetrics) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Human-identifying attributes of the span, such as name, type, etc.
type ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributes struct {
	// Name of the span, for display purposes only
	Name param.Field[string] `json:"name"`
	// Type of the span, for display purposes only
	Type        param.Field[ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesType] `json:"type"`
	ExtraFields map[string]interface{}                                                                   `json:"-,extras"`
}

func (r ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the span, for display purposes only
type ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesType string

const (
	ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesTypeLlm      ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesType = "llm"
	ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesTypeScore    ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesType = "score"
	ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesTypeFunction ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesType = "function"
	ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesTypeEval     ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesType = "eval"
	ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesTypeTask     ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesType = "task"
	ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesTypeTool     ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesType = "tool"
)

type ProjectLogInsertParamsEventsInsertProjectLogsEventMerge struct {
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
	IsMerge param.Field[ProjectLogInsertParamsEventsInsertProjectLogsEventMergeIsMerge] `json:"_is_merge,required"`
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
	Context param.Field[ProjectLogInsertParamsEventsInsertProjectLogsEventMergeContext] `json:"context"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not.
	// Braintrust currently does not compare `output` to `expected` for you, since
	// there are so many different ways to do that correctly. Instead, these values are
	// just used to help you navigate while digging into analyses. However, we may
	// later use these values to re-score outputs or fine-tune your models.
	Expected param.Field[interface{}] `json:"expected"`
	// The arguments that uniquely define a user input(an arbitrary, JSON serializable
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
	Metrics param.Field[ProjectLogInsertParamsEventsInsertProjectLogsEventMergeMetrics] `json:"metrics"`
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
	SpanAttributes param.Field[ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributes] `json:"span_attributes"`
}

func (r ProjectLogInsertParamsEventsInsertProjectLogsEventMerge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectLogInsertParamsEventsInsertProjectLogsEventMerge) implementsProjectLogInsertParamsEvent() {
}

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
type ProjectLogInsertParamsEventsInsertProjectLogsEventMergeIsMerge bool

const (
	ProjectLogInsertParamsEventsInsertProjectLogsEventMergeIsMergeTrue ProjectLogInsertParamsEventsInsertProjectLogsEventMergeIsMerge = true
)

// Context is additional information about the code that produced the project logs
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the project
// logs event
type ProjectLogInsertParamsEventsInsertProjectLogsEventMergeContext struct {
	// Name of the file in code where the project logs event was created
	CallerFilename param.Field[string] `json:"caller_filename"`
	// The function in code which created the project logs event
	CallerFunctionname param.Field[string] `json:"caller_functionname"`
	// Line of code where the project logs event was created
	CallerLineno param.Field[int64]     `json:"caller_lineno"`
	ExtraFields  map[string]interface{} `json:"-,extras"`
}

func (r ProjectLogInsertParamsEventsInsertProjectLogsEventMergeContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the project logs event. Use "start" and "end" to track the time span
// over which the project logs event was produced
type ProjectLogInsertParamsEventsInsertProjectLogsEventMergeMetrics struct {
	// A unix timestamp recording when the section of code which produced the project
	// logs event finished
	End param.Field[float64] `json:"end"`
	// A unix timestamp recording when the section of code which produced the project
	// logs event started
	Start       param.Field[float64]   `json:"start"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ProjectLogInsertParamsEventsInsertProjectLogsEventMergeMetrics) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Human-identifying attributes of the span, such as name, type, etc.
type ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributes struct {
	// Name of the span, for display purposes only
	Name param.Field[string] `json:"name"`
	// Type of the span, for display purposes only
	Type        param.Field[ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributesType] `json:"type"`
	ExtraFields map[string]interface{}                                                                 `json:"-,extras"`
}

func (r ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the span, for display purposes only
type ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributesType string

const (
	ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributesTypeLlm      ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributesType = "llm"
	ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributesTypeScore    ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributesType = "score"
	ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributesTypeFunction ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributesType = "function"
	ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributesTypeEval     ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributesType = "eval"
	ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributesTypeTask     ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributesType = "task"
	ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributesTypeTool     ProjectLogInsertParamsEventsInsertProjectLogsEventMergeSpanAttributesType = "tool"
)
