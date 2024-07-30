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
		PromptData: braintrust.F(braintrust.FunctionNewParamsPromptData{
			Prompt: braintrust.F[braintrust.FunctionNewParamsPromptDataPromptUnion](braintrust.FunctionNewParamsPromptDataPromptPromptDataPrompt0{
				Type:    braintrust.F(braintrust.FunctionNewParamsPromptDataPromptPromptDataPrompt0TypeCompletion),
				Content: braintrust.F("content"),
			}),
			Options: braintrust.F(braintrust.FunctionNewParamsPromptDataOptions{
				Model: braintrust.F("model"),
				Params: braintrust.F[braintrust.FunctionNewParamsPromptDataOptionsParamsUnion](braintrust.FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0{
					UseCache:         braintrust.F(true),
					Temperature:      braintrust.F(0.000000),
					TopP:             braintrust.F(0.000000),
					MaxTokens:        braintrust.F(0.000000),
					FrequencyPenalty: braintrust.F(0.000000),
					PresencePenalty:  braintrust.F(0.000000),
					ResponseFormat: braintrust.F(braintrust.FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat{
						Type: braintrust.F(braintrust.FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject),
					}),
					ToolChoice:   braintrust.F[braintrust.FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion](braintrust.FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0(braintrust.FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto)),
					FunctionCall: braintrust.F[braintrust.FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion](braintrust.FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0(braintrust.FunctionNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto)),
					N:            braintrust.F(0.000000),
					Stop:         braintrust.F([]string{"string", "string", "string"}),
				}),
				Position: braintrust.F("position"),
			}),
			Origin: braintrust.F(braintrust.FunctionNewParamsPromptDataOrigin{
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
			PromptData: braintrust.F(braintrust.FunctionUpdateParamsPromptData{
				Prompt: braintrust.F[braintrust.FunctionUpdateParamsPromptDataPromptUnion](braintrust.FunctionUpdateParamsPromptDataPromptPromptDataPrompt0{
					Type:    braintrust.F(braintrust.FunctionUpdateParamsPromptDataPromptPromptDataPrompt0TypeCompletion),
					Content: braintrust.F("content"),
				}),
				Options: braintrust.F(braintrust.FunctionUpdateParamsPromptDataOptions{
					Model: braintrust.F("model"),
					Params: braintrust.F[braintrust.FunctionUpdateParamsPromptDataOptionsParamsUnion](braintrust.FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0{
						UseCache:         braintrust.F(true),
						Temperature:      braintrust.F(0.000000),
						TopP:             braintrust.F(0.000000),
						MaxTokens:        braintrust.F(0.000000),
						FrequencyPenalty: braintrust.F(0.000000),
						PresencePenalty:  braintrust.F(0.000000),
						ResponseFormat: braintrust.F(braintrust.FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat{
							Type: braintrust.F(braintrust.FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject),
						}),
						ToolChoice:   braintrust.F[braintrust.FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion](braintrust.FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0(braintrust.FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto)),
						FunctionCall: braintrust.F[braintrust.FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion](braintrust.FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0(braintrust.FunctionUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto)),
						N:            braintrust.F(0.000000),
						Stop:         braintrust.F([]string{"string", "string", "string"}),
					}),
					Position: braintrust.F("position"),
				}),
				Origin: braintrust.F(braintrust.FunctionUpdateParamsPromptDataOrigin{
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
	_, err := client.Functions.List(context.TODO(), braintrust.FunctionListParams{
		EndingBefore:  braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		FunctionName:  braintrust.F("function_name"),
		IDs:           braintrust.F[braintrust.FunctionListParamsIDsUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
		Limit:         braintrust.F(int64(0)),
		OrgName:       braintrust.F("org_name"),
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
		PromptData: braintrust.F(braintrust.FunctionReplaceParamsPromptData{
			Prompt: braintrust.F[braintrust.FunctionReplaceParamsPromptDataPromptUnion](braintrust.FunctionReplaceParamsPromptDataPromptPromptDataPrompt0{
				Type:    braintrust.F(braintrust.FunctionReplaceParamsPromptDataPromptPromptDataPrompt0TypeCompletion),
				Content: braintrust.F("content"),
			}),
			Options: braintrust.F(braintrust.FunctionReplaceParamsPromptDataOptions{
				Model: braintrust.F("model"),
				Params: braintrust.F[braintrust.FunctionReplaceParamsPromptDataOptionsParamsUnion](braintrust.FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0{
					UseCache:         braintrust.F(true),
					Temperature:      braintrust.F(0.000000),
					TopP:             braintrust.F(0.000000),
					MaxTokens:        braintrust.F(0.000000),
					FrequencyPenalty: braintrust.F(0.000000),
					PresencePenalty:  braintrust.F(0.000000),
					ResponseFormat: braintrust.F(braintrust.FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat{
						Type: braintrust.F(braintrust.FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject),
					}),
					ToolChoice:   braintrust.F[braintrust.FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion](braintrust.FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0(braintrust.FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto)),
					FunctionCall: braintrust.F[braintrust.FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion](braintrust.FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0(braintrust.FunctionReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto)),
					N:            braintrust.F(0.000000),
					Stop:         braintrust.F([]string{"string", "string", "string"}),
				}),
				Position: braintrust.F("position"),
			}),
			Origin: braintrust.F(braintrust.FunctionReplaceParamsPromptDataOrigin{
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
