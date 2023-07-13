package schema

import (
	"github.com/codebdy/entify-graphql-schema/resolve"
	"github.com/codebdy/entify/model/graph"
	"github.com/codebdy/entify/shared"
	"github.com/graphql-go/graphql"
)

func (m *MetaProcessor) mutationFields() graphql.Fields {
	mutationFields := graphql.Fields{}

	for _, entity := range m.Repo.Model.Graph.RootEnities() {
		if entity.Domain.Root {
			m.appendEntityMutationToFields(entity, mutationFields)
		}
	}
	for _, script := range m.Repo.Model.Meta.ScriptLogics {
		if script.OperateType == shared.MUTATION {
			m.appendScriptMethodsToFields(script, mutationFields)
		}
	}
	for _, logicFLow := range m.Repo.Model.Meta.GraphLogics {
		if logicFLow.OperateType == shared.MUTATION {
			m.appendLogicFlowMethodsToFields(logicFLow, mutationFields, getSubFlowMetas(m.Repo.Model.Meta.GraphLogics))
		}
	}
	return mutationFields
}

func (m *MetaProcessor) appendEntityMutationToFields(entity *graph.Entity, feilds graphql.Fields) {
	(feilds)[entity.DeleteName()] = &graphql.Field{
		Type:    m.modelParser.MutationResponse(entity.Name()),
		Args:    m.modelParser.DeleteArgs(entity),
		Resolve: resolve.DeleteResolveFn(entity.Name(), m.Repo),
	}
	(feilds)[entity.DeleteByIdName()] = &graphql.Field{
		Type:    m.modelParser.OutputType(entity.Name()),
		Args:    m.modelParser.DeleteByIdArgs(),
		Resolve: resolve.DeleteByIdResolveFn(entity.Name(), m.Repo),
	}
	(feilds)[entity.UpsertName()] = &graphql.Field{
		Type:    &graphql.List{OfType: m.modelParser.OutputType(entity.Name())},
		Args:    m.modelParser.UpsertArgs(entity),
		Resolve: resolve.PostResolveFn(entity.Name(), m.Repo),
	}
	(feilds)[entity.UpsertOneName()] = &graphql.Field{
		Type:    m.modelParser.OutputType(entity.Name()),
		Args:    m.modelParser.UpsertOneArgs(entity),
		Resolve: resolve.PostOneResolveFn(entity.Name(), m.Repo),
	}

	updateInput := m.modelParser.SetInput(entity.Name())
	if len(updateInput.Fields()) > 0 {
		(feilds)[entity.SetName()] = &graphql.Field{
			Type:    m.modelParser.MutationResponse(entity.Name()),
			Args:    m.modelParser.SetArgs(entity),
			Resolve: resolve.SetResolveFn(entity.Name(), m.Repo),
		}
	}
}
