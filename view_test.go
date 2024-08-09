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

func TestViewNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.New(context.TODO(), braintrust.ViewNewParams{
		CreateView: shared.CreateViewParam{
			ObjectType: braintrust.F(shared.CreateViewObjectTypeOrganization),
			ObjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ViewType:   braintrust.F(shared.CreateViewViewTypeProjects),
			Name:       braintrust.F("name"),
			ViewData: braintrust.F(shared.ViewDataParam{
				Search: braintrust.F(shared.ViewDataSearchParam{
					Filter: braintrust.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
					Tag:    braintrust.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
					Match:  braintrust.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
					Sort:   braintrust.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				}),
			}),
			Options: braintrust.F(shared.ViewOptionsParam{
				ColumnVisibility: braintrust.F(map[string]bool{
					"foo": true,
				}),
				ColumnOrder: braintrust.F([]string{"string", "string", "string"}),
				ColumnSizing: braintrust.F(map[string]float64{
					"foo": 0.000000,
				}),
			}),
			UserID:    braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DeletedAt: braintrust.F(time.Now()),
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

func TestViewGet(t *testing.T) {
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
	_, err := client.Views.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ViewGetParams{
			ObjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ObjectType: braintrust.F(braintrust.ViewGetParamsObjectTypeOrganization),
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

func TestViewUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ViewUpdateParams{
			PatchView: shared.PatchViewParam{
				ObjectType: braintrust.F(shared.PatchViewObjectTypeOrganization),
				ObjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ViewType:   braintrust.F(shared.PatchViewViewTypeProjects),
				Name:       braintrust.F("name"),
				ViewData: braintrust.F(shared.ViewDataParam{
					Search: braintrust.F(shared.ViewDataSearchParam{
						Filter: braintrust.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
						Tag:    braintrust.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
						Match:  braintrust.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
						Sort:   braintrust.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
					}),
				}),
				Options: braintrust.F(shared.ViewOptionsParam{
					ColumnVisibility: braintrust.F(map[string]bool{
						"foo": true,
					}),
					ColumnOrder: braintrust.F([]string{"string", "string", "string"}),
					ColumnSizing: braintrust.F(map[string]float64{
						"foo": 0.000000,
					}),
				}),
				UserID: braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestViewListWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.List(context.TODO(), braintrust.ViewListParams{
		ObjectID:      braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ObjectType:    braintrust.F(braintrust.ViewListParamsObjectTypeOrganization),
		EndingBefore:  braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IDs:           braintrust.F[braintrust.ViewListParamsIDsUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
		Limit:         braintrust.F(int64(0)),
		StartingAfter: braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ViewName:      braintrust.F("view_name"),
		ViewType:      braintrust.F(braintrust.ViewListParamsViewTypeProjects),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestViewDelete(t *testing.T) {
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
	_, err := client.Views.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ViewDeleteParams{
			DeleteView: shared.DeleteViewParam{
				ObjectType: braintrust.F(shared.DeleteViewObjectTypeOrganization),
				ObjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestViewReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.Replace(context.TODO(), braintrust.ViewReplaceParams{
		CreateView: shared.CreateViewParam{
			ObjectType: braintrust.F(shared.CreateViewObjectTypeOrganization),
			ObjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ViewType:   braintrust.F(shared.CreateViewViewTypeProjects),
			Name:       braintrust.F("name"),
			ViewData: braintrust.F(shared.ViewDataParam{
				Search: braintrust.F(shared.ViewDataSearchParam{
					Filter: braintrust.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
					Tag:    braintrust.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
					Match:  braintrust.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
					Sort:   braintrust.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				}),
			}),
			Options: braintrust.F(shared.ViewOptionsParam{
				ColumnVisibility: braintrust.F(map[string]bool{
					"foo": true,
				}),
				ColumnOrder: braintrust.F([]string{"string", "string", "string"}),
				ColumnSizing: braintrust.F(map[string]float64{
					"foo": 0.000000,
				}),
			}),
			UserID:    braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DeletedAt: braintrust.F(time.Now()),
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
