// File generated from our OpenAPI spec by Stainless.

package braintrust_test

import (
	"context"
	"os"
	"testing"

	"github.com/braintrustdata/braintrust-go"
	"github.com/braintrustdata/braintrust-go/internal/testutil"
	"github.com/braintrustdata/braintrust-go/option"
)

func TestUsage(t *testing.T) {
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
	project, err := client.Project.New(context.TODO(), braintrust.ProjectNewParams{
		Name: braintrust.F("string"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", project.ID)
}
