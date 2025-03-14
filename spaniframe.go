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
	"github.com/braintrustdata/braintrust-go/packages/pagination"
	"github.com/braintrustdata/braintrust-go/shared"
)

// SpanIframeService contains methods and other services that help with interacting
// with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpanIframeService] method instead.
type SpanIframeService struct {
	Options []option.RequestOption
}

// NewSpanIframeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpanIframeService(opts ...option.RequestOption) (r *SpanIframeService) {
	r = &SpanIframeService{}
	r.Options = opts
	return
}

// Create a new span_iframe. If there is an existing span_iframe with the same name
// as the one specified in the request, will return the existing span_iframe
// unmodified
func (r *SpanIframeService) New(ctx context.Context, body SpanIframeNewParams, opts ...option.RequestOption) (res *shared.SpanIFrame, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/span_iframe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a span_iframe object by its id
func (r *SpanIframeService) Get(ctx context.Context, spanIframeID string, opts ...option.RequestOption) (res *shared.SpanIFrame, err error) {
	opts = append(r.Options[:], opts...)
	if spanIframeID == "" {
		err = errors.New("missing required span_iframe_id parameter")
		return
	}
	path := fmt.Sprintf("v1/span_iframe/%s", spanIframeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update a span_iframe object. Specify the fields to update in the
// payload. Any object-type fields will be deep-merged with existing content.
// Currently we do not support removing fields or setting them to null.
func (r *SpanIframeService) Update(ctx context.Context, spanIframeID string, body SpanIframeUpdateParams, opts ...option.RequestOption) (res *shared.SpanIFrame, err error) {
	opts = append(r.Options[:], opts...)
	if spanIframeID == "" {
		err = errors.New("missing required span_iframe_id parameter")
		return
	}
	path := fmt.Sprintf("v1/span_iframe/%s", spanIframeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all span_iframes. The span_iframes are sorted by creation date, with
// the most recently-created span_iframes coming first
func (r *SpanIframeService) List(ctx context.Context, query SpanIframeListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.SpanIFrame], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/span_iframe"
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

// List out all span_iframes. The span_iframes are sorted by creation date, with
// the most recently-created span_iframes coming first
func (r *SpanIframeService) ListAutoPaging(ctx context.Context, query SpanIframeListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.SpanIFrame] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a span_iframe object by its id
func (r *SpanIframeService) Delete(ctx context.Context, spanIframeID string, opts ...option.RequestOption) (res *shared.SpanIFrame, err error) {
	opts = append(r.Options[:], opts...)
	if spanIframeID == "" {
		err = errors.New("missing required span_iframe_id parameter")
		return
	}
	path := fmt.Sprintf("v1/span_iframe/%s", spanIframeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create or replace span_iframe. If there is an existing span_iframe with the same
// name as the one specified in the request, will replace the existing span_iframe
// with the provided fields
func (r *SpanIframeService) Replace(ctx context.Context, body SpanIframeReplaceParams, opts ...option.RequestOption) (res *shared.SpanIFrame, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/span_iframe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SpanIframeNewParams struct {
	// Name of the span iframe
	Name param.Field[string] `json:"name,required"`
	// Unique identifier for the project that the span iframe belongs under
	ProjectID param.Field[string] `json:"project_id,required" format:"uuid"`
	// URL to embed the project viewer in an iframe
	URL param.Field[string] `json:"url,required"`
	// Textual description of the span iframe
	Description param.Field[string] `json:"description"`
	// Whether to post messages to the iframe containing the span's data. This is
	// useful when you want to render more data than fits in the URL.
	PostMessage param.Field[bool] `json:"post_message"`
}

func (r SpanIframeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SpanIframeUpdateParams struct {
	// Textual description of the span iframe
	Description param.Field[string] `json:"description"`
	// Name of the span iframe
	Name param.Field[string] `json:"name"`
	// Whether to post messages to the iframe containing the span's data. This is
	// useful when you want to render more data than fits in the URL.
	PostMessage param.Field[bool] `json:"post_message"`
	// URL to embed the project viewer in an iframe
	URL param.Field[string] `json:"url"`
}

func (r SpanIframeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SpanIframeListParams struct {
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[SpanIframeListParamsIDsUnion] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[int64] `query:"limit"`
	// Filter search results to within a particular organization
	OrgName param.Field[string] `query:"org_name"`
	// Name of the span_iframe to search for
	SpanIframeName param.Field[string] `query:"span_iframe_name"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [SpanIframeListParams]'s query parameters as `url.Values`.
func (r SpanIframeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [SpanIframeListParamsIDsArray].
type SpanIframeListParamsIDsUnion interface {
	ImplementsSpanIframeListParamsIDsUnion()
}

type SpanIframeListParamsIDsArray []string

func (r SpanIframeListParamsIDsArray) ImplementsSpanIframeListParamsIDsUnion() {}

type SpanIframeReplaceParams struct {
	// Name of the span iframe
	Name param.Field[string] `json:"name,required"`
	// Unique identifier for the project that the span iframe belongs under
	ProjectID param.Field[string] `json:"project_id,required" format:"uuid"`
	// URL to embed the project viewer in an iframe
	URL param.Field[string] `json:"url,required"`
	// Textual description of the span iframe
	Description param.Field[string] `json:"description"`
	// Whether to post messages to the iframe containing the span's data. This is
	// useful when you want to render more data than fits in the URL.
	PostMessage param.Field[bool] `json:"post_message"`
}

func (r SpanIframeReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
