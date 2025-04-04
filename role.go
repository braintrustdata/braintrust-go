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

// RoleService contains methods and other services that help with interacting with
// the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoleService] method instead.
type RoleService struct {
	Options []option.RequestOption
}

// NewRoleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRoleService(opts ...option.RequestOption) (r RoleService) {
	r = RoleService{}
	r.Options = opts
	return
}

// Create a new role. If there is an existing role with the same name as the one
// specified in the request, will return the existing role unmodified
func (r *RoleService) New(ctx context.Context, body RoleNewParams, opts ...option.RequestOption) (res *shared.Role, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/role"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a role object by its id
func (r *RoleService) Get(ctx context.Context, roleID string, opts ...option.RequestOption) (res *shared.Role, err error) {
	opts = append(r.Options[:], opts...)
	if roleID == "" {
		err = errors.New("missing required role_id parameter")
		return
	}
	path := fmt.Sprintf("v1/role/%s", roleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update a role object. Specify the fields to update in the payload. Any
// object-type fields will be deep-merged with existing content. Currently we do
// not support removing fields or setting them to null.
func (r *RoleService) Update(ctx context.Context, roleID string, body RoleUpdateParams, opts ...option.RequestOption) (res *shared.Role, err error) {
	opts = append(r.Options[:], opts...)
	if roleID == "" {
		err = errors.New("missing required role_id parameter")
		return
	}
	path := fmt.Sprintf("v1/role/%s", roleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all roles. The roles are sorted by creation date, with the most
// recently-created roles coming first
func (r *RoleService) List(ctx context.Context, query RoleListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.Role], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/role"
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

// List out all roles. The roles are sorted by creation date, with the most
// recently-created roles coming first
func (r *RoleService) ListAutoPaging(ctx context.Context, query RoleListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.Role] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a role object by its id
func (r *RoleService) Delete(ctx context.Context, roleID string, opts ...option.RequestOption) (res *shared.Role, err error) {
	opts = append(r.Options[:], opts...)
	if roleID == "" {
		err = errors.New("missing required role_id parameter")
		return
	}
	path := fmt.Sprintf("v1/role/%s", roleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create or replace role. If there is an existing role with the same name as the
// one specified in the request, will replace the existing role with the provided
// fields
func (r *RoleService) Replace(ctx context.Context, body RoleReplaceParams, opts ...option.RequestOption) (res *shared.Role, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/role"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type RoleNewParams struct {
	// Name of the role
	Name string `json:"name,required"`
	// Textual description of the role
	Description param.Opt[string] `json:"description,omitzero"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the role belongs in.
	OrgName param.Opt[string] `json:"org_name,omitzero"`
	// (permission, restrict_object_type) tuples which belong to this role
	MemberPermissions []RoleNewParamsMemberPermission `json:"member_permissions,omitzero"`
	// Ids of the roles this role inherits from
	//
	// An inheriting role has all the permissions contained in its member roles, as
	// well as all of their inherited permissions
	MemberRoles []string `json:"member_roles,omitzero" format:"uuid"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RoleNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r RoleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RoleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property Permission is required.
type RoleNewParamsMemberPermission struct {
	// Each permission permits a certain type of operation on an object in the system
	//
	// Permissions can be assigned to to objects on an individual basis, or grouped
	// into roles
	//
	// Any of "create", "read", "update", "delete", "create_acls", "read_acls",
	// "update_acls", "delete_acls".
	Permission shared.Permission `json:"permission,omitzero,required"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	RestrictObjectType shared.ACLObjectType `json:"restrict_object_type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RoleNewParamsMemberPermission) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r RoleNewParamsMemberPermission) MarshalJSON() (data []byte, err error) {
	type shadow RoleNewParamsMemberPermission
	return param.MarshalObject(r, (*shadow)(&r))
}

type RoleUpdateParams struct {
	// Textual description of the role
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the role
	Name param.Opt[string] `json:"name,omitzero"`
	// A list of permissions to add to the role
	AddMemberPermissions []RoleUpdateParamsAddMemberPermission `json:"add_member_permissions,omitzero"`
	// A list of role IDs to add to the role's inheriting-from set
	AddMemberRoles []string `json:"add_member_roles,omitzero" format:"uuid"`
	// A list of permissions to remove from the role
	RemoveMemberPermissions []RoleUpdateParamsRemoveMemberPermission `json:"remove_member_permissions,omitzero"`
	// A list of role IDs to remove from the role's inheriting-from set
	RemoveMemberRoles []string `json:"remove_member_roles,omitzero" format:"uuid"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RoleUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r RoleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RoleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property Permission is required.
type RoleUpdateParamsAddMemberPermission struct {
	// Each permission permits a certain type of operation on an object in the system
	//
	// Permissions can be assigned to to objects on an individual basis, or grouped
	// into roles
	//
	// Any of "create", "read", "update", "delete", "create_acls", "read_acls",
	// "update_acls", "delete_acls".
	Permission shared.Permission `json:"permission,omitzero,required"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	RestrictObjectType shared.ACLObjectType `json:"restrict_object_type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RoleUpdateParamsAddMemberPermission) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r RoleUpdateParamsAddMemberPermission) MarshalJSON() (data []byte, err error) {
	type shadow RoleUpdateParamsAddMemberPermission
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property Permission is required.
type RoleUpdateParamsRemoveMemberPermission struct {
	// Each permission permits a certain type of operation on an object in the system
	//
	// Permissions can be assigned to to objects on an individual basis, or grouped
	// into roles
	//
	// Any of "create", "read", "update", "delete", "create_acls", "read_acls",
	// "update_acls", "delete_acls".
	Permission shared.Permission `json:"permission,omitzero,required"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	RestrictObjectType shared.ACLObjectType `json:"restrict_object_type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RoleUpdateParamsRemoveMemberPermission) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r RoleUpdateParamsRemoveMemberPermission) MarshalJSON() (data []byte, err error) {
	type shadow RoleUpdateParamsRemoveMemberPermission
	return param.MarshalObject(r, (*shadow)(&r))
}

type RoleListParams struct {
	// Limit the number of objects to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Opt[string] `query:"ending_before,omitzero" format:"uuid" json:"-"`
	// Filter search results to within a particular organization
	OrgName param.Opt[string] `query:"org_name,omitzero" json:"-"`
	// Name of the role to search for
	RoleName param.Opt[string] `query:"role_name,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Opt[string] `query:"starting_after,omitzero" format:"uuid" json:"-"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs RoleListParamsIDsUnion `query:"ids,omitzero" format:"uuid" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RoleListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [RoleListParams]'s query parameters as `url.Values`.
func (r RoleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RoleListParamsIDsUnion struct {
	OfString            param.Opt[string] `query:",omitzero,inline"`
	OfRoleListsIDsArray []string          `query:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u RoleListParamsIDsUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }

func (u *RoleListParamsIDsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfRoleListsIDsArray) {
		return &u.OfRoleListsIDsArray
	}
	return nil
}

type RoleReplaceParams struct {
	// Name of the role
	Name string `json:"name,required"`
	// Textual description of the role
	Description param.Opt[string] `json:"description,omitzero"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the role belongs in.
	OrgName param.Opt[string] `json:"org_name,omitzero"`
	// (permission, restrict_object_type) tuples which belong to this role
	MemberPermissions []RoleReplaceParamsMemberPermission `json:"member_permissions,omitzero"`
	// Ids of the roles this role inherits from
	//
	// An inheriting role has all the permissions contained in its member roles, as
	// well as all of their inherited permissions
	MemberRoles []string `json:"member_roles,omitzero" format:"uuid"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RoleReplaceParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r RoleReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow RoleReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property Permission is required.
type RoleReplaceParamsMemberPermission struct {
	// Each permission permits a certain type of operation on an object in the system
	//
	// Permissions can be assigned to to objects on an individual basis, or grouped
	// into roles
	//
	// Any of "create", "read", "update", "delete", "create_acls", "read_acls",
	// "update_acls", "delete_acls".
	Permission shared.Permission `json:"permission,omitzero,required"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	RestrictObjectType shared.ACLObjectType `json:"restrict_object_type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RoleReplaceParamsMemberPermission) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r RoleReplaceParamsMemberPermission) MarshalJSON() (data []byte, err error) {
	type shadow RoleReplaceParamsMemberPermission
	return param.MarshalObject(r, (*shadow)(&r))
}
