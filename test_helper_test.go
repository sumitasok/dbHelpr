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
	defer db.Close()

	assert.NotNil(db.Instance)
}

/* a dummy logger for testing purpose
in real, you can use both package log and testing.T */
type tLogger struct{}

func (l tLogger) Fatal(args ...interface{}) {}

func (l tLogger) Fatalf(s string, args ...interface{}) {}

func TestDbConnNotPanicsWhenLoggerSet(t *testing.T) {
	assert := assert.New(t)

	db := New("nodb", "root", "mice").Log(tLogger{})

	assert.NotPanics(func() {
		db.Conn()
		defer db.Close()
	})
}

func TestDbPing(t *testing.T) {
	assert := assert.New(t)

	db := New("nodb", "root", "mice")

	assert.Panics(func() {
		db.Conn()
		defer db.Close()
	})

	db = New("ark_test", "root", "mice")

	assert.NotPanics(func() {
		db.Conn()
		defer db.Close()
	})
}

func TestDbTruncate(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		db := New("ark_test", "root", "mice")
		db.Conn()
		defer db.Close()
		db.Truncate("noTable")
	})

	assert.NotPanics(func() {
		db := New("ark_test", "root", "mice")
		db.Conn()
		defer db.Close()
		db.Truncate("event_venue")
	})
}
