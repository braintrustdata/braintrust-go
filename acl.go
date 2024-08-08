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
func (r *ACLService) New(ctx context.Context, body ACLNewParams, opts ...option.RequestOption) (res *ACL, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/acl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an acl object by its id
func (r *ACLService) Get(ctx context.Context, aclID string, opts ...option.RequestOption) (res *ACL, err error) {
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
func (r *ACLService) List(ctx context.Context, query ACLListParams, opts ...option.RequestOption) (res *pagination.ListObjects[ACL], err error) {
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
func (r *ACLService) ListAutoPaging(ctx context.Context, query ACLListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[ACL] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete an acl object by its id
func (r *ACLService) Delete(ctx context.Context, aclID string, opts ...option.RequestOption) (res *ACL, err error) {
	opts = append(r.Options[:], opts...)
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("v1/acl/%s", aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
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
type ACL struct {
	// Unique identifier for the acl
	ID string `json:"id,required" format:"uuid"`
	// The organization the ACL's referred object belongs to
	ObjectOrgID string `json:"_object_org_id,required" format:"uuid"`
	// The id of the object the ACL applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType ACLObjectType `json:"object_type,required,nullable"`
	// Date of acl creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Id of the group the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	GroupID string `json:"group_id,nullable" format:"uuid"`
	// Permission the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	Permission ACLPermission `json:"permission,nullable"`
	// When setting a permission directly, optionally restricts the permission grant to
	// just the specified object type. Cannot be set alongside a `role_id`.
	RestrictObjectType ACLRestrictObjectType `json:"restrict_object_type,nullable"`
	// Id of the role the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	RoleID string `json:"role_id,nullable" format:"uuid"`
	// Id of the user the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	UserID string  `json:"user_id,nullable" format:"uuid"`
	JSON   aclJSON `json:"-"`
}

// aclJSON contains the JSON metadata for the struct [ACL]
type aclJSON struct {
	ID                 apijson.Field
	ObjectOrgID        apijson.Field
	ObjectID           apijson.Field
	ObjectType         apijson.Field
	Created            apijson.Field
	GroupID            apijson.Field
	Permission         apijson.Field
	RestrictObjectType apijson.Field
	RoleID             apijson.Field
	UserID             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclJSON) RawJSON() string {
	return r.raw
}

// The object type that the ACL applies to
type ACLObjectType string

const (
	ACLObjectTypeOrganization  ACLObjectType = "organization"
	ACLObjectTypeProject       ACLObjectType = "project"
	ACLObjectTypeExperiment    ACLObjectType = "experiment"
	ACLObjectTypeDataset       ACLObjectType = "dataset"
	ACLObjectTypePrompt        ACLObjectType = "prompt"
	ACLObjectTypePromptSession ACLObjectType = "prompt_session"
	ACLObjectTypeGroup         ACLObjectType = "group"
	ACLObjectTypeRole          ACLObjectType = "role"
	ACLObjectTypeOrgMember     ACLObjectType = "org_member"
	ACLObjectTypeProjectLog    ACLObjectType = "project_log"
	ACLObjectTypeOrgProject    ACLObjectType = "org_project"
)

func (r ACLObjectType) IsKnown() bool {
	switch r {
	case ACLObjectTypeOrganization, ACLObjectTypeProject, ACLObjectTypeExperiment, ACLObjectTypeDataset, ACLObjectTypePrompt, ACLObjectTypePromptSession, ACLObjectTypeGroup, ACLObjectTypeRole, ACLObjectTypeOrgMember, ACLObjectTypeProjectLog, ACLObjectTypeOrgProject:
		return true
	}
	return false
}

// Permission the ACL grants. Exactly one of `permission` and `role_id` will be
// provided
type ACLPermission string

const (
	ACLPermissionCreate     ACLPermission = "create"
	ACLPermissionRead       ACLPermission = "read"
	ACLPermissionUpdate     ACLPermission = "update"
	ACLPermissionDelete     ACLPermission = "delete"
	ACLPermissionCreateACLs ACLPermission = "create_acls"
	ACLPermissionReadACLs   ACLPermission = "read_acls"
	ACLPermissionUpdateACLs ACLPermission = "update_acls"
	ACLPermissionDeleteACLs ACLPermission = "delete_acls"
)

func (r ACLPermission) IsKnown() bool {
	switch r {
	case ACLPermissionCreate, ACLPermissionRead, ACLPermissionUpdate, ACLPermissionDelete, ACLPermissionCreateACLs, ACLPermissionReadACLs, ACLPermissionUpdateACLs, ACLPermissionDeleteACLs:
		return true
	}
	return false
}

// When setting a permission directly, optionally restricts the permission grant to
// just the specified object type. Cannot be set alongside a `role_id`.
type ACLRestrictObjectType string

const (
	ACLRestrictObjectTypeOrganization  ACLRestrictObjectType = "organization"
	ACLRestrictObjectTypeProject       ACLRestrictObjectType = "project"
	ACLRestrictObjectTypeExperiment    ACLRestrictObjectType = "experiment"
	ACLRestrictObjectTypeDataset       ACLRestrictObjectType = "dataset"
	ACLRestrictObjectTypePrompt        ACLRestrictObjectType = "prompt"
	ACLRestrictObjectTypePromptSession ACLRestrictObjectType = "prompt_session"
	ACLRestrictObjectTypeGroup         ACLRestrictObjectType = "group"
	ACLRestrictObjectTypeRole          ACLRestrictObjectType = "role"
	ACLRestrictObjectTypeOrgMember     ACLRestrictObjectType = "org_member"
	ACLRestrictObjectTypeProjectLog    ACLRestrictObjectType = "project_log"
	ACLRestrictObjectTypeOrgProject    ACLRestrictObjectType = "org_project"
)

func (r ACLRestrictObjectType) IsKnown() bool {
	switch r {
	case ACLRestrictObjectTypeOrganization, ACLRestrictObjectTypeProject, ACLRestrictObjectTypeExperiment, ACLRestrictObjectTypeDataset, ACLRestrictObjectTypePrompt, ACLRestrictObjectTypePromptSession, ACLRestrictObjectTypeGroup, ACLRestrictObjectTypeRole, ACLRestrictObjectTypeOrgMember, ACLRestrictObjectTypeProjectLog, ACLRestrictObjectTypeOrgProject:
		return true
	}
	return false
}

type ACLNewParams struct {
	// The id of the object the ACL applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ACLNewParamsObjectType] `json:"object_type,required"`
	// Id of the group the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	GroupID param.Field[string] `json:"group_id" format:"uuid"`
	// Permission the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	Permission param.Field[ACLNewParamsPermission] `json:"permission"`
	// When setting a permission directly, optionally restricts the permission grant to
	// just the specified object type. Cannot be set alongside a `role_id`.
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

// Permission the ACL grants. Exactly one of `permission` and `role_id` will be
// provided
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

// When setting a permission directly, optionally restricts the permission grant to
// just the specified object type. Cannot be set alongside a `role_id`.
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
