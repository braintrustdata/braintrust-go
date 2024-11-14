// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"net/http"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/internal/param"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
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
func NewEvalService(opts ...option.RequestOption) (r *EvalService) {
	r = &EvalService{}
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
	Data param.Field[EvalNewParamsDataUnion] `json:"data,required"`
	// Unique identifier for the project to run the eval in
	ProjectID param.Field[string] `json:"project_id,required"`
	// The functions to score the eval on
	Scores param.Field[[]EvalNewParamsScoreUnion] `json:"scores,required"`
	// The function to evaluate
	Task param.Field[EvalNewParamsTaskUnion] `json:"task,required"`
	// An optional experiment id to use as a base. If specified, the new experiment
	// will be summarized and compared to this experiment.
	BaseExperimentID param.Field[string] `json:"base_experiment_id"`
	// An optional experiment name to use as a base. If specified, the new experiment
	// will be summarized and compared to this experiment.
	BaseExperimentName param.Field[string] `json:"base_experiment_name"`
	// An optional name for the experiment created by this eval. If it conflicts with
	// an existing experiment, it will be suffixed with a unique identifier.
	ExperimentName param.Field[string] `json:"experiment_name"`
	// Optional settings for collecting git metadata. By default, will collect all git
	// metadata fields allowed in org-level settings.
	GitMetadataSettings param.Field[EvalNewParamsGitMetadataSettings] `json:"git_metadata_settings"`
	// Whether the experiment should be public. Defaults to false.
	IsPublic param.Field[bool] `json:"is_public"`
	// The maximum number of tasks/scorers that will be run concurrently. Defaults to
	// undefined, in which case there is no max concurrency.
	MaxConcurrency param.Field[float64] `json:"max_concurrency"`
	// Optional experiment-level metadata to store about the evaluation. You can later
	// use this to slice & dice across experiments.
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// Metadata about the state of the repo when the experiment was created
	RepoInfo param.Field[shared.RepoInfoParam] `json:"repo_info"`
	// Whether to stream the results of the eval. If true, the request will return two
	// events: one to indicate the experiment has started, and another upon completion.
	// If false, the request will return the evaluation's summary upon completion.
	Stream param.Field[bool] `json:"stream"`
	// The maximum duration, in milliseconds, to run the evaluation. Defaults to
	// undefined, in which case there is no timeout.
	Timeout param.Field[float64] `json:"timeout"`
	// The number of times to run the evaluator per input. This is useful for
	// evaluating applications that have non-deterministic behavior and gives you both
	// a stronger aggregate measure and a sense of the variance in the results.
	TrialCount param.Field[float64] `json:"trial_count"`
}

func (r EvalNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The dataset to use
type EvalNewParamsData struct {
	DatasetID   param.Field[string] `json:"dataset_id"`
	DatasetName param.Field[string] `json:"dataset_name"`
	ProjectName param.Field[string] `json:"project_name"`
}

func (r EvalNewParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsData) implementsEvalNewParamsDataUnion() {}

// The dataset to use
//
// Satisfied by [EvalNewParamsDataDatasetID],
// [EvalNewParamsDataProjectDatasetName], [EvalNewParamsData].
type EvalNewParamsDataUnion interface {
	implementsEvalNewParamsDataUnion()
}

// Dataset id
type EvalNewParamsDataDatasetID struct {
	DatasetID param.Field[string] `json:"dataset_id,required"`
}

func (r EvalNewParamsDataDatasetID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsDataDatasetID) implementsEvalNewParamsDataUnion() {}

// Project and dataset name
type EvalNewParamsDataProjectDatasetName struct {
	DatasetName param.Field[string] `json:"dataset_name,required"`
	ProjectName param.Field[string] `json:"project_name,required"`
}

func (r EvalNewParamsDataProjectDatasetName) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsDataProjectDatasetName) implementsEvalNewParamsDataUnion() {}

// The function to evaluate
type EvalNewParamsScore struct {
	// The inline code to execute
	Code param.Field[string] `json:"code"`
	// The ID of the function
	FunctionID param.Field[string] `json:"function_id"`
	// The name of the global function. Currently, the global namespace includes the
	// functions in autoevals
	GlobalFunction param.Field[string]      `json:"global_function"`
	InlineContext  param.Field[interface{}] `json:"inline_context"`
	// The prompt, model, and its parameters
	InlinePrompt param.Field[shared.PromptDataParam] `json:"inline_prompt"`
	// The name of the inline code function
	Name param.Field[string] `json:"name"`
	// The name of the project containing the function
	ProjectName param.Field[string] `json:"project_name"`
	// The ID of the function in the prompt session
	PromptSessionFunctionID param.Field[string] `json:"prompt_session_function_id"`
	// The ID of the prompt session
	PromptSessionID param.Field[string] `json:"prompt_session_id"`
	// The slug of the function
	Slug param.Field[string] `json:"slug"`
	// The version of the function
	Version param.Field[string] `json:"version"`
}

func (r EvalNewParamsScore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsScore) implementsEvalNewParamsScoreUnion() {}

// The function to evaluate
//
// Satisfied by [EvalNewParamsScoresFunctionID], [EvalNewParamsScoresProjectSlug],
// [EvalNewParamsScoresGlobalFunction], [EvalNewParamsScoresPromptSessionID],
// [EvalNewParamsScoresInlineCode], [EvalNewParamsScoresInlinePrompt],
// [EvalNewParamsScore].
type EvalNewParamsScoreUnion interface {
	implementsEvalNewParamsScoreUnion()
}

// Function id
type EvalNewParamsScoresFunctionID struct {
	// The ID of the function
	FunctionID param.Field[string] `json:"function_id,required"`
	// The version of the function
	Version param.Field[string] `json:"version"`
}

func (r EvalNewParamsScoresFunctionID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsScoresFunctionID) implementsEvalNewParamsScoreUnion() {}

// Project name and slug
type EvalNewParamsScoresProjectSlug struct {
	// The name of the project containing the function
	ProjectName param.Field[string] `json:"project_name,required"`
	// The slug of the function
	Slug param.Field[string] `json:"slug,required"`
	// The version of the function
	Version param.Field[string] `json:"version"`
}

func (r EvalNewParamsScoresProjectSlug) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsScoresProjectSlug) implementsEvalNewParamsScoreUnion() {}

// Global function name
type EvalNewParamsScoresGlobalFunction struct {
	// The name of the global function. Currently, the global namespace includes the
	// functions in autoevals
	GlobalFunction param.Field[string] `json:"global_function,required"`
}

func (r EvalNewParamsScoresGlobalFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsScoresGlobalFunction) implementsEvalNewParamsScoreUnion() {}

// Prompt session id
type EvalNewParamsScoresPromptSessionID struct {
	// The ID of the function in the prompt session
	PromptSessionFunctionID param.Field[string] `json:"prompt_session_function_id,required"`
	// The ID of the prompt session
	PromptSessionID param.Field[string] `json:"prompt_session_id,required"`
	// The version of the function
	Version param.Field[string] `json:"version"`
}

func (r EvalNewParamsScoresPromptSessionID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsScoresPromptSessionID) implementsEvalNewParamsScoreUnion() {}

// Inline code function
type EvalNewParamsScoresInlineCode struct {
	// The inline code to execute
	Code          param.Field[string]                                     `json:"code,required"`
	InlineContext param.Field[EvalNewParamsScoresInlineCodeInlineContext] `json:"inline_context,required"`
	// The name of the inline code function
	Name param.Field[string] `json:"name"`
}

func (r EvalNewParamsScoresInlineCode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsScoresInlineCode) implementsEvalNewParamsScoreUnion() {}

type EvalNewParamsScoresInlineCodeInlineContext struct {
	Runtime param.Field[EvalNewParamsScoresInlineCodeInlineContextRuntime] `json:"runtime,required"`
	Version param.Field[string]                                            `json:"version,required"`
}

func (r EvalNewParamsScoresInlineCodeInlineContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EvalNewParamsScoresInlineCodeInlineContextRuntime string

const (
	EvalNewParamsScoresInlineCodeInlineContextRuntimeNode   EvalNewParamsScoresInlineCodeInlineContextRuntime = "node"
	EvalNewParamsScoresInlineCodeInlineContextRuntimePython EvalNewParamsScoresInlineCodeInlineContextRuntime = "python"
)

func (r EvalNewParamsScoresInlineCodeInlineContextRuntime) IsKnown() bool {
	switch r {
	case EvalNewParamsScoresInlineCodeInlineContextRuntimeNode, EvalNewParamsScoresInlineCodeInlineContextRuntimePython:
		return true
	}
	return false
}

// Inline prompt definition
type EvalNewParamsScoresInlinePrompt struct {
	// The prompt, model, and its parameters
	InlinePrompt param.Field[shared.PromptDataParam] `json:"inline_prompt,required"`
	// The name of the inline prompt
	Name param.Field[string] `json:"name"`
}

func (r EvalNewParamsScoresInlinePrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsScoresInlinePrompt) implementsEvalNewParamsScoreUnion() {}

// The function to evaluate
type EvalNewParamsTask struct {
	// The inline code to execute
	Code param.Field[string] `json:"code"`
	// The ID of the function
	FunctionID param.Field[string] `json:"function_id"`
	// The name of the global function. Currently, the global namespace includes the
	// functions in autoevals
	GlobalFunction param.Field[string]      `json:"global_function"`
	InlineContext  param.Field[interface{}] `json:"inline_context"`
	// The prompt, model, and its parameters
	InlinePrompt param.Field[shared.PromptDataParam] `json:"inline_prompt"`
	// The name of the inline code function
	Name param.Field[string] `json:"name"`
	// The name of the project containing the function
	ProjectName param.Field[string] `json:"project_name"`
	// The ID of the function in the prompt session
	PromptSessionFunctionID param.Field[string] `json:"prompt_session_function_id"`
	// The ID of the prompt session
	PromptSessionID param.Field[string] `json:"prompt_session_id"`
	// The slug of the function
	Slug param.Field[string] `json:"slug"`
	// The version of the function
	Version param.Field[string] `json:"version"`
}

func (r EvalNewParamsTask) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsTask) implementsEvalNewParamsTaskUnion() {}

// The function to evaluate
//
// Satisfied by [EvalNewParamsTaskFunctionID], [EvalNewParamsTaskProjectSlug],
// [EvalNewParamsTaskGlobalFunction], [EvalNewParamsTaskPromptSessionID],
// [EvalNewParamsTaskInlineCode], [EvalNewParamsTaskInlinePrompt],
// [EvalNewParamsTask].
type EvalNewParamsTaskUnion interface {
	implementsEvalNewParamsTaskUnion()
}

// Function id
type EvalNewParamsTaskFunctionID struct {
	// The ID of the function
	FunctionID param.Field[string] `json:"function_id,required"`
	// The version of the function
	Version param.Field[string] `json:"version"`
}

func (r EvalNewParamsTaskFunctionID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsTaskFunctionID) implementsEvalNewParamsTaskUnion() {}

// Project name and slug
type EvalNewParamsTaskProjectSlug struct {
	// The name of the project containing the function
	ProjectName param.Field[string] `json:"project_name,required"`
	// The slug of the function
	Slug param.Field[string] `json:"slug,required"`
	// The version of the function
	Version param.Field[string] `json:"version"`
}

func (r EvalNewParamsTaskProjectSlug) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsTaskProjectSlug) implementsEvalNewParamsTaskUnion() {}

// Global function name
type EvalNewParamsTaskGlobalFunction struct {
	// The name of the global function. Currently, the global namespace includes the
	// functions in autoevals
	GlobalFunction param.Field[string] `json:"global_function,required"`
}

func (r EvalNewParamsTaskGlobalFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsTaskGlobalFunction) implementsEvalNewParamsTaskUnion() {}

// Prompt session id
type EvalNewParamsTaskPromptSessionID struct {
	// The ID of the function in the prompt session
	PromptSessionFunctionID param.Field[string] `json:"prompt_session_function_id,required"`
	// The ID of the prompt session
	PromptSessionID param.Field[string] `json:"prompt_session_id,required"`
	// The version of the function
	Version param.Field[string] `json:"version"`
}

func (r EvalNewParamsTaskPromptSessionID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsTaskPromptSessionID) implementsEvalNewParamsTaskUnion() {}

// Inline code function
type EvalNewParamsTaskInlineCode struct {
	// The inline code to execute
	Code          param.Field[string]                                   `json:"code,required"`
	InlineContext param.Field[EvalNewParamsTaskInlineCodeInlineContext] `json:"inline_context,required"`
	// The name of the inline code function
	Name param.Field[string] `json:"name"`
}

func (r EvalNewParamsTaskInlineCode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsTaskInlineCode) implementsEvalNewParamsTaskUnion() {}

type EvalNewParamsTaskInlineCodeInlineContext struct {
	Runtime param.Field[EvalNewParamsTaskInlineCodeInlineContextRuntime] `json:"runtime,required"`
	Version param.Field[string]                                          `json:"version,required"`
}

func (r EvalNewParamsTaskInlineCodeInlineContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EvalNewParamsTaskInlineCodeInlineContextRuntime string

const (
	EvalNewParamsTaskInlineCodeInlineContextRuntimeNode   EvalNewParamsTaskInlineCodeInlineContextRuntime = "node"
	EvalNewParamsTaskInlineCodeInlineContextRuntimePython EvalNewParamsTaskInlineCodeInlineContextRuntime = "python"
)

func (r EvalNewParamsTaskInlineCodeInlineContextRuntime) IsKnown() bool {
	switch r {
	case EvalNewParamsTaskInlineCodeInlineContextRuntimeNode, EvalNewParamsTaskInlineCodeInlineContextRuntimePython:
		return true
	}
	return false
}

// Inline prompt definition
type EvalNewParamsTaskInlinePrompt struct {
	// The prompt, model, and its parameters
	InlinePrompt param.Field[shared.PromptDataParam] `json:"inline_prompt,required"`
	// The name of the inline prompt
	Name param.Field[string] `json:"name"`
}

func (r EvalNewParamsTaskInlinePrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvalNewParamsTaskInlinePrompt) implementsEvalNewParamsTaskUnion() {}

// Optional settings for collecting git metadata. By default, will collect all git
// metadata fields allowed in org-level settings.
type EvalNewParamsGitMetadataSettings struct {
	Collect param.Field[EvalNewParamsGitMetadataSettingsCollect] `json:"collect,required"`
	Fields  param.Field[[]EvalNewParamsGitMetadataSettingsField] `json:"fields"`
}

func (r EvalNewParamsGitMetadataSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EvalNewParamsGitMetadataSettingsCollect string

const (
	EvalNewParamsGitMetadataSettingsCollectAll  EvalNewParamsGitMetadataSettingsCollect = "all"
	EvalNewParamsGitMetadataSettingsCollectNone EvalNewParamsGitMetadataSettingsCollect = "none"
	EvalNewParamsGitMetadataSettingsCollectSome EvalNewParamsGitMetadataSettingsCollect = "some"
)

func (r EvalNewParamsGitMetadataSettingsCollect) IsKnown() bool {
	switch r {
	case EvalNewParamsGitMetadataSettingsCollectAll, EvalNewParamsGitMetadataSettingsCollectNone, EvalNewParamsGitMetadataSettingsCollectSome:
		return true
	}
	return false
}

type EvalNewParamsGitMetadataSettingsField string

const (
	EvalNewParamsGitMetadataSettingsFieldCommit        EvalNewParamsGitMetadataSettingsField = "commit"
	EvalNewParamsGitMetadataSettingsFieldBranch        EvalNewParamsGitMetadataSettingsField = "branch"
	EvalNewParamsGitMetadataSettingsFieldTag           EvalNewParamsGitMetadataSettingsField = "tag"
	EvalNewParamsGitMetadataSettingsFieldDirty         EvalNewParamsGitMetadataSettingsField = "dirty"
	EvalNewParamsGitMetadataSettingsFieldAuthorName    EvalNewParamsGitMetadataSettingsField = "author_name"
	EvalNewParamsGitMetadataSettingsFieldAuthorEmail   EvalNewParamsGitMetadataSettingsField = "author_email"
	EvalNewParamsGitMetadataSettingsFieldCommitMessage EvalNewParamsGitMetadataSettingsField = "commit_message"
	EvalNewParamsGitMetadataSettingsFieldCommitTime    EvalNewParamsGitMetadataSettingsField = "commit_time"
	EvalNewParamsGitMetadataSettingsFieldGitDiff       EvalNewParamsGitMetadataSettingsField = "git_diff"
)

func (r EvalNewParamsGitMetadataSettingsField) IsKnown() bool {
	switch r {
	case EvalNewParamsGitMetadataSettingsFieldCommit, EvalNewParamsGitMetadataSettingsFieldBranch, EvalNewParamsGitMetadataSettingsFieldTag, EvalNewParamsGitMetadataSettingsFieldDirty, EvalNewParamsGitMetadataSettingsFieldAuthorName, EvalNewParamsGitMetadataSettingsFieldAuthorEmail, EvalNewParamsGitMetadataSettingsFieldCommitMessage, EvalNewParamsGitMetadataSettingsFieldCommitTime, EvalNewParamsGitMetadataSettingsFieldGitDiff:
		return true
	}
	return false
}
