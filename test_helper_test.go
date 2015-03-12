package dbhelpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDB(t *testing.T) {
	assert := assert.New(t)

	db := New("db_name", "user_android", "android_password")

	assert.Equal("db_name", db.Name)
}

func TestDbConn(t *testing.T) {
	assert := assert.New(t)

	db := New("ark_test", "root", "mice")

	assert.Nil(db.Instance)

	db.Conn()

	assert.NotNil(db.Instance)
}

func TestDbPing(t *testing.T) {
	assert := assert.New(t)

	db := New("nodb", "root", "mice")

	assert.Panics(func() {
		db.Conn()
	})

	db = New("ark_test", "root", "mice")

	assert.NotPanics(func() {
		db.Conn()
	})
}
