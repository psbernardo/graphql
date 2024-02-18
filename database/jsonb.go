package database

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/psbernardo/graphql/graph/model"
)

// JSONB Interface for JSONB Field of yourTableName Table
type QueryResultJSONB struct {
	value model.QueryResult
}

func NewQueryResultJSONB(value model.QueryResult) QueryResultJSONB {
	return QueryResultJSONB{value: value}
}

func (c QueryResultJSONB) Value() (driver.Value, error) {

	// jsonStr, err := json.Marshal(c.value)
	// if err != nil {
	// 	return nil, err
	// }
	// msgStr := "[" + string(jsonStr) + "]"
	return json.Marshal(c.value)
}

func (c *QueryResultJSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &c.value)
}
