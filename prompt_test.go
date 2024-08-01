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

func TestPromptNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Prompts.New(context.TODO(), braintrust.PromptNewParams{
		Name:        braintrust.F("name"),
		ProjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Slug:        braintrust.F("slug"),
		Description: braintrust.F("description"),
		PromptData: braintrust.F(braintrust.PromptNewParamsPromptData{
			Prompt: braintrust.F[braintrust.PromptNewParamsPromptDataPromptUnion](braintrust.PromptNewParamsPromptDataPromptCompletion{
				Type:    braintrust.F(braintrust.PromptNewParamsPromptDataPromptCompletionTypeCompletion),
				Content: braintrust.F("content"),
			}),
			Options: braintrust.F(braintrust.PromptNewParamsPromptDataOptions{
				Model: braintrust.F("model"),
				Params: braintrust.F[braintrust.PromptNewParamsPromptDataOptionsParamsUnion](braintrust.PromptNewParamsPromptDataOptionsParamsOpenAIModelParams{
					UseCache:         braintrust.F(true),
					Temperature:      braintrust.F(0.000000),
					TopP:             braintrust.F(0.000000),
					MaxTokens:        braintrust.F(0.000000),
					FrequencyPenalty: braintrust.F(0.000000),
					PresencePenalty:  braintrust.F(0.000000),
					ResponseFormat: braintrust.F(braintrust.PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat{
						Type: braintrust.F(braintrust.PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject),
					}),
					ToolChoice:   braintrust.F[braintrust.PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion](braintrust.PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto(braintrust.PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto)),
					FunctionCall: braintrust.F[braintrust.PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion](braintrust.PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto(braintrust.PromptNewParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto)),
					N:            braintrust.F(0.000000),
					Stop:         braintrust.F([]string{"string", "string", "string"}),
				}),
				Position: braintrust.F("position"),
			}),
			Origin: braintrust.F(braintrust.PromptNewParamsPromptDataOrigin{
				PromptID:      braintrust.F("prompt_id"),
				ProjectID:     braintrust.F("project_id"),
				PromptVersion: braintrust.F("prompt_version"),
			}),
		}),
		Tags: braintrust.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPromptGet(t *testing.T) {
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
	_, err := client.Prompts.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPromptUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Prompts.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.PromptUpdateParams{
			Description: braintrust.F("description"),
			Name:        braintrust.F("name"),
			PromptData: braintrust.F(braintrust.PromptUpdateParamsPromptData{
				Prompt: braintrust.F[braintrust.PromptUpdateParamsPromptDataPromptUnion](braintrust.PromptUpdateParamsPromptDataPromptCompletion{
					Type:    braintrust.F(braintrust.PromptUpdateParamsPromptDataPromptCompletionTypeCompletion),
					Content: braintrust.F("content"),
				}),
				Options: braintrust.F(braintrust.PromptUpdateParamsPromptDataOptions{
					Model: braintrust.F("model"),
					Params: braintrust.F[braintrust.PromptUpdateParamsPromptDataOptionsParamsUnion](braintrust.PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParams{
						UseCache:         braintrust.F(true),
						Temperature:      braintrust.F(0.000000),
						TopP:             braintrust.F(0.000000),
						MaxTokens:        braintrust.F(0.000000),
						FrequencyPenalty: braintrust.F(0.000000),
						PresencePenalty:  braintrust.F(0.000000),
						ResponseFormat: braintrust.F(braintrust.PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat{
							Type: braintrust.F(braintrust.PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject),
						}),
						ToolChoice:   braintrust.F[braintrust.PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion](braintrust.PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto(braintrust.PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto)),
						FunctionCall: braintrust.F[braintrust.PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion](braintrust.PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto(braintrust.PromptUpdateParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto)),
						N:            braintrust.F(0.000000),
						Stop:         braintrust.F([]string{"string", "string", "string"}),
					}),
					Position: braintrust.F("position"),
				}),
				Origin: braintrust.F(braintrust.PromptUpdateParamsPromptDataOrigin{
					PromptID:      braintrust.F("prompt_id"),
					ProjectID:     braintrust.F("project_id"),
					PromptVersion: braintrust.F("prompt_version"),
				}),
			}),
			Tags: braintrust.F([]string{"string", "string", "string"}),
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

func TestPromptListWithOptionalParams(t *testing.T) {
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
	_, err := client.Prompts.List(context.TODO(), braintrust.PromptListParams{
		EndingBefore:  braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IDs:           braintrust.F[braintrust.PromptListParamsIDsUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
		Limit:         braintrust.F(int64(0)),
		OrgName:       braintrust.F("org_name"),
		ProjectName:   braintrust.F("project_name"),
		PromptName:    braintrust.F("prompt_name"),
		Slug:          braintrust.F("slug"),
		StartingAfter: braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Version:       braintrust.F("version"),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPromptDelete(t *testing.T) {
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
	_, err := client.Prompts.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPromptReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Prompts.Replace(context.TODO(), braintrust.PromptReplaceParams{
		Name:        braintrust.F("name"),
		ProjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Slug:        braintrust.F("slug"),
		Description: braintrust.F("description"),
		PromptData: braintrust.F(braintrust.PromptReplaceParamsPromptData{
			Prompt: braintrust.F[braintrust.PromptReplaceParamsPromptDataPromptUnion](braintrust.PromptReplaceParamsPromptDataPromptCompletion{
				Type:    braintrust.F(braintrust.PromptReplaceParamsPromptDataPromptCompletionTypeCompletion),
				Content: braintrust.F("content"),
			}),
			Options: braintrust.F(braintrust.PromptReplaceParamsPromptDataOptions{
				Model: braintrust.F("model"),
				Params: braintrust.F[braintrust.PromptReplaceParamsPromptDataOptionsParamsUnion](braintrust.PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParams{
					UseCache:         braintrust.F(true),
					Temperature:      braintrust.F(0.000000),
					TopP:             braintrust.F(0.000000),
					MaxTokens:        braintrust.F(0.000000),
					FrequencyPenalty: braintrust.F(0.000000),
					PresencePenalty:  braintrust.F(0.000000),
					ResponseFormat: braintrust.F(braintrust.PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormat{
						Type: braintrust.F(braintrust.PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject),
					}),
					ToolChoice:   braintrust.F[braintrust.PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceUnion](braintrust.PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto(braintrust.PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto)),
					FunctionCall: braintrust.F[braintrust.PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallUnion](braintrust.PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto(braintrust.PromptReplaceParamsPromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto)),
					N:            braintrust.F(0.000000),
					Stop:         braintrust.F([]string{"string", "string", "string"}),
				}),
				Position: braintrust.F("position"),
			}),
			Origin: braintrust.F(braintrust.PromptReplaceParamsPromptDataOrigin{
				PromptID:      braintrust.F("prompt_id"),
				ProjectID:     braintrust.F("project_id"),
				PromptVersion: braintrust.F("prompt_version"),
			}),
		}),
		Tags: braintrust.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
