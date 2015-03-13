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

This will exit as `t.Fatal` is called internally

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


