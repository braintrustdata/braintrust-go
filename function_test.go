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

func TestFunctionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Function.New(context.TODO(), braintrust.FunctionNewParams{
		FunctionData: braintrust.F[braintrust.FunctionNewParamsFunctionDataUnion](braintrust.FunctionNewParamsFunctionDataPrompt{
			Type: braintrust.F(braintrust.FunctionNewParamsFunctionDataPromptTypePrompt),
		}),
		Name:         braintrust.F("name"),
		ProjectID:    braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Slug:         braintrust.F("slug"),
		Description:  braintrust.F("description"),
		FunctionType: braintrust.F(braintrust.FunctionNewParamsFunctionTypeTask),
		Origin: braintrust.F(braintrust.FunctionNewParamsOrigin{
			ObjectType: braintrust.F(braintrust.FunctionNewParamsOriginObjectTypeOrganization),
			ObjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Internal:   braintrust.F(true),
		}),
		PromptData: braintrust.F(shared.PromptDataParam{
			Prompt: braintrust.F[shared.PromptDataPromptUnionParam](shared.PromptDataPromptCompletionParam{
				Type:    braintrust.F(shared.PromptDataPromptCompletionTypeCompletion),
				Content: braintrust.F("content"),
			}),
			Options: braintrust.F(shared.PromptDataOptionsParam{
				Model: braintrust.F("model"),
				Params: braintrust.F[shared.PromptDataOptionsParamsUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsParam{
					UseCache:         braintrust.F(true),
					Temperature:      braintrust.F(0.000000),
					TopP:             braintrust.F(0.000000),
					MaxTokens:        braintrust.F(0.000000),
					FrequencyPenalty: braintrust.F(0.000000),
					PresencePenalty:  braintrust.F(0.000000),
					ResponseFormat: braintrust.F(shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam{
						Type: braintrust.F(shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject),
					}),
					ToolChoice:   braintrust.F[shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto(shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto)),
					FunctionCall: braintrust.F[shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto(shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto)),
					N:            braintrust.F(0.000000),
					Stop:         braintrust.F([]string{"string", "string", "string"}),
				}),
				Position: braintrust.F("position"),
			}),
			Parser: braintrust.F(shared.PromptDataParserParam{
				Type:   braintrust.F(shared.PromptDataParserTypeLlmClassifier),
				UseCot: braintrust.F(true),
				ChoiceScores: braintrust.F(map[string]float64{
					"foo": 0.000000,
				}),
			}),
			Origin: braintrust.F(shared.PromptDataOriginParam{
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

func TestFunctionGet(t *testing.T) {
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
	_, err := client.Function.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFunctionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Function.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.FunctionUpdateParams{
			Description: braintrust.F("description"),
			FunctionData: braintrust.F[braintrust.FunctionUpdateParamsFunctionDataUnion](braintrust.FunctionUpdateParamsFunctionDataPrompt{
				Type: braintrust.F(braintrust.FunctionUpdateParamsFunctionDataPromptTypePrompt),
			}),
			Name: braintrust.F("name"),
			PromptData: braintrust.F(shared.PromptDataParam{
				Prompt: braintrust.F[shared.PromptDataPromptUnionParam](shared.PromptDataPromptCompletionParam{
					Type:    braintrust.F(shared.PromptDataPromptCompletionTypeCompletion),
					Content: braintrust.F("content"),
				}),
				Options: braintrust.F(shared.PromptDataOptionsParam{
					Model: braintrust.F("model"),
					Params: braintrust.F[shared.PromptDataOptionsParamsUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsParam{
						UseCache:         braintrust.F(true),
						Temperature:      braintrust.F(0.000000),
						TopP:             braintrust.F(0.000000),
						MaxTokens:        braintrust.F(0.000000),
						FrequencyPenalty: braintrust.F(0.000000),
						PresencePenalty:  braintrust.F(0.000000),
						ResponseFormat: braintrust.F(shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam{
							Type: braintrust.F(shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject),
						}),
						ToolChoice:   braintrust.F[shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto(shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto)),
						FunctionCall: braintrust.F[shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto(shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto)),
						N:            braintrust.F(0.000000),
						Stop:         braintrust.F([]string{"string", "string", "string"}),
					}),
					Position: braintrust.F("position"),
				}),
				Parser: braintrust.F(shared.PromptDataParserParam{
					Type:   braintrust.F(shared.PromptDataParserTypeLlmClassifier),
					UseCot: braintrust.F(true),
					ChoiceScores: braintrust.F(map[string]float64{
						"foo": 0.000000,
					}),
				}),
				Origin: braintrust.F(shared.PromptDataOriginParam{
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

func TestFunctionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Function.List(context.TODO(), braintrust.FunctionListParams{
		EndingBefore:  braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		FunctionName:  braintrust.F("function_name"),
		IDs:           braintrust.F[braintrust.FunctionListParamsIDsUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
		Limit:         braintrust.F(int64(0)),
		OrgName:       braintrust.F("org_name"),
		ProjectID:     braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectName:   braintrust.F("project_name"),
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

func TestFunctionDelete(t *testing.T) {
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
	_, err := client.Function.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFunctionReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Function.Replace(context.TODO(), braintrust.FunctionReplaceParams{
		FunctionData: braintrust.F[braintrust.FunctionReplaceParamsFunctionDataUnion](braintrust.FunctionReplaceParamsFunctionDataPrompt{
			Type: braintrust.F(braintrust.FunctionReplaceParamsFunctionDataPromptTypePrompt),
		}),
		Name:         braintrust.F("name"),
		ProjectID:    braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Slug:         braintrust.F("slug"),
		Description:  braintrust.F("description"),
		FunctionType: braintrust.F(braintrust.FunctionReplaceParamsFunctionTypeTask),
		Origin: braintrust.F(braintrust.FunctionReplaceParamsOrigin{
			ObjectType: braintrust.F(braintrust.FunctionReplaceParamsOriginObjectTypeOrganization),
			ObjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Internal:   braintrust.F(true),
		}),
		PromptData: braintrust.F(shared.PromptDataParam{
			Prompt: braintrust.F[shared.PromptDataPromptUnionParam](shared.PromptDataPromptCompletionParam{
				Type:    braintrust.F(shared.PromptDataPromptCompletionTypeCompletion),
				Content: braintrust.F("content"),
			}),
			Options: braintrust.F(shared.PromptDataOptionsParam{
				Model: braintrust.F("model"),
				Params: braintrust.F[shared.PromptDataOptionsParamsUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsParam{
					UseCache:         braintrust.F(true),
					Temperature:      braintrust.F(0.000000),
					TopP:             braintrust.F(0.000000),
					MaxTokens:        braintrust.F(0.000000),
					FrequencyPenalty: braintrust.F(0.000000),
					PresencePenalty:  braintrust.F(0.000000),
					ResponseFormat: braintrust.F(shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam{
						Type: braintrust.F(shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject),
					}),
					ToolChoice:   braintrust.F[shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto(shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto)),
					FunctionCall: braintrust.F[shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto(shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto)),
					N:            braintrust.F(0.000000),
					Stop:         braintrust.F([]string{"string", "string", "string"}),
				}),
				Position: braintrust.F("position"),
			}),
			Parser: braintrust.F(shared.PromptDataParserParam{
				Type:   braintrust.F(shared.PromptDataParserTypeLlmClassifier),
				UseCot: braintrust.F(true),
				ChoiceScores: braintrust.F(map[string]float64{
					"foo": 0.000000,
				}),
			}),
			Origin: braintrust.F(shared.PromptDataOriginParam{
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
