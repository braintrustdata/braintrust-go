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

func TestProjectLogFeedback(t *testing.T) {
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
	_, err := client.Projects.Logs.Feedback(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ProjectLogFeedbackParams{
			Feedback: braintrust.F([]shared.FeedbackProjectLogsItemParam{{
				ID:       braintrust.F("id"),
				Comment:  braintrust.F("comment"),
				Expected: braintrust.F[any](map[string]interface{}{}),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": "bar",
				}),
				Scores: braintrust.F(map[string]float64{
					"foo": 0.000000,
				}),
				Source: braintrust.F(shared.FeedbackProjectLogsItemSourceApp),
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

func TestProjectLogFetchWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.Logs.Fetch(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ProjectLogFetchParams{
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

func TestProjectLogFetchPostWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.Logs.FetchPost(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ProjectLogFetchPostParams{
			Cursor: braintrust.F("cursor"),
			Filters: braintrust.F([]shared.PathLookupFilterParam{{
				Path:  braintrust.F([]string{"string"}),
				Type:  braintrust.F(shared.PathLookupFilterTypePathLookup),
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

func TestProjectLogInsert(t *testing.T) {
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
	_, err := client.Projects.Logs.Insert(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ProjectLogInsertParams{
			Events: braintrust.F([]shared.InsertProjectLogsEventParam{{
				ID:           braintrust.F("id"),
				IsMerge:      braintrust.F(true),
				MergePaths:   braintrust.F([][]string{{"string"}}),
				ObjectDelete: braintrust.F(true),
				ParentID:     braintrust.F("_parent_id"),
				Context: braintrust.F(shared.InsertProjectLogsEventContextParam{
					CallerFilename:     braintrust.F("caller_filename"),
					CallerFunctionname: braintrust.F("caller_functionname"),
					CallerLineno:       braintrust.F(int64(0)),
				}),
				Created:  braintrust.F(time.Now()),
				Error:    braintrust.F[any](map[string]interface{}{}),
				Expected: braintrust.F[any](map[string]interface{}{}),
				Input:    braintrust.F[any](map[string]interface{}{}),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": "bar",
				}),
				Metrics: braintrust.F(shared.InsertProjectLogsEventMetricsParam{
					CallerFilename:     braintrust.F[any](map[string]interface{}{}),
					CallerFunctionname: braintrust.F[any](map[string]interface{}{}),
					CallerLineno:       braintrust.F[any](map[string]interface{}{}),
					CompletionTokens:   braintrust.F(int64(0)),
					End:                braintrust.F(0.000000),
					PromptTokens:       braintrust.F(int64(0)),
					Start:              braintrust.F(0.000000),
					Tokens:             braintrust.F(int64(0)),
				}),
				Output:     braintrust.F[any](map[string]interface{}{}),
				RootSpanID: braintrust.F("root_span_id"),
				Scores: braintrust.F(map[string]float64{
					"foo": 0.000000,
				}),
				SpanAttributes: braintrust.F(shared.SpanAttributesParam{
					Name: braintrust.F("name"),
					Type: braintrust.F(shared.SpanAttributesTypeLlm),
				}),
				SpanID:      braintrust.F("span_id"),
				SpanParents: braintrust.F([]string{"string"}),
				Tags:        braintrust.F([]string{"string"}),
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
