package templateloader

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
)

type preloader struct{}

func (preloader) Load(name string, h HandlerFunc, templateFiles ...string) (http.HandlerFunc, error) {
	t, err := template.New(name).ParseFiles(templateFiles...)
	return func(w http.ResponseWriter, r *http.Request) {
		h(t, w, r)
	}, errors.Wrap(err, "failed preloading template")
}

// Creates a Loader that loads and parses the template when Load is called.
// If the template cannot be parsed, an error is returned as the result of Load.
func NewPreloader() Loader {
	return preloader{}
}
