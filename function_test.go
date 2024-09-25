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
	_, err := client.Functions.New(context.TODO(), braintrust.FunctionNewParams{
		FunctionData: braintrust.F[braintrust.FunctionNewParamsFunctionDataUnion](braintrust.FunctionNewParamsFunctionDataPrompt{
			Type: braintrust.F(braintrust.FunctionNewParamsFunctionDataPromptTypePrompt),
		}),
		Name:        braintrust.F("name"),
		ProjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Slug:        braintrust.F("slug"),
		Description: braintrust.F("description"),
		FunctionSchema: braintrust.F(braintrust.FunctionNewParamsFunctionSchema{
			Parameters: braintrust.F[any](map[string]interface{}{}),
			Returns:    braintrust.F[any](map[string]interface{}{}),
		}),
		FunctionType: braintrust.F(braintrust.FunctionNewParamsFunctionTypeLlm),
		Origin: braintrust.F(braintrust.FunctionNewParamsOrigin{
			ObjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ObjectType: braintrust.F(braintrust.FunctionNewParamsOriginObjectTypeOrganization),
			Internal:   braintrust.F(true),
		}),
		PromptData: braintrust.F(shared.PromptDataParam{
			Options: braintrust.F(shared.PromptDataOptionsParam{
				Model: braintrust.F("model"),
				Params: braintrust.F[shared.PromptDataOptionsParamsUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsParam{
					FrequencyPenalty: braintrust.F(0.000000),
					FunctionCall:     braintrust.F[shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto(shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto)),
					MaxTokens:        braintrust.F(0.000000),
					N:                braintrust.F(0.000000),
					PresencePenalty:  braintrust.F(0.000000),
					ResponseFormat: braintrust.F(shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam{
						Type: braintrust.F(shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject),
					}),
					Stop:        braintrust.F([]string{"string", "string", "string"}),
					Temperature: braintrust.F(0.000000),
					ToolChoice:  braintrust.F[shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto(shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto)),
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
			}, shared.PromptDataToolFunctionsFunctionParam{
				ID:   braintrust.F("id"),
				Type: braintrust.F(shared.PromptDataToolFunctionsFunctionTypeFunction),
			}, shared.PromptDataToolFunctionsFunctionParam{
				ID:   braintrust.F("id"),
				Type: braintrust.F(shared.PromptDataToolFunctionsFunctionTypeFunction),
			}}),
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
	_, err := client.Functions.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
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
	_, err := client.Functions.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.FunctionUpdateParams{
			Description: braintrust.F("description"),
			FunctionData: braintrust.F[braintrust.FunctionUpdateParamsFunctionDataUnion](braintrust.FunctionUpdateParamsFunctionDataPrompt{
				Type: braintrust.F(braintrust.FunctionUpdateParamsFunctionDataPromptTypePrompt),
			}),
			Name: braintrust.F("name"),
			PromptData: braintrust.F(shared.PromptDataParam{
				Options: braintrust.F(shared.PromptDataOptionsParam{
					Model: braintrust.F("model"),
					Params: braintrust.F[shared.PromptDataOptionsParamsUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsParam{
						FrequencyPenalty: braintrust.F(0.000000),
						FunctionCall:     braintrust.F[shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto(shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto)),
						MaxTokens:        braintrust.F(0.000000),
						N:                braintrust.F(0.000000),
						PresencePenalty:  braintrust.F(0.000000),
						ResponseFormat: braintrust.F(shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam{
							Type: braintrust.F(shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject),
						}),
						Stop:        braintrust.F([]string{"string", "string", "string"}),
						Temperature: braintrust.F(0.000000),
						ToolChoice:  braintrust.F[shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto(shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto)),
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
				}, shared.PromptDataToolFunctionsFunctionParam{
					ID:   braintrust.F("id"),
					Type: braintrust.F(shared.PromptDataToolFunctionsFunctionTypeFunction),
				}, shared.PromptDataToolFunctionsFunctionParam{
					ID:   braintrust.F("id"),
					Type: braintrust.F(shared.PromptDataToolFunctionsFunctionTypeFunction),
				}}),
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
	_, err := client.Functions.List(context.TODO(), braintrust.FunctionListParams{
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
	_, err := client.Functions.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *braintrust.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFunctionInvokeWithOptionalParams(t *testing.T) {
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
	_, err := client.Functions.Invoke(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.FunctionInvokeParams{
			Input: braintrust.F[any](map[string]interface{}{}),
			Messages: braintrust.F([]shared.MessagesUnionParam{shared.MessagesSystemParam{
				Role:    braintrust.F(shared.MessagesSystemRoleSystem),
				Content: braintrust.F("content"),
				Name:    braintrust.F("name"),
			}, shared.MessagesSystemParam{
				Role:    braintrust.F(shared.MessagesSystemRoleSystem),
				Content: braintrust.F("content"),
				Name:    braintrust.F("name"),
			}, shared.MessagesSystemParam{
				Role:    braintrust.F(shared.MessagesSystemRoleSystem),
				Content: braintrust.F("content"),
				Name:    braintrust.F("name"),
			}}),
			Mode: braintrust.F(braintrust.FunctionInvokeParamsModeAuto),
			Parent: braintrust.F[braintrust.FunctionInvokeParamsParentUnion](braintrust.FunctionInvokeParamsParentSpanParentStruct{
				ObjectID:   braintrust.F("object_id"),
				ObjectType: braintrust.F(braintrust.FunctionInvokeParamsParentSpanParentStructObjectTypeProjectLogs),
				PropagatedEvent: braintrust.F(map[string]interface{}{
					"foo": "bar",
				}),
				RowIDs: braintrust.F(braintrust.FunctionInvokeParamsParentSpanParentStructRowIDs{
					ID:         braintrust.F("id"),
					RootSpanID: braintrust.F("root_span_id"),
					SpanID:     braintrust.F("span_id"),
				}),
			}),
			Stream:  braintrust.F(true),
			Version: braintrust.F("version"),
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
	_, err := client.Functions.Replace(context.TODO(), braintrust.FunctionReplaceParams{
		FunctionData: braintrust.F[braintrust.FunctionReplaceParamsFunctionDataUnion](braintrust.FunctionReplaceParamsFunctionDataPrompt{
			Type: braintrust.F(braintrust.FunctionReplaceParamsFunctionDataPromptTypePrompt),
		}),
		Name:        braintrust.F("name"),
		ProjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Slug:        braintrust.F("slug"),
		Description: braintrust.F("description"),
		FunctionSchema: braintrust.F(braintrust.FunctionReplaceParamsFunctionSchema{
			Parameters: braintrust.F[any](map[string]interface{}{}),
			Returns:    braintrust.F[any](map[string]interface{}{}),
		}),
		FunctionType: braintrust.F(braintrust.FunctionReplaceParamsFunctionTypeLlm),
		Origin: braintrust.F(braintrust.FunctionReplaceParamsOrigin{
			ObjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ObjectType: braintrust.F(braintrust.FunctionReplaceParamsOriginObjectTypeOrganization),
			Internal:   braintrust.F(true),
		}),
		PromptData: braintrust.F(shared.PromptDataParam{
			Options: braintrust.F(shared.PromptDataOptionsParam{
				Model: braintrust.F("model"),
				Params: braintrust.F[shared.PromptDataOptionsParamsUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsParam{
					FrequencyPenalty: braintrust.F(0.000000),
					FunctionCall:     braintrust.F[shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAuto(shared.PromptDataOptionsParamsOpenAIModelParamsFunctionCallAutoAuto)),
					MaxTokens:        braintrust.F(0.000000),
					N:                braintrust.F(0.000000),
					PresencePenalty:  braintrust.F(0.000000),
					ResponseFormat: braintrust.F(shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormatParam{
						Type: braintrust.F(shared.PromptDataOptionsParamsOpenAIModelParamsResponseFormatTypeJsonObject),
					}),
					Stop:        braintrust.F([]string{"string", "string", "string"}),
					Temperature: braintrust.F(0.000000),
					ToolChoice:  braintrust.F[shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceUnionParam](shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAuto(shared.PromptDataOptionsParamsOpenAIModelParamsToolChoiceAutoAuto)),
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
			}, shared.PromptDataToolFunctionsFunctionParam{
				ID:   braintrust.F("id"),
				Type: braintrust.F(shared.PromptDataToolFunctionsFunctionTypeFunction),
			}, shared.PromptDataToolFunctionsFunctionParam{
				ID:   braintrust.F("id"),
				Type: braintrust.F(shared.PromptDataToolFunctionsFunctionTypeFunction),
			}}),
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
