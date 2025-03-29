// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"
	"time"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/packages/param"
	"github.com/braintrustdata/braintrust-go/packages/resp"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type AISecret struct {
	// Unique identifier for the AI secret
	ID string `json:"id,required" format:"uuid"`
	// Name of the AI secret
	Name string `json:"name,required"`
	// Unique identifier for the organization
	OrgID string `json:"org_id,required" format:"uuid"`
	// Date of AI secret creation
	Created       time.Time              `json:"created,nullable" format:"date-time"`
	Metadata      map[string]interface{} `json:"metadata,nullable"`
	PreviewSecret string                 `json:"preview_secret,nullable"`
	Type          string                 `json:"type,nullable"`
	// Date of last AI secret update
	UpdatedAt time.Time `json:"updated_at,nullable" format:"date-time"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID            resp.Field
		Name          resp.Field
		OrgID         resp.Field
		Created       resp.Field
		Metadata      resp.Field
		PreviewSecret resp.Field
		Type          resp.Field
		UpdatedAt     resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AISecret) RawJSON() string { return r.JSON.raw }
func (r *AISecret) UnmarshalJSON(data []byte) error {
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
type ACL struct {
	// Unique identifier for the acl
	ID string `json:"id,required" format:"uuid"`
	// The organization the ACL's referred object belongs to
	ObjectOrgID string `json:"_object_org_id,required" format:"uuid"`
	// The id of the object the ACL applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType ACLObjectType `json:"object_type,required"`
	// Date of acl creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Id of the group the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	GroupID string `json:"group_id,nullable" format:"uuid"`
	// Permission the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	//
	// Any of "create", "read", "update", "delete", "create_acls", "read_acls",
	// "update_acls", "delete_acls".
	Permission Permission `json:"permission,nullable"`
	// When setting a permission directly, optionally restricts the permission grant to
	// just the specified object type. Cannot be set alongside a `role_id`.
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	RestrictObjectType ACLObjectType `json:"restrict_object_type,nullable"`
	// Id of the role the ACL grants. Exactly one of `permission` and `role_id` will be
	// provided
	RoleID string `json:"role_id,nullable" format:"uuid"`
	// Id of the user the ACL applies to. Exactly one of `user_id` and `group_id` will
	// be provided
	UserID string `json:"user_id,nullable" format:"uuid"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                 resp.Field
		ObjectOrgID        resp.Field
		ObjectID           resp.Field
		ObjectType         resp.Field
		Created            resp.Field
		GroupID            resp.Field
		Permission         resp.Field
		RestrictObjectType resp.Field
		RoleID             resp.Field
		UserID             resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ACL) RawJSON() string { return r.JSON.raw }
func (r *ACL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ACLBatchUpdateResponse struct {
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
	AddedACLs []ACL `json:"added_acls,required"`
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
	RemovedACLs []ACL `json:"removed_acls,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AddedACLs   resp.Field
		RemovedACLs resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ACLBatchUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ACLBatchUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type APIKey struct {
	// Unique identifier for the api key
	ID string `json:"id,required" format:"uuid"`
	// Name of the api key
	Name        string `json:"name,required"`
	PreviewName string `json:"preview_name,required"`
	// Date of api key creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Unique identifier for the organization
	OrgID string `json:"org_id,nullable" format:"uuid"`
	// Unique identifier for the user
	UserID string `json:"user_id,nullable" format:"uuid"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Name        resp.Field
		PreviewName resp.Field
		Created     resp.Field
		OrgID       resp.Field
		UserID      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKey) RawJSON() string { return r.JSON.raw }
func (r *APIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionContentPartImage struct {
	ImageURL ChatCompletionContentPartImageImageURL `json:"image_url,required"`
	// Any of "image_url".
	Type ChatCompletionContentPartImageType `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ImageURL    resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionContentPartImage) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionContentPartImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ChatCompletionContentPartImage to a
// ChatCompletionContentPartImageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChatCompletionContentPartImageParam.IsOverridden()
func (r ChatCompletionContentPartImage) ToParam() ChatCompletionContentPartImageParam {
	return param.OverrideObj[ChatCompletionContentPartImageParam](r.RawJSON())
}

type ChatCompletionContentPartImageImageURL struct {
	URL string `json:"url,required"`
	// Any of "auto", "low", "high".
	Detail ChatCompletionContentPartImageImageURLDetail `json:"detail"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		URL         resp.Field
		Detail      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionContentPartImageImageURL) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionContentPartImageImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionContentPartImageImageURLDetail string

const (
	ChatCompletionContentPartImageImageURLDetailAuto ChatCompletionContentPartImageImageURLDetail = "auto"
	ChatCompletionContentPartImageImageURLDetailLow  ChatCompletionContentPartImageImageURLDetail = "low"
	ChatCompletionContentPartImageImageURLDetailHigh ChatCompletionContentPartImageImageURLDetail = "high"
)

type ChatCompletionContentPartImageType string

const (
	ChatCompletionContentPartImageTypeImageURL ChatCompletionContentPartImageType = "image_url"
)

// The properties ImageURL, Type are required.
type ChatCompletionContentPartImageParam struct {
	ImageURL ChatCompletionContentPartImageImageURLParam `json:"image_url,omitzero,required"`
	// Any of "image_url".
	Type ChatCompletionContentPartImageType `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ChatCompletionContentPartImageParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ChatCompletionContentPartImageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionContentPartImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property URL is required.
type ChatCompletionContentPartImageImageURLParam struct {
	URL string `json:"url,required"`
	// Any of "auto", "low", "high".
	Detail ChatCompletionContentPartImageImageURLDetail `json:"detail,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ChatCompletionContentPartImageImageURLParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ChatCompletionContentPartImageImageURLParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionContentPartImageImageURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ChatCompletionContentPartText struct {
	// Any of "text".
	Type ChatCompletionContentPartTextType `json:"type,required"`
	Text string                            `json:"text"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Type        resp.Field
		Text        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionContentPartText) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionContentPartText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ChatCompletionContentPartText to a
// ChatCompletionContentPartTextParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChatCompletionContentPartTextParam.IsOverridden()
func (r ChatCompletionContentPartText) ToParam() ChatCompletionContentPartTextParam {
	return param.OverrideObj[ChatCompletionContentPartTextParam](r.RawJSON())
}

type ChatCompletionContentPartTextType string

const (
	ChatCompletionContentPartTextTypeText ChatCompletionContentPartTextType = "text"
)

// The property Type is required.
type ChatCompletionContentPartTextParam struct {
	// Any of "text".
	Type ChatCompletionContentPartTextType `json:"type,omitzero,required"`
	Text param.Opt[string]                 `json:"text,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ChatCompletionContentPartTextParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ChatCompletionContentPartTextParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionContentPartTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ChatCompletionMessageToolCall struct {
	ID       string                                `json:"id,required"`
	Function ChatCompletionMessageToolCallFunction `json:"function,required"`
	// Any of "function".
	Type ChatCompletionMessageToolCallType `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Function    resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionMessageToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionMessageToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ChatCompletionMessageToolCall to a
// ChatCompletionMessageToolCallParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChatCompletionMessageToolCallParam.IsOverridden()
func (r ChatCompletionMessageToolCall) ToParam() ChatCompletionMessageToolCallParam {
	return param.OverrideObj[ChatCompletionMessageToolCallParam](r.RawJSON())
}

type ChatCompletionMessageToolCallFunction struct {
	Arguments string `json:"arguments,required"`
	Name      string `json:"name,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Arguments   resp.Field
		Name        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionMessageToolCallFunction) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionMessageToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionMessageToolCallType string

const (
	ChatCompletionMessageToolCallTypeFunction ChatCompletionMessageToolCallType = "function"
)

// The properties ID, Function, Type are required.
type ChatCompletionMessageToolCallParam struct {
	ID       string                                     `json:"id,required"`
	Function ChatCompletionMessageToolCallFunctionParam `json:"function,omitzero,required"`
	// Any of "function".
	Type ChatCompletionMessageToolCallType `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ChatCompletionMessageToolCallParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ChatCompletionMessageToolCallParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionMessageToolCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties Arguments, Name are required.
type ChatCompletionMessageToolCallFunctionParam struct {
	Arguments string `json:"arguments,required"`
	Name      string `json:"name,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ChatCompletionMessageToolCallFunctionParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ChatCompletionMessageToolCallFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionMessageToolCallFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type CodeBundle struct {
	BundleID       string                   `json:"bundle_id,required"`
	Location       CodeBundleLocationUnion  `json:"location,required"`
	RuntimeContext CodeBundleRuntimeContext `json:"runtime_context,required"`
	// A preview of the code
	Preview string `json:"preview,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		BundleID       resp.Field
		Location       resp.Field
		RuntimeContext resp.Field
		Preview        resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CodeBundle) RawJSON() string { return r.JSON.raw }
func (r *CodeBundle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CodeBundle to a CodeBundleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CodeBundleParam.IsOverridden()
func (r CodeBundle) ToParam() CodeBundleParam {
	return param.OverrideObj[CodeBundleParam](r.RawJSON())
}

// CodeBundleLocationUnion contains all possible properties and values from
// [CodeBundleLocationExperiment], [CodeBundleLocationFunction].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CodeBundleLocationUnion struct {
	// This field is from variant [CodeBundleLocationExperiment].
	EvalName string `json:"eval_name"`
	// This field is from variant [CodeBundleLocationExperiment].
	Position CodeBundleLocationExperimentPositionUnion `json:"position"`
	Type     string                                    `json:"type"`
	// This field is from variant [CodeBundleLocationFunction].
	Index int64 `json:"index"`
	JSON  struct {
		EvalName resp.Field
		Position resp.Field
		Type     resp.Field
		Index    resp.Field
		raw      string
	} `json:"-"`
}

func (u CodeBundleLocationUnion) AsExperiment() (v CodeBundleLocationExperiment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CodeBundleLocationUnion) AsFunction() (v CodeBundleLocationFunction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CodeBundleLocationUnion) RawJSON() string { return u.JSON.raw }

func (r *CodeBundleLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CodeBundleLocationExperiment struct {
	EvalName string                                    `json:"eval_name,required"`
	Position CodeBundleLocationExperimentPositionUnion `json:"position,required"`
	// Any of "experiment".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		EvalName    resp.Field
		Position    resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CodeBundleLocationExperiment) RawJSON() string { return r.JSON.raw }
func (r *CodeBundleLocationExperiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CodeBundleLocationExperimentPositionUnion contains all possible properties and
// values from [CodeBundleLocationExperimentPositionType],
// [CodeBundleLocationExperimentPositionScorer].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CodeBundleLocationExperimentPositionUnion struct {
	Type string `json:"type"`
	// This field is from variant [CodeBundleLocationExperimentPositionScorer].
	Index int64 `json:"index"`
	JSON  struct {
		Type  resp.Field
		Index resp.Field
		raw   string
	} `json:"-"`
}

func (u CodeBundleLocationExperimentPositionUnion) AsCodeBundleLocationExperimentPositionType() (v CodeBundleLocationExperimentPositionType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CodeBundleLocationExperimentPositionUnion) AsScorer() (v CodeBundleLocationExperimentPositionScorer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CodeBundleLocationExperimentPositionUnion) RawJSON() string { return u.JSON.raw }

func (r *CodeBundleLocationExperimentPositionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CodeBundleLocationExperimentPositionType struct {
	// Any of "task".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CodeBundleLocationExperimentPositionType) RawJSON() string { return r.JSON.raw }
func (r *CodeBundleLocationExperimentPositionType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CodeBundleLocationExperimentPositionScorer struct {
	Index int64 `json:"index,required"`
	// Any of "scorer".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Index       resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CodeBundleLocationExperimentPositionScorer) RawJSON() string { return r.JSON.raw }
func (r *CodeBundleLocationExperimentPositionScorer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CodeBundleLocationFunction struct {
	Index int64 `json:"index,required"`
	// Any of "function".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Index       resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CodeBundleLocationFunction) RawJSON() string { return r.JSON.raw }
func (r *CodeBundleLocationFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CodeBundleRuntimeContext struct {
	// Any of "node", "python".
	Runtime string `json:"runtime,required"`
	Version string `json:"version,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Runtime     resp.Field
		Version     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CodeBundleRuntimeContext) RawJSON() string { return r.JSON.raw }
func (r *CodeBundleRuntimeContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BundleID, Location, RuntimeContext are required.
type CodeBundleParam struct {
	BundleID       string                        `json:"bundle_id,required"`
	Location       CodeBundleLocationUnionParam  `json:"location,omitzero,required"`
	RuntimeContext CodeBundleRuntimeContextParam `json:"runtime_context,omitzero,required"`
	// A preview of the code
	Preview param.Opt[string] `json:"preview,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f CodeBundleParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r CodeBundleParam) MarshalJSON() (data []byte, err error) {
	type shadow CodeBundleParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CodeBundleLocationUnionParam struct {
	OfExperiment *CodeBundleLocationExperimentParam `json:",omitzero,inline"`
	OfFunction   *CodeBundleLocationFunctionParam   `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u CodeBundleLocationUnionParam) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u CodeBundleLocationUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[CodeBundleLocationUnionParam](u.OfExperiment, u.OfFunction)
}

func (u *CodeBundleLocationUnionParam) asAny() any {
	if !param.IsOmitted(u.OfExperiment) {
		return u.OfExperiment
	} else if !param.IsOmitted(u.OfFunction) {
		return u.OfFunction
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CodeBundleLocationUnionParam) GetEvalName() *string {
	if vt := u.OfExperiment; vt != nil {
		return &vt.EvalName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CodeBundleLocationUnionParam) GetPosition() *CodeBundleLocationExperimentPositionUnionParam {
	if vt := u.OfExperiment; vt != nil {
		return &vt.Position
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CodeBundleLocationUnionParam) GetIndex() *int64 {
	if vt := u.OfFunction; vt != nil {
		return &vt.Index
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CodeBundleLocationUnionParam) GetType() *string {
	if vt := u.OfExperiment; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFunction; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The properties EvalName, Position, Type are required.
type CodeBundleLocationExperimentParam struct {
	EvalName string                                         `json:"eval_name,required"`
	Position CodeBundleLocationExperimentPositionUnionParam `json:"position,omitzero,required"`
	// Any of "experiment".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f CodeBundleLocationExperimentParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r CodeBundleLocationExperimentParam) MarshalJSON() (data []byte, err error) {
	type shadow CodeBundleLocationExperimentParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[CodeBundleLocationExperimentParam](
		"Type", false, "experiment",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CodeBundleLocationExperimentPositionUnionParam struct {
	OfCodeBundleLocationExperimentPositionType *CodeBundleLocationExperimentPositionTypeParam   `json:",omitzero,inline"`
	OfScorer                                   *CodeBundleLocationExperimentPositionScorerParam `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u CodeBundleLocationExperimentPositionUnionParam) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u CodeBundleLocationExperimentPositionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[CodeBundleLocationExperimentPositionUnionParam](u.OfCodeBundleLocationExperimentPositionType, u.OfScorer)
}

func (u *CodeBundleLocationExperimentPositionUnionParam) asAny() any {
	if !param.IsOmitted(u.OfCodeBundleLocationExperimentPositionType) {
		return u.OfCodeBundleLocationExperimentPositionType
	} else if !param.IsOmitted(u.OfScorer) {
		return u.OfScorer
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CodeBundleLocationExperimentPositionUnionParam) GetIndex() *int64 {
	if vt := u.OfScorer; vt != nil {
		return &vt.Index
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CodeBundleLocationExperimentPositionUnionParam) GetType() *string {
	if vt := u.OfCodeBundleLocationExperimentPositionType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfScorer; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The property Type is required.
type CodeBundleLocationExperimentPositionTypeParam struct {
	// Any of "task".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f CodeBundleLocationExperimentPositionTypeParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r CodeBundleLocationExperimentPositionTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow CodeBundleLocationExperimentPositionTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[CodeBundleLocationExperimentPositionTypeParam](
		"Type", false, "task",
	)
}

// The properties Index, Type are required.
type CodeBundleLocationExperimentPositionScorerParam struct {
	Index int64 `json:"index,required"`
	// Any of "scorer".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f CodeBundleLocationExperimentPositionScorerParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r CodeBundleLocationExperimentPositionScorerParam) MarshalJSON() (data []byte, err error) {
	type shadow CodeBundleLocationExperimentPositionScorerParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[CodeBundleLocationExperimentPositionScorerParam](
		"Type", false, "scorer",
	)
}

// The properties Index, Type are required.
type CodeBundleLocationFunctionParam struct {
	Index int64 `json:"index,required"`
	// Any of "function".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f CodeBundleLocationFunctionParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r CodeBundleLocationFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow CodeBundleLocationFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[CodeBundleLocationFunctionParam](
		"Type", false, "function",
	)
}

// The properties Runtime, Version are required.
type CodeBundleRuntimeContextParam struct {
	// Any of "node", "python".
	Runtime string `json:"runtime,omitzero,required"`
	Version string `json:"version,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f CodeBundleRuntimeContextParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r CodeBundleRuntimeContextParam) MarshalJSON() (data []byte, err error) {
	type shadow CodeBundleRuntimeContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[CodeBundleRuntimeContextParam](
		"Runtime", false, "node", "python",
	)
}

type CreateAPIKeyOutput struct {
	// Unique identifier for the api key
	ID string `json:"id,required" format:"uuid"`
	// The raw API key. It will only be exposed this one time
	Key string `json:"key,required"`
	// Name of the api key
	Name        string `json:"name,required"`
	PreviewName string `json:"preview_name,required"`
	// Date of api key creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Unique identifier for the organization
	OrgID string `json:"org_id,nullable" format:"uuid"`
	// Unique identifier for the user
	UserID string `json:"user_id,nullable" format:"uuid"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Key         resp.Field
		Name        resp.Field
		PreviewName resp.Field
		Created     resp.Field
		OrgID       resp.Field
		UserID      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateAPIKeyOutput) RawJSON() string { return r.JSON.raw }
func (r *CreateAPIKeyOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary of a dataset's data
type DataSummary struct {
	// Total number of records in the dataset
	TotalRecords int64 `json:"total_records,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		TotalRecords resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataSummary) RawJSON() string { return r.JSON.raw }
func (r *DataSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Dataset struct {
	// Unique identifier for the dataset
	ID string `json:"id,required" format:"uuid"`
	// Name of the dataset. Within a project, dataset names are unique
	Name string `json:"name,required"`
	// Unique identifier for the project that the dataset belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// Date of dataset creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Date of dataset deletion, or null if the dataset is still active
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Textual description of the dataset
	Description string `json:"description,nullable"`
	// User-controlled metadata about the dataset
	Metadata map[string]interface{} `json:"metadata,nullable"`
	// Identifies the user who created the dataset
	UserID string `json:"user_id,nullable" format:"uuid"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Name        resp.Field
		ProjectID   resp.Field
		Created     resp.Field
		DeletedAt   resp.Field
		Description resp.Field
		Metadata    resp.Field
		UserID      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Dataset) RawJSON() string { return r.JSON.raw }
func (r *Dataset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatasetEvent struct {
	// A unique identifier for the dataset event. If you don't provide one, BrainTrust
	// will generate one for you
	ID string `json:"id,required"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the dataset (see the `version`
	// parameter)
	XactID string `json:"_xact_id,required"`
	// The timestamp the dataset event was created
	Created time.Time `json:"created,required" format:"date-time"`
	// Unique identifier for the dataset
	DatasetID string `json:"dataset_id,required" format:"uuid"`
	// Unique identifier for the project that the dataset belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// A unique identifier for the trace this dataset event belongs to
	RootSpanID string `json:"root_span_id,required"`
	// A unique identifier used to link different dataset events together as part of a
	// full trace. See the
	// [tracing guide](https://www.braintrust.dev/docs/guides/tracing) for full details
	// on tracing
	SpanID string `json:"span_id,required"`
	// The output of your application, including post-processing (an arbitrary, JSON
	// serializable object)
	Expected interface{} `json:"expected"`
	// The argument that uniquely define an input case (an arbitrary, JSON serializable
	// object)
	Input interface{} `json:"input"`
	// Whether this span is a root span
	IsRoot bool `json:"is_root,nullable"`
	// A dictionary with additional data about the test example, model outputs, or just
	// about anything else that's relevant, that you can use to help find and analyze
	// examples later. For example, you could log the `prompt`, example's `id`, or
	// anything else that would be useful to slice/dice later. The values in `metadata`
	// can be any JSON-serializable type, but its keys must be strings
	Metadata DatasetEventMetadata `json:"metadata,nullable"`
	// Indicates the event was copied from another object.
	Origin ObjectReference `json:"origin,nullable"`
	// A list of tags to log
	Tags []string `json:"tags,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		XactID      resp.Field
		Created     resp.Field
		DatasetID   resp.Field
		ProjectID   resp.Field
		RootSpanID  resp.Field
		SpanID      resp.Field
		Expected    resp.Field
		Input       resp.Field
		IsRoot      resp.Field
		Metadata    resp.Field
		Origin      resp.Field
		Tags        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DatasetEvent) RawJSON() string { return r.JSON.raw }
func (r *DatasetEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dictionary with additional data about the test example, model outputs, or just
// about anything else that's relevant, that you can use to help find and analyze
// examples later. For example, you could log the `prompt`, example's `id`, or
// anything else that would be useful to slice/dice later. The values in `metadata`
// can be any JSON-serializable type, but its keys must be strings
type DatasetEventMetadata struct {
	// The model used for this example
	Model       string                 `json:"model,nullable"`
	ExtraFields map[string]interface{} `json:",extras"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Model       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DatasetEventMetadata) RawJSON() string { return r.JSON.raw }
func (r *DatasetEventMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnvVar struct {
	// Unique identifier for the environment variable
	ID string `json:"id,required" format:"uuid"`
	// The name of the environment variable
	Name string `json:"name,required"`
	// The id of the object the environment variable is scoped for
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The type of the object the environment variable is scoped for
	//
	// Any of "organization", "project", "function".
	ObjectType EnvVarObjectType `json:"object_type,required"`
	// Date of environment variable creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Date the environment variable was last used
	Used time.Time `json:"used,nullable" format:"date-time"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Name        resp.Field
		ObjectID    resp.Field
		ObjectType  resp.Field
		Created     resp.Field
		Used        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnvVar) RawJSON() string { return r.JSON.raw }
func (r *EnvVar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the object the environment variable is scoped for
type EnvVarObjectType string

const (
	EnvVarObjectTypeOrganization EnvVarObjectType = "organization"
	EnvVarObjectTypeProject      EnvVarObjectType = "project"
	EnvVarObjectTypeFunction     EnvVarObjectType = "function"
)

type Experiment struct {
	// Unique identifier for the experiment
	ID string `json:"id,required" format:"uuid"`
	// Name of the experiment. Within a project, experiment names are unique
	Name string `json:"name,required"`
	// Unique identifier for the project that the experiment belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// Whether or not the experiment is public. Public experiments can be viewed by
	// anybody inside or outside the organization
	Public bool `json:"public,required"`
	// Id of default base experiment to compare against when viewing this experiment
	BaseExpID string `json:"base_exp_id,nullable" format:"uuid"`
	// Commit, taken directly from `repo_info.commit`
	Commit string `json:"commit,nullable"`
	// Date of experiment creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Identifier of the linked dataset, or null if the experiment is not linked to a
	// dataset
	DatasetID string `json:"dataset_id,nullable" format:"uuid"`
	// Version number of the linked dataset the experiment was run against. This can be
	// used to reproduce the experiment after the dataset has been modified.
	DatasetVersion string `json:"dataset_version,nullable"`
	// Date of experiment deletion, or null if the experiment is still active
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Textual description of the experiment
	Description string `json:"description,nullable"`
	// User-controlled metadata about the experiment
	Metadata map[string]interface{} `json:"metadata,nullable"`
	// Metadata about the state of the repo when the experiment was created
	RepoInfo RepoInfo `json:"repo_info,nullable"`
	// Identifies the user who created the experiment
	UserID string `json:"user_id,nullable" format:"uuid"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID             resp.Field
		Name           resp.Field
		ProjectID      resp.Field
		Public         resp.Field
		BaseExpID      resp.Field
		Commit         resp.Field
		Created        resp.Field
		DatasetID      resp.Field
		DatasetVersion resp.Field
		DeletedAt      resp.Field
		Description    resp.Field
		Metadata       resp.Field
		RepoInfo       resp.Field
		UserID         resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Experiment) RawJSON() string { return r.JSON.raw }
func (r *Experiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentEvent struct {
	// A unique identifier for the experiment event. If you don't provide one,
	// BrainTrust will generate one for you
	ID string `json:"id,required"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the experiment (see the
	// `version` parameter)
	XactID string `json:"_xact_id,required"`
	// The timestamp the experiment event was created
	Created time.Time `json:"created,required" format:"date-time"`
	// Unique identifier for the experiment
	ExperimentID string `json:"experiment_id,required" format:"uuid"`
	// Unique identifier for the project that the experiment belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// A unique identifier for the trace this experiment event belongs to
	RootSpanID string `json:"root_span_id,required"`
	// A unique identifier used to link different experiment events together as part of
	// a full trace. See the
	// [tracing guide](https://www.braintrust.dev/docs/guides/tracing) for full details
	// on tracing
	SpanID string `json:"span_id,required"`
	// Context is additional information about the code that produced the experiment
	// event. It is essentially the textual counterpart to `metrics`. Use the
	// `caller_*` attributes to track the location in code which produced the
	// experiment event
	Context ExperimentEventContext `json:"context,nullable"`
	// The error that occurred, if any.
	Error interface{} `json:"error"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not.
	// Braintrust currently does not compare `output` to `expected` for you, since
	// there are so many different ways to do that correctly. Instead, these values are
	// just used to help you navigate your experiments while digging into analyses.
	// However, we may later use these values to re-score outputs or fine-tune your
	// models
	Expected interface{} `json:"expected"`
	// The arguments that uniquely define a test case (an arbitrary, JSON serializable
	// object). Later on, Braintrust will use the `input` to know whether two test
	// cases are the same between experiments, so they should not contain
	// experiment-specific state. A simple rule of thumb is that if you run the same
	// experiment twice, the `input` should be identical
	Input interface{} `json:"input"`
	// Whether this span is a root span
	IsRoot bool `json:"is_root,nullable"`
	// A dictionary with additional data about the test example, model outputs, or just
	// about anything else that's relevant, that you can use to help find and analyze
	// examples later. For example, you could log the `prompt`, example's `id`, or
	// anything else that would be useful to slice/dice later. The values in `metadata`
	// can be any JSON-serializable type, but its keys must be strings
	Metadata ExperimentEventMetadata `json:"metadata,nullable"`
	// Metrics are numerical measurements tracking the execution of the code that
	// produced the experiment event. Use "start" and "end" to track the time span over
	// which the experiment event was produced
	Metrics ExperimentEventMetrics `json:"metrics,nullable"`
	// Indicates the event was copied from another object.
	Origin ObjectReference `json:"origin,nullable"`
	// The output of your application, including post-processing (an arbitrary, JSON
	// serializable object), that allows you to determine whether the result is correct
	// or not. For example, in an app that generates SQL queries, the `output` should
	// be the _result_ of the SQL query generated by the model, not the query itself,
	// because there may be multiple valid queries that answer a single question
	Output interface{} `json:"output"`
	// A dictionary of numeric values (between 0 and 1) to log. The scores should give
	// you a variety of signals that help you determine how accurate the outputs are
	// compared to what you expect and diagnose failures. For example, a summarization
	// app might have one score that tells you how accurate the summary is, and another
	// that measures the word similarity between the generated and grouth truth
	// summary. The word similarity score could help you determine whether the
	// summarization was covering similar concepts or not. You can use these scores to
	// help you sort, filter, and compare experiments
	Scores map[string]float64 `json:"scores,nullable"`
	// Human-identifying attributes of the span, such as name, type, etc.
	SpanAttributes SpanAttributes `json:"span_attributes,nullable"`
	// An array of the parent `span_ids` of this experiment event. This should be empty
	// for the root span of a trace, and should most often contain just one parent
	// element for subspans
	SpanParents []string `json:"span_parents,nullable"`
	// A list of tags to log
	Tags []string `json:"tags,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID             resp.Field
		XactID         resp.Field
		Created        resp.Field
		ExperimentID   resp.Field
		ProjectID      resp.Field
		RootSpanID     resp.Field
		SpanID         resp.Field
		Context        resp.Field
		Error          resp.Field
		Expected       resp.Field
		Input          resp.Field
		IsRoot         resp.Field
		Metadata       resp.Field
		Metrics        resp.Field
		Origin         resp.Field
		Output         resp.Field
		Scores         resp.Field
		SpanAttributes resp.Field
		SpanParents    resp.Field
		Tags           resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentEvent) RawJSON() string { return r.JSON.raw }
func (r *ExperimentEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Context is additional information about the code that produced the experiment
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the
// experiment event
type ExperimentEventContext struct {
	// Name of the file in code where the experiment event was created
	CallerFilename string `json:"caller_filename,nullable"`
	// The function in code which created the experiment event
	CallerFunctionname string `json:"caller_functionname,nullable"`
	// Line of code where the experiment event was created
	CallerLineno int64                  `json:"caller_lineno,nullable"`
	ExtraFields  map[string]interface{} `json:",extras"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CallerFilename     resp.Field
		CallerFunctionname resp.Field
		CallerLineno       resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentEventContext) RawJSON() string { return r.JSON.raw }
func (r *ExperimentEventContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dictionary with additional data about the test example, model outputs, or just
// about anything else that's relevant, that you can use to help find and analyze
// examples later. For example, you could log the `prompt`, example's `id`, or
// anything else that would be useful to slice/dice later. The values in `metadata`
// can be any JSON-serializable type, but its keys must be strings
type ExperimentEventMetadata struct {
	// The model used for this example
	Model       string                 `json:"model,nullable"`
	ExtraFields map[string]interface{} `json:",extras"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Model       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentEventMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExperimentEventMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the experiment event. Use "start" and "end" to track the time span over
// which the experiment event was produced
type ExperimentEventMetrics struct {
	// This metric is deprecated
	CallerFilename interface{} `json:"caller_filename"`
	// This metric is deprecated
	CallerFunctionname interface{} `json:"caller_functionname"`
	// This metric is deprecated
	CallerLineno interface{} `json:"caller_lineno"`
	// The number of tokens in the completion generated by the model (only set if this
	// is an LLM span)
	CompletionTokens int64 `json:"completion_tokens,nullable"`
	// A unix timestamp recording when the section of code which produced the
	// experiment event finished
	End float64 `json:"end,nullable"`
	// The number of tokens in the prompt used to generate the experiment event (only
	// set if this is an LLM span)
	PromptTokens int64 `json:"prompt_tokens,nullable"`
	// A unix timestamp recording when the section of code which produced the
	// experiment event started
	Start float64 `json:"start,nullable"`
	// The total number of tokens in the input and output of the experiment event.
	Tokens      int64              `json:"tokens,nullable"`
	ExtraFields map[string]float64 `json:",extras"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CallerFilename     resp.Field
		CallerFunctionname resp.Field
		CallerLineno       resp.Field
		CompletionTokens   resp.Field
		End                resp.Field
		PromptTokens       resp.Field
		Start              resp.Field
		Tokens             resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentEventMetrics) RawJSON() string { return r.JSON.raw }
func (r *ExperimentEventMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type FeedbackDatasetItemParam struct {
	// The id of the dataset event to log feedback for. This is the row `id` returned
	// by `POST /v1/dataset/{dataset_id}/insert`
	ID string `json:"id,required"`
	// An optional comment string to log about the dataset event
	Comment param.Opt[string] `json:"comment,omitzero"`
	// A dictionary with additional data about the feedback. If you have a `user_id`,
	// you can log it here and access it in the Braintrust UI. Note, this metadata does
	// not correspond to the main event itself, but rather the audit log attached to
	// the event.
	Metadata map[string]interface{} `json:"metadata,omitzero"`
	// The source of the feedback. Must be one of "external" (default), "app", or "api"
	//
	// Any of "app", "api", "external".
	Source FeedbackDatasetItemSource `json:"source,omitzero"`
	// A list of tags to log
	Tags []string `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FeedbackDatasetItemParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r FeedbackDatasetItemParam) MarshalJSON() (data []byte, err error) {
	type shadow FeedbackDatasetItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The source of the feedback. Must be one of "external" (default), "app", or "api"
type FeedbackDatasetItemSource string

const (
	FeedbackDatasetItemSourceApp      FeedbackDatasetItemSource = "app"
	FeedbackDatasetItemSourceAPI      FeedbackDatasetItemSource = "api"
	FeedbackDatasetItemSourceExternal FeedbackDatasetItemSource = "external"
)

// The property ID is required.
type FeedbackExperimentItemParam struct {
	// The id of the experiment event to log feedback for. This is the row `id`
	// returned by `POST /v1/experiment/{experiment_id}/insert`
	ID string `json:"id,required"`
	// An optional comment string to log about the experiment event
	Comment param.Opt[string] `json:"comment,omitzero"`
	// A dictionary with additional data about the feedback. If you have a `user_id`,
	// you can log it here and access it in the Braintrust UI. Note, this metadata does
	// not correspond to the main event itself, but rather the audit log attached to
	// the event.
	Metadata map[string]interface{} `json:"metadata,omitzero"`
	// A dictionary of numeric values (between 0 and 1) to log. These scores will be
	// merged into the existing scores for the experiment event
	Scores map[string]float64 `json:"scores,omitzero"`
	// The source of the feedback. Must be one of "external" (default), "app", or "api"
	//
	// Any of "app", "api", "external".
	Source FeedbackExperimentItemSource `json:"source,omitzero"`
	// A list of tags to log
	Tags []string `json:"tags,omitzero"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not
	Expected interface{} `json:"expected,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FeedbackExperimentItemParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r FeedbackExperimentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow FeedbackExperimentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The source of the feedback. Must be one of "external" (default), "app", or "api"
type FeedbackExperimentItemSource string

const (
	FeedbackExperimentItemSourceApp      FeedbackExperimentItemSource = "app"
	FeedbackExperimentItemSourceAPI      FeedbackExperimentItemSource = "api"
	FeedbackExperimentItemSourceExternal FeedbackExperimentItemSource = "external"
)

// The property ID is required.
type FeedbackProjectLogsItemParam struct {
	// The id of the project logs event to log feedback for. This is the row `id`
	// returned by `POST /v1/project_logs/{project_id}/insert`
	ID string `json:"id,required"`
	// An optional comment string to log about the project logs event
	Comment param.Opt[string] `json:"comment,omitzero"`
	// A dictionary with additional data about the feedback. If you have a `user_id`,
	// you can log it here and access it in the Braintrust UI. Note, this metadata does
	// not correspond to the main event itself, but rather the audit log attached to
	// the event.
	Metadata map[string]interface{} `json:"metadata,omitzero"`
	// A dictionary of numeric values (between 0 and 1) to log. These scores will be
	// merged into the existing scores for the project logs event
	Scores map[string]float64 `json:"scores,omitzero"`
	// The source of the feedback. Must be one of "external" (default), "app", or "api"
	//
	// Any of "app", "api", "external".
	Source FeedbackProjectLogsItemSource `json:"source,omitzero"`
	// A list of tags to log
	Tags []string `json:"tags,omitzero"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not
	Expected interface{} `json:"expected,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FeedbackProjectLogsItemParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r FeedbackProjectLogsItemParam) MarshalJSON() (data []byte, err error) {
	type shadow FeedbackProjectLogsItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The source of the feedback. Must be one of "external" (default), "app", or "api"
type FeedbackProjectLogsItemSource string

const (
	FeedbackProjectLogsItemSourceApp      FeedbackProjectLogsItemSource = "app"
	FeedbackProjectLogsItemSourceAPI      FeedbackProjectLogsItemSource = "api"
	FeedbackProjectLogsItemSourceExternal FeedbackProjectLogsItemSource = "external"
)

type FeedbackResponseSchema struct {
	// Any of "success".
	Status FeedbackResponseSchemaStatus `json:"status,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Status      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeedbackResponseSchema) RawJSON() string { return r.JSON.raw }
func (r *FeedbackResponseSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeedbackResponseSchemaStatus string

const (
	FeedbackResponseSchemaStatusSuccess FeedbackResponseSchemaStatus = "success"
)

type FetchDatasetEventsResponse struct {
	// A list of fetched events
	Events []DatasetEvent `json:"events,required"`
	// Pagination cursor
	//
	// Pass this string directly as the `cursor` param to your next fetch request to
	// get the next page of results. Not provided if the returned result set is empty.
	Cursor string `json:"cursor,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Events      resp.Field
		Cursor      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FetchDatasetEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *FetchDatasetEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FetchExperimentEventsResponse struct {
	// A list of fetched events
	Events []ExperimentEvent `json:"events,required"`
	// Pagination cursor
	//
	// Pass this string directly as the `cursor` param to your next fetch request to
	// get the next page of results. Not provided if the returned result set is empty.
	Cursor string `json:"cursor,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Events      resp.Field
		Cursor      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FetchExperimentEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *FetchExperimentEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FetchProjectLogsEventsResponse struct {
	// A list of fetched events
	Events []ProjectLogsEvent `json:"events,required"`
	// Pagination cursor
	//
	// Pass this string directly as the `cursor` param to your next fetch request to
	// get the next page of results. Not provided if the returned result set is empty.
	Cursor string `json:"cursor,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Events      resp.Field
		Cursor      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FetchProjectLogsEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *FetchProjectLogsEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Function struct {
	// Unique identifier for the prompt
	ID string `json:"id,required" format:"uuid"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the prompt (see the `version`
	// parameter)
	XactID       string                    `json:"_xact_id,required"`
	FunctionData FunctionFunctionDataUnion `json:"function_data,required"`
	// A literal 'p' which identifies the object as a project prompt
	//
	// Any of "p".
	LogID FunctionLogID `json:"log_id,required"`
	// Name of the prompt
	Name string `json:"name,required"`
	// Unique identifier for the organization
	OrgID string `json:"org_id,required" format:"uuid"`
	// Unique identifier for the project that the prompt belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// Unique identifier for the prompt
	Slug string `json:"slug,required"`
	// Date of prompt creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Textual description of the prompt
	Description string `json:"description,nullable"`
	// JSON schema for the function's parameters and return type
	FunctionSchema FunctionFunctionSchema `json:"function_schema,nullable"`
	// Any of "llm", "scorer", "task", "tool".
	FunctionType FunctionFunctionType `json:"function_type,nullable"`
	// User-controlled metadata about the prompt
	Metadata map[string]interface{} `json:"metadata,nullable"`
	Origin   FunctionOrigin         `json:"origin,nullable"`
	// The prompt, model, and its parameters
	PromptData PromptData `json:"prompt_data,nullable"`
	// A list of tags for the prompt
	Tags []string `json:"tags,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID             resp.Field
		XactID         resp.Field
		FunctionData   resp.Field
		LogID          resp.Field
		Name           resp.Field
		OrgID          resp.Field
		ProjectID      resp.Field
		Slug           resp.Field
		Created        resp.Field
		Description    resp.Field
		FunctionSchema resp.Field
		FunctionType   resp.Field
		Metadata       resp.Field
		Origin         resp.Field
		PromptData     resp.Field
		Tags           resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Function) RawJSON() string { return r.JSON.raw }
func (r *Function) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunctionFunctionDataUnion contains all possible properties and values from
// [FunctionFunctionDataPrompt], [FunctionFunctionDataCode],
// [FunctionFunctionDataGlobal].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunctionFunctionDataUnion struct {
	Type string `json:"type"`
	// This field is from variant [FunctionFunctionDataCode].
	Data FunctionFunctionDataCodeDataUnion `json:"data"`
	// This field is from variant [FunctionFunctionDataGlobal].
	Name string `json:"name"`
	JSON struct {
		Type resp.Field
		Data resp.Field
		Name resp.Field
		raw  string
	} `json:"-"`
}

func (u FunctionFunctionDataUnion) AsPrompt() (v FunctionFunctionDataPrompt) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionFunctionDataUnion) AsCode() (v FunctionFunctionDataCode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionFunctionDataUnion) AsGlobal() (v FunctionFunctionDataGlobal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunctionFunctionDataUnion) RawJSON() string { return u.JSON.raw }

func (r *FunctionFunctionDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionFunctionDataPrompt struct {
	// Any of "prompt".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionFunctionDataPrompt) RawJSON() string { return r.JSON.raw }
func (r *FunctionFunctionDataPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionFunctionDataCode struct {
	Data FunctionFunctionDataCodeDataUnion `json:"data,required"`
	// Any of "code".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data        resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionFunctionDataCode) RawJSON() string { return r.JSON.raw }
func (r *FunctionFunctionDataCode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunctionFunctionDataCodeDataUnion contains all possible properties and values
// from [FunctionFunctionDataCodeDataBundle], [FunctionFunctionDataCodeDataInline].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunctionFunctionDataCodeDataUnion struct {
	// This field is from variant [FunctionFunctionDataCodeDataBundle].
	BundleID string `json:"bundle_id"`
	// This field is from variant [FunctionFunctionDataCodeDataBundle].
	Location CodeBundleLocationUnion `json:"location"`
	// This field is a union of [CodeBundleRuntimeContext],
	// [FunctionFunctionDataCodeDataInlineRuntimeContext]
	RuntimeContext FunctionFunctionDataCodeDataUnionRuntimeContext `json:"runtime_context"`
	// This field is from variant [FunctionFunctionDataCodeDataBundle].
	Preview string `json:"preview"`
	Type    string `json:"type"`
	// This field is from variant [FunctionFunctionDataCodeDataInline].
	Code string `json:"code"`
	JSON struct {
		BundleID       resp.Field
		Location       resp.Field
		RuntimeContext resp.Field
		Preview        resp.Field
		Type           resp.Field
		Code           resp.Field
		raw            string
	} `json:"-"`
}

func (u FunctionFunctionDataCodeDataUnion) AsBundle() (v FunctionFunctionDataCodeDataBundle) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionFunctionDataCodeDataUnion) AsInline() (v FunctionFunctionDataCodeDataInline) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunctionFunctionDataCodeDataUnion) RawJSON() string { return u.JSON.raw }

func (r *FunctionFunctionDataCodeDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunctionFunctionDataCodeDataUnionRuntimeContext is an implicit subunion of
// [FunctionFunctionDataCodeDataUnion].
// FunctionFunctionDataCodeDataUnionRuntimeContext provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [FunctionFunctionDataCodeDataUnion].
type FunctionFunctionDataCodeDataUnionRuntimeContext struct {
	Runtime string `json:"runtime"`
	Version string `json:"version"`
	JSON    struct {
		Runtime resp.Field
		Version resp.Field
		raw     string
	} `json:"-"`
}

func (r *FunctionFunctionDataCodeDataUnionRuntimeContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionFunctionDataCodeDataBundle struct {
	// Any of "bundle".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
	CodeBundle
}

// Returns the unmodified JSON received from the API
func (r FunctionFunctionDataCodeDataBundle) RawJSON() string { return r.JSON.raw }
func (r *FunctionFunctionDataCodeDataBundle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionFunctionDataCodeDataInline struct {
	Code           string                                           `json:"code,required"`
	RuntimeContext FunctionFunctionDataCodeDataInlineRuntimeContext `json:"runtime_context,required"`
	// Any of "inline".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Code           resp.Field
		RuntimeContext resp.Field
		Type           resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionFunctionDataCodeDataInline) RawJSON() string { return r.JSON.raw }
func (r *FunctionFunctionDataCodeDataInline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionFunctionDataCodeDataInlineRuntimeContext struct {
	// Any of "node", "python".
	Runtime string `json:"runtime,required"`
	Version string `json:"version,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Runtime     resp.Field
		Version     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionFunctionDataCodeDataInlineRuntimeContext) RawJSON() string { return r.JSON.raw }
func (r *FunctionFunctionDataCodeDataInlineRuntimeContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionFunctionDataGlobal struct {
	Name string `json:"name,required"`
	// Any of "global".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionFunctionDataGlobal) RawJSON() string { return r.JSON.raw }
func (r *FunctionFunctionDataGlobal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A literal 'p' which identifies the object as a project prompt
type FunctionLogID string

const (
	FunctionLogIDP FunctionLogID = "p"
)

// JSON schema for the function's parameters and return type
type FunctionFunctionSchema struct {
	Parameters interface{} `json:"parameters"`
	Returns    interface{} `json:"returns"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Parameters  resp.Field
		Returns     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionFunctionSchema) RawJSON() string { return r.JSON.raw }
func (r *FunctionFunctionSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionFunctionType string

const (
	FunctionFunctionTypeLlm    FunctionFunctionType = "llm"
	FunctionFunctionTypeScorer FunctionFunctionType = "scorer"
	FunctionFunctionTypeTask   FunctionFunctionType = "task"
	FunctionFunctionTypeTool   FunctionFunctionType = "tool"
)

type FunctionOrigin struct {
	// Id of the object the function is originating from
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType ACLObjectType `json:"object_type,required"`
	// The function exists for internal purposes and should not be displayed in the
	// list of functions.
	Internal bool `json:"internal,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ObjectID    resp.Field
		ObjectType  resp.Field
		Internal    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionOrigin) RawJSON() string { return r.JSON.raw }
func (r *FunctionOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A group is a collection of users which can be assigned an ACL
//
// Groups can consist of individual users, as well as a set of groups they inherit
// from
type Group struct {
	// Unique identifier for the group
	ID string `json:"id,required" format:"uuid"`
	// Name of the group
	Name string `json:"name,required"`
	// Unique id for the organization that the group belongs under
	//
	// It is forbidden to change the org after creating a group
	OrgID string `json:"org_id,required" format:"uuid"`
	// Date of group creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Date of group deletion, or null if the group is still active
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Textual description of the group
	Description string `json:"description,nullable"`
	// Ids of the groups this group inherits from
	//
	// An inheriting group has all the users contained in its member groups, as well as
	// all of their inherited users
	MemberGroups []string `json:"member_groups,nullable" format:"uuid"`
	// Ids of users which belong to this group
	MemberUsers []string `json:"member_users,nullable" format:"uuid"`
	// Identifies the user who created the group
	UserID string `json:"user_id,nullable" format:"uuid"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID           resp.Field
		Name         resp.Field
		OrgID        resp.Field
		Created      resp.Field
		DeletedAt    resp.Field
		Description  resp.Field
		MemberGroups resp.Field
		MemberUsers  resp.Field
		UserID       resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Group) RawJSON() string { return r.JSON.raw }
func (r *Group) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dataset event
type InsertDatasetEventParam struct {
	// A unique identifier for the dataset event. If you don't provide one, BrainTrust
	// will generate one for you
	ID param.Opt[string] `json:"id,omitzero"`
	// The `_is_merge` field controls how the row is merged with any existing row with
	// the same id in the DB. By default (or when set to `false`), the existing row is
	// completely replaced by the new row. When set to `true`, the new row is
	// deep-merged into the existing row, if one is found. If no existing row is found,
	// the new row is inserted as is.
	//
	// For example, say there is an existing row in the DB
	// `{"id": "foo", "input": {"a": 5, "b": 10}}`. If we merge a new row as
	// `{"_is_merge": true, "id": "foo", "input": {"b": 11, "c": 20}}`, the new row
	// will be `{"id": "foo", "input": {"a": 5, "b": 11, "c": 20}}`. If we replace the
	// new row as `{"id": "foo", "input": {"b": 11, "c": 20}}`, the new row will be
	// `{"id": "foo", "input": {"b": 11, "c": 20}}`
	IsMerge param.Opt[bool] `json:"_is_merge,omitzero"`
	// Pass `_object_delete=true` to mark the dataset event deleted. Deleted events
	// will not show up in subsequent fetches for this dataset
	ObjectDelete param.Opt[bool] `json:"_object_delete,omitzero"`
	// DEPRECATED: The `_parent_id` field is deprecated and should not be used. Support
	// for `_parent_id` will be dropped in a future version of Braintrust. Log
	// `span_id`, `root_span_id`, and `span_parents` explicitly instead.
	//
	// Use the `_parent_id` field to create this row as a subspan of an existing row.
	// Tracking hierarchical relationships are important for tracing (see the
	// [guide](https://www.braintrust.dev/docs/guides/tracing) for full details).
	//
	// For example, say we have logged a row
	// `{"id": "abc", "input": "foo", "output": "bar", "expected": "boo", "scores": {"correctness": 0.33}}`.
	// We can create a sub-span of the parent row by logging
	// `{"_parent_id": "abc", "id": "llm_call", "input": {"prompt": "What comes after foo?"}, "output": "bar", "metrics": {"tokens": 1}}`.
	// In the webapp, only the root span row `"abc"` will show up in the summary view.
	// You can view the full trace hierarchy (in this case, the `"llm_call"` row) by
	// clicking on the "abc" row.
	//
	// If the row is being merged into an existing row, this field will be ignored.
	ParentID param.Opt[string] `json:"_parent_id,omitzero"`
	// The timestamp the dataset event was created
	Created param.Opt[time.Time] `json:"created,omitzero" format:"date-time"`
	// Use `span_id`, `root_span_id`, and `span_parents` instead of `_parent_id`, which
	// is now deprecated. The span_id is a unique identifier describing the row's place
	// in the a trace, and the root_span_id is a unique identifier for the whole trace.
	// See the [guide](https://www.braintrust.dev/docs/guides/tracing) for full
	// details.
	//
	// For example, say we have logged a row
	// `{"id": "abc", "span_id": "span0", "root_span_id": "root_span0", "input": "foo", "output": "bar", "expected": "boo", "scores": {"correctness": 0.33}}`.
	// We can create a sub-span of the parent row by logging
	// `{"id": "llm_call", "span_id": "span1", "root_span_id": "root_span0", "span_parents": ["span0"], "input": {"prompt": "What comes after foo?"}, "output": "bar", "metrics": {"tokens": 1}}`.
	// In the webapp, only the root span row `"abc"` will show up in the summary view.
	// You can view the full trace hierarchy (in this case, the `"llm_call"` row) by
	// clicking on the "abc" row.
	//
	// If the row is being merged into an existing row, this field will be ignored.
	RootSpanID param.Opt[string] `json:"root_span_id,omitzero"`
	// Use `span_id`, `root_span_id`, and `span_parents` instead of `_parent_id`, which
	// is now deprecated. The span_id is a unique identifier describing the row's place
	// in the a trace, and the root_span_id is a unique identifier for the whole trace.
	// See the [guide](https://www.braintrust.dev/docs/guides/tracing) for full
	// details.
	//
	// For example, say we have logged a row
	// `{"id": "abc", "span_id": "span0", "root_span_id": "root_span0", "input": "foo", "output": "bar", "expected": "boo", "scores": {"correctness": 0.33}}`.
	// We can create a sub-span of the parent row by logging
	// `{"id": "llm_call", "span_id": "span1", "root_span_id": "root_span0", "span_parents": ["span0"], "input": {"prompt": "What comes after foo?"}, "output": "bar", "metrics": {"tokens": 1}}`.
	// In the webapp, only the root span row `"abc"` will show up in the summary view.
	// You can view the full trace hierarchy (in this case, the `"llm_call"` row) by
	// clicking on the "abc" row.
	//
	// If the row is being merged into an existing row, this field will be ignored.
	SpanID param.Opt[string] `json:"span_id,omitzero"`
	// The `_merge_paths` field allows controlling the depth of the merge, when
	// `_is_merge=true`. `_merge_paths` is a list of paths, where each path is a list
	// of field names. The deep merge will not descend below any of the specified merge
	// paths.
	//
	// For example, say there is an existing row in the DB
	// `{"id": "foo", "input": {"a": {"b": 10}, "c": {"d": 20}}, "output": {"a": 20}}`.
	// If we merge a new row as
	// `{"_is_merge": true, "_merge_paths": [["input", "a"], ["output"]], "input": {"a": {"q": 30}, "c": {"e": 30}, "bar": "baz"}, "output": {"d": 40}}`,
	// the new row will be
	// `{"id": "foo": "input": {"a": {"q": 30}, "c": {"d": 20, "e": 30}, "bar": "baz"}, "output": {"d": 40}}`.
	// In this case, due to the merge paths, we have replaced `input.a` and `output`,
	// but have still deep-merged `input` and `input.c`.
	MergePaths [][]string `json:"_merge_paths,omitzero"`
	// A dictionary with additional data about the test example, model outputs, or just
	// about anything else that's relevant, that you can use to help find and analyze
	// examples later. For example, you could log the `prompt`, example's `id`, or
	// anything else that would be useful to slice/dice later. The values in `metadata`
	// can be any JSON-serializable type, but its keys must be strings
	Metadata InsertDatasetEventMetadataParam `json:"metadata,omitzero"`
	// Indicates the event was copied from another object.
	Origin ObjectReferenceParam `json:"origin,omitzero"`
	// Use `span_id`, `root_span_id`, and `span_parents` instead of `_parent_id`, which
	// is now deprecated. The span_id is a unique identifier describing the row's place
	// in the a trace, and the root_span_id is a unique identifier for the whole trace.
	// See the [guide](https://www.braintrust.dev/docs/guides/tracing) for full
	// details.
	//
	// For example, say we have logged a row
	// `{"id": "abc", "span_id": "span0", "root_span_id": "root_span0", "input": "foo", "output": "bar", "expected": "boo", "scores": {"correctness": 0.33}}`.
	// We can create a sub-span of the parent row by logging
	// `{"id": "llm_call", "span_id": "span1", "root_span_id": "root_span0", "span_parents": ["span0"], "input": {"prompt": "What comes after foo?"}, "output": "bar", "metrics": {"tokens": 1}}`.
	// In the webapp, only the root span row `"abc"` will show up in the summary view.
	// You can view the full trace hierarchy (in this case, the `"llm_call"` row) by
	// clicking on the "abc" row.
	//
	// If the row is being merged into an existing row, this field will be ignored.
	SpanParents []string `json:"span_parents,omitzero"`
	// A list of tags to log
	Tags []string `json:"tags,omitzero"`
	// The output of your application, including post-processing (an arbitrary, JSON
	// serializable object)
	Expected interface{} `json:"expected,omitzero"`
	// The argument that uniquely define an input case (an arbitrary, JSON serializable
	// object)
	Input interface{} `json:"input,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InsertDatasetEventParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r InsertDatasetEventParam) MarshalJSON() (data []byte, err error) {
	type shadow InsertDatasetEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A dictionary with additional data about the test example, model outputs, or just
// about anything else that's relevant, that you can use to help find and analyze
// examples later. For example, you could log the `prompt`, example's `id`, or
// anything else that would be useful to slice/dice later. The values in `metadata`
// can be any JSON-serializable type, but its keys must be strings
type InsertDatasetEventMetadataParam struct {
	// The model used for this example
	Model       param.Opt[string]      `json:"model,omitzero"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InsertDatasetEventMetadataParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r InsertDatasetEventMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow InsertDatasetEventMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type InsertEventsResponse struct {
	// The ids of all rows that were inserted, aligning one-to-one with the rows
	// provided as input
	RowIDs []string `json:"row_ids,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		RowIDs      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsertEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *InsertEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An experiment event
type InsertExperimentEventParam struct {
	// A unique identifier for the experiment event. If you don't provide one,
	// BrainTrust will generate one for you
	ID param.Opt[string] `json:"id,omitzero"`
	// The `_is_merge` field controls how the row is merged with any existing row with
	// the same id in the DB. By default (or when set to `false`), the existing row is
	// completely replaced by the new row. When set to `true`, the new row is
	// deep-merged into the existing row, if one is found. If no existing row is found,
	// the new row is inserted as is.
	//
	// For example, say there is an existing row in the DB
	// `{"id": "foo", "input": {"a": 5, "b": 10}}`. If we merge a new row as
	// `{"_is_merge": true, "id": "foo", "input": {"b": 11, "c": 20}}`, the new row
	// will be `{"id": "foo", "input": {"a": 5, "b": 11, "c": 20}}`. If we replace the
	// new row as `{"id": "foo", "input": {"b": 11, "c": 20}}`, the new row will be
	// `{"id": "foo", "input": {"b": 11, "c": 20}}`
	IsMerge param.Opt[bool] `json:"_is_merge,omitzero"`
	// Pass `_object_delete=true` to mark the experiment event deleted. Deleted events
	// will not show up in subsequent fetches for this experiment
	ObjectDelete param.Opt[bool] `json:"_object_delete,omitzero"`
	// DEPRECATED: The `_parent_id` field is deprecated and should not be used. Support
	// for `_parent_id` will be dropped in a future version of Braintrust. Log
	// `span_id`, `root_span_id`, and `span_parents` explicitly instead.
	//
	// Use the `_parent_id` field to create this row as a subspan of an existing row.
	// Tracking hierarchical relationships are important for tracing (see the
	// [guide](https://www.braintrust.dev/docs/guides/tracing) for full details).
	//
	// For example, say we have logged a row
	// `{"id": "abc", "input": "foo", "output": "bar", "expected": "boo", "scores": {"correctness": 0.33}}`.
	// We can create a sub-span of the parent row by logging
	// `{"_parent_id": "abc", "id": "llm_call", "input": {"prompt": "What comes after foo?"}, "output": "bar", "metrics": {"tokens": 1}}`.
	// In the webapp, only the root span row `"abc"` will show up in the summary view.
	// You can view the full trace hierarchy (in this case, the `"llm_call"` row) by
	// clicking on the "abc" row.
	//
	// If the row is being merged into an existing row, this field will be ignored.
	ParentID param.Opt[string] `json:"_parent_id,omitzero"`
	// The timestamp the experiment event was created
	Created param.Opt[time.Time] `json:"created,omitzero" format:"date-time"`
	// Use `span_id`, `root_span_id`, and `span_parents` instead of `_parent_id`, which
	// is now deprecated. The span_id is a unique identifier describing the row's place
	// in the a trace, and the root_span_id is a unique identifier for the whole trace.
	// See the [guide](https://www.braintrust.dev/docs/guides/tracing) for full
	// details.
	//
	// For example, say we have logged a row
	// `{"id": "abc", "span_id": "span0", "root_span_id": "root_span0", "input": "foo", "output": "bar", "expected": "boo", "scores": {"correctness": 0.33}}`.
	// We can create a sub-span of the parent row by logging
	// `{"id": "llm_call", "span_id": "span1", "root_span_id": "root_span0", "span_parents": ["span0"], "input": {"prompt": "What comes after foo?"}, "output": "bar", "metrics": {"tokens": 1}}`.
	// In the webapp, only the root span row `"abc"` will show up in the summary view.
	// You can view the full trace hierarchy (in this case, the `"llm_call"` row) by
	// clicking on the "abc" row.
	//
	// If the row is being merged into an existing row, this field will be ignored.
	RootSpanID param.Opt[string] `json:"root_span_id,omitzero"`
	// Use `span_id`, `root_span_id`, and `span_parents` instead of `_parent_id`, which
	// is now deprecated. The span_id is a unique identifier describing the row's place
	// in the a trace, and the root_span_id is a unique identifier for the whole trace.
	// See the [guide](https://www.braintrust.dev/docs/guides/tracing) for full
	// details.
	//
	// For example, say we have logged a row
	// `{"id": "abc", "span_id": "span0", "root_span_id": "root_span0", "input": "foo", "output": "bar", "expected": "boo", "scores": {"correctness": 0.33}}`.
	// We can create a sub-span of the parent row by logging
	// `{"id": "llm_call", "span_id": "span1", "root_span_id": "root_span0", "span_parents": ["span0"], "input": {"prompt": "What comes after foo?"}, "output": "bar", "metrics": {"tokens": 1}}`.
	// In the webapp, only the root span row `"abc"` will show up in the summary view.
	// You can view the full trace hierarchy (in this case, the `"llm_call"` row) by
	// clicking on the "abc" row.
	//
	// If the row is being merged into an existing row, this field will be ignored.
	SpanID param.Opt[string] `json:"span_id,omitzero"`
	// The `_merge_paths` field allows controlling the depth of the merge, when
	// `_is_merge=true`. `_merge_paths` is a list of paths, where each path is a list
	// of field names. The deep merge will not descend below any of the specified merge
	// paths.
	//
	// For example, say there is an existing row in the DB
	// `{"id": "foo", "input": {"a": {"b": 10}, "c": {"d": 20}}, "output": {"a": 20}}`.
	// If we merge a new row as
	// `{"_is_merge": true, "_merge_paths": [["input", "a"], ["output"]], "input": {"a": {"q": 30}, "c": {"e": 30}, "bar": "baz"}, "output": {"d": 40}}`,
	// the new row will be
	// `{"id": "foo": "input": {"a": {"q": 30}, "c": {"d": 20, "e": 30}, "bar": "baz"}, "output": {"d": 40}}`.
	// In this case, due to the merge paths, we have replaced `input.a` and `output`,
	// but have still deep-merged `input` and `input.c`.
	MergePaths [][]string `json:"_merge_paths,omitzero"`
	// Context is additional information about the code that produced the experiment
	// event. It is essentially the textual counterpart to `metrics`. Use the
	// `caller_*` attributes to track the location in code which produced the
	// experiment event
	Context InsertExperimentEventContextParam `json:"context,omitzero"`
	// A dictionary with additional data about the test example, model outputs, or just
	// about anything else that's relevant, that you can use to help find and analyze
	// examples later. For example, you could log the `prompt`, example's `id`, or
	// anything else that would be useful to slice/dice later. The values in `metadata`
	// can be any JSON-serializable type, but its keys must be strings
	Metadata InsertExperimentEventMetadataParam `json:"metadata,omitzero"`
	// Metrics are numerical measurements tracking the execution of the code that
	// produced the experiment event. Use "start" and "end" to track the time span over
	// which the experiment event was produced
	Metrics InsertExperimentEventMetricsParam `json:"metrics,omitzero"`
	// Indicates the event was copied from another object.
	Origin ObjectReferenceParam `json:"origin,omitzero"`
	// A dictionary of numeric values (between 0 and 1) to log. The scores should give
	// you a variety of signals that help you determine how accurate the outputs are
	// compared to what you expect and diagnose failures. For example, a summarization
	// app might have one score that tells you how accurate the summary is, and another
	// that measures the word similarity between the generated and grouth truth
	// summary. The word similarity score could help you determine whether the
	// summarization was covering similar concepts or not. You can use these scores to
	// help you sort, filter, and compare experiments
	Scores map[string]float64 `json:"scores,omitzero"`
	// Human-identifying attributes of the span, such as name, type, etc.
	SpanAttributes SpanAttributesParam `json:"span_attributes,omitzero"`
	// Use `span_id`, `root_span_id`, and `span_parents` instead of `_parent_id`, which
	// is now deprecated. The span_id is a unique identifier describing the row's place
	// in the a trace, and the root_span_id is a unique identifier for the whole trace.
	// See the [guide](https://www.braintrust.dev/docs/guides/tracing) for full
	// details.
	//
	// For example, say we have logged a row
	// `{"id": "abc", "span_id": "span0", "root_span_id": "root_span0", "input": "foo", "output": "bar", "expected": "boo", "scores": {"correctness": 0.33}}`.
	// We can create a sub-span of the parent row by logging
	// `{"id": "llm_call", "span_id": "span1", "root_span_id": "root_span0", "span_parents": ["span0"], "input": {"prompt": "What comes after foo?"}, "output": "bar", "metrics": {"tokens": 1}}`.
	// In the webapp, only the root span row `"abc"` will show up in the summary view.
	// You can view the full trace hierarchy (in this case, the `"llm_call"` row) by
	// clicking on the "abc" row.
	//
	// If the row is being merged into an existing row, this field will be ignored.
	SpanParents []string `json:"span_parents,omitzero"`
	// A list of tags to log
	Tags []string `json:"tags,omitzero"`
	// The error that occurred, if any.
	Error interface{} `json:"error,omitzero"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not.
	// Braintrust currently does not compare `output` to `expected` for you, since
	// there are so many different ways to do that correctly. Instead, these values are
	// just used to help you navigate your experiments while digging into analyses.
	// However, we may later use these values to re-score outputs or fine-tune your
	// models
	Expected interface{} `json:"expected,omitzero"`
	// The arguments that uniquely define a test case (an arbitrary, JSON serializable
	// object). Later on, Braintrust will use the `input` to know whether two test
	// cases are the same between experiments, so they should not contain
	// experiment-specific state. A simple rule of thumb is that if you run the same
	// experiment twice, the `input` should be identical
	Input interface{} `json:"input,omitzero"`
	// The output of your application, including post-processing (an arbitrary, JSON
	// serializable object), that allows you to determine whether the result is correct
	// or not. For example, in an app that generates SQL queries, the `output` should
	// be the _result_ of the SQL query generated by the model, not the query itself,
	// because there may be multiple valid queries that answer a single question
	Output interface{} `json:"output,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InsertExperimentEventParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r InsertExperimentEventParam) MarshalJSON() (data []byte, err error) {
	type shadow InsertExperimentEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Context is additional information about the code that produced the experiment
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the
// experiment event
type InsertExperimentEventContextParam struct {
	// Name of the file in code where the experiment event was created
	CallerFilename param.Opt[string] `json:"caller_filename,omitzero"`
	// The function in code which created the experiment event
	CallerFunctionname param.Opt[string] `json:"caller_functionname,omitzero"`
	// Line of code where the experiment event was created
	CallerLineno param.Opt[int64]       `json:"caller_lineno,omitzero"`
	ExtraFields  map[string]interface{} `json:"-,extras"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InsertExperimentEventContextParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InsertExperimentEventContextParam) MarshalJSON() (data []byte, err error) {
	type shadow InsertExperimentEventContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A dictionary with additional data about the test example, model outputs, or just
// about anything else that's relevant, that you can use to help find and analyze
// examples later. For example, you could log the `prompt`, example's `id`, or
// anything else that would be useful to slice/dice later. The values in `metadata`
// can be any JSON-serializable type, but its keys must be strings
type InsertExperimentEventMetadataParam struct {
	// The model used for this example
	Model       param.Opt[string]      `json:"model,omitzero"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InsertExperimentEventMetadataParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InsertExperimentEventMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow InsertExperimentEventMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the experiment event. Use "start" and "end" to track the time span over
// which the experiment event was produced
type InsertExperimentEventMetricsParam struct {
	// The number of tokens in the completion generated by the model (only set if this
	// is an LLM span)
	CompletionTokens param.Opt[int64] `json:"completion_tokens,omitzero"`
	// A unix timestamp recording when the section of code which produced the
	// experiment event finished
	End param.Opt[float64] `json:"end,omitzero"`
	// The number of tokens in the prompt used to generate the experiment event (only
	// set if this is an LLM span)
	PromptTokens param.Opt[int64] `json:"prompt_tokens,omitzero"`
	// A unix timestamp recording when the section of code which produced the
	// experiment event started
	Start param.Opt[float64] `json:"start,omitzero"`
	// The total number of tokens in the input and output of the experiment event.
	Tokens param.Opt[int64] `json:"tokens,omitzero"`
	// This metric is deprecated
	CallerFilename interface{} `json:"caller_filename,omitzero"`
	// This metric is deprecated
	CallerFunctionname interface{} `json:"caller_functionname,omitzero"`
	// This metric is deprecated
	CallerLineno interface{}        `json:"caller_lineno,omitzero"`
	ExtraFields  map[string]float64 `json:"-,extras"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InsertExperimentEventMetricsParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InsertExperimentEventMetricsParam) MarshalJSON() (data []byte, err error) {
	type shadow InsertExperimentEventMetricsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A project logs event
type InsertProjectLogsEventParam struct {
	// A unique identifier for the project logs event. If you don't provide one,
	// BrainTrust will generate one for you
	ID param.Opt[string] `json:"id,omitzero"`
	// The `_is_merge` field controls how the row is merged with any existing row with
	// the same id in the DB. By default (or when set to `false`), the existing row is
	// completely replaced by the new row. When set to `true`, the new row is
	// deep-merged into the existing row, if one is found. If no existing row is found,
	// the new row is inserted as is.
	//
	// For example, say there is an existing row in the DB
	// `{"id": "foo", "input": {"a": 5, "b": 10}}`. If we merge a new row as
	// `{"_is_merge": true, "id": "foo", "input": {"b": 11, "c": 20}}`, the new row
	// will be `{"id": "foo", "input": {"a": 5, "b": 11, "c": 20}}`. If we replace the
	// new row as `{"id": "foo", "input": {"b": 11, "c": 20}}`, the new row will be
	// `{"id": "foo", "input": {"b": 11, "c": 20}}`
	IsMerge param.Opt[bool] `json:"_is_merge,omitzero"`
	// Pass `_object_delete=true` to mark the project logs event deleted. Deleted
	// events will not show up in subsequent fetches for this project logs
	ObjectDelete param.Opt[bool] `json:"_object_delete,omitzero"`
	// DEPRECATED: The `_parent_id` field is deprecated and should not be used. Support
	// for `_parent_id` will be dropped in a future version of Braintrust. Log
	// `span_id`, `root_span_id`, and `span_parents` explicitly instead.
	//
	// Use the `_parent_id` field to create this row as a subspan of an existing row.
	// Tracking hierarchical relationships are important for tracing (see the
	// [guide](https://www.braintrust.dev/docs/guides/tracing) for full details).
	//
	// For example, say we have logged a row
	// `{"id": "abc", "input": "foo", "output": "bar", "expected": "boo", "scores": {"correctness": 0.33}}`.
	// We can create a sub-span of the parent row by logging
	// `{"_parent_id": "abc", "id": "llm_call", "input": {"prompt": "What comes after foo?"}, "output": "bar", "metrics": {"tokens": 1}}`.
	// In the webapp, only the root span row `"abc"` will show up in the summary view.
	// You can view the full trace hierarchy (in this case, the `"llm_call"` row) by
	// clicking on the "abc" row.
	//
	// If the row is being merged into an existing row, this field will be ignored.
	ParentID param.Opt[string] `json:"_parent_id,omitzero"`
	// The timestamp the project logs event was created
	Created param.Opt[time.Time] `json:"created,omitzero" format:"date-time"`
	// Use `span_id`, `root_span_id`, and `span_parents` instead of `_parent_id`, which
	// is now deprecated. The span_id is a unique identifier describing the row's place
	// in the a trace, and the root_span_id is a unique identifier for the whole trace.
	// See the [guide](https://www.braintrust.dev/docs/guides/tracing) for full
	// details.
	//
	// For example, say we have logged a row
	// `{"id": "abc", "span_id": "span0", "root_span_id": "root_span0", "input": "foo", "output": "bar", "expected": "boo", "scores": {"correctness": 0.33}}`.
	// We can create a sub-span of the parent row by logging
	// `{"id": "llm_call", "span_id": "span1", "root_span_id": "root_span0", "span_parents": ["span0"], "input": {"prompt": "What comes after foo?"}, "output": "bar", "metrics": {"tokens": 1}}`.
	// In the webapp, only the root span row `"abc"` will show up in the summary view.
	// You can view the full trace hierarchy (in this case, the `"llm_call"` row) by
	// clicking on the "abc" row.
	//
	// If the row is being merged into an existing row, this field will be ignored.
	RootSpanID param.Opt[string] `json:"root_span_id,omitzero"`
	// Use `span_id`, `root_span_id`, and `span_parents` instead of `_parent_id`, which
	// is now deprecated. The span_id is a unique identifier describing the row's place
	// in the a trace, and the root_span_id is a unique identifier for the whole trace.
	// See the [guide](https://www.braintrust.dev/docs/guides/tracing) for full
	// details.
	//
	// For example, say we have logged a row
	// `{"id": "abc", "span_id": "span0", "root_span_id": "root_span0", "input": "foo", "output": "bar", "expected": "boo", "scores": {"correctness": 0.33}}`.
	// We can create a sub-span of the parent row by logging
	// `{"id": "llm_call", "span_id": "span1", "root_span_id": "root_span0", "span_parents": ["span0"], "input": {"prompt": "What comes after foo?"}, "output": "bar", "metrics": {"tokens": 1}}`.
	// In the webapp, only the root span row `"abc"` will show up in the summary view.
	// You can view the full trace hierarchy (in this case, the `"llm_call"` row) by
	// clicking on the "abc" row.
	//
	// If the row is being merged into an existing row, this field will be ignored.
	SpanID param.Opt[string] `json:"span_id,omitzero"`
	// The `_merge_paths` field allows controlling the depth of the merge, when
	// `_is_merge=true`. `_merge_paths` is a list of paths, where each path is a list
	// of field names. The deep merge will not descend below any of the specified merge
	// paths.
	//
	// For example, say there is an existing row in the DB
	// `{"id": "foo", "input": {"a": {"b": 10}, "c": {"d": 20}}, "output": {"a": 20}}`.
	// If we merge a new row as
	// `{"_is_merge": true, "_merge_paths": [["input", "a"], ["output"]], "input": {"a": {"q": 30}, "c": {"e": 30}, "bar": "baz"}, "output": {"d": 40}}`,
	// the new row will be
	// `{"id": "foo": "input": {"a": {"q": 30}, "c": {"d": 20, "e": 30}, "bar": "baz"}, "output": {"d": 40}}`.
	// In this case, due to the merge paths, we have replaced `input.a` and `output`,
	// but have still deep-merged `input` and `input.c`.
	MergePaths [][]string `json:"_merge_paths,omitzero"`
	// Context is additional information about the code that produced the project logs
	// event. It is essentially the textual counterpart to `metrics`. Use the
	// `caller_*` attributes to track the location in code which produced the project
	// logs event
	Context InsertProjectLogsEventContextParam `json:"context,omitzero"`
	// A dictionary with additional data about the test example, model outputs, or just
	// about anything else that's relevant, that you can use to help find and analyze
	// examples later. For example, you could log the `prompt`, example's `id`, or
	// anything else that would be useful to slice/dice later. The values in `metadata`
	// can be any JSON-serializable type, but its keys must be strings
	Metadata InsertProjectLogsEventMetadataParam `json:"metadata,omitzero"`
	// Metrics are numerical measurements tracking the execution of the code that
	// produced the project logs event. Use "start" and "end" to track the time span
	// over which the project logs event was produced
	Metrics InsertProjectLogsEventMetricsParam `json:"metrics,omitzero"`
	// Indicates the event was copied from another object.
	Origin ObjectReferenceParam `json:"origin,omitzero"`
	// A dictionary of numeric values (between 0 and 1) to log. The scores should give
	// you a variety of signals that help you determine how accurate the outputs are
	// compared to what you expect and diagnose failures. For example, a summarization
	// app might have one score that tells you how accurate the summary is, and another
	// that measures the word similarity between the generated and grouth truth
	// summary. The word similarity score could help you determine whether the
	// summarization was covering similar concepts or not. You can use these scores to
	// help you sort, filter, and compare logs.
	Scores map[string]float64 `json:"scores,omitzero"`
	// Human-identifying attributes of the span, such as name, type, etc.
	SpanAttributes SpanAttributesParam `json:"span_attributes,omitzero"`
	// Use `span_id`, `root_span_id`, and `span_parents` instead of `_parent_id`, which
	// is now deprecated. The span_id is a unique identifier describing the row's place
	// in the a trace, and the root_span_id is a unique identifier for the whole trace.
	// See the [guide](https://www.braintrust.dev/docs/guides/tracing) for full
	// details.
	//
	// For example, say we have logged a row
	// `{"id": "abc", "span_id": "span0", "root_span_id": "root_span0", "input": "foo", "output": "bar", "expected": "boo", "scores": {"correctness": 0.33}}`.
	// We can create a sub-span of the parent row by logging
	// `{"id": "llm_call", "span_id": "span1", "root_span_id": "root_span0", "span_parents": ["span0"], "input": {"prompt": "What comes after foo?"}, "output": "bar", "metrics": {"tokens": 1}}`.
	// In the webapp, only the root span row `"abc"` will show up in the summary view.
	// You can view the full trace hierarchy (in this case, the `"llm_call"` row) by
	// clicking on the "abc" row.
	//
	// If the row is being merged into an existing row, this field will be ignored.
	SpanParents []string `json:"span_parents,omitzero"`
	// A list of tags to log
	Tags []string `json:"tags,omitzero"`
	// The error that occurred, if any.
	Error interface{} `json:"error,omitzero"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not.
	// Braintrust currently does not compare `output` to `expected` for you, since
	// there are so many different ways to do that correctly. Instead, these values are
	// just used to help you navigate while digging into analyses. However, we may
	// later use these values to re-score outputs or fine-tune your models.
	Expected interface{} `json:"expected,omitzero"`
	// The arguments that uniquely define a user input (an arbitrary, JSON serializable
	// object).
	Input interface{} `json:"input,omitzero"`
	// The output of your application, including post-processing (an arbitrary, JSON
	// serializable object), that allows you to determine whether the result is correct
	// or not. For example, in an app that generates SQL queries, the `output` should
	// be the _result_ of the SQL query generated by the model, not the query itself,
	// because there may be multiple valid queries that answer a single question.
	Output interface{} `json:"output,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InsertProjectLogsEventParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r InsertProjectLogsEventParam) MarshalJSON() (data []byte, err error) {
	type shadow InsertProjectLogsEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Context is additional information about the code that produced the project logs
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the project
// logs event
type InsertProjectLogsEventContextParam struct {
	// Name of the file in code where the project logs event was created
	CallerFilename param.Opt[string] `json:"caller_filename,omitzero"`
	// The function in code which created the project logs event
	CallerFunctionname param.Opt[string] `json:"caller_functionname,omitzero"`
	// Line of code where the project logs event was created
	CallerLineno param.Opt[int64]       `json:"caller_lineno,omitzero"`
	ExtraFields  map[string]interface{} `json:"-,extras"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InsertProjectLogsEventContextParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InsertProjectLogsEventContextParam) MarshalJSON() (data []byte, err error) {
	type shadow InsertProjectLogsEventContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A dictionary with additional data about the test example, model outputs, or just
// about anything else that's relevant, that you can use to help find and analyze
// examples later. For example, you could log the `prompt`, example's `id`, or
// anything else that would be useful to slice/dice later. The values in `metadata`
// can be any JSON-serializable type, but its keys must be strings
type InsertProjectLogsEventMetadataParam struct {
	// The model used for this example
	Model       param.Opt[string]      `json:"model,omitzero"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InsertProjectLogsEventMetadataParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InsertProjectLogsEventMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow InsertProjectLogsEventMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the project logs event. Use "start" and "end" to track the time span
// over which the project logs event was produced
type InsertProjectLogsEventMetricsParam struct {
	// The number of tokens in the completion generated by the model (only set if this
	// is an LLM span)
	CompletionTokens param.Opt[int64] `json:"completion_tokens,omitzero"`
	// A unix timestamp recording when the section of code which produced the project
	// logs event finished
	End param.Opt[float64] `json:"end,omitzero"`
	// The number of tokens in the prompt used to generate the project logs event (only
	// set if this is an LLM span)
	PromptTokens param.Opt[int64] `json:"prompt_tokens,omitzero"`
	// A unix timestamp recording when the section of code which produced the project
	// logs event started
	Start param.Opt[float64] `json:"start,omitzero"`
	// The total number of tokens in the input and output of the project logs event.
	Tokens param.Opt[int64] `json:"tokens,omitzero"`
	// This metric is deprecated
	CallerFilename interface{} `json:"caller_filename,omitzero"`
	// This metric is deprecated
	CallerFunctionname interface{} `json:"caller_functionname,omitzero"`
	// This metric is deprecated
	CallerLineno interface{}        `json:"caller_lineno,omitzero"`
	ExtraFields  map[string]float64 `json:"-,extras"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InsertProjectLogsEventMetricsParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InsertProjectLogsEventMetricsParam) MarshalJSON() (data []byte, err error) {
	type shadow InsertProjectLogsEventMetricsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Summary of a metric's performance
type MetricSummary struct {
	// Number of improvements in the metric
	Improvements int64 `json:"improvements,required"`
	// Average metric across all examples
	Metric float64 `json:"metric,required"`
	// Name of the metric
	Name string `json:"name,required"`
	// Number of regressions in the metric
	Regressions int64 `json:"regressions,required"`
	// Unit label for the metric
	Unit string `json:"unit,required"`
	// Difference in metric between the current and comparison experiment
	Diff float64 `json:"diff"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Improvements resp.Field
		Metric       resp.Field
		Name         resp.Field
		Regressions  resp.Field
		Unit         resp.Field
		Diff         resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricSummary) RawJSON() string { return r.JSON.raw }
func (r *MetricSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the event was copied from another object.
type ObjectReference struct {
	// ID of the original event.
	ID string `json:"id,required"`
	// Transaction ID of the original event.
	XactID string `json:"_xact_id,required"`
	// ID of the object the event is originating from.
	ObjectID string `json:"object_id,required" format:"uuid"`
	// Type of the object the event is originating from.
	//
	// Any of "experiment", "dataset", "prompt", "function", "prompt_session",
	// "project_logs".
	ObjectType ObjectReferenceObjectType `json:"object_type,required"`
	// Created timestamp of the original event. Used to help sort in the UI
	Created string `json:"created,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		XactID      resp.Field
		ObjectID    resp.Field
		ObjectType  resp.Field
		Created     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectReference) RawJSON() string { return r.JSON.raw }
func (r *ObjectReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ObjectReference to a ObjectReferenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ObjectReferenceParam.IsOverridden()
func (r ObjectReference) ToParam() ObjectReferenceParam {
	return param.OverrideObj[ObjectReferenceParam](r.RawJSON())
}

// Type of the object the event is originating from.
type ObjectReferenceObjectType string

const (
	ObjectReferenceObjectTypeExperiment    ObjectReferenceObjectType = "experiment"
	ObjectReferenceObjectTypeDataset       ObjectReferenceObjectType = "dataset"
	ObjectReferenceObjectTypePrompt        ObjectReferenceObjectType = "prompt"
	ObjectReferenceObjectTypeFunction      ObjectReferenceObjectType = "function"
	ObjectReferenceObjectTypePromptSession ObjectReferenceObjectType = "prompt_session"
	ObjectReferenceObjectTypeProjectLogs   ObjectReferenceObjectType = "project_logs"
)

// Indicates the event was copied from another object.
//
// The properties ID, XactID, ObjectID, ObjectType are required.
type ObjectReferenceParam struct {
	// ID of the original event.
	ID string `json:"id,required"`
	// Transaction ID of the original event.
	XactID string `json:"_xact_id,required"`
	// ID of the object the event is originating from.
	ObjectID string `json:"object_id,required" format:"uuid"`
	// Type of the object the event is originating from.
	//
	// Any of "experiment", "dataset", "prompt", "function", "prompt_session",
	// "project_logs".
	ObjectType ObjectReferenceObjectType `json:"object_type,omitzero,required"`
	// Created timestamp of the original event. Used to help sort in the UI
	Created param.Opt[string] `json:"created,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ObjectReferenceParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ObjectReferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow ObjectReferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type OnlineScoreConfig struct {
	// The sampling rate for online scoring
	SamplingRate float64 `json:"sampling_rate,required"`
	// The list of scorers to use for online scoring
	Scorers []OnlineScoreConfigScorerUnion `json:"scorers,required"`
	// Whether to trigger online scoring on the root span of each trace
	ApplyToRootSpan bool `json:"apply_to_root_span,nullable"`
	// Trigger online scoring on any spans with a name in this list
	ApplyToSpanNames []string `json:"apply_to_span_names,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		SamplingRate     resp.Field
		Scorers          resp.Field
		ApplyToRootSpan  resp.Field
		ApplyToSpanNames resp.Field
		ExtraFields      map[string]resp.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnlineScoreConfig) RawJSON() string { return r.JSON.raw }
func (r *OnlineScoreConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OnlineScoreConfig to a OnlineScoreConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OnlineScoreConfigParam.IsOverridden()
func (r OnlineScoreConfig) ToParam() OnlineScoreConfigParam {
	return param.OverrideObj[OnlineScoreConfigParam](r.RawJSON())
}

// OnlineScoreConfigScorerUnion contains all possible properties and values from
// [OnlineScoreConfigScorerFunction], [OnlineScoreConfigScorerGlobal].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OnlineScoreConfigScorerUnion struct {
	// This field is from variant [OnlineScoreConfigScorerFunction].
	ID   string `json:"id"`
	Type string `json:"type"`
	// This field is from variant [OnlineScoreConfigScorerGlobal].
	Name string `json:"name"`
	JSON struct {
		ID   resp.Field
		Type resp.Field
		Name resp.Field
		raw  string
	} `json:"-"`
}

func (u OnlineScoreConfigScorerUnion) AsFunction() (v OnlineScoreConfigScorerFunction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OnlineScoreConfigScorerUnion) AsGlobal() (v OnlineScoreConfigScorerGlobal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OnlineScoreConfigScorerUnion) RawJSON() string { return u.JSON.raw }

func (r *OnlineScoreConfigScorerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnlineScoreConfigScorerFunction struct {
	ID string `json:"id,required"`
	// Any of "function".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnlineScoreConfigScorerFunction) RawJSON() string { return r.JSON.raw }
func (r *OnlineScoreConfigScorerFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnlineScoreConfigScorerGlobal struct {
	Name string `json:"name,required"`
	// Any of "global".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnlineScoreConfigScorerGlobal) RawJSON() string { return r.JSON.raw }
func (r *OnlineScoreConfigScorerGlobal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties SamplingRate, Scorers are required.
type OnlineScoreConfigParam struct {
	// The sampling rate for online scoring
	SamplingRate float64 `json:"sampling_rate,required"`
	// The list of scorers to use for online scoring
	Scorers []OnlineScoreConfigScorerUnionParam `json:"scorers,omitzero,required"`
	// Whether to trigger online scoring on the root span of each trace
	ApplyToRootSpan param.Opt[bool] `json:"apply_to_root_span,omitzero"`
	// Trigger online scoring on any spans with a name in this list
	ApplyToSpanNames []string `json:"apply_to_span_names,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f OnlineScoreConfigParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r OnlineScoreConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow OnlineScoreConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OnlineScoreConfigScorerUnionParam struct {
	OfFunction *OnlineScoreConfigScorerFunctionParam `json:",omitzero,inline"`
	OfGlobal   *OnlineScoreConfigScorerGlobalParam   `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u OnlineScoreConfigScorerUnionParam) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u OnlineScoreConfigScorerUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[OnlineScoreConfigScorerUnionParam](u.OfFunction, u.OfGlobal)
}

func (u *OnlineScoreConfigScorerUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFunction) {
		return u.OfFunction
	} else if !param.IsOmitted(u.OfGlobal) {
		return u.OfGlobal
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OnlineScoreConfigScorerUnionParam) GetID() *string {
	if vt := u.OfFunction; vt != nil {
		return &vt.ID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OnlineScoreConfigScorerUnionParam) GetName() *string {
	if vt := u.OfGlobal; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OnlineScoreConfigScorerUnionParam) GetType() *string {
	if vt := u.OfFunction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfGlobal; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The properties ID, Type are required.
type OnlineScoreConfigScorerFunctionParam struct {
	ID string `json:"id,required"`
	// Any of "function".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f OnlineScoreConfigScorerFunctionParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r OnlineScoreConfigScorerFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow OnlineScoreConfigScorerFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[OnlineScoreConfigScorerFunctionParam](
		"Type", false, "function",
	)
}

// The properties Name, Type are required.
type OnlineScoreConfigScorerGlobalParam struct {
	Name string `json:"name,required"`
	// Any of "global".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f OnlineScoreConfigScorerGlobalParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r OnlineScoreConfigScorerGlobalParam) MarshalJSON() (data []byte, err error) {
	type shadow OnlineScoreConfigScorerGlobalParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[OnlineScoreConfigScorerGlobalParam](
		"Type", false, "global",
	)
}

type Organization struct {
	// Unique identifier for the organization
	ID string `json:"id,required" format:"uuid"`
	// Name of the organization
	Name   string `json:"name,required"`
	APIURL string `json:"api_url,nullable"`
	// Date of organization creation
	Created        time.Time `json:"created,nullable" format:"date-time"`
	IsUniversalAPI bool      `json:"is_universal_api,nullable"`
	ProxyURL       string    `json:"proxy_url,nullable"`
	RealtimeURL    string    `json:"realtime_url,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID             resp.Field
		Name           resp.Field
		APIURL         resp.Field
		Created        resp.Field
		IsUniversalAPI resp.Field
		ProxyURL       resp.Field
		RealtimeURL    resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Organization) RawJSON() string { return r.JSON.raw }
func (r *Organization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PatchOrganizationMembersOutput struct {
	// The id of the org that was modified.
	OrgID string `json:"org_id,required"`
	// Any of "success".
	Status PatchOrganizationMembersOutputStatus `json:"status,required"`
	// If invite emails failed to send for some reason, the patch operation will still
	// complete, but we will return an error message here
	SendEmailError string `json:"send_email_error,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		OrgID          resp.Field
		Status         resp.Field
		SendEmailError resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PatchOrganizationMembersOutput) RawJSON() string { return r.JSON.raw }
func (r *PatchOrganizationMembersOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PatchOrganizationMembersOutputStatus string

const (
	PatchOrganizationMembersOutputStatusSuccess PatchOrganizationMembersOutputStatus = "success"
)

// Each permission permits a certain type of operation on an object in the system
//
// Permissions can be assigned to to objects on an individual basis, or grouped
// into roles
type Permission string

const (
	PermissionCreate     Permission = "create"
	PermissionRead       Permission = "read"
	PermissionUpdate     Permission = "update"
	PermissionDelete     Permission = "delete"
	PermissionCreateACLs Permission = "create_acls"
	PermissionReadACLs   Permission = "read_acls"
	PermissionUpdateACLs Permission = "update_acls"
	PermissionDeleteACLs Permission = "delete_acls"
)

type Project struct {
	// Unique identifier for the project
	ID string `json:"id,required" format:"uuid"`
	// Name of the project
	Name string `json:"name,required"`
	// Unique id for the organization that the project belongs under
	OrgID string `json:"org_id,required" format:"uuid"`
	// Date of project creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Date of project deletion, or null if the project is still active
	DeletedAt time.Time       `json:"deleted_at,nullable" format:"date-time"`
	Settings  ProjectSettings `json:"settings,nullable"`
	// Identifies the user who created the project
	UserID string `json:"user_id,nullable" format:"uuid"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Name        resp.Field
		OrgID       resp.Field
		Created     resp.Field
		DeletedAt   resp.Field
		Settings    resp.Field
		UserID      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Project) RawJSON() string { return r.JSON.raw }
func (r *Project) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectLogsEvent struct {
	// A unique identifier for the project logs event. If you don't provide one,
	// BrainTrust will generate one for you
	ID string `json:"id,required"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the project logs (see the
	// `version` parameter)
	XactID string `json:"_xact_id,required"`
	// The timestamp the project logs event was created
	Created time.Time `json:"created,required" format:"date-time"`
	// A literal 'g' which identifies the log as a project log
	//
	// Any of "g".
	LogID ProjectLogsEventLogID `json:"log_id,required"`
	// Unique id for the organization that the project belongs under
	OrgID string `json:"org_id,required" format:"uuid"`
	// Unique identifier for the project
	ProjectID string `json:"project_id,required" format:"uuid"`
	// A unique identifier for the trace this project logs event belongs to
	RootSpanID string `json:"root_span_id,required"`
	// A unique identifier used to link different project logs events together as part
	// of a full trace. See the
	// [tracing guide](https://www.braintrust.dev/docs/guides/tracing) for full details
	// on tracing
	SpanID string `json:"span_id,required"`
	// Context is additional information about the code that produced the project logs
	// event. It is essentially the textual counterpart to `metrics`. Use the
	// `caller_*` attributes to track the location in code which produced the project
	// logs event
	Context ProjectLogsEventContext `json:"context,nullable"`
	// The error that occurred, if any.
	Error interface{} `json:"error"`
	// The ground truth value (an arbitrary, JSON serializable object) that you'd
	// compare to `output` to determine if your `output` value is correct or not.
	// Braintrust currently does not compare `output` to `expected` for you, since
	// there are so many different ways to do that correctly. Instead, these values are
	// just used to help you navigate while digging into analyses. However, we may
	// later use these values to re-score outputs or fine-tune your models.
	Expected interface{} `json:"expected"`
	// The arguments that uniquely define a user input (an arbitrary, JSON serializable
	// object).
	Input interface{} `json:"input"`
	// Whether this span is a root span
	IsRoot bool `json:"is_root,nullable"`
	// A dictionary with additional data about the test example, model outputs, or just
	// about anything else that's relevant, that you can use to help find and analyze
	// examples later. For example, you could log the `prompt`, example's `id`, or
	// anything else that would be useful to slice/dice later. The values in `metadata`
	// can be any JSON-serializable type, but its keys must be strings
	Metadata ProjectLogsEventMetadata `json:"metadata,nullable"`
	// Metrics are numerical measurements tracking the execution of the code that
	// produced the project logs event. Use "start" and "end" to track the time span
	// over which the project logs event was produced
	Metrics ProjectLogsEventMetrics `json:"metrics,nullable"`
	// Indicates the event was copied from another object.
	Origin ObjectReference `json:"origin,nullable"`
	// The output of your application, including post-processing (an arbitrary, JSON
	// serializable object), that allows you to determine whether the result is correct
	// or not. For example, in an app that generates SQL queries, the `output` should
	// be the _result_ of the SQL query generated by the model, not the query itself,
	// because there may be multiple valid queries that answer a single question.
	Output interface{} `json:"output"`
	// A dictionary of numeric values (between 0 and 1) to log. The scores should give
	// you a variety of signals that help you determine how accurate the outputs are
	// compared to what you expect and diagnose failures. For example, a summarization
	// app might have one score that tells you how accurate the summary is, and another
	// that measures the word similarity between the generated and grouth truth
	// summary. The word similarity score could help you determine whether the
	// summarization was covering similar concepts or not. You can use these scores to
	// help you sort, filter, and compare logs.
	Scores map[string]float64 `json:"scores,nullable"`
	// Human-identifying attributes of the span, such as name, type, etc.
	SpanAttributes SpanAttributes `json:"span_attributes,nullable"`
	// An array of the parent `span_ids` of this project logs event. This should be
	// empty for the root span of a trace, and should most often contain just one
	// parent element for subspans
	SpanParents []string `json:"span_parents,nullable"`
	// A list of tags to log
	Tags []string `json:"tags,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID             resp.Field
		XactID         resp.Field
		Created        resp.Field
		LogID          resp.Field
		OrgID          resp.Field
		ProjectID      resp.Field
		RootSpanID     resp.Field
		SpanID         resp.Field
		Context        resp.Field
		Error          resp.Field
		Expected       resp.Field
		Input          resp.Field
		IsRoot         resp.Field
		Metadata       resp.Field
		Metrics        resp.Field
		Origin         resp.Field
		Output         resp.Field
		Scores         resp.Field
		SpanAttributes resp.Field
		SpanParents    resp.Field
		Tags           resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectLogsEvent) RawJSON() string { return r.JSON.raw }
func (r *ProjectLogsEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A literal 'g' which identifies the log as a project log
type ProjectLogsEventLogID string

const (
	ProjectLogsEventLogIDG ProjectLogsEventLogID = "g"
)

// Context is additional information about the code that produced the project logs
// event. It is essentially the textual counterpart to `metrics`. Use the
// `caller_*` attributes to track the location in code which produced the project
// logs event
type ProjectLogsEventContext struct {
	// Name of the file in code where the project logs event was created
	CallerFilename string `json:"caller_filename,nullable"`
	// The function in code which created the project logs event
	CallerFunctionname string `json:"caller_functionname,nullable"`
	// Line of code where the project logs event was created
	CallerLineno int64                  `json:"caller_lineno,nullable"`
	ExtraFields  map[string]interface{} `json:",extras"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CallerFilename     resp.Field
		CallerFunctionname resp.Field
		CallerLineno       resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectLogsEventContext) RawJSON() string { return r.JSON.raw }
func (r *ProjectLogsEventContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dictionary with additional data about the test example, model outputs, or just
// about anything else that's relevant, that you can use to help find and analyze
// examples later. For example, you could log the `prompt`, example's `id`, or
// anything else that would be useful to slice/dice later. The values in `metadata`
// can be any JSON-serializable type, but its keys must be strings
type ProjectLogsEventMetadata struct {
	// The model used for this example
	Model       string                 `json:"model,nullable"`
	ExtraFields map[string]interface{} `json:",extras"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Model       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectLogsEventMetadata) RawJSON() string { return r.JSON.raw }
func (r *ProjectLogsEventMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metrics are numerical measurements tracking the execution of the code that
// produced the project logs event. Use "start" and "end" to track the time span
// over which the project logs event was produced
type ProjectLogsEventMetrics struct {
	// This metric is deprecated
	CallerFilename interface{} `json:"caller_filename"`
	// This metric is deprecated
	CallerFunctionname interface{} `json:"caller_functionname"`
	// This metric is deprecated
	CallerLineno interface{} `json:"caller_lineno"`
	// The number of tokens in the completion generated by the model (only set if this
	// is an LLM span)
	CompletionTokens int64 `json:"completion_tokens,nullable"`
	// A unix timestamp recording when the section of code which produced the project
	// logs event finished
	End float64 `json:"end,nullable"`
	// The number of tokens in the prompt used to generate the project logs event (only
	// set if this is an LLM span)
	PromptTokens int64 `json:"prompt_tokens,nullable"`
	// A unix timestamp recording when the section of code which produced the project
	// logs event started
	Start float64 `json:"start,nullable"`
	// The total number of tokens in the input and output of the project logs event.
	Tokens      int64              `json:"tokens,nullable"`
	ExtraFields map[string]float64 `json:",extras"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CallerFilename     resp.Field
		CallerFunctionname resp.Field
		CallerLineno       resp.Field
		CompletionTokens   resp.Field
		End                resp.Field
		PromptTokens       resp.Field
		Start              resp.Field
		Tokens             resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectLogsEventMetrics) RawJSON() string { return r.JSON.raw }
func (r *ProjectLogsEventMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A project score is a user-configured score, which can be manually-labeled
// through the UI
type ProjectScore struct {
	// Unique identifier for the project score
	ID string `json:"id,required" format:"uuid"`
	// Name of the project score
	Name string `json:"name,required"`
	// Unique identifier for the project that the project score belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// The type of the configured score
	//
	// Any of "slider", "categorical", "weighted", "minimum", "maximum", "online",
	// "free-form".
	ScoreType ProjectScoreType `json:"score_type,required"`
	UserID    string           `json:"user_id,required" format:"uuid"`
	// For categorical-type project scores, the list of all categories
	Categories ProjectScoreCategoriesUnion `json:"categories,nullable"`
	Config     ProjectScoreConfig          `json:"config,nullable"`
	// Date of project score creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Textual description of the project score
	Description string `json:"description,nullable"`
	// An optional LexoRank-based string that sets the sort position for the score in
	// the UI
	Position string `json:"position,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Name        resp.Field
		ProjectID   resp.Field
		ScoreType   resp.Field
		UserID      resp.Field
		Categories  resp.Field
		Config      resp.Field
		Created     resp.Field
		Description resp.Field
		Position    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectScore) RawJSON() string { return r.JSON.raw }
func (r *ProjectScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProjectScoreCategoriesUnion contains all possible properties and values from
// [[]ProjectScoreCategory], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfCategorical OfMinimum]
type ProjectScoreCategoriesUnion struct {
	// This field will be present if the value is a [[]ProjectScoreCategory] instead of
	// an object.
	OfCategorical []ProjectScoreCategory `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfMinimum []string `json:",inline"`
	JSON      struct {
		OfCategorical resp.Field
		OfMinimum     resp.Field
		raw           string
	} `json:"-"`
}

func (u ProjectScoreCategoriesUnion) AsCategorical() (v []ProjectScoreCategory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProjectScoreCategoriesUnion) AsMinimum() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProjectScoreCategoriesUnion) RawJSON() string { return u.JSON.raw }

func (r *ProjectScoreCategoriesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// For categorical-type project scores, defines a single category
type ProjectScoreCategory struct {
	// Name of the category
	Name string `json:"name,required"`
	// Numerical value of the category. Must be between 0 and 1, inclusive
	Value float64 `json:"value,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		Value       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectScoreCategory) RawJSON() string { return r.JSON.raw }
func (r *ProjectScoreCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ProjectScoreCategory to a ProjectScoreCategoryParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ProjectScoreCategoryParam.IsOverridden()
func (r ProjectScoreCategory) ToParam() ProjectScoreCategoryParam {
	return param.OverrideObj[ProjectScoreCategoryParam](r.RawJSON())
}

// For categorical-type project scores, defines a single category
//
// The properties Name, Value are required.
type ProjectScoreCategoryParam struct {
	// Name of the category
	Name string `json:"name,required"`
	// Numerical value of the category. Must be between 0 and 1, inclusive
	Value float64 `json:"value,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectScoreCategoryParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ProjectScoreCategoryParam) MarshalJSON() (data []byte, err error) {
	type shadow ProjectScoreCategoryParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ProjectScoreConfig struct {
	Destination string            `json:"destination,nullable"`
	MultiSelect bool              `json:"multi_select,nullable"`
	Online      OnlineScoreConfig `json:"online,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Destination resp.Field
		MultiSelect resp.Field
		Online      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectScoreConfig) RawJSON() string { return r.JSON.raw }
func (r *ProjectScoreConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ProjectScoreConfig to a ProjectScoreConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ProjectScoreConfigParam.IsOverridden()
func (r ProjectScoreConfig) ToParam() ProjectScoreConfigParam {
	return param.OverrideObj[ProjectScoreConfigParam](r.RawJSON())
}

type ProjectScoreConfigParam struct {
	Destination param.Opt[string]      `json:"destination,omitzero"`
	MultiSelect param.Opt[bool]        `json:"multi_select,omitzero"`
	Online      OnlineScoreConfigParam `json:"online,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectScoreConfigParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ProjectScoreConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow ProjectScoreConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The type of the configured score
type ProjectScoreType string

const (
	ProjectScoreTypeSlider      ProjectScoreType = "slider"
	ProjectScoreTypeCategorical ProjectScoreType = "categorical"
	ProjectScoreTypeWeighted    ProjectScoreType = "weighted"
	ProjectScoreTypeMinimum     ProjectScoreType = "minimum"
	ProjectScoreTypeMaximum     ProjectScoreType = "maximum"
	ProjectScoreTypeOnline      ProjectScoreType = "online"
	ProjectScoreTypeFreeForm    ProjectScoreType = "free-form"
)

type ProjectSettings struct {
	// The id of the experiment to use as the default baseline for comparisons
	BaselineExperimentID string `json:"baseline_experiment_id,nullable" format:"uuid"`
	// The key used to join two experiments (defaults to `input`)
	ComparisonKey string `json:"comparison_key,nullable"`
	// The order of the fields to display in the trace view
	SpanFieldOrder []ProjectSettingsSpanFieldOrder `json:"spanFieldOrder,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		BaselineExperimentID resp.Field
		ComparisonKey        resp.Field
		SpanFieldOrder       resp.Field
		ExtraFields          map[string]resp.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectSettings) RawJSON() string { return r.JSON.raw }
func (r *ProjectSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ProjectSettings to a ProjectSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ProjectSettingsParam.IsOverridden()
func (r ProjectSettings) ToParam() ProjectSettingsParam {
	return param.OverrideObj[ProjectSettingsParam](r.RawJSON())
}

type ProjectSettingsSpanFieldOrder struct {
	ColumnID   string `json:"column_id,required"`
	ObjectType string `json:"object_type,required"`
	Position   string `json:"position,required"`
	// Any of "full", "two_column".
	Layout ProjectSettingsSpanFieldOrderLayout `json:"layout,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ColumnID    resp.Field
		ObjectType  resp.Field
		Position    resp.Field
		Layout      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectSettingsSpanFieldOrder) RawJSON() string { return r.JSON.raw }
func (r *ProjectSettingsSpanFieldOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectSettingsSpanFieldOrderLayout string

const (
	ProjectSettingsSpanFieldOrderLayoutFull      ProjectSettingsSpanFieldOrderLayout = "full"
	ProjectSettingsSpanFieldOrderLayoutTwoColumn ProjectSettingsSpanFieldOrderLayout = "two_column"
)

type ProjectSettingsParam struct {
	// The id of the experiment to use as the default baseline for comparisons
	BaselineExperimentID param.Opt[string] `json:"baseline_experiment_id,omitzero" format:"uuid"`
	// The key used to join two experiments (defaults to `input`)
	ComparisonKey param.Opt[string] `json:"comparison_key,omitzero"`
	// The order of the fields to display in the trace view
	SpanFieldOrder []ProjectSettingsSpanFieldOrderParam `json:"spanFieldOrder,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectSettingsParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ProjectSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow ProjectSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties ColumnID, ObjectType, Position are required.
type ProjectSettingsSpanFieldOrderParam struct {
	ColumnID   string `json:"column_id,required"`
	ObjectType string `json:"object_type,required"`
	Position   string `json:"position,required"`
	// Any of "full", "two_column".
	Layout ProjectSettingsSpanFieldOrderLayout `json:"layout,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectSettingsSpanFieldOrderParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ProjectSettingsSpanFieldOrderParam) MarshalJSON() (data []byte, err error) {
	type shadow ProjectSettingsSpanFieldOrderParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A project tag is a user-configured tag for tracking and filtering your
// experiments, logs, and other data
type ProjectTag struct {
	// Unique identifier for the project tag
	ID string `json:"id,required" format:"uuid"`
	// Name of the project tag
	Name string `json:"name,required"`
	// Unique identifier for the project that the project tag belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	UserID    string `json:"user_id,required" format:"uuid"`
	// Color of the tag for the UI
	Color string `json:"color,nullable"`
	// Date of project tag creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Textual description of the project tag
	Description string `json:"description,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Name        resp.Field
		ProjectID   resp.Field
		UserID      resp.Field
		Color       resp.Field
		Created     resp.Field
		Description resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectTag) RawJSON() string { return r.JSON.raw }
func (r *ProjectTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Prompt struct {
	// Unique identifier for the prompt
	ID string `json:"id,required" format:"uuid"`
	// The transaction id of an event is unique to the network operation that processed
	// the event insertion. Transaction ids are monotonically increasing over time and
	// can be used to retrieve a versioned snapshot of the prompt (see the `version`
	// parameter)
	XactID string `json:"_xact_id,required"`
	// A literal 'p' which identifies the object as a project prompt
	//
	// Any of "p".
	LogID PromptLogID `json:"log_id,required"`
	// Name of the prompt
	Name string `json:"name,required"`
	// Unique identifier for the organization
	OrgID string `json:"org_id,required" format:"uuid"`
	// Unique identifier for the project that the prompt belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// Unique identifier for the prompt
	Slug string `json:"slug,required"`
	// Date of prompt creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Textual description of the prompt
	Description string `json:"description,nullable"`
	// Any of "llm", "scorer", "task", "tool".
	FunctionType PromptFunctionType `json:"function_type,nullable"`
	// User-controlled metadata about the prompt
	Metadata map[string]interface{} `json:"metadata,nullable"`
	// The prompt, model, and its parameters
	PromptData PromptData `json:"prompt_data,nullable"`
	// A list of tags for the prompt
	Tags []string `json:"tags,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID           resp.Field
		XactID       resp.Field
		LogID        resp.Field
		Name         resp.Field
		OrgID        resp.Field
		ProjectID    resp.Field
		Slug         resp.Field
		Created      resp.Field
		Description  resp.Field
		FunctionType resp.Field
		Metadata     resp.Field
		PromptData   resp.Field
		Tags         resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Prompt) RawJSON() string { return r.JSON.raw }
func (r *Prompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A literal 'p' which identifies the object as a project prompt
type PromptLogID string

const (
	PromptLogIDP PromptLogID = "p"
)

type PromptFunctionType string

const (
	PromptFunctionTypeLlm    PromptFunctionType = "llm"
	PromptFunctionTypeScorer PromptFunctionType = "scorer"
	PromptFunctionTypeTask   PromptFunctionType = "task"
	PromptFunctionTypeTool   PromptFunctionType = "tool"
)

// The prompt, model, and its parameters
type PromptData struct {
	Options       PromptOptions                 `json:"options,nullable"`
	Origin        PromptDataOrigin              `json:"origin,nullable"`
	Parser        PromptDataParser              `json:"parser,nullable"`
	Prompt        PromptDataPromptUnion         `json:"prompt,nullable"`
	ToolFunctions []PromptDataToolFunctionUnion `json:"tool_functions,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Options       resp.Field
		Origin        resp.Field
		Parser        resp.Field
		Prompt        resp.Field
		ToolFunctions resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptData) RawJSON() string { return r.JSON.raw }
func (r *PromptData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PromptData to a PromptDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PromptDataParam.IsOverridden()
func (r PromptData) ToParam() PromptDataParam {
	return param.OverrideObj[PromptDataParam](r.RawJSON())
}

type PromptDataOrigin struct {
	ProjectID     string `json:"project_id"`
	PromptID      string `json:"prompt_id"`
	PromptVersion string `json:"prompt_version"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ProjectID     resp.Field
		PromptID      resp.Field
		PromptVersion resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptDataOrigin) RawJSON() string { return r.JSON.raw }
func (r *PromptDataOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptDataParser struct {
	ChoiceScores map[string]float64 `json:"choice_scores,required"`
	// Any of "llm_classifier".
	Type   string `json:"type,required"`
	UseCot bool   `json:"use_cot,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ChoiceScores resp.Field
		Type         resp.Field
		UseCot       resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptDataParser) RawJSON() string { return r.JSON.raw }
func (r *PromptDataParser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PromptDataPromptUnion contains all possible properties and values from
// [PromptDataPromptCompletion], [PromptDataPromptChat].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PromptDataPromptUnion struct {
	// This field is from variant [PromptDataPromptCompletion].
	Content string `json:"content"`
	Type    string `json:"type"`
	// This field is from variant [PromptDataPromptChat].
	Messages []PromptDataPromptChatMessageUnion `json:"messages"`
	// This field is from variant [PromptDataPromptChat].
	Tools string `json:"tools"`
	JSON  struct {
		Content  resp.Field
		Type     resp.Field
		Messages resp.Field
		Tools    resp.Field
		raw      string
	} `json:"-"`
}

func (u PromptDataPromptUnion) AsCompletion() (v PromptDataPromptCompletion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptDataPromptUnion) AsChat() (v PromptDataPromptChat) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PromptDataPromptUnion) RawJSON() string { return u.JSON.raw }

func (r *PromptDataPromptUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptDataPromptCompletion struct {
	Content string `json:"content,required"`
	// Any of "completion".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Content     resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptDataPromptCompletion) RawJSON() string { return r.JSON.raw }
func (r *PromptDataPromptCompletion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptDataPromptChat struct {
	Messages []PromptDataPromptChatMessageUnion `json:"messages,required"`
	// Any of "chat".
	Type  string `json:"type,required"`
	Tools string `json:"tools"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Messages    resp.Field
		Type        resp.Field
		Tools       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptDataPromptChat) RawJSON() string { return r.JSON.raw }
func (r *PromptDataPromptChat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PromptDataPromptChatMessageUnion contains all possible properties and values
// from [PromptDataPromptChatMessageSystem], [PromptDataPromptChatMessageUser],
// [PromptDataPromptChatMessageAssistant], [PromptDataPromptChatMessageTool],
// [PromptDataPromptChatMessageFunction], [PromptDataPromptChatMessageFallback].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PromptDataPromptChatMessageUnion struct {
	Role string `json:"role"`
	// This field is a union of [string],
	// [PromptDataPromptChatMessageUserContentUnion], [string], [string], [string],
	// [string]
	Content PromptDataPromptChatMessageUnionContent `json:"content"`
	Name    string                                  `json:"name"`
	// This field is from variant [PromptDataPromptChatMessageAssistant].
	FunctionCall PromptDataPromptChatMessageAssistantFunctionCall `json:"function_call"`
	// This field is from variant [PromptDataPromptChatMessageAssistant].
	ToolCalls []ChatCompletionMessageToolCall `json:"tool_calls"`
	// This field is from variant [PromptDataPromptChatMessageTool].
	ToolCallID string `json:"tool_call_id"`
	JSON       struct {
		Role         resp.Field
		Content      resp.Field
		Name         resp.Field
		FunctionCall resp.Field
		ToolCalls    resp.Field
		ToolCallID   resp.Field
		raw          string
	} `json:"-"`
}

func (u PromptDataPromptChatMessageUnion) AsSystem() (v PromptDataPromptChatMessageSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptDataPromptChatMessageUnion) AsUser() (v PromptDataPromptChatMessageUser) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptDataPromptChatMessageUnion) AsAssistant() (v PromptDataPromptChatMessageAssistant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptDataPromptChatMessageUnion) AsTool() (v PromptDataPromptChatMessageTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptDataPromptChatMessageUnion) AsFunction() (v PromptDataPromptChatMessageFunction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptDataPromptChatMessageUnion) AsFallback() (v PromptDataPromptChatMessageFallback) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PromptDataPromptChatMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *PromptDataPromptChatMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PromptDataPromptChatMessageUnionContent is an implicit subunion of
// [PromptDataPromptChatMessageUnion]. PromptDataPromptChatMessageUnionContent
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PromptDataPromptChatMessageUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfArray]
type PromptDataPromptChatMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]PromptDataPromptChatMessageUserContentArrayItemUnion] instead of an object.
	OfArray []PromptDataPromptChatMessageUserContentArrayItemUnion `json:",inline"`
	JSON    struct {
		OfString resp.Field
		OfArray  resp.Field
		raw      string
	} `json:"-"`
}

func (r *PromptDataPromptChatMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptDataPromptChatMessageSystem struct {
	// Any of "system".
	Role    string `json:"role,required"`
	Content string `json:"content"`
	Name    string `json:"name"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Role        resp.Field
		Content     resp.Field
		Name        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptDataPromptChatMessageSystem) RawJSON() string { return r.JSON.raw }
func (r *PromptDataPromptChatMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptDataPromptChatMessageUser struct {
	// Any of "user".
	Role    string                                      `json:"role,required"`
	Content PromptDataPromptChatMessageUserContentUnion `json:"content"`
	Name    string                                      `json:"name"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Role        resp.Field
		Content     resp.Field
		Name        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptDataPromptChatMessageUser) RawJSON() string { return r.JSON.raw }
func (r *PromptDataPromptChatMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PromptDataPromptChatMessageUserContentUnion contains all possible properties and
// values from [string], [[]PromptDataPromptChatMessageUserContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfArray]
type PromptDataPromptChatMessageUserContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]PromptDataPromptChatMessageUserContentArrayItemUnion] instead of an object.
	OfArray []PromptDataPromptChatMessageUserContentArrayItemUnion `json:",inline"`
	JSON    struct {
		OfString resp.Field
		OfArray  resp.Field
		raw      string
	} `json:"-"`
}

func (u PromptDataPromptChatMessageUserContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptDataPromptChatMessageUserContentUnion) AsArray() (v []PromptDataPromptChatMessageUserContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PromptDataPromptChatMessageUserContentUnion) RawJSON() string { return u.JSON.raw }

func (r *PromptDataPromptChatMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PromptDataPromptChatMessageUserContentArrayItemUnion contains all possible
// properties and values from [ChatCompletionContentPartText],
// [ChatCompletionContentPartImage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PromptDataPromptChatMessageUserContentArrayItemUnion struct {
	Type string `json:"type"`
	// This field is from variant [ChatCompletionContentPartText].
	Text string `json:"text"`
	// This field is from variant [ChatCompletionContentPartImage].
	ImageURL ChatCompletionContentPartImageImageURL `json:"image_url"`
	JSON     struct {
		Type     resp.Field
		Text     resp.Field
		ImageURL resp.Field
		raw      string
	} `json:"-"`
}

func (u PromptDataPromptChatMessageUserContentArrayItemUnion) AsText() (v ChatCompletionContentPartText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptDataPromptChatMessageUserContentArrayItemUnion) AsImageURL() (v ChatCompletionContentPartImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PromptDataPromptChatMessageUserContentArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *PromptDataPromptChatMessageUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptDataPromptChatMessageAssistant struct {
	// Any of "assistant".
	Role         string                                           `json:"role,required"`
	Content      string                                           `json:"content,nullable"`
	FunctionCall PromptDataPromptChatMessageAssistantFunctionCall `json:"function_call,nullable"`
	Name         string                                           `json:"name,nullable"`
	ToolCalls    []ChatCompletionMessageToolCall                  `json:"tool_calls,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Role         resp.Field
		Content      resp.Field
		FunctionCall resp.Field
		Name         resp.Field
		ToolCalls    resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptDataPromptChatMessageAssistant) RawJSON() string { return r.JSON.raw }
func (r *PromptDataPromptChatMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptDataPromptChatMessageAssistantFunctionCall struct {
	Arguments string `json:"arguments,required"`
	Name      string `json:"name,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Arguments   resp.Field
		Name        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptDataPromptChatMessageAssistantFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *PromptDataPromptChatMessageAssistantFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptDataPromptChatMessageTool struct {
	// Any of "tool".
	Role       string `json:"role,required"`
	Content    string `json:"content"`
	ToolCallID string `json:"tool_call_id"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Role        resp.Field
		Content     resp.Field
		ToolCallID  resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptDataPromptChatMessageTool) RawJSON() string { return r.JSON.raw }
func (r *PromptDataPromptChatMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptDataPromptChatMessageFunction struct {
	Name string `json:"name,required"`
	// Any of "function".
	Role    string `json:"role,required"`
	Content string `json:"content"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		Role        resp.Field
		Content     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptDataPromptChatMessageFunction) RawJSON() string { return r.JSON.raw }
func (r *PromptDataPromptChatMessageFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptDataPromptChatMessageFallback struct {
	// Any of "model".
	Role    string `json:"role,required"`
	Content string `json:"content,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Role        resp.Field
		Content     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptDataPromptChatMessageFallback) RawJSON() string { return r.JSON.raw }
func (r *PromptDataPromptChatMessageFallback) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PromptDataToolFunctionUnion contains all possible properties and values from
// [PromptDataToolFunctionFunction], [PromptDataToolFunctionGlobal].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PromptDataToolFunctionUnion struct {
	// This field is from variant [PromptDataToolFunctionFunction].
	ID   string `json:"id"`
	Type string `json:"type"`
	// This field is from variant [PromptDataToolFunctionGlobal].
	Name string `json:"name"`
	JSON struct {
		ID   resp.Field
		Type resp.Field
		Name resp.Field
		raw  string
	} `json:"-"`
}

func (u PromptDataToolFunctionUnion) AsFunction() (v PromptDataToolFunctionFunction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptDataToolFunctionUnion) AsGlobal() (v PromptDataToolFunctionGlobal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PromptDataToolFunctionUnion) RawJSON() string { return u.JSON.raw }

func (r *PromptDataToolFunctionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptDataToolFunctionFunction struct {
	ID string `json:"id,required"`
	// Any of "function".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptDataToolFunctionFunction) RawJSON() string { return r.JSON.raw }
func (r *PromptDataToolFunctionFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptDataToolFunctionGlobal struct {
	Name string `json:"name,required"`
	// Any of "global".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptDataToolFunctionGlobal) RawJSON() string { return r.JSON.raw }
func (r *PromptDataToolFunctionGlobal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The prompt, model, and its parameters
type PromptDataParam struct {
	Options       PromptOptionsParam                 `json:"options,omitzero"`
	Origin        PromptDataOriginParam              `json:"origin,omitzero"`
	Parser        PromptDataParserParam              `json:"parser,omitzero"`
	Prompt        PromptDataPromptUnionParam         `json:"prompt,omitzero"`
	ToolFunctions []PromptDataToolFunctionUnionParam `json:"tool_functions,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptDataParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r PromptDataParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type PromptDataOriginParam struct {
	ProjectID     param.Opt[string] `json:"project_id,omitzero"`
	PromptID      param.Opt[string] `json:"prompt_id,omitzero"`
	PromptVersion param.Opt[string] `json:"prompt_version,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptDataOriginParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r PromptDataOriginParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptDataOriginParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties ChoiceScores, Type, UseCot are required.
type PromptDataParserParam struct {
	ChoiceScores map[string]float64 `json:"choice_scores,omitzero,required"`
	// Any of "llm_classifier".
	Type   string `json:"type,omitzero,required"`
	UseCot bool   `json:"use_cot,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptDataParserParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r PromptDataParserParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptDataParserParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptDataParserParam](
		"Type", false, "llm_classifier",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PromptDataPromptUnionParam struct {
	OfCompletion *PromptDataPromptCompletionParam `json:",omitzero,inline"`
	OfChat       *PromptDataPromptChatParam       `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u PromptDataPromptUnionParam) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u PromptDataPromptUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[PromptDataPromptUnionParam](u.OfCompletion, u.OfChat)
}

func (u *PromptDataPromptUnionParam) asAny() any {
	if !param.IsOmitted(u.OfCompletion) {
		return u.OfCompletion
	} else if !param.IsOmitted(u.OfChat) {
		return u.OfChat
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptDataPromptUnionParam) GetContent() *string {
	if vt := u.OfCompletion; vt != nil {
		return &vt.Content
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptDataPromptUnionParam) GetMessages() []PromptDataPromptChatMessageUnionParam {
	if vt := u.OfChat; vt != nil {
		return vt.Messages
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptDataPromptUnionParam) GetTools() *string {
	if vt := u.OfChat; vt != nil && vt.Tools.IsPresent() {
		return &vt.Tools.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptDataPromptUnionParam) GetType() *string {
	if vt := u.OfCompletion; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfChat; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The properties Content, Type are required.
type PromptDataPromptCompletionParam struct {
	Content string `json:"content,required"`
	// Any of "completion".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptDataPromptCompletionParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r PromptDataPromptCompletionParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptDataPromptCompletionParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptDataPromptCompletionParam](
		"Type", false, "completion",
	)
}

// The properties Messages, Type are required.
type PromptDataPromptChatParam struct {
	Messages []PromptDataPromptChatMessageUnionParam `json:"messages,omitzero,required"`
	// Any of "chat".
	Type  string            `json:"type,omitzero,required"`
	Tools param.Opt[string] `json:"tools,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptDataPromptChatParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r PromptDataPromptChatParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptDataPromptChatParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptDataPromptChatParam](
		"Type", false, "chat",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PromptDataPromptChatMessageUnionParam struct {
	OfSystem    *PromptDataPromptChatMessageSystemParam    `json:",omitzero,inline"`
	OfUser      *PromptDataPromptChatMessageUserParam      `json:",omitzero,inline"`
	OfAssistant *PromptDataPromptChatMessageAssistantParam `json:",omitzero,inline"`
	OfTool      *PromptDataPromptChatMessageToolParam      `json:",omitzero,inline"`
	OfFunction  *PromptDataPromptChatMessageFunctionParam  `json:",omitzero,inline"`
	OfFallback  *PromptDataPromptChatMessageFallbackParam  `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u PromptDataPromptChatMessageUnionParam) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u PromptDataPromptChatMessageUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[PromptDataPromptChatMessageUnionParam](u.OfSystem,
		u.OfUser,
		u.OfAssistant,
		u.OfTool,
		u.OfFunction,
		u.OfFallback)
}

func (u *PromptDataPromptChatMessageUnionParam) asAny() any {
	if !param.IsOmitted(u.OfSystem) {
		return u.OfSystem
	} else if !param.IsOmitted(u.OfUser) {
		return u.OfUser
	} else if !param.IsOmitted(u.OfAssistant) {
		return u.OfAssistant
	} else if !param.IsOmitted(u.OfTool) {
		return u.OfTool
	} else if !param.IsOmitted(u.OfFunction) {
		return u.OfFunction
	} else if !param.IsOmitted(u.OfFallback) {
		return u.OfFallback
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptDataPromptChatMessageUnionParam) GetFunctionCall() *PromptDataPromptChatMessageAssistantFunctionCallParam {
	if vt := u.OfAssistant; vt != nil {
		return &vt.FunctionCall
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptDataPromptChatMessageUnionParam) GetToolCalls() []ChatCompletionMessageToolCallParam {
	if vt := u.OfAssistant; vt != nil {
		return vt.ToolCalls
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptDataPromptChatMessageUnionParam) GetToolCallID() *string {
	if vt := u.OfTool; vt != nil && vt.ToolCallID.IsPresent() {
		return &vt.ToolCallID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptDataPromptChatMessageUnionParam) GetRole() *string {
	if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfUser; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfAssistant; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfTool; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfFunction; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfFallback; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptDataPromptChatMessageUnionParam) GetName() *string {
	if vt := u.OfSystem; vt != nil && vt.Name.IsPresent() {
		return &vt.Name.Value
	} else if vt := u.OfUser; vt != nil && vt.Name.IsPresent() {
		return &vt.Name.Value
	} else if vt := u.OfAssistant; vt != nil && vt.Name.IsPresent() {
		return &vt.Name.Value
	} else if vt := u.OfFunction; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u PromptDataPromptChatMessageUnionParam) GetContent() (res promptDataPromptChatMessageUnionParamContent) {
	if vt := u.OfSystem; vt != nil && vt.Content.IsPresent() {
		res.ofString = &vt.Content.Value
	} else if vt := u.OfUser; vt != nil {
		res.ofPromptDataPromptChatMessageUserContentUnion = &vt.Content
	} else if vt := u.OfAssistant; vt != nil && vt.Content.IsPresent() {
		res.ofString = &vt.Content.Value
	} else if vt := u.OfTool; vt != nil && vt.Content.IsPresent() {
		res.ofString = &vt.Content.Value
	} else if vt := u.OfFunction; vt != nil && vt.Content.IsPresent() {
		res.ofString = &vt.Content.Value
	} else if vt := u.OfFallback; vt != nil && vt.Content.IsPresent() {
		res.ofString = &vt.Content.Value
	}
	return
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type promptDataPromptChatMessageUnionParamContent struct {
	ofString                                      *string
	ofPromptDataPromptChatMessageUserContentUnion *PromptDataPromptChatMessageUserContentUnionParam
}

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]shared.PromptDataPromptChatMessageUserContentArrayItemUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u promptDataPromptChatMessageUnionParamContent) AsAny() any {
	if !param.IsOmitted(u.ofString) {
		return u.ofString
	} else if !param.IsOmitted(u.ofPromptDataPromptChatMessageUserContentUnion) {
		return u.ofPromptDataPromptChatMessageUserContentUnion.asAny()
	} else if !param.IsOmitted(u.ofString) {
		return u.ofString
	} else if !param.IsOmitted(u.ofString) {
		return u.ofString
	} else if !param.IsOmitted(u.ofString) {
		return u.ofString
	} else if !param.IsOmitted(u.ofString) {
		return u.ofString
	}
	return nil
}

// The property Role is required.
type PromptDataPromptChatMessageSystemParam struct {
	// Any of "system".
	Role    string            `json:"role,omitzero,required"`
	Content param.Opt[string] `json:"content,omitzero"`
	Name    param.Opt[string] `json:"name,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptDataPromptChatMessageSystemParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptDataPromptChatMessageSystemParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptDataPromptChatMessageSystemParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptDataPromptChatMessageSystemParam](
		"Role", false, "system",
	)
}

// The property Role is required.
type PromptDataPromptChatMessageUserParam struct {
	// Any of "user".
	Role    string                                           `json:"role,omitzero,required"`
	Name    param.Opt[string]                                `json:"name,omitzero"`
	Content PromptDataPromptChatMessageUserContentUnionParam `json:"content,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptDataPromptChatMessageUserParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptDataPromptChatMessageUserParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptDataPromptChatMessageUserParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptDataPromptChatMessageUserParam](
		"Role", false, "user",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PromptDataPromptChatMessageUserContentUnionParam struct {
	OfString param.Opt[string]                                           `json:",omitzero,inline"`
	OfArray  []PromptDataPromptChatMessageUserContentArrayItemUnionParam `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u PromptDataPromptChatMessageUserContentUnionParam) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u PromptDataPromptChatMessageUserContentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[PromptDataPromptChatMessageUserContentUnionParam](u.OfString, u.OfArray)
}

func (u *PromptDataPromptChatMessageUserContentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfArray) {
		return &u.OfArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PromptDataPromptChatMessageUserContentArrayItemUnionParam struct {
	OfText     *ChatCompletionContentPartTextParam  `json:",omitzero,inline"`
	OfImageURL *ChatCompletionContentPartImageParam `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u PromptDataPromptChatMessageUserContentArrayItemUnionParam) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u PromptDataPromptChatMessageUserContentArrayItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[PromptDataPromptChatMessageUserContentArrayItemUnionParam](u.OfText, u.OfImageURL)
}

func (u *PromptDataPromptChatMessageUserContentArrayItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfImageURL) {
		return u.OfImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptDataPromptChatMessageUserContentArrayItemUnionParam) GetText() *string {
	if vt := u.OfText; vt != nil && vt.Text.IsPresent() {
		return &vt.Text.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptDataPromptChatMessageUserContentArrayItemUnionParam) GetImageURL() *ChatCompletionContentPartImageImageURLParam {
	if vt := u.OfImageURL; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptDataPromptChatMessageUserContentArrayItemUnionParam) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImageURL; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The property Role is required.
type PromptDataPromptChatMessageAssistantParam struct {
	// Any of "assistant".
	Role         string                                                `json:"role,omitzero,required"`
	Content      param.Opt[string]                                     `json:"content,omitzero"`
	Name         param.Opt[string]                                     `json:"name,omitzero"`
	FunctionCall PromptDataPromptChatMessageAssistantFunctionCallParam `json:"function_call,omitzero"`
	ToolCalls    []ChatCompletionMessageToolCallParam                  `json:"tool_calls,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptDataPromptChatMessageAssistantParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptDataPromptChatMessageAssistantParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptDataPromptChatMessageAssistantParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptDataPromptChatMessageAssistantParam](
		"Role", false, "assistant",
	)
}

// The properties Arguments, Name are required.
type PromptDataPromptChatMessageAssistantFunctionCallParam struct {
	Arguments string `json:"arguments,required"`
	Name      string `json:"name,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptDataPromptChatMessageAssistantFunctionCallParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptDataPromptChatMessageAssistantFunctionCallParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptDataPromptChatMessageAssistantFunctionCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property Role is required.
type PromptDataPromptChatMessageToolParam struct {
	// Any of "tool".
	Role       string            `json:"role,omitzero,required"`
	Content    param.Opt[string] `json:"content,omitzero"`
	ToolCallID param.Opt[string] `json:"tool_call_id,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptDataPromptChatMessageToolParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptDataPromptChatMessageToolParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptDataPromptChatMessageToolParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptDataPromptChatMessageToolParam](
		"Role", false, "tool",
	)
}

// The properties Name, Role are required.
type PromptDataPromptChatMessageFunctionParam struct {
	Name string `json:"name,required"`
	// Any of "function".
	Role    string            `json:"role,omitzero,required"`
	Content param.Opt[string] `json:"content,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptDataPromptChatMessageFunctionParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptDataPromptChatMessageFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptDataPromptChatMessageFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptDataPromptChatMessageFunctionParam](
		"Role", false, "function",
	)
}

// The property Role is required.
type PromptDataPromptChatMessageFallbackParam struct {
	// Any of "model".
	Role    string            `json:"role,omitzero,required"`
	Content param.Opt[string] `json:"content,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptDataPromptChatMessageFallbackParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptDataPromptChatMessageFallbackParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptDataPromptChatMessageFallbackParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptDataPromptChatMessageFallbackParam](
		"Role", false, "model",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PromptDataToolFunctionUnionParam struct {
	OfFunction *PromptDataToolFunctionFunctionParam `json:",omitzero,inline"`
	OfGlobal   *PromptDataToolFunctionGlobalParam   `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u PromptDataToolFunctionUnionParam) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u PromptDataToolFunctionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[PromptDataToolFunctionUnionParam](u.OfFunction, u.OfGlobal)
}

func (u *PromptDataToolFunctionUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFunction) {
		return u.OfFunction
	} else if !param.IsOmitted(u.OfGlobal) {
		return u.OfGlobal
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptDataToolFunctionUnionParam) GetID() *string {
	if vt := u.OfFunction; vt != nil {
		return &vt.ID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptDataToolFunctionUnionParam) GetName() *string {
	if vt := u.OfGlobal; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptDataToolFunctionUnionParam) GetType() *string {
	if vt := u.OfFunction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfGlobal; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The properties ID, Type are required.
type PromptDataToolFunctionFunctionParam struct {
	ID string `json:"id,required"`
	// Any of "function".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptDataToolFunctionFunctionParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptDataToolFunctionFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptDataToolFunctionFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptDataToolFunctionFunctionParam](
		"Type", false, "function",
	)
}

// The properties Name, Type are required.
type PromptDataToolFunctionGlobalParam struct {
	Name string `json:"name,required"`
	// Any of "global".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptDataToolFunctionGlobalParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptDataToolFunctionGlobalParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptDataToolFunctionGlobalParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptDataToolFunctionGlobalParam](
		"Type", false, "global",
	)
}

type PromptOptions struct {
	Model    string                   `json:"model"`
	Params   PromptOptionsParamsUnion `json:"params"`
	Position string                   `json:"position"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Model       resp.Field
		Params      resp.Field
		Position    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptOptions) RawJSON() string { return r.JSON.raw }
func (r *PromptOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PromptOptions to a PromptOptionsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PromptOptionsParam.IsOverridden()
func (r PromptOptions) ToParam() PromptOptionsParam {
	return param.OverrideObj[PromptOptionsParam](r.RawJSON())
}

// PromptOptionsParamsUnion contains all possible properties and values from
// [PromptOptionsParamsOpenAIModelParams],
// [PromptOptionsParamsAnthropicModelParams],
// [PromptOptionsParamsGoogleModelParams],
// [PromptOptionsParamsWindowAIModelParams],
// [PromptOptionsParamsJsCompletionParams].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PromptOptionsParamsUnion struct {
	// This field is from variant [PromptOptionsParamsOpenAIModelParams].
	FrequencyPenalty float64 `json:"frequency_penalty"`
	// This field is from variant [PromptOptionsParamsOpenAIModelParams].
	FunctionCall PromptOptionsParamsOpenAIModelParamsFunctionCallUnion `json:"function_call"`
	// This field is from variant [PromptOptionsParamsOpenAIModelParams].
	MaxCompletionTokens float64 `json:"max_completion_tokens"`
	MaxTokens           float64 `json:"max_tokens"`
	// This field is from variant [PromptOptionsParamsOpenAIModelParams].
	N float64 `json:"n"`
	// This field is from variant [PromptOptionsParamsOpenAIModelParams].
	PresencePenalty float64 `json:"presence_penalty"`
	// This field is from variant [PromptOptionsParamsOpenAIModelParams].
	ReasoningEffort string `json:"reasoning_effort"`
	// This field is from variant [PromptOptionsParamsOpenAIModelParams].
	ResponseFormat PromptOptionsParamsOpenAIModelParamsResponseFormatUnion `json:"response_format"`
	// This field is from variant [PromptOptionsParamsOpenAIModelParams].
	Stop        []string `json:"stop"`
	Temperature float64  `json:"temperature"`
	// This field is from variant [PromptOptionsParamsOpenAIModelParams].
	ToolChoice PromptOptionsParamsOpenAIModelParamsToolChoiceUnion `json:"tool_choice"`
	TopP       float64                                             `json:"top_p"`
	UseCache   bool                                                `json:"use_cache"`
	// This field is from variant [PromptOptionsParamsAnthropicModelParams].
	MaxTokensToSample float64 `json:"max_tokens_to_sample"`
	// This field is from variant [PromptOptionsParamsAnthropicModelParams].
	StopSequences []string `json:"stop_sequences"`
	TopK          float64  `json:"top_k"`
	// This field is from variant [PromptOptionsParamsGoogleModelParams].
	MaxOutputTokens float64 `json:"maxOutputTokens"`
	JSON            struct {
		FrequencyPenalty    resp.Field
		FunctionCall        resp.Field
		MaxCompletionTokens resp.Field
		MaxTokens           resp.Field
		N                   resp.Field
		PresencePenalty     resp.Field
		ReasoningEffort     resp.Field
		ResponseFormat      resp.Field
		Stop                resp.Field
		Temperature         resp.Field
		ToolChoice          resp.Field
		TopP                resp.Field
		UseCache            resp.Field
		MaxTokensToSample   resp.Field
		StopSequences       resp.Field
		TopK                resp.Field
		MaxOutputTokens     resp.Field
		raw                 string
	} `json:"-"`
}

func (u PromptOptionsParamsUnion) AsOpenAIModelParams() (v PromptOptionsParamsOpenAIModelParams) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptOptionsParamsUnion) AsAnthropicModelParams() (v PromptOptionsParamsAnthropicModelParams) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptOptionsParamsUnion) AsGoogleModelParams() (v PromptOptionsParamsGoogleModelParams) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptOptionsParamsUnion) AsWindowAIModelParams() (v PromptOptionsParamsWindowAIModelParams) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptOptionsParamsUnion) AsJsCompletionParams() (v PromptOptionsParamsJsCompletionParams) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PromptOptionsParamsUnion) RawJSON() string { return u.JSON.raw }

func (r *PromptOptionsParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptOptionsParamsOpenAIModelParams struct {
	FrequencyPenalty float64                                               `json:"frequency_penalty"`
	FunctionCall     PromptOptionsParamsOpenAIModelParamsFunctionCallUnion `json:"function_call"`
	// The successor to max_tokens
	MaxCompletionTokens float64 `json:"max_completion_tokens"`
	MaxTokens           float64 `json:"max_tokens"`
	N                   float64 `json:"n"`
	PresencePenalty     float64 `json:"presence_penalty"`
	// Any of "low", "medium", "high".
	ReasoningEffort string                                                  `json:"reasoning_effort"`
	ResponseFormat  PromptOptionsParamsOpenAIModelParamsResponseFormatUnion `json:"response_format,nullable"`
	Stop            []string                                                `json:"stop"`
	Temperature     float64                                                 `json:"temperature"`
	ToolChoice      PromptOptionsParamsOpenAIModelParamsToolChoiceUnion     `json:"tool_choice"`
	TopP            float64                                                 `json:"top_p"`
	UseCache        bool                                                    `json:"use_cache"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		FrequencyPenalty    resp.Field
		FunctionCall        resp.Field
		MaxCompletionTokens resp.Field
		MaxTokens           resp.Field
		N                   resp.Field
		PresencePenalty     resp.Field
		ReasoningEffort     resp.Field
		ResponseFormat      resp.Field
		Stop                resp.Field
		Temperature         resp.Field
		ToolChoice          resp.Field
		TopP                resp.Field
		UseCache            resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptOptionsParamsOpenAIModelParams) RawJSON() string { return r.JSON.raw }
func (r *PromptOptionsParamsOpenAIModelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PromptOptionsParamsOpenAIModelParamsFunctionCallUnion contains all possible
// properties and values from
// [PromptOptionsParamsOpenAIModelParamsFunctionCallString],
// [PromptOptionsParamsOpenAIModelParamsFunctionCallFunction].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPromptOptionssOpenAIModelParamsFunctionCallString]
type PromptOptionsParamsOpenAIModelParamsFunctionCallUnion struct {
	// This field will be present if the value is a
	// [PromptOptionsParamsOpenAIModelParamsFunctionCallString] instead of an object.
	OfPromptOptionssOpenAIModelParamsFunctionCallString PromptOptionsParamsOpenAIModelParamsFunctionCallString `json:",inline"`
	// This field is from variant
	// [PromptOptionsParamsOpenAIModelParamsFunctionCallFunction].
	Name string `json:"name"`
	JSON struct {
		OfPromptOptionssOpenAIModelParamsFunctionCallString resp.Field
		Name                                                resp.Field
		raw                                                 string
	} `json:"-"`
}

func (u PromptOptionsParamsOpenAIModelParamsFunctionCallUnion) AsPromptOptionsParamsOpenAIModelParamsFunctionCallString() (v PromptOptionsParamsOpenAIModelParamsFunctionCallString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptOptionsParamsOpenAIModelParamsFunctionCallUnion) AsFunction() (v PromptOptionsParamsOpenAIModelParamsFunctionCallFunction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PromptOptionsParamsOpenAIModelParamsFunctionCallUnion) RawJSON() string { return u.JSON.raw }

func (r *PromptOptionsParamsOpenAIModelParamsFunctionCallUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptOptionsParamsOpenAIModelParamsFunctionCallString string

const (
	PromptOptionsParamsOpenAIModelParamsFunctionCallStringAuto PromptOptionsParamsOpenAIModelParamsFunctionCallString = "auto"
	PromptOptionsParamsOpenAIModelParamsFunctionCallStringNone PromptOptionsParamsOpenAIModelParamsFunctionCallString = "none"
)

type PromptOptionsParamsOpenAIModelParamsFunctionCallFunction struct {
	Name string `json:"name,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptOptionsParamsOpenAIModelParamsFunctionCallFunction) RawJSON() string { return r.JSON.raw }
func (r *PromptOptionsParamsOpenAIModelParamsFunctionCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PromptOptionsParamsOpenAIModelParamsResponseFormatUnion contains all possible
// properties and values from
// [PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObject],
// [PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchema],
// [PromptOptionsParamsOpenAIModelParamsResponseFormatText].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PromptOptionsParamsOpenAIModelParamsResponseFormatUnion struct {
	Type string `json:"type"`
	// This field is from variant
	// [PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchema].
	JsonSchema PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchema `json:"json_schema"`
	JSON       struct {
		Type       resp.Field
		JsonSchema resp.Field
		raw        string
	} `json:"-"`
}

func (u PromptOptionsParamsOpenAIModelParamsResponseFormatUnion) AsJsonObject() (v PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptOptionsParamsOpenAIModelParamsResponseFormatUnion) AsJsonSchema() (v PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptOptionsParamsOpenAIModelParamsResponseFormatUnion) AsText() (v PromptOptionsParamsOpenAIModelParamsResponseFormatText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PromptOptionsParamsOpenAIModelParamsResponseFormatUnion) RawJSON() string { return u.JSON.raw }

func (r *PromptOptionsParamsOpenAIModelParamsResponseFormatUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObject struct {
	// Any of "json_object".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObject) RawJSON() string {
	return r.JSON.raw
}
func (r *PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchema struct {
	JsonSchema PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchema `json:"json_schema,required"`
	// Any of "json_schema".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		JsonSchema  resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchema) RawJSON() string {
	return r.JSON.raw
}
func (r *PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchema struct {
	Name        string `json:"name,required"`
	Description string `json:"description"`
	Schema      string `json:"schema"`
	Strict      bool   `json:"strict,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		Description resp.Field
		Schema      resp.Field
		Strict      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchema) RawJSON() string {
	return r.JSON.raw
}
func (r *PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptOptionsParamsOpenAIModelParamsResponseFormatText struct {
	// Any of "text".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptOptionsParamsOpenAIModelParamsResponseFormatText) RawJSON() string { return r.JSON.raw }
func (r *PromptOptionsParamsOpenAIModelParamsResponseFormatText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PromptOptionsParamsOpenAIModelParamsToolChoiceUnion contains all possible
// properties and values from
// [PromptOptionsParamsOpenAIModelParamsToolChoiceString],
// [PromptOptionsParamsOpenAIModelParamsToolChoiceFunction].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPromptOptionssOpenAIModelParamsToolChoiceString]
type PromptOptionsParamsOpenAIModelParamsToolChoiceUnion struct {
	// This field will be present if the value is a
	// [PromptOptionsParamsOpenAIModelParamsToolChoiceString] instead of an object.
	OfPromptOptionssOpenAIModelParamsToolChoiceString PromptOptionsParamsOpenAIModelParamsToolChoiceString `json:",inline"`
	// This field is from variant
	// [PromptOptionsParamsOpenAIModelParamsToolChoiceFunction].
	Function PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction `json:"function"`
	// This field is from variant
	// [PromptOptionsParamsOpenAIModelParamsToolChoiceFunction].
	Type string `json:"type"`
	JSON struct {
		OfPromptOptionssOpenAIModelParamsToolChoiceString resp.Field
		Function                                          resp.Field
		Type                                              resp.Field
		raw                                               string
	} `json:"-"`
}

func (u PromptOptionsParamsOpenAIModelParamsToolChoiceUnion) AsPromptOptionsParamsOpenAIModelParamsToolChoiceString() (v PromptOptionsParamsOpenAIModelParamsToolChoiceString) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PromptOptionsParamsOpenAIModelParamsToolChoiceUnion) AsFunction() (v PromptOptionsParamsOpenAIModelParamsToolChoiceFunction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PromptOptionsParamsOpenAIModelParamsToolChoiceUnion) RawJSON() string { return u.JSON.raw }

func (r *PromptOptionsParamsOpenAIModelParamsToolChoiceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptOptionsParamsOpenAIModelParamsToolChoiceString string

const (
	PromptOptionsParamsOpenAIModelParamsToolChoiceStringAuto     PromptOptionsParamsOpenAIModelParamsToolChoiceString = "auto"
	PromptOptionsParamsOpenAIModelParamsToolChoiceStringNone     PromptOptionsParamsOpenAIModelParamsToolChoiceString = "none"
	PromptOptionsParamsOpenAIModelParamsToolChoiceStringRequired PromptOptionsParamsOpenAIModelParamsToolChoiceString = "required"
)

type PromptOptionsParamsOpenAIModelParamsToolChoiceFunction struct {
	Function PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction `json:"function,required"`
	// Any of "function".
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Function    resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptOptionsParamsOpenAIModelParamsToolChoiceFunction) RawJSON() string { return r.JSON.raw }
func (r *PromptOptionsParamsOpenAIModelParamsToolChoiceFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction struct {
	Name string `json:"name,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction) RawJSON() string {
	return r.JSON.raw
}
func (r *PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptOptionsParamsAnthropicModelParams struct {
	MaxTokens   float64 `json:"max_tokens,required"`
	Temperature float64 `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample float64  `json:"max_tokens_to_sample"`
	StopSequences     []string `json:"stop_sequences"`
	TopK              float64  `json:"top_k"`
	TopP              float64  `json:"top_p"`
	UseCache          bool     `json:"use_cache"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		MaxTokens         resp.Field
		Temperature       resp.Field
		MaxTokensToSample resp.Field
		StopSequences     resp.Field
		TopK              resp.Field
		TopP              resp.Field
		UseCache          resp.Field
		ExtraFields       map[string]resp.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptOptionsParamsAnthropicModelParams) RawJSON() string { return r.JSON.raw }
func (r *PromptOptionsParamsAnthropicModelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptOptionsParamsGoogleModelParams struct {
	MaxOutputTokens float64 `json:"maxOutputTokens"`
	Temperature     float64 `json:"temperature"`
	TopK            float64 `json:"topK"`
	TopP            float64 `json:"topP"`
	UseCache        bool    `json:"use_cache"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		MaxOutputTokens resp.Field
		Temperature     resp.Field
		TopK            resp.Field
		TopP            resp.Field
		UseCache        resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptOptionsParamsGoogleModelParams) RawJSON() string { return r.JSON.raw }
func (r *PromptOptionsParamsGoogleModelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptOptionsParamsWindowAIModelParams struct {
	Temperature float64 `json:"temperature"`
	TopK        float64 `json:"topK"`
	UseCache    bool    `json:"use_cache"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Temperature resp.Field
		TopK        resp.Field
		UseCache    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptOptionsParamsWindowAIModelParams) RawJSON() string { return r.JSON.raw }
func (r *PromptOptionsParamsWindowAIModelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptOptionsParamsJsCompletionParams struct {
	UseCache bool `json:"use_cache"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		UseCache    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptOptionsParamsJsCompletionParams) RawJSON() string { return r.JSON.raw }
func (r *PromptOptionsParamsJsCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptOptionsParam struct {
	Model    param.Opt[string]             `json:"model,omitzero"`
	Position param.Opt[string]             `json:"position,omitzero"`
	Params   PromptOptionsParamsUnionParam `json:"params,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptOptionsParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r PromptOptionsParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptOptionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PromptOptionsParamsUnionParam struct {
	OfOpenAIModels    *PromptOptionsParamsOpenAIModelParamsParam    `json:",omitzero,inline"`
	OfAnthropicModels *PromptOptionsParamsAnthropicModelParamsParam `json:",omitzero,inline"`
	OfGoogleModels    *PromptOptionsParamsGoogleModelParamsParam    `json:",omitzero,inline"`
	OfWindowAIModels  *PromptOptionsParamsWindowAIModelParamsParam  `json:",omitzero,inline"`
	OfJsCompletions   *PromptOptionsParamsJsCompletionParamsParam   `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u PromptOptionsParamsUnionParam) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u PromptOptionsParamsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[PromptOptionsParamsUnionParam](u.OfOpenAIModels,
		u.OfAnthropicModels,
		u.OfGoogleModels,
		u.OfWindowAIModels,
		u.OfJsCompletions)
}

func (u *PromptOptionsParamsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfOpenAIModels) {
		return u.OfOpenAIModels
	} else if !param.IsOmitted(u.OfAnthropicModels) {
		return u.OfAnthropicModels
	} else if !param.IsOmitted(u.OfGoogleModels) {
		return u.OfGoogleModels
	} else if !param.IsOmitted(u.OfWindowAIModels) {
		return u.OfWindowAIModels
	} else if !param.IsOmitted(u.OfJsCompletions) {
		return u.OfJsCompletions
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetFrequencyPenalty() *float64 {
	if vt := u.OfOpenAIModels; vt != nil && vt.FrequencyPenalty.IsPresent() {
		return &vt.FrequencyPenalty.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetFunctionCall() *PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam {
	if vt := u.OfOpenAIModels; vt != nil {
		return &vt.FunctionCall
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetMaxCompletionTokens() *float64 {
	if vt := u.OfOpenAIModels; vt != nil && vt.MaxCompletionTokens.IsPresent() {
		return &vt.MaxCompletionTokens.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetN() *float64 {
	if vt := u.OfOpenAIModels; vt != nil && vt.N.IsPresent() {
		return &vt.N.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetPresencePenalty() *float64 {
	if vt := u.OfOpenAIModels; vt != nil && vt.PresencePenalty.IsPresent() {
		return &vt.PresencePenalty.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetReasoningEffort() *string {
	if vt := u.OfOpenAIModels; vt != nil {
		return &vt.ReasoningEffort
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetResponseFormat() *PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam {
	if vt := u.OfOpenAIModels; vt != nil {
		return &vt.ResponseFormat
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetStop() []string {
	if vt := u.OfOpenAIModels; vt != nil {
		return vt.Stop
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetToolChoice() *PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam {
	if vt := u.OfOpenAIModels; vt != nil {
		return &vt.ToolChoice
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetMaxTokensToSample() *float64 {
	if vt := u.OfAnthropicModels; vt != nil && vt.MaxTokensToSample.IsPresent() {
		return &vt.MaxTokensToSample.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetStopSequences() []string {
	if vt := u.OfAnthropicModels; vt != nil {
		return vt.StopSequences
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetMaxOutputTokens() *float64 {
	if vt := u.OfGoogleModels; vt != nil && vt.MaxOutputTokens.IsPresent() {
		return &vt.MaxOutputTokens.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetMaxTokens() *float64 {
	if vt := u.OfOpenAIModels; vt != nil && vt.MaxTokens.IsPresent() {
		return &vt.MaxTokens.Value
	} else if vt := u.OfAnthropicModels; vt != nil {
		return (*float64)(&vt.MaxTokens)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetTemperature() *float64 {
	if vt := u.OfOpenAIModels; vt != nil && vt.Temperature.IsPresent() {
		return &vt.Temperature.Value
	} else if vt := u.OfAnthropicModels; vt != nil {
		return (*float64)(&vt.Temperature)
	} else if vt := u.OfGoogleModels; vt != nil && vt.Temperature.IsPresent() {
		return &vt.Temperature.Value
	} else if vt := u.OfWindowAIModels; vt != nil && vt.Temperature.IsPresent() {
		return &vt.Temperature.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetTopP() *float64 {
	if vt := u.OfOpenAIModels; vt != nil && vt.TopP.IsPresent() {
		return &vt.TopP.Value
	} else if vt := u.OfAnthropicModels; vt != nil && vt.TopP.IsPresent() {
		return &vt.TopP.Value
	} else if vt := u.OfGoogleModels; vt != nil && vt.TopP.IsPresent() {
		return &vt.TopP.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetUseCache() *bool {
	if vt := u.OfOpenAIModels; vt != nil && vt.UseCache.IsPresent() {
		return &vt.UseCache.Value
	} else if vt := u.OfAnthropicModels; vt != nil && vt.UseCache.IsPresent() {
		return &vt.UseCache.Value
	} else if vt := u.OfGoogleModels; vt != nil && vt.UseCache.IsPresent() {
		return &vt.UseCache.Value
	} else if vt := u.OfWindowAIModels; vt != nil && vt.UseCache.IsPresent() {
		return &vt.UseCache.Value
	} else if vt := u.OfJsCompletions; vt != nil && vt.UseCache.IsPresent() {
		return &vt.UseCache.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsUnionParam) GetTopK() *float64 {
	if vt := u.OfAnthropicModels; vt != nil && vt.TopK.IsPresent() {
		return &vt.TopK.Value
	} else if vt := u.OfGoogleModels; vt != nil && vt.TopK.IsPresent() {
		return &vt.TopK.Value
	} else if vt := u.OfWindowAIModels; vt != nil && vt.TopK.IsPresent() {
		return &vt.TopK.Value
	}
	return nil
}

type PromptOptionsParamsOpenAIModelParamsParam struct {
	FrequencyPenalty param.Opt[float64] `json:"frequency_penalty,omitzero"`
	// The successor to max_tokens
	MaxCompletionTokens param.Opt[float64]                                           `json:"max_completion_tokens,omitzero"`
	MaxTokens           param.Opt[float64]                                           `json:"max_tokens,omitzero"`
	N                   param.Opt[float64]                                           `json:"n,omitzero"`
	PresencePenalty     param.Opt[float64]                                           `json:"presence_penalty,omitzero"`
	Temperature         param.Opt[float64]                                           `json:"temperature,omitzero"`
	TopP                param.Opt[float64]                                           `json:"top_p,omitzero"`
	UseCache            param.Opt[bool]                                              `json:"use_cache,omitzero"`
	ResponseFormat      PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam `json:"response_format,omitzero"`
	FunctionCall        PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam   `json:"function_call,omitzero"`
	// Any of "low", "medium", "high".
	ReasoningEffort string                                                   `json:"reasoning_effort,omitzero"`
	Stop            []string                                                 `json:"stop,omitzero"`
	ToolChoice      PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam `json:"tool_choice,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptOptionsParamsOpenAIModelParamsParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptOptionsParamsOpenAIModelParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptOptionsParamsOpenAIModelParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptOptionsParamsOpenAIModelParamsParam](
		"ReasoningEffort", false, "low", "medium", "high",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfPromptOptionssOpenAIModelParamsFunctionCallString)
	OfPromptOptionssOpenAIModelParamsFunctionCallString param.Opt[PromptOptionsParamsOpenAIModelParamsFunctionCallString] `json:",omitzero,inline"`
	OfFunction                                          *PromptOptionsParamsOpenAIModelParamsFunctionCallFunctionParam    `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam](u.OfPromptOptionssOpenAIModelParamsFunctionCallString, u.OfFunction)
}

func (u *PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam) asAny() any {
	if !param.IsOmitted(u.OfPromptOptionssOpenAIModelParamsFunctionCallString) {
		return &u.OfPromptOptionssOpenAIModelParamsFunctionCallString
	} else if !param.IsOmitted(u.OfFunction) {
		return u.OfFunction
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam) GetName() *string {
	if vt := u.OfFunction; vt != nil {
		return &vt.Name
	}
	return nil
}

// The property Name is required.
type PromptOptionsParamsOpenAIModelParamsFunctionCallFunctionParam struct {
	Name string `json:"name,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptOptionsParamsOpenAIModelParamsFunctionCallFunctionParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptOptionsParamsOpenAIModelParamsFunctionCallFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptOptionsParamsOpenAIModelParamsFunctionCallFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam struct {
	OfJsonObject *PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectParam `json:",omitzero,inline"`
	OfJsonSchema *PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaParam `json:",omitzero,inline"`
	OfText       *PromptOptionsParamsOpenAIModelParamsResponseFormatTextParam       `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam](u.OfJsonObject, u.OfJsonSchema, u.OfText)
}

func (u *PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam) asAny() any {
	if !param.IsOmitted(u.OfJsonObject) {
		return u.OfJsonObject
	} else if !param.IsOmitted(u.OfJsonSchema) {
		return u.OfJsonSchema
	} else if !param.IsOmitted(u.OfText) {
		return u.OfText
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam) GetJsonSchema() *PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchemaParam {
	if vt := u.OfJsonSchema; vt != nil {
		return &vt.JsonSchema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam) GetType() *string {
	if vt := u.OfJsonObject; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfJsonSchema; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The property Type is required.
type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectParam struct {
	// Any of "json_object".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectParam](
		"Type", false, "json_object",
	)
}

// The properties JsonSchema, Type are required.
type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaParam struct {
	JsonSchema PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchemaParam `json:"json_schema,omitzero,required"`
	// Any of "json_schema".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaParam](
		"Type", false, "json_schema",
	)
}

// The property Name is required.
type PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchemaParam struct {
	Name        string            `json:"name,required"`
	Strict      param.Opt[bool]   `json:"strict,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	Schema      string            `json:"schema,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchemaParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptOptionsParamsOpenAIModelParamsResponseFormatJsonSchemaJsonSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property Type is required.
type PromptOptionsParamsOpenAIModelParamsResponseFormatTextParam struct {
	// Any of "text".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptOptionsParamsOpenAIModelParamsResponseFormatTextParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptOptionsParamsOpenAIModelParamsResponseFormatTextParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptOptionsParamsOpenAIModelParamsResponseFormatTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptOptionsParamsOpenAIModelParamsResponseFormatTextParam](
		"Type", false, "text",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfPromptOptionssOpenAIModelParamsToolChoiceString)
	OfPromptOptionssOpenAIModelParamsToolChoiceString param.Opt[PromptOptionsParamsOpenAIModelParamsToolChoiceString] `json:",omitzero,inline"`
	OfFunction                                        *PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionParam    `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam](u.OfPromptOptionssOpenAIModelParamsToolChoiceString, u.OfFunction)
}

func (u *PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfPromptOptionssOpenAIModelParamsToolChoiceString) {
		return &u.OfPromptOptionssOpenAIModelParamsToolChoiceString
	} else if !param.IsOmitted(u.OfFunction) {
		return u.OfFunction
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam) GetFunction() *PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionParam {
	if vt := u.OfFunction; vt != nil {
		return &vt.Function
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam) GetType() *string {
	if vt := u.OfFunction; vt != nil {
		return &vt.Type
	}
	return nil
}

// The properties Function, Type are required.
type PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionParam struct {
	Function PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionParam `json:"function,omitzero,required"`
	// Any of "function".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionParam](
		"Type", false, "function",
	)
}

// The property Name is required.
type PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionParam struct {
	Name string `json:"name,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptOptionsParamsOpenAIModelParamsToolChoiceFunctionFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties MaxTokens, Temperature are required.
type PromptOptionsParamsAnthropicModelParamsParam struct {
	MaxTokens   float64 `json:"max_tokens,required"`
	Temperature float64 `json:"temperature,required"`
	// This is a legacy parameter that should not be used.
	MaxTokensToSample param.Opt[float64] `json:"max_tokens_to_sample,omitzero"`
	TopK              param.Opt[float64] `json:"top_k,omitzero"`
	TopP              param.Opt[float64] `json:"top_p,omitzero"`
	UseCache          param.Opt[bool]    `json:"use_cache,omitzero"`
	StopSequences     []string           `json:"stop_sequences,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptOptionsParamsAnthropicModelParamsParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptOptionsParamsAnthropicModelParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptOptionsParamsAnthropicModelParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type PromptOptionsParamsGoogleModelParamsParam struct {
	MaxOutputTokens param.Opt[float64] `json:"maxOutputTokens,omitzero"`
	Temperature     param.Opt[float64] `json:"temperature,omitzero"`
	TopK            param.Opt[float64] `json:"topK,omitzero"`
	TopP            param.Opt[float64] `json:"topP,omitzero"`
	UseCache        param.Opt[bool]    `json:"use_cache,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptOptionsParamsGoogleModelParamsParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptOptionsParamsGoogleModelParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptOptionsParamsGoogleModelParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type PromptOptionsParamsWindowAIModelParamsParam struct {
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	TopK        param.Opt[float64] `json:"topK,omitzero"`
	UseCache    param.Opt[bool]    `json:"use_cache,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptOptionsParamsWindowAIModelParamsParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptOptionsParamsWindowAIModelParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptOptionsParamsWindowAIModelParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type PromptOptionsParamsJsCompletionParamsParam struct {
	UseCache param.Opt[bool] `json:"use_cache,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PromptOptionsParamsJsCompletionParamsParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r PromptOptionsParamsJsCompletionParamsParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptOptionsParamsJsCompletionParamsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Metadata about the state of the repo when the experiment was created
type RepoInfo struct {
	// Email of the author of the most recent commit
	AuthorEmail string `json:"author_email,nullable"`
	// Name of the author of the most recent commit
	AuthorName string `json:"author_name,nullable"`
	// Name of the branch the most recent commit belongs to
	Branch string `json:"branch,nullable"`
	// SHA of most recent commit
	Commit string `json:"commit,nullable"`
	// Most recent commit message
	CommitMessage string `json:"commit_message,nullable"`
	// Time of the most recent commit
	CommitTime string `json:"commit_time,nullable"`
	// Whether or not the repo had uncommitted changes when snapshotted
	Dirty bool `json:"dirty,nullable"`
	// If the repo was dirty when run, this includes the diff between the current state
	// of the repo and the most recent commit.
	GitDiff string `json:"git_diff,nullable"`
	// Name of the tag on the most recent commit
	Tag string `json:"tag,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AuthorEmail   resp.Field
		AuthorName    resp.Field
		Branch        resp.Field
		Commit        resp.Field
		CommitMessage resp.Field
		CommitTime    resp.Field
		Dirty         resp.Field
		GitDiff       resp.Field
		Tag           resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RepoInfo) RawJSON() string { return r.JSON.raw }
func (r *RepoInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RepoInfo to a RepoInfoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RepoInfoParam.IsOverridden()
func (r RepoInfo) ToParam() RepoInfoParam {
	return param.OverrideObj[RepoInfoParam](r.RawJSON())
}

// Metadata about the state of the repo when the experiment was created
type RepoInfoParam struct {
	// Email of the author of the most recent commit
	AuthorEmail param.Opt[string] `json:"author_email,omitzero"`
	// Name of the author of the most recent commit
	AuthorName param.Opt[string] `json:"author_name,omitzero"`
	// Name of the branch the most recent commit belongs to
	Branch param.Opt[string] `json:"branch,omitzero"`
	// SHA of most recent commit
	Commit param.Opt[string] `json:"commit,omitzero"`
	// Most recent commit message
	CommitMessage param.Opt[string] `json:"commit_message,omitzero"`
	// Time of the most recent commit
	CommitTime param.Opt[string] `json:"commit_time,omitzero"`
	// Whether or not the repo had uncommitted changes when snapshotted
	Dirty param.Opt[bool] `json:"dirty,omitzero"`
	// If the repo was dirty when run, this includes the diff between the current state
	// of the repo and the most recent commit.
	GitDiff param.Opt[string] `json:"git_diff,omitzero"`
	// Name of the tag on the most recent commit
	Tag param.Opt[string] `json:"tag,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RepoInfoParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r RepoInfoParam) MarshalJSON() (data []byte, err error) {
	type shadow RepoInfoParam
	return param.MarshalObject(r, (*shadow)(&r))
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
	UserID string `json:"user_id,nullable" format:"uuid"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                resp.Field
		Name              resp.Field
		Created           resp.Field
		DeletedAt         resp.Field
		Description       resp.Field
		MemberPermissions resp.Field
		MemberRoles       resp.Field
		OrgID             resp.Field
		UserID            resp.Field
		ExtraFields       map[string]resp.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Role) RawJSON() string { return r.JSON.raw }
func (r *Role) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoleMemberPermission struct {
	// Each permission permits a certain type of operation on an object in the system
	//
	// Permissions can be assigned to to objects on an individual basis, or grouped
	// into roles
	//
	// Any of "create", "read", "update", "delete", "create_acls", "read_acls",
	// "update_acls", "delete_acls".
	Permission Permission `json:"permission,required"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	RestrictObjectType ACLObjectType `json:"restrict_object_type,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Permission         resp.Field
		RestrictObjectType resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoleMemberPermission) RawJSON() string { return r.JSON.raw }
func (r *RoleMemberPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary of a score's performance
type ScoreSummary struct {
	// Number of improvements in the score
	Improvements int64 `json:"improvements,required"`
	// Name of the score
	Name string `json:"name,required"`
	// Number of regressions in the score
	Regressions int64 `json:"regressions,required"`
	// Average score across all examples
	Score float64 `json:"score,required"`
	// Difference in score between the current and comparison experiment
	Diff float64 `json:"diff"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Improvements resp.Field
		Name         resp.Field
		Regressions  resp.Field
		Score        resp.Field
		Diff         resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoreSummary) RawJSON() string { return r.JSON.raw }
func (r *ScoreSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Human-identifying attributes of the span, such as name, type, etc.
type SpanAttributes struct {
	// Name of the span, for display purposes only
	Name string `json:"name,nullable"`
	// Type of the span, for display purposes only
	//
	// Any of "llm", "score", "function", "eval", "task", "tool".
	Type        SpanType               `json:"type,nullable"`
	ExtraFields map[string]interface{} `json:",extras"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpanAttributes) RawJSON() string { return r.JSON.raw }
func (r *SpanAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SpanAttributes to a SpanAttributesParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SpanAttributesParam.IsOverridden()
func (r SpanAttributes) ToParam() SpanAttributesParam {
	return param.OverrideObj[SpanAttributesParam](r.RawJSON())
}

// Human-identifying attributes of the span, such as name, type, etc.
type SpanAttributesParam struct {
	// Name of the span, for display purposes only
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of the span, for display purposes only
	//
	// Any of "llm", "score", "function", "eval", "task", "tool".
	Type        SpanType               `json:"type,omitzero"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SpanAttributesParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r SpanAttributesParam) MarshalJSON() (data []byte, err error) {
	type shadow SpanAttributesParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type SpanIFrame struct {
	// Unique identifier for the span iframe
	ID string `json:"id,required" format:"uuid"`
	// Name of the span iframe
	Name string `json:"name,required"`
	// Unique identifier for the project that the span iframe belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// URL to embed the project viewer in an iframe
	URL string `json:"url,required"`
	// Date of span iframe creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Date of span iframe deletion, or null if the span iframe is still active
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Textual description of the span iframe
	Description string `json:"description,nullable"`
	// Whether to post messages to the iframe containing the span's data. This is
	// useful when you want to render more data than fits in the URL.
	PostMessage bool `json:"post_message,nullable"`
	// Identifies the user who created the span iframe
	UserID string `json:"user_id,nullable" format:"uuid"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Name        resp.Field
		ProjectID   resp.Field
		URL         resp.Field
		Created     resp.Field
		DeletedAt   resp.Field
		Description resp.Field
		PostMessage resp.Field
		UserID      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpanIFrame) RawJSON() string { return r.JSON.raw }
func (r *SpanIFrame) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the span, for display purposes only
type SpanType string

const (
	SpanTypeLlm      SpanType = "llm"
	SpanTypeScore    SpanType = "score"
	SpanTypeFunction SpanType = "function"
	SpanTypeEval     SpanType = "eval"
	SpanTypeTask     SpanType = "task"
	SpanTypeTool     SpanType = "tool"
)

// Summary of a dataset
type SummarizeDatasetResponse struct {
	// Name of the dataset
	DatasetName string `json:"dataset_name,required"`
	// URL to the dataset's page in the Braintrust app
	DatasetURL string `json:"dataset_url,required" format:"uri"`
	// Name of the project that the dataset belongs to
	ProjectName string `json:"project_name,required"`
	// URL to the project's page in the Braintrust app
	ProjectURL string `json:"project_url,required" format:"uri"`
	// Summary of a dataset's data
	DataSummary DataSummary `json:"data_summary,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		DatasetName resp.Field
		DatasetURL  resp.Field
		ProjectName resp.Field
		ProjectURL  resp.Field
		DataSummary resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SummarizeDatasetResponse) RawJSON() string { return r.JSON.raw }
func (r *SummarizeDatasetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary of an experiment
type SummarizeExperimentResponse struct {
	// Name of the experiment
	ExperimentName string `json:"experiment_name,required"`
	// URL to the experiment's page in the Braintrust app
	ExperimentURL string `json:"experiment_url,required" format:"uri"`
	// Name of the project that the experiment belongs to
	ProjectName string `json:"project_name,required"`
	// URL to the project's page in the Braintrust app
	ProjectURL string `json:"project_url,required" format:"uri"`
	// The experiment which scores are baselined against
	ComparisonExperimentName string `json:"comparison_experiment_name,nullable"`
	// Summary of the experiment's metrics
	Metrics map[string]MetricSummary `json:"metrics,nullable"`
	// Summary of the experiment's scores
	Scores map[string]ScoreSummary `json:"scores,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ExperimentName           resp.Field
		ExperimentURL            resp.Field
		ProjectName              resp.Field
		ProjectURL               resp.Field
		ComparisonExperimentName resp.Field
		Metrics                  resp.Field
		Scores                   resp.Field
		ExtraFields              map[string]resp.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SummarizeExperimentResponse) RawJSON() string { return r.JSON.raw }
func (r *SummarizeExperimentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type User struct {
	// Unique identifier for the user
	ID string `json:"id,required" format:"uuid"`
	// URL of the user's Avatar image
	AvatarURL string `json:"avatar_url,nullable"`
	// Date of user creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// The user's email
	Email string `json:"email,nullable"`
	// Family name of the user
	FamilyName string `json:"family_name,nullable"`
	// Given name of the user
	GivenName string `json:"given_name,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		AvatarURL   resp.Field
		Created     resp.Field
		Email       resp.Field
		FamilyName  resp.Field
		GivenName   resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type View struct {
	// Unique identifier for the view
	ID string `json:"id,required" format:"uuid"`
	// Name of the view
	Name string `json:"name,required"`
	// The id of the object the view applies to
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType ACLObjectType `json:"object_type,required"`
	// Type of table that the view corresponds to.
	//
	// Any of "projects", "experiments", "experiment", "playgrounds", "playground",
	// "datasets", "dataset", "prompts", "tools", "scorers", "logs".
	ViewType ViewViewType `json:"view_type,required"`
	// Date of view creation
	Created time.Time `json:"created,nullable" format:"date-time"`
	// Date of role deletion, or null if the role is still active
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Options for the view in the app
	Options ViewOptions `json:"options,nullable"`
	// Identifies the user who created the view
	UserID string `json:"user_id,nullable" format:"uuid"`
	// The view definition
	ViewData ViewData `json:"view_data,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Name        resp.Field
		ObjectID    resp.Field
		ObjectType  resp.Field
		ViewType    resp.Field
		Created     resp.Field
		DeletedAt   resp.Field
		Options     resp.Field
		UserID      resp.Field
		ViewData    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r View) RawJSON() string { return r.JSON.raw }
func (r *View) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of table that the view corresponds to.
type ViewViewType string

const (
	ViewViewTypeProjects    ViewViewType = "projects"
	ViewViewTypeExperiments ViewViewType = "experiments"
	ViewViewTypeExperiment  ViewViewType = "experiment"
	ViewViewTypePlaygrounds ViewViewType = "playgrounds"
	ViewViewTypePlayground  ViewViewType = "playground"
	ViewViewTypeDatasets    ViewViewType = "datasets"
	ViewViewTypeDataset     ViewViewType = "dataset"
	ViewViewTypePrompts     ViewViewType = "prompts"
	ViewViewTypeTools       ViewViewType = "tools"
	ViewViewTypeScorers     ViewViewType = "scorers"
	ViewViewTypeLogs        ViewViewType = "logs"
)

// The view definition
type ViewData struct {
	Search ViewDataSearch `json:"search,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Search      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewData) RawJSON() string { return r.JSON.raw }
func (r *ViewData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ViewData to a ViewDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ViewDataParam.IsOverridden()
func (r ViewData) ToParam() ViewDataParam {
	return param.OverrideObj[ViewDataParam](r.RawJSON())
}

// The view definition
type ViewDataParam struct {
	Search ViewDataSearchParam `json:"search,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ViewDataParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ViewDataParam) MarshalJSON() (data []byte, err error) {
	type shadow ViewDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ViewDataSearch struct {
	Filter []interface{} `json:"filter,nullable"`
	Match  []interface{} `json:"match,nullable"`
	Sort   []interface{} `json:"sort,nullable"`
	Tag    []interface{} `json:"tag,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Filter      resp.Field
		Match       resp.Field
		Sort        resp.Field
		Tag         resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewDataSearch) RawJSON() string { return r.JSON.raw }
func (r *ViewDataSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ViewDataSearch to a ViewDataSearchParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ViewDataSearchParam.IsOverridden()
func (r ViewDataSearch) ToParam() ViewDataSearchParam {
	return param.OverrideObj[ViewDataSearchParam](r.RawJSON())
}

type ViewDataSearchParam struct {
	Filter []interface{} `json:"filter,omitzero"`
	Match  []interface{} `json:"match,omitzero"`
	Sort   []interface{} `json:"sort,omitzero"`
	Tag    []interface{} `json:"tag,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ViewDataSearchParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ViewDataSearchParam) MarshalJSON() (data []byte, err error) {
	type shadow ViewDataSearchParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Options for the view in the app
type ViewOptions struct {
	ColumnOrder      []string           `json:"columnOrder,nullable"`
	ColumnSizing     map[string]float64 `json:"columnSizing,nullable"`
	ColumnVisibility map[string]bool    `json:"columnVisibility,nullable"`
	Grouping         string             `json:"grouping,nullable"`
	Layout           string             `json:"layout,nullable"`
	RowHeight        string             `json:"rowHeight,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ColumnOrder      resp.Field
		ColumnSizing     resp.Field
		ColumnVisibility resp.Field
		Grouping         resp.Field
		Layout           resp.Field
		RowHeight        resp.Field
		ExtraFields      map[string]resp.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewOptions) RawJSON() string { return r.JSON.raw }
func (r *ViewOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ViewOptions to a ViewOptionsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ViewOptionsParam.IsOverridden()
func (r ViewOptions) ToParam() ViewOptionsParam {
	return param.OverrideObj[ViewOptionsParam](r.RawJSON())
}

// Options for the view in the app
type ViewOptionsParam struct {
	Grouping         param.Opt[string]  `json:"grouping,omitzero"`
	Layout           param.Opt[string]  `json:"layout,omitzero"`
	RowHeight        param.Opt[string]  `json:"rowHeight,omitzero"`
	ColumnOrder      []string           `json:"columnOrder,omitzero"`
	ColumnSizing     map[string]float64 `json:"columnSizing,omitzero"`
	ColumnVisibility map[string]bool    `json:"columnVisibility,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ViewOptionsParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ViewOptionsParam) MarshalJSON() (data []byte, err error) {
	type shadow ViewOptionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Type of table that the view corresponds to.
type ViewType string

const (
	ViewTypeProjects    ViewType = "projects"
	ViewTypeExperiments ViewType = "experiments"
	ViewTypeExperiment  ViewType = "experiment"
	ViewTypePlaygrounds ViewType = "playgrounds"
	ViewTypePlayground  ViewType = "playground"
	ViewTypeDatasets    ViewType = "datasets"
	ViewTypeDataset     ViewType = "dataset"
	ViewTypePrompts     ViewType = "prompts"
	ViewTypeTools       ViewType = "tools"
	ViewTypeScorers     ViewType = "scorers"
	ViewTypeLogs        ViewType = "logs"
)
