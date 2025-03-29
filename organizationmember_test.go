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
		InviteUsers: braintrust.OrganizationMemberUpdateParamsInviteUsers{
			Emails:           []string{"string"},
			GroupID:          braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			GroupIDs:         []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			GroupName:        braintrust.String("group_name"),
			GroupNames:       []string{"string"},
			IDs:              []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			SendInviteEmails: braintrust.Bool(true),
		},
		OrgID:   braintrust.String("org_id"),
		OrgName: braintrust.String("org_name"),
		RemoveUsers: braintrust.OrganizationMemberUpdateParamsRemoveUsers{
			Emails: []string{"string"},
			IDs:    []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
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
