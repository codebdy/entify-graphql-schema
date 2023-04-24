package schema

import (
	"github.com/codebdy/entify-graphql-schema/resolve"
	"github.com/codebdy/entify/model/graph"
	"github.com/codebdy/entify/model/observer/consts"
	"github.com/codebdy/entify/shared"
	"github.com/graphql-go/graphql"
)

func (m *MetaProcessor) mutationFields() []*graphql.Field {
	mutationFields := graphql.Fields{}

	for _, entity := range m.Repo.Model.Graph.RootEnities() {
		if entity.Domain.Root {
			m.appendEntityMutationToFields(entity, mutationFields)
		}
	}
	for _, scripts := range m.Repo.Model.Meta.ScriptLogics {
		if scripts.OperateType == shared.MUTATION {
			m.appendMethodsToFields(scripts, mutationFields)
		}
	}
	return convertFieldsArray(mutationFields)
}

func (m *MetaProcessor) deleteArgs(entity *graph.Entity) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		consts.ARG_WHERE: &graphql.ArgumentConfig{
			Type: m.modelParser.WhereExp(entity.Name()),
		},
	}
}

func deleteByIdArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		consts.ID: &graphql.ArgumentConfig{
			Type: graphql.ID,
		},
	}
}

func (m *MetaProcessor) upsertArgs(entity *graph.Entity) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		consts.ARG_OBJECTS: &graphql.ArgumentConfig{
			Type: &graphql.NonNull{
				OfType: &graphql.List{
					OfType: &graphql.NonNull{
						OfType: m.modelParser.SaveInput(entity.Name()),
					},
				},
			},
		},
	}
}

func (m *MetaProcessor) upsertOneArgs(entity *graph.Entity) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		consts.ARG_OBJECT: &graphql.ArgumentConfig{
			Type: &graphql.NonNull{
				OfType: m.modelParser.SaveInput(entity.Name()),
			},
		},
	}
}

func (m *MetaProcessor) setArgs(entity *graph.Entity) graphql.FieldConfigArgument {
	updateInput := m.modelParser.SetInput(entity.Name())
	return graphql.FieldConfigArgument{
		consts.ARG_SET: &graphql.ArgumentConfig{
			Type: &graphql.NonNull{
				OfType: updateInput,
			},
		},
		consts.ARG_WHERE: &graphql.ArgumentConfig{
			Type: m.modelParser.WhereExp(entity.Name()),
		},
	}
}

func (m *MetaProcessor) appendEntityMutationToFields(entity *graph.Entity, feilds graphql.Fields) {
	(feilds)[entity.DeleteName()] = &graphql.Field{
		Type:    m.modelParser.MutationResponse(entity.Name()),
		Args:    m.deleteArgs(entity),
		Resolve: resolve.DeleteResolveFn(entity.Name(), m.Repo),
	}
	(feilds)[entity.DeleteByIdName()] = &graphql.Field{
		Type:    m.modelParser.OutputType(entity.Name()),
		Args:    deleteByIdArgs(),
		Resolve: resolve.DeleteByIdResolveFn(entity.Name(), m.Repo),
	}
	(feilds)[entity.UpsertName()] = &graphql.Field{
		Type:    &graphql.List{OfType: m.modelParser.OutputType(entity.Name())},
		Args:    m.upsertArgs(entity),
		Resolve: resolve.PostResolveFn(entity.Name(), m.Repo),
	}
	(feilds)[entity.UpsertOneName()] = &graphql.Field{
		Type:    m.modelParser.OutputType(entity.Name()),
		Args:    m.upsertOneArgs(entity),
		Resolve: resolve.PostOneResolveFn(entity.Name(), m.Repo),
	}

	updateInput := m.modelParser.SetInput(entity.Name())
	if len(updateInput.Fields()) > 0 {
		(feilds)[entity.SetName()] = &graphql.Field{
			Type:    m.modelParser.MutationResponse(entity.Name()),
			Args:    m.setArgs(entity),
			Resolve: resolve.SetResolveFn(entity.Name(), m.Repo),
		}
	}
}
