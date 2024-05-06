// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"github.com/braintrustdata/braintrust-go/option"
)

// PromptSessionService contains methods and other services that help with
// interacting with the braintrust API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPromptSessionService] method
// instead.
type PromptSessionService struct {
	Options []option.RequestOption
}

// NewPromptSessionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPromptSessionService(opts ...option.RequestOption) (r *PromptSessionService) {
	r = &PromptSessionService{}
	r.Options = opts
	return
}
