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
		Name:         braintrust.F("x"),
		ProjectID:    braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Slug:         braintrust.F("x"),
		Description:  braintrust.F("description"),
		FunctionType: braintrust.F(braintrust.PromptNewParamsFunctionTypeLlm),
		PromptData: braintrust.F(shared.PromptDataParam{
			Options: braintrust.F(shared.PromptOptionsParam{
				Model: braintrust.F("model"),
				Params: braintrust.F[shared.PromptOptionsParamsUnionParam](shared.PromptOptionsParamsOpenAIModelParamsParam{
					FrequencyPenalty:    braintrust.F(0.000000),
					FunctionCall:        braintrust.F[shared.PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam](shared.PromptOptionsParamsOpenAIModelParamsFunctionCallString(shared.PromptOptionsParamsOpenAIModelParamsFunctionCallStringAuto)),
					MaxCompletionTokens: braintrust.F(0.000000),
					MaxTokens:           braintrust.F(0.000000),
					N:                   braintrust.F(0.000000),
					PresencePenalty:     braintrust.F(0.000000),
					ReasoningEffort:     braintrust.F(shared.PromptOptionsParamsOpenAIModelParamsReasoningEffortLow),
					ResponseFormat: braintrust.F[shared.PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam](shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectParam{
						Type: braintrust.F(shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectTypeJsonObject),
					}),
					Stop:        braintrust.F([]string{"string"}),
					Temperature: braintrust.F(0.000000),
					ToolChoice:  braintrust.F[shared.PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam](shared.PromptOptionsParamsOpenAIModelParamsToolChoiceString(shared.PromptOptionsParamsOpenAIModelParamsToolChoiceStringAuto)),
					TopP:        braintrust.F(0.000000),
					UseCache:    braintrust.F(true),
				}),
				Position: braintrust.F("position"),
			}),
			Origin: braintrust.F(shared.PromptDataOriginParam{
				ProjectID:     braintrust.F("project_id"),
				PromptID:      braintrust.F("prompt_id"),
				PromptVersion: braintrust.F("prompt_version"),
			}),
			Parser: braintrust.F(shared.PromptDataParserParam{
				ChoiceScores: braintrust.F(map[string]float64{
					"foo": 0.000000,
				}),
				Type:   braintrust.F(shared.PromptDataParserTypeLlmClassifier),
				UseCot: braintrust.F(true),
			}),
			Prompt: braintrust.F[shared.PromptDataPromptUnionParam](shared.PromptDataPromptCompletionParam{
				Content: braintrust.F("content"),
				Type:    braintrust.F(shared.PromptDataPromptCompletionTypeCompletion),
			}),
			ToolFunctions: braintrust.F([]shared.PromptDataToolFunctionsUnionParam{shared.PromptDataToolFunctionsFunctionParam{
				ID:   braintrust.F("id"),
				Type: braintrust.F(shared.PromptDataToolFunctionsFunctionTypeFunction),
			}}),
		}),
		Tags: braintrust.F([]string{"string"}),
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
			PromptData: braintrust.F(shared.PromptDataParam{
				Options: braintrust.F(shared.PromptOptionsParam{
					Model: braintrust.F("model"),
					Params: braintrust.F[shared.PromptOptionsParamsUnionParam](shared.PromptOptionsParamsOpenAIModelParamsParam{
						FrequencyPenalty:    braintrust.F(0.000000),
						FunctionCall:        braintrust.F[shared.PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam](shared.PromptOptionsParamsOpenAIModelParamsFunctionCallString(shared.PromptOptionsParamsOpenAIModelParamsFunctionCallStringAuto)),
						MaxCompletionTokens: braintrust.F(0.000000),
						MaxTokens:           braintrust.F(0.000000),
						N:                   braintrust.F(0.000000),
						PresencePenalty:     braintrust.F(0.000000),
						ReasoningEffort:     braintrust.F(shared.PromptOptionsParamsOpenAIModelParamsReasoningEffortLow),
						ResponseFormat: braintrust.F[shared.PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam](shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectParam{
							Type: braintrust.F(shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectTypeJsonObject),
						}),
						Stop:        braintrust.F([]string{"string"}),
						Temperature: braintrust.F(0.000000),
						ToolChoice:  braintrust.F[shared.PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam](shared.PromptOptionsParamsOpenAIModelParamsToolChoiceString(shared.PromptOptionsParamsOpenAIModelParamsToolChoiceStringAuto)),
						TopP:        braintrust.F(0.000000),
						UseCache:    braintrust.F(true),
					}),
					Position: braintrust.F("position"),
				}),
				Origin: braintrust.F(shared.PromptDataOriginParam{
					ProjectID:     braintrust.F("project_id"),
					PromptID:      braintrust.F("prompt_id"),
					PromptVersion: braintrust.F("prompt_version"),
				}),
				Parser: braintrust.F(shared.PromptDataParserParam{
					ChoiceScores: braintrust.F(map[string]float64{
						"foo": 0.000000,
					}),
					Type:   braintrust.F(shared.PromptDataParserTypeLlmClassifier),
					UseCot: braintrust.F(true),
				}),
				Prompt: braintrust.F[shared.PromptDataPromptUnionParam](shared.PromptDataPromptCompletionParam{
					Content: braintrust.F("content"),
					Type:    braintrust.F(shared.PromptDataPromptCompletionTypeCompletion),
				}),
				ToolFunctions: braintrust.F([]shared.PromptDataToolFunctionsUnionParam{shared.PromptDataToolFunctionsFunctionParam{
					ID:   braintrust.F("id"),
					Type: braintrust.F(shared.PromptDataToolFunctionsFunctionTypeFunction),
				}}),
			}),
			Slug: braintrust.F("slug"),
			Tags: braintrust.F([]string{"string"}),
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
		ProjectID:     braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
		Name:         braintrust.F("x"),
		ProjectID:    braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Slug:         braintrust.F("x"),
		Description:  braintrust.F("description"),
		FunctionType: braintrust.F(braintrust.PromptReplaceParamsFunctionTypeLlm),
		PromptData: braintrust.F(shared.PromptDataParam{
			Options: braintrust.F(shared.PromptOptionsParam{
				Model: braintrust.F("model"),
				Params: braintrust.F[shared.PromptOptionsParamsUnionParam](shared.PromptOptionsParamsOpenAIModelParamsParam{
					FrequencyPenalty:    braintrust.F(0.000000),
					FunctionCall:        braintrust.F[shared.PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam](shared.PromptOptionsParamsOpenAIModelParamsFunctionCallString(shared.PromptOptionsParamsOpenAIModelParamsFunctionCallStringAuto)),
					MaxCompletionTokens: braintrust.F(0.000000),
					MaxTokens:           braintrust.F(0.000000),
					N:                   braintrust.F(0.000000),
					PresencePenalty:     braintrust.F(0.000000),
					ReasoningEffort:     braintrust.F(shared.PromptOptionsParamsOpenAIModelParamsReasoningEffortLow),
					ResponseFormat: braintrust.F[shared.PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam](shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectParam{
						Type: braintrust.F(shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectTypeJsonObject),
					}),
					Stop:        braintrust.F([]string{"string"}),
					Temperature: braintrust.F(0.000000),
					ToolChoice:  braintrust.F[shared.PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam](shared.PromptOptionsParamsOpenAIModelParamsToolChoiceString(shared.PromptOptionsParamsOpenAIModelParamsToolChoiceStringAuto)),
					TopP:        braintrust.F(0.000000),
					UseCache:    braintrust.F(true),
				}),
				Position: braintrust.F("position"),
			}),
			Origin: braintrust.F(shared.PromptDataOriginParam{
				ProjectID:     braintrust.F("project_id"),
				PromptID:      braintrust.F("prompt_id"),
				PromptVersion: braintrust.F("prompt_version"),
			}),
			Parser: braintrust.F(shared.PromptDataParserParam{
				ChoiceScores: braintrust.F(map[string]float64{
					"foo": 0.000000,
				}),
				Type:   braintrust.F(shared.PromptDataParserTypeLlmClassifier),
				UseCot: braintrust.F(true),
			}),
			Prompt: braintrust.F[shared.PromptDataPromptUnionParam](shared.PromptDataPromptCompletionParam{
				Content: braintrust.F("content"),
				Type:    braintrust.F(shared.PromptDataPromptCompletionTypeCompletion),
			}),
			ToolFunctions: braintrust.F([]shared.PromptDataToolFunctionsUnionParam{shared.PromptDataToolFunctionsFunctionParam{
				ID:   braintrust.F("id"),
				Type: braintrust.F(shared.PromptDataToolFunctionsFunctionTypeFunction),
			}}),
		}),
		Tags: braintrust.F([]string{"string"}),
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
