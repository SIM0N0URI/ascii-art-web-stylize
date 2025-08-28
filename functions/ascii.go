package asciiart

import (
	"bytes"
	"net/http"
	"text/template"
)

// Ascii handles POST requests to generate ASCII art from input text and banner style.
// It validates input, renders the output using a template, and returns the result.
func Ascii(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		RenderError(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		RenderError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		RenderError(w, "Internal Server Error: failed to load template", http.StatusInternalServerError)
		return
	}

	Text := r.FormValue("text")
	Banner := r.FormValue("banner")

	allowedBanners := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}

	if !allowedBanners[Banner] {
		RenderError(w, "Invalid banner selected", http.StatusBadRequest)
		return
	}

	textOutput, err := GenerateASCII(Text, Banner)
	if err != nil {
		if err.Error() == "failed to read banner file" {
			RenderError(w, "Failed to read banner file", http.StatusInternalServerError)
			return
		}
		RenderError(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, textOutput); err != nil {
		RenderError(w, "Internal Server Error: failed to render template", http.StatusInternalServerError)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		RenderError(w, "Internal Server Error: failed to write response", http.StatusInternalServerError)
		return
	}
}
