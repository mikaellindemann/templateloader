package templateloader

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/pkg/errors"
)

type onRequestLoader struct{}

func (onRequestLoader) Load(name string, h HandlerFunc, templateFiles ...string) (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New(name).ParseFiles(templateFiles...)
		if err != nil {
			http.Error(w, fmt.Sprintf("%+v", errors.WithStack(errors.Wrap(err, http.StatusText(http.StatusInternalServerError)))), http.StatusInternalServerError)
			return
		}
		h(t, w, r)
	}, nil
}

// NewOnRequestLoader creates a Loader that loads and parses the template on each request.
// Useful for development, as changes to template files are reflected on the next request.
// If the template cannot be parsed, the errors will show when the resulting handler is invoked.
func NewOnRequestLoader() Loader {
	return onRequestLoader{}
}
