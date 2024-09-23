// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/braintrustdata/braintrust-go/internal/apiquery"
	"github.com/braintrustdata/braintrust-go/internal/pagination"
	"github.com/braintrustdata/braintrust-go/internal/param"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/braintrustdata/braintrust-go/shared"
)

// UserService contains methods and other services that help with interacting with
// the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	return
}

// Get a user object by its id
func (r *UserService) Get(ctx context.Context, userID string, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/user/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List out all users. The users are sorted by creation date, with the most
// recently-created users coming first
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.User], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/user"
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

// List out all users. The users are sorted by creation date, with the most
// recently-created users coming first
func (r *UserService) ListAutoPaging(ctx context.Context, query UserListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.User] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

type UserListParams struct {
	// Email of the user to search for. You may pass the param multiple times to filter
	// for more than one email
	Email param.Field[UserListParamsEmailUnion] `query:"email"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Family name of the user to search for. You may pass the param multiple times to
	// filter for more than one family name
	FamilyName param.Field[UserListParamsFamilyNameUnion] `query:"family_name"`
	// Given name of the user to search for. You may pass the param multiple times to
	// filter for more than one given name
	GivenName param.Field[UserListParamsGivenNameUnion] `query:"given_name"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[UserListParamsIDsUnion] `query:"ids" format:"uuid"`
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

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Email of the user to search for. You may pass the param multiple times to filter
// for more than one email
//
// Satisfied by [shared.UnionString], [UserListParamsEmailEmailMultiple].
type UserListParamsEmailUnion interface {
	ImplementsUserListParamsEmailUnion()
}

type UserListParamsEmailEmailMultiple []string

func (r UserListParamsEmailEmailMultiple) ImplementsUserListParamsEmailUnion() {}

// Family name of the user to search for. You may pass the param multiple times to
// filter for more than one family name
//
// Satisfied by [shared.UnionString], [UserListParamsFamilyNameFamilyNameMultiple].
type UserListParamsFamilyNameUnion interface {
	ImplementsUserListParamsFamilyNameUnion()
}

type UserListParamsFamilyNameFamilyNameMultiple []string

func (r UserListParamsFamilyNameFamilyNameMultiple) ImplementsUserListParamsFamilyNameUnion() {}

// Given name of the user to search for. You may pass the param multiple times to
// filter for more than one given name
//
// Satisfied by [shared.UnionString], [UserListParamsGivenNameGivenNameMultiple].
type UserListParamsGivenNameUnion interface {
	ImplementsUserListParamsGivenNameUnion()
}

type UserListParamsGivenNameGivenNameMultiple []string

func (r UserListParamsGivenNameGivenNameMultiple) ImplementsUserListParamsGivenNameUnion() {}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [UserListParamsIDsArray].
type UserListParamsIDsUnion interface {
	ImplementsUserListParamsIDsUnion()
}

type UserListParamsIDsArray []string

func (r UserListParamsIDsArray) ImplementsUserListParamsIDsUnion() {}
