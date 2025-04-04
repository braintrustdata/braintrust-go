// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/braintrustdata/braintrust-go/internal/apiquery"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/braintrustdata/braintrust-go/packages/pagination"
	"github.com/braintrustdata/braintrust-go/packages/param"
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
func NewExperimentService(opts ...option.RequestOption) (r ExperimentService) {
	r = ExperimentService{}
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
// but with the parameters in the URL query rather than in the request body. For
// more complex queries, use the `POST /btql` endpoint.
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
// but with the parameters in the request body rather than in the URL query. For
// more complex queries, use the `POST /btql` endpoint.
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
	ProjectID string `json:"project_id,required" format:"uuid"`
	// Id of default base experiment to compare against when viewing this experiment
	BaseExpID param.Opt[string] `json:"base_exp_id,omitzero" format:"uuid"`
	// Identifier of the linked dataset, or null if the experiment is not linked to a
	// dataset
	DatasetID param.Opt[string] `json:"dataset_id,omitzero" format:"uuid"`
	// Version number of the linked dataset the experiment was run against. This can be
	// used to reproduce the experiment after the dataset has been modified.
	DatasetVersion param.Opt[string] `json:"dataset_version,omitzero"`
	// Textual description of the experiment
	Description param.Opt[string] `json:"description,omitzero"`
	// Normally, creating an experiment with the same name as an existing experiment
	// will return the existing one un-modified. But if `ensure_new` is true,
	// registration will generate a new experiment with a unique name in case of a
	// conflict.
	EnsureNew param.Opt[bool] `json:"ensure_new,omitzero"`
	// Name of the experiment. Within a project, experiment names are unique
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether or not the experiment is public. Public experiments can be viewed by
	// anybody inside or outside the organization
	Public param.Opt[bool] `json:"public,omitzero"`
	// User-controlled metadata about the experiment
	Metadata map[string]interface{} `json:"metadata,omitzero"`
	// Metadata about the state of the repo when the experiment was created
	RepoInfo shared.RepoInfoParam `json:"repo_info,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ExperimentNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ExperimentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type ExperimentUpdateParams struct {
	// Id of default base experiment to compare against when viewing this experiment
	BaseExpID param.Opt[string] `json:"base_exp_id,omitzero" format:"uuid"`
	// Identifier of the linked dataset, or null if the experiment is not linked to a
	// dataset
	DatasetID param.Opt[string] `json:"dataset_id,omitzero" format:"uuid"`
	// Version number of the linked dataset the experiment was run against. This can be
	// used to reproduce the experiment after the dataset has been modified.
	DatasetVersion param.Opt[string] `json:"dataset_version,omitzero"`
	// Textual description of the experiment
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the experiment. Within a project, experiment names are unique
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether or not the experiment is public. Public experiments can be viewed by
	// anybody inside or outside the organization
	Public param.Opt[bool] `json:"public,omitzero"`
	// User-controlled metadata about the experiment
	Metadata map[string]interface{} `json:"metadata,omitzero"`
	// Metadata about the state of the repo when the experiment was created
	RepoInfo shared.RepoInfoParam `json:"repo_info,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ExperimentUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ExperimentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type ExperimentListParams struct {
	// Limit the number of objects to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Opt[string] `query:"ending_before,omitzero" format:"uuid" json:"-"`
	// Name of the experiment to search for
	ExperimentName param.Opt[string] `query:"experiment_name,omitzero" json:"-"`
	// Filter search results to within a particular organization
	OrgName param.Opt[string] `query:"org_name,omitzero" json:"-"`
	// Project id
	ProjectID param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Name of the project to search for
	ProjectName param.Opt[string] `query:"project_name,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Opt[string] `query:"starting_after,omitzero" format:"uuid" json:"-"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs ExperimentListParamsIDsUnion `query:"ids,omitzero" format:"uuid" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ExperimentListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ExperimentListParams]'s query parameters as `url.Values`.
func (r ExperimentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExperimentListParamsIDsUnion struct {
	OfString                  param.Opt[string] `query:",omitzero,inline"`
	OfExperimentListsIDsArray []string          `query:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u ExperimentListParamsIDsUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }

func (u *ExperimentListParamsIDsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExperimentListsIDsArray) {
		return &u.OfExperimentListsIDsArray
	}
	return nil
}

type ExperimentFeedbackParams struct {
	// A list of experiment feedback items
	Feedback []shared.FeedbackExperimentItemParam `json:"feedback,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ExperimentFeedbackParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ExperimentFeedbackParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentFeedbackParams
	return param.MarshalObject(r, (*shadow)(&r))
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
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
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
	MaxRootSpanID param.Opt[string] `query:"max_root_span_id,omitzero" json:"-"`
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
	MaxXactID param.Opt[string] `query:"max_xact_id,omitzero" json:"-"`
	// Retrieve a snapshot of events from a past time
	//
	// The version id is essentially a filter on the latest event transaction id. You
	// can use the `max_xact_id` returned by a past fetch as the version to reproduce
	// that exact fetch.
	Version param.Opt[string] `query:"version,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ExperimentFetchParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ExperimentFetchParams]'s query parameters as `url.Values`.
func (r ExperimentFetchParams) URLQuery() (v url.Values, err error) {
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
	Cursor param.Opt[string] `json:"cursor,omitzero"`
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
	Limit param.Opt[int64] `json:"limit,omitzero"`
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
	MaxRootSpanID param.Opt[string] `json:"max_root_span_id,omitzero"`
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
	MaxXactID param.Opt[string] `json:"max_xact_id,omitzero"`
	// Retrieve a snapshot of events from a past time
	//
	// The version id is essentially a filter on the latest event transaction id. You
	// can use the `max_xact_id` returned by a past fetch as the version to reproduce
	// that exact fetch.
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ExperimentFetchPostParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ExperimentFetchPostParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentFetchPostParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type ExperimentInsertParams struct {
	// A list of experiment events to insert
	Events []shared.InsertExperimentEventParam `json:"events,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ExperimentInsertParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ExperimentInsertParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentInsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type ExperimentSummarizeParams struct {
	// Whether to summarize the scores and metrics. If false (or omitted), only the
	// metadata will be returned.
	SummarizeScores param.Opt[bool] `query:"summarize_scores,omitzero" json:"-"`
	// The experiment to compare against, if summarizing scores and metrics. If
	// omitted, will fall back to the `base_exp_id` stored in the experiment metadata,
	// and then to the most recent experiment run in the same project. Must pass
	// `summarize_scores=true` for this id to be used
	ComparisonExperimentID param.Opt[string] `query:"comparison_experiment_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ExperimentSummarizeParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ExperimentSummarizeParams]'s query parameters as
// `url.Values`.
func (r ExperimentSummarizeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
