package npngraphql

import (
	"net/http"

	"github.com/kyleu/libnpn/npntemplate/gen/npntemplate"

	"github.com/kyleu/libnpn/npncontroller"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npnweb"
)

var (
	gqlQueryName    = "Query"
	gqlMutationName = "Mutation"
)

const (
	graphiqlName = "GraphiQL"
	voyagerName  = "GraphQL Voyager"
)

func GraphiQL(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(npncore.KeyGraphiQL)}
		ctx.Title = graphiqlName
		return npncontroller.T(npntemplate.GraphiQL(ctx, w))
	})
}

func AdminGraphiQL(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Breadcrumbs = npncontroller.AdminBC(ctx, npncore.KeyGraphiQL, npncore.KeyGraphQL)
		ctx.Title = graphiqlName
		return npncontroller.T(npntemplate.GraphiQL(ctx, w))
	})
}

func GraphQLVoyagerQuery(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		bc := npnweb.BreadcrumbsSimple(npncore.KeyGraphiQL, npncore.KeyGraphQL)
		bc = append(bc, npnweb.BreadcrumbSelf("query"))
		ctx.Breadcrumbs = bc
		ctx.Title = voyagerName
		return npncontroller.T(npntemplate.GraphQLVoyager(gqlQueryName, ctx, w))
	})
}

func AdminGraphQLVoyagerQuery(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		bc := npncontroller.AdminBC(ctx, npncore.KeyGraphiQL, npncore.KeyGraphQL)
		bc = append(bc, npnweb.BreadcrumbSelf("query"))
		ctx.Breadcrumbs = bc
		ctx.Title = voyagerName
		return npncontroller.T(npntemplate.GraphQLVoyager(gqlQueryName, ctx, w))
	})
}

func GraphQLVoyagerMutation(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		bc := npnweb.BreadcrumbsSimple(npncore.KeyGraphiQL, npncore.KeyGraphQL)
		bc = append(bc, npnweb.BreadcrumbSelf("mutation"))
		ctx.Breadcrumbs = bc
		ctx.Title = voyagerName
		return npncontroller.T(npntemplate.GraphQLVoyager(gqlMutationName, ctx, w))
	})
}

func AdminGraphQLVoyagerMutation(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		bc := npncontroller.AdminBC(ctx, npncore.KeyGraphiQL, npncore.KeyGraphQL)
		bc = append(bc, npnweb.BreadcrumbSelf("mutation"))
		ctx.Breadcrumbs = bc
		ctx.Title = voyagerName
		return npncontroller.T(npntemplate.GraphQLVoyager(gqlMutationName, ctx, w))
	})
}
