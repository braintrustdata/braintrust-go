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
	"github.com/braintrustdata/braintrust-go/internal/pagination"
	"github.com/braintrustdata/braintrust-go/internal/param"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/braintrustdata/braintrust-go/shared"
)

// ViewService contains methods and other services that help with interacting with
// the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewViewService] method instead.
type ViewService struct {
	Options []option.RequestOption
}

// NewViewService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewViewService(opts ...option.RequestOption) (r *ViewService) {
	r = &ViewService{}
	r.Options = opts
	return
}

// Create a new view. If there is an existing view with the same name as the one
// specified in the request, will return the existing view unmodified
func (r *ViewService) New(ctx context.Context, body ViewNewParams, opts ...option.RequestOption) (res *shared.View, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/view"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a view object by its id
func (r *ViewService) Get(ctx context.Context, viewID shared.ViewIDParam, query ViewGetParams, opts ...option.RequestOption) (res *shared.View, err error) {
	opts = append(r.Options[:], opts...)
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("v1/view/%s", viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Partially update a view object. Specify the fields to update in the payload. Any
// object-type fields will be deep-merged with existing content. Currently we do
// not support removing fields or setting them to null.
func (r *ViewService) Update(ctx context.Context, viewID shared.ViewIDParam, body ViewUpdateParams, opts ...option.RequestOption) (res *shared.View, err error) {
	opts = append(r.Options[:], opts...)
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("v1/view/%s", viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all views. The views are sorted by creation date, with the most
// recently-created views coming first
func (r *ViewService) List(ctx context.Context, query ViewListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.View], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/view"
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

// List out all views. The views are sorted by creation date, with the most
// recently-created views coming first
func (r *ViewService) ListAutoPaging(ctx context.Context, query ViewListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.View] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a view object by its id
func (r *ViewService) Delete(ctx context.Context, viewID shared.ViewIDParam, body ViewDeleteParams, opts ...option.RequestOption) (res *shared.View, err error) {
	opts = append(r.Options[:], opts...)
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("v1/view/%s", viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Create or replace view. If there is an existing view with the same name as the
// one specified in the request, will replace the existing view with the provided
// fields
func (r *ViewService) Replace(ctx context.Context, body ViewReplaceParams, opts ...option.RequestOption) (res *shared.View, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/view"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ViewNewParams struct {
	CreateView shared.CreateViewParam `json:"create_view,required"`
}

func (r ViewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateView)
}

type ViewGetParams struct {
	// The id of the object the ACL applies to
	ObjectID param.Field[shared.ACLObjectIDParam] `query:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[shared.ACLObjectType] `query:"object_type,required"`
}

// URLQuery serializes [ViewGetParams]'s query parameters as `url.Values`.
func (r ViewGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ViewUpdateParams struct {
	PatchView shared.PatchViewParam `json:"patch_view,required"`
}

func (r ViewUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PatchView)
}

type ViewListParams struct {
	// The id of the object the ACL applies to
	ObjectID param.Field[shared.ACLObjectIDParam] `query:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[shared.ACLObjectType] `query:"object_type,required"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[shared.EndingBeforeParam] `query:"ending_before" format:"uuid"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[shared.IDsUnionParam] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[shared.AppLimitParam] `query:"limit"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Field[shared.StartingAfterParam] `query:"starting_after" format:"uuid"`
	// Name of the view to search for
	ViewName param.Field[shared.ViewNameParam] `query:"view_name"`
	// Type of table that the view corresponds to.
	ViewType param.Field[shared.ViewType] `query:"view_type"`
}

// URLQuery serializes [ViewListParams]'s query parameters as `url.Values`.
func (r ViewListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ViewDeleteParams struct {
	DeleteView shared.DeleteViewParam `json:"delete_view,required"`
}

func (r ViewDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DeleteView)
}

type ViewReplaceParams struct {
	CreateView shared.CreateViewParam `json:"create_view,required"`
}

func (r ViewReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateView)
}
