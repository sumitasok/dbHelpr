Usage

Mysql support

This is how you provide the dataSourceName
```
db := New("database_name", "user_name", "password")
```

When you want to Conn
```
db.Conn()
```

If you don't want your application to panic in error
Or, if you are using this in test files, and want to couple the errors with test errors

```
func TestXXX(t *testing.T) {
	db := New("database_name", "user_name", "password").Conn()
	db.Log(t)	
}
```
You are setting **pointer to testing.T instance as Logger to dbHelper
This will let the test exit when any database error happens as `t.Fatal` is called internally

Alternatively, you can pass `log` packages `log.New(io.Writer, string, int) *Logger`


---

**Wrapper function** to effectively truncate the data created while testing

```
db := New("ark_test", "root", "mice")
db.Conn()
db.Wrap(t, func(tIn *testing.T, dIn *Db) {
//You tests, use t for test, and dIn.Instance.DataBaseSqlQueries()
}, "table_name1", "table_name2",...)
```

---

# Build - insert into mysql using just a struct

All you need is a struct with mysql tags

```
type Table struct {
	Name      string    `mysql:"name"`
	CreatedAt time.Time `mysql:"created_at" // any name that ends with _at is considered datetime and is autofilled it with time now.
}

func (t Table) ResourceName() string {
	return "table_name" // mysql table name
}
```

note: `name` and `created_at` are the column names in the table.

```

tRow := Table{"Sumit"}

dbhelpr.Build(tRow)

```

This will,
- insert the data into `table_name`
- and the `created_at` field is automatically filled in with datetime as now. (regex .*_at, which means, the column name should end with "_at" for this feature to work)

