package npncontroller

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kyleu/libnpn/npncontroller/routes"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npntemplate/gen/npntemplate"
	"github.com/kyleu/libnpn/npnweb"
)

// Common utility routes, for viewing the installed routes and modules, and providing sitemap.xml, OPTIONS, and 404 handlers
func RoutesUtil(app npnweb.AppInfo, r *mux.Router) {
	RoutesRouteList(app, r)
	RoutesModuleList(app, r)
	RoutesSitemap(app, r)
	RoutesOptions(app, r)
	RoutesNotFound(app, r)
}

func RoutesRouteList(app npnweb.AppInfo, r *mux.Router) {
	rts := r.Path(routes.Path(npncore.KeyRoutes)).Subrouter()
	rts.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(RouteList))).Name(routes.Name(npncore.KeyRoutes))
}

func RoutesModuleList(app npnweb.AppInfo, r *mux.Router) {
	modules := r.Path(routes.Path(npncore.KeyModules)).Subrouter()
	modules.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(ModuleList))).Name(routes.Name(npncore.KeyModules))
}

func RoutesSitemap(app npnweb.AppInfo, r *mux.Router) {
	sitemap := r.Path(routes.Path("sitemap.xml")).Subrouter()
	sitemap.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(SitemapXML))).Name(routes.Name("sitemap"))
}

func RoutesOptions(app npnweb.AppInfo, r *mux.Router) {
	r.PathPrefix("").Methods(http.MethodOptions).Handler(routes.AddContext(r, app, http.HandlerFunc(Options)))
}

func RoutesNotFound(app npnweb.AppInfo, r *mux.Router) {
	r.PathPrefix("").Handler(routes.AddContext(r, app, http.HandlerFunc(NotFound)))
}

// List the configured HTTP routes
func RouteList(w http.ResponseWriter, r *http.Request) {
	Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = npncore.Title(npncore.KeyRoutes)
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(npncore.KeyRoutes)}
		return T(npntemplate.RoutesList(ctx, w))
	})
}

// Lists the Go modules used
func ModuleList(w http.ResponseWriter, r *http.Request) {
	Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = npncore.Title(npncore.KeyModules)
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(npncore.KeyModules)}
		return T(npntemplate.ModulesList(ctx, w))
	})
}

// Handles the standard /sitemap.xml request for the configured routes
func SitemapXML(w http.ResponseWriter, r *http.Request) {
	Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ret := make([]string, 0)
		ret = append(ret, `<?xml version="1.0" encoding="UTF-8"?>`)
		ret = append(ret, `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)
		for _, rt := range npnweb.ExtractRoutes(ctx.Routes) {
			if routeMatches(rt) {
				url := rt.Path
				ret = append(ret, `  <url>`)
				ret = append(ret, `     <loc>`+url+`</loc>`)
				ret = append(ret, `     <changefreq>always</changefreq>`)
				ret = append(ret, `  </url>`)
			}
		}
		ret = append(ret, `</urlset>`)
		_, _ = w.Write([]byte(strings.Join(ret, "\n")))
		return "", nil
	})
}

func routeMatches(rt *npnweb.RouteDescription) bool {
	pathCheck := func(s ...string) bool {
		for _, x := range s {
			if strings.Contains(rt.Path, x) {
				return false
			}
		}
		return true
	}
	if !pathCheck("admin", "assets", "sitemap", "robots", "{") {
		return false
	}
	if rt.Path == "/ws" {
		return false
	}
	if !strings.Contains(rt.Methods, http.MethodGet) {
		return false
	}
	return true
}
