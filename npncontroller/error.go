package npncontroller

import (
	"fmt"
	"net/http"

	"github.com/kyleu/libnpn/npntemplate/gen/npntemplate"
	"github.com/kyleu/libnpn/npnweb"
)

// 404 handler
func NotFound(w http.ResponseWriter, r *http.Request) {
	WriteCORS(w)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	ctx := npnweb.ExtractContext(w, r, false)
	ctx.Title = "Not Found"
	ctx.Breadcrumbs = npnweb.BreadcrumbsSimple(r.URL.Path, "not found")
	ctx.Logger.Info(fmt.Sprintf("%v %v returned [%d]", r.Method, r.URL.Path, http.StatusNotFound))
	_, _ = npntemplate.NotFound(r, ctx, w)
}

// Handles HTTP requests of type OPTIONS, adds CORS headers
func Options(w http.ResponseWriter, r *http.Request) {
	WriteCORS(w)
	w.WriteHeader(http.StatusOK)
}
