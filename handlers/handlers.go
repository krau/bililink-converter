package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/krau/bililink-converter/common"
)

// @Summary Convert Bv to Av
// @Description Convert Bv to Av
// @Tags Convert
// @Accept json
// @Produce json
// @Param bv path string true "Bv"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Invalid Bv"
// @Router /bv2av/{bv} [get]
func Bv2avEndpoint(c *gin.Context) {
	bv := c.Param("bv")
	av, err := common.DecodeBV(bv)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"link": fmt.Sprintf("https://www.bilibili.com/video/av%d", av),
	})
}

// @Summary Convert Av to Bv
// @Description Convert Av to Bv
// @Tags Convert
// @Accept json
// @Produce json
// @Param av path string true "Av"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Invalid Av"
// @Router /av2bv/{av} [get]
func Av2bvEndpoint(c *gin.Context) {
	av, err := strconv.Atoi(c.Param("av"))
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	bv, err := common.EncodeBV(av)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"link": fmt.Sprintf("https://www.bilibili.com/video/%s", bv),
	})
}

// @Summary Convert B23 to Av
// @Description Convert B23 to Av
// @Tags Convert
// @Accept json
// @Produce json
// @Param b23code path string true "B23 code"
// @Param av query bool false "Convert to Av"
// @Param iframe query bool false "Convert to iframe"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Invalid B23 code"
// @Router /b23/{b23code} [get]
func B23ConvertEndpoint(c *gin.Context) {
	b23code := c.Param("b23code")
	var convertedLink string
	if c.Query("iframe") == "true" {
		convertedLink, err := common.B23toIframe(b23code)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		c.JSON(200, gin.H{
			"link": convertedLink,
		})
		return
	}
	if c.Query("av") == "true" {
		convertedLink, err := common.B23toAv(b23code)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		c.JSON(200, gin.H{
			"link": convertedLink,
		})
		return
	}
	convertedLink, err := common.ConvertB23(b23code)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"link": convertedLink,
	})
}
