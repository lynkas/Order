package main

import (
	"github.com/go-bongo/bongo"
	"log"
)

var config = &bongo.Config{
	ConnectionString: "0.0.0.0",
	Database:         "test",
}

var All struct{
	Connection *bongo.Connection
}



func main() {
	var err error
	All.Connection, err = bongo.Connect(config)
	if err != nil {
		log.Fatal(err)
	}
	//_ =CONNECTION


	r:=routers()

	r.Run() // listen and serve on 0.0.0.0:8080

}
