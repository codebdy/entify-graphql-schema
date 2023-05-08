package parser

import (
	"fmt"
	"strings"

	"github.com/codebdy/entify-graphql-schema/consts"
	"github.com/codebdy/entify/model/graph"
	"github.com/graphql-go/graphql"
)

var queryFieldSDL = "\t%s(%s) : %s \n"

var objectWithKeySDL = `
type %s%s @key(fields: "id"){
	%s
}
`

var objectSDL = `
type %s%s {
	%s
}
`

var enumSDL = `
enum %s{
	%s
}
`

var inputSDL = `
input %s{
	%s
}
`

var comparisonSDL = `
input %s{
	%s
}
`

func (p *ModelParser) QuerySDL() (string, string) {
	queryFields := ""
	types := ""
	for _, enum := range p.model.Graph.Enums {
		types = types + enumToSDL(p.EnumType(enum.Name))
	}

	for _, enum := range p.DistinctOnEnums() {
		types = types + enumToSDL(enum)
	}

	types = types + enumToSDL(EnumOrderBy)
	for _, orderBy := range p.orderByMap {
		types = types + inputToSDL(orderBy)
	}

	types = types + comparisonToSDL(&BooleanComparisonExp)
	types = types + comparisonToSDL(&DateTimeComparisonExp)
	types = types + comparisonToSDL(&FloatComparisonExp)
	types = types + comparisonToSDL(&IntComparisonExp)
	types = types + comparisonToSDL(&IdComparisonExp)
	types = types + comparisonToSDL(&StringComparisonExp)

	for _, comparision := range p.enumComparisonExpMap {
		types = types + comparisonToSDL(comparision)
	}

	for _, where := range p.whereExpMap {
		types = types + inputToSDL(where)
	}

	for _, entity := range p.model.Graph.Entities {
		types = types + objectToSDL(p.EntityeOutputType(entity.Name()), true)
	}
	for _, entity := range p.model.Graph.RootEnities() {
		queryFields = queryFields + p.makeEntitySDL(entity)
	}

	for _, aggregate := range p.aggregateMap {
		types = types + objectToSDL(aggregate, false)
		fieldsType := aggregate.Fields()[consts.AGGREGATE].Type.(*graphql.Object)
		types = types + objectToSDL(fieldsType, false)

		for key := range fieldsType.Fields() {
			field := fieldsType.Fields()[key]
			if field.Name != consts.ARG_COUNT {
				types = types + objectToSDL(field.Type.(*graphql.Object), false)
			}
		}
	}

	for _, selectColumn := range p.selectColumnsMap {
		types = types + inputToSDL(selectColumn)
	}
	return queryFields, types
}

func (p *ModelParser) makeEntitySDL(entity *graph.Entity) string {
	sdl := ""
	sdl = sdl + fmt.Sprintf(queryFieldSDL,
		entity.QueryName(),
		makeArgsSDL(p.QueryArgs(entity.Name())),
		p.EntityListType(entity).String(),
	)

	sdl = sdl + fmt.Sprintf(queryFieldSDL,
		entity.QueryOneName(),
		makeArgsSDL(p.QueryArgs(entity.Name())),
		p.OutputType(entity.Name()).String(),
	)

	sdl = sdl + fmt.Sprintf(queryFieldSDL,
		entity.QueryAggregateName(),
		makeArgsSDL(p.QueryArgs(entity.Name())),
		(*p.aggregateType(entity)).String(),
	)

	return sdl
}

func makeArgsSDL(args graphql.FieldConfigArgument) string {
	var sdls []string
	for key := range args {
		sdls = append(sdls, key+":"+args[key].Type.Name())
	}
	return strings.Join(sdls, ",")
}

func makeArgArraySDL(args []*graphql.Argument) string {
	var sdls []string
	for _, arg := range args {
		sdls = append(sdls, arg.Name()+":"+arg.Type.Name())
	}
	return strings.Join(sdls, ",")
}

func objectToSDL(obj *graphql.Object, withKey bool) string {
	var intfNames []string
	implString := ""

	for _, intf := range obj.Interfaces() {
		intfNames = append(intfNames, intf.Name())
	}
	if len(intfNames) > 0 {
		implString = " implements " + strings.Join(intfNames, " & ")
	}

	sdl := objectSDL
	if withKey {
		sdl = objectWithKeySDL
	}
	return fmt.Sprintf(sdl, obj.Name(), implString, fieldsToSDL(obj.Fields()))
}

func enumToSDL(enum *graphql.Enum) string {
	var values []string

	sdl := enumSDL
	for _, value := range enum.Values() {
		values = append(values, value.Name)
	}
	return fmt.Sprintf(sdl, enum.Name(), strings.Join(values, "\n\t"))
}

func inputToSDL(input *graphql.InputObject) string {
	sdl := inputSDL
	return fmt.Sprintf(sdl, input.Name(), inputFieldsToSDL(input.Fields()))
}

func inputFieldsToSDL(fields graphql.InputObjectFieldMap) string {
	var fieldsStrings []string
	for key := range fields {
		field := fields[key]
		fieldsStrings = append(fieldsStrings, key+":"+field.Type.String())
	}

	return strings.Join(fieldsStrings, "\n\t")
}

func comparisonToSDL(comarison *graphql.InputObjectFieldConfig) string {
	sdl := comparisonSDL
	var comType *graphql.InputObject
	comType = comarison.Type.(*graphql.InputObject)
	return fmt.Sprintf(sdl, comType.Name(), inputFieldsToSDL(comType.Fields()))
}

func fieldsToSDL(fields graphql.FieldDefinitionMap) string {
	var fieldsStrings []string
	for i := range fields {
		field := fields[i]
		if len(field.Args) > 0 {
			fieldsStrings = append(fieldsStrings, fmt.Sprintf("%s(%s):%s", field.Name, makeArgArraySDL(field.Args), field.Type.String()))
		} else {
			fieldsStrings = append(fieldsStrings, field.Name+":"+field.Type.String())
		}
	}

	return strings.Join(fieldsStrings, "\n\t")
}
