package dbhelpr

import (
	"reflect"
	"strings"
	"time"
)

type Resource interface {
	ResourceName() string
}

func (db Db) Build(r Resource) {
	numberOfColumns := reflect.TypeOf(r).NumField()

	queryString := "INSERT INTO " + r.ResourceName() + " (" + listOfTags(r) + ") VALUES(" + valueQuery(numberOfColumns) + ")"
	println(queryString)
	stmtCreate, sCError := db.Instance.Prepare(queryString)
	x := []interface{}{
		"eventID", "venueID", "venueName", time.Now().UTC(), time.Now().UTC(),
	}
	stmtCreate.Exec(x...)

	if sCError != nil {
		db.Logger.Fatal(sCError)
	}
}

func listOfTags(r Resource) string {
	n := reflect.TypeOf(r).NumField()
	columnList := []string{}
	for i := 0; i < n; i++ {
		tagStr := reflect.TypeOf(r).FieldByIndex([]int{i}).Tag.Get("mysql")
		tagList := strings.Split(tagStr, ",")
		if len(tagList) == 0 {
			break
		}
		columnList = append(columnList, tagList[0])
	}
	return strings.Join(columnList, ",")
}

func valueQuery(n int) string {
	valueQuery := []string{}
	for i := 0; i < n; i++ {
		valueQuery = append(valueQuery, "?")
	}
	return strings.Join(valueQuery, ",")
}
