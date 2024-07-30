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
func (r *RoleService) Update(ctx context.Context, roleID string, body RoleUpdateParams, opts ...option.RequestOption) (res *Role, err error) {
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
func (r *RoleService) List(ctx context.Context, query RoleListParams, opts ...option.RequestOption) (res *pagination.ListObjects[Role], err error) {
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
func (r *RoleService) ListAutoPaging(ctx context.Context, query RoleListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[Role] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a role object by its id
func (r *RoleService) Delete(ctx context.Context, roleID string, opts ...option.RequestOption) (res *Role, err error) {
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
	// (permission, restrict_object_type) tuples which belong to this role
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

type RoleMemberPermission struct {
	// Each permission permits a certain type of operation on an object in the system
	//
	// Permissions can be assigned to to objects on an individual basis, or grouped
	// into roles
	Permission RoleMemberPermissionsPermission `json:"permission,required,nullable"`
	// The object type that the ACL applies to
	RestrictObjectType RoleMemberPermissionsRestrictObjectType `json:"restrict_object_type,nullable"`
	JSON               roleMemberPermissionJSON                `json:"-"`
}

// roleMemberPermissionJSON contains the JSON metadata for the struct
// [RoleMemberPermission]
type roleMemberPermissionJSON struct {
	Permission         apijson.Field
	RestrictObjectType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RoleMemberPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleMemberPermissionJSON) RawJSON() string {
	return r.raw
}

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
type RoleMemberPermissionsPermission string

const (
	RoleMemberPermissionsPermissionCreate     RoleMemberPermissionsPermission = "create"
	RoleMemberPermissionsPermissionRead       RoleMemberPermissionsPermission = "read"
	RoleMemberPermissionsPermissionUpdate     RoleMemberPermissionsPermission = "update"
	RoleMemberPermissionsPermissionDelete     RoleMemberPermissionsPermission = "delete"
	RoleMemberPermissionsPermissionCreateACLs RoleMemberPermissionsPermission = "create_acls"
	RoleMemberPermissionsPermissionReadACLs   RoleMemberPermissionsPermission = "read_acls"
	RoleMemberPermissionsPermissionUpdateACLs RoleMemberPermissionsPermission = "update_acls"
	RoleMemberPermissionsPermissionDeleteACLs RoleMemberPermissionsPermission = "delete_acls"
)

func (r RoleMemberPermissionsPermission) IsKnown() bool {
	switch r {
	case RoleMemberPermissionsPermissionCreate, RoleMemberPermissionsPermissionRead, RoleMemberPermissionsPermissionUpdate, RoleMemberPermissionsPermissionDelete, RoleMemberPermissionsPermissionCreateACLs, RoleMemberPermissionsPermissionReadACLs, RoleMemberPermissionsPermissionUpdateACLs, RoleMemberPermissionsPermissionDeleteACLs:
		return true
	}
	return false
}

// The object type that the ACL applies to
type RoleMemberPermissionsRestrictObjectType string

const (
	RoleMemberPermissionsRestrictObjectTypeOrganization  RoleMemberPermissionsRestrictObjectType = "organization"
	RoleMemberPermissionsRestrictObjectTypeProject       RoleMemberPermissionsRestrictObjectType = "project"
	RoleMemberPermissionsRestrictObjectTypeExperiment    RoleMemberPermissionsRestrictObjectType = "experiment"
	RoleMemberPermissionsRestrictObjectTypeDataset       RoleMemberPermissionsRestrictObjectType = "dataset"
	RoleMemberPermissionsRestrictObjectTypePrompt        RoleMemberPermissionsRestrictObjectType = "prompt"
	RoleMemberPermissionsRestrictObjectTypePromptSession RoleMemberPermissionsRestrictObjectType = "prompt_session"
	RoleMemberPermissionsRestrictObjectTypeGroup         RoleMemberPermissionsRestrictObjectType = "group"
	RoleMemberPermissionsRestrictObjectTypeRole          RoleMemberPermissionsRestrictObjectType = "role"
	RoleMemberPermissionsRestrictObjectTypeOrgMember     RoleMemberPermissionsRestrictObjectType = "org_member"
	RoleMemberPermissionsRestrictObjectTypeProjectLog    RoleMemberPermissionsRestrictObjectType = "project_log"
	RoleMemberPermissionsRestrictObjectTypeOrgProject    RoleMemberPermissionsRestrictObjectType = "org_project"
)

func (r RoleMemberPermissionsRestrictObjectType) IsKnown() bool {
	switch r {
	case RoleMemberPermissionsRestrictObjectTypeOrganization, RoleMemberPermissionsRestrictObjectTypeProject, RoleMemberPermissionsRestrictObjectTypeExperiment, RoleMemberPermissionsRestrictObjectTypeDataset, RoleMemberPermissionsRestrictObjectTypePrompt, RoleMemberPermissionsRestrictObjectTypePromptSession, RoleMemberPermissionsRestrictObjectTypeGroup, RoleMemberPermissionsRestrictObjectTypeRole, RoleMemberPermissionsRestrictObjectTypeOrgMember, RoleMemberPermissionsRestrictObjectTypeProjectLog, RoleMemberPermissionsRestrictObjectTypeOrgProject:
		return true
	}
	return false
}

type RoleNewParams struct {
	// Name of the role
	Name param.Field[string] `json:"name,required"`
	// Textual description of the role
	Description param.Field[string] `json:"description"`
	// (permission, restrict_object_type) tuples which belong to this role
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

type RoleNewParamsMemberPermission struct {
	// Each permission permits a certain type of operation on an object in the system
	//
	// Permissions can be assigned to to objects on an individual basis, or grouped
	// into roles
	Permission param.Field[RoleNewParamsMemberPermissionsPermission] `json:"permission,required"`
	// The object type that the ACL applies to
	RestrictObjectType param.Field[RoleNewParamsMemberPermissionsRestrictObjectType] `json:"restrict_object_type"`
}

func (r RoleNewParamsMemberPermission) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
type RoleNewParamsMemberPermissionsPermission string

const (
	RoleNewParamsMemberPermissionsPermissionCreate     RoleNewParamsMemberPermissionsPermission = "create"
	RoleNewParamsMemberPermissionsPermissionRead       RoleNewParamsMemberPermissionsPermission = "read"
	RoleNewParamsMemberPermissionsPermissionUpdate     RoleNewParamsMemberPermissionsPermission = "update"
	RoleNewParamsMemberPermissionsPermissionDelete     RoleNewParamsMemberPermissionsPermission = "delete"
	RoleNewParamsMemberPermissionsPermissionCreateACLs RoleNewParamsMemberPermissionsPermission = "create_acls"
	RoleNewParamsMemberPermissionsPermissionReadACLs   RoleNewParamsMemberPermissionsPermission = "read_acls"
	RoleNewParamsMemberPermissionsPermissionUpdateACLs RoleNewParamsMemberPermissionsPermission = "update_acls"
	RoleNewParamsMemberPermissionsPermissionDeleteACLs RoleNewParamsMemberPermissionsPermission = "delete_acls"
)

func (r RoleNewParamsMemberPermissionsPermission) IsKnown() bool {
	switch r {
	case RoleNewParamsMemberPermissionsPermissionCreate, RoleNewParamsMemberPermissionsPermissionRead, RoleNewParamsMemberPermissionsPermissionUpdate, RoleNewParamsMemberPermissionsPermissionDelete, RoleNewParamsMemberPermissionsPermissionCreateACLs, RoleNewParamsMemberPermissionsPermissionReadACLs, RoleNewParamsMemberPermissionsPermissionUpdateACLs, RoleNewParamsMemberPermissionsPermissionDeleteACLs:
		return true
	}
	return false
}

// The object type that the ACL applies to
type RoleNewParamsMemberPermissionsRestrictObjectType string

const (
	RoleNewParamsMemberPermissionsRestrictObjectTypeOrganization  RoleNewParamsMemberPermissionsRestrictObjectType = "organization"
	RoleNewParamsMemberPermissionsRestrictObjectTypeProject       RoleNewParamsMemberPermissionsRestrictObjectType = "project"
	RoleNewParamsMemberPermissionsRestrictObjectTypeExperiment    RoleNewParamsMemberPermissionsRestrictObjectType = "experiment"
	RoleNewParamsMemberPermissionsRestrictObjectTypeDataset       RoleNewParamsMemberPermissionsRestrictObjectType = "dataset"
	RoleNewParamsMemberPermissionsRestrictObjectTypePrompt        RoleNewParamsMemberPermissionsRestrictObjectType = "prompt"
	RoleNewParamsMemberPermissionsRestrictObjectTypePromptSession RoleNewParamsMemberPermissionsRestrictObjectType = "prompt_session"
	RoleNewParamsMemberPermissionsRestrictObjectTypeGroup         RoleNewParamsMemberPermissionsRestrictObjectType = "group"
	RoleNewParamsMemberPermissionsRestrictObjectTypeRole          RoleNewParamsMemberPermissionsRestrictObjectType = "role"
	RoleNewParamsMemberPermissionsRestrictObjectTypeOrgMember     RoleNewParamsMemberPermissionsRestrictObjectType = "org_member"
	RoleNewParamsMemberPermissionsRestrictObjectTypeProjectLog    RoleNewParamsMemberPermissionsRestrictObjectType = "project_log"
	RoleNewParamsMemberPermissionsRestrictObjectTypeOrgProject    RoleNewParamsMemberPermissionsRestrictObjectType = "org_project"
)

func (r RoleNewParamsMemberPermissionsRestrictObjectType) IsKnown() bool {
	switch r {
	case RoleNewParamsMemberPermissionsRestrictObjectTypeOrganization, RoleNewParamsMemberPermissionsRestrictObjectTypeProject, RoleNewParamsMemberPermissionsRestrictObjectTypeExperiment, RoleNewParamsMemberPermissionsRestrictObjectTypeDataset, RoleNewParamsMemberPermissionsRestrictObjectTypePrompt, RoleNewParamsMemberPermissionsRestrictObjectTypePromptSession, RoleNewParamsMemberPermissionsRestrictObjectTypeGroup, RoleNewParamsMemberPermissionsRestrictObjectTypeRole, RoleNewParamsMemberPermissionsRestrictObjectTypeOrgMember, RoleNewParamsMemberPermissionsRestrictObjectTypeProjectLog, RoleNewParamsMemberPermissionsRestrictObjectTypeOrgProject:
		return true
	}
	return false
}

type RoleUpdateParams struct {
	// A list of permissions to add to the role
	AddMemberPermissions param.Field[[]RoleUpdateParamsAddMemberPermission] `json:"add_member_permissions"`
	// A list of role IDs to add to the role's inheriting-from set
	AddMemberRoles param.Field[[]string] `json:"add_member_roles" format:"uuid"`
	// Textual description of the role
	Description param.Field[string] `json:"description"`
	// Name of the role
	Name param.Field[string] `json:"name"`
	// A list of permissions to remove from the role
	RemoveMemberPermissions param.Field[[]RoleUpdateParamsRemoveMemberPermission] `json:"remove_member_permissions"`
	// A list of role IDs to remove from the role's inheriting-from set
	RemoveMemberRoles param.Field[[]string] `json:"remove_member_roles" format:"uuid"`
}

func (r RoleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RoleUpdateParamsAddMemberPermission struct {
	// Each permission permits a certain type of operation on an object in the system
	//
	// Permissions can be assigned to to objects on an individual basis, or grouped
	// into roles
	Permission param.Field[RoleUpdateParamsAddMemberPermissionsPermission] `json:"permission,required"`
	// The object type that the ACL applies to
	RestrictObjectType param.Field[RoleUpdateParamsAddMemberPermissionsRestrictObjectType] `json:"restrict_object_type"`
}

func (r RoleUpdateParamsAddMemberPermission) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
type RoleUpdateParamsAddMemberPermissionsPermission string

const (
	RoleUpdateParamsAddMemberPermissionsPermissionCreate     RoleUpdateParamsAddMemberPermissionsPermission = "create"
	RoleUpdateParamsAddMemberPermissionsPermissionRead       RoleUpdateParamsAddMemberPermissionsPermission = "read"
	RoleUpdateParamsAddMemberPermissionsPermissionUpdate     RoleUpdateParamsAddMemberPermissionsPermission = "update"
	RoleUpdateParamsAddMemberPermissionsPermissionDelete     RoleUpdateParamsAddMemberPermissionsPermission = "delete"
	RoleUpdateParamsAddMemberPermissionsPermissionCreateACLs RoleUpdateParamsAddMemberPermissionsPermission = "create_acls"
	RoleUpdateParamsAddMemberPermissionsPermissionReadACLs   RoleUpdateParamsAddMemberPermissionsPermission = "read_acls"
	RoleUpdateParamsAddMemberPermissionsPermissionUpdateACLs RoleUpdateParamsAddMemberPermissionsPermission = "update_acls"
	RoleUpdateParamsAddMemberPermissionsPermissionDeleteACLs RoleUpdateParamsAddMemberPermissionsPermission = "delete_acls"
)

func (r RoleUpdateParamsAddMemberPermissionsPermission) IsKnown() bool {
	switch r {
	case RoleUpdateParamsAddMemberPermissionsPermissionCreate, RoleUpdateParamsAddMemberPermissionsPermissionRead, RoleUpdateParamsAddMemberPermissionsPermissionUpdate, RoleUpdateParamsAddMemberPermissionsPermissionDelete, RoleUpdateParamsAddMemberPermissionsPermissionCreateACLs, RoleUpdateParamsAddMemberPermissionsPermissionReadACLs, RoleUpdateParamsAddMemberPermissionsPermissionUpdateACLs, RoleUpdateParamsAddMemberPermissionsPermissionDeleteACLs:
		return true
	}
	return false
}

// The object type that the ACL applies to
type RoleUpdateParamsAddMemberPermissionsRestrictObjectType string

const (
	RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeOrganization  RoleUpdateParamsAddMemberPermissionsRestrictObjectType = "organization"
	RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeProject       RoleUpdateParamsAddMemberPermissionsRestrictObjectType = "project"
	RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeExperiment    RoleUpdateParamsAddMemberPermissionsRestrictObjectType = "experiment"
	RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeDataset       RoleUpdateParamsAddMemberPermissionsRestrictObjectType = "dataset"
	RoleUpdateParamsAddMemberPermissionsRestrictObjectTypePrompt        RoleUpdateParamsAddMemberPermissionsRestrictObjectType = "prompt"
	RoleUpdateParamsAddMemberPermissionsRestrictObjectTypePromptSession RoleUpdateParamsAddMemberPermissionsRestrictObjectType = "prompt_session"
	RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeGroup         RoleUpdateParamsAddMemberPermissionsRestrictObjectType = "group"
	RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeRole          RoleUpdateParamsAddMemberPermissionsRestrictObjectType = "role"
	RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeOrgMember     RoleUpdateParamsAddMemberPermissionsRestrictObjectType = "org_member"
	RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeProjectLog    RoleUpdateParamsAddMemberPermissionsRestrictObjectType = "project_log"
	RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeOrgProject    RoleUpdateParamsAddMemberPermissionsRestrictObjectType = "org_project"
)

func (r RoleUpdateParamsAddMemberPermissionsRestrictObjectType) IsKnown() bool {
	switch r {
	case RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeOrganization, RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeProject, RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeExperiment, RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeDataset, RoleUpdateParamsAddMemberPermissionsRestrictObjectTypePrompt, RoleUpdateParamsAddMemberPermissionsRestrictObjectTypePromptSession, RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeGroup, RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeRole, RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeOrgMember, RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeProjectLog, RoleUpdateParamsAddMemberPermissionsRestrictObjectTypeOrgProject:
		return true
	}
	return false
}

type RoleUpdateParamsRemoveMemberPermission struct {
	// Each permission permits a certain type of operation on an object in the system
	//
	// Permissions can be assigned to to objects on an individual basis, or grouped
	// into roles
	Permission param.Field[RoleUpdateParamsRemoveMemberPermissionsPermission] `json:"permission,required"`
	// The object type that the ACL applies to
	RestrictObjectType param.Field[RoleUpdateParamsRemoveMemberPermissionsRestrictObjectType] `json:"restrict_object_type"`
}

func (r RoleUpdateParamsRemoveMemberPermission) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
type RoleUpdateParamsRemoveMemberPermissionsPermission string

const (
	RoleUpdateParamsRemoveMemberPermissionsPermissionCreate     RoleUpdateParamsRemoveMemberPermissionsPermission = "create"
	RoleUpdateParamsRemoveMemberPermissionsPermissionRead       RoleUpdateParamsRemoveMemberPermissionsPermission = "read"
	RoleUpdateParamsRemoveMemberPermissionsPermissionUpdate     RoleUpdateParamsRemoveMemberPermissionsPermission = "update"
	RoleUpdateParamsRemoveMemberPermissionsPermissionDelete     RoleUpdateParamsRemoveMemberPermissionsPermission = "delete"
	RoleUpdateParamsRemoveMemberPermissionsPermissionCreateACLs RoleUpdateParamsRemoveMemberPermissionsPermission = "create_acls"
	RoleUpdateParamsRemoveMemberPermissionsPermissionReadACLs   RoleUpdateParamsRemoveMemberPermissionsPermission = "read_acls"
	RoleUpdateParamsRemoveMemberPermissionsPermissionUpdateACLs RoleUpdateParamsRemoveMemberPermissionsPermission = "update_acls"
	RoleUpdateParamsRemoveMemberPermissionsPermissionDeleteACLs RoleUpdateParamsRemoveMemberPermissionsPermission = "delete_acls"
)

func (r RoleUpdateParamsRemoveMemberPermissionsPermission) IsKnown() bool {
	switch r {
	case RoleUpdateParamsRemoveMemberPermissionsPermissionCreate, RoleUpdateParamsRemoveMemberPermissionsPermissionRead, RoleUpdateParamsRemoveMemberPermissionsPermissionUpdate, RoleUpdateParamsRemoveMemberPermissionsPermissionDelete, RoleUpdateParamsRemoveMemberPermissionsPermissionCreateACLs, RoleUpdateParamsRemoveMemberPermissionsPermissionReadACLs, RoleUpdateParamsRemoveMemberPermissionsPermissionUpdateACLs, RoleUpdateParamsRemoveMemberPermissionsPermissionDeleteACLs:
		return true
	}
	return false
}

// The object type that the ACL applies to
type RoleUpdateParamsRemoveMemberPermissionsRestrictObjectType string

const (
	RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeOrganization  RoleUpdateParamsRemoveMemberPermissionsRestrictObjectType = "organization"
	RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeProject       RoleUpdateParamsRemoveMemberPermissionsRestrictObjectType = "project"
	RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeExperiment    RoleUpdateParamsRemoveMemberPermissionsRestrictObjectType = "experiment"
	RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeDataset       RoleUpdateParamsRemoveMemberPermissionsRestrictObjectType = "dataset"
	RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypePrompt        RoleUpdateParamsRemoveMemberPermissionsRestrictObjectType = "prompt"
	RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypePromptSession RoleUpdateParamsRemoveMemberPermissionsRestrictObjectType = "prompt_session"
	RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeGroup         RoleUpdateParamsRemoveMemberPermissionsRestrictObjectType = "group"
	RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeRole          RoleUpdateParamsRemoveMemberPermissionsRestrictObjectType = "role"
	RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeOrgMember     RoleUpdateParamsRemoveMemberPermissionsRestrictObjectType = "org_member"
	RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeProjectLog    RoleUpdateParamsRemoveMemberPermissionsRestrictObjectType = "project_log"
	RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeOrgProject    RoleUpdateParamsRemoveMemberPermissionsRestrictObjectType = "org_project"
)

func (r RoleUpdateParamsRemoveMemberPermissionsRestrictObjectType) IsKnown() bool {
	switch r {
	case RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeOrganization, RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeProject, RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeExperiment, RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeDataset, RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypePrompt, RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypePromptSession, RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeGroup, RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeRole, RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeOrgMember, RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeProjectLog, RoleUpdateParamsRemoveMemberPermissionsRestrictObjectTypeOrgProject:
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
	// (permission, restrict_object_type) tuples which belong to this role
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

type RoleReplaceParamsMemberPermission struct {
	// Each permission permits a certain type of operation on an object in the system
	//
	// Permissions can be assigned to to objects on an individual basis, or grouped
	// into roles
	Permission param.Field[RoleReplaceParamsMemberPermissionsPermission] `json:"permission,required"`
	// The object type that the ACL applies to
	RestrictObjectType param.Field[RoleReplaceParamsMemberPermissionsRestrictObjectType] `json:"restrict_object_type"`
}

func (r RoleReplaceParamsMemberPermission) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
type RoleReplaceParamsMemberPermissionsPermission string

const (
	RoleReplaceParamsMemberPermissionsPermissionCreate     RoleReplaceParamsMemberPermissionsPermission = "create"
	RoleReplaceParamsMemberPermissionsPermissionRead       RoleReplaceParamsMemberPermissionsPermission = "read"
	RoleReplaceParamsMemberPermissionsPermissionUpdate     RoleReplaceParamsMemberPermissionsPermission = "update"
	RoleReplaceParamsMemberPermissionsPermissionDelete     RoleReplaceParamsMemberPermissionsPermission = "delete"
	RoleReplaceParamsMemberPermissionsPermissionCreateACLs RoleReplaceParamsMemberPermissionsPermission = "create_acls"
	RoleReplaceParamsMemberPermissionsPermissionReadACLs   RoleReplaceParamsMemberPermissionsPermission = "read_acls"
	RoleReplaceParamsMemberPermissionsPermissionUpdateACLs RoleReplaceParamsMemberPermissionsPermission = "update_acls"
	RoleReplaceParamsMemberPermissionsPermissionDeleteACLs RoleReplaceParamsMemberPermissionsPermission = "delete_acls"
)

func (r RoleReplaceParamsMemberPermissionsPermission) IsKnown() bool {
	switch r {
	case RoleReplaceParamsMemberPermissionsPermissionCreate, RoleReplaceParamsMemberPermissionsPermissionRead, RoleReplaceParamsMemberPermissionsPermissionUpdate, RoleReplaceParamsMemberPermissionsPermissionDelete, RoleReplaceParamsMemberPermissionsPermissionCreateACLs, RoleReplaceParamsMemberPermissionsPermissionReadACLs, RoleReplaceParamsMemberPermissionsPermissionUpdateACLs, RoleReplaceParamsMemberPermissionsPermissionDeleteACLs:
		return true
	}
	return false
}

// The object type that the ACL applies to
type RoleReplaceParamsMemberPermissionsRestrictObjectType string

const (
	RoleReplaceParamsMemberPermissionsRestrictObjectTypeOrganization  RoleReplaceParamsMemberPermissionsRestrictObjectType = "organization"
	RoleReplaceParamsMemberPermissionsRestrictObjectTypeProject       RoleReplaceParamsMemberPermissionsRestrictObjectType = "project"
	RoleReplaceParamsMemberPermissionsRestrictObjectTypeExperiment    RoleReplaceParamsMemberPermissionsRestrictObjectType = "experiment"
	RoleReplaceParamsMemberPermissionsRestrictObjectTypeDataset       RoleReplaceParamsMemberPermissionsRestrictObjectType = "dataset"
	RoleReplaceParamsMemberPermissionsRestrictObjectTypePrompt        RoleReplaceParamsMemberPermissionsRestrictObjectType = "prompt"
	RoleReplaceParamsMemberPermissionsRestrictObjectTypePromptSession RoleReplaceParamsMemberPermissionsRestrictObjectType = "prompt_session"
	RoleReplaceParamsMemberPermissionsRestrictObjectTypeGroup         RoleReplaceParamsMemberPermissionsRestrictObjectType = "group"
	RoleReplaceParamsMemberPermissionsRestrictObjectTypeRole          RoleReplaceParamsMemberPermissionsRestrictObjectType = "role"
	RoleReplaceParamsMemberPermissionsRestrictObjectTypeOrgMember     RoleReplaceParamsMemberPermissionsRestrictObjectType = "org_member"
	RoleReplaceParamsMemberPermissionsRestrictObjectTypeProjectLog    RoleReplaceParamsMemberPermissionsRestrictObjectType = "project_log"
	RoleReplaceParamsMemberPermissionsRestrictObjectTypeOrgProject    RoleReplaceParamsMemberPermissionsRestrictObjectType = "org_project"
)

func (r RoleReplaceParamsMemberPermissionsRestrictObjectType) IsKnown() bool {
	switch r {
	case RoleReplaceParamsMemberPermissionsRestrictObjectTypeOrganization, RoleReplaceParamsMemberPermissionsRestrictObjectTypeProject, RoleReplaceParamsMemberPermissionsRestrictObjectTypeExperiment, RoleReplaceParamsMemberPermissionsRestrictObjectTypeDataset, RoleReplaceParamsMemberPermissionsRestrictObjectTypePrompt, RoleReplaceParamsMemberPermissionsRestrictObjectTypePromptSession, RoleReplaceParamsMemberPermissionsRestrictObjectTypeGroup, RoleReplaceParamsMemberPermissionsRestrictObjectTypeRole, RoleReplaceParamsMemberPermissionsRestrictObjectTypeOrgMember, RoleReplaceParamsMemberPermissionsRestrictObjectTypeProjectLog, RoleReplaceParamsMemberPermissionsRestrictObjectTypeOrgProject:
		return true
	}
	return false
}