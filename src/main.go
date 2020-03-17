package main

import (
	"fmt"
	"net/http"
	"github.com/njoysubho/go-mux-example/src/routers"
	"github.com/njoysubho/go-mux-example/src/connection"
)

func main() {
	connection.CreateConnection()
	fmt.Println("Starting Server at port 80...")
	r := routers.Routers()
	http.ListenAndServe(":80", r)
    
}
