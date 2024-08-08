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
	_, err := client.Datasets.New(context.TODO(), braintrust.DatasetNewParams{
		CreateDataset: shared.CreateDatasetParam{
			ProjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:        braintrust.F("name"),
			Description: braintrust.F("description"),
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
	_, err := client.Datasets.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
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
	_, err := client.Datasets.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.DatasetUpdateParams{
			PatchDataset: shared.PatchDatasetParam{
				Name:        braintrust.F("name"),
				Description: braintrust.F("description"),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": map[string]interface{}{},
				}),
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
	_, err := client.Datasets.List(context.TODO(), braintrust.DatasetListParams{
		DatasetName:   braintrust.F("dataset_name"),
		EndingBefore:  braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IDs:           braintrust.F[shared.IDsUnionParam](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
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
	_, err := client.Datasets.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
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
	err := client.Datasets.Feedback(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.DatasetFeedbackParams{
			FeedbackDatasetEventRequest: shared.FeedbackDatasetEventRequestParam{
				Feedback: braintrust.F([]shared.FeedbackDatasetItemParam{{
					ID:      braintrust.F("id"),
					Comment: braintrust.F("comment"),
					Metadata: braintrust.F(map[string]interface{}{
						"foo": map[string]interface{}{},
					}),
					Source: braintrust.F(shared.FeedbackDatasetItemSourceApp),
				}, {
					ID:      braintrust.F("id"),
					Comment: braintrust.F("comment"),
					Metadata: braintrust.F(map[string]interface{}{
						"foo": map[string]interface{}{},
					}),
					Source: braintrust.F(shared.FeedbackDatasetItemSourceApp),
				}, {
					ID:      braintrust.F("id"),
					Comment: braintrust.F("comment"),
					Metadata: braintrust.F(map[string]interface{}{
						"foo": map[string]interface{}{},
					}),
					Source: braintrust.F(shared.FeedbackDatasetItemSourceApp),
				}}),
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
	_, err := client.Datasets.Fetch(
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
	_, err := client.Datasets.FetchPost(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.DatasetFetchPostParams{
			FetchEventsRequest: shared.FetchEventsRequestParam{
				Limit:         braintrust.F(int64(0)),
				Cursor:        braintrust.F("cursor"),
				MaxXactID:     braintrust.F("max_xact_id"),
				MaxRootSpanID: braintrust.F("max_root_span_id"),
				Filters: braintrust.F([]shared.PathLookupFilterParam{{
					Type:  braintrust.F(shared.PathLookupFilterTypePathLookup),
					Path:  braintrust.F([]string{"string", "string", "string"}),
					Value: braintrust.F[any](map[string]interface{}{}),
				}, {
					Type:  braintrust.F(shared.PathLookupFilterTypePathLookup),
					Path:  braintrust.F([]string{"string", "string", "string"}),
					Value: braintrust.F[any](map[string]interface{}{}),
				}, {
					Type:  braintrust.F(shared.PathLookupFilterTypePathLookup),
					Path:  braintrust.F([]string{"string", "string", "string"}),
					Value: braintrust.F[any](map[string]interface{}{}),
				}}),
				Version: braintrust.F("version"),
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
	_, err := client.Datasets.Insert(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.DatasetInsertParams{
			InsertDatasetEventRequest: shared.InsertDatasetEventRequestParam{
				Events: braintrust.F([]shared.InsertDatasetEventUnionParam{shared.InsertDatasetEventReplaceParam{
					Input:    braintrust.F[any](map[string]interface{}{}),
					Expected: braintrust.F[any](map[string]interface{}{}),
					Metadata: braintrust.F(map[string]interface{}{
						"foo": map[string]interface{}{},
					}),
					Tags:         braintrust.F([]string{"string", "string", "string"}),
					ID:           braintrust.F("id"),
					Created:      braintrust.F(time.Now()),
					ObjectDelete: braintrust.F(true),
					IsMerge:      braintrust.F(true),
					ParentID:     braintrust.F("_parent_id"),
				}, shared.InsertDatasetEventReplaceParam{
					Input:    braintrust.F[any](map[string]interface{}{}),
					Expected: braintrust.F[any](map[string]interface{}{}),
					Metadata: braintrust.F(map[string]interface{}{
						"foo": map[string]interface{}{},
					}),
					Tags:         braintrust.F([]string{"string", "string", "string"}),
					ID:           braintrust.F("id"),
					Created:      braintrust.F(time.Now()),
					ObjectDelete: braintrust.F(true),
					IsMerge:      braintrust.F(true),
					ParentID:     braintrust.F("_parent_id"),
				}, shared.InsertDatasetEventReplaceParam{
					Input:    braintrust.F[any](map[string]interface{}{}),
					Expected: braintrust.F[any](map[string]interface{}{}),
					Metadata: braintrust.F(map[string]interface{}{
						"foo": map[string]interface{}{},
					}),
					Tags:         braintrust.F([]string{"string", "string", "string"}),
					ID:           braintrust.F("id"),
					Created:      braintrust.F(time.Now()),
					ObjectDelete: braintrust.F(true),
					IsMerge:      braintrust.F(true),
					ParentID:     braintrust.F("_parent_id"),
				}}),
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
	_, err := client.Datasets.Summarize(
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
