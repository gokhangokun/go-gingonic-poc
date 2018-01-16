package main

import (
	"net/http"

	"../go-gingonic-poc/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	sample := new(controllers.SampleController) //Controller instance
	r.GET("/api/sample", sample.GetSample)      //Simple route
	http.ListenAndServe("localhost:3000", r)
}
