package dbhelpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tTable struct {
	Field string `mysql:"column_name"`
}

func TestColumnNameSuccess(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("column_name", ColumnName(tTable{}))
	assert.True(true)
}
