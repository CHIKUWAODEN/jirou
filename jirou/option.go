package jirou

import "flag"

// Represents command line option
type Option struct {
	Help  bool
	Setup bool
	Port  int
}

func ParseOption() (option *Option) {
	var o = new(Option)
	flag.BoolVar(&o.Help, "help", false, "print help message")
	flag.BoolVar(&o.Setup, "setup", false, "insert records to database from file")
	flag.IntVar(&o.Port, "port", 8080, "set port number")
	flag.Parse()
	return o
}
