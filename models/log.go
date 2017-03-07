package models

// Log is the collection of methods used for logging
type Log interface {
	Infoln(args ...interface{})
	Errorln(args ...interface{})
	Println(args ...interface{})
	Printf(format string, args ...interface{})
	Fatalln(args ...interface{})
}
