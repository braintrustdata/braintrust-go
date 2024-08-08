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
func (r *ProjectLogService) Feedback(ctx context.Context, projectID shared.ProjectIDParam, body ProjectLogFeedbackParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/project_logs/%s/feedback", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Fetch the events in a project logs. Equivalent to the POST form of the same
// path, but with the parameters in the URL query rather than in the request body
func (r *ProjectLogService) Fetch(ctx context.Context, projectID shared.ProjectIDParam, query ProjectLogFetchParams, opts ...option.RequestOption) (res *shared.FetchProjectLogsEventsResponse, err error) {
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
func (r *ProjectLogService) FetchPost(ctx context.Context, projectID shared.ProjectIDParam, body ProjectLogFetchPostParams, opts ...option.RequestOption) (res *shared.FetchProjectLogsEventsResponse, err error) {
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
func (r *ProjectLogService) Insert(ctx context.Context, projectID shared.ProjectIDParam, body ProjectLogInsertParams, opts ...option.RequestOption) (res *shared.InsertEventsResponse, err error) {
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
	FeedbackProjectLogsEventRequest shared.FeedbackProjectLogsEventRequestParam `json:"feedback_project_logs_event_request,required"`
}

func (r ProjectLogFeedbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.FeedbackProjectLogsEventRequest)
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
	Limit param.Field[shared.FetchLimitParam] `query:"limit"`
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
	MaxRootSpanID param.Field[shared.MaxRootSpanIDParam] `query:"max_root_span_id"`
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
	MaxXactID param.Field[shared.MaxXactIDParam] `query:"max_xact_id"`
	// Retrieve a snapshot of events from a past time
	//
	// The version id is essentially a filter on the latest event transaction id. You
	// can use the `max_xact_id` returned by a past fetch as the version to reproduce
	// that exact fetch.
	Version param.Field[shared.VersionParam] `query:"version"`
}

// URLQuery serializes [ProjectLogFetchParams]'s query parameters as `url.Values`.
func (r ProjectLogFetchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectLogFetchPostParams struct {
	FetchEventsRequest shared.FetchEventsRequestParam `json:"fetch_events_request,required"`
}

func (r ProjectLogFetchPostParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.FetchEventsRequest)
}

type ProjectLogInsertParams struct {
	InsertProjectLogsEventRequest shared.InsertProjectLogsEventRequestParam `json:"insert_project_logs_event_request,required"`
}

func (r ProjectLogInsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.InsertProjectLogsEventRequest)
}
