package main

import (
	"fmt"
	"github.com/omid-h70/bookstore/users-api/adapters/router"
	"github.com/omid-h70/bookstore/users-api/app"
)

func Test(r *router.Router) {

}

func main() {
	fmt.Println("Hi i'm up")
	app.StartServer()
}
