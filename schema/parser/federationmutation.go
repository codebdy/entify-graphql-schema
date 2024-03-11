package parser

import (
	"fmt"

	"github.com/codebdy/entify-core/model/graph"
	"github.com/codebdy/entify-core/shared"
	"github.com/graphql-go/graphql"
)

var mutationFieldSDL = "\t%s(%s) : %s \n"

func (p *ModelParser) MutationSDL() (string, string) {
	mutationFields := ""
	types := ""

	for _, entity := range p.model.Graph.RootEnities() {
		mutationFields = mutationFields + p.makeEntityMutationSDL(entity)
		types = types + objectToSDL(p.MutationResponse(entity.Name()), false)
	}

	for _, logic := range p.model.Meta.ScriptLogics {
		if logic.OperateType == shared.MUTATION {
			mutationFields = mutationFields + p.makeApiSDL(logic)
		}
	}

	for _, api := range p.model.Meta.APIs {
		if api.OperateType == shared.MUTATION {
			mutationFields = mutationFields + p.makeApiSDL(api)
		}
	}

	for _, input := range p.setInputMap {
		if len(input.Fields()) > 0 {
			types = types + inputToSDL(input)
		}

	}
	for _, input := range p.saveInputMap {
		types = types + inputToSDL(input)
	}
	for _, input := range p.hasManyInputMap {
		types = types + inputToSDL(input)
	}
	for _, input := range p.hasOneInputMap {
		types = types + inputToSDL(input)
	}

	return mutationFields, types
}

func (p *ModelParser) makeEntityMutationSDL(entity *graph.Entity) string {
	sdl := ""
	sdl = sdl + fmt.Sprintf(mutationFieldSDL,
		entity.DeleteName(),
		makeArgsSDL(p.DeleteArgs(entity)),
		p.MutationResponse(entity.Name()).Name(),
	)

	sdl = sdl + fmt.Sprintf(mutationFieldSDL,
		entity.DeleteByIdName(),
		makeArgsSDL(p.DeleteByIdArgs()),
		p.OutputType(entity.Name()).String(),
	)

	updateInput := p.SetInput(entity.Name())
	if len(updateInput.Fields()) > 0 {
		sdl = sdl + fmt.Sprintf(mutationFieldSDL,
			entity.SetName(),
			makeArgsSDL(p.SetArgs(entity)),
			p.MutationResponse(entity.Name()).String(),
		)
	}

	sdl = sdl + fmt.Sprintf(mutationFieldSDL,
		entity.UpsertName(),
		makeArgsSDL(p.UpsertArgs(entity)),
		(&graphql.List{OfType: p.OutputType(entity.Name())}).String(),
	)

	sdl = sdl + fmt.Sprintf(mutationFieldSDL,
		entity.UpsertOneName(),
		makeArgsSDL(p.UpsertOneArgs(entity)),
		p.OutputType(entity.Name()).String(),
	)

	return sdl
}
