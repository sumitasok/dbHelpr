package dbhelpr

import (
	"reflect"
	"strings"
)

type DbDetails struct {
	TagIdentifier string
}

// ColumnName will return the name of the column in mysql database
func (d *DbDetails) ColumnName(table interface{}) string {
	field := "Field"
	tagValue, ok := reflect.TypeOf(table).FieldByName(field)
	if ok != true {
		panic("Field with name '" + field + "'' Doesn't exist")
	}
	tagStr := tagValue.Tag.Get(d.TagIdentifier)
	sep := ","
	index := 0
	for i, tag := range strings.Split(tagStr, sep) {
		if i == index {
			return tag
		}
	}

	panic("tag not found! for " + field)
}

func fieldName(field interface{}) string {
	// t := reflect.TypeOf(field)
	// v := reflect.ValueOf(field)
	return ""
}
