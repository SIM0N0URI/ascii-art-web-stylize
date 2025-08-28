package asciiart

import (
	"net/http"
	"os"
)

// Style serves static style files based on the request URL.
// It checks file existence and access permissions before serving the file.
func Style(w http.ResponseWriter, r *http.Request) {
	fileinfo, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		RenderError(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if fileinfo.IsDir() {
		RenderError(w, "Access denied", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
