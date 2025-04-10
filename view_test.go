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
		Name:       "name",
		ObjectID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		ObjectType: shared.ACLObjectTypeOrganization,
		ViewType:   braintrust.ViewNewParamsViewTypeProjects,
		DeletedAt:  braintrust.Time(time.Now()),
		Options: shared.ViewOptionsParam{
			ColumnOrder: []string{"string"},
			ColumnSizing: map[string]float64{
				"foo": 0,
			},
			ColumnVisibility: map[string]bool{
				"foo": true,
			},
			Grouping:  braintrust.String("grouping"),
			Layout:    braintrust.String("layout"),
			RowHeight: braintrust.String("rowHeight"),
		},
		UserID: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ViewData: shared.ViewDataParam{
			Search: shared.ViewDataSearchParam{
				Filter: []interface{}{map[string]interface{}{}},
				Match:  []interface{}{map[string]interface{}{}},
				Sort:   []interface{}{map[string]interface{}{}},
				Tag:    []interface{}{map[string]interface{}{}},
			},
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
			ObjectID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ObjectType: shared.ACLObjectTypeOrganization,
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
			ObjectID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ObjectType: shared.ACLObjectTypeOrganization,
			Name:       braintrust.String("name"),
			Options: shared.ViewOptionsParam{
				ColumnOrder: []string{"string"},
				ColumnSizing: map[string]float64{
					"foo": 0,
				},
				ColumnVisibility: map[string]bool{
					"foo": true,
				},
				Grouping:  braintrust.String("grouping"),
				Layout:    braintrust.String("layout"),
				RowHeight: braintrust.String("rowHeight"),
			},
			UserID: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ViewData: shared.ViewDataParam{
				Search: shared.ViewDataSearchParam{
					Filter: []interface{}{map[string]interface{}{}},
					Match:  []interface{}{map[string]interface{}{}},
					Sort:   []interface{}{map[string]interface{}{}},
					Tag:    []interface{}{map[string]interface{}{}},
				},
			},
			ViewType: braintrust.ViewUpdateParamsViewTypeProjects,
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
		ObjectID:     "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		ObjectType:   shared.ACLObjectTypeOrganization,
		EndingBefore: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IDs: braintrust.ViewListParamsIDsUnion{
			OfString: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
		Limit:         braintrust.Int(0),
		StartingAfter: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ViewName:      braintrust.String("view_name"),
		ViewType:      shared.ViewTypeProjects,
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
			ObjectID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ObjectType: shared.ACLObjectTypeOrganization,
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
		Name:       "name",
		ObjectID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		ObjectType: shared.ACLObjectTypeOrganization,
		ViewType:   braintrust.ViewReplaceParamsViewTypeProjects,
		DeletedAt:  braintrust.Time(time.Now()),
		Options: shared.ViewOptionsParam{
			ColumnOrder: []string{"string"},
			ColumnSizing: map[string]float64{
				"foo": 0,
			},
			ColumnVisibility: map[string]bool{
				"foo": true,
			},
			Grouping:  braintrust.String("grouping"),
			Layout:    braintrust.String("layout"),
			RowHeight: braintrust.String("rowHeight"),
		},
		UserID: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ViewData: shared.ViewDataParam{
			Search: shared.ViewDataSearchParam{
				Filter: []interface{}{map[string]interface{}{}},
				Match:  []interface{}{map[string]interface{}{}},
				Sort:   []interface{}{map[string]interface{}{}},
				Tag:    []interface{}{map[string]interface{}{}},
			},
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
