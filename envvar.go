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
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/braintrustdata/braintrust-go/packages/param"
	"github.com/braintrustdata/braintrust-go/packages/respjson"
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
func NewEnvVarService(opts ...option.RequestOption) (r EnvVarService) {
	r = EnvVarService{}
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
	Objects []shared.EnvVar `json:"objects,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Objects     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnvVarListResponse) RawJSON() string { return r.JSON.raw }
func (r *EnvVarListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnvVarNewParams struct {
	// The name of the environment variable
	Name string `json:"name,required"`
	// The id of the object the environment variable is scoped for
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The type of the object the environment variable is scoped for
	//
	// Any of "organization", "project", "function".
	ObjectType EnvVarNewParamsObjectType `json:"object_type,omitzero,required"`
	// The value of the environment variable. Will be encrypted at rest.
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r EnvVarNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EnvVarNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// The type of the object the environment variable is scoped for
type EnvVarNewParamsObjectType string

const (
	EnvVarNewParamsObjectTypeOrganization EnvVarNewParamsObjectType = "organization"
	EnvVarNewParamsObjectTypeProject      EnvVarNewParamsObjectType = "project"
	EnvVarNewParamsObjectTypeFunction     EnvVarNewParamsObjectType = "function"
)

type EnvVarUpdateParams struct {
	// The name of the environment variable
	Name string `json:"name,required"`
	// The value of the environment variable. Will be encrypted at rest.
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r EnvVarUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EnvVarUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type EnvVarListParams struct {
	// Limit the number of objects to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Name of the env_var to search for
	EnvVarName param.Opt[string] `query:"env_var_name,omitzero" json:"-"`
	// The id of the object the environment variable is scoped for
	ObjectID param.Opt[string] `query:"object_id,omitzero" format:"uuid" json:"-"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs EnvVarListParamsIDsUnion `query:"ids,omitzero" format:"uuid" json:"-"`
	// The type of the object the environment variable is scoped for
	//
	// Any of "organization", "project", "function".
	ObjectType shared.EnvVarObjectType `query:"object_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EnvVarListParams]'s query parameters as `url.Values`.
func (r EnvVarListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EnvVarListParamsIDsUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *EnvVarListParamsIDsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type EnvVarReplaceParams struct {
	// The name of the environment variable
	Name string `json:"name,required"`
	// The id of the object the environment variable is scoped for
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The type of the object the environment variable is scoped for
	//
	// Any of "organization", "project", "function".
	ObjectType EnvVarReplaceParamsObjectType `json:"object_type,omitzero,required"`
	// The value of the environment variable. Will be encrypted at rest.
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r EnvVarReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow EnvVarReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// The type of the object the environment variable is scoped for
type EnvVarReplaceParamsObjectType string

const (
	EnvVarReplaceParamsObjectTypeOrganization EnvVarReplaceParamsObjectType = "organization"
	EnvVarReplaceParamsObjectTypeProject      EnvVarReplaceParamsObjectType = "project"
	EnvVarReplaceParamsObjectTypeFunction     EnvVarReplaceParamsObjectType = "function"
)
