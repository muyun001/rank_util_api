package request_build_service

import (
	"fmt"
	"net/url"
	"gitlab.fxt.cn/fxt/rank-util/structs/request_builder"
	"gitlab.fxt.cn/fxt/rank-util/utils"
)

func BuildDcRequestSug360Pc(rbr *request_builder.RequestBuilderRequest) *request_builder.DcRequest {
	var responseTypes []string
	if ! (rbr.Capture) {
		responseTypes = []string{"body"}
	} else {
		responseTypes = []string{"body", "capture"}
	}

	requestUrl := fmt.Sprintf("https://sug.so.360.cn/suggest?encodein=utf-8&encodeout=utf-8&format=json&fields=word&word=%s", url.QueryEscape(rbr.SearchWord))

	dcRequest := &request_builder.DcRequest{}
	dcRequest.UniqueKey = utils.DcUniqueKey(requestUrl, rbr.Capture, rbr.SearchCycle)
	dcRequest.Request.Url = requestUrl
	dcRequest.Request.UserAgent = ""
	dcRequest.Request.Cookie = ""
	dcRequest.Request.Body = ""
	dcRequest.Config.District = ""
	dcRequest.Config.ResponseTypes = responseTypes
	dcRequest.Config.FollowRedirect = false
	dcRequest.Config.Priority = "normal"
	dcRequest.Status = 0

	return dcRequest
}
