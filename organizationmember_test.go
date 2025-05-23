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
)

func TestOrganizationMemberUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.Members.Update(context.TODO(), braintrust.OrganizationMemberUpdateParams{
		InviteUsers: braintrust.F(braintrust.OrganizationMemberUpdateParamsInviteUsers{
			Emails:           braintrust.F([]string{"string"}),
			GroupID:          braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			GroupIDs:         braintrust.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			GroupName:        braintrust.F("group_name"),
			GroupNames:       braintrust.F([]string{"string"}),
			IDs:              braintrust.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			SendInviteEmails: braintrust.F(true),
		}),
		OrgID:   braintrust.F("org_id"),
		OrgName: braintrust.F("org_name"),
		RemoveUsers: braintrust.F(braintrust.OrganizationMemberUpdateParamsRemoveUsers{
			Emails: braintrust.F([]string{"string"}),
			IDs:    braintrust.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		}),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
