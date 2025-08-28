package asciiart

import (
	"bytes"
	"html/template"
	"net/http"
)

// Home handles GET requests to the root path and renders the main HTML template.
// It validates the request method and path before serving the homepage.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderError(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		RenderError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		RenderError(w, "Internal Server Error: failed to load template", http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, ""); err != nil {
		RenderError(w, "Internal Server Error: failed to render template", http.StatusInternalServerError)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		RenderError(w, "Internal Server Error: failed to write response", http.StatusInternalServerError)
		return
	}
}
