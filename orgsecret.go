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

// OrgSecretService contains methods and other services that help with interacting
// with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgSecretService] method instead.
type OrgSecretService struct {
	Options []option.RequestOption
}

// NewOrgSecretService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrgSecretService(opts ...option.RequestOption) (r *OrgSecretService) {
	r = &OrgSecretService{}
	r.Options = opts
	return
}

// Create a new org_secret. If there is an existing org_secret with the same name
// as the one specified in the request, will return the existing org_secret
// unmodified
func (r *OrgSecretService) New(ctx context.Context, body OrgSecretNewParams, opts ...option.RequestOption) (res *shared.OrgSecret, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/org_secret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an org_secret object by its id
func (r *OrgSecretService) Get(ctx context.Context, orgSecretID string, opts ...option.RequestOption) (res *shared.OrgSecret, err error) {
	opts = append(r.Options[:], opts...)
	if orgSecretID == "" {
		err = errors.New("missing required org_secret_id parameter")
		return
	}
	path := fmt.Sprintf("v1/org_secret/%s", orgSecretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update an org_secret object. Specify the fields to update in the
// payload. Any object-type fields will be deep-merged with existing content.
// Currently we do not support removing fields or setting them to null.
func (r *OrgSecretService) Update(ctx context.Context, orgSecretID string, body OrgSecretUpdateParams, opts ...option.RequestOption) (res *shared.OrgSecret, err error) {
	opts = append(r.Options[:], opts...)
	if orgSecretID == "" {
		err = errors.New("missing required org_secret_id parameter")
		return
	}
	path := fmt.Sprintf("v1/org_secret/%s", orgSecretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all org_secrets. The org_secrets are sorted by creation date, with the
// most recently-created org_secrets coming first
func (r *OrgSecretService) List(ctx context.Context, query OrgSecretListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.OrgSecret], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/org_secret"
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

// List out all org_secrets. The org_secrets are sorted by creation date, with the
// most recently-created org_secrets coming first
func (r *OrgSecretService) ListAutoPaging(ctx context.Context, query OrgSecretListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.OrgSecret] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete an org_secret object by its id
func (r *OrgSecretService) Delete(ctx context.Context, orgSecretID string, opts ...option.RequestOption) (res *shared.OrgSecret, err error) {
	opts = append(r.Options[:], opts...)
	if orgSecretID == "" {
		err = errors.New("missing required org_secret_id parameter")
		return
	}
	path := fmt.Sprintf("v1/org_secret/%s", orgSecretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create or replace org_secret. If there is an existing org_secret with the same
// name as the one specified in the request, will replace the existing org_secret
// with the provided fields
func (r *OrgSecretService) Replace(ctx context.Context, body OrgSecretReplaceParams, opts ...option.RequestOption) (res *shared.OrgSecret, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/org_secret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type OrgSecretNewParams struct {
	// Name of the org secret
	Name     param.Field[string]                 `json:"name,required"`
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the Org Secret belongs in.
	OrgName param.Field[string] `json:"org_name"`
	// Secret value. If omitted in a PUT request, the existing secret value will be
	// left intact, not replaced with null.
	Secret param.Field[string] `json:"secret"`
	Type   param.Field[string] `json:"type"`
}

func (r OrgSecretNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrgSecretUpdateParams struct {
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// Name of the org secret
	Name   param.Field[string] `json:"name"`
	Secret param.Field[string] `json:"secret"`
	Type   param.Field[string] `json:"type"`
}

func (r OrgSecretUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrgSecretListParams struct {
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[OrgSecretListParamsIDsUnion] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[int64] `query:"limit"`
	// Filter search results to within a particular organization
	OrgName param.Field[string] `query:"org_name"`
	// Name of the org_secret to search for
	OrgSecretName param.Field[string]                                `query:"org_secret_name"`
	OrgSecretType param.Field[OrgSecretListParamsOrgSecretTypeUnion] `query:"org_secret_type"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [OrgSecretListParams]'s query parameters as `url.Values`.
func (r OrgSecretListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [OrgSecretListParamsIDsArray].
type OrgSecretListParamsIDsUnion interface {
	ImplementsOrgSecretListParamsIDsUnion()
}

type OrgSecretListParamsIDsArray []string

func (r OrgSecretListParamsIDsArray) ImplementsOrgSecretListParamsIDsUnion() {}

// Satisfied by [shared.UnionString], [OrgSecretListParamsOrgSecretTypeArray].
type OrgSecretListParamsOrgSecretTypeUnion interface {
	ImplementsOrgSecretListParamsOrgSecretTypeUnion()
}

type OrgSecretListParamsOrgSecretTypeArray []string

func (r OrgSecretListParamsOrgSecretTypeArray) ImplementsOrgSecretListParamsOrgSecretTypeUnion() {}

type OrgSecretReplaceParams struct {
	// Name of the org secret
	Name     param.Field[string]                 `json:"name,required"`
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the Org Secret belongs in.
	OrgName param.Field[string] `json:"org_name"`
	// Secret value. If omitted in a PUT request, the existing secret value will be
	// left intact, not replaced with null.
	Secret param.Field[string] `json:"secret"`
	Type   param.Field[string] `json:"type"`
}

func (r OrgSecretReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
