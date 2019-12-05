package request_build_service

import (
	"fmt"
	"net/url"
	"gitlab.fxt.cn/fxt/rank-util/structs/request_builder"
	"gitlab.fxt.cn/fxt/rank-util/utils"
)

func BuildDcRequestBaiduMiniProgram(rbr *request_builder.RequestBuilderRequest) *request_builder.DcRequest {
	var responseTypes []string
	if !(rbr.Capture) {
		responseTypes = []string{"body"}
	} else {
		responseTypes = []string{"body", "capture"}
	}

	var requestUrl string
	if rbr.Page == 1 {
		requestUrl = fmt.Sprintf("https://m.baidu.com/s?wd=%s&ie=utf-8", url.QueryEscape(rbr.SearchWord))
	} else {
		requestUrl = fmt.Sprintf("https://m.baidu.com/s?wd=%s&pn=%d&ie=utf-8", url.QueryEscape(rbr.SearchWord), (rbr.Page-1)*10)
	}

	dcRequest := &request_builder.DcRequest{}
	dcRequest.UniqueKey = utils.DcUniqueKey(requestUrl, rbr.Capture, rbr.SearchCycle)
	dcRequest.Request.Url = requestUrl
	dcRequest.Request.UserAgent = utils.RandomUserAgentForEngine("baidu_mini_program")
	dcRequest.Request.Cookie = fmt.Sprintf("BAIDUID=%s", utils.BaiduCookie())
	dcRequest.Request.Body = ""
	dcRequest.Config.District = ""
	dcRequest.Config.ResponseTypes = responseTypes
	dcRequest.Config.FollowRedirect = false
	dcRequest.Config.Priority = rbr.Priority
	dcRequest.Status = 0

	return dcRequest
}
