// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
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

// RoleService contains methods and other services that help with interacting with
// the braintrust API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewRoleService] method instead.
type RoleService struct {
	Options []option.RequestOption
}

// NewRoleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRoleService(opts ...option.RequestOption) (r *RoleService) {
	r = &RoleService{}
	r.Options = opts
	return
}

// Create a new role. If there is an existing role with the same name as the one
// specified in the request, will return the existing role unmodified
func (r *RoleService) New(ctx context.Context, body RoleNewParams, opts ...option.RequestOption) (res *Role, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/role"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a role object by its id
func (r *RoleService) Get(ctx context.Context, roleID string, opts ...option.RequestOption) (res *Role, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/role/%s", roleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update a role object. Specify the fields to update in the payload. Any
// object-type fields will be deep-merged with existing content. Currently we do
// not support removing fields or setting them to null.
func (r *RoleService) Update(ctx context.Context, roleID string, body RoleUpdateParams, opts ...option.RequestOption) (res *Role, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/role/%s", roleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all roles. The roles are sorted by creation date, with the most
// recently-created roles coming first
func (r *RoleService) List(ctx context.Context, query RoleListParams, opts ...option.RequestOption) (res *pagination.ListObjects[Role], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *RoleService) ListAutoPaging(ctx context.Context, query RoleListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[Role] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a role object by its id
func (r *RoleService) Delete(ctx context.Context, roleID string, opts ...option.RequestOption) (res *Role, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/role/%s", roleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// NOTE: This operation is deprecated and will be removed in a future revision of
// the API. Create or replace a new role. If there is an existing role with the
// same name as the one specified in the request, will return the existing role
// unmodified, will replace the existing role with the provided fields
func (r *RoleService) Replace(ctx context.Context, body RoleReplaceParams, opts ...option.RequestOption) (res *Role, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/role"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// A role is a collection of permissions which can be granted as part of an ACL
//
// Roles can consist of individual permissions, as well as a set of roles they
// inherit from
type Role struct {
	// Unique identifier for the role
	ID string `json:"id,required" format:"uuid"`
	// Name of the role
	Name string `json:"name,required"`
	// Date of role creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Date of role deletion, or null if the role is still active
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Textual description of the role
	Description string `json:"description,nullable"`
	// Permissions which belong to this role
	MemberPermissions []RoleMemberPermission `json:"member_permissions,nullable"`
	// Ids of the roles this role inherits from
	//
	// An inheriting role has all the permissions contained in its member roles, as
	// well as all of their inherited permissions
	MemberRoles []string `json:"member_roles,nullable" format:"uuid"`
	// Unique id for the organization that the role belongs under
	//
	// A null org_id indicates a system role, which may be assigned to anybody and
	// inherited by any other role, but cannot be edited.
	//
	// It is forbidden to change the org after creating a role
	OrgID string `json:"org_id,nullable" format:"uuid"`
	// Identifies the user who created the role
	UserID string   `json:"user_id,nullable" format:"uuid"`
	JSON   roleJSON `json:"-"`
}

// roleJSON contains the JSON metadata for the struct [Role]
type roleJSON struct {
	ID                apijson.Field
	Name              apijson.Field
	Created           apijson.Field
	DeletedAt         apijson.Field
	Description       apijson.Field
	MemberPermissions apijson.Field
	MemberRoles       apijson.Field
	OrgID             apijson.Field
	UserID            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Role) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleJSON) RawJSON() string {
	return r.raw
}

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
type RoleMemberPermission string

const (
	RoleMemberPermissionCreate     RoleMemberPermission = "create"
	RoleMemberPermissionRead       RoleMemberPermission = "read"
	RoleMemberPermissionUpdate     RoleMemberPermission = "update"
	RoleMemberPermissionDelete     RoleMemberPermission = "delete"
	RoleMemberPermissionCreateACLs RoleMemberPermission = "create_acls"
	RoleMemberPermissionReadACLs   RoleMemberPermission = "read_acls"
	RoleMemberPermissionUpdateACLs RoleMemberPermission = "update_acls"
	RoleMemberPermissionDeleteACLs RoleMemberPermission = "delete_acls"
)

func (r RoleMemberPermission) IsKnown() bool {
	switch r {
	case RoleMemberPermissionCreate, RoleMemberPermissionRead, RoleMemberPermissionUpdate, RoleMemberPermissionDelete, RoleMemberPermissionCreateACLs, RoleMemberPermissionReadACLs, RoleMemberPermissionUpdateACLs, RoleMemberPermissionDeleteACLs:
		return true
	}
	return false
}

type RoleNewParams struct {
	// Name of the role
	Name param.Field[string] `json:"name,required"`
	// Textual description of the role
	Description param.Field[string] `json:"description"`
	// Permissions which belong to this role
	MemberPermissions param.Field[[]RoleNewParamsMemberPermission] `json:"member_permissions"`
	// Ids of the roles this role inherits from
	//
	// An inheriting role has all the permissions contained in its member roles, as
	// well as all of their inherited permissions
	MemberRoles param.Field[[]string] `json:"member_roles" format:"uuid"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the role belongs in.
	OrgName param.Field[string] `json:"org_name"`
}

func (r RoleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
type RoleNewParamsMemberPermission string

const (
	RoleNewParamsMemberPermissionCreate     RoleNewParamsMemberPermission = "create"
	RoleNewParamsMemberPermissionRead       RoleNewParamsMemberPermission = "read"
	RoleNewParamsMemberPermissionUpdate     RoleNewParamsMemberPermission = "update"
	RoleNewParamsMemberPermissionDelete     RoleNewParamsMemberPermission = "delete"
	RoleNewParamsMemberPermissionCreateACLs RoleNewParamsMemberPermission = "create_acls"
	RoleNewParamsMemberPermissionReadACLs   RoleNewParamsMemberPermission = "read_acls"
	RoleNewParamsMemberPermissionUpdateACLs RoleNewParamsMemberPermission = "update_acls"
	RoleNewParamsMemberPermissionDeleteACLs RoleNewParamsMemberPermission = "delete_acls"
)

func (r RoleNewParamsMemberPermission) IsKnown() bool {
	switch r {
	case RoleNewParamsMemberPermissionCreate, RoleNewParamsMemberPermissionRead, RoleNewParamsMemberPermissionUpdate, RoleNewParamsMemberPermissionDelete, RoleNewParamsMemberPermissionCreateACLs, RoleNewParamsMemberPermissionReadACLs, RoleNewParamsMemberPermissionUpdateACLs, RoleNewParamsMemberPermissionDeleteACLs:
		return true
	}
	return false
}

type RoleUpdateParams struct {
	// Textual description of the role
	Description param.Field[string] `json:"description"`
	// Permissions which belong to this role
	MemberPermissions param.Field[[]RoleUpdateParamsMemberPermission] `json:"member_permissions"`
	// Ids of the roles this role inherits from
	//
	// An inheriting role has all the permissions contained in its member roles, as
	// well as all of their inherited permissions
	MemberRoles param.Field[[]string] `json:"member_roles" format:"uuid"`
	// Name of the role
	Name param.Field[string] `json:"name"`
}

func (r RoleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
type RoleUpdateParamsMemberPermission string

const (
	RoleUpdateParamsMemberPermissionCreate     RoleUpdateParamsMemberPermission = "create"
	RoleUpdateParamsMemberPermissionRead       RoleUpdateParamsMemberPermission = "read"
	RoleUpdateParamsMemberPermissionUpdate     RoleUpdateParamsMemberPermission = "update"
	RoleUpdateParamsMemberPermissionDelete     RoleUpdateParamsMemberPermission = "delete"
	RoleUpdateParamsMemberPermissionCreateACLs RoleUpdateParamsMemberPermission = "create_acls"
	RoleUpdateParamsMemberPermissionReadACLs   RoleUpdateParamsMemberPermission = "read_acls"
	RoleUpdateParamsMemberPermissionUpdateACLs RoleUpdateParamsMemberPermission = "update_acls"
	RoleUpdateParamsMemberPermissionDeleteACLs RoleUpdateParamsMemberPermission = "delete_acls"
)

func (r RoleUpdateParamsMemberPermission) IsKnown() bool {
	switch r {
	case RoleUpdateParamsMemberPermissionCreate, RoleUpdateParamsMemberPermissionRead, RoleUpdateParamsMemberPermissionUpdate, RoleUpdateParamsMemberPermissionDelete, RoleUpdateParamsMemberPermissionCreateACLs, RoleUpdateParamsMemberPermissionReadACLs, RoleUpdateParamsMemberPermissionUpdateACLs, RoleUpdateParamsMemberPermissionDeleteACLs:
		return true
	}
	return false
}

type RoleListParams struct {
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[RoleListParamsIDsUnion] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[int64] `query:"limit"`
	// Filter search results to within a particular organization
	OrgName param.Field[string] `query:"org_name"`
	// Name of the role to search for
	RoleName param.Field[string] `query:"role_name"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [RoleListParams]'s query parameters as `url.Values`.
func (r RoleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [RoleListParamsIDsArray].
type RoleListParamsIDsUnion interface {
	ImplementsRoleListParamsIDsUnion()
}

type RoleListParamsIDsArray []string

func (r RoleListParamsIDsArray) ImplementsRoleListParamsIDsUnion() {}

type RoleReplaceParams struct {
	// Name of the role
	Name param.Field[string] `json:"name,required"`
	// Textual description of the role
	Description param.Field[string] `json:"description"`
	// Permissions which belong to this role
	MemberPermissions param.Field[[]RoleReplaceParamsMemberPermission] `json:"member_permissions"`
	// Ids of the roles this role inherits from
	//
	// An inheriting role has all the permissions contained in its member roles, as
	// well as all of their inherited permissions
	MemberRoles param.Field[[]string] `json:"member_roles" format:"uuid"`
	// For nearly all users, this parameter should be unnecessary. But in the rare case
	// that your API key belongs to multiple organizations, you may specify the name of
	// the organization the role belongs in.
	OrgName param.Field[string] `json:"org_name"`
}

func (r RoleReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
type RoleReplaceParamsMemberPermission string

const (
	RoleReplaceParamsMemberPermissionCreate     RoleReplaceParamsMemberPermission = "create"
	RoleReplaceParamsMemberPermissionRead       RoleReplaceParamsMemberPermission = "read"
	RoleReplaceParamsMemberPermissionUpdate     RoleReplaceParamsMemberPermission = "update"
	RoleReplaceParamsMemberPermissionDelete     RoleReplaceParamsMemberPermission = "delete"
	RoleReplaceParamsMemberPermissionCreateACLs RoleReplaceParamsMemberPermission = "create_acls"
	RoleReplaceParamsMemberPermissionReadACLs   RoleReplaceParamsMemberPermission = "read_acls"
	RoleReplaceParamsMemberPermissionUpdateACLs RoleReplaceParamsMemberPermission = "update_acls"
	RoleReplaceParamsMemberPermissionDeleteACLs RoleReplaceParamsMemberPermission = "delete_acls"
)

func (r RoleReplaceParamsMemberPermission) IsKnown() bool {
	switch r {
	case RoleReplaceParamsMemberPermissionCreate, RoleReplaceParamsMemberPermissionRead, RoleReplaceParamsMemberPermissionUpdate, RoleReplaceParamsMemberPermissionDelete, RoleReplaceParamsMemberPermissionCreateACLs, RoleReplaceParamsMemberPermissionReadACLs, RoleReplaceParamsMemberPermissionUpdateACLs, RoleReplaceParamsMemberPermissionDeleteACLs:
		return true
	}
	return false
}
