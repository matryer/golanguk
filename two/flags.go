package main

import "flag"

// START OMIT
func main() {
	var (
		mongoAddr = flag.String("mongo", "", "mongodb addr")
	)
	flag.Parse()
	//...
}

// END OMIT
