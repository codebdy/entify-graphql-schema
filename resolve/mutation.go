package resolve

import (
	"log"

	"github.com/codebdy/entify-core"
	"github.com/codebdy/entify-core/model/data"
	"github.com/codebdy/entify-core/model/observer"
	"github.com/codebdy/entify-core/shared"
	"github.com/codebdy/entify-graphql-schema/service"
	"github.com/graphql-go/graphql"
)

func PostResolveFn(entityName string, r *entify.Repository) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer shared.PrintErrorStack()
		objects := p.Args[shared.ARG_OBJECTS].([]map[string]interface{})

		s := service.New(p.Context, r)
		returing, err := s.Save(entityName, objects)

		if err != nil {
			return nil, err
		}
		observer.EmitObjectMultiPosted(returing, entityName, p.Context)
		return returing, nil
	}
}

// 未实现
func SetResolveFn(entityName string, r *entify.Repository) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer shared.PrintErrorStack()
		// s := service.New(p.Context, r)
		// objs := s.QueryEntityList(entityName, p.Args, []string{}).Nodes
		// convertedObjs := objs
		//instances := []*data.Instance{}

		// for i := range convertedObjs {
		// 	obj := convertedObjs[i]
		// 	object := map[string]interface{}{}

		// 	object[consts.ID] = obj[consts.ID]

		// 	for key := range set {
		// 		object[key] = set[key]
		// 	}
		// }
		//returing, err := s.Save(entityName, convertedObjs)

		// if err != nil {
		// 	return nil, err
		// }

		//logs.WriteModelLog(model.Graph, &entity.Class, p.Context, logs.SET, logs.SUCCESS, "", p.Context.Value("gql"))

		// return map[string]interface{}{
		// 	shared.RESPONSE_RETURNING:    returing,
		// 	shared.RESPONSE_AFFECTEDROWS: len(instances),
		// }, nil
		return nil, nil
	}
}

func PostOneResolveFn(entityName string, r *entify.Repository) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer shared.PrintErrorStack()
		object := p.Args[shared.ARG_OBJECT].(map[string]interface{})
		data.ConvertObjectId(object)

		s := service.New(p.Context, r)
		result, err := s.SaveOne(entityName, object)
		observer.EmitObjectPosted(result.(map[string]interface{}), entityName, p.Context)
		return result, err
	}
}

func DeleteByIdResolveFn(entityName string, r *entify.Repository) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer shared.PrintErrorStack()
		argId := p.Args[shared.ID_NAME]

		s := service.New(p.Context, r)
		result, err := s.DeleteInstance(entityName, data.ConvertId(argId))
		observer.EmitObjectDeleted(result.(map[string]interface{}), entityName, p.Context)
		return result, err
	}
}

func DeleteResolveFn(entityName string, r *entify.Repository) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		defer shared.PrintErrorStack()
		s := service.New(p.Context, r)
		objs := s.QueryEntityList(entityName, p.Args, []string{shared.ID_NAME}).Nodes

		if objs == nil || len(objs) == 0 {
			return map[string]interface{}{
				shared.RESPONSE_RETURNING:    []interface{}{},
				shared.RESPONSE_AFFECTEDROWS: 0,
			}, nil
		}

		convertedObjs := objs

		ids := []shared.ID{}
		for i := range convertedObjs {
			ids = append(ids, data.ConvertId(convertedObjs[i][shared.ID_NAME]))
		}

		_, err := s.DeleteInstances(entityName, ids)
		if err != nil {
			log.Panic(err.Error())
		}
		observer.EmitObjectMultiDeleted(objs, entityName, p.Context)
		return map[string]interface{}{
			shared.RESPONSE_RETURNING:    objs,
			shared.RESPONSE_AFFECTEDROWS: len(ids),
		}, nil
	}
}
