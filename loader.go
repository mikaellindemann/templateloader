// Package templateloader provides different strategies for loading html templates in a web application.
package templateloader

import (
	"html/template"
	"net/http"
)

// HandlerFunc is an extension to http.HandlerFunc that takes an html template as first argument.
// Used in conjunction with the Loader interface.
type HandlerFunc func(*template.Template, http.ResponseWriter, *http.Request)

// The Loader describes a common interface for loading html templates that can be used with a HandlerFunc.
type Loader interface {

	// Load returns an http.HandlerFunc that passes a loaded html template to the supplied HandlerFunc.
	// name should correspond to the main template to load for this handler.
	Load(name string, h HandlerFunc, templateFiles ...string) (http.HandlerFunc, error)
}
