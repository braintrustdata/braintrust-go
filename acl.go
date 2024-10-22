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
func NewACLService(opts ...option.RequestOption) (r *ACLService) {
	r = &ACLService{}
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
	path := "v1/acl/batch-update"
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
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ACLNewParamsObjectType] `json:"object_type,required"`
	// Id of the group the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	GroupID param.Field[string] `json:"group_id" format:"uuid"`
	// Each permission permits a certain type of operation on an object in the system
	//
	// Permissions can be assigned to to objects on an individual basis, or grouped
	// into roles
	Permission param.Field[ACLNewParamsPermission] `json:"permission"`
	// The object type that the ACL applies to
	RestrictObjectType param.Field[ACLNewParamsRestrictObjectType] `json:"restrict_object_type"`
	// Id of the role the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	RoleID param.Field[string] `json:"role_id" format:"uuid"`
	// Id of the user the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	UserID param.Field[string] `json:"user_id" format:"uuid"`
}

func (r ACLNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The object type that the ACL applies to
type ACLNewParamsObjectType string

const (
	ACLNewParamsObjectTypeOrganization  ACLNewParamsObjectType = "organization"
	ACLNewParamsObjectTypeProject       ACLNewParamsObjectType = "project"
	ACLNewParamsObjectTypeExperiment    ACLNewParamsObjectType = "experiment"
	ACLNewParamsObjectTypeDataset       ACLNewParamsObjectType = "dataset"
	ACLNewParamsObjectTypePrompt        ACLNewParamsObjectType = "prompt"
	ACLNewParamsObjectTypePromptSession ACLNewParamsObjectType = "prompt_session"
	ACLNewParamsObjectTypeGroup         ACLNewParamsObjectType = "group"
	ACLNewParamsObjectTypeRole          ACLNewParamsObjectType = "role"
	ACLNewParamsObjectTypeOrgMember     ACLNewParamsObjectType = "org_member"
	ACLNewParamsObjectTypeProjectLog    ACLNewParamsObjectType = "project_log"
	ACLNewParamsObjectTypeOrgProject    ACLNewParamsObjectType = "org_project"
)

func (r ACLNewParamsObjectType) IsKnown() bool {
	switch r {
	case ACLNewParamsObjectTypeOrganization, ACLNewParamsObjectTypeProject, ACLNewParamsObjectTypeExperiment, ACLNewParamsObjectTypeDataset, ACLNewParamsObjectTypePrompt, ACLNewParamsObjectTypePromptSession, ACLNewParamsObjectTypeGroup, ACLNewParamsObjectTypeRole, ACLNewParamsObjectTypeOrgMember, ACLNewParamsObjectTypeProjectLog, ACLNewParamsObjectTypeOrgProject:
		return true
	}
	return false
}

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
type ACLNewParamsPermission string

const (
	ACLNewParamsPermissionCreate     ACLNewParamsPermission = "create"
	ACLNewParamsPermissionRead       ACLNewParamsPermission = "read"
	ACLNewParamsPermissionUpdate     ACLNewParamsPermission = "update"
	ACLNewParamsPermissionDelete     ACLNewParamsPermission = "delete"
	ACLNewParamsPermissionCreateACLs ACLNewParamsPermission = "create_acls"
	ACLNewParamsPermissionReadACLs   ACLNewParamsPermission = "read_acls"
	ACLNewParamsPermissionUpdateACLs ACLNewParamsPermission = "update_acls"
	ACLNewParamsPermissionDeleteACLs ACLNewParamsPermission = "delete_acls"
)

func (r ACLNewParamsPermission) IsKnown() bool {
	switch r {
	case ACLNewParamsPermissionCreate, ACLNewParamsPermissionRead, ACLNewParamsPermissionUpdate, ACLNewParamsPermissionDelete, ACLNewParamsPermissionCreateACLs, ACLNewParamsPermissionReadACLs, ACLNewParamsPermissionUpdateACLs, ACLNewParamsPermissionDeleteACLs:
		return true
	}
	return false
}

// The object type that the ACL applies to
type ACLNewParamsRestrictObjectType string

const (
	ACLNewParamsRestrictObjectTypeOrganization  ACLNewParamsRestrictObjectType = "organization"
	ACLNewParamsRestrictObjectTypeProject       ACLNewParamsRestrictObjectType = "project"
	ACLNewParamsRestrictObjectTypeExperiment    ACLNewParamsRestrictObjectType = "experiment"
	ACLNewParamsRestrictObjectTypeDataset       ACLNewParamsRestrictObjectType = "dataset"
	ACLNewParamsRestrictObjectTypePrompt        ACLNewParamsRestrictObjectType = "prompt"
	ACLNewParamsRestrictObjectTypePromptSession ACLNewParamsRestrictObjectType = "prompt_session"
	ACLNewParamsRestrictObjectTypeGroup         ACLNewParamsRestrictObjectType = "group"
	ACLNewParamsRestrictObjectTypeRole          ACLNewParamsRestrictObjectType = "role"
	ACLNewParamsRestrictObjectTypeOrgMember     ACLNewParamsRestrictObjectType = "org_member"
	ACLNewParamsRestrictObjectTypeProjectLog    ACLNewParamsRestrictObjectType = "project_log"
	ACLNewParamsRestrictObjectTypeOrgProject    ACLNewParamsRestrictObjectType = "org_project"
)

func (r ACLNewParamsRestrictObjectType) IsKnown() bool {
	switch r {
	case ACLNewParamsRestrictObjectTypeOrganization, ACLNewParamsRestrictObjectTypeProject, ACLNewParamsRestrictObjectTypeExperiment, ACLNewParamsRestrictObjectTypeDataset, ACLNewParamsRestrictObjectTypePrompt, ACLNewParamsRestrictObjectTypePromptSession, ACLNewParamsRestrictObjectTypeGroup, ACLNewParamsRestrictObjectTypeRole, ACLNewParamsRestrictObjectTypeOrgMember, ACLNewParamsRestrictObjectTypeProjectLog, ACLNewParamsRestrictObjectTypeOrgProject:
		return true
	}
	return false
}

type ACLListParams struct {
	// The id of the object the ACL applies to
	ObjectID param.Field[string] `query:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ACLListParamsObjectType] `query:"object_type,required"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Field[string] `query:"ending_before" format:"uuid"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs param.Field[ACLListParamsIDsUnion] `query:"ids" format:"uuid"`
	// Limit the number of objects to return
	Limit param.Field[int64] `query:"limit"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Field[string] `query:"starting_after" format:"uuid"`
}

// URLQuery serializes [ACLListParams]'s query parameters as `url.Values`.
func (r ACLListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The object type that the ACL applies to
type ACLListParamsObjectType string

const (
	ACLListParamsObjectTypeOrganization  ACLListParamsObjectType = "organization"
	ACLListParamsObjectTypeProject       ACLListParamsObjectType = "project"
	ACLListParamsObjectTypeExperiment    ACLListParamsObjectType = "experiment"
	ACLListParamsObjectTypeDataset       ACLListParamsObjectType = "dataset"
	ACLListParamsObjectTypePrompt        ACLListParamsObjectType = "prompt"
	ACLListParamsObjectTypePromptSession ACLListParamsObjectType = "prompt_session"
	ACLListParamsObjectTypeGroup         ACLListParamsObjectType = "group"
	ACLListParamsObjectTypeRole          ACLListParamsObjectType = "role"
	ACLListParamsObjectTypeOrgMember     ACLListParamsObjectType = "org_member"
	ACLListParamsObjectTypeProjectLog    ACLListParamsObjectType = "project_log"
	ACLListParamsObjectTypeOrgProject    ACLListParamsObjectType = "org_project"
)

func (r ACLListParamsObjectType) IsKnown() bool {
	switch r {
	case ACLListParamsObjectTypeOrganization, ACLListParamsObjectTypeProject, ACLListParamsObjectTypeExperiment, ACLListParamsObjectTypeDataset, ACLListParamsObjectTypePrompt, ACLListParamsObjectTypePromptSession, ACLListParamsObjectTypeGroup, ACLListParamsObjectTypeRole, ACLListParamsObjectTypeOrgMember, ACLListParamsObjectTypeProjectLog, ACLListParamsObjectTypeOrgProject:
		return true
	}
	return false
}

// Filter search results to a particular set of object IDs. To specify a list of
// IDs, include the query param multiple times
//
// Satisfied by [shared.UnionString], [ACLListParamsIDsArray].
type ACLListParamsIDsUnion interface {
	ImplementsACLListParamsIDsUnion()
}

type ACLListParamsIDsArray []string

func (r ACLListParamsIDsArray) ImplementsACLListParamsIDsUnion() {}

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
	AddACLs param.Field[[]ACLBatchUpdateParamsAddACL] `json:"add_acls"`
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
	RemoveACLs param.Field[[]ACLBatchUpdateParamsRemoveACL] `json:"remove_acls"`
}

func (r ACLBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
type ACLBatchUpdateParamsAddACL struct {
	// The id of the object the ACL applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ACLBatchUpdateParamsAddACLsObjectType] `json:"object_type,required"`
	// Id of the group the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	GroupID param.Field[string] `json:"group_id" format:"uuid"`
	// Each permission permits a certain type of operation on an object in the system
	//
	// Permissions can be assigned to to objects on an individual basis, or grouped
	// into roles
	Permission param.Field[ACLBatchUpdateParamsAddACLsPermission] `json:"permission"`
	// The object type that the ACL applies to
	RestrictObjectType param.Field[ACLBatchUpdateParamsAddACLsRestrictObjectType] `json:"restrict_object_type"`
	// Id of the role the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	RoleID param.Field[string] `json:"role_id" format:"uuid"`
	// Id of the user the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	UserID param.Field[string] `json:"user_id" format:"uuid"`
}

func (r ACLBatchUpdateParamsAddACL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The object type that the ACL applies to
type ACLBatchUpdateParamsAddACLsObjectType string

const (
	ACLBatchUpdateParamsAddACLsObjectTypeOrganization  ACLBatchUpdateParamsAddACLsObjectType = "organization"
	ACLBatchUpdateParamsAddACLsObjectTypeProject       ACLBatchUpdateParamsAddACLsObjectType = "project"
	ACLBatchUpdateParamsAddACLsObjectTypeExperiment    ACLBatchUpdateParamsAddACLsObjectType = "experiment"
	ACLBatchUpdateParamsAddACLsObjectTypeDataset       ACLBatchUpdateParamsAddACLsObjectType = "dataset"
	ACLBatchUpdateParamsAddACLsObjectTypePrompt        ACLBatchUpdateParamsAddACLsObjectType = "prompt"
	ACLBatchUpdateParamsAddACLsObjectTypePromptSession ACLBatchUpdateParamsAddACLsObjectType = "prompt_session"
	ACLBatchUpdateParamsAddACLsObjectTypeGroup         ACLBatchUpdateParamsAddACLsObjectType = "group"
	ACLBatchUpdateParamsAddACLsObjectTypeRole          ACLBatchUpdateParamsAddACLsObjectType = "role"
	ACLBatchUpdateParamsAddACLsObjectTypeOrgMember     ACLBatchUpdateParamsAddACLsObjectType = "org_member"
	ACLBatchUpdateParamsAddACLsObjectTypeProjectLog    ACLBatchUpdateParamsAddACLsObjectType = "project_log"
	ACLBatchUpdateParamsAddACLsObjectTypeOrgProject    ACLBatchUpdateParamsAddACLsObjectType = "org_project"
)

func (r ACLBatchUpdateParamsAddACLsObjectType) IsKnown() bool {
	switch r {
	case ACLBatchUpdateParamsAddACLsObjectTypeOrganization, ACLBatchUpdateParamsAddACLsObjectTypeProject, ACLBatchUpdateParamsAddACLsObjectTypeExperiment, ACLBatchUpdateParamsAddACLsObjectTypeDataset, ACLBatchUpdateParamsAddACLsObjectTypePrompt, ACLBatchUpdateParamsAddACLsObjectTypePromptSession, ACLBatchUpdateParamsAddACLsObjectTypeGroup, ACLBatchUpdateParamsAddACLsObjectTypeRole, ACLBatchUpdateParamsAddACLsObjectTypeOrgMember, ACLBatchUpdateParamsAddACLsObjectTypeProjectLog, ACLBatchUpdateParamsAddACLsObjectTypeOrgProject:
		return true
	}
	return false
}

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
type ACLBatchUpdateParamsAddACLsPermission string

const (
	ACLBatchUpdateParamsAddACLsPermissionCreate     ACLBatchUpdateParamsAddACLsPermission = "create"
	ACLBatchUpdateParamsAddACLsPermissionRead       ACLBatchUpdateParamsAddACLsPermission = "read"
	ACLBatchUpdateParamsAddACLsPermissionUpdate     ACLBatchUpdateParamsAddACLsPermission = "update"
	ACLBatchUpdateParamsAddACLsPermissionDelete     ACLBatchUpdateParamsAddACLsPermission = "delete"
	ACLBatchUpdateParamsAddACLsPermissionCreateACLs ACLBatchUpdateParamsAddACLsPermission = "create_acls"
	ACLBatchUpdateParamsAddACLsPermissionReadACLs   ACLBatchUpdateParamsAddACLsPermission = "read_acls"
	ACLBatchUpdateParamsAddACLsPermissionUpdateACLs ACLBatchUpdateParamsAddACLsPermission = "update_acls"
	ACLBatchUpdateParamsAddACLsPermissionDeleteACLs ACLBatchUpdateParamsAddACLsPermission = "delete_acls"
)

func (r ACLBatchUpdateParamsAddACLsPermission) IsKnown() bool {
	switch r {
	case ACLBatchUpdateParamsAddACLsPermissionCreate, ACLBatchUpdateParamsAddACLsPermissionRead, ACLBatchUpdateParamsAddACLsPermissionUpdate, ACLBatchUpdateParamsAddACLsPermissionDelete, ACLBatchUpdateParamsAddACLsPermissionCreateACLs, ACLBatchUpdateParamsAddACLsPermissionReadACLs, ACLBatchUpdateParamsAddACLsPermissionUpdateACLs, ACLBatchUpdateParamsAddACLsPermissionDeleteACLs:
		return true
	}
	return false
}

// The object type that the ACL applies to
type ACLBatchUpdateParamsAddACLsRestrictObjectType string

const (
	ACLBatchUpdateParamsAddACLsRestrictObjectTypeOrganization  ACLBatchUpdateParamsAddACLsRestrictObjectType = "organization"
	ACLBatchUpdateParamsAddACLsRestrictObjectTypeProject       ACLBatchUpdateParamsAddACLsRestrictObjectType = "project"
	ACLBatchUpdateParamsAddACLsRestrictObjectTypeExperiment    ACLBatchUpdateParamsAddACLsRestrictObjectType = "experiment"
	ACLBatchUpdateParamsAddACLsRestrictObjectTypeDataset       ACLBatchUpdateParamsAddACLsRestrictObjectType = "dataset"
	ACLBatchUpdateParamsAddACLsRestrictObjectTypePrompt        ACLBatchUpdateParamsAddACLsRestrictObjectType = "prompt"
	ACLBatchUpdateParamsAddACLsRestrictObjectTypePromptSession ACLBatchUpdateParamsAddACLsRestrictObjectType = "prompt_session"
	ACLBatchUpdateParamsAddACLsRestrictObjectTypeGroup         ACLBatchUpdateParamsAddACLsRestrictObjectType = "group"
	ACLBatchUpdateParamsAddACLsRestrictObjectTypeRole          ACLBatchUpdateParamsAddACLsRestrictObjectType = "role"
	ACLBatchUpdateParamsAddACLsRestrictObjectTypeOrgMember     ACLBatchUpdateParamsAddACLsRestrictObjectType = "org_member"
	ACLBatchUpdateParamsAddACLsRestrictObjectTypeProjectLog    ACLBatchUpdateParamsAddACLsRestrictObjectType = "project_log"
	ACLBatchUpdateParamsAddACLsRestrictObjectTypeOrgProject    ACLBatchUpdateParamsAddACLsRestrictObjectType = "org_project"
)

func (r ACLBatchUpdateParamsAddACLsRestrictObjectType) IsKnown() bool {
	switch r {
	case ACLBatchUpdateParamsAddACLsRestrictObjectTypeOrganization, ACLBatchUpdateParamsAddACLsRestrictObjectTypeProject, ACLBatchUpdateParamsAddACLsRestrictObjectTypeExperiment, ACLBatchUpdateParamsAddACLsRestrictObjectTypeDataset, ACLBatchUpdateParamsAddACLsRestrictObjectTypePrompt, ACLBatchUpdateParamsAddACLsRestrictObjectTypePromptSession, ACLBatchUpdateParamsAddACLsRestrictObjectTypeGroup, ACLBatchUpdateParamsAddACLsRestrictObjectTypeRole, ACLBatchUpdateParamsAddACLsRestrictObjectTypeOrgMember, ACLBatchUpdateParamsAddACLsRestrictObjectTypeProjectLog, ACLBatchUpdateParamsAddACLsRestrictObjectTypeOrgProject:
		return true
	}
	return false
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
type ACLBatchUpdateParamsRemoveACL struct {
	// The id of the object the ACL applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ACLBatchUpdateParamsRemoveACLsObjectType] `json:"object_type,required"`
	// Id of the group the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	GroupID param.Field[string] `json:"group_id" format:"uuid"`
	// Each permission permits a certain type of operation on an object in the system
	//
	// Permissions can be assigned to to objects on an individual basis, or grouped
	// into roles
	Permission param.Field[ACLBatchUpdateParamsRemoveACLsPermission] `json:"permission"`
	// The object type that the ACL applies to
	RestrictObjectType param.Field[ACLBatchUpdateParamsRemoveACLsRestrictObjectType] `json:"restrict_object_type"`
	// Id of the role the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	RoleID param.Field[string] `json:"role_id" format:"uuid"`
	// Id of the user the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	UserID param.Field[string] `json:"user_id" format:"uuid"`
}

func (r ACLBatchUpdateParamsRemoveACL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The object type that the ACL applies to
type ACLBatchUpdateParamsRemoveACLsObjectType string

const (
	ACLBatchUpdateParamsRemoveACLsObjectTypeOrganization  ACLBatchUpdateParamsRemoveACLsObjectType = "organization"
	ACLBatchUpdateParamsRemoveACLsObjectTypeProject       ACLBatchUpdateParamsRemoveACLsObjectType = "project"
	ACLBatchUpdateParamsRemoveACLsObjectTypeExperiment    ACLBatchUpdateParamsRemoveACLsObjectType = "experiment"
	ACLBatchUpdateParamsRemoveACLsObjectTypeDataset       ACLBatchUpdateParamsRemoveACLsObjectType = "dataset"
	ACLBatchUpdateParamsRemoveACLsObjectTypePrompt        ACLBatchUpdateParamsRemoveACLsObjectType = "prompt"
	ACLBatchUpdateParamsRemoveACLsObjectTypePromptSession ACLBatchUpdateParamsRemoveACLsObjectType = "prompt_session"
	ACLBatchUpdateParamsRemoveACLsObjectTypeGroup         ACLBatchUpdateParamsRemoveACLsObjectType = "group"
	ACLBatchUpdateParamsRemoveACLsObjectTypeRole          ACLBatchUpdateParamsRemoveACLsObjectType = "role"
	ACLBatchUpdateParamsRemoveACLsObjectTypeOrgMember     ACLBatchUpdateParamsRemoveACLsObjectType = "org_member"
	ACLBatchUpdateParamsRemoveACLsObjectTypeProjectLog    ACLBatchUpdateParamsRemoveACLsObjectType = "project_log"
	ACLBatchUpdateParamsRemoveACLsObjectTypeOrgProject    ACLBatchUpdateParamsRemoveACLsObjectType = "org_project"
)

func (r ACLBatchUpdateParamsRemoveACLsObjectType) IsKnown() bool {
	switch r {
	case ACLBatchUpdateParamsRemoveACLsObjectTypeOrganization, ACLBatchUpdateParamsRemoveACLsObjectTypeProject, ACLBatchUpdateParamsRemoveACLsObjectTypeExperiment, ACLBatchUpdateParamsRemoveACLsObjectTypeDataset, ACLBatchUpdateParamsRemoveACLsObjectTypePrompt, ACLBatchUpdateParamsRemoveACLsObjectTypePromptSession, ACLBatchUpdateParamsRemoveACLsObjectTypeGroup, ACLBatchUpdateParamsRemoveACLsObjectTypeRole, ACLBatchUpdateParamsRemoveACLsObjectTypeOrgMember, ACLBatchUpdateParamsRemoveACLsObjectTypeProjectLog, ACLBatchUpdateParamsRemoveACLsObjectTypeOrgProject:
		return true
	}
	return false
}

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
type ACLBatchUpdateParamsRemoveACLsPermission string

const (
	ACLBatchUpdateParamsRemoveACLsPermissionCreate     ACLBatchUpdateParamsRemoveACLsPermission = "create"
	ACLBatchUpdateParamsRemoveACLsPermissionRead       ACLBatchUpdateParamsRemoveACLsPermission = "read"
	ACLBatchUpdateParamsRemoveACLsPermissionUpdate     ACLBatchUpdateParamsRemoveACLsPermission = "update"
	ACLBatchUpdateParamsRemoveACLsPermissionDelete     ACLBatchUpdateParamsRemoveACLsPermission = "delete"
	ACLBatchUpdateParamsRemoveACLsPermissionCreateACLs ACLBatchUpdateParamsRemoveACLsPermission = "create_acls"
	ACLBatchUpdateParamsRemoveACLsPermissionReadACLs   ACLBatchUpdateParamsRemoveACLsPermission = "read_acls"
	ACLBatchUpdateParamsRemoveACLsPermissionUpdateACLs ACLBatchUpdateParamsRemoveACLsPermission = "update_acls"
	ACLBatchUpdateParamsRemoveACLsPermissionDeleteACLs ACLBatchUpdateParamsRemoveACLsPermission = "delete_acls"
)

func (r ACLBatchUpdateParamsRemoveACLsPermission) IsKnown() bool {
	switch r {
	case ACLBatchUpdateParamsRemoveACLsPermissionCreate, ACLBatchUpdateParamsRemoveACLsPermissionRead, ACLBatchUpdateParamsRemoveACLsPermissionUpdate, ACLBatchUpdateParamsRemoveACLsPermissionDelete, ACLBatchUpdateParamsRemoveACLsPermissionCreateACLs, ACLBatchUpdateParamsRemoveACLsPermissionReadACLs, ACLBatchUpdateParamsRemoveACLsPermissionUpdateACLs, ACLBatchUpdateParamsRemoveACLsPermissionDeleteACLs:
		return true
	}
	return false
}

// The object type that the ACL applies to
type ACLBatchUpdateParamsRemoveACLsRestrictObjectType string

const (
	ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeOrganization  ACLBatchUpdateParamsRemoveACLsRestrictObjectType = "organization"
	ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeProject       ACLBatchUpdateParamsRemoveACLsRestrictObjectType = "project"
	ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeExperiment    ACLBatchUpdateParamsRemoveACLsRestrictObjectType = "experiment"
	ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeDataset       ACLBatchUpdateParamsRemoveACLsRestrictObjectType = "dataset"
	ACLBatchUpdateParamsRemoveACLsRestrictObjectTypePrompt        ACLBatchUpdateParamsRemoveACLsRestrictObjectType = "prompt"
	ACLBatchUpdateParamsRemoveACLsRestrictObjectTypePromptSession ACLBatchUpdateParamsRemoveACLsRestrictObjectType = "prompt_session"
	ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeGroup         ACLBatchUpdateParamsRemoveACLsRestrictObjectType = "group"
	ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeRole          ACLBatchUpdateParamsRemoveACLsRestrictObjectType = "role"
	ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeOrgMember     ACLBatchUpdateParamsRemoveACLsRestrictObjectType = "org_member"
	ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeProjectLog    ACLBatchUpdateParamsRemoveACLsRestrictObjectType = "project_log"
	ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeOrgProject    ACLBatchUpdateParamsRemoveACLsRestrictObjectType = "org_project"
)

func (r ACLBatchUpdateParamsRemoveACLsRestrictObjectType) IsKnown() bool {
	switch r {
	case ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeOrganization, ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeProject, ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeExperiment, ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeDataset, ACLBatchUpdateParamsRemoveACLsRestrictObjectTypePrompt, ACLBatchUpdateParamsRemoveACLsRestrictObjectTypePromptSession, ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeGroup, ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeRole, ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeOrgMember, ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeProjectLog, ACLBatchUpdateParamsRemoveACLsRestrictObjectTypeOrgProject:
		return true
	}
	return false
}

type ACLFindAndDeleteParams struct {
	// The id of the object the ACL applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ACLFindAndDeleteParamsObjectType] `json:"object_type,required"`
	// Id of the group the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	GroupID param.Field[string] `json:"group_id" format:"uuid"`
	// Each permission permits a certain type of operation on an object in the system
	//
	// Permissions can be assigned to to objects on an individual basis, or grouped
	// into roles
	Permission param.Field[ACLFindAndDeleteParamsPermission] `json:"permission"`
	// The object type that the ACL applies to
	RestrictObjectType param.Field[ACLFindAndDeleteParamsRestrictObjectType] `json:"restrict_object_type"`
	// Id of the role the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	RoleID param.Field[string] `json:"role_id" format:"uuid"`
	// Id of the user the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	UserID param.Field[string] `json:"user_id" format:"uuid"`
}

func (r ACLFindAndDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The object type that the ACL applies to
type ACLFindAndDeleteParamsObjectType string

const (
	ACLFindAndDeleteParamsObjectTypeOrganization  ACLFindAndDeleteParamsObjectType = "organization"
	ACLFindAndDeleteParamsObjectTypeProject       ACLFindAndDeleteParamsObjectType = "project"
	ACLFindAndDeleteParamsObjectTypeExperiment    ACLFindAndDeleteParamsObjectType = "experiment"
	ACLFindAndDeleteParamsObjectTypeDataset       ACLFindAndDeleteParamsObjectType = "dataset"
	ACLFindAndDeleteParamsObjectTypePrompt        ACLFindAndDeleteParamsObjectType = "prompt"
	ACLFindAndDeleteParamsObjectTypePromptSession ACLFindAndDeleteParamsObjectType = "prompt_session"
	ACLFindAndDeleteParamsObjectTypeGroup         ACLFindAndDeleteParamsObjectType = "group"
	ACLFindAndDeleteParamsObjectTypeRole          ACLFindAndDeleteParamsObjectType = "role"
	ACLFindAndDeleteParamsObjectTypeOrgMember     ACLFindAndDeleteParamsObjectType = "org_member"
	ACLFindAndDeleteParamsObjectTypeProjectLog    ACLFindAndDeleteParamsObjectType = "project_log"
	ACLFindAndDeleteParamsObjectTypeOrgProject    ACLFindAndDeleteParamsObjectType = "org_project"
)

func (r ACLFindAndDeleteParamsObjectType) IsKnown() bool {
	switch r {
	case ACLFindAndDeleteParamsObjectTypeOrganization, ACLFindAndDeleteParamsObjectTypeProject, ACLFindAndDeleteParamsObjectTypeExperiment, ACLFindAndDeleteParamsObjectTypeDataset, ACLFindAndDeleteParamsObjectTypePrompt, ACLFindAndDeleteParamsObjectTypePromptSession, ACLFindAndDeleteParamsObjectTypeGroup, ACLFindAndDeleteParamsObjectTypeRole, ACLFindAndDeleteParamsObjectTypeOrgMember, ACLFindAndDeleteParamsObjectTypeProjectLog, ACLFindAndDeleteParamsObjectTypeOrgProject:
		return true
	}
	return false
}

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
type ACLFindAndDeleteParamsPermission string

const (
	ACLFindAndDeleteParamsPermissionCreate     ACLFindAndDeleteParamsPermission = "create"
	ACLFindAndDeleteParamsPermissionRead       ACLFindAndDeleteParamsPermission = "read"
	ACLFindAndDeleteParamsPermissionUpdate     ACLFindAndDeleteParamsPermission = "update"
	ACLFindAndDeleteParamsPermissionDelete     ACLFindAndDeleteParamsPermission = "delete"
	ACLFindAndDeleteParamsPermissionCreateACLs ACLFindAndDeleteParamsPermission = "create_acls"
	ACLFindAndDeleteParamsPermissionReadACLs   ACLFindAndDeleteParamsPermission = "read_acls"
	ACLFindAndDeleteParamsPermissionUpdateACLs ACLFindAndDeleteParamsPermission = "update_acls"
	ACLFindAndDeleteParamsPermissionDeleteACLs ACLFindAndDeleteParamsPermission = "delete_acls"
)

func (r ACLFindAndDeleteParamsPermission) IsKnown() bool {
	switch r {
	case ACLFindAndDeleteParamsPermissionCreate, ACLFindAndDeleteParamsPermissionRead, ACLFindAndDeleteParamsPermissionUpdate, ACLFindAndDeleteParamsPermissionDelete, ACLFindAndDeleteParamsPermissionCreateACLs, ACLFindAndDeleteParamsPermissionReadACLs, ACLFindAndDeleteParamsPermissionUpdateACLs, ACLFindAndDeleteParamsPermissionDeleteACLs:
		return true
	}
	return false
}

// The object type that the ACL applies to
type ACLFindAndDeleteParamsRestrictObjectType string

const (
	ACLFindAndDeleteParamsRestrictObjectTypeOrganization  ACLFindAndDeleteParamsRestrictObjectType = "organization"
	ACLFindAndDeleteParamsRestrictObjectTypeProject       ACLFindAndDeleteParamsRestrictObjectType = "project"
	ACLFindAndDeleteParamsRestrictObjectTypeExperiment    ACLFindAndDeleteParamsRestrictObjectType = "experiment"
	ACLFindAndDeleteParamsRestrictObjectTypeDataset       ACLFindAndDeleteParamsRestrictObjectType = "dataset"
	ACLFindAndDeleteParamsRestrictObjectTypePrompt        ACLFindAndDeleteParamsRestrictObjectType = "prompt"
	ACLFindAndDeleteParamsRestrictObjectTypePromptSession ACLFindAndDeleteParamsRestrictObjectType = "prompt_session"
	ACLFindAndDeleteParamsRestrictObjectTypeGroup         ACLFindAndDeleteParamsRestrictObjectType = "group"
	ACLFindAndDeleteParamsRestrictObjectTypeRole          ACLFindAndDeleteParamsRestrictObjectType = "role"
	ACLFindAndDeleteParamsRestrictObjectTypeOrgMember     ACLFindAndDeleteParamsRestrictObjectType = "org_member"
	ACLFindAndDeleteParamsRestrictObjectTypeProjectLog    ACLFindAndDeleteParamsRestrictObjectType = "project_log"
	ACLFindAndDeleteParamsRestrictObjectTypeOrgProject    ACLFindAndDeleteParamsRestrictObjectType = "org_project"
)

func (r ACLFindAndDeleteParamsRestrictObjectType) IsKnown() bool {
	switch r {
	case ACLFindAndDeleteParamsRestrictObjectTypeOrganization, ACLFindAndDeleteParamsRestrictObjectTypeProject, ACLFindAndDeleteParamsRestrictObjectTypeExperiment, ACLFindAndDeleteParamsRestrictObjectTypeDataset, ACLFindAndDeleteParamsRestrictObjectTypePrompt, ACLFindAndDeleteParamsRestrictObjectTypePromptSession, ACLFindAndDeleteParamsRestrictObjectTypeGroup, ACLFindAndDeleteParamsRestrictObjectTypeRole, ACLFindAndDeleteParamsRestrictObjectTypeOrgMember, ACLFindAndDeleteParamsRestrictObjectTypeProjectLog, ACLFindAndDeleteParamsRestrictObjectTypeOrgProject:
		return true
	}
	return false
}
