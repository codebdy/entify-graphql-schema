package schema

import (
	"github.com/codebdy/entify-graphql-schema/resolve"
	"github.com/codebdy/entify/model/graph"
	"github.com/graphql-go/graphql"
)

func (a *MetaProcessor) QueryFields() []*graphql.Field {
	queryFields := graphql.Fields{}

	for _, entity := range a.Repo.Model.Graph.RootEnities() {
		a.appendEntityToQueryFields(entity, queryFields)
	}
	return convertFieldsArray(queryFields)
}

func (a *MetaProcessor) EntityQueryResponseType(entity *graph.Entity) graphql.Output {
	return a.modelParser.EntityListType(entity)
}
func (a *MetaProcessor) ClassQueryResponseType(cls *graph.Class) graphql.Output {
	return a.modelParser.ClassListType(cls)
}

func (a *MetaProcessor) appendEntityToQueryFields(entity *graph.Entity, fields graphql.Fields) {
	(fields)[entity.QueryName()] = &graphql.Field{
		Type:    a.EntityQueryResponseType(entity),
		Args:    a.modelParser.QueryArgs(entity.Name()),
		Resolve: resolve.QueryEntityResolveFn(entity.Name(), a.Repo),
	}
	(fields)[entity.QueryOneName()] = &graphql.Field{
		Type:    a.modelParser.OutputType(entity.Name()),
		Args:    a.modelParser.QueryArgs(entity.Name()),
		Resolve: resolve.QueryOneEntityResolveFn(entity.Name(), a.Repo),
	}

}
