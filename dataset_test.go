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

func TestDatasetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Dataset.New(context.TODO(), braintrust.DatasetNewParams{
		Name:        braintrust.F("name"),
		Description: braintrust.F("description"),
		ProjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDatasetGet(t *testing.T) {
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
	_, err := client.Dataset.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDatasetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Dataset.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.DatasetUpdateParams{
			Description: braintrust.F("description"),
			Metadata: braintrust.F(map[string]interface{}{
				"foo": "bar",
			}),
			Name: braintrust.F("name"),
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

func TestDatasetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Dataset.List(context.TODO(), braintrust.DatasetListParams{
		DatasetName:   braintrust.F("dataset_name"),
		EndingBefore:  braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IDs:           braintrust.F[braintrust.DatasetListParamsIDsUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
		Limit:         braintrust.F(int64(0)),
		OrgName:       braintrust.F("org_name"),
		ProjectID:     braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectName:   braintrust.F("project_name"),
		StartingAfter: braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDatasetDelete(t *testing.T) {
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
	_, err := client.Dataset.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDatasetFeedback(t *testing.T) {
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
	_, err := client.Dataset.Feedback(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.DatasetFeedbackParams{
			Feedback: braintrust.F([]shared.FeedbackDatasetItemParam{{
				ID:      braintrust.F("id"),
				Comment: braintrust.F("comment"),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": "bar",
				}),
				Source: braintrust.F(shared.FeedbackDatasetItemSourceApp),
			}, {
				ID:      braintrust.F("id"),
				Comment: braintrust.F("comment"),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": "bar",
				}),
				Source: braintrust.F(shared.FeedbackDatasetItemSourceApp),
			}, {
				ID:      braintrust.F("id"),
				Comment: braintrust.F("comment"),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": "bar",
				}),
				Source: braintrust.F(shared.FeedbackDatasetItemSourceApp),
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

func TestDatasetFetchWithOptionalParams(t *testing.T) {
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
	_, err := client.Dataset.Fetch(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.DatasetFetchParams{
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

func TestDatasetFetchPostWithOptionalParams(t *testing.T) {
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
	_, err := client.Dataset.FetchPost(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.DatasetFetchPostParams{
			Cursor: braintrust.F("cursor"),
			Filters: braintrust.F([]shared.PathLookupFilterParam{{
				Path:  braintrust.F([]string{"string", "string", "string"}),
				Type:  braintrust.F(shared.PathLookupFilterTypePathLookup),
				Value: braintrust.F[any](map[string]interface{}{}),
			}, {
				Path:  braintrust.F([]string{"string", "string", "string"}),
				Type:  braintrust.F(shared.PathLookupFilterTypePathLookup),
				Value: braintrust.F[any](map[string]interface{}{}),
			}, {
				Path:  braintrust.F([]string{"string", "string", "string"}),
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

func TestDatasetInsert(t *testing.T) {
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
	_, err := client.Dataset.Insert(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.DatasetInsertParams{
			Events: braintrust.F([]braintrust.DatasetInsertParamsEventUnion{shared.InsertDatasetEventReplaceParam{
				ID:           braintrust.F("id"),
				IsMerge:      braintrust.F(true),
				ObjectDelete: braintrust.F(true),
				ParentID:     braintrust.F("_parent_id"),
				Created:      braintrust.F(time.Now()),
				Expected:     braintrust.F[any](map[string]interface{}{}),
				Input:        braintrust.F[any](map[string]interface{}{}),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": "bar",
				}),
				Tags: braintrust.F([]string{"string", "string", "string"}),
			}, shared.InsertDatasetEventReplaceParam{
				ID:           braintrust.F("id"),
				IsMerge:      braintrust.F(true),
				ObjectDelete: braintrust.F(true),
				ParentID:     braintrust.F("_parent_id"),
				Created:      braintrust.F(time.Now()),
				Expected:     braintrust.F[any](map[string]interface{}{}),
				Input:        braintrust.F[any](map[string]interface{}{}),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": "bar",
				}),
				Tags: braintrust.F([]string{"string", "string", "string"}),
			}, shared.InsertDatasetEventReplaceParam{
				ID:           braintrust.F("id"),
				IsMerge:      braintrust.F(true),
				ObjectDelete: braintrust.F(true),
				ParentID:     braintrust.F("_parent_id"),
				Created:      braintrust.F(time.Now()),
				Expected:     braintrust.F[any](map[string]interface{}{}),
				Input:        braintrust.F[any](map[string]interface{}{}),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": "bar",
				}),
				Tags: braintrust.F([]string{"string", "string", "string"}),
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

func TestDatasetSummarizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Dataset.Summarize(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.DatasetSummarizeParams{
			SummarizeData: braintrust.F(true),
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
