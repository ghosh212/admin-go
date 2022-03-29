package main

import (
	"admin-go/base"
	mongoConfig "admin-go/config"
	"fmt"
	"log"
)

const basePath = "/" + "admin" + "/"
const port = ":8080"

func main() {
	fmt.Println("Let Us Begin")
	var s base.Service
	//making server endpoints
	e := base.MakeServerEndpoints(s)
	go func() {
		log.Default().Println("APPLICATION STARTING UP")
		//http.ListenAndServe()
	}()
	base.MakeHTTPPHandler(e, basePath, port)
	//mongo db connection
	client, ctx, cancel, err := mongoConfig.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	//mongo fdb connection close
	defer mongoConfig.Close(client, ctx, cancel)

	//to check the connectivity status for database
	//mongoConfig.Ping(client, ctx)

}
