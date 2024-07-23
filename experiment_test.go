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
		ProjectID:      braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		BaseExpID:      braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		DatasetID:      braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		DatasetVersion: braintrust.F("dataset_version"),
		Description:    braintrust.F("description"),
		EnsureNew:      braintrust.F(true),
		Metadata: braintrust.F(map[string]interface{}{
			"foo": map[string]interface{}{},
		}),
		Name:   braintrust.F("name"),
		Public: braintrust.F(true),
		RepoInfo: braintrust.F(braintrust.ExperimentNewParamsRepoInfo{
			Commit:        braintrust.F("commit"),
			Branch:        braintrust.F("branch"),
			Tag:           braintrust.F("tag"),
			Dirty:         braintrust.F(true),
			AuthorName:    braintrust.F("author_name"),
			AuthorEmail:   braintrust.F("author_email"),
			CommitMessage: braintrust.F("commit_message"),
			CommitTime:    braintrust.F("commit_time"),
			GitDiff:       braintrust.F("git_diff"),
		}),
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
			BaseExpID:      braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DatasetID:      braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DatasetVersion: braintrust.F("dataset_version"),
			Description:    braintrust.F("description"),
			Metadata: braintrust.F(map[string]interface{}{
				"foo": map[string]interface{}{},
			}),
			Name:   braintrust.F("name"),
			Public: braintrust.F(true),
			RepoInfo: braintrust.F(braintrust.ExperimentUpdateParamsRepoInfo{
				Commit:        braintrust.F("commit"),
				Branch:        braintrust.F("branch"),
				Tag:           braintrust.F("tag"),
				Dirty:         braintrust.F(true),
				AuthorName:    braintrust.F("author_name"),
				AuthorEmail:   braintrust.F("author_email"),
				CommitMessage: braintrust.F("commit_message"),
				CommitTime:    braintrust.F("commit_time"),
				GitDiff:       braintrust.F("git_diff"),
			}),
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
		EndingBefore:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ExperimentName: braintrust.F("experiment_name"),
		IDs:            braintrust.F[braintrust.ExperimentListParamsIDsUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
		Limit:          braintrust.F(int64(0)),
		OrgName:        braintrust.F("org_name"),
		ProjectName:    braintrust.F("project_name"),
		StartingAfter:  braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
	err := client.Experiments.Feedback(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ExperimentFeedbackParams{
			Feedback: braintrust.F([]braintrust.ExperimentFeedbackParamsFeedback{{
				ID: braintrust.F("id"),
				Scores: braintrust.F(map[string]float64{
					"foo": 0.000000,
				}),
				Expected: braintrust.F[any](map[string]interface{}{}),
				Comment:  braintrust.F("comment"),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": map[string]interface{}{},
				}),
				Source: braintrust.F(braintrust.ExperimentFeedbackParamsFeedbackSourceApp),
			}, {
				ID: braintrust.F("id"),
				Scores: braintrust.F(map[string]float64{
					"foo": 0.000000,
				}),
				Expected: braintrust.F[any](map[string]interface{}{}),
				Comment:  braintrust.F("comment"),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": map[string]interface{}{},
				}),
				Source: braintrust.F(braintrust.ExperimentFeedbackParamsFeedbackSourceApp),
			}, {
				ID: braintrust.F("id"),
				Scores: braintrust.F(map[string]float64{
					"foo": 0.000000,
				}),
				Expected: braintrust.F[any](map[string]interface{}{}),
				Comment:  braintrust.F("comment"),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": map[string]interface{}{},
				}),
				Source: braintrust.F(braintrust.ExperimentFeedbackParamsFeedbackSourceApp),
			}}),
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
			Limit:         braintrust.F(int64(0)),
			MaxRootSpanID: braintrust.F("max_root_span_id"),
			MaxXactID:     braintrust.F("max_xact_id"),
			Version:       braintrust.F("version"),
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
			Cursor: braintrust.F("cursor"),
			Filters: braintrust.F([]braintrust.ExperimentFetchPostParamsFilter{{
				Type:  braintrust.F(braintrust.ExperimentFetchPostParamsFiltersTypePathLookup),
				Path:  braintrust.F([]string{"string", "string", "string"}),
				Value: braintrust.F[any](map[string]interface{}{}),
			}, {
				Type:  braintrust.F(braintrust.ExperimentFetchPostParamsFiltersTypePathLookup),
				Path:  braintrust.F([]string{"string", "string", "string"}),
				Value: braintrust.F[any](map[string]interface{}{}),
			}, {
				Type:  braintrust.F(braintrust.ExperimentFetchPostParamsFiltersTypePathLookup),
				Path:  braintrust.F([]string{"string", "string", "string"}),
				Value: braintrust.F[any](map[string]interface{}{}),
			}}),
			Limit:         braintrust.F(int64(0)),
			MaxRootSpanID: braintrust.F("max_root_span_id"),
			MaxXactID:     braintrust.F("max_xact_id"),
			Version:       braintrust.F("version"),
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
			Events: braintrust.F([]braintrust.ExperimentInsertParamsEventUnion{braintrust.ExperimentInsertParamsEventsInsertExperimentEventReplace{
				Input:    braintrust.F[any](map[string]interface{}{}),
				Output:   braintrust.F[any](map[string]interface{}{}),
				Expected: braintrust.F[any](map[string]interface{}{}),
				Scores: braintrust.F(map[string]float64{
					"foo": 0.000000,
				}),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": map[string]interface{}{},
				}),
				Tags: braintrust.F([]string{"string", "string", "string"}),
				Metrics: braintrust.F(braintrust.ExperimentInsertParamsEventsInsertExperimentEventReplaceMetrics{
					Start:            braintrust.F(0.000000),
					End:              braintrust.F(0.000000),
					PromptTokens:     braintrust.F(int64(0)),
					CompletionTokens: braintrust.F(int64(0)),
					Tokens:           braintrust.F(int64(0)),
				}),
				Context: braintrust.F(braintrust.ExperimentInsertParamsEventsInsertExperimentEventReplaceContext{
					CallerFunctionname: braintrust.F("caller_functionname"),
					CallerFilename:     braintrust.F("caller_filename"),
					CallerLineno:       braintrust.F(int64(0)),
				}),
				SpanAttributes: braintrust.F(braintrust.ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributes{
					Name: braintrust.F("name"),
					Type: braintrust.F(braintrust.ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesTypeLlm),
				}),
				ID:              braintrust.F("id"),
				DatasetRecordID: braintrust.F("dataset_record_id"),
				Created:         braintrust.F(time.Now()),
				ObjectDelete:    braintrust.F(true),
				IsMerge:         braintrust.F(true),
				ParentID:        braintrust.F("_parent_id"),
			}, braintrust.ExperimentInsertParamsEventsInsertExperimentEventReplace{
				Input:    braintrust.F[any](map[string]interface{}{}),
				Output:   braintrust.F[any](map[string]interface{}{}),
				Expected: braintrust.F[any](map[string]interface{}{}),
				Scores: braintrust.F(map[string]float64{
					"foo": 0.000000,
				}),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": map[string]interface{}{},
				}),
				Tags: braintrust.F([]string{"string", "string", "string"}),
				Metrics: braintrust.F(braintrust.ExperimentInsertParamsEventsInsertExperimentEventReplaceMetrics{
					Start:            braintrust.F(0.000000),
					End:              braintrust.F(0.000000),
					PromptTokens:     braintrust.F(int64(0)),
					CompletionTokens: braintrust.F(int64(0)),
					Tokens:           braintrust.F(int64(0)),
				}),
				Context: braintrust.F(braintrust.ExperimentInsertParamsEventsInsertExperimentEventReplaceContext{
					CallerFunctionname: braintrust.F("caller_functionname"),
					CallerFilename:     braintrust.F("caller_filename"),
					CallerLineno:       braintrust.F(int64(0)),
				}),
				SpanAttributes: braintrust.F(braintrust.ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributes{
					Name: braintrust.F("name"),
					Type: braintrust.F(braintrust.ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesTypeLlm),
				}),
				ID:              braintrust.F("id"),
				DatasetRecordID: braintrust.F("dataset_record_id"),
				Created:         braintrust.F(time.Now()),
				ObjectDelete:    braintrust.F(true),
				IsMerge:         braintrust.F(true),
				ParentID:        braintrust.F("_parent_id"),
			}, braintrust.ExperimentInsertParamsEventsInsertExperimentEventReplace{
				Input:    braintrust.F[any](map[string]interface{}{}),
				Output:   braintrust.F[any](map[string]interface{}{}),
				Expected: braintrust.F[any](map[string]interface{}{}),
				Scores: braintrust.F(map[string]float64{
					"foo": 0.000000,
				}),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": map[string]interface{}{},
				}),
				Tags: braintrust.F([]string{"string", "string", "string"}),
				Metrics: braintrust.F(braintrust.ExperimentInsertParamsEventsInsertExperimentEventReplaceMetrics{
					Start:            braintrust.F(0.000000),
					End:              braintrust.F(0.000000),
					PromptTokens:     braintrust.F(int64(0)),
					CompletionTokens: braintrust.F(int64(0)),
					Tokens:           braintrust.F(int64(0)),
				}),
				Context: braintrust.F(braintrust.ExperimentInsertParamsEventsInsertExperimentEventReplaceContext{
					CallerFunctionname: braintrust.F("caller_functionname"),
					CallerFilename:     braintrust.F("caller_filename"),
					CallerLineno:       braintrust.F(int64(0)),
				}),
				SpanAttributes: braintrust.F(braintrust.ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributes{
					Name: braintrust.F("name"),
					Type: braintrust.F(braintrust.ExperimentInsertParamsEventsInsertExperimentEventReplaceSpanAttributesTypeLlm),
				}),
				ID:              braintrust.F("id"),
				DatasetRecordID: braintrust.F("dataset_record_id"),
				Created:         braintrust.F(time.Now()),
				ObjectDelete:    braintrust.F(true),
				IsMerge:         braintrust.F(true),
				ParentID:        braintrust.F("_parent_id"),
			}}),
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
			ComparisonExperimentID: braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			SummarizeScores:        braintrust.F(true),
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
