package schema

import (
	"github.com/codebdy/entify-core/model/graph"
	"github.com/codebdy/entify-core/model/meta"
	"github.com/codebdy/entify-core/shared"
	"github.com/codebdy/entify-graphql-schema/resolve"
	"github.com/codebdy/minions-go/dsl"
	"github.com/graphql-go/graphql"
)

func (m *MetaProcessor) QueryFields() graphql.Fields {
	// EntityType = graphql.NewUnion(
	// 	graphql.UnionConfig{
	// 		Name:        consts.ENTITY_TYPE,
	// 		Types:       m.modelParser.EntityObjects(),
	// 		ResolveType: m.modelParser.ResolveTypeFn,
	// 	},
	// )
	queryFields := graphql.Fields{
		// consts.SERVICE: m.serviceField(),
		// consts.ENTITIES: &graphql.Field{
		// 	Type: &graphql.NonNull{
		// 		OfType: &graphql.List{
		// 			OfType: EntityType,
		// 		},
		// 	},
		// 	Args: graphql.FieldConfigArgument{
		// 		consts.REPRESENTATIONS: &graphql.ArgumentConfig{
		// 			Type: &graphql.NonNull{
		// 				OfType: &graphql.List{
		// 					OfType: &graphql.NonNull{
		// 						OfType: scalars.AnyType,
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// },
	}

	for _, entity := range m.Repo.Model.Graph.RootEnities() {
		m.appendEntityToQueryFields(entity, queryFields)
	}

	for _, script := range m.Repo.Model.Meta.ScriptLogics {
		if script.OperateType == shared.QUERY {
			m.appendScriptMethodsToFields(script, queryFields)
		}
	}

	for _, logicFLow := range m.Repo.Model.Meta.GraphLogics {
		if logicFLow.OperateType == shared.QUERY {
			m.appendLogicFlowMethodsToFields(logicFLow, queryFields, getSubFlowMetas(m.Repo.Model.Meta.GraphLogics))
		}
	}
	return queryFields
}

func (m *MetaProcessor) EntityQueryResponseType(entity *graph.Entity) graphql.Output {
	return m.modelParser.EntityListType(entity)
}

// func (m *MetaProcessor) ClassQueryResponseType(cls *graph.Class) graphql.Output {
// 	return m.modelParser.ClassListType(cls)
// }

func (m *MetaProcessor) appendEntityToQueryFields(entity *graph.Entity, fields graphql.Fields) {
	(fields)[entity.QueryListName()] = &graphql.Field{
		Type:    m.EntityQueryResponseType(entity),
		Args:    m.modelParser.QueryArgs(entity.Name()),
		Resolve: resolve.QueryEntityListResolveFn(entity.Name(), m.Repo),
	}
	(fields)[entity.QueryOneName()] = &graphql.Field{
		Type:    m.modelParser.OutputType(entity.Name()),
		Args:    m.modelParser.QueryArgs(entity.Name()),
		Resolve: resolve.QueryOneEntityResolveFn(entity.Name(), m.Repo),
	}
}

func (m *MetaProcessor) appendScriptMethodsToFields(method *meta.MethodMeta, fields graphql.Fields) {
	fields[method.Name] = &graphql.Field{
		Type:        m.modelParser.MethodType(method),
		Args:        m.modelParser.MethodArgs(method.Args),
		Description: method.Description,
		Resolve:     resolve.ScriptMethodResolveFn(method, m.Repo),
	}
}

func (m *MetaProcessor) appendLogicFlowMethodsToFields(method *meta.MethodMeta, fields graphql.Fields, subFlows *[]dsl.SubLogicFlowMeta) {
	fields[method.Name] = &graphql.Field{
		Type:        m.modelParser.MethodType(method),
		Args:        m.modelParser.MethodArgs(method.Args),
		Description: method.Description,
		Resolve:     resolve.LogicFlowMethodResolveFn(method, m.Repo, subFlows),
	}
}

func getSubFlowMetas(methodMetas []*meta.MethodMeta) *[]dsl.SubLogicFlowMeta {
	subLogicFlows := []dsl.SubLogicFlowMeta{}
	for _, method := range methodMetas {
		if method.OperateType == shared.SUBMETHOD {
			subLogicFlows = append(subLogicFlows, dsl.SubLogicFlowMeta{Id: method.Uuid, LogicFlowMeta: method.LogicMetas})
		}
	}

	return &subLogicFlows
}
