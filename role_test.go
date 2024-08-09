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

func TestRoleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Roles.New(context.TODO(), braintrust.RoleNewParams{
		CreateRole: shared.CreateRoleParam{
			Name:        braintrust.F("name"),
			Description: braintrust.F("description"),
			MemberPermissions: braintrust.F([]shared.CreateRoleMemberPermissionParam{{
				Permission:         braintrust.F(shared.CreateRoleMemberPermissionsPermissionCreate),
				RestrictObjectType: braintrust.F(shared.CreateRoleMemberPermissionsRestrictObjectTypeOrganization),
			}, {
				Permission:         braintrust.F(shared.CreateRoleMemberPermissionsPermissionCreate),
				RestrictObjectType: braintrust.F(shared.CreateRoleMemberPermissionsRestrictObjectTypeOrganization),
			}, {
				Permission:         braintrust.F(shared.CreateRoleMemberPermissionsPermissionCreate),
				RestrictObjectType: braintrust.F(shared.CreateRoleMemberPermissionsRestrictObjectTypeOrganization),
			}}),
			MemberRoles: braintrust.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			OrgName:     braintrust.F("org_name"),
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

func TestRoleGet(t *testing.T) {
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
	_, err := client.Roles.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Roles.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.RoleUpdateParams{
			PatchRole: shared.PatchRoleParam{
				Description: braintrust.F("description"),
				Name:        braintrust.F("name"),
				AddMemberPermissions: braintrust.F([]shared.PatchRoleAddMemberPermissionParam{{
					Permission:         braintrust.F(shared.PatchRoleAddMemberPermissionsPermissionCreate),
					RestrictObjectType: braintrust.F(shared.PatchRoleAddMemberPermissionsRestrictObjectTypeOrganization),
				}, {
					Permission:         braintrust.F(shared.PatchRoleAddMemberPermissionsPermissionCreate),
					RestrictObjectType: braintrust.F(shared.PatchRoleAddMemberPermissionsRestrictObjectTypeOrganization),
				}, {
					Permission:         braintrust.F(shared.PatchRoleAddMemberPermissionsPermissionCreate),
					RestrictObjectType: braintrust.F(shared.PatchRoleAddMemberPermissionsRestrictObjectTypeOrganization),
				}}),
				RemoveMemberPermissions: braintrust.F([]shared.PatchRoleRemoveMemberPermissionParam{{
					Permission:         braintrust.F(shared.PatchRoleRemoveMemberPermissionsPermissionCreate),
					RestrictObjectType: braintrust.F(shared.PatchRoleRemoveMemberPermissionsRestrictObjectTypeOrganization),
				}, {
					Permission:         braintrust.F(shared.PatchRoleRemoveMemberPermissionsPermissionCreate),
					RestrictObjectType: braintrust.F(shared.PatchRoleRemoveMemberPermissionsRestrictObjectTypeOrganization),
				}, {
					Permission:         braintrust.F(shared.PatchRoleRemoveMemberPermissionsPermissionCreate),
					RestrictObjectType: braintrust.F(shared.PatchRoleRemoveMemberPermissionsRestrictObjectTypeOrganization),
				}}),
				AddMemberRoles:    braintrust.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
				RemoveMemberRoles: braintrust.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
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

func TestRoleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Roles.List(context.TODO(), braintrust.RoleListParams{
		EndingBefore:  braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IDs:           braintrust.F[braintrust.RoleListParamsIDsUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
		Limit:         braintrust.F(int64(0)),
		OrgName:       braintrust.F("org_name"),
		RoleName:      braintrust.F("role_name"),
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

func TestRoleDelete(t *testing.T) {
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
	_, err := client.Roles.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoleReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Roles.Replace(context.TODO(), braintrust.RoleReplaceParams{
		CreateRole: shared.CreateRoleParam{
			Name:        braintrust.F("name"),
			Description: braintrust.F("description"),
			MemberPermissions: braintrust.F([]shared.CreateRoleMemberPermissionParam{{
				Permission:         braintrust.F(shared.CreateRoleMemberPermissionsPermissionCreate),
				RestrictObjectType: braintrust.F(shared.CreateRoleMemberPermissionsRestrictObjectTypeOrganization),
			}, {
				Permission:         braintrust.F(shared.CreateRoleMemberPermissionsPermissionCreate),
				RestrictObjectType: braintrust.F(shared.CreateRoleMemberPermissionsRestrictObjectTypeOrganization),
			}, {
				Permission:         braintrust.F(shared.CreateRoleMemberPermissionsPermissionCreate),
				RestrictObjectType: braintrust.F(shared.CreateRoleMemberPermissionsRestrictObjectTypeOrganization),
			}}),
			MemberRoles: braintrust.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			OrgName:     braintrust.F("org_name"),
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
