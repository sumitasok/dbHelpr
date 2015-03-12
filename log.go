package dbhelpr

type logger interface {
	Fatal(...interface{})
	Fatalf(string, ...interface{})
}
