package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"zhakhangers.net/snippetbox/internal/models"
	"zhakhangers.net/snippetbox/ui"
)

type templateData struct {
	CurrentYear int
	Snippet *models.Snippet
	Snippets []*models.Snippet
	Form any
	Flash string
	IsAuthenticated bool
	CSRFToken string
}

func humanDate(t time.Time) string {

	if t.IsZero() {
		return ""
	}
	
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap {
	"humanDate": humanDate,
}



func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		// Create a slice containing the filepath patterns for the templates we
		// want to parse.
		patterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}
		// Use ParseFS() instead of ParseFiles() to parse the template files
		// from the ui.Files embedded filesystem.
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}
		// Add the template set to the map as normal...
		cache[name] = ts
	}
	return cache, nil
	}