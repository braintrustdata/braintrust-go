// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/internal/apiquery"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/braintrustdata/braintrust-go/packages/pagination"
	"github.com/braintrustdata/braintrust-go/packages/param"
	"github.com/braintrustdata/braintrust-go/shared"
)

// DatasetService contains methods and other services that help with interacting
// with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatasetService] method instead.
type DatasetService struct {
	Options []option.RequestOption
}

// NewDatasetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDatasetService(opts ...option.RequestOption) (r DatasetService) {
	r = DatasetService{}
	r.Options = opts
	return
}

// Create a new dataset. If there is an existing dataset in the project with the
// same name as the one specified in the request, will return the existing dataset
// unmodified
func (r *DatasetService) New(ctx context.Context, body DatasetNewParams, opts ...option.RequestOption) (res *shared.Dataset, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/dataset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a dataset object by its id
func (r *DatasetService) Get(ctx context.Context, datasetID string, opts ...option.RequestOption) (res *shared.Dataset, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dataset/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update a dataset object. Specify the fields to update in the payload.
// Any object-type fields will be deep-merged with existing content. Currently we
// do not support removing fields or setting them to null.
func (r *DatasetService) Update(ctx context.Context, datasetID string, body DatasetUpdateParams, opts ...option.RequestOption) (res *shared.Dataset, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dataset/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all datasets. The datasets are sorted by creation date, with the most
// recently-created datasets coming first
func (r *DatasetService) List(ctx context.Context, query DatasetListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.Dataset], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/dataset"
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

// List out all datasets. The datasets are sorted by creation date, with the most
// recently-created datasets coming first
func (r *DatasetService) ListAutoPaging(ctx context.Context, query DatasetListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.Dataset] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a dataset object by its id
func (r *DatasetService) Delete(ctx context.Context, datasetID string, opts ...option.RequestOption) (res *shared.Dataset, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dataset/%s", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Log feedback for a set of dataset events
func (r *DatasetService) Feedback(ctx context.Context, datasetID string, body DatasetFeedbackParams, opts ...option.RequestOption) (res *shared.FeedbackResponseSchema, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dataset/%s/feedback", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch the events in a dataset. Equivalent to the POST form of the same path, but
// with the parameters in the URL query rather than in the request body. For more
// complex queries, use the `POST /btql` endpoint.
func (r *DatasetService) Fetch(ctx context.Context, datasetID string, query DatasetFetchParams, opts ...option.RequestOption) (res *shared.FetchDatasetEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dataset/%s/fetch", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch the events in a dataset. Equivalent to the GET form of the same path, but
// with the parameters in the request body rather than in the URL query. For more
// complex queries, use the `POST /btql` endpoint.
func (r *DatasetService) FetchPost(ctx context.Context, datasetID string, body DatasetFetchPostParams, opts ...option.RequestOption) (res *shared.FetchDatasetEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dataset/%s/fetch", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Insert a set of events into the dataset
func (r *DatasetService) Insert(ctx context.Context, datasetID string, body DatasetInsertParams, opts ...option.RequestOption) (res *shared.InsertEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dataset/%s/insert", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Summarize dataset
func (r *DatasetService) Summarize(ctx context.Context, datasetID string, query DatasetSummarizeParams, opts ...option.RequestOption) (res *shared.SummarizeDatasetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dataset/%s/summarize", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DatasetNewParams struct {
	// Name of the dataset. Within a project, dataset names are unique
	Name string `json:"name,required"`
	// Unique identifier for the project that the dataset belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// Textual description of the dataset
	Description param.Opt[string] `json:"description,omitzero"`
	// User-controlled metadata about the dataset
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r DatasetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DatasetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatasetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatasetUpdateParams struct {
	// Textual description of the dataset
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the dataset. Within a project, dataset names are unique
	Name param.Opt[string] `json:"name,omitzero"`
	// User-controlled metadata about the dataset
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r DatasetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DatasetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatasetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatasetListParams struct {
	// Limit the number of objects to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Name of the dataset to search for
	DatasetName param.Opt[string] `query:"dataset_name,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Opt[string] `query:"ending_before,omitzero" format:"uuid" json:"-"`
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
	IDs DatasetListParamsIDsUnion `query:"ids,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [DatasetListParams]'s query parameters as `url.Values`.
func (r DatasetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DatasetListParamsIDsUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *DatasetListParamsIDsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type DatasetFeedbackParams struct {
	// A list of dataset feedback items
	Feedback []shared.FeedbackDatasetItemParam `json:"feedback,omitzero,required"`
	paramObj
}

func (r DatasetFeedbackParams) MarshalJSON() (data []byte, err error) {
	type shadow DatasetFeedbackParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatasetFeedbackParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatasetFetchParams struct {
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

// URLQuery serializes [DatasetFetchParams]'s query parameters as `url.Values`.
func (r DatasetFetchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DatasetFetchPostParams struct {
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

func (r DatasetFetchPostParams) MarshalJSON() (data []byte, err error) {
	type shadow DatasetFetchPostParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatasetFetchPostParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatasetInsertParams struct {
	// A list of dataset events to insert
	Events []shared.InsertDatasetEventParam `json:"events,omitzero,required"`
	paramObj
}

func (r DatasetInsertParams) MarshalJSON() (data []byte, err error) {
	type shadow DatasetInsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatasetInsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatasetSummarizeParams struct {
	// Whether to summarize the data. If false (or omitted), only the metadata will be
	// returned.
	SummarizeData param.Opt[bool] `query:"summarize_data,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DatasetSummarizeParams]'s query parameters as `url.Values`.
func (r DatasetSummarizeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
