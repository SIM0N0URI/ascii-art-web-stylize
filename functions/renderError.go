package asciiart

import (
	"bytes"
	"html/template"
	"net/http"
)

type ErrorPage struct {
	Code    int
	Message string
}

// RenderError renders a custom error page with the given status code and message.
// It loads and executes the error template, then writes it to the response.
func RenderError(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}

	data := ErrorPage{
		Code:    code,
		Message: message,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, "Template write error", http.StatusInternalServerError)
		return
	}
}
