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

**Wrapper function** to effectively Delete the data created while testing, for every test so that the tests are run on fresh database everytime.

```
db := New("ark_test", "root", "mice")
db.Conn()
db.Wrap(t, func(tIn *testing.T, dIn *Db) {
//You tests, use t for test, and dIn.Instance.DataBaseSqlQueries()
}, "table_name1", "table_name2",...)
```

---

# Build - insert your struct data into mysql table in the easiest way possible

Currently tested only with one non-nested structs

All you need is a struct with mysql tags

```
type Table struct {
	Name      string    `mysql:"name"`
	CreatedAt time.Time `mysql:"created_at" // any name that ends with _at is considered datetime and is autofilled with time now.
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
A Query is generated like this,

```
INSERT INTO table_name (name, created_at) VALUES ("Sumit", "2015-03-16 20:23:22")
```

This will,
- insert the data into `table_name`
- and the `created_at` field is automatically filled in with datetime as now. (regex .*_at, which means, the column name should end with "_at" for this feature to work)

