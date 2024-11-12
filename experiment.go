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
	"github.com/braintrustdata/braintrust-go/packages/pagination"
	"github.com/braintrustdata/braintrust-go/shared"
)

// ExperimentService contains methods and other services that help with interacting
// with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentService] method instead.
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
// the same name as the one specified in the request, will return the existing
// experiment unmodified
func (r *ExperimentService) New(ctx context.Context, body ExperimentNewParams, opts ...option.RequestOption) (res *shared.Experiment, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/experiment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an experiment object by its id
func (r *ExperimentService) Get(ctx context.Context, experimentID string, opts ...option.RequestOption) (res *shared.Experiment, err error) {
	opts = append(r.Options[:], opts...)
	if experimentID == "" {
		err = errors.New("missing required experiment_id parameter")
		return
	}
	path := fmt.Sprintf("v1/experiment/%s", experimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update an experiment object. Specify the fields to update in the
// payload. Any object-type fields will be deep-merged with existing content.
// Currently we do not support removing fields or setting them to null.
func (r *ExperimentService) Update(ctx context.Context, experimentID string, body ExperimentUpdateParams, opts ...option.RequestOption) (res *shared.Experiment, err error) {
	opts = append(r.Options[:], opts...)
	if experimentID == "" {
		err = errors.New("missing required experiment_id parameter")
		return
	}
	path := fmt.Sprintf("v1/experiment/%s", experimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all experiments. The experiments are sorted by creation date, with the
// most recently-created experiments coming first
func (r *ExperimentService) List(ctx context.Context, query ExperimentListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.Experiment], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *ExperimentService) ListAutoPaging(ctx context.Context, query ExperimentListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.Experiment] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete an experiment object by its id
func (r *ExperimentService) Delete(ctx context.Context, experimentID string, opts ...option.RequestOption) (res *shared.Experiment, err error) {
	opts = append(r.Options[:], opts...)
	if experimentID == "" {
		err = errors.New("missing required experiment_id parameter")
		return
	}
	path := fmt.Sprintf("v1/experiment/%s", experimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Log feedback for a set of experiment events
func (r *ExperimentService) Feedback(ctx context.Context, experimentID string, body ExperimentFeedbackParams, opts ...option.RequestOption) (res *shared.FeedbackResponseSchema, err error) {
	opts = append(r.Options[:], opts...)
	if experimentID == "" {
		err = errors.New("missing required experiment_id parameter")
		return
	}
	path := fmt.Sprintf("v1/experiment/%s/feedback", experimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch the events in an experiment. Equivalent to the POST form of the same path,
// but with the parameters in the URL query rather than in the request body
func (r *ExperimentService) Fetch(ctx context.Context, experimentID string, query ExperimentFetchParams, opts ...option.RequestOption) (res *shared.FetchExperimentEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if experimentID == "" {
		err = errors.New("missing required experiment_id parameter")
		return
	}
	path := fmt.Sprintf("v1/experiment/%s/fetch", experimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch the events in an experiment. Equivalent to the GET form of the same path,
// but with the parameters in the request body rather than in the URL query
func (r *ExperimentService) FetchPost(ctx context.Context, experimentID string, body ExperimentFetchPostParams, opts ...option.RequestOption) (res *shared.FetchExperimentEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if experimentID == "" {
		err = errors.New("missing required experiment_id parameter")
		return
	}
	path := fmt.Sprintf("v1/experiment/%s/fetch", experimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Insert a set of events into the experiment
func (r *ExperimentService) Insert(ctx context.Context, experimentID string, body ExperimentInsertParams, opts ...option.RequestOption) (res *shared.InsertEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if experimentID == "" {
		err = errors.New("missing required experiment_id parameter")
		return
	}
	path := fmt.Sprintf("v1/experiment/%s/insert", experimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Summarize experiment
func (r *ExperimentService) Summarize(ctx context.Context, experimentID string, query ExperimentSummarizeParams, opts ...option.RequestOption) (res *shared.SummarizeExperimentResponse, err error) {
	opts = append(r.Options[:], opts...)
	if experimentID == "" {
		err = errors.New("missing required experiment_id parameter")
		return
	}
	path := fmt.Sprintf("v1/experiment/%s/summarize", experimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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
	// Normally, creating an experiment with the same name as an existing experiment
	// will return the existing one un-modified. But if `ensure_new` is true,
	// registration will generate a new experiment with a unique name in case of a
	// conflict.
	EnsureNew param.Field[bool] `json:"ensure_new"`
	// User-controlled metadata about the experiment
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// Name of the experiment. Within a project, experiment names are unique
	Name param.Field[string] `json:"name"`
	// Whether or not the experiment is public. Public experiments can be viewed by
	// anybody inside or outside the organization
	Public param.Field[bool] `json:"public"`
	// Metadata about the state of the repo when the experiment was created
	RepoInfo param.Field[shared.RepoInfoParam] `json:"repo_info"`
}

func (r ExperimentNewParams) MarshalJSON() (data []byte, err error) {
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
	RepoInfo param.Field[shared.RepoInfoParam] `json:"repo_info"`
}

func (r ExperimentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExperimentListParams struct {
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Name of the experiment to search for
	ExperimentName param.Field[string] `query:"experiment_name"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[ExperimentListParamsIDsUnion] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[int64] `query:"limit"`
	// Filter search results to within a particular organization
	OrgName param.Field[string] `query:"org_name"`
	// Project id
	ProjectID param.Field[string] `query:"project_id" format:"uuid"`
	// Name of the project to search for
	ProjectName param.Field[string] `query:"project_name"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [ExperimentListParams]'s query parameters as `url.Values`.
func (r ExperimentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [ExperimentListParamsIDsArray].
type ExperimentListParamsIDsUnion interface {
	ImplementsExperimentListParamsIDsUnion()
}

type ExperimentListParamsIDsArray []string

func (r ExperimentListParamsIDsArray) ImplementsExperimentListParamsIDsUnion() {}

type ExperimentFeedbackParams struct {
	// A list of experiment feedback items
	Feedback param.Field[[]shared.FeedbackExperimentItemParam] `json:"feedback,required"`
}

func (r ExperimentFeedbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExperimentFetchParams struct {
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

// URLQuery serializes [ExperimentFetchParams]'s query parameters as `url.Values`.
func (r ExperimentFetchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExperimentFetchPostParams struct {
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

func (r ExperimentFetchPostParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExperimentInsertParams struct {
	// A list of experiment events to insert
	Events param.Field[[]ExperimentInsertParamsEventUnion] `json:"events,required"`
}

func (r ExperimentInsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An experiment event
type ExperimentInsertParamsEvent struct {
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
	IsMerge    param.Field[bool]        `json:"_is_merge"`
	MergePaths param.Field[interface{}] `json:"_merge_paths"`
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
	ParentID param.Field[string]      `json:"_parent_id"`
	Context  param.Field[interface{}] `json:"context"`
	// The timestamp the experiment event was created
	Created param.Field[time.Time] `json:"created" format:"date-time"`
	// If the experiment is associated to a dataset, this is the event-level dataset id
	// this experiment event is tied to
	DatasetRecordID param.Field[string]      `json:"dataset_record_id"`
	Error           param.Field[interface{}] `json:"error"`
	Expected        param.Field[interface{}] `json:"expected"`
	Input           param.Field[interface{}] `json:"input"`
	Metadata        param.Field[interface{}] `json:"metadata"`
	Metrics         param.Field[interface{}] `json:"metrics"`
	Output          param.Field[interface{}] `json:"output"`
	Scores          param.Field[interface{}] `json:"scores"`
	SpanAttributes  param.Field[interface{}] `json:"span_attributes"`
	Tags            param.Field[interface{}] `json:"tags"`
}

func (r ExperimentInsertParamsEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExperimentInsertParamsEvent) ImplementsExperimentInsertParamsEventUnion() {}

// An experiment event
//
// Satisfied by [shared.InsertExperimentEventReplaceParam],
// [shared.InsertExperimentEventMergeParam], [ExperimentInsertParamsEvent].
type ExperimentInsertParamsEventUnion interface {
	ImplementsExperimentInsertParamsEventUnion()
}

type ExperimentSummarizeParams struct {
	// The experiment to compare against, if summarizing scores and metrics. If
	// omitted, will fall back to the `base_exp_id` stored in the experiment metadata,
	// and then to the most recent experiment run in the same project. Must pass
	// `summarize_scores=true` for this id to be used
	ComparisonExperimentID param.Field[string] `query:"comparison_experiment_id" format:"uuid"`
	// Whether to summarize the scores and metrics. If false (or omitted), only the
	// metadata will be returned.
	SummarizeScores param.Field[bool] `query:"summarize_scores"`
}

// URLQuery serializes [ExperimentSummarizeParams]'s query parameters as
// `url.Values`.
func (r ExperimentSummarizeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
