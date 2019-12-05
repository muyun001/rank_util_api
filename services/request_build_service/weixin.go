package request_build_service

import (
	"fmt"
	"net/url"
	"gitlab.fxt.cn/fxt/rank-util/structs/article_request_builder"
	"gitlab.fxt.cn/fxt/rank-util/utils"
)

func BuildDcRequestWeixin(rbr *article_request_builder.RequestBuilderRequest) *article_request_builder.DcRequest {
	var requestUrl string
	if rbr.Url == "" {
		requestUrl = fmt.Sprintf("https://weixin.sogou.com/weixin?query=%s&_sug_type_=&s_from=input&_sug_=n&type=2&ie=utf8", url.QueryEscape(rbr.SearchWord))
	} else {
		requestUrl = rbr.Url
	}

	dcRequest := &article_request_builder.DcRequest{}
	dcRequest.UniqueKey = utils.DcArticleUniqueKey(requestUrl)
	dcRequest.Request.Url = requestUrl
	dcRequest.Request.UserAgent = utils.RandomUserAgentForEngine("sogou_pc")
	dcRequest.Request.Cookie = ""
	dcRequest.Request.Body = ""
	dcRequest.Config.District = ""
	dcRequest.Config.ResponseTypes = []string{"body"}
	dcRequest.Config.FollowRedirect = false
	dcRequest.Config.Priority = "normal"
	dcRequest.Status = 0

	return dcRequest
}
