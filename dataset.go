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
	"github.com/braintrustdata/braintrust-go/internal/pagination"
	"github.com/braintrustdata/braintrust-go/internal/param"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
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
func NewDatasetService(opts ...option.RequestOption) (r *DatasetService) {
	r = &DatasetService{}
	r.Options = opts
	return
}

// Create a new dataset. If there is an existing dataset in the project with the
// same name as the one specified in the request, will return the existing dataset
// unmodified
func (r *DatasetService) New(ctx context.Context, body DatasetNewParams, opts ...option.RequestOption) (res *Dataset, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/dataset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a dataset object by its id
func (r *DatasetService) Get(ctx context.Context, datasetID string, opts ...option.RequestOption) (res *Dataset, err error) {
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
func (r *DatasetService) Update(ctx context.Context, datasetID string, body DatasetUpdateParams, opts ...option.RequestOption) (res *Dataset, err error) {
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
func (r *DatasetService) List(ctx context.Context, query DatasetListParams, opts ...option.RequestOption) (res *pagination.ListObjects[Dataset], err error) {
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
func (r *DatasetService) ListAutoPaging(ctx context.Context, query DatasetListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[Dataset] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a dataset object by its id
func (r *DatasetService) Delete(ctx context.Context, datasetID string, opts ...option.RequestOption) (res *Dataset, err error) {
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
func (r *DatasetService) Feedback(ctx context.Context, datasetID string, body DatasetFeedbackParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dataset/%s/feedback", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Fetch the events in a dataset. Equivalent to the POST form of the same path, but
// with the parameters in the URL query rather than in the request body
func (r *DatasetService) Fetch(ctx context.Context, datasetID string, query DatasetFetchParams, opts ...option.RequestOption) (res *DatasetFetchResponse, err error) {
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
// with the parameters in the request body rather than in the URL query
func (r *DatasetService) FetchPost(ctx context.Context, datasetID string, body DatasetFetchPostParams, opts ...option.RequestOption) (res *DatasetFetchPostResponse, err error) {
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
func (r *DatasetService) Insert(ctx context.Context, datasetID string, body DatasetInsertParams, opts ...option.RequestOption) (res *DatasetInsertResponse, err error) {
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
func (r *DatasetService) Summarize(ctx context.Context, datasetID string, query DatasetSummarizeParams, opts ...option.RequestOption) (res *DatasetSummarizeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dataset/%s/summarize", datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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

type DatasetFetchResponse struct {
	// A list of fetched events
	Events []DatasetFetchResponseEvent `json:"events,required"`
	// Pagination cursor
	//
	// Pass this string directly as the `cursor` param to your next fetch request to
	// get the next page of results. Not provided if the returned result set is empty.
	Cursor string                   `json:"cursor,nullable"`
	JSON   datasetFetchResponseJSON `json:"-"`
}

// datasetFetchResponseJSON contains the JSON metadata for the struct
// [DatasetFetchResponse]
type datasetFetchResponseJSON struct {
	Events      apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetFetchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetFetchResponseJSON) RawJSON() string {
	return r.raw
}

type DatasetFetchResponseEvent struct {
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
	RootSpanID string `json:"root_span_id,required" format:"uuid"`
	// A unique identifier used to link different dataset events together as part of a
	// full trace. See the
	// [tracing guide](https://www.braintrust.dev/docs/guides/tracing) for full details
	// on tracing
	SpanID string `json:"span_id,required" format:"uuid"`
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
	Tags []string                      `json:"tags,nullable"`
	JSON datasetFetchResponseEventJSON `json:"-"`
}

// datasetFetchResponseEventJSON contains the JSON metadata for the struct
// [DatasetFetchResponseEvent]
type datasetFetchResponseEventJSON struct {
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

func (r *DatasetFetchResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetFetchResponseEventJSON) RawJSON() string {
	return r.raw
}

type DatasetFetchPostResponse struct {
	// A list of fetched events
	Events []DatasetFetchPostResponseEvent `json:"events,required"`
	// Pagination cursor
	//
	// Pass this string directly as the `cursor` param to your next fetch request to
	// get the next page of results. Not provided if the returned result set is empty.
	Cursor string                       `json:"cursor,nullable"`
	JSON   datasetFetchPostResponseJSON `json:"-"`
}

// datasetFetchPostResponseJSON contains the JSON metadata for the struct
// [DatasetFetchPostResponse]
type datasetFetchPostResponseJSON struct {
	Events      apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetFetchPostResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetFetchPostResponseJSON) RawJSON() string {
	return r.raw
}

type DatasetFetchPostResponseEvent struct {
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
	RootSpanID string `json:"root_span_id,required" format:"uuid"`
	// A unique identifier used to link different dataset events together as part of a
	// full trace. See the
	// [tracing guide](https://www.braintrust.dev/docs/guides/tracing) for full details
	// on tracing
	SpanID string `json:"span_id,required" format:"uuid"`
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
	Tags []string                          `json:"tags,nullable"`
	JSON datasetFetchPostResponseEventJSON `json:"-"`
}

// datasetFetchPostResponseEventJSON contains the JSON metadata for the struct
// [DatasetFetchPostResponseEvent]
type datasetFetchPostResponseEventJSON struct {
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

func (r *DatasetFetchPostResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetFetchPostResponseEventJSON) RawJSON() string {
	return r.raw
}

type DatasetInsertResponse struct {
	// The ids of all rows that were inserted, aligning one-to-one with the rows
	// provided as input
	RowIDs []string                  `json:"row_ids,required"`
	JSON   datasetInsertResponseJSON `json:"-"`
}

// datasetInsertResponseJSON contains the JSON metadata for the struct
// [DatasetInsertResponse]
type datasetInsertResponseJSON struct {
	RowIDs      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetInsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetInsertResponseJSON) RawJSON() string {
	return r.raw
}

// Summary of a dataset
type DatasetSummarizeResponse struct {
	// Name of the dataset
	DatasetName string `json:"dataset_name,required"`
	// URL to the dataset's page in the Braintrust app
	DatasetURL string `json:"dataset_url,required" format:"uri"`
	// Name of the project that the dataset belongs to
	ProjectName string `json:"project_name,required"`
	// URL to the project's page in the Braintrust app
	ProjectURL string `json:"project_url,required" format:"uri"`
	// Summary of a dataset's data
	DataSummary DatasetSummarizeResponseDataSummary `json:"data_summary,nullable"`
	JSON        datasetSummarizeResponseJSON        `json:"-"`
}

// datasetSummarizeResponseJSON contains the JSON metadata for the struct
// [DatasetSummarizeResponse]
type datasetSummarizeResponseJSON struct {
	DatasetName apijson.Field
	DatasetURL  apijson.Field
	ProjectName apijson.Field
	ProjectURL  apijson.Field
	DataSummary apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetSummarizeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetSummarizeResponseJSON) RawJSON() string {
	return r.raw
}

// Summary of a dataset's data
type DatasetSummarizeResponseDataSummary struct {
	// Total number of records in the dataset
	TotalRecords int64                                   `json:"total_records,required"`
	JSON         datasetSummarizeResponseDataSummaryJSON `json:"-"`
}

// datasetSummarizeResponseDataSummaryJSON contains the JSON metadata for the
// struct [DatasetSummarizeResponseDataSummary]
type datasetSummarizeResponseDataSummaryJSON struct {
	TotalRecords apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DatasetSummarizeResponseDataSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetSummarizeResponseDataSummaryJSON) RawJSON() string {
	return r.raw
}

type DatasetNewParams struct {
	// Name of the dataset. Within a project, dataset names are unique
	Name param.Field[string] `json:"name,required"`
	// Textual description of the dataset
	Description param.Field[string] `json:"description"`
	// Unique identifier for the project that the dataset belongs under
	ProjectID param.Field[string] `json:"project_id" format:"uuid"`
}

func (r DatasetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatasetUpdateParams struct {
	// Textual description of the dataset
	Description param.Field[string] `json:"description"`
	// User-controlled metadata about the dataset
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// Name of the dataset. Within a project, dataset names are unique
	Name param.Field[string] `json:"name"`
}

func (r DatasetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatasetListParams struct {
	// Name of the dataset to search for
	DatasetName param.Field[string] `query:"dataset_name"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[DatasetListParamsIDsUnion] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[int64] `query:"limit"`
	// Filter search results to within a particular organization
	OrgName param.Field[string] `query:"org_name"`
	// Name of the project to search for
	ProjectName param.Field[string] `query:"project_name"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [DatasetListParams]'s query parameters as `url.Values`.
func (r DatasetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [DatasetListParamsIDsArray].
type DatasetListParamsIDsUnion interface {
	ImplementsDatasetListParamsIDsUnion()
}

type DatasetListParamsIDsArray []string

func (r DatasetListParamsIDsArray) ImplementsDatasetListParamsIDsUnion() {}

type DatasetFeedbackParams struct {
	// A list of dataset feedback items
	Feedback param.Field[[]DatasetFeedbackParamsFeedback] `json:"feedback,required"`
}

func (r DatasetFeedbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatasetFeedbackParamsFeedback struct {
	// The id of the dataset event to log feedback for. This is the row `id` returned
	// by `POST /v1/dataset/{dataset_id}/insert`
	ID param.Field[string] `json:"id,required"`
	// An optional comment string to log about the dataset event
	Comment param.Field[string] `json:"comment"`
	// A dictionary with additional data about the feedback. If you have a `user_id`,
	// you can log it here and access it in the Braintrust UI.
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// The source of the feedback. Must be one of "external" (default), "app", or "api"
	Source param.Field[DatasetFeedbackParamsFeedbackSource] `json:"source"`
}

func (r DatasetFeedbackParamsFeedback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The source of the feedback. Must be one of "external" (default), "app", or "api"
type DatasetFeedbackParamsFeedbackSource string

const (
	DatasetFeedbackParamsFeedbackSourceApp      DatasetFeedbackParamsFeedbackSource = "app"
	DatasetFeedbackParamsFeedbackSourceAPI      DatasetFeedbackParamsFeedbackSource = "api"
	DatasetFeedbackParamsFeedbackSourceExternal DatasetFeedbackParamsFeedbackSource = "external"
)

func (r DatasetFeedbackParamsFeedbackSource) IsKnown() bool {
	switch r {
	case DatasetFeedbackParamsFeedbackSourceApp, DatasetFeedbackParamsFeedbackSourceAPI, DatasetFeedbackParamsFeedbackSourceExternal:
		return true
	}
	return false
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

// URLQuery serializes [DatasetFetchParams]'s query parameters as `url.Values`.
func (r DatasetFetchParams) URLQuery() (v url.Values) {
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
	Cursor param.Field[string] `json:"cursor"`
	// A list of filters on the events to fetch. Currently, only path-lookup type
	// filters are supported, but we may add more in the future
	Filters param.Field[[]DatasetFetchPostParamsFilter] `json:"filters"`
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

func (r DatasetFetchPostParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A path-lookup filter describes an equality comparison against a specific
// sub-field in the event row. For instance, if you wish to filter on the value of
// `c` in `{"input": {"a": {"b": {"c": "hello"}}}}`, pass
// `path=["input", "a", "b", "c"]` and `value="hello"`
type DatasetFetchPostParamsFilter struct {
	// List of fields describing the path to the value to be checked against. For
	// instance, if you wish to filter on the value of `c` in
	// `{"input": {"a": {"b": {"c": "hello"}}}}`, pass `path=["input", "a", "b", "c"]`
	Path param.Field[[]string] `json:"path,required"`
	// Denotes the type of filter as a path-lookup filter
	Type param.Field[DatasetFetchPostParamsFiltersType] `json:"type,required"`
	// The value to compare equality-wise against the event value at the specified
	// `path`. The value must be a "primitive", that is, any JSON-serializable object
	// except for objects and arrays. For instance, if you wish to filter on the value
	// of "input.a.b.c" in the object `{"input": {"a": {"b": {"c": "hello"}}}}`, pass
	// `value="hello"`
	Value param.Field[interface{}] `json:"value"`
}

func (r DatasetFetchPostParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Denotes the type of filter as a path-lookup filter
type DatasetFetchPostParamsFiltersType string

const (
	DatasetFetchPostParamsFiltersTypePathLookup DatasetFetchPostParamsFiltersType = "path_lookup"
)

func (r DatasetFetchPostParamsFiltersType) IsKnown() bool {
	switch r {
	case DatasetFetchPostParamsFiltersTypePathLookup:
		return true
	}
	return false
}

type DatasetInsertParams struct {
	// A list of dataset events to insert
	Events param.Field[[]DatasetInsertParamsEventUnion] `json:"events,required"`
}

func (r DatasetInsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A dataset event
type DatasetInsertParamsEvent struct {
	Input    param.Field[interface{}] `json:"input,required"`
	Expected param.Field[interface{}] `json:"expected,required"`
	Metadata param.Field[interface{}] `json:"metadata,required"`
	Tags     param.Field[interface{}] `json:"tags,required"`
	// A unique identifier for the dataset event. If you don't provide one, BrainTrust
	// will generate one for you
	ID param.Field[string] `json:"id"`
	// The timestamp the dataset event was created
	Created param.Field[time.Time] `json:"created" format:"date-time"`
	// Pass `_object_delete=true` to mark the dataset event deleted. Deleted events
	// will not show up in subsequent fetches for this dataset
	ObjectDelete param.Field[bool] `json:"_object_delete"`
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
	ParentID   param.Field[string]      `json:"_parent_id"`
	MergePaths param.Field[interface{}] `json:"_merge_paths,required"`
}

func (r DatasetInsertParamsEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DatasetInsertParamsEvent) implementsDatasetInsertParamsEventUnion() {}

// A dataset event
//
// Satisfied by [DatasetInsertParamsEventsInsertDatasetEventReplace],
// [DatasetInsertParamsEventsInsertDatasetEventMerge], [DatasetInsertParamsEvent].
type DatasetInsertParamsEventUnion interface {
	implementsDatasetInsertParamsEventUnion()
}

type DatasetInsertParamsEventsInsertDatasetEventReplace struct {
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

func (r DatasetInsertParamsEventsInsertDatasetEventReplace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DatasetInsertParamsEventsInsertDatasetEventReplace) implementsDatasetInsertParamsEventUnion() {
}

type DatasetInsertParamsEventsInsertDatasetEventMerge struct {
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

func (r DatasetInsertParamsEventsInsertDatasetEventMerge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DatasetInsertParamsEventsInsertDatasetEventMerge) implementsDatasetInsertParamsEventUnion() {}

type DatasetSummarizeParams struct {
	// Whether to summarize the data. If false (or omitted), only the metadata will be
	// returned.
	SummarizeData param.Field[bool] `query:"summarize_data"`
}

// URLQuery serializes [DatasetSummarizeParams]'s query parameters as `url.Values`.
func (r DatasetSummarizeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}