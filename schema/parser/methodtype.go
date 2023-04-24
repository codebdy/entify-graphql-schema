package parser

import (
	"log"

	"github.com/codebdy/entify/model/meta"
	"github.com/graphql-go/graphql"
)

func (p *ModelParser) MethodType(method *meta.MethodMeta) graphql.Output {
	switch method.Type {
	case meta.ENTITY:
		entity := p.model.Graph.GetEntityByUuid(method.TypeUuid)
		if entity == nil {
			log.Panic("Can not find entity by uuid:" + method.TypeUuid)
		}
		return p.OutputType(entity.Name())
	case meta.ENTITY_ARRAY:
		entity := p.model.Graph.GetEntityByUuid(method.TypeUuid)
		if entity == nil {
			log.Panic("Can not find entity by uuid:" + method.TypeUuid)
		}
		return p.EntityListType(entity)
	}
	return PropertyType(method.Type)
}
