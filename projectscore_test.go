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

func TestProjectScoreNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ProjectScore.New(context.TODO(), braintrust.ProjectScoreNewParams{
		Name:      braintrust.F("string"),
		ProjectID: braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ScoreType: braintrust.F(braintrust.ProjectScoreNewParamsScoreTypeSlider),
		Categories: braintrust.F[braintrust.ProjectScoreNewParamsCategoriesUnion](braintrust.ProjectScoreNewParamsCategoriesArray([]braintrust.ProjectScoreNewParamsCategoriesArray{{
			Name:  braintrust.F("string"),
			Value: braintrust.F(0.000000),
		}, {
			Name:  braintrust.F("string"),
			Value: braintrust.F(0.000000),
		}, {
			Name:  braintrust.F("string"),
			Value: braintrust.F(0.000000),
		}})),
		Description: braintrust.F("string"),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectScoreGet(t *testing.T) {
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
	_, err := client.ProjectScore.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectScoreUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ProjectScore.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ProjectScoreUpdateParams{
			Categories: braintrust.F[braintrust.ProjectScoreUpdateParamsCategoriesUnion](braintrust.ProjectScoreUpdateParamsCategoriesArray([]braintrust.ProjectScoreUpdateParamsCategoriesArray{{
				Name:  braintrust.F("string"),
				Value: braintrust.F(0.000000),
			}, {
				Name:  braintrust.F("string"),
				Value: braintrust.F(0.000000),
			}, {
				Name:  braintrust.F("string"),
				Value: braintrust.F(0.000000),
			}})),
			Description: braintrust.F("string"),
			Name:        braintrust.F("string"),
			ScoreType:   braintrust.F(braintrust.ProjectScoreUpdateParamsScoreTypeSlider),
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

func TestProjectScoreListWithOptionalParams(t *testing.T) {
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
	_, err := client.ProjectScore.List(context.TODO(), braintrust.ProjectScoreListParams{
		EndingBefore:     braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IDs:              braintrust.F[braintrust.ProjectScoreListParamsIDsUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
		Limit:            braintrust.F(int64(0)),
		OrgName:          braintrust.F("string"),
		ProjectName:      braintrust.F("string"),
		ProjectScoreName: braintrust.F("string"),
		StartingAfter:    braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectScoreDelete(t *testing.T) {
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
	_, err := client.ProjectScore.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectScoreReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.ProjectScore.Replace(context.TODO(), braintrust.ProjectScoreReplaceParams{
		Name:      braintrust.F("string"),
		ProjectID: braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ScoreType: braintrust.F(braintrust.ProjectScoreReplaceParamsScoreTypeSlider),
		Categories: braintrust.F[braintrust.ProjectScoreReplaceParamsCategoriesUnion](braintrust.ProjectScoreReplaceParamsCategoriesArray([]braintrust.ProjectScoreReplaceParamsCategoriesArray{{
			Name:  braintrust.F("string"),
			Value: braintrust.F(0.000000),
		}, {
			Name:  braintrust.F("string"),
			Value: braintrust.F(0.000000),
		}, {
			Name:  braintrust.F("string"),
			Value: braintrust.F(0.000000),
		}})),
		Description: braintrust.F("string"),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
