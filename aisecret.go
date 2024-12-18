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

// AISecretService contains methods and other services that help with interacting
// with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAISecretService] method instead.
type AISecretService struct {
	Options []option.RequestOption
}

// NewAISecretService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAISecretService(opts ...option.RequestOption) (r *AISecretService) {
	r = &AISecretService{}
	r.Options = opts
	return
}

// Create a new ai_secret. If there is an existing ai_secret with the same name as
// the one specified in the request, will return the existing ai_secret unmodified
func (r *AISecretService) New(ctx context.Context, body AISecretNewParams, opts ...option.RequestOption) (res *shared.AISecret, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/ai_secret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an ai_secret object by its id
func (r *AISecretService) Get(ctx context.Context, aiSecretID string, opts ...option.RequestOption) (res *shared.AISecret, err error) {
	opts = append(r.Options[:], opts...)
	if aiSecretID == "" {
		err = errors.New("missing required ai_secret_id parameter")
		return
	}
	path := fmt.Sprintf("v1/ai_secret/%s", aiSecretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update an ai_secret object. Specify the fields to update in the
// payload. Any object-type fields will be deep-merged with existing content.
// Currently we do not support removing fields or setting them to null.
func (r *AISecretService) Update(ctx context.Context, aiSecretID string, body AISecretUpdateParams, opts ...option.RequestOption) (res *shared.AISecret, err error) {
	opts = append(r.Options[:], opts...)
	if aiSecretID == "" {
		err = errors.New("missing required ai_secret_id parameter")
		return
	}
	path := fmt.Sprintf("v1/ai_secret/%s", aiSecretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all ai_secrets. The ai_secrets are sorted by creation date, with the
// most recently-created ai_secrets coming first
func (r *AISecretService) List(ctx context.Context, query AISecretListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.AISecret], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/ai_secret"
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

// List out all ai_secrets. The ai_secrets are sorted by creation date, with the
// most recently-created ai_secrets coming first
func (r *AISecretService) ListAutoPaging(ctx context.Context, query AISecretListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.AISecret] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete an ai_secret object by its id
func (r *AISecretService) Delete(ctx context.Context, aiSecretID string, opts ...option.RequestOption) (res *shared.AISecret, err error) {
	opts = append(r.Options[:], opts...)
	if aiSecretID == "" {
		err = errors.New("missing required ai_secret_id parameter")
		return
	}
	path := fmt.Sprintf("v1/ai_secret/%s", aiSecretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Delete a single ai_secret
func (r *AISecretService) FindAndDelete(ctx context.Context, body AISecretFindAndDeleteParams, opts ...option.RequestOption) (res *shared.AISecret, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/ai_secret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Create or replace ai_secret. If there is an existing ai_secret with the same
// name as the one specified in the request, will replace the existing ai_secret
// with the provided fields
func (r *AISecretService) Replace(ctx context.Context, body AISecretReplaceParams, opts ...option.RequestOption) (res *shared.AISecret, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/ai_secret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AISecretNewParams struct {
	// Name of the AI secret
	Name     param.Field[string]                 `json:"name,required"`
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the AI Secret belongs in.
	OrgName param.Field[string] `json:"org_name"`
	// Secret value. If omitted in a PUT request, the existing secret value will be
	// left intact, not replaced with null.
	Secret param.Field[string] `json:"secret"`
	Type   param.Field[string] `json:"type"`
}

func (r AISecretNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AISecretUpdateParams struct {
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// Name of the AI secret
	Name   param.Field[string] `json:"name"`
	Secret param.Field[string] `json:"secret"`
	Type   param.Field[string] `json:"type"`
}

func (r AISecretUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AISecretListParams struct {
	// Name of the ai_secret to search for
	AISecretName param.Field[string]                              `query:"ai_secret_name"`
	AISecretType param.Field[AISecretListParamsAISecretTypeUnion] `query:"ai_secret_type"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[AISecretListParamsIDsUnion] `query:"ids" format:"uuid"`
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

// URLQuery serializes [AISecretListParams]'s query parameters as `url.Values`.
func (r AISecretListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Satisfied by [shared.UnionString], [AISecretListParamsAISecretTypeArray].
type AISecretListParamsAISecretTypeUnion interface {
	ImplementsAISecretListParamsAISecretTypeUnion()
}

type AISecretListParamsAISecretTypeArray []string

func (r AISecretListParamsAISecretTypeArray) ImplementsAISecretListParamsAISecretTypeUnion() {}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [AISecretListParamsIDsArray].
type AISecretListParamsIDsUnion interface {
	ImplementsAISecretListParamsIDsUnion()
}

type AISecretListParamsIDsArray []string

func (r AISecretListParamsIDsArray) ImplementsAISecretListParamsIDsUnion() {}

type AISecretFindAndDeleteParams struct {
	// Name of the AI secret
	Name param.Field[string] `json:"name,required"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the AI Secret belongs in.
	OrgName param.Field[string] `json:"org_name"`
}

func (r AISecretFindAndDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AISecretReplaceParams struct {
	// Name of the AI secret
	Name     param.Field[string]                 `json:"name,required"`
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the AI Secret belongs in.
	OrgName param.Field[string] `json:"org_name"`
	// Secret value. If omitted in a PUT request, the existing secret value will be
	// left intact, not replaced with null.
	Secret param.Field[string] `json:"secret"`
	Type   param.Field[string] `json:"type"`
}

func (r AISecretReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
