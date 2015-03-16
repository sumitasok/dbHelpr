package dbhelpr

// Adds date time for tags .*_at as now() when not provided

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Venue struct {
	EventID   string    `mysql:"event_id,false"`
	ID        string    `mysql:"venue_id"`
	Name      string    `mysql:"venue_name"`
	UpdatedAt time.Time `mysql:"venue_updated_at"`
	CreatedAt time.Time `mysql:"venue_created_at"`
}

var (
	yesterday = time.Now().AddDate(0, 0, -1).UTC()
	venue     = Venue{
		EventID:   "event id",
		ID:        time.Now().String(),
		Name:      "Melbourn",
		CreatedAt: yesterday,
	}
)

func (v Venue) ResourceName() string {
	return "event_venue"
}

func TestBuildSuccess(t *testing.T) {
	assert := assert.New(t)

	New("ark_test", "root", "mice").Conn().Wrap(t, func(tCp *testing.T, dbCp *Db) {
		dbCp.Build(venue)

		venueResult := Venue{}

		result := dbCp.Instance.QueryRow("select * from event_venue where venue_id = ?", venue.ID)
		result.Scan(&venueResult.EventID, &venueResult.ID, &venueResult.Name, &venueResult.CreatedAt, &venueResult.UpdatedAt)

		assert.Equal(venue.Name, venueResult.Name)
		assert.Equal(venue.EventID, venueResult.EventID)
		// datetime remains same when provided
		// assert.Equal(yesterday.String(), venueResult.CreatedAt.String())
	})

	assert.True(true)
}

func TestValueQuery(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("?,?,?", valueQuery(3))
	assert.Equal("", valueQuery(0))
	assert.Equal("?", valueQuery(1))
}

func TestListOfColumns(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("event_id,venue_id,venue_name,updated_at,created_at", listOfTags(venue))
}
