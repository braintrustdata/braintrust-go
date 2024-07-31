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
		PromptData: braintrust.F(braintrust.PromptDataParam{
			Prompt: braintrust.F[braintrust.PromptDataPromptUnionParam](braintrust.PromptDataPromptCompletionParam{
				Type:    braintrust.F(braintrust.PromptDataPromptCompletionTypeCompletion),
				Content: braintrust.F("content"),
			}),
			Options: braintrust.F(braintrust.PromptDataOptionsParam{
				Model: braintrust.F("model"),
				Params: braintrust.F[braintrust.PromptDataOptionsParamsUnionParam](braintrust.PromptDataOptionsParamsOpenAIModelParamsParam{
					UseCache:         braintrust.F(true),
					Temperature:      braintrust.F(0.000000),
					TopP:             braintrust.F(0.000000),
					MaxTokens:        braintrust.F(0.000000),
					FrequencyPenalty: braintrust.F(0.000000),
					PresencePenalty:  braintrust.F(0.000000),
					ResponseFormat: braintrust.F(braintrust.PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam{
						Type: braintrust.F(braintrust.PromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject),
					}),
					ToolChoice:   braintrust.F[braintrust.PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam](braintrust.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto(braintrust.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto)),
					FunctionCall: braintrust.F[braintrust.PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam](braintrust.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto(braintrust.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto)),
					N:            braintrust.F(0.000000),
					Stop:         braintrust.F([]string{"string", "string", "string"}),
				}),
				Position: braintrust.F("position"),
			}),
			Origin: braintrust.F(braintrust.PromptDataOriginParam{
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
			PromptData: braintrust.F(braintrust.PromptDataParam{
				Prompt: braintrust.F[braintrust.PromptDataPromptUnionParam](braintrust.PromptDataPromptCompletionParam{
					Type:    braintrust.F(braintrust.PromptDataPromptCompletionTypeCompletion),
					Content: braintrust.F("content"),
				}),
				Options: braintrust.F(braintrust.PromptDataOptionsParam{
					Model: braintrust.F("model"),
					Params: braintrust.F[braintrust.PromptDataOptionsParamsUnionParam](braintrust.PromptDataOptionsParamsOpenAIModelParamsParam{
						UseCache:         braintrust.F(true),
						Temperature:      braintrust.F(0.000000),
						TopP:             braintrust.F(0.000000),
						MaxTokens:        braintrust.F(0.000000),
						FrequencyPenalty: braintrust.F(0.000000),
						PresencePenalty:  braintrust.F(0.000000),
						ResponseFormat: braintrust.F(braintrust.PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam{
							Type: braintrust.F(braintrust.PromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject),
						}),
						ToolChoice:   braintrust.F[braintrust.PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam](braintrust.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto(braintrust.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto)),
						FunctionCall: braintrust.F[braintrust.PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam](braintrust.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto(braintrust.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto)),
						N:            braintrust.F(0.000000),
						Stop:         braintrust.F([]string{"string", "string", "string"}),
					}),
					Position: braintrust.F("position"),
				}),
				Origin: braintrust.F(braintrust.PromptDataOriginParam{
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
		PromptData: braintrust.F(braintrust.PromptDataParam{
			Prompt: braintrust.F[braintrust.PromptDataPromptUnionParam](braintrust.PromptDataPromptCompletionParam{
				Type:    braintrust.F(braintrust.PromptDataPromptCompletionTypeCompletion),
				Content: braintrust.F("content"),
			}),
			Options: braintrust.F(braintrust.PromptDataOptionsParam{
				Model: braintrust.F("model"),
				Params: braintrust.F[braintrust.PromptDataOptionsParamsUnionParam](braintrust.PromptDataOptionsParamsOpenAIModelParamsParam{
					UseCache:         braintrust.F(true),
					Temperature:      braintrust.F(0.000000),
					TopP:             braintrust.F(0.000000),
					MaxTokens:        braintrust.F(0.000000),
					FrequencyPenalty: braintrust.F(0.000000),
					PresencePenalty:  braintrust.F(0.000000),
					ResponseFormat: braintrust.F(braintrust.PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam{
						Type: braintrust.F(braintrust.PromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject),
					}),
					ToolChoice:   braintrust.F[braintrust.PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam](braintrust.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto(braintrust.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto)),
					FunctionCall: braintrust.F[braintrust.PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam](braintrust.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto(braintrust.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto)),
					N:            braintrust.F(0.000000),
					Stop:         braintrust.F([]string{"string", "string", "string"}),
				}),
				Position: braintrust.F("position"),
			}),
			Origin: braintrust.F(braintrust.PromptDataOriginParam{
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
