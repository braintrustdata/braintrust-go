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
func NewSpanIframeService(opts ...option.RequestOption) (r SpanIframeService) {
	r = SpanIframeService{}
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
	Name string `json:"name,required"`
	// Unique identifier for the project that the span iframe belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// URL to embed the project viewer in an iframe
	URL string `json:"url,required"`
	// Textual description of the span iframe
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether to post messages to the iframe containing the span's data. This is
	// useful when you want to render more data than fits in the URL.
	PostMessage param.Opt[bool] `json:"post_message,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SpanIframeNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SpanIframeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SpanIframeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type SpanIframeUpdateParams struct {
	// Textual description of the span iframe
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the span iframe
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether to post messages to the iframe containing the span's data. This is
	// useful when you want to render more data than fits in the URL.
	PostMessage param.Opt[bool] `json:"post_message,omitzero"`
	// URL to embed the project viewer in an iframe
	URL param.Opt[string] `json:"url,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SpanIframeUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SpanIframeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SpanIframeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type SpanIframeListParams struct {
	// Limit the number of objects to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Opt[string] `query:"ending_before,omitzero" format:"uuid" json:"-"`
	// Filter search results to within a particular organization
	OrgName param.Opt[string] `query:"org_name,omitzero" json:"-"`
	// Name of the span_iframe to search for
	SpanIframeName param.Opt[string] `query:"span_iframe_name,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Opt[string] `query:"starting_after,omitzero" format:"uuid" json:"-"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs SpanIframeListParamsIDsUnion `query:"ids,omitzero" format:"uuid" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SpanIframeListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SpanIframeListParams]'s query parameters as `url.Values`.
func (r SpanIframeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SpanIframeListParamsIDsUnion struct {
	OfString                  param.Opt[string] `json:",omitzero,inline"`
	OfSpanIframeListsIDsArray []string          `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u SpanIframeListParamsIDsUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u SpanIframeListParamsIDsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[SpanIframeListParamsIDsUnion](u.OfString, u.OfSpanIframeListsIDsArray)
}

func (u *SpanIframeListParamsIDsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfSpanIframeListsIDsArray) {
		return &u.OfSpanIframeListsIDsArray
	}
	return nil
}

type SpanIframeReplaceParams struct {
	// Name of the span iframe
	Name string `json:"name,required"`
	// Unique identifier for the project that the span iframe belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// URL to embed the project viewer in an iframe
	URL string `json:"url,required"`
	// Textual description of the span iframe
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether to post messages to the iframe containing the span's data. This is
	// useful when you want to render more data than fits in the URL.
	PostMessage param.Opt[bool] `json:"post_message,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SpanIframeReplaceParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SpanIframeReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow SpanIframeReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
