package controllers

import (
	"github.com/gin-gonic/gin"
)

//SampleController struct
type SampleController struct{}

//GetSample function
func (ctrl SampleController) GetSample(c *gin.Context) {
	c.JSON(200, "sample response")
	return
}
