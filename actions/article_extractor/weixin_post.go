package article_extractor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gitlab.fxt.cn/fxt/rank-util/services/article_extractor_service"
	"gitlab.fxt.cn/fxt/rank-util/structs/article_extractor"
)

func ArticleExtractorWeixin(c *gin.Context) {
	req := &article_extractor.ArticleExtractorRequest{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求格式不正确",
		})
		return
	}

	res, err := article_extractor_service.ArticleExtractorWeixin(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}

	c.JSON(http.StatusOK, res)
}
