package dbhelpr

/*
	If Conn gets closed accidentaly, you can call conn on the existing object any time
*/

import (
	sql "database/sql"
	// needed for dtabase sql connection
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

type Db struct {
	Name     string
	Username string
	Password string
	Instance *sql.DB
	Logger   logger
	Tables   []table
}

type table struct {
	Name string
}

func New(name string, user string, pswd string) *Db {
	return &Db{
		Name:     name,
		Username: user,
		Password: pswd,
	}
}

func (d *Db) Log(l logger) *Db {
	d.Logger = l
	return d
}

func (d *Db) log(args ...interface{}) {
	if d.Logger != nil {
		d.Logger.Fatal(args)
		return
	}
	panic(args)
}

func (d *Db) Conn() *Db {
	db, err := sql.Open("mysql", d.Username+":"+d.Password+"@/"+d.Name)

	if err != nil {
		d.log(err.Error())
	}

	d.Instance = db
	// defer d.Instance.Close()

	pingErr := d.Instance.Ping()

	if pingErr != nil {
		d.log(pingErr)
	}

	return d
}

func (d *Db) Close() {
	d.Instance.Close()
}

func (d *Db) Wrap(t *testing.T, testFn func(*testing.T, *Db), tables ...string) *Db {
	d.clean(tables...)
	testFn(t, d)
	d.clean(tables...)
	d.Close()
	return d
}

func (d *Db) clean(tables ...string) *Db {

	if len(tables) == 0 {
		for _, table := range d.Tables {
			d.Truncate(table.Name)
		}
		return d
	}

	for _, table := range tables {
		d.Truncate(table)
	}
	return d
}

func (d *Db) Truncate(table string) *Db {
	stmt, _ := d.Instance.Prepare("truncate " + table)
	defer stmt.Close()
	_, err := stmt.Exec()
	if err != nil {
		d.log(err)
	}
	return d
}
