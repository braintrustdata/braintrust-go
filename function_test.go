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
		FunctionData: braintrust.F[braintrust.FunctionNewParamsFunctionDataUnion](braintrust.FunctionNewParamsFunctionDataObject{
			Type: braintrust.F(braintrust.FunctionNewParamsFunctionDataObjectTypePrompt),
		}),
		Name:        braintrust.F("string"),
		ProjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Slug:        braintrust.F("string"),
		Description: braintrust.F("string"),
		PromptData: braintrust.F(braintrust.FunctionNewParamsPromptData{
			Prompt: braintrust.F[braintrust.FunctionNewParamsPromptDataPromptUnion](braintrust.FunctionNewParamsPromptDataPromptObject{
				Type:    braintrust.F(braintrust.FunctionNewParamsPromptDataPromptObjectTypeCompletion),
				Content: braintrust.F("string"),
			}),
			Options: braintrust.F(braintrust.FunctionNewParamsPromptDataOptions{
				Model: braintrust.F("string"),
				Params: braintrust.F[braintrust.FunctionNewParamsPromptDataOptionsParamsUnion](braintrust.FunctionNewParamsPromptDataOptionsParamsObject{
					UseCache:         braintrust.F(true),
					Temperature:      braintrust.F(0.000000),
					TopP:             braintrust.F(0.000000),
					MaxTokens:        braintrust.F(0.000000),
					FrequencyPenalty: braintrust.F(0.000000),
					PresencePenalty:  braintrust.F(0.000000),
					ResponseFormat: braintrust.F(braintrust.FunctionNewParamsPromptDataOptionsParamsObjectResponseFormat{
						Type: braintrust.F(braintrust.FunctionNewParamsPromptDataOptionsParamsObjectResponseFormatTypeJsonObject),
					}),
					ToolChoice:   braintrust.F[braintrust.FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceUnion](braintrust.FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceString(braintrust.FunctionNewParamsPromptDataOptionsParamsObjectToolChoiceStringAuto)),
					FunctionCall: braintrust.F[braintrust.FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallUnion](braintrust.FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallString(braintrust.FunctionNewParamsPromptDataOptionsParamsObjectFunctionCallStringAuto)),
					N:            braintrust.F(0.000000),
					Stop:         braintrust.F([]string{"string", "string", "string"}),
				}),
				Position: braintrust.F("string"),
			}),
			Origin: braintrust.F(braintrust.FunctionNewParamsPromptDataOrigin{
				PromptID:      braintrust.F("string"),
				ProjectID:     braintrust.F("string"),
				PromptVersion: braintrust.F("string"),
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
			Description: braintrust.F("string"),
			FunctionData: braintrust.F[braintrust.FunctionUpdateParamsFunctionDataUnion](braintrust.FunctionUpdateParamsFunctionDataObject{
				Type: braintrust.F(braintrust.FunctionUpdateParamsFunctionDataObjectTypePrompt),
			}),
			Name: braintrust.F("string"),
			PromptData: braintrust.F(braintrust.FunctionUpdateParamsPromptData{
				Prompt: braintrust.F[braintrust.FunctionUpdateParamsPromptDataPromptUnion](braintrust.FunctionUpdateParamsPromptDataPromptObject{
					Type:    braintrust.F(braintrust.FunctionUpdateParamsPromptDataPromptObjectTypeCompletion),
					Content: braintrust.F("string"),
				}),
				Options: braintrust.F(braintrust.FunctionUpdateParamsPromptDataOptions{
					Model: braintrust.F("string"),
					Params: braintrust.F[braintrust.FunctionUpdateParamsPromptDataOptionsParamsUnion](braintrust.FunctionUpdateParamsPromptDataOptionsParamsObject{
						UseCache:         braintrust.F(true),
						Temperature:      braintrust.F(0.000000),
						TopP:             braintrust.F(0.000000),
						MaxTokens:        braintrust.F(0.000000),
						FrequencyPenalty: braintrust.F(0.000000),
						PresencePenalty:  braintrust.F(0.000000),
						ResponseFormat: braintrust.F(braintrust.FunctionUpdateParamsPromptDataOptionsParamsObjectResponseFormat{
							Type: braintrust.F(braintrust.FunctionUpdateParamsPromptDataOptionsParamsObjectResponseFormatTypeJsonObject),
						}),
						ToolChoice:   braintrust.F[braintrust.FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceUnion](braintrust.FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceString(braintrust.FunctionUpdateParamsPromptDataOptionsParamsObjectToolChoiceStringAuto)),
						FunctionCall: braintrust.F[braintrust.FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallUnion](braintrust.FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallString(braintrust.FunctionUpdateParamsPromptDataOptionsParamsObjectFunctionCallStringAuto)),
						N:            braintrust.F(0.000000),
						Stop:         braintrust.F([]string{"string", "string", "string"}),
					}),
					Position: braintrust.F("string"),
				}),
				Origin: braintrust.F(braintrust.FunctionUpdateParamsPromptDataOrigin{
					PromptID:      braintrust.F("string"),
					ProjectID:     braintrust.F("string"),
					PromptVersion: braintrust.F("string"),
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
		FunctionName:  braintrust.F("string"),
		IDs:           braintrust.F[braintrust.FunctionListParamsIDsUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
		Limit:         braintrust.F(int64(0)),
		OrgName:       braintrust.F("string"),
		ProjectName:   braintrust.F("string"),
		Slug:          braintrust.F("string"),
		StartingAfter: braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Version:       braintrust.F("string"),
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

func TestFunctionFeedback(t *testing.T) {
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
	err := client.Function.Feedback(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		braintrust.FunctionFeedbackParams{
			Feedback: braintrust.F([]braintrust.FunctionFeedbackParamsFeedback{{
				ID:      braintrust.F("string"),
				Comment: braintrust.F("string"),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": map[string]interface{}{},
				}),
				Source: braintrust.F(braintrust.FunctionFeedbackParamsFeedbackSourceApp),
			}, {
				ID:      braintrust.F("string"),
				Comment: braintrust.F("string"),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": map[string]interface{}{},
				}),
				Source: braintrust.F(braintrust.FunctionFeedbackParamsFeedbackSourceApp),
			}, {
				ID:      braintrust.F("string"),
				Comment: braintrust.F("string"),
				Metadata: braintrust.F(map[string]interface{}{
					"foo": map[string]interface{}{},
				}),
				Source: braintrust.F(braintrust.FunctionFeedbackParamsFeedbackSourceApp),
			}}),
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
	_, err := client.Function.Replace(context.TODO(), braintrust.FunctionReplaceParams{
		FunctionData: braintrust.F[braintrust.FunctionReplaceParamsFunctionDataUnion](braintrust.FunctionReplaceParamsFunctionDataObject{
			Type: braintrust.F(braintrust.FunctionReplaceParamsFunctionDataObjectTypePrompt),
		}),
		Name:        braintrust.F("string"),
		ProjectID:   braintrust.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Slug:        braintrust.F("string"),
		Description: braintrust.F("string"),
		PromptData: braintrust.F(braintrust.FunctionReplaceParamsPromptData{
			Prompt: braintrust.F[braintrust.FunctionReplaceParamsPromptDataPromptUnion](braintrust.FunctionReplaceParamsPromptDataPromptObject{
				Type:    braintrust.F(braintrust.FunctionReplaceParamsPromptDataPromptObjectTypeCompletion),
				Content: braintrust.F("string"),
			}),
			Options: braintrust.F(braintrust.FunctionReplaceParamsPromptDataOptions{
				Model: braintrust.F("string"),
				Params: braintrust.F[braintrust.FunctionReplaceParamsPromptDataOptionsParamsUnion](braintrust.FunctionReplaceParamsPromptDataOptionsParamsObject{
					UseCache:         braintrust.F(true),
					Temperature:      braintrust.F(0.000000),
					TopP:             braintrust.F(0.000000),
					MaxTokens:        braintrust.F(0.000000),
					FrequencyPenalty: braintrust.F(0.000000),
					PresencePenalty:  braintrust.F(0.000000),
					ResponseFormat: braintrust.F(braintrust.FunctionReplaceParamsPromptDataOptionsParamsObjectResponseFormat{
						Type: braintrust.F(braintrust.FunctionReplaceParamsPromptDataOptionsParamsObjectResponseFormatTypeJsonObject),
					}),
					ToolChoice:   braintrust.F[braintrust.FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceUnion](braintrust.FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceString(braintrust.FunctionReplaceParamsPromptDataOptionsParamsObjectToolChoiceStringAuto)),
					FunctionCall: braintrust.F[braintrust.FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallUnion](braintrust.FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallString(braintrust.FunctionReplaceParamsPromptDataOptionsParamsObjectFunctionCallStringAuto)),
					N:            braintrust.F(0.000000),
					Stop:         braintrust.F([]string{"string", "string", "string"}),
				}),
				Position: braintrust.F("string"),
			}),
			Origin: braintrust.F(braintrust.FunctionReplaceParamsPromptDataOrigin{
				PromptID:      braintrust.F("string"),
				ProjectID:     braintrust.F("string"),
				PromptVersion: braintrust.F("string"),
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
