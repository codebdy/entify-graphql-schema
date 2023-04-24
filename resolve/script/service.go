package script

import (
	"context"
	"log"

	"github.com/codebdy/entify"
	"github.com/codebdy/entify/model/observer"
	"github.com/codebdy/entify/model/observer/consts"
	"github.com/codebdy/entify/orm"
)

type ScriptService struct {
	ctx context.Context
	//roleIds []uint64
	repository *entify.Repository
	session    *orm.Session
}

func NewService(ctx context.Context, repository *entify.Repository) *ScriptService {

	return &ScriptService{
		ctx:        ctx,
		repository: repository,
		//roleIds: service.QueryRoleIds(ctx, model),
	}
}

func (s *ScriptService) SetSession(session *orm.Session) {
	s.session = session
}

func (s *ScriptService) BeginTx() {
	session, err := orm.Open(s.repository.DbConfig, s.repository.Model)
	if err != nil {
		log.Panic(err.Error())
	}
	s.session = session
	err = session.BeginTx()
	if err != nil {
		log.Panic(err.Error())
	}
}

func (s *ScriptService) Commit() {
	if s.session == nil {
		log.Panic("No session to commit")
	}
	err := s.session.Commit()

	if err != nil {
		log.Panic(err.Error())
	}
}

func (s *ScriptService) ClearTx() {
	if s.session == nil {
		log.Panic("No session to ClearTx")
	}
	s.session.ClearTx()
	s.session = nil
}

func (s *ScriptService) Rollback() {
	if s.session == nil {
		log.Panic("No session to Rollback")
	}

	err := s.session.Dbx.Rollback()
	if err != nil {
		log.Panic(err.Error())
	}
	s.session = nil
}

func (s *ScriptService) checkSession() {
	if s.session == nil {
		session, err := orm.Open(s.repository.DbConfig, s.repository.Model)
		if err != nil {
			log.Panic(err.Error())
		}
		s.session = session
	}
}

func (s *ScriptService) Save(objects []interface{}, entityName string) []orm.InsanceData {
	s.checkSession()

	if len(objects) > 0 {
		objects := s.session.QueryByIds(entityName, objects)
		observer.EmitObjectMultiPosted(objects, entityName, s.ctx)
	}

	return []orm.InsanceData{}
}

func (s *ScriptService) SaveOne(object interface{}, entityName string) interface{} {
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

func (s *ScriptService) WriteLog(
	operate string,
	result string,
	message string,
) {
	//logs.WriteBusinessLog(s.ctx, operate, result, message)
}

func (s *ScriptService) EmitNotification(text string, noticeType string, userId uint64) {
	s.SaveOne(
		map[string]interface{}{
			"text":       text,
			"noticeType": noticeType,
			"user": map[string]interface{}{
				"sync": map[string]interface{}{
					consts.ID: userId,
				},
			},
			"app": map[string]interface{}{
				"sync": map[string]interface{}{
					//consts.ID: contexts.Values(s.ctx).AppId,
				},
			},
		},
		"Notification",
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
