// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/braintrustdata/braintrust-go"
	"github.com/braintrustdata/braintrust-go/internal/testutil"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/braintrustdata/braintrust-go/shared"
)

func TestEvalNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := braintrust.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Evals.New(context.TODO(), braintrust.EvalNewParams{
		Data: braintrust.EvalNewParamsDataUnion{
			OfDatasetID: &braintrust.EvalNewParamsDataDatasetID{
				DatasetID: "dataset_id",
				InternalBtql: map[string]any{
					"foo": "bar",
				},
			},
		},
		ProjectID: "project_id",
		Scores: []braintrust.EvalNewParamsScoreUnion{{
			OfFunctionID: &braintrust.EvalNewParamsScoreFunctionID{
				FunctionID: "function_id",
				Version:    braintrust.String("version"),
			},
		}},
		Task: braintrust.EvalNewParamsTaskUnion{
			OfFunctionID: &braintrust.EvalNewParamsTaskFunctionID{
				FunctionID: "function_id",
				Version:    braintrust.String("version"),
			},
		},
		BaseExperimentID:   braintrust.String("base_experiment_id"),
		BaseExperimentName: braintrust.String("base_experiment_name"),
		ExperimentName:     braintrust.String("experiment_name"),
		GitMetadataSettings: braintrust.EvalNewParamsGitMetadataSettings{
			Collect: "all",
			Fields:  []string{"commit"},
		},
		IsPublic:       braintrust.Bool(true),
		MaxConcurrency: braintrust.Float(0),
		Metadata: map[string]any{
			"foo": "bar",
		},
		Parent: braintrust.EvalNewParamsParentUnion{
			OfSpanParentStruct: &braintrust.EvalNewParamsParentSpanParentStruct{
				ObjectID:   "object_id",
				ObjectType: "project_logs",
				PropagatedEvent: map[string]any{
					"foo": "bar",
				},
				RowIDs: braintrust.EvalNewParamsParentSpanParentStructRowIDs{
					ID:         "id",
					RootSpanID: "root_span_id",
					SpanID:     "span_id",
				},
			},
		},
		RepoInfo: shared.RepoInfoParam{
			AuthorEmail:   braintrust.String("author_email"),
			AuthorName:    braintrust.String("author_name"),
			Branch:        braintrust.String("branch"),
			Commit:        braintrust.String("commit"),
			CommitMessage: braintrust.String("commit_message"),
			CommitTime:    braintrust.String("commit_time"),
			Dirty:         braintrust.Bool(true),
			GitDiff:       braintrust.String("git_diff"),
			Tag:           braintrust.String("tag"),
		},
		Stream:     braintrust.Bool(true),
		Timeout:    braintrust.Float(0),
		TrialCount: braintrust.Float(0),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
