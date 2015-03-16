package dbhelpr

import (
	"reflect"
	"regexp"
	"strings"
	"time"
)

type Resource interface {
	ResourceName() string
}

func (db Db) Build(r Resource) {
	numberOfColumns := reflect.TypeOf(r).NumField()

	queryString := "INSERT INTO " + r.ResourceName() + " (" + listOfTags(r) + ") VALUES(" + valueQuery(numberOfColumns) + ")"

	_, sCError := db.Instance.Exec(queryString, listOfValues(r)...)

	if sCError != nil {
		println(sCError.Error())
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

func listOfValues(r Resource) []interface{} {
	n := reflect.TypeOf(r).NumField()
	columnList := []interface{}{}
	for i := 0; i < n; i++ {
		// Check and Fill time.Time
		tagStr := reflect.TypeOf(r).FieldByIndex([]int{i}).Tag.Get("mysql")
		tagList := strings.Split(tagStr, ",")
		if len(tagList) != 0 {
			if condition, _ := regexp.MatchString(".*_at", tagList[0]); condition {
				if v, ok := reflect.ValueOf(r).FieldByIndex([]int{i}).Interface().(time.Time); ok {
					const layout = "0001-01-01 00:00:00 +0000 UTC"
					if v.String() == layout {
						columnList = append(columnList, time.Now().UTC())
					} else {
						value := reflect.ValueOf(r).FieldByIndex([]int{i}).Interface()
						columnList = append(columnList, value)
					}
				}
			} else {
				value := reflect.ValueOf(r).FieldByIndex([]int{i}).Interface()
				columnList = append(columnList, value)
			}
		}

	}
	return columnList

}

func valueQuery(n int) string {
	valueQuery := []string{}
	for i := 0; i < n; i++ {
		valueQuery = append(valueQuery, "?")
	}
	return strings.Join(valueQuery, ",")
}
