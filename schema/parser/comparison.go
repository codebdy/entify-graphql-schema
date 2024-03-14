package parser

import (
	"github.com/codebdy/entify-core/model/graph"
	"github.com/codebdy/entify-core/shared"
	"github.com/graphql-go/graphql"
)

var BooleanComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "BooleanComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				shared.ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				shared.ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				shared.ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				shared.ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Boolean),
				},
				shared.ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				shared.ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				shared.ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				shared.ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				shared.ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Boolean),
				},
			},
		},
	),
}

var DateTimeComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "DateTimeComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				shared.ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				shared.ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				shared.ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				shared.ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.DateTime),
				},
				shared.ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				shared.ARG_ISNOTNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				shared.ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				shared.ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				shared.ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.DateTime,
				},
				shared.ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.DateTime),
				},
			},
		},
	),
}

var FloatComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "FloatComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				shared.ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				shared.ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				shared.ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				shared.ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Float),
				},
				shared.ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				shared.ARG_ISNOTNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				shared.ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				shared.ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				shared.ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				shared.ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Float),
				},
			},
		},
	),
}

var IntComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "IntComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				shared.ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				shared.ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				shared.ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				shared.ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Int),
				},
				shared.ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				shared.ARG_ISNOTNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				shared.ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				shared.ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				shared.ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				shared.ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.Int),
				},
			},
		},
	),
}

var IdComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "IdComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				shared.ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.ID,
				},
				shared.ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.ID,
				},
				shared.ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.ID,
				},
				shared.ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.ID),
				},
				shared.ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.ID,
				},
				shared.ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.ID,
				},
				shared.ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.ID,
				},
				shared.ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.ID),
				},
				shared.ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				shared.ARG_ISNOTNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
			},
		},
	),
}

var StringComparisonExp = graphql.InputObjectFieldConfig{
	Type: graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "StringComparisonExp",
			Fields: graphql.InputObjectConfigFieldMap{
				shared.ARG_EQ: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				shared.ARG_GT: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				shared.ARG_GTE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				shared.ARG_ILIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				shared.ARG_IN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
				// shared.ARG_IREGEX: &graphql.InputObjectFieldConfig{
				// 	Type: graphql.String,
				// },
				shared.ARG_ISNULL: &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				shared.ARG_LIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				shared.ARG_LT: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				shared.ARG_LTE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				shared.ARG_NOTEQ: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				shared.ARG_NOTILIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				shared.ARG_NOTIN: &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
				// shared.ARG_NOTIREGEX: &graphql.InputObjectFieldConfig{
				// 	Type: graphql.String,
				// },
				shared.ARG_NOTLIKE: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				shared.ARG_NOTREGEX: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				// shared.ARG_NOTSIMILAR: &graphql.InputObjectFieldConfig{
				// 	Type: graphql.String,
				// },
				shared.ARG_REGEX: &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				// shared.ARG_SIMILAR: &graphql.InputObjectFieldConfig{
				// 	Type: graphql.String,
				// },
			},
		},
	),
}

func (p *ModelParser) EnumComparisonExp(attr *graph.Attribute) *graphql.InputObjectFieldConfig {
	enumEntity := attr.EumnType
	if enumEntity == nil {
		panic("Can not find enum entity, please check the attribute enum type. entity name: " + attr.EnityType.Name() + " attribute name: " + attr.Name)
	}
	if p.enumComparisonExpMap[enumEntity.Name] != nil {
		return p.enumComparisonExpMap[enumEntity.Name]
	}
	enumType := graphql.String //p.EnumType(enumEntity.Name)
	enumxp := graphql.InputObjectFieldConfig{
		Type: graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name: enumEntity.Name + "EnumComparisonExp",
				Fields: graphql.InputObjectConfigFieldMap{
					shared.ARG_EQ: &graphql.InputObjectFieldConfig{
						Type: enumType,
					},
					shared.ARG_IN: &graphql.InputObjectFieldConfig{
						Type: graphql.NewList(enumType),
					},
					shared.ARG_ISNULL: &graphql.InputObjectFieldConfig{
						Type: graphql.Boolean,
					},
					shared.ARG_NOTEQ: &graphql.InputObjectFieldConfig{
						Type: enumType,
					},
					shared.ARG_NOTIN: &graphql.InputObjectFieldConfig{
						Type: graphql.NewList(enumType),
					},
				},
			},
		),
	}
	p.enumComparisonExpMap[enumEntity.Name] = &enumxp
	return &enumxp
}
