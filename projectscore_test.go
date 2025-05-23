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
	_, err := client.ProjectScores.New(context.TODO(), braintrust.ProjectScoreNewParams{
		Name:      braintrust.F("name"),
		ProjectID: braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ScoreType: braintrust.F(shared.ProjectScoreTypeSlider),
		Categories: braintrust.F[braintrust.ProjectScoreNewParamsCategoriesUnion](braintrust.ProjectScoreNewParamsCategoriesCategorical([]shared.ProjectScoreCategoryParam{{
			Name:  braintrust.F("name"),
			Value: braintrust.F(0.000000),
		}})),
		Config: braintrust.F(shared.ProjectScoreConfigParam{
			Destination: braintrust.F("destination"),
			MultiSelect: braintrust.F(true),
			Online: braintrust.F(shared.OnlineScoreConfigParam{
				SamplingRate: braintrust.F(0.000000),
				Scorers: braintrust.F([]shared.OnlineScoreConfigScorersUnionParam{shared.OnlineScoreConfigScorersFunctionParam{
					ID:   braintrust.F("id"),
					Type: braintrust.F(shared.OnlineScoreConfigScorersFunctionTypeFunction),
				}}),
				ApplyToRootSpan:  braintrust.F(true),
				ApplyToSpanNames: braintrust.F([]string{"string"}),
			}),
		}),
		Description: braintrust.F("description"),
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
	_, err := client.ProjectScores.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
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
	_, err := client.ProjectScores.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.ProjectScoreUpdateParams{
			Categories: braintrust.F[braintrust.ProjectScoreUpdateParamsCategoriesUnion](braintrust.ProjectScoreUpdateParamsCategoriesCategorical([]shared.ProjectScoreCategoryParam{{
				Name:  braintrust.F("name"),
				Value: braintrust.F(0.000000),
			}})),
			Config: braintrust.F(shared.ProjectScoreConfigParam{
				Destination: braintrust.F("destination"),
				MultiSelect: braintrust.F(true),
				Online: braintrust.F(shared.OnlineScoreConfigParam{
					SamplingRate: braintrust.F(0.000000),
					Scorers: braintrust.F([]shared.OnlineScoreConfigScorersUnionParam{shared.OnlineScoreConfigScorersFunctionParam{
						ID:   braintrust.F("id"),
						Type: braintrust.F(shared.OnlineScoreConfigScorersFunctionTypeFunction),
					}}),
					ApplyToRootSpan:  braintrust.F(true),
					ApplyToSpanNames: braintrust.F([]string{"string"}),
				}),
			}),
			Description: braintrust.F("description"),
			Name:        braintrust.F("name"),
			ScoreType:   braintrust.F(shared.ProjectScoreTypeSlider),
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
	_, err := client.ProjectScores.List(context.TODO(), braintrust.ProjectScoreListParams{
		EndingBefore:     braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IDs:              braintrust.F[braintrust.ProjectScoreListParamsIDsUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
		Limit:            braintrust.F(int64(0)),
		OrgName:          braintrust.F("org_name"),
		ProjectID:        braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectName:      braintrust.F("project_name"),
		ProjectScoreName: braintrust.F("project_score_name"),
		ScoreType:        braintrust.F[braintrust.ProjectScoreListParamsScoreTypeUnion](shared.ProjectScoreType(shared.ProjectScoreTypeSlider)),
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
	_, err := client.ProjectScores.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
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
	_, err := client.ProjectScores.Replace(context.TODO(), braintrust.ProjectScoreReplaceParams{
		Name:      braintrust.F("name"),
		ProjectID: braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ScoreType: braintrust.F(shared.ProjectScoreTypeSlider),
		Categories: braintrust.F[braintrust.ProjectScoreReplaceParamsCategoriesUnion](braintrust.ProjectScoreReplaceParamsCategoriesCategorical([]shared.ProjectScoreCategoryParam{{
			Name:  braintrust.F("name"),
			Value: braintrust.F(0.000000),
		}})),
		Config: braintrust.F(shared.ProjectScoreConfigParam{
			Destination: braintrust.F("destination"),
			MultiSelect: braintrust.F(true),
			Online: braintrust.F(shared.OnlineScoreConfigParam{
				SamplingRate: braintrust.F(0.000000),
				Scorers: braintrust.F([]shared.OnlineScoreConfigScorersUnionParam{shared.OnlineScoreConfigScorersFunctionParam{
					ID:   braintrust.F("id"),
					Type: braintrust.F(shared.OnlineScoreConfigScorersFunctionTypeFunction),
				}}),
				ApplyToRootSpan:  braintrust.F(true),
				ApplyToSpanNames: braintrust.F([]string{"string"}),
			}),
		}),
		Description: braintrust.F("description"),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
