package schema

import (
	"github.com/codebdy/entify-graphql-schema/resolve"
	"github.com/codebdy/entify/model/graph"
	"github.com/codebdy/entify/model/meta"
	"github.com/codebdy/entify/shared"
	"github.com/graphql-go/graphql"
)

func (m *MetaProcessor) QueryFields() []*graphql.Field {
	queryFields := graphql.Fields{}

	for _, entity := range m.Repo.Model.Graph.RootEnities() {
		m.appendEntityToQueryFields(entity, queryFields)
	}

	for _, scripts := range m.Repo.Model.Meta.ScriptLogics {
		if scripts.OperateType == shared.QUERY {
			m.appendMethodsToFields(scripts, queryFields)
		}
	}
	return convertFieldsArray(queryFields)
}

func (m *MetaProcessor) EntityQueryResponseType(entity *graph.Entity) graphql.Output {
	return m.modelParser.EntityListType(entity)
}
func (m *MetaProcessor) ClassQueryResponseType(cls *graph.Class) graphql.Output {
	return m.modelParser.ClassListType(cls)
}

func (m *MetaProcessor) appendEntityToQueryFields(entity *graph.Entity, fields graphql.Fields) {
	(fields)[entity.QueryName()] = &graphql.Field{
		Type:    m.EntityQueryResponseType(entity),
		Args:    m.modelParser.QueryArgs(entity.Name()),
		Resolve: resolve.QueryEntityResolveFn(entity.Name(), m.Repo),
	}
	(fields)[entity.QueryOneName()] = &graphql.Field{
		Type:    m.modelParser.OutputType(entity.Name()),
		Args:    m.modelParser.QueryArgs(entity.Name()),
		Resolve: resolve.QueryOneEntityResolveFn(entity.Name(), m.Repo),
	}
}

func (m *MetaProcessor) appendMethodsToFields(method *meta.MethodMeta, fields graphql.Fields) {
	fields[method.Name] = &graphql.Field{
		Type:        m.modelParser.MethodType(method),
		Args:        m.modelParser.MethodArgs(method.Args),
		Description: method.Description,
		Resolve:     resolve.ScriptMethodResolveFn(method.LogicScript, method.Args, m.Repo),
	}
}
