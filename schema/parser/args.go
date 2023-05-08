package parser

import (
	"github.com/codebdy/entify/model/graph"
	"github.com/codebdy/entify/model/observer/consts"
	"github.com/graphql-go/graphql"
)

func (p *ModelParser) DeleteArgs(entity *graph.Entity) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		consts.ARG_WHERE: &graphql.ArgumentConfig{
			Type: p.WhereExp(entity.Name()),
		},
	}
}

func (p *ModelParser) DeleteByIdArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		consts.ID: &graphql.ArgumentConfig{
			Type: graphql.ID,
		},
	}
}

func (p *ModelParser) UpsertArgs(entity *graph.Entity) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		consts.ARG_OBJECTS: &graphql.ArgumentConfig{
			Type: &graphql.NonNull{
				OfType: &graphql.List{
					OfType: &graphql.NonNull{
						OfType: p.SaveInput(entity.Name()),
					},
				},
			},
		},
	}
}

func (p *ModelParser) UpsertOneArgs(entity *graph.Entity) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		consts.ARG_OBJECT: &graphql.ArgumentConfig{
			Type: &graphql.NonNull{
				OfType: p.SaveInput(entity.Name()),
			},
		},
	}
}

func (p *ModelParser) SetArgs(entity *graph.Entity) graphql.FieldConfigArgument {
	updateInput := p.SetInput(entity.Name())
	return graphql.FieldConfigArgument{
		consts.ARG_SET: &graphql.ArgumentConfig{
			Type: &graphql.NonNull{
				OfType: updateInput,
			},
		},
		consts.ARG_WHERE: &graphql.ArgumentConfig{
			Type: p.WhereExp(entity.Name()),
		},
	}
}
