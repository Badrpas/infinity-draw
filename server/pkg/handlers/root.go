package handlers

import (
	"net/http"
	"regexp"
)

var fs http.Handler

func init() {
	fs = http.FileServer(http.Dir("./static/"))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if ok, err := regexp.MatchString("\\.js$", r.URL.Path); ok && err == nil {
		w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	}
	fs.ServeHTTP(w, r)
}
