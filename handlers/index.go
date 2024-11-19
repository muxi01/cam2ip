package handlers

import (
	"fmt"
	"net/http"
)

// Index handler.
type Index struct {
	Baseurl string
}

// NewIndex returns new Index handler.
func NewIndex(baseurl string) *Index {
	index := &Index{}
	index.Baseurl = baseurl
	return index
}

// ServeHTTP handles requests on incoming connections.
func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.Method != "HEAD" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	contents := fmt.Sprintf(`<html>
                <head><title>cam2ip</title></head>
                <body>
                <h1>cam2ip</h1>
                <p><a href='%shtml'>html</a></p>
                <p><a href='%sjpeg'>jpeg</a></p>
                <p><a href='%smjpeg'>mjpeg</a></p>
                </body>
                </html>`, i.Baseurl, i.Baseurl, i.Baseurl)

	w.Write([]byte(contents))
}
