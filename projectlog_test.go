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
	err := client.Projects.Logs.Feedback(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ProjectLogFeedbackParams{
			Feedback: braintrust.F([]braintrust.ProjectLogFeedbackParamsFeedback{{
				ID: braintrust.F("id"),
				Scores: braintrust.F(map[string]float64{
					"foo": 0.000000,
				}),
				Expected: braintrust.F[any](map[string]interface{}{}),
				Comment:  braintrust.F("comment"),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": map[string]interface{}{},
				}),
				Source: braintrust.F(braintrust.ProjectLogFeedbackParamsFeedbackSourceApp),
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
				Source: braintrust.F(braintrust.ProjectLogFeedbackParamsFeedbackSourceApp),
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
				Source: braintrust.F(braintrust.ProjectLogFeedbackParamsFeedbackSourceApp),
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
			Filters: braintrust.F([]braintrust.ProjectLogFetchPostParamsFilter{{
				Type:  braintrust.F(braintrust.ProjectLogFetchPostParamsFiltersTypePathLookup),
				Path:  braintrust.F([]string{"string", "string", "string"}),
				Value: braintrust.F[any](map[string]interface{}{}),
			}, {
				Type:  braintrust.F(braintrust.ProjectLogFetchPostParamsFiltersTypePathLookup),
				Path:  braintrust.F([]string{"string", "string", "string"}),
				Value: braintrust.F[any](map[string]interface{}{}),
			}, {
				Type:  braintrust.F(braintrust.ProjectLogFetchPostParamsFiltersTypePathLookup),
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
			Events: braintrust.F([]braintrust.ProjectLogInsertParamsEventUnion{braintrust.ProjectLogInsertParamsEventsInsertProjectLogsEventReplace{
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
				Metrics: braintrust.F(braintrust.ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceMetrics{
					Start:            braintrust.F(0.000000),
					End:              braintrust.F(0.000000),
					PromptTokens:     braintrust.F(int64(0)),
					CompletionTokens: braintrust.F(int64(0)),
					Tokens:           braintrust.F(int64(0)),
				}),
				Context: braintrust.F(braintrust.ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceContext{
					CallerFunctionname: braintrust.F("caller_functionname"),
					CallerFilename:     braintrust.F("caller_filename"),
					CallerLineno:       braintrust.F(int64(0)),
				}),
				SpanAttributes: braintrust.F(braintrust.ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributes{
					Name: braintrust.F("name"),
					Type: braintrust.F(braintrust.ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesTypeLlm),
				}),
				ID:           braintrust.F("id"),
				Created:      braintrust.F(time.Now()),
				ObjectDelete: braintrust.F(true),
				IsMerge:      braintrust.F(true),
				ParentID:     braintrust.F("_parent_id"),
			}, braintrust.ProjectLogInsertParamsEventsInsertProjectLogsEventReplace{
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
				Metrics: braintrust.F(braintrust.ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceMetrics{
					Start:            braintrust.F(0.000000),
					End:              braintrust.F(0.000000),
					PromptTokens:     braintrust.F(int64(0)),
					CompletionTokens: braintrust.F(int64(0)),
					Tokens:           braintrust.F(int64(0)),
				}),
				Context: braintrust.F(braintrust.ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceContext{
					CallerFunctionname: braintrust.F("caller_functionname"),
					CallerFilename:     braintrust.F("caller_filename"),
					CallerLineno:       braintrust.F(int64(0)),
				}),
				SpanAttributes: braintrust.F(braintrust.ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributes{
					Name: braintrust.F("name"),
					Type: braintrust.F(braintrust.ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesTypeLlm),
				}),
				ID:           braintrust.F("id"),
				Created:      braintrust.F(time.Now()),
				ObjectDelete: braintrust.F(true),
				IsMerge:      braintrust.F(true),
				ParentID:     braintrust.F("_parent_id"),
			}, braintrust.ProjectLogInsertParamsEventsInsertProjectLogsEventReplace{
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
				Metrics: braintrust.F(braintrust.ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceMetrics{
					Start:            braintrust.F(0.000000),
					End:              braintrust.F(0.000000),
					PromptTokens:     braintrust.F(int64(0)),
					CompletionTokens: braintrust.F(int64(0)),
					Tokens:           braintrust.F(int64(0)),
				}),
				Context: braintrust.F(braintrust.ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceContext{
					CallerFunctionname: braintrust.F("caller_functionname"),
					CallerFilename:     braintrust.F("caller_filename"),
					CallerLineno:       braintrust.F(int64(0)),
				}),
				SpanAttributes: braintrust.F(braintrust.ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributes{
					Name: braintrust.F("name"),
					Type: braintrust.F(braintrust.ProjectLogInsertParamsEventsInsertProjectLogsEventReplaceSpanAttributesTypeLlm),
				}),
				ID:           braintrust.F("id"),
				Created:      braintrust.F(time.Now()),
				ObjectDelete: braintrust.F(true),
				IsMerge:      braintrust.F(true),
				ParentID:     braintrust.F("_parent_id"),
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
