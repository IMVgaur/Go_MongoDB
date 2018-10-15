package main

import (
	"net/http"

	"github.com/Go_Mongo/TestApp/dao"

	"github.com/Go_Mongo/TestApp/controller"
)

func main() {
	dao.Init()
	http.ListenAndServe(":3000", controllers.Handlers())
}
