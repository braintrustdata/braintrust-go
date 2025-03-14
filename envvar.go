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
	"github.com/braintrustdata/braintrust-go/shared"
)

// EnvVarService contains methods and other services that help with interacting
// with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvVarService] method instead.
type EnvVarService struct {
	Options []option.RequestOption
}

// NewEnvVarService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEnvVarService(opts ...option.RequestOption) (r *EnvVarService) {
	r = &EnvVarService{}
	r.Options = opts
	return
}

// Create a new env_var. If there is an existing env_var with the same name as the
// one specified in the request, will return the existing env_var unmodified
func (r *EnvVarService) New(ctx context.Context, body EnvVarNewParams, opts ...option.RequestOption) (res *shared.EnvVar, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/env_var"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an env_var object by its id
func (r *EnvVarService) Get(ctx context.Context, envVarID string, opts ...option.RequestOption) (res *shared.EnvVar, err error) {
	opts = append(r.Options[:], opts...)
	if envVarID == "" {
		err = errors.New("missing required env_var_id parameter")
		return
	}
	path := fmt.Sprintf("v1/env_var/%s", envVarID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update an env_var object. Specify the fields to update in the payload.
// Any object-type fields will be deep-merged with existing content. Currently we
// do not support removing fields or setting them to null.
func (r *EnvVarService) Update(ctx context.Context, envVarID string, body EnvVarUpdateParams, opts ...option.RequestOption) (res *shared.EnvVar, err error) {
	opts = append(r.Options[:], opts...)
	if envVarID == "" {
		err = errors.New("missing required env_var_id parameter")
		return
	}
	path := fmt.Sprintf("v1/env_var/%s", envVarID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all env_vars. The env_vars are sorted by creation date, with the most
// recently-created env_vars coming first
func (r *EnvVarService) List(ctx context.Context, query EnvVarListParams, opts ...option.RequestOption) (res *EnvVarListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/env_var"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an env_var object by its id
func (r *EnvVarService) Delete(ctx context.Context, envVarID string, opts ...option.RequestOption) (res *shared.EnvVar, err error) {
	opts = append(r.Options[:], opts...)
	if envVarID == "" {
		err = errors.New("missing required env_var_id parameter")
		return
	}
	path := fmt.Sprintf("v1/env_var/%s", envVarID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create or replace env_var. If there is an existing env_var with the same name as
// the one specified in the request, will replace the existing env_var with the
// provided fields
func (r *EnvVarService) Replace(ctx context.Context, body EnvVarReplaceParams, opts ...option.RequestOption) (res *shared.EnvVar, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/env_var"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type EnvVarListResponse struct {
	// A list of env_var objects
	Objects []shared.EnvVar        `json:"objects,required"`
	JSON    envVarListResponseJSON `json:"-"`
}

// envVarListResponseJSON contains the JSON metadata for the struct
// [EnvVarListResponse]
type envVarListResponseJSON struct {
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvVarListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r envVarListResponseJSON) RawJSON() string {
	return r.raw
}

type EnvVarNewParams struct {
	// The name of the environment variable
	Name param.Field[string] `json:"name,required"`
	// The id of the object the environment variable is scoped for
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The type of the object the environment variable is scoped for
	ObjectType param.Field[EnvVarNewParamsObjectType] `json:"object_type,required"`
	// The value of the environment variable. Will be encrypted at rest.
	Value param.Field[string] `json:"value"`
}

func (r EnvVarNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the object the environment variable is scoped for
type EnvVarNewParamsObjectType string

const (
	EnvVarNewParamsObjectTypeOrganization EnvVarNewParamsObjectType = "organization"
	EnvVarNewParamsObjectTypeProject      EnvVarNewParamsObjectType = "project"
	EnvVarNewParamsObjectTypeFunction     EnvVarNewParamsObjectType = "function"
)

func (r EnvVarNewParamsObjectType) IsKnown() bool {
	switch r {
	case EnvVarNewParamsObjectTypeOrganization, EnvVarNewParamsObjectTypeProject, EnvVarNewParamsObjectTypeFunction:
		return true
	}
	return false
}

type EnvVarUpdateParams struct {
	// The name of the environment variable
	Name param.Field[string] `json:"name,required"`
	// The value of the environment variable. Will be encrypted at rest.
	Value param.Field[string] `json:"value"`
}

func (r EnvVarUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvVarListParams struct {
	// Name of the env_var to search for
	EnvVarName param.Field[string] `query:"env_var_name"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[EnvVarListParamsIDsUnion] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[int64] `query:"limit"`
	// The id of the object the environment variable is scoped for
	ObjectID param.Field[string] `query:"object_id" format:"uuid"`
	// The type of the object the environment variable is scoped for
	ObjectType param.Field[shared.EnvVarObjectType] `query:"object_type"`
}

// URLQuery serializes [EnvVarListParams]'s query parameters as `url.Values`.
func (r EnvVarListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [EnvVarListParamsIDsArray].
type EnvVarListParamsIDsUnion interface {
	ImplementsEnvVarListParamsIDsUnion()
}

type EnvVarListParamsIDsArray []string

func (r EnvVarListParamsIDsArray) ImplementsEnvVarListParamsIDsUnion() {}

type EnvVarReplaceParams struct {
	// The name of the environment variable
	Name param.Field[string] `json:"name,required"`
	// The id of the object the environment variable is scoped for
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The type of the object the environment variable is scoped for
	ObjectType param.Field[EnvVarReplaceParamsObjectType] `json:"object_type,required"`
	// The value of the environment variable. Will be encrypted at rest.
	Value param.Field[string] `json:"value"`
}

func (r EnvVarReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the object the environment variable is scoped for
type EnvVarReplaceParamsObjectType string

const (
	EnvVarReplaceParamsObjectTypeOrganization EnvVarReplaceParamsObjectType = "organization"
	EnvVarReplaceParamsObjectTypeProject      EnvVarReplaceParamsObjectType = "project"
	EnvVarReplaceParamsObjectTypeFunction     EnvVarReplaceParamsObjectType = "function"
)

func (r EnvVarReplaceParamsObjectType) IsKnown() bool {
	switch r {
	case EnvVarReplaceParamsObjectTypeOrganization, EnvVarReplaceParamsObjectTypeProject, EnvVarReplaceParamsObjectTypeFunction:
		return true
	}
	return false
}
