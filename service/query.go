package service

import (
	"log"

	"github.com/codebdy/entify-core/model/graph"
	"github.com/codebdy/entify-core/orm"
	"github.com/codebdy/entify-core/shared"
)

func mergeWhereArgs(whereArgs, authArgs graph.QueryArg) graph.QueryArg {
	if len(whereArgs) == 0 {
		return authArgs
	}

	if len(authArgs) == 0 {
		return whereArgs
	}

	return graph.QueryArg{
		shared.ARG_AND: []graph.QueryArg{
			whereArgs,
			authArgs,
		},
	}
}

func (s *Service) QueryEntity(entityName string, args graph.QueryArg, fieldNames []string) orm.QueryResponse {
	canRead, authArgs := s.canReadEntity(entityName)
	if !canRead {
		log.Panic("No access")
	}
	session, err := s.repository.OpenSession()
	if err != nil {
		panic(err.Error())
	}

	return session.Query(entityName, mergeWhereArgs(args, authArgs), fieldNames)
}

func (s *Service) QueryOneEntity(entityName string, args graph.QueryArg) interface{} {
	canRead, authArgs := s.canReadEntity(entityName)
	if !canRead {
		log.Panic("No access")
	}
	session, err := s.repository.OpenSession()
	if err != nil {
		log.Panic(err.Error())
	}
	return session.QueryOne(entityName, mergeWhereArgs(args, authArgs))
}

func (s *Service) QueryById(entityName string, id uint64) interface{} {
	canRead, authArgs := s.canReadEntity(entityName)
	if !canRead {
		log.Panic("No access")
	}
	return s.QueryOneEntity(entityName, mergeWhereArgs(graph.QueryArg{
		shared.ARG_WHERE: graph.QueryArg{
			shared.ID_NAME: graph.QueryArg{
				shared.ARG_EQ: id,
			},
		},
	}, authArgs))
}

func (s *Service) BatchQueryAssociations(
	association *graph.Association,
	ids []uint64,
	args graph.QueryArg,
) []map[string]interface{} {
	session, err := s.repository.OpenSession()
	if err != nil {
		panic(err.Error())
	}
	return session.BatchRealAssociations(association, ids, args)
}
