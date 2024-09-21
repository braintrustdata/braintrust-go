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

// APIKeyResourceService contains methods and other services that help with
// interacting with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyResourceService] method instead.
type APIKeyResourceService struct {
	Options []option.RequestOption
}

// NewAPIKeyResourceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIKeyResourceService(opts ...option.RequestOption) (r *APIKeyResourceService) {
	r = &APIKeyResourceService{}
	r.Options = opts
	return
}

// Create a new api_key. It is possible to have multiple API keys with the same
// name. There is no de-duplication
func (r *APIKeyResourceService) New(ctx context.Context, body APIKeyResourceNewParams, opts ...option.RequestOption) (res *shared.CreateAPIKeyOutput, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/api_key"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an api_key object by its id
func (r *APIKeyResourceService) Get(ctx context.Context, apiKeyID string, opts ...option.RequestOption) (res *shared.APIKey, err error) {
	opts = append(r.Options[:], opts...)
	if apiKeyID == "" {
		err = errors.New("missing required api_key_id parameter")
		return
	}
	path := fmt.Sprintf("v1/api_key/%s", apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List out all api_keys. The api_keys are sorted by creation date, with the most
// recently-created api_keys coming first
func (r *APIKeyResourceService) List(ctx context.Context, query APIKeyResourceListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.APIKey], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/api_key"
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

// List out all api_keys. The api_keys are sorted by creation date, with the most
// recently-created api_keys coming first
func (r *APIKeyResourceService) ListAutoPaging(ctx context.Context, query APIKeyResourceListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.APIKey] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete an api_key object by its id
func (r *APIKeyResourceService) Delete(ctx context.Context, apiKeyID string, opts ...option.RequestOption) (res *shared.APIKey, err error) {
	opts = append(r.Options[:], opts...)
	if apiKeyID == "" {
		err = errors.New("missing required api_key_id parameter")
		return
	}
	path := fmt.Sprintf("v1/api_key/%s", apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type APIKeyResourceNewParams struct {
	// Name of the api key. Does not have to be unique
	Name param.Field[string] `json:"name,required"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the API key belongs in.
	OrgName param.Field[string] `json:"org_name"`
}

func (r APIKeyResourceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type APIKeyResourceListParams struct {
	// Name of the api_key to search for
	APIKeyName param.Field[string] `query:"api_key_name"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[APIKeyResourceListParamsIDsUnion] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[int64] `query:"limit"`
	// Filter search results to within a particular organization
	OrgName param.Field[string] `query:"org_name"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [APIKeyResourceListParams]'s query parameters as
// `url.Values`.
func (r APIKeyResourceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [APIKeyResourceListParamsIDsArray].
type APIKeyResourceListParamsIDsUnion interface {
	ImplementsAPIKeyResourceListParamsIDsUnion()
}

type APIKeyResourceListParamsIDsArray []string

func (r APIKeyResourceListParamsIDsArray) ImplementsAPIKeyResourceListParamsIDsUnion() {}
