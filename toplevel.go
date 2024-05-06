// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package braintrust

import (
	"context"
	"net/http"

	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
)

// TopLevelService contains methods and other services that help with interacting
// with the braintrust API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewTopLevelService] method instead.
type TopLevelService struct {
	Options []option.RequestOption
}

// NewTopLevelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTopLevelService(opts ...option.RequestOption) (r *TopLevelService) {
	r = &TopLevelService{}
	r.Options = opts
	return
}

// Default endpoint. Simply replies with 'Hello, World!'. Authorization is not
// required
func (r *TopLevelService) HelloWorld(ctx context.Context, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
