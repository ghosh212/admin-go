package main

import (
	mongoConfig "admin-go/config"
	"fmt"
)

func main() {
	fmt.Println("Hellow World from Go")

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
