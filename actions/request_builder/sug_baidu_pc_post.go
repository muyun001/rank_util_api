package request_builder

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gitlab.fxt.cn/fxt/rank-util/services/request_build_service"
	"gitlab.fxt.cn/fxt/rank-util/structs/request_builder"
)

func RequestBuilderSugBaiduPc(c *gin.Context) {
	rbr := &request_builder.RequestBuilderRequest{}
	err := c.BindJSON(rbr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求格式不正确",
		})
		return
	}

	dcRequest := request_build_service.BuildDcRequestSugBaiduPc(rbr)

	c.JSON(http.StatusOK, dcRequest)
}