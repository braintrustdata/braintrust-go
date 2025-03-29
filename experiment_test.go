// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/braintrustdata/braintrust-go"
	"github.com/braintrustdata/braintrust-go/internal/testutil"
	"github.com/braintrustdata/braintrust-go/option"
	"github.com/braintrustdata/braintrust-go/shared"
)

func TestExperimentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.New(context.TODO(), braintrust.ExperimentNewParams{
		ProjectID:      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		BaseExpID:      braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		DatasetID:      braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		DatasetVersion: braintrust.String("dataset_version"),
		Description:    braintrust.String("description"),
		EnsureNew:      braintrust.Bool(true),
		Metadata: map[string]interface{}{
			"foo": "bar",
		},
		Name:   braintrust.String("x"),
		Public: braintrust.Bool(true),
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
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentGet(t *testing.T) {
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
	_, err := client.Experiments.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ExperimentUpdateParams{
			BaseExpID:      braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DatasetID:      braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DatasetVersion: braintrust.String("dataset_version"),
			Description:    braintrust.String("description"),
			Metadata: map[string]interface{}{
				"foo": "bar",
			},
			Name:   braintrust.String("name"),
			Public: braintrust.Bool(true),
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
		},
	)
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.List(context.TODO(), braintrust.ExperimentListParams{
		EndingBefore:   braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ExperimentName: braintrust.String("experiment_name"),
		IDs: braintrust.ExperimentListParamsIDsUnion{
			OfString: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
		Limit:         braintrust.Int(0),
		OrgName:       braintrust.String("org_name"),
		ProjectID:     braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectName:   braintrust.String("project_name"),
		StartingAfter: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentDelete(t *testing.T) {
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
	_, err := client.Experiments.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentFeedback(t *testing.T) {
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
	_, err := client.Experiments.Feedback(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ExperimentFeedbackParams{
			Feedback: []shared.FeedbackExperimentItemParam{{
				ID:       "id",
				Comment:  braintrust.String("comment"),
				Expected: map[string]interface{}{},
				Metadata: map[string]interface{}{
					"foo": "bar",
				},
				Scores: map[string]float64{
					"foo": 0,
				},
				Source: shared.FeedbackExperimentItemSourceApp,
				Tags:   []string{"string"},
			}},
		},
	)
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentFetchWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.Fetch(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ExperimentFetchParams{
			Limit:         braintrust.Int(0),
			MaxRootSpanID: braintrust.String("max_root_span_id"),
			MaxXactID:     braintrust.String("max_xact_id"),
			Version:       braintrust.String("version"),
		},
	)
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentFetchPostWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.FetchPost(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ExperimentFetchPostParams{
			Cursor:        braintrust.String("cursor"),
			Limit:         braintrust.Int(0),
			MaxRootSpanID: braintrust.String("max_root_span_id"),
			MaxXactID:     braintrust.String("max_xact_id"),
			Version:       braintrust.String("version"),
		},
	)
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentInsert(t *testing.T) {
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
	_, err := client.Experiments.Insert(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ExperimentInsertParams{
			Events: []shared.InsertExperimentEventParam{{
				ID:           braintrust.String("id"),
				IsMerge:      braintrust.Bool(true),
				MergePaths:   [][]string{{"string"}},
				ObjectDelete: braintrust.Bool(true),
				ParentID:     braintrust.String("_parent_id"),
				Context: shared.InsertExperimentEventContextParam{
					CallerFilename:     braintrust.String("caller_filename"),
					CallerFunctionname: braintrust.String("caller_functionname"),
					CallerLineno:       braintrust.Int(0),
				},
				Created:  braintrust.Time(time.Now()),
				Error:    map[string]interface{}{},
				Expected: map[string]interface{}{},
				Input:    map[string]interface{}{},
				Metadata: shared.InsertExperimentEventMetadataParam{
					Model: braintrust.String("model"),
				},
				Metrics: shared.InsertExperimentEventMetricsParam{
					CallerFilename:     map[string]interface{}{},
					CallerFunctionname: map[string]interface{}{},
					CallerLineno:       map[string]interface{}{},
					CompletionTokens:   braintrust.Int(0),
					End:                braintrust.Float(0),
					PromptTokens:       braintrust.Int(0),
					Start:              braintrust.Float(0),
					Tokens:             braintrust.Int(0),
				},
				Origin: shared.ObjectReferenceParam{
					ID:         "id",
					XactID:     "_xact_id",
					ObjectID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
					ObjectType: shared.ObjectReferenceObjectTypeExperiment,
					Created:    braintrust.String("created"),
				},
				Output:     map[string]interface{}{},
				RootSpanID: braintrust.String("root_span_id"),
				Scores: map[string]float64{
					"foo": 0,
				},
				SpanAttributes: shared.SpanAttributesParam{
					Name: braintrust.String("name"),
					Type: shared.SpanTypeLlm,
				},
				SpanID:      braintrust.String("span_id"),
				SpanParents: []string{"string"},
				Tags:        []string{"string"},
			}},
		},
	)
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentSummarizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.Summarize(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ExperimentSummarizeParams{
			ComparisonExperimentID: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			SummarizeScores:        braintrust.Bool(true),
		},
	)
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
