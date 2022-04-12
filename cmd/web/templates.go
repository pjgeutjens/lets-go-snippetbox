package main

import "leapconsulting.be/snippetbox/pkg/models"

// templateData is a type that acts as a holding structure
// for dynamic data we want to pass to the HTML templates
// starts out only containing a single field
type templateData struct {
	Snippet *models.Snippet
}
