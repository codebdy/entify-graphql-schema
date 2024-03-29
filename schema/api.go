package schema

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/codebdy/entify-core/model/meta"
	"github.com/codebdy/entify-core/shared"
	"github.com/graphql-go/graphql"
)

func (mp *MetaProcessor) QueryApiFields(resolver interface{}) graphql.Fields {
	queryFields := graphql.Fields{}
	for _, api := range mp.Repo.Model.Meta.APIs {
		if api.OperateType == shared.QUERY {
			mp.appendApiToFields(api, queryFields, resolver)
		}
	}
	return queryFields
}

func (mp *MetaProcessor) MutationApiFields(resolver interface{}) graphql.Fields {
	mutationFields := graphql.Fields{}
	for _, api := range mp.Repo.Model.Meta.APIs {
		if api.OperateType == shared.MUTATION {
			mp.appendApiToFields(api, mutationFields, resolver)
		}
	}
	return mutationFields
}

func (mp *MetaProcessor) appendApiToFields(method *meta.MethodMeta, fields graphql.Fields, resolver interface{}) {
	op := shared.FirstUpper(method.Name)

	fields[method.Name] = &graphql.Field{
		Type:        mp.modelParser.MethodType(method),
		Args:        mp.modelParser.MethodArgs(method.Args),
		Description: method.Description,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			rv := reflect.ValueOf(resolver)
			m := rv.MethodByName(op)
			if m.IsValid() {
				mt := m.Type()
				if mt.NumIn() != len(method.Args) && mt.NumIn() != (len(method.Args)+1) {
					return nil, fmt.Errorf("method %q of %v must accept less %d or %d arguments, got %d", op, rv.Type(), len(method.Args), len(method.Args)+1, mt.NumIn())
				}
				// if mt.NumIn() > 2 {
				// 	return nil, fmt.Errorf("method %q of %v must accept less 2 arguments, got %d", op, rv.Type(), mt.NumIn())
				// }
				if mt.NumOut() != 1 {
					return nil, fmt.Errorf("method %q of %v must have 1 return value, got %d", op, rv.Type(), mt.NumOut())
				}
				//ot := mt.Out(0)
				// if ot.Kind() != reflect.Pointer && ot.Kind() != reflect.Interface {
				// 	return nil, fmt.Errorf("method %q of %v must return an interface or a pointer, got %+v", op, rv.Type(), ot)
				// }
				inputs := make([]reflect.Value, mt.NumIn())
				for i := range method.Args {
					arg := method.Args[i]
					inputs[i] = reflect.ValueOf(p.Args[arg.Name])
				}
				if mt.NumIn() > len(method.Args) {
					inputs[len(method.Args)] = reflect.ValueOf(p.Context)
				}

				out := m.Call(inputs)
				res := out[0]
				// if res.IsNil() {
				// 	return nil, fmt.Errorf("method %q of %v must return a non-nil result, got %v", op, rv.Type(), res)
				// }
				// switch res.Kind() {
				// case reflect.Pointer:
				// 	resolvers[op] = res.Elem().Addr().Interface()
				// case reflect.Interface:
				// 	resolvers[op] = res.Elem().Interface()
				// default:
				// 	panic("ureachable")
				// }
				return res, nil
			}

			return nil, errors.New("Can not find resolver method:" + op)
		},
	}
}
