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

// FunctionService contains methods and other services that help with interacting
// with the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionService] method instead.
type FunctionService struct {
	Options []option.RequestOption
}

// NewFunctionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFunctionService(opts ...option.RequestOption) (r FunctionService) {
	r = FunctionService{}
	r.Options = opts
	return
}

// Create a new function. If there is an existing function in the project with the
// same slug as the one specified in the request, will return the existing function
// unmodified
func (r *FunctionService) New(ctx context.Context, body FunctionNewParams, opts ...option.RequestOption) (res *shared.Function, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/function"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a function object by its id
func (r *FunctionService) Get(ctx context.Context, functionID string, opts ...option.RequestOption) (res *shared.Function, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required function_id parameter")
		return
	}
	path := fmt.Sprintf("v1/function/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update a function object. Specify the fields to update in the payload.
// Any object-type fields will be deep-merged with existing content. Currently we
// do not support removing fields or setting them to null.
func (r *FunctionService) Update(ctx context.Context, functionID string, body FunctionUpdateParams, opts ...option.RequestOption) (res *shared.Function, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required function_id parameter")
		return
	}
	path := fmt.Sprintf("v1/function/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List out all functions. The functions are sorted by creation date, with the most
// recently-created functions coming first
func (r *FunctionService) List(ctx context.Context, query FunctionListParams, opts ...option.RequestOption) (res *pagination.ListObjects[shared.Function], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/function"
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

// List out all functions. The functions are sorted by creation date, with the most
// recently-created functions coming first
func (r *FunctionService) ListAutoPaging(ctx context.Context, query FunctionListParams, opts ...option.RequestOption) *pagination.ListObjectsAutoPager[shared.Function] {
	return pagination.NewListObjectsAutoPager(r.List(ctx, query, opts...))
}

// Delete a function object by its id
func (r *FunctionService) Delete(ctx context.Context, functionID string, opts ...option.RequestOption) (res *shared.Function, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required function_id parameter")
		return
	}
	path := fmt.Sprintf("v1/function/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Invoke a function.
func (r *FunctionService) Invoke(ctx context.Context, functionID string, body FunctionInvokeParams, opts ...option.RequestOption) (res *FunctionInvokeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required function_id parameter")
		return
	}
	path := fmt.Sprintf("v1/function/%s/invoke", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create or replace function. If there is an existing function in the project with
// the same slug as the one specified in the request, will replace the existing
// function with the provided fields
func (r *FunctionService) Replace(ctx context.Context, body FunctionReplaceParams, opts ...option.RequestOption) (res *shared.Function, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/function"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type FunctionInvokeResponse = any

type FunctionNewParams struct {
	FunctionData FunctionNewParamsFunctionDataUnion `json:"function_data,omitzero,required"`
	// Name of the prompt
	Name string `json:"name,required"`
	// Unique identifier for the project that the prompt belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// Unique identifier for the prompt
	Slug string `json:"slug,required"`
	// Textual description of the prompt
	Description param.Opt[string] `json:"description,omitzero"`
	// JSON schema for the function's parameters and return type
	FunctionSchema FunctionNewParamsFunctionSchema `json:"function_schema,omitzero"`
	// Any of "llm", "scorer", "task", "tool".
	FunctionType FunctionNewParamsFunctionType `json:"function_type,omitzero"`
	Origin       FunctionNewParamsOrigin       `json:"origin,omitzero"`
	// The prompt, model, and its parameters
	PromptData shared.PromptDataParam `json:"prompt_data,omitzero"`
	// A list of tags for the prompt
	Tags []string `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r FunctionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunctionNewParamsFunctionDataUnion struct {
	OfPrompt *FunctionNewParamsFunctionDataPrompt `json:",omitzero,inline"`
	OfCode   *FunctionNewParamsFunctionDataCode   `json:",omitzero,inline"`
	OfGlobal *FunctionNewParamsFunctionDataGlobal `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u FunctionNewParamsFunctionDataUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u FunctionNewParamsFunctionDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[FunctionNewParamsFunctionDataUnion](u.OfPrompt, u.OfCode, u.OfGlobal)
}

func (u *FunctionNewParamsFunctionDataUnion) asAny() any {
	if !param.IsOmitted(u.OfPrompt) {
		return u.OfPrompt
	} else if !param.IsOmitted(u.OfCode) {
		return u.OfCode
	} else if !param.IsOmitted(u.OfGlobal) {
		return u.OfGlobal
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionNewParamsFunctionDataUnion) GetData() *FunctionNewParamsFunctionDataCodeDataUnion {
	if vt := u.OfCode; vt != nil {
		return &vt.Data
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionNewParamsFunctionDataUnion) GetName() *string {
	if vt := u.OfGlobal; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionNewParamsFunctionDataUnion) GetType() *string {
	if vt := u.OfPrompt; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCode; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfGlobal; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The property Type is required.
type FunctionNewParamsFunctionDataPrompt struct {
	// Any of "prompt".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionNewParamsFunctionDataPrompt) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionNewParamsFunctionDataPrompt) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsFunctionDataPrompt
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionNewParamsFunctionDataPrompt](
		"Type", false, "prompt",
	)
}

// The properties Data, Type are required.
type FunctionNewParamsFunctionDataCode struct {
	Data FunctionNewParamsFunctionDataCodeDataUnion `json:"data,omitzero,required"`
	// Any of "code".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionNewParamsFunctionDataCode) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionNewParamsFunctionDataCode) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsFunctionDataCode
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionNewParamsFunctionDataCode](
		"Type", false, "code",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunctionNewParamsFunctionDataCodeDataUnion struct {
	OfBundle *FunctionNewParamsFunctionDataCodeDataBundle `json:",omitzero,inline"`
	OfInline *FunctionNewParamsFunctionDataCodeDataInline `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u FunctionNewParamsFunctionDataCodeDataUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u FunctionNewParamsFunctionDataCodeDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[FunctionNewParamsFunctionDataCodeDataUnion](u.OfBundle, u.OfInline)
}

func (u *FunctionNewParamsFunctionDataCodeDataUnion) asAny() any {
	if !param.IsOmitted(u.OfBundle) {
		return u.OfBundle
	} else if !param.IsOmitted(u.OfInline) {
		return u.OfInline
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionNewParamsFunctionDataCodeDataUnion) GetBundleID() *string {
	if vt := u.OfBundle; vt != nil {
		return &vt.BundleID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionNewParamsFunctionDataCodeDataUnion) GetLocation() *shared.CodeBundleLocationUnionParam {
	if vt := u.OfBundle; vt != nil {
		return &vt.Location
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionNewParamsFunctionDataCodeDataUnion) GetPreview() *string {
	if vt := u.OfBundle; vt != nil && vt.Preview.IsPresent() {
		return &vt.Preview.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionNewParamsFunctionDataCodeDataUnion) GetCode() *string {
	if vt := u.OfInline; vt != nil {
		return &vt.Code
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionNewParamsFunctionDataCodeDataUnion) GetType() *string {
	if vt := u.OfBundle; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInline; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u FunctionNewParamsFunctionDataCodeDataUnion) GetRuntimeContext() (res functionNewParamsFunctionDataCodeDataUnionRuntimeContext) {
	if vt := u.OfBundle; vt != nil {
		res.any = &vt.RuntimeContext
	} else if vt := u.OfInline; vt != nil {
		res.any = &vt.RuntimeContext
	}
	return
}

// Can have the runtime types [*shared.CodeBundleRuntimeContextParam],
// [*FunctionNewParamsFunctionDataCodeDataInlineRuntimeContext]
type functionNewParamsFunctionDataCodeDataUnionRuntimeContext struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *shared.CodeBundleRuntimeContextParam:
//	case *braintrust.FunctionNewParamsFunctionDataCodeDataInlineRuntimeContext:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u functionNewParamsFunctionDataCodeDataUnionRuntimeContext) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u functionNewParamsFunctionDataCodeDataUnionRuntimeContext) GetRuntime() *string {
	switch vt := u.any.(type) {
	case *shared.CodeBundleRuntimeContextParam:
		return (*string)(&vt.Runtime)
	case *FunctionNewParamsFunctionDataCodeDataInlineRuntimeContext:
		return (*string)(&vt.Runtime)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u functionNewParamsFunctionDataCodeDataUnionRuntimeContext) GetVersion() *string {
	switch vt := u.any.(type) {
	case *shared.CodeBundleRuntimeContextParam:
		return (*string)(&vt.Version)
	case *FunctionNewParamsFunctionDataCodeDataInlineRuntimeContext:
		return (*string)(&vt.Version)
	}
	return nil
}

type FunctionNewParamsFunctionDataCodeDataBundle struct {
	Type string `json:"type,omitzero,required"`
	shared.CodeBundleParam
}

func (r FunctionNewParamsFunctionDataCodeDataBundle) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsFunctionDataCodeDataBundle
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties Code, RuntimeContext, Type are required.
type FunctionNewParamsFunctionDataCodeDataInline struct {
	Code           string                                                    `json:"code,required"`
	RuntimeContext FunctionNewParamsFunctionDataCodeDataInlineRuntimeContext `json:"runtime_context,omitzero,required"`
	// Any of "inline".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionNewParamsFunctionDataCodeDataInline) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionNewParamsFunctionDataCodeDataInline) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsFunctionDataCodeDataInline
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionNewParamsFunctionDataCodeDataInline](
		"Type", false, "inline",
	)
}

// The properties Runtime, Version are required.
type FunctionNewParamsFunctionDataCodeDataInlineRuntimeContext struct {
	// Any of "node", "python".
	Runtime string `json:"runtime,omitzero,required"`
	Version string `json:"version,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionNewParamsFunctionDataCodeDataInlineRuntimeContext) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionNewParamsFunctionDataCodeDataInlineRuntimeContext) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsFunctionDataCodeDataInlineRuntimeContext
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionNewParamsFunctionDataCodeDataInlineRuntimeContext](
		"Runtime", false, "node", "python",
	)
}

// The properties Name, Type are required.
type FunctionNewParamsFunctionDataGlobal struct {
	Name string `json:"name,required"`
	// Any of "global".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionNewParamsFunctionDataGlobal) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionNewParamsFunctionDataGlobal) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsFunctionDataGlobal
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionNewParamsFunctionDataGlobal](
		"Type", false, "global",
	)
}

// JSON schema for the function's parameters and return type
type FunctionNewParamsFunctionSchema struct {
	Parameters any `json:"parameters,omitzero"`
	Returns    any `json:"returns,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionNewParamsFunctionSchema) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r FunctionNewParamsFunctionSchema) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsFunctionSchema
	return param.MarshalObject(r, (*shadow)(&r))
}

type FunctionNewParamsFunctionType string

const (
	FunctionNewParamsFunctionTypeLlm    FunctionNewParamsFunctionType = "llm"
	FunctionNewParamsFunctionTypeScorer FunctionNewParamsFunctionType = "scorer"
	FunctionNewParamsFunctionTypeTask   FunctionNewParamsFunctionType = "task"
	FunctionNewParamsFunctionTypeTool   FunctionNewParamsFunctionType = "tool"
)

// The properties ObjectID, ObjectType are required.
type FunctionNewParamsOrigin struct {
	// Id of the object the function is originating from
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType shared.ACLObjectType `json:"object_type,omitzero,required"`
	// The function exists for internal purposes and should not be displayed in the
	// list of functions.
	Internal param.Opt[bool] `json:"internal,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionNewParamsOrigin) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r FunctionNewParamsOrigin) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsOrigin
	return param.MarshalObject(r, (*shadow)(&r))
}

type FunctionUpdateParams struct {
	// Textual description of the prompt
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the prompt
	Name         param.Opt[string]                     `json:"name,omitzero"`
	FunctionData FunctionUpdateParamsFunctionDataUnion `json:"function_data,omitzero"`
	// The prompt, model, and its parameters
	PromptData shared.PromptDataParam `json:"prompt_data,omitzero"`
	// A list of tags for the prompt
	Tags []string `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r FunctionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunctionUpdateParamsFunctionDataUnion struct {
	OfPrompt *FunctionUpdateParamsFunctionDataPrompt `json:",omitzero,inline"`
	OfCode   *FunctionUpdateParamsFunctionDataCode   `json:",omitzero,inline"`
	OfGlobal *FunctionUpdateParamsFunctionDataGlobal `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u FunctionUpdateParamsFunctionDataUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u FunctionUpdateParamsFunctionDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[FunctionUpdateParamsFunctionDataUnion](u.OfPrompt, u.OfCode, u.OfGlobal)
}

func (u *FunctionUpdateParamsFunctionDataUnion) asAny() any {
	if !param.IsOmitted(u.OfPrompt) {
		return u.OfPrompt
	} else if !param.IsOmitted(u.OfCode) {
		return u.OfCode
	} else if !param.IsOmitted(u.OfGlobal) {
		return u.OfGlobal
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionUpdateParamsFunctionDataUnion) GetData() *FunctionUpdateParamsFunctionDataCodeDataUnion {
	if vt := u.OfCode; vt != nil {
		return &vt.Data
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionUpdateParamsFunctionDataUnion) GetName() *string {
	if vt := u.OfGlobal; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionUpdateParamsFunctionDataUnion) GetType() *string {
	if vt := u.OfPrompt; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCode; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfGlobal; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The property Type is required.
type FunctionUpdateParamsFunctionDataPrompt struct {
	// Any of "prompt".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionUpdateParamsFunctionDataPrompt) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionUpdateParamsFunctionDataPrompt) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsFunctionDataPrompt
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionUpdateParamsFunctionDataPrompt](
		"Type", false, "prompt",
	)
}

// The properties Data, Type are required.
type FunctionUpdateParamsFunctionDataCode struct {
	Data FunctionUpdateParamsFunctionDataCodeDataUnion `json:"data,omitzero,required"`
	// Any of "code".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionUpdateParamsFunctionDataCode) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionUpdateParamsFunctionDataCode) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsFunctionDataCode
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionUpdateParamsFunctionDataCode](
		"Type", false, "code",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunctionUpdateParamsFunctionDataCodeDataUnion struct {
	OfBundle *FunctionUpdateParamsFunctionDataCodeDataBundle `json:",omitzero,inline"`
	OfInline *FunctionUpdateParamsFunctionDataCodeDataInline `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u FunctionUpdateParamsFunctionDataCodeDataUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u FunctionUpdateParamsFunctionDataCodeDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[FunctionUpdateParamsFunctionDataCodeDataUnion](u.OfBundle, u.OfInline)
}

func (u *FunctionUpdateParamsFunctionDataCodeDataUnion) asAny() any {
	if !param.IsOmitted(u.OfBundle) {
		return u.OfBundle
	} else if !param.IsOmitted(u.OfInline) {
		return u.OfInline
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionUpdateParamsFunctionDataCodeDataUnion) GetBundleID() *string {
	if vt := u.OfBundle; vt != nil {
		return &vt.BundleID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionUpdateParamsFunctionDataCodeDataUnion) GetLocation() *shared.CodeBundleLocationUnionParam {
	if vt := u.OfBundle; vt != nil {
		return &vt.Location
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionUpdateParamsFunctionDataCodeDataUnion) GetPreview() *string {
	if vt := u.OfBundle; vt != nil && vt.Preview.IsPresent() {
		return &vt.Preview.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionUpdateParamsFunctionDataCodeDataUnion) GetCode() *string {
	if vt := u.OfInline; vt != nil {
		return &vt.Code
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionUpdateParamsFunctionDataCodeDataUnion) GetType() *string {
	if vt := u.OfBundle; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInline; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u FunctionUpdateParamsFunctionDataCodeDataUnion) GetRuntimeContext() (res functionUpdateParamsFunctionDataCodeDataUnionRuntimeContext) {
	if vt := u.OfBundle; vt != nil {
		res.any = &vt.RuntimeContext
	} else if vt := u.OfInline; vt != nil {
		res.any = &vt.RuntimeContext
	}
	return
}

// Can have the runtime types [*shared.CodeBundleRuntimeContextParam],
// [*FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContext]
type functionUpdateParamsFunctionDataCodeDataUnionRuntimeContext struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *shared.CodeBundleRuntimeContextParam:
//	case *braintrust.FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContext:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u functionUpdateParamsFunctionDataCodeDataUnionRuntimeContext) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u functionUpdateParamsFunctionDataCodeDataUnionRuntimeContext) GetRuntime() *string {
	switch vt := u.any.(type) {
	case *shared.CodeBundleRuntimeContextParam:
		return (*string)(&vt.Runtime)
	case *FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContext:
		return (*string)(&vt.Runtime)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u functionUpdateParamsFunctionDataCodeDataUnionRuntimeContext) GetVersion() *string {
	switch vt := u.any.(type) {
	case *shared.CodeBundleRuntimeContextParam:
		return (*string)(&vt.Version)
	case *FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContext:
		return (*string)(&vt.Version)
	}
	return nil
}

type FunctionUpdateParamsFunctionDataCodeDataBundle struct {
	Type string `json:"type,omitzero,required"`
	shared.CodeBundleParam
}

func (r FunctionUpdateParamsFunctionDataCodeDataBundle) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsFunctionDataCodeDataBundle
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties Code, RuntimeContext, Type are required.
type FunctionUpdateParamsFunctionDataCodeDataInline struct {
	Code           string                                                       `json:"code,required"`
	RuntimeContext FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContext `json:"runtime_context,omitzero,required"`
	// Any of "inline".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionUpdateParamsFunctionDataCodeDataInline) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionUpdateParamsFunctionDataCodeDataInline) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsFunctionDataCodeDataInline
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionUpdateParamsFunctionDataCodeDataInline](
		"Type", false, "inline",
	)
}

// The properties Runtime, Version are required.
type FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContext struct {
	// Any of "node", "python".
	Runtime string `json:"runtime,omitzero,required"`
	Version string `json:"version,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContext) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContext) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContext
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionUpdateParamsFunctionDataCodeDataInlineRuntimeContext](
		"Runtime", false, "node", "python",
	)
}

// The properties Name, Type are required.
type FunctionUpdateParamsFunctionDataGlobal struct {
	Name string `json:"name,required"`
	// Any of "global".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionUpdateParamsFunctionDataGlobal) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionUpdateParamsFunctionDataGlobal) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsFunctionDataGlobal
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionUpdateParamsFunctionDataGlobal](
		"Type", false, "global",
	)
}

type FunctionListParams struct {
	// Limit the number of objects to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the initial item in the last page you fetched had an id of
	// `foo`, pass `ending_before=foo` to fetch the previous page. Note: you may only
	// pass one of `starting_after` and `ending_before`
	EndingBefore param.Opt[string] `query:"ending_before,omitzero" format:"uuid" json:"-"`
	// Name of the function to search for
	FunctionName param.Opt[string] `query:"function_name,omitzero" json:"-"`
	// Filter search results to within a particular organization
	OrgName param.Opt[string] `query:"org_name,omitzero" json:"-"`
	// Project id
	ProjectID param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Name of the project to search for
	ProjectName param.Opt[string] `query:"project_name,omitzero" json:"-"`
	// Retrieve prompt with a specific slug
	Slug param.Opt[string] `query:"slug,omitzero" json:"-"`
	// Pagination cursor id.
	//
	// For example, if the final item in the last page you fetched had an id of `foo`,
	// pass `starting_after=foo` to fetch the next page. Note: you may only pass one of
	// `starting_after` and `ending_before`
	StartingAfter param.Opt[string] `query:"starting_after,omitzero" format:"uuid" json:"-"`
	// Retrieve prompt at a specific version.
	//
	// The version id can either be a transaction id (e.g. '1000192656880881099') or a
	// version identifier (e.g. '81cd05ee665fdfb3').
	Version param.Opt[string] `query:"version,omitzero" json:"-"`
	// Filter search results to a particular set of object IDs. To specify a list of
	// IDs, include the query param multiple times
	IDs FunctionListParamsIDsUnion `query:"ids,omitzero" format:"uuid" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [FunctionListParams]'s query parameters as `url.Values`.
func (r FunctionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunctionListParamsIDsUnion struct {
	OfString                param.Opt[string] `query:",omitzero,inline"`
	OfFunctionListsIDsArray []string          `query:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u FunctionListParamsIDsUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }

func (u *FunctionListParamsIDsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFunctionListsIDsArray) {
		return &u.OfFunctionListsIDsArray
	}
	return nil
}

type FunctionInvokeParams struct {
	// Whether to stream the response. If true, results will be returned in the
	// Braintrust SSE format.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// The version of the function
	Version param.Opt[string] `json:"version,omitzero"`
	// Any relevant metadata
	Metadata map[string]any `json:"metadata,omitzero"`
	// The mode format of the returned value (defaults to 'auto')
	//
	// Any of "auto", "parallel".
	Mode FunctionInvokeParamsMode `json:"mode,omitzero"`
	// The expected output of the function
	Expected any `json:"expected,omitzero"`
	// Argument to the function, which can be any JSON serializable value
	Input any `json:"input,omitzero"`
	// If the function is an LLM, additional messages to pass along to it
	Messages []FunctionInvokeParamsMessageUnion `json:"messages,omitzero"`
	// Options for tracing the function call
	Parent FunctionInvokeParamsParentUnion `json:"parent,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionInvokeParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r FunctionInvokeParams) MarshalJSON() (data []byte, err error) {
	type shadow FunctionInvokeParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunctionInvokeParamsMessageUnion struct {
	OfSystem    *FunctionInvokeParamsMessageSystem    `json:",omitzero,inline"`
	OfUser      *FunctionInvokeParamsMessageUser      `json:",omitzero,inline"`
	OfAssistant *FunctionInvokeParamsMessageAssistant `json:",omitzero,inline"`
	OfTool      *FunctionInvokeParamsMessageTool      `json:",omitzero,inline"`
	OfFunction  *FunctionInvokeParamsMessageFunction  `json:",omitzero,inline"`
	OfFallback  *FunctionInvokeParamsMessageFallback  `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u FunctionInvokeParamsMessageUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u FunctionInvokeParamsMessageUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[FunctionInvokeParamsMessageUnion](u.OfSystem,
		u.OfUser,
		u.OfAssistant,
		u.OfTool,
		u.OfFunction,
		u.OfFallback)
}

func (u *FunctionInvokeParamsMessageUnion) asAny() any {
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
func (u FunctionInvokeParamsMessageUnion) GetFunctionCall() *FunctionInvokeParamsMessageAssistantFunctionCall {
	if vt := u.OfAssistant; vt != nil {
		return &vt.FunctionCall
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionInvokeParamsMessageUnion) GetToolCalls() []shared.ChatCompletionMessageToolCallParam {
	if vt := u.OfAssistant; vt != nil {
		return vt.ToolCalls
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionInvokeParamsMessageUnion) GetToolCallID() *string {
	if vt := u.OfTool; vt != nil && vt.ToolCallID.IsPresent() {
		return &vt.ToolCallID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionInvokeParamsMessageUnion) GetRole() *string {
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
func (u FunctionInvokeParamsMessageUnion) GetName() *string {
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
func (u FunctionInvokeParamsMessageUnion) GetContent() (res functionInvokeParamsMessageUnionContent) {
	if vt := u.OfSystem; vt != nil && vt.Content.IsPresent() {
		res.any = &vt.Content.Value
	} else if vt := u.OfUser; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfAssistant; vt != nil && vt.Content.IsPresent() {
		res.any = &vt.Content.Value
	} else if vt := u.OfTool; vt != nil && vt.Content.IsPresent() {
		res.any = &vt.Content.Value
	} else if vt := u.OfFunction; vt != nil && vt.Content.IsPresent() {
		res.any = &vt.Content.Value
	} else if vt := u.OfFallback; vt != nil && vt.Content.IsPresent() {
		res.any = &vt.Content.Value
	}
	return
}

// Can have the runtime types [*string],
// [\*[]FunctionInvokeParamsMessageUserContentArrayItemUnion]
type functionInvokeParamsMessageUnionContent struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]braintrust.FunctionInvokeParamsMessageUserContentArrayItemUnion:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u functionInvokeParamsMessageUnionContent) AsAny() any { return u.any }

// The property Role is required.
type FunctionInvokeParamsMessageSystem struct {
	// Any of "system".
	Role    string            `json:"role,omitzero,required"`
	Content param.Opt[string] `json:"content,omitzero"`
	Name    param.Opt[string] `json:"name,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionInvokeParamsMessageSystem) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionInvokeParamsMessageSystem) MarshalJSON() (data []byte, err error) {
	type shadow FunctionInvokeParamsMessageSystem
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionInvokeParamsMessageSystem](
		"Role", false, "system",
	)
}

// The property Role is required.
type FunctionInvokeParamsMessageUser struct {
	// Any of "user".
	Role    string                                      `json:"role,omitzero,required"`
	Name    param.Opt[string]                           `json:"name,omitzero"`
	Content FunctionInvokeParamsMessageUserContentUnion `json:"content,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionInvokeParamsMessageUser) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r FunctionInvokeParamsMessageUser) MarshalJSON() (data []byte, err error) {
	type shadow FunctionInvokeParamsMessageUser
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionInvokeParamsMessageUser](
		"Role", false, "user",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunctionInvokeParamsMessageUserContentUnion struct {
	OfString param.Opt[string]                                      `json:",omitzero,inline"`
	OfArray  []FunctionInvokeParamsMessageUserContentArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u FunctionInvokeParamsMessageUserContentUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u FunctionInvokeParamsMessageUserContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[FunctionInvokeParamsMessageUserContentUnion](u.OfString, u.OfArray)
}

func (u *FunctionInvokeParamsMessageUserContentUnion) asAny() any {
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
type FunctionInvokeParamsMessageUserContentArrayItemUnion struct {
	OfText     *shared.ChatCompletionContentPartTextParam  `json:",omitzero,inline"`
	OfImageURL *shared.ChatCompletionContentPartImageParam `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u FunctionInvokeParamsMessageUserContentArrayItemUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u FunctionInvokeParamsMessageUserContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[FunctionInvokeParamsMessageUserContentArrayItemUnion](u.OfText, u.OfImageURL)
}

func (u *FunctionInvokeParamsMessageUserContentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfImageURL) {
		return u.OfImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionInvokeParamsMessageUserContentArrayItemUnion) GetText() *string {
	if vt := u.OfText; vt != nil && vt.Text.IsPresent() {
		return &vt.Text.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionInvokeParamsMessageUserContentArrayItemUnion) GetImageURL() *shared.ChatCompletionContentPartImageImageURLParam {
	if vt := u.OfImageURL; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionInvokeParamsMessageUserContentArrayItemUnion) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImageURL; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The property Role is required.
type FunctionInvokeParamsMessageAssistant struct {
	// Any of "assistant".
	Role         string                                           `json:"role,omitzero,required"`
	Content      param.Opt[string]                                `json:"content,omitzero"`
	Name         param.Opt[string]                                `json:"name,omitzero"`
	FunctionCall FunctionInvokeParamsMessageAssistantFunctionCall `json:"function_call,omitzero"`
	ToolCalls    []shared.ChatCompletionMessageToolCallParam      `json:"tool_calls,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionInvokeParamsMessageAssistant) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionInvokeParamsMessageAssistant) MarshalJSON() (data []byte, err error) {
	type shadow FunctionInvokeParamsMessageAssistant
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionInvokeParamsMessageAssistant](
		"Role", false, "assistant",
	)
}

// The properties Arguments, Name are required.
type FunctionInvokeParamsMessageAssistantFunctionCall struct {
	Arguments string `json:"arguments,required"`
	Name      string `json:"name,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionInvokeParamsMessageAssistantFunctionCall) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionInvokeParamsMessageAssistantFunctionCall) MarshalJSON() (data []byte, err error) {
	type shadow FunctionInvokeParamsMessageAssistantFunctionCall
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property Role is required.
type FunctionInvokeParamsMessageTool struct {
	// Any of "tool".
	Role       string            `json:"role,omitzero,required"`
	Content    param.Opt[string] `json:"content,omitzero"`
	ToolCallID param.Opt[string] `json:"tool_call_id,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionInvokeParamsMessageTool) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r FunctionInvokeParamsMessageTool) MarshalJSON() (data []byte, err error) {
	type shadow FunctionInvokeParamsMessageTool
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionInvokeParamsMessageTool](
		"Role", false, "tool",
	)
}

// The properties Name, Role are required.
type FunctionInvokeParamsMessageFunction struct {
	Name string `json:"name,required"`
	// Any of "function".
	Role    string            `json:"role,omitzero,required"`
	Content param.Opt[string] `json:"content,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionInvokeParamsMessageFunction) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionInvokeParamsMessageFunction) MarshalJSON() (data []byte, err error) {
	type shadow FunctionInvokeParamsMessageFunction
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionInvokeParamsMessageFunction](
		"Role", false, "function",
	)
}

// The property Role is required.
type FunctionInvokeParamsMessageFallback struct {
	// Any of "model".
	Role    string            `json:"role,omitzero,required"`
	Content param.Opt[string] `json:"content,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionInvokeParamsMessageFallback) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionInvokeParamsMessageFallback) MarshalJSON() (data []byte, err error) {
	type shadow FunctionInvokeParamsMessageFallback
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionInvokeParamsMessageFallback](
		"Role", false, "model",
	)
}

// The mode format of the returned value (defaults to 'auto')
type FunctionInvokeParamsMode string

const (
	FunctionInvokeParamsModeAuto     FunctionInvokeParamsMode = "auto"
	FunctionInvokeParamsModeParallel FunctionInvokeParamsMode = "parallel"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunctionInvokeParamsParentUnion struct {
	OfSpanParentStruct *FunctionInvokeParamsParentSpanParentStruct `json:",omitzero,inline"`
	OfString           param.Opt[string]                           `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u FunctionInvokeParamsParentUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u FunctionInvokeParamsParentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[FunctionInvokeParamsParentUnion](u.OfSpanParentStruct, u.OfString)
}

func (u *FunctionInvokeParamsParentUnion) asAny() any {
	if !param.IsOmitted(u.OfSpanParentStruct) {
		return u.OfSpanParentStruct
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Span parent properties
//
// The properties ObjectID, ObjectType are required.
type FunctionInvokeParamsParentSpanParentStruct struct {
	// The id of the container object you are logging to
	ObjectID string `json:"object_id,required"`
	// Any of "project_logs", "experiment", "playground_logs".
	ObjectType string `json:"object_type,omitzero,required"`
	// Include these properties in every span created under this parent
	PropagatedEvent map[string]any `json:"propagated_event,omitzero"`
	// Identifiers for the row to to log a subspan under
	RowIDs FunctionInvokeParamsParentSpanParentStructRowIDs `json:"row_ids,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionInvokeParamsParentSpanParentStruct) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionInvokeParamsParentSpanParentStruct) MarshalJSON() (data []byte, err error) {
	type shadow FunctionInvokeParamsParentSpanParentStruct
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionInvokeParamsParentSpanParentStruct](
		"ObjectType", false, "project_logs", "experiment", "playground_logs",
	)
}

// Identifiers for the row to to log a subspan under
//
// The properties ID, RootSpanID, SpanID are required.
type FunctionInvokeParamsParentSpanParentStructRowIDs struct {
	// The id of the row
	ID string `json:"id,required"`
	// The root_span_id of the row
	RootSpanID string `json:"root_span_id,required"`
	// The span_id of the row
	SpanID string `json:"span_id,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionInvokeParamsParentSpanParentStructRowIDs) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionInvokeParamsParentSpanParentStructRowIDs) MarshalJSON() (data []byte, err error) {
	type shadow FunctionInvokeParamsParentSpanParentStructRowIDs
	return param.MarshalObject(r, (*shadow)(&r))
}

type FunctionReplaceParams struct {
	FunctionData FunctionReplaceParamsFunctionDataUnion `json:"function_data,omitzero,required"`
	// Name of the prompt
	Name string `json:"name,required"`
	// Unique identifier for the project that the prompt belongs under
	ProjectID string `json:"project_id,required" format:"uuid"`
	// Unique identifier for the prompt
	Slug string `json:"slug,required"`
	// Textual description of the prompt
	Description param.Opt[string] `json:"description,omitzero"`
	// JSON schema for the function's parameters and return type
	FunctionSchema FunctionReplaceParamsFunctionSchema `json:"function_schema,omitzero"`
	// Any of "llm", "scorer", "task", "tool".
	FunctionType FunctionReplaceParamsFunctionType `json:"function_type,omitzero"`
	Origin       FunctionReplaceParamsOrigin       `json:"origin,omitzero"`
	// The prompt, model, and its parameters
	PromptData shared.PromptDataParam `json:"prompt_data,omitzero"`
	// A list of tags for the prompt
	Tags []string `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionReplaceParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r FunctionReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow FunctionReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunctionReplaceParamsFunctionDataUnion struct {
	OfPrompt *FunctionReplaceParamsFunctionDataPrompt `json:",omitzero,inline"`
	OfCode   *FunctionReplaceParamsFunctionDataCode   `json:",omitzero,inline"`
	OfGlobal *FunctionReplaceParamsFunctionDataGlobal `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u FunctionReplaceParamsFunctionDataUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u FunctionReplaceParamsFunctionDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[FunctionReplaceParamsFunctionDataUnion](u.OfPrompt, u.OfCode, u.OfGlobal)
}

func (u *FunctionReplaceParamsFunctionDataUnion) asAny() any {
	if !param.IsOmitted(u.OfPrompt) {
		return u.OfPrompt
	} else if !param.IsOmitted(u.OfCode) {
		return u.OfCode
	} else if !param.IsOmitted(u.OfGlobal) {
		return u.OfGlobal
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionReplaceParamsFunctionDataUnion) GetData() *FunctionReplaceParamsFunctionDataCodeDataUnion {
	if vt := u.OfCode; vt != nil {
		return &vt.Data
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionReplaceParamsFunctionDataUnion) GetName() *string {
	if vt := u.OfGlobal; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionReplaceParamsFunctionDataUnion) GetType() *string {
	if vt := u.OfPrompt; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCode; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfGlobal; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The property Type is required.
type FunctionReplaceParamsFunctionDataPrompt struct {
	// Any of "prompt".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionReplaceParamsFunctionDataPrompt) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionReplaceParamsFunctionDataPrompt) MarshalJSON() (data []byte, err error) {
	type shadow FunctionReplaceParamsFunctionDataPrompt
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionReplaceParamsFunctionDataPrompt](
		"Type", false, "prompt",
	)
}

// The properties Data, Type are required.
type FunctionReplaceParamsFunctionDataCode struct {
	Data FunctionReplaceParamsFunctionDataCodeDataUnion `json:"data,omitzero,required"`
	// Any of "code".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionReplaceParamsFunctionDataCode) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionReplaceParamsFunctionDataCode) MarshalJSON() (data []byte, err error) {
	type shadow FunctionReplaceParamsFunctionDataCode
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionReplaceParamsFunctionDataCode](
		"Type", false, "code",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunctionReplaceParamsFunctionDataCodeDataUnion struct {
	OfBundle *FunctionReplaceParamsFunctionDataCodeDataBundle `json:",omitzero,inline"`
	OfInline *FunctionReplaceParamsFunctionDataCodeDataInline `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u FunctionReplaceParamsFunctionDataCodeDataUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u FunctionReplaceParamsFunctionDataCodeDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[FunctionReplaceParamsFunctionDataCodeDataUnion](u.OfBundle, u.OfInline)
}

func (u *FunctionReplaceParamsFunctionDataCodeDataUnion) asAny() any {
	if !param.IsOmitted(u.OfBundle) {
		return u.OfBundle
	} else if !param.IsOmitted(u.OfInline) {
		return u.OfInline
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionReplaceParamsFunctionDataCodeDataUnion) GetBundleID() *string {
	if vt := u.OfBundle; vt != nil {
		return &vt.BundleID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionReplaceParamsFunctionDataCodeDataUnion) GetLocation() *shared.CodeBundleLocationUnionParam {
	if vt := u.OfBundle; vt != nil {
		return &vt.Location
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionReplaceParamsFunctionDataCodeDataUnion) GetPreview() *string {
	if vt := u.OfBundle; vt != nil && vt.Preview.IsPresent() {
		return &vt.Preview.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionReplaceParamsFunctionDataCodeDataUnion) GetCode() *string {
	if vt := u.OfInline; vt != nil {
		return &vt.Code
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FunctionReplaceParamsFunctionDataCodeDataUnion) GetType() *string {
	if vt := u.OfBundle; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInline; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u FunctionReplaceParamsFunctionDataCodeDataUnion) GetRuntimeContext() (res functionReplaceParamsFunctionDataCodeDataUnionRuntimeContext) {
	if vt := u.OfBundle; vt != nil {
		res.any = &vt.RuntimeContext
	} else if vt := u.OfInline; vt != nil {
		res.any = &vt.RuntimeContext
	}
	return
}

// Can have the runtime types [*shared.CodeBundleRuntimeContextParam],
// [*FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContext]
type functionReplaceParamsFunctionDataCodeDataUnionRuntimeContext struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *shared.CodeBundleRuntimeContextParam:
//	case *braintrust.FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContext:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u functionReplaceParamsFunctionDataCodeDataUnionRuntimeContext) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u functionReplaceParamsFunctionDataCodeDataUnionRuntimeContext) GetRuntime() *string {
	switch vt := u.any.(type) {
	case *shared.CodeBundleRuntimeContextParam:
		return (*string)(&vt.Runtime)
	case *FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContext:
		return (*string)(&vt.Runtime)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u functionReplaceParamsFunctionDataCodeDataUnionRuntimeContext) GetVersion() *string {
	switch vt := u.any.(type) {
	case *shared.CodeBundleRuntimeContextParam:
		return (*string)(&vt.Version)
	case *FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContext:
		return (*string)(&vt.Version)
	}
	return nil
}

type FunctionReplaceParamsFunctionDataCodeDataBundle struct {
	Type string `json:"type,omitzero,required"`
	shared.CodeBundleParam
}

func (r FunctionReplaceParamsFunctionDataCodeDataBundle) MarshalJSON() (data []byte, err error) {
	type shadow FunctionReplaceParamsFunctionDataCodeDataBundle
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties Code, RuntimeContext, Type are required.
type FunctionReplaceParamsFunctionDataCodeDataInline struct {
	Code           string                                                        `json:"code,required"`
	RuntimeContext FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContext `json:"runtime_context,omitzero,required"`
	// Any of "inline".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionReplaceParamsFunctionDataCodeDataInline) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionReplaceParamsFunctionDataCodeDataInline) MarshalJSON() (data []byte, err error) {
	type shadow FunctionReplaceParamsFunctionDataCodeDataInline
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionReplaceParamsFunctionDataCodeDataInline](
		"Type", false, "inline",
	)
}

// The properties Runtime, Version are required.
type FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContext struct {
	// Any of "node", "python".
	Runtime string `json:"runtime,omitzero,required"`
	Version string `json:"version,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContext) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContext) MarshalJSON() (data []byte, err error) {
	type shadow FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContext
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionReplaceParamsFunctionDataCodeDataInlineRuntimeContext](
		"Runtime", false, "node", "python",
	)
}

// The properties Name, Type are required.
type FunctionReplaceParamsFunctionDataGlobal struct {
	Name string `json:"name,required"`
	// Any of "global".
	Type string `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionReplaceParamsFunctionDataGlobal) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionReplaceParamsFunctionDataGlobal) MarshalJSON() (data []byte, err error) {
	type shadow FunctionReplaceParamsFunctionDataGlobal
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FunctionReplaceParamsFunctionDataGlobal](
		"Type", false, "global",
	)
}

// JSON schema for the function's parameters and return type
type FunctionReplaceParamsFunctionSchema struct {
	Parameters any `json:"parameters,omitzero"`
	Returns    any `json:"returns,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionReplaceParamsFunctionSchema) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FunctionReplaceParamsFunctionSchema) MarshalJSON() (data []byte, err error) {
	type shadow FunctionReplaceParamsFunctionSchema
	return param.MarshalObject(r, (*shadow)(&r))
}

type FunctionReplaceParamsFunctionType string

const (
	FunctionReplaceParamsFunctionTypeLlm    FunctionReplaceParamsFunctionType = "llm"
	FunctionReplaceParamsFunctionTypeScorer FunctionReplaceParamsFunctionType = "scorer"
	FunctionReplaceParamsFunctionTypeTask   FunctionReplaceParamsFunctionType = "task"
	FunctionReplaceParamsFunctionTypeTool   FunctionReplaceParamsFunctionType = "tool"
)

// The properties ObjectID, ObjectType are required.
type FunctionReplaceParamsOrigin struct {
	// Id of the object the function is originating from
	ObjectID string `json:"object_id,required" format:"uuid"`
	// The object type that the ACL applies to
	//
	// Any of "organization", "project", "experiment", "dataset", "prompt",
	// "prompt_session", "group", "role", "org_member", "project_log", "org_project".
	ObjectType shared.ACLObjectType `json:"object_type,omitzero,required"`
	// The function exists for internal purposes and should not be displayed in the
	// list of functions.
	Internal param.Opt[bool] `json:"internal,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FunctionReplaceParamsOrigin) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r FunctionReplaceParamsOrigin) MarshalJSON() (data []byte, err error) {
	type shadow FunctionReplaceParamsOrigin
	return param.MarshalObject(r, (*shadow)(&r))
}
