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
	"github.com/braintrustdata/braintrust-go/internal/pagination"
	"github.com/braintrustdata/braintrust-go/internal/param"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
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
