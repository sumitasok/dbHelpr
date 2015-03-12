package dbhelpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumnNameSuccess(t *testing.T) {
	assert := assert.New(t)

	type tTable struct {
		Field string `mysql:"column_name"`
	}

	d := DbDetails{
		TagIdentifier: "mysql",
	}

	assert.Equal("column_name", d.ColumnName(tTable{}))
	assert.True(true)
}
