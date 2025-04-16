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
			Feedback: []shared.FeedbackProjectLogsItemParam{{
				ID:       "id",
				Comment:  braintrust.String("comment"),
				Expected: map[string]interface{}{},
				Metadata: map[string]any{
					"foo": "bar",
				},
				Scores: map[string]float64{
					"foo": 0,
				},
				Source: shared.FeedbackProjectLogsItemSourceApp,
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
			Events: []shared.InsertProjectLogsEventParam{{
				ID:           braintrust.String("id"),
				IsMerge:      braintrust.Bool(true),
				MergePaths:   [][]string{{"string"}},
				ObjectDelete: braintrust.Bool(true),
				ParentID:     braintrust.String("_parent_id"),
				Context: shared.InsertProjectLogsEventContextParam{
					CallerFilename:     braintrust.String("caller_filename"),
					CallerFunctionname: braintrust.String("caller_functionname"),
					CallerLineno:       braintrust.Int(0),
				},
				Created:  braintrust.Time(time.Now()),
				Error:    map[string]interface{}{},
				Expected: map[string]interface{}{},
				Input:    map[string]interface{}{},
				Metadata: shared.InsertProjectLogsEventMetadataParam{
					Model: braintrust.String("model"),
				},
				Metrics: shared.InsertProjectLogsEventMetricsParam{
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
