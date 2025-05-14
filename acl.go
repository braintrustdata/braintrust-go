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
	"github.com/braintrustdata/braintrust-go/packages/pagination"
	"github.com/braintrustdata/braintrust-go/packages/param"
	"github.com/braintrustdata/braintrust-go/shared"
)

// ACLService contains methods and other services that help with interacting with
// the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewACLService] method instead.
type ACLService struct {
	Options []option.RequestOption
}

// NewACLService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewACLService(opts ...option.RequestOption) (r ACLService) {
	r = ACLService{}
	r.Options = opts
	return
}

// Create a new acl. If there is an existing acl with the same contents as the one
// specified in the request, will return the existing acl unmodified
func (r *ACLService) New(ctx context.Context, body ACLNewParams, opts ...option.RequestOption) (res *shared.ACL, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/acl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an acl object by its id
func (r *ACLService) Get(ctx context.Context, aclID string, opts ...option.RequestOption) (res *shared.ACL, err error) {
	opts = append(r.Options[:], opts...)
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("v1/acl/%s", aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List out all acls. The acls are sorted by creation date, with the most
// recently-created acls coming first
func (r *ACLService) List(ctx context.Context, query ACLListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.ACL], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/acl"
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

// List out all acls. The acls are sorted by creation date, with the most
// recently-created acls coming first
func (r *ACLService) ListAutoPaging(ctx context.Context, query ACLListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.ACL] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete an acl object by its id
func (r *ACLService) Delete(ctx context.Context, aclID string, opts ...option.RequestOption) (res *shared.ACL, err error) {
	opts = append(r.Options[:], opts...)
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("v1/acl/%s", aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Batch update acls. This operation is idempotent, so adding acls which already
// exist will have no effect, and removing acls which do not exist will have no
// effect.
func (r *ACLService) BatchUpdate(ctx context.Context, body ACLBatchUpdateParams, opts ...option.RequestOption) (res *shared.ACLBatchUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/acl/batch_update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a single acl
func (r *ACLService) FindAndDelete(ctx context.Context, body ACLFindAndDeleteParams, opts ...option.RequestOption) (res *shared.ACL, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/acl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type ACLNewParams struct {
	// The id of the object the ACL applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType shared.ACLObjectType `json:"object_type,omitzero,required"`
	// Id of the group the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	GroupID param.Opt[string] `json:"group_id,omitzero" format:"uuid"`
	// Id of the role the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	RoleID param.Opt[string] `json:"role_id,omitzero" format:"uuid"`
	// Id of the user the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	UserID param.Opt[string] `json:"user_id,omitzero" format:"uuid"`
	// Permission the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	//
	// Any of "create", "read", "update", "delete", "create_acls", "read_acls",
	// "update_acls", "delete_acls".
	Permission shared.Permission `json:"permission,omitzero"`
	// When setting a permission directly, optionally restricts the permission grant to
	// just the specified object type. Cannot be set alongside a `role_id`.
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	RestrictObjectType shared.ACLObjectType `json:"restrict_object_type,omitzero"`
	paramObj
}

func (r ACLNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ACLNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ACLNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ACLListParams struct {
	// The id of the object the ACL applies to
	ObjectID string `query:"object_id,required" format:"uuid" json:"-"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType shared.ACLObjectType `query:"object_type,omitzero,required" json:"-"`
	// Limit the number of objects to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Opt[string] `query:"ending_before,omitzero" format:"uuid" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Opt[string] `query:"starting_after,omitzero" format:"uuid" json:"-"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs ACLListParamsIDsUnion `query:"ids,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ACLListParams]'s query parameters as `url.Values`.
func (r ACLListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ACLListParamsIDsUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *ACLListParamsIDsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type ACLBatchUpdateParams struct {
	// An ACL grants a certain permission or role to a certain user or group on an
	// object.
	//
	// ACLs are inherited across the object hierarchy. So for example, if a user has
	// read permissions on a project, they will also have read permissions on any
	// experiment, dataset, etc. created within that project.
	//
	// To restrict a grant to a particular sub-object, you may specify
	// `restrict_object_type` in the ACL, as part of a direct permission grant or as
	// part of a role.
	AddACLs []ACLBatchUpdateParamsAddACL `json:"add_acls,omitzero"`
	// An ACL grants a certain permission or role to a certain user or group on an
	// object.
	//
	// ACLs are inherited across the object hierarchy. So for example, if a user has
	// read permissions on a project, they will also have read permissions on any
	// experiment, dataset, etc. created within that project.
	//
	// To restrict a grant to a particular sub-object, you may specify
	// `restrict_object_type` in the ACL, as part of a direct permission grant or as
	// part of a role.
	RemoveACLs []ACLBatchUpdateParamsRemoveACL `json:"remove_acls,omitzero"`
	paramObj
}

func (r ACLBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ACLBatchUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ACLBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An ACL grants a certain permission or role to a certain user or group on an
// object.
//
// ACLs are inherited across the object hierarchy. So for example, if a user has
// read permissions on a project, they will also have read permissions on any
// experiment, dataset, etc. created within that project.
//
// To restrict a grant to a particular sub-object, you may specify
// `restrict_object_type` in the ACL, as part of a direct permission grant or as
// part of a role.
//
// The properties ObjectID, ObjectType are required.
type ACLBatchUpdateParamsAddACL struct {
	// The id of the object the ACL applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType shared.ACLObjectType `json:"object_type,omitzero,required"`
	// Id of the group the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	GroupID param.Opt[string] `json:"group_id,omitzero" format:"uuid"`
	// Id of the role the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	RoleID param.Opt[string] `json:"role_id,omitzero" format:"uuid"`
	// Id of the user the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	UserID param.Opt[string] `json:"user_id,omitzero" format:"uuid"`
	// Permission the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	//
	// Any of "create", "read", "update", "delete", "create_acls", "read_acls",
	// "update_acls", "delete_acls".
	Permission shared.Permission `json:"permission,omitzero"`
	// When setting a permission directly, optionally restricts the permission grant to
	// just the specified object type. Cannot be set alongside a `role_id`.
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	RestrictObjectType shared.ACLObjectType `json:"restrict_object_type,omitzero"`
	paramObj
}

func (r ACLBatchUpdateParamsAddACL) MarshalJSON() (data []byte, err error) {
	type shadow ACLBatchUpdateParamsAddACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ACLBatchUpdateParamsAddACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An ACL grants a certain permission or role to a certain user or group on an
// object.
//
// ACLs are inherited across the object hierarchy. So for example, if a user has
// read permissions on a project, they will also have read permissions on any
// experiment, dataset, etc. created within that project.
//
// To restrict a grant to a particular sub-object, you may specify
// `restrict_object_type` in the ACL, as part of a direct permission grant or as
// part of a role.
//
// The properties ObjectID, ObjectType are required.
type ACLBatchUpdateParamsRemoveACL struct {
	// The id of the object the ACL applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType shared.ACLObjectType `json:"object_type,omitzero,required"`
	// Id of the group the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	GroupID param.Opt[string] `json:"group_id,omitzero" format:"uuid"`
	// Id of the role the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	RoleID param.Opt[string] `json:"role_id,omitzero" format:"uuid"`
	// Id of the user the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	UserID param.Opt[string] `json:"user_id,omitzero" format:"uuid"`
	// Permission the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	//
	// Any of "create", "read", "update", "delete", "create_acls", "read_acls",
	// "update_acls", "delete_acls".
	Permission shared.Permission `json:"permission,omitzero"`
	// When setting a permission directly, optionally restricts the permission grant to
	// just the specified object type. Cannot be set alongside a `role_id`.
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	RestrictObjectType shared.ACLObjectType `json:"restrict_object_type,omitzero"`
	paramObj
}

func (r ACLBatchUpdateParamsRemoveACL) MarshalJSON() (data []byte, err error) {
	type shadow ACLBatchUpdateParamsRemoveACL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ACLBatchUpdateParamsRemoveACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ACLFindAndDeleteParams struct {
	// The id of the object the ACL applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType shared.ACLObjectType `json:"object_type,omitzero,required"`
	// Id of the group the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	GroupID param.Opt[string] `json:"group_id,omitzero" format:"uuid"`
	// Id of the role the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	RoleID param.Opt[string] `json:"role_id,omitzero" format:"uuid"`
	// Id of the user the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	UserID param.Opt[string] `json:"user_id,omitzero" format:"uuid"`
	// Permission the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	//
	// Any of "create", "read", "update", "delete", "create_acls", "read_acls",
	// "update_acls", "delete_acls".
	Permission shared.Permission `json:"permission,omitzero"`
	// When setting a permission directly, optionally restricts the permission grant to
	// just the specified object type. Cannot be set alongside a `role_id`.
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	RestrictObjectType shared.ACLObjectType `json:"restrict_object_type,omitzero"`
	paramObj
}

func (r ACLFindAndDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow ACLFindAndDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ACLFindAndDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
