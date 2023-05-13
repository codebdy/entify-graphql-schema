package consts

const (
	ROOT_QUERY_NAME        = "Query"
	ROOT_MUTATION_NAME     = "Mutation"
	ROOT_SUBSCRIPTION_NAME = "Subscription"
)

const (
	ONE        = "one"
	QUERY      = "query"
	MUTATION   = "mutation"
	AGGREGATE  = "aggregate"
	LIST       = "list"
	FIELDS     = "Fields"
	NODES      = "nodes"
	TOTAL      = "total"
	INPUT      = "Input"
	SET_INPUT  = "Set"
	UPSERT     = "upsert"
	UPSERT_ONE = "upsertOne"
	INSERT     = "insert"
	INSERT_ONE = "insertOne"
	UPDATE     = "update"
	UPDATE_ONE = "updateOne"
	DELETE     = "delete"
	BY_ID      = "ById"
	SET        = "set"
	HAS_MANY   = "HasMany"
	HAS_ONE    = "HasOne"
	ENTITY     = "Entity"

	ARG_DISTINCTON string = "distinctOn"
	ARG_LIMIT      string = "limit"
	ARG_OFFSET     string = "offset"
	ARG_ORDERBY    string = "orderBy"
	ARG_WHERE      string = "where"

	ARG_ADD     string = "add"
	ARG_DELETE  string = "delete"
	ARG_UPDATE  string = "update"
	ARG_SYNC    string = "sync"
	ARG_CASCADE string = "cascade"
	ARG_CLEAR   string = "clear"

	ARG_AND string = "_and"
	ARG_NOT string = "_not"
	ARG_OR  string = "_or"
)

//EQ("="), GTE(">="), GT(">"), LT("<"), LTE("<=");
const (
	ARG_EQ     string = "_eq"
	ARG_GT     string = "_gt"
	ARG_GTE    string = "_gte"
	ARG_IN     string = "_in"
	ARG_ISNULL string = "_isNull"
	ARG_LT     string = "_lt"
	ARG_LTE    string = "_lte"
	ARG_NOTEQ  string = "_notEq"
	ARG_NOTIN  string = "_notIn"

	ARG_ILIKE string = "_iLike"
	// ARG_IREGEX     string = "_iregex"
	ARG_LIKE     string = "_like"
	ARG_NOTILIKE string = "_notILike"
	// ARG_NOTIREGEX  string = "_notIRegex"
	ARG_NOTLIKE  string = "_notLike"
	ARG_NOTREGEX string = "_notRegexp"
	// ARG_NOTSIMILAR string = "_notSimilar"
	ARG_REGEX string = "_regexp"
	// ARG_SIMILAR    string = "_similar"
)

const (
	ARG_COUNT    string = "count"
	ARG_COLUMNS  string = "columns"
	ARG_DISTINCT string = "distinct"
)

const (
	ARG_OBJECT            string = "object"
	ARG_OBJECTS           string = "objects"
	RESPONSE_RETURNING    string = "returning"
	RESPONSE_AFFECTEDROWS string = "affectedRows"
	ARG_SET               string = "set"
	ARG_FILE              string = "file"
	ARG_FOLDER            string = "folder"
)

const (
	UUID    string = "uuid"
	INNERID string = "innerId"
	TYPE    string = "type"
)

const (
	ID_SUFFIX     string = "_id"
	PIVOT         string = "pivot"
	INDEX_SUFFIX  string = "_idx"
	SUFFIX_SOURCE string = "_source"
	SUFFIX_TARGET string = "_target"
)

const (
	BOOLEXP           string = "BoolExp"
	ORDERBY           string = "OrderBy"
	DISTINCTEXP       string = "DistinctExp"
	MUTATION_RESPONSE string = "MutationResponse"
)
