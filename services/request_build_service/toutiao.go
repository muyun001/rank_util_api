package request_build_service

import (
	"fmt"
	"net/url"
	"gitlab.fxt.cn/fxt/rank-util/structs/article_request_builder"
	"gitlab.fxt.cn/fxt/rank-util/structs/request_builder"
	"gitlab.fxt.cn/fxt/rank-util/utils"
)

func BuildDcRequestToutiao(rbr *article_request_builder.RequestBuilderRequest) *request_builder.DcRequest {
	var requestUrl string
	if rbr.Url == "" {
		requestUrl = fmt.Sprintf("https://www.toutiao.com/api/search/content/?aid=24&app_name=web_search&offset=0&format=json&keyword=%s&pd=synthesis", url.QueryEscape(rbr.SearchWord))
	} else {
		requestUrl = rbr.Url
	}

	dcRequest := &request_builder.DcRequest{}
	dcRequest.UniqueKey = utils.DcArticleUniqueKey(requestUrl)
	dcRequest.Request.Url = requestUrl
	dcRequest.Request.UserAgent = utils.RandomUserAgentForEngine("toutiao")
	dcRequest.Request.Cookie = ""
	dcRequest.Request.Body = ""
	dcRequest.Config.District = ""
	dcRequest.Config.ResponseTypes = []string{"body"}
	dcRequest.Config.FollowRedirect = false
	dcRequest.Config.Priority = "normal"
	dcRequest.Status = 0

	return dcRequest
}
