package gin

import (
	"product-service/presentation"

	"github.com/gin-gonic/gin"
)

type GinSampleController struct {
}

func (grc *GinSampleController) GetSampleHandler(c *gin.Context) {
	dto := presentation.SampleDto{Message: "Hello World!!!"}
	c.JSON(200, dto)
}
