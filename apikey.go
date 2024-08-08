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

// APIKeyService contains methods and other services that help with interacting
// with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyService] method instead.
type APIKeyService struct {
	Options []option.RequestOption
}

// NewAPIKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIKeyService(opts ...option.RequestOption) (r *APIKeyService) {
	r = &APIKeyService{}
	r.Options = opts
	return
}

// Create a new api_key. It is possible to have multiple API keys with the same
// name. There is no de-duplication
func (r *APIKeyService) New(ctx context.Context, body APIKeyNewParams, opts ...option.RequestOption) (res *APIKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/api_key"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an api_key object by its id
func (r *APIKeyService) Get(ctx context.Context, apiKeyID string, opts ...option.RequestOption) (res *APIKey, err error) {
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
func (r *APIKeyService) List(ctx context.Context, query APIKeyListParams, opts ...option.RequestOption) (res *pagination.ListObjects[APIKey], err error) {
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
func (r *APIKeyService) ListAutoPaging(ctx context.Context, query APIKeyListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[APIKey] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete an api_key object by its id
func (r *APIKeyService) Delete(ctx context.Context, apiKeyID string, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = append(r.Options[:], opts...)
	if apiKeyID == "" {
		err = errors.New("missing required api_key_id parameter")
		return
	}
	path := fmt.Sprintf("v1/api_key/%s", apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type APIKey struct {
	// Unique identifier for the api key
	ID string `json:"id,required" format:"uuid"`
	// Name of the api key
	Name        string `json:"name,required"`
	PreviewName string `json:"preview_name,required"`
	// Date of api key creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Unique identifier for the organization
	OrgID string `json:"org_id,nullable" format:"uuid"`
	// Unique identifier for the user
	UserID string     `json:"user_id,nullable" format:"uuid"`
	JSON   apiKeyJSON `json:"-"`
}

// apiKeyJSON contains the JSON metadata for the struct [APIKey]
type apiKeyJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	PreviewName apijson.Field
	Created     apijson.Field
	OrgID       apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyJSON) RawJSON() string {
	return r.raw
}

type APIKeyNewResponse struct {
	// Unique identifier for the api key
	ID string `json:"id,required" format:"uuid"`
	// The raw API key. It will only be exposed this one time
	Key string `json:"key,required"`
	// Name of the api key
	Name        string `json:"name,required"`
	PreviewName string `json:"preview_name,required"`
	// Date of api key creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Unique identifier for the organization
	OrgID string `json:"org_id,nullable" format:"uuid"`
	// Unique identifier for the user
	UserID string                `json:"user_id,nullable" format:"uuid"`
	JSON   apiKeyNewResponseJSON `json:"-"`
}

// apiKeyNewResponseJSON contains the JSON metadata for the struct
// [APIKeyNewResponse]
type apiKeyNewResponseJSON struct {
	ID          apijson.Field
	Key         apijson.Field
	Name        apijson.Field
	PreviewName apijson.Field
	Created     apijson.Field
	OrgID       apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

type APIKeyNewParams struct {
	// Name of the api key. Does not have to be unique
	Name param.Field[string] `json:"name,required"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the API key belongs in.
	OrgName param.Field[string] `json:"org_name"`
}

func (r APIKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type APIKeyListParams struct {
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
	IDs param.Field[APIKeyListParamsIDsUnion] `query:"ids" format:"uuid"`
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

// URLQuery serializes [APIKeyListParams]'s query parameters as `url.Values`.
func (r APIKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [APIKeyListParamsIDsArray].
type APIKeyListParamsIDsUnion interface {
	ImplementsAPIKeyListParamsIDsUnion()
}

type APIKeyListParamsIDsArray []string

func (r APIKeyListParamsIDsArray) ImplementsAPIKeyListParamsIDsUnion() {}
