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

// OrganizationService contains methods and other services that help with
// interacting with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	return
}

// Get a organization object by its id
func (r *OrganizationService) Get(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *Organization, err error) {
	opts = append(r.Options[:], opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("v1/organization/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update a organization object. Specify the fields to update in the
// payload. Any object-type fields will be deep-merged with existing content.
// Currently we do not support removing fields or setting them to null.
func (r *OrganizationService) Update(ctx context.Context, organizationID string, body OrganizationUpdateParams, opts ...option.RequestOption) (res *Organization, err error) {
	opts = append(r.Options[:], opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("v1/organization/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all organizations. The organizations are sorted by creation date, with
// the most recently-created organizations coming first
func (r *OrganizationService) List(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) (res *pagination.ListObjects[Organization], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/organization"
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

// List out all organizations. The organizations are sorted by creation date, with
// the most recently-created organizations coming first
func (r *OrganizationService) ListAutoPaging(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[Organization] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a organization object by its id
func (r *OrganizationService) Delete(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *Organization, err error) {
	opts = append(r.Options[:], opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("v1/organization/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Organization struct {
	// Unique identifier for the organization
	ID string `json:"id,required" format:"uuid"`
	// Name of the organization
	Name   string `json:"name,required"`
	APIURL string `json:"api_url,nullable"`
	// Date of organization creation
	Created     time.Time        `json:"created,nullable" format:"date-time"`
	ProxyURL    string           `json:"proxy_url,nullable"`
	RealtimeURL string           `json:"realtime_url,nullable"`
	JSON        organizationJSON `json:"-"`
}

// organizationJSON contains the JSON metadata for the struct [Organization]
type organizationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	APIURL      apijson.Field
	Created     apijson.Field
	ProxyURL    apijson.Field
	RealtimeURL apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Organization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateParams struct {
	APIURL param.Field[string] `json:"api_url"`
	// Name of the organization
	Name        param.Field[string] `json:"name"`
	ProxyURL    param.Field[string] `json:"proxy_url"`
	RealtimeURL param.Field[string] `json:"realtime_url"`
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationListParams struct {
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[OrganizationListParamsIDsUnion] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[int64] `query:"limit"`
	// Filter search results to within a particular organization
	OrgName param.Field[string] `query:"org_name"`
	// Name of the organization to search for
	OrganizationName param.Field[string] `query:"organization_name"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [OrganizationListParams]'s query parameters as `url.Values`.
func (r OrganizationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [OrganizationListParamsIDsArray].
type OrganizationListParamsIDsUnion interface {
	ImplementsOrganizationListParamsIDsUnion()
}

type OrganizationListParamsIDsArray []string

func (r OrganizationListParamsIDsArray) ImplementsOrganizationListParamsIDsUnion() {}