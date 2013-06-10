package nullobj

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

type NullFloat64 struct {
	sql.NullFloat64
}

type NullInt64 struct {
	sql.NullInt64
}

type NullBool struct {
	sql.NullBool
}

func (c *NullString) MarshalJSON() ([]byte, error) {
	if c.String == "" && !c.Valid {
		return json.Marshal("")
	}
	return json.Marshal(c.String)
}

func (c *NullInt64) MarshalJSON() ([]byte, error) {
	if !c.Valid {
		return json.Marshal("")
	}
	return json.Marshal(c.Int64)
}

func (c *NullFloat64) MarshalJSON() ([]byte, error) {
	if !c.Valid {
		return json.Marshal("")
	}
	return json.Marshal(c.Float64)
}

func (c *NullBool) MarshalJSON() ([]byte, error) {
	if !c.Valid {
		return json.Marshal("")
	}
	return json.Marshal(c.Bool)
}
