package script

import (
	"context"
	"log"

	"github.com/codebdy/entify-core"
	"github.com/codebdy/entify-core/model/observer"
	"github.com/codebdy/entify-core/orm"
	"github.com/codebdy/entify-core/shared"
)

type EntifyService struct {
	ctx context.Context
	//roleIds []uint64
	repository *entify.Repository
	session    *orm.Session
}

func NewService(ctx context.Context, repository *entify.Repository) *EntifyService {
	session, err := repository.OpenSession()
	if err != nil {
		panic(err.Error())
	}
	return &EntifyService{
		ctx:        ctx,
		repository: repository,
		session:    session,
	}
}

func (s *EntifyService) BeginTx() {
	// session, err := orm.Open(s.repository.DbConfig, s.repository.Model)
	// if err != nil {
	// 	log.Panic(err.Error())
	// }
	// s.session = session
	err := s.session.BeginTx()
	if err != nil {
		log.Panic(err.Error())
	}
}

func (s *EntifyService) Commit() {
	if s.session == nil {
		log.Panic("No session to commit")
	}
	err := s.session.Commit()

	if err != nil {
		log.Panic(err.Error())
	}
}

func (s *EntifyService) ClearTx() {
	if s.session == nil {
		log.Panic("No session to ClearTx")
	}
	s.session.ClearTx()
	s.session = nil
}

func (s *EntifyService) Rollback() {
	if s.session == nil {
		log.Panic("No session to Rollback")
	}

	err := s.session.Dbx.Rollback()
	if err != nil {
		log.Panic(err.Error())
	}
	s.session = nil
}

func (s *EntifyService) checkSession() {
	if s.session == nil {
		session, err := orm.Open(s.repository.DbConfig, s.repository.Model)
		if err != nil {
			log.Panic(err.Error())
		}
		s.session = session
	}
}

func (s *EntifyService) QueryOne(entityName string, args map[string]interface{}) interface{} {
	return s.session.QueryOne(entityName, args)
}

func (s *EntifyService) Save(entityName string, objects []interface{}) []orm.InsanceData {
	s.checkSession()

	if len(objects) > 0 {
		objects := s.session.QueryByIds(entityName, objects)
		observer.EmitObjectMultiPosted(objects, entityName, s.ctx)
	}

	return []orm.InsanceData{}
}

func (s *EntifyService) SaveOne(entityName string, object interface{}) interface{} {
	s.checkSession()

	if object == nil {
		log.Panic("Object to save is nil")
	}

	id, err := s.session.SaveOne(entityName, object.(map[string]interface{}))
	if err != nil {
		log.Panic(err.Error())
	}

	result := s.session.QueryOneById(entityName, id)
	observer.EmitObjectPosted(result.(map[string]interface{}), entityName, s.ctx)
	return result
}

func (s *EntifyService) WriteLog(
	operate string,
	result string,
	message string,
) {
	//logs.WriteBusinessLog(s.ctx, operate, result, message)
}

func (s *EntifyService) EmitNotification(text string, noticeType string, userId uint64) {
	s.SaveOne(
		"Notification",
		map[string]interface{}{
			"text":       text,
			"noticeType": noticeType,
			"user": map[string]interface{}{
				"sync": map[string]interface{}{
					shared.ID_NAME: userId,
				},
			},
			"app": map[string]interface{}{
				"sync": map[string]interface{}{
					//consts.ID: contexts.Values(s.ctx).AppId,
				},
			},
		},
	)
}

//切换成对象形式的接口
// func (s *ScriptService) Query(gql string, variables interface{}) interface{} {
// 	var newVariables map[string]interface{}

// 	if variables != nil {
// 		newVariables = variables.(map[string]interface{})
// 	}
// 	params := graphql.Params{
// 		Schema:         register.GetSchema(s.ctx),
// 		RequestString:  gql,
// 		VariableValues: newVariables,
// 		Context:        context.WithValue(s.ctx, "gql", gql),
// 	}

// 	r := graphql.Do(params)
// 	if len(r.Errors) > 0 {
// 		log.Printf("failed to execute graphql operation, errors: %+v", r.Errors)
// 		log.Panic(r.Errors[0].Error())
// 	}

// 	return r.Data
// }
