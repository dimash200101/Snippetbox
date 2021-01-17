package main

import (
	"html/template"
	"net/url" // New import
	"path/filepath"
	"time"
	"alexedwards.net/snippetbox/pkg/models"
)
// Add FormData and FormErrors fields to the templateData struct.
type templateData struct {
	CurrentYear int
	FormData url.Values
	FormErrors map[string]string
	Snippet *models.Snippet
	Snippets []*models.Snippet
}
