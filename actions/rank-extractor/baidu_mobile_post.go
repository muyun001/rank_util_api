package rank_extractor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gitlab.fxt.cn/fxt/rank-util/services/rank_extractor_service"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
)

func RankExtractorBaiduMobile(c *gin.Context) {
	req := &rank_extractor.RankExtractorRequest{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求格式不正确",
		})
		return
	}

	if req.CheckMatch == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "CheckMatch不能为空",
		})
		return
	}

	res, err := rank_extractor_service.RankExtractorBaiduMobile(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}

	c.JSON(http.StatusOK, res)
}
