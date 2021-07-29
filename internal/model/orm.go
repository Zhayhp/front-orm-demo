package model

type TableName = string // 表名
type Condition = string
type Value = string
type SearchType = string

type RetrieveReq struct {
	List map[TableName]map[Condition]Value
}

//type RetrieveReq map[SearchType]map[TableName]map[Condition]Value

