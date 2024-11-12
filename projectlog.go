// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/internal/apiquery"
	"github.com/braintrustdata/braintrust-go/internal/param"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/braintrustdata/braintrust-go/shared"
)

// ProjectLogService contains methods and other services that help with interacting
// with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectLogService] method instead.
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
func (r *ProjectLogService) Feedback(ctx context.Context, projectID string, body ProjectLogFeedbackParams, opts ...option.RequestOption) (res *shared.FeedbackResponseSchema, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/project_logs/%s/feedback", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch the events in a project logs. Equivalent to the POST form of the same
// path, but with the parameters in the URL query rather than in the request body
func (r *ProjectLogService) Fetch(ctx context.Context, projectID string, query ProjectLogFetchParams, opts ...option.RequestOption) (res *shared.FetchProjectLogsEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/project_logs/%s/fetch", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch the events in a project logs. Equivalent to the GET form of the same path,
// but with the parameters in the request body rather than in the URL query
func (r *ProjectLogService) FetchPost(ctx context.Context, projectID string, body ProjectLogFetchPostParams, opts ...option.RequestOption) (res *shared.FetchProjectLogsEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/project_logs/%s/fetch", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Insert a set of events into the project logs
func (r *ProjectLogService) Insert(ctx context.Context, projectID string, body ProjectLogInsertParams, opts ...option.RequestOption) (res *shared.InsertEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/project_logs/%s/insert", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ProjectLogFeedbackParams struct {
	// A list of project logs feedback items
	Feedback param.Field[[]shared.FeedbackProjectLogsItemParam] `json:"feedback,required"`
}

func (r ProjectLogFeedbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectLogFetchParams struct {
	// limit the number of traces fetched
	//
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
	// DEPRECATION NOTICE: The manually-constructed pagination cursor is deprecated in
	// favor of the explicit 'cursor' returned by object fetch requests. Please prefer
	// the 'cursor' argument going forwards.
	//
	// Together, `max_xact_id` and `max_root_span_id` form a pagination cursor
	//
	// Since a paginated fetch query returns results in order from latest to earliest,
	// the cursor for the next page can be found as the row with the minimum (earliest)
	// value of the tuple `(_xact_id, root_span_id)`. See the documentation of `limit`
	// for an overview of paginating fetch queries.
	MaxRootSpanID param.Field[string] `query:"max_root_span_id"`
	// DEPRECATION NOTICE: The manually-constructed pagination cursor is deprecated in
	// favor of the explicit 'cursor' returned by object fetch requests. Please prefer
	// the 'cursor' argument going forwards.
	//
	// Together, `max_xact_id` and `max_root_span_id` form a pagination cursor
	//
	// Since a paginated fetch query returns results in order from latest to earliest,
	// the cursor for the next page can be found as the row with the minimum (earliest)
	// value of the tuple `(_xact_id, root_span_id)`. See the documentation of `limit`
	// for an overview of paginating fetch queries.
	MaxXactID param.Field[string] `query:"max_xact_id"`
	// Retrieve a snapshot of events from a past time
	//
	// The version id is essentially a filter on the latest event transaction id. You
	// can use the `max_xact_id` returned by a past fetch as the version to reproduce
	// that exact fetch.
	Version param.Field[string] `query:"version"`
}

// URLQuery serializes [ProjectLogFetchParams]'s query parameters as `url.Values`.
func (r ProjectLogFetchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectLogFetchPostParams struct {
	// An opaque string to be used as a cursor for the next page of results, in order
	// from latest to earliest.
	//
	// The string can be obtained directly from the `cursor` property of the previous
	// fetch query
	Cursor param.Field[string] `json:"cursor"`
	// NOTE: This parameter is deprecated and will be removed in a future revision.
	// Consider using the `/btql` endpoint
	// (https://www.braintrust.dev/docs/reference/btql) for more advanced filtering.
	//
	// A list of filters on the events to fetch. Currently, only path-lookup type
	// filters are supported.
	Filters param.Field[[]shared.PathLookupFilterParam] `json:"filters"`
	// limit the number of traces fetched
	//
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
	// DEPRECATION NOTICE: The manually-constructed pagination cursor is deprecated in
	// favor of the explicit 'cursor' returned by object fetch requests. Please prefer
	// the 'cursor' argument going forwards.
	//
	// Together, `max_xact_id` and `max_root_span_id` form a pagination cursor
	//
	// Since a paginated fetch query returns results in order from latest to earliest,
	// the cursor for the next page can be found as the row with the minimum (earliest)
	// value of the tuple `(_xact_id, root_span_id)`. See the documentation of `limit`
	// for an overview of paginating fetch queries.
	MaxRootSpanID param.Field[string] `json:"max_root_span_id"`
	// DEPRECATION NOTICE: The manually-constructed pagination cursor is deprecated in
	// favor of the explicit 'cursor' returned by object fetch requests. Please prefer
	// the 'cursor' argument going forwards.
	//
	// Together, `max_xact_id` and `max_root_span_id` form a pagination cursor
	//
	// Since a paginated fetch query returns results in order from latest to earliest,
	// the cursor for the next page can be found as the row with the minimum (earliest)
	// value of the tuple `(_xact_id, root_span_id)`. See the documentation of `limit`
	// for an overview of paginating fetch queries.
	MaxXactID param.Field[string] `json:"max_xact_id"`
	// Retrieve a snapshot of events from a past time
	//
	// The version id is essentially a filter on the latest event transaction id. You
	// can use the `max_xact_id` returned by a past fetch as the version to reproduce
	// that exact fetch.
	Version param.Field[string] `json:"version"`
}

func (r ProjectLogFetchPostParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectLogInsertParams struct {
	// A list of project logs events to insert
	Events param.Field[[]ProjectLogInsertParamsEventUnion] `json:"events,required"`
}

func (r ProjectLogInsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A project logs event
type ProjectLogInsertParamsEvent struct {
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
	IsMerge    param.Field[bool]        `json:"_is_merge"`
	MergePaths param.Field[interface{}] `json:"_merge_paths"`
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
	ParentID param.Field[string]      `json:"_parent_id"`
	Context  param.Field[interface{}] `json:"context"`
	// The timestamp the project logs event was created
	Created        param.Field[time.Time]   `json:"created" format:"date-time"`
	Error          param.Field[interface{}] `json:"error"`
	Expected       param.Field[interface{}] `json:"expected"`
	Input          param.Field[interface{}] `json:"input"`
	Metadata       param.Field[interface{}] `json:"metadata"`
	Metrics        param.Field[interface{}] `json:"metrics"`
	Output         param.Field[interface{}] `json:"output"`
	Scores         param.Field[interface{}] `json:"scores"`
	SpanAttributes param.Field[interface{}] `json:"span_attributes"`
	Tags           param.Field[interface{}] `json:"tags"`
}

func (r ProjectLogInsertParamsEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectLogInsertParamsEvent) ImplementsProjectLogInsertParamsEventUnion() {}

// A project logs event
//
// Satisfied by [shared.InsertProjectLogsEventReplaceParam],
// [shared.InsertProjectLogsEventMergeParam], [ProjectLogInsertParamsEvent].
type ProjectLogInsertParamsEventUnion interface {
	ImplementsProjectLogInsertParamsEventUnion()
}
