package dbhelpr

import (
	sql "database/sql"
	// needed for dtabase sql connection
	_ "github.com/go-sql-driver/mysql"
)

type Db struct {
	Name     string
	Username string
	Password string
	Instance *sql.DB
}

func New(name string, user string, pswd string) *Db {
	return &Db{
		Name:     name,
		Username: user,
		Password: pswd,
	}
}

func (d *Db) Conn() *Db {
	db, err := sql.Open("mysql", d.Username+":"+d.Password+"@/"+d.Name)

	if err != nil {
		panic(err.Error())
	}

	d.Instance = db
	defer d.Instance.Close()

	pingErr := d.Instance.Ping()

	if pingErr != nil {
		panic(pingErr)
	}

	return d
}
