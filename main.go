// Command fileserver serves stored files by name. It is a launch-gate fixture
// with one deliberate, known flaw: the name from the URL reaches the file
// system unchecked.
package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// storeDir is where files are kept. A variable so a test can point it at a
// temporary directory.
var storeDir = "store"

func serveFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	data, err := os.ReadFile(filepath.Join(storeDir, name))
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(data)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("GET /download", serveFile)
	http.HandleFunc("GET /healthz", health)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
