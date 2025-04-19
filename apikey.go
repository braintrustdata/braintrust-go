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
func NewAPIKeyService(opts ...option.RequestOption) (r APIKeyService) {
	r = APIKeyService{}
	r.Options = opts
	return
}

// Create a new api_key. It is possible to have multiple API keys with the same
// name. There is no de-duplication
func (r *APIKeyService) New(ctx context.Context, body APIKeyNewParams, opts ...option.RequestOption) (res *shared.CreateAPIKeyOutput, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/api_key"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an api_key object by its id
func (r *APIKeyService) Get(ctx context.Context, apiKeyID string, opts ...option.RequestOption) (res *shared.APIKey, err error) {
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
func (r *APIKeyService) List(ctx context.Context, query APIKeyListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.APIKey], err error) {
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
func (r *APIKeyService) ListAutoPaging(ctx context.Context, query APIKeyListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.APIKey] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete an api_key object by its id
func (r *APIKeyService) Delete(ctx context.Context, apiKeyID string, opts ...option.RequestOption) (res *shared.APIKey, err error) {
	opts = append(r.Options[:], opts...)
	if apiKeyID == "" {
		err = errors.New("missing required api_key_id parameter")
		return
	}
	path := fmt.Sprintf("v1/api_key/%s", apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type APIKeyNewParams struct {
	// Name of the api key. Does not have to be unique
	Name string `json:"name,required"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the API key belongs in.
	OrgName param.Opt[string] `json:"org_name,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f APIKeyNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r APIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type APIKeyListParams struct {
	// Limit the number of objects to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Name of the api_key to search for
	APIKeyName param.Opt[string] `query:"api_key_name,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Opt[string] `query:"ending_before,omitzero" format:"uuid" json:"-"`
	// Filter search results to within a particular organization
	OrgName param.Opt[string] `query:"org_name,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Opt[string] `query:"starting_after,omitzero" format:"uuid" json:"-"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs APIKeyListParamsIDsUnion `query:"ids,omitzero" format:"uuid" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f APIKeyListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [APIKeyListParams]'s query parameters as `url.Values`.
func (r APIKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type APIKeyListParamsIDsUnion struct {
	OfString              param.Opt[string] `query:",omitzero,inline"`
	OfAPIKeyListsIDsArray []string          `query:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u APIKeyListParamsIDsUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }

func (u *APIKeyListParamsIDsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAPIKeyListsIDsArray) {
		return &u.OfAPIKeyListsIDsArray
	}
	return nil
}
