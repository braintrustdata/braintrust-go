// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"reflect"

	"github.com/braintrustdata/braintrust-go/internal/apijson"
	"github.com/braintrustdata/braintrust-go/internal/requestconfig"
	"github.com/braintrustdata/braintrust-go/option"
)

type ListObjects[T any] struct {
	Objects []T             `json:"objects"`
	JSON    listObjectsJSON `json:"-"`
	cfg     *requestconfig.RequestConfig
	res     *http.Response
}

// listObjectsJSON contains the JSON metadata for the struct [ListObjects[T]]
type listObjectsJSON struct {
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListObjects[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listObjectsJSON) RawJSON() string {
	return r.raw
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *ListObjects[T]) GetNextPage() (res *ListObjects[T], err error) {
	items := r.Objects
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	cfg.Apply(option.WithQuery("starting_after", field.Interface().(string)))
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ListObjects[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	r.cfg = cfg
	r.res = res
}

type ListObjectsAutoPager[T any] struct {
	page *ListObjects[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewListObjectsAutoPager[T any](page *ListObjects[T], err error) *ListObjectsAutoPager[T] {
	return &ListObjectsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ListObjectsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Objects) == 0 {
		return false
	}
	if r.idx >= len(r.page.Objects) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Objects) == 0 {
			return false
		}
	}
	r.cur = r.page.Objects[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ListObjectsAutoPager[T]) Current() T {
	return r.cur
}

func (r *ListObjectsAutoPager[T]) Err() error {
	return r.err
}

func (r *ListObjectsAutoPager[T]) Index() int {
	return r.run
}
