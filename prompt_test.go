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
			Prompt: braintrust.F[braintrust.PromptNewParamsPromptDataPromptUnion](braintrust.PromptNewParamsPromptDataPromptPromptDataPrompt0{
				Type:    braintrust.F(braintrust.PromptNewParamsPromptDataPromptPromptDataPrompt0TypeCompletion),
				Content: braintrust.F("content"),
			}),
			Options: braintrust.F(braintrust.PromptNewParamsPromptDataOptions{
				Model: braintrust.F("model"),
				Params: braintrust.F[braintrust.PromptNewParamsPromptDataOptionsParamsUnion](braintrust.PromptNewParamsPromptDataOptionsParamsPromptDataOptions0{
					UseCache:         braintrust.F(true),
					Temperature:      braintrust.F(0.000000),
					TopP:             braintrust.F(0.000000),
					MaxTokens:        braintrust.F(0.000000),
					FrequencyPenalty: braintrust.F(0.000000),
					PresencePenalty:  braintrust.F(0.000000),
					ResponseFormat: braintrust.F(braintrust.PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat{
						Type: braintrust.F(braintrust.PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject),
					}),
					ToolChoice:   braintrust.F[braintrust.PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion](braintrust.PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0(braintrust.PromptNewParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto)),
					FunctionCall: braintrust.F[braintrust.PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion](braintrust.PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0(braintrust.PromptNewParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto)),
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
				Prompt: braintrust.F[braintrust.PromptUpdateParamsPromptDataPromptUnion](braintrust.PromptUpdateParamsPromptDataPromptPromptDataPrompt0{
					Type:    braintrust.F(braintrust.PromptUpdateParamsPromptDataPromptPromptDataPrompt0TypeCompletion),
					Content: braintrust.F("content"),
				}),
				Options: braintrust.F(braintrust.PromptUpdateParamsPromptDataOptions{
					Model: braintrust.F("model"),
					Params: braintrust.F[braintrust.PromptUpdateParamsPromptDataOptionsParamsUnion](braintrust.PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0{
						UseCache:         braintrust.F(true),
						Temperature:      braintrust.F(0.000000),
						TopP:             braintrust.F(0.000000),
						MaxTokens:        braintrust.F(0.000000),
						FrequencyPenalty: braintrust.F(0.000000),
						PresencePenalty:  braintrust.F(0.000000),
						ResponseFormat: braintrust.F(braintrust.PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat{
							Type: braintrust.F(braintrust.PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject),
						}),
						ToolChoice:   braintrust.F[braintrust.PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion](braintrust.PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0(braintrust.PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto)),
						FunctionCall: braintrust.F[braintrust.PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion](braintrust.PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0(braintrust.PromptUpdateParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto)),
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
			Prompt: braintrust.F[braintrust.PromptReplaceParamsPromptDataPromptUnion](braintrust.PromptReplaceParamsPromptDataPromptPromptDataPrompt0{
				Type:    braintrust.F(braintrust.PromptReplaceParamsPromptDataPromptPromptDataPrompt0TypeCompletion),
				Content: braintrust.F("content"),
			}),
			Options: braintrust.F(braintrust.PromptReplaceParamsPromptDataOptions{
				Model: braintrust.F("model"),
				Params: braintrust.F[braintrust.PromptReplaceParamsPromptDataOptionsParamsUnion](braintrust.PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0{
					UseCache:         braintrust.F(true),
					Temperature:      braintrust.F(0.000000),
					TopP:             braintrust.F(0.000000),
					MaxTokens:        braintrust.F(0.000000),
					FrequencyPenalty: braintrust.F(0.000000),
					PresencePenalty:  braintrust.F(0.000000),
					ResponseFormat: braintrust.F(braintrust.PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormat{
						Type: braintrust.F(braintrust.PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ResponseFormatTypeJsonObject),
					}),
					ToolChoice:   braintrust.F[braintrust.PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoiceUnion](braintrust.PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0(braintrust.PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0ToolChoicePromptDataToolChoice0Auto)),
					FunctionCall: braintrust.F[braintrust.PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallUnion](braintrust.PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0(braintrust.PromptReplaceParamsPromptDataOptionsParamsPromptDataOptions0FunctionCallPromptDataFunctionCall0Auto)),
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
