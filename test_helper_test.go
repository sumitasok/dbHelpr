package dbhelpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func (l tLogger) Fatal(args ...interface{}) {
	println("error mesage successfully printing - No Error")
}

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

func TestWrapTest(t *testing.T) {
	// assert := assert.New(t)
	db := New("ark_test", "root", "mice")
	db.Conn()
	db.Wrap(t, func(t *testing.T, d *Db) {
		db := &d.Instance

		randomEventID := "eventId"
		stmtCreate, sCError := db.Prepare("INSERT INTO " + "event_venue" + " VALUES(?,?,?,?,?)")
		stmtCreate.Exec(randomEventID, "venueID", "venueName", time.Now().UTC(), time.Now().UTC())

		if sCError != nil {
			t.Fatal(sCError)
		}

		// Need to make sure this completes before next query is executed.

		var eventId string
		err := db.QueryRow("SELECT event_id FROM event_venue WHERE event_id = ?", "eventId").Scan(&eventId)

		switch {
		case err == d.ErrNoRow():
			t.Fatal("No Event returned")
			break
		case err != nil:
			t.Fatal(err)
			break
		default:
			break
		}

		if randomEventID != eventId {
			t.Fatal("randomEventId doesn't match with returned eventId", eventId)
		}
	}, "event_venue")
	defer db.Close()

	// var eventId string
	// err := New("ark_test", "root", "mice").Conn().Instance.QueryRow("SELECT event_id FROM event_venue WHERE event_id = ?", "eventId").Scan(&eventId)
	// assert.Error(err)

}

func TestCleanSuccess(t *testing.T) {
	db := New("ark_test", "root", "mice").Conn()
	db.clean("event_venue", "venue_audi")
}
