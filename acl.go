// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/internal/apiquery"
	"github.com/braintrustdata/braintrust-go/internal/pagination"
	"github.com/braintrustdata/braintrust-go/internal/param"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/tidwall/gjson"
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
	path := fmt.Sprintf("v1/acl/%s", aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List out all acls. The acls are sorted by creation date, with the most
// recently-created acls coming first
func (r *ACLService) List(ctx context.Context, query ACLListParams, opts ...option.RequestOption) (res *pagination.ListObjects[ACL], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
	path := fmt.Sprintf("v1/acl/%s", aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// NOTE: This operation is deprecated and will be removed in a future revision of
// the API. Create or replace a new acl. If there is an existing acl with the same
// contents as the one specified in the request, will return the existing acl
// unmodified, will replace the existing acl with the provided fields
func (r *ACLService) Replace(ctx context.Context, body ACLReplaceParams, opts ...option.RequestOption) (res *ACL, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/acl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
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
// `restrict_object_type` in the ACL.
type ACL struct {
	// Unique identifier for the acl
	ID string `json:"id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType ACLObjectType `json:"object_type,required"`
	// The id of the object the ACL applies to
	ObjectID           string      `json:"object_id,required" format:"uuid"`
	RestrictObjectType interface{} `json:"restrict_object_type,required"`
	// The organization the ACL's referred object belongs to
	ObjectOrgID string `json:"_object_org_id,required" format:"uuid"`
	// Date of acl creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Id of the user the ACL applies to
	UserID string `json:"user_id" format:"uuid"`
	// Permission the ACL grants
	Permission ACLPermission `json:"permission"`
	// Id of the role the ACL grants
	RoleID string `json:"role_id" format:"uuid"`
	// Id of the group the ACL applies to
	GroupID string  `json:"group_id" format:"uuid"`
	JSON    aclJSON `json:"-"`
	union   ACLUnion
}

// aclJSON contains the JSON metadata for the struct [ACL]
type aclJSON struct {
	ID                 apijson.Field
	ObjectType         apijson.Field
	ObjectID           apijson.Field
	RestrictObjectType apijson.Field
	ObjectOrgID        apijson.Field
	Created            apijson.Field
	UserID             apijson.Field
	Permission         apijson.Field
	RoleID             apijson.Field
	GroupID            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r aclJSON) RawJSON() string {
	return r.raw
}

func (r *ACL) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r ACL) AsUnion() ACLUnion {
	return r.union
}

// An ACL grants a certain permission or role to a certain user or group on an
// object.
//
// ACLs are inherited across the object hierarchy. So for example, if a user has
// read permissions on a project, they will also have read permissions on any
// experiment, dataset, etc. created within that project.
//
// To restrict a grant to a particular sub-object, you may specify
// `restrict_object_type` in the ACL.
//
// Union satisfied by [ACLUserPermissionACL], [ACLUserRoleACL],
// [ACLGroupPermissionACL] or [ACLGroupRoleACL].
type ACLUnion interface {
	implementsACL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ACLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ACLUserPermissionACL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ACLUserRoleACL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ACLGroupPermissionACL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ACLGroupRoleACL{}),
		},
	)
}

type ACLUserPermissionACL struct {
	// Unique identifier for the acl
	ID string `json:"id,required" format:"uuid"`
	// The organization the ACL's referred object belongs to
	ObjectOrgID string `json:"_object_org_id,required" format:"uuid"`
	// The id of the object the ACL applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType ACLUserPermissionACLObjectType `json:"object_type,required"`
	// Permission the ACL grants
	Permission ACLUserPermissionACLPermission `json:"permission,required"`
	// Id of the user the ACL applies to
	UserID string `json:"user_id,required" format:"uuid"`
	// Date of acl creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Optionally restricts the permission grant to just the specified object type
	RestrictObjectType ACLUserPermissionACLRestrictObjectTypeUnion `json:"restrict_object_type"`
	JSON               aclUserPermissionACLJSON                    `json:"-"`
}

// aclUserPermissionACLJSON contains the JSON metadata for the struct
// [ACLUserPermissionACL]
type aclUserPermissionACLJSON struct {
	ID                 apijson.Field
	ObjectOrgID        apijson.Field
	ObjectID           apijson.Field
	ObjectType         apijson.Field
	Permission         apijson.Field
	UserID             apijson.Field
	Created            apijson.Field
	RestrictObjectType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ACLUserPermissionACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclUserPermissionACLJSON) RawJSON() string {
	return r.raw
}

func (r ACLUserPermissionACL) implementsACL() {}

// The object type that the ACL applies to
type ACLUserPermissionACLObjectType string

const (
	ACLUserPermissionACLObjectTypeOrganization  ACLUserPermissionACLObjectType = "organization"
	ACLUserPermissionACLObjectTypeProject       ACLUserPermissionACLObjectType = "project"
	ACLUserPermissionACLObjectTypeExperiment    ACLUserPermissionACLObjectType = "experiment"
	ACLUserPermissionACLObjectTypeDataset       ACLUserPermissionACLObjectType = "dataset"
	ACLUserPermissionACLObjectTypePrompt        ACLUserPermissionACLObjectType = "prompt"
	ACLUserPermissionACLObjectTypePromptSession ACLUserPermissionACLObjectType = "prompt_session"
	ACLUserPermissionACLObjectTypeProjectScore  ACLUserPermissionACLObjectType = "project_score"
	ACLUserPermissionACLObjectTypeProjectTag    ACLUserPermissionACLObjectType = "project_tag"
	ACLUserPermissionACLObjectTypeGroup         ACLUserPermissionACLObjectType = "group"
	ACLUserPermissionACLObjectTypeRole          ACLUserPermissionACLObjectType = "role"
)

func (r ACLUserPermissionACLObjectType) IsKnown() bool {
	switch r {
	case ACLUserPermissionACLObjectTypeOrganization, ACLUserPermissionACLObjectTypeProject, ACLUserPermissionACLObjectTypeExperiment, ACLUserPermissionACLObjectTypeDataset, ACLUserPermissionACLObjectTypePrompt, ACLUserPermissionACLObjectTypePromptSession, ACLUserPermissionACLObjectTypeProjectScore, ACLUserPermissionACLObjectTypeProjectTag, ACLUserPermissionACLObjectTypeGroup, ACLUserPermissionACLObjectTypeRole:
		return true
	}
	return false
}

// Permission the ACL grants
type ACLUserPermissionACLPermission string

const (
	ACLUserPermissionACLPermissionCreate     ACLUserPermissionACLPermission = "create"
	ACLUserPermissionACLPermissionRead       ACLUserPermissionACLPermission = "read"
	ACLUserPermissionACLPermissionUpdate     ACLUserPermissionACLPermission = "update"
	ACLUserPermissionACLPermissionDelete     ACLUserPermissionACLPermission = "delete"
	ACLUserPermissionACLPermissionCreateACLs ACLUserPermissionACLPermission = "create_acls"
	ACLUserPermissionACLPermissionReadACLs   ACLUserPermissionACLPermission = "read_acls"
	ACLUserPermissionACLPermissionUpdateACLs ACLUserPermissionACLPermission = "update_acls"
	ACLUserPermissionACLPermissionDeleteACLs ACLUserPermissionACLPermission = "delete_acls"
)

func (r ACLUserPermissionACLPermission) IsKnown() bool {
	switch r {
	case ACLUserPermissionACLPermissionCreate, ACLUserPermissionACLPermissionRead, ACLUserPermissionACLPermissionUpdate, ACLUserPermissionACLPermissionDelete, ACLUserPermissionACLPermissionCreateACLs, ACLUserPermissionACLPermissionReadACLs, ACLUserPermissionACLPermissionUpdateACLs, ACLUserPermissionACLPermissionDeleteACLs:
		return true
	}
	return false
}

// Optionally restricts the permission grant to just the specified object type
//
// Union satisfied by [ACLUserPermissionACLRestrictObjectTypeString] or
// [ACLUserPermissionACLRestrictObjectTypeObject].
type ACLUserPermissionACLRestrictObjectTypeUnion interface {
	ImplementsACLUserPermissionACLRestrictObjectTypeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ACLUserPermissionACLRestrictObjectTypeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ACLUserPermissionACLRestrictObjectTypeString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ACLUserPermissionACLRestrictObjectTypeObject{}),
		},
	)
}

// The object type that the ACL applies to
type ACLUserPermissionACLRestrictObjectTypeString string

const (
	ACLUserPermissionACLRestrictObjectTypeStringOrganization  ACLUserPermissionACLRestrictObjectTypeString = "organization"
	ACLUserPermissionACLRestrictObjectTypeStringProject       ACLUserPermissionACLRestrictObjectTypeString = "project"
	ACLUserPermissionACLRestrictObjectTypeStringExperiment    ACLUserPermissionACLRestrictObjectTypeString = "experiment"
	ACLUserPermissionACLRestrictObjectTypeStringDataset       ACLUserPermissionACLRestrictObjectTypeString = "dataset"
	ACLUserPermissionACLRestrictObjectTypeStringPrompt        ACLUserPermissionACLRestrictObjectTypeString = "prompt"
	ACLUserPermissionACLRestrictObjectTypeStringPromptSession ACLUserPermissionACLRestrictObjectTypeString = "prompt_session"
	ACLUserPermissionACLRestrictObjectTypeStringProjectScore  ACLUserPermissionACLRestrictObjectTypeString = "project_score"
	ACLUserPermissionACLRestrictObjectTypeStringProjectTag    ACLUserPermissionACLRestrictObjectTypeString = "project_tag"
	ACLUserPermissionACLRestrictObjectTypeStringGroup         ACLUserPermissionACLRestrictObjectTypeString = "group"
	ACLUserPermissionACLRestrictObjectTypeStringRole          ACLUserPermissionACLRestrictObjectTypeString = "role"
)

func (r ACLUserPermissionACLRestrictObjectTypeString) IsKnown() bool {
	switch r {
	case ACLUserPermissionACLRestrictObjectTypeStringOrganization, ACLUserPermissionACLRestrictObjectTypeStringProject, ACLUserPermissionACLRestrictObjectTypeStringExperiment, ACLUserPermissionACLRestrictObjectTypeStringDataset, ACLUserPermissionACLRestrictObjectTypeStringPrompt, ACLUserPermissionACLRestrictObjectTypeStringPromptSession, ACLUserPermissionACLRestrictObjectTypeStringProjectScore, ACLUserPermissionACLRestrictObjectTypeStringProjectTag, ACLUserPermissionACLRestrictObjectTypeStringGroup, ACLUserPermissionACLRestrictObjectTypeStringRole:
		return true
	}
	return false
}

type ACLUserPermissionACLRestrictObjectTypeObject struct {
	// This is just a placeholder nullable object. Only pass null, not the object
	// itself
	ReservedOnlyAllowNull ACLUserPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull `json:"__reserved_only_allow_null,required"`
	JSON                  aclUserPermissionACLRestrictObjectTypeObjectJSON                  `json:"-"`
}

// aclUserPermissionACLRestrictObjectTypeObjectJSON contains the JSON metadata for
// the struct [ACLUserPermissionACLRestrictObjectTypeObject]
type aclUserPermissionACLRestrictObjectTypeObjectJSON struct {
	ReservedOnlyAllowNull apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ACLUserPermissionACLRestrictObjectTypeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclUserPermissionACLRestrictObjectTypeObjectJSON) RawJSON() string {
	return r.raw
}

func (r ACLUserPermissionACLRestrictObjectTypeObject) ImplementsACLUserPermissionACLRestrictObjectTypeUnion() {
}

// This is just a placeholder nullable object. Only pass null, not the object
// itself
type ACLUserPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull struct {
	JSON aclUserPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON `json:"-"`
}

// aclUserPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON contains
// the JSON metadata for the struct
// [ACLUserPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull]
type aclUserPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLUserPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclUserPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON) RawJSON() string {
	return r.raw
}

type ACLUserRoleACL struct {
	// Unique identifier for the acl
	ID string `json:"id,required" format:"uuid"`
	// The organization the ACL's referred object belongs to
	ObjectOrgID string `json:"_object_org_id,required" format:"uuid"`
	// The id of the object the ACL applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType ACLUserRoleACLObjectType `json:"object_type,required"`
	// Id of the role the ACL grants
	RoleID string `json:"role_id,required" format:"uuid"`
	// Id of the user the ACL applies to
	UserID string `json:"user_id,required" format:"uuid"`
	// Date of acl creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Optionally restricts the permission grant to just the specified object type
	RestrictObjectType ACLUserRoleACLRestrictObjectTypeUnion `json:"restrict_object_type"`
	JSON               aclUserRoleACLJSON                    `json:"-"`
}

// aclUserRoleACLJSON contains the JSON metadata for the struct [ACLUserRoleACL]
type aclUserRoleACLJSON struct {
	ID                 apijson.Field
	ObjectOrgID        apijson.Field
	ObjectID           apijson.Field
	ObjectType         apijson.Field
	RoleID             apijson.Field
	UserID             apijson.Field
	Created            apijson.Field
	RestrictObjectType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ACLUserRoleACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclUserRoleACLJSON) RawJSON() string {
	return r.raw
}

func (r ACLUserRoleACL) implementsACL() {}

// The object type that the ACL applies to
type ACLUserRoleACLObjectType string

const (
	ACLUserRoleACLObjectTypeOrganization  ACLUserRoleACLObjectType = "organization"
	ACLUserRoleACLObjectTypeProject       ACLUserRoleACLObjectType = "project"
	ACLUserRoleACLObjectTypeExperiment    ACLUserRoleACLObjectType = "experiment"
	ACLUserRoleACLObjectTypeDataset       ACLUserRoleACLObjectType = "dataset"
	ACLUserRoleACLObjectTypePrompt        ACLUserRoleACLObjectType = "prompt"
	ACLUserRoleACLObjectTypePromptSession ACLUserRoleACLObjectType = "prompt_session"
	ACLUserRoleACLObjectTypeProjectScore  ACLUserRoleACLObjectType = "project_score"
	ACLUserRoleACLObjectTypeProjectTag    ACLUserRoleACLObjectType = "project_tag"
	ACLUserRoleACLObjectTypeGroup         ACLUserRoleACLObjectType = "group"
	ACLUserRoleACLObjectTypeRole          ACLUserRoleACLObjectType = "role"
)

func (r ACLUserRoleACLObjectType) IsKnown() bool {
	switch r {
	case ACLUserRoleACLObjectTypeOrganization, ACLUserRoleACLObjectTypeProject, ACLUserRoleACLObjectTypeExperiment, ACLUserRoleACLObjectTypeDataset, ACLUserRoleACLObjectTypePrompt, ACLUserRoleACLObjectTypePromptSession, ACLUserRoleACLObjectTypeProjectScore, ACLUserRoleACLObjectTypeProjectTag, ACLUserRoleACLObjectTypeGroup, ACLUserRoleACLObjectTypeRole:
		return true
	}
	return false
}

// Optionally restricts the permission grant to just the specified object type
//
// Union satisfied by [ACLUserRoleACLRestrictObjectTypeString] or
// [ACLUserRoleACLRestrictObjectTypeObject].
type ACLUserRoleACLRestrictObjectTypeUnion interface {
	ImplementsACLUserRoleACLRestrictObjectTypeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ACLUserRoleACLRestrictObjectTypeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ACLUserRoleACLRestrictObjectTypeString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ACLUserRoleACLRestrictObjectTypeObject{}),
		},
	)
}

// The object type that the ACL applies to
type ACLUserRoleACLRestrictObjectTypeString string

const (
	ACLUserRoleACLRestrictObjectTypeStringOrganization  ACLUserRoleACLRestrictObjectTypeString = "organization"
	ACLUserRoleACLRestrictObjectTypeStringProject       ACLUserRoleACLRestrictObjectTypeString = "project"
	ACLUserRoleACLRestrictObjectTypeStringExperiment    ACLUserRoleACLRestrictObjectTypeString = "experiment"
	ACLUserRoleACLRestrictObjectTypeStringDataset       ACLUserRoleACLRestrictObjectTypeString = "dataset"
	ACLUserRoleACLRestrictObjectTypeStringPrompt        ACLUserRoleACLRestrictObjectTypeString = "prompt"
	ACLUserRoleACLRestrictObjectTypeStringPromptSession ACLUserRoleACLRestrictObjectTypeString = "prompt_session"
	ACLUserRoleACLRestrictObjectTypeStringProjectScore  ACLUserRoleACLRestrictObjectTypeString = "project_score"
	ACLUserRoleACLRestrictObjectTypeStringProjectTag    ACLUserRoleACLRestrictObjectTypeString = "project_tag"
	ACLUserRoleACLRestrictObjectTypeStringGroup         ACLUserRoleACLRestrictObjectTypeString = "group"
	ACLUserRoleACLRestrictObjectTypeStringRole          ACLUserRoleACLRestrictObjectTypeString = "role"
)

func (r ACLUserRoleACLRestrictObjectTypeString) IsKnown() bool {
	switch r {
	case ACLUserRoleACLRestrictObjectTypeStringOrganization, ACLUserRoleACLRestrictObjectTypeStringProject, ACLUserRoleACLRestrictObjectTypeStringExperiment, ACLUserRoleACLRestrictObjectTypeStringDataset, ACLUserRoleACLRestrictObjectTypeStringPrompt, ACLUserRoleACLRestrictObjectTypeStringPromptSession, ACLUserRoleACLRestrictObjectTypeStringProjectScore, ACLUserRoleACLRestrictObjectTypeStringProjectTag, ACLUserRoleACLRestrictObjectTypeStringGroup, ACLUserRoleACLRestrictObjectTypeStringRole:
		return true
	}
	return false
}

type ACLUserRoleACLRestrictObjectTypeObject struct {
	// This is just a placeholder nullable object. Only pass null, not the object
	// itself
	ReservedOnlyAllowNull ACLUserRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull `json:"__reserved_only_allow_null,required"`
	JSON                  aclUserRoleACLRestrictObjectTypeObjectJSON                  `json:"-"`
}

// aclUserRoleACLRestrictObjectTypeObjectJSON contains the JSON metadata for the
// struct [ACLUserRoleACLRestrictObjectTypeObject]
type aclUserRoleACLRestrictObjectTypeObjectJSON struct {
	ReservedOnlyAllowNull apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ACLUserRoleACLRestrictObjectTypeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclUserRoleACLRestrictObjectTypeObjectJSON) RawJSON() string {
	return r.raw
}

func (r ACLUserRoleACLRestrictObjectTypeObject) ImplementsACLUserRoleACLRestrictObjectTypeUnion() {}

// This is just a placeholder nullable object. Only pass null, not the object
// itself
type ACLUserRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull struct {
	JSON aclUserRoleACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON `json:"-"`
}

// aclUserRoleACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON contains the
// JSON metadata for the struct
// [ACLUserRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull]
type aclUserRoleACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLUserRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclUserRoleACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON) RawJSON() string {
	return r.raw
}

type ACLGroupPermissionACL struct {
	// Unique identifier for the acl
	ID string `json:"id,required" format:"uuid"`
	// The organization the ACL's referred object belongs to
	ObjectOrgID string `json:"_object_org_id,required" format:"uuid"`
	// Id of the group the ACL applies to
	GroupID string `json:"group_id,required" format:"uuid"`
	// The id of the object the ACL applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType ACLGroupPermissionACLObjectType `json:"object_type,required"`
	// Permission the ACL grants
	Permission ACLGroupPermissionACLPermission `json:"permission,required"`
	// Date of acl creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Optionally restricts the permission grant to just the specified object type
	RestrictObjectType ACLGroupPermissionACLRestrictObjectTypeUnion `json:"restrict_object_type"`
	JSON               aclGroupPermissionACLJSON                    `json:"-"`
}

// aclGroupPermissionACLJSON contains the JSON metadata for the struct
// [ACLGroupPermissionACL]
type aclGroupPermissionACLJSON struct {
	ID                 apijson.Field
	ObjectOrgID        apijson.Field
	GroupID            apijson.Field
	ObjectID           apijson.Field
	ObjectType         apijson.Field
	Permission         apijson.Field
	Created            apijson.Field
	RestrictObjectType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ACLGroupPermissionACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclGroupPermissionACLJSON) RawJSON() string {
	return r.raw
}

func (r ACLGroupPermissionACL) implementsACL() {}

// The object type that the ACL applies to
type ACLGroupPermissionACLObjectType string

const (
	ACLGroupPermissionACLObjectTypeOrganization  ACLGroupPermissionACLObjectType = "organization"
	ACLGroupPermissionACLObjectTypeProject       ACLGroupPermissionACLObjectType = "project"
	ACLGroupPermissionACLObjectTypeExperiment    ACLGroupPermissionACLObjectType = "experiment"
	ACLGroupPermissionACLObjectTypeDataset       ACLGroupPermissionACLObjectType = "dataset"
	ACLGroupPermissionACLObjectTypePrompt        ACLGroupPermissionACLObjectType = "prompt"
	ACLGroupPermissionACLObjectTypePromptSession ACLGroupPermissionACLObjectType = "prompt_session"
	ACLGroupPermissionACLObjectTypeProjectScore  ACLGroupPermissionACLObjectType = "project_score"
	ACLGroupPermissionACLObjectTypeProjectTag    ACLGroupPermissionACLObjectType = "project_tag"
	ACLGroupPermissionACLObjectTypeGroup         ACLGroupPermissionACLObjectType = "group"
	ACLGroupPermissionACLObjectTypeRole          ACLGroupPermissionACLObjectType = "role"
)

func (r ACLGroupPermissionACLObjectType) IsKnown() bool {
	switch r {
	case ACLGroupPermissionACLObjectTypeOrganization, ACLGroupPermissionACLObjectTypeProject, ACLGroupPermissionACLObjectTypeExperiment, ACLGroupPermissionACLObjectTypeDataset, ACLGroupPermissionACLObjectTypePrompt, ACLGroupPermissionACLObjectTypePromptSession, ACLGroupPermissionACLObjectTypeProjectScore, ACLGroupPermissionACLObjectTypeProjectTag, ACLGroupPermissionACLObjectTypeGroup, ACLGroupPermissionACLObjectTypeRole:
		return true
	}
	return false
}

// Permission the ACL grants
type ACLGroupPermissionACLPermission string

const (
	ACLGroupPermissionACLPermissionCreate     ACLGroupPermissionACLPermission = "create"
	ACLGroupPermissionACLPermissionRead       ACLGroupPermissionACLPermission = "read"
	ACLGroupPermissionACLPermissionUpdate     ACLGroupPermissionACLPermission = "update"
	ACLGroupPermissionACLPermissionDelete     ACLGroupPermissionACLPermission = "delete"
	ACLGroupPermissionACLPermissionCreateACLs ACLGroupPermissionACLPermission = "create_acls"
	ACLGroupPermissionACLPermissionReadACLs   ACLGroupPermissionACLPermission = "read_acls"
	ACLGroupPermissionACLPermissionUpdateACLs ACLGroupPermissionACLPermission = "update_acls"
	ACLGroupPermissionACLPermissionDeleteACLs ACLGroupPermissionACLPermission = "delete_acls"
)

func (r ACLGroupPermissionACLPermission) IsKnown() bool {
	switch r {
	case ACLGroupPermissionACLPermissionCreate, ACLGroupPermissionACLPermissionRead, ACLGroupPermissionACLPermissionUpdate, ACLGroupPermissionACLPermissionDelete, ACLGroupPermissionACLPermissionCreateACLs, ACLGroupPermissionACLPermissionReadACLs, ACLGroupPermissionACLPermissionUpdateACLs, ACLGroupPermissionACLPermissionDeleteACLs:
		return true
	}
	return false
}

// Optionally restricts the permission grant to just the specified object type
//
// Union satisfied by [ACLGroupPermissionACLRestrictObjectTypeString] or
// [ACLGroupPermissionACLRestrictObjectTypeObject].
type ACLGroupPermissionACLRestrictObjectTypeUnion interface {
	ImplementsACLGroupPermissionACLRestrictObjectTypeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ACLGroupPermissionACLRestrictObjectTypeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ACLGroupPermissionACLRestrictObjectTypeString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ACLGroupPermissionACLRestrictObjectTypeObject{}),
		},
	)
}

// The object type that the ACL applies to
type ACLGroupPermissionACLRestrictObjectTypeString string

const (
	ACLGroupPermissionACLRestrictObjectTypeStringOrganization  ACLGroupPermissionACLRestrictObjectTypeString = "organization"
	ACLGroupPermissionACLRestrictObjectTypeStringProject       ACLGroupPermissionACLRestrictObjectTypeString = "project"
	ACLGroupPermissionACLRestrictObjectTypeStringExperiment    ACLGroupPermissionACLRestrictObjectTypeString = "experiment"
	ACLGroupPermissionACLRestrictObjectTypeStringDataset       ACLGroupPermissionACLRestrictObjectTypeString = "dataset"
	ACLGroupPermissionACLRestrictObjectTypeStringPrompt        ACLGroupPermissionACLRestrictObjectTypeString = "prompt"
	ACLGroupPermissionACLRestrictObjectTypeStringPromptSession ACLGroupPermissionACLRestrictObjectTypeString = "prompt_session"
	ACLGroupPermissionACLRestrictObjectTypeStringProjectScore  ACLGroupPermissionACLRestrictObjectTypeString = "project_score"
	ACLGroupPermissionACLRestrictObjectTypeStringProjectTag    ACLGroupPermissionACLRestrictObjectTypeString = "project_tag"
	ACLGroupPermissionACLRestrictObjectTypeStringGroup         ACLGroupPermissionACLRestrictObjectTypeString = "group"
	ACLGroupPermissionACLRestrictObjectTypeStringRole          ACLGroupPermissionACLRestrictObjectTypeString = "role"
)

func (r ACLGroupPermissionACLRestrictObjectTypeString) IsKnown() bool {
	switch r {
	case ACLGroupPermissionACLRestrictObjectTypeStringOrganization, ACLGroupPermissionACLRestrictObjectTypeStringProject, ACLGroupPermissionACLRestrictObjectTypeStringExperiment, ACLGroupPermissionACLRestrictObjectTypeStringDataset, ACLGroupPermissionACLRestrictObjectTypeStringPrompt, ACLGroupPermissionACLRestrictObjectTypeStringPromptSession, ACLGroupPermissionACLRestrictObjectTypeStringProjectScore, ACLGroupPermissionACLRestrictObjectTypeStringProjectTag, ACLGroupPermissionACLRestrictObjectTypeStringGroup, ACLGroupPermissionACLRestrictObjectTypeStringRole:
		return true
	}
	return false
}

type ACLGroupPermissionACLRestrictObjectTypeObject struct {
	// This is just a placeholder nullable object. Only pass null, not the object
	// itself
	ReservedOnlyAllowNull ACLGroupPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull `json:"__reserved_only_allow_null,required"`
	JSON                  aclGroupPermissionACLRestrictObjectTypeObjectJSON                  `json:"-"`
}

// aclGroupPermissionACLRestrictObjectTypeObjectJSON contains the JSON metadata for
// the struct [ACLGroupPermissionACLRestrictObjectTypeObject]
type aclGroupPermissionACLRestrictObjectTypeObjectJSON struct {
	ReservedOnlyAllowNull apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ACLGroupPermissionACLRestrictObjectTypeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclGroupPermissionACLRestrictObjectTypeObjectJSON) RawJSON() string {
	return r.raw
}

func (r ACLGroupPermissionACLRestrictObjectTypeObject) ImplementsACLGroupPermissionACLRestrictObjectTypeUnion() {
}

// This is just a placeholder nullable object. Only pass null, not the object
// itself
type ACLGroupPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull struct {
	JSON aclGroupPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON `json:"-"`
}

// aclGroupPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON contains
// the JSON metadata for the struct
// [ACLGroupPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull]
type aclGroupPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLGroupPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclGroupPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON) RawJSON() string {
	return r.raw
}

type ACLGroupRoleACL struct {
	// Unique identifier for the acl
	ID string `json:"id,required" format:"uuid"`
	// The organization the ACL's referred object belongs to
	ObjectOrgID string `json:"_object_org_id,required" format:"uuid"`
	// Id of the group the ACL applies to
	GroupID string `json:"group_id,required" format:"uuid"`
	// The id of the object the ACL applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType ACLGroupRoleACLObjectType `json:"object_type,required"`
	// Id of the role the ACL grants
	RoleID string `json:"role_id,required" format:"uuid"`
	// Date of acl creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Optionally restricts the permission grant to just the specified object type
	RestrictObjectType ACLGroupRoleACLRestrictObjectTypeUnion `json:"restrict_object_type"`
	JSON               aclGroupRoleACLJSON                    `json:"-"`
}

// aclGroupRoleACLJSON contains the JSON metadata for the struct [ACLGroupRoleACL]
type aclGroupRoleACLJSON struct {
	ID                 apijson.Field
	ObjectOrgID        apijson.Field
	GroupID            apijson.Field
	ObjectID           apijson.Field
	ObjectType         apijson.Field
	RoleID             apijson.Field
	Created            apijson.Field
	RestrictObjectType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ACLGroupRoleACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclGroupRoleACLJSON) RawJSON() string {
	return r.raw
}

func (r ACLGroupRoleACL) implementsACL() {}

// The object type that the ACL applies to
type ACLGroupRoleACLObjectType string

const (
	ACLGroupRoleACLObjectTypeOrganization  ACLGroupRoleACLObjectType = "organization"
	ACLGroupRoleACLObjectTypeProject       ACLGroupRoleACLObjectType = "project"
	ACLGroupRoleACLObjectTypeExperiment    ACLGroupRoleACLObjectType = "experiment"
	ACLGroupRoleACLObjectTypeDataset       ACLGroupRoleACLObjectType = "dataset"
	ACLGroupRoleACLObjectTypePrompt        ACLGroupRoleACLObjectType = "prompt"
	ACLGroupRoleACLObjectTypePromptSession ACLGroupRoleACLObjectType = "prompt_session"
	ACLGroupRoleACLObjectTypeProjectScore  ACLGroupRoleACLObjectType = "project_score"
	ACLGroupRoleACLObjectTypeProjectTag    ACLGroupRoleACLObjectType = "project_tag"
	ACLGroupRoleACLObjectTypeGroup         ACLGroupRoleACLObjectType = "group"
	ACLGroupRoleACLObjectTypeRole          ACLGroupRoleACLObjectType = "role"
)

func (r ACLGroupRoleACLObjectType) IsKnown() bool {
	switch r {
	case ACLGroupRoleACLObjectTypeOrganization, ACLGroupRoleACLObjectTypeProject, ACLGroupRoleACLObjectTypeExperiment, ACLGroupRoleACLObjectTypeDataset, ACLGroupRoleACLObjectTypePrompt, ACLGroupRoleACLObjectTypePromptSession, ACLGroupRoleACLObjectTypeProjectScore, ACLGroupRoleACLObjectTypeProjectTag, ACLGroupRoleACLObjectTypeGroup, ACLGroupRoleACLObjectTypeRole:
		return true
	}
	return false
}

// Optionally restricts the permission grant to just the specified object type
//
// Union satisfied by [ACLGroupRoleACLRestrictObjectTypeString] or
// [ACLGroupRoleACLRestrictObjectTypeObject].
type ACLGroupRoleACLRestrictObjectTypeUnion interface {
	ImplementsACLGroupRoleACLRestrictObjectTypeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ACLGroupRoleACLRestrictObjectTypeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ACLGroupRoleACLRestrictObjectTypeString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ACLGroupRoleACLRestrictObjectTypeObject{}),
		},
	)
}

// The object type that the ACL applies to
type ACLGroupRoleACLRestrictObjectTypeString string

const (
	ACLGroupRoleACLRestrictObjectTypeStringOrganization  ACLGroupRoleACLRestrictObjectTypeString = "organization"
	ACLGroupRoleACLRestrictObjectTypeStringProject       ACLGroupRoleACLRestrictObjectTypeString = "project"
	ACLGroupRoleACLRestrictObjectTypeStringExperiment    ACLGroupRoleACLRestrictObjectTypeString = "experiment"
	ACLGroupRoleACLRestrictObjectTypeStringDataset       ACLGroupRoleACLRestrictObjectTypeString = "dataset"
	ACLGroupRoleACLRestrictObjectTypeStringPrompt        ACLGroupRoleACLRestrictObjectTypeString = "prompt"
	ACLGroupRoleACLRestrictObjectTypeStringPromptSession ACLGroupRoleACLRestrictObjectTypeString = "prompt_session"
	ACLGroupRoleACLRestrictObjectTypeStringProjectScore  ACLGroupRoleACLRestrictObjectTypeString = "project_score"
	ACLGroupRoleACLRestrictObjectTypeStringProjectTag    ACLGroupRoleACLRestrictObjectTypeString = "project_tag"
	ACLGroupRoleACLRestrictObjectTypeStringGroup         ACLGroupRoleACLRestrictObjectTypeString = "group"
	ACLGroupRoleACLRestrictObjectTypeStringRole          ACLGroupRoleACLRestrictObjectTypeString = "role"
)

func (r ACLGroupRoleACLRestrictObjectTypeString) IsKnown() bool {
	switch r {
	case ACLGroupRoleACLRestrictObjectTypeStringOrganization, ACLGroupRoleACLRestrictObjectTypeStringProject, ACLGroupRoleACLRestrictObjectTypeStringExperiment, ACLGroupRoleACLRestrictObjectTypeStringDataset, ACLGroupRoleACLRestrictObjectTypeStringPrompt, ACLGroupRoleACLRestrictObjectTypeStringPromptSession, ACLGroupRoleACLRestrictObjectTypeStringProjectScore, ACLGroupRoleACLRestrictObjectTypeStringProjectTag, ACLGroupRoleACLRestrictObjectTypeStringGroup, ACLGroupRoleACLRestrictObjectTypeStringRole:
		return true
	}
	return false
}

type ACLGroupRoleACLRestrictObjectTypeObject struct {
	// This is just a placeholder nullable object. Only pass null, not the object
	// itself
	ReservedOnlyAllowNull ACLGroupRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull `json:"__reserved_only_allow_null,required"`
	JSON                  aclGroupRoleACLRestrictObjectTypeObjectJSON                  `json:"-"`
}

// aclGroupRoleACLRestrictObjectTypeObjectJSON contains the JSON metadata for the
// struct [ACLGroupRoleACLRestrictObjectTypeObject]
type aclGroupRoleACLRestrictObjectTypeObjectJSON struct {
	ReservedOnlyAllowNull apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ACLGroupRoleACLRestrictObjectTypeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclGroupRoleACLRestrictObjectTypeObjectJSON) RawJSON() string {
	return r.raw
}

func (r ACLGroupRoleACLRestrictObjectTypeObject) ImplementsACLGroupRoleACLRestrictObjectTypeUnion() {}

// This is just a placeholder nullable object. Only pass null, not the object
// itself
type ACLGroupRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull struct {
	JSON aclGroupRoleACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON `json:"-"`
}

// aclGroupRoleACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON contains the
// JSON metadata for the struct
// [ACLGroupRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull]
type aclGroupRoleACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLGroupRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclGroupRoleACLRestrictObjectTypeObjectReservedOnlyAllowNullJSON) RawJSON() string {
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
	ACLObjectTypeProjectScore  ACLObjectType = "project_score"
	ACLObjectTypeProjectTag    ACLObjectType = "project_tag"
	ACLObjectTypeGroup         ACLObjectType = "group"
	ACLObjectTypeRole          ACLObjectType = "role"
)

func (r ACLObjectType) IsKnown() bool {
	switch r {
	case ACLObjectTypeOrganization, ACLObjectTypeProject, ACLObjectTypeExperiment, ACLObjectTypeDataset, ACLObjectTypePrompt, ACLObjectTypePromptSession, ACLObjectTypeProjectScore, ACLObjectTypeProjectTag, ACLObjectTypeGroup, ACLObjectTypeRole:
		return true
	}
	return false
}

// Permission the ACL grants
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

type ACLNewParams struct {
	Body ACLNewParamsBodyUnion `json:"body,required"`
}

func (r ACLNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ACLNewParamsBody struct {
	// The object type that the ACL applies to
	ObjectType param.Field[ACLNewParamsBodyObjectType] `json:"object_type,required"`
	// The id of the object the ACL applies to
	ObjectID           param.Field[string]      `json:"object_id,required" format:"uuid"`
	RestrictObjectType param.Field[interface{}] `json:"restrict_object_type,required"`
	// Id of the user the ACL applies to
	UserID param.Field[string] `json:"user_id" format:"uuid"`
	// Permission the ACL grants
	Permission param.Field[ACLNewParamsBodyPermission] `json:"permission"`
	// Id of the role the ACL grants
	RoleID param.Field[string] `json:"role_id" format:"uuid"`
	// Id of the group the ACL applies to
	GroupID param.Field[string] `json:"group_id" format:"uuid"`
}

func (r ACLNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLNewParamsBody) implementsACLNewParamsBodyUnion() {}

// Satisfied by [ACLNewParamsBodyCreateUserPermissionACL],
// [ACLNewParamsBodyCreateUserRoleACL], [ACLNewParamsBodyCreateGroupPermissionACL],
// [ACLNewParamsBodyCreateGroupRoleACL], [ACLNewParamsBody].
type ACLNewParamsBodyUnion interface {
	implementsACLNewParamsBodyUnion()
}

type ACLNewParamsBodyCreateUserPermissionACL struct {
	// The id of the object the ACL applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ACLNewParamsBodyCreateUserPermissionACLObjectType] `json:"object_type,required"`
	// Permission the ACL grants
	Permission param.Field[ACLNewParamsBodyCreateUserPermissionACLPermission] `json:"permission,required"`
	// Id of the user the ACL applies to
	UserID param.Field[string] `json:"user_id,required" format:"uuid"`
	// Optionally restricts the permission grant to just the specified object type
	RestrictObjectType param.Field[ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeUnion] `json:"restrict_object_type"`
}

func (r ACLNewParamsBodyCreateUserPermissionACL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLNewParamsBodyCreateUserPermissionACL) implementsACLNewParamsBodyUnion() {}

// The object type that the ACL applies to
type ACLNewParamsBodyCreateUserPermissionACLObjectType string

const (
	ACLNewParamsBodyCreateUserPermissionACLObjectTypeOrganization  ACLNewParamsBodyCreateUserPermissionACLObjectType = "organization"
	ACLNewParamsBodyCreateUserPermissionACLObjectTypeProject       ACLNewParamsBodyCreateUserPermissionACLObjectType = "project"
	ACLNewParamsBodyCreateUserPermissionACLObjectTypeExperiment    ACLNewParamsBodyCreateUserPermissionACLObjectType = "experiment"
	ACLNewParamsBodyCreateUserPermissionACLObjectTypeDataset       ACLNewParamsBodyCreateUserPermissionACLObjectType = "dataset"
	ACLNewParamsBodyCreateUserPermissionACLObjectTypePrompt        ACLNewParamsBodyCreateUserPermissionACLObjectType = "prompt"
	ACLNewParamsBodyCreateUserPermissionACLObjectTypePromptSession ACLNewParamsBodyCreateUserPermissionACLObjectType = "prompt_session"
	ACLNewParamsBodyCreateUserPermissionACLObjectTypeProjectScore  ACLNewParamsBodyCreateUserPermissionACLObjectType = "project_score"
	ACLNewParamsBodyCreateUserPermissionACLObjectTypeProjectTag    ACLNewParamsBodyCreateUserPermissionACLObjectType = "project_tag"
	ACLNewParamsBodyCreateUserPermissionACLObjectTypeGroup         ACLNewParamsBodyCreateUserPermissionACLObjectType = "group"
	ACLNewParamsBodyCreateUserPermissionACLObjectTypeRole          ACLNewParamsBodyCreateUserPermissionACLObjectType = "role"
)

func (r ACLNewParamsBodyCreateUserPermissionACLObjectType) IsKnown() bool {
	switch r {
	case ACLNewParamsBodyCreateUserPermissionACLObjectTypeOrganization, ACLNewParamsBodyCreateUserPermissionACLObjectTypeProject, ACLNewParamsBodyCreateUserPermissionACLObjectTypeExperiment, ACLNewParamsBodyCreateUserPermissionACLObjectTypeDataset, ACLNewParamsBodyCreateUserPermissionACLObjectTypePrompt, ACLNewParamsBodyCreateUserPermissionACLObjectTypePromptSession, ACLNewParamsBodyCreateUserPermissionACLObjectTypeProjectScore, ACLNewParamsBodyCreateUserPermissionACLObjectTypeProjectTag, ACLNewParamsBodyCreateUserPermissionACLObjectTypeGroup, ACLNewParamsBodyCreateUserPermissionACLObjectTypeRole:
		return true
	}
	return false
}

// Permission the ACL grants
type ACLNewParamsBodyCreateUserPermissionACLPermission string

const (
	ACLNewParamsBodyCreateUserPermissionACLPermissionCreate     ACLNewParamsBodyCreateUserPermissionACLPermission = "create"
	ACLNewParamsBodyCreateUserPermissionACLPermissionRead       ACLNewParamsBodyCreateUserPermissionACLPermission = "read"
	ACLNewParamsBodyCreateUserPermissionACLPermissionUpdate     ACLNewParamsBodyCreateUserPermissionACLPermission = "update"
	ACLNewParamsBodyCreateUserPermissionACLPermissionDelete     ACLNewParamsBodyCreateUserPermissionACLPermission = "delete"
	ACLNewParamsBodyCreateUserPermissionACLPermissionCreateACLs ACLNewParamsBodyCreateUserPermissionACLPermission = "create_acls"
	ACLNewParamsBodyCreateUserPermissionACLPermissionReadACLs   ACLNewParamsBodyCreateUserPermissionACLPermission = "read_acls"
	ACLNewParamsBodyCreateUserPermissionACLPermissionUpdateACLs ACLNewParamsBodyCreateUserPermissionACLPermission = "update_acls"
	ACLNewParamsBodyCreateUserPermissionACLPermissionDeleteACLs ACLNewParamsBodyCreateUserPermissionACLPermission = "delete_acls"
)

func (r ACLNewParamsBodyCreateUserPermissionACLPermission) IsKnown() bool {
	switch r {
	case ACLNewParamsBodyCreateUserPermissionACLPermissionCreate, ACLNewParamsBodyCreateUserPermissionACLPermissionRead, ACLNewParamsBodyCreateUserPermissionACLPermissionUpdate, ACLNewParamsBodyCreateUserPermissionACLPermissionDelete, ACLNewParamsBodyCreateUserPermissionACLPermissionCreateACLs, ACLNewParamsBodyCreateUserPermissionACLPermissionReadACLs, ACLNewParamsBodyCreateUserPermissionACLPermissionUpdateACLs, ACLNewParamsBodyCreateUserPermissionACLPermissionDeleteACLs:
		return true
	}
	return false
}

// Optionally restricts the permission grant to just the specified object type
type ACLNewParamsBodyCreateUserPermissionACLRestrictObjectType struct {
	ReservedOnlyAllowNull param.Field[interface{}] `json:"__reserved_only_allow_null"`
}

func (r ACLNewParamsBodyCreateUserPermissionACLRestrictObjectType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLNewParamsBodyCreateUserPermissionACLRestrictObjectType) ImplementsACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeUnion() {
}

// Optionally restricts the permission grant to just the specified object type
//
// Satisfied by [ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeString],
// [ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeObject],
// [ACLNewParamsBodyCreateUserPermissionACLRestrictObjectType].
type ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeUnion interface {
	ImplementsACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeUnion()
}

// The object type that the ACL applies to
type ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeString string

const (
	ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringOrganization  ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "organization"
	ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringProject       ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "project"
	ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringExperiment    ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "experiment"
	ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringDataset       ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "dataset"
	ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringPrompt        ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "prompt"
	ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringPromptSession ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "prompt_session"
	ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringProjectScore  ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "project_score"
	ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringProjectTag    ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "project_tag"
	ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringGroup         ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "group"
	ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringRole          ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "role"
)

func (r ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeString) IsKnown() bool {
	switch r {
	case ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringOrganization, ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringProject, ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringExperiment, ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringDataset, ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringPrompt, ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringPromptSession, ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringProjectScore, ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringProjectTag, ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringGroup, ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeStringRole:
		return true
	}
	return false
}

type ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeObject struct {
	// This is just a placeholder nullable object. Only pass null, not the object
	// itself
	ReservedOnlyAllowNull param.Field[ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull] `json:"__reserved_only_allow_null,required"`
}

func (r ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeObject) ImplementsACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeUnion() {
}

// This is just a placeholder nullable object. Only pass null, not the object
// itself
type ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull struct {
}

func (r ACLNewParamsBodyCreateUserPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ACLNewParamsBodyCreateUserRoleACL struct {
	// The id of the object the ACL applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ACLNewParamsBodyCreateUserRoleACLObjectType] `json:"object_type,required"`
	// Id of the role the ACL grants
	RoleID param.Field[string] `json:"role_id,required" format:"uuid"`
	// Id of the user the ACL applies to
	UserID param.Field[string] `json:"user_id,required" format:"uuid"`
	// Optionally restricts the permission grant to just the specified object type
	RestrictObjectType param.Field[ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeUnion] `json:"restrict_object_type"`
}

func (r ACLNewParamsBodyCreateUserRoleACL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLNewParamsBodyCreateUserRoleACL) implementsACLNewParamsBodyUnion() {}

// The object type that the ACL applies to
type ACLNewParamsBodyCreateUserRoleACLObjectType string

const (
	ACLNewParamsBodyCreateUserRoleACLObjectTypeOrganization  ACLNewParamsBodyCreateUserRoleACLObjectType = "organization"
	ACLNewParamsBodyCreateUserRoleACLObjectTypeProject       ACLNewParamsBodyCreateUserRoleACLObjectType = "project"
	ACLNewParamsBodyCreateUserRoleACLObjectTypeExperiment    ACLNewParamsBodyCreateUserRoleACLObjectType = "experiment"
	ACLNewParamsBodyCreateUserRoleACLObjectTypeDataset       ACLNewParamsBodyCreateUserRoleACLObjectType = "dataset"
	ACLNewParamsBodyCreateUserRoleACLObjectTypePrompt        ACLNewParamsBodyCreateUserRoleACLObjectType = "prompt"
	ACLNewParamsBodyCreateUserRoleACLObjectTypePromptSession ACLNewParamsBodyCreateUserRoleACLObjectType = "prompt_session"
	ACLNewParamsBodyCreateUserRoleACLObjectTypeProjectScore  ACLNewParamsBodyCreateUserRoleACLObjectType = "project_score"
	ACLNewParamsBodyCreateUserRoleACLObjectTypeProjectTag    ACLNewParamsBodyCreateUserRoleACLObjectType = "project_tag"
	ACLNewParamsBodyCreateUserRoleACLObjectTypeGroup         ACLNewParamsBodyCreateUserRoleACLObjectType = "group"
	ACLNewParamsBodyCreateUserRoleACLObjectTypeRole          ACLNewParamsBodyCreateUserRoleACLObjectType = "role"
)

func (r ACLNewParamsBodyCreateUserRoleACLObjectType) IsKnown() bool {
	switch r {
	case ACLNewParamsBodyCreateUserRoleACLObjectTypeOrganization, ACLNewParamsBodyCreateUserRoleACLObjectTypeProject, ACLNewParamsBodyCreateUserRoleACLObjectTypeExperiment, ACLNewParamsBodyCreateUserRoleACLObjectTypeDataset, ACLNewParamsBodyCreateUserRoleACLObjectTypePrompt, ACLNewParamsBodyCreateUserRoleACLObjectTypePromptSession, ACLNewParamsBodyCreateUserRoleACLObjectTypeProjectScore, ACLNewParamsBodyCreateUserRoleACLObjectTypeProjectTag, ACLNewParamsBodyCreateUserRoleACLObjectTypeGroup, ACLNewParamsBodyCreateUserRoleACLObjectTypeRole:
		return true
	}
	return false
}

// Optionally restricts the permission grant to just the specified object type
type ACLNewParamsBodyCreateUserRoleACLRestrictObjectType struct {
	ReservedOnlyAllowNull param.Field[interface{}] `json:"__reserved_only_allow_null"`
}

func (r ACLNewParamsBodyCreateUserRoleACLRestrictObjectType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLNewParamsBodyCreateUserRoleACLRestrictObjectType) ImplementsACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeUnion() {
}

// Optionally restricts the permission grant to just the specified object type
//
// Satisfied by [ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeString],
// [ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeObject],
// [ACLNewParamsBodyCreateUserRoleACLRestrictObjectType].
type ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeUnion interface {
	ImplementsACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeUnion()
}

// The object type that the ACL applies to
type ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeString string

const (
	ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringOrganization  ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeString = "organization"
	ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringProject       ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeString = "project"
	ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringExperiment    ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeString = "experiment"
	ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringDataset       ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeString = "dataset"
	ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringPrompt        ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeString = "prompt"
	ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringPromptSession ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeString = "prompt_session"
	ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringProjectScore  ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeString = "project_score"
	ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringProjectTag    ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeString = "project_tag"
	ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringGroup         ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeString = "group"
	ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringRole          ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeString = "role"
)

func (r ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeString) IsKnown() bool {
	switch r {
	case ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringOrganization, ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringProject, ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringExperiment, ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringDataset, ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringPrompt, ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringPromptSession, ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringProjectScore, ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringProjectTag, ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringGroup, ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeStringRole:
		return true
	}
	return false
}

type ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeObject struct {
	// This is just a placeholder nullable object. Only pass null, not the object
	// itself
	ReservedOnlyAllowNull param.Field[ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull] `json:"__reserved_only_allow_null,required"`
}

func (r ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeObject) ImplementsACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeUnion() {
}

// This is just a placeholder nullable object. Only pass null, not the object
// itself
type ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull struct {
}

func (r ACLNewParamsBodyCreateUserRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ACLNewParamsBodyCreateGroupPermissionACL struct {
	// Id of the group the ACL applies to
	GroupID param.Field[string] `json:"group_id,required" format:"uuid"`
	// The id of the object the ACL applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ACLNewParamsBodyCreateGroupPermissionACLObjectType] `json:"object_type,required"`
	// Permission the ACL grants
	Permission param.Field[ACLNewParamsBodyCreateGroupPermissionACLPermission] `json:"permission,required"`
	// Optionally restricts the permission grant to just the specified object type
	RestrictObjectType param.Field[ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeUnion] `json:"restrict_object_type"`
}

func (r ACLNewParamsBodyCreateGroupPermissionACL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLNewParamsBodyCreateGroupPermissionACL) implementsACLNewParamsBodyUnion() {}

// The object type that the ACL applies to
type ACLNewParamsBodyCreateGroupPermissionACLObjectType string

const (
	ACLNewParamsBodyCreateGroupPermissionACLObjectTypeOrganization  ACLNewParamsBodyCreateGroupPermissionACLObjectType = "organization"
	ACLNewParamsBodyCreateGroupPermissionACLObjectTypeProject       ACLNewParamsBodyCreateGroupPermissionACLObjectType = "project"
	ACLNewParamsBodyCreateGroupPermissionACLObjectTypeExperiment    ACLNewParamsBodyCreateGroupPermissionACLObjectType = "experiment"
	ACLNewParamsBodyCreateGroupPermissionACLObjectTypeDataset       ACLNewParamsBodyCreateGroupPermissionACLObjectType = "dataset"
	ACLNewParamsBodyCreateGroupPermissionACLObjectTypePrompt        ACLNewParamsBodyCreateGroupPermissionACLObjectType = "prompt"
	ACLNewParamsBodyCreateGroupPermissionACLObjectTypePromptSession ACLNewParamsBodyCreateGroupPermissionACLObjectType = "prompt_session"
	ACLNewParamsBodyCreateGroupPermissionACLObjectTypeProjectScore  ACLNewParamsBodyCreateGroupPermissionACLObjectType = "project_score"
	ACLNewParamsBodyCreateGroupPermissionACLObjectTypeProjectTag    ACLNewParamsBodyCreateGroupPermissionACLObjectType = "project_tag"
	ACLNewParamsBodyCreateGroupPermissionACLObjectTypeGroup         ACLNewParamsBodyCreateGroupPermissionACLObjectType = "group"
	ACLNewParamsBodyCreateGroupPermissionACLObjectTypeRole          ACLNewParamsBodyCreateGroupPermissionACLObjectType = "role"
)

func (r ACLNewParamsBodyCreateGroupPermissionACLObjectType) IsKnown() bool {
	switch r {
	case ACLNewParamsBodyCreateGroupPermissionACLObjectTypeOrganization, ACLNewParamsBodyCreateGroupPermissionACLObjectTypeProject, ACLNewParamsBodyCreateGroupPermissionACLObjectTypeExperiment, ACLNewParamsBodyCreateGroupPermissionACLObjectTypeDataset, ACLNewParamsBodyCreateGroupPermissionACLObjectTypePrompt, ACLNewParamsBodyCreateGroupPermissionACLObjectTypePromptSession, ACLNewParamsBodyCreateGroupPermissionACLObjectTypeProjectScore, ACLNewParamsBodyCreateGroupPermissionACLObjectTypeProjectTag, ACLNewParamsBodyCreateGroupPermissionACLObjectTypeGroup, ACLNewParamsBodyCreateGroupPermissionACLObjectTypeRole:
		return true
	}
	return false
}

// Permission the ACL grants
type ACLNewParamsBodyCreateGroupPermissionACLPermission string

const (
	ACLNewParamsBodyCreateGroupPermissionACLPermissionCreate     ACLNewParamsBodyCreateGroupPermissionACLPermission = "create"
	ACLNewParamsBodyCreateGroupPermissionACLPermissionRead       ACLNewParamsBodyCreateGroupPermissionACLPermission = "read"
	ACLNewParamsBodyCreateGroupPermissionACLPermissionUpdate     ACLNewParamsBodyCreateGroupPermissionACLPermission = "update"
	ACLNewParamsBodyCreateGroupPermissionACLPermissionDelete     ACLNewParamsBodyCreateGroupPermissionACLPermission = "delete"
	ACLNewParamsBodyCreateGroupPermissionACLPermissionCreateACLs ACLNewParamsBodyCreateGroupPermissionACLPermission = "create_acls"
	ACLNewParamsBodyCreateGroupPermissionACLPermissionReadACLs   ACLNewParamsBodyCreateGroupPermissionACLPermission = "read_acls"
	ACLNewParamsBodyCreateGroupPermissionACLPermissionUpdateACLs ACLNewParamsBodyCreateGroupPermissionACLPermission = "update_acls"
	ACLNewParamsBodyCreateGroupPermissionACLPermissionDeleteACLs ACLNewParamsBodyCreateGroupPermissionACLPermission = "delete_acls"
)

func (r ACLNewParamsBodyCreateGroupPermissionACLPermission) IsKnown() bool {
	switch r {
	case ACLNewParamsBodyCreateGroupPermissionACLPermissionCreate, ACLNewParamsBodyCreateGroupPermissionACLPermissionRead, ACLNewParamsBodyCreateGroupPermissionACLPermissionUpdate, ACLNewParamsBodyCreateGroupPermissionACLPermissionDelete, ACLNewParamsBodyCreateGroupPermissionACLPermissionCreateACLs, ACLNewParamsBodyCreateGroupPermissionACLPermissionReadACLs, ACLNewParamsBodyCreateGroupPermissionACLPermissionUpdateACLs, ACLNewParamsBodyCreateGroupPermissionACLPermissionDeleteACLs:
		return true
	}
	return false
}

// Optionally restricts the permission grant to just the specified object type
type ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectType struct {
	ReservedOnlyAllowNull param.Field[interface{}] `json:"__reserved_only_allow_null"`
}

func (r ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectType) ImplementsACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeUnion() {
}

// Optionally restricts the permission grant to just the specified object type
//
// Satisfied by [ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeString],
// [ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeObject],
// [ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectType].
type ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeUnion interface {
	ImplementsACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeUnion()
}

// The object type that the ACL applies to
type ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeString string

const (
	ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringOrganization  ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "organization"
	ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringProject       ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "project"
	ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringExperiment    ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "experiment"
	ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringDataset       ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "dataset"
	ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringPrompt        ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "prompt"
	ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringPromptSession ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "prompt_session"
	ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringProjectScore  ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "project_score"
	ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringProjectTag    ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "project_tag"
	ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringGroup         ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "group"
	ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringRole          ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "role"
)

func (r ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeString) IsKnown() bool {
	switch r {
	case ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringOrganization, ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringProject, ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringExperiment, ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringDataset, ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringPrompt, ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringPromptSession, ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringProjectScore, ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringProjectTag, ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringGroup, ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringRole:
		return true
	}
	return false
}

type ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeObject struct {
	// This is just a placeholder nullable object. Only pass null, not the object
	// itself
	ReservedOnlyAllowNull param.Field[ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull] `json:"__reserved_only_allow_null,required"`
}

func (r ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeObject) ImplementsACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeUnion() {
}

// This is just a placeholder nullable object. Only pass null, not the object
// itself
type ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull struct {
}

func (r ACLNewParamsBodyCreateGroupPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ACLNewParamsBodyCreateGroupRoleACL struct {
	// Id of the group the ACL applies to
	GroupID param.Field[string] `json:"group_id,required" format:"uuid"`
	// The id of the object the ACL applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ACLNewParamsBodyCreateGroupRoleACLObjectType] `json:"object_type,required"`
	// Id of the role the ACL grants
	RoleID param.Field[string] `json:"role_id,required" format:"uuid"`
	// Optionally restricts the permission grant to just the specified object type
	RestrictObjectType param.Field[ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeUnion] `json:"restrict_object_type"`
}

func (r ACLNewParamsBodyCreateGroupRoleACL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLNewParamsBodyCreateGroupRoleACL) implementsACLNewParamsBodyUnion() {}

// The object type that the ACL applies to
type ACLNewParamsBodyCreateGroupRoleACLObjectType string

const (
	ACLNewParamsBodyCreateGroupRoleACLObjectTypeOrganization  ACLNewParamsBodyCreateGroupRoleACLObjectType = "organization"
	ACLNewParamsBodyCreateGroupRoleACLObjectTypeProject       ACLNewParamsBodyCreateGroupRoleACLObjectType = "project"
	ACLNewParamsBodyCreateGroupRoleACLObjectTypeExperiment    ACLNewParamsBodyCreateGroupRoleACLObjectType = "experiment"
	ACLNewParamsBodyCreateGroupRoleACLObjectTypeDataset       ACLNewParamsBodyCreateGroupRoleACLObjectType = "dataset"
	ACLNewParamsBodyCreateGroupRoleACLObjectTypePrompt        ACLNewParamsBodyCreateGroupRoleACLObjectType = "prompt"
	ACLNewParamsBodyCreateGroupRoleACLObjectTypePromptSession ACLNewParamsBodyCreateGroupRoleACLObjectType = "prompt_session"
	ACLNewParamsBodyCreateGroupRoleACLObjectTypeProjectScore  ACLNewParamsBodyCreateGroupRoleACLObjectType = "project_score"
	ACLNewParamsBodyCreateGroupRoleACLObjectTypeProjectTag    ACLNewParamsBodyCreateGroupRoleACLObjectType = "project_tag"
	ACLNewParamsBodyCreateGroupRoleACLObjectTypeGroup         ACLNewParamsBodyCreateGroupRoleACLObjectType = "group"
	ACLNewParamsBodyCreateGroupRoleACLObjectTypeRole          ACLNewParamsBodyCreateGroupRoleACLObjectType = "role"
)

func (r ACLNewParamsBodyCreateGroupRoleACLObjectType) IsKnown() bool {
	switch r {
	case ACLNewParamsBodyCreateGroupRoleACLObjectTypeOrganization, ACLNewParamsBodyCreateGroupRoleACLObjectTypeProject, ACLNewParamsBodyCreateGroupRoleACLObjectTypeExperiment, ACLNewParamsBodyCreateGroupRoleACLObjectTypeDataset, ACLNewParamsBodyCreateGroupRoleACLObjectTypePrompt, ACLNewParamsBodyCreateGroupRoleACLObjectTypePromptSession, ACLNewParamsBodyCreateGroupRoleACLObjectTypeProjectScore, ACLNewParamsBodyCreateGroupRoleACLObjectTypeProjectTag, ACLNewParamsBodyCreateGroupRoleACLObjectTypeGroup, ACLNewParamsBodyCreateGroupRoleACLObjectTypeRole:
		return true
	}
	return false
}

// Optionally restricts the permission grant to just the specified object type
type ACLNewParamsBodyCreateGroupRoleACLRestrictObjectType struct {
	ReservedOnlyAllowNull param.Field[interface{}] `json:"__reserved_only_allow_null"`
}

func (r ACLNewParamsBodyCreateGroupRoleACLRestrictObjectType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLNewParamsBodyCreateGroupRoleACLRestrictObjectType) ImplementsACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeUnion() {
}

// Optionally restricts the permission grant to just the specified object type
//
// Satisfied by [ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeString],
// [ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeObject],
// [ACLNewParamsBodyCreateGroupRoleACLRestrictObjectType].
type ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeUnion interface {
	ImplementsACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeUnion()
}

// The object type that the ACL applies to
type ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeString string

const (
	ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringOrganization  ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "organization"
	ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringProject       ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "project"
	ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringExperiment    ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "experiment"
	ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringDataset       ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "dataset"
	ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringPrompt        ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "prompt"
	ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringPromptSession ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "prompt_session"
	ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringProjectScore  ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "project_score"
	ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringProjectTag    ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "project_tag"
	ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringGroup         ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "group"
	ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringRole          ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "role"
)

func (r ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeString) IsKnown() bool {
	switch r {
	case ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringOrganization, ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringProject, ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringExperiment, ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringDataset, ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringPrompt, ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringPromptSession, ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringProjectScore, ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringProjectTag, ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringGroup, ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeStringRole:
		return true
	}
	return false
}

type ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeObject struct {
	// This is just a placeholder nullable object. Only pass null, not the object
	// itself
	ReservedOnlyAllowNull param.Field[ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull] `json:"__reserved_only_allow_null,required"`
}

func (r ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeObject) ImplementsACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeUnion() {
}

// This is just a placeholder nullable object. Only pass null, not the object
// itself
type ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull struct {
}

func (r ACLNewParamsBodyCreateGroupRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The object type that the ACL applies to
type ACLNewParamsBodyObjectType string

const (
	ACLNewParamsBodyObjectTypeOrganization  ACLNewParamsBodyObjectType = "organization"
	ACLNewParamsBodyObjectTypeProject       ACLNewParamsBodyObjectType = "project"
	ACLNewParamsBodyObjectTypeExperiment    ACLNewParamsBodyObjectType = "experiment"
	ACLNewParamsBodyObjectTypeDataset       ACLNewParamsBodyObjectType = "dataset"
	ACLNewParamsBodyObjectTypePrompt        ACLNewParamsBodyObjectType = "prompt"
	ACLNewParamsBodyObjectTypePromptSession ACLNewParamsBodyObjectType = "prompt_session"
	ACLNewParamsBodyObjectTypeProjectScore  ACLNewParamsBodyObjectType = "project_score"
	ACLNewParamsBodyObjectTypeProjectTag    ACLNewParamsBodyObjectType = "project_tag"
	ACLNewParamsBodyObjectTypeGroup         ACLNewParamsBodyObjectType = "group"
	ACLNewParamsBodyObjectTypeRole          ACLNewParamsBodyObjectType = "role"
)

func (r ACLNewParamsBodyObjectType) IsKnown() bool {
	switch r {
	case ACLNewParamsBodyObjectTypeOrganization, ACLNewParamsBodyObjectTypeProject, ACLNewParamsBodyObjectTypeExperiment, ACLNewParamsBodyObjectTypeDataset, ACLNewParamsBodyObjectTypePrompt, ACLNewParamsBodyObjectTypePromptSession, ACLNewParamsBodyObjectTypeProjectScore, ACLNewParamsBodyObjectTypeProjectTag, ACLNewParamsBodyObjectTypeGroup, ACLNewParamsBodyObjectTypeRole:
		return true
	}
	return false
}

// Permission the ACL grants
type ACLNewParamsBodyPermission string

const (
	ACLNewParamsBodyPermissionCreate     ACLNewParamsBodyPermission = "create"
	ACLNewParamsBodyPermissionRead       ACLNewParamsBodyPermission = "read"
	ACLNewParamsBodyPermissionUpdate     ACLNewParamsBodyPermission = "update"
	ACLNewParamsBodyPermissionDelete     ACLNewParamsBodyPermission = "delete"
	ACLNewParamsBodyPermissionCreateACLs ACLNewParamsBodyPermission = "create_acls"
	ACLNewParamsBodyPermissionReadACLs   ACLNewParamsBodyPermission = "read_acls"
	ACLNewParamsBodyPermissionUpdateACLs ACLNewParamsBodyPermission = "update_acls"
	ACLNewParamsBodyPermissionDeleteACLs ACLNewParamsBodyPermission = "delete_acls"
)

func (r ACLNewParamsBodyPermission) IsKnown() bool {
	switch r {
	case ACLNewParamsBodyPermissionCreate, ACLNewParamsBodyPermissionRead, ACLNewParamsBodyPermissionUpdate, ACLNewParamsBodyPermissionDelete, ACLNewParamsBodyPermissionCreateACLs, ACLNewParamsBodyPermissionReadACLs, ACLNewParamsBodyPermissionUpdateACLs, ACLNewParamsBodyPermissionDeleteACLs:
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
	ACLListParamsObjectTypeProjectScore  ACLListParamsObjectType = "project_score"
	ACLListParamsObjectTypeProjectTag    ACLListParamsObjectType = "project_tag"
	ACLListParamsObjectTypeGroup         ACLListParamsObjectType = "group"
	ACLListParamsObjectTypeRole          ACLListParamsObjectType = "role"
)

func (r ACLListParamsObjectType) IsKnown() bool {
	switch r {
	case ACLListParamsObjectTypeOrganization, ACLListParamsObjectTypeProject, ACLListParamsObjectTypeExperiment, ACLListParamsObjectTypeDataset, ACLListParamsObjectTypePrompt, ACLListParamsObjectTypePromptSession, ACLListParamsObjectTypeProjectScore, ACLListParamsObjectTypeProjectTag, ACLListParamsObjectTypeGroup, ACLListParamsObjectTypeRole:
		return true
	}
	return false
}

type ACLReplaceParams struct {
	Body ACLReplaceParamsBodyUnion `json:"body,required"`
}

func (r ACLReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ACLReplaceParamsBody struct {
	// The object type that the ACL applies to
	ObjectType param.Field[ACLReplaceParamsBodyObjectType] `json:"object_type,required"`
	// The id of the object the ACL applies to
	ObjectID           param.Field[string]      `json:"object_id,required" format:"uuid"`
	RestrictObjectType param.Field[interface{}] `json:"restrict_object_type,required"`
	// Id of the user the ACL applies to
	UserID param.Field[string] `json:"user_id" format:"uuid"`
	// Permission the ACL grants
	Permission param.Field[ACLReplaceParamsBodyPermission] `json:"permission"`
	// Id of the role the ACL grants
	RoleID param.Field[string] `json:"role_id" format:"uuid"`
	// Id of the group the ACL applies to
	GroupID param.Field[string] `json:"group_id" format:"uuid"`
}

func (r ACLReplaceParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLReplaceParamsBody) implementsACLReplaceParamsBodyUnion() {}

// Satisfied by [ACLReplaceParamsBodyCreateUserPermissionACL],
// [ACLReplaceParamsBodyCreateUserRoleACL],
// [ACLReplaceParamsBodyCreateGroupPermissionACL],
// [ACLReplaceParamsBodyCreateGroupRoleACL], [ACLReplaceParamsBody].
type ACLReplaceParamsBodyUnion interface {
	implementsACLReplaceParamsBodyUnion()
}

type ACLReplaceParamsBodyCreateUserPermissionACL struct {
	// The id of the object the ACL applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ACLReplaceParamsBodyCreateUserPermissionACLObjectType] `json:"object_type,required"`
	// Permission the ACL grants
	Permission param.Field[ACLReplaceParamsBodyCreateUserPermissionACLPermission] `json:"permission,required"`
	// Id of the user the ACL applies to
	UserID param.Field[string] `json:"user_id,required" format:"uuid"`
	// Optionally restricts the permission grant to just the specified object type
	RestrictObjectType param.Field[ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeUnion] `json:"restrict_object_type"`
}

func (r ACLReplaceParamsBodyCreateUserPermissionACL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLReplaceParamsBodyCreateUserPermissionACL) implementsACLReplaceParamsBodyUnion() {}

// The object type that the ACL applies to
type ACLReplaceParamsBodyCreateUserPermissionACLObjectType string

const (
	ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeOrganization  ACLReplaceParamsBodyCreateUserPermissionACLObjectType = "organization"
	ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeProject       ACLReplaceParamsBodyCreateUserPermissionACLObjectType = "project"
	ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeExperiment    ACLReplaceParamsBodyCreateUserPermissionACLObjectType = "experiment"
	ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeDataset       ACLReplaceParamsBodyCreateUserPermissionACLObjectType = "dataset"
	ACLReplaceParamsBodyCreateUserPermissionACLObjectTypePrompt        ACLReplaceParamsBodyCreateUserPermissionACLObjectType = "prompt"
	ACLReplaceParamsBodyCreateUserPermissionACLObjectTypePromptSession ACLReplaceParamsBodyCreateUserPermissionACLObjectType = "prompt_session"
	ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeProjectScore  ACLReplaceParamsBodyCreateUserPermissionACLObjectType = "project_score"
	ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeProjectTag    ACLReplaceParamsBodyCreateUserPermissionACLObjectType = "project_tag"
	ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeGroup         ACLReplaceParamsBodyCreateUserPermissionACLObjectType = "group"
	ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeRole          ACLReplaceParamsBodyCreateUserPermissionACLObjectType = "role"
)

func (r ACLReplaceParamsBodyCreateUserPermissionACLObjectType) IsKnown() bool {
	switch r {
	case ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeOrganization, ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeProject, ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeExperiment, ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeDataset, ACLReplaceParamsBodyCreateUserPermissionACLObjectTypePrompt, ACLReplaceParamsBodyCreateUserPermissionACLObjectTypePromptSession, ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeProjectScore, ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeProjectTag, ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeGroup, ACLReplaceParamsBodyCreateUserPermissionACLObjectTypeRole:
		return true
	}
	return false
}

// Permission the ACL grants
type ACLReplaceParamsBodyCreateUserPermissionACLPermission string

const (
	ACLReplaceParamsBodyCreateUserPermissionACLPermissionCreate     ACLReplaceParamsBodyCreateUserPermissionACLPermission = "create"
	ACLReplaceParamsBodyCreateUserPermissionACLPermissionRead       ACLReplaceParamsBodyCreateUserPermissionACLPermission = "read"
	ACLReplaceParamsBodyCreateUserPermissionACLPermissionUpdate     ACLReplaceParamsBodyCreateUserPermissionACLPermission = "update"
	ACLReplaceParamsBodyCreateUserPermissionACLPermissionDelete     ACLReplaceParamsBodyCreateUserPermissionACLPermission = "delete"
	ACLReplaceParamsBodyCreateUserPermissionACLPermissionCreateACLs ACLReplaceParamsBodyCreateUserPermissionACLPermission = "create_acls"
	ACLReplaceParamsBodyCreateUserPermissionACLPermissionReadACLs   ACLReplaceParamsBodyCreateUserPermissionACLPermission = "read_acls"
	ACLReplaceParamsBodyCreateUserPermissionACLPermissionUpdateACLs ACLReplaceParamsBodyCreateUserPermissionACLPermission = "update_acls"
	ACLReplaceParamsBodyCreateUserPermissionACLPermissionDeleteACLs ACLReplaceParamsBodyCreateUserPermissionACLPermission = "delete_acls"
)

func (r ACLReplaceParamsBodyCreateUserPermissionACLPermission) IsKnown() bool {
	switch r {
	case ACLReplaceParamsBodyCreateUserPermissionACLPermissionCreate, ACLReplaceParamsBodyCreateUserPermissionACLPermissionRead, ACLReplaceParamsBodyCreateUserPermissionACLPermissionUpdate, ACLReplaceParamsBodyCreateUserPermissionACLPermissionDelete, ACLReplaceParamsBodyCreateUserPermissionACLPermissionCreateACLs, ACLReplaceParamsBodyCreateUserPermissionACLPermissionReadACLs, ACLReplaceParamsBodyCreateUserPermissionACLPermissionUpdateACLs, ACLReplaceParamsBodyCreateUserPermissionACLPermissionDeleteACLs:
		return true
	}
	return false
}

// Optionally restricts the permission grant to just the specified object type
type ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectType struct {
	ReservedOnlyAllowNull param.Field[interface{}] `json:"__reserved_only_allow_null"`
}

func (r ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectType) ImplementsACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeUnion() {
}

// Optionally restricts the permission grant to just the specified object type
//
// Satisfied by
// [ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeString],
// [ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeObject],
// [ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectType].
type ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeUnion interface {
	ImplementsACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeUnion()
}

// The object type that the ACL applies to
type ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeString string

const (
	ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringOrganization  ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "organization"
	ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringProject       ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "project"
	ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringExperiment    ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "experiment"
	ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringDataset       ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "dataset"
	ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringPrompt        ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "prompt"
	ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringPromptSession ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "prompt_session"
	ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringProjectScore  ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "project_score"
	ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringProjectTag    ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "project_tag"
	ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringGroup         ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "group"
	ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringRole          ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeString = "role"
)

func (r ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeString) IsKnown() bool {
	switch r {
	case ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringOrganization, ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringProject, ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringExperiment, ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringDataset, ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringPrompt, ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringPromptSession, ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringProjectScore, ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringProjectTag, ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringGroup, ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeStringRole:
		return true
	}
	return false
}

type ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeObject struct {
	// This is just a placeholder nullable object. Only pass null, not the object
	// itself
	ReservedOnlyAllowNull param.Field[ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull] `json:"__reserved_only_allow_null,required"`
}

func (r ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeObject) ImplementsACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeUnion() {
}

// This is just a placeholder nullable object. Only pass null, not the object
// itself
type ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull struct {
}

func (r ACLReplaceParamsBodyCreateUserPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ACLReplaceParamsBodyCreateUserRoleACL struct {
	// The id of the object the ACL applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ACLReplaceParamsBodyCreateUserRoleACLObjectType] `json:"object_type,required"`
	// Id of the role the ACL grants
	RoleID param.Field[string] `json:"role_id,required" format:"uuid"`
	// Id of the user the ACL applies to
	UserID param.Field[string] `json:"user_id,required" format:"uuid"`
	// Optionally restricts the permission grant to just the specified object type
	RestrictObjectType param.Field[ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeUnion] `json:"restrict_object_type"`
}

func (r ACLReplaceParamsBodyCreateUserRoleACL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLReplaceParamsBodyCreateUserRoleACL) implementsACLReplaceParamsBodyUnion() {}

// The object type that the ACL applies to
type ACLReplaceParamsBodyCreateUserRoleACLObjectType string

const (
	ACLReplaceParamsBodyCreateUserRoleACLObjectTypeOrganization  ACLReplaceParamsBodyCreateUserRoleACLObjectType = "organization"
	ACLReplaceParamsBodyCreateUserRoleACLObjectTypeProject       ACLReplaceParamsBodyCreateUserRoleACLObjectType = "project"
	ACLReplaceParamsBodyCreateUserRoleACLObjectTypeExperiment    ACLReplaceParamsBodyCreateUserRoleACLObjectType = "experiment"
	ACLReplaceParamsBodyCreateUserRoleACLObjectTypeDataset       ACLReplaceParamsBodyCreateUserRoleACLObjectType = "dataset"
	ACLReplaceParamsBodyCreateUserRoleACLObjectTypePrompt        ACLReplaceParamsBodyCreateUserRoleACLObjectType = "prompt"
	ACLReplaceParamsBodyCreateUserRoleACLObjectTypePromptSession ACLReplaceParamsBodyCreateUserRoleACLObjectType = "prompt_session"
	ACLReplaceParamsBodyCreateUserRoleACLObjectTypeProjectScore  ACLReplaceParamsBodyCreateUserRoleACLObjectType = "project_score"
	ACLReplaceParamsBodyCreateUserRoleACLObjectTypeProjectTag    ACLReplaceParamsBodyCreateUserRoleACLObjectType = "project_tag"
	ACLReplaceParamsBodyCreateUserRoleACLObjectTypeGroup         ACLReplaceParamsBodyCreateUserRoleACLObjectType = "group"
	ACLReplaceParamsBodyCreateUserRoleACLObjectTypeRole          ACLReplaceParamsBodyCreateUserRoleACLObjectType = "role"
)

func (r ACLReplaceParamsBodyCreateUserRoleACLObjectType) IsKnown() bool {
	switch r {
	case ACLReplaceParamsBodyCreateUserRoleACLObjectTypeOrganization, ACLReplaceParamsBodyCreateUserRoleACLObjectTypeProject, ACLReplaceParamsBodyCreateUserRoleACLObjectTypeExperiment, ACLReplaceParamsBodyCreateUserRoleACLObjectTypeDataset, ACLReplaceParamsBodyCreateUserRoleACLObjectTypePrompt, ACLReplaceParamsBodyCreateUserRoleACLObjectTypePromptSession, ACLReplaceParamsBodyCreateUserRoleACLObjectTypeProjectScore, ACLReplaceParamsBodyCreateUserRoleACLObjectTypeProjectTag, ACLReplaceParamsBodyCreateUserRoleACLObjectTypeGroup, ACLReplaceParamsBodyCreateUserRoleACLObjectTypeRole:
		return true
	}
	return false
}

// Optionally restricts the permission grant to just the specified object type
type ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectType struct {
	ReservedOnlyAllowNull param.Field[interface{}] `json:"__reserved_only_allow_null"`
}

func (r ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectType) ImplementsACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeUnion() {
}

// Optionally restricts the permission grant to just the specified object type
//
// Satisfied by [ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeString],
// [ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeObject],
// [ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectType].
type ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeUnion interface {
	ImplementsACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeUnion()
}

// The object type that the ACL applies to
type ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeString string

const (
	ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringOrganization  ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeString = "organization"
	ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringProject       ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeString = "project"
	ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringExperiment    ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeString = "experiment"
	ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringDataset       ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeString = "dataset"
	ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringPrompt        ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeString = "prompt"
	ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringPromptSession ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeString = "prompt_session"
	ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringProjectScore  ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeString = "project_score"
	ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringProjectTag    ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeString = "project_tag"
	ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringGroup         ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeString = "group"
	ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringRole          ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeString = "role"
)

func (r ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeString) IsKnown() bool {
	switch r {
	case ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringOrganization, ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringProject, ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringExperiment, ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringDataset, ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringPrompt, ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringPromptSession, ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringProjectScore, ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringProjectTag, ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringGroup, ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeStringRole:
		return true
	}
	return false
}

type ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeObject struct {
	// This is just a placeholder nullable object. Only pass null, not the object
	// itself
	ReservedOnlyAllowNull param.Field[ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull] `json:"__reserved_only_allow_null,required"`
}

func (r ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeObject) ImplementsACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeUnion() {
}

// This is just a placeholder nullable object. Only pass null, not the object
// itself
type ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull struct {
}

func (r ACLReplaceParamsBodyCreateUserRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ACLReplaceParamsBodyCreateGroupPermissionACL struct {
	// Id of the group the ACL applies to
	GroupID param.Field[string] `json:"group_id,required" format:"uuid"`
	// The id of the object the ACL applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ACLReplaceParamsBodyCreateGroupPermissionACLObjectType] `json:"object_type,required"`
	// Permission the ACL grants
	Permission param.Field[ACLReplaceParamsBodyCreateGroupPermissionACLPermission] `json:"permission,required"`
	// Optionally restricts the permission grant to just the specified object type
	RestrictObjectType param.Field[ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeUnion] `json:"restrict_object_type"`
}

func (r ACLReplaceParamsBodyCreateGroupPermissionACL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLReplaceParamsBodyCreateGroupPermissionACL) implementsACLReplaceParamsBodyUnion() {}

// The object type that the ACL applies to
type ACLReplaceParamsBodyCreateGroupPermissionACLObjectType string

const (
	ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeOrganization  ACLReplaceParamsBodyCreateGroupPermissionACLObjectType = "organization"
	ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeProject       ACLReplaceParamsBodyCreateGroupPermissionACLObjectType = "project"
	ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeExperiment    ACLReplaceParamsBodyCreateGroupPermissionACLObjectType = "experiment"
	ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeDataset       ACLReplaceParamsBodyCreateGroupPermissionACLObjectType = "dataset"
	ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypePrompt        ACLReplaceParamsBodyCreateGroupPermissionACLObjectType = "prompt"
	ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypePromptSession ACLReplaceParamsBodyCreateGroupPermissionACLObjectType = "prompt_session"
	ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeProjectScore  ACLReplaceParamsBodyCreateGroupPermissionACLObjectType = "project_score"
	ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeProjectTag    ACLReplaceParamsBodyCreateGroupPermissionACLObjectType = "project_tag"
	ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeGroup         ACLReplaceParamsBodyCreateGroupPermissionACLObjectType = "group"
	ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeRole          ACLReplaceParamsBodyCreateGroupPermissionACLObjectType = "role"
)

func (r ACLReplaceParamsBodyCreateGroupPermissionACLObjectType) IsKnown() bool {
	switch r {
	case ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeOrganization, ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeProject, ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeExperiment, ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeDataset, ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypePrompt, ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypePromptSession, ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeProjectScore, ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeProjectTag, ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeGroup, ACLReplaceParamsBodyCreateGroupPermissionACLObjectTypeRole:
		return true
	}
	return false
}

// Permission the ACL grants
type ACLReplaceParamsBodyCreateGroupPermissionACLPermission string

const (
	ACLReplaceParamsBodyCreateGroupPermissionACLPermissionCreate     ACLReplaceParamsBodyCreateGroupPermissionACLPermission = "create"
	ACLReplaceParamsBodyCreateGroupPermissionACLPermissionRead       ACLReplaceParamsBodyCreateGroupPermissionACLPermission = "read"
	ACLReplaceParamsBodyCreateGroupPermissionACLPermissionUpdate     ACLReplaceParamsBodyCreateGroupPermissionACLPermission = "update"
	ACLReplaceParamsBodyCreateGroupPermissionACLPermissionDelete     ACLReplaceParamsBodyCreateGroupPermissionACLPermission = "delete"
	ACLReplaceParamsBodyCreateGroupPermissionACLPermissionCreateACLs ACLReplaceParamsBodyCreateGroupPermissionACLPermission = "create_acls"
	ACLReplaceParamsBodyCreateGroupPermissionACLPermissionReadACLs   ACLReplaceParamsBodyCreateGroupPermissionACLPermission = "read_acls"
	ACLReplaceParamsBodyCreateGroupPermissionACLPermissionUpdateACLs ACLReplaceParamsBodyCreateGroupPermissionACLPermission = "update_acls"
	ACLReplaceParamsBodyCreateGroupPermissionACLPermissionDeleteACLs ACLReplaceParamsBodyCreateGroupPermissionACLPermission = "delete_acls"
)

func (r ACLReplaceParamsBodyCreateGroupPermissionACLPermission) IsKnown() bool {
	switch r {
	case ACLReplaceParamsBodyCreateGroupPermissionACLPermissionCreate, ACLReplaceParamsBodyCreateGroupPermissionACLPermissionRead, ACLReplaceParamsBodyCreateGroupPermissionACLPermissionUpdate, ACLReplaceParamsBodyCreateGroupPermissionACLPermissionDelete, ACLReplaceParamsBodyCreateGroupPermissionACLPermissionCreateACLs, ACLReplaceParamsBodyCreateGroupPermissionACLPermissionReadACLs, ACLReplaceParamsBodyCreateGroupPermissionACLPermissionUpdateACLs, ACLReplaceParamsBodyCreateGroupPermissionACLPermissionDeleteACLs:
		return true
	}
	return false
}

// Optionally restricts the permission grant to just the specified object type
type ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectType struct {
	ReservedOnlyAllowNull param.Field[interface{}] `json:"__reserved_only_allow_null"`
}

func (r ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectType) ImplementsACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeUnion() {
}

// Optionally restricts the permission grant to just the specified object type
//
// Satisfied by
// [ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeString],
// [ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeObject],
// [ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectType].
type ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeUnion interface {
	ImplementsACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeUnion()
}

// The object type that the ACL applies to
type ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeString string

const (
	ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringOrganization  ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "organization"
	ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringProject       ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "project"
	ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringExperiment    ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "experiment"
	ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringDataset       ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "dataset"
	ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringPrompt        ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "prompt"
	ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringPromptSession ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "prompt_session"
	ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringProjectScore  ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "project_score"
	ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringProjectTag    ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "project_tag"
	ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringGroup         ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "group"
	ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringRole          ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeString = "role"
)

func (r ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeString) IsKnown() bool {
	switch r {
	case ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringOrganization, ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringProject, ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringExperiment, ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringDataset, ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringPrompt, ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringPromptSession, ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringProjectScore, ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringProjectTag, ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringGroup, ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeStringRole:
		return true
	}
	return false
}

type ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeObject struct {
	// This is just a placeholder nullable object. Only pass null, not the object
	// itself
	ReservedOnlyAllowNull param.Field[ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull] `json:"__reserved_only_allow_null,required"`
}

func (r ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeObject) ImplementsACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeUnion() {
}

// This is just a placeholder nullable object. Only pass null, not the object
// itself
type ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull struct {
}

func (r ACLReplaceParamsBodyCreateGroupPermissionACLRestrictObjectTypeObjectReservedOnlyAllowNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ACLReplaceParamsBodyCreateGroupRoleACL struct {
	// Id of the group the ACL applies to
	GroupID param.Field[string] `json:"group_id,required" format:"uuid"`
	// The id of the object the ACL applies to
	ObjectID param.Field[string] `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	ObjectType param.Field[ACLReplaceParamsBodyCreateGroupRoleACLObjectType] `json:"object_type,required"`
	// Id of the role the ACL grants
	RoleID param.Field[string] `json:"role_id,required" format:"uuid"`
	// Optionally restricts the permission grant to just the specified object type
	RestrictObjectType param.Field[ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeUnion] `json:"restrict_object_type"`
}

func (r ACLReplaceParamsBodyCreateGroupRoleACL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLReplaceParamsBodyCreateGroupRoleACL) implementsACLReplaceParamsBodyUnion() {}

// The object type that the ACL applies to
type ACLReplaceParamsBodyCreateGroupRoleACLObjectType string

const (
	ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeOrganization  ACLReplaceParamsBodyCreateGroupRoleACLObjectType = "organization"
	ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeProject       ACLReplaceParamsBodyCreateGroupRoleACLObjectType = "project"
	ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeExperiment    ACLReplaceParamsBodyCreateGroupRoleACLObjectType = "experiment"
	ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeDataset       ACLReplaceParamsBodyCreateGroupRoleACLObjectType = "dataset"
	ACLReplaceParamsBodyCreateGroupRoleACLObjectTypePrompt        ACLReplaceParamsBodyCreateGroupRoleACLObjectType = "prompt"
	ACLReplaceParamsBodyCreateGroupRoleACLObjectTypePromptSession ACLReplaceParamsBodyCreateGroupRoleACLObjectType = "prompt_session"
	ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeProjectScore  ACLReplaceParamsBodyCreateGroupRoleACLObjectType = "project_score"
	ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeProjectTag    ACLReplaceParamsBodyCreateGroupRoleACLObjectType = "project_tag"
	ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeGroup         ACLReplaceParamsBodyCreateGroupRoleACLObjectType = "group"
	ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeRole          ACLReplaceParamsBodyCreateGroupRoleACLObjectType = "role"
)

func (r ACLReplaceParamsBodyCreateGroupRoleACLObjectType) IsKnown() bool {
	switch r {
	case ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeOrganization, ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeProject, ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeExperiment, ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeDataset, ACLReplaceParamsBodyCreateGroupRoleACLObjectTypePrompt, ACLReplaceParamsBodyCreateGroupRoleACLObjectTypePromptSession, ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeProjectScore, ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeProjectTag, ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeGroup, ACLReplaceParamsBodyCreateGroupRoleACLObjectTypeRole:
		return true
	}
	return false
}

// Optionally restricts the permission grant to just the specified object type
type ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectType struct {
	ReservedOnlyAllowNull param.Field[interface{}] `json:"__reserved_only_allow_null"`
}

func (r ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectType) ImplementsACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeUnion() {
}

// Optionally restricts the permission grant to just the specified object type
//
// Satisfied by [ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeString],
// [ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeObject],
// [ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectType].
type ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeUnion interface {
	ImplementsACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeUnion()
}

// The object type that the ACL applies to
type ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeString string

const (
	ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringOrganization  ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "organization"
	ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringProject       ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "project"
	ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringExperiment    ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "experiment"
	ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringDataset       ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "dataset"
	ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringPrompt        ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "prompt"
	ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringPromptSession ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "prompt_session"
	ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringProjectScore  ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "project_score"
	ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringProjectTag    ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "project_tag"
	ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringGroup         ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "group"
	ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringRole          ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeString = "role"
)

func (r ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeString) IsKnown() bool {
	switch r {
	case ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringOrganization, ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringProject, ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringExperiment, ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringDataset, ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringPrompt, ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringPromptSession, ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringProjectScore, ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringProjectTag, ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringGroup, ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeStringRole:
		return true
	}
	return false
}

type ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeObject struct {
	// This is just a placeholder nullable object. Only pass null, not the object
	// itself
	ReservedOnlyAllowNull param.Field[ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull] `json:"__reserved_only_allow_null,required"`
}

func (r ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeObject) ImplementsACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeUnion() {
}

// This is just a placeholder nullable object. Only pass null, not the object
// itself
type ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull struct {
}

func (r ACLReplaceParamsBodyCreateGroupRoleACLRestrictObjectTypeObjectReservedOnlyAllowNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The object type that the ACL applies to
type ACLReplaceParamsBodyObjectType string

const (
	ACLReplaceParamsBodyObjectTypeOrganization  ACLReplaceParamsBodyObjectType = "organization"
	ACLReplaceParamsBodyObjectTypeProject       ACLReplaceParamsBodyObjectType = "project"
	ACLReplaceParamsBodyObjectTypeExperiment    ACLReplaceParamsBodyObjectType = "experiment"
	ACLReplaceParamsBodyObjectTypeDataset       ACLReplaceParamsBodyObjectType = "dataset"
	ACLReplaceParamsBodyObjectTypePrompt        ACLReplaceParamsBodyObjectType = "prompt"
	ACLReplaceParamsBodyObjectTypePromptSession ACLReplaceParamsBodyObjectType = "prompt_session"
	ACLReplaceParamsBodyObjectTypeProjectScore  ACLReplaceParamsBodyObjectType = "project_score"
	ACLReplaceParamsBodyObjectTypeProjectTag    ACLReplaceParamsBodyObjectType = "project_tag"
	ACLReplaceParamsBodyObjectTypeGroup         ACLReplaceParamsBodyObjectType = "group"
	ACLReplaceParamsBodyObjectTypeRole          ACLReplaceParamsBodyObjectType = "role"
)

func (r ACLReplaceParamsBodyObjectType) IsKnown() bool {
	switch r {
	case ACLReplaceParamsBodyObjectTypeOrganization, ACLReplaceParamsBodyObjectTypeProject, ACLReplaceParamsBodyObjectTypeExperiment, ACLReplaceParamsBodyObjectTypeDataset, ACLReplaceParamsBodyObjectTypePrompt, ACLReplaceParamsBodyObjectTypePromptSession, ACLReplaceParamsBodyObjectTypeProjectScore, ACLReplaceParamsBodyObjectTypeProjectTag, ACLReplaceParamsBodyObjectTypeGroup, ACLReplaceParamsBodyObjectTypeRole:
		return true
	}
	return false
}

// Permission the ACL grants
type ACLReplaceParamsBodyPermission string

const (
	ACLReplaceParamsBodyPermissionCreate     ACLReplaceParamsBodyPermission = "create"
	ACLReplaceParamsBodyPermissionRead       ACLReplaceParamsBodyPermission = "read"
	ACLReplaceParamsBodyPermissionUpdate     ACLReplaceParamsBodyPermission = "update"
	ACLReplaceParamsBodyPermissionDelete     ACLReplaceParamsBodyPermission = "delete"
	ACLReplaceParamsBodyPermissionCreateACLs ACLReplaceParamsBodyPermission = "create_acls"
	ACLReplaceParamsBodyPermissionReadACLs   ACLReplaceParamsBodyPermission = "read_acls"
	ACLReplaceParamsBodyPermissionUpdateACLs ACLReplaceParamsBodyPermission = "update_acls"
	ACLReplaceParamsBodyPermissionDeleteACLs ACLReplaceParamsBodyPermission = "delete_acls"
)

func (r ACLReplaceParamsBodyPermission) IsKnown() bool {
	switch r {
	case ACLReplaceParamsBodyPermissionCreate, ACLReplaceParamsBodyPermissionRead, ACLReplaceParamsBodyPermissionUpdate, ACLReplaceParamsBodyPermissionDelete, ACLReplaceParamsBodyPermissionCreateACLs, ACLReplaceParamsBodyPermissionReadACLs, ACLReplaceParamsBodyPermissionUpdateACLs, ACLReplaceParamsBodyPermissionDeleteACLs:
		return true
	}
	return false
}
