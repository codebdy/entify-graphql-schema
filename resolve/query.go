package resolve

import (
	"fmt"

	"github.com/codebdy/entify-core"
	"github.com/codebdy/entify-core/model"
	"github.com/codebdy/entify-core/model/graph"
	"github.com/codebdy/entify-core/model/meta"
	"github.com/codebdy/entify-core/shared"
	"github.com/codebdy/entify-graphql-schema/service"
	"github.com/graphql-go/graphql"
)

func QueryOneEntityResolveFn(entityName string, r *entify.Repository) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer shared.PrintErrorStack()
		s := service.New(p.Context, r)
		instance := s.QueryOneEntity(entityName, p.Args)
		return instance, nil
	}
}

func QueryEntityResolveFn(entityName string, r *entify.Repository) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer shared.PrintErrorStack()
		s := service.New(p.Context, r)
		fields := parseListFields(p.Info)
		result := s.QueryEntity(entityName, p.Args, fields)
		return result, nil
	}
}

func QueryAssociationFn(asso *graph.Association, r *entify.Repository) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		var (
			source      = p.Source.(map[string]interface{})
			v           = p.Context.Value
			loaders     = v(shared.LOADERS).(*Loaders)
			handleError = func(err error) error {
				return fmt.Errorf(err.Error())
			}
		)
		defer shared.PrintErrorStack()

		if loaders == nil {
			panic("Data loaders is nil")
		}
		loader := loaders.GetLoader(p, asso, p.Args, r)
		thunk := loader.Load(p.Context, NewKey(source[shared.ID_NAME].(uint64)))
		return func() (interface{}, error) {
			data, err := thunk()
			if err != nil {
				return nil, handleError(err)
			}

			var retValue interface{}
			if data == nil {
				if asso.IsArray() {
					retValue = []map[string]interface{}{}
				} else {
					retValue = nil
				}
			} else {
				retValue = data
			}
			return retValue, nil
		}, nil
	}
}

func AttributeResolveFn(attr *graph.Attribute, model *model.Model) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer shared.PrintErrorStack()
		source := p.Source.(map[string]interface{})
		if attr.Type == meta.PASSWORD {
			return nil, nil
		}
		return source[attr.Name], nil
	}
}
