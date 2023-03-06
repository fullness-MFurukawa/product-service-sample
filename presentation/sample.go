package presentation

import "github.com/gin-gonic/gin"

type GinSampleController struct {
}

func (grc *GinSampleController) GetSampleHandler(c *gin.Context) {
	dto := SampleDto{Message: "Hello World!!!"}
	c.JSON(200, dto)
}
