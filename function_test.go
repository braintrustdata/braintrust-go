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
		FunctionData: braintrust.FunctionNewParamsFunctionDataUnion{
			OfPrompt: &braintrust.FunctionNewParamsFunctionDataPrompt{
				Type: "prompt",
			},
		},
		Name:        "x",
		ProjectID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		Slug:        "x",
		Description: braintrust.String("description"),
		FunctionSchema: braintrust.FunctionNewParamsFunctionSchema{
			Parameters: map[string]interface{}{},
			Returns:    map[string]interface{}{},
		},
		FunctionType: braintrust.FunctionNewParamsFunctionTypeLlm,
		Origin: braintrust.FunctionNewParamsOrigin{
			ObjectID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ObjectType: shared.ACLObjectTypeOrganization,
			Internal:   braintrust.Bool(true),
		},
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
			Description: braintrust.String("description"),
			FunctionData: braintrust.FunctionUpdateParamsFunctionDataUnion{
				OfPrompt: &braintrust.FunctionUpdateParamsFunctionDataPrompt{
					Type: "prompt",
				},
			},
			Name: braintrust.String("name"),
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
		EndingBefore: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		FunctionName: braintrust.String("function_name"),
		IDs: braintrust.FunctionListParamsIDsUnion{
			OfString: braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
		Limit:         braintrust.Int(0),
		OrgName:       braintrust.String("org_name"),
		ProjectID:     braintrust.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectName:   braintrust.String("project_name"),
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
			Expected: map[string]interface{}{},
			Input:    map[string]interface{}{},
			Messages: []braintrust.FunctionInvokeParamsMessageUnion{{
				OfSystem: &braintrust.FunctionInvokeParamsMessageSystem{
					Role:    "system",
					Content: braintrust.String("content"),
					Name:    braintrust.String("name"),
				},
			}},
			Metadata: map[string]interface{}{
				"foo": "bar",
			},
			Mode: braintrust.FunctionInvokeParamsModeAuto,
			Parent: braintrust.FunctionInvokeParamsParentUnion{
				OfSpanParentStruct: &braintrust.FunctionInvokeParamsParentSpanParentStruct{
					ObjectID:   "object_id",
					ObjectType: "project_logs",
					PropagatedEvent: map[string]interface{}{
						"foo": "bar",
					},
					RowIDs: braintrust.FunctionInvokeParamsParentSpanParentStructRowIDs{
						ID:         "id",
						RootSpanID: "root_span_id",
						SpanID:     "span_id",
					},
				},
			},
			Stream:  braintrust.Bool(true),
			Version: braintrust.String("version"),
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
		FunctionData: braintrust.FunctionReplaceParamsFunctionDataUnion{
			OfPrompt: &braintrust.FunctionReplaceParamsFunctionDataPrompt{
				Type: "prompt",
			},
		},
		Name:        "x",
		ProjectID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		Slug:        "x",
		Description: braintrust.String("description"),
		FunctionSchema: braintrust.FunctionReplaceParamsFunctionSchema{
			Parameters: map[string]interface{}{},
			Returns:    map[string]interface{}{},
		},
		FunctionType: braintrust.FunctionReplaceParamsFunctionTypeLlm,
		Origin: braintrust.FunctionReplaceParamsOrigin{
			ObjectID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ObjectType: shared.ACLObjectTypeOrganization,
			Internal:   braintrust.Bool(true),
		},
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
