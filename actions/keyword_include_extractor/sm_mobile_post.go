package keyword_include_extractor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gitlab.fxt.cn/fxt/rank-util/services/keyword_include_extractor_service"
	"gitlab.fxt.cn/fxt/rank-util/structs/include_extractor"
)

func KeywordIncludeExtractorSmMobile(c *gin.Context)  {
	req := &include_extractor.IncludeExtractorRequest{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求格式不正确",
		})
		return
	}

	res, err := keyword_include_extractor_service.KeywordIncludeExtractorSmMobile(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}