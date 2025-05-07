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

func TestACLNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ACLs.New(context.TODO(), braintrust.ACLNewParams{
		ObjectID:           "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		ObjectType:         shared.ACLObjectTypeOrganization,
		GroupID:            braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Permission:         shared.PermissionCreate,
		RestrictObjectType: shared.ACLObjectTypeOrganization,
		RoleID:             braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		UserID:             braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACLGet(t *testing.T) {
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
	_, err := client.ACLs.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACLListWithOptionalParams(t *testing.T) {
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
	_, err := client.ACLs.List(context.TODO(), braintrust.ACLListParams{
		ObjectID:     "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		ObjectType:   shared.ACLObjectTypeOrganization,
		EndingBefore: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IDs: braintrust.ACLListParamsIDsUnion{
			OfString: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
		Limit:         braintrust.Int(0),
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

func TestACLDelete(t *testing.T) {
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
	_, err := client.ACLs.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACLBatchUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ACLs.BatchUpdate(context.TODO(), braintrust.ACLBatchUpdateParams{
		AddACLs: []braintrust.ACLBatchUpdateParamsAddACL{{
			ObjectID:           "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ObjectType:         shared.ACLObjectTypeOrganization,
			GroupID:            braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Permission:         shared.PermissionCreate,
			RestrictObjectType: shared.ACLObjectTypeOrganization,
			RoleID:             braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			UserID:             braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}},
		RemoveACLs: []braintrust.ACLBatchUpdateParamsRemoveACL{{
			ObjectID:           "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ObjectType:         shared.ACLObjectTypeOrganization,
			GroupID:            braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Permission:         shared.PermissionCreate,
			RestrictObjectType: shared.ACLObjectTypeOrganization,
			RoleID:             braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			UserID:             braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}},
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACLFindAndDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.ACLs.FindAndDelete(context.TODO(), braintrust.ACLFindAndDeleteParams{
		ObjectID:           "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		ObjectType:         shared.ACLObjectTypeOrganization,
		GroupID:            braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Permission:         shared.PermissionCreate,
		RestrictObjectType: shared.ACLObjectTypeOrganization,
		RoleID:             braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		UserID:             braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
