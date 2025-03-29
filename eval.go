// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"net/http"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/braintrustdata/braintrust-go/packages/param"
	"github.com/braintrustdata/braintrust-go/shared"
)

// EvalService contains methods and other services that help with interacting with
// the braintrust API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvalService] method instead.
type EvalService struct {
	Options []option.RequestOption
}

// NewEvalService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEvalService(opts ...option.RequestOption) (r EvalService) {
	r = EvalService{}
	r.Options = opts
	return
}

// Launch an evaluation. This is the API-equivalent of the `Eval` function that is
// built into the Braintrust SDK. In the Eval API, you provide pointers to a
// dataset, task function, and scoring functions. The API will then run the
// evaluation, create an experiment, and return the results along with a link to
// the experiment. To learn more about evals, see the
// [Evals guide](https://www.braintrust.dev/docs/guides/evals).
func (r *EvalService) New(ctx context.Context, body EvalNewParams, opts ...option.RequestOption) (res *shared.SummarizeExperimentResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/eval"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvalNewParams struct {
	// The dataset to use
	Data EvalNewParamsDataUnion `json:"data,omitzero,required"`
	// Unique identifier for the project to run the eval in
	ProjectID string `json:"project_id,required"`
	// The functions to score the eval on
	Scores []EvalNewParamsScoreUnion `json:"scores,omitzero,required"`
	// The function to evaluate
	Task EvalNewParamsTaskUnion `json:"task,omitzero,required"`
	// An optional experiment id to use as a base. If specified, the new experiment
	// will be summarized and compared to this experiment.
	BaseExperimentID param.Opt[string] `json:"base_experiment_id,omitzero"`
	// An optional experiment name to use as a base. If specified, the new experiment
	// will be summarized and compared to this experiment.
	BaseExperimentName param.Opt[string] `json:"base_experiment_name,omitzero"`
	// Whether the experiment should be public. Defaults to false.
	IsPublic param.Opt[bool] `json:"is_public,omitzero"`
	// The maximum number of tasks/scorers that will be run concurrently. Defaults to
	// undefined, in which case there is no max concurrency.
	MaxConcurrency param.Opt[float64] `json:"max_concurrency,omitzero"`
	// The maximum duration, in milliseconds, to run the evaluation. Defaults to
	// undefined, in which case there is no timeout.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// The number of times to run the evaluator per input. This is useful for
	// evaluating applications that have non-deterministic behavior and gives you both
	// a stronger aggregate measure and a sense of the variance in the results.
	TrialCount param.Opt[float64] `json:"trial_count,omitzero"`
	// An optional name for the experiment created by this eval. If it conflicts with
	// an existing experiment, it will be suffixed with a unique identifier.
	ExperimentName param.Opt[string] `json:"experiment_name,omitzero"`
	// Whether to stream the results of the eval. If true, the request will return two
	// events: one to indicate the experiment has started, and another upon completion.
	// If false, the request will return the evaluation's summary upon completion.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// Optional settings for collecting git metadata. By default, will collect all git
	// metadata fields allowed in org-level settings.
	GitMetadataSettings EvalNewParamsGitMetadataSettings `json:"git_metadata_settings,omitzero"`
	// Metadata about the state of the repo when the experiment was created
	RepoInfo shared.RepoInfoParam `json:"repo_info,omitzero"`
	// Optional experiment-level metadata to store about the evaluation. You can later
	// use this to slice & dice across experiments.
	Metadata map[string]interface{} `json:"metadata,omitzero"`
	// Options for tracing the evaluation
	Parent EvalNewParamsParentUnion `json:"parent,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r EvalNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalNewParamsDataUnion struct {
	OfDatasetID          *EvalNewParamsDataDatasetID          `json:",omitzero,inline"`
	OfProjectDatasetName *EvalNewParamsDataProjectDatasetName `json:",omitzero,inline"`
	OfDatasetRows        *EvalNewParamsDataDatasetRows        `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u EvalNewParamsDataUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u EvalNewParamsDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[EvalNewParamsDataUnion](u.OfDatasetID, u.OfProjectDatasetName, u.OfDatasetRows)
}

func (u *EvalNewParamsDataUnion) asAny() any {
	if !param.IsOmitted(u.OfDatasetID) {
		return u.OfDatasetID
	} else if !param.IsOmitted(u.OfProjectDatasetName) {
		return u.OfProjectDatasetName
	} else if !param.IsOmitted(u.OfDatasetRows) {
		return u.OfDatasetRows
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsDataUnion) GetDatasetID() *string {
	if vt := u.OfDatasetID; vt != nil {
		return &vt.DatasetID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsDataUnion) GetDatasetName() *string {
	if vt := u.OfProjectDatasetName; vt != nil {
		return &vt.DatasetName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsDataUnion) GetProjectName() *string {
	if vt := u.OfProjectDatasetName; vt != nil {
		return &vt.ProjectName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsDataUnion) GetData() []interface{} {
	if vt := u.OfDatasetRows; vt != nil {
		return vt.Data
	}
	return nil
}

// Returns a pointer to the underlying variant's InternalBtql property, if present.
func (u EvalNewParamsDataUnion) GetInternalBtql() map[string]interface{} {
	if vt := u.OfDatasetID; vt != nil {
		return vt.InternalBtql
	} else if vt := u.OfProjectDatasetName; vt != nil {
		return vt.InternalBtql
	}
	return nil
}

// Dataset id
//
// The property DatasetID is required.
type EvalNewParamsDataDatasetID struct {
	DatasetID    string                 `json:"dataset_id,required"`
	InternalBtql map[string]interface{} `json:"_internal_btql,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsDataDatasetID) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvalNewParamsDataDatasetID) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsDataDatasetID
	return param.MarshalObject(r, (*shadow)(&r))
}

// Project and dataset name
//
// The properties DatasetName, ProjectName are required.
type EvalNewParamsDataProjectDatasetName struct {
	DatasetName  string                 `json:"dataset_name,required"`
	ProjectName  string                 `json:"project_name,required"`
	InternalBtql map[string]interface{} `json:"_internal_btql,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsDataProjectDatasetName) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvalNewParamsDataProjectDatasetName) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsDataProjectDatasetName
	return param.MarshalObject(r, (*shadow)(&r))
}

// Dataset rows
//
// The property Data is required.
type EvalNewParamsDataDatasetRows struct {
	Data []interface{} `json:"data,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsDataDatasetRows) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvalNewParamsDataDatasetRows) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsDataDatasetRows
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalNewParamsScoreUnion struct {
	OfFunctionID      *EvalNewParamsScoreFunctionID      `json:",omitzero,inline"`
	OfProjectSlug     *EvalNewParamsScoreProjectSlug     `json:",omitzero,inline"`
	OfGlobalFunction  *EvalNewParamsScoreGlobalFunction  `json:",omitzero,inline"`
	OfPromptSessionID *EvalNewParamsScorePromptSessionID `json:",omitzero,inline"`
	OfInlineCode      *EvalNewParamsScoreInlineCode      `json:",omitzero,inline"`
	OfInlinePrompt    *EvalNewParamsScoreInlinePrompt    `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u EvalNewParamsScoreUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u EvalNewParamsScoreUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[EvalNewParamsScoreUnion](u.OfFunctionID,
		u.OfProjectSlug,
		u.OfGlobalFunction,
		u.OfPromptSessionID,
		u.OfInlineCode,
		u.OfInlinePrompt)
}

func (u *EvalNewParamsScoreUnion) asAny() any {
	if !param.IsOmitted(u.OfFunctionID) {
		return u.OfFunctionID
	} else if !param.IsOmitted(u.OfProjectSlug) {
		return u.OfProjectSlug
	} else if !param.IsOmitted(u.OfGlobalFunction) {
		return u.OfGlobalFunction
	} else if !param.IsOmitted(u.OfPromptSessionID) {
		return u.OfPromptSessionID
	} else if !param.IsOmitted(u.OfInlineCode) {
		return u.OfInlineCode
	} else if !param.IsOmitted(u.OfInlinePrompt) {
		return u.OfInlinePrompt
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsScoreUnion) GetFunctionID() *string {
	if vt := u.OfFunctionID; vt != nil {
		return &vt.FunctionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsScoreUnion) GetProjectName() *string {
	if vt := u.OfProjectSlug; vt != nil {
		return &vt.ProjectName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsScoreUnion) GetSlug() *string {
	if vt := u.OfProjectSlug; vt != nil {
		return &vt.Slug
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsScoreUnion) GetGlobalFunction() *string {
	if vt := u.OfGlobalFunction; vt != nil {
		return &vt.GlobalFunction
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsScoreUnion) GetPromptSessionFunctionID() *string {
	if vt := u.OfPromptSessionID; vt != nil {
		return &vt.PromptSessionFunctionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsScoreUnion) GetPromptSessionID() *string {
	if vt := u.OfPromptSessionID; vt != nil {
		return &vt.PromptSessionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsScoreUnion) GetCode() *string {
	if vt := u.OfInlineCode; vt != nil {
		return &vt.Code
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsScoreUnion) GetInlineContext() *EvalNewParamsScoreInlineCodeInlineContext {
	if vt := u.OfInlineCode; vt != nil {
		return &vt.InlineContext
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsScoreUnion) GetInlinePrompt() *shared.PromptDataParam {
	if vt := u.OfInlinePrompt; vt != nil {
		return &vt.InlinePrompt
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsScoreUnion) GetVersion() *string {
	if vt := u.OfFunctionID; vt != nil && vt.Version.IsPresent() {
		return &vt.Version.Value
	} else if vt := u.OfProjectSlug; vt != nil && vt.Version.IsPresent() {
		return &vt.Version.Value
	} else if vt := u.OfPromptSessionID; vt != nil && vt.Version.IsPresent() {
		return &vt.Version.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsScoreUnion) GetName() *string {
	if vt := u.OfInlineCode; vt != nil && vt.Name.IsPresent() {
		return &vt.Name.Value
	} else if vt := u.OfInlinePrompt; vt != nil && vt.Name.IsPresent() {
		return &vt.Name.Value
	}
	return nil
}

// Function id
//
// The property FunctionID is required.
type EvalNewParamsScoreFunctionID struct {
	// The ID of the function
	FunctionID string `json:"function_id,required"`
	// The version of the function
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsScoreFunctionID) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvalNewParamsScoreFunctionID) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsScoreFunctionID
	return param.MarshalObject(r, (*shadow)(&r))
}

// Project name and slug
//
// The properties ProjectName, Slug are required.
type EvalNewParamsScoreProjectSlug struct {
	// The name of the project containing the function
	ProjectName string `json:"project_name,required"`
	// The slug of the function
	Slug string `json:"slug,required"`
	// The version of the function
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsScoreProjectSlug) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvalNewParamsScoreProjectSlug) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsScoreProjectSlug
	return param.MarshalObject(r, (*shadow)(&r))
}

// Global function name
//
// The property GlobalFunction is required.
type EvalNewParamsScoreGlobalFunction struct {
	// The name of the global function. Currently, the global namespace includes the
	// functions in autoevals
	GlobalFunction string `json:"global_function,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsScoreGlobalFunction) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvalNewParamsScoreGlobalFunction) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsScoreGlobalFunction
	return param.MarshalObject(r, (*shadow)(&r))
}

// Prompt session id
//
// The properties PromptSessionFunctionID, PromptSessionID are required.
type EvalNewParamsScorePromptSessionID struct {
	// The ID of the function in the prompt session
	PromptSessionFunctionID string `json:"prompt_session_function_id,required"`
	// The ID of the prompt session
	PromptSessionID string `json:"prompt_session_id,required"`
	// The version of the function
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsScorePromptSessionID) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvalNewParamsScorePromptSessionID) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsScorePromptSessionID
	return param.MarshalObject(r, (*shadow)(&r))
}

// Inline code function
//
// The properties Code, InlineContext are required.
type EvalNewParamsScoreInlineCode struct {
	// The inline code to execute
	Code          string                                    `json:"code,required"`
	InlineContext EvalNewParamsScoreInlineCodeInlineContext `json:"inline_context,omitzero,required"`
	// The name of the inline code function
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsScoreInlineCode) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvalNewParamsScoreInlineCode) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsScoreInlineCode
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties Runtime, Version are required.
type EvalNewParamsScoreInlineCodeInlineContext struct {
	// Any of "node", "python".
	Runtime string `json:"runtime,omitzero,required"`
	Version string `json:"version,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsScoreInlineCodeInlineContext) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvalNewParamsScoreInlineCodeInlineContext) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsScoreInlineCodeInlineContext
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[EvalNewParamsScoreInlineCodeInlineContext](
		"Runtime", false, "node", "python",
	)
}

// Inline prompt definition
//
// The property InlinePrompt is required.
type EvalNewParamsScoreInlinePrompt struct {
	// The prompt, model, and its parameters
	InlinePrompt shared.PromptDataParam `json:"inline_prompt,omitzero,required"`
	// The name of the inline prompt
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsScoreInlinePrompt) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvalNewParamsScoreInlinePrompt) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsScoreInlinePrompt
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalNewParamsTaskUnion struct {
	OfFunctionID      *EvalNewParamsTaskFunctionID      `json:",omitzero,inline"`
	OfProjectSlug     *EvalNewParamsTaskProjectSlug     `json:",omitzero,inline"`
	OfGlobalFunction  *EvalNewParamsTaskGlobalFunction  `json:",omitzero,inline"`
	OfPromptSessionID *EvalNewParamsTaskPromptSessionID `json:",omitzero,inline"`
	OfInlineCode      *EvalNewParamsTaskInlineCode      `json:",omitzero,inline"`
	OfInlinePrompt    *EvalNewParamsTaskInlinePrompt    `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u EvalNewParamsTaskUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u EvalNewParamsTaskUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[EvalNewParamsTaskUnion](u.OfFunctionID,
		u.OfProjectSlug,
		u.OfGlobalFunction,
		u.OfPromptSessionID,
		u.OfInlineCode,
		u.OfInlinePrompt)
}

func (u *EvalNewParamsTaskUnion) asAny() any {
	if !param.IsOmitted(u.OfFunctionID) {
		return u.OfFunctionID
	} else if !param.IsOmitted(u.OfProjectSlug) {
		return u.OfProjectSlug
	} else if !param.IsOmitted(u.OfGlobalFunction) {
		return u.OfGlobalFunction
	} else if !param.IsOmitted(u.OfPromptSessionID) {
		return u.OfPromptSessionID
	} else if !param.IsOmitted(u.OfInlineCode) {
		return u.OfInlineCode
	} else if !param.IsOmitted(u.OfInlinePrompt) {
		return u.OfInlinePrompt
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsTaskUnion) GetFunctionID() *string {
	if vt := u.OfFunctionID; vt != nil {
		return &vt.FunctionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsTaskUnion) GetProjectName() *string {
	if vt := u.OfProjectSlug; vt != nil {
		return &vt.ProjectName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsTaskUnion) GetSlug() *string {
	if vt := u.OfProjectSlug; vt != nil {
		return &vt.Slug
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsTaskUnion) GetGlobalFunction() *string {
	if vt := u.OfGlobalFunction; vt != nil {
		return &vt.GlobalFunction
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsTaskUnion) GetPromptSessionFunctionID() *string {
	if vt := u.OfPromptSessionID; vt != nil {
		return &vt.PromptSessionFunctionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsTaskUnion) GetPromptSessionID() *string {
	if vt := u.OfPromptSessionID; vt != nil {
		return &vt.PromptSessionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsTaskUnion) GetCode() *string {
	if vt := u.OfInlineCode; vt != nil {
		return &vt.Code
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsTaskUnion) GetInlineContext() *EvalNewParamsTaskInlineCodeInlineContext {
	if vt := u.OfInlineCode; vt != nil {
		return &vt.InlineContext
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsTaskUnion) GetInlinePrompt() *shared.PromptDataParam {
	if vt := u.OfInlinePrompt; vt != nil {
		return &vt.InlinePrompt
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsTaskUnion) GetVersion() *string {
	if vt := u.OfFunctionID; vt != nil && vt.Version.IsPresent() {
		return &vt.Version.Value
	} else if vt := u.OfProjectSlug; vt != nil && vt.Version.IsPresent() {
		return &vt.Version.Value
	} else if vt := u.OfPromptSessionID; vt != nil && vt.Version.IsPresent() {
		return &vt.Version.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsTaskUnion) GetName() *string {
	if vt := u.OfInlineCode; vt != nil && vt.Name.IsPresent() {
		return &vt.Name.Value
	} else if vt := u.OfInlinePrompt; vt != nil && vt.Name.IsPresent() {
		return &vt.Name.Value
	}
	return nil
}

// Function id
//
// The property FunctionID is required.
type EvalNewParamsTaskFunctionID struct {
	// The ID of the function
	FunctionID string `json:"function_id,required"`
	// The version of the function
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsTaskFunctionID) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvalNewParamsTaskFunctionID) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsTaskFunctionID
	return param.MarshalObject(r, (*shadow)(&r))
}

// Project name and slug
//
// The properties ProjectName, Slug are required.
type EvalNewParamsTaskProjectSlug struct {
	// The name of the project containing the function
	ProjectName string `json:"project_name,required"`
	// The slug of the function
	Slug string `json:"slug,required"`
	// The version of the function
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsTaskProjectSlug) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvalNewParamsTaskProjectSlug) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsTaskProjectSlug
	return param.MarshalObject(r, (*shadow)(&r))
}

// Global function name
//
// The property GlobalFunction is required.
type EvalNewParamsTaskGlobalFunction struct {
	// The name of the global function. Currently, the global namespace includes the
	// functions in autoevals
	GlobalFunction string `json:"global_function,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsTaskGlobalFunction) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvalNewParamsTaskGlobalFunction) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsTaskGlobalFunction
	return param.MarshalObject(r, (*shadow)(&r))
}

// Prompt session id
//
// The properties PromptSessionFunctionID, PromptSessionID are required.
type EvalNewParamsTaskPromptSessionID struct {
	// The ID of the function in the prompt session
	PromptSessionFunctionID string `json:"prompt_session_function_id,required"`
	// The ID of the prompt session
	PromptSessionID string `json:"prompt_session_id,required"`
	// The version of the function
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsTaskPromptSessionID) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvalNewParamsTaskPromptSessionID) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsTaskPromptSessionID
	return param.MarshalObject(r, (*shadow)(&r))
}

// Inline code function
//
// The properties Code, InlineContext are required.
type EvalNewParamsTaskInlineCode struct {
	// The inline code to execute
	Code          string                                   `json:"code,required"`
	InlineContext EvalNewParamsTaskInlineCodeInlineContext `json:"inline_context,omitzero,required"`
	// The name of the inline code function
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsTaskInlineCode) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvalNewParamsTaskInlineCode) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsTaskInlineCode
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties Runtime, Version are required.
type EvalNewParamsTaskInlineCodeInlineContext struct {
	// Any of "node", "python".
	Runtime string `json:"runtime,omitzero,required"`
	Version string `json:"version,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsTaskInlineCodeInlineContext) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvalNewParamsTaskInlineCodeInlineContext) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsTaskInlineCodeInlineContext
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[EvalNewParamsTaskInlineCodeInlineContext](
		"Runtime", false, "node", "python",
	)
}

// Inline prompt definition
//
// The property InlinePrompt is required.
type EvalNewParamsTaskInlinePrompt struct {
	// The prompt, model, and its parameters
	InlinePrompt shared.PromptDataParam `json:"inline_prompt,omitzero,required"`
	// The name of the inline prompt
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsTaskInlinePrompt) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvalNewParamsTaskInlinePrompt) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsTaskInlinePrompt
	return param.MarshalObject(r, (*shadow)(&r))
}

// Optional settings for collecting git metadata. By default, will collect all git
// metadata fields allowed in org-level settings.
//
// The property Collect is required.
type EvalNewParamsGitMetadataSettings struct {
	// Any of "all", "none", "some".
	Collect string   `json:"collect,omitzero,required"`
	Fields  []string `json:"fields,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsGitMetadataSettings) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EvalNewParamsGitMetadataSettings) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsGitMetadataSettings
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[EvalNewParamsGitMetadataSettings](
		"Collect", false, "all", "none", "some",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalNewParamsParentUnion struct {
	OfSpanParentStruct *EvalNewParamsParentSpanParentStruct `json:",omitzero,inline"`
	OfString           param.Opt[string]                    `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u EvalNewParamsParentUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u EvalNewParamsParentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[EvalNewParamsParentUnion](u.OfSpanParentStruct, u.OfString)
}

func (u *EvalNewParamsParentUnion) asAny() any {
	if !param.IsOmitted(u.OfSpanParentStruct) {
		return u.OfSpanParentStruct
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsParentUnion) GetObjectID() *string {
	if vt := u.OfSpanParentStruct; vt != nil {
		return &vt.ObjectID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsParentUnion) GetObjectType() *string {
	if vt := u.OfSpanParentStruct; vt != nil {
		return &vt.ObjectType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsParentUnion) GetPropagatedEvent() map[string]interface{} {
	if vt := u.OfSpanParentStruct; vt != nil {
		return vt.PropagatedEvent
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsParentUnion) GetRowIDs() *EvalNewParamsParentSpanParentStructRowIDs {
	if vt := u.OfSpanParentStruct; vt != nil {
		return &vt.RowIDs
	}
	return nil
}

// Span parent properties
//
// The properties ObjectID, ObjectType are required.
type EvalNewParamsParentSpanParentStruct struct {
	// The id of the container object you are logging to
	ObjectID string `json:"object_id,required"`
	// Any of "project_logs", "experiment", "playground_logs".
	ObjectType string `json:"object_type,omitzero,required"`
	// Include these properties in every span created under this parent
	PropagatedEvent map[string]interface{} `json:"propagated_event,omitzero"`
	// Identifiers for the row to to log a subspan under
	RowIDs EvalNewParamsParentSpanParentStructRowIDs `json:"row_ids,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EvalNewParamsParentSpanParentStruct) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvalNewParamsParentSpanParentStruct) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsParentSpanParentStruct
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[EvalNewParamsParentSpanParentStruct](
		"ObjectType", false, "project_logs", "experiment", "playground_logs",
	)
}

// Identifiers for the row to to log a subspan under
//
// The properties ID, RootSpanID, SpanID are required.
type EvalNewParamsParentSpanParentStructRowIDs struct {
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
func (f EvalNewParamsParentSpanParentStructRowIDs) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r EvalNewParamsParentSpanParentStructRowIDs) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsParentSpanParentStructRowIDs
	return param.MarshalObject(r, (*shadow)(&r))
}
