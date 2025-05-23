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
		ObjectID:           braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ObjectType:         braintrust.F(shared.ACLObjectTypeOrganization),
		GroupID:            braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Permission:         braintrust.F(shared.PermissionCreate),
		RestrictObjectType: braintrust.F(shared.ACLObjectTypeOrganization),
		RoleID:             braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		UserID:             braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
		ObjectID:      braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ObjectType:    braintrust.F(shared.ACLObjectTypeOrganization),
		EndingBefore:  braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IDs:           braintrust.F[braintrust.ACLListParamsIDsUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
		Limit:         braintrust.F(int64(0)),
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
		AddACLs: braintrust.F([]braintrust.ACLBatchUpdateParamsAddACL{{
			ObjectID:           braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ObjectType:         braintrust.F(shared.ACLObjectTypeOrganization),
			GroupID:            braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Permission:         braintrust.F(shared.PermissionCreate),
			RestrictObjectType: braintrust.F(shared.ACLObjectTypeOrganization),
			RoleID:             braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			UserID:             braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}}),
		RemoveACLs: braintrust.F([]braintrust.ACLBatchUpdateParamsRemoveACL{{
			ObjectID:           braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ObjectType:         braintrust.F(shared.ACLObjectTypeOrganization),
			GroupID:            braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Permission:         braintrust.F(shared.PermissionCreate),
			RestrictObjectType: braintrust.F(shared.ACLObjectTypeOrganization),
			RoleID:             braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			UserID:             braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}}),
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
		ObjectID:           braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ObjectType:         braintrust.F(shared.ACLObjectTypeOrganization),
		GroupID:            braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Permission:         braintrust.F(shared.PermissionCreate),
		RestrictObjectType: braintrust.F(shared.ACLObjectTypeOrganization),
		RoleID:             braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		UserID:             braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
