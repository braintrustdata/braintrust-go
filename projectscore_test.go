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
		Name:      "name",
		ProjectID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		ScoreType: shared.ProjectScoreTypeSlider,
		Categories: braintrust.ProjectScoreNewParamsCategoriesUnion{
			OfCategorical: []shared.ProjectScoreCategoryParam{{
				Name:  "name",
				Value: 0,
			}},
		},
		Config: shared.ProjectScoreConfigParam{
			Destination: braintrust.String("destination"),
			MultiSelect: braintrust.Bool(true),
			Online: shared.OnlineScoreConfigParam{
				SamplingRate: 0,
				Scorers: []shared.OnlineScoreConfigScorerUnionParam{{
					OfFunction: &shared.OnlineScoreConfigScorerFunctionParam{
						ID:   "id",
						Type: "function",
					},
				}},
				ApplyToRootSpan:  braintrust.Bool(true),
				ApplyToSpanNames: []string{"string"},
			},
		},
		Description: braintrust.String("description"),
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
			Categories: braintrust.ProjectScoreUpdateParamsCategoriesUnion{
				OfCategorical: []shared.ProjectScoreCategoryParam{{
					Name:  "name",
					Value: 0,
				}},
			},
			Config: shared.ProjectScoreConfigParam{
				Destination: braintrust.String("destination"),
				MultiSelect: braintrust.Bool(true),
				Online: shared.OnlineScoreConfigParam{
					SamplingRate: 0,
					Scorers: []shared.OnlineScoreConfigScorerUnionParam{{
						OfFunction: &shared.OnlineScoreConfigScorerFunctionParam{
							ID:   "id",
							Type: "function",
						},
					}},
					ApplyToRootSpan:  braintrust.Bool(true),
					ApplyToSpanNames: []string{"string"},
				},
			},
			Description: braintrust.String("description"),
			Name:        braintrust.String("name"),
			ScoreType:   shared.ProjectScoreTypeSlider,
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
		EndingBefore: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IDs: braintrust.ProjectScoreListParamsIDsUnion{
			OfString: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
		Limit:            braintrust.Int(0),
		OrgName:          braintrust.String("org_name"),
		ProjectID:        braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectName:      braintrust.String("project_name"),
		ProjectScoreName: braintrust.String("project_score_name"),
		ScoreType: braintrust.ProjectScoreListParamsScoreTypeUnion{
			OfProjectScoreTypeSingle: braintrust.Opt(shared.ProjectScoreTypeSlider),
		},
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
		Name:      "name",
		ProjectID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		ScoreType: shared.ProjectScoreTypeSlider,
		Categories: braintrust.ProjectScoreReplaceParamsCategoriesUnion{
			OfCategorical: []shared.ProjectScoreCategoryParam{{
				Name:  "name",
				Value: 0,
			}},
		},
		Config: shared.ProjectScoreConfigParam{
			Destination: braintrust.String("destination"),
			MultiSelect: braintrust.Bool(true),
			Online: shared.OnlineScoreConfigParam{
				SamplingRate: 0,
				Scorers: []shared.OnlineScoreConfigScorerUnionParam{{
					OfFunction: &shared.OnlineScoreConfigScorerFunctionParam{
						ID:   "id",
						Type: "function",
					},
				}},
				ApplyToRootSpan:  braintrust.Bool(true),
				ApplyToSpanNames: []string{"string"},
			},
		},
		Description: braintrust.String("description"),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
