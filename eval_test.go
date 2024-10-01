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

func TestEvalNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Evals.New(context.TODO(), braintrust.EvalNewParams{
		Data: braintrust.F[braintrust.EvalNewParamsDataUnion](braintrust.EvalNewParamsDataDatasetID{
			DatasetID: braintrust.F("dataset_id"),
		}),
		ProjectID: braintrust.F("project_id"),
		Scores: braintrust.F([]braintrust.EvalNewParamsScoreUnion{braintrust.EvalNewParamsScoresFunctionID{
			FunctionID: braintrust.F("function_id"),
			Version:    braintrust.F("version"),
		}, braintrust.EvalNewParamsScoresFunctionID{
			FunctionID: braintrust.F("function_id"),
			Version:    braintrust.F("version"),
		}, braintrust.EvalNewParamsScoresFunctionID{
			FunctionID: braintrust.F("function_id"),
			Version:    braintrust.F("version"),
		}}),
		Task: braintrust.F[braintrust.EvalNewParamsTaskUnion](braintrust.EvalNewParamsTaskFunctionID{
			FunctionID: braintrust.F("function_id"),
			Version:    braintrust.F("version"),
		}),
		ExperimentName: braintrust.F("experiment_name"),
		Metadata: braintrust.F(map[string]interface{}{
			"foo": "bar",
		}),
		Stream: braintrust.F(true),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
