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
		Name:         "x",
		ProjectID:    "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		Slug:         "x",
		Description:  braintrust.String("description"),
		FunctionType: braintrust.PromptNewParamsFunctionTypeLlm,
		PromptData: shared.PromptDataParam{
			Options: shared.PromptOptionsParam{
				Model: braintrust.String("model"),
				Params: shared.PromptOptionsParamsUnionParam{
					OfOpenAIModels: &shared.PromptOptionsParamsOpenAIModelParamsParam{
						FrequencyPenalty: braintrust.Float(0),
						FunctionCall: shared.PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam{
							OfPromptOptionssOpenAIModelParamsFunctionCallString: braintrust.Opt(shared.PromptOptionsParamsOpenAIModelParamsFunctionCallStringAuto),
						},
						MaxCompletionTokens: braintrust.Float(0),
						MaxTokens:           braintrust.Float(0),
						N:                   braintrust.Float(0),
						PresencePenalty:     braintrust.Float(0),
						ReasoningEffort:     "low",
						ResponseFormat: shared.PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam{
							OfJsonObject: &shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectParam{
								Type: "json_object",
							},
						},
						Stop:        []string{"string"},
						Temperature: braintrust.Float(0),
						ToolChoice: shared.PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam{
							OfPromptOptionssOpenAIModelParamsToolChoiceString: braintrust.Opt(shared.PromptOptionsParamsOpenAIModelParamsToolChoiceStringAuto),
						},
						TopP:     braintrust.Float(0),
						UseCache: braintrust.Bool(true),
					},
				},
				Position: braintrust.String("position"),
			},
			Origin: shared.PromptDataOriginParam{
				ProjectID:     braintrust.String("project_id"),
				PromptID:      braintrust.String("prompt_id"),
				PromptVersion: braintrust.String("prompt_version"),
			},
			Parser: shared.PromptDataParserParam{
				ChoiceScores: map[string]float64{
					"foo": 0,
				},
				Type:   "llm_classifier",
				UseCot: true,
			},
			Prompt: shared.PromptDataPromptUnionParam{
				OfCompletion: &shared.PromptDataPromptCompletionParam{
					Content: "content",
					Type:    "completion",
				},
			},
			ToolFunctions: []shared.PromptDataToolFunctionUnionParam{{
				OfFunction: &shared.PromptDataToolFunctionFunctionParam{
					ID:   "id",
					Type: "function",
				},
			}},
		},
		Tags: []string{"string"},
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
			Description: braintrust.String("description"),
			Name:        braintrust.String("name"),
			PromptData: shared.PromptDataParam{
				Options: shared.PromptOptionsParam{
					Model: braintrust.String("model"),
					Params: shared.PromptOptionsParamsUnionParam{
						OfOpenAIModels: &shared.PromptOptionsParamsOpenAIModelParamsParam{
							FrequencyPenalty: braintrust.Float(0),
							FunctionCall: shared.PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam{
								OfPromptOptionssOpenAIModelParamsFunctionCallString: braintrust.Opt(shared.PromptOptionsParamsOpenAIModelParamsFunctionCallStringAuto),
							},
							MaxCompletionTokens: braintrust.Float(0),
							MaxTokens:           braintrust.Float(0),
							N:                   braintrust.Float(0),
							PresencePenalty:     braintrust.Float(0),
							ReasoningEffort:     "low",
							ResponseFormat: shared.PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam{
								OfJsonObject: &shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectParam{
									Type: "json_object",
								},
							},
							Stop:        []string{"string"},
							Temperature: braintrust.Float(0),
							ToolChoice: shared.PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam{
								OfPromptOptionssOpenAIModelParamsToolChoiceString: braintrust.Opt(shared.PromptOptionsParamsOpenAIModelParamsToolChoiceStringAuto),
							},
							TopP:     braintrust.Float(0),
							UseCache: braintrust.Bool(true),
						},
					},
					Position: braintrust.String("position"),
				},
				Origin: shared.PromptDataOriginParam{
					ProjectID:     braintrust.String("project_id"),
					PromptID:      braintrust.String("prompt_id"),
					PromptVersion: braintrust.String("prompt_version"),
				},
				Parser: shared.PromptDataParserParam{
					ChoiceScores: map[string]float64{
						"foo": 0,
					},
					Type:   "llm_classifier",
					UseCot: true,
				},
				Prompt: shared.PromptDataPromptUnionParam{
					OfCompletion: &shared.PromptDataPromptCompletionParam{
						Content: "content",
						Type:    "completion",
					},
				},
				ToolFunctions: []shared.PromptDataToolFunctionUnionParam{{
					OfFunction: &shared.PromptDataToolFunctionFunctionParam{
						ID:   "id",
						Type: "function",
					},
				}},
			},
			Slug: braintrust.String("slug"),
			Tags: []string{"string"},
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
		EndingBefore: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IDs: braintrust.PromptListParamsIDsUnion{
			OfString: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
		Limit:         braintrust.Int(0),
		OrgName:       braintrust.String("org_name"),
		ProjectID:     braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectName:   braintrust.String("project_name"),
		PromptName:    braintrust.String("prompt_name"),
		Slug:          braintrust.String("slug"),
		StartingAfter: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Version:       braintrust.String("version"),
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
		Name:         "x",
		ProjectID:    "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		Slug:         "x",
		Description:  braintrust.String("description"),
		FunctionType: braintrust.PromptReplaceParamsFunctionTypeLlm,
		PromptData: shared.PromptDataParam{
			Options: shared.PromptOptionsParam{
				Model: braintrust.String("model"),
				Params: shared.PromptOptionsParamsUnionParam{
					OfOpenAIModels: &shared.PromptOptionsParamsOpenAIModelParamsParam{
						FrequencyPenalty: braintrust.Float(0),
						FunctionCall: shared.PromptOptionsParamsOpenAIModelParamsFunctionCallUnionParam{
							OfPromptOptionssOpenAIModelParamsFunctionCallString: braintrust.Opt(shared.PromptOptionsParamsOpenAIModelParamsFunctionCallStringAuto),
						},
						MaxCompletionTokens: braintrust.Float(0),
						MaxTokens:           braintrust.Float(0),
						N:                   braintrust.Float(0),
						PresencePenalty:     braintrust.Float(0),
						ReasoningEffort:     "low",
						ResponseFormat: shared.PromptOptionsParamsOpenAIModelParamsResponseFormatUnionParam{
							OfJsonObject: &shared.PromptOptionsParamsOpenAIModelParamsResponseFormatJsonObjectParam{
								Type: "json_object",
							},
						},
						Stop:        []string{"string"},
						Temperature: braintrust.Float(0),
						ToolChoice: shared.PromptOptionsParamsOpenAIModelParamsToolChoiceUnionParam{
							OfPromptOptionssOpenAIModelParamsToolChoiceString: braintrust.Opt(shared.PromptOptionsParamsOpenAIModelParamsToolChoiceStringAuto),
						},
						TopP:     braintrust.Float(0),
						UseCache: braintrust.Bool(true),
					},
				},
				Position: braintrust.String("position"),
			},
			Origin: shared.PromptDataOriginParam{
				ProjectID:     braintrust.String("project_id"),
				PromptID:      braintrust.String("prompt_id"),
				PromptVersion: braintrust.String("prompt_version"),
			},
			Parser: shared.PromptDataParserParam{
				ChoiceScores: map[string]float64{
					"foo": 0,
				},
				Type:   "llm_classifier",
				UseCot: true,
			},
			Prompt: shared.PromptDataPromptUnionParam{
				OfCompletion: &shared.PromptDataPromptCompletionParam{
					Content: "content",
					Type:    "completion",
				},
			},
			ToolFunctions: []shared.PromptDataToolFunctionUnionParam{{
				OfFunction: &shared.PromptDataToolFunctionFunctionParam{
					ID:   "id",
					Type: "function",
				},
			}},
		},
		Tags: []string{"string"},
	})
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
